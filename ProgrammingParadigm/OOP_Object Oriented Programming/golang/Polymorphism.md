> Poorphism describes a pattern in object oriented programming in which classes have different functionality while sharing a common interface.

Như chúng ta thường thấy ở các ngôn ngữ lập trình hướng đối tượng, tính đa hình thể hiện khi các lớp kế thừa cùng một lớp. Với việc sử dụng interface, mặc dù không có khái niệm kế thừa nhưng Go cũng hỗ trợ tính đa hình theo cách riêng của nó.

```go
package polymorphism 
import "fmt"
type SloganSayer interface {
    Slogan()
}

// SaySlogan truyền vào một tham số kiểu SloganSayer
func SaySlogan(sloganSayer SloganSayer) {
    sloganSayer.Slogan()
}

// Hillary thỏa mãn SloganSayer interfa
// bằng việc thực thi function Slogan.
// Vì vậy, Hillary cũng là một kiểu của SloganSayer.
type Hillary struct{}
func (h *Hillary) Slogan() {
    fmt.Println("Stronger together.")
}

// tương tự với struct Trump
type Trump struct{}
func (t *Trump) Slogan() {
    fmt.Println("Make America great again.")
}
```

```go
package main 
import "github.com/amy/tech-talk/polymorphism"
func main() {
    hillary := new(polymorphism.Hillary)
    hillary.Slogan()                  // "Stronger together."
    polymorphism.SaySlogan(hillary)   // "Stronger together."
    trump := new(polymorphism.Trump)
    polymorphism.SaySlogan(trump)     // "Make America great again."
}
```

Trong ví dụ trên, ta không cần quan tâm ứng cử viên nào đang nói khẩu hiệu. Miễn là một kiểu implements của SloganSayer interface, chúng ta có thể truyền nó vào SaySlogan.