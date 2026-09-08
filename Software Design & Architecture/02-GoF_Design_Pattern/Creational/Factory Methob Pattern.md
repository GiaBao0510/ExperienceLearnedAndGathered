## **Factory Pattern là gì?**

**Factory Pattern** là một trong 9 mẫu thiết kế của lập trình hướng đối tượng (OOP). Nó giúp cho chúng ta có thể định nghĩa nhiều đối tượng và cho phép các lớp con tự quyết định là cái nào sẽ được khởi tạo.

**Factory Pattern** là một các thiết kế để tạo ra một thực thể (Object) để tạo ra nhiều thực thể khác. Có thể nói Factory là một lớp (Class) có nhiều phương thức, mỗi phương thức như vậy sẽ tạo ra các thực thể khác nhau dựa trên giá trị mà chúng ta truyền vào.

Trong **Factory Design Pattern**, chúng ta tạo Object mà không để lộ ra cách mà nó tạo ra như thế nào đến phía người dùng. Theo nguyên lý này thì người dùng chỉ cần gọi và nhận được Object mà họ mong muốn, mà không cần quan tâm đến cách nó tạo ra như thế nào.

**Factory Pattern** cung cấp cách tốt nhất để tạo ra các đối tượng. Và quan trọng là nó chia ra 3 loại mô hình đó là:

- Simple factory  
- Factory Method
- Abstract factory
---
## Problem

Hãy tưởng tượng bạn đang xây dựng một ứng dụng quản lý logistics. Phiên bản đầu tiên của ứng dụng chỉ hỗ trợ vận tải bằng xe tải, vì vậy phần lớn mã nguồn nằm bên trong lớp `Truck` (Xe tải).

Sau một thời gian, ứng dụng trở nên khá phổ biến. Mỗi ngày, bạn nhận được hàng chục yêu cầu từ các công ty vận tải biển về việc tích hợp dịch vụ logistics đường biển vào ứng dụng.

![](https://refactoring.guru/images/patterns/diagrams/factory-method/problem1-en-1.5x.png)

Tin tuyệt vời phải không? Nhưng còn phần mã nguồn thì sao? Hiện tại, phần lớn mã nguồn của bạn đang bị ràng buộc chặt chẽ với lớp `Truck`. Việc thêm `Ship` vào ứng dụng sẽ đòi hỏi phải thay đổi toàn bộ hệ thống mã nguồn. Hơn nữa, nếu sau này bạn quyết định thêm một loại hình vận chuyển khác vào ứng dụng, có lẽ bạn sẽ lại phải thực hiện tất cả những thay đổi đó một lần nữa.

Kết quả là bạn sẽ có một đoạn mã rất tệ, chằng chịt các câu lệnh điều kiện để thay đổi hành vi của ứng dụng tùy thuộc vào kiểu của các đối tượng vận chuyển.

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

For example, both `Truck` and `Ship` classes should implement the `Transport` interface, which declares a method called `deliver`. Each class implements this method differently: trucks deliver cargo by land, ships deliver cargo by sea. The factory method in the `RoadLogistics` class returns truck objects, whereas the factory method in the `SeaLogistics` class returns ships.

![](https://refactoring.guru/images/patterns/diagrams/factory-method/solution3-en-1.5x.png)

Đoạn mã sử dụng phương thức factory (thường được gọi là mã khách - client code) không nhận thấy sự khác biệt giữa các sản phẩm thực tế do các lớp con khác nhau trả về. Mã khách coi tất cả các sản phẩm này đều là kiểu trừu tượng `Transport`. Mã khách biết rằng mọi đối tượng vận chuyển đều có phương thức `deliver`, nhưng cơ chế hoạt động cụ thể của phương thức đó không quan trọng đối với mã khách.


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

1. Product (Sản phẩm) định nghĩa giao diện chung cho tất cả các đối tượng có thể được tạo ra bởi Creator (Người tạo) và các lớp con của nó.

2. Concrete Products (Các sản phẩm cụ thể) là những cách triển khai khác nhau của giao diện sản phẩm này.

3. Lớp Creator khai báo phương thức factory (factory method) trả về các đối tượng sản phẩm mới. Điều quan trọng là kiểu dữ liệu trả về của phương thức này phải khớp với giao diện sản phẩm.

4. Bạn có thể khai báo phương thức factory dưới dạng trừu tượng (abstract) để buộc tất cả các lớp con phải tự triển khai phiên bản riêng của phương thức đó. Một cách khác là để phương thức factory cơ sở trả về một loại sản phẩm mặc định.

5. Cần lưu ý rằng, bất chấp tên gọi của nó, việc tạo ra sản phẩm không phải là trách nhiệm chính của Creator. Thông thường, lớp Creator đã chứa sẵn một số logic nghiệp vụ cốt lõi liên quan đến sản phẩm. Phương thức factory giúp tách biệt logic này khỏi các lớp sản phẩm cụ thể. Hãy hình dung ví dụ sau: một công ty phát triển phần mềm lớn có thể có bộ phận đào tạo lập trình viên. Tuy nhiên, chức năng chính của toàn bộ công ty vẫn là viết mã nguồn chứ không phải là sản xuất ra các lập trình viên.

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
