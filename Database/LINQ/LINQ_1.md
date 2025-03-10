# **Phần 1: Tổng quan về LINQ**


![](https://s3-sgn09.fptcloud.com/codelearnstorage/Upload/Blog/linq-la-gi-63731082354.1023.jpg)
## 1.**LINQ là gì?**

**LINQ (Language Integrated Query)** là một công cụ mạnh mẽ trong c# cho **phép truy vấn và thao tác dữ liệu** một cách dễ dàng và trực quan, giống như SQL nhưng ngay cả trong mã c#.
Nó giúp tiết kiệm thời gian truy vấn dữ liệu ngắn gọn, dễ đọc, dễ bảo trì hơn so với cách truyền thống (dùng vòng lặp, điều kiện if ...).
**🔹LINQ hoạt động trên nhiều nguồn dữ liệu khác nhau,** bao gồm:
✅ **Collection (List, Array, Dictionary,...) →** LINQ to Objects.
✅ **Cơ sở dữ liệu quan hệ (SQL Server, MySQL, PostgreSQL, SQLite,...) →** LINQ to Entities (Entities Framwork).
✅ **Dữ liệu XML** → LINQ to XML.
**✅ Dữ liệu NoSQL(MongoDB, Firebase,...) →** LINQ to NoSQL (thông qua drive hỗ trợ).
**✅ Dữ luệ từ dịch vụ web.**

---
## 2.**Lợi ích của LINQ**

🔹 **Cú pháp dễ đọc, dễ hiểu**
- So với cách truy vấn dữ liệu truyền thống (dùng vòng lặp và điều kiện if), LINQ giúp code gọn hơn.

🔹 **Đồng nhất khi làm việc với nhiều loại dữ liệu**
- Dù bạn truy vấn dữ liệu từ Collection, Database, hay XML, cú pháp LINQ vẫn giống nhau.

🔹 **Tăng hiệu suất và tối ưu hóa**
- LINQ hỗ trợ **Deferred Execution (Thực thi trì hoãn)** giúp tăng hiệu suất bằng cách chỉ thực hiện truy vấn khi cần thiết.

🔹 **Dễ bảo trì**
- Code LINQ rõ ràng hơn vòng lặp `for`, `foreach`, giúp bảo trì dễ dàng hơn.

---
## **LINQ làm việc với các loại cơ sở dữ liệu nào?**

Có thể làm việc với hầu hết các loại cơ sở dữ liệu, nhưng cách thực hiện có thể khác nhau:

🔹**SQL Server, MySQL, PostgreSQL, SQLite**
- sử dụng **Entity Framework Core (LINQ to Entities).**
- khi viết câu truy vấn, **EF Core** sẽ chuyển LINQ thành SQL tương ứng.

🔹 **MongoDB**
- Dùng thư viện `MongoDB.Drive` và `LINQ` để truy vấn dữ liệu.
- LINQ không được chuyển đổi thành SQL mà thành truy vấn MongoDB.

🔹 **Firebase (Realtime Database, Firestore)**
- LINQ không hỗ trợ trực tiếp, nhưng có thể dùng với `.AsQueryable()`

**🔹DynamoDB, CosmosDB:**
- Một số NoSQL DB có hỗ trợ LINQ thông qua SDK riêng.

---
## **Cách thiết lập dự án để sử dụng LINQ với Database:**

Nếu việc dùng LINQ với Collection(`List`,  `Array`) không cần thêm gì cả vì LINQ có sẳn trong **.NET Core**.
Nếu muốn áp dụng LINQ với Database, thì cần thêm package phù hợp sau.

### **🔹Cài đặt LINQ với SQL Server (Entity Framework Core):**
**Bước 1:** Vào dự án và mở terminal để tải **package Entity Framework Core:**
```
dotnet add package Microsoft.EntityFrameworkCore.SqlServer
dotnet add package Microsoft.EntityFrameworkCore.Tools
```

**Bước 2:** Cấu hình `DBContext` trong `program.cs`:
```
using Microsoft.EntityFrameworkCore;

var builder = WebApplication.CreateBuilder(args);
builder.Services.AddDbContext<AppDbContext>(options =>
    options.UseSqlServer("Server=localhost;Database=TestDB;User Id=sa;Password=yourpassword;"));

var app = builder.Build();
app.Run();
```

---
## **Tổng kết:**