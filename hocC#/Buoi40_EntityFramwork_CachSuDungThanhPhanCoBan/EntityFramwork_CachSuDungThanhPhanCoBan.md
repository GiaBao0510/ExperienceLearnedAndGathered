## **Entity Framework Core**
- là Framework (thư viện khung) để ánh xạ các đơn vị dữ liệu mô tả bằng lớp (đối tượng) vào trong CSDL quan hệ, nó cho phép ánh xạ vào các bảng csdl, tạo csdl, truy vấn với LINQ tạo và cập  nhật vào database.
- Thêm các gói cần thiết:
```
dotnet add package System.Data.SqlClient
dotnet add package Microsoft.EntityFrameworkCore
dotnet add package Microsoft.EntityFrameworkCore.SqlServer
dotnet add package Microsoft.EntityFrameworkCore.Design
dotnet add package Microsoft.Extensions.DependencyInjection
dotnet add package Microsoft.Extensions.Logging
dotnet add package Microsoft.Extensions.Logging.Console
dotnet add package Microsoft.EntityFrameworkCore.Tools
dotnet add package Pomelo.EntityFrameworkCore.MySql
```
- Các namespace có thể dùng:
```
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.ChangeTracking;
using Microsoft.EntityFrameworkCore.Diagnostics;
using Microsoft.EntityFrameworkCore.Infrastructure;
using Microsoft.EntityFrameworkCore.Metadata.Builders;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Logging;
using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using System.Linq;
using System.Threading.Tasks;
```
- ##### Tạo Model đơn giản, ánh xạ bảng CSDL
- Model là mô hình hóa các đối tượng dữ liệu trong hệ quản trị CSDL thành các đối tượng lập trình, đó là các lớp (class) tương ứng với các bảng ...
- Trước khi sử dụng model trong EF Core, hãy bổ sung các thiết lập thông qua các Attribute

---
### **Tạo Context - DbContext**
- **DbContext** trong EF là ngữ cảnh làm việc, nó biểu diễn, chứa các thông tin cần thiết của một phiên làm việc với CSDL.
- Bất kỳ khi nào đối tượng DbContext được tạo mới, thì nó sẽ thi hành phương thức override
- Chúng ta thường làm tròn phương thức này để cấu hình kết nối đến cơ sở dữ liệu
- Trong đó cần nạp chồng `OnConfiguring` để cấu hình (thiết lập chuỗi kết nối ...), và tạo ra thuộc tính có kiểu `DbSet<Product>` chính là bảng trong CSDL
- Để đọc tên cơ sở dữ liệu từ đối tượng kế thừa từ lớp DbContext như sau: 
```
string DBname = dbContext.Database.GetDbConnection().Database;
```
- Phương thức **EnsureCreated()** trong đối tượng kế thừ từ **DBContext** dùng để tạo ra cơ sở dữ liệu nếu csdl này chưa có, ngược lại thì không tạo.

### **DbSet:**
- Kiểu thuộc tính là DbSet là kiểu dữ liệu bảng của csdl mỗi dòng của bảng csdl của bảng nào đó biểu diễn cho một đối tượng lớp nào đó

**- Add, AddAsync**: Chèn dữ liệu vào bảng
**- Add, AddAsync**: Chèn dữ liệu vào bảng
**- Add, AddAsync**: Chèn dữ liệu vào bảng
**- SaveChanges**: Cập nhật câu lệnh sql server

## **Logging**:

- Có thể truy cập đến csdl bằng DBContext
- Không cần làm việc trực tiếp đến các câu lệnh sql, mà thư viện này tự động sinh ra câu lệnh truy vấn sql tương ứng.
- Cần đảm bảo thêm các gói sau: 
```
dotnet add package Microsoft.Extensions.DependencyInjection
dotnet add package Microsoft.Extensions.Logging
dotnet add package Microsoft.Extensions.Logging.Console
```
- Phương thức **AddFilter** dùng để hiển thị thông tin gốc