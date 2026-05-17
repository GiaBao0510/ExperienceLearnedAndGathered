# Vòng Lặp Trong Go

Trong Go, **chỉ có một loại vòng lặp duy nhất là `for`**. Không có `while` hay `do-while` như các ngôn ngữ khác. Tuy nhiên, bạn có thể dùng `for` linh hoạt để thay thế hoàn toàn chúng.

### 1. Vòng lặp `for` cổ điển (giống C/Java)

![](https://www.programiz.com/sites/tutorial2program/files/golang-for-loop-flowcontrol.png)

Dùng để lặp một số lần xác định.

**Cú pháp**:
```go
for khởi_tạo; điều_kiện; bước_nhảy {
    // Khối lệnh thực thi
}
```

**Giải thích**:
- `khởi_tạo`: Thường là khai báo biến đếm (ví dụ: `i := 0`).
- `điều_kiện`: Kiểm tra trước mỗi lần lặp (nếu sai → dừng).
- `bước_nhảy`: Thường là tăng/giảm biến đếm (ví dụ: `i++`).

**Ví dụ**:
```go
package main
import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        fmt.Println("Lần thứ", i)
    }
}
```

**Kết quả**:
```
Lần thứ 0
Lần thứ 1
Lần thứ 2
Lần thứ 3
Lần thứ 4
```

---
### 2. Vòng lặp `for` kiểu `while` (lặp khi điều kiện đúng)

Bỏ phần khởi tạo và bước nhảy, chỉ giữ điều kiện.

**Cú pháp**:
```go
for điều_kiện {
    // Khối lệnh
}
```

**Ví dụ**: In số từ 1 đến 10, nhưng dùng kiểu while:
```go
i := 1
for i <= 10 {
    fmt.Println(i)
    i++  // Phải tự tăng, nếu không sẽ lặp vô hạn!
}
```

---
### 3. Vòng lặp vô hạn (tương đương `while(true)`)

**Cú pháp**:
```go
for {
    // Khối lệnh lặp mãi mãi
    // Dùng break để thoát
}
```

**Ví dụ**: Chương trình hỏi tên đến khi người dùng nhập "quit"
```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    for {
        fmt.Print("Nhập tên của bạn (gõ 'quit' để thoát): ")
        var name string
        fmt.Scanln(&name)

        if strings.TrimSpace(name) == "quit" {
            break  // Thoát vòng lặp
        }
        fmt.Println("Xin chào,", name)
    }
    fmt.Println("Tạm biệt!")
}
```

### 4. Dùng `for` với `range` để duyệt mảng/slice/string/map

Rất phổ biến trong Go!

**Ví dụ 1**: Duyệt slice
```go
numbers := []int{10, 20, 30}

for index, value := range numbers {
    fmt.Printf("Phần tử thứ %d = %d\n", index, value)
}
```

**Ví dụ 2**: Chỉ lấy giá trị (bỏ index)
```go
fruits := []string{"táo", "chuối", "cam"}

for _, fruit := range fruits {  // Dùng _ để bỏ qua index
    fmt.Println("Trái cây:", fruit)
}
```

### Lưu Ý Quan Trọng

- Go **không có** `while` hay `do-while`. Tất cả đều dùng `for`.
- Luôn nhớ tăng/giảm biến trong vòng lặp kiểu while để tránh **lặp vô hạn**.
- Dùng `break` để thoát vòng lặp sớm.
- Dùng `continue` để bỏ qua lần lặp hiện tại và sang lần tiếp theo.
