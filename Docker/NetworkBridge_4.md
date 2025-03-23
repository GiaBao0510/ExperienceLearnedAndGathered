### **Network Bridge trong Docker là gì?**

**Network Bridge** là một loại mạng ảo trong **Docker**, sử dụng driver ==bridge== (cầu nối), cho phép các container kết nối với nhau và với máy host thông qua một giao diện mạng chung. Theo mặc định, khi bạn chạy một container mà không chỉ định mạng, nó sẽ tự động kết nối vào mạng ==bridge== mặc định của Docker (thường có tên là ==bridge==). Đây là cơ chế cơ bản để đảm bảo các container có thể giao tiếp với nhau hoặc với thế giới bên ngoài (qua máy host).

![](https://towardsdatascience.com/wp-content/uploads/2020/07/1DQ7oHTHIgHOvfTqk58NR-A-1536x880.png)

---
##### **Cách hoạt động**:
- **Network Bridge** tạo ra một ==mạng nội bộ (private network)== trong máy host.
- Mỗi **container** kết nối vào mạng này được gán một địa chỉ IP riêng trong dải mạng nội bộ (ví dụ: 172.17.0.x trong mạng bridge mặc định).
- Các **container** trong cùng một mạng **bridge** có thể giao tiếp trực tiếp với nhau bằng tên **container** (nhờ DNS nội bộ của Docker) hoặc địa chỉ IP.
- Để giao tiếp với máy **host** hoặc mạng bên ngoài, các cổng của **container** cần được ánh xạ (**port mapping**) ra **host**.

![](https://miro.medium.com/v2/resize:fit:720/format:webp/1*oOEiI1Ssa9h8QNVkgCzlgw.png)

---
### **Tại sao cần Network Bridge?**

- **Giao tiếp giữa các container**: Trong ứng dụng đa container (ví dụ: một web server và một database), các container cần trao đổi dữ liệu nhanh chóng và an toàn.
- **Cách ly mạng**: Mỗi mạng bridge là một môi trường biệt lập, giúp tăng tính bảo mật và tránh xung đột giữa các ứng dụng.
- **Kết nối với thế giới bên ngoài**: Ánh xạ cổng từ container ra host cho phép người dùng hoặc dịch vụ bên ngoài truy cập container.
- **Quản lý đơn giản**: Không cần cấu hình mạng phức tạp, Docker tự động xử lý NAT (Network Address Translation) và DNS cho các container trong cùng mạng.

---
### **So sánh Network Bridge với các loại mạng khác trong Docker**

|Loại mạng|Đặc điểm|Trường hợp sử dụng|
|---|---|---|
|**Bridge**|Mạng nội bộ, cách ly, hỗ trợ ánh xạ cổng|Ứng dụng đa container trên 1 host|
|**Host**|Container dùng mạng của host trực tiếp|Khi cần hiệu suất mạng tối đa|
|**Overlay**|Kết nối container trên nhiều host (Docker Swarm)|Hệ thống phân tán, cluster|
|**None**|Không có mạng, container bị cô lập hoàn toàn|Trường hợp đặc biệt, bảo mật cao|

---
#### **Cú pháp:**

- Mạng bridge mặc định được tạo sẵn trong Docker. Khi không chỉ định mạng, container sẽ tự động kết nối vào đây

- Liệt kê các network đang có:  
```
docker network ls
```

- Kiểm tra xem network có bao nhiêu container kết nối vào:
```
docker network inspect [tên network]
```

- Lệnh kiểm tra container kết nối đến mạng nào
```
docker inspect [Tên container / container ID]
```

- **Xóa tất cả các network không được sử dụng**:
```
docker network prune
```

- **Xem thông tin chi tiết về network**:
```
docker network inspect [Tên network]
```

##### ***Ví dụ:*** Chạy một container Nginx và kiểm tra mạng.

```
docker run -d --name web_server nginx
docker network inspect bridge
```
- **Kết quả**: Trong output của ==docker network inspect bridge==, bạn sẽ thấy *web_server* được liệt kê trong phần "**Containers**", với một địa chỉ IP như 172.17.0.2.

**Kiểm tra container dùng mạng nào**:
```
docker inspect web_server | findstr NetworkMode
```
- **Kết quả**: ==NetworkMode: default== (tức là dùng mạng bridge mặc định).

---
### **Ánh xạ cổng mạng trong Docker:**

- Ánh xạ cổng từ container đến cổng của máy host
```
docker run -it --name [Tên container tự đặt] -p [Số cổng muốn ánh xạ]:[Cổng container đang tạo] [Tên Image]
```

**_Ví dụ:_**  Ánh xạ cổng 80 của container có tên là B2 vào máy host(127.0.0.1) có cổng là  8888
```
docker run -it --name B2 -p 8888:80 busybox
```
- **Giải thích**: Cổng 80 của container được ánh xạ ra cổng 8080 của host. Truy cập ==localhost:8080== trên trình duyệt để thấy trang mặc định của busybox.

- Tạo mạng cầu:
```
docker network create --driver bridge [Tên mạng]
```

- Xóa mạng đã tạo trên docker:
```
docker network rm [Tên mạng cần xóa]
```

- Chỉ định container kết nối mạng nào đó khi khởi tạo:
```
docker run -it --name [Đặt tên container] --network [Tên mạng có sẵn] [image]
```

- Tạo container chạy trên mạng cụ thể có ánh xạ đến cổng của máy chủ
```
docker run -it --name [Đặt tên container] --network [Tên mạng có sẳn] -p [Cổng host]:[Cổng container] [Image]
```

##### ***Ví dụ:*** Tạo và sử dụng mạng Bridge tùy chỉnh

**Lý thuyết**: Bạn có thể tạo mạng bridge riêng để cách ly các nhóm container, tăng tính bảo mật và quản lý tốt hơn.

- **Bước 1**: Tạo mạng bridge mới.
```
docker network create --driver bridge my_network
```

- **Bước 2**: Chạy hai container trong mạng vừa tạo.
```
docker run -d --name web --network my_network nginx
docker run -it --name client --network my_network busybox
```

 **Kiểm tra giao tiếp:**
- Trong container client, chạy:
```
ping web
```
- **Kết quả**: Container client có thể ping được web bằng tên nhờ DNS nội bộ của Docker.

**Kiểm tra thông tin mạng**:
```
docker network inspect my_network
```

**Kết quả mẫu**:
```
{
  "Name": "my_network",
  "Driver": "bridge",
  "Containers": {
    "web": {"IPv4Address": "172.18.0.2/16"},
    "client": {"IPv4Address": "172.18.0.3/16"}
  }
}
```

---
### **Nối container đang chạy vào một network:**

Bạn có thể kết nối một container đang chạy vào một mạng khác mà không cần khởi động lại, giúp linh hoạt trong quản lý.

- Cú pháp nối một container đang chạy trên mạng khác kết nối với một mạng khác nữa:
```
docker network connect [Tên mạng cần kết nối] [Tên container/ container ID]
```

##### ***Ví dụ:*** Kết nối container đang chạy vào mạng khác

- **Bước 1**: Chạy container trong mạng mặc định.
```
docker run -d --name db_server mysql
```

- **Bước 2**: Kết nối container vào mạng tùy chỉnh.
```
docker network connect my_network db_server
```

- **Kiểm tra**:
```
docker network inspect my_network
docker inspect db_server
```
- **Kết quả**: db_server sẽ xuất hiện trong my_network và vẫn giữ kết nối với mạng bridge mặc định. Container giờ thuộc cả hai mạng.

---
### **Ví dụ:**

- **Ví dụ về `docker run` với network**:
```
docker run -it --name my_container --network my_network image_ID
```
*Lệnh này sẽ tạo một container mới và kết nối nó vào network `my_network`.*

- **Ví dụ về `docker network create`**:
```
docker network create --driver bridge my_network
```
*Lệnh này sẽ tạo một network mới có tên `my_network` với driver `bridge`.*

- **Ví dụ về `docker network connect`**:
```
docker network connect my_network my_container
```
*Lệnh này sẽ kết nối container `my_container` vào network `my_network`.*

---
#### **Ứng dụng thực tế**

Dựa trên danh sách container bạn cung cấp trước đó (==mongodb, my_mysql, myredis==), đây là cách áp dụng Network Bridge:

**Tạo mạng cho ứng dụng**:
```
docker network create app_network
```

**Kết nối các container**:
```
docker network connect app_network mongodb
docker network connect app_network my_mysql
docker network connect app_network myredis
```

**Kiểm tra giao tiếp**:
- Từ một container bất kỳ (ví dụ: my_mysql), bạn có thể dùng ping mongodb hoặc kết nối trực tiếp qua cổng (ví dụ: 27017 cho MongoDB).

---
## **Kết luận**

- **Network Bridge mặc định**: Đơn giản, phù hợp cho các ứng dụng cơ bản.
- **Network Bridge tùy chỉnh**: Tốt hơn cho ứng dụng phức tạp, cần cách ly và quản lý.
- **Ánh xạ cổng**: Cầu nối giữa container và thế giới bên ngoài.