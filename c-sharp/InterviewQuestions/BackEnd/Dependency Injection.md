
## **1. Khái niệm:

Trước khi tìm hiểu về **Dependency Injection (DI)**, thì cần phải nắm gõ về hai khái niệm nền tảng: **Dependency Inversion** và **Inversion of Control Pattern (IoC)**. Đây là 2 khái niệm cốt lõi trong lập trình hướng đối tượng và thiết kế phần mềm giúp hiểu gõ về tại sao và cách DI được áp dụng trong thực tế.
##### **Dependency Inversion:

![](https://images.viblo.asia/f32fc786-388f-4f8d-b08d-fe5c30c2729a.jpg)
![]()

**Định nghĩa:**
DIP là một trong 5 [nguyên tắc SOLID](https://toidicodedao.com/2015/03/24/solid-la-gi-ap-dung-cac-nguyen-ly-solid-de-tro-thanh-lap-trinh-vien-code-cung/) và được định nghĩa như sau:
- Các module cấp cao (high-level modules) không nên quá phụ thuộc vào các module cấp thấp (low-level modules). Cả hai phụ thuộc vào abstraction (trừu tượng)
- Abstraction không nên phụ thuộc vào chi tiết. Ngược lại, chi tiết nên phụ thuộc vào abstraction.

**Mục tiêu chính:**
- Giảm sự phụ thuộc chặc chẽ (tight coupling) giữa lớp/module
- Tăng tính linh hoạt và khả năng mở rộng.

**Ví dụ:**

- Không tuân theo DIP:
```csharp
public class Engine{
	public void Start(){
		Console.WriteLine("Engine started");
	}
}

public class Car{
	public Engine _engnine = new Engine(); //Phụ thuộc vào 1 lớp cụ thể
	public void Start(){
		_engnine.Start();
	}
}
```

- Tuân theo DIP:
```csharp
public interface IEngine{
	void Start(){}
}

public class GasEngine: IEngine{
	public void Start(){
		Console.WriteLine("Gas engine started");
	}
}

public class ElectricEngine: IEngine{
	public void Start(){
		Console.WriteLine("Electric engine started");
	}
}

public class car{
	private readonly IEngine _engine;
	public Car(IEngine engine) => _engine = engine; //phụ thuộc vào abstraction

	public void Start(){
		_engine.Start();
	}
}
```

##### **Inversion of Control (IoC):

![](https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ2mJXUUhR-lhvGi7gyfJIFB_iEVGtR4xcz-Q&s)

**Định nghĩa:**
- IoC là một **mẫu thiết kế (design patten)**, trong đó **việc điều kiển (control)** luồng của chương trình hoặc việc khởi tạo và quản lý các phụ thuộc (dependency) không được thực hiện bởi chính lớp đó, mà được **bên thứ 3 quản lý nó.**

**Mục tiêu chính:**
- Dịch chuyển trách nhiệm khởi tạo đối tượng va quản lý vòng đời của chúng sang một **IoC container hoặc framework.**

**Các phương pháp IoC phổ biến:**
1. **Dependency Injection (phổ biến nhất).**
2. **Service Locator (ít được khuyến kích hơn)**
3. **Event-based IoC.**

**Ví dụ IoC trong DI:**
```csharp
public class Program{
	static void Main(string[] args){
		IEngine engine = new ElectricEngine(); //IoC container sẽ dảm nhận công việc này
		Car car = new Car(engine);
		car.Start();
	}
}
```

Sử dụng IoC container: Một IoC container như ``Microsft.Extensions.DependencyInjection`` có thể tự động khởi tạo và cung cấp ``IEngine``:
```csharp
public void ConfigureServices(IServiceCollection services){
	services.AddSingleton<IEngine, Engine>();
	services.AddTransient<Car>();
}
```

##### **Dependency Injection (DI):**

![](https://images.viblo.asia/226c1711-7a88-45ba-b705-9f197d3a71dd.png)

**Định nghĩa:**
- **Dependency Injection** là một thiết kế trong lập trình giúp giải quyết vấn đề **phụ thuộc giữa các lớp**. Nó cho phép **"inject" (cung cấp)** các phụ thuộc (dependency) vào một lớp thay vì để lớp đó tự tạo hoặc quản lý chúng.

Mục tiêu chính:
- **Giảm độ kết dính (tight coupling).**
- Các module **không giao tiếp trực tiếp với nhau**, mà phải **thông qua interface**. Module cấp thấp sẽ implement interface, module cấp cao sẽ gọi module cấp thấp thông quan interface.
- **Tăng tính linh hoạt và khả năng mở rộng (extensibility).**
- Dễ dàng **kiểm thử** với các mock dependencies.
- Ví dụ: Để giao tiếp với database, thì sẽ có interface IDatabase, các module cấp thấp XMLDatabase, SQLDatabase. Module cấp cao là CustomerBusiness sẽ chỉ sử dụng interface IDatabase.
- Công việc khởi tạo các module cấp thấp sẽ do **DI Container thực hiện**. Ví dụ trong module CustomerBusiness, sẽ không khởi tạo ``IDatabase db = new XMLDatabase();`` ,**Việc này sẽ do DI Container thực hiện**. Module CustomersBusiness sẽ không biết gì về module XMLDatabase hay SQLDatabase.
- Việc Module nào gắn vào với Interface nào sẽ được config trong code.
- DI được dùng để làm giảm sự phụ thuộc giữa các module, dễ dàng hơn trong việc thay đổi module, bảo trì code và testing.


---
## **2. Các loại Dependency Injection:

- **Construction injection:** Các dependency (biến phụ thuộc) được cung cấp thông qua constructor(hàm tạo lớp).
	```csharp
	public class Car{
		private readonly IEngine _engine;
		public Car(IEngine engine){
			_engine = engine;
		}
	}
	```
	- **Ưu điểm:** Phổ biến nhất, đảm bảo đối tượng luôn được cung cấp phụ thuộc.

- **Property Injection**: Inject qua thuộc tính (properties).
	```csharp
	public class Car{
		public IEngine Engine {set; get;}
	}
	```
	- **Hạn chế** không đảm bảo phụ thuộc luôn được cung cấp.
	
- Method injection: Inject qua tham số của phương thức
	```csharp
	public void StartEngine(IEngine engine){
		engine.Start();
	}
	```
	- Hạn chế là chỉ áp dụng trong phạm vi phương thức
	
- **Setter injection:** Các dependency sẽ được truyền vào 1 class thông qua setter method (hàm setter).
	```csharp
	public class ServiceA{
		private DaoB daoB;
		
		//setter ịnection
		public void setDaoB(DaoB _daoB){
			this.daoB = _daoB;
		}
	}
	```
	- Setter injection linh hoạt hơn Constructor injection vì cho phép thay đổi dependencies sau khi đối tượng đã được tạo.
	
- **Interface injection:** Dependency sẽ cung cấp một Interface, trong đó có chứa hàm Inject. Các Client phải triển khai một Interface mà có một setter method dành cho việc nhận dependency và truyền nó vào class thông qua việc gọi hàm Inject của Interface đó.


---
## **3. Ưu ,nhược điểm của Denpendency Injecion:

**Ưu điểm:**
- Giảm sự kết dính giữa các module
- Code dễ bảo trì, dễ thay thế module
- Dễ test và Unit Test.
- Dễ dàng thấy quan hệ giữa các module (vì các dependency điều được inject vào constructor).

**Nhược điểm:**
- Về khái niệm "Dependency" rất phức tạp khi học.
- Sử dụng interface nên đôi khi khó debug, do khó biết chính xác module nào được gọi
- Làm tăng độ phức tạp của code.
- Các object được khởi tạo toàn bộ ngay từ đầu, có thể làm giảm performance.


---
## **4. Sử dụng DI trong .NET CORE

Sử dụng Dependency Injection thông qua các bước sau:

- Sử dụng một interface hoặc base class để trừu tượng hóa việc triển khai phụ thuộc.
	*Ví dụ:*
	**IMyDependency interface để xác định phương thức Write:**
	```
	public interface IMyDependency{
		void WriteMessage(string text);
	}
	```

	Interface này được MyDependency triển khai:
	```
	public class MyDependency: IMyDependency{
		public void WriteMessage(string text){
			Console.WriteLine($"Text: {text}");
		}
	}
	```

- Đăng ký phần phụ thuộc trong service container. [ASP.NET Core](http://asp.net/) cho phép chúng ta có thể đăng ký các dịch vụ ứng dụng của mình với IoC container, trương phương thức ConfigureServices của lớp Startup. Phương thức  ConfigureServices bao gồm một tham số kiểu IServiceCollection. Được sử dụng để đăng ký các dịch vụ ứng dụng.
	*Ví dụ:*
	**Phương thức AddScope đăng ký service với scoped lifetime, lifetime của một singleton request**
	```
	using DependencyInjectionSample.Interfaces;
	using DependencyInjectionSmaple.Services;

	var builder = WebApplication.CreateBuilder(args);
	builder.Services.AddRazorPages();
	builder.Services.AddScoped<IMyDependency, MyDependency>();
	var app = builder.Build();
	```

- Đưa service vào phương thức khởi tạo của lớp mà nó được sử dụng. Framework sẽ tạo thể hiện của sự phụ thuộc và loại bỏ khi nó không cần thiết nữa
	*Ví dụ:*
	**IMyDependency service được request và sử dụng để gọi phuong thức Write**
	```
	public class Index2Model: PageModel{
		private readonly IMyDependency _myDependency;
		public Index2Model(IMyDependency myDependency){
			_myDependency = myDependency;
		}

		public void OnGet(){
			_myDependency.WriteMessage("Thank for watching!");
		}
	}
	```

Bằng các sử dụng DI pattern, service sẽ: không sử dụng MyDependency, chỉ sử dụng IMyDependency interface thực hiện nó. Điều này giúp cho dễ dàng thay đổi việc thực thi của Controller mà không cần sửa đổi Controller. Không tạo ra instance của MessageWrite, bởi vì nó được tạo bởi DI Container.


---
## **5. Các loại Service Lifetime khi đăng ký DI:

Bất cứ khi nào yêu cầu Service, DI Container sẽ quyết định xem là có nên tạo ra instance mới hay sử dụng lại instance đã tạo ra trước đó. ***Vòng đời của Service sẽ phụ thuộc vào khi khởi tạo instance và nó tồn tại bao lâu***. Có 3 mức độ vòng đời: addTransient, addScoped, addSingleton.

- **Transient:** Instance được khởi tạo mỗi lần tạo service.
- **Scope**: Instance được khởi tạo mỗi scope. (Scope ở đây chính là mỗi lần request gửi đến ứng dụng). Trong cùng một scope thì service sẽ được tái sử dụng
- **Singleton**: Instance của service sẽ được tạo duy nhất từ lúc khởi chạy ứng dụng và được dùng ở mọi nơi.


---
## **Tài liệu tham khảo:
[1]. [6 Phút Để Hiểu Rĩ về Dependency Injection](https://codelearn.io/sharing/hieu-ro-ve-dependency-injection)

[2]. [Giới thiệu về Inversion of Control và Dependency Injection](https://shareprogramming.net/gioi-thieu-inversion-of-control-va-dependency-injection-trong-spring/)

[3]. [Document Dependency Injection Microsoft .NET CORE](https://docs.microsoft.com/en-us/aspnet/core/fundamentals/dependency-injection?view=aspnetcore-5.0)

[4]. [# Dependency Injection là gì? Tìm hiểu về DI trong lập trình](https://bkhost.vn/blog/dependency-injection-la-gi/)
