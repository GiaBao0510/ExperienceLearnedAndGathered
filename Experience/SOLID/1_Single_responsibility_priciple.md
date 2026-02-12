# Single Responsibility Principle (SRP)

## Định nghĩa

**Single Responsibility Principle (SRP)** là nguyên lý đầu tiên trong bộ nguyên tắc **SOLID** của lập trình hướng đối tượng.

![](https://images.viblo.asia/7ae08c24-26d9-447a-b460-4a23f8c85cf0.png)

> **"Mỗi class chỉ nên có một trách nhiệm duy nhất"**
> 
> Hay nói cách khác: **"Mỗi class chỉ nên có một lý do duy nhất để thay đổi"**

---

## Tại sao cần SRP?

Khi một class đảm nhận **nhiều trách nhiệm**:

- ❌ Khó quản lý và bảo trì
- ❌ Khó kiểm tra lỗi (testing)
- ❌ Khó mở rộng
- ❌ Thay đổi một chức năng có thể ảnh hưởng đến chức năng khác

Khi tuân theo SRP:

- ✅ Code dễ đọc, dễ hiểu
- ✅ Dễ bảo trì và mở rộng
- ✅ Dễ test (unit testing)
- ✅ Giảm coupling (ràng buộc giữa các class)

---

## Ví dụ thực tế

> **Ví dụ xây nhà:**
> 
> ❌ **Không tuân theo SRP:** Thuê 1 người thợ làm tất cả (xây tường, lắp điện, sửa ống nước, sơn nhà)
> 
> - Khi cần sửa điện → phải gọi người này
> - Nhưng anh ta đang bận sơn → phải đợi
> - Sửa điện có thể ảnh hưởng đến tường đã xây
> 
> ✅ **Tuân theo SRP:** Thuê từng thợ chuyên môn
> 
> - Thợ xây → chỉ xây tường
> - Thợ điện → chỉ lắp điện
> - Thợ ống nước → chỉ sửa nước
> - Thợ sơn → chỉ sơn nhà
> - Khi cần sửa điện → gọi thợ điện, không ảnh hưởng phần khác

---

## Ví dụ với C#

### ❌ Vi phạm SRP

```csharp
namespace SRP
{
    // Class Employee đảm nhận NHIỀU trách nhiệm
    public class Employee
    {
        public int EmployeeId { get; set; }
        public string EmployeeName { get; set; }
        public string Email { get; set; }

        // Trách nhiệm 1: Quản lý dữ liệu Employee
        public bool AddEmployee(Employee employee)
        {
            // Logic thêm employee vào database
            return true;
        }

        // Trách nhiệm 2: Tạo báo cáo
        public void GenerateReport(Employee employee)
        {
            // Logic tạo báo cáo với Crystal Report
            Console.WriteLine("Generating report...");
        }

        // Trách nhiệm 3: Gửi email
        public void SendEmail(Employee employee, string message)
        {
            // Logic gửi email
            Console.WriteLine($"Sending email to {employee.Email}");
        }
    }
}
```

**Vấn đề:**

- Class `Employee` có **3 trách nhiệm**: quản lý dữ liệu, tạo báo cáo, gửi email
- Nếu cần thay đổi format báo cáo (PDF, Excel) → phải sửa class `Employee`
- Nếu cần thay đổi cách gửi email → phải sửa class `Employee`
- Khó test riêng từng chức năng

### ✅ Tuân theo SRP

```csharp
namespace SRP
{
    // Class Employee chỉ quản lý THÔNG TIN employee
    public class Employee
    {
        public int EmployeeId { get; set; }
        public string EmployeeName { get; set; }
        public string Email { get; set; }
    }

    // Class EmployeeRepository chỉ chịu trách nhiệm LƯU TRỮ dữ liệu
    public class EmployeeRepository
    {
        public bool AddEmployee(Employee employee)
        {
            // Insert into database
            Console.WriteLine($"Added employee: {employee.EmployeeName}");
            return true;
        }

        public Employee GetEmployeeById(int id)
        {
            // Get from database
            return new Employee { EmployeeId = id };
        }
    }

    // Class ReportGenerator chỉ chịu trách nhiệm TẠO BÁO CÁO
    public class ReportGenerator
    {
        public void GenerateReport(Employee employee)
        {
            // Generate report
            Console.WriteLine($"Generating report for: {employee.EmployeeName}");
        }

        public void GeneratePdfReport(Employee employee)
        {
            // Generate PDF report
        }

        public void GenerateExcelReport(Employee employee)
        {
            // Generate Excel report
        }
    }

    // Class EmailService chỉ chịu trách nhiệm GỬI EMAIL
    public class EmailService
    {
        public void SendEmail(string email, string message)
        {
            // Send email logic
            Console.WriteLine($"Email sent to: {email}");
        }
    }

    // Sử dụng
    class Program
    {
        static void Main()
        {
            var employee = new Employee 
            { 
                EmployeeId = 1, 
                EmployeeName = "John Doe",
                Email = "john@example.com"
            };

            var repository = new EmployeeRepository();
            var reportGenerator = new ReportGenerator();
            var emailService = new EmailService();

            repository.AddEmployee(employee);
            reportGenerator.GenerateReport(employee);
            emailService.SendEmail(employee.Email, "Welcome!");
        }
    }
}
```

**Lợi ích:**

- Mỗi class có **1 trách nhiệm duy nhất**
- Muốn thay đổi báo cáo → chỉ sửa `ReportGenerator`
- Muốn thay đổi email → chỉ sửa `EmailService`
- Dễ test từng class độc lập

---

## Ví dụ với Golang

### ❌ Vi phạm SRP

```go
package main

import (
    "database/sql"
    "errors"
    "fmt"
)

// Struct Employee đảm nhận NHIỀU trách nhiệm
type Employee struct {
    ID    int
    Name  string
    Email string
    db    *sql.DB // Database connection
}

// Trách nhiệm 1: Lưu trữ dữ liệu
func (e *Employee) Save() error {
    _, err := e.db.Exec(
        "INSERT INTO employees (id, name, email) VALUES (?, ?, ?)",
        e.ID, e.Name, e.Email,
    )
    return err
}

// Trách nhiệm 2: Validate dữ liệu
func (e *Employee) Validate() error {
    if e.Name == "" {
        return errors.New("name cannot be empty")
    }
    if e.Email == "" {
        return errors.New("email cannot be empty")
    }
    return nil
}

// Trách nhiệm 3: Tạo báo cáo
func (e *Employee) GenerateReport() string {
    return fmt.Sprintf("Report for %s (ID: %d)", e.Name, e.ID)
}

// Trách nhiệm 4: Gửi email
func (e *Employee) SendWelcomeEmail() error {
    // Logic gửi email
    fmt.Printf("Sending email to %s\n", e.Email)
    return nil
}
```

**Vấn đề:**

- Struct `Employee` có **4 trách nhiệm**: lưu trữ, validate, tạo báo cáo, gửi email
- Khó test riêng từng chức năng
- Thay đổi một phần ảnh hưởng toàn bộ

### ✅ Tuân theo SRP

```go
package main

import (
    "database/sql"
    "errors"
    "fmt"
)

// Employee chỉ chứa THÔNG TIN
type Employee struct {
    ID    int
    Name  string
    Email string
}

// EmployeeRepository chỉ chịu trách nhiệm LƯU TRỮ dữ liệu
type EmployeeRepository struct {
    db *sql.DB
}

func NewEmployeeRepository(db *sql.DB) *EmployeeRepository {
    return &EmployeeRepository{db: db}
}

func (r *EmployeeRepository) Save(employee *Employee) error {
    _, err := r.db.Exec(
        "INSERT INTO employees (id, name, email) VALUES (?, ?, ?)",
        employee.ID, employee.Name, employee.Email,
    )
    return err
}

func (r *EmployeeRepository) FindByID(id int) (*Employee, error) {
    var emp Employee
    err := r.db.QueryRow(
        "SELECT id, name, email FROM employees WHERE id = ?", id,
    ).Scan(&emp.ID, &emp.Name, &emp.Email)
    
    if err != nil {
        return nil, err
    }
    return &emp, nil
}

// EmployeeValidator chỉ chịu trách nhiệm VALIDATE dữ liệu
type EmployeeValidator struct{}

func (v *EmployeeValidator) Validate(employee *Employee) error {
    if employee.Name == "" {
        return errors.New("name cannot be empty")
    }
    if employee.Email == "" {
        return errors.New("email cannot be empty")
    }
    // Validate email format
    if !isValidEmail(employee.Email) {
        return errors.New("invalid email format")
    }
    return nil
}

func isValidEmail(email string) bool {
    // Simple email validation
    return len(email) > 3 && contains(email, "@")
}

func contains(str, substr string) bool {
    for i := 0; i < len(str); i++ {
        if str[i:i+1] == substr {
            return true
        }
    }
    return false
}

// ReportGenerator chỉ chịu trách nhiệm TẠO BÁO CÁO
type ReportGenerator struct{}

func (rg *ReportGenerator) GenerateReport(employee *Employee) string {
    return fmt.Sprintf("=== Employee Report ===\nID: %d\nName: %s\nEmail: %s\n", 
        employee.ID, employee.Name, employee.Email)
}

func (rg *ReportGenerator) GeneratePDFReport(employee *Employee) {
    // Logic tạo PDF report
    fmt.Println("Generating PDF report...")
}

// EmailService chỉ chịu trách nhiệm GỬI EMAIL
type EmailService struct {
    smtpServer string
}

func NewEmailService(smtpServer string) *EmailService {
    return &EmailService{smtpServer: smtpServer}
}

func (es *EmailService) SendWelcomeEmail(employee *Employee) error {
    message := fmt.Sprintf("Welcome %s!", employee.Name)
    fmt.Printf("Sending email to %s: %s\n", employee.Email, message)
    // Logic gửi email qua SMTP
    return nil
}

func (es *EmailService) SendEmail(to, subject, body string) error {
    fmt.Printf("Sending email to %s\nSubject: %s\n", to, subject)
    return nil
}

// Sử dụng
func main() {
    // Giả sử có database connection
    var db *sql.DB // Initialize database connection
    
    // Tạo employee
    employee := &Employee{
        ID:    1,
        Name:  "John Doe",
        Email: "john@example.com",
    }

    // Sử dụng các service độc lập
    validator := &EmployeeValidator{}
    repository := NewEmployeeRepository(db)
    reportGenerator := &ReportGenerator{}
    emailService := NewEmailService("smtp.example.com")

    // Validate
    if err := validator.Validate(employee); err != nil {
        fmt.Println("Validation error:", err)
        return
    }

    // Save to database
    if err := repository.Save(employee); err != nil {
        fmt.Println("Save error:", err)
        return
    }

    // Generate report
    report := reportGenerator.GenerateReport(employee)
    fmt.Println(report)

    // Send welcome email
    emailService.SendWelcomeEmail(employee)
}
```

**Lợi ích:**

- Mỗi struct/type có **1 trách nhiệm rõ ràng**
- `Employee` → chỉ chứa data
- `EmployeeRepository` → chỉ lo database
- `EmployeeValidator` → chỉ lo validation
- `ReportGenerator` → chỉ lo báo cáo
- `EmailService` → chỉ lo gửi email
- Dễ test, dễ maintain, dễ mở rộng

---

## So sánh trực quan

### Vi phạm SRP

```
┌─────────────────────────────────────┐
│          Employee Class             │
│                                     │
│  - Data (ID, Name, Email)          │
│  - Save to Database                 │
│  - Validate Data                    │
│  - Generate Report                  │
│  - Send Email                       │
│                                     │
│  → 5 TRÁCH NHIỆM trong 1 class!    │
└─────────────────────────────────────┘
```

### Tuân theo SRP

```
┌──────────────┐   ┌──────────────────┐   ┌─────────────────┐
│   Employee   │   │EmployeeRepository│   │EmployeeValidator│
│              │   │                  │   │                 │
│ - ID         │   │ - Save()         │   │ - Validate()    │
│ - Name       │   │ - FindByID()     │   │                 │
│ - Email      │   │ - Update()       │   │                 │
└──────────────┘   └──────────────────┘   └─────────────────┘

┌──────────────────┐   ┌─────────────────┐
│ ReportGenerator  │   │  EmailService   │
│                  │   │                 │
│ - Generate()     │   │ - SendEmail()   │
│ - GeneratePDF()  │   │ - SendWelcome() │
└──────────────────┘   └─────────────────┘

Mỗi class → 1 TRÁCH NHIỆM duy nhất!
```

---

## Kết luận

### Nguyên tắc vàng của SRP:

> **"One class, one responsibility, one reason to change"**
> 
> **"Một class, một trách nhiệm, một lý do để thay đổi"**

### Cách áp dụng SRP:

1. **Hỏi bản thân:** Class này đang làm bao nhiêu việc?
2. **Nếu > 1 việc:** Tách thành nhiều class
3. **Mỗi class chỉ nên có 1 lý do để thay đổi**

### Lợi ích:

- Code dễ đọc, dễ hiểu
- Dễ bảo trì và mở rộng
- Dễ test (unit testing)
- Giảm bug khi thay đổi code
- Tăng khả năng tái sử dụng code

### Lưu ý:

- Đừng tách quá nhỏ → quá phức tạp
- Cân bằng giữa SRP và pragmatism (thực tế)
- Áp dụng khi cần, không phải lúc nào cũng bắt buộc