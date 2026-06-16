> "Encapsulation is the mechanism of hiding data implementation by restricting access to public methods."

## Tính đóng gói trong Go

Trong hầu hết các ngôn ngữ OOP phổ biến, tính đóng gói được thực hiện thông qua từ khóa `private` / `public` / `protected` trên các trường và phương thức của lớp. **Trong Go, đóng gói được thực hiện ở cấp độ package**, không phải cấp độ lớp.

Quy tắc rất đơn giản:

- **Exported (công khai):** Tên bắt đầu bằng **chữ hoa** → có thể truy cập từ bên ngoài package.
- **Unexported (nội bộ):** Tên bắt đầu bằng **chữ thường** → chỉ có thể truy cập trong cùng package.

Quy tắc này áp dụng nhất quán cho tất cả: hằng số, biến, hàm, kiểu dữ liệu, trường struct, phương thức, v.v.

> **Lưu ý thuật ngữ:** Go không dùng "public/private" — thuật ngữ chính xác là **exported** và **unexported**. Trong nội bộ một package, mọi identifier đều có thể truy cập tự do dù viết hoa hay viết thường.

---

## Ví dụ

### Định nghĩa (package `encapsulation`)

```go
package encapsulation

import "fmt"

// Encapsulation — exported struct (tên viết hoa)
type Encapsulation struct{}

// Expose — exported method, có thể gọi từ package khác
func (e *Encapsulation) Expose() {
    fmt.Println("AHHHH! I'm exposed!")
}

// hide — unexported method, chỉ dùng được trong package này
func (e *Encapsulation) hide() {
    fmt.Println("Shhhh... this is super secret")
}

// Unhide — exported method gọi nội bộ hide()
func (e *Encapsulation) Unhide() {
    e.hide()
    fmt.Println("...jk")
}
```

Trong package `encapsulation`: `Encapsulation` (struct), `Expose` (method) và `Unhide` (method) đều exported — có thể dùng từ package khác. `hide` là unexported — chỉ dùng được trong nội bộ package.

### Sử dụng (package `main`)

```go
import "github.com/amy/tech-talk/encapsulation"

func main() {
    e := encapsulation.Encapsulation{}

    e.Expose()   // OK — exported method

    // e.hide() // Lỗi compile: cannot refer to unexported method
                // encapsulation.(*Encapsulation).hide

    e.Unhide()   // OK — Unhide là exported, nó gọi hide() nội bộ
}
```

---

## Đóng gói với trường struct và Getter/Setter

Trong thực tế, tính đóng gói thường được áp dụng trên **trường của struct**: giữ trường ở dạng unexported và cung cấp phương thức exported để đọc/ghi có kiểm soát.

```go
package account

// BankAccount — trường balance được bảo vệ (unexported)
type BankAccount struct {
    balance float64
}

// Deposit — kiểm tra điều kiện trước khi thay đổi dữ liệu
func (a *BankAccount) Deposit(amount float64) error {
    if amount <= 0 {
        return fmt.Errorf("số tiền nạp phải lớn hơn 0")
    }
    a.balance += amount
    return nil
}

// Balance — chỉ cho phép đọc, không cho phép ghi trực tiếp
func (a *BankAccount) Balance() float64 {
    return a.balance
}
```

Nhờ đóng gói, không có code nào bên ngoài package có thể gán trực tiếp `a.balance = -999` — toàn bộ logic kiểm tra nằm trong package.

---

## Câu hỏi phỏng vấn thường gặp

### 1. Tại sao Go không đi theo hướng OOP truyền thống như Java hay C#?

Go được thiết kế với triết lý **đơn giản và thực dụng**. Các tác giả (Rob Pike, Ken Thompson, Robert Griesemer) nhận thấy rằng OOP dựa trên kế thừa lớp (class inheritance) trong các ngôn ngữ lớn thường dẫn đến:

- Cây kế thừa phức tạp, khó bảo trì.
- Sự phụ thuộc chặt chẽ (tight coupling) giữa lớp cha và lớp con.
- Overhead khái niệm không cần thiết cho hầu hết bài toán thực tế.

Go chọn **composition over inheritance** — thay vì kế thừa, bạn nhúng struct (embedding) và định nghĩa hành vi qua interface. Kết quả là code linh hoạt hơn, ít ràng buộc hơn, và dễ test hơn.

### 2. Tại sao Go không dùng lớp (class)?

Go không có `class`, nhưng có `struct` + `method`. Về bản chất, bạn vẫn có thể gắn hành vi (method) vào dữ liệu (struct) — đây là cốt lõi của OOP. Điều Go loại bỏ là **kế thừa dựa trên lớp** (class-based inheritance), không phải khái niệm đóng gói hay đa hình.

Thay vào đó, Go dùng:

- **Struct embedding** để tái sử dụng code (thay cho kế thừa).
- **Interface** để đạt tính đa hình và trừu tượng (thay cho abstract class).

Điều này giúp tránh vấn đề kinh điển của OOP: _"fragile base class problem"_ — lớp con bị ảnh hưởng ngoài ý muốn khi lớp cha thay đổi.

### 3. Tại sao phạm vi truy cập trong Go không dùng private/public/protected như các ngôn ngữ khác?

Go đơn giản hóa mô hình truy cập xuống còn **hai mức**: exported và unexported, dựa trên quy ước đặt tên thay vì từ khóa. Lý do:

- **Không cần `protected`:** Go không có kế thừa, nên khái niệm "chỉ lớp con mới truy cập được" không tồn tại.
- **Package là đơn vị đóng gói tự nhiên:** Thay vì kiểm soát ở từng lớp, Go kiểm soát ở cấp độ package — toàn bộ code trong một package được xem là cùng một "đơn vị tin cậy".
- **Đơn giản hóa ngôn ngữ:** Một quy tắc duy nhất (viết hoa/viết thường) áp dụng nhất quán cho mọi thứ, giảm tải nhận thức cho lập trình viên.