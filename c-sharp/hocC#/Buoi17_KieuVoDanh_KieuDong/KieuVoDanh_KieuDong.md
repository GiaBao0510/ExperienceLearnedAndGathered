## **1. Kiểu vô danh(Anonymous type):**

- Đối tượng thuộc kiểu này **chỉ có thể đọc giá trị**, chứ không được phép thay đổi giá trị.
- _Cú pháp:_
	new { thuoctinh1 = giatri1, thuoctinh2 = giatri2,...}
- Kiểu vô danh luôn có kiểu dữ liệu là **var**.

## **2. Kiểu động(dynamic):**
- Khai báo kiểu với từ khóa **dynamic**.
- Kiểu dynamic và kiểu var giống ở chỗ cùng là **kiểu ngầm định**, khác cái là ở kiểu var khi khai báo biến kiểu var thì phải gán ngay giá trị cho biến đó để xác định kiểu dữ liệu. Còn kiểu dynamic thì để xác định được biến này có kiểu gì thì lúc biên dịch chương trình mới biết biến đó có kiểu như thế nào.
- Nếu có lỗi xuất hiện thì lúc chạy chương trình mới báo lỗi.
- Khi tạo biến có kiểu dynamic thì chúng ta không thể nào xác định được phương thức cụ thể. Do đó chúng ta thoải mái truy cập vào biến phương thức của biến đó. 
