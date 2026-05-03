
![](https://logyscal.com/wp-content/uploads/2022/05/Encapsulation3.jpg)

**Tính đóng gói** là quy trình giữ và che giấu một hoặc nhiều thành phần thành với một gói vật lý hoặc logic. Trong phương pháp lập trình hướng đối tượng, nó ngăn việc truy cập vào chi tiết triển khai bên trong đối tượng hay thư viện.

Đóng gói không chỉ ẩn dữ liệu mà còn:
- Đảm bảo dữ liệu ở trạng thái hợp lệ thông qua kiểm tra trong **getter/setter** (c#).
- Cung cấp giao diện an toàn để tương tác với lớp.
- Giảm sự phụ thuộc giữa các thành phần trong hệ thống (loose coupling).

> ***"Tính đóng gói giúp bảo vệ tính toàn vẹn của đối tượng bằng cách ngăn chặn truy cập trực tiếp vào dữ liệu nội bộ. Thay vào đó, dữ liệu chỉ có thể được truy cập hoặc thay đổi thông qua các phương thức hoặc thuộc tính được kiểm soát, đảm bảo logic nghiệp vụ được thực thi đúng."***

==Tính đóng gói được triển khai== bằng cách sử dụng các từ khóa truy cập (**access specifiers**). Một chỉ định truy cập định nghĩa phạm vi và sự ẩn hiện các thành phần bên trong class. Các từ khóa bao gồm: public, private, protected, internal,...

| Access specifiers      | mean                                                                                                                                                                                                                                                                                                                               |
| ---------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **private**            | Chỉ có thể truy cập ở bên trong class                                                                                                                                                                                                                                                                                              |
| **public**             | Chỉ có thể truy cập từ bất cứ đâu                                                                                                                                                                                                                                                                                                  |
| **protected**          | Có thể truy cập bên trong lớp và các lớp con kế thừa (dù ở cùng hay khác Assembly).                                                                                                                                                                                                                                                |
| **internal**           | Giống như public nhưng chỉ hạn chế trong 1 assembly. Hiểu đơn giản là ta có 1 ứng dụng **App** gọi thư viện **Lib** bên ngoài. Trong thư viện **LIb** này có **class C** khai báo **internal**. Các class khác trong thư viện Lib này có thể truy cập **class C** này nhưng ứng dụng **App** không thể là bởi vì khác **Assembly** |
| **protected internal** | Có thể truy cập trong cùng Assembly hoặc trong các lớp con kế thừa (dù ở cùng hay khác Assembly).                                                                                                                                                                                                                                  |
- **Giải thích thêm về Assembly:** Assembly là đơn vị triển khai trong .NET (thường là một file .dll hoặc .exe). Các thành viên internal chỉ có thể truy cập trong cùng Assembly, giúp giới hạn phạm vi truy cập trong một dự án hoặc thư viện.

Ví dụ:
```csharp
public class Bike{
	public int mileage = 65;
	public string color = "Blue";
	private string formula = "a*b";

	//Public - có thể truy xuất từ bên ngoài class
	public int GetMileage(){
		return mileage;
	}
	
	//Public - có thể truy xuất từ bên ngoài class
	public string GetColor(){
		return color;
	}
	
	
	//private - chỉ được phép truy cập bên trong một lớp
	public string GetEngineMakeFomula(){
		return formula;
	}
}

public class Program
{
	public static void Main(string[] args)
	{
		Bike objBike = new Bike();

		Console.WriteLine($"Bike mileage is {objBike.GetMileage()}");
		Console.WriteLine($"Bike Color is {objBike.GetColor()}");
		
	}
}
```

---
### **Phân biệt getter/setter thủ công và auto-implemented properties:**

- C# cung cấp **auto-implemented properties** để đơn giản hóa việc tạo getter/setter khi không cần logic phức tạp.
- **Ví dụ:**
```csharp
public class Person
{
    // Auto-implemented property
    public string Name { get; set; }

    // Property với logic kiểm tra
    private int _age;
    public int Age
    {
        get => _age;
        set
        {
            if (value >= 0 && value <= 150)
                _age = value;
            else
                throw new 
	                ArgumentException("Age must be between 0 and 150.");
        }
    }
}
```

Nói các khác ==tính đóng gói cho phép kiểm soát quyền truy cập== (và thay đổi) giá trị của thuộc tính hoặc quyền gọi phương thức của đối tượng (hoặc lớp) và đối tượng (hoặc lớp) con.

==Một lớp được đóng gói đầy đủ có các hàm getter và setter== được sử dụng để đọc và ghi dữ liệu. Lớp này không cho phép truy cập dữ liệu trực tiếp.

Một thuộc tính định nghĩa 2 phương thức truy cập (acccessor methods):
- Một phương thức `get()`, định nghĩa các mà các trường liên quan có thể được truy cập.
- Một phương thức `set()`, định nghĩa các mà cùng một trường có để được sửa đổi.
Phương thức truy cập mặc định sẽ là `public` Khi không có bộ điều chỉnh truy cập nào được chỉ định.

***Ví dụ:***

```csharp
namespace AccessSpecifiers{
	class Student
	{
		//Creating setter and getter for each property
		public string ID{get; set;}
		public string Name{get; set;}
		public string Email{get; set;}
	}
}
```

```csharp
namespace AccessSpecifiers{
	class Program
	{
		static void Main(string[] args)
		{
			Student student = new Student();

			//Setting values
			student.ID = "B2016947";
			student.Name = "PhamGiaBao";
			student.Email = "PhamGiaBao123@gmail.com";
			
			//Getting values
			Console.WriteLine($"ID: {student.ID}");
			Console.WriteLine($"Name: {student.Name}");
			Console.WriteLine($"Email: {student.Email}");
		}
	}
}
```

---
### **Che giấu thông tin (information Hiding):**

Che giấu thông tin là khái niệm giới hạn trực tiếp. Dữ liệu được truy cập gián tiếp sử dụng cơ chế an toàn, các phương thức trong lập trình hướng đối tượng. 

---
### **Question**

#### 1.**“Bảo vệ khỏi sự truy cập trái phép từ bên ngoài bởi ai?”**

👉 **Trả lời:**
Là bảo vệ khỏi **các phần mã khác bên ngoài lớp** – ví dụ như các hàm, lớp khác **không nên được phép truy cập trực tiếp vào dữ liệu nội bộ của lớp**.

>Tưởng tượng như:  
Bạn có một chiếc két sắt (class), bên trong có tiền (dữ liệu). Bạn không muốn ai đó mở két và lấy tiền trực tiếp, mà phải dùng giao diện như “rút tiền”, “gửi tiền” – chính là getter và setter hoặc các method do bạn cho phép.

##### 2. **“Nếu có getter và setter rồi thì cần gì private nữa? Vì setter vẫn thay đổi được mà?”**

👉 **Trả lời:**
Đúng là setter có thể thay đổi giá trị, **nhưng điểm mấu chốt nằm ở việc bạn kiểm soát được setter sẽ làm gì.**

### 🔐 Mục đích của `private`:

- Không cho người khác **truy cập hoặc thay đổi trực tiếp dữ liệu**.
- Thay vào đó, bạn **định nghĩa rõ cách dữ liệu đó được truy cập hoặc thay đổi** thông qua getter/setter.

**✅ Ví dụ dễ hiểu:**
```csharp
class BankAccount
{
    private decimal _balance;

    public decimal Balance
    {
        get { return _balance; }
        private set 
        {
            if (value >= 0)
                _balance = value;
        }
    }

    public void Deposit(decimal amount)
    {
        if (amount > 0)
            Balance += amount;
    }
}
```

### 🎯 Giải thích:

- `Balance` là `private set`, tức bên ngoài **chỉ được đọc**, không được gán trực tiếp.
- Muốn nạp tiền thì phải đi qua `Deposit()` – tức là **bạn kiểm soát được quá trình thay đổi dữ liệu**.
- Nếu không dùng `private`, người ta có thể viết:
```csharp
acc.Balance = -1000; // Sai logic, âm tiền!
```

⏩ **=> `private` là để ngăn việc truy cập không hợp lệ** và ép buộc người dùng phải sử dụng phương thức mà bạn thiết kế để đảm bảo an toàn logic nghiệp vụ.

#### **3."Tính đóng gói có ảnh hưởng đến hiệu suất không?"**

**Trả lời:** Tính đóng gói không ảnh hưởng đáng kể đến hiệu suất, vì getter/setter thường được biên dịch thành các lệnh truy cập trực tiếp (inline). Tuy nhiên, việc thêm logic kiểm tra trong setter có thể làm tăng nhẹ chi phí tính toán, nhưng điều này cần thiết để đảm bảo tính toàn vẹn dữ liệu.