## **Kiểu dữ liệu cấu trúc(Struct) và kiểu dữ liệu liệt kê**

### **1. Struct:**
- struct là kiểu duex liệu cấu trúc, do chúng ta tự định nghĩa. 
- Khai báo kiễu dữ liệu struct gần giống khai báo kiểu dữ liệu lớp.
- struct có kiểu dữ liệu là tham trị khác với lớp(Kiểu tham chiếu).
- Trong **_struct_** nếu thực hiện phép gán giữa 2 đổi tượng cùng lớp thì, 2 đối tượng đó sẽ _2 đối tượng độc lập nhau_. Ngược lại là **lớp** thì thực hiện phép gán giữa 2 đối tượng cùng lớp thì _2 đối tượng đó sẽ tham chiếu đến cùng 1 địa chỉ trong bộ nhớ (Tức là 2 đối tượng này cùng tham chiếu đến 1 đối tượng)_.
- Kiểu dữ liệu _đơn giản nhỏ_ thì dùng struct. Còn kiểu dữ liệu lớn thì nên dùng lớp.

_Cú pháp:_
	<Phạm vi truy cập> struct <tên struct>{
		//Dữ liệu
		//Phương thức
	}
### **2. Kiểu liệt kê (enum):**
