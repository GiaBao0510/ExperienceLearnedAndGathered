## **1. Generic:**
- Thường được gọi là kiểu đại diện, giúp tránh việc viết code lại nhiều lần (Tức là tránh viết lại một giải thuật giống nhau lại nhiều lần nhưng lại khác kiểu dữ liệu).
- _Cú pháp kiểu Generic trên **phương thức**:_
	<Kiểu trả về> <Tên phương thức> **<** <Tên kiểu đại diện> **>**().
- Trong **phương thức Generic** có ít nhất 1 biến có kiểu đại diện.
- Khi nếu dùng phương thức Generic có quy định kiểu dữ liệu trước thì bắt buột tham số truyền vào phải đúng theo kiểu đó.
-  _Cú pháp kiểu Generic trên **lớp**:_
	<Phạm vi truy cập> class <tên lớp> **<** <Tên kiểu đại diện> **>**{
		_//Khối lệnh_
	}
- Nếu mà gọi lớp có kiểu Generic thì phải quy định kiểu cho lớp đó trước.
