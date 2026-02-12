# Open/Closed Principle (OCP)

## Định nghĩa

**Open/Closed Principle (OCP)** là nguyên lý thứ hai trong bộ nguyên tắc **SOLID**.

> **"Software entities should be OPEN for extension but CLOSED for modification"**
> 
> **"Các thực thể phần mềm nên MỞ cho việc mở rộng nhưng ĐÓNG cho việc sửa đổi"**

![](https://statics.cdn.200lab.io/2024/08/solid-la-gi-single-open-closed-principle.png?width=1200)

**Giải thích:**

- **OPEN for extension (Mở để mở rộng):** Có thể thêm tính năng mới dễ dàng
- **CLOSED for modification (Đóng để sửa đổi):** Không cần thay đổi code hiện có

---

## Tại sao cần OCP?

Khi **KHÔNG tuân theo OCP** (sửa code cũ mỗi lần thêm tính năng):

- ❌ Dễ gây ra bug ở code đang hoạt động tốt
- ❌ Phải test lại toàn bộ code cũ
- ❌ Code ngày càng phức tạp (nhiều if-else)
- ❌ Khó bảo trì và mở rộng

Khi **tuân theo OCP** (thêm code mới, không sửa code cũ):

- ✅ Code cũ vẫn ổn định, không bị ảnh hưởng
- ✅ Chỉ cần test code mới
- ✅ Dễ thêm tính năng mới
- ✅ Code sạch, dễ maintain

---

## Ví dụ thực tế

> **Ví dụ động cơ ô tô:**
> 
> ❌ **Thiết kế không linh hoạt:**
> 
> - Ban đầu xe dùng động cơ xăng
> - Muốn thêm động cơ điện → phải **tháo rời toàn bộ** xe
> - Phải **thay đổi kết cấu** xe
> - Tốn thời gian, chi phí cao, rủi ro lớn
> 
> ✅ **Thiết kế linh hoạt (OCP):**
> 
> - Thiết kế xe với **khung chuẩn** cho động cơ
> - Động cơ xăng, điện, hybrid đều **cắm vào cùng khung**
> - Muốn đổi động cơ → chỉ cần **thay module** động cơ
> - Không cần thay đổi kết cấu xe
> - **Mở rộng** (thêm động cơ mới) nhưng **không sửa** (không đụng vào khung xe)

---

## Ví dụ với C-Sharp

### ❌ Vi phạm OCP

```csharp
namespace OCP
{
    public class Employee
    {
        public int Id { get; set; }
        public string Name { get; set; }
    }

    // Class vi phạm OCP - phải SỬA CODE mỗi lần thêm format mới
    public class ReportGenerator
    {
        public string ReportType { get; set; }

        // Phải SỬA method này mỗi khi thêm format mới
        public void GenerateReport(Employee employee)
        {
            if (ReportType == "PDF")
            {
                // Logic tạo PDF report
                Console.WriteLine($"Generating PDF report for {employee.Name}");
            }
            else if (ReportType == "Excel")
            {
                // Logic tạo Excel report
                Console.WriteLine($"Generating Excel report for {employee.Name}");
            }
            else if (ReportType == "Word")  // Phải THÊM if này → Vi phạm OCP
            {
                // Logic tạo Word report
                Console.WriteLine($"Generating Word report for {employee.Name}");
            }
            // Nếu cần thêm CSV, JSON... → phải tiếp tục thêm if-else
        }
    }
}
```

**Vấn đề:**

- Mỗi lần thêm format mới (Word, CSV, JSON...) → phải **SỬA** method `GenerateReport`
- Nhiều if-else → code phức tạp
- Test lại toàn bộ logic mỗi lần thay đổi
- Vi phạm OCP: **không ĐÓNG cho sửa đổi**

### ✅ Tuân theo OCP

```csharp
namespace OCP
{
    public class Employee
    {
        public int Id { get; set; }
        public string Name { get; set; }
    }

    // Interface định nghĩa hợp đồng cho việc tạo report
    public interface IReportGenerator
    {
        void GenerateReport(Employee employee);
    }

    // PDF Report - Class mới, KHÔNG SỬA code cũ
    public class PdfReportGenerator : IReportGenerator
    {
        public void GenerateReport(Employee employee)
        {
            Console.WriteLine($"[PDF] Generating report for {employee.Name}");
            // Logic cụ thể cho PDF
        }
    }

    // Excel Report - Class mới, KHÔNG SỬA code cũ
    public class ExcelReportGenerator : IReportGenerator
    {
        public void GenerateReport(Employee employee)
        {
            Console.WriteLine($"[Excel] Generating report for {employee.Name}");
            // Logic cụ thể cho Excel
        }
    }

    // Word Report - THÊM class mới, KHÔNG SỬA code cũ
    public class WordReportGenerator : IReportGenerator
    {
        public void GenerateReport(Employee employee)
        {
            Console.WriteLine($"[Word] Generating report for {employee.Name}");
            // Logic cụ thể cho Word
        }
    }

    // CSV Report - THÊM class mới trong tương lai
    public class CsvReportGenerator : IReportGenerator
    {
        public void GenerateReport(Employee employee)
        {
            Console.WriteLine($"[CSV] Generating report for {employee.Name}");
            // Logic cụ thể cho CSV
        }
    }

    // Service để sử dụng các report generators
    public class ReportService
    {
        private readonly IReportGenerator _reportGenerator;

        // Dependency Injection - nhận IReportGenerator
        public ReportService(IReportGenerator reportGenerator)
        {
            _reportGenerator = reportGenerator;
        }

        public void CreateReport(Employee employee)
        {
            _reportGenerator.GenerateReport(employee);
        }
    }

    // Sử dụng
    class Program
    {
        static void Main()
        {
            var employee = new Employee { Id = 1, Name = "John Doe" };

            // Tạo PDF report
            var pdfService = new ReportService(new PdfReportGenerator());
            pdfService.CreateReport(employee);

            // Tạo Excel report
            var excelService = new ReportService(new ExcelReportGenerator());
            excelService.CreateReport(employee);

            // Tạo Word report - THÊM MỚI, không sửa code cũ
            var wordService = new ReportService(new WordReportGenerator());
            wordService.CreateReport(employee);
        }
    }
}
```

**Lợi ích:**

- Thêm format mới (Word, CSV, JSON...) → chỉ cần **THÊM class mới**
- **KHÔNG SỬA** code cũ (PdfReportGenerator, ExcelReportGenerator vẫn nguyên)
- Code cũ vẫn hoạt động tốt, không bị ảnh hưởng
- Tuân theo OCP: **MỞ cho mở rộng, ĐÓNG cho sửa đổi**

---

## Ví dụ với Golang

### ❌ Vi phạm OCP

```go
package main

import "fmt"

type Employee struct {
    ID   int
    Name string
}

// Vi phạm OCP - phải SỬA CODE mỗi lần thêm format mới
type ReportGenerator struct {
    ReportType string
}

// Phải SỬA method này mỗi khi thêm format mới
func (rg *ReportGenerator) GenerateReport(employee Employee) {
    if rg.ReportType == "PDF" {
        fmt.Printf("Generating PDF report for %s\n", employee.Name)
    } else if rg.ReportType == "Excel" {
        fmt.Printf("Generating Excel report for %s\n", employee.Name)
    } else if rg.ReportType == "Word" { // Phải THÊM if này → Vi phạm OCP
        fmt.Printf("Generating Word report for %s\n", employee.Name)
    }
    // Nếu cần CSV, JSON... → phải tiếp tục thêm if-else
}

func main() {
    employee := Employee{ID: 1, Name: "John Doe"}
    
    generator := &ReportGenerator{ReportType: "PDF"}
    generator.GenerateReport(employee)
}
```

**Vấn đề:**
- Mỗi lần thêm format → phải **SỬA** method `GenerateReport`
- Nhiều if-else → code phức tạp
- Vi phạm OCP

### ✅ Tuân theo OCP

```go
package main

import "fmt"

type Employee struct {
    ID   int
    Name string
}

// Interface định nghĩa hợp đồng cho việc tạo report
type ReportGenerator interface {
    GenerateReport(employee Employee)
}

// PDF Report - Struct mới, KHÔNG SỬA code cũ
type PdfReportGenerator struct{}

func (p *PdfReportGenerator) GenerateReport(employee Employee) {
    fmt.Printf("[PDF] Generating report for %s\n", employee.Name)
    // Logic cụ thể cho PDF
}

// Excel Report - Struct mới, KHÔNG SỬA code cũ
type ExcelReportGenerator struct{}

func (e *ExcelReportGenerator) GenerateReport(employee Employee) {
    fmt.Printf("[Excel] Generating report for %s\n", employee.Name)
    // Logic cụ thể cho Excel
}

// Word Report - THÊM struct mới, KHÔNG SỬA code cũ
type WordReportGenerator struct{}

func (w *WordReportGenerator) GenerateReport(employee Employee) {
    fmt.Printf("[Word] Generating report for %s\n", employee.Name)
    // Logic cụ thể cho Word
}

// CSV Report - THÊM struct mới trong tương lai
type CsvReportGenerator struct{}

func (c *CsvReportGenerator) GenerateReport(employee Employee) {
    fmt.Printf("[CSV] Generating report for %s\n", employee.Name)
    // Logic cụ thể cho CSV
}

// Service để sử dụng các report generators
type ReportService struct {
    generator ReportGenerator
}

func NewReportService(generator ReportGenerator) *ReportService {
    return &ReportService{generator: generator}
}

func (rs *ReportService) CreateReport(employee Employee) {
    rs.generator.GenerateReport(employee)
}

func main() {
    employee := Employee{ID: 1, Name: "John Doe"}

    // Tạo PDF report
    pdfService := NewReportService(&PdfReportGenerator{})
    pdfService.CreateReport(employee)

    // Tạo Excel report
    excelService := NewReportService(&ExcelReportGenerator{})
    excelService.CreateReport(employee)

    // Tạo Word report - THÊM MỚI, không sửa code cũ
    wordService := NewReportService(&WordReportGenerator{})
    wordService.CreateReport(employee)

    // Tạo CSV report
    csvService := NewReportService(&CsvReportGenerator{})
    csvService.CreateReport(employee)
}
```

**Lợi ích:**

- Thêm format mới → chỉ cần **THÊM struct mới** implement interface
- **KHÔNG SỬA** code cũ
- Code cũ vẫn hoạt động tốt
- Tuân theo OCP

---

## Ví dụ thực tế: Payment System

### C# - Payment System

```csharp
// Interface cho payment methods
public interface IPaymentProcessor
{
    void ProcessPayment(decimal amount);
}

// Credit Card Payment
public class CreditCardPayment : IPaymentProcessor
{
    public void ProcessPayment(decimal amount)
    {
        Console.WriteLine($"Processing ${amount} via Credit Card");
    }
}

// PayPal Payment
public class PayPalPayment : IPaymentProcessor
{
    public void ProcessPayment(decimal amount)
    {
        Console.WriteLine($"Processing ${amount} via PayPal");
    }
}

// THÊM MỚI: Crypto Payment - KHÔNG SỬA code cũ
public class CryptoPayment : IPaymentProcessor
{
    public void ProcessPayment(decimal amount)
    {
        Console.WriteLine($"Processing ${amount} via Cryptocurrency");
    }
}

// Payment Service
public class PaymentService
{
    public void Pay(IPaymentProcessor processor, decimal amount)
    {
        processor.ProcessPayment(amount);
    }
}

// Sử dụng
var paymentService = new PaymentService();
paymentService.Pay(new CreditCardPayment(), 100.00m);
paymentService.Pay(new PayPalPayment(), 50.00m);
paymentService.Pay(new CryptoPayment(), 200.00m); // THÊM MỚI
```

### Golang - Payment System

```go
// Interface cho payment methods
type PaymentProcessor interface {
    ProcessPayment(amount float64)
}

// Credit Card Payment
type CreditCardPayment struct{}

func (c *CreditCardPayment) ProcessPayment(amount float64) {
    fmt.Printf("Processing $%.2f via Credit Card\n", amount)
}

// PayPal Payment
type PayPalPayment struct{}

func (p *PayPalPayment) ProcessPayment(amount float64) {
    fmt.Printf("Processing $%.2f via PayPal\n", amount)
}

// THÊM MỚI: Crypto Payment - KHÔNG SỬA code cũ
type CryptoPayment struct{}

func (c *CryptoPayment) ProcessPayment(amount float64) {
    fmt.Printf("Processing $%.2f via Cryptocurrency\n", amount)
}

// Payment Service
type PaymentService struct{}

func (ps *PaymentService) Pay(processor PaymentProcessor, amount float64) {
    processor.ProcessPayment(amount)
}

// Sử dụng
func main() {
    paymentService := &PaymentService{}
    
    paymentService.Pay(&CreditCardPayment{}, 100.00)
    paymentService.Pay(&PayPalPayment{}, 50.00)
    paymentService.Pay(&CryptoPayment{}, 200.00) // THÊM MỚI
}
```

---

## So sánh trực quan

### Vi phạm OCP

```
┌─────────────────────────────────────────┐
│      ReportGenerator (class/struct)     │
│                                         │
│  GenerateReport(employee):              │
│    if type == "PDF":                    │
│       // PDF logic                      │
│    else if type == "Excel":             │
│       // Excel logic                    │
│    else if type == "Word":   ← THÊM IF  │
│       // Word logic                     │
│    else if type == "CSV":    ← THÊM IF  │
│       // CSV logic                      │
│                                         │
│  → Phải SỬA CODE mỗi lần thêm format   │
└─────────────────────────────────────────┘
```

### Tuân theo OCP

```
         ┌──────────────────────┐
         │  ReportGenerator     │
         │   (Interface)        │
         │                      │
         │  GenerateReport()    │
         └──────────┬───────────┘
                    │
         ┌──────────┴───────────────────┬──────────────┐
         │                              │              │
┌────────▼─────────┐    ┌──────────────▼──┐   ┌──────▼─────────┐
│ PdfReportGen     │    │ ExcelReportGen  │   │ WordReportGen  │
│                  │    │                 │   │   (THÊM MỚI)   │
│ GenerateReport() │    │ GenerateReport()│   │ GenerateReport()│
└──────────────────┘    └─────────────────┘   └────────────────┘

→ THÊM class/struct mới, KHÔNG SỬA code cũ
```

---

## Khi nào nên áp dụng OCP?

### ✅ Nên áp dụng OCP khi:

1. **Có nhiều biến thể của cùng một chức năng**
    
    - Ví dụ: Nhiều format export (PDF, Excel, CSV...)
    - Nhiều payment methods (Credit Card, PayPal, Crypto...)
    - Nhiều notification channels (Email, SMS, Push...)
2. **Chức năng có thể thay đổi/mở rộng trong tương lai**
    
    - Khách hàng có thể yêu cầu thêm format mới
    - Có thể tích hợp thêm payment gateway
3. **Muốn code ổn định, dễ maintain**
    
    - Tránh sửa code đang hoạt động tốt
    - Giảm risk của regression bugs

### ❌ Không cần quá áp dụng khi:

1. **Chức năng đơn giản, không thay đổi**
    
    - Chỉ có 1-2 biến thể
    - Không có kế hoạch mở rộng
2. **Over-engineering**
    
    - Tạo quá nhiều abstraction không cần thiết
    - Code trở nên phức tạp hơn mà không có lợi ích

---

## Cách áp dụng OCP

### Các bước thực hiện:

1. **Xác định phần có thể thay đổi**
    - Tìm những chỗ có nhiều if-else/switch-case
    - Tìm những chức năng có nhiều biến thể

2. **Tạo abstraction (Interface/Abstract Class)**
    - Định nghĩa contract chung cho các biến thể
    - Interface trong C#/Golang

3. **Implement các biến thể cụ thể**
    - Mỗi biến thể là một class/struct riêng
    - Implement interface

4. **Sử dụng Dependency Injection**
    - Truyền interface vào, không truyền concrete class
    - Dễ dàng swap implementation

---

## Kết luận

### Nguyên tắc vàng của OCP:

> **"Open for Extension, Closed for Modification"**
> 
> **"Thêm tính năng mới bằng cách THÊM code, không SỬA code cũ"**

### Lợi ích chính:

1. **Code ổn định**
    - Code cũ không bị sửa → ít bug hơn
    - Chỉ cần test code mới

2. **Dễ mở rộng**
    - Thêm tính năng mới dễ dàng
    - Không ảnh hưởng code hiện tại

3. **Maintainable**
    - Code rõ ràng, dễ hiểu
    - Mỗi class có trách nhiệm riêng

### Công cụ để đạt được OCP:

- **Interface** (C# & Golang)
- **Abstract Class** (C#)
- **Polymorphism** (Đa hình)
- **Dependency Injection**

### Lưu ý:

- OCP và SRP thường đi cùng nhau
- Đừng over-engineer - chỉ áp dụng khi cần
- Cân bằng giữa flexibility và simplicity

---

**Tài liệu này được biên soạn cho mục đích học tập SOLID principles.**