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

 **Gắn ổ đĩa volume vào container:**
```
docker run -it --mount source=DISK,target=pathContainer ImageID
```

> Dữ liệu sẽ không mất đi khi xóa container

 **Tạo ổ đĩa ánh xạ đến thư mục host**
```
docker volume create --opt device=pathHOST --opt type=none --otp o=bind DISKNAME
```
