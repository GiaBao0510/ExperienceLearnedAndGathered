## Giới thiệu về ADO.NET:
- **ADO.NET(ActiveX Data Object):** là một tập hợp các thư viện cho phép ứng dụng có thể( lấy, thêm , sửa ,xóa) với các nguồn dữ liệu như là (SQL server, XML, MySQL, Oracle Database).
- Kiến trúc để thực hiện truy cập dữ liệu với _ADO.NET_ được phân ra nhiều phần rời rạc. Mỗi phần có thể sử dụng độc lập hoặc đồng thời nhiều thành phần được sử dụng. Ví dụ:
![](https://raw.githubusercontent.com/xuanthulabnet/learn-cs-netcore/master/imgs/cs058.png)

##### 1. **DataProvider** là các thư viện lớp cung cấp chức năng tạo kết nối đến nguồn dữ liệu, thi hành các lệnh trên nguồn dữ liệu đó _Insert, update, delete_.
- **SQL Server**: là loại DataProvider mặc định trong **.NET core** là SqlClient ở _namespace System.Data.SqlClient_ cung cấp khả năng kết nối đến **SQL server**.
- **MySQL** :là loại DataProvider chưa có sẳn trong **.NET core** .Nên phải tải package _MySQL.Data_, sẽ có DataProvider _MySql.Data.MySqlClient_
- **SQLite:** phải tải gói Microft.Data.SQLite.
##### 2.  **DataSet** là các thư viện lớp (độc lập với Data Provider) tạo ra các đối tượng để quản lý dữ liệu không phụ thuộc ngồn dữ liệu đến từ đâu, đã ở trong ứng dụng (local) hay từ nguồn XML. - DataSet thường gồm nhiều DataTable, trong DataTable lại gồm DataColumn, các dàng buộc, các khóa chính ... Vậy DataSet là sự trừu tượng hóa một CSDL thực.

- Thêm gói SqlClient:
```
dotnet add package System.Data.SqlClient
```
- Thêm namespace:
```
using System.Data;
using System.Data.SqlClient;
```

# **Buoi38_Lớp SqlCommand - khởi tạo đối tượng SqlCommand**
- Lớp _SqlCommand_ triển khai từ _DbCommand_ cho phép tạo ra đối tượng, các đối tượng này có thể thực hiện các lệnh SQL để tương tác đến CSDL, như các mệnh đề **select | update | create tables | ...** cũng như cho phép thi hành các hàm, các **stored procedure** cuả Database.
- Để thi hành được _SqlCommand_ tương tác đến cơ sở dữ liệu, thì trước hết thiết lập cho nó một cái kết nối đến CSDL (**SqlConnection**), sau đó thêm câu truy vấn và các tham số cho câu lệnh truy vấn.
- Để thi hành lệnh SQL với SqlCommand, thì cần có một kết nối trước ([SqlConnection](https://xuanthulab.net/ado-net-gioi-thieu-ado-net-va-ket-noi-sql-server-voi-sqlconnection.html)), rồi tạo ra đối tượng SqlCommand, gán cho nó kết nối, câu lệnh SQL sau đó mới thi hành được. Để thi hành, gọi một trong các phương thức như `ExecuteScalar`, `ExecuteNonQuery`, `ExecuteReader` `...`

#### **1. Thiết lập các tham số SqlCommand**:
- Trong các câu lệnh SQL có thể viết chứa **tên tham số** trong nó, sao đó giá trị thực của tham số này được SqlCommand gán thay vào để có mệnh đề SQl thực sự. Tham số trong chuỗi câu lệnh truy vấn SQL ký hiện là **@tenthamso** (Cần phải kèm theo dấu @).

#### **2. Các cách thi hành SqlCommand và lấy kết quả truy vấn:**
Có các phương thức khác nhau để thi hành SqlCommad tùy theo ngữ cảnh và mục đích:
- `ExecuteNonQuery()` thi hành truy vấn - không cần trả về dữ liệu gì, phù hợp thực hiện các truy vấn như `Update`, `Delete`,....
-  `ExecuteReader()` thi hành lệnh - trả về đối tượng giao diện `IDataReader` như `SqlDataReader`, từ đó đọc được dữ liệu trả về
-  `ExecuteScalar()` thì hành và trả về một giá trị duy nhất - ở hàng đầu tiên, cột đầu tiên


# **Buoi39: # DataAdapter DataSet và DataTable tìm hiểu và sử dụng** 

### **DataSet:**
- là một cấu trúc phức tạp, là thành phần cơ bản trong ADO.NET nó ánh xạ CSDL nguồn(SQL Database) vào thành phần đối tượng trong bộ nhớ. **DataSet** chứa trong nó là một tập hợp các đối tượng **DataTable**
- Khởi tạo: ``` DataSet dataser = new DataSet ```