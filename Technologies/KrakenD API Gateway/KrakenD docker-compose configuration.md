## Cần chuẩn bị:
- Kiến thức về Docker
- tải docker
- thiết lập Dockercompose trong thiết bị

tạo file docker-compose
```yaml
version: '3'
service:
  krakend:
    image: krakend:latest
    container_name: krakend
    command: run -dc /etc/krakend/krakend.tmpl  #
    volumes:
      - ./krakend_config:/etc/krakend/
    dns: 127.0.0.1
    ports:
      - "8080:8080"
    enviroment:
      - FC_ENABLE=1 # Enable flexible configuration
      - FC_SETTINGS=/etc/krakend/config/settings/ # Correct path to settings
      - FC_PARTIALS=/etc/krakend/config/partials/ # Path for partials
      - FC_TEMPLATES=/etc/krakend/config/templates/ # Path for templates
      - FC_OUT=/etc/krakend/output.json # Optional: Output the final configuration for debugging
    networks:
      - myservice
    restart: unless-stopped
```

Giải thích
• image: Sử dụng image chính thức từ Docker Hub - KrakenD.
• ports: Mở cổng 8080 của KrakenD ra máy local (địa chỉ mặc định truy cập API: http://localhost:8080).
• volumes: Mount file cấu hình krakend.json từ thư mục hiện tại của bạn vào đường dẫn /etc/krakend/krakend.json bên trong container

Giải thích các biến môi trường
- FC_ENABLE:
- FC_SETTINGS
- FC_PARTIALS:
- FC_TEMPLATES
- FC_OUT

---
trong thư mục krakend_config tạo thêm 1 số file

krakend.json
```json
{
  "$schema": "https://www.krakend.io/schema/v3.json",
  "version": 3,
  "timeout": "20000ms",
  "cache_ttl": "300s",
  "extra_config": {
    "telemetry/logging": {
      "level": "INFO",
      "prefix": "[KRAKEND]",
      "syslog": false,
      "stdout": true,
      "format": "default"
    },
    "security/cors": {
      "allow_origins": [
        "*"
      ],
      "allow_methods": [
        "GET",
        "POST",
        "PUT",
        "DELETE"
      ],
      "expose_headers": [
        "Content-Length",
        "Content-Type"
      ],
      "allow_headers": [
        "Accept-Language"
      ],
      "max_age": "12h",
      "allow_credentials": false,
      "debug": false
    },
    "auth/api-keys": {
      "strategy": "header",
      "identifier": "X-Key",
      "hash": "fnv128",
      "salt": "Senior",
      "@hashed-with": "https://go.dev/play/p/hHksLLCUJy3",
      "keys": [
          {
              "@key-plain": "YourKeyPlain",
              "key": "YourKey",
              "roles": ["all"],
              "@description": "ACME Inc."
          }
      ]
    }
  },
  "endpoints": []
}
```

---
Chạy câu lệnh docker-compose up -d

---
Kiểm tra: docker compose logs -f krakend

---

Test Your API Gateway
Once KrakenD is up and running, you can test your API gateway by sending a GET request to http://localhost:8080/hello


[Đoạn yêu cầu (lưu ý tạo bản mởi thì xoá đoạn này đi, chủ yếu bạn cũng nên thực hiện theo hãy yêu cầu này)]: Đây là tài liệu mà tôi thiết kế triển khai Krakend với docker-compose thì bạn xem qua cách tôi thiết kế có cấu trúc, đúng đăng, đúng logic hay chưa. Xem kỹ lại
