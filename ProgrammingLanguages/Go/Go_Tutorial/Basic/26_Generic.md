# Generics Trong Go (Từ Phiên Bản 1.18+)

**Generics** (tính năng tổng quát) cho phép viết hàm, struct hoặc interface có thể làm việc với **nhiều kiểu dữ liệu khác nhau** mà không cần lặp lại code. Trước Go 1.18, bạn phải viết hàm riêng cho từng kiểu (ví dụ: hàm riêng cho int, string...). Generics giúp code ngắn gọn, tái sử dụng cao và type-safe (kiểm tra kiểu tại compile-time).

![]()https://200lab.io/blog/_next/image?url=https%3A%2F%2Fstatics.cdn.200lab.io%2F2021%2F12%2Fgo-generics.png&w=640&q=75

**Lợi ích**:
- Giảm duplicate code.
- An toàn kiểu dữ liệu (lỗi phát hiện sớm).
- Hiệu suất tốt (không boxing như Java).

**Cú pháp cơ bản**:
- `[T Constraint]` là type parameter (T là tên placeholder, Constraint giới hạn kiểu).
- Constraint phổ biến: `any` (bất kỳ kiểu nào, tương đương interface{}), `comparable` (hỗ trợ == và !=).

### 1. Hàm Generic Cơ Bản

![](https://images.squarespace-cdn.com/content/v1/5e10bdc20efb8f0d169f85f9/2262dafd-1b89-46f2-8642-6129983c7160/generics-fish.png?format=750w)

**Ví dụ**: Hàm in và trả về giá trị bất kỳ.
```go
package main
import "fmt"

func PrintValue[T any](value T) T {
    fmt.Println("Giá trị:", value)
    return value
}

func main() {
    PrintValue("Hello All")  // T là string
    PrintValue(123)          // T là int
    PrintValue(true)         // T là bool
}
```

**Kết quả**:
```
Giá trị: Hello All
Giá trị: 123
Giá trị: true
```

### 2. Constraint `comparable` (Cho So Sánh == và !=)

Dùng khi cần so sánh bằng.

**Ví dụ**:
```go
package main
import "fmt"

func IsEqual[T comparable](a, b T) bool {
    return a == b
}

func main() {
    fmt.Println(IsEqual(5, 5))          // true
    fmt.Println(IsEqual("fmt", "fmt"))  // true
    fmt.Println(IsEqual(3.14, 3.14))    // true
    fmt.Println(IsEqual(true, false))   // false
}
```

**Lưu ý**: Slice, map, func **không** comparable (không hỗ trợ ==).

### 3. So Sánh Lớn/Nhỏ Hơn (Dùng `cmp.Ordered` – Từ Go 1.21)

`cmp.Ordered` hỗ trợ các toán tử >, <, >=, <= (int, float, string...).

**Ví dụ**: Tìm số lớn hơn.
```go
package main
import (
    "cmp"
    "fmt"
)

func Max[T cmp.Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}

func main() {
    fmt.Println(Max(10.0001, 10.0002))  // 10.0002
    fmt.Println(Max("apple", "banana")) // banana
    fmt.Println(Max(5, 10))            // 10
}
```

### 4. Struct Generic

**Ví dụ**: Hộp chứa bất kỳ giá trị nào.
```go
package main
import "fmt"

type Box[T any] struct {
    Value T
}

func main() {
    stringBox := Box[string]{Value: "Hi Guy"}
    fmt.Println(stringBox.Value)  // Hi Guy

    intBox := Box[int]{Value: 105}
    fmt.Println(intBox.Value)     // 105
}
```

**Ứng dụng thực tế**: Stack generic, Pair<K, V> như map entry.

### 5. Tự Định Nghĩa Constraint (Type Union)

Dùng `|` để giới hạn các kiểu cụ thể (không cần package ngoài).

**Ví dụ**: Chỉ cho phép số (int hoặc float64).
```go
package main
import "fmt"

type Number interface {
    int | float64
}

func Sum[T Number](a, b T) T {
    return a + b
}

func main() {
    fmt.Println(Sum(3, 5))       // 8 (int)
    fmt.Println(Sum(15.5, 6.9))  // 22.4 (float64)
    // Sum("a", "b") // Lỗi compile: string không thỏa Number
}
```

**Ví dụ mở rộng** (thêm float32, int64):
```go
type Numeric interface {
    int | int64 | float32 | float64
}
```

### Ví Dụ Thực Tế: Hàm Generic Với Slice

**Map** – Biến đổi từng phần tử slice.
```go
package main
import "fmt"

func MapSlice[T any, U any](slice []T, f func(T) U) []U {
    result := make([]U, len(slice))
    for i, v := range slice {
        result[i] = f(v)
    }
    return result
}

func main() {
    nums := []int{1, 2, 3}
    doubled := MapSlice(nums, func(x int) int { return x * 2 })
    fmt.Println(doubled)  // [2 4 6]

    strs := []string{"a", "b"}
    upper := MapSlice(strs, func(s string) string { return strings.ToUpper(s) })
    fmt.Println(upper)    // [A B] (cần import "strings")
}
```

### Lưu Ý Quan Trọng Cho Sinh Viên

- **Khi nào dùng**: Khi code lặp lại cho nhiều kiểu (slice functions, container như stack/queue, utils).
- **Không lạm dụng**: Generics làm code phức tạp hơn, khó đọc nếu dùng quá nhiều. Ưu tiên code đơn giản trước.
- **Instantiation**: Go tự suy luận kiểu khi gọi hàm (như ví dụ trên). Có thể chỉ định rõ: `PrintValue[int](123)`.
- **Hạn chế**: Không hỗ trợ specialization (hàm riêng cho kiểu cụ thể). Một số kiểu (chan, func) có hạn chế.
- **Bổ sung**: Từ Go 1.21+, có `cmp` package cho so sánh tiện lợi. Thực hành với `go mod init` và chạy `go run .` (Go >= 1.18).
- **Best practice**: Giới hạn constraint chặt chẽ để tránh lỗi runtime ẩn.