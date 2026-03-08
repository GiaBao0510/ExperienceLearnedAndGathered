# 🐳 Dockerfile – Hướng Dẫn Toàn Diện

## 📖 1. Dockerfile là gì?

**Dockerfile** là một tệp văn bản thuần túy chứa tập hợp các **câu lệnh (chỉ thị)** mà Docker sẽ đọc và thực thi tuần tự để tự động tạo ra một **Docker Image** theo nhu cầu.

```
Dockerfile  ──(docker build)──▶  Docker Image  ──(docker run)──▶  Container
```

**Hình dung đơn giản:**

- Dockerfile giống như một **công thức nấu ăn**
- Docker Image giống như **món ăn đã hoàn thành**
- Container giống như **phần ăn được dọn ra bàn để dùng**

---

## ⚙️ 2. Các lệnh Docker cơ bản liên quan

### 🔨 Build image từ Dockerfile

```shell
docker build -t <tên_image>:<phiên_bản> --force-rm -f <tên_dockerfile> .
```

|Tham số|Ý nghĩa|
|---|---|
|`-t`|Đặt tên và tag cho image (ví dụ: `myapp:1.0`)|
|`--force-rm`|Xóa các container trung gian ngay cả khi build thất bại|
|`-f`|Chỉ định tên file Dockerfile (mặc định là `Dockerfile`)|
|`.`|Đường dẫn **build context** – thư mục chứa source code|

> ⚠️ **Lưu ý:** Dấu `.` ở cuối rất quan trọng, nó chỉ định thư mục hiện tại là build context.

**Ví dụ thực tế:**

```shell
docker build -t backend-aspnet:1.0 --force-rm -f BackEndAspNet .
```

---

### 🧹 Dọn dẹp image tạm

Khi build, Docker có thể tạo ra các **image trung gian (dangling images)** gây tốn dung lượng. Dọn dẹp bằng:

```shell
docker image prune
```

---

### ▶️ Tạo và chạy container từ image

```shell
docker run -it --name <tên_container> -p <cổng_host>:<cổng_container> <tên_image>
```

|Tham số|Ý nghĩa|
|---|---|
|`-it`|Chạy ở chế độ tương tác (interactive terminal)|
|`--name`|Đặt tên cho container|
|`-p`|Ánh xạ cổng: cổng máy host → cổng trong container|

**Ví dụ:**

```shell
docker run -it --name my-backend -p 8080:80 backend-aspnet:1.0
```

> Truy cập ứng dụng qua `http://localhost:8080`

---

## 📋 3. Cú pháp viết chỉ thị trong Dockerfile

```dockerfile
<TÊN_CHỈ_THỊ> <tham số>
```

> Ví dụ: `RUN yum update -y`

- **Tên chỉ thị** thường viết **HOA** (theo quy ước)
- Mỗi chỉ thị tạo ra một **layer** mới trong image
- Thứ tự các chỉ thị **rất quan trọng** vì Docker build từ trên xuống

---

## 🧱 4. Các chỉ thị quan trọng trong Dockerfile

### 4.1 `FROM` – Chỉ định image gốc _(Bắt buộc)_

```dockerfile
FROM <image>:<tag>
```

- **Mọi Dockerfile đều phải bắt đầu bằng `FROM`**
- Chỉ định image nền để xây dựng image mới
- Nên chọn image gốc **chính thức và nhỏ gọn** (ví dụ: `alpine`)

```dockerfile
FROM ubuntu:22.04
FROM node:20-alpine
FROM mcr.microsoft.com/dotnet/aspnet:8.0
```

> 💡 Hãy đọc tài liệu của image gốc trên Docker Hub để biết nó hỗ trợ các lệnh gì.

---

### 4.2 `WORKDIR` – Thiết lập thư mục làm việc

```dockerfile
WORKDIR <đường_dẫn>
```

- Đặt thư mục làm việc cho tất cả các chỉ thị phía sau: `RUN`, `COPY`, `ADD`, `CMD`, `ENTRYPOINT`
- Nếu thư mục chưa tồn tại, Docker sẽ **tự động tạo**
- Nên dùng thay vì `RUN cd /app` để code rõ ràng hơn

```dockerfile
WORKDIR /app
```

---

### 4.3 `COPY` và `ADD` – Sao chép file vào image

```dockerfile
COPY <nguồn> <đích>
ADD  <nguồn> <đích>
```

|Chỉ thị|Chức năng|
|---|---|
|`COPY`|Chỉ sao chép file/thư mục từ máy host vào image|
|`ADD`|Giống `COPY` nhưng **thêm tính năng**: tự giải nén `.tar`, hỗ trợ URL|

> ✅ **Khuyến nghị:** Ưu tiên dùng `COPY` vì rõ ràng hơn. Chỉ dùng `ADD` khi cần giải nén tự động.

```dockerfile
# Sao chép file cụ thể
COPY package.json ./

# Sao chép toàn bộ thư mục hiện tại vào /app
COPY . /app

# ADD với URL (ít dùng)
ADD https://example.com/file.tar.gz /tmp/
```

---

### 4.4 `RUN` – Chạy lệnh trong quá trình build

```dockerfile
RUN <lệnh shell>
RUN ["executable", "param1", "param2"]
```

- Thực thi lệnh trong một **shell mới** bên trong container tạm
- Kết quả được **lưu vào layer** của image
- Thường dùng để cài đặt phần mềm, cấu hình hệ thống

```dockerfile
# Cài gói trên Ubuntu
RUN apt-get update && apt-get install -y curl git

# Cài dependencies Node.js
RUN npm install

# Build ứng dụng Go
RUN go build -o main .
```

> ⚡ **Mẹo hiệu suất:** Gộp nhiều lệnh vào một `RUN` bằng `&&` để giảm số layer:
> 
> ```dockerfile
> # ❌ Không nên (tạo 3 layer)
> RUN apt-get update
> RUN apt-get install -y curl
> RUN apt-get clean
> 
> # ✅ Nên dùng (chỉ 1 layer)
> RUN apt-get update && apt-get install -y curl && apt-get clean
> ```

---

### 4.5 `ENV` – Khai báo biến môi trường

```dockerfile
ENV <tên_biến>=<giá_trị>
```

- Biến môi trường được **kế thừa** vào container khi chạy
- Có thể ghi đè bằng `docker run -e VAR=value`

```dockerfile
ENV APP_ENV=production
ENV PORT=8080
ENV DB_HOST=localhost DB_PORT=5432
```

---

### 4.6 `ARG` – Biến khi build (Build-time Variable)

```dockerfile
ARG <tên>=<giá_trị_mặc_định>
```

- Khác với `ENV`: biến `ARG` **chỉ tồn tại trong quá trình build**, không có trong container khi chạy
- Có thể truyền giá trị khi build bằng `--build-arg`

```dockerfile
ARG APP_VERSION=1.0
ARG GO_VERSION=1.22

FROM golang:${GO_VERSION}-alpine
```

```shell
# Truyền giá trị khi build
docker build --build-arg APP_VERSION=2.0 -t myapp:2.0 .
```

---

### 4.7 `EXPOSE` – Khai báo cổng lắng nghe

```dockerfile
EXPOSE <cổng>/<giao_thức>
```

- **Khai báo** cổng mà ứng dụng bên trong container lắng nghe
- Đây chỉ là **tài liệu hóa** (documentation), không tự động publish cổng ra ngoài
- Để truy cập từ bên ngoài, vẫn cần dùng `-p` khi `docker run`

```dockerfile
EXPOSE 8080
EXPOSE 443/tcp
EXPOSE 5353/udp
```

---

### 4.8 `CMD` và `ENTRYPOINT` – Lệnh khởi chạy container

Cả hai đều chỉ định lệnh chạy khi container khởi động, nhưng có sự khác biệt quan trọng:

||`CMD`|`ENTRYPOINT`|
|---|---|---|
|Mục đích|Lệnh mặc định, **có thể ghi đè** khi `docker run`|Lệnh chính, **khó ghi đè** hơn|
|Ghi đè|`docker run image <lệnh_mới>`|Cần dùng `--entrypoint`|
|Dùng khi|Muốn cho phép user thay lệnh|Muốn container luôn chạy một chương trình cố định|

```dockerfile
# Dạng shell (chạy qua /bin/sh -c)
CMD go run main.go

# Dạng exec (khuyến nghị - không dùng shell)
CMD ["go", "run", "main.go"]
ENTRYPOINT ["./myapp"]
```

> ✅ **Kết hợp ENTRYPOINT + CMD:**
> 
> ```dockerfile
> ENTRYPOINT ["./myapp"]
> CMD ["--port", "8080"]   # tham số mặc định, có thể ghi đè
> ```

---

### 4.9 `VOLUME` – Gắn ổ đĩa chia sẻ

```dockerfile
VOLUME ["/data"]
```

- Tạo **mount point** để lưu trữ dữ liệu bền vững (persist data)
- Dữ liệu trong Volume **không bị mất** khi container bị xóa

```dockerfile
VOLUME ["/var/log/app"]
VOLUME ["/data/db"]
```

---

### 4.10 `LABEL` – Thêm metadata cho image

```dockerfile
LABEL <key>="<value>"
```

- Thêm thông tin mô tả cho image (tác giả, phiên bản, mô tả...)
- Hữu ích cho việc quản lý và tìm kiếm image

```dockerfile
LABEL maintainer="nguyenvana@gmail.com"
LABEL version="1.0"
LABEL description="Backend service cho hệ thống bầu cử trực tuyến"
```

---

### 4.11 `USER` – Chỉ định user thực thi

```dockerfile
USER <username>
```

- Chỉ định user sẽ thực thi các lệnh `RUN`, `CMD`, `ENTRYPOINT` phía sau
- **Quan trọng về bảo mật:** Không nên chạy ứng dụng với quyền `root`

```dockerfile
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser
```

---

### 4.12 `HEALTHCHECK` – Kiểm tra sức khỏe container

```dockerfile
HEALTHCHECK --interval=<thời_gian> --timeout=<thời_gian> CMD <lệnh>
```

- Docker sẽ định kỳ chạy lệnh này để kiểm tra container còn hoạt động không
- Kết quả: `healthy`, `unhealthy`, hoặc `starting`

```dockerfile
HEALTHCHECK --interval=30s --timeout=10s --retries=3 \
  CMD curl -f http://localhost:8080/health || exit 1
```

---

### 4.13 `ONBUILD` – Chỉ thị trì hoãn

```dockerfile
ONBUILD <chỉ_thị_khác>
```

- Các chỉ thị này **chưa thực thi ngay**, mà chỉ thực thi khi image này được dùng làm **base image** cho image khác
- Thường dùng cho các **base image tái sử dụng**

```dockerfile
ONBUILD COPY . /app
ONBUILD RUN go build -o main .
```

---

## 🔄 5. Multi-Stage Build – Tối ưu kích thước image

**Vấn đề:** Image chứa cả SDK/compiler thường rất nặng (hàng GB), trong khi chạy production chỉ cần binary/runtime.

**Giải pháp:** Dùng **multi-stage build** – build trong một container, chỉ lấy kết quả sang container khác nhỏ hơn.

```
Stage 1 (builder)     Stage 2 (runtime)
──────────────────    ──────────────────
SDK + source code  →  Chỉ lấy binary/output
(~1GB)                (~50MB)
```

```dockerfile
# Stage 1: Build
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o myapp .

# Stage 2: Runtime (image rất nhỏ)
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/myapp .   # Chỉ lấy binary từ stage build
ENTRYPOINT ["./myapp"]
```

---

## 🚫 6. File `.dockerignore` – Loại trừ file không cần thiết

Tương tự `.gitignore`, file `.dockerignore` giúp:

- Giảm **build context** gửi lên Docker daemon
- Tránh copy các file nhạy cảm vào image
- **Tăng tốc** quá trình build

```dockerignore
# Dependencies
node_modules/
vendor/

# Version control
.git/
.gitignore

# Logs
*.log
logs/

# Environment files (bảo mật!)
.env
.env.local

# Build artifacts
dist/
build/
*.exe
```

---

## 💡 7. Ví dụ thực tế

### 7.1 Ví dụ – Backend ASP.NET Core (C#)

```dockerfile
# ========== Stage 1: Build ==========
FROM mcr.microsoft.com/dotnet/sdk:8.0 AS build
WORKDIR /app

# Sao chép file project và restore packages (tận dụng layer cache)
COPY ["BackEnd.csproj", "./"]
RUN dotnet restore "BackEnd.csproj"

# Sao chép toàn bộ source và publish
COPY . ./
RUN dotnet publish "BackEnd.csproj" -c Release -o /app/out

# ========== Stage 2: Runtime ==========
FROM mcr.microsoft.com/dotnet/aspnet:8.0
WORKDIR /app

LABEL maintainer="dev@example.com"
LABEL description="Backend hệ thống bầu cử trực tuyến"

# Chỉ sao chép output từ stage build
COPY --from=build /app/out ./

EXPOSE 80

ENTRYPOINT ["dotnet", "BackEnd.dll"]
```

---

### 7.2 Ví dụ – Backend Golang (REST API)

Giả sử bạn có một project Go đơn giản với cấu trúc:

```
my-go-api/
├── main.go
├── go.mod
├── go.sum
├── handlers/
│   └── user.go
└── Dockerfile
```

**`main.go` (minh họa):**

```go
package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello from Go API!")
    })
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "OK")
    })
    log.Println("Server running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

**`Dockerfile` cho project Go:**

```dockerfile
# ========== Stage 1: Build ==========
FROM golang:1.22-alpine AS builder

# Cài đặt các công cụ cần thiết
RUN apk add --no-cache git

WORKDIR /app

# Copy go.mod và go.sum trước để tận dụng Docker layer cache
# (chỉ re-download modules khi go.mod thay đổi)
COPY go.mod go.sum ./
RUN go mod download

# Copy toàn bộ source code
COPY . .

# Build binary, tắt CGO để chạy trên Alpine không cần thư viện C
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# ========== Stage 2: Runtime ==========
# Dùng image cực nhỏ (~5MB) – chỉ chứa binary
FROM alpine:3.19

# Cài ca-certificates để hỗ trợ HTTPS
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Thêm metadata
LABEL maintainer="dev@example.com"
LABEL version="1.0"
LABEL description="REST API viết bằng Golang"

# Sao chép binary từ stage build
COPY --from=builder /app/main .

# Khai báo cổng ứng dụng lắng nghe
EXPOSE 8080

# Kiểm tra sức khỏe mỗi 30 giây
HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 \
  CMD wget -qO- http://localhost:8080/health || exit 1

# Chạy ứng dụng
ENTRYPOINT ["./main"]
```

**Build và chạy:**

```shell
# Build image
docker build -t my-go-api:1.0 .

# Chạy container
docker run -d --name go-api -p 8080:8080 my-go-api:1.0

# Kiểm tra
curl http://localhost:8080/
# → Hello from Go API!

curl http://localhost:8080/health
# → OK
```

**So sánh kích thước:**

```
golang:1.22-alpine (full)  → ~270MB
my-go-api:1.0 (multi-stage) → ~15MB  ✅ Nhỏ hơn ~18 lần!
```

---

## 🐛 8. Troubleshooting – Xử lý lỗi thường gặp

### Kiểm tra lỗi khi build

```shell
docker build -t my_image:1.0 -f Dockerfile .
```

> Docker sẽ hiển thị từng bước, lỗi sẽ xuất hiện tại bước thất bại

### Xem logs của container để debug

```shell
docker logs <tên_container hoặc container_ID>

# Theo dõi logs realtime
docker logs -f <tên_container>
```

### Xem các layer của image

```shell
docker history <image_ID hoặc tên_image>
```

> Hiển thị từng layer, kích thước, và lệnh tạo ra nó

### Vào bên trong container để kiểm tra

```shell
docker exec -it <tên_container> /bin/sh
# hoặc nếu có bash:
docker exec -it <tên_container> /bin/bash
```

### Kiểm tra thông tin chi tiết của image

```shell
docker inspect <tên_image>
```

---

## ✅ 9. Tổng hợp các chỉ thị Dockerfile

|Chỉ thị|Chức năng chính|
|---|---|
|`FROM`|Image gốc để xây dựng _(bắt buộc)_|
|`WORKDIR`|Thư mục làm việc trong container|
|`COPY`|Sao chép file từ máy host vào image|
|`ADD`|Như COPY, thêm hỗ trợ giải nén và URL|
|`RUN`|Thực thi lệnh khi **build** image|
|`CMD`|Lệnh mặc định khi **chạy** container (ghi đè được)|
|`ENTRYPOINT`|Lệnh chính khi **chạy** container (cố định hơn)|
|`ENV`|Biến môi trường (tồn tại cả khi chạy)|
|`ARG`|Biến build-time (chỉ tồn tại lúc build)|
|`EXPOSE`|Khai báo cổng lắng nghe (tài liệu hóa)|
|`VOLUME`|Tạo mount point cho dữ liệu bền vững|
|`USER`|Chỉ định user thực thi lệnh|
|`LABEL`|Thêm metadata (tác giả, phiên bản...)|
|`HEALTHCHECK`|Định nghĩa cách kiểm tra sức khỏe container|
|`ONBUILD`|Chỉ thị thực thi khi image được dùng làm base|

---

> 📚 **Tài liệu tham khảo:**
> 
> - [Docker Official Documentation](https://docs.docker.com/engine/reference/builder/)
> - [Docker Hub](https://hub.docker.com/) – Tìm kiếm các image chính thức
> - [Best practices for writing Dockerfiles](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/)