_Kiểm tra các Images:_
```shell
docker images -a
```

_Tải Image nào đó:_
```shell
docker pull [name image]:[tag]
```

_Kiểm tra các container đang chạy:_
```shell
docker ps
```

_Liệt kê tất cả các container:_
```shell
docker container ls --all
```

_Xóa image theo ID:_
```shell
docker rmi [image id]
```

_Xóa image theo tên và tag:_
```shell
docker rmi [imagename:tag]
```

_Tạo và chạy container:_
```shell
docker run [options] image [command] [Arg...]
```
- ``[options]: thiết lập trước khi tạo container
- ``[command]: cho biết các lệnh, sau khi chạy container
- ``[Arg...]: các tham số khi chạy container

_Chạy 1 container đã dừng dựa trên ID_
```shell
docker container start -i [container ID]
```

_Vào terminal của container đang chạy:_
```shell
docker container attach [container ID]
```

_Chạy một câu lệnh trên container đang chạy:_
```shell
docker exec -it [container ID] command
```

Dừng container nào đó khi đang chạy dựa trên Container ID
```shell
docker stop [container ID]
```

Xóa container đang chạy dựa trên tên container
```shell
docker rm [tên container]
```

**Xóa tất cả các container đang chạy**:
```shell
docker rm -f ${docker ps -aq}
```

**Xóa tất cả các images**:
```shell
docker rmi -f ${docker images -aq}
```

**Xóa các container đã dừng**:
```shell
docker container prune
```

**Xóa các images không được sử dụng**:
```shell
docker image prune
```

**Xem logs của container**:
```shell
docker logs [container ID]
```

**Xem thông tin chi tiết về container**:
```shell
docker inspect [Container ID]
```

**Tạo một network mới**:
```shell
docker network create [network name]
```

**Kết nối container vào một network**:
```shell
docker network connect [network name] [container ID]
```

**Xem danh sách các network**:
```shell
docker network ls
```

---
## **Cập nhật Image, lưu image ra file và nạp image từ file vào docker**

_Lưu container thành một image:_
```shell
docker commit [tên tự đặt] [myimage]:[version]
```

_Lưu ra file tại nơi đang đứng trên thư mục:_
```shell
docker save --output [MyImage].tar myimage
```

_Nạp file vào docker để tạo ra image:_
```shell
docker load -i MyImage.tar
```

_Đổi tên image cụ thể:_
```shell
docker tag [image id] [image name]:[version]
```

_Lấy địa chỉ IPv4 của một container đang chạy cụ thể:_
```shell
docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' <container_name/container_id>
```
