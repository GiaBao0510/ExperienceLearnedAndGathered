
- Docker file là một tệp tin văn bản ,bao gồm các câu lệnh để Docker có thể đọc các câu lệnh trên và chạy các câu lệnh đó cho ra một image mới theo nhu cầu.
- Chạy tệp tin Dockerfile .Ví dụ tên tệp tin là BackEndAspNet:
```
docker build -t <Tên image>:<Version> --force-rm -f BackEndAspNet
```

**_Quy Tắc viết câu chỉ thị:_**

```
<Tên chỉ thị> <Tham số>
```

> Ví dụ: RUN yum update -y

- _Do quá trình Docker build image mới từ Dockerfile có thể tạo ra các image tạm, gây rác cho hệ thống.Đây là lệnh xóa Image tạm:_
```
docker image prune
```

- _Tạo và chạy container từ image mới:_
```
docker run -it --name <tên container> -p<Cổng Host>:<Cổng Container> -h <> <>
```

#####  _**Các chỉ thị trong docker:**_

- _FROM:_ mọi Docker file điều có chỉ thị này, chỉnh định image cơ sở
>Chỉ thị này chỉ ra image cở sở để xây dựng nên image mới .Để xây dựng từ imagr nào đó thì cần đọc document tuef image đó để biết rong đó đang chứa gì, có thể chạy lệnh gì trong đó

- *_COPY  & ADD_*: sao chép dữ liệu
>Được dùng để them thư mục file vào image
```
ADD thư_mục_nguồn thư_mục_đích
```

trong đó:
	- Thư mục ngồn là thư mục của máy chạy Dockerfile
	- Thư mục đính là nơi thêm dữ liệu vào container

- _ENV:_ thiết lập biến môi trường
> Chỉ thị này dùng để thiết lập biến môi trường như biến môi trường PATH,... tùy hệ thống, ứng dụng yêu câu biến môi trường nào thì căn cứ vào đó để thiết lập
```
ENV biến giá trị
```

- _RUN:_ chạy các lệnh
> Thi hành các lệnh, tương tự như lệnh shell trên OS từ terminal
```
RUN lệnh cần chạy
```

- _VOLUME: gắn ổ đĩa, thư mục
> Chỉ thị tạo một ổ đĩa chia sẻ được giữa các container

- _User:_ user
> Thêm user được dùng khi chạy các chỉ thị RUN CMD WORKDIR

- _WORKDIR:_ thư mục làm việc
> Thiết lập thư mục làm việc hiện tại chỉ các chỉ thị CMD, ENTRYPOINT, ADD thi hành
```
WOEKDIR path_curent_dir
```

- _EXPOSE:_ thiết lập cổng
> Thiết lập cổng mà container lắng nghe, cho phép container khác trên cùng mạng liên lạc qua cổng này hoặc đi ánh xạ cổng host vào cổng này.

- ENTRYPOINT, CMD trong Docker:
> chạy lệnh này khi container được chạy
```
ENTRYPOINT command_script
ENTRYPOINT ['command', 'tham số',...]
```
CMD ý nghĩa tương tự như ENTRYPOINT, khác là lệnh chạy bằng shell của container