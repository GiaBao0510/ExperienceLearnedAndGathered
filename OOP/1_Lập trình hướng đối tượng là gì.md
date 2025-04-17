
```
Lập trình hướng đối tượng(Object-Oriented Programming - OOP) là một phương pháp, kỹ thuật lập trình cho phép tạo ra các đối tượng trong code bằng cách trừu tượng hóa các đối tượng thực tế trong cuộc sống.
```
---
**_Lớp (Class):_** Có thể xem như là một khuôn mẫu(template) dùng để tạo ra ra các đối tượng. Trong lớp sẽ gồm có thuộc tính(Attribute) và phương thức(Methob).
- **Attributes/ Fields/ Properties:** định nghĩa đặc trưng, đặc tính của đối tượng.
- **Methobs**: định nghĩa các hành vi, các hành động của object.

![](https://statics.cdn.200lab.io/2023/08/oop-building-blocks-1.jpg?width=800)

***Ví dụ:*** ta có class Person với các attributes gồm: họ tên, tuổi tác, địa chỉ. Với Methobs gồm: di chuyển, ăn, nghỉ ngơi.

---
**_Đối tượng:_** Có thể được xem là một thể hiện cụ thể của lớp, vì thế khi đối tượng được khởi tạo sẽ mang đầy đủ thông tin cụ thể.

![](https://statics.cdn.200lab.io/2023/08/oop-building-blocks-1.jpg?width=800)


---
### **Tính chất cơ bản trong OOP:**

![](https://statics.cdn.200lab.io/2023/08/oop-4-tinh-chat.jpg?width=800)

Trong OOP sẽ bao gồm có 4 tính chất, gồm: 
- **Tính đóng gói (Encapsulation):** Đóng gói là tính chất cho phép bảo vệ dữ liệu của đối tượng bằng cách giới hạn quyền truy cập. Dữ  liệu và các phương thức thao tác với cơ sở dữ liệu được gộp lại trong một lớp, và chỉ có những phương thức công khai (public) được phép truy cập hoặc sửa đổ dữ liệu đó.
- **Tính trừu tượng (Abstraction):** Cho phép bạn tạo ra các lớp hoặc giao diện mà chỉ định các phương thức mà không cần cung cấp thông tin chi tiết cụ thể về phương thức. Điều này giúp người dùng chỉ cần biết cách sử dụng mà không cần hiểu sâu.  
- **Tính đa hình (Polymorphism):** Cho phép các đối tượng thuộc các lớp khác nhau có thể được xử lý thông qua cùng một giao diện. Điều này giúp cho mã linh hoạt hơn
- **Tính kế thừa (Inheritance):** Cho phép tạo ra một lớp mới dựa một lớp đã có sẳn, giúp tái sử dụng mã và mở rông chức năng của lớp cha mà không cần viết lại mã

---
### **Ưu điểm của lập trình hướng đối tượng:**
- Tái sử dụng lại mã nguồn (nhờ tính kế thừa)
- Dễ bảo trì và mở rộng
- phù hợp cho những dự án phức tạp

---
#### **Khi nào đối tượng được tạo ra ?**
- Đối tượng được tạo ra là khi sử dụng từ khóa **new** để tạo ra một thể hiện của một lớp trong bộ nhớ.

#### **Một lớp có thể tạo ra bao nhiêu đối tượng?**
- Một lớp có thể tạo ra không giới hạn số lượng đối tượng

#### **Sự khác biệt giữa lớp và đối tượng**

| Tiêu chí      | Lớp(class)                               | Đối tượng(object)                                                  |
| ------------- | ---------------------------------------- | ------------------------------------------------------------------ |
| **Khái niệm** | là một khuôn mẫu để tạo ra đối tượng     | Là một thể hiện cụ thể được tạo ra từ lớp                          |
| **Tồn tại**   | Là một khái niệm trừu tượng(logical).    | Là một thực thể cụ thể trong bộ nhớ                                |
| **Công dụng** | Định nghĩa các thuộc tính và phương thức | Lưu trữ dữ liệu thông tin cụ thể và thực hiện hành vi dựa trên lớp |
#### **Sự khác biệt giữa lớp và đối tượng**
- **Lớp:** là kiểu tham chiếu(reference type), lưu trữ trên heap, hỗ trợ kế thừa
- **Cấu trúc dữ liệu:** là kiểu giá trị (value type), lưu trữ trên stack , không hỗ trợ kế thừa

---
### **5W  and 1H?**

**What: Lập trình hướng đối tượng là gì?**
**Why: Tại sao cần lập trình hướng đối tượng?**
**Where: Lập trình hướng đối tượng áp dụng ở đâu>**
**When: Lập trình hướng đối tượng áp dụng khi nào?**
**Who: Ai sẽ là người áp dụng lập trình hướng đối tượng?**
**How: Lập trình hướng đối tượng áp dụng bằng cách nào?** 