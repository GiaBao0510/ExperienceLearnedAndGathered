# Hàm (Function) Trong Go

Hàm là một khối mã độc lập, được viết để thực hiện một nhiệm vụ cụ thể. Thay vì lặp lại cùng một đoạn code nhiều lần, bạn chỉ cần định nghĩa hàm một lần và gọi lại khi cần. Điều này giúp code sạch hơn, dễ bảo trì và tái sử dụng.

### Cú Pháp Cơ Bản

#### 1. Hàm không trả về giá trị
```go
func tênHàm(tham_số1 kiểu, tham_số2 kiểu) {
    // Khối lệnh
}
```

#### 2. Hàm có trả về giá trị
```go
func tênHàm(tham_số1 kiểu, tham_số2 kiểu) kiểu_trả_về {
    // Khối lệnh
    return giá_trị
}
```

#### 3. Gọi hàm
```go
tênHàm(giá_trị1, giá_trị2)
```

### Lưu Ý Quan Trọng
- Số lượng và kiểu của tham số khi **gọi hàm** phải **đúng** với khi **khai báo**.
- Go **không hỗ trợ giá trị mặc định** cho tham số (khác với nhiều ngôn ngữ khác).
- Không được viết: `func sum(a int, b int = 10)` → **lỗi biên dịch**.

### Ví Dụ Thực Tế

```go
package main
import "fmt"

// 1. Hàm tính tổng - trả về giá trị
func tong(a int, b int) int {
    return a + b
}

// 2. Cách viết ngắn gọn hơn (named return)
func tong2(a int, b int) (ketQua int) {
    ketQua = a + b
    return // Có thể bỏ giá trị vì đã đặt tên
}

// 3. Hàm in thông báo - không trả về
func inChao(ten string) {
    fmt.Println("Xin chào,", ten)
}

// 4. Hàm với nhiều giá trị trả về (rất phổ biến trong Go)
func chiaDu(a, b int) (thuong int, du int) {
    thuong = a / b
    du = a % b
    return
}

func main() {
    // Gọi hàm tính tổng
    fmt.Printf("5 + 10 = %d\n", tong(5, 10))     // Kết quả: 15
    fmt.Printf("7 + 8 = %d\n", tong2(7, 8))      // Kết quả: 15

    // Gọi hàm in
    inChao("GoLang")                            // Kết quả: Xin chào, GoLang

    // Nhận nhiều giá trị trả về
    t, d := chiaDu(17, 5)
    fmt.Printf("17 chia 5: thương = %d, dư = %d\n", t, d)
    // Kết quả: thương = 3, dư = 2
}
```

### Cách Xử Lý Khi Muốn Có Giá Trị Mặc Định

Vì Go không cho phép giá trị mặc định trực tiếp, bạn có thể xử lý bên trong hàm:

```go
func chaoTen(ten string) {
    if ten == "" {
        ten = "Bạn"  // Giá trị mặc định nếu không truyền
    }
    fmt.Println("Xin chào,", ten)
}

func main() {
    chaoTen("Nam")    // Xin chào, Nam
    chaoTen("")       // Xin chào, Bạn
}
```

---
Hàm (function) là thành phần cơ bản của chương trình. Các hàm trong ngôn ngữ Go có thể có tên hoặc ẩn danh (anonymous function):

```go
// hàm được đặt tên
func Add(a, b int) int {
    return a+b
}

// hàm ẩn danh
var Add = func(a, b int) int {
    return a+b
}
```

Một hàm trong ngôn ngữ Go có thể có nhiều tham số và nhiều giá trị trả về. Cả tham số và giá trị trả về trao đổi dữ liệu với hàm theo cách truyền vào giá trị (pass by value). Về mặt cú pháp, hàm cũng hỗ trợ số lượng tham số thay đổi, biến số lượng tham số phải là tham số cuối cùng và biến này phải là kiểu slice.
```go
// Nhiều tham số và nhiều giá trị trả về
func Swap(a, b int) (int, int) {
    return b, a
}

// Biến số lượng tham số 'more'
// Tương ứng với kiểu [] int, là một slice
func Sum(a int, more ...int) int {
    for _, v := range more {
        a += v
    }
    return a
}
```

Khi đối số có thể thay đổi là một kiểu interface null, việc người gọi có phân giải (unpack) đối số đó hay không sẽ dẫn đến những kết quả khác nhau:

```go
func main() {
    var a = []interface{}{123, "abc"}

    // tương đương với lời gọi trực tiếp `Print(123, "abc")`
    Print(a...) // 123 abc

    // tương đương với lời gọi `Print([]interface{}{123, "abc"})`
    Print(a)    // [123 abc]
}

func Print(a ...interface{}) {
    fmt.Println(a...)
}
```

Cả tham số truyền vào và các giá trị trả về đều có thể được đặt tên:

```go
func Find(m map[int]int, key int) (value int, ok bool) {
    value, ok = m[key]
    return
}
```