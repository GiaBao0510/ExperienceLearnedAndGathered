> Encapsulation is the mechanism of hiding of data implementation by restricting access to public methods.

Trong hầu hết các ngôn ngữ lập trình phổ biến, việc đóng gói trong OOP dựa trên lớp đạt được thông qua private và public class variables / methods. Trong Go, đóng gói được thực hiện trên các **package** level.

Các thành phần "public" có thể được xuất ra bên ngoài các package và được trình bày bằng cách viết hoa chữ cái đầu tiên. Ở đây, publlic được đặt trong dấu ngoặc kép bởi vì thuật ngữ chính xác hơn là exported và unexported elements, tuy nhiên dùng từ public sẽ giúp chúng ta nắm bắt nhanh hơn. Unexported elements được chỉ định bằng chữ cái đầu tiên và chỉ có thể truy cập được trong package tương ứng.

> Public/protected/private là những từ khóa liên quan đến các lớp, trong khi exporting/importing liên quan đến các packages.

```go
package encapsulation
import "fmt"
// Encapsulation struct có thể exported ra bên ngoài pagekage này (Encapsulation viết hoa chữ cái đầu)
type Encapsulation struct{}

// Hàm Expose có thể exported ra bên ngoài pagekage này (Expose viết hoa chữ cái đầu)
func (e *Encapsulation) Expose() {
    fmt.Println("AHHHH! I'm exposed!")
}

// hàm hide chỉ có thể sử dụng trong package này (hide viết thường chữ cái đầu)
func (e *Encapsulation) hide() {
     fmt.Println("Shhhh... this is super secret")
}

// Unhide sử dụng hàm hide chưa được exported
func (e *Encapsulation) Unhide() {
     e.hide()
     fmt.Println("...jk")
}
```

Trong package _encapsulation_, Encapsulation (struct), Expose (method), và Unhide (method) tất cả đều có thể được sử dụng từ các packages khác.

```go
import "github.com/amy/tech-talk/encapsulation"
func main() {
    e := encapsulation.Encapsulation{}    
    e.Expose()    
    // e.hide()    //nếu bỏ comment,  xuất hiện lỗi
                   // ./main.go:10: e.hide undefined (cannot refer
                   // to unexported field or method encapsulation.
                   // (*Encapsulation)."".hide)
    
    e.Unhide()
}
```

Chung quy lại, trong Go khái niệm đóng gói khá đơn giản: chữ cái đầu tên viết hoa thì mở, còn viết thường thì đóng. Quy tắt này áp dụng cho hằng, biến, hàm, trường, phương thức, v.v... Có điều trong Go, khái niệm mở hay đóng chỉ áp dụng bên ngoài package. Trong package, mọi cái đều mở dù tên viết hoa hay viết thường.