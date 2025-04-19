### **Constructor:**

![](https://flatcoding.com/wp-content/uploads/2025/03/php-oop-constructor.jpg)

**Constructor** là một phương thức đặt biệt, có cùng tên với lớp như không có kiểu trả về, phương thức này được sử dụng để khởi tạo ra một thể hiện của một lớp.

Một khi đã định nghĩa ra hàm khởi tạo riêng trong **class**, thì **C# Compiler** sẽ không sinh ra hàm tạo mặc định nữa.

**Các điểm quan trong về Constructor**:
- Hàm khởi tạo phái có tên trùng với tên lớp.
- Hàm khởi tạo không có kiểu trả về
- Hàm khởi tạo có hoặc không có danh sách tham số đầu vào. (**Mặc định** thì sẽ không có danh sách tham số đầu vào).
- Từ khóa `this` dùng để chỉ rõ đâu là thành viên của class
- Hàm khởi tạo có thể được nạp chồng (**overload**).
- Nếu không có **constructor** nào được định nghĩa thì **CLR** (Commom Language Runtime) sẽ tự khởi tạo một constructor mặc định.
- Constructor không được khai báo với từ khóa **virtual**.
- Chúng ta không sử dụng ***tham chiếu hoặc con trỏ*** trong **constructor** vì chúng không cấp địa chỉ

**Ví dụ:**
```csharp
public class User{

	public string Id_user {set; get;}
	public string Name {set; get;}

	//Có tham số
	public User(string id_user, string name){
		Id_user = id_user;
		Name = name;
	}

	//Không có tham số
	public User(){
		Id_user = null;
		Name = null;
	}
}
```

---
### **Khởi tạo Object với property:**

Trong C# cũng có cung cấp một cách để khởi tạo object: sử dụng bộ khởi tạo (**object initializer**).  Cú pháp khởi tạo này cùng sử dụng property.

Khi tạo đối tượng của lớp thì phải cần dùng đến từ khoá `new` thì nó sẽ tạo đối tượng và thi hành phương thức khởi tạo tương ứng với tham số phù hợp

***Ví dụ:***
```csharp
public class Book{

	//Backing field
	private int _idBook;
	private string _bookName;
	private string _title;

	public string IdBook { get => _idBook; set => _idBook = value;}
	public string BookName { get => _bookName; set => _bookName = value;}
	public string Title { get => _title; set => _title = value;}

	//Constructor
	public Book(){
		_idBook = 0;
		_bookName = "null";
		_title = "null";
	}
	
	public Book(int idBook, string bookName, string title){
		_idBook = idBook;
		_bookName = bookName;
		_title = title;
	}
}

internal class Program{
	public static void Main(string[] args){
		
		var book1 = new Book{
			IdBook = 5,
			BookName = "Conan",
			Title = "Thám tử lừng danh"
		};

		var Book2 = new Book(2, "Doremon", "Chú mèo máy");
	}
}
```

---
### **Phương thức khởi tạo riêng tư**

Khi xây dựng một số lớp đặc biệt, ==nhất là những lớp tiện ích== chỉ chứa thành viên tĩnh, bạn mong muốn chỉ ra một cách tường minh lớp này không được phép tạo đối tượng, lớp này không được phép kế thừa thì hãy chọn phương thức khởi tạo không tham số và chỉ ra trạng thái truy cập là `private`

```csharp
class MyLib{

	public static double PI = 3.14;
	
	private MyLib(){}
}
```

Sẽ gặp lỗi nếu như:
- Khởi tạo đối tượng.
- Nếu dùng làm lớp cơ sở cũng gặp lỗi.

---
### **Phương thức khởi tạo tĩnh:**

Bạn có thể xây dựng một phương thức khởi tạo không có tham số, chỉ có phạm vi truy cập là `static`, phương thức khởi tạo này dùng để khởi tạo các thành viên dữ liệu tĩnh, nó tự dộng gọi khi truy cập một thành viên dữ liệu tĩnh lần đầu. 

**Constructor** tĩnh được gọi một lần duy nhất khi lớp được truy cập lần đầu tiên, thường dùng để khởi tạo dữ liệu tĩnh hoặc thiết lập trạng thái ban đầu.

**Constructor** tĩnh ==không thể có tham số== và không được gọi trực tiếp bỏi lập trình viên

**Ví dụ:**

```csharp
class MyColorCode{

	public static string color_primary;
	public static string color_success;
	public static string color_warning;
	public static string color_info;
	public static string color_danger;

	//Phương thức khởi tạo tĩnh, được gọi khi lần đầu truy cập một thành viên tĩnh
	static MyColorCode(){
		color_primary = "Navy";
		color_danger = "red";
		color_success = "green";
		color_warning = "yellow";
		color_info = "white";
	}

}
```

**Ví dụ:**
```csharp
public class Configuration{
	public string ConnectionString {get; private set;}
	public string int MaxRetries {get; private set; }

	//Consstructor tĩnh
	static Configuration(){
		ConnectionString = "Server=localhost;Database=MyApp;Trusted_Connection=True;";
		MaxRetries = 3;
	}
}

class Program{
	static void Main(){

		//Gọi Constructor tĩnh lần đầu
		Console.WriteLine(Configuration.ConnectionString); 
	}
}
```

---
#### **`base` và `this`**

**constructor** có thể gọi constructor khác trong cùng lớp (dùng `this`) hoặc **constructor** của lớp cha (dùng `base`)

```csharp
public class Employee
{
    public string Id { get; set; }
    public string Name { get; set; }
    public decimal Salary { get; set; }

    // Constructor không tham số
    public Employee() : this("Unknown", "Unknown", 0)
    {
    }

    // Constructor có tham số
    public Employee(string id, string name, decimal salary)
    {
        Id = id;
        Name = name;
        Salary = salary;
    }
}

public class Manager : Employee
{
    public string Department { get; set; }

    public Manager(string id, string name, decimal salary, string department)
        : base(id, name, salary)
    {
        Department = department;
    }
}

class Program
{
    static void Main()
    {
        var manager = new Manager("M001", "Alice", 100000, "IT");
        Console.WriteLine($"{manager.Name} manages {manager.Department}");
    }
}
```

---
### **Hàm hủy (Destructor):**

Vì trình dọn rác GC (Garbage Collection) sẽ tự động làm sạch hệ thống, ==nó sẽ hủy các đối tượng lâu không sử dụng== nhưng có nhiều khi chúng ta cần làm một số thứ bằng tay. Trong trường hợp này chúng ta có thể sử dụng một Destructor để hủy các đối tượng không còn sử dụng.

Một hàm hủy là một phương thức được gọi khi đối tượng hủy, có thể sử dụng để giải phóng tài nguyên được sử dụng bởi đối tượng. Hàm huy không giống như phương thức khác

```csharp
public class Bike{

	//Constructor
	public Bike(){
	
	}
	
	//Destructor
	~Bike(){
	
	}
}
```

