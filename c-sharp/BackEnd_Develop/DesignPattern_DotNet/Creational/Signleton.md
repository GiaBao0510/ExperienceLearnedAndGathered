### **Singleton là gì?**

- **Singleton** là một mẫu thiếu kế thuộc nhóm **Creational Patterns**, đảm bảo **một lớp chỉ có một thể hiệm duy nhất (instance)** trong suốt vòng đời của ứng dụng và cung cấp một điểm truy cập toàn cụ đến thể hiện đó. 

![](https://images.viblo.asia/8cc36217-fa29-496b-a2ab-03a5286d8b6b.png)

---

### **Kiến trúc**

- Đầu tiên đặt ==**Constructor** là private== để không cho client có thể khởi tạo object của lớp.
- Tạo một biến ==static private với tên là instance của lớp đó==, để đảm bảo rằng nó là duy nhất và chỉ được tạo ra trong lớp đó.
- Tạo một hàm ==**public static method trả về instance vừa khởi tạo bên trên==,** đây là cách duy nhất để các lớp khác có thể truy cập vào instance của lớp này.

![](https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSxsdEhte4FyIrNEWYp-yIWPANF3mwxSV5cXw&s)

---

### **Đặc điểm của Singleton?**

- **Duy nhất (unique):** chỉ tồn tại một thể hiện duy nhất của lớp.
- **Toàn cục (Global Access):**  có thể truy cập đến thể hiện này ở bất kỳ đâu trong ứng dụng.
- **Kiểm soát việc khởi tạo:** thể hiện chỉ được tạo ra khi cần thiết (Lazy Initialization).

---

### **Nên sử dụng Singleton khi nào?**

- Khi cần một đối tượng duy nhất trong toàn bộ ứng dụng.
- **Quản lý cấu hình:** Cấu hình ứng dụng cần 1 thể hiện.
- **Quản lý kết nối cơ sở dữ liệu:** Đảm bảo việc tái sử dụng kết nối và tối ưu tài nguyên.
- **Hệ thống ghi log:** một logger dùng chung trong toàn bộ chương trình.

---
### **Ví dụ minh họa với c#:**

**A. Triển khai Singleton với Lazy initialization**

```
public class Singleton1{

	//Biến static giữ thể hiện duy nhất
	public static Singleton1 _instance;

	//Constructor private để ngăn việc khởi tạo từ bên ngoài
	private Singleton1(){}

	//Phương thức public để truy cập thể hiện duy nhất
	public static Singleton1 GetInstance(){
		if(_instance == null){
			_instance = new Singleton1();
		}
		return _instance;
	}

	//Phương thức ví dụ
	public void DoSomething(){
		Console.WriteLine("Singleton Instance is working!");
	}
}
```

- Cách này chỉ sử dụng tốt trong trường hợp đơn luồng, trường hợp nếu có 2 luồng cùng chạy và cùng gọi đến ==hàm GetInstance()== vào trong cùng 1 thời điểm thì đương nhiên sẽ có ít nhất 2 thể hiện của instance.

**B.Triển khai Singleton với Thread-Safety:**

```
public class Singleton1{

	//Biến static giữ thể hiện duy nhất
	public static Singleton1 _instance;
	private static readonly object _lock = new object();

	//Constructor private để ngăn việc khởi tạo từ bên ngoài
	private Singleton1(){}

	//Phương thức public để truy cập thể hiện duy nhất
	public static Singleton1 GetInstance(){
		if(_instance == null){
			lock(_lock){   //Đảm bảo chỉ có một luồng truy cập
				if(_instance == null){
					_instance = new Singleton1();
				}
			}
		}
		return _instance;
	}

	//Phương thức ví dụ
	public void DoSomething(){
		Console.WriteLine("Singleton Instance is working!");
	}
}
```

- Cách này được áp dụng trong trường hợp đa luồng.

---

### **Ưu & nhược điểm**

**Ưu điểm:**
- **Đảm bảo tính duy nhất:** Chỉ có một thể hiện trong toàn bộ ứng dụng.
- **Tiết kiệm tài nguyên:** Đặc biệt quan trọng khi khởi tạo đối tượng tốn kém như kết nối cơ sở dữ liệu.
- Đối tượng singleton chỉ được khởi tạo duy nhất trong một lần khi nó được yêu cầu lần đầu.
- Kiểm soát việc truy cập đến instance duy nhất.
- Giảm namespace.
- **Truy cập toàn cục:** Đơn giản hóa việc truy cập đối tượng dùng chung.

 **Nhược điểm:**
 - **Phá vỡ nguyên tắt SRP (Single Responsibility Principle):** Singleton vừa quản lý công việc khởi tạo, vừa cung cấp logic nghiệp vụ
 - **Khó kiểm tra (Unit Test):** Vì Singleton sử dụng trạng thái toàn cục, việc kiểm tra có thể khó khăn hơn.
 - Có thể gây ra lỗi trong môi trường đa luồng nếu không được triển khai đúng cách
 
---
## **Tài liệu tham khảo:**

1. https://viblo.asia/p/singleton-design-pattern-tro-thu-dac-luc-cua-developers-Qbq5QBkJKD8
2. https://viblo.asia/p/hoc-singleton-pattern-trong-5-phut-4P856goOKY3