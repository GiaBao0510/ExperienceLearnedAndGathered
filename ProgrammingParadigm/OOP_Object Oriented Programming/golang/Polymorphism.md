# Tính Đa Hình (Polymorphism) trong Go

> Polymorphism describes a pattern in object-oriented programming in which classes have different functionality while sharing a common interface.

---

## Tính Đa Hình Trong Go Hoạt Động Như Thế Nào?

Như ta thường thấy ở các ngôn ngữ lập trình hướng đối tượng (OOP) như Java hay C#, tính đa hình thường được thể hiện thông qua kế thừa (inheritance): các lớp con ghi đè (override) phương thức của lớp cha, nhưng đều có thể được xử lý thống nhất qua kiểu của lớp cha.

Go không có khái niệm kế thừa hay lớp (class). Thay vào đó, Go đạt được tính đa hình thông qua **interface**. Một kiểu (type) được coi là "thỏa mãn" một interface nếu nó implement đủ tất cả các method mà interface đó khai báo — **không cần khai báo tường minh** bằng từ khóa như `implements` (khác với Java/C#). Đây là cơ chế **implicit interface implementation**, còn được gọi là **duck typing**.

> cơ chế "implicit interface implementation" — đây là điểm khác biệt cốt lõi và thường gây nhầm lẫn cho người mới từ Java/C# chuyển sang Go. Việc hiểu rõ điều này giúp người đọc nắm được _tại sao_ code trong ví dụ hoạt động, thay vì chỉ thấy _nó_ hoạt động.

---

## Ví Dụ

Xét ví dụ dưới đây, trong đó `Hillary` và `Trump` là hai struct khác nhau, nhưng đều thỏa mãn interface `SloganSayer` bằng cách implement method `Slogan()`.

```go
// File: polymorphism/polymorphism.go
package polymorphism

import "fmt"

// SloganSayer là interface định nghĩa hành vi "nói khẩu hiệu"
type SloganSayer interface {
    Slogan()
}

// SaySlogan nhận vào bất kỳ kiểu nào thỏa mãn SloganSayer.
// Đây chính là polymorphism: hàm không quan tâm kiểu cụ thể
// là gì, miễn là nó implement interface đúng quy định.
func SaySlogan(sloganSayer SloganSayer) {
    sloganSayer.Slogan()
}

// Hillary thỏa mãn SloganSayer interface
// bằng cách implement method Slogan.
type Hillary struct{}

func (h *Hillary) Slogan() {
    fmt.Println("Stronger together.")
}

// Trump cũng thỏa mãn SloganSayer interface theo cách tương tự.
type Trump struct{}

func (t *Trump) Slogan() {
    fmt.Println("Make America great again.")
}
```

```go
// File: main.go
package main

import "github.com/amy/tech-talk/polymorphism"

func main() {
    hillary := new(polymorphism.Hillary)
    hillary.Slogan()                  // Output: "Stronger together."
    polymorphism.SaySlogan(hillary)   // Output: "Stronger together."

    trump := new(polymorphism.Trump)
    polymorphism.SaySlogan(trump)     // Output: "Make America great again."
}
```

> **Lưu ý:** Import path `"github.com/amy/tech-talk/polymorphism"` là đường dẫn ví dụ minh họa. Khi áp dụng vào project thực tế, cần thay bằng module path tương ứng trong `go.mod` của dự án.

---

## Phân Tích

Trong ví dụ trên, hàm `SaySlogan` nhận tham số kiểu `SloganSayer` — một interface — thay vì nhận trực tiếp `*Hillary` hay `*Trump`. Nhờ đó:

- Khi truyền `hillary` vào `SaySlogan`, Go gọi `(*Hillary).Slogan()`.
- Khi truyền `trump` vào `SaySlogan`, Go gọi `(*Trump).Slogan()`.

Cùng một lời gọi hàm `SaySlogan(...)`, nhưng hành vi thực tế phụ thuộc vào kiểu cụ thể của đối số được truyền vào lúc runtime. Đây chính là **tính đa hình**: một interface, nhiều hành vi.

Ta không cần biết hay quan tâm đến kiểu cụ thể của đối số. Miễn là một kiểu implement đầy đủ các method của `SloganSayer` interface, nó đều có thể được truyền vào `SaySlogan` mà không cần thay đổi gì ở phía hàm nhận.

---

## Mở Rộng

1. **Interface lồng nhau (Interface embedding):** Go cho phép một interface bao gồm (embed) các interface khác, tạo ra interface phức hợp. Ví dụ: `io.ReadWriter` là tổ hợp của `io.Reader` và `io.Writer`. Đây là cách Go xây dựng hệ thống interface có tầng lớp mà không cần kế thừa.
    
2. **Empty interface và type assertion:** `interface{}` (hoặc `any` từ Go 1.18) là interface không có method nào — mọi kiểu đều thỏa mãn nó. Kết hợp với **type assertion** (`x.(T)`) và **type switch**, đây là cách Go xử lý các trường hợp cần làm việc với nhiều kiểu khác nhau mà không biết trước kiểu cụ thể.
    
3. **Interface và testability:** Tính đa hình thông qua interface là nền tảng để viết code dễ test trong Go. Bằng cách inject dependency dưới dạng interface thay vì struct cụ thể, bạn có thể dễ dàng thay thế bằng mock object khi viết unit test — mà không cần framework mock phức tạp.
    
4. **So sánh với generics (Go 1.18+):** Kể từ Go 1.18, Go có generics với type constraint. Trong một số trường hợp, generics và interface đều có thể giải quyết bài toán "xử lý nhiều kiểu", nhưng với trade-off khác nhau về hiệu năng runtime và tính rõ ràng của code. Đây là chủ đề đáng tìm hiểu tiếp theo sau khi nắm vững polymorphism qua interface.