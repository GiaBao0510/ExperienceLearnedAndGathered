Bài 01: Xây dựng 1 hàm đệ quy để tính tổng dãy số N
VD: Người dùng nhập vào 10 thì sẽ tính tổng toàn bộ số từ 1 > 10

```go
func TongDaySoN(num int) int {
    if num == 0 {
        return 0
    }

    return num + TongDaySoN(num-1)
}
```

Bài 02: Xây dựng 1 hàm để hiển thị dãy số Fibonacci
- Cho phép người dùng nhập vào số lượng hiển thị dãy số Fibonacci
- Hiển thị dãy số Fibonacci
- Tỉnh tổng dãy số Fibonacci đó

```go
func HienThiDaySo_Fibonacci(num int) {

    if num <= 0 {
        fmt.Println("Vui lòng nhập số nguyên dương.")

    } else if num == 1 {
        fmt.Println("Day so Fibonacci: 0")

    } else if num == 2 {
        fmt.Println("Day so Fibonacci: 0 1")

    } else {
        fmt.Print("Day so Fibonacci: ")

        sum := 0
        total := 0
        term1, term2 := 0, 1

        fmt.Printf("%d ", term1)
        fmt.Printf("%d ", term2)

        total += (term1 + term2)

        for i := 3; i <= num; i++ {

            sum = term1 + term2
            total += sum

            fmt.Printf("%d ", sum)
            term1, term2 = term2, sum
        }

        fmt.Printf("\nTong Day so Fibonacci: %d", total)
    }
}
```

Bài 03: Xây dựng tập hợp chức năng
- Chuyển toàn bộ bài tập đã làm thành function
- Xây dựng menu để có thể gọi hàm từ menu
```go

```