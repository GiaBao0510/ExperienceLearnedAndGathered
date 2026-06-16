> "Abstraction means working with something we know how to use without knowing how it works internally."

Trong Golang, **tính trừu tượng** (Abstraction) không được thực hiện thông qua lớp trừu tượng (abstract class) hay kế thừa (inheritance) như các ngôn ngữ OOP truyền thống. Thay vào đó, Go sử dụng **Interface** để định nghĩa hợp đồng hành vi (behavioral contract), giúp ẩn đi chi tiết triển khai bên trong và tăng tính linh hoạt cho mã nguồn.

## Cơ chế hoạt động của Tính trừu tượng trong Go

- **Định nghĩa hành vi (Interface):** Là tập hợp các chữ ký phương thức (method signatures) không có phần thân — chỉ khai báo _hành vi_, không khai báo _cách thực hiện_.
- **Triển khai ngầm định (Implicit Implementation):** Bất kỳ `struct` nào định nghĩa đầy đủ các phương thức có trong `interface` đều được coi là tự động thực thi interface đó — không cần từ khóa `implements` như trong Java hay C#.
- **Giao tiếp qua interface:** Code nghiệp vụ chỉ phụ thuộc vào interface, không phụ thuộc vào bất kỳ kiểu cụ thể nào.

> **Lưu ý:** Interface trong Go là kiểu dữ liệu, không phải khai báo lớp. Một `struct` có thể thực thi nhiều interface cùng lúc.

## Ví dụ thực tế

Giả sử bạn đang xây dựng một hệ thống thanh toán. Bạn muốn ẩn đi cách thức thanh toán cụ thể (bằng thẻ hay qua ví điện tử) và chỉ quan tâm đến hành động "thanh toán".

```go
package main

import "fmt"

// 1. Định nghĩa Interface (Phần trừu tượng - Chỉ quan tâm đến hành vi)
type PaymentGateway interface {
	Pay(amount float64) string
}

// 2. Triển khai 1: Thanh toán bằng Paypal (Chi tiết triển khai)
type Paypal struct {
	Email string
}

func (p *Paypal) Pay(amount float64) string {
	return fmt.Sprintf("Đã thanh toán $%.2f qua tài khoản Paypal %s", amount, p.Email)
}

// 3. Triển khai 2: Thanh toán bằng Credit Card (Chi tiết triển khai)
type CreditCard struct {
	CardNumber string
}

func (c *CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("Đã thanh toán $%.2f bằng thẻ tín dụng có số đuôi %s", amount, c.CardNumber[12:])
}

// 4. Hàm xử lý nghiệp vụ (Chỉ phụ thuộc vào Interface, không phụ thuộc vào chi tiết)
func ProcessPayment(pg PaymentGateway, amount float64) {
	result := pg.Pay(amount)
	fmt.Println(result)
}

func main() {
	paypalMethod := &Paypal{Email: "user@example.com"}
	cardMethod := &CreditCard{CardNumber: "1234567890123456"}

	ProcessPayment(paypalMethod, 150.0) // Kết quả: Đã thanh toán $150.00 qua tài khoản Paypal user@example.com
	ProcessPayment(cardMethod, 200.0)   // Kết quả: Đã thanh toán $200.00 bằng thẻ tín dụng có số đuôi 3456
}
```

## Tại sao tính trừu tượng bằng Interface lại quan trọng trong Go?

- **Tách biệt mã nguồn (Decoupling):** Code nghiệp vụ (như hàm `ProcessPayment`) chỉ giao tiếp qua `interface`. Bạn có thể thay đổi hoặc thêm mới một phương thức thanh toán mà không cần sửa bất kỳ dòng code nghiệp vụ nào.
- **Dễ dàng kiểm thử (Unit Testing):** Bạn có thể tạo `mock struct` giả lập hành vi để kiểm thử mà không cần kết nối đến hệ thống thật (database, cổng thanh toán,...).
- **Mở rộng dễ dàng (Open/Closed Principle):** Thêm phương thức thanh toán mới chỉ cần tạo một `struct` mới triển khai `PaymentGateway` — không cần sửa code hiện có.

## Interface lồng nhau và Interface rỗng

### Interface lồng nhau (Interface Embedding)

Một interface có thể nhúng interface khác để tạo hợp đồng hành vi phức hợp:

```go
type Reader interface {
	Read() string
}

type Writer interface {
	Write(data string)
}

// ReadWriter nhúng cả Reader và Writer
type ReadWriter interface {
	Reader
	Writer
}
```

Bất kỳ `struct` nào triển khai đầy đủ cả `Read()` lẫn `Write()` đều thỏa mãn interface `ReadWriter`.

### Interface rỗng (Empty Interface)

`interface{}` (hoặc `any` từ Go 1.18 trở đi) không yêu cầu bất kỳ phương thức nào, do đó mọi kiểu dữ liệu đều thỏa mãn nó:

```go
func PrintAnything(v any) {
	fmt.Println(v)
}

PrintAnything(42)        // int
PrintAnything("hello")   // string
PrintAnything(true)      // bool
```

> **Lưu ý:** Nên hạn chế dùng `any` vì mất đi kiểm tra kiểu tĩnh tại compile-time. Chỉ dùng khi thực sự cần xử lý nhiều kiểu không xác định trước.

## Kiểm tra kiểu tại runtime (Type Assertion & Type Switch)

Khi làm việc với interface, đôi khi bạn cần lấy lại kiểu cụ thể bên trong.

### Type Assertion

```go
var pg PaymentGateway = &Paypal{Email: "user@example.com"}

// Dạng an toàn (không panic nếu sai kiểu)
if p, ok := pg.(*Paypal); ok {
	fmt.Println("Email:", p.Email)
}
```

### Type Switch

```go
func Describe(pg PaymentGateway) {
	switch v := pg.(type) {
	case *Paypal:
		fmt.Println("Paypal, email:", v.Email)
	case *CreditCard:
		fmt.Println("Credit Card, số:", v.CardNumber)
	default:
		fmt.Println("Phương thức thanh toán không xác định")
	}
}
```

## Những lưu ý quan trọng

- **Pointer receiver vs Value receiver:** Nếu phương thức của `struct` được định nghĩa với **pointer receiver** (`func (p *Paypal) Pay(...)`), thì chỉ `*Paypal` (con trỏ) mới thỏa mãn interface — không phải `Paypal` (giá trị). Đây là lỗi phổ biến khi mới học Go.

```go
// ❌ Sai — Paypal (value) không thỏa mãn interface khi Pay dùng pointer receiver
var pg PaymentGateway = Paypal{Email: "user@example.com"} // compile error

// ✅ Đúng — dùng con trỏ
var pg PaymentGateway = &Paypal{Email: "user@example.com"}
```

- **Interface nên nhỏ gọn:** Theo triết lý Go, interface tốt thường chỉ có 1–3 phương thức. Interface `io.Reader` của thư viện chuẩn chỉ có một phương thức `Read()` nhưng được dùng ở khắp nơi trong hệ sinh thái Go.
- **Không nên định nghĩa interface quá sớm:** Chỉ tạo interface khi bạn thực sự có từ 2 triển khai trở lên, hoặc khi cần mock để kiểm thử.