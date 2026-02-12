# SOLID Principles trong Golang

## **SOLID là gì?**

**SOLID** là tập hợp 5 nguyên lý thiết kế phần mềm trong lập trình hướng đối tượng, giúp tạo ra các hệ thống:

- ✅ Dễ bảo trì
- ✅ Dễ mở rộng
- ✅ Ổn định và linh hoạt

**SOLID** là chữ viết tắt ghép từ 5 nguyên tắc:

| Chữ cái | Nguyên tắc                          | Ý nghĩa ngắn gọn                       |
| ------- | ----------------------------------- | -------------------------------------- |
| **S**   | **S**ingle Responsibility Principle | Một struct/type - một trách nhiệm      |
| **O**   | **O**pen/Closed Principle           | Mở cho mở rộng, đóng cho sửa đổi       |
| **L**   | **L**iskov Substitution Principle   | Subtype phải thay thế được basetype    |
| **I**   | **I**nterface Segregation Principle | Interface nhỏ, không ép implement thừa |
| **D**   | **D**ependency Inversion Principle  | Phụ thuộc vào abstraction (interface)  |

---

## **1. Single Responsibility Principle (SRP)**

> **"Một struct/type chỉ nên có một trách nhiệm duy nhất"**

### ❌ Vi phạm SRP

```go
package main

import "fmt"

// Vi phạm SRP - Employee làm quá nhiều việc
type Employee struct {
    ID    int
    Name  string
    Email string
}

// Trách nhiệm 1: Quản lý dữ liệu
func (e *Employee) Save() {
    fmt.Println("Saving employee to database")
}

// Trách nhiệm 2: Tạo báo cáo
func (e *Employee) GenerateReport() {
    fmt.Println("Generating employee report")
}

// Trách nhiệm 3: Gửi email
func (e *Employee) SendEmail(message string) {
    fmt.Printf("Sending email to %s\n", e.Email)
}
```

**Vấn đề:** Employee có 3 trách nhiệm → khó maintain

### ✅ Tuân theo SRP

```go
package main

import "fmt"

// Employee chỉ chứa DATA
type Employee struct {
    ID    int
    Name  string
    Email string
}

// EmployeeRepository - chịu trách nhiệm LƯU TRỮ
type EmployeeRepository struct{}

func (r *EmployeeRepository) Save(emp *Employee) {
    fmt.Printf("Saving employee %s to database\n", emp.Name)
}

// ReportGenerator - chịu trách nhiệm TẠO BÁO CÁO
type ReportGenerator struct{}

func (rg *ReportGenerator) Generate(emp *Employee) {
    fmt.Printf("Generating report for %s\n", emp.Name)
}

// EmailService - chịu trách nhiệm GỬI EMAIL
type EmailService struct{}

func (es *EmailService) Send(email, message string) {
    fmt.Printf("Sending email to %s: %s\n", email, message)
}

func main() {
    emp := &Employee{ID: 1, Name: "John", Email: "john@example.com"}
    
    repo := &EmployeeRepository{}
    repo.Save(emp)
    
    report := &ReportGenerator{}
    report.Generate(emp)
    
    email := &EmailService{}
    email.Send(emp.Email, "Welcome!")
}
```

**Lợi ích:** Mỗi struct có 1 trách nhiệm rõ ràng

---

## **2. Open/Closed Principle (OCP)**

> **"Mở cho mở rộng, đóng cho sửa đổi"**

### ❌ Vi phạm OCP

```go
package main

import "fmt"

type Employee struct {
    Name string
}

// Vi phạm OCP - phải SỬA CODE mỗi lần thêm format
type ReportGenerator struct {
    ReportType string
}

func (rg *ReportGenerator) Generate(emp Employee) {
    if rg.ReportType == "PDF" {
        fmt.Printf("Generating PDF report for %s\n", emp.Name)
    } else if rg.ReportType == "Excel" {
        fmt.Printf("Generating Excel report for %s\n", emp.Name)
    } else if rg.ReportType == "Word" { // Thêm format mới → phải SỬA
        fmt.Printf("Generating Word report for %s\n", emp.Name)
    }
    // Thêm CSV, JSON... → phải tiếp tục sửa
}
```

**Vấn đề:** Mỗi lần thêm format → phải sửa `Generate()`

### ✅ Tuân theo OCP

```go
package main

import "fmt"

type Employee struct {
    Name string
}

// Interface định nghĩa contract
type ReportGenerator interface {
    Generate(emp Employee)
}

// PDF Report - struct mới, KHÔNG sửa code cũ
type PdfReportGenerator struct{}

func (p *PdfReportGenerator) Generate(emp Employee) {
    fmt.Printf("[PDF] Generating report for %s\n", emp.Name)
}

// Excel Report - struct mới, KHÔNG sửa code cũ
type ExcelReportGenerator struct{}

func (e *ExcelReportGenerator) Generate(emp Employee) {
    fmt.Printf("[Excel] Generating report for %s\n", emp.Name)
}

// Word Report - THÊM MỚI, không sửa code cũ
type WordReportGenerator struct{}

func (w *WordReportGenerator) Generate(emp Employee) {
    fmt.Printf("[Word] Generating report for %s\n", emp.Name)
}

// Service sử dụng interface
type ReportService struct {
    generator ReportGenerator
}

func NewReportService(gen ReportGenerator) *ReportService {
    return &ReportService{generator: gen}
}

func (rs *ReportService) CreateReport(emp Employee) {
    rs.generator.Generate(emp)
}

func main() {
    emp := Employee{Name: "John Doe"}
    
    // PDF
    pdfService := NewReportService(&PdfReportGenerator{})
    pdfService.CreateReport(emp)
    
    // Excel
    excelService := NewReportService(&ExcelReportGenerator{})
    excelService.CreateReport(emp)
    
    // Word - THÊM MỚI mà không sửa code cũ
    wordService := NewReportService(&WordReportGenerator{})
    wordService.CreateReport(emp)
}
```

**Lợi ích:** Thêm format mới → chỉ thêm struct, không sửa code cũ

---

## **3. Liskov Substitution Principle (LSP)**

> **"Subtype phải có thể thay thế basetype mà không gây lỗi"**

### ❌ Vi phạm LSP

```go
package main

import "fmt"

// Basetype
type Bird interface {
    Fly()
}

type Sparrow struct{}

func (s *Sparrow) Fly() {
    fmt.Println("Sparrow is flying")
}

// VI PHẠM LSP - Penguin không bay được!
type Penguin struct{}

func (p *Penguin) Fly() {
    panic("Penguins can't fly!") // Exception!
}

func MakeBirdFly(bird Bird) {
    bird.Fly() // Nếu bird là Penguin → PANIC!
}
```

**Vấn đề:** `Penguin` không thể thay thế `Bird`

### ✅ Tuân theo LSP

```go
package main

import "fmt"

// Interface chung
type Bird interface {
    Eat()
}

// Interface cho birds bay được
type FlyingBird interface {
    Bird
    Fly()
}

type Sparrow struct{}

func (s *Sparrow) Eat() {
    fmt.Println("Sparrow is eating")
}

func (s *Sparrow) Fly() {
    fmt.Println("Sparrow is flying")
}

// Penguin chỉ implement Bird
type Penguin struct{}

func (p *Penguin) Eat() {
    fmt.Println("Penguin is eating")
}

func (p *Penguin) Swim() {
    fmt.Println("Penguin is swimming")
}

func MakeFlyingBirdFly(bird FlyingBird) {
    bird.Fly() // An toàn - chỉ nhận birds bay được
}

func FeedBird(bird Bird) {
    bird.Eat() // Hoạt động với mọi bird
}

func main() {
    sparrow := &Sparrow{}
    penguin := &Penguin{}
    
    FeedBird(sparrow) // OK
    FeedBird(penguin) // OK
    
    MakeFlyingBirdFly(sparrow) // OK
    // MakeFlyingBirdFly(penguin) // Compile error - đúng!
}
```

**Lợi ích:** Type-safe, không có runtime panic

---

## **4. Interface Segregation Principle (ISP)**

> **"Không ép buộc implement methods không dùng"**

### ❌ Vi phạm ISP

```go
package main

import "fmt"

// FAT INTERFACE - Vi phạm ISP
type Worker interface {
    Work()
    Eat()
    Sleep()
}

type HumanWorker struct{}

func (h *HumanWorker) Work()  { fmt.Println("Human working") }
func (h *HumanWorker) Eat()   { fmt.Println("Human eating") }
func (h *HumanWorker) Sleep() { fmt.Println("Human sleeping") }

// VI PHẠM ISP - Robot phải implement Eat/Sleep không dùng
type RobotWorker struct{}

func (r *RobotWorker) Work()  { fmt.Println("Robot working") }
func (r *RobotWorker) Eat()   { panic("Robots don't eat") }
func (r *RobotWorker) Sleep() { panic("Robots don't sleep") }
```

**Vấn đề:** Robot phải implement methods không cần

### ✅ Tuân theo ISP

```go
package main

import "fmt"

// Tách thành nhiều interface nhỏ
type Workable interface {
    Work()
}

type Eatable interface {
    Eat()
}

type Sleepable interface {
    Sleep()
}

// Human implement TẤT CẢ
type HumanWorker struct{}

func (h *HumanWorker) Work()  { fmt.Println("Human working") }
func (h *HumanWorker) Eat()   { fmt.Println("Human eating") }
func (h *HumanWorker) Sleep() { fmt.Println("Human sleeping") }

// Robot chỉ implement Workable
type RobotWorker struct{}

func (r *RobotWorker) Work() { fmt.Println("Robot working") }
// KHÔNG CẦN implement Eat/Sleep

type WorkManager struct{}

func (wm *WorkManager) ManageWork(worker Workable) {
    worker.Work()
}

func (wm *WorkManager) ManageLunch(worker Eatable) {
    worker.Eat()
}

func main() {
    manager := &WorkManager{}
    human := &HumanWorker{}
    robot := &RobotWorker{}
    
    manager.ManageWork(human) // OK
    manager.ManageWork(robot) // OK
    
    manager.ManageLunch(human) // OK
    // manager.ManageLunch(robot) // Compile error - đúng!
}
```

**Lợi ích:** Mỗi struct chỉ implement những gì cần

---

## **5. Dependency Inversion Principle (DIP)**

> **"Phụ thuộc vào interface, không phụ thuộc vào concrete struct"**

### ❌ Vi phạm DIP

```go
package main

import "fmt"

// LOW-LEVEL - Concrete implementation
type EmailService struct{}

func (e *EmailService) SendEmail(message string) {
    fmt.Printf("Sending EMAIL: %s\n", message)
}

// HIGH-LEVEL - VI PHẠM DIP
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

// Vấn đề: Muốn đổi sang SMS → phải SỬA NotificationService
```

**Vấn đề:** Tight coupling, khó thay đổi implementation

### ✅ Tuân theo DIP

```go
package main

import "fmt"

// ABSTRACTION - Interface
type MessageSender interface {
    SendMessage(message string)
}

// LOW-LEVEL 1 - Email
type EmailService struct{}

func (e *EmailService) SendMessage(message string) {
    fmt.Printf("Sending EMAIL: %s\n", message)
}

// LOW-LEVEL 2 - SMS
type SmsService struct{}

func (s *SmsService) SendMessage(message string) {
    fmt.Printf("Sending SMS: %s\n", message)
}

// LOW-LEVEL 3 - Push
type PushService struct{}

func (p *PushService) SendMessage(message string) {
    fmt.Printf("Sending PUSH: %s\n", message)
}

// HIGH-LEVEL - TUÂN THEO DIP
// Phụ thuộc vào MessageSender interface
type NotificationService struct {
    sender MessageSender
}

// DEPENDENCY INJECTION qua Constructor
func NewNotificationService(sender MessageSender) *NotificationService {
    return &NotificationService{sender: sender}
}

func (n *NotificationService) SendNotification(message string) {
    n.sender.SendMessage(message)
}

func main() {
    // Dễ dàng thay đổi implementation
    
    emailService := NewNotificationService(&EmailService{})
    emailService.SendNotification("Hello via Email")
    
    smsService := NewNotificationService(&SmsService{})
    smsService.SendNotification("Hello via SMS")
    
    pushService := NewNotificationService(&PushService{})
    pushService.SendNotification("Hello via Push")
}
```

**Lợi ích:** Loose coupling, dễ thay đổi, dễ test

---

## **Ví dụ tổng hợp: User Service**

Áp dụng TẤT CẢ SOLID principles:

```go
package main

import "fmt"

// Domain Model
type User struct {
    ID    int
    Name  string
    Email string
}

// ===== SRP: Tách trách nhiệm =====

// Repository Interface (DIP)
type UserRepository interface {
    Save(user *User) error
    FindByID(id int) (*User, error)
}

// Validator (SRP)
type UserValidator struct{}

func (v *UserValidator) Validate(user *User) error {
    if user.Name == "" {
        return fmt.Errorf("name is required")
    }
    if user.Email == "" {
        return fmt.Errorf("email is required")
    }
    return nil
}

// Notification Interface (ISP - tách nhỏ)
type UserNotifier interface {
    NotifyUserCreated(user *User)
}

// ===== OCP: Dễ mở rộng =====

// Concrete Repository 1 - PostgreSQL
type PostgresUserRepository struct{}

func (r *PostgresUserRepository) Save(user *User) error {
    fmt.Printf("Saving user to PostgreSQL: %s\n", user.Name)
    return nil
}

func (r *PostgresUserRepository) FindByID(id int) (*User, error) {
    fmt.Printf("Finding user from PostgreSQL: %d\n", id)
    return &User{ID: id, Name: "John"}, nil
}

// Concrete Repository 2 - MongoDB (OCP - thêm mới không sửa code cũ)
type MongoUserRepository struct{}

func (r *MongoUserRepository) Save(user *User) error {
    fmt.Printf("Saving user to MongoDB: %s\n", user.Name)
    return nil
}

func (r *MongoUserRepository) FindByID(id int) (*User, error) {
    fmt.Printf("Finding user from MongoDB: %d\n", id)
    return &User{ID: id, Name: "Jane"}, nil
}

// Email Notifier
type EmailNotifier struct{}

func (n *EmailNotifier) NotifyUserCreated(user *User) {
    fmt.Printf("Sending welcome email to %s\n", user.Email)
}

// SMS Notifier (OCP - thêm mới)
type SmsNotifier struct{}

func (n *SmsNotifier) NotifyUserCreated(user *User) {
    fmt.Printf("Sending welcome SMS to %s\n", user.Email)
}

// ===== DIP: Business Logic phụ thuộc vào Interface =====

type UserService struct {
    repository UserRepository // DIP
    validator  *UserValidator
    notifier   UserNotifier // DIP
}

func NewUserService(
    repo UserRepository,
    validator *UserValidator,
    notifier UserNotifier,
) *UserService {
    return &UserService{
        repository: repo,
        validator:  validator,
        notifier:   notifier,
    }
}

func (s *UserService) CreateUser(user *User) error {
    // Validate
    if err := s.validator.Validate(user); err != nil {
        return err
    }
    
    // Save
    if err := s.repository.Save(user); err != nil {
        return err
    }
    
    // Notify
    s.notifier.NotifyUserCreated(user)
    
    return nil
}

// ===== Main - Demo SOLID =====

func main() {
    user := &User{
        ID:    1,
        Name:  "Alice",
        Email: "alice@example.com",
    }
    
    validator := &UserValidator{}
    
    // Sử dụng PostgreSQL + Email
    postgresRepo := &PostgresUserRepository{}
    emailNotifier := &EmailNotifier{}
    service1 := NewUserService(postgresRepo, validator, emailNotifier)
    service1.CreateUser(user)
    
    fmt.Println("---")
    
    // Dễ dàng đổi sang MongoDB + SMS (OCP, DIP)
    mongoRepo := &MongoUserRepository{}
    smsNotifier := &SmsNotifier{}
    service2 := NewUserService(mongoRepo, validator, smsNotifier)
    service2.CreateUser(user)
}
```

**Output:**

```
Saving user to PostgreSQL: Alice
Sending welcome email to alice@example.com
---
Saving user to MongoDB: Alice
Sending welcome SMS to alice@example.com
```

**Phân tích:**

- ✅ **SRP:** `UserValidator`, `UserRepository`, `UserNotifier` - mỗi type 1 trách nhiệm
- ✅ **OCP:** Thêm `MongoUserRepository`, `SmsNotifier` mà không sửa code cũ
- ✅ **LSP:** Mọi repository đều thay thế được cho nhau
- ✅ **ISP:** Interface nhỏ, cụ thể (`UserRepository`, `UserNotifier`)
- ✅ **DIP:** `UserService` phụ thuộc vào interface, không phụ thuộc concrete

---

## **Tóm tắt SOLID**

|Nguyên tắc|Câu hỏi kiểm tra|Giải pháp|
|---|---|---|
|**SRP**|Struct này làm bao nhiêu việc?|Nếu > 1 → tách ra|
|**OCP**|Thêm feature mới có phải sửa code cũ?|Nếu có → dùng interface|
|**LSP**|Subtype thay thế basetype có lỗi?|Nếu có → sửa lại hierarchy|
|**ISP**|Có method nào không dùng?|Nếu có → tách interface|
|**DIP**|Phụ thuộc vào concrete struct?|Nếu có → dùng interface|

---

## **Lợi ích khi áp dụng SOLID**

Khi code tuân theo SOLID:

- ✅ **Dễ đọc** - Mỗi struct có trách nhiệm rõ ràng
- ✅ **Dễ test** - Có thể mock interface
- ✅ **Dễ mở rộng** - Thêm tính năng mà không sửa code cũ
- ✅ **Ít bug** - Thay đổi không ảnh hưởng code khác
- ✅ **Flexible** - Dễ thay đổi implementation
- ✅ **Maintainable** - Team làm việc hiệu quả hơn

---

## **Lưu ý khi áp dụng SOLID trong Golang**

### 1. **Golang không có class, dùng struct + interface**

```go
// Interface định nghĩa behavior
type Reader interface {
    Read() string
}

// Struct implement interface
type FileReader struct{}

func (f *FileReader) Read() string {
    return "reading file"
}
```

### 2. **Interface trong Go là implicit**

```go
// Không cần khai báo "implements"
type Writer interface {
    Write(data string)
}

type FileWriter struct{}

// Tự động implement Writer interface
func (f *FileWriter) Write(data string) {
    fmt.Println("Writing:", data)
}
```

### 3. **Prefer small interfaces**

```go
// ✅ Tốt - interface nhỏ
type Reader interface {
    Read() ([]byte, error)
}

type Writer interface {
    Write([]byte) (int, error)
}

// ❌ Tránh - interface quá lớn
type FileHandler interface {
    Read() ([]byte, error)
    Write([]byte) (int, error)
    Close() error
    Seek(int64, int) (int64, error)
    // ... nhiều methods
}
```

### 4. **Constructor pattern cho DI**

```go
// Sử dụng constructor để inject dependencies
func NewUserService(repo UserRepository) *UserService {
    return &UserService{repository: repo}
}
```

---

## **Kết luận**

SOLID không phải là quy tắc cứng nhắc mà là **nguyên tắc hướng dẫn** giúp:

- Thiết kế code tốt hơn
- Tránh common mistakes
- Tạo codebase bền vững

**Áp dụng khi cần, không over-engineer!**

> **"Make it work, make it right, make it fast"**
> 
> — Kent Beck

**Học tốt SOLID → Code Golang chuyên nghiệp hơn!**

---

**Tài liệu này được biên soạn cho mục đích học tập SOLID principles với Golang.**