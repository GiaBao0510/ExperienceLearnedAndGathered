- Trong Docker cho phép tạo ra một network (giao tiếp mạng), sau đó các container kết nốt vào với network. Khi mà các container mà kết nối cùng network thì chúng có thể giao tiếp với nhau một cách nhanh chóng thông qua tên container và cổng được lắng nghe của container trên mạng đó.

- Các container thường được khởi tạo sẽ mặt định kết nối đến network Bridge.

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

---
### **Ánh xạ cổng mạng trong Docker:**

- Ánh xạ cổng từ container đến cổng của máy host
```
docker run -it --name [Tên container tự đặt] -p [Số cổng muốn ánh xạ]:[Cổng container đang tạo] [Tên Image]
```

_Ví dụ:_ Ánh xạ cổng 80 của container có thên là B2 vào máy host(127.0.0.1) có cổng là  8888
```
docker run -it --name B2 -p 8888:80 busybox
```

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

### **Nối container đang chạy vào một network:**

- Cú pháp nối một container đang chạy trên mạng khác kết nối với một mạng khác nữa:
```
docker network connect [Tên mạng cần kết nối] [Tên container/ container ID]
```


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
