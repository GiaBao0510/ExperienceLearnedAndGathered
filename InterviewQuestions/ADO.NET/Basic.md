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
- Các component ngắt kiết nối xây dựng cơ sở kiến trúc cơ bản của ADO.NET .Có thể sử dụng các Component này (hoặc lớp) có hoặc là không có các nhà cung ấp dữ liệu (data providers). _Ví dụ: Có thể sử dụng đối tượng DataTable có hoặc không có dữ liệu từ nhà cung cấp_. Các component chung hoặc chia sẻ này được sử dụng chung bới tất cả các nhà cung cấp dữ liệu. Các thành phần nhà cung cấp dữ liệu được thiết kế đặt biết để làm việc với các nguồn dữ liệu khác nhau. _Ví dụ: nhà cung cấp dữ liệu ODBC hoạt động với các nguồn dữ liệu ODBC và các nhà hoạt động với nguồn dữ liệu ODBC và các nhà cung cấp dữ liệu ODBC và nhà cung cấp dữ liệu OleDB hoạt động với nguồn dữ liệu OLE-DB.

### **4.  Trình bày cấu trúc của DataSet?**

![[dataset.gif]] 
- Một đối tượng **DataSet** thuộc vào dòng thành phần ngắt kết nối (disconnected components). ==DataSet== bao gồm một tập hợp các bảng (Tables), Hàng(rows), cột (columns) và mối quan hệ (relationship).
- ==DataSet== chứa tập hợp các ==DataTable== và ==DataTable== chứa một tập hợp các ==DataRow==, ==DataRelation== và ==DataColumn==. Một ==DataTable== tương ứng với một bảng trong cơ sở dữ liệu.

### **5. Connection Pooling trong ADO.NET là gì?**

- ADO.NET sử dụng kỹ thuật kết nối gọi là **Connection Pooling**, giúp giảm thiểu chi phí của việc mở rộng và đóng kết nối lặp đi lặp lại nhiều lần
- _" Conection Pooling tái sử dụng các kết nối hiện đang hoạt động có cùng chuỗi kết nối thay vì tạo các kết nối mới khi có yêu cầu tới cơ sở dữ liệu."_
- Kỹ thuật này bao gồm việc sử dụng quản lý kết nối(connection manager) có nhiệm vụ duy trì một danh sách hoặc pool, các kết nối sẳn cho một chuỗi cụ thể. Nhiều Pool kết nối tồn tại nếu các chuỗi kết nối khác nhau yêu cầu sử dụng connection pooling.
![[Pasted image 20241213102959.png]]

### **6. Connection Pooling trong ADO.NET là gì?**

