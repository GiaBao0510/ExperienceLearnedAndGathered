
Khi làm việc với Entity Framework (EF) trong **ASP.NET Core**, có hai cách để tiếp cận chính để thiết kế cơ sở dữ liệu: **Database First** và **Code First**. Mỗi cách tiếp cận điều có ưu và nhược điểm riêng.

---
## **1️⃣ Database First là gì?**

**Database First (DB First)** đầu tiên là sẽ bắt đầu với ==cơ sở dữ liệu có sẳn== trước. Sau đó **Entity Framework** sẽ ==tự động tạo ra các lớp C# (model) tương ứng với các bảng trong csdl==.
### 🔹 **Cách hoạt động:**

1. Thiết kế database trước.
2. Sử dụng **EF Scaffold-DbContext** để tạo ra model từ database.
3. Viết code để tương tác với DB thông qua **EF**.

***Ví dụ tạo model từ DB (EF Core - DB First):***

```powershell
dotnet ef dbcontext scaffold "Server=.;Database=MyDb;Trusted_Connection=True;" Microsoft.EntityFrameworkCore.SqlServer -o Models
```

---
## **2️⃣ Code First là gì?**

**CodeFirst** là cách tiếp cận ==bắt đầu với code trước==, tức là định nghĩa ra các model trước, Sau đó **EF** sẽ ==tự động tạo database== và ==các bảng tương ứng dựa trên model.==

### 🔹 **Cách hoạt động:**

1. Định nghĩa các class C# đại diện cho bảng.
2. Cấu hình `DbContext`.
3. Dùng Migration để tự động tạo hoặc cập nhật DB.

***Ví dụ database từ model (EF Core - CodeFirst):***

```csharp
public class Product
{
	[Key]
	public int Id {set; get;}

	[Required]
	public string Name {set; get;}

	[Required]
	public decimal Price {set; get;}
}
```
Tại ví dụ này:
- class `Product`  sẽ ánh xạ tới một bảng có tên mặc định là `Product`.
- Thuộc tính `Id` sẽ là khóa chính của bảng (EF Core nhận diện theo quy ước đặt tên `Id` hoặc `[TênEntity]Id`). Attribute `[Key]` giúp làm rõ điều này.
- Thuộc tính `Name` sẽ là một cột kiểu chuỗi trong bảng. Attribute `[Required]` sẽ làm cho cột này không được phép NULL ở cấp độ database.
- Thuộc tính `Price` sẽ là một cột kiểu decimal.

#### **Tạo DbContext - Cửa ngõ tới Database:**

`DbContext` là class quan trong trong EF Core. Nó ==đại diện cho một phiên làm việc với database== và là ==nơi cấu hình model, thực hiện việc truy vấn và lưu thay đổi==. Cần tạo một lớp kế thừ từ thư viện `Microsoft.EntityFrameworkCore.DbContext`.

```csharp
public class MyDbContext: DbContext{

	//Một tập hợp các product cần quản lý và nó ánh xạ đến database
	public DbSet<Product> Product{get; set;}

	// Phương thực này được dùng để cấu hinh DbContext, bao gồm chuỗi kết nối
	protected override void OnConfiguring(DbContextOptionsBuilder opt){
		//Lưu ý chuỗi kết nối nên đặt trong file appsetting.json hoặc trong biến môi trường
		opt.UseSqlServicer("Server=.;Database=MyDb;Trusted_Connection=True;");
	}
}
```

#### **Áp dụng Migrations: Tạo Database Schema**

Sau khi đã định nghĩa `model` và `DbContext` xong thì cần thông báo cho **EF Core** biết về `model` này và yêu cầu nó tạo ra script để tạo **schema database** tương ứng. Quá trình này được thực hiện thông qua **Migrations**.

**Migrations** giúp quản lý các thay đổi về **schema database** theo thời gian. Mỗi khi thay đổi model C# ( thêm/xóa class, thêm/xóa thuộc tính, thay đổi kiểu dữ liệu, thêm quan hệ,...), thì sẽ tạo ra **migration** mới, **Migration** này chứa đựng các đoạn mã C# mô tả thao tác database cần thiết để chuyển đổi từ schema cũ sang schema mới.

Sử dụng công cụ dòng lệnh *"dotnet ef"* để cài đặt và chạy tại thư mục gốc:
```bash
dotnet ef migrations add InitialCreate
```
- **dotnet ef migrations**: yêu cầu EF Core tạo một migration mới.
- **InitialCreate**: tên của Migration (nên đặt tên có nghĩa).

Sau khi chạy lệnh này, EF Core sẽ:
1. Scan project để tìm DbContext.
2. Phân tích model được định nghĩa trong **DbContext** (class `product`).
3. So sánh với model hiện tại với trạng thái cuối cùng được ghi nhận trong folder **Migrations** (Nếu chưa có **migration** nào, nó sẽ so sánh với trạng thái database trống).
4. Tạo một folder `Migrations` trong project.
5. Tạo **2 file .cs** bên trong folder `Migrations`: một file chứa các thao tác `Up` (áp dụng migration) và `Down` (rollback migration), và một file snapshot của model hiện tại.

Kiểm tra folder `Migrations`,sẽ thấy các file được tạo ra. File migration (ví dụ: `20231027100000_InitialCreate.cs` – phần số là timestamp)

```csharp
using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace EFCoreGettingStarted.Migrations
{
    /// <inheritdoc />
    public partial class InitialCreate : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.CreateTable(
                name: "Products",
                columns: table => new
                {
                    Id = table.Column<int>(type: "int", nullable: false)
                        .Annotation("SqlServer:Identity", "1, 1"), // Cột tự tăng
                    Name = table.Column<string>(type: "nvarchar(max)", nullable: false), // NOT NULL do [Required]
                    Price = table.Column<decimal>(type: "decimal(18,2)", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_Products", x => x.Id); // Khóa chính
                });
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropTable(
                name: "Products");
        }
    }
}
```
File này mô tả việc tạo bảng “Products” trong phương thức `Up` và xóa bảng đó trong phương thức `Down`. EF Core đã dịch các thuộc tính trong class `Product` thành các cột với kiểu dữ liệu và ràng buộc tương ứng.

Bước kế tiếp là áp dụng migration này vào database. Qua câu lệnh sau:
```bash
dotnet ef database update
```

Lệnh này sẽ:
1. Tìm **DbContext**.
2. Đọc chuỗi kết nối từ **DbContext**.
3. Kết nối tới Database (nếu database chưa tồn tại, EF Core sẽ cố gắng tạo mới nó),
4. Kiểm tra bảng lịch sử migrations trong database (`__EFMigrationsHistory`).
5. Thực thi các **migration** chưa được áp dụng. Trong trường hợp này, nó sẽ chạy phương thức `Up` của migration `InitialCreate`, tạo bảng `Products`.

---
## **3️⃣ So sánh chi tiết giữa Database First và Code First**

| Tiêu chí                     | Database First                                      | Code First                                              |
| ---------------------------- | --------------------------------------------------- | ------------------------------------------------------- |
| **Khởi điểm**                | Thiết kế databse trước rồi tạo mode sau             | Tạo model trước rồi sinh database sau.                  |
| **Tạo database**             | Đã có DB sẵn, EF chỉ kết nối vào                    | EF sẽ tự tạo ra DB dựa trên model.                      |
| **Công cụ hỗ trợ**           | EDMX, wizard trong Visual Studio.                   | Migration, Fluent API, Data Annotations.                |
| Kiểm soát cấu trúc           | Cơ sở dữ liệu quyết định                            | Code quyết địn                                          |
| **Dễ bảo trì**               | Khó hơn khi CSDL phức tạp.                          | Dễ hơn nhờ vào Migration                                |
| **Linh hoạt**                | Dễ sửa đổi nếu đã có DB phức tạp (Ít linh hoạt hơn) | Dễ sử đổi nếu bắt đầu từ đầu (Rất linh hoạt)            |
| **Quản lý Schema**           | Phù hợp nếu cần giữ nguyên cấu trúc DB cũ           | Có thể dễ dàng thay đổi schema với migration            |
| **Tích hợp với hệ thống cũ** | Rất phù hợp nếu có database cũ                      | Khó tích hợp nếu đã có database cũ.                     |
| **Phù hợp với team nào?**    | DBA (Database Administration) quản lý tốt hơn       | Dev (Developer) dễ làm việc hơn                         |
| **Khó khắn khi phát triển**  | Mất công update code nếu schemal thay đổi.          | Mất công update databse nếu model thay đổi.             |
| **Hỗ trợ Migration**         | Không hỗ trợ tốt, cần chỉnh sửa thủ công            | Hỗ trợ tốt với `Migration`                              |
| **Hiệu suất**                | Tối ưu hơn nếu DB được thiết kế chuẩn               | Có thể tạo nhiều bảng không cần thiết nếu không tối ưu, |


---
## **4️⃣ Khi nào nên dùng cái nào ? **

### **✅Dùng Database First khi:**

✔ Đã có **database sẳn** hoặc cần kết nối với hệ thống cũ.
✔ **Team có DBA** chuyên thiết kế và quản lý database.
✔ **Database phức tạp**, có stored procedure, trigger, view, index,....
✔ Cần làm việc với **hệ thống lớn**, nơi mà dữ liệu là điều quan trong và cần kiểm soát tốt.

🔹 ***Ví dụ:***
- Một hệ thống ERP đã có database sẵn, chỉ cần tạo ứng dụng ASP.NET Core kết nối vào để xử lý dữ liệu.
- Một hệ thống thương mại điện tử lớn đã có nhiều bảng, quan hệ phức tạp.

### **✅Dùng Code First khi:**

✔ Dự án mới, chưa có database, muốn linh hoạt thay đổi schema.
✔ Team không có DBA chuyên biệt, lập trình viên tự quản lý database.
✔ Muốn tận dụng Migration để dễ dàng cập nhật database từ code.
✔ Muốn **đơn giản hóa quá trình phát triển** mà không cần tạo CSDL bằng tay.

🔹 ***Ví dụ:***
- Một startup mới phát triển một ứng dụng quản lý kho hàng, chưa có database.
- Một dự án internal tool nhỏ, chỉ cần CRUD đơn giản, không yêu cầu tối ưu database quá mức.

---
## **5️⃣ Ưu điểm và Nhược điểm?**

#### ***Database First***

- **Ưu điểm**:
    1. **Phù hợp với cơ sở dữ liệu có sẵn**: Nếu dự án kế thừa cơ sở dữ liệu cũ hoặc làm việc với DBA (Database Administrator) đã thiết kế sẵn.
    2. **Nhanh chóng khởi đầu**: Công cụ sinh mã tự động giúp tiết kiệm thời gian khi cơ sở dữ liệu lớn.
    3. **Dễ hình dung**: Các nhà phát triển quen với SQL và thiết kế cơ sở dữ liệu sẽ thấy dễ tiếp cận.
- **Nhược điểm**:
    1. **Ít kiểm soát code**: Code sinh tự động có thể rối, khó tùy chỉnh.
    2. **Khó bảo trì**: Khi cơ sở dữ liệu thay đổi, bạn phải cập nhật EDMX thủ công, dễ gây lỗi.
    3. **Không linh hoạt**: Khó áp dụng các thay đổi logic trong code mà không sửa cơ sở dữ liệu trước.
    4. **Phụ thuộc công cụ**: Dựa vào EDMX, có thể không tối ưu với EF Core (ít hỗ trợ EDMX).

#### ***Code First***

- **Ưu điểm**:
    1. **Kiểm soát hoàn toàn**: Bạn định nghĩa mô hình trong code, dễ tùy chỉnh (dùng Data Annotations hoặc Fluent API).
    2. **Migration mạnh mẽ**: Dễ dàng thay đổi cấu trúc cơ sở dữ liệu qua migration (Add-Migration, Update-Database).
    3. **Linh hoạt**: Hỗ trợ cả tạo mới cơ sở dữ liệu và ánh xạ tới cơ sở dữ liệu cũ.
    4. **Tích hợp DevOps**: Migration dễ đưa vào pipeline CI/CD, quản lý phiên bản cơ sở dữ liệu.
    5. **Hỗ trợ EF Core tốt hơn**: EF Core khuyến khích Code First, không còn EDMX.
- **Nhược điểm**:
    1. **Khó với cơ sở dữ liệu phức tạp có sẵn**: Nếu cơ sở dữ liệu cũ lớn và phức tạp, việc ánh xạ thủ công mất thời gian.
    2. **Yêu cầu hiểu biết**: Cần nắm rõ cách cấu hình migration và Fluent API.
    3. **Khởi đầu chậm hơn**: Viết code và migration từ đầu tốn công sức ban đầu.

---
## **6️⃣ Vậy cái nào tốt hơn?**

❌ **Không có cái nào tốt hơn hoàn toàn**, **tùy vào dự án** mà chọn cái phù hợp nhất!

📌 **Nếu đã có database → Chọn Database First**.  
📌 **Nếu chưa có database, muốn linh hoạt → Chọn Code First**.  
📌 **Nếu hệ thống lớn, nhiều logic SQL → Chọn Database First**.  
📌 **Nếu muốn dễ dàng mở rộng, thay đổi nhanh → Chọn Code First**.