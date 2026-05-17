# Break và Continue Trong Go

Hai câu lệnh `break` và `continue` dùng để **điều khiển luồng thực thi** bên trong vòng lặp `for` (Go chỉ có vòng lặp `for`). Chúng giúp bạn thoát sớm hoặc bỏ qua một phần của vòng lặp một cách linh hoạt.

---
### 1. `break` – Thoát khỏi vòng lặp ngay lập tức

Khi gặp `break`, chương trình sẽ **dừng toàn bộ vòng lặp** và chuyển đến lệnh ngay sau vòng lặp.

**Ví dụ**: Tìm số đầu tiên chia hết cho 5 trong dãy từ 1 đến 20
```go
package main
import "fmt"

func main() {
    for i := 1; i <= 20; i++ {
        if i%5 == 0 {
            fmt.Println("Tìm thấy số đầu tiên chia hết cho 5:", i)
            break  // Thoát vòng lặp ngay lập tức
        }
        fmt.Println("Đang kiểm tra:", i)
    }
    fmt.Println("Kết thúc chương trình")
}
```

**Kết quả**:
```
Đang kiểm tra: 1
Đang kiểm tra: 2
Đang kiểm tra: 3
Đang kiểm tra: 4
Tìm thấy số đầu tiên chia hết cho 5: 5
Kết thúc chương trình
```

→ Vòng lặp dừng ngay khi tìm thấy số 5.

---
### 2. `continue` – Bỏ qua phần còn lại của lần lặp hiện tại

Khi gặp `continue`, chương trình sẽ **bỏ qua các lệnh phía dưới** trong lần lặp hiện tại và chuyển sang lần lặp tiếp theo.

**Ví dụ**: In các số lẻ từ 1 đến 10 (bỏ qua số chẵn)
```go
package main
import "fmt"

func main() {
    for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            continue  // Bỏ qua số chẵn, sang số tiếp theo
        }
        fmt.Println("Số lẻ:", i)
    }
}
```

**Kết quả**:
```
Số lẻ: 1
Số lẻ: 3
Số lẻ: 5
Số lẻ: 7
Số lẻ: 9
```

→ Các số chẵn bị bỏ qua, không in ra.

### Lưu Ý Quan Trọng

- `break` và `continue` chỉ ảnh hưởng đến **vòng lặp gần nhất** bao quanh nó.
- Có thể dùng với vòng lặp lồng nhau, kết hợp với **nhãn (label)** để điều khiển vòng lặp bên ngoài (nâng cao).
- Không dùng `break`/`continue` ngoài vòng lặp → lỗi biên dịch.

**Tóm tắt nhanh**:
- `break`    → **Dừng hẳn** vòng lặp.
- `continue` → **Bỏ qua** lần lặp hiện tại, làm lần tiếp theo.