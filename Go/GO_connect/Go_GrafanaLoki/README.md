Tại đây sẽ kết hợp Grafana, Loki để truy vấn nội dung trong những file đã log (.log)

## Trực quan hóa về Log

![](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/ezpdcwsweednvtyqtbi0.png)

Thì chúng ta quan sát hình ảnh ở trên.

- Logs: nơi lưu trữ các thông tin trong projetc
- Promtail: Đọc/quét dữ liệu từ file Logs. Sau đó gửi lên Loki
- Loki: Nơi này để truy vấn log.
- Grafana: Thì 3 thành phần trên nằm rong Grafana

Theo mô tả luồng như sau:
App (ghi log: app.log, http.log, sql.log)
⬇️ (mount volume)
Promtail đọc file log
⬇️
Promtail gửi log -> Loki (http://loki:3100)
⬇️
Grafana Query Loki để hiển thị log

Lưu ý quan trọng:
- Trong file `promtail.config.yml`, Promtail phải đẩy log tới `http://loki:3100/loki/api/v1/push`.
- Khi thêm Data source trong Grafana, URL của Loki cũng պետք là `http://loki:3100`, không phải `http://localhost:3100`.
- Truy vấn `| json | level="ERROR"` chỉ hoạt động nếu mỗi dòng log là JSON hợp lệ và có field `level`.

ban đầu khởi tạo cấu trúc thư mục như sau:
```shell
Go_GrafanaLoki> tree /f             
Folder PATH listing
Volume serial number is F61A-D446
D:.
│   docker-compose.yaml
│   go.mod
│   main.go
│   README.md
│   
├───config
├───internal
│   └───logs
│           http.log
│           
└───system
    └───promtail
            promtail.config.yml
```

Nội dung từ docker-compose.yaml:
```yaml
version: '3.8'

services:
  # Grafana: nền tảng giám sát và trực quan hóa dữ liệu
  grafana:
    image: grafana/grafana:11.0.0
    container_name: go-grafana
    restart: unless-stopped

    ports:
      - "3000:3000"

    environment:
      - GF_SECURITY_ADMIN_USER=admin      # thiết lập user admin
      - GF_SECURITY_ADMIN_PASSWORD=admin  # thiết lập password admin

    # depend_on: đảm bảo rằng Loki đã sẵn sàng trước khi Grafana khởi động
    depends_on:
      - loki

    volumes:
      - grafana-data:/var/lib/grafana  # Lưu trữ dữ liệu Grafana trên ổ đĩa local của máy host

  # Loki: lưu trữ log từ các container và ứng dụng
  loki:
    image: grafana/loki:2.9.0
    container_name: go-loki
    restart: unless-stopped
    ports:
      - "3100:3100"

    command: -config.file=/etc/loki/local-config.yaml
    volumes:
      - loki-data:/loki  # Lưu trữ dữ liệu Loki trên ổ đĩa local của máy host

  # Promtail: thu thập log từ các container và gửi đến Loki
  promtail:
    image: grafana/promtail:2.9.0
    container_name: go-promtail
    restart: unless-stopped
    volumes:
      - ./internal/logs:/var/log  # Thu thập log từ host

      - ./system/promtail/promtail.config.yml:/etc/promtail/config.yml  # Cấu hình Promtail
    command: -config.file=/etc/promtail/config.yml
    depends_on:
      - loki

# Dữ liệu sẽ không bị mất khi container bị xóa hoặc rebuild
volumes:
  grafana-data: # Lưu trữ trên ổ đĩa local của máy host
  loki-data:     # Lưu trữ trên ổ đĩa local của máy host
```

Nội dung từ promtail.config.yml:
```yaml
server:
  http_listen_port: 9080
  grpc_listen_port: 0

positions:
  filename: /tmp/positions.yaml

# sau khi prommtail đọc xong thì nó sẽ gửi đến loki
clients:
  - url: http://localhost:3100/loki/api/v1/push
  
# Đọc file
scrape_configs:
  - job_name: golang-logs
    static_configs:
      - targets:
          - localhost
        labels:
          job: golang-logs
          __path__: /var/log/*.log  # Đường dẫn đến file log cần đọc


```

Sau đó chạy lệnh docker compose up để build lên

Chúng ta vào http://localhost:3000 để đăng nhập Grafana

Trong Grafana, thử truy vấn theo từng bước để dễ kiểm tra:
```logql
{job="golang-logs"}
```
Nếu đã thấy log, thử tiếp:
```logql
{job="golang-logs"} | json
```
Cuối cùng mới lọc level:
```logql
{job="golang-logs"} | json | level="ERROR"
```



