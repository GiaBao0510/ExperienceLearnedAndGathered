# Composition (Thay thế cho Inheritance)

**Golang không hỗ trợ tính kế thừa (inheritance)** như các ngôn ngữ lập trình hướng đối tượng truyền thống (Java, C++, C#). Để đạt mục đích tái sử dụng mã và đa hình tương tự, Go sử dụng kỹ thuật **Composition (thành phần)** kết hợp **Struct Embedding (nhúng struct)** và **Interface**.

> "Favor composition over inheritance." — Nguyên tắc thiết kế phần mềm mà Go áp dụng triệt để.

---

## Struct Embedding (Nhúng struct)

Trong Go, bạn có thể nhúng một struct vào bên trong một struct khác mà không cần đặt tên trường — gọi là **anonymous field** (trường ẩn danh). Khi đó, mọi trường và phương thức của struct được nhúng sẽ được **thăng cấp** (promoted) lên struct ngoài, cho phép truy cập trực tiếp như thể chúng thuộc về struct ngoài.

> **Lưu ý quan trọng:** Khi một phương thức được thăng cấp và được gọi, **receiver vẫn là struct bên trong**, không phải struct bên ngoài. Đây là điểm khác biệt then chốt so với kế thừa trong OOP truyền thống.

### Ví dụ 1 — Nhiều struct nhúng cùng một kiểu

```go
package main

import "fmt"

type Human struct {
    FirstName string
    LastName  string
    CanSwim   bool
}

func (h *Human) Name() {
    fmt.Printf("Hello! My name is %v %v\n", h.FirstName, h.LastName)
}

func (h *Human) Swim() {
    if h.CanSwim {
        fmt.Println("I can swim!")
    } else {
        fmt.Println("I cannot swim.")
    }
}

// Amy và Alan đều nhúng Human — chia sẻ hành vi mà không cần kế thừa
type Amy struct {
    Human
}

type Alan struct {
    Human
}

func main() {
    amy := Amy{Human: Human{FirstName: "Amy", LastName: "Nguyen", CanSwim: true}}
    alan := Alan{Human: Human{FirstName: "Alan", LastName: "Tran", CanSwim: false}}

    amy.Name()  // Hello! My name is Amy Nguyen
    amy.Swim()  // I can swim!

    alan.Name() // Hello! My name is Alan Tran
    alan.Swim() // I cannot swim.
}
```

### Ví dụ 2 — Struct embedding cơ bản

```go
package main

import "fmt"

// Struct cơ sở
type Animal struct {
    Name string
}

func (a Animal) Speak() {
    fmt.Println("...", a.Name, "lên tiếng...")
}

// Dog nhúng Animal (anonymous field — không đặt tên trường)
type Dog struct {
    Animal
    Breed string
}

func main() {
    myDog := Dog{
        Animal: Animal{Name: "Buddy"},
        Breed:  "Golden Retriever",
    }

    // Trường và phương thức của Animal được thăng cấp lên Dog
    fmt.Println("Tên:", myDog.Name)   // truy cập trực tiếp myDog.Name thay vì myDog.Animal.Name
    fmt.Println("Giống:", myDog.Breed)
    myDog.Speak()                     // gọi trực tiếp phương thức của Animal
}
```

---

## Ghi đè phương thức (Method Overriding)

Không giống kế thừa trong Java/C++, Go không có từ khóa `override`. Tuy nhiên bạn có thể **định nghĩa lại phương thức cùng tên** trên struct ngoài để "che" phương thức của struct được nhúng:

```go
package main

import "fmt"

type Animal struct {
    Name string
}

func (a Animal) Speak() {
    fmt.Println(a.Name, "kêu: ...")
}

type Cat struct {
    Animal
}

// Cat định nghĩa lại Speak — che phương thức Speak của Animal
func (c Cat) Speak() {
    fmt.Println(c.Name, "kêu: Meow!")
}

func main() {
    myCat := Cat{Animal: Animal{Name: "Whiskers"}}
    myCat.Speak()        // Whiskers kêu: Meow!  (phương thức của Cat)
    myCat.Animal.Speak() // Whiskers kêu: ...    (truy cập trực tiếp phương thức gốc)
}
```

> **Lưu ý:** Phương thức của Animal không bị xóa — bạn vẫn có thể gọi tường minh qua `myCat.Animal.Speak()`.

---

## Composition với Interface

Sức mạnh thực sự của composition trong Go đến từ việc kết hợp **struct embedding** với **interface**. Đây là cách Go đạt được tính đa hình mà không cần kế thừa:

```go
package main

import "fmt"

type Speaker interface {
    Speak()
}

type Animal struct {
    Name string
}

type Dog struct {
    Animal
}

func (d Dog) Speak() {
    fmt.Println(d.Name, "kêu: Woof!")
}

type Cat struct {
    Animal
}

func (c Cat) Speak() {
    fmt.Println(c.Name, "kêu: Meow!")
}

// Hàm chỉ phụ thuộc vào interface — không quan tâm kiểu cụ thể
func MakeNoise(s Speaker) {
    s.Speak()
}

func main() {
    dog := Dog{Animal: Animal{Name: "Buddy"}}
    cat := Cat{Animal: Animal{Name: "Whiskers"}}

    MakeNoise(dog) // Buddy kêu: Woof!
    MakeNoise(cat) // Whiskers kêu: Meow!
}
```

---

## Tại sao Go thiết kế như vậy?

- **Tránh cây phân cấp phức tạp:** Ngăn ngừa "diamond problem" (vấn đề đa kế thừa hình thoi) và các cây kế thừa nhiều tầng khó bảo trì.
- **Đơn giản và minh bạch:** Truy cập trực tiếp vào phương thức được thăng cấp mà không cần từ khóa `super` hay `this`. Không có hành vi ẩn nào từ lớp cha.
- **Khuyến khích lắp ghép (Composition):** Tạo ra các đơn vị code nhỏ, độc lập, dễ hiểu và dễ bảo trì — thay vì buộc mọi hành vi phải gắn vào một cây cấu trúc kế thừa cứng nhắc.
- **Tránh "fragile base class problem":** Trong OOP truyền thống, thay đổi lớp cha có thể làm hỏng lớp con ngoài ý muốn. Composition tránh được rủi ro này vì không có mối quan hệ IS-A cứng giữa các kiểu.

---

## Tổng kết: So sánh Inheritance vs Composition

|Tiêu chí|Inheritance (Java/C++)|Composition (Go)|
|---|---|---|
|Cơ chế|Lớp con kế thừa lớp cha (`extends`)|Nhúng struct (`embedding`)|
|Quan hệ|IS-A (là một loại)|HAS-A (có chứa)|
|Đa kế thừa|Phức tạp / hạn chế|Nhúng nhiều struct tùy ý|
|Ghi đè|`override` tường minh|Định nghĩa lại phương thức cùng tên|
|Gọi phương thức gốc|`super.method()`|`outer.Inner.method()`|
|Tính đa hình|Qua kế thừa + override|Qua interface|
