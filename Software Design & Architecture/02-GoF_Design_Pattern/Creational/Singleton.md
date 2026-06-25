### **Singleton là gì?**

- **Singleton** là một mẫu thiếu kế thuộc nhóm **Creational Patterns**, đảm bảo **một lớp chỉ có một thể hiện duy nhất (instance)** trong suốt vòng đời của ứng dụng và cung cấp  một phương thức truy cập đến **Instance** đó từ mọi nơi (global access). 

![](https://images.viblo.asia/8cc36217-fa29-496b-a2ab-03a5286d8b6b.png)

---

### **Kiến trúc**

- Đầu tiên đặt ==**Constructor** là private== để không cho client có thể khởi tạo object của lớp.
- Tạo một biến ==static private với tên là instance của lớp đó==, để đảm bảo rằng nó là duy nhất và chỉ được tạo ra trong lớp đó.
- Tạo một hàm ==**public static method trả về instance vừa khởi tạo bên trên==,** đây là cách duy nhất để các lớp khác có thể truy cập vào instance của lớp này.

![](https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSxsdEhte4FyIrNEWYp-yIWPANF3mwxSV5cXw&s)

---

### **Đặc điểm của Singleton?**

- **Duy nhất (unique):** chỉ tồn tại một thể hiện duy nhất của lớp. Ví dụ: quản lý cấu hình, kết nối cơ sở dữ liệu, logging, hoặc thread pool.
- **Toàn cục (Global Access):**  có thể truy cập đến thể hiện này ở bất kỳ đâu trong ứng dụng.
- **Kiểm soát việc khởi tạo:** thể hiện chỉ được tạo ra khi cần thiết (Lazy Initialization).

---

### **Nên sử dụng Singleton khi nào?**

- Khi cần một đối tượng duy nhất trong toàn bộ ứng dụng.
- **Quản lý cấu hình:** Cấu hình ứng dụng cần 1 thể hiện.
- **Quản lý kết nối cơ sở dữ liệu:** Đảm bảo việc tái sử dụng kết nối và tối ưu tài nguyên.
- **Hệ thống ghi log:** một logger dùng chung trong toàn bộ chương trình.

### **Khi nào không nên áp dụng Singleton?**

- Khi lớp cần hỗ trợ nhiều instance trong các ngữ cảnh khác nhau.
- Khi việc kiểm tra unit test trở nên phức tạp (Singleton có thể gây khó khăn trong việc mock hoặc thay thế instace).
- Khi nguy cơ lạm dụng, dẫn đến việc thiết kế giống như biến toàn cục (global variable), làm giảm tính linh hoạt và khó bảo trì.

---

### **Ưu & nhược điểm**

**Ưu điểm:**

- **Đảm bảo tính duy nhất:** Chỉ có một thể hiện trong toàn bộ ứng dụng.

- **Tiết kiệm tài nguyên:** Đặc biệt quan trọng khi khởi tạo đối tượng tốn kém như kết nối cơ sở dữ liệu.

- Đối tượng **singleton** chỉ được khởi tạo duy nhất trong một lần khi nó được yêu cầu lần đầu.

- Kiểm soát việc truy cập đến instance duy nhất.

- Giảm namespace.

- **Truy cập toàn cục:** Đơn giản hóa việc truy cập đối tượng dùng chung.
  
  **Nhược điểm:**
  
  - **Phá vỡ nguyên tắt SRP (Single Responsibility Principle):** Singleton vừa quản lý công việc khởi tạo, vừa cung cấp logic nghiệp vụ
  - **Khó kiểm tra (Unit Test):** Vì Singleton sử dụng trạng thái toàn cục, việc kiểm tra có thể khó khăn hơn.
  - Có thể gây ra lỗi trong môi trường đa luồng nếu không được triển khai đúng cách

---

### **So sánh giữa Singleton Pattern với Service Singleton trong .NET**

Trong .NET (Đặc biệt là ASP.NET Core), thì có thể đăng ký dịch vụ vòng đời **Singleton**. Sau đây là sự so sánh giữa **Singleton Pattern** và Service Singleton **trong** .NET:

| Tiêu chí                        | Singleton Pattern (truyền thống)                                                                  | Service Singleton trong .NET                                                                  |
| ------------------------------- | ------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| **Cách triển khai**             | Tự triển khai logic trong lớp (`constructor` `private`, biến static, phương thức lấy `instance`). | Đăng ký dịch vụ với `AddSingleton` trong `IServiceCollection`, container DI quản lý instance. |
| **Kiểm soát instance**          | Lớp tự quản lý `instance` duy nhất, thường sử dụng static và `lock` để thread-safe.               | Container DI của .NET đảm bảo chỉ có một instance duy nhất, tự động thread-safe.              |
| **Tính linh hoạt**              | Ít linh hoạt hơn, khó thay thế hoặc mock instance khi unit testing.                               | Linh hoạt hơn, dễ dàng thay thế hoặc mock thông qua DI (Dependency Injection).                |
| **Khả năng kiểm tra (Testing)** | Khó kiểm tra vì instance là cố định, khó mock hoặc thay thế trong unit test.                      | Dễ dàng kiểm tra nhờ DI, có thể inject mock hoặc implementation khác.                         |
| **Quản lý vòng đời**            | Nhà phát triển phải tự quản lý vòng đời của instance (khởi tạo, hủy).                             | Container DI tự động quản lý vòng đời, tích hợp với scope của ứng dụng.                       |
| **Phạm vi sử dụng**             | Thường sử dụng trong ứng dụng không có DI container hoặc cần kiểm soát chặt chẽ.                  | Phù hợp với ứng dụng .NET hiện đại sử dụng DI (như ASP.NET Core).                             |
| **Ví dụ sử dụng**               | Quản lý cấu hình, logging, hoặc kết nối cơ sở dữ liệu trong ứng dụng không DI.                    | Đăng ký dịch vụ như DbContext, HttpClientFactory, hoặc các dịch vụ dùng chung.                |

---

## **Tài liệu tham khảo:**

1. https://viblo.asia/p/singleton-design-pattern-tro-thu-dac-luc-cua-developers-Qbq5QBkJKD8
2. https://viblo.asia/p/hoc-singleton-pattern-trong-5-phut-4P856goOKY3