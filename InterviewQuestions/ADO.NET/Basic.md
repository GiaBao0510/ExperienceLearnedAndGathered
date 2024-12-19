### **1. ADO.NET là gì?
- **ADO(Active Data Object)** và ADO.NET là một bộ thư viện .NET cho ADO.NET là một bộ các thư viện quản lý được sử dụng bởi các ứng dụng .NET để giao tiếp với nguồn dữ liệu bằng cách sử dụng một trình điều khiển(drive) hoặc provider:
	-    Các ứng dụng doanh nghiệp thường xử lý với nguồn dữ liệu lớn. Dữ liệu này chủ yếu được lưu trữ dạng cơ sở dữ liệu quan hệ( MySQL, SQL Server, Oracle, Access, ...). Các cơ sở dữ liệu này thường dùng ngôn ngữ truy vấn SQL để truy xuất dữ liệu.
	-    Để truy cập vào dữ liệu từ một ứng dụng .NET, cần phải có một giao diện. Giao diện này hoạt động như một cầu nối giữa hệ thống và quản lý cơ sở dữ liệu(RDBMS - Relational Database Managemanet System) và ứng dụng .NET .ADO.NET chính là một giao diện, được tạo ra để kết nối ứng dụng .NET với hệ thống RDBMS.
	- Bất kỳ ứng dụng nào từ .NET dù trên máy tính hay website điều có thể tương tác với cơ sở dữ liệu bằng cách sử dụng các lớp của thư viện ADO.NET .Dữ liệu có thể được truy xuất bất kỳ cơ sở dữ liệu nào với kiến trúc kết nối (connected) hoặc ngắt kết nối (disconnected).

### **2.  Ý nghĩa chính xác của phương pháp disconnected và connected trong ADO.NET là gì?

Tóm lại:
- **Disconnected:** thực hiện kết nối -> truy xuất dữ liệu -> đóng kết nối.
- **Connected:** thực hiện kết nối -> giữ kết nối và đóng kết nối khi được gọi.

> Kiến trúc ADO.NET, trong đó kết nối phải được mở cho đến cuối cùng để truy xuất và truy cập dữ liệu từ cơ sở dữ liệu, được gọi là kiến trúc kết nối (connected architecture). Kiến trúc kết nối được xây dựng trên các loại: Kết nối(==connection==), lệnh(==command==), đọc dữ liệu (==datareader==).

>Kiến trúc ADO.NET ,trong đó kết nối chỉ được giữ mở cho đến khi dữ liệu được truy xuất từ cơ sở dữ liệu và sau đó có thể truy cập ngay cả khi kết nối đến cơ sở dữ liệu đã được đóng, được gọi là kiến trúc ngắt kết nối (disconnected architecture). Kiến trúc ngắt kết nối đươc xây dựng trên các lệnh sau: kêt nối (==connection==), trình điều kiển (DataAdapter), xây dựng lệnh (==commandbuilder==), tập dữ liệu và xem dữ liệu (==dataview==). 

### **3.  Các component của ADO.NET là gì?
- Các Component của ADO.NE được phân loại vào 3 chế độ:
	- Ngắt kết nối (Disconnected).
	- Chung hoặc được chia sẻ (Common or Shared), và
	- Các nhà cung cấp dữ liệu .NET.
- Các component ngắt kết nối xây dựng cơ sở kiến trúc cơ bản của ADO.NET .Có thể sử dụng các Component này (hoặc lớp) có hoặc là không có các nhà cung cấp dữ liệu (data providers). _Ví dụ: Có thể sử dụng đối tượng DataTable có hoặc không có dữ liệu từ nhà cung cấp_. Các component chung hoặc chia sẻ này được sử dụng chung bới tất cả các nhà cung cấp dữ liệu. Các thành phần nhà cung cấp dữ liệu được thiết kế đặt biết để làm việc với các nguồn dữ liệu khác nhau. _Ví dụ: nhà cung cấp dữ liệu ODBC hoạt động với các nguồn dữ liệu ODBC và các nhà hoạt động với nguồn dữ liệu ODBC và các nhà cung cấp dữ liệu ODBC và nhà cung cấp dữ liệu OleDB hoạt động với nguồn dữ liệu OLE-DB.

### **4.  Trình bày cấu trúc của DataSet?**

![[dataset.gif]] 
- Một đối tượng **DataSet** thuộc vào dòng thành phần ngắt kết nối (disconnected components). ==DataSet== bao gồm một tập hợp các bảng (Tables), Hàng(rows), cột (columns) và mối quan hệ (relationship).
- ==DataSet== chứa tập hợp các ==DataTable== và ==DataTable== chứa một tập hợp các ==DataRow==, ==DataRelation== và ==DataColumn==. Một ==DataTable== tương ứng với một bảng trong cơ sở dữ liệu.

### **5. Connection Pooling trong ADO.NET là gì?**

- ADO.NET sử dụng kỹ thuật kết nối gọi là **Connection Pooling**, giúp giảm thiểu chi phí của việc mở rộng và đóng kết nối lặp đi lặp lại nhiều lần
- _" Conection Pooling tái sử dụng các kết nối hiện đang hoạt động có cùng chuỗi kết nối thay vì tạo các kết nối mới khi có yêu cầu tới cơ sở dữ liệu."_
- Kỹ thuật này bao gồm việc sử dụng quản lý kết nối(connection manager) có nhiệm vụ duy trì một danh sách hoặc pool, các kết nối sẳn cho một chuỗi cụ thể. Nhiều Pool kết nối tồn tại nếu các chuỗi kết nối khác nhau yêu cầu sử dụng connection pooling.
![[Pasted image 20241213102959.png]]

### **6. SqlCommand object là gì?**
- **SqlCommand** chứa câu lệnh SQL cần được thực thi trên cở sở dữ liệu. đối tượng SqlComand chứa 2 tham số là: **CommandText, Connection**
	- **CommandText:** tham số này thường sẽ chứa đoạn văn bản truy vấn đế cơ sở dữ liệu SQL.
	- **Connection:** tham số này sẽ dựa vào nơi kết nối đến cơ sở dữ liệu để thực hiện câu truy vấn SQL.
	- Các truy vấn ở dạng văn bản trực tiếp (Inline text), thủ tục lưu trữ (Stored Procedures) hoặc truy cập trực tiếp bảng dữ liệu(Table access).
	- Tính năng quan trọng của đối tượng ==Command ==là có thể sử dụng thực thi các truy vấn và thủ tục lưu trữ tham số(Parameters).
	- Nếu một truy vấn _select_ được thực thi thì, tập kết quả trả về thường được lưu trữ trong đối tượng ==DataSet== hoặc ==DataReader==.
- Từ hình ảnh bên dưới mô tả 3 phương thức quan trọng mà đối tượng SqlCommand cung cấp được hiển thị: ==ExecuteNonQuery, ExecuteReader,  ExecuteScalar==
![[Pasted image 20241215203911.jpg]]
_Ví dụ:_
```
// Tạo đối tượng DbCommand
using var command = new SqlCommand();
command.Connection = connection;

// Câu truy vấn gồm: chèn dữ liệu vào và lấy định danh(Primary key) mới chèn vào
string queryString = @"INSERT INTO Shippers (Hoten, Sodienthoai) VALUES (@Hoten, @Sodienthoai);
                       SELECT CAST(scope_identity() AS int)";

command.CommandText = queryString;
command.Parameters.AddWithValue("@Hoten", "Abc");
command.Parameters.AddWithValue("@Sodienthoai", 123456);

var ShipperID = command.ExecuteScalar();
```
### **7. Bạn hiểu lớp DataRelation như thế nào?**
- **DataRelation** là một lớp kiến trúc của Disconnected(ngắt kết nối) của .NET Framework. Lớp này thuộc **namespace** tên **System.Data**. **DataRelation** đại diện cho một mối quan hệ giữa các bảng cơ sở dữ liệu và liên kết các bảng dựa trên cột phù hợp
```
DataSet dataset = LoadData();
DataTable userTable = dataset.Table["User"];
DataTable groupTable = dataset.Table["Group"];

DataRelation relation = DataRelation("Group_user", groupTable.Columns["GroupID"], userTable.Columns["GroupID"]);

dataset.Relations.Add(relation);
```

![DataRelation](https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSghLSuozJRX32wJ7zB9FPty5w_JT8Ikn6JXg&s)

### **8. DataAdapter trong ADO.NET là gì?**

- DataAdapter được sử dụng để lấy dữ liệu từ nguồn dữ liệu và điền vào các bảng trong DataSet. DataAdapter tạo sự kết nối giữa nguồn dữ liệu và và datatset. DataAdapter cũng giải quyết các thay đổi được thực hiện trên DataSet trở lại nguồn dữ liệu. DataAdapter sử dụng đối tượng Connection của nhà cung cấp dữ liệu .NET Framework để kết nối từ nguồn dữ liệu, và nó sử dụng đối tượng Command để lấy dữ liệu từ nguồn dữ liệu và giải quyết các thay đổi
- DataAdapter chủ yếu hỗ trợ 2 phương thức sau:
	- **Fill():** phương thức Fill điền dữ liệu vào một đối tượng dataset hoặc datatable từ cơ sở dữ liệu. Nó lấy các hàng từ nguồn dữ liệu bằng các sử dụng lệnh **SELECT** được chỉ định bởi thuộc tính select command liên kết .Phương thức Fill để lại kết nối trong cùng trạng thái mà nó gặp trước khi điền dữ liệu
	- **Update():** phương thức Update xác nhận thay đổi trở lại csdl. Nó cũng phân tích trạng thái của mỗi bản ghi trong DataSet và gọi các câu lệnh INSERT, UPDATE và DELETE
![DataApdater](https://www.c-sharpcorner.com/UploadFile/mahesh/command-object-command-builder-data-adapter-object-ado-net/Images/Figure-3.6.gif)

_Ví dụ:_
```
//Tạo kết nối
var sqlconnectstring = @"Data Source=localhost,1433;
                         Initial Catalog=xtlab;
                         User ID=SA;Password=Password123";
var connection = new SqlConnection(sqlconnectstring);
connection.Open();

//Tạo DataAdapter 
SqlDataAdapter adapter = new SqlDataAdapter();

//Thiết lập bản trong DataSet ánh xạ tương ứng có tên là NhanVien
adapter.SelectCommand = new SqlCommand(@"SELECT NhanviennID,Ten,Ho FROM Nhanvien", connection);

// InsertCommand - Thực khi khi gọi Update, nếu DataSet có chèn dòng mới (vào DataTable)
// Dữ liệu lấy từ DataTable, như cột Ten tương  ứng với tham số @Ten

adapter.InsertCommand = new SqlCommand(@"INSERT INTO Nhanvien (Ten, Ho) VALUES (@Ten, @Ho)", connection);
adapter.InsertCommand.Parameters.Add("@Ten", SqlDbType.NVarChar, 255, "Ten");
adapter.InsertCommand.Parameters.Add("@Ho", SqlDbType.NVarChar, 255, "Ho");

// DeleteCommand  - Thực thi khi gọi Update, nếu có remove dòng nào đó của DataTable
adapter.DeleteCommand = new SqlCommand(@"DELETE FROM Nhanvien WHERE NhanviennID = @NhanviennID", connection);
var pr1 = adapter.DeleteCommand.Parameters.Add(new SqlParameter("@NhanviennID", SqlDbType.Int));
   pr1.SourceColumn = "NhanviennID";
   pr1.SourceVersion = DataRowVersion.Original;  // Giá trị ban đầu

// UpdateCommand -  Thực thi khi gọi Update, nếu có chỉnh sửa trường dữ liệu nào đó
adapter.UpdateCommand = new SqlCommand(@"UPDATE Nhanvien SET Ho=@Ho, Ten = @Ten
                                         WHERE NhanviennID = @NhanviennID", connection);
adapter.UpdateCommand.Parameters.Add("@Ten", SqlDbType.NVarChar, 255, "Ten");
adapter.UpdateCommand.Parameters.Add("@Ho", SqlDbType.NVarChar, 255, "Ho");
var pr2 = adapter.UpdateCommand.Parameters.Add(
    new SqlParameter("@NhanviennID", SqlDbType.Int)
    { SourceColumn = "NhanviennID" });
pr2.SourceVersion  = DataRowVersion.Original;

DataSet dataSet = new DataSet();

// Thực hiện lấy dữ liệu từ nguồn về DataSet
adapter.Fill(dataSet);
// Lấy DataTable kết quả và hiện thị
DataTable dataTable = dataSet.Tables["Nhanvien"];
ShowDataTable(dataTable);
```

### **9. DataAdapter trong ADO.NET là gì?**