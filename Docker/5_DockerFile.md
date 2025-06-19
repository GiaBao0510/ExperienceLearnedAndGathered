![](https://miro.medium.com/v2/resize:fit:600/format:webp/0*nX1z5vaygpdwyukK.jpeg)

- Docker file là một tệp tin văn bản ,bao gồm các câu lệnh để Docker có thể đọc các câu lệnh trên và chạy các câu lệnh đó cho ra một image mới theo nhu cầu.
- Chạy tệp tin Dockerfile .Ví dụ tên tệp tin là BackEndAspNet:
```shell
docker build -t <Tên image>:<Version> --force-rm -f BackEndAspNet
```

**_Quy Tắc viết câu chỉ thị:_**

```shell
<Tên chỉ thị> <Tham số>
```

> Ví dụ: `RUN yum update -y`

- _Do quá trình Docker build image mới từ Dockerfile có thể tạo ra các image tạm, gây rác cho hệ thống.Đây là lệnh xóa Image tạm:_
```shell
docker image prune
```

- _Tạo và chạy container từ image mới:_
```shell
docker run -it --name <tên container> -p<Cổng Host>:<Cổng Container> -h <> <>
```


---
#####  _**Các chỉ thị trong docker:**_

![](https://i.geekflare.com/cdn-cgi/image/width=697,height=270,fit=crop,quality=90,format=auto,onerror=redirect,metadata=none/wp-content/uploads/2019/07/dockerfile-697x270.png)

- _FROM:_ mọi Docker file điều có chỉ thị này, chỉnh định image cơ sở
>Chỉ thị này chỉ ra image cở sở để xây dựng nên image mới .Để xây dựng từ image nào đó thì cần đọc document từ image đó để biết trong đó đang chứa gì, có thể chạy lệnh gì trong đó

- *_COPY  & ADD_*: sao chép dữ liệu
>Được dùng để thêm thư mục file vào image
```dockerfile
ADD thư_mục_nguồn thư_mục_đích
```

trong đó:
	- Thư mục ngồn là thư mục của máy chạy Dockerfile
	- Thư mục đính là nơi thêm dữ liệu vào container

- _ENV:_ thiết lập biến môi trường
> Chỉ thị này dùng để thiết lập biến môi trường như biến môi trường PATH,... tùy hệ thống, ứng dụng yêu câu biến môi trường nào thì căn cứ vào đó để thiết lập
```dockerfile
ENV biến giá trị
```

- **_RUN**:_ chạy các lệnh
> Thi hành các lệnh, tương tự như lệnh shell trên OS từ terminal
```dockerfile
RUN lệnh cần chạy
```

- **_VOLUME**: gắn ổ đĩa, thư mục
> Chỉ thị tạo một ổ đĩa chia sẻ được giữa các container

- **_User**:_ user
> Thêm user được dùng khi chạy các chỉ thị RUN CMD WORKDIR

- **_WORKDIR**:_ thư mục làm việc
> Thiết lập thư mục làm việc hiện tại chỉ các chỉ thị CMD, ENTRYPOINT, ADD thi hành
```dockerfile
WOEKDIR path_curent_dir
```

- **_EXPOSE**:_ thiết lập cổng
> Thiết lập cổng mà container lắng nghe, cho phép container khác trên cùng mạng liên lạc qua cổng này hoặc đi ánh xạ cổng host vào cổng này.

- **ENTRYPOINT**, CMD trong Docker:
> chạy lệnh này khi container được chạy
```dockerfile
ENTRYPOINT command_script
ENTRYPOINT ['command', 'tham số',...]
```
CMD ý nghĩa tương tự như ENTRYPOINT, khác là lệnh chạy bằng shell của container

- **LABEL**: Thêm metadata vào image.
```dockerfile
LABEL maintainer="your-email@example.com"
LABLE version="1.0"
LABEL description="This is a sample Docker image"
```

- **ARG**: Định nghĩa các biến build-time.
```dockerfile
ARG APP_VERSION=1.0
ENV APP_VERSION=${APP_VERSION}
```

- **HEALTHCHECK**: Kiểm tra sức khỏe của container.
```dockerfile
HEALTHCHECK --interval=5m --timeout=3s \ CMD curl -f http://localhost/ || exit 1
```

- **ONBUILD**: Thêm các chỉ thị sẽ được thực thi khi image được sử dụng làm base image cho một image khác.
```dockerfile
ONBUILD COPY . /app
ONBUILD RUN make /app
```


---
## **❓Ví dụ về tạo một image bằng cách đọc dockerfile từ dự án xây dựng hệ thống bầu cử trực tuyến phía BackEnd**
```dockerfile
# Bước 1: Sử dụng SDK .NET để build ứng dụng
FROM mcr.microsoft.com/dotnet/sdk:8.0 AS build
WORKDIR /app

# Sao chép tệp dự án và khôi phục gói nuget
COPY ["BackEnd.csproj", "./"]
RUN dotnet restore "BackEnd.csproj"

# Sao chép toàn bộ mã nguồn và build ứng dụng
COPY . ./
RUN dotnet publish "BackEnd.csproj" -c Release -o /app/out

# Bước 2: Sử dụng Runtime .NET để chạy ứng dụng
FROM mcr.microsoft.com/dotnet/aspnet:8.0
WORKDIR /app

# Sao chép kết quả từ giai đoạn build
COPY --from=build /app/out ./

# Mở cổng 80 nếu cần
EXPOSE 80

# Khởi chạy ứng dụng
ENTRYPOINT ["dotnet", "BackEnd.dll"]
```


---
## **🧠Thực tiễn tốt nhất:**

- **Sử dụng multi-stage builds**: Giảm kích thước image cuối cùng bằng cách chỉ sao chép các file cần thiết từ các giai đoạn build trước.
```dockerfile
FROM mcr.microsoft.com/dotnet/sdk:8.0 AS build
WORKDIR /app
COPY . .
RUN dotnet publish -c Release -o out

FROM mrc.microsoft.com/dotnet/aspnet:8.0
WORKDIR /app
COPY --from=build /app/out .
ENTRYPOINT ["dotnet", "myapp.dll"]
```

- **Sử dụng `.dockerignore`**: Bỏ qua các file và thư mục không cần thiết khi build image.
```dockerignore
node_modules
.git
*.log
```

---
## **Phần Troubleshooting cơ bản:**

- **Kiểm tra lỗi build**:
```shell
docker build -t my_image:1.0 -f Dockerfile
```

- **Xem logs của container để debug**:
```shell
docker logs <Container name/ container ID>
```

- **Kiểm tra các layer của image**:
```shell
docker history <Image ID>
```

