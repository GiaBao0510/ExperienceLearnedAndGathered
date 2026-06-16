# Tập trung và trực quan hóa log Docker với Loki và Grafana

## Giới thiệu

Tài liệu này hướng dẫn cách thiết lập hệ thống tập trung log (centralized logging) cho Docker container sử dụng Grafana Loki và Grafana. Hệ thống này giúp thu thập log từ nhiều container về một nơi, truy vấn và trực quan hóa trên giao diện Grafana, hỗ trợ mở rộng theo chiều ngang khi cần.

---

## Kiến trúc tổng quan

Có hai phương án triển khai:

**Phương án 1 — Docker Loki Plugin (tài liệu này sử dụng):**

```
Docker daemon --> Loki (qua Docker logging driver) --> Grafana
```

Docker daemon gửi log trực tiếp đến Loki thông qua plugin logging driver mà không cần thành phần trung gian. Phù hợp với môi trường Docker thuần túy.

**Phương án 2 — Promtail làm agent trung gian:**

```
Log files / Docker containers --> Promtail --> Loki --> Grafana
```

Promtail đọc log từ file hoặc Docker socket và đẩy vào Loki. Phù hợp khi cần thu thập log từ nhiều nguồn khác nhau (file log ứng dụng, PM2, hệ thống).

> Bản gốc mô tả không nhất quán: phần mở đầu giới thiệu phương án dùng Docker daemon gửi thẳng đến Loki (không cần Promtail), nhưng phần hướng dẫn lại bao gồm cả thiết lập Promtail mà không phân biệt rõ hai cách tiếp cận. Đã tái cấu trúc để người đọc hiểu rõ mình đang dùng phương án nào và lý do.

---

## Cây thư mục tổng quan

```
loki-lab/
├── loki-config.yaml          # Cấu hình Loki
├── docker-compose.yaml       # Khởi động Loki + Grafana
└── logs/                     # (Nếu dùng Promtail) Thư mục chứa file log ứng dụng
    └── app.log

promtail/                     # (Tùy chọn) Nếu dùng Promtail
├── promtail-config.yaml
└── promtail-compose.yaml
```

---

## Các thành phần

### Grafana

Nền tảng quan sát mã nguồn mở, cung cấp giao diện web để truy vấn và trực quan hóa dữ liệu log và metric.

Chức năng chính:
- Trực quan hóa dữ liệu chuỗi thời gian và log
- Truy vấn log bằng ngôn ngữ **LogQL**
- Tích hợp với nhiều nguồn dữ liệu: Loki, Prometheus, Elasticsearch, và các nguồn khác
- Tạo dashboard, thiết lập alert

### Loki

Hệ thống tổng hợp log có khả năng mở rộng theo chiều ngang, lấy cảm hứng từ Prometheus nhưng dành cho log thay vì metric.

Đặc điểm kỹ thuật quan trọng:
- **Không lập chỉ mục toàn văn (no full-text indexing):** Loki chỉ lập chỉ mục các label (metadata có cấu trúc), không lập chỉ mục nội dung log. Đây là lý do Loki tiêu tốn ít tài nguyên hơn Elasticsearch nhưng cũng giới hạn khả năng tìm kiếm toàn văn bản.
- **Nén log:** Log được nén và lưu theo chunk, tiết kiệm dung lượng lưu trữ đáng kể.
- **Tương thích tốt với Docker, PM2, microservices**

> Bản gốc mô tả Loki "lập chỉ mục nhật ký với sử dụng tài nguyên thấp" — thiếu chính xác. Điểm cốt lõi của Loki là **không lập chỉ mục nội dung log**, chỉ lập chỉ mục label. Đây là trade-off quan trọng cần hiểu: tốn ít resource hơn nhưng truy vấn toàn văn chậm hơn Elasticsearch. Đã bổ sung giải thích.

### Promtail

Agent thu thập log chạy cùng với ứng dụng, đọc log từ file hoặc Docker socket rồi đẩy vào Loki.

Chức năng chính:
- Đọc log từ file (`.log` của PM2, ứng dụng) hoặc Docker container logs
- Thêm label (tên container, job, tên file) vào từng dòng log trước khi gửi
- Đẩy log đến Loki qua HTTP
- Lưu vị trí đọc (positions) để không đọc lặp lại khi restart

> Bản gốc dịch Promtail là "đại lý thu thập nhật ký vận chuyển các khúc gỗ cho Loki" — đây là bản dịch máy sai nghĩa ("log" bị dịch thành "khúc gỗ", "agent" thành "đại lý"). Đã viết lại toàn bộ.

---

## Bảng so sánh thành phần

| Thành phần | Vai trò | Ghi chú |
|---|---|---|
| Grafana | Giao diện truy vấn, trực quan hóa và cảnh báo | UI duy nhất người dùng tương tác |
| Loki | Lưu trữ và lập chỉ mục label cho log | Tương đương Elasticsearch nhưng nhẹ hơn |
| Promtail | Agent đọc log từ file/Docker và đẩy vào Loki | Tùy chọn — không cần nếu dùng Docker plugin |
| Docker logging driver | Gửi log container trực tiếp đến Loki | Thay thế Promtail trong môi trường Docker thuần |
| PM2 | Process manager cho Node.js, tạo ra file log stdout/stderr | Nguồn log phổ biến khi không dùng Docker |

---

## Tại sao cần tập trung log?

Trong môi trường có nhiều server hoặc nhiều container, log bị phân tán khắp nơi. Tập trung log giải quyết vấn đề:

- Xem log từ tất cả container và server tại một nơi duy nhất
- Tạo dashboard theo dõi log theo dịch vụ, thời gian, hoặc môi trường
- Tìm kiếm và lọc log nhanh chóng không cần SSH vào từng server
- Thiết lập alert khi log xuất hiện pattern lỗi nhất định

---

## Thiết lập Loki và Grafana

### Bước 1: Tạo thư mục làm việc

```bash
mkdir loki-lab
cd loki-lab
```

### Bước 2: Tạo file cấu hình Loki

Trên Linux/macOS:
```bash
touch loki-config.yaml
```

Trên Windows (PowerShell):
```powershell
New-Item -Name loki-config.yaml -ItemType File
```

Nội dung `loki-config.yaml`:

```yaml
# Tắt xác thực — chỉ dùng trong môi trường lab/dev
# Trong production, cần bật xác thực hoặc đặt Loki sau reverse proxy có auth
auth_enabled: false

server:
  http_listen_port: 3100
  grpc_listen_port: 9096

common:
  instance_addr: 127.0.0.1
  # Thư mục gốc cho tất cả dữ liệu Loki
  path_prefix: /tmp/loki

  storage:
    filesystem:
      chunks_directory: /tmp/loki/chunks
      rules_directory: /tmp/loki/rules

  # Single instance — không cần replication trong lab
  replication_factor: 1

  ring:
    kvstore:
      # inmemory: dữ liệu ring mất khi restart, phù hợp cho lab
      # Production nên dùng: consul, etcd, memberlist
      store: inmemory

query_range:
  results_cache:
    cache:
      embedded_cache:
        enabled: true
        max_size_mb: 100

schema_config:
  configs:
    - from: 2020-10-24
      # TSDB: backend index mặc định từ Loki 2.8+, hiệu năng tốt hơn BoltDB
      store: tsdb
      object_store: filesystem
      schema: v13
      index:
        prefix: index_
        period: 24h

ruler:
  alertmanager_url: http://localhost:9093
```

> Bản gốc có comment `# T0` không có nghĩa. Đây có vẻ là ghi chú chưa hoàn thiện, đã xóa. Đồng thời bổ sung giải thích cho `store: inmemory` và gợi ý thay thế cho production.

**Giải thích các thông số quan trọng:**

- `auth_enabled: false`: Tắt xác thực. Chỉ dùng trong lab hoặc khi Loki đã được bảo vệ bằng lớp khác (VPN, reverse proxy).
- `store: inmemory`: Loki dùng ring (hash ring) để phân phối dữ liệu giữa các instance. Trong lab single-node, `inmemory` là đủ; production cần `memberlist` hoặc `consul`.
- `store: tsdb` với `schema: v13`: TSDB là backend index mặc định từ Loki 2.8+. Nếu dùng Loki cũ hơn, có thể cần đổi sang `boltdb-shipper`.

---

### Bước 3: Tạo file Docker Compose cho Loki và Grafana

Tạo `docker-compose.yaml`:

```yaml
services:
  loki:
    image: grafana/loki:3.0.0
    container_name: loki
    volumes:
      # Mount file cấu hình vào container
      - ./loki-config.yaml:/etc/loki/loki-config.yaml
      # Named volume để dữ liệu log không mất khi container restart
      - loki-data:/tmp/loki
    ports:
      - "3100:3100"
    restart: unless-stopped
    command: -config.file=/etc/loki/loki-config.yaml

  grafana:
    image: grafana/grafana:11.0.0
    container_name: grafana
    ports:
      - "3000:3000"
    restart: unless-stopped
    depends_on:
      - loki
    volumes:
      - grafana-data:/var/lib/grafana

volumes:
  loki-data:
  grafana-data:
```

> Bản gốc dùng `image: grafana/loki:latest` và `grafana/grafana:latest` — tag `latest` không ổn định cho môi trường thực và khó tái hiện lỗi khi debug. Đã đổi sang version cụ thể. Điều chỉnh sang version phù hợp với nhu cầu thực tế.

> Bản gốc mount `./loki-config.yml` (đuôi `.yml`) nhưng file tạo ra là `loki-config.yaml` (đuôi `.yaml`) — mismatch tên file khiến Loki không đọc được config và báo lỗi. Đã thống nhất dùng `.yaml`.

---

### Bước 4: Khởi động Loki và Grafana

```bash
docker compose up -d
```

Kiểm tra container đang chạy:

```bash
docker compose ps
docker logs loki
docker logs grafana
```

Truy cập:
- Loki health check: `http://localhost:3100/ready` (trả về `ready` khi sẵn sàng)
- Grafana UI: `http://localhost:3000`

Thông tin đăng nhập Grafana mặc định:
- Username: `admin`
- Password: `admin`

Grafana sẽ yêu cầu đổi mật khẩu ở lần đăng nhập đầu tiên.

> Bản gốc ghi kiểm tra Loki tại `http://localhost:3100` nhưng endpoint gốc trả về trang HTML trống. Endpoint đúng để kiểm tra trạng thái là `http://localhost:3100/ready`. Đã sửa.

---

## Kết nối Grafana với Loki

1. Đăng nhập Grafana tại `http://localhost:3000`
2. Vào **Connections** → **Data Sources** → **Add new data source**
3. Chọn **Loki**
4. Nhập URL: `http://loki:3100`
   - Dùng tên service `loki` (không phải `localhost`) vì Grafana và Loki chạy trong cùng Docker network của Compose.
5. Nhấn **Save & Test** — kết quả nên là "Data source successfully connected"

> Bản gốc hướng dẫn nhập "URL của Loki" mà không chỉ rõ URL cụ thể là gì. Trong môi trường Docker Compose, hai container giao tiếp với nhau qua tên service, không qua `localhost`. Đã bổ sung.

---

## Phương án 2: Thu log qua Promtail

Dùng Promtail khi cần thu log từ file ứng dụng (PM2, ứng dụng Go ghi ra file) hoặc khi muốn kiểm soát chi tiết hơn về label và pipeline xử lý log.

### Thiết lập Promtail thu log từ file

**Bước 1: Tạo thư mục Promtail và thư mục chứa log**

```bash
mkdir -p promtail/logs
cd promtail

# Tạo file log mẫu để kiểm tra
echo "2024-01-15T10:00:00Z INFO application started" > logs/app.log
echo "2024-01-15T10:00:01Z INFO listening on :8080" >> logs/app.log
```

**Bước 2: Tạo `promtail-config.yaml`**

```yaml
server:
  http_listen_port: 9080
  grpc_listen_port: 0

positions:
  # Lưu vị trí đọc log để không đọc lặp lại sau khi Promtail restart
  filename: /tmp/positions.yaml

clients:
  # Địa chỉ Loki — dùng tên service nếu chạy trong cùng Docker network
  - url: http://loki:3100/loki/api/v1/push

scrape_configs:
  - job_name: app-logs
    static_configs:
      - targets:
          - localhost
        labels:
          job: my-app
          environment: development
          # Đọc tất cả file .log trong thư mục được mount
          __path__: /mnt/logs/*.log
```

> Bản gốc dùng `url: http://localhost:3100/loki/api/v1/push` trong config Promtail — không hoạt động khi Promtail chạy trong Docker container vì `localhost` trong container trỏ đến container đó, không phải host machine hay container Loki. Đúng phải là `http://loki:3100/loki/api/v1/push` (tên service trong Docker network). Đã sửa.

> Bản gốc trình bày file YAML cho Promtail Windows dưới dạng một dòng không có xuống dòng — không phải YAML hợp lệ, gây lỗi parse. Đã sửa thành YAML đúng định dạng.

**Bước 3: Tạo `promtail-compose.yaml`**

```yaml
services:
  promtail:
    image: grafana/promtail:3.0.0
    container_name: promtail
    restart: unless-stopped
    volumes:
      # Mount thư mục chứa log — :ro để container chỉ đọc, không ghi
      - ./logs:/mnt/logs:ro
      # Mount file cấu hình Promtail
      - ./promtail-config.yaml:/etc/promtail/promtail-config.yaml
    command: -config.file=/etc/promtail/promtail-config.yaml
    networks:
      - loki-net

networks:
  loki-net:
    external: true
```

> Promtail cần kết nối được với Loki. Nếu chạy file Compose riêng, cần khai báo cùng Docker network. Đã thêm `networks` để đảm bảo Promtail và Loki giao tiếp được với nhau. Tên network `loki-net` cần khớp với network trong file Compose của Loki+Grafana.

**Bước 4: Khởi động Promtail**

```bash
docker compose -f promtail-compose.yaml up -d
```

**Bước 5: Kiểm tra log trong Grafana**

1. Mở Grafana tại `http://localhost:3000`
2. Vào **Explore**
3. Chọn datasource **Loki**
4. Nhập truy vấn LogQL:

```logql
{job="my-app"}
```

Hoặc lọc theo environment:

```logql
{job="my-app", environment="development"}
```

---

## Thu log Docker container trực tiếp qua Loki Docker Plugin

Nếu muốn Docker daemon tự động gửi log của tất cả container đến Loki mà không cần Promtail:

### Bước 1: Cài đặt Loki Docker Plugin

```bash
docker plugin install grafana/loki-docker-driver:latest --alias loki --grant-all-permissions
```

### Bước 2: Cấu hình Docker daemon dùng Loki làm logging driver mặc định

Sửa file `/etc/docker/daemon.json` (Linux) hoặc Docker Desktop Settings → Docker Engine (Windows/macOS):

```json
{
  "log-driver": "loki",
  "log-opts": {
    "loki-url": "http://localhost:3100/loki/api/v1/push",
    "loki-batch-size": "400",
    "loki-retries": "5",
    "loki-max-backoff": "800ms",
    "loki-timeout": "1s"
  }
}
```

Restart Docker daemon:

```bash
sudo systemctl restart docker
```

Sau khi cấu hình, mọi container khởi động mới sẽ tự động gửi log đến Loki. Truy vấn log trong Grafana theo label container:

```logql
{container_name="my-container"}
```

---

## Các truy vấn LogQL cơ bản

LogQL là ngôn ngữ truy vấn của Loki, cú pháp tương tự PromQL.

```logql
# Lọc theo label
{job="my-app"}

# Lọc theo label và tìm chuỗi trong nội dung log
{job="my-app"} |= "ERROR"

# Tìm log chứa "timeout" hoặc "connection refused"
{job="my-app"} |~ "timeout|connection refused"

# Đếm số dòng log lỗi theo thời gian (rate)
rate({job="my-app"} |= "ERROR" [5m])

# Lọc log trong khoảng thời gian cụ thể
{job="my-app"} | json | level="error"
```

---

## Thông tin bổ sung

### 1. So sánh Loki với ELK Stack (Elasticsearch + Logstash + Kibana)

| Tiêu chí | Loki | ELK Stack |
|---|---|---|
| Tài nguyên | Nhẹ, RAM thấp | Tốn nhiều RAM (ES cần ít nhất 2–4 GB) |
| Tìm kiếm toàn văn | Không hỗ trợ natively | Hỗ trợ đầy đủ |
| Chi phí vận hành | Thấp | Cao hơn đáng kể |
| Phù hợp với | Log có cấu trúc, microservices, Docker | Log cần full-text search, phân tích phức tạp |
| Tích hợp | Grafana (native) | Kibana |

### 2. Lưu ý quan trọng về label trong Loki

Loki lập chỉ mục theo label — số lượng giá trị unique của label ảnh hưởng trực tiếp đến hiệu năng:

- **Nên dùng label có cardinality thấp:** `job`, `environment`, `service`, `container_name`
- **Tránh dùng label có cardinality cao:** user_id, request_id, IP address — những giá trị này nên nằm trong nội dung log, không phải label

```logql
# Sai — tạo hàng triệu stream
{user_id="12345"}

# Đúng — tìm trong nội dung log
{job="api"} |= "user_id=12345"
```

### 3. Loki trong production

Các điểm cần thay đổi khi chuyển từ lab lên production:

- **Storage:** Thay `filesystem` bằng object storage (AWS S3, GCS, MinIO) để có khả năng mở rộng và độ bền cao hơn.
- **Authentication:** Bật `auth_enabled: true` và cấu hình multi-tenancy, hoặc đặt Loki sau reverse proxy (Nginx, Traefik) có xác thực.
- **Ring store:** Thay `inmemory` bằng `memberlist` (cho cluster Loki nhiều node) hoặc `consul`.
- **Retention:** Cấu hình thời gian lưu log để kiểm soát dung lượng storage.
- **High Availability:** Triển khai Loki theo microservice mode với nhiều component (ingester, querier, distributor) riêng biệt.
