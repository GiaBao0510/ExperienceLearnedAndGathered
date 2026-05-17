# Return, Break, và Continue Trong Go

Ba từ khóa `return`, `break`, và `continue` dùng để **điều khiển luồng thực thi** trong hàm, vòng lặp (`for`), hoặc cấu trúc điều kiện (`switch`, `if`). Chúng giúp code linh hoạt hơn, nhưng cần dùng đúng để tránh lỗi

### 1. `return` – Thoát khỏi hàm hiện tại

- **Mô tả**: Dừng hàm ngay lập tức và trả về giá trị (nếu có). Nếu dùng trong vòng lặp bên trong hàm, nó sẽ dừng vòng lặp và thoát hàm, **không phải thoát chương trình toàn bộ** (trừ khi ở `main()`).
- **Ứng dụng**: Thường dùng để trả kết quả sớm khi tìm thấy điều kiện.

**Ví dụ**: Tìm số chẵn đầu tiên trong mảng và trả về.
```go
package main
import "fmt"

func timSoChan(danhSach []int) int {
    for _, num := range danhSach {
        if num%2 == 0 {
            return num  // Dừng vòng lặp và thoát hàm
        }
    }
    return -1  // Không tìm thấy
}

func main() {
    nums := []int{1, 3, 5, 2, 4}
    fmt.Println("Số chẵn đầu tiên:", timSoChan(nums))  // Kết quả: 2
}
```

### 2. `break` – Thoát khỏi vòng lặp hoặc switch gần nhất

- **Mô tả**: Dừng vòng lặp (`for`) hoặc `switch` ngay lập tức và chuyển đến lệnh sau đó. Trong `switch` của Go, không cần `break` vì không tự động fallthrough (chuyển sang case tiếp), nhưng có thể dùng để thoát sớm nếu cần.
- **Sai sót phổ biến**: Không dùng `break` trong `if-else` đơn lẻ (vì `if` không phải vòng lặp). Nếu cần thay thế `switch` bằng `if-else`, bạn không cần `break` vì `if` tự dừng.

**Ví dụ 1 (với vòng lặp)**: Thoát khi tìm thấy số 5.
```go
for i := 1; i <= 10; i++ {
    if i == 5 {
        break  // Thoát vòng lặp
    }
    fmt.Println(i)  // In: 1 2 3 4
}
```

**Ví dụ 2 (với switch)**: Xử lý điểm số (không cần `break` ở mỗi case).
```go
diem := 8
switch diem {
case 10, 9:
    fmt.Println("Xuất sắc")
case 8, 7:
    fmt.Println("Giỏi")  // Tự dừng, không cần break
default:
    fmt.Println("Khác")
}
```

### 3. `continue` – Bỏ qua phần còn lại của lần lặp hiện tại

- **Mô tả**: Chỉ dùng trong vòng lặp (`for`). Bỏ qua các lệnh phía dưới và chuyển sang lần lặp tiếp theo. Vòng lặp vẫn tiếp tục bình thường.
- **Ứng dụng**: Lọc dữ liệu, bỏ qua trường hợp không mong muốn.

**Ví dụ**: In số lẻ từ 1 đến 10 (bỏ qua số chẵn).
```go
for i := 1; i <= 10; i++ {
    if i%2 == 0 {
        continue  // Bỏ qua số chẵn, sang i tiếp theo
    }
    fmt.Println("Số lẻ:", i)  // In: 1 3 5 7 9
}
```

### Lưu Ý Quan Trọng 

- `return`: Ảnh hưởng toàn hàm, không chỉ vòng lặp.
- `break` và `continue`: Chỉ trong vòng lặp gần nhất. Với vòng lặp lồng nhau, dùng **nhãn (label)** để chỉ định (nâng cao).
- Trong `switch`: Go khác C/Java – không fallthrough mặc định, dùng `fallthrough` nếu muốn tiếp tục case sau.
- Tránh lạm dụng: Ưu tiên code rõ ràng, dễ đọc.
- Lỗi thường gặp: Quên điều kiện dừng → lặp vô hạn.
