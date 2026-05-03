# Composition (có thể hiểu như inheritance)

Trong **Go**, thừa kế là không tồn tại. Thay vào đó, chúng ta xây dựng các cấu trúc với các yếu tố tổng hợp và tái sử dựng thông qua **embedding** (nhúng).

Go cho phép chúng ta embed các loại bên trong interface hoặc structs. Thông qua embedding, chúng ta có thể biến các phương thức được included từ bên trong hay bên ngoài.

> When we embed a type, the methods of that type become methods of the outer type, but when they are invoked the receiver of the method is the inner type, not the outer one.

```go 
package composition 
import "fmt"
type Human struct {
    FirstName   string
    LastName    string
    CanSwim     bool
}
// Amy được embedded với kiểu Human
// và do đó Amy có thể gọi các phương thức của Human
type Amy struct {
    Human
}
// Alan được embedded với kiểu Human 
type Alan struct {
    Human
}
func (h *Human) Name() { 
    fmt.Printf("Hello! My name is %v %v", h.FirstName, h.LastName)
}

func (h *Human) Swim() {
    
    if h.CanSwim == true {
        fmt.Println("I can swim!")
    } else {
        fmt.Println("I can not swim.")
    }
}
```