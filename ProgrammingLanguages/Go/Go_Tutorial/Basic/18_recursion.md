# Đệ Quy (Recursion) Trong Go

**Đệ quy** là kỹ thuật một hàm **gọi lại chính nó** để giải quyết bài toán bằng cách chia nhỏ thành các bài toán con giống nhau, cho đến khi gặp **điều kiện dừng** (base case).

Nếu không có điều kiện dừng rõ ràng, chương trình sẽ gọi hàm mãi mãi → **lỗi stack overflow** (tràn bộ nhớ).

### Cú Pháp Cơ Bản

```go
func tênHàm(tham_số) kiểu_trả_về {
    // Điều kiện dừng (rất quan trọng!)
    if điều_kiện_dừng {
        return giá_trị_cơ_sở
    }

    // Gọi lại chính hàm với bài toán con nhỏ hơn
    return tênHàm(tham_số_nhỏ_hơn)
}
```

### Ví Dụ 1: Đếm ngược từ n về 1

```go
package main
import "fmt"

func countDown(n int) {
    if n <= 0 {              // Điều kiện dừng
        fmt.Println("Kết thúc!")
        return
    }

    fmt.Println(n)
    countDown(n - 1)         // Gọi lại với n nhỏ hơn
}

func main() {
    countDown(5)
}
```

**Kết quả**:
```
5
4
3
2
1
Kết thúc!
```

### Ví Dụ 2: Tính giai thừa (factorial) – rất kinh điển

```go
package main
import "fmt"

func factorial(n int) int {
    if n == 0 || n == 1 {    // Điều kiện dừng
        return 1
    }

    return n * factorial(n - 1)  // Đệ quy
}

func main() {
    fmt.Println("5! =", factorial(5))  // 120
    fmt.Println("0! =", factorial(0))  // 1
}
```

**Kết quả**:
```
5! = 120
0! = 1
```

### Lưu Ý Quan Trọng

- **Luôn phải có điều kiện dừng** rõ ràng.
- Mỗi lần gọi đệ quy, bài toán phải **nhỏ hơn** (tiến gần đến điều kiện dừng).
- Đệ quy dễ hiểu nhưng có thể **chậm và tốn bộ nhớ** nếu n lớn (vì mỗi lần gọi tạo một khung stack mới).
- Với bài toán lớn, nên cân nhắc dùng vòng lặp (iterative) thay vì đệ quy.

### Khi Nào Nên Dùng Đệ Quy?

- Bài toán có cấu trúc tự nhiên chia nhỏ (giai thừa, Fibonacci, duyệt cây, backtracking...).
- Code ngắn gọn, dễ đọc hơn so với vòng lặp.
