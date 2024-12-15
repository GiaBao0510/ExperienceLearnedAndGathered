## **1. Lập tình hướng sự kiện:**
- là xây dựng các lớp mà nó phát sinh sự kiện.
- Publisher => class => Phát sự kiện: Những lớp phát ra sựn kiện.
- Subsriber => class => Nhận sự kiện: Những lớp này đăng ký sự kiện thì sẽ nhận thông tin sự kiện.
- Nếu 2 lớp đăng ký nhận cùng 1 sự kiện thì hệ thống chỉ cho 1 lớp cuối cùng nhận sự kiện thôi. => **Solution**: Sử dụng từ khóa **event** đứng trước tên biến có kiểu **delegate**.
- Nếu vậy thì thuộc tính kiểu **delegate** dùng có từ khóa **event** thì để nhận giá trị thì chỉ dùng phép toán += (_Đăng ký nhận sự kiện_), -=(_Đăng ký hủy sự kiện_) không được dùng phép toán gán.
- Cấu trúc **Delegate** chuyên dùng để khai báo sự kiện **Eventhandller** ,được cung cấp sẵn.