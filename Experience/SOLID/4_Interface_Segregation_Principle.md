# Interface Segregation Principle (ISP)

## Định nghĩa

**Interface Segregation Principle (ISP)** là nguyên lý thứ tư trong bộ nguyên tắc **SOLID**.

![](https://statics.cdn.200lab.io/2024/08/solid-la-gi-interface-segregation-principle.png?width=800)

> **"No client should be forced to depend on methods it does not use"**
> 
> **"Không nên ép buộc client phải implement những method mà nó không sử dụng"**

**Nói đơn giản:**

- Không tạo **"Fat Interface"** (Interface béo/lớn) với nhiều methods
- Tách thành **nhiều interface nhỏ**, mỗi interface có mục đích riêng
- Class chỉ implement những interface mà nó thực sự cần

---

## Tại sao cần ISP?

Khi **KHÔNG tuân theo ISP** (Fat Interface):

- ❌ Class phải implement methods không cần thiết
- ❌ Code thừa, không sử dụng
- ❌ Khó bảo trì (thay đổi interface ảnh hưởng nhiều class)
- ❌ Vi phạm SRP (Single Responsibility)

Khi **tuân theo ISP** (Small Interfaces):

- ✅ Class chỉ implement những gì cần
- ✅ Code gọn, không thừa
- ✅ Dễ maintain và extend
- ✅ Tuân theo SRP

---

## Ví dụ thực tế

> **Ví dụ nhân viên công ty:**
> 
> ❌ **Vi phạm ISP:**
> 
> - Bạn là nhân viên **xử lý đơn hàng**
> - Nhưng bị ép phải làm: **Quản lý tài chính, Điều hành nhân sự, Phát triển sản phẩm**
> - Bạn quá tải, không tập trung được
> 
> ✅ **Tuân theo ISP:**
> 
> - Bạn chỉ làm **xử lý đơn hàng**
> - Các phòng ban khác lo phần của họ
> - Mỗi người tập trung vào chuyên môn của mình

---

## Ví dụ kinh điển: Worker Interface

### ❌ Vi phạm ISP (C#)

```csharp
namespace ISP
{
    // FAT INTERFACE - Vi phạm ISP
    public interface IWorker
    {
        void Work();
        void Eat();
        void Sleep();
    }

    // Human worker - OK với tất cả methods
    public class HumanWorker : IWorker
    {
        public void Work()
        {
            Console.WriteLine("Human is working");
        }

        public void Eat()
        {
            Console.WriteLine("Human is eating");
        }

        public void Sleep()
        {
            Console.WriteLine("Human is sleeping");
        }
    }

    // Robot worker - VI PHẠM ISP!
    // Robot KHÔNG ĂN, KHÔNG NGỦ nhưng phải implement Eat() và Sleep()
    public class RobotWorker : IWorker
    {
        public void Work()
        {
            Console.WriteLine("Robot is working");
        }

        // Phải implement method KHÔNG DÙNG - VI PHẠM ISP
        public void Eat()
        {
            // Robot không ăn!
            throw new NotImplementedException("Robots don't eat");
        }

        // Phải implement method KHÔNG DÙNG - VI PHẠM ISP
        public void Sleep()
        {
            // Robot không ngủ!
            throw new NotImplementedException("Robots don't sleep");
        }
    }

    // Sử dụng
    class Program
    {
        static void Main()
        {
            IWorker human = new HumanWorker();
            IWorker robot = new RobotWorker();

            human.Work();
            human.Eat(); // OK

            robot.Work();
            robot.Eat(); // EXCEPTION! Vi phạm ISP
        }
    }
}
```

**Vấn đề:**

- `RobotWorker` phải implement `Eat()` và `Sleep()` mà nó không dùng
- Phải throw exception hoặc để empty → code thừa, không tốt
- Vi phạm ISP

### ✅ Tuân theo ISP (C#)

```csharp
namespace ISP
{
    // Tách thành nhiều interface nhỏ
    public interface IWorkable
    {
        void Work();
    }

    public interface IEatable
    {
        void Eat();
    }

    public interface ISleepable
    {
        void Sleep();
    }

    // Human implement TẤT CẢ interface
    public class HumanWorker : IWorkable, IEatable, ISleepable
    {
        public void Work()
        {
            Console.WriteLine("Human is working");
        }

        public void Eat()
        {
            Console.WriteLine("Human is eating");
        }

        public void Sleep()
        {
            Console.WriteLine("Human is sleeping");
        }
    }

    // Robot chỉ implement IWorkable
    public class RobotWorker : IWorkable
    {
        public void Work()
        {
            Console.WriteLine("Robot is working");
        }
        // KHÔNG CẦN implement Eat() và Sleep() - Tuân theo ISP!
    }

    // Manager để quản lý workers
    public class WorkManager
    {
        public void ManageWork(IWorkable worker)
        {
            worker.Work();
        }

        public void ManageLunch(IEatable worker)
        {
            worker.Eat();
        }

        public void ManageRest(ISleepable worker)
        {
            worker.Sleep();
        }
    }

    // Sử dụng
    class Program
    {
        static void Main()
        {
            var manager = new WorkManager();

            HumanWorker human = new HumanWorker();
            RobotWorker robot = new RobotWorker();

            // Quản lý work - cả human và robot đều OK
            manager.ManageWork(human);
            manager.ManageWork(robot);

            // Quản lý lunch - chỉ human
            manager.ManageLunch(human);
            // manager.ManageLunch(robot); // Compile error - robot không có IEatable

            // Tuân theo ISP: mỗi class chỉ implement những gì cần!
        }
    }
}
```

**Lợi ích:**

- `RobotWorker` chỉ implement `IWorkable` - không phải implement Eat/Sleep
- Code gọn, không thừa
- Type-safe: không thể gọi Eat() trên Robot

---

## Ví dụ kinh điển: Worker Interface (Golang)

### ❌ Vi phạm ISP (Golang)

```go
package main

import "fmt"

// FAT INTERFACE - Vi phạm ISP
type Worker interface {
    Work()
    Eat()
    Sleep()
}

// Human worker - OK với tất cả methods
type HumanWorker struct{}

func (h *HumanWorker) Work() {
    fmt.Println("Human is working")
}

func (h *HumanWorker) Eat() {
    fmt.Println("Human is eating")
}

func (h *HumanWorker) Sleep() {
    fmt.Println("Human is sleeping")
}

// Robot worker - VI PHẠM ISP!
// Robot KHÔNG ĂN, KHÔNG NGỦ nhưng phải implement Eat() và Sleep()
type RobotWorker struct{}

func (r *RobotWorker) Work() {
    fmt.Println("Robot is working")
}

// Phải implement method KHÔNG DÙNG - VI PHẠM ISP
func (r *RobotWorker) Eat() {
    panic("Robots don't eat")
}

// Phải implement method KHÔNG DÙNG - VI PHẠM ISP
func (r *RobotWorker) Sleep() {
    panic("Robots don't sleep")
}

func main() {
    var human Worker = &HumanWorker{}
    var robot Worker = &RobotWorker{}

    human.Work()
    human.Eat() // OK

    robot.Work()
    // robot.Eat() // PANIC! Vi phạm ISP
}
```

### ✅ Tuân theo ISP (Golang)

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

// Human implement TẤT CẢ interface
type HumanWorker struct{}

func (h *HumanWorker) Work() {
    fmt.Println("Human is working")
}

func (h *HumanWorker) Eat() {
    fmt.Println("Human is eating")
}

func (h *HumanWorker) Sleep() {
    fmt.Println("Human is sleeping")
}

// Robot chỉ implement Workable
type RobotWorker struct{}

func (r *RobotWorker) Work() {
    fmt.Println("Robot is working")
}
// KHÔNG CẦN implement Eat() và Sleep() - Tuân theo ISP!

// Manager để quản lý workers
type WorkManager struct{}

func (wm *WorkManager) ManageWork(worker Workable) {
    worker.Work()
}

func (wm *WorkManager) ManageLunch(worker Eatable) {
    worker.Eat()
}

func (wm *WorkManager) ManageRest(worker Sleepable) {
    worker.Sleep()
}

func main() {
    manager := &WorkManager{}

    human := &HumanWorker{}
    robot := &RobotWorker{}

    // Quản lý work - cả human và robot đều OK
    manager.ManageWork(human)
    manager.ManageWork(robot)

    // Quản lý lunch - chỉ human
    manager.ManageLunch(human)
    // manager.ManageLunch(robot) // Compile error - robot không implement Eatable

    // Tuân theo ISP: mỗi struct chỉ implement những gì cần!
}
```

---

## Ví dụ thực tế: Printer System

### ❌ Vi phạm ISP (C#)

```csharp
// FAT INTERFACE - Vi phạm ISP
public interface IPrinter
{
    void Print(string document);
    void Scan(string document);
    void Fax(string document);
}

// All-in-one printer - OK
public class AllInOnePrinter : IPrinter
{
    public void Print(string document)
    {
        Console.WriteLine($"Printing: {document}");
    }

    public void Scan(string document)
    {
        Console.WriteLine($"Scanning: {document}");
    }

    public void Fax(string document)
    {
        Console.WriteLine($"Faxing: {document}");
    }
}

// Simple printer - VI PHẠM ISP!
// Chỉ cần Print nhưng phải implement Scan và Fax
public class SimplePrinter : IPrinter
{
    public void Print(string document)
    {
        Console.WriteLine($"Printing: {document}");
    }

    // VI PHẠM ISP - phải implement method không dùng
    public void Scan(string document)
    {
        throw new NotImplementedException("This printer can't scan");
    }

    // VI PHẠM ISP - phải implement method không dùng
    public void Fax(string document)
    {
        throw new NotImplementedException("This printer can't fax");
    }
}
```

### ✅ Tuân theo ISP (C#)

```csharp
// Tách thành nhiều interface nhỏ
public interface IPrintable
{
    void Print(string document);
}

public interface IScannable
{
    void Scan(string document);
}

public interface IFaxable
{
    void Fax(string document);
}

// All-in-one printer implement tất cả
public class AllInOnePrinter : IPrintable, IScannable, IFaxable
{
    public void Print(string document)
    {
        Console.WriteLine($"Printing: {document}");
    }

    public void Scan(string document)
    {
        Console.WriteLine($"Scanning: {document}");
    }

    public void Fax(string document)
    {
        Console.WriteLine($"Faxing: {document}");
    }
}

// Simple printer chỉ implement IPrintable
public class SimplePrinter : IPrintable
{
    public void Print(string document)
    {
        Console.WriteLine($"Printing: {document}");
    }
    // KHÔNG CẦN implement Scan và Fax - Tuân theo ISP!
}

// Scanner chỉ implement IScannable
public class Scanner : IScannable
{
    public void Scan(string document)
    {
        Console.WriteLine($"Scanning: {document}");
    }
}
```

### ✅ Tuân theo ISP (Golang)

```go
// Tách thành nhiều interface nhỏ
type Printable interface {
    Print(document string)
}

type Scannable interface {
    Scan(document string)
}

type Faxable interface {
    Fax(document string)
}

// All-in-one printer
type AllInOnePrinter struct{}

func (p *AllInOnePrinter) Print(document string) {
    fmt.Printf("Printing: %s\n", document)
}

func (p *AllInOnePrinter) Scan(document string) {
    fmt.Printf("Scanning: %s\n", document)
}

func (p *AllInOnePrinter) Fax(document string) {
    fmt.Printf("Faxing: %s\n", document)
}

// Simple printer chỉ implement Printable
type SimplePrinter struct{}

func (p *SimplePrinter) Print(document string) {
    fmt.Printf("Printing: %s\n", document)
}
// KHÔNG CẦN implement Scan và Fax - Tuân theo ISP!

// Scanner chỉ implement Scannable
type Scanner struct{}

func (s *Scanner) Scan(document string) {
    fmt.Printf("Scanning: %s\n", document)
}
```

---

## Ví dụ: Report Export System

### ❌ Vi phạm ISP (C#)

```csharp
// FAT INTERFACE - Vi phạm ISP
public interface IReportExporter
{
    void ExportToPdf();
    void ExportToExcel();
    void ExportToWord();
    void ExportToCsv();
}

// Employee chỉ cần Excel nhưng phải implement tất cả
public class EmployeeReportExporter : IReportExporter
{
    public void ExportToExcel()
    {
        Console.WriteLine("Exporting to Excel");
    }

    // VI PHẠM ISP - không dùng nhưng phải implement
    public void ExportToPdf()
    {
        throw new NotImplementedException();
    }

    public void ExportToWord()
    {
        throw new NotImplementedException();
    }

    public void ExportToCsv()
    {
        throw new NotImplementedException();
    }
}
```

### ✅ Tuân theo ISP (C# & Golang)

**C#:**

```csharp
public interface IPdfExportable { void ExportToPdf(); }
public interface IExcelExportable { void ExportToExcel(); }
public interface IWordExportable { void ExportToWord(); }
public interface ICsvExportable { void ExportToCsv(); }

// Employee chỉ implement Excel
public class EmployeeReportExporter : IExcelExportable
{
    public void ExportToExcel()
    {
        Console.WriteLine("Exporting to Excel");
    }
}

// Financial report implement cả PDF và Excel
public class FinancialReportExporter : IPdfExportable, IExcelExportable
{
    public void ExportToPdf()
    {
        Console.WriteLine("Exporting to PDF");
    }

    public void ExportToExcel()
    {
        Console.WriteLine("Exporting to Excel");
    }
}
```

**Golang:**

```go
type PdfExportable interface { ExportToPdf() }
type ExcelExportable interface { ExportToExcel() }
type WordExportable interface { ExportToWord() }
type CsvExportable interface { ExportToCsv() }

// Employee chỉ implement Excel
type EmployeeReportExporter struct{}

func (e *EmployeeReportExporter) ExportToExcel() {
    fmt.Println("Exporting to Excel")
}

// Financial report implement cả PDF và Excel
type FinancialReportExporter struct{}

func (f *FinancialReportExporter) ExportToPdf() {
    fmt.Println("Exporting to PDF")
}

func (f *FinancialReportExporter) ExportToExcel() {
    fmt.Println("Exporting to Excel")
}
```

---

## So sánh trực quan

### Vi phạm ISP (Fat Interface)

```
┌─────────────────────────────────┐
│     IWorker (Fat Interface)     │
│                                 │
│  + Work()                       │
│  + Eat()                        │
│  + Sleep()                      │
└────────────┬────────────────────┘
             │
     ┌───────┴───────┬────────────────┐
     │               │                │
┌────▼────┐    ┌────▼────┐    ┌─────▼──────┐
│  Human  │    │  Robot  │    │   AI       │
│         │    │         │    │            │
│ Work() ✓│    │ Work() ✓│    │ Work() ✓   │
│ Eat() ✓ │    │ Eat() ✗ │    │ Eat() ✗    │
│ Sleep()✓│    │ Sleep()✗│    │ Sleep() ✗  │
└─────────┘    │throw!   │    │throw!      │
               └─────────┘    └────────────┘
               VI PHẠM ISP!   VI PHẠM ISP!
```

### Tuân theo ISP (Segregated Interfaces)

```
┌──────────┐  ┌──────────┐  ┌───────────┐
│IWorkable │  │ IEatable │  │ISleepable │
│          │  │          │  │           │
│+ Work()  │  │+ Eat()   │  │+ Sleep()  │
└────┬─────┘  └────┬─────┘  └─────┬─────┘
     │             │              │
     │        ┌────┴────┬─────────┘
     │        │         │
┌────▼────────▼─────────▼──┐
│         Human            │
│  implements all 3        │
└──────────────────────────┘

┌─────────────┐
│IWorkable    │
│             │
│+ Work()     │
└──────┬──────┘
       │
┌──────▼──────┐
│   Robot     │
│ chỉ cần 1   │
└─────────────┘

TUÂN THEO ISP - Mỗi class chỉ implement những gì cần!
```

---

## Nguyên tắc áp dụng ISP

### 1. Nhận diện Fat Interface

**Dấu hiệu Fat Interface:**

- Interface có nhiều methods (> 3-5 methods)
- Một số implementations throw `NotImplementedException`
- Một số implementations có empty methods
- Methods không liên quan đến nhau

### 2. Cách tách Interface

**Theo chức năng:**

```csharp
// Thay vì
interface IAnimal {
    void Walk();
    void Fly();
    void Swim();
}

// Tách thành
interface IWalkable { void Walk(); }
interface IFlyable { void Fly(); }
interface ISwimmable { void Swim(); }
```

**Theo role:**

```csharp
// Thay vì
interface IUser {
    void Login();
    void ViewContent();
    void EditContent();
    void DeleteContent();
}

// Tách thành
interface IReadOnlyUser { void ViewContent(); }
interface IContentEditor : IReadOnlyUser { void EditContent(); }
interface IContentAdmin : IContentEditor { void DeleteContent(); }
```

---

## Kết luận

### Nguyên tắc vàng của ISP:

> **"Many client-specific interfaces are better than one general-purpose interface"**
> 
> **"Nhiều interface nhỏ, cụ thể tốt hơn một interface lớn, chung chung"**

### Lợi ích chính:

1. **Code gọn, không thừa**
    
    - Class chỉ implement những gì cần
    - Không có empty methods hay exceptions
2. **Dễ maintain**
    
    - Thay đổi interface chỉ ảnh hưởng đúng những class cần
    - Không ảnh hưởng class không liên quan
3. **Tuân theo SRP**
    
    - Mỗi interface có trách nhiệm rõ ràng
    - Class chỉ implement những trách nhiệm của nó

### Khi nào nên tách interface:

✅ **Nên tách khi:**

- Interface có nhiều methods không liên quan
- Một số class không cần tất cả methods
- Xuất hiện empty implementation hoặc exceptions

❌ **Không cần tách khi:**

- Interface nhỏ (2-3 methods liên quan)
- Tất cả implementations đều cần tất cả methods
- Tách quá nhỏ làm phức tạp không cần thiết

### Mối quan hệ với các nguyên tắc khác:

- **ISP + SRP:** ISP hỗ trợ SRP - interface nhỏ → class có trách nhiệm rõ ràng
- **ISP + LSP:** ISP giúp LSP dễ dàng hơn - không bắt buộc implement method không cần