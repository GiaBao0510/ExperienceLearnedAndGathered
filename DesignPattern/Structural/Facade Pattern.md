### **Giới thiệu:**

![](https://refactoring.guru/images/patterns/content/facade/facade.png?id=1f4be17305b6316fbd548edf1937ac3b)
- **Facade** là một mẫu thiết kế thuộc nhóm cấu trúc (**Structural pattern**).
- **Facade Pattern** cung cấp cho chúng ta có một ==giao diện đơn giản thay cho một nhóm các giao diện trong một hệ thống con== (**subsystem**). **Facade** định nghĩa một giao diện ở cấp độ cao hơn để giúp người dùng có thể dễ dàng sử dụng hệ thống con này.
- **Facade Pattern** cho phép các đối tượng có thể trực tiếp truy cập trực tiếp giao diện chung này để giao tiếp với các giao diện trong hệ thống con khác. Mục tiêu là ==che giấu đi các hoạt động phức tạp bên trong hệ thống con== ,làm cho hệ thống con dễ sử dụng hơn.
- **Facade Design pattern** xác định higher-level instance giúp hệ thống con dễ sử dụng hơn.
- **Tần suất sử dụng:** khá cao.

---
### **Mục đích ra đời:**

#### **Vấn đề gặp phải (Problem):**

Trong một ứng dụng phức tạp, có nhiều hành động cần được thực hiện theo một thứ tự nhất định. Những hành động này thường được sử dụng ở nhiều nơi khác nhau trong hệ thống. Ví dụ:

- Khi làm việc với một thư viện bên thứ ba phức tạp, chúng ta cần khởi tạo nhiều đối tượng, theo dõi trạng thái của chúng và đảm bảo thực hiện đúng thứ tự logic.
- Khi thao tác với một tập hợp lớn các đối tượng, việc lặp lại quy trình khởi tạo và quản lý trạng thái trở nên rườm rà.

Điều này dẫn đến tình trạng:
- Phải viết lại cùng một đoạn mã nhiều lần ở các vị trí khác nhau trong ứng dụng.
- Nếu có thay đổi về quy trình, chúng ta phải chỉnh sửa tất cả các đoạn mã liên quan, rất mất thời gian và dễ gây lỗi.

Hệ quả:
- Mã nguồn trở nên khó bảo trì.
- Khả năng xuất hiện lỗi cao khi có sự thay đổi.
- Logic nghiệp vụ bị gắn chặt với các thành phần bên thứ ba, làm giảm tính linh hoạt của hệ thống.

![](https://images.viblo.asia/ac5b8b8c-0aad-46c5-94ae-5400c1990d9e.png)
#### **Giải pháp (Solution):**

Để giải quyết vấn đề này, chúng ta có thể sử dụng **Facade Pattern**. Facade đóng vai trò như một lớp trung gian, cung cấp các phương thức đơn giản hóa thao tác với hệ thống phức tạp bên trong.

Cách hoạt động:
- Thay vì gọi trực tiếp từng thành phần bên trong, ta chỉ cần gọi một phương thức trong Facade.
- Facade sẽ xử lý toàn bộ các bước cần thiết và trả về kết quả.
- Nếu cần thay đổi quy trình, ta chỉ cần sửa đổi Facade mà không ảnh hưởng đến các phần còn lại của ứng dụng.

Lợi ích:
- Giảm sự phụ thuộc của các thành phần trong hệ thống vào chi tiết cài đặt bên trong.
- Dễ dàng bảo trì và mở rộng.
- Giúp mã nguồn gọn gàng, dễ đọc hơn.

![](https://images.viblo.asia/51de245b-2a3a-43d8-b88f-f68e847e032c.png)

---
### **Kiến trúc:**

![](https://images.viblo.asia/e0ce5777-c04e-4ac4-8303-c44d0901168d.png)

Các thành phần trong mô hình:
- ***Facade:*** Nắm rõ được hệ thống con nào đảm bảo nhiệm việc đáp ứng yêu cầu của client, nó sẽ chuyển yêu cầu của client đến các đối tượng trong hệ thống con tương ứng. 
- ***Additional Facade:*** Có thể được tạo để tránh việc làm phức tạp một **facade**. Có thể sử dụng bởi client và **facade**.
- ***Complex Sybsystem:*** Bao gồm nhiều object khác nhau, được cài đặt các chức năng của hệ thống con, xử lý công việc được gọi bởi **Facade**. Các lớp này không cần biết **Facade** và không tham chiếu đến nó.
- ***Client:*** Đối tượng sử dụng Facade để tương tác với các **subsystem** thay vì gọi **subsystem** trực tiếp.

Các đối tượng **Facade** thường là **Singleton** bởi vì chỉ cần duy nhất một đối tượng **Facade**.

---
### **Ưu/nhược điểm:**

**Ưu điểm:**
- Ta có thể thực hiện tách mã nguồn của mình ra khỏi sự phức tạp của hệ thống con.
- Hệ thống tích hợp thông qua Facade sẽ đơn giản hơn vì chỉ cần tương tác với Facade thay vì hàng loạt đối tượng khác.
- Tăng khả năng độc lập và khả chuyển, giảm sự phụ thuộc.
- Có thể đóng gói nhiều hàm được thiết kế không tốt bằng 1 hàm có sự thiết kế tốt hơn.

**Nhược điểm:**
- Class Facade của bạn có thể trở nên quá lớn, làm quá nhiều nhiệm vụ với nhiều hàm chức năng trong nó.
- Dễ bị phá vỡ các quy tắc trong **SOLID**.
- Việc sử dụng Facade cho hệ thống đơn giản, không quá phức tạp trở nên dư thừa.

---
### **Khi nào nên sử dụng:**

Dưới đây chúng ta sẽ liệt kê ra một số trường hợp mà khi gặp sẽ phải cân nhắc sử dụng **Facade pattern**:

- *Muốn gom nhóm chức năng lại để Client dễ sử dụng.* khi hệ thống có rất nhiều lớp làm người sử dụng rất có thể khó hiểu được quy trình xử lý của chương trình. Và khi có rất nhiều hệ thống con mà mỗi hệ thống con đó lại có những giao diện riêng lẻ của nó rất khó cho việc sử dụng phối hợp. Khi đó có thể sử dụng **Facade Pattern** để tạo ra một giao diện đơn giản cho người sử dụng hệ thống phức tạp.
- *Giảm sự phụ thuộc.* Khi bạn muốn phân lớp các hệ thống con. Dùng **Facade Pattern** để định nghĩa cổng giao tiếp chung cho mỗi hệ thống con, do đó giúp giảm sự phụ thuộc của các hệ thống con vì các hệ thống này chỉ giao tiếp thông qua các cổng giao diện chung.
- *Tăng khả năng độc lập và khả chuyển.*
- *Khi người dùng phụ thuộc nhiều vào các lớp cài đặt,* Việc áp dụng **Facade Pattern** sẽ tách biệt hệ thống con của người dùng và các hệ thống con khác, do đó tăng khả năng độc lập và khả chuyển của hệ thống con, dễ chuyển nổi và nâng cấp trong tương lai,
- *Đóng gói nhiều chức năng, che giấu thuật toán phức tạp.*
- *Cần một interface không rắc rối mà dễ sử dụng*

***Ví dụ:*** Khi bạn gọi điện đến shop để đặt hàng. Khi đó tổng đài sẽ là Facade của tất cả các dịch vụ và phòng ban của cửa hàng. Hệ thống sẽ cung cấp cho bạn một giao diện đơn giản qua điện thoại để đặt hàng, giao hàng hay nhiều công việc khác nhau.

![](https://images.viblo.asia/25947d05-05cc-429b-9547-815ba7edeee3.png)

---
### **Ví dụ minh họa - C#:**

**Tạo Subsystem:**
```csharp
public class AccountService
{
	public void GetAccout(string email)
	{
		Console.WriteLine("Getting the account of " + email);
	}
}

public class EmailService
{
	public void SendMail(string mailTo)
	{
		Console.WriteLine("Sending an email to " + mailTo);
	}
}

public class PaymentService
{
	public void PaymentByPaypal()
	{
		Console.WriteLine("Payment by Paypal");
	}
	public void PaymentByCreditCard()
	{
		Console.WriteLine("Payment by Credit Card");
	}
	public void PaymentByEBankingAccount()
	{
		Console.WriteLine("Payment by E-banking account");
	}
	public void PaymentByCash()
	{
		Console.WriteLine("Payment by cash");
	}
}

public class ShippingService
{
	public void FreeShipping()
	{
		Console.WriteLine("Free Shipping");
	}

	public void StandardShipping()
	{
		Console.WriteLine("Standard Shipping");
	}

	public void ExpressShipping()
	{
		Console.WriteLine("Express Shipping");
	}
}

public class SmsService
{
	public void sendSMS(string mobilePhone)
	{
		Console.WriteLine("Sending an message to " + mobilePhone);
	}
}
```


**Tạo Facade:**
```csharp
public class ShopFacade{
	
	private static ShopFacade _instance;

	private AccountService accountservice;
	private EmailService emailservice;
	private PaymentService paymentservice;
	private ShippingService shippingservice;
	private SmsService smsservice;

	private ShopFacade(){
		accountservice = new AccountService();
		emailservice = new EmailService();
		paymentservice = new PaymentService();
		shippingservice = new ShippingService();
		smsservice = new SmsService();
	}

	public static ShopFacade getInstance(){
		if(_instance == null);
			_instance = new ShopFacade();
		return _instance;
	}

	public void buyProductByCashWithFreeShipping(string email){
		accountservice.GetAccout(email);
		paymentservice.PaymentByCash();
		shippingservice.FreeShipping();
		emailservice.SendMail(email);
		Console.WriteLine("Done\n");
	}

	public void buyProductByPaypalWithStandardShipping(string email, string mobilePhone){
		accountservice.GetAccout(email);
		paymentservice.PaymentByPaypal();
		shippingservice.StandardShipping();
		emailservice.SendMail(email);
		smsservice.sendSMS(mobilePhone);
		Console.WriteLine("Done\n");
	}
}
```

**Client gọi Facade:**
```csharp
class Client{
	static void Main(string[] args){
		ShopFacade.getInstance()
			.buyProductByCashWithFreeShipping("baob2016947@student.ctu.edu.vn");
		ShopFacade.getInstance()
			.buyProductByPaypalWithStandardShipping(
				"baob2016947@student.ctu.edu.vn",
				"0123456789"
			);
	}
}
```

---
### **Tài liệu tham khảo:**

1. [Facade Design Pattern - Trợ thủ đắc lực của Developers](https://viblo.asia/p/facade-design-pattern-tro-thu-dac-luc-cua-developers-924lJBLNlPM)
2. [Facade Pattern – Đơn giản hóa tất cả](https://topdev.vn/blog/facade-pattern-don-gian-hoa-tat-ca/)
3. [Facade Design Pattern trong C# - Cách triển khai và ví dụ](https://freetuts.net/facade-design-pattern-trong-c-sharp-5624.html)