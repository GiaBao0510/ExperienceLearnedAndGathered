# Type Assertion Và Type Switch Trong Go

Trong Go, **interface{}** (empty interface) có thể lưu trữ **bất kỳ giá trị nào** vì nó không yêu cầu method nào. Tuy nhiên, để truy cập giá trị cụ thể, bạn cần kiểm tra hoặc chuyển đổi kiểu (type checking/conversion). Hai công cụ chính là **Type Assertion** và **Type Switch**:

- **Type Assertion**: Kiểm tra và lấy giá trị từ interface{} bằng cách "khẳng định" (assert) kiểu cụ thể. Nếu sai, có thể gây panic.
- **Type Switch**: Kiểm tra kiểu qua cấu trúc switch, an toàn hơn khi xử lý nhiều kiểu.

Chúng giúp xử lý dữ liệu động, như từ JSON, API, hoặc hàm generic (trước Go 1.18). Bổ sung: Go không có runtime type checking như dynamic languages, nên hai công cụ này là cách chính để xử lý interface{}.

### 1. Type Assertion

- **Mô tả**: Chuyển đổi interface{} sang kiểu cụ thể. Cú pháp: `value := i.(T)` (có thể panic nếu sai kiểu) hoặc `value, ok := i.(T)` (an toàn, ok=false nếu sai).
- **Lưu ý**: 
  - Luôn dùng dạng comma-ok để tránh panic.
  - Chỉ dùng với interface{} (hoặc interface có underlying type).
  - Nếu interface là nil, assertion sẽ panic.

**Ví dụ thực tế**: Xử lý giá trị từ map (như dữ liệu JSON).
```go
package main
import "fmt"

func main() {
    var data interface{} = "Hello, world!"  // Có thể là bất kỳ kiểu nào

    // Assertion thành string
    s, ok := data.(string)
    if ok {
        fmt.Println("Là string:", s)
    } else {
        fmt.Println("Không phải string")
    }

    // Assertion thành float64 (sai kiểu)
    f, ok := data.(float64)
    if ok {
        fmt.Println("Là float64:", f)
    } else {
        fmt.Println("Không phải float64")
    }

    // Ví dụ gây panic (không dùng ok)
    // wrong := data.(int)  // Panic: interface conversion: interface {} is string, not int
}
```

**Kết quả**:
```
Là string: Hello, world!
Không phải float64
```

**Ứng dụng**: Khi unmarshal JSON, giá trị có thể là string/int/float, dùng assertion để lấy và xử lý.

### 2. Type Switch

- **Mô tả**: Kiểm tra kiểu của interface{} qua switch. Cú pháp: `switch x := v.(type) { case T: ... }`. Biến `x` sẽ giữ giá trị với kiểu đã assert.
- **Lưu ý**: 
  - Chỉ dùng `. (type)` trong switch (không dùng riêng lẻ).
  - Case có thể là kiểu cụ thể (int, string,...), hoặc interface (nâng cao).
  - Default xử lý kiểu không khớp.
  - An toàn hơn assertion vì không panic, và xử lý nhiều case một lúc.

**Ví dụ thực tế**: Hàm in thông tin dựa trên kiểu (như xử lý dữ liệu đa dạng từ input).
```go
package main
import "fmt"

func CheckType(v interface{}) {
    switch x := v.(type) {
    case int:
        fmt.Printf("Là int: %d\n", x)
    case string:
        fmt.Printf("Là string: %s\n", x)
    case bool:
        fmt.Printf("Là bool: %t\n", x)
    case float64:
        fmt.Printf("Là float64: %f\n", x)
    default:
        fmt.Printf("Kiểu không xác định: %T\n", x)  // %T in kiểu
    }
}

func main() {
    CheckType("Hello, guys")
    CheckType(42)
    CheckType(true)
    CheckType(3.14)
    CheckType('k')  // rune (int32)
}
```

**Kết quả**:
```
Là string: Hello, guys
Là int: 42
Là bool: true
Là float64: 3.140000
Kiểu không xác định: int32
```

**Ứng dụng**: Trong hàm generic xử lý slice/map với phần tử đa kiểu, hoặc debug dữ liệu không xác định.

### So Sánh Và Khi Nào Dùng?

- **Type Assertion**: Dùng khi bạn chắc chắn (hoặc kiểm tra) kiểu cụ thể, chỉ một kiểu. Nhanh nhưng rủi ro panic nếu sai.
- **Type Switch**: Dùng khi cần kiểm tra nhiều kiểu, an toàn hơn, dễ mở rộng (thêm case).
- **Bổ sung**: Với generics (Go 1.18+), ít cần interface{} hơn, nhưng hai công cụ này vẫn quan trọng cho code cũ hoặc dữ liệu động.