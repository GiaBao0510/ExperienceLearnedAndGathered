 **Ánh xạ dữ liệu giữa máy host và container**

_Cú pháp:_ 
``` 
docker run -it -v [pathHost]:[pathContainer] image_ID
```

 **Chia sẻ dữ liệu giữa các container**
 - Chia sẻ dữ liệu giữa các container dựa trên container có sẵn và image ID của container đó
```
docker run -it --name [Tên container mới] --volumes-from [Tên container có sẵn] [image ID]
```
 
 **Chia sẻ qua volume**
 - Kiểm tra xem hiện tại có những ổ đĩa nào
 ```
 docker volume ls
```

**Tạo volume**
```
docker volume create [NameDisk]
```

- Kiểm tra thông tin ổ đĩa:
```
docker volume inspect [NameDisk]
```

- Xóa ổ đĩa:
```
docker volume rm [NameDisk]
```

**Xem danh sách các volume**:
```
docker volume ls
```

**Xóa tất cả các volume không được sử dụng**:
```
docker volume prune
```

**Xem thông tin chi tiết về volume**:
```
docker volume inspect [NameDisk]
```

 **Gắn ổ đĩa volume vào container:**
```
docker run -it --mount source=DISK,target=pathContainer ImageID
```

> Dữ liệu sẽ không mất đi khi xóa container

 **Tạo ổ đĩa ánh xạ đến thư mục host**
```
docker volume create --opt device=pathHOST --opt type=none --otp o=bind DISKNAME
```

**Kiểm tra trạng thái volume**:
```
docker volume inspect [NameDisk]
```

---

### **Ví dụ minh họa:**

- **Ví dụ về `docker run` với volume**:
```
docker run -it -v /host/path:/container/path image_ID
```
*Lệnh này sẽ mount thư mục `/host/path` từ host vào thư mục `/container/path` trong container.*

- **Ví dụ về `docker run` với `--volumes-from`**:
```
docker run -it --name new_container --volumes-from existing_container -image_ID
```
*Lệnh này sẽ tạo một container mới và chia sẻ volume từ container `existing_container`.*

- **Ví dụ về `docker volume create`**:
```
docker volume create my_volume
```
*Lệnh này sẽ tạo một volume mới có tên `my_volume`.*