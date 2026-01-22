# Interface Trong Go

**Interface** là một kiểu dữ liệu trừu tượng trong Go, định nghĩa một tập hợp các phương thức (methods) mà bất kỳ kiểu dữ liệu nào cũng có thể triển khai để thỏa mãn interface đó. Interface không chứa dữ liệu hay triển khai cụ thể – nó chỉ mô tả "hành vi" (behavior). Điều này hỗ trợ **polymorphism** (đa hình), giúp code linh hoạt, dễ mở rộng và decoupling (tách rời) giữa các phần code. Không giống OOP truyền thống, Go không có class hay kế thừa, nhưng interface cho phép composition (kết hợp) và **type-safe polymorphism** tại compile-time.

**Lợi ích chính**:
- **Decoupling**: Các hàm chỉ cần biết interface, không cần biết kiểu cụ thể (dễ test, mock).
- **Type safety**: Go kiểm tra tại **compile-time** xem kiểu có triển khai đủ methods không.
- **Dễ mở rộng**: Thêm kiểu mới mà không sửa code cũ.

**Lưu ý tên gọi** (best practice):
- Nếu interface chỉ có **một method**, tên interface nên kết thúc bằng "er" (ví dụ: `Reader`, `Writer`).
- Tên method nên giống tên interface nhưng bỏ "er" (ví dụ: method `Read()` cho `Reader`).

### Cú Pháp Cơ Bản

```go
type TênInterface interface {
    Method1(arg kiểu) kiểu_trả_về
    Method2()  // Không tham số
}
```

Bất kỳ kiểu nào (thường là struct) triển khai **tất cả methods** của interface sẽ tự động thỏa mãn nó (không cần khai báo explicit như "implements").

### Ví Dụ Thực Tế: Interface Cho Các Hình Học

Giả sử chúng ta mô hình hóa các hình (circle, square) với hành vi chung: Lấy thông tin và áp dụng thực tế.

#### 1. Định nghĩa Struct Và Methods

Tạo package `circle` (file `circle/circle.go`):
```go
package circle

import (
    "errors"
    "strings"
)

type Circle struct {
    Name string `json:"ten_hinh_tron"`
}

// Constructor với validation
func New(name string) (*Circle, error) {
    name = strings.TrimSpace(name)
    if name == "" {
        return nil, errors.New("Tên không được để trống")
    }
    if len(name) > 255 {
        return nil, errors.New("Số ký tự tối đa là 255")
    }
    return &Circle{Name: name}, nil
}

func (c *Circle) GetInfo() string {
    return c.Name
}

func (c *Circle) Apply() string {
    return "Tạo bánh răng cửa"
}

// Method thêm cho circle
func (c *Circle) PerimeterFormula() string {
    return "C = 2 * π * r"
}
```

Tạo package `square` (file `square/square.go`): Tương tự, nhưng không có `PerimeterFormula()`.

```go
package square

import (
    "errors"
    "strings"
)

type Square struct {
    Name string `json:"ten_hinh_vuong"`
}

func New(name string) (*Square, error) {
    name = strings.TrimSpace(name)
    if name == "" {
        return nil, errors.New("Tên không được để trống")
    }
    if len(name) > 255 {
        return nil, errors.New("Số ký tự tối đa là 255")
    }
    return &Square{Name: name}, nil
}

func (s *Square) GetInfo() string {
    return s.Name
}

func (s *Square) Apply() string {
    return "Tạo gạch lát vuông"
}
```

#### 2. Định Nghĩa Interface

Tạo package `services` (file `services/services.go`) để quản lý interface riêng biệt (best practice: Tách interface khỏi implementation để dễ reuse).

```go
package services

type Shape interface {
    Apply() string
    GetInfo() string
}

// Interface mở rộng (embedding): Kế thừa từ Shape
type ShapePlus interface {
    Shape  // Nhúng Shape, tự động có Apply() và GetInfo()
    PerimeterFormula() string
}
```

#### 3. Sử Dụng Trong Main (file `main.go`)

Giả sử `go.mod`: `module hello` (và đã `go get` nếu cần).

```go
package main

import (
    "fmt"
    "hello/circle"
    "hello/services"
    "hello/square"
)

// Hàm sử dụng interface cơ bản
func PracticalApplication(s services.Shape) {
    fmt.Printf("Tên: %s\n", s.GetInfo())
    fmt.Printf("Ứng dụng: %s\n\n", s.Apply())
}

// Hàm sử dụng interface mở rộng
func PracticalApplicationPlus(sp services.ShapePlus) {
    fmt.Printf("Tên: %s\n", sp.GetInfo())
    fmt.Printf("Ứng dụng: %s\n", sp.Apply())
    fmt.Printf("Công thức chu vi: %s\n\n", sp.PerimeterFormula())
}

func main() {
    circleObj, err := circle.New("Bánh xe")
    if err != nil {
        panic(err)
    }

    squareObj, err := square.New("Cánh diều")
    if err != nil {
        panic(err)
    }

    PracticalApplicationPlus(circleObj)  // Circle có PerimeterFormula()
    PracticalApplication(squareObj)      // Square chỉ dùng Shape
}
```

**Kết quả**:
```
Tên: Bánh xe
Ứng dụng: Tạo bánh răng cửa
Công thức chu vi: C = 2 * π * r

Tên: Cánh diều
Ứng dụng: Tạo gạch lát vuông
```

**Giải thích**: 
- Circle thỏa mãn cả `Shape` và `ShapePlus` (vì có thêm method).
- Square chỉ thỏa mãn `Shape` (không có `PerimeterFormula()` → compile error nếu dùng `ShapePlus`).
- Embedding (`ShapePlus` nhúng `Shape`) giúp tránh lặp code, giống "kế thừa" interface.

### Quản Lý Interface Hiệu Quả

- **Sử dụng pointer receiver**: Để sửa dữ liệu gốc và tránh copy lớn (như trong ví dụ).
- **Constructor với error handling**: Đảm bảo validation sớm (trim space, check length).
- **Tách package**: Interface ở package riêng (services) để dễ import và reuse.
- **Embedding**: Khi cần mở rộng, nhúng interface cũ vào mới thay vì copy methods.
- **Type assertion (bổ sung)**: Để kiểm tra kiểu cụ thể tại runtime nếu cần.
  Ví dụ: `if plus, ok := s.(services.ShapePlus); ok { fmt.Println(plus.PerimeterFormula()) }`.

---
### Empty Interface (interface{})

- **Empty interface** (`interface{}`) không có method nào, nên có thể chứa **bất kỳ giá trị nào** (tương tự `any` trong Go 1.18+ hoặc `Object` ở Java).
- **Ứng dụng**: Hàm generic đơn giản, lưu dữ liệu không biết kiểu trước (như JSON unmarshal).
- **Nhược điểm**: Mất type safety → cần type assertion để lấy giá trị cụ thể, có thể panic nếu sai kiểu.

**Ví dụ**:
```go
package main
import "fmt"

func PrintInfo(v interface{}) {  // Hoặc any từ Go 1.18
    fmt.Println(v)
}

func main() {
    PrintInfo("Hello")  // string
    PrintInfo(123)      // int
    PrintInfo(3.14)     // float64

    // Type assertion: Lấy giá trị cụ thể
    var x interface{} = 42
    if num, ok := x.(int); ok {
        fmt.Println("Là int:", num * 2)  // 84
    } else {
        fmt.Println("Không phải int")
    }

    // Type switch: Xử lý nhiều kiểu
    switch val := x.(type) {
    case int:
        fmt.Println("Int:", val)
    case string:
        fmt.Println("String:", val)
    default:
        fmt.Println("Kiểu khác")
    }
}
```

**Kết quả**:
```
Hello
123
3.14
Là int: 84
Int: 42
```

**Bổ sung: Type Switch**: Dùng để xử lý empty interface an toàn hơn assertion.

### Lưu Ý Quan Trọng

- ==**Interface** được kiểm tra tại compile-time==: Nếu struct thiếu method, code không build.
- Không overusing empty interface: Ưu tiên interface cụ thể để giữ type safety.
- Go 1.18+ có generics (`any` thay `interface{}`), nhưng interface vẫn mạnh cho polymorphism.
- Thực hành: Tạo interface `Animal` với methods `Speak()` và `Move()`, implement cho `Dog` và `Cat`.
- Debug: Nếu panic, kiểm tra nil pointer hoặc type assertion sai.