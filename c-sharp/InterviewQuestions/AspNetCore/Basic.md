### 1. **ASP.NET Core là gì?**

**ASP.NET Core** không phải là một phiên bản nâng cấp của **ASP.NET** .**ASP.NET Core** hoàn toàn được tạo ra để làm việc với **.NET Core framework.**

Nó nhanh hơn, dễ cấu hình chi tiết hơn, ==theo hướng module, dễ mở rộng và hỗ trợ đa nền tảng==. Nó có thể làm việc với cả **.NET Core** và **.NET Framework** thông qua **.NET standard framework**.

**ASP.NET Core** là thích hợp nhất để phát triển các ứng dụng cloud-based như web, mobile, IoT.

---
### 2. **ASP.NET Core cung cấp các tính năng nào?**

1. Hỗ trợ **Dependency Injection**.
2. Hỗ trợ **Logging Framework** và có thể được mở rộng.
3. Giới thiệu web server mới: Kestrel, nó nhanh và hỗ trợ đa nền tảng. Vì vậy, một ứng dụng web có thể chạy nhanh mà không cần **IIS**, **Apache** và **Nginx**.
4. Hỗ trợ nhiều cách **Hosting**.
5. Hỗ trợ hướng module, vì vậy các developer cần import các module mình cần trong ứng dụng. Tuy nhiên, .NET Core framework cũng cung cấp Meta package trong các thư viện.
6. Command-line hỗ trợ tạo, build và chạy ứng dụng.
7. không có **file web.config** .Chúng ta có thể lưu trữ các thông tin cấu hình vào trong f**ile appsettings.json**
8. Không có **file** **Global.asax**. Chúng ta có thể đăng ký và sử dụng các **services** trong **startup class**.
9. Hõ trợ cho lập trình bất đồng bộ.
10. Hỗ trợ **WebSocket** và **SignalR**.
11. Cung cấp bảo vệ chống **CSRF (Cross-Site Request Forgery).**

---
### 3. **Những ưu điểm nào của ASP.NET Core so với ASP.NET?**

1. Hỗ trợ đa nền tảng, có thể chạy được trên Windows, Linux, Mac.
2. Không có Dependency trong cài đặt framework bởi vì tất cả dependency bắt buộc gửi cùng ứng dụng.
3. ASP.NET Core có thể xử lý nhiều request hơn ASP.NET.
4. Nhiều tùy chọn deloy web có sẵn trong ASP.NET Core.

---
### 4. **Metapackage là gì?**

Nó bao gồm tất cả các **package** được hỗ trỡ bỏi mã ASP.NET với các dependency của chúng vào trong 1 package. Nó giúp chúng ta phát triển nhanh vì chúng ta không yêu cầu include các package ASP.NET Core riêng lẻ. **Microsoft.AspNetCore.All** một **Metapackage** được cung cấp bởi ASP.NET Core.


---
### 5. **Ứng dụng ASP.NET Core có thể làm việc với framework .NET 4.x không?**

***Có mà phải thông qua thư viện chuẩn .NET***

---
### 6. **Startup class trong ASP.NET Core là gì?**

**Startup Class** là điểm đầu vào của ứng dụng ASP.NET Core. Mỗi ứng dụng **.Net Core** phải có class này. Nó không phải bắt buộc có tên class là "**Startup**", ta có thể định nghĩa tên nào đó bất kỳ cũng được. Chúng ta ==có thể cấu hình startup class trong Program class.==

```csharp
public class Program{

	public static void Main(string[] args){
		CreateWebHostBuilder(args).Build().Run();
	}

	public static IWebHostBuilder CreateWebHostBuilder(string[] args) 
		=> WebHost.CreateDefaultBuilder(args)
			.UseStartup<TestClass>();
}
```

---
### 7. **Hàm Configure Services trong startup class dùng để làm gì?**

Đây là một ==hàm tùy chọn== của **startup class**. Nó ==được dùng để cấu hình các service== được sử dụng trong ứng dụng. Hàm này được gọi đầu tiên khi ứng dụng được request lần đầu. Sử dụng hàm, chúng ta có thể thêm các service vào **DI-container**, như vậy các service sẽ có sẳn như 1 **dependency** trong constructor của controller.  


---
### 8. **Hàm Configure trong startup class dùng để làm gì?**

Nó ==định nghĩa cách mà ứng dụng sẽ hồi đáp== tới mỗi **HTTP request**. Chúng ta có thể cấu hình **Request Pipeline** bằng cấu hình **Middleware**. Nó chấp nhận **IApplication Builder** như là một parameter và nó cũng thêm 2 parameter tùy chọn: ==IHostingEnvironment và ILoggerFactory==. Sử dụng hàm này, chúng ta có thể cấu hình Middleware có sẵn như Routing, authentication, session,... cũng như các Third-party Middleware.

---
### 9. **Middleware là gì?**

Là thành phần của một phần mềm ==đóng vai trò tác động vào request pipeline== (Luồng request) để xử lý chúng và tạo ra **response** phản hồi lại **client**. Mỗi một tiến trình middleware thao tác với các **request** nhân được từ **middleware** trước nó. ==Nó cũng quyết định gọi middleware tiếp theo trong pipeline hoặc trả về response== cho **middleware** ngay trước nó (Ngắt pipeline).


---
### 10. **IApplicationBuilder.Use() và IApplicationBuilder.Run() khác nhau như thế nào?**

Chúng ta có thể sử dụng cả 2 phương thức trên vào trong phương thức **Configure** của **startup class**.

Cả hai phương thức được sử dụng để thêm **Middleware Delegate** tới **Request Pipeline** của ứng dụng.

Việc thêm **Middleware** bằng cách dùng **IApplicationBuilder.Use** có thể gọi Middleware tiếp theo trong **Pipeline**, trong khi việc sử dụng **IApplicationBuilder.Run** để thêm Middleware thì sẽ không bao giờ gọi đọc Middleware tiếp theo.

Sau khi gọi phương thức **IApplicationBuilder.Run**, hệ thống sẽ ngưng thêm **Middleware** trong **Request Pipeline**.

---
### 11. **Công dụng của hàm "Map" trong IApplicationBuilder trong việc  thêm Middleware vào Asp.Net Core Pipeline là gì?**

Nó được dùng để chia nhánh **Pipeline**. Nó phân nhánh **Asp.NET Core Pipeline** dựa vào việc khớp đường dẫn Request. Nếu đường dẫn Request bắt đầu với đường dẫn đã cho, **Middleware** trên nhánh đó sẽ được thực thi.

```csharp
public void Configure(IApplicationBuilder app){
	app.Map("/path1", Middleware1);
	app.Map("/path2", Middleware2);
}
```


---
### 12. **Trình bày về routing trong Asp.Net Core?**

**Routing** là chức năng ánh xạ các **request** đến bộ xử lý định tuyến. **Route** có thể có nhiều giá trị (được trích xuất ra từ URL) thường được dùng để xử lý **request**. Tất cả các Route được đăng ký khi ứng dụng bắt đầu. **ASP.NET Core** hộ trợ 2 loại Routing:
- Định tuyến thông thường.
- Định tuyến thuộc tính (Attribute routing).

**Routing** sử dụng các **Route** ==để ánh xạ các **Request** với bộ định tuyến== và **Generate URL** được dùng trong việc hồi đáp request. Hầu hết ứng dụng có 1 bộ tập hợp Route và ==bộ tập hợp này được dùng trong việc xử lý request==. Phương thức **RouteAssync** thường được sử dụng để ánh xạ request (ứng với URL) có sẵn trong bộ tập hợp Route

---
### 13. **Làm gì để thiết lập Session trong ASP.NET Core?**

**Middleware** cho **Session** được cung cấp bởi gói **Microsoft.AspNetCore.Sesion**. Để ứng dụng **Session** trong **Asp.net Core**, chúng ta cần thêm gói này tới **file .csproj** và thêm **Session Middleware** tới **Asp.Net Core Request Pipeline**.

```csharp
public class Startup{

	public void ConfigureServices(IServiceCollection services){
		services.AddSession();;
		services.AddMvc();
	}

	public void Configure(IApplicationBuilder app, IHostingEnviroment env){
		app.UseSession();
	}
}
```

---
### 14. **Các file JSON nào có sẳn trong ASP.NET Core**

- **global.json**
- **launchsettings.json**
- **appsettings.json**
- **bundleconfig.json**
- **bower.json**
- **package.json**

---
### 15. **Tag-Help trong ASP.NET Core là gì ?**

Nó là một tính năng được cung cấp bởi **Razor view-engine** cho phép chúng ta viết mã **server-side** để tạo và hiển thị các phần tử HTML trong view (Razor). **Tag-helper** là các lớp C# thường được dùng để tạo **view** bằng cách thêm các phần tử HTML. Chức năng của tag-helper rất giống với HTML-helper của Asp.net MVC.

```csharp
//HTML Helper
@Html.TextBoxFor(model => model.FirstName, 
new { @class = "form-control", placeholder = "Enter Your First Name" }) 

//content with tag helper
<input asp-for="FirstName" placeholder="Enter Your First Name" 
class="form-control" /> 

//Equivalent HTML
<input placeholder="Enter Your First Name" class="form-control" 
id="FirstName" name="FirstName" value="" type="text"> 
```

---
### 16. **Làm sao để disable tag-helper ở cấp phần tử?**

Chúng ta có thể **disable tag-helper** ở cấp phần tử bằng cách sử dụng ký tự **“!”**. Ký tự này phải được dùng ở tag đóng và mở.

```csharp
<!span asp-validation-for="phone" class="divPhone"></!span>
```

---
### 17. **Razor Pages trong Asp.Net Core là gì?**

Đây là một tính năng mới được giới thiệu trong **Asp.net Core 2.0**. Nó tuân theo một **mô hình phát triển tập trung** như Asp.net Web forms. Nó hỗ trợ tất cả tính năng của Asp.net Core.

```csharp
@page 
<h1> Hello, Book Reader!</h1> 
<h2> This is Razor Pages </h2>
```

**Razor Pages bắt đầu với directive “@page”**. Directive này xử lý các yêu cầu trực tiếp mà không phải thông qua Controller. Razor Pages có thể có file code-behind, nhưng không thật sự là file code-behind. **Nó là lớp kế thừa từ lớp PageModel**.

---
### 18. **Làm sao chúng ta có thể thực hiện model-binding tự động trong Razor Pages?**

**Razor Pages** cung cấp tùy chọn để liên kết thuộc tính một cách tự động khi truyền dữ liệu sử dụng thuộc tính BindProperty. Mặc định nó chỉ liên kết các thuộc tính với các method không phải GET. Chúng ta cần gán thuộc tính **SupportsGet** thành “true” để liên kết 1 thuộc tính với method GET.

```csharp
public class Test1Model : PageModel
{
   [BindProperty]
   public string Name { get; set; }
}
```

---
### 19. **Làm sao để thêm một Service Dependency vào Controller?**

Có 3 bước đơn giản để thêm mộ custom Service Dependency vào một Controller.

**Bước 1: Tạo Service.**
```csharp
public interface IHelloWorldService{
	string SayHello();
}

public class HelloWorldService: IHelloWorldService{
	public string SayHello(){
		return "Hello";
	}
}
```


**Bước 2: thêm service này vào Service container (Bằng Singleton hoặc Scoped hoặc Transient)**.
```csharp
public void ConfigureServices(IServiceCollection service){
	service.AddTransient<IHelloWorldService, HelloWorldService>();
}
```


**Bước 3: Sửu dụng sẻvice này trong Controller.**
```csharp
public class HomeController: Controller {
	private IHelloWorldService _helloWorldService;

	public HomeController(IHelloWorldService helloWorldService)
		=> _helloWorldService = helloWorldService;
}
```