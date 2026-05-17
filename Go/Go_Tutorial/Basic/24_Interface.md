# Interface Trong Go

## 1. Tổng Quan

**Interface** là một kiểu dữ liệu trừu tượng trong Go, dùng để định nghĩa một tập hợp các **phương thức (methods)** mà bất kỳ kiểu dữ liệu nào cũng có thể triển khai. Interface không chứa dữ liệu hay logic cụ thể — nó chỉ mô tả **"hành vi" (behavior)**.

Go không có class hay kế thừa như OOP truyền thống, nhưng interface cho phép **composition (kết hợp)** và **type-safe polymorphism** được kiểm tra tại compile-time.

> Rob Pike, cha đẻ của ngôn ngữ Go, từng nói: _"Languages that try to disallow idiocy become themselves idiotic"_
> 
> Ý tưởng: Go tin tưởng lập trình viên — interface được thỏa mãn **tự động** (implicit), không cần khai báo `implements` như Java hay C#.

### Lợi ích chính

|Lợi ích|Mô tả|
|---|---|
|**Decoupling**|Hàm chỉ cần biết interface, không cần biết kiểu cụ thể → dễ test, mock|
|**Type safety**|Go kiểm tra tại compile-time xem kiểu có đủ methods không|
|**Dễ mở rộng**|Thêm kiểu mới mà không cần sửa code cũ|
|**Polymorphism**|Nhiều kiểu khác nhau có thể dùng chung một hàm|

---

## 2. Cú Pháp Cơ Bản

```go
type TênInterface interface {
    Method1(arg KiểuThamSố) KiểuTrảVề
    Method2() // Không tham số, không trả về
}
```

Bất kỳ kiểu nào triển khai **tất cả methods** của interface sẽ tự động thỏa mãn nó — không cần khai báo tường minh như `implements`.

### Quy tắc đặt tên (best practice)

- Nếu interface chỉ có **một method**, nên đặt tên kết thúc bằng `"er"`.
    - Ví dụ: `Reader`, `Writer`, `Stringer`, `Closer`
- Tên method thường trùng với tên interface nhưng bỏ đuôi `"er"`.
    - Ví dụ: interface `Reader` → method `Read()`

---

## 3. Ví Dụ Thực Tế: Interface Cho Các Hình Học

Giả sử chúng ta mô hình hóa các hình học (hình tròn, hình vuông) với hành vi chung: lấy thông tin và ứng dụng thực tế.

### 3.1. Định Nghĩa Struct và Methods

**Package `circle`** — file `circle/circle.go`:

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

// Method riêng của Circle (không có trong Square)
func (c *Circle) PerimeterFormula() string {
    return "C = 2 * π * r"
}
```

**Package `square`** — file `square/square.go`:

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

### 3.2. Định Nghĩa Interface

**Package `services`** — file `services/services.go`:

> **Best practice:** Tách interface ra package riêng để dễ reuse và tránh phụ thuộc vòng.

```go
package services

// Interface cơ bản
type Shape interface {
    Apply() string
    GetInfo() string
}

// Interface mở rộng: nhúng Shape (embedding)
// Tự động kế thừa Apply() và GetInfo() từ Shape
type ShapePlus interface {
    Shape
    PerimeterFormula() string
}
```

### 3.3. Sử Dụng Trong `main.go`

```go
package main

import (
    "fmt"
    "hello/circle"
    "hello/services"
    "hello/square"
)

// Hàm nhận interface cơ bản
func PracticalApplication(s services.Shape) {
    fmt.Printf("Tên: %s\n", s.GetInfo())
    fmt.Printf("Ứng dụng: %s\n\n", s.Apply())
}

// Hàm nhận interface mở rộng
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

    PracticalApplicationPlus(circleObj) // Circle thỏa mãn ShapePlus
    PracticalApplication(squareObj)     // Square chỉ thỏa mãn Shape
}
```

**Kết quả:**

```
Tên: Bánh xe
Ứng dụng: Tạo bánh răng cửa
Công thức chu vi: C = 2 * π * r

Tên: Cánh diều
Ứng dụng: Tạo gạch lát vuông
```

**Giải thích:**

- `Circle` thỏa mãn cả `Shape` lẫn `ShapePlus` (vì có thêm `PerimeterFormula()`).
- `Square` chỉ thỏa mãn `Shape` — nếu truyền vào `PracticalApplicationPlus()` sẽ bị **lỗi compile**.
- Embedding (`ShapePlus` nhúng `Shape`) giúp tránh lặp code, tương tự "kế thừa" interface.

---

## 4. Interface Rỗng (`interface{}` / `any`)

**Interface rỗng** không có method nào, vì vậy **mọi kiểu dữ liệu** đều tự động thỏa mãn nó.

```go
// Hai cách viết tương đương nhau (từ Go 1.18+)
var x interface{} = 42
var y any = "hello"
```

### Ứng dụng phổ biến

- Hàm nhận tham số không xác định kiểu trước
- Lưu dữ liệu JSON chưa biết cấu trúc (`json.Unmarshal`)
- Các cấu trúc dữ liệu generic đơn giản

### Ví dụ

```go
package main

import "fmt"

func PrintInfo(v interface{}) {
    fmt.Println(v)
}

func main() {
    PrintInfo("Hello") // string
    PrintInfo(123)     // int
    PrintInfo(3.14)    // float64

    // Type assertion: lấy giá trị cụ thể từ interface{}
    var x interface{} = 42
    if num, ok := x.(int); ok {
        fmt.Println("Là int:", num*2) // 84
    } else {
        fmt.Println("Không phải int")
    }

    // Type switch: xử lý nhiều kiểu an toàn hơn
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

**Kết quả:**

```
Hello
123
3.14
Là int: 84
Int: 42
```

> **Lưu ý:** Interface rỗng mất đi type safety. Ưu tiên dùng interface cụ thể khi có thể. Từ Go 1.18+, nên dùng **generics** thay thế cho nhiều trường hợp.

---

## 5. Chuyển Đổi Kiểu Với Interface

Go **không hỗ trợ chuyển đổi kiểu ngầm định** với các kiểu cơ bản. Ví dụ, không thể gán trực tiếp `int` cho `int64`:

```go
var a int = 10
var b int64 = a // ❌ Lỗi compile: cannot use a (type int) as type int64
var c int64 = int64(a) // ✅ Phải chuyển đổi tường minh
```

Tuy nhiên, **với interface thì linh hoạt hơn** — Go cho phép chuyển đổi ngầm định giữa object và interface, hoặc giữa hai interface với nhau, miễn là kiểu nguồn thỏa mãn interface đích:

```go
var (
    // Chuyển đổi ngầm định: *os.File thỏa mãn io.ReadCloser
    a io.ReadCloser = (*os.File)(f)

    // Chuyển đổi ngầm định: io.ReadCloser thỏa mãn io.Reader
    b io.Reader = a

    // Chuyển đổi ngầm định: io.ReadCloser thỏa mãn io.Closer
    c io.Closer = a

    // Chuyển đổi tường minh (type assertion): cần khi compiler không chắc chắn
    d io.Reader = c.(io.Reader)
)
```

**Tóm tắt quy tắc:**

|Trường hợp|Loại chuyển đổi|
|---|---|
|Kiểu cơ bản → kiểu cơ bản khác|**Tường minh** (bắt buộc)|
|Object → interface (object thỏa mãn interface)|**Ngầm định**|
|Interface → interface rộng hơn|**Ngầm định**|
|Interface → interface hẹp hơn / cụ thể hơn|**Tường minh** (type assertion)|

---

## 6. Nhúng Interface Vào Struct (Interface Embedding Trong Struct)

Không chỉ interface mới có thể nhúng interface khác — **struct cũng có thể nhúng interface** như một trường ẩn danh. Kỹ thuật này cho phép struct "kế thừa" tất cả methods của interface đó.

### Ví dụ: Làm giả phương thức `private` của `testing.TB`

```go
package main

import (
    "fmt"
    "testing"
)

// TB nhúng interface testing.TB như trường ẩn danh
// → tự động "kế thừa" tất cả methods của testing.TB
// (kể cả các method private chỉ dùng nội bộ package testing)
type TB struct {
    testing.TB
}

// Override method Fatal của testing.TB
func (p *TB) Fatal(args ...interface{}) {
    fmt.Println("TB.Fatal disabled!")
}

func main() {
    var tb testing.TB = new(TB)
    tb.Fatal("Hello, playground") // Gọi Fatal đã override
}
```

**Kết quả:**

```
TB.Fatal disabled!
```

**Giải thích:** Bằng cách nhúng `testing.TB` vào struct `TB`, ta "thừa hưởng" toàn bộ interface đó và có thể override từng method theo ý muốn. Đây là kỹ thuật hữu ích khi viết mock object trong unit test.

---

## 7. Bảo Vệ Interface Khỏi Bị "Giả Mạo"

Vì interface trong Go được thỏa mãn tự động (implicit), đôi khi một struct **vô tình** thỏa mãn một interface mà không cố ý. Điều này có thể gây ra lỗi khó phát hiện.

### Giải pháp: Thêm method "dấu hiệu" (marker method)

Định nghĩa một method rỗng với tên đặc biệt để phân biệt interface:

```go
// runtime.Error chỉ được thỏa mãn bởi các kiểu
// có method RuntimeError() — tránh nhầm với error thông thường
type runtime.Error interface {
    error
    RuntimeError() // Method rỗng, chỉ dùng để phân biệt
}
```

Trong protobuf, `proto.Message` áp dụng tương tự:

```go
type proto.Message interface {
    Reset()
    String() string
    ProtoMessage() // Marker method — chỉ protobuf mới implement
}
```

**Lợi ích:** Chỉ những kiểu **cố ý** implement marker method mới thỏa mãn interface, tránh việc struct khác vô tình "giả mạo" interface.

---

## 8. Quản Lý Interface Hiệu Quả — Best Practices

|Thực hành|Lý do|
|---|---|
|Dùng **pointer receiver** (`*T`)|Sửa được dữ liệu gốc, tránh copy struct lớn|
|**Constructor + error handling**|Validation sớm, tránh object ở trạng thái không hợp lệ|
|**Tách interface ra package riêng**|Dễ reuse, tránh phụ thuộc vòng (circular import)|
|**Embedding** thay vì copy method|Khi mở rộng interface, nhúng interface cũ vào mới|
|**Không lạm dụng `interface{}`**|Ưu tiên interface cụ thể để giữ type safety|
|**Dùng type assertion có kiểm tra** (`val, ok := x.(T)`)|Tránh panic khi assertion sai kiểu|

---

## 9. Lưu Ý Quan Trọng

- **Interface được kiểm tra tại compile-time**: Nếu struct thiếu method, code không build được.
- **Go 1.18+ có `any`** là alias của `interface{}` — hai cách viết hoàn toàn tương đương.
- **Go 1.18+ có Generics**: Nhiều trường hợp dùng `interface{}` trước đây nay nên thay bằng generics để an toàn kiểu hơn.
- **Khi `main()` kết thúc, toàn bộ chương trình dừng** — bao gồm cả các Goroutine đang chạy.