# Dependency Inversion Principle (DIP)

## Định nghĩa

**Dependency Inversion Principle (DIP)** là nguyên lý thứ năm và cuối cùng trong bộ nguyên tắc **SOLID**.

![](https://statics.cdn.200lab.io/2024/08/solid-la-gi-dependency-inversion-principle.png?width=800)

> **"High-level modules should not depend on low-level modules. Both should depend on abstractions"**
> 
> **"Module cấp cao không nên phụ thuộc vào module cấp thấp. Cả hai nên phụ thuộc vào abstraction (interface/abstract class)"**

**Nói đơn giản:**

- **KHÔNG** tạo dependency trực tiếp vào concrete class
- Cả hai module nên phụ thuộc vào **interface**
- Giảm sự liên kết chặt chẽ (tight coupling) giữa các module

---

## DIP vs DI - Khác nhau gì?

### **DIP (Dependency Inversion Principle)**

- Là một **nguyên tắc thiết kế**
- Nói về **cách thiết kế** dependencies
- "Depend on abstractions, not concretions"

### **DI (Dependency Injection)**

- Là một **kỹ thuật triển khai**
- Cách **cài đặt** để đạt được DIP
- Inject dependencies vào class thay vì tự tạo

**Quan hệ:** DI là một trong những cách để đạt được DIP.

---

## Giải thích thuật ngữ

### **High-level module (Module cấp cao)**

- Business logic, use cases
- **Ví dụ:** NotificationService, OrderService, UserService

### **Low-level module (Module cấp thấp)**

- Chi tiết kỹ thuật, implementation
- **Ví dụ:** EmailSender, SmsSender, DatabaseRepository

### **Abstraction (Trừu tượng)**

- Interface hoặc Abstract class
- **Ví dụ:** IMessageSender, IRepository

---

## Tại sao cần DIP?

Khi **KHÔNG tuân theo DIP** (tight coupling):

- ❌ High-level module phụ thuộc trực tiếp vào low-level module
- ❌ Khó thay đổi implementation
- ❌ Khó test (không thể mock)
- ❌ Code cứng nhắc, khó mở rộng

Khi **tuân theo DIP** (loose coupling):

- ✅ Dễ thay đổi implementation
- ✅ Dễ test (có thể mock interface)
- ✅ Code linh hoạt, dễ mở rộng
- ✅ Tuân theo Open/Closed Principle

---

## Ví dụ thực tế

> **Ví dụ hệ thống điện và bóng đèn:**
> 
> ❌ **Vi phạm DIP:**
> 
> - Hệ thống điện được thiết kế **chỉ cho bóng đèn sợi đốt**
> - Muốn đổi sang LED → phải **thay đổi hệ thống điện**
> - Tight coupling: Hệ thống điện ← Bóng đèn cụ thể
> 
> ✅ **Tuân theo DIP:**
> 
> - Hệ thống điện thiết kế với **ổ cắm tiêu chuẩn** (abstraction)
> - Bóng đèn nào cũng dùng **cùng chuẩn ổ cắm**
> - Muốn đổi LED → chỉ cần **cắm bóng mới**, không sửa hệ thống
> - Loose coupling: Hệ thống điện → Ổ cắm chuẩn ← Bóng đèn

---

## Ví dụ với C#

### ❌ Vi phạm DIP

```csharp
namespace DIP
{
    // LOW-LEVEL MODULE - Concrete implementation
    public class EmailService
    {
        public void SendEmail(string message)
        {
            Console.WriteLine($"Sending EMAIL: {message}");
        }
    }

    // HIGH-LEVEL MODULE - VI PHẠM DIP
    // Phụ thuộc trực tiếp vào EmailService (low-level)
    public class NotificationService
    {
        private EmailService _emailService;

        public NotificationService()
        {
            // Tạo dependency trực tiếp - TIGHT COUPLING
            _emailService = new EmailService();
        }

        public void SendNotification(string message)
        {
            _emailService.SendEmail(message);
        }
    }

    // Sử dụng
    class Program
    {
        static void Main()
        {
            var notification = new NotificationService();
            notification.SendNotification("Hello!");
            
            // Vấn đề:
            // 1. Muốn đổi sang SMS → phải SỬA NotificationService
            // 2. Không thể test NotificationService độc lập
            // 3. Tight coupling
        }
    }
}
```

**Vấn đề:**

- `NotificationService` (high-level) phụ thuộc trực tiếp vào `EmailService` (low-level)
- Muốn thay đổi sang SMS → phải sửa code `NotificationService`
- Không thể mock EmailService để test

### ✅ Tuân theo DIP

```csharp
namespace DIP
{
    // ABSTRACTION - Interface
    public interface IMessageSender
    {
        void SendMessage(string message);
    }

    // LOW-LEVEL MODULE 1 - Email implementation
    public class EmailService : IMessageSender
    {
        public void SendMessage(string message)
        {
            Console.WriteLine($"Sending EMAIL: {message}");
        }
    }

    // LOW-LEVEL MODULE 2 - SMS implementation
    public class SmsService : IMessageSender
    {
        public void SendMessage(string message)
        {
            Console.WriteLine($"Sending SMS: {message}");
        }
    }

    // LOW-LEVEL MODULE 3 - Push notification
    public class PushNotificationService : IMessageSender
    {
        public void SendMessage(string message)
        {
            Console.WriteLine($"Sending PUSH: {message}");
        }
    }

    // HIGH-LEVEL MODULE - TUÂN THEO DIP
    // Phụ thuộc vào IMessageSender (abstraction)
    public class NotificationService
    {
        private readonly IMessageSender _messageSender;

        // DEPENDENCY INJECTION qua Constructor
        public NotificationService(IMessageSender messageSender)
        {
            _messageSender = messageSender;
        }

        public void SendNotification(string message)
        {
            _messageSender.SendMessage(message);
        }
    }

    // Sử dụng
    class Program
    {
        static void Main()
        {
            // Có thể dễ dàng thay đổi implementation
            
            // Gửi qua Email
            var emailNotification = new NotificationService(new EmailService());
            emailNotification.SendNotification("Hello via Email!");

            // Gửi qua SMS - KHÔNG CẦN sửa NotificationService
            var smsNotification = new NotificationService(new SmsService());
            smsNotification.SendNotification("Hello via SMS!");

            // Gửi qua Push - KHÔNG CẦN sửa NotificationService
            var pushNotification = new NotificationService(new PushNotificationService());
            pushNotification.SendNotification("Hello via Push!");
        }
    }
}
```

**Lợi ích:**

- `NotificationService` phụ thuộc vào `IMessageSender` (abstraction)
- Dễ dàng thay đổi implementation (Email, SMS, Push...)
- Dễ test với mock `IMessageSender`
- Tuân theo DIP và OCP

---

## 3 Cách Dependency Injection trong C#

### 1. Constructor Injection (Khuyên dùng)

```csharp
public class NotificationService
{
    private readonly IMessageSender _messageSender;

    // Inject qua constructor
    public NotificationService(IMessageSender messageSender)
    {
        _messageSender = messageSender;
    }

    public void SendNotification(string message)
    {
        _messageSender.SendMessage(message);
    }
}

// Sử dụng
var service = new NotificationService(new EmailService());
```

**Ưu điểm:** Dependencies rõ ràng, immutable, dễ test

### 2. Property Injection

```csharp
public class NotificationService
{
    // Inject qua property
    public IMessageSender MessageSender { get; set; }

    public void SendNotification(string message)
    {
        MessageSender?.SendMessage(message);
    }
}

// Sử dụng
var service = new NotificationService
{
    MessageSender = new EmailService()
};
```

**Ưu điểm:** Optional dependencies **Nhược điểm:** Có thể null, mutable

### 3. Method Injection

```csharp
public class NotificationService
{
    // Inject qua method parameter
    public void SendNotification(IMessageSender messageSender, string message)
    {
        messageSender.SendMessage(message);
    }
}

// Sử dụng
var service = new NotificationService();
service.SendNotification(new EmailService(), "Hello!");
```

**Ưu điểm:** Linh hoạt per-call **Nhược điểm:** Phải truyền mỗi lần gọi

---

## Ví dụ với Golang

### ❌ Vi phạm DIP (Golang)

```go
package main

import "fmt"

// LOW-LEVEL MODULE - Concrete implementation
type EmailService struct{}

func (e *EmailService) SendEmail(message string) {
    fmt.Printf("Sending EMAIL: %s\n", message)
}

// HIGH-LEVEL MODULE - VI PHẠM DIP
// Phụ thuộc trực tiếp vào EmailService
type NotificationService struct {
    emailService *EmailService
}

func NewNotificationService() *NotificationService {
    return &NotificationService{
        emailService: &EmailService{}, // TIGHT COUPLING
    }
}

func (n *NotificationService) SendNotification(message string) {
    n.emailService.SendEmail(message)
}

func main() {
    service := NewNotificationService()
    service.SendNotification("Hello!")
    
    // Vấn đề: Muốn đổi sang SMS → phải SỬA NotificationService
}
```

### ✅ Tuân theo DIP (Golang)

```go
package main

import "fmt"

// ABSTRACTION - Interface
type MessageSender interface {
    SendMessage(message string)
}

// LOW-LEVEL MODULE 1 - Email implementation
type EmailService struct{}

func (e *EmailService) SendMessage(message string) {
    fmt.Printf("Sending EMAIL: %s\n", message)
}

// LOW-LEVEL MODULE 2 - SMS implementation
type SmsService struct{}

func (s *SmsService) SendMessage(message string) {
    fmt.Printf("Sending SMS: %s\n", message)
}

// LOW-LEVEL MODULE 3 - Push notification
type PushNotificationService struct{}

func (p *PushNotificationService) SendMessage(message string) {
    fmt.Printf("Sending PUSH: %s\n", message)
}

// HIGH-LEVEL MODULE - TUÂN THEO DIP
// Phụ thuộc vào MessageSender (abstraction)
type NotificationService struct {
    messageSender MessageSender
}

// DEPENDENCY INJECTION qua Constructor
func NewNotificationService(messageSender MessageSender) *NotificationService {
    return &NotificationService{
        messageSender: messageSender,
    }
}

func (n *NotificationService) SendNotification(message string) {
    n.messageSender.SendMessage(message)
}

func main() {
    // Có thể dễ dàng thay đổi implementation
    
    // Gửi qua Email
    emailService := NewNotificationService(&EmailService{})
    emailService.SendNotification("Hello via Email!")

    // Gửi qua SMS - KHÔNG CẦN sửa NotificationService
    smsService := NewNotificationService(&SmsService{})
    smsService.SendNotification("Hello via SMS!")

    // Gửi qua Push - KHÔNG CẦN sửa NotificationService
    pushService := NewNotificationService(&PushNotificationService{})
    pushService.SendNotification("Hello via Push!")
}
```

**Lợi ích:** Giống như C#, dễ test và mở rộng

---

## Ví dụ thực tế: Database Repository

### ✅ C# - Repository Pattern với DIP

```csharp
// ABSTRACTION
public interface IUserRepository
{
    User GetById(int id);
    void Save(User user);
}

// LOW-LEVEL - SQL implementation
public class SqlUserRepository : IUserRepository
{
    public User GetById(int id)
    {
        Console.WriteLine($"Getting user {id} from SQL database");
        return new User { Id = id, Name = "John" };
    }

    public void Save(User user)
    {
        Console.WriteLine($"Saving user {user.Name} to SQL database");
    }
}

// LOW-LEVEL - MongoDB implementation
public class MongoUserRepository : IUserRepository
{
    public User GetById(int id)
    {
        Console.WriteLine($"Getting user {id} from MongoDB");
        return new User { Id = id, Name = "Jane" };
    }

    public void Save(User user)
    {
        Console.WriteLine($"Saving user {user.Name} to MongoDB");
    }
}

// HIGH-LEVEL - Business logic
public class UserService
{
    private readonly IUserRepository _repository;

    public UserService(IUserRepository repository)
    {
        _repository = repository;
    }

    public void RegisterUser(User user)
    {
        // Business logic
        Console.WriteLine($"Registering user: {user.Name}");
        _repository.Save(user);
    }

    public User GetUser(int id)
    {
        return _repository.GetById(id);
    }
}

// Sử dụng
var sqlService = new UserService(new SqlUserRepository());
sqlService.RegisterUser(new User { Name = "Alice" });

var mongoService = new UserService(new MongoUserRepository());
mongoService.RegisterUser(new User { Name = "Bob" });
```

### ✅ Golang - Repository Pattern với DIP

```go
package main

import "fmt"

// Domain model
type User struct {
    ID   int
    Name string
}

// ABSTRACTION
type UserRepository interface {
    GetByID(id int) *User
    Save(user *User)
}

// LOW-LEVEL - SQL implementation
type SqlUserRepository struct{}

func (r *SqlUserRepository) GetByID(id int) *User {
    fmt.Printf("Getting user %d from SQL database\n", id)
    return &User{ID: id, Name: "John"}
}

func (r *SqlUserRepository) Save(user *User) {
    fmt.Printf("Saving user %s to SQL database\n", user.Name)
}

// LOW-LEVEL - MongoDB implementation
type MongoUserRepository struct{}

func (r *MongoUserRepository) GetByID(id int) *User {
    fmt.Printf("Getting user %d from MongoDB\n", id)
    return &User{ID: id, Name: "Jane"}
}

func (r *MongoUserRepository) Save(user *User) {
    fmt.Printf("Saving user %s to MongoDB\n", user.Name)
}

// HIGH-LEVEL - Business logic
type UserService struct {
    repository UserRepository
}

func NewUserService(repository UserRepository) *UserService {
    return &UserService{repository: repository}
}

func (s *UserService) RegisterUser(user *User) {
    fmt.Printf("Registering user: %s\n", user.Name)
    s.repository.Save(user)
}

func (s *UserService) GetUser(id int) *User {
    return s.repository.GetByID(id)
}

func main() {
    // Dễ dàng thay đổi database
    sqlService := NewUserService(&SqlUserRepository{})
    sqlService.RegisterUser(&User{Name: "Alice"})

    mongoService := NewUserService(&MongoUserRepository{})
    mongoService.RegisterUser(&User{Name: "Bob"})
}
```

---

## So sánh trực quan

### Vi phạm DIP (Tight Coupling)

```
┌─────────────────────────┐
│  NotificationService    │ ← High-level module
│  (Business Logic)       │
└────────────┬────────────┘
             │ depends on (trực tiếp)
             ↓
┌────────────────────────┐
│    EmailService        │ ← Low-level module
│  (Implementation)      │
└────────────────────────┘

Problem: Muốn thay đổi EmailService → phải sửa NotificationService
```

### Tuân theo DIP (Loose Coupling)

```
┌─────────────────────────┐
│  NotificationService    │ ← High-level module
│  (Business Logic)       │
└────────────┬────────────┘
             │ depends on
             ↓
┌────────────────────────┐
│   IMessageSender       │ ← ABSTRACTION (Interface)
└────────────┬───────────┘
             │ implemented by
      ┌──────┴──────┬──────────────┐
      ↓             ↓              ↓
┌──────────┐  ┌──────────┐  ┌────────────┐
│  Email   │  │   SMS    │  │   Push     │ ← Low-level modules
│ Service  │  │ Service  │  │ Service    │
└──────────┘  └──────────┘  └────────────┘

Solution: Cả high-level và low-level đều phụ thuộc vào abstraction
```

---

## Lợi ích của DIP

### 1. Dễ thay đổi implementation

```csharp
// Chỉ cần thay đổi ở chỗ khởi tạo
var service1 = new NotificationService(new EmailService());
var service2 = new NotificationService(new SmsService());
// NotificationService KHÔNG CẦN sửa code
```

### 2. Dễ test với Mock

```csharp
// Mock implementation cho testing
public class MockMessageSender : IMessageSender
{
    public bool WasCalled { get; private set; }
    
    public void SendMessage(string message)
    {
        WasCalled = true;
    }
}

// Test
var mock = new MockMessageSender();
var service = new NotificationService(mock);
service.SendNotification("Test");
Assert.IsTrue(mock.WasCalled);
```

### 3. Tuân theo Open/Closed Principle

- Thêm implementation mới (SMS, Push...) → KHÔNG sửa code cũ
- Chỉ cần implement interface

### 4. Code linh hoạt, dễ mở rộng

- Có thể swap implementation runtime
- Có thể dùng nhiều implementation khác nhau

---

## Kết luận

### Nguyên tắc vàng của DIP:

> **"Depend on abstractions, not on concretions"**
> 
> **"Phụ thuộc vào abstraction (interface), không phụ thuộc vào concrete class"**

### Cách áp dụng DIP:

1. **Tạo interface cho dependencies**
    
    - Định nghĩa contract rõ ràng
2. **High-level module phụ thuộc vào interface**
    
    - KHÔNG tạo `new ConcreteClass()` trong high-level module
3. **Inject dependencies từ bên ngoài**
    
    - Constructor Injection (khuyên dùng)
    - Property/Method Injection khi cần
4. **Low-level module implement interface**
    
    - Nhiều implementation khác nhau

### Mối quan hệ với các SOLID principles:

- **DIP + SRP:** Mỗi module có trách nhiệm riêng, interface rõ ràng
- **DIP + OCP:** Thêm implementation mới mà không sửa code cũ
- **DIP + LSP:** Subclass có thể thay thế interface
- **DIP + ISP:** Interface nhỏ, cụ thể → dễ implement

### Tóm tắt SOLID:

|Nguyên tắc|Ý nghĩa|
|---|---|
|**S**RP|Một class, một trách nhiệm|
|**O**CP|Mở cho mở rộng, đóng cho sửa đổi|
|**L**SP|Subclass phải thay thế được superclass|
|**I**SP|Interface nhỏ, không ép buộc implement thừa|
|**D**IP|Phụ thuộc vào abstraction, không phụ thuộc vào concrete|

Khi áp dụng đúng SOLID, code sẽ:

- Dễ đọc, dễ hiểu
- Dễ bảo trì và mở rộng
- Ít bug hơn
- Dễ test
- Flexible và scalable