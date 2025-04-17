### **1. Entity Framework là gì?**
- **Entity Framework (EF)** là một công nghệ **ORM (Object-Relactional Mapping)** của Microsoft giúp cho các lập trình viên có thể thao tác với cơ sở dữ liệu một cách dễ dàng hơn. Thay vì phải viết các câu lệnh SQL thủ công thì có thể sử dụng đối tượng trong mã nguồn để tương tác với cơ sở dữ liệu.
- Entity Framework sử dụng mô hình ánh xạ (Mapping) để liên kết các thực thể trong ứng dụng với các bảng trong cơ sở dữ liệu giữa:
	- **Bảng trong cơ sở dữ liệu ↔ Lớp trong ứng dụng**
	- **Cột trong bảng ↔ Thuộc tính của lớp**
![EntityFramwork](https://www.tutorialspoint.com/entity_framework/images/conceptual_model.jpg)
### **2. Các thành phần trong Entity Framework?**
1. **Model:**
	- Biễu diễn các đối tượng ánh xạ đến cơ sở dữ liệu
	- Gồm các lớp (classes) và quan hệ giữa chúng thường, được gọi là **Entity Classes.**
	- Mang lại sự tách biệt giữa chương trình và CSDL với 3 thành phàn ==Conceptual Model, Mapping và Storage Model.== 
1. **DbContext:** 
	- Là một lớp quan trong để tương tác với cơ sở dữ liệu.
	- Chịu trách nhiệm về:
		- Kết nối cơ sở dữ liệu.
		- Quản lý các thực thể (Entity).
		- Thực thi các thao tác như ==Insert, Update, Delete, Query.==
2. **LINQ to Entities và Entity SQL:**
	- Đây là 2 ngôn ngữ truy vấn
	- Dùng để thực hiện truy vấn ngôn ngữ lập trình thay vì SQL thuần.
3. **Change Tracker:**
	- Theo dõi các thay đổi trên các thực thể (Entities) và đồng bộ chúng với cơ sở dữ liệu khi lưu.
4. **Database Provider:**
	- Thành phần làm việc với các hệ quản trị csdl khác nhau (SQL Server, MySQL, PostgreSQL, mariaDB, SQLite,...).
### **3. Đặc điểm của Entity Framework?**

1. **Hỗ trợ ORM (Object-Relational Mapping)**
	- ORM (Object Relation Mapping) là kỹ thuật ánh xạ CSLD sang các đối tượng trong các ngôn ngữ lập trình: java, c#, PHP,...
_Ví dụ:_
![ViDuORM](https://images.viblo.asia/37b9229f-f30f-4a1e-a1f9-929103bb0874.png)

1. **Truy vấn dữ liệu bằng LINQ:**
	- LINQ thay thế các câu lệnh SQL, giúp truy vấn dữ liệu dễ đọc và dễ bảo trì hơn
2. **Hỗ trợ Code-First, Database-First, Model-First:**
	- Code-First: Tạo cơ sở dữ liệu từ các lớp trong mã nguồn 
	- Database-First: Tạo lớp từ csdl có sẳn
	- Model-First: Tạo mô hình trong ứng dụng, từ đó tạo cơ sở dữ liệu.
3. **Quản lý trạng thái của thực thể:**
	- EF tự động theo dõi trạng thái của thực thể để thực hiện cập nhật CSDL chính xác
4. **Tương thích đa nền tảng (EF Core):**
	- Có thể chạy trên Window, Linux, macOS với **EntityFramwork Core.** 
### **4. Lợi ích khi sử dụng Entity Framework là gì?**
**Tăng năng suất:**
- Giảm lượng mã phải viết (ít SQL thủ công hơn).
- Tập trung vào logic ứng dụng thay vì quản lý CSDL.
**Dễ bảo trì:**
- Thay đổi mô hình hoặc CSDL dễ dàng, EF tự động xử lý ánh xạ.
**Truy vấn dễ đọc:**
- LINQ giúp truy vấn rõ ràng và ít lỗi hơn so với SQL thuần.
**Độc lập với hệ quản trị cơ sở dữ liệu:**
- Dễ dàng chuyển đổi giữa các hệ CSDL (SQL Server, MySQL, PostgreSQL, mariaDB, SQLite,...). bằng cách thay đổi **DB Provider.**
**Tích hợp tốt với .NET:**
- Là công cụ chính thức của Microsoft cho các ứng dụng .NET và ASP.NET Core. 
### **5. Hạn chế của Entity Framework là gì?**
**Hiệu năng thấp hơn so với SQL thuần:**
- Do EF thêm lớp trừu tượng nên hiệu năng có thể chậm hơn SQL thuần, đặc biệt với những câu truy vấn phức tạp.
**Hạn chế với truy vấn phức tạp:**
- LINQ có thể không đử mạnh hoặc rõ rảng để viết các truy vấn SQL rất phức tạp.
**Phụ thuộc nhiều vào DbContext:**
- Việc sử dụng không tối ưu DBContext có thể dẫn đến hiệu năng và bộ nhớ
**Không kiểm soát toàn diện:**
- Mặc dù tiện lợi, FE không cung cấp toàn bộ quyền kiểm soát như viết SQL thuần.
**Độ phức tạp với ứng dụng lớn:**
- EF có thể gây khó khăn khi sử dụng có quá nhiều ánh xạ và quan hệ phức tạp.

