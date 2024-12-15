## **1. Sử dụng Delegate:**
- Delegate là kiểu dữ liệu được dùng đề khai báo tạo ra các biến và các biến này dùng để **tham chiếu đến phương thức**.
- Có thể sử dụng biến có kiểu Delegate để **thi hành phương thức** được gán cho biến đó.
- Delegate có thể khai báo bên trong _namespace_ hoặc _class._
- Biến kiểu Delegate **có thể nhận giá trị null**. và biến này khi nhận phương thức làm giá trị, thì phương thức đó phải có cùng kiểu trả về và cùng tham số đầu vào.
- Có thể gọi **hàm Invoke()** từ biến có kiểu dữ liệu **Delegate** để thi hành phương thức mà nó đang nhận;
- Phải kiểm tra biến có kiểu dữ liệu **delegate** phải khác null .Mới chạy không bị lỗi.
- Có thể dùng toán tử **"+="** cho biến có kiểu dữ liệu delegate có thể thì hành **một loạt** các phương thức (có kiểu trả về và thông số đầu vào) mà nó nhận vào. 

## **2. Khai báo Delegate Action:**
- Đây cũng là kiểu delegate luôn.
- Khi dùng dụng từ khóa **Action** để khai báo biến thì  tương đương việc khai báo biến kiểu **Delegate**.
- Example:
	- _**Action** hanhdong1; // delegate void hanhdong1();_
	- _**Action<string, int>** hanhdong2; // delegate void hanhdong1(string a, int b);_

## **3. Delegate func:**
- Đây cũng là kiểu delegate luôn.
- Khi dùng dụng từ khóa **Func** để khai báo biến thì  tương đương việc khai báo biến kiểu **Delegate**. **_Nhưng phải xác định kiểu trả về_**.
- Và kiểu này sẽ dựa vào kiểu dữ liệu của tham số cuối cùng để trả về.
- Example: 
	- **_Func<Kiểu trả về>_** <tên biến>;
	- **_Func< int>_** h;  //delegate int  h(int x);
	- **_Func< int, string, double>_** h2;  //delegate double  h(int x, string y, double z);
