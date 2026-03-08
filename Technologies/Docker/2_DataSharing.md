### **Data Sharing trong Docker là gì?**

![](https://miro.medium.com/v2/resize:fit:720/format:webp/1*xONk464vW-xNYxzE_HsSkw.png)

**Data Sharing (Chia sẻ dữ liệu)** trong Docker là quá trình cho phép các container hoặc giữa container và máy host ==truy cập, lưu trữ hoặc sử dụng chung dữ liệu==. **Docker container** theo thiết kế là **stateless** (không lưu trạng thái), ==nghĩa là khi container bị xóa, mọi dữ liệu bên trong nó cũng mất đi trừ khi được lưu trữ ra bên ngoài.== Data Sharing giải quyết vấn đề này bằng cách cung cấp cơ chế để:

1. **Lưu trữ dữ liệu lâu dài**: Đảm bảo dữ liệu không bị mất khi container dừng hoặc bị xóa.
2. **Chia sẻ dữ liệu giữa các container**: Cho phép nhiều container cùng truy cập vào một nguồn dữ liệu chung.
3. **Tích hợp với hệ thống host**: Kết nối dữ liệu từ máy host vào container hoặc ngược lại.

---
### **Tại sao cần Data Sharing?**

- **Tính bền vững (Persistence)**: Các ứng dụng như cơ sở dữ liệu (MySQL, MongoDB) cần lưu trữ dữ liệu lâu dài. Nếu không có Data Sharing, dữ liệu sẽ mất khi container dừng.
- **Hiệu quả tài nguyên**: Nhiều container có thể dùng chung dữ liệu mà không cần sao chép dư thừa.
- **Tính linh hoạt**: Dễ dàng di chuyển dữ liệu giữa các môi trường (host sang container, container sang container).
- **Quản lý ứng dụng thực tế**: Trong môi trường sản xuất, các microservices thường cần chia sẻ cấu hình, log, hoặc dữ liệu giao dịch.

---
### **Các phương pháp chính để chia sẻ dữ liệu trong Docker**

Docker hỗ trợ 3 cơ chế chính cho Data Sharing:

1. **Bind Mounts**: Ánh xạ trực tiếp một thư mục/tệp từ máy host vào container.
2. **Volumes**: Sử dụng các ổ đĩa ảo do Docker quản lý để lưu trữ và chia sẻ dữ liệu.
3. **Volumes-from**: Chia sẻ dữ liệu từ một container khác sang container mới.

##### **So sánh các phương pháp**

| Phương pháp      | Ưu điểm                              | Nhược điểm                           | Trường hợp sử dụng                  |
| ---------------- | ------------------------------------ | ------------------------------------ | ----------------------------------- |
| **Bind Mounts**  | Dễ cấu hình, linh hoạt với host      | Phụ thuộc vào hệ thống file của host | Phát triển, kiểm thử, log trực tiếp |
| **Volumes**      | Độc lập với host, được Docker tối ưu | Cần quản lý thêm volume              | Ứng dụng sản xuất, cơ sở dữ liệu    |
| **Volumes-from** | Dễ chia sẻ giữa container            | Phụ thuộc vào container nguồn        | Nhiều container dùng chung dữ liệu  |

---
##### 1.**Ánh xạ dữ liệu giữa máy host và container**

![](https://docker-docs.uclv.cu/storage/images/types-of-mounts-bind.png)

**Lý thuyết**: Bind Mounts cho phép bạn ánh xạ trực tiếp một thư mục hoặc tệp từ máy host vào container. Điều này hữu ích khi bạn muốn sử dụng dữ liệu từ host (ví dụ: file cấu hình, log) hoặc lưu kết quả từ container ra host.

_Cú pháp:_ 
``` powershell
docker run -it -v [pathHost]:[pathContainer] image_ID
```

***Ví dụ:*** Chạy một container Nginx và ánh xạ thư mục chứa file HTML từ host.
```powershell
docker run -d -p 8080:80 -v /home/user/html:/usr/share/nginx/html nginx:latest
```
- **Giải thích**:
	- ==/home/user/html== (host) được ánh xạ vào ==/usr/share/nginx/html== (container).
	- Khi truy cập ==localhost:8080==, bạn sẽ thấy nội dung từ thư mục host.

***Ví dụ 2**:* Lưu log từ container ra host.
```powershell
docker run -d -v /var/logs/myapp:/app/logs my-app-image
```
- **Kết quả**: Log sinh ra trong ==/app/logs== của container sẽ được lưu vào ==/var/logs/myapp== trên host.

---
##### 2. **Chia sẻ dữ liệu giữa các container với --volumes-from**
 
 - Phương pháp `--volumes-from` cho phép một container mới kế thừa các **volume** từ một **container** đã có. Điều này hữu ích khi bạn muốn nhiều container cùng truy cập dữ liệu mà không cần tạo volume riêng.
 - Chia sẻ dữ liệu giữa các container dựa trên container có sẵn và image ID của container đó
```shell
docker run -it --name [Tên container mới] --volumes-from [Tên container có sẵn] [image ID]
```

***Ví dụ:***
**Bước 1**: Tạo container nguồn với volume.
```shell
docker run -d --name data_container -v /shared-data busybox
```
- **Giải thích**: Container *data_container* tạo một *volume* tại */shared-data*

**Bước 2**: Tạo container mới chia sẻ volume từ container nguồn.
```shell
docker run -it --name new_container --volumes-from data_container busybox
```

- **Kiểm tra**: Trong *new_container*, chạy *ls /shared-data* để thấy dữ liệu từ container nguồn.
- **Ứng dụng thực tế**: Chạy một container ghi dữ liệu và một container đọc dữ liệu.

```shell
docker run -d --name writer --volumes-from data_container busybox sh -c "echo 'Hello' > /shared-data/test.txt"

docker run -it --name reader --volumes-from data_container busybox cat /shared-data/test.txt
```
- **Kết quả**: Container reader sẽ in ra Hello.

---
##### 3.**Chia sẻ qua volume**

![](https://i0.wp.com/collabnix.com/wp-content/uploads/2024/04/Feature-image-Intro-to-docker-volumes.png?resize=1536%2C806&ssl=1)

Volume là cách lưu trữ dữ liệu bền vững và độc lập với **host**, được Docker quản lý trong thư mục mặc định (thường là /var/lib/docker/volumes). **Volume** ==không mất khi container bị xóa và có thể gắn vào nhiều container==.

 - Kiểm tra xem hiện tại có những ổ đĩa nào
 ```shell
 docker volume ls
```

**Tạo volume**
```
docker volume create [NameDisk]
```

- Kiểm tra thông tin ổ đĩa:
```shell
docker volume inspect [NameDisk]
```

- Xóa ổ đĩa:
```shell
docker volume rm [NameDisk]
```

**Xem danh sách các volume**:
```shell
docker volume ls
```

**Xóa tất cả các volume không được sử dụng**:
```shell
docker volume prune
```

**Xem thông tin chi tiết về volume**:
```shell
docker volume inspect [NameDisk]
```

- **Kết quả mẫu:** 
```json
[
    {
        "Name": "my_volume",
        "Driver": "local",
        "Mountpoint": "/var/lib/docker/volumes/my_volume/_data"
    }
]
```

 **Gắn ổ đĩa volume vào container:**
```shell
docker run -it --mount source=DISK,target=pathContainer ImageID
```

> Dữ liệu sẽ không mất đi khi xóa container

 
 **Tạo ổ đĩa ánh xạ đến thư mục host**
```shell
docker volume create --opt device=pathHOST --opt type=none --otp o=bind DISKNAME
```


---
### **Ví dụ minh họa:**

- **Ví dụ về `docker run` với volume**:
```shell
docker run -it -v /host/path:/container/path image_ID
```
*Lệnh này sẽ mount thư mục `/host/path` từ host vào thư mục `/container/path` trong container.*

- **Ví dụ về `docker run` với `--volumes-from`**:
```shell
docker run -it --name new_container --volumes-from existing_container -image_ID
```
*Lệnh này sẽ tạo một container mới và chia sẻ volume từ container `existing_container`.*

- **Ví dụ về `docker volume create`**:
```shell
docker volume create my_volume
```
*Lệnh này sẽ tạo một volume mới có tên `my_volume`.*

---
## **Kết luận**

- **Bind Mounts**: Phù hợp cho phát triển hoặc khi cần truy cập trực tiếp dữ liệu host.
- **Volumes**: Tốt nhất cho môi trường sản xuất, lưu trữ dữ liệu bền vững.
- **Volumes-from**: Hữu ích khi cần chia sẻ nhanh giữa các container