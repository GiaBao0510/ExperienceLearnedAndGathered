- Liệt kê danh sách lịch sử các thao tác đã xây dựng lênh image:
```shell
docker history <Tên hoặc ID của image>
```

- Thông tin chi tiết về một image/ container
```shell
docker inspect <Image ID/ container ID>
```

- Lấy những thông tin bên trong container hiện tại khác với lúc mới tạo:
```shell
docker diff <Container name/ container ID>
```

- Kiểm tra nhật ký trên container:
```shell
docker logs <Container name/ container ID>
```

- Kiểm tra nhật ký theo số lượng dòng cuối cùng (Ví dụ lấy số lượng dòng cần kiểm tra là 15)
```shell
docker logs --tail 10 <Container name/ container ID>
```

- Lấy thông tin nhật ký khi containet nó đang chạy
```shell
docker logs -f <Container name/ container ID>
```

- Mức độ sử dụng của các container trên máy tính (Như: cpu, ram, tốc độ mạng,...)
```shell
docker stats <Container1 name> <Container2 name> ...
```

- Xem tất cả các container đang sử dụng tài nguyên trên máy
```shell
docker stats
```
