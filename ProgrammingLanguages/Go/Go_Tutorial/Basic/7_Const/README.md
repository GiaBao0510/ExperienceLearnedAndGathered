# Hằng số trong Golang

## 1. Hằng số là gì?

- Hằng số (`const`) là một vùng bộ nhớ dùng để lưu trữ dữ liệu, tương tự biến, nhưng giá trị của hằng số **không thể thay đổi** sau khi được khai báo.
- Hằng số chỉ hỗ trợ các kiểu dữ liệu cơ bản (`int`, `float`, `string`, `bool`).
- Hằng số được tính toán tại thời điểm biên dịch (compile-time), không thể gán giá trị từ hàm hoặc biểu thức runtime.
- **Quy ước đặt tên**:
    - Hằng số xuất (exported): Viết hoa chữ cái đầu (ví dụ: `MaxRetries`).
    - Hằng số không xuất (unexported): Dùng `camelCase` (ví dụ: `appVersion`).
    - Tránh dùng toàn bộ chữ hoa (ví dụ: `FULLNAME`) trừ khi theo quy ước cụ thể của dự án.

---
## 2. Khai báo hằng số

- **Cú pháp**:
```go
const constantName type = value
```

- **Đặc điểm**:
    - Có thể khai báo mà không cần chỉ định kiểu dữ liệu (type inference).
    - Phải gán giá trị tại thời điểm khai báo, không thể để trống.
    - Có thể khai báo trong hoặc ngoài hàm (phạm vi toàn cục hoặc cục bộ).
    - Có thể khai báo nhiều hằng số cùng lúc trong khối `const`.

- **Ví dụ khai báo đơn**:
```go
const Pi float32 = 3.14
const AppName = "MyApp" // Type inference: string
```

- **Ví dụ khai báo khối**:
```go
const (
    MaxRetries = 3
    AppVersion = "1.0.0"
    DebugMode  = false
)
```

---
## 3. Phạm vi hằng số

- Hằng số khai báo ngoài hàm (toàn cục) có thể được truy cập trong toàn bộ package.
- Hằng số khai báo trong hàm (cục bộ) chỉ tồn tại trong phạm vi hàm đó.
- Nếu hằng số trong hàm trùng tên với hằng số toàn cục, hằng số trong hàm sẽ che giấu (shadow) hằng số toàn cục mà không gây lỗi.
- **Ví dụ**:

```go
const Pi = "3.14" // Toàn cục, type inference: string

func main() {
    const Pi float32 = 3.14 // Cục bộ, che giấu hằng số toàn cục
    fmt.Println("Hằng số Pi trong hàm:", Pi) // In: 3.14
    fmt.Println("Hằng số Pi toàn cục:", Pi) // Lỗi nếu cố truy cập trực tiếp
}
```

- **Lưu ý**: Nếu khai báo hai hằng số trùng tên trong cùng một phạm vi (ví dụ: trong cùng một hàm hoặc khối `const`), sẽ gây lỗi biên dịch: `redeclared in this block`.

---
## 4. Hằng số không kiểu (Untyped Constants)

- Hằng số không kiểu là hằng số không chỉ định kiểu dữ liệu, cho phép sử dụng linh hoạt với các kiểu tương thích.
- **Ví dụ**:

```go
const Pi = 3.14
var x float32 = Pi // Pi tự động chuyển thành float32
var y float64 = Pi // Pi tự động chuyển thành float64
fmt.Println(x, y) // In: 3.14 3.14
```

---
## 5. Lưu ý quan trọng

- **Không thể gán lại giá trị**: Cố gắng thay đổi giá trị hằng số sẽ gây lỗi biên dịch.
    ```go
    const Pi = 3.14
    Pi = 3.14159 // Lỗi: cannot assign to Pi
    ```
    
- **Chỉ hỗ trợ kiểu cơ bản**: Hằng số không thể là slice, map, hoặc struct.
- **Trường hợp sử dụng**:
    - Dùng hằng số cho các giá trị cố định (ví dụ: số Pi, số lần thử tối đa, tên ứng dụng).
    - Giúp mã dễ đọc, bảo trì, và tránh thay đổi giá trị ngoài ý muốn.

---
## 6. Ví dụ tổng hợp

```go
package main

import "fmt"

const MaxRetries = 3 // Exported constant

const (
    AppVersion = "1.0.0" // Unexported constant
    DebugMode  = false
)

func main() {
    const Pi float32 = 3.14 // Cục bộ
    fmt.Println("Hằng số Pi:", Pi)
    fmt.Println("Max Retries:", MaxRetries)
    fmt.Println("App Version:", AppVersion)
    fmt.Println("Debug Mode:", DebugMode)

    // Untyped constant
    const Value = 42
    var x int = Value
    var y float64 = Value
    fmt.Println("Untyped constant:", x, y)
}
```

- **Kết quả**:
```shell
Hằng số Pi: 3.14
Max Retries: 3
App Version: 1.0.0
Debug Mode: false
Untyped constant: 42 42
```