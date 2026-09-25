_NuGet là nơi phân phối các packets. Các thư viện được gọi là gói._
- Nếu trong dự án .Net có sử dụng các gói thì sẽ được cấu hình trong file có đuôi là **.csproj**
- Lệnh để thêm gói mới vào trong dự án: **dotnet add package <Tên gói muốn cài>**
- Lệnh để loại bỏ gói nào đó ra khỏi dự án thì chạy câu lệnh **dotnet remove package <tên gói đã cài>**
- Khi trên dự án có một số gói nào đó bị lỗi thì dùng câu lệnh: **dotnet restore**  để nó kiểm tra và phục hồi các package
- Để thêm thư viện vừa được tạo vào trong dự án hiện tại thì dùng cú pháp sau: **dotnet add <_Path/_tên file.csproj của dự án hiện tại> reference <_Path/_tên file.csproj của thư viện vừa mới tạo>**.
- Các bước để đưa thư viện lên NuGet:
	- 1. Đóng gói thư viện cần đóng bằng lệnh: **"dotnet pack"** .Sau khi thực hiện câu lệnh trên nó sẽ tạo ra 1 file có đuôi là **.mupkg**
	- 2. Đẩy lên bằng lệnh: **"dotnet nuget push <đường dẫn đến file .mupkg>"** --api-key <dán api đã tạo vào đây> --source [https://api.nuget.org/v3/index.json](https://api.nuget.org/v3/index.json)