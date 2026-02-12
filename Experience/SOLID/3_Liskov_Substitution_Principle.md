# Liskov Substitution Principle (LSP)

## Định nghĩa

**Liskov Substitution Principle (LSP)** là nguyên lý thứ ba trong bộ nguyên tắc **SOLID**.

Nguyên lý này được đặt theo tên của **nhà khoa học máy tính Barbara Liskov** (1987).

> **"Objects of a superclass should be replaceable with objects of a subclass without breaking the application"**
> 
> **"Các đối tượng của class con phải có thể THAY THẾ cho các đối tượng của class cha mà không làm thay đổi tính đúng đắn của chương trình"**

**Nói đơn giản:**

- Class con phải có thể **thay thế** class cha mà không gây lỗi
- Class con **không được phá vỡ** các hành vi của class cha
- Nếu code expect class cha → class con phải hoạt động đúng

---

## Tại sao cần LSP?

Khi **KHÔNG tuân theo LSP**:

- ❌ Class con throw exception khi gọi method của class cha
- ❌ Phải check type trước khi sử dụng (if-else checks)
- ❌ Code không đáng tin cậy, dễ gây lỗi runtime
- ❌ Vi phạm nguyên tắc đa hình (polymorphism)

Khi **tuân theo LSP**:

- ✅ Class con có thể thay thế class cha an toàn
- ✅ Code đáng tin cậy, không cần check type
- ✅ Tận dụng được đa hình
- ✅ Dễ mở rộng và bảo trì

---

## Ví dụ thực tế

> **Ví dụ phương tiện giao thông:**
> 
> ❌ **Vi phạm LSP:**
> 
> - Bạn có class `Vehicle` (phương tiện) với method `StartEngine()`
> - Tạo class `Bicycle` (xe đạp) kế thừa `Vehicle`
> - Nhưng xe đạp **không có động cơ** → `StartEngine()` throw exception
> - Khi code expect `Vehicle` và nhận `Bicycle` → **LỖI!**
> 
> ✅ **Tuân theo LSP:**
> 
> - Tạo interface/abstract class phù hợp với từng loại
> - `MotorVehicle` (có động cơ) → `Car`, `Motorcycle`
> - `ManualVehicle` (không động cơ) → `Bicycle`, `Skateboard`
> - Mỗi class con **đáp ứng đầy đủ** contract của class cha

---

## Ví dụ kinh điển: Rectangle - Square Problem

### ❌ Vi phạm LSP (C#)

```csharp
namespace LSP
{
    // Class cha - Rectangle (Hình chữ nhật)
    public class Rectangle
    {
        public virtual int Width { get; set; }
        public virtual int Height { get; set; }

        public int GetArea()
        {
            return Width * Height;
        }
    }

    // Class con - Square (Hình vuông) kế thừa Rectangle
    // VI PHẠM LSP: Square có ràng buộc Width = Height
    public class Square : Rectangle
    {
        private int _side;

        public override int Width
        {
            get { return _side; }
            set { _side = value; }  // Set cả Width và Height
        }

        public override int Height
        {
            get { return _side; }
            set { _side = value; }  // Set cả Width và Height
        }
    }

    // Test code - VI PHẠM LSP
    class Program
    {
        static void Main()
        {
            Rectangle rect = new Rectangle();
            rect.Width = 5;
            rect.Height = 10;
            Console.WriteLine($"Rectangle Area: {rect.GetArea()}"); // 50 ✓

            // Thay thế Rectangle bằng Square
            Rectangle square = new Square();  // LSP: Square thay thế Rectangle
            square.Width = 5;
            square.Height = 10;
            Console.WriteLine($"Square Area: {square.GetArea()}"); // 100 ✗
            // Expect: 50, nhưng Result: 100 → VI PHẠM LSP!
        }
    }
}
```

**Vấn đề:**

- `Square` **không thể thay thế** `Rectangle` một cách an toàn
- Khi set `Width = 5`, `Height = 10` → Square vẫn có area = 100 (10×10)
- Vi phạm LSP: hành vi không như mong đợi

### ✅ Tuân theo LSP (C#)

```csharp
namespace LSP
{
    // Interface chung cho các hình
    public interface IShape
    {
        int GetArea();
    }

    // Rectangle - độc lập
    public class Rectangle : IShape
    {
        public int Width { get; set; }
        public int Height { get; set; }

        public int GetArea()
        {
            return Width * Height;
        }
    }

    // Square - độc lập, KHÔNG kế thừa Rectangle
    public class Square : IShape
    {
        public int Side { get; set; }

        public int GetArea()
        {
            return Side * Side;
        }
    }

    // Sử dụng
    class Program
    {
        static void Main()
        {
            IShape rect = new Rectangle { Width = 5, Height = 10 };
            Console.WriteLine($"Rectangle Area: {rect.GetArea()}"); // 50 ✓

            IShape square = new Square { Side = 10 };
            Console.WriteLine($"Square Area: {square.GetArea()}"); // 100 ✓

            // Không có confusion, mỗi shape có cách dùng riêng
        }

        // Method nhận IShape - hoạt động với mọi shape
        static void PrintArea(IShape shape)
        {
            Console.WriteLine($"Area: {shape.GetArea()}");
        }
    }
}
```

**Lợi ích:**

- `Rectangle` và `Square` **độc lập**
- Không có confusion về hành vi
- Mỗi class có contract rõ ràng

---

## Ví dụ kinh điển: Rectangle - Square Problem (Golang)

### ❌ Vi phạm LSP (Golang)

```go
package main

import "fmt"

// Struct Rectangle
type Rectangle struct {
    width  int
    height int
}

func (r *Rectangle) SetWidth(width int) {
    r.width = width
}

func (r *Rectangle) SetHeight(height int) {
    r.height = height
}

func (r *Rectangle) GetArea() int {
    return r.width * r.height
}

// Struct Square kế thừa behavior của Rectangle
// VI PHẠM LSP: Square có ràng buộc width = height
type Square struct {
    Rectangle
}

func (s *Square) SetWidth(width int) {
    s.width = width
    s.height = width  // Force cả width và height bằng nhau
}

func (s *Square) SetHeight(height int) {
    s.width = height  // Force cả width và height bằng nhau
    s.height = height
}

// Test - VI PHẠM LSP
func main() {
    rect := &Rectangle{}
    rect.SetWidth(5)
    rect.SetHeight(10)
    fmt.Printf("Rectangle Area: %d\n", rect.GetArea()) // 50 ✓

    // Thay thế Rectangle bằng Square
    var rect2 *Rectangle = &Rectangle{} // Expect Rectangle behavior
    square := &Square{}
    rect2 = &square.Rectangle // "Thay thế" Rectangle bằng Square
    
    square.SetWidth(5)
    square.SetHeight(10)
    fmt.Printf("Square Area: %d\n", square.GetArea()) // 100 ✗
    // Expect: 50, nhưng Result: 100 → VI PHẠM LSP!
}
```

**Vấn đề:**

- `Square` không thể thay thế `Rectangle` đúng cách
- Hành vi không như mong đợi

### ✅ Tuân theo LSP (Golang)

```go
package main

import "fmt"

// Interface chung cho các hình
type Shape interface {
    GetArea() int
}

// Rectangle - độc lập
type Rectangle struct {
    Width  int
    Height int
}

func (r *Rectangle) GetArea() int {
    return r.Width * r.Height
}

// Square - độc lập, KHÔNG kế thừa Rectangle
type Square struct {
    Side int
}

func (s *Square) GetArea() int {
    return s.Side * s.Side
}

// Sử dụng
func main() {
    var rect Shape = &Rectangle{Width: 5, Height: 10}
    fmt.Printf("Rectangle Area: %d\n", rect.GetArea()) // 50 ✓

    var square Shape = &Square{Side: 10}
    fmt.Printf("Square Area: %d\n", square.GetArea()) // 100 ✓

    // Không có confusion, mỗi shape có cách dùng riêng
    PrintArea(rect)
    PrintArea(square)
}

// Function nhận Shape - hoạt động với mọi shape
func PrintArea(shape Shape) {
    fmt.Printf("Area: %d\n", shape.GetArea())
}
```

**Lợi ích:**

- Mỗi shape độc lập
- Contract rõ ràng qua interface

---

## Ví dụ thực tế: Bird - Penguin Problem

### ❌ Vi phạm LSP

**C#:**

```csharp
// VI PHẠM LSP
public class Bird
{
    public virtual void Fly()
    {
        Console.WriteLine("Flying...");
    }
}

public class Sparrow : Bird
{
    public override void Fly()
    {
        Console.WriteLine("Sparrow is flying");
    }
}

// PENGUIN KHÔNG BAY ĐƯỢC!
public class Penguin : Bird
{
    public override void Fly()
    {
        throw new NotImplementedException("Penguins can't fly!");
        // VI PHẠM LSP: throw exception
    }
}

// Test
void MakeBirdFly(Bird bird)
{
    bird.Fly(); // Nếu bird là Penguin → EXCEPTION!
}
```

**Golang:**

```go
// VI PHẠM LSP
type Bird interface {
    Fly()
}

type Sparrow struct{}

func (s *Sparrow) Fly() {
    fmt.Println("Sparrow is flying")
}

type Penguin struct{}

func (p *Penguin) Fly() {
    panic("Penguins can't fly!") // VI PHẠM LSP: panic
}

// Test
func MakeBirdFly(bird Bird) {
    bird.Fly() // Nếu bird là Penguin → PANIC!
}
```

### ✅ Tuân theo LSP

**C#:**

```csharp
// TUÂN THEO LSP
public interface IBird
{
    void Eat();
}

public interface IFlyingBird : IBird
{
    void Fly();
}

public class Sparrow : IFlyingBird
{
    public void Eat() { Console.WriteLine("Eating..."); }
    public void Fly() { Console.WriteLine("Sparrow is flying"); }
}

public class Penguin : IBird  // Chỉ implement IBird, KHÔNG implement IFlyingBird
{
    public void Eat() { Console.WriteLine("Eating..."); }
    public void Swim() { Console.WriteLine("Penguin is swimming"); }
}

// Sử dụng
void MakeFlyingBirdFly(IFlyingBird bird)
{
    bird.Fly(); // An toàn, chỉ nhận birds có thể bay
}

void FeedBird(IBird bird)
{
    bird.Eat(); // Hoạt động với mọi bird
}
```

**Golang:**

```go
// TUÂN THEO LSP
type Bird interface {
    Eat()
}

type FlyingBird interface {
    Bird
    Fly()
}

type Sparrow struct{}

func (s *Sparrow) Eat() { fmt.Println("Eating...") }
func (s *Sparrow) Fly() { fmt.Println("Sparrow is flying") }

type Penguin struct{}

func (p *Penguin) Eat() { fmt.Println("Eating...") }
func (p *Penguin) Swim() { fmt.Println("Penguin is swimming") }

// Sử dụng
func MakeFlyingBirdFly(bird FlyingBird) {
    bird.Fly() // An toàn, chỉ nhận birds có thể bay
}

func FeedBird(bird Bird) {
    bird.Eat() // Hoạt động với mọi bird
}
```

---

## Ví dụ Employee System

### ❌ Vi phạm LSP (C#)

```csharp
public abstract class Employee
{
    public abstract string GetEmployeeDetails(int employeeId);
    public abstract string GetProjectDetails(int employeeId);
}

public class PermanentEmployee : Employee
{
    public override string GetEmployeeDetails(int employeeId)
    {
        return "Permanent Employee Details";
    }

    public override string GetProjectDetails(int employeeId)
    {
        return "Permanent Employee Projects";
    }
}

// VI PHẠM LSP: ContractEmployee không có project
public class ContractEmployee : Employee
{
    public override string GetEmployeeDetails(int employeeId)
    {
        return "Contract Employee Details";
    }

    // VI PHẠM LSP: throw exception
    public override string GetProjectDetails(int employeeId)
    {
        throw new NotImplementedException("Contract employees don't have projects");
    }
}

// Test
List<Employee> employees = new List<Employee>
{
    new PermanentEmployee(),
    new ContractEmployee()
};

foreach (var emp in employees)
{
    emp.GetProjectDetails(1); // Exception với ContractEmployee!
}
```

### ✅ Tuân theo LSP (C#)

```csharp
// Interface cho employee base
public interface IEmployee
{
    string GetEmployeeDetails(int employeeId);
}

// Interface cho employees có project
public interface IProjectEmployee : IEmployee
{
    string GetProjectDetails(int employeeId);
}

public class PermanentEmployee : IProjectEmployee
{
    public string GetEmployeeDetails(int employeeId)
    {
        return "Permanent Employee Details";
    }

    public string GetProjectDetails(int employeeId)
    {
        return "Permanent Employee Projects";
    }
}

// Contract employee chỉ implement IEmployee
public class ContractEmployee : IEmployee
{
    public string GetEmployeeDetails(int employeeId)
    {
        return "Contract Employee Details";
    }
    // Không có GetProjectDetails - không vi phạm LSP
}

// Sử dụng
void ProcessProjectEmployees(List<IProjectEmployee> employees)
{
    foreach (var emp in employees)
    {
        emp.GetProjectDetails(1); // An toàn, tất cả đều có projects
    }
}

void ProcessAllEmployees(List<IEmployee> employees)
{
    foreach (var emp in employees)
    {
        emp.GetEmployeeDetails(1); // Hoạt động với mọi employee
    }
}
```

### ✅ Tuân theo LSP (Golang)

```go
package main

import "fmt"

// Interface cho employee base
type Employee interface {
    GetEmployeeDetails(employeeID int) string
}

// Interface cho employees có project
type ProjectEmployee interface {
    Employee
    GetProjectDetails(employeeID int) string
}

// Permanent Employee
type PermanentEmployee struct{}

func (p *PermanentEmployee) GetEmployeeDetails(employeeID int) string {
    return "Permanent Employee Details"
}

func (p *PermanentEmployee) GetProjectDetails(employeeID int) string {
    return "Permanent Employee Projects"
}

// Contract Employee - chỉ implement Employee
type ContractEmployee struct{}

func (c *ContractEmployee) GetEmployeeDetails(employeeID int) string {
    return "Contract Employee Details"
}
// Không có GetProjectDetails - không vi phạm LSP

// Sử dụng
func ProcessProjectEmployees(employees []ProjectEmployee) {
    for _, emp := range employees {
        fmt.Println(emp.GetProjectDetails(1)) // An toàn
    }
}

func ProcessAllEmployees(employees []Employee) {
    for _, emp := range employees {
        fmt.Println(emp.GetEmployeeDetails(1)) // Hoạt động với mọi employee
    }
}

func main() {
    permanent := &PermanentEmployee{}
    contract := &ContractEmployee{}

    // Sử dụng đúng type
    ProcessProjectEmployees([]ProjectEmployee{permanent})
    ProcessAllEmployees([]Employee{permanent, contract})
}
```

---

## Nguyên tắc để tuân theo LSP

### 1. Contract Rules (Quy tắc hợp đồng)

**Preconditions (Điều kiện tiên quyết):**

- Class con **không được** yêu cầu điều kiện đầu vào **mạnh hơn** class cha
- **Ví dụ:** Nếu class cha chấp nhận `age >= 0`, class con không được yêu cầu `age >= 18`

**Postconditions (Điều kiện hậu quyết):**

- Class con **phải** đảm bảo kết quả **ít nhất bằng** class cha
- **Ví dụ:** Nếu class cha return `non-null`, class con không được return `null`

### 2. Không throw Exception bất ngờ

```csharp
// ❌ VI PHẠM
public class Parent
{
    public virtual void DoSomething() { }
}

public class Child : Parent
{
    public override void DoSomething()
    {
        throw new NotImplementedException(); // SAII
    }
}

// ✅ ĐÚNG
public class Child : Parent
{
    public override void DoSomething()
    {
        // Implement đầy đủ, không throw exception
    }
}
```

### 3. Không thay đổi hành vi mong đợi

```go
// ❌ VI PHẠM
type Stack interface {
    Push(item int)
    Pop() int
}

type LimitedStack struct {
    items []int
    maxSize int
}

func (s *LimitedStack) Push(item int) {
    if len(s.items) >= s.maxSize {
        panic("Stack overflow") // Thay đổi hành vi - VI PHẠM LSP
    }
    s.items = append(s.items, item)
}

// ✅ ĐÚNG - Reflect limitation trong interface
type LimitedStack interface {
    Push(item int) error // Return error thay vì panic
    Pop() (int, error)
}
```

---

## Kết luận

### Nguyên tắc vàng của LSP:

> **"Subclasses must be substitutable for their base classes"**
> 
> **"Class con phải có thể thay thế class cha mà không gây lỗi"**

### Cách áp dụng LSP:

1. **Design interface/abstract class cẩn thận**
    
    - Chỉ định nghĩa những gì THẬT SỰ chung
    - Không ép buộc class con implement những gì không phù hợp
2. **Tách interface khi cần**
    
    - Nếu không phải tất cả subclass đều cần một method → tách interface
    - ISP (Interface Segregation Principle) hỗ trợ LSP
3. **Không throw exception trong override**
    
    - Nếu phải throw → đó là dấu hiệu vi phạm LSP
    - Cân nhắc lại thiết kế
4. **Test substitutability**
    
    - Nếu code expect base class → test với tất cả subclasses
    - Tất cả phải hoạt động đúng

### Lợi ích:

- Code đáng tin cậy, ít bug
- Tận dụng polymorphism
- Dễ mở rộng
- Maintainable

### Mối quan hệ với các nguyên tắc khác:

- **LSP + OCP:** LSP giúp OCP hoạt động tốt (có thể thay thế subclass)
- **LSP + ISP:** Interface segregation giúp tuân theo LSP dễ dàng hơn