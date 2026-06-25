## **Factory Pattern là gì?**

**Factory Pattern** là một trong 9 mẫu thiết kế của lập trình hướng đối tượng (OOP). Nó giúp cho chúng ta có thể định nghĩa nhiều đối tượng và cho phép các lớp con tự quyết định là cái nào sẽ được khởi tạo.

**Factory Pattern** là một các thiết kế để tạo ra một thực thể (Object) để tạo ra nhiều thực thể khác. Có thể nói Factory là một lớp (Class) có nhiều phương thức, mỗi phương thức như vậy sẽ tạo ra các thực thể khác nhau dựa trên giá trị mà chúng ta truyền vào.

Trong **Factory Design Pattern**, chúng ta tạo Object mà không để lộ ra cách mà nó tạo ra như thế nào đến phía người dùng. Theo nguyên lý này thì người dùng chỉ cần gọi và nhận được Object mà họ mong muốn, mà không cần quan tâm đến cách nó tạo ra như thế nào.

**Factory Pattern** cung cấp cách tốt nhất để tạo ra các đối tượng. Và quan trọng là nó chia ra 3 loại mô hình đó là:

- Simple factory  
- Factory Method
- Abstract factory

---
## **Factory Method Pattern là gì?**

![](https://images.viblo.asia/fa6170a7-ffd8-44a4-940f-2ccc47849fcd.png)

**Factory Method** là một pattern thuộc nhóm Creational Patterns - và mẫu này nằm trong nhóm Polymorphic Factory (đa hình). Và vì tính trừu tượng của nó, Factory Method còn được gọi là **Virtual Constructor**. **Factory Method** giải quyết vấn đề khởi tạo đối tượng mà không chỉ ra chính xác lớp nào sẽ khởi tạo, ủy quyền cho lớp con.

**Factory Method** cung cấp một interface, phương thức trong việc tạo nên một đối tượng (Object) trong class. Nhưng để cho class con kế thừa nó có thể ghi đè để chỉ rõ đối tượng (Object) nào được tạo. **Factory Method** giao việc khởi tạo đối tượng cụ thể cho lớp con.

**Mục đích:** 
- Tạo ra một cách khởi tạo object mới thông qua một interface chung.
- Che giấu quá trình xử lý logic của phương thức khởi tạo.
- Giảm sự phụ thuộc, dễ dàng mở rộng.
- Giảm khả năng gây lỗi compile.

![](https://images.viblo.asia/6ed7d8a5-7e91-4666-8156-1a0676b2c912.png)

---
## **Mục đích ra đời?**

Giả sử ta có 3 class `Dog`, `Cat`, `Duck` cùng implement interface `IAnimal`.
Khi cần tạo đối tượng `IAnimal` mà chưa biết trước sẽ là con gì (tùy thuộc vào điều kiện cụ thể), thì code thường như sau:
![](https://images.viblo.asia/1ea931d6-4432-4990-ab2a-94ea05b47913.png)

```csharp
IAnimal animal;

if(...){
	animal = new Dog();
}
else if(...){
	animal = new Cat();
}
else if(...){
	animal = new Duck();
}
```

Cách này gây ra: 
- **Trùng lặp logic khởi tạo** nếu cần áp dụng ở nhiều nơi.
- **Khó bảo trì** khi muốn sửa đổi hoặc mở rộng.

![](https://images.viblo.asia/02bc95d1-e578-4cd3-9853-1a35e0dd25e9.png)
##### **$\to$ Giải pháp: Factory Method**

**Factory Method**  gom toàn bộ logic khởi tạo vào một nơi duy nhất -- Giúp mã ngắn gọn, dễ quản lý và hỗ trợ mở rộng tốt hơn.

```csharp
public class AnimalFactory{
	public static IAnimal CreateAnimal(AnimalType type){
		switch(type){
			case AnimalType.Cat: return new Cat();
			case AnimalType.Dog: return new Dog();
			case AnimalType.Duck: return new Duck();
			default: return null;
		}
	}
}
```

*Lợi ích:*
- giảm lặp code.
- Dễ thay đổi/ tùy biến logic khởi tạo.
- Tăng tính đa hình, linh hoạt theo ngữ cảnh sử dụng.
---
## **Kiến trúc**

![](https://images.viblo.asia/87b847da-a31e-47ba-83c5-3b4090d80893.png)

Các thành phần trong mô hình:
- **Product:** Định nghĩa một khuôn mẫu (interface) của các đối tượng mà factory method tạo ra.
- **ConcreteProduct:** Các lớp được cài đặt khuôn mẫu product.
- **Creator:** 
	- Khai báo factory method, trả về kiểu đối tượng thuộc kiểu product. Creator cũng có thể định nghĩa một cài đặt mặc định của factory method mà giá trị trả về là một đối tượng **ConcreteProduct** mặc định.
	- Gọi factory method để tạo đối tượng kiểu product.
- **ConcreteCreator:** Ghi đè factory method để trả về một instace của **ConcreteProduct**.

---
## **Ưu & nhược điểm**

##### **Ưu điểm:**
- Che giấu quá trình xử lý logic của phương thức khởi tạo
- Hạn chế sự phụ thuộc giữa creator và concrete products.
- Dễ dàng mở rộng, thêm những đoạn code mới vào chương trình mà không cần phá vỡ các đối tượng ban đầu.
- Giúp gom các đoạn code tạo ra product vào một nơi trong chương trình, nhờ đó giúp dễ theo dõi và thao tác.
- Giảm khả năng gây lỗi compile, trong trường hợp chúng ta cần tạo một đối tượng mà quên khai báo lớp, chúng ta cũng có thể xử lý lỗi trong Factory và khai báo lớp cho chúng sau.

=> Vì những ưu điểm trên nên **Factory method pattern** thường được sử dụng trong các thư viện (người dùng đạt được mục đích là tạo ra đối tượng mà không cần quan tâm đến các nó được tạo ra như thế nào)

##### **Nhược điểm:**
- Sorce code có thể trở nên phức tạp hơn mức bình thường vì phải đòi hỏi phải sử dụng nhiều class mới có thể cài đặt được pattern này.
- Việc refactoring (tái cấu trúc) một class bình thường có sẵn thành một class có Factory Method có thể dẫn đến nhiều lỗi trong hệ thống, phá vỡ sự tồn tại của client.
- Factory method pattern lệ thuộc vào việc sử dụng private constructor nên các class không thể mở rộng và kế thừa

---
## **Khi nào thì sử dụng?**

Factory method được sử dụng khi:
- Chúng ta có một super class với nhiều class con và dựa trên dữ liệu đầu vào để trả về một class con. Mô hình này chịu trách nhiệm cho việc khởi tạo một lớp tư phía người dùng (client) sang lớp Factory, giúp tiết kiệm tài nguyên hệ thống vì nhờ vào việc tái sử dụng các object đã có thay vì xây dựng lại mỗi phần có thêm product.
- Do là không biết sau này sẽ cần những lớp con nào nữa. Khi cần mở rộng, hãy tạo ra sub class và implement thêm vào factory method cho việc khởi tạo sub class này.

---
## **Code minh họa:**

**Ví dụ: Với bài toán mua bánh Pizza:**
```csharp
//Interface
public interface INotificationService{
	void Send(string message, string recipient);
}

//Concrete implementation
public class EmailService: INotificationService
{
	public void Send(string message, string recipient){
		Console.WriteLine($"Email sent to: {recipient}: {message}");
	}
}

public class SmsService: INotificationService
{
	public void Send(string message, string recipient){
		Console.WriteLine($"SMS sent to: {recipient}: {message}");
	}
}

//Factory
public class NotificationFactory
{
	public static INotificationService CreateNotification(string type){
		return type.ToLower() switch{
			"email" => new EmailService(),
			"sms" => new SmsService(),
			_=> throw new ArgumentException($"Unknown type: {type}")
		};
	}
}

class Program
{
	static void Main(string[])
	{
		//Phải gọi đến static method
		var emailService = NotificationFactory.CreateNotification("email");
		emailService.Send("Hello", "user123@gmail.com");

		var smsService = NotificationFactory.CreateNotification("sms");
		smsService.Send("Hello", "0123456987");
	}
}
```

Ví dụ bài toán thông báo (Factory method pattern + DI)
```csharp
public interface INotificationService
{
    Task Send(string message, string recipient);
    string NotificationType { get; }
}


#region ==== Concrete implementation =====
public class EmailService: INotificationService
{
    public string NotificationType => "Email";

    public async Task Send(string message, string recipient)
    {
        await Task.Delay(500); // Simulate async operation
        Console.WriteLine($"Email sent to: {recipient}: {message}");
    }
}

public class SmsService : INotificationService
{

    public string NotificationType => "SMS";
  
    public async Task Send(string message, string recipient)
    {
        await Task.Delay(500); // Simulate async operation
        Console.WriteLine($"SMS sent to: {recipient}: {message}");
    }
}
#endregion

#region  ===== Factory với DI =====
public interface INotificationFactory
{
    INotificationService CreateNotification(string type);
    IEnumerable<string> GetAvailableTypes();
}

public class NotificationFactory : INotificationFactory
{

    public readonly IServiceProvider _serviceProvider;
    public NotificationFactory(IServiceProvider serviceProvider)
    {
        _serviceProvider = serviceProvider;
    }

    public IEnumerable<string> GetAvailableTypes()
        => new List<string> { "email", "sms" };

    public INotificationService CreateNotification(string type)
    {
        return type.ToLower() switch
        {
            "email" => _serviceProvider.GetRequiredService<EmailService>(),
            "sms" => _serviceProvider.GetRequiredService<SmsService>(),
            _ => throw new ArgumentException($"Unknown type: {type}")
        };
    }
}
#endregion

  

#region  ==== Service sử dụng Factory =====
public interface INotificationManager
{
    Task SendNotificationAsync(string type, string message, string recipient);
    Task SendToAllChannelsAsync(string message, string recipient);
}

public class NotificationManager : INotificationManager
{
    private readonly INotificationFactory _factory;
    public NotificationManager(INotificationFactory factory)
    {
        _factory = factory;
    }


    public async Task SendNotificationAsync(string type, string message, string recipient)
    {
        try
        {
            var service = _factory.CreateNotification(type);
            await service.Send(message, recipient);
        }
        catch (Exception ex)
        {
            Console.WriteLine($"Error sending notification: {ex.Message}");
        }
    }

    public async Task SendToAllChannelsAsync(string message, string recipient)
    {
        var availableTyoe = _factory.GetAvailableTypes();
        var task = availableTyoe.Select(type =>
            SendNotificationAsync(type, message, recipient)
        );
        await Task.WhenAll(task);
    }
}
#endregion

#region  ==== Dependency Injection Setup =====
public static class NotificationServiceExtensions
{
    public static IServiceCollection AddNotificationService(this IServiceCollection services)
    {

        //Đăng ký Concrete services
        services.AddTransient<EmailService>();
        services.AddTransient<SmsService>();

        //Đăng ký Factory
        services.AddScoped<INotificationFactory, NotificationFactory>();

        //Đăng ký Service sử dụng Factory
        services.AddScoped<INotificationManager, NotificationManager>();

        return services;
    }
}
#endregion
```

---
