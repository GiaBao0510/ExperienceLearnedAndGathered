
```
Lập trình hướng đối tượng(Object-Oriented Programming - OOP) là một phương pháp, kỹ thuật lập trình cho phép tạo ra các đối tượng trong code bằng cách trừu tượng hóa các đối tượng thực tế trong cuộc sống.
```

**_Lớp (Class):_** Có thể xem như là một khuôn mẫu(template) dùng để tạo ra ra các đối tượng. Trong lớp sẽ gồm có thuộc tính(Attribute) và phương thức(Methob).
- **Attributes:** định nghĩa đặc trưng, đặc tính của đối tượng.
- **Methobs**: định nghĩa các hành vi, các hành động của object.

**_Đối tượng:_** Có thể được xem là một thể hiện cụ thể của lớp, vì thế khi đối tượng được khởi tạo sẽ mang đầy đủ thông tin cụ thể.

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

| Tiêu chí      | Lớp(class)                               | Đối tượng(Attribute)                                               |
| ------------- | ---------------------------------------- | ------------------------------------------------------------------ |
| **Khái niệm** | là một khuôn mẫu để tạo ra đối tượng     | Là một thể hiện cụ thể được tạo ra từ lớp                          |
| **Tồn tại**   | Là một khái niệm trừu tượng(logical).    | Là một thực thể cụ thể trong bộ nhớ                                |
| **Công dụng** | Định nghĩa các thuộc tính và phương thức | Lưu trữ dữ liệu thông tin cụ thể và thực hiện hành vi dựa trên lớp |
#### **Sự khác biệt giữa lớp và đối tượng**
- **Lớp:** là kiểu tham chiếu(reference type), lưu trữ trên heap, hỗ trợ kế thừa
- **Cấu trúc dữ liệu:** là kiểu giá trị (value type), lưu trữ trên stack , không hỗ trợ kế thừa

