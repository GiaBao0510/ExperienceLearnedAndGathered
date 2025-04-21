
![](https://logicmojo.com/assets/dist/new_pages/images/polymorphism_example.png)

Từ đa hình (polymorphism) nghĩa là có nhiều khuôn mẫu. Tổng quan, đa hình xuất hiện khi một danh sách class chúng có liên quan đến nhau thông quan kế thừa.

> *Hãy lấy ví dụ về Bike, một Bike có thể có 2 khuôn mẫu như xe đề và xe đạp nổ. Chúng ta có thể quyết định phương thức vận hành nào muốn sử dụng để khởi động (nghĩa là lúc chạy).*

**Có hai kiểu đa hình:** Đa hình ở thời điểm biên dịch và Đa hình ở thời điểm thực thi. 

![](https://s3-sgn09.fptcloud.com/codelearnstorage/Upload/Blog/overiding-va-overloading-63739896651.9348.jpg)

---
#### **Đa hình ở thời điểm biên dịch/ Đa hình tĩnh (Compile time polymorphism):**

Đây là kiểu đa hình mà trình biên dịch sẽ nhận dạng khuôn mẫu nào có thể thực hiện ở thời điểm biên dịch gọi là **compile time polymorphism** hay còn gọi là **early binding.** Ví dụ chúng ta gọi là **Nạp chồng phương thức (Methob Overloading)** và **nạp chồng toán tử (Operator Overloading)**

**Method Overloading** có nghĩa là có nhiều hơn một phương thức cùng tên, nhưng **==khác kiểu dữ liệu tham số, khác thứ tự tham số và khác số lượng tham số==** trong cùng hoặc khác class
- **Lợi ích:** thực thi nhanh bởi vì tất cả các phương thức được nhận dạng ở thời điểm biên dịch.
- **Điểm yếu:** Không mềm dẻo.

***Ví dụ - Methob Overloading:***
```csharp
public class Product{
	public string prod_id {get; set;}
	public string prod_name {get; set;}
	public DateTime prod_expiry {get; set;}

	public Product(){
		prod_id = "null";
		prod_name = "null";
	}
	
	public Product(){
		prod_id = "null";
		prod_name = "null";
		prod_expiry = DateTime.UtcNow;
	}

	public Product(string Prod_id, string Prod_name, DateTime Prod_expiry){
		prod_id = Prod_id;
		prod_name = Prod_name;
		prod_expiry = Prod_expiry;
	}
}

public class Program{
	public static void Main(string[] args)
	{
		var product1 = new Product();
		var product1 = new Product("1", "Bàn chải", DateTime.UtcNow);
	}
}
```

***Ví dụ - Methob Overloading bằng các thay đổi kiểu dữ liệu của tham số:***
```csharp
public class Cal{

	public static int Sum(int a, int b)
		=> a + b;
		
	public static float Sum(float a, float b)
		=> a + b;
}

public class Program{
	public static void Main(string[] args)
	{
		Console.WriteLine(Cal.Sum(5, 7));
		Console.WriteLine(Cal.Sum(14.7f, 25.3f));
	}
}
```

**⚠️Chú ý:** Khi nạp chống phương thức, ==một quy tắc tuân theo là nạp chồng phương thức phải so sánh số lượng đối số hoặc kiểu dữ liệu của ít nhất một đối số.== Chúng ta có thể xem Method Overriding như là một ví dụ của việc đa hình ở thời điểm biên dịch được gọi trực tiếp bởi đối tượng được khởi tạo.

---
#### **Đa hình ở thời điểm thực thi/ Đa hình động (Runtime polymorphism):**

![](https://cdn.educba.com/academy/wp-content/uploads/2019/10/Overriding-in-OOPs.png)

Trong kiểu đa hình này, ==trình biên dịch sẽ nhận dạng khuôn mẫu đa hình nào sẽ được thực thi ở thời điểm thực thi, nhưng không phải ở thời điểm biên dịch== được gọi là đa hình ở thời điểm thực thi hoặc **late binding**. Ví dụ early binding là **Method Overriding**.

**Method Overriding** nghĩa là có hai phương thức cùng tên và cùng **signature**, một phương thức trong **class Base**, phương thức còn lại trong **class Child**. Nó được yêu cầu thay đổi hành vi của phương thức trong **class Base** khi **class Child** sử dụng.
- **Lợi ích:** Nó có thể mềm dẻo để điều chỉnh các kiểu đối tượng ở thời điểm thực thi.
- **Hạn chế:** Thực thi sẽ chậm hơn vì phải mất thời gian lấy thông tin về phương thức ở thời điểm chạy.

**Quy tắc sử dụng phương thức ghi đè:**
- Các phương thức static thì không ghi đè nhưng được mô tả lại.
- Các phương thức không kế thừa sẽ không được overriden(hiển nhiên)

Chúng ta cần sử dụng **một trong 2 cách là virtual method và abstract method** ==cho phép lớp dẫn xuất ghi đè một phương thức trong class Base==.

***Ví dụ - Methob Overriding sử dụng phương thức Virtual:*** (Từ khóa `virtual`, `override`)

```csharp
public class Base{
	public virtual string BlogName()
	{
		return "Hello VietNam";
	}
}

public class Child: Base{
	public override string BlogName()
	{
		return "Hello VietNam - 123";
	}
}

public class Program{
	public static void Main(string[] args)
	{
		Base objBase = new Child();
		Console.WriteLine(objBase.BlogName); 
	}
}
```

***Ví dụ - Methob Overriding sử dụng phương thức abstract:*** (Từ khóa `abstract`, `override`)

```csharp
public class Base{
	public abstract string BlogName();
}

public class Child: Base{
	public override string BlogName()
	{
		return "Hello VietNam - 123";
	}
}

public class Program{
	public static void Main(string[] args)
	{
		Base objBase = new Child();
		Console.WriteLine(objBase.BlogName); 
	}
}
```

**⚠️Chú ý:** Cả ghi đè phương thức và nạp chồng phương thức là ác khái niệm khác biệt và rất quan trong của OOP. Nên chú ý kỹ vì tên nó tương tự nhau

---
## ✅ **Bảng so sánh Method Overloading vs Method Overriding**

| Tiêu chí                     | **Overloading (Nạp chồng)**                                              | **Overriding (Ghi đè)**                                                      |
| ---------------------------- | ------------------------------------------------------------------------ | ---------------------------------------------------------------------------- |
| ***Xảy ra khi nào?***        | Thời điểm **biên dịch (Compile time)**                                   | Thời điểm **thực thi (Runtime)**                                             |
| ***Yêu cầu về từ khóa***     | Không cần từ khóa đặc biệt                                               | Dùng `virtual` hoặc `abstract` ở lớp cha, `override` ở lớp con               |
| ***Điều kiện***              | Cùng tên, **khác tham số** (kiểu, số lượng, thứ tự)                      | Cùng tên, **cùng tham số** (signature), trong class con kế thừa từ class cha |
| ***Mục đích***               | Tăng tính linh hoạt khi gọi phương thức với nhiều kiểu dữ liệu khác nhau | Thay đổi hoặc mở rộng hành vi của phương thức từ lớp cha                     |
| ***Liên quan đến kế thừa?*** | ❌ Không bắt buộc có kế thừa                                              | ✔ Bắt buộc phải có kế thừa                                                   |
| ***Có hỗ trợ đa hình?***     | ✔ Đa hình thời gian biên dịch                                            | ✔ Đa hình thời gian thực thi                                                 |
| ***Hiệu suất***              | Nhanh hơn (được xử lý sẵn bởi compiler)                                  | Chậm hơn (quyết định tại runtime – late binding)                             |
| ***Ví dụ từ bài bạn:***      | `Sum(int a, int b)` và `Sum(float a, float b)`                           | `BlogName()` ở class cha và `override BlogName()` ở class con                |

## ✅ **So sánh giữa từ khóa `abstract` và `virtual` trong C#**

|Tiêu chí|`abstract`|`virtual`|
|---|---|---|
|Mục đích|Bắt buộc lớp con **phải ghi đè**|Cho phép lớp con **tùy chọn ghi đè**|
|Có triển khai mặc định?|❌ Không có – không có thân hàm|✔ Có – có thể có sẵn phần thân hàm|
|Được dùng ở đâu?|Chỉ dùng trong **class abstract**|Dùng trong **class thường hoặc class abstract**|
|Lớp con bắt buộc `override`?|✔ Bắt buộc|❌ Không bắt buộc|
|Ví dụ đơn giản|`public abstract void Speak();`|`public virtual void Speak() => Console.WriteLine("Hello");`|