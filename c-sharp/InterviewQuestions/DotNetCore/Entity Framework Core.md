## **1. Entity Framework Core là gì?**

Entity Framework Core (EF Core) là phiên bản mới, hiện đại và nhẹ hơn của **Entity Framework (EF)**, được phát triển bởi Microsoft. EF Core là một **ORM (Object-Relational Mapping)** hỗ trợ các nhà phát triển .NET làm việc với cơ sở dữ liệu bằng cách sử dụng đối tượng lớp thay vì SQL thuần.

EF Core được thiết kế để hỗ trợ:
- **Đa nền tảng:** window, macOS, linux.
- **Đa hệ quản trị CSDL:** mySQL, SQL server, SQLite, v.v.
- Các ứng dụng .NET hiện tại như ASP.NET Core.

_Minh họa:_
![EFCore](https://miro.medium.com/v2/resize:fit:1400/0*N9Zrc7IA5Zf1Jt8s.jpg)

![EFCore](https://images.viblo.asia/d724d0e6-2a1d-4d8c-908a-5151a135bcaf.png)
## **2. Đặc điểm của Entity Framework Core là gì?**

Sau đây sẽ nêu các đặc điểm về  EF Core đã có mà EF chưa có.

**Nhẹ và hiệu năng tốt hơn:**
- ít tài nguyên hơn so với EF truyền thống.
- Tốc độ nhanh hơn(performance).
**Đa nền tảng:**
- Có thể chạy trên **.NET Core** (hỗ trợ Windows, Linux, macOS) và **.NET Framework**.
**Hỗ trợ cả Code-First, Database-First, Model-First**.
**Open source.**
**Khả năng mở rộng:**
- Có thể tùy chỉnh và mở rộng dễ dàng thông qua các tính năng như ==Interceptors, Custom Conventions==.
**Hỗ trợ cơ chế Loading:**
- Eager Loading, Lazy Loading và Explicit Loading.
## **3. DBContext là gì?**

DbContext là một lớp trung tâm trong Entity Framework Core, được sử dụng để:
- **Quản lý kết nối CSDL.**
- **Theo dõi các thực thể (Entities) và trạng thái của chúng.**
- **Thực hiện các thao tác CRUD (Create, Read, Update, Delete).**
- **Tương tác với CSDL bằng LINQ.**

_Ví dụ:_
```csharp
public class MyAppContext : Dbcontext{
	public DbSet<User> Users {get; set;}
	public DbSet<Product> Products {get; set;}
	protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder){
		optionsBuilder.UseSqlServer("ConnectString");
	}
}
```

## **4. DBSET là gì?**

**==DbSet==** đại diện cho một thực thể trong CSDL.
- Mỗi ==DbSet< T >== ánh xạ đến một bảng (table) trong cơ sở dữ liệu, và ==T== là lớp thực thể trong (entity).
_Ví dụ:_
```
public DbSet<User> Users {get; set;}
```
## **5. So sánh entity framework core với entity framework?**

|         Tiêu chí         | EntityFramework                                         | EntityFramework Core                                    |
| :----------------------: | :------------------------------------------------------ | :------------------------------------------------------ |
|       **Nền tảng**       | Chỉ hỗ trợ .NET Framework                               | Hỗ trợ .NET Core, .NET 5+                               |
|     **Đa nền tảng**      | Không                                                   | Window, Linux, macOS                                    |
|      **Hiệu năng**       | Chậm hơn do sự phức tạp và cồng kềnh                    | Tốt hơn                                                 |
| **Hỗ trợ cơ sở dữ liệu** | Hỗ trợ nhiều hệ quản trị (MySQL, SQLite, PostgreSQL...) | Hỗ trợ nhiều hệ quản trị (MySQL, SQLite, PostgreSQL...) |
|  **Tính năng hiện đại**  | Hạn chế                                                 | Có thên nhiều tính năng mới                             |
|      **Tương lai**       | Không còn được phát triển                               | Được phát triển tích cực                                |

## **6. Phân biệt code first và database first trong EF core?**

|     Tiêu chí      | Code-First                     | Database-First                         |
| :---------------: | ------------------------------ | -------------------------------------- |
| **Cách làm việc** | Tạo CSDL từ mã nguồn           | Tạo mã nguồn từ CSDL sẳn có            |
|   **Dùng khi**    | Dùng khi chưa có CSDL.         | Dùng khi đã có CSDL sẳn                |
|    **Ưu điểm**    | Kiểm soát mã nguồn dễ dàng.    | Nhanh chóng tích hợp CSDL cũ.          |
|  **Nhược điểm**   | Khó làm việc với CSDL phức tạp | Mã nguồn sinh ra có thể khó điều chỉnh |
_Hình minh họa:_
![DB-First vs Code-First](https://images.viblo.asia/5890f299-da38-485b-afc4-fea83e2f7caf.png)
## **7. Phân biệt eager loading, lazy loading và explicit loading?**

| Cách tải dữ liệu  | Eager Loading                              | Lazy Loading                                    | Explicit Loading                                     |
| :---------------: | ------------------------------------------ | ----------------------------------------------- | ---------------------------------------------------- |
|     Khái niệm     | Tải toàn bộ dữ liệu liên quan ngay lập tức | Chỉ tải dữ liệu liên quan khi cần thiết         | Chỉ tải dữ liệu liên quan khi lập trình viên yêu cầu |
|  Cách thực hiện   | Sử dụng ==Include()== trong truy vấn       | Tự động, không cần chỉ định                     | Sử dụng ==Load()== hoặc truy vấn cụ thể              |
|     HIệu năng     | Có thể gây tải nặng nếu dữ liệu nhiều      | Tốt hơn cho dữ liệu lớn nhưng cần dùng truy vấn | Kiểm soát tối ưu hóa nhưng yêu cầu lập trình thêm    |
| **Khi nào dùng**? | Khi cần lấy tất cả dữ liệu ngay lập tức    | Khi dữ liệu phụ không phải lúc nào cần          | Khi kiểm soát cách và lúc tải dữ liệu.               |
_Ví dụ:_
- **Eager Loading:**
```
var order = context.Order.Include(o => o.customer).ToList();
```

- **Lazy Loading:**
```
public class Order{
	public virtual Customer Customer{get; set;}
}
```

- **Explicit Loaing:**
```
var order = context.Order.Find(orderID);
context.Entry(order).Reference(o => o.Customer).Load();
```

## **8. Ánh xạ kiểu dữ liệu từ c# sang Sql server và ngược lại?**

![](https://images.viblo.asia/24f6b1db-3831-4aaa-adb9-2eb4057fc4e9.png)

## **7. Phân biệt eager loading, lazy loading và explicit loading?**

## **7. Phân biệt eager loading, lazy loading và explicit loading?**