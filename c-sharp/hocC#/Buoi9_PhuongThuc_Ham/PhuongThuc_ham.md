
#### **Phương thức - Hàm:**
là một nhóm các câu lệnh cùng thực hiện công việc gì đó. Đặc biệt là phương thức thì không trả về giá trị, ngược lại thì hàm có trả về giá trị gì đó.

Trong chương trình thì bên trong lớp có một phương thức Main, phương thức này sẽ chạy ngay lập tức khi cho chạy chương trình.

Phương thức có kiểu trả về là **void**, còn hàm thì có kiểu trả về như là: **int, float, bool, double, ...**

Nếu muốn _gọi_ phương thức bên trong một phương thức nào đó thì ta phải gọi lớp(Đối tượng) chứa phương thức đó thì mới gọi được, hoặc có thể đặt từ khóa _static_ ở trước phương thức thì nếu gọi phương thức này thì không cần gọi đến lớp chứa phương thức đó.

_Cú pháp phương thức:_
<Phạm vi truy cập> <kiểu trả về> <Tên phương thức> ([tham số]){
	//Khối lệnh
}

_Cú pháp hàm:_
<Phạm vi truy cập> <kiểu trả về> <Tên hàm> ([tham số]){
	//Khối lệnh
	**return ... ;**
}

**Quá tải phương thức(methob overloading):**
- phương thức có thể _cùng tên_ nhưng khác **kiểu tham số**.
- Phương thức có thể _cùng tên_ nhưng khác **số lượng tham số.**
- phương thức có thể _cùng tên_ nhưng khác **kiểu**.

**Tham chiếu, tham trị:**
- truyền tham chiếu là làm thay đổi giá trị của một biến khi truyền vào phương thức, mà trước tham số nhận vào bên trong phương thức thì trước tham số phải có từ khóa **_ref_** .
- truyền tham trị là khi chuyền 1 biến vào trong 1 tham số cụ bộ của phương thức, thì nó chỉ làm thay đổi giá trị cục bộ thôi, không ảnh hưởng đến giá trị của biến đã được truyền vào phương thức.

!**Nếu tham số là đối tượng lớp thì sẽ mặc định là tham chiếu.**
#### **Phạm vi truy cập:**
- **_public:_** cho phép các lớp bên ngoài có thể truy cập đến.
-  **_private:_** chỉ cho phép nội bộ bên trong một lớp có thể truy cập được. Các lớp khác không thể truy cập đến.
- **_protected:_** chỉ cho phép nội bộ bên trong một lớp có thể truy cập và các lớp kế thừa có thể truy cập. Các lớp khác không thể truy cập đến.