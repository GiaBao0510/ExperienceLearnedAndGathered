# Docker Compose

Docker Compose là một công cụ mạnh mẽ được sử dụng để **định nghĩa và quản lý các ứng dụng Docker multi-container**. Docker Compose cho phép bạn sử dụng một file YAML để cấu hình các dịch vụ của ứng dụng (`docker-compose.yml`), sau đó chỉ với **một lệnh duy nhất**, bạn có thể tạo và khởi động tất cả các dịch vụ từ cấu hình đó.

## Tổng Quan Về Docker Compose

Docker Compose giúp đơn giản hóa quá trình quản lý các ứng dụng phức tạp bằng cách cho phép bạn định nghĩa tất cả các dịch vụ, mạng và volume cần thiết trong **một file duy nhất**. Điều này giúp bạn dễ dàng tái tạo môi trường phát triển, kiểm thử và sản xuất một cách nhất quán.

Quy trình sử dụng Docker Compose chỉ gồm **ba bước đơn giản**:

1. Khai báo môi trường của ứng dụng trong `Dockerfile`.
2. Khai báo các services cần thiết để chạy ứng dụng trong file `docker-compose.yml`.
3. Chạy lệnh `docker-compose up` để khởi động toàn bộ ứng dụng.

![](https://ttth.tlu.edu.vn/wp-content/uploads/2025/03/Gioi-thieu-ve-Docker-Compose.webp)

Ví dụ, một ứng dụng web gồm **backend**, **frontend** và **cơ sở dữ liệu** có thể được khởi động chỉ bằng một lệnh nhờ Docker Compose, thay vì phải quản lý từng container riêng lẻ. Điều này không chỉ tiết kiệm thời gian mà còn đảm bảo tính nhất quán giữa các môi trường.

---

## Lợi Ích Của Docker Compose

|Lợi ích|Mô tả|
|---|---|
|**Đơn Giản Hóa Quản Lý**|Quản lý ứng dụng multi-container dễ dàng chỉ với một file cấu hình duy nhất.|
|**Tính Nhất Quán**|Đảm bảo môi trường phát triển, kiểm thử và sản xuất luôn đồng nhất, giảm thiểu lỗi do sự khác biệt giữa các môi trường.|
|**Tự Động Hóa**|Tự động hóa quá trình xây dựng, khởi động và dừng các dịch vụ của ứng dụng.|
|**Khả Năng Mở Rộng**|Dễ dàng scale up/down các dịch vụ theo nhu cầu.|

---

## Các Thành Phần Cơ Bản Của Docker Compose

- **Compose file**: Tệp YAML (`docker-compose.yml`) chứa toàn bộ thông tin cần thiết về container để triển khai ứng dụng.
    
- **Services (Dịch vụ)**: Tập hợp các container mà bạn muốn tạo ra. Mỗi service đại diện cho một thành phần của ứng dụng, ví dụ: service `backend`, service `frontend`, service `database`...
    
- **Container**: Đối tượng chứa ứng dụng và các thành phần của nó. Docker Compose sử dụng container để triển khai các service được định nghĩa trong `docker-compose.yml`.
    
- **Networks (Mạng)**: Định nghĩa các mạng ảo để các container trong cùng hoặc khác service có thể liên lạc với nhau bằng tên hoặc địa chỉ IP.
    
- **Volumes (Khối lưu trữ)**: Cho phép lưu trữ dữ liệu của ứng dụng **độc lập với vòng đời của container**, giúp việc sao lưu, phục hồi và quản lý dữ liệu dễ dàng hơn.
    
    > ⚠️ **Lưu ý:** Nếu không dùng volume, dữ liệu sẽ **bị mất** khi container bị xóa.
    
- **Environment Variables (Biến môi trường)**: Các biến được dùng để cấu hình linh hoạt cho từng container, ví dụ: tên database, mật khẩu, chế độ chạy...
    
- **CLI (Command Line Interface)**: Giao diện dòng lệnh để tạo, triển khai, quản lý và xóa các container và service thông qua các lệnh `docker-compose`.
    

---

## Tại Sao Docker Compose Quan Trọng Với Lập Trình Viên?

**1. Tạo nhiều môi trường độc lập (Isolated Environments)**

Compose giúp cô lập môi trường hoạt động của từng project, đảm bảo chúng không xung đột với nhau. Bạn cũng có thể nhanh chóng tạo bản sao của một môi trường để thử nghiệm.

**2. Chỉ tạo lại container khi có thay đổi**

Compose nhận biết được service nào chưa thay đổi và **tái sử dụng container** tương ứng thay vì tạo lại từ đầu, giúp tiết kiệm thời gian.

**3. Hỗ trợ biến môi trường linh hoạt**

Compose cho phép dùng các biến trong file cấu hình. Với mỗi môi trường hoặc người dùng khác nhau, bạn có thể dễ dàng điều chỉnh các biến mà không cần sửa toàn bộ cấu hình.

---

## Cấu Trúc File `docker-compose.yml`

File `docker-compose.yml` là file cấu hình chính, chứa định nghĩa về các service, mạng và volume của ứng dụng. Dưới đây là ví dụ minh họa:

```yaml
version: "3"
services:
  web:
    image: nginx
    ports:
      - "8080:80"

  app:
    build: .
    ports:
      - "3000:3000"
    environment:
      NODE_ENV: production
    volumes:
      - .:/app
    depends_on:
      - db

  db:
    image: postgres
    environment:
      POSTGRES_USER: example
      POSTGRES_PASSWORD: example
    volumes:
      - db_data:/var/lib/postgresql/data

volumes:
  db_data:
```

**Giải thích các thành phần trong ví dụ trên:**

| Thành phần    | Ý nghĩa                                                                                                                                                   |
| ------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `version`     | Phiên bản cú pháp của Docker Compose (hiện tại phổ biến là `"3"` hoặc `"3.x"`).                                                                           |
| `services`    | Khai báo tất cả các dịch vụ cần chạy.                                                                                                                     |
| `web`         | Dịch vụ web dùng image `nginx`, ánh xạ cổng `8080` của máy host sang cổng `80` của container.                                                             |
| `app`         | Dịch vụ ứng dụng, build từ `Dockerfile` ở thư mục hiện tại, mở cổng `3000`, thiết lập biến môi trường và gắn thư mục làm việc vào `/app` trong container. |
| `db`          | Dịch vụ cơ sở dữ liệu dùng image `postgres`, thiết lập tài khoản đăng nhập qua biến môi trường và lưu dữ liệu vào volume `db_data`.                       |
| `image`       | Dùng image có sẵn từ Docker Hub.                                                                                                                          |
| `build`       | Tự build image từ mã nguồn cục bộ (yêu cầu có `Dockerfile`).                                                                                              |
| `ports`       | Ánh xạ cổng theo dạng `"host:container"`.                                                                                                                 |
| `volumes`     | Gắn dữ liệu từ host hoặc volume vào container.                                                                                                            |
| `environment` | Thiết lập biến môi trường bên trong container.                                                                                                            |
| `depends_on`  | Khai báo thứ tự khởi động — service `app` sẽ chờ service `db` khởi động trước.                                                                            |

> ⚠️ **Lưu ý về `depends_on`:** `depends_on` chỉ đảm bảo **thứ tự khởi động container**, không đảm bảo service bên trong đã sẵn sàng nhận kết nối (ví dụ: PostgreSQL chưa hoàn tất khởi tạo). Để xử lý vấn đề này, bạn có thể dùng thêm `healthcheck` hoặc thư viện hỗ trợ như `wait-for-it`.

---

## Các Lệnh Docker Compose Cơ Bản

#### 1. `docker-compose up`

Tạo và khởi động tất cả các dịch vụ được định nghĩa trong `docker-compose.yml`.

```bash
docker-compose up
```

Thêm tùy chọn `-d` để chạy ở **chế độ nền** (detached mode):

```bash
docker-compose up -d
```

#### 2. `docker-compose down`

Dừng và **xóa** tất cả các container, mạng được tạo bởi `docker-compose up`.

```bash
docker-compose down
```

Thêm `--volumes` để xóa cả volume đi kèm:

```bash
docker-compose down --volumes
```

#### 3. `docker-compose build`

Xây dựng hoặc tái xây dựng image cho các dịch vụ có khai báo `build`.

```bash
docker-compose build
```

#### 4. `docker-compose ps`

Liệt kê trạng thái hiện tại của các container đang được quản lý.

```bash
docker-compose ps
```

#### 5. `docker-compose logs`

Hiển thị log của tất cả hoặc một dịch vụ cụ thể.

```bash
docker-compose logs
docker-compose logs <service_name>
```

Thêm `-f` để theo dõi log **theo thời gian thực**:

```bash
docker-compose logs -f <service_name>
```

#### 6. `docker-compose exec`

Thực thi một lệnh bên trong một container **đang chạy**.

```bash
docker-compose exec <service_name> <command>
```

_Ví dụ — mở terminal trong container `app`:_

```bash
docker-compose exec app /bin/bash
```

#### 7. `docker-compose stop` / `docker-compose start`

Dừng các container mà **không xóa** chúng, và khởi động lại khi cần.

```bash
docker-compose stop
docker-compose start
```

#### 8. `docker-compose restart`

Khởi động lại một hoặc tất cả các dịch vụ.

```bash
docker-compose restart <service_name>
```

---

## Khi Nào Nên Sử Dụng Docker Compose?

Để tận dụng tối đa Docker Compose, bạn cần hiểu rõ bối cảnh áp dụng:

✅ **Nên dùng Docker Compose khi:**

- **Giai đoạn phát triển**: Tạo môi trường local nhanh chóng để thử nghiệm ứng dụng.
- **Dự án nhỏ và vừa**: Quản lý hệ thống với số lượng container hạn chế.
- **Kiểm thử (Testing)**: Mô phỏng môi trường gần với production để chạy test tự động.

❌ **Không nên dùng Docker Compose khi:**

- Dự án yêu cầu **khả năng mở rộng lớn** hoặc **tính sẵn sàng cao** (high availability).
- Cần triển khai trên **nhiều máy chủ** (multi-host deployment).

> 💡 Trong các trường hợp trên, hãy cân nhắc chuyển sang **Kubernetes** — một nền tảng điều phối container mạnh mẽ hơn — sau khi đã hoàn thiện giai đoạn phát triển với Docker Compose.

---

## Tóm Tắt

|Khái niệm|Vai trò|
|---|---|
|`docker-compose.yml`|File cấu hình trung tâm của toàn bộ ứng dụng|
|`services`|Các thành phần/container tạo nên ứng dụng|
|`networks`|Kênh giao tiếp giữa các container|
|`volumes`|Lưu trữ dữ liệu bền vững, không bị mất khi container tắt|
|`docker-compose up`|Lệnh khởi động toàn bộ ứng dụng|
|`docker-compose down`|Lệnh dừng và dọn dẹp toàn bộ ứng dụng|