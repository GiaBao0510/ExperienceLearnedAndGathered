> “Abstraction means working with something we know how to use without knowing how it works internally.”

Tương tự như embedding các structs bên trong một struct, chúng ta cũng có thể embed các interfaces trong các structs. Bất kỳ kiểu nào thỏa mãn interface nào cũng sẽ sử dụng được interface đó.

```go
package abstraction
import "fmt"
type SloganSayer interface {
    Slogan()
}
// Campaign có thể accept a SloganSayer trong quá trình khởi tạo
// Campaign cũng là một SloganSayer bởi vì nó cũng implements SloganSayer interface.
type Campaign struct{
    SloganSayer
}
// SaySlogan cũng có thể accept Campaign như là một tham số truyền vào!
func SaySlogan(s SloganSayer) {
    s.Slogan()
}
// Hillary implements the SloganSayer interface
// Hillary là một SloganSayer
type Hillary struct{}
func (h *Hillary) Slogan() {
    fmt.Println("Stronger together.")
}
// Tương tự với Trump 
type Trump struct{}
func (t *Trump) Slogan() {
    fmt.Println("Make American great again.")
}
```

```go
package main
import "github.com/amy/tech-talk/abstraction"
func main() {
    hillary := new(abstraction.Hillary)
    trump := new(abstraction.Trump)
    h := abstraction.Campaign{hillary}
    t := abstraction.Campaign{trump}
    // Triển khai slogan tranh cử của Trump và hilary được trừu tượng hóa đi.
    // Thay vào đó. Campaign chỉ biết rằng có đó là một SloganSayer
    // và do đó có thể gọi Slogan.
    h.Slogan()  // "Stronger together."
    t.Slogan()  // "Make America great again."
    // Chúng ta có thể inject một  SloganSayer vào tham số SaySlogan
    abstraction.SaySlogan(hillary)  // "Stronger together."
    abstraction.SaySlogan(trump)    // "Make America great again."
    // h và t cũng là một loại campaign
    abstraction.SaySlogan(h)  // "Stronger together."
    abstraction.SaySlogan(t)  // "Make America great again."
}
```