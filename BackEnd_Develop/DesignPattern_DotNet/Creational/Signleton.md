### **Singleton là gì?**

- **Singleton** là một mẫu thiếu kế thuộc nhóm **Creational Patterns**, đảm bảo **một lớp chỉ có một thể hiệm duy nhất (instance)** trong suốt vòng đời của ứng dụng và cung cấp một điểm truy cập toàn cụ đến thể hiện đó. 

![](https://images.viblo.asia/8cc36217-fa29-496b-a2ab-03a5286d8b6b.png)

---

### **Mục đích ra đời**

---

### **Kiến trúc**

![](https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSxsdEhte4FyIrNEWYp-yIWPANF3mwxSV5cXw&s)

---

### **Đặc điểm của Singleton?**

---

### **Đặc điểm của Singleton?**

- **Duy nhất (unique):** chỉ tồn tại một thể hiện duy nhất của lớp.
- **Toàn cục (Global Access):**  có thể truy cập đến thể hiện này ở bất kỳ đâu trong ứng dụng.
- **Kiểm soát việc khởi tạo:** thể hiện chỉ được tạo ra khi cần thiết(Lazy Initialization).
---

### **Nên sử dụng Singleton khi nào?**

- Khi cần một đối tượng duy nhất trong toàn bộ ứng dụng.
- Quản lý cấu hình: Cấu hình ứng dụng cần 1 thể hiện.
- Quản lý kết nối cơ sở dữ liệu: Đảm bảo việc tái sử dụng kết nối và tối ưu tài nguyên.
- Hệ thống ghi log: một logger dùng chung trong toàn bộ chương trình.
---

### **Ví dụ minh họa với c#:**

**A. Triển khai Singleton cơ bản**
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

---

### **Ưu & nhược điểm**

**Ưu điểm:**

 **Nhược điểm:**
 
---
## **Tài liệu tham khảo:**