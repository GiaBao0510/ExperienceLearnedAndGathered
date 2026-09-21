Tài liệu này sẽ nói cách thiết lập và sao lưu Datatbase và workflow

Quá trình sao lưu database PostgreSQL vào local và Cloud (Cloudflare)


#### **Thường sẽ sao lưu dưới định dạng `.dump` lý do là vì:**
- Sao lưu cơ sở dữ liệu (database) ở dạng **dump** (như dùng `mysqldump` hay `pg_dump`) ==[giúp chuyển đổi dữ liệu thành các câu lệnh SQL hoặc tệp văn bản gọn nhẹ để dễ dàng di chuyển và phục hồi](https://erpsaigon.com/backup-database-dump-load-sage-300/#:~:text=G%E1%BB%ADi%20d%E1%BB%AF,to%C3%A0n%20v%E1%BA%B9n%3B)==
- **Độc lập phiên bản và hệ thống:** Dạng dump lưu cấu trúc và dữ liệu dưới dạng mã nguồn SQL thuần túy, không phụ thuộc vào đường dẫn vật lý hay định dạng bộ nhớ thô của hệ điều hành gốc.
- **Tính nhất quán cao:** Quá trình dump tạo ra một bản sao đồng nhất tại thời điểm chạy lệnh mà không làm gián đoạn nặng nề đến hoạt động đọc ghi của hệ thống

#### **Vì sao sao lưu trên Cloudflare, vì:**
- **Tăng tốc độ website (CDN):** Hệ thống mạng phân phối nội dung (CDN) lưu bản sao nội dung tĩnh (hình ảnh, CSS, JS) tại các máy chủ gần người dùng nhất, giúp giảm độ trễ khi tải trang
- **Giảm tải cho máy chủ gốc:** Nhờ tính năng lưu bộ nhớ đệm (Caching), phần lớn các yêu cầu truy cập được Cloudflare xử lý thay, giúp máy chủ của bạn đỡ quá tải và tiết kiệm băng thông
- **Tiết kiệm chi phí:** Cung cấp nhiều tính năng quan trọng ở gói miễn phí (Free tier), giúp các cá nhân và doanh nghiệp nhỏ vận hành website ổn định mà không tốn nhiều phí duy trì

---
Để thực hành quá trình sao lưu tự động thì đầu tiên tạo ra một service riêng chạy trên container thông qua docker, thì quá trình này mô tả kỹ như sau:

Tạo file .env với các thông tin đầu vào như sau:
```evn
# cloudflare R2 configuration
R2_ACCOUNT_ID=xxxxxxxxxxxxxxxxxxxxxxxxxx
R2_ACCESS_KEY_ID=xxxxxxxxxxxxxxxxxxxxxxxxxx
R2_SECRET_ACCESS_KEY=xxxxxxxxxxxxxxxxxxxxxxxxxx
R2_BUCKET_NAME=ecommerce-backups

# database configuration
DB_HOST=postgresql
DB_USER=USER
DB_PASSWORD=USER
DB_NAME=e-commerce

# position for config local file
ConfigPath=configs
ConfigName=local
ConfigType=yaml
```

Đầu tiên tạo một service trong docker-compose: như sau:
```yml
# ============================================
# Cấu hình logging dùng chung cho các service
# ============================================
x-logging: &default-logging
  driver: "json-file"
  options:
    max-size: "10m"
    max-file: "3"
services:
	#...
	# Service Backup: sao lưu dữ liệu PostgreSQL sang nơi an toàn

  backup-service:
    build:
      context: .    # Thư mục gốc của dự án (nơi chứa docker-compose.yaml)
      dockerfile: docker/backup/Dockerfile
    container_name: backup.ecommerce
    restart: unless-stopped
    environment:

      # Cấu hình cho ứng dụng Go
      ConfigPath: /app/configs
      ConfigName: docker
      ConfigType: yaml

      DB_HOST: ${DB_HOST}                 # Tên service Postgres trong CÙNG docker network — Docker tự phân giải DNS nội bộ
      DB_PORT: "5432"
      DB_USER: ${DB_USER}
      DB_PASSWORD: ${DB_PASSWORD}
      DB_NAME: e-commerce
      BACKUP_CRON: "0 2 * * *"            # Cú pháp cron chuẩn: phút giờ ngày tháng thứ — "0 2 * * *" = 2h00 sáng mỗi ngày
      BACKUP_RETENTION_DAYS: "7"          # Giữ lại backup trong 7 ngày gần nhất
      BACKUP_DIR: "/backups"             # Thư mục lưu trữ backup
      R2_ACCOUNT_ID: ${R2_ACCOUNT_ID}        # Lấy từ biến môi trường .env
      R2_ACCESS_KEY_ID: ${R2_ACCESS_KEY_ID}
      R2_SECRET_ACCESS_KEY: ${R2_SECRET_ACCESS_KEY}
      R2_BUCKET_NAME: ${R2_BUCKET_NAME}
      APP_ENV: docker
      TZ: "Asia/Ho_Chi_Minh"                  # Thiết lập múi giờ cho cronjob
    volumes:
      - backup_data:/backups  # Lưu trữ backup bền vững trên host
    networks:
      - ecommerce_net
    depends_on:
      postgresql:
        condition: service_healthy  # Chỉ chạy backup khi Postgres đã sẵn sàng
    deploy:
      resources:
        limits:
          memory: 256M
          cpus: "0.5"
    logging: *default-logging
```

Tạo 1 file Dockerfile lưu ở một nơi nào đó (ví dụ lưu tại: `docker/backup/Dockerfile`) để cho backupservice dựa vào các thông tin ở đây để mà có thể tự động thực hiện backup:
```
# Dockerfile riêng cho backup-service — TÁCH BIỆT hoàn toàn khỏi
# Dockerfile của API server chính (image nhẹ, chỉ chứa đúng thứ cần).

# ------ Giai đoạn 1: Build ứng dụng Go ------
FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /backup-service ./cmd/backup

# ------ Giai đoạn 2: Tạo image nhẹ cho backup-service ------
FROM alpine:3.20

# Cài pg_dump/pg_restore — client tools của PostgreSQL.
# [LƯU Ý] Nên chọn version client KHỚP với version server 
RUN apk add --no-cache postgresql15-client tzdata

WORKDIR /app
COPY --from=builder /backup-service .
COPY --from=builder /app/configs ./configs

# Thư mục lưu backup — sẽ được mount ra Docker volume trong docker-compose
RUN mkdir -p /backups

ENTRYPOINT ["/app/backup-service"]
```

Tạo một file docker.yml (lưu tại thư mục configs) riêng để lưu các thông tin cấu hình cho backupservice đọc và thiết lập như sau:
```yml
server:
  port: 8080
  host: localhost
  mode: "dev"

postgres:
  host: postgresql 
  port: 5432
  user: USER
  password: USER
  dbname: e-commerce
  
  # Số kết nối nhàn rỗi được giữ lại trong pool khi không có request
  # Giá trị cao giúp tái sử dụng connection nhanh hơn, nhưng tốn tài nguyên hơn
  maxIdleConns: 10

  # Số kết nối mở tối đa đến database cùng một lúc
  # Vượt ngưỡng này, các request mới sẽ phải chờ có connection rảnh
  maxOpenConns: 100

  connMaxLifetime: 3600 # Đặt vòng đời tối đa của một kết nối là 1 phút


authentication:
 #...
  cloudflare:
    r2_token_name: "ecommerce-backups"
    r2_bucket_name: "ecommerce-backups"
    r2_token_value: "xxxxxxxxxxxxxxxxxxx"
    r2_account_id: "xxxxxxxxxxxxxxxxxxxx"
    r2_access_key_id: "xxxxxxxxxxxxxxxxxxxx"
    r2_secret_access_key: "xxxxxxxxxxxxxxxxxxx"
    r2_endpoint: "https://xxxxxxxxxxxxxxxxxxx.r2.cloudflarestorage.com"

#.....

# Cấu hình cho các cronjob
cronjob:
  backup_cron: "@every 1h"
  backup_retention_days: 7
  backup_dir: "/backups"
```

Tạo `cmd\cronjob\backupsDB.go` để chạy backup ngầm đã thiết lập sẳn trước đó

---
## **Backup - workflow**

Quy trình hoạt động như sau:


---
Vì sao lưu có thể ở nhiều dạng như ở cục bộ (local), cloud (AWS S3, cloudflare,...) thì có thể tổ chức code theo dạng Strategy Design Pattern thuộc nhóm Behavioral trong Design pattern


Kết quả sau khi sao lưu trên Cloudflare. Thì nó nằm ở mục **Storage & databases > R2 Object Storage > Bukets**
![](https://i.ibb.co/jPZbf4Vt/image.png)

![](![[Pasted image 20260920160736.png]])