
## **Giới thiệu về Dependency Injection (DI) trong ASP.NET Core:**

**Dependency Injection (DI)** là một kỹ thuật giúp ==quản lý sự phụ thuộc giữa các đối tượng== trong ứng dụng, giúp ==tăng tính linh hoạt, dễ mở rộng và kiểm soát==. **ASP.NET** Core có sẵn một **DI Container** để quản lý các vòng đời (**Lifetime**) của **Services**.

![](https://www.code4it.dev/blog/dependency-injection-lifetimes/featuredImage.jpeg)

Ba loại **Lifetime** trong **DI** của **ASP.NET Core:**

1. **Transient** (Luôn tạo mới).
2. **Scoped** (Dùng chung trong request).
3. **Singleton** (Dùng chung trong toàn ứng dụng).

---
## **Tiêu chí so sánh giữa Transient, Scoped & Singleton:**

| Loại          | Vòng đời            | Khi nào được tạo mới?                    | Khi nào nên dùng?                                                                  | Ưu điểm                                           | Nhược điểm                                                       |
| ------------- | ------------------- | ---------------------------------------- | ---------------------------------------------------------------------------------- | ------------------------------------------------- | ---------------------------------------------------------------- |
| **Transient** | Luôn tạo mới        | Mỗi lần được gọi                         | Khi service không lưu trạng thái (Stateless)                                       | Không bị ảnh hưởng bởi dữ liệu cũ                 | Tốn tài nguyên nếu được tạo quá nhiều lần                        |
| **Scoped**    | Trong suốt request  | Một instace trong mỗi request            | Khi cần lưu trạng thái trong request, ví dụ trong **Unit of work**, **Repository** | Dùng chung trog request giúp tiết kiệm tài nguyên | Không phù hợp khi cần dùng chung toàn bộ ứng dụng                |
| **Singleton** | Xuyên suốt ứng dụng | Một lần duy nhất khi ứng dụng khởi động. | Khi cần duy trì trạng thái chung, ví dụ: Logger, Cache, Configuration.             | Tiết kiệm tài nguyên, dùng chung toàn bộ ứng dụng | Nếu chứa dữ liệu thay đổi, có thể gây xung đột (race condition). |

---
## **Khi nào nên dùng từng loại Lifetime?**

##### 1. **Transient - Khi nào nên dùng ?**

👉 **Dùng khi service không cần giữ trạng thái (stateless) và không cần dùng lại instace cũ.**

*Ví dụ:*
- **Service gửi mail.**
- **Service format dữ liệu.**
- **Service xử lý logic đơn giản mà không cần lưu trữ dữ liệu lâu dài.**

🚀 **Ví dụ triển khai Transient trong ASP.NET Core**
```csharp
services.AddTransient<IEmailService, EmailService>();
```

```csharp
public class EmailService : IEmailService
{
    public void SendEmail(string to, string subject, string body)
    {
        Console.WriteLine($"Gửi email đến {to} với tiêu đề {subject}");
    }
}
```
⛔ **Cảnh báo:** Không nên dùng nếu service cần lưu trạng thái, vì mỗi lần gọi sẽ tạo một object mới.

##### 2. **Scoped - Khi nào nên dùng ?**

👉 **Dùng khi cần chia sẻ một service trong suốt request nhưng không cần lưu trữ lâu dài.**

*Ví dụ:*
- **Unit of Work (EF Core DbContext).**
- **Repository Pattern.**

🚀 **Ví dụ triển khai Scoped trong ASP.NET Core**
```csharp
services.AddScoped<IOrderRepository, OrderRepository>();
```

```csharp
public class OrderRepository : IOrderRepository
{
    private readonly ApplicationDbContext _context;

    public OrderRepository(ApplicationDbContext context)
    {
        _context = context;
    }

    public void AddOrder(Order order)
    {
        _context.Orders.Add(order);
        _context.SaveChanges();
    }
}
```
⛔ **Cảnh báo:** Không dùng cho **background tasks** hoặc **Singleton**, vì instance Scoped có thể bị giải phóng trước khi background task kết thúc.

##### 3. **Singleton - Khi nào nên dùng ?**

👉 **Dùng khi cần giữ nguyên một instance xuyên suốt vòng đời của ứng dụng.**

Ví dụ:
- **Logger**
- **Cache** (Lưu trữ danh sách sản phẩm, danh mục, v.v)
- **Configuration Service**

🚀 **Ví dụ triển khai Singleton trong ASP.NET Core**
```csharp
services.AddSingleton<ICacheService, CacheService>();
```

```csharp
public class CacheService : ICacheService
{
    private readonly Dictionary<string, object> _cache = new Dictionary<string, object>();

    public void Set(string key, object value)
    {
        _cache[key] = value;
    }

    public object Get(string key)
    {
        return _cache.TryGetValue(key, out var value) ? value : null;
    }
}
```
⛔ **Cảnh báo:** Nếu lưu trữ dữ liệu thay đổi, cần xử lý **concurrency** (race condition) bằng **lock hoặc thread-safe mechanism**.

---
## **Cách áp dụng DI trong hệ thống lớn**

Khi xây dựng hệ thống lớn, **chọn dúng Lifetime giúp tối ưu hiệu suất và tài nguyên**. Dưới đây là hướng dẫn:

| **Loại Service**                                 | **Nên dùng Lifetime nào?**       |
| ------------------------------------------------ | -------------------------------- |
| **Service stateless (Không lưu trạng thái)**     | Transient                        |
| **Repository / Unit of Work (Giao tiếp với DB)** | Scoped                           |
| **Cache / Logger / Configuration**               | Singleton                        |
| **Background tasks**                             | Singleton hoặc Scoped (cẩn thận) |
#### 1. **Ứng dụng trong Microserrvices:**

- **Request xử lý dữ liệu:** Scoped (Tránh xung đột dữ liệu)
- **Service cache toàn hệ thống:** Singleton
- **Service gủi thông báo:** Transient

#### 2. **Ứng dụng trong E-commerce (Web bán hàng):**

- **Repository/ Unit of work:** Scoped.
- **Lưu thông tin giỏ hàng vào session:** Scoped.
- **Lưu danh sách sản phẩm vào cache:** Singleton.
- **Service thanh toán (Paypal, momo, v.v):** Transient.

#### 3. **Ứng dụng trong hệ thống real-time (Websocket, SignalR, v.v):**

- **Quản lý kết nối WebSocket:** Singleton
- **Gửi tin nhắn real-time:** Transient.

***Ví dụ:***

###### 1. **Database Access:**
```csharp
services.AddScoped<DbContext, AppDbContext>();
services.AddScoped<IUnitOfWork, UnitOfWork>();
```
###### 2. **Business Logic Service:**
```csharp
services.AddScoped<IOrderService, OrderService>();
services.AddScoped<IPaymentService, PaymentService>();
```
###### 3. **Shared Resources:**
```csharp
services.AddSingleton<ICacheManager, DistributedCacheManager>();
services.AddSingleton<ILogger, CentralizedLogger>();
```
###### 4. **Utilities:**
```csharp
services.AddTransient<IEmailFormatter, EmailFormatter>();
services.AddTransient<IFileParser, CsvFileParser>();
```

---
## **🎯Vậy nên sử dụng cái nào?**

**Transient service** là cách an toàn nhất để tạo, vì bạn luôn tạo mới một thể hiện. Nhưng vì thế mà nó sẽ tạo mỗi lần bạn yêu cầu như vậy sẽ ==dùng nhiều bộ nhớ và tài nguyên==. Điều này có thể gây ảnh hưởng không tốt đến hiệu năng nếu quá nhiều thể hiện được tạo.

Sử dụng **Transient Services** ==sẽ phù hợp khi bạn muốn dùng cho các service nhẹ và nhỏ== cũng nhưng không có trạng thái.

**Scope service** thì tốt hơn khi bạn muốn duy trì trạng thái trong một **request**.

**Singleton** được tạo chỉ một lần, nó không bị hủy cho đến khi ứng dụng tắt. Bất cứ việc chiếm bộ nhớ nào với các service này điều tích lại theo thời gina và nó đầy lên. Nhưng cũng ==giúp chúng ta tiết kiệm bộ nhớ== nếu xử lý tốt vì chúng được tạo chỉ 1 lần và sử dụng mọi nơi.

Sử dụng **Singleton** khi bạn cần duy trì ==trạng thái trong hệ thống==. Cấu hình, tham số ứng dụng, các service logging, caching dữ liệu,... là các ví dụ thường dùng Singleton.


---
## **⚠️ Lưu ý:**

Các **service** với mức độ vòng đời thấp hơn được **inject** vào **service** có vòng đời cao hơn sẽ thay đổi service vòng đời thấp hơn thành cao hơn. Điều này sẽ là việc **debug** trở nên khó khăn hơn và nên tránh. Từ thấp đến cao là: **Transient**, **Scoped** và **Singleton**.

##### ***Vì thế hãy nhớ quy tắc:***

###### 1. Không bao giờ **Inject Scoped & Transient service** vào **Singleton service**.
###### 2. Không bao giờ  Inject **Transient service** vào **Scoped service**.
###### 3. **Disposable Objects:**
- .NET tự giải phóng **Scoped/Transient services** khi scope kết thúc
- **Singleton service**s cần tự xử lý việc cleanup nếu cần

---
# **Kết luận – Khi nào tối ưu và an toàn nhất?**

**Transient**: Nhẹ, độc lập, dùng cho tác vụ nhỏ.
**Scoped**: Tốt nhất cho logic request trong Web API.
**Singleton**: Tiết kiệm tài nguyên, dùng cho toàn cục.

✅ **Dùng Transient** nếu bạn **cần một instance mới mỗi lần gọi**, giúp tránh việc giữ trạng thái không cần thiết.  
✅ **Dùng Scoped** nếu bạn **muốn giữ service trong request**, giúp tiết kiệm tài nguyên và tránh conflict dữ liệu.  
✅ **Dùng Singleton** nếu service **cần dùng chung cho toàn ứng dụng** nhưng đảm bảo **không có race condition**.

⛔ **Lưu ý:**

- Không dùng Singleton cho service chứa **DbContext** vì dễ gây lỗi **concurrency**.
- Không dùng Scoped trong **background service** nếu không kiểm soát được vòng đời request.