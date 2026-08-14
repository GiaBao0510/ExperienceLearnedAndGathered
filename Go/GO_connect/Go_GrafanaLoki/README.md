# Trực quan hóa Log với Grafana, Loki và Promtail

## 1. Mục tiêu

Tài liệu này hướng dẫn cách dựng một hệ thống xem và truy vấn log đơn giản bằng:

- **Grafana**: giao diện trực quan hóa và truy vấn.
- **Loki**: nơi lưu trữ và truy vấn log.
- **Promtail**: đọc file log từ ứng dụng và đẩy lên Loki.

Mô hình này phù hợp để:

- Xem log của ứng dụng theo thời gian gần thực.
- Lọc log theo `level`, `service`, `trace_id`, `status_code`, `duration_ms`.
- Tạo dashboard theo dõi lỗi, request chậm, số lượng log theo thời gian.
- Học cách xây dựng hệ thống logging tập trung ở mức cơ bản.

## 2. Kiến trúc tổng quan

```text
App ghi log ra file
        │
        ▼
internal/logs/*.log
        │
        ▼
Promtail đọc file log
        │
        ▼
Gửi log đến Loki
        │
        ▼
Grafana truy vấn Loki
        │
        ▼
Hiển thị log, biểu đồ, dashboard
```

### 2.1 Vai trò từng thành phần

| Thành phần | Vai trò |
|---|---|
| Application | Ghi log ra file, ví dụ `app.log`, `http.log`, `sql.log` |
| Promtail | Đọc file log, gắn label, gửi log tới Loki |
| Loki | Lưu trữ log và cung cấp API truy vấn LogQL |
| Grafana | Giao diện để truy vấn, xem log và vẽ biểu đồ |

### 2.2 Lưu ý quan trọng

Trong hệ thống này:

- **Grafana không chứa log**.
- **Loki mới là nơi lưu trữ log**.
- **Promtail không phải nơi truy vấn log**, nó chỉ thu thập và đẩy log.
- Grafana sẽ truy vấn Loki thông qua datasource.

## 3. Yêu cầu

Bạn cần có:

- Docker Desktop hoặc Docker Engine.
- Docker Compose.
- Một ứng dụng có khả năng ghi log ra file.
- Log nên ở dạng JSON nếu muốn dùng các truy vấn `| json`.

Ví dụ đường dẫn log trong tài liệu này:

```text
internal/logs/http.log
internal/logs/app.log
internal/logs/sql.log
```

## 4. Cấu trúc thư mục đề xuất

```text
.
├── docker-compose.yaml
├── go.mod
├── main.go
├── README.md
├── internal/
│   └── logs/
│       ├── app.log
│       ├── http.log
│       └── sql.log
└── system/
    └── promtail/
        └── promtail.config.yml
```

Nếu chưa có file log, bạn có thể tạo trước:

```bash
mkdir -p internal/logs
touch internal/logs/http.log
touch internal/logs/app.log
touch internal/logs/sql.log
```

Trên Windows PowerShell:

```powershell
New-Item -ItemType Directory -Force -Path internal/logs
New-Item -ItemType File -Force -Path internal/logs/http.log
New-Item -ItemType File -Force -Path internal/logs/app.log
New-Item -ItemType File -Force -Path internal/logs/sql.log
```

## 5. File `docker-compose.yaml`

```yaml
services:
  grafana:
    image: grafana/grafana:11.0.0
    container_name: go-grafana
    restart: unless-stopped
    ports:
      - "3000:3000"
    environment:
      - GF_SECURITY_ADMIN_USER=admin
      - GF_SECURITY_ADMIN_PASSWORD=admin
    volumes:
      - grafana-data:/var/lib/grafana
    depends_on:
      - loki

  loki:
    image: grafana/loki:2.9.0
    container_name: go-loki
    restart: unless-stopped
    ports:
      - "3100:3100"
    command: -config.file=/etc/loki/local-config.yaml
    volumes:
      - loki-data:/loki

  promtail:
    image: grafana/promtail:2.9.0
    container_name: go-promtail
    restart: unless-stopped
    volumes:
      - ./internal/logs:/var/log
      - ./system/promtail/promtail.config.yml:/etc/promtail/config.yml
    command: -config.file=/etc/promtail/config.yml
    depends_on:
      - loki

volumes:
  grafana-data:
  loki-data:
```

### 5.1 Giải thích ngắn

| Service | Ý nghĩa |
|---|---|
| `grafana` | Chạy web UI tại `http://localhost:3000` |
| `loki` | Chạy Loki tại `http://loki:3100` trong Docker network |
| `promtail` | Mount thư mục `./internal/logs` vào `/var/log` để đọc log |
| `grafana-data` | Lưu cấu hình/dashboard của Grafana |
| `loki-data` | Lưu dữ liệu Loki |

### 5.2 Vì sao không dùng `http://localhost:3100` cho datasource trong Grafana?

Trong Docker Compose, các container nói chuyện với nhau bằng tên service.

Vì vậy:

- Từ **Promtail container**, muốn gửi log tới Loki, phải dùng:

```text
http://loki:3100/loki/api/v1/push
```

- Từ **Grafana container**, khi cấu hình datasource Loki, cũng nên dùng:

```text
http://loki:3100
```

Nếu bạn dùng `http://localhost:3100` bên trong Grafana container, `localhost` đó là chính container Grafana, không phải container Loki.

Nếu bạn truy cập từ máy host, ví dụ dùng trình duyệt hoặc curl trên máy bạn, thì có thể dùng:

```text
http://localhost:3100
```

## 6. File `promtail.config.yml`

Đường dẫn:

```text
system/promtail/promtail.config.yml
```

Nội dung:

```yaml
server:
  http_listen_port: 9080
  grpc_listen_port: 0

positions:
  filename: /tmp/positions.yaml

clients:
  - url: http://loki:3100/loki/api/v1/push

scrape_configs:
  - job_name: golang-logs
    static_configs:
      - targets:
          - localhost
        labels:
          job: golang-logs
          __path__: /var/log/*.log
```

### 6.1 Giải thích cấu hình

| Cấu hình | Ý nghĩa |
|---|---|
| `http_listen_port: 9080` | Port API của Promtail |
| `positions.filename` | Nơi Promtail lưu vị trí đã đọc của từng file log |
| `clients.url` | Địa chỉ Loki nhận log |
| `job_name: golang-logs` | Tên job, dùng để truy vấn trong Grafana |
| `job: golang-logs` | Label được gắn vào log |
| `__path__: /var/log/*.log` | Promtail đọc tất cả file có đuôi `.log` trong `/var/log` |

Vì trong `docker-compose.yaml`, thư mục `./internal/logs` được mount vào `/var/log`, nên:

```text
internal/logs/http.log  -> /var/log/http.log
internal/logs/app.log   -> /var/log/app.log
internal/logs/sql.log   -> /var/log/sql.log
```

## 7. Dữ liệu log mẫu

Để truy vấn `| json` hoạt động tốt, mỗi dòng trong file log nên là một JSON object hợp lệ.

Ví dụ file:

```text
internal/logs/http.log
```

Nội dung mẫu:

```json
{"timestamp":"2026-06-22T10:00:00.000Z","level":"INFO","trace_id":"trace-001","span_id":"span-001","request_id":"req-001","service":"user-service","host":"api-01","env":"dev","protocol":"HTTP/1.1","method":"GET","scheme":"http","path":"/api/v1/users","query":null,"status_code":200,"duration_ms":45.2,"client_ip":"127.0.0.1","user_id":"user_001","user_agent":"curl/8.8.0","referer":null,"bytes_in":120,"bytes_out":980,"region":"ap-southeast-1","route":"/users","error":null}
{"timestamp":"2026-06-22T10:00:05.000Z","level":"WARN","trace_id":"trace-002","span_id":"span-002","request_id":"req-002","service":"auth-service","host":"api-01","env":"dev","protocol":"HTTP/1.1","method":"POST","scheme":"http","path":"/api/v1/auth/login","query":null,"status_code":401,"duration_ms":64.5,"client_ip":"127.0.0.1","user_id":null,"user_agent":"curl/8.8.0","referer":null,"bytes_in":220,"bytes_out":160,"region":"ap-southeast-1","route":"/auth/login","error":null}
{"timestamp":"2026-06-22T10:00:10.000Z","level":"ERROR","trace_id":"trace-003","span_id":"span-003","request_id":"req-003","service":"order-service","host":"api-01","env":"dev","protocol":"HTTP/1.1","method":"POST","scheme":"http","path":"/api/v1/orders","query":null,"status_code":500,"duration_ms":320.7,"client_ip":"127.0.0.1","user_id":"user_002","user_agent":"curl/8.8.0","referer":null,"bytes_in":540,"bytes_out":210,"region":"ap-southeast-1","route":"/orders","error":"database_error"}
```

### 7.1 Các field hữu ích

| Field | Ý nghĩa |
|---|---|
| `timestamp` | Thời điểm ứng dụng ghi log |
| `level` | Mức log: INFO, WARN, ERROR |
| `trace_id` | ID để truy vết request |
| `service` | Tên service ghi log |
| `method` | HTTP method |
| `path` | Đường dẫn API |
| `status_code` | HTTP status code |
| `duration_ms` | Thời gian xử lý request |
| `user_id` | ID người dùng |
| `error` | Thông tin lỗi |

### 7.2 Lưu ý về thời gian log

Loki thường dùng thời điểm nó nhận log làm mốc thời gian truy vấn trong các trường hợp mặc định. Vì vậy:

- Nếu log cũ từ nhiều ngày trước, bạn cần chọn time range phù hợp.
- Nếu muốn test nhanh, hãy ghi log mới vào file.
- Khi truy vấn trong Grafana, nên chọn khoảng thời gian rộng như `Last 1 hour` hoặc `Last 6 hours`.

## 8. Chạy hệ thống

Tại thư mục gốc của project:

```bash
docker compose up -d
```

Kiểm tra container:

```bash
docker compose ps
```

Xem log của Promtail:

```bash
docker compose logs -f promtail
```

Xem log của Loki:

```bash
docker compose logs -f loki
```

Xem log của Grafana:

```bash
docker compose logs -f grafana
```

Kiểm tra file log có tồn tại trong container Promtail hay không:

```bash
docker compose exec promtail ls -la /var/log
```

## 9. Đăng nhập Grafana

Truy cập:

```text
http://localhost:3000
```

Đăng nhập bằng:

```text
Username: admin
Password: admin
```

Đây là cấu hình mặc định trong `docker-compose.yaml`, chỉ nên dùng cho môi trường dev.

## 10. Thêm datasource Loki trong Grafana

Các bước:

1. Vào **Connections** hoặc **Configuration**.
2. Chọn **Data sources**.
3. Chọn **Add data source**.
4. Chọn **Loki**.
5. Nhập URL:

```text
http://loki:3100
```

6. Chọn **Save & test**.

Nếu thành công, Grafana sẽ báo datasource hoạt động.

### 10.1 Lưu ý quan trọng

Trong ô URL của datasource, bạn nên nhập:

```text
http://loki:3100
```

Không nên nhập:

```text
http://localhost:3100
```

nếu Grafana cũng đang chạy trong Docker Compose.

## 11. Truy vấn log trong Grafana

Vào **Explore**, chọn datasource Loki.

### 11.1 Xem toàn bộ log của job `golang-logs`

```logql
{job="golang-logs"}
```

Ý nghĩa:

```text
Lấy tất cả log có label job = golang-logs
```

Tương tự câu lệnh SQL:

```sql
SELECT *
FROM logs
WHERE job = 'golang-logs';
```

### 11.2 Parse log dạng JSON

```logql
{job="golang-logs"} | json
```

Câu lệnh này sẽ parse từng dòng log JSON để trích xuất các field như:

- `level`
- `service`
- `trace_id`
- `status_code`
- `duration_ms`
- `error`

### 11.3 Lọc log theo level ERROR

```logql
{job="golang-logs"} | json | level="ERROR"
```

Ý nghĩa:

```text
Chỉ lấy các dòng log có level là ERROR
```

### 11.4 Lọc log theo trace_id

```logql
{job="golang-logs"} | json | trace_id="trace-003"
```

Câu lệnh này hữu ích khi bạn muốn truy vết một request cụ thể.

### 11.5 Lọc log theo service

```logql
{job="golang-logs"} | json | service="order-service"
```

### 11.6 Lọc log theo status_code

```logql
{job="golang-logs"} | json | status_code=500
```

Hoặc lọc các lỗi 5xx:

```logql
{job="golang-logs"} | json | status_code >= 500
```

### 11.7 Tìm request chậm

```logql
{job="golang-logs"} | json | duration_ms > 300
```

Câu lệnh này lấy các request có thời gian xử lý lớn hơn 300ms.

### 11.8 Tìm log có chứa chuỗi cụ thể

```logql
{job="golang-logs"} |= "database_error"
```

Câu lệnh này tìm tất cả dòng log có chứa chuỗi `database_error`.

### 11.9 Lọc theo file log cụ thể

Nếu Promtail có label `filename`, bạn có thể truy vấn:

```logql
{filename="/var/log/http.log"}
```

Nếu truy vấn không hoạt động, hãy kiểm tra label thực tế trong Grafana hoặc dùng truy vấn tổng quát trước:

```logql
{job="golang-logs"}
```

### 11.10 Định dạng lại dòng hiển thị

```logql
{job="golang-logs"}
| json
| line_format "{{.timestamp}} | {{.level}} | {{.service}} | {{.method}} {{.path}} | status={{.status_code}} | duration={{.duration_ms}}ms"
```

Câu lệnh này giúp hiển thị log gọn hơn trong Grafana.

### 11.11 Đếm số lượng log theo level

```logql
sum(rate({job="golang-logs"} | json [1m])) by (level)
```

Câu lệnh này phù hợp để vẽ biểu đồ số lượng log theo từng level trong mỗi phút.

### 11.12 Đếm số log ERROR theo thời gian

```logql
sum(rate({job="golang-logs"} | json | level="ERROR" [1m]))
```

Câu lệnh này phù hợp để cảnh báo hoặc theo dõi số lượng lỗi tăng bất thường.

## 12. Tạo dashboard đơn giản

Bạn có thể tạo dashboard với các panel sau:

### 12.1 Panel xem log gần nhất

- Kiểu hiển thị: **Logs**.
- Query:

```logql
{job="golang-logs"} | json
```

### 12.2 Panel số lượng log theo level

- Kiểu hiển thị: **Time series**.
- Query:

```logql
sum(rate({job="golang-logs"} | json [1m])) by (level)
```

### 12.3 Panel số lượng log ERROR

- Kiểu hiển thị: **Time series**.
- Query:

```logql
sum(rate({job="golang-logs"} | json | level="ERROR" [1m]))
```

### 12.4 Panel request chậm

- Kiểu hiển thị: **Table** hoặc **Logs**.
- Query:

```logql
{job="golang-logs"} | json | duration_ms > 300
```

### 12.5 Panel lỗi theo service

- Kiểu hiển thị: **Time series**.
- Query:

```logql
sum(rate({job="golang-logs"} | json | level="ERROR" [1m])) by (service)
```

## 13. Gợi ý tạo log mẫu bằng Go

Nếu bạn đang dùng Go, có thể tạo chương trình đơn giản để ghi log JSON vào file.

Ví dụ file:

```text
main.go
```

```go
package main

import (
	"encoding/json"
	"math/rand"
	"os"
	"time"
)

type HTTPLog struct {
	Timestamp  string  `json:"timestamp"`
	Level      string  `json:"level"`
	TraceID    string  `json:"trace_id"`
	Service    string  `json:"service"`
	Method     string  `json:"method"`
	Path       string  `json:"path"`
	StatusCode int     `json:"status_code"`
	DurationMs float64 `json:"duration_ms"`
	Error      *string `json:"error"`
}

func randomLevel() string {
	levels := []string{"INFO", "WARN", "ERROR"}
	return levels[rand.Intn(len(levels))]
}

func randomPath() string {
	paths := []string{
		"/api/v1/users",
		"/api/v1/products",
		"/api/v1/orders",
		"/api/v1/auth/login",
	}
	return paths[rand.Intn(len(paths))]
}

func main() {
	file, err := os.OpenFile(
		"internal/logs/http.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644,
	)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	for i := 0; i < 20; i++ {
		level := randomLevel()

		var errorMessage *string
		if level == "ERROR" {
			msg := "database_error"
			errorMessage = &msg
		}

		entry := HTTPLog{
			Timestamp:  time.Now().UTC().Format(time.RFC3339Nano),
			Level:      level,
			TraceID:    time.Now().Format("20060102150405.000000"),
			Service:    "demo-service",
			Method:     "GET",
			Path:       randomPath(),
			StatusCode: 200,
			DurationMs: rand.Float64() * 500,
			Error:      errorMessage,
		}

		if level == "ERROR" {
			entry.StatusCode = 500
		}

		if err := encoder.Encode(entry); err != nil {
			panic(err)
		}

		time.Sleep(2 * time.Second)
	}
}
```

Chạy chương trình:

```bash
go run main.go
```

Sau khi chạy, file `internal/logs/http.log` sẽ có thêm các dòng log JSON mới.

## 14. Troubleshooting

### 14.1 Grafana không truy vấn được Loki

Kiểm tra datasource trong Grafana:

```text
http://loki:3100
```

Không dùng:

```text
http://localhost:3100
```

nếu Grafana đang chạy trong Docker Compose.

### 14.2 Không thấy log xuất hiện

Kiểm tra lần lượt:

1. File log có tồn tại không?

```bash
ls -la internal/logs
```

2. File log có được mount vào Promtail không?

```bash
docker compose exec promtail ls -la /var/log
```

3. Promtail có lỗi không?

```bash
docker compose logs -f promtail
```

4. Trong Grafana, chọn time range đủ rộng.

Ví dụ:

```text
Last 1 hour
Last 6 hours
```

5. Truy vấn thử:

```logql
{job="golang-logs"}
```

### 14.3 Truy vấn `| json` không hoạt động

Điều kiện:

- Mỗi dòng log phải là một JSON object hợp lệ.
- Không có nhiều JSON object trên cùng một dòng.
- Không có ký tự rác trước hoặc sau JSON.
- Các field dùng để lọc phải tồn tại sau khi parse JSON.

Ví dụ hợp lệ:

```json
{"level":"ERROR","service":"order-service","trace_id":"trace-003"}
```

Không hợp lệ:

```text
INFO 2026-06-22 something happened
```

Vì dòng này không phải JSON.

### 14.4 Muốn lọc theo field nhưng báo lỗi

Ví dụ:

```logql
{job="golang-logs"} | json | duration_ms > 300
```

Nếu lỗi, hãy kiểm tra:

- Field `duration_ms` có tồn tại trong log không?
- Giá trị có phải số không?
- JSON có hợp lệ không?

Bạn có thể kiểm tra từng bước:

```logql
{job="golang-logs"} | json
```

Sau đó mới thêm điều kiện:

```logql
{job="golang-logs"} | json | duration_ms > 300
```

### 14.5 Promtail không đọc được file log

Có thể do:

- File log chưa tồn tại.
- Đường dẫn mount sai.
- Quyền đọc file không đủ.

Trên Linux/macOS, có thể thử:

```bash
chmod -R 755 internal/logs
```

Sau đó restart:

```bash
docker compose restart promtail
```

### 14.6 Loki không nhận log từ Promtail

Kiểm tra URL trong `promtail.config.yml`:

```yaml
clients:
  - url: http://loki:3100/loki/api/v1/push
```

Không nên dùng:

```yaml
clients:
  - url: http://localhost:3100/loki/api/v1/push
```

vì từ container Promtail, `localhost` không trỏ đến container Loki.

## 15. Ghi chú khi dùng trong production

Cấu hình trong tài liệu này phù hợp cho dev/local. Nếu triển khai production, bạn cần thêm:

- Không đặt password Grafana là `admin/admin`.
- Không expose port Loki ra internet nếu không cần thiết.
- Bật HTTPS cho Grafana.
- Dùng reverse proxy như Nginx, Traefik hoặc Caddy.
- Giới hạn quyền đọc file log cho Promtail.
- Cấu hình retention cho Loki để tránh đầy đĩa.
- Giới hạn dung lượng log ứng dụng.
- Tách riêng log nhạy cảm.
- Không log secret, password, token, số thẻ, dữ liệu cá nhân nhạy cảm.
- Có alert khi log ERROR tăng bất thường.
- Có backup hoặc lưu trữ dự phòng nếu log quan trọng.

## 16. Kết luận

Hệ thống Grafana + Loki + Promtail là một cách tốt để bắt đầu làm quen với logging tập trung và observability.

Luồng cần nhớ:

```text
App ghi log JSON vào file
Promtail đọc file và gửi log đến Loki
Grafana truy vấn Loki để hiển thị log
```

Khi truy vấn, bạn cần nhớ:

- `{job="golang-logs"}` để chọn stream log.
- `| json` để parse JSON.
- Các bộ lọc như `level="ERROR"`, `duration_ms > 300`, `service="order-service"` được dùng sau khi parse JSON.
- Trong Docker Compose, các service nên gọi nhau bằng tên service, ví dụ `http://loki:3100`.

### Mở rộng

Sau khi hoàn thành hệ thống cơ bản, bạn có thể tìm hiểu thêm các chủ đề sau:

- **LogQL**: ngôn ngữ truy vấn của Loki.
- **Label trong Loki**: cách thiết kế label hiệu quả, tránh high cardinality.
- **High cardinality labels**: vì sao không nên dùng `user_id`, `request_id`, `trace_id` làm label Loki.
- **Structured logging**: chuẩn hóa log JSON trong ứng dụng.
- **Correlation ID / Trace ID**: truy vết request qua nhiều service.
- **Distributed Tracing**: kết hợp Loki với Tempo hoặc Jaeger.
- **Grafana Tempo**: lưu trữ và truy vết distributed trace.
- **Grafana Alloy**: công cụ thu thập telemetry mới hơn trong hệ sinh thái Grafana.
- **Promtail pipeline stages**: parse, rewrite, drop, filter log trước khi gửi tới Loki.
- **Loki retention**: cấu hình tự động xóa log cũ.
- **Loki alerting**: tạo cảnh báo dựa trên log, ví dụ số lượng ERROR tăng cao.
- **Grafana dashboard provisioning**: tự động nạp dashboard bằng file YAML/JSON.
- **Service discovery**: để Promtail tự động phát hiện file log trong nhiều môi trường.
- **Kubernetes logging**: thu thập log từ pod/container bằng Promtail hoặc Grafana Alloy.
- **Multi-tenancy trong Loki**: phân tách log theo tenant.
- **Log security**: masking dữ liệu nhạy cảm trước khi ghi log.
- **OpenTelemetry**: chuẩn thu thập log, metric, trace.
- **Metric từ log**: dùng LogQL để tạo metric từ log, ví dụ đếm lỗi theo phút.
- **Loki vs Elasticsearch**: so sánh mô hình indexing và chi phí vận hành.
- **Loki vs ELK/EFK**: so sánh với Elasticsearch, Logstash/Fluentd, Kibana.