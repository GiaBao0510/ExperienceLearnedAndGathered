# Câu Lệnh `switch...case` Trong Go

`switch...case` là một cấu trúc điều khiển giúp kiểm tra giá trị của một biểu thức và thực hiện khối lệnh tương ứng. Nó thay thế cho chuỗi `if...else if...else` dài dòng, giúp code ngắn gọn, rõ ràng và dễ đọc hơn – đặc biệt khi có nhiều điều kiện kiểm tra cùng một biến.

### Cú Pháp Cơ Bản

```go
switch biểu_thức {  // Có thể bỏ biểu_thức nếu dùng switch không điều kiện
case giá_trị_1:
    // Khối lệnh thực hiện khi biểu_thức == giá_trị_1
case giá_trị_2:
    // Khối lệnh thực hiện khi biểu_thức == giá_trị_2
case giá_trị_3, giá_trị_4:  // Nhiều giá trị trong một case
    // Thực hiện khi biểu_thức bằng giá_trị_3 hoặc giá_trị_4
default:  // Tùy chọn – thực hiện khi không khớp case nào
    // Khối lệnh mặc định
}
```

### Ví Dụ Minh Họa

#### 1. Switch thông thường (có biểu thức)

```go
package main
import "fmt"

func main() {
    diem := 8

    switch diem {
    case 10, 9:
        fmt.Println("Xuất sắc")
    case 8, 7:
        fmt.Println("Giỏi")
    case 6, 5:
        fmt.Println("Khá")
    case 4, 3, 2, 1, 0:
        fmt.Println("Yếu")
    default:
        fmt.Println("Điểm không hợp lệ")
    }
}
```

**Kết quả**: `Giỏi`

#### 2. Switch không có biểu thức (tương đương if...else if)

```go
tuoi := 20

switch {
case tuoi < 13:
    fmt.Println("Thiếu nhi")
case tuoi < 18:
    fmt.Println("Thiếu niên")
case tuoi < 60:
    fmt.Println("Trung niên")
default:
    fmt.Println("Người cao tuổi")
}
```

**Kết quả**: `Trung niên`

#### 3. Sử dụng `fallthrough` (chuyển tiếp sang case tiếp theo)

Mặc định Go **không** tự động chạy sang case tiếp theo (khác với C/C++). Nếu muốn, dùng từ khóa `fallthrough`:

```go
diem := 10

switch diem {
case 10:
    fmt.Println("Hoàn hảo")
    fallthrough  // Chạy tiếp sang case 9
case 9:
    fmt.Println("Gần hoàn hảo")
}
```

**Kết quả**:
```
Hoàn hảo
Gần hoàn hảo
```

### Lưu Ý Quan Trọng

- Mỗi `case` không cần `break` (Go tự động break).
- Có thể gộp nhiều giá trị trong một `case` bằng dấu phẩy.
- `default` không bắt buộc, có thể đặt ở bất kỳ vị trí nào (nhưng thường để cuối).
- Có thể dùng biểu thức phức tạp hơn trong `case` nếu dùng switch kiểu "type switch" hoặc với điều kiện (xem tài liệu nâng cao).

### Khi Nào Nên Dùng `switch`?

- Khi kiểm tra **cùng một biến** với **nhiều giá trị cụ thể**.
- Khi có hơn 3–4 nhánh `if...else` liên tiếp.
- Muốn code sạch, dễ bảo trì.