# Nhập Dữ liệu từ Bàn phím trong Golang

Nhập liệu từ bàn phím là cách chương trình tương tác với người dùng. Trong Golang, hai gói phổ biến để nhập liệu là `fmt` (đơn giản, phù hợp cho dữ liệu cơ bản) và `bufio` (linh hoạt, xử lý chuỗi phức tạp). Tài liệu này giới thiệu hai cách nhập liệu chính: `fmt.Scan` và `bufio.NewScanner`.

---
## 1. Sử dụng `fmt.Scan`

- **Mô tả**: `fmt.Scan` đọc dữ liệu từ luồng đầu vào chuẩn (`os.Stdin`) và gán vào các biến. Dữ liệu được phân tách bởi khoảng trắng, và ==hàm dừng khi gặp ký tự xuống dòng (`\n`) hoặc khoảng trắng==.
- **Hạn chế**: Chỉ đọc được một token (phần dữ liệu trước khoảng trắng hoặc `\n`), ==không phù hợp với chuỗi có khoảng trắng==.

- **Cú pháp**:
```go
var variable type
fmt.Scan(&variable)
```

- **Ví dụ**:
```go
package main

import "fmt"

func main() {
    var fullName string
    fmt.Print("Nhập họ tên của bạn: ")
    _, err := fmt.Scan(&fullName)
    if err != nil {
        fmt.Println("Lỗi nhập liệu:", err)
        returnP
    }
    fmt.Println("Họ tên của bạn là:", fullName)
}
```

- **Kết quả**:
```shell
Nhập họ tên của bạn: KimHo
Họ tên của bạn là: KimHo

Nhập họ tên của bạn: Pham Gia Bao
Họ tên của bạn là: Pham
```

- **Lưu ý**:
    - `fmt.Scan` chỉ lấy token đầu tiên (dừng khi gặp khoảng trắng hoặc `\n`).
    - Kiểm tra lỗi bằng cách kiểm tra giá trị trả về (số lượng giá trị được đọc thành công).

---
## 2. Sử dụng `bufio.NewScanner`

- **Mô tả**: `bufio.NewScanner` đọc dữ liệu theo dòng từ `os.Stdin`, phù hợp để nhập chuỗi dài hoặc chuỗi có khoảng trắng.
- **Ưu điểm**: Có thể đọc toàn bộ dòng, bao gồm khoảng trắng, và xử lý lỗi tốt hơn.

- **Cú pháp**:
```go
scanner := bufio.NewScanner(os.Stdin)
if scanner.Scan() {
    variable = scanner.Text()
}
```

- **Ví dụ**:
```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    var fullName string
    fmt.Print("Nhập họ tên của bạn: ")
    scanner := bufio.NewScanner(os.Stdin)
    if scanner.Scan() {
        fullName = scanner.Text()
    }
    if err := scanner.Err(); err != nil {
        fmt.Println("Lỗi nhập liệu:", err)
        return
    }
    fmt.Println("Họ tên của bạn là:", fullName)
}
```

- **Kết quả**:
```shell
Nhập họ tên của bạn: Pham Gia Bao
Họ tên của bạn là: Pham Gia Bao
```

*Lưu ý:*
- `os`: là hệ điều hành
- `stdin`: là cho phép nhập vô

---
## 3. Nhập các kiểu dữ liệu khác

- **Nhập số nguyên hoặc số thực**:
    - Với `fmt.Scan`:
        ```go
        var number int
        fmt.Print("Nhập số nguyên: ")
        _, err := fmt.Scan(&number)
        if err != nil {
            fmt.Println("Lỗi nhập liệu:", err)
            return
        }
        fmt.Println("Số bạn nhập:", number)
        ```
        
    - Với `bufio.NewScanner` và `strconv`:
        ```go
        scanner := bufio.NewScanner(os.Stdin)
        fmt.Print("Nhập số nguyên: ")
        scanner.Scan()
        input := scanner.Text()
        number, err := strconv.Atoi(input)
        if err != nil {
            fmt.Println("Lỗi: Craven): Vui lòng nhập số nguyên")
            return
        }
        fmt.Println("Số bạn nhập:", number)
        ```
        
- **Lưu ý**: Sử dụng `strconv` để chuyển đổi chuỗi thành số (`Atoi` cho số nguyên, `ParseFloat` cho số thực).

---
## 4. So sánh `fmt.Scan` và `bufio.NewScanner`

| `fmt.Scan`                 | `bufio.NewScanner`                     |
| -------------------------- | -------------------------------------- |
| Đơn giản, dễ dùng          | Linh hoạt, đọc cả dòng                 |
| Chỉ đọc token đầu tiên     | Đọc toàn bộ dòng, bao gồm khoảng trắng |
| Phù hợp với số, chuỗi ngắn | Phù hợp với chuỗi dài, phức tạp        |
| Xử lý lỗi đơn giản         | Cần kiểm tra lỗi bằng `scanner.Err()`  |

---
## 5. Lưu ý quan trọng

- **Xử lý lỗi**: Luôn kiểm tra lỗi khi nhập liệu để tránh crash chương trình.
- **Hiệu suất**: `bufio.NewScanner` hiệu quả hơn khi đọc dữ liệu lớn hoặc phức tạp.
- **Trường hợp sử dụng**:
    - Dùng `fmt.Scan` cho dữ liệu đơn giản (số, chuỗi ngắn).
    - Dùng `bufio.NewScanner` cho chuỗi dài hoặc có khoảng trắng.
- **Ký tự xuống dòng**: `bufio.NewScanner` tự động loại bỏ `\n`, trong khi `fmt.Scan` dừng tại `\n`.

---
## 6. Ví dụ tổng hợp

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    // Nhập chuỗi bằng bufio.NewScanner
    fmt.Print("Nhập họ tên: ")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    fullName := scanner.Text()
    fmt.Println("Họ tên:", fullName)

    // Nhập số bằng fmt.Scan
    var age int
    fmt.Print("Nhập tuổi: ")
    _, err := fmt.Scan(&age)
    if err != nil {
        fmt.Println("Lỗi nhập liệu:", err)
        return
    }
    fmt.Println("Tuổi:", age)
}
```

- **Kết quả**:

```shell
Nhập họ tên: Pham Gia Bao
Họ tên: Pham Gia Bao
Nhập tuổi: 25
Tuổi: 25
```