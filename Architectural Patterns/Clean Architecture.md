**Clean Architecture** là mô hình kiến trúc phần mềm được đề xuất bởi **Robert C.Martin (Uncle Bob),** giúp để tạo ra một hệ thống phần mềm dễ bảo trì và mở rộng và kiểm thử.

![](https://images.viblo.asia/b926dbd9-ca7e-4519-a05e-0bb21a767171.png)

![](https://images.viblo.asia/c10e98e7-999e-4336-8ebd-e2feafae922b.png)

---
## **Nguyên lý cơ bản về Clean Architecture:**

- **Mô tả:** Clean Architecture chia hệ thống thành các lớp (layers) độc lập, tuân theo nguyên tắc SOLID và hướng đến việc tách biệt các mối quan tâm (separation of concerns).
- **Mục tiêu:**
	- Độc lập với framework, cơ sở dữ liệu, giao diện người dùng, và các công cụ bên ngoài.
	- Dễ dàng bảo trì và mở rộng.
	- 📌 **Luồng phụ thuộc phải đi từ ngoài vào trong (Outer Layers → Inner Layers)**.  
	- 📌 **Không có vòng tròn phụ thuộc giữa các lớp** (giảm coupling, tăng cohesion).

---
## **Các lớp bên trong Clean Architecture:**

#### **🎗️Sơ đồ tổng quan:**
```
┌──────────────────────────┐  
│ UI / Presentation Layer  │  (Giao diện, API, Controller)  
└───────────▲──────────────┘  
            │  
┌───────────▼──────────────┐  
│ Application Layer        │  (Use Cases - Luồng nghiệp vụ)  
└───────────▲──────────────┘  
            │  
┌───────────▼──────────────┐  
│ Domain Layer             │  (Business Logic, Entities)  
└───────────▲──────────────┘  
            │  
┌───────────▼──────────────┐  
│ Infrastructure Layer      │  (Database, Framework, External Services)  
└──────────────────────────┘  

```

### 🔹 **1.Domain Layer (Lớp miền – Core Business Logic)**

✅ Chứa các **Entities (Thực thể)** đại diện cho dữ liệu và quy tắc nghiệp vụ.  
✅ Không phụ thuộc vào UI, Database, hay framework nào.

- ##### **Entities (Lớp thực thể):**
	- **Mô tả:** Chứa các đối tượng nghiệp vụ cốt lõi và các quy tắc nghiệp vụ
	- **Ví dụ:** Các lớp như `User`, `Product`, `Order`

### 🔹 **2.Application Layer (Lớp luồng nghiệp vụ – Use Cases)**

✅ Chứa **Use Cases (tương tác giữa Domain và UI)**.  
✅ Định nghĩa cách dữ liệu di chuyển qua hệ thống.  
✅ Không chứa logic trình bày (UI) hoặc chi tiết Database.
- ##### **Use case (Lớp tác vụ):**
	- **Mô tả:** Chứa các quy tắc nghiệp vụ cụ thể và logic ứng dụng
	- **Ví dụ:** Các tác vụ như `CreateUser`, `PlaceOrder`, `ProcessPayment`

### **🔹3.Infrastructure Layer (Lớp cơ sở hạ tầng)**

✅ Chứa **repository, database, framework, external APIs**.  
✅ **Chịu trách nhiệm thực hiện giao tiếp với các hệ thống bên ngoài.**  
✅ Có thể thay thế (ví dụ: đổi SQL Server sang MongoDB mà không ảnh hưởng đến business logic).

##### **Interface Adapters (Lớp điều phối giao diện):**
- **Mô tả:** Chuyển đổi dữ liệu giữa các lớp bên ngoài (như giao diện người dùng, cơ sở dữ liệu) và các lớp bên trong (Entities & Use case)
- **Ví dụ:** Các controller, presenter, gateway.

##### **Frameworks and Drivers (Lớp framework và công cụ):**
- **Mô tả:** Chứa các công cụ bên ngoài như cơ sở dữ liệu, web framework, giao diện người dùng
- **Ví dụ:** Cơ sở dữ liệu MySQL , web framework ASP.NET Core.

---
## **Các nguyên tắc quan trọng:**

**Dependency Rule (Quy tắc phụ thuộc):**
- **Mô tả:** Các lớp bên trong (Entities và Use Cases) không được phụ thuộc vào các lớp bên ngoài (Interface Adapters và Frameworks). Thay vào đó, các lớp bên ngoài phụ thuộc vào các lớp bên trong.
- **Ví dụ:** Use Cases không biết gì về cơ sở dữ liệu, nhưng cơ sở dữ liệu phải tuân theo các interface được định nghĩa bởi Use Case.

**Separation of Concerns (Tách biệt các mối quan tâm):**
- **Mô tả:** Mỗi lớp chỉ nên đảm nhận một nhiệm vụ cụ thể, giúp hệ thống dễ bảo trì và mở rộng.
- **Ví dụ:** Lớp Entities chỉ chứa các đối tượng nghiệp vụ, không chứa  

**Testability (Khả năng kiểm thử):**
- **Mô tả:** hệ thống được thiết kế để dễ dàng kiểm thử phần độc lập.
- **Ví dụ:** Các Use Cases có thể được kiểm thử mà không cần phụ thuộc vào cơ sở dữ liệu thực tế.

---
## **Ví dụ minh họa:**

- Ví dụ một nghiệp vụ của **API Update Product:**
![](https://statics.cdn.200lab.io/2022/05/example-basic-clean-architecture-update-product-api-1.png?width=800)

#### **a. Domain Layer – Định nghĩa Entity**
```
public class Event
{
    public int Id { get; set; }
    public string Name { get; set; }
    public DateTime Date { get; set; }
    
    public Event(string name, DateTime date)
    {
        Name = name;
        Date = date;
    }
}

```
#### **b. Application Layer – Use Case**
```
public class CreateEventHandler
{
    private readonly IEventRepository _eventRepository;

    public CreateEventHandler(IEventRepository eventRepository)
    {
        _eventRepository = eventRepository;
    }

    public async Task Handle(CreateEventCommand command)
    {
        var newEvent = new Event(command.Name, command.Date);
        await _eventRepository.AddAsync(newEvent);
    }
}

```
#### **c. Infrastructure Layer – Repository**
```
public class EventRepository : IEventRepository
{
    private readonly AppDbContext _context;

    public EventRepository(AppDbContext context)
    {
        _context = context;
    }

    public async Task AddAsync(Event eventEntity)
    {
        _context.Events.Add(eventEntity);
        await _context.SaveChangesAsync();
    }
}
```
#### **d. Presentation Layer – API Controller**
```
[ApiController]
[Route("api/events")]
public class EventController : ControllerBase
{
    private readonly CreateEventHandler _handler;

    public EventController(CreateEventHandler handler)
    {
        _handler = handler;
    }

    [HttpPost]
    public async Task<IActionResult> CreateEvent(CreateEventCommand command)
    {
        await _handler.Handle(command);
        return Ok("Event Created Successfully");
    }
}
```


#### **🖼️Cấu trúc thư mục gợi ý:**
```
📂 MyProject
│── 📂 Core
│   ├── 📂 Domain (Entities, Aggregates, Value Objects)
│   ├── 📂 Application (Use Cases, DTOs, Interfaces)
│
│── 📂 Infrastructure
│   ├── 📂 Persistence (Repositories, Database, Migrations)
│   ├── 📂 ExternalServices (Email, APIs, Caching)
│
│── 📂 Presentation
│   ├── 📂 API (Controllers, View Models)
│   ├── 📂 UI (Razor Pages, Blazor, Angular, React)
│
│── 📂 Tests (Unit Tests, Integration Tests)

```

---
## **Ưu/nhược điểm:**

##### **Ưu điểm:**
- Dễ bảo trì và mở rộng.
- Độc lập với framework và công cụ bên ngoài.
- Dễ dàng kiểm thử từng phần độc lập
##### **Nhược điểm:**
- Yêu cầu hiểu biết sâu về thiết kế phần mềm.
- Có thể phức tạp hơn so với các phương pháp thiết kế truyền thống
---
## **Tài liệu:**
1. [Clean Architecture là gì - Ưu nhược và cách dùng hợp lý](https://200lab.io/blog/clean-architecture-uu-nhuoc-va-cach-dung-hop-ly)
2. [Giới thiệu Clean Architecture](https://tedu.com.vn/video/bai-9-gioi-thieu-clean-architecture-1267.html)