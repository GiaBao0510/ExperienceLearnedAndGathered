## **1. Lớp, Lập trình hướng đối tượng:**

**Lập trình hướng đối tượng(OOP):** là kỹ thuật lập trình yêu cầu trừu tượng hóa các vấn đề về thành các đối tượng. Kỹ thuật lập trình hướng đối tượng có 4 tính chất:
- **Tính trừu tượng:** là một cách tổng quát hóa các thông tin không chi tiết về các đối tượng không gắn cúng với một đối tượng cụ thể cần mô tả.
- **Tính đóng gói:** là dữ liệu đối tượng cố gắng như được đặt trong 1 hộp đen, các thành viên khác ngoài đổi tượng không thể truy cập trực tiếp đến dữ liệu.
- **Tính đa hình:** đối tượng có thể ứng xử khác nhau tùy theo trường hợp cụ thể.
- **Tính kế thừa:** đặc tính của lớp được kế thừa từ lớp khác.

#### **Khái niệm về lớp:** 
- Lớp là một kiểu dữ liệu tham chiếu. Trong kiểu dữ liệu này có thể chứa các thuộc tính và các phương thức .Gọi chung là các thành viên trong một lớp.
- Lớp được xây dựng bằng cách phản ảnh trừu tượng hóa, tổng quát hóa những sự vật hiện tượng xuất hiện trong đời sống.
- Trong các lớp luôn có phương thức khởi tạo. Các lớp này luôn được đầu tiên, khi lớp được tạo ra. Phương thức khởi tạo có cùng tên với tên lớp.

**_Cú pháp xây dựng lớp:_**
**<Phạm vi truy cập>** class **<Tên lớp>**{
	//Khai báo các thuộc tính
	//Khai báo các phương thức
}
#### **Phạm vi truy cập của lớp:**
- **_public:_** Lớp này có thể được truy cập bởi bất kỳ đâu.
- **_internal:_** lớp này chỉ được truy cập khi cùng assemply (_Nghĩa là cùng chương trình_), những chương trình khác sẽ không thể truy cập vào lớp này một các trực tiếp . Nếu khai báo lớp mà không đặt phạm vi truy cập thì mặc định nó tự đặt là internal.

#### **2.Thuộc tính(properties):**
- Là các thành viên trong lớp, dùng để hiển thị , sửa đổi, gán dữ liệu bên trong một lớp.
- Trong 1 thuộc tính có cả hai set và get hoặc chỉ có 1 set hoặc get.

**_Cú pháp:_** <phạm vi truy cập> <kiểu dữ liệu> <tên thuộc tính> {set; get;}

#### **4.Hàm hủy:**
- hàm này được chạy một cách tự động, khi đối tượng đó được giải phóng khỏi bộ nhớ.
- Hàm hủy này được gọi khi hệ thống cảm thấy thiếu bộ nhớ.
- Một lớp chỉ được phép có 1 phương thức hủy.

**_Cú pháp:_** hàm hủy có cùng tên với lớp và bắt đầu bằng dấu ngã, không có kiểu trả về
~<tên lớp>(){}

#### **5. Triển khai IDisaposable và từ khóa using:**
- Trong trường hợp muốn quản lý giải phóng bộ nhớ thì kế thừa IDisaposable, dùng để giải phóng tài nguyên mà đối tượng chiếm giữa.