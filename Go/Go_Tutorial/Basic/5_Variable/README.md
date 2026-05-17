![](https://miro.medium.com/v2/resize:fit:720/format:webp/1*9TfM27i8XsSnTHfAEA92Ww.png)
# Biến trong Golang

## Biến là gì?

- Biến là một vùng bộ nhớ dùng để lưu trữ dữ liệu.
- Giá trị của biến có thể được thay đổi (ghi đè) trong quá trình chạy chương trình.
- Cú pháp cơ bản: 
	- `[var] [Tên biến] [Kiểu dữ liệu] = [Giá trị]`.
	- `[Tên biến] := [Giá trị]`

## Kiểu dữ liệu trong Golang

- **Type** là kiểu dữ liệu của biến. Golang hỗ trợ nhiều kiểu dữ liệu cơ bản:
    - **Số nguyên**: `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`.
    - **Số thực**: `float32`, `float64`.
    - **Chuỗi**: `string` (lưu trữ chuỗi ký tự Unicode).
    - **Logic**: `bool` (giá trị `true` hoặc `false`).
    - **Đặc biệt**:
        - `byte`: Bí danh của `uint8`, dùng để biểu diễn dữ liệu thô (1 byte).
        - `rune`: Bí danh của `int32`, dùng để lưu ký tự Unicode (ví dụ: chữ có dấu, emoji 😎).
- **Lưu ý**:
    - Chọn kiểu dữ liệu phù hợp giúp tối ưu hiệu suất chương trình.
    - Kích thước của `int` và `uint` phụ thuộc vào hệ thống (32-bit hoặc 64-bit).

---
## Quy tắc đặt tên biến

- Nên đặt tên biến theo kiểu `camelCase` (ví dụ: `userName`, `maxRetries`).
- Tên biến nên ngắn gọn, rõ nghĩa, và phản ánh đúng mục đích sử dụng.
- Có thể khai báo nhiều biến cùng một lần.

---
## Cách khai báo biến

### 1. Sử dụng từ khóa `var`

- Cú pháp:
```go
var variableName type = value
```

- Đặc điểm:
    - Có thể khai báo mà không cần gán giá trị ban đầu hoặc chỉ định kiểu dữ liệu.
    - Biến khai báo bằng `var` sẽ nhận **giá trị mặc định (zero value)** nếu không gán giá trị:
        - `int`: 0
        - `float32`, `float64`: 0.0
        - `string`: "" (chuỗi rỗng)
        - `bool`: false
    - Có thể dùng trong hoặc ngoài hàm (toàn cục hoặc cục bộ).

- Ví dụ:
```go
var age int // Zero value: 0
var name string // Zero value: ""
var (
    score int
    subject string
)
```

### 2. Sử dụng dấu `:=`

- Cú pháp:
```go
variableName := value
```

- Đặc điểm:
    - Kiểu dữ liệu được trình biên dịch suy ra từ giá trị (type inference).
    - Phải gán giá trị ngay khi khai báo.
    - Chỉ sử dụng trong hàm.

- Ví dụ:
```go
func main() {
    name := "Golang" // string
    score, rank := 95, 1 // int, int
    fmt.Println(name, score, rank) // In: Golang 95 1
}
```

---
### So sánh `var` và `:=`

| `var`                                | `:=`                                   |
| ------------------------------------ | -------------------------------------- |
| Dùng trong và ngoài hàm              | Chỉ dùng trong hàm                     |
| Có thể khai báo và gán giá trị riêng | Phải khai báo và gán giá trị cùng lúc  |
| Có thể khai báo kiểu dữ liệu         | Không khai báo kiểu dữ liệu, tự suy ra |

---
## Hằng số (`const`)

- Hằng số là giá trị không thể thay đổi sau khi khai báo.

- Cú pháp:
```go
const constantName type = value
```

- Đặc điểm:
    - Chỉ hỗ trợ các kiểu dữ liệu cơ bản (`int`, `float`, `string`, `bool`).
    - Có thể khai báo ở phạm vi toàn cục hoặc trong hàm.
    - Không thể sử dụng `:=` để khai báo hằng số.

- Khi nào nên dùng `const`:
    - Dùng cho các giá trị cố định, không thay đổi trong suốt chương trình (ví dụ: số Pi, tên ứng dụng, số lần thử tối đa).

- Ví dụ:
```go
const Pi float64 = 3.14159
const AppName = "MyApp"
```

## Lưu ý quan trọng

- **Phạm vi biến**:
    - **Toàn cục**: Biến khai báo bằng `var` ngoài hàm, có thể truy cập trong toàn bộ package.
    - **Cục bộ**: Biến khai báo trong hàm (`var` hoặc `:=`), chỉ tồn tại trong phạm vi hàm đó.
    
	- Ví dụ:    
        ```go
        var GlobalVar = "I'm global"
        
        func main() {
            localVar := "I'm local"
            fmt.Println(GlobalVar) // Có thể truy cập
            fmt.Println(localVar)  // Chỉ truy cập trong main
        }
        ```
        
- **Biến không sử dụng**: Golang yêu cầu mọi biến khai báo phải được sử dụng, nếu không sẽ báo lỗi biên dịch. Có thể dùng `_` để bỏ qua biến không cần thiết.

    - Ví dụ:
        ```go
        func main() {
            var x int = 10
            _ = x // Tránh lỗi không sử dụng biến
        }
        ```
        
- **Type inference**: Khi dùng `:=`, Golang tự suy ra kiểu dữ liệu dựa trên giá trị (ví dụ: `x := 1.0` là `float64`, `x := 42` là `int`).

## Ví dụ tổng hợp
```go
package main

import "fmt"

const MaxRetries = 3

var GlobalVar = "I'm global"

func main() {
    // Sử dụng var
    var name string = "Golang"
    var score int
    score = 90
    fmt.Println("Name:", name, "Score:", score)

    // Sử dụng :=
    age := 25
    subject, grade := "Programming", 8
    fmt.Println("Subject:", subject, "Grade:", grade, "Age:", age)

    // Hằng số
    fmt.Println("Max Retries:", MaxRetries)

    // Biến toàn cục
    fmt.Println("Global:", GlobalVar)

    // Zero value
    var x int
    var s string
    var b bool
    fmt.Println("Zero values:", x, s, b)
}
```

**Kết quả**:

```
Name: Golang Score: 90
Subject: Programming Grade: 8 Age: 25
Max Retries: 3
Global: I'm global
Zero values: 0 "" false
```

---
### **Gán nhiều biến cùng một lúc trong Golang**


Ví dụ:
```go
func main() {
	var a = 10
	var b = 15

	a, b = b, a

	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
```



