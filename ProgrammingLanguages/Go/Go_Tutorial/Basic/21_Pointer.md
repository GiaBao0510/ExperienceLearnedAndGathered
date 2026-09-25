# Pointer (Con Trỏ) trong Go - Hướng dẫn Chi tiết

## 📋 Mục lục

1. [Pointer là gì?](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#1-pointer-l%C3%A0-g%C3%AC)
2. [Cách hoạt động của Pointer](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#2-c%C3%A1ch-ho%E1%BA%A1t-%C4%91%E1%BB%99ng-c%E1%BB%A7a-pointer)
3. [Lấy địa chỉ và Dereference](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#3-l%E1%BA%A5y-%C4%91%E1%BB%8Ba-ch%E1%BB%89-v%C3%A0-dereference)
4. [Pass by Value vs Pass by Reference](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#4-pass-by-value-vs-pass-by-reference)
5. [Pointer to Pointer](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#5-pointer-to-pointer)
6. [Pointer với Struct](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#6-pointer-v%E1%BB%9Bi-struct)
7. [Pointer với Slice, Map](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#7-pointer-v%E1%BB%9Bi-slice-map)
8. [Nil Pointer](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#8-nil-pointer)
9. [Best Practices](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#9-best-practices)
10. [Ví dụ thực tế](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#10-v%C3%AD-d%E1%BB%A5-th%E1%BB%B1c-t%E1%BA%BF)

---

## 1. Pointer là gì?

**Pointer (Con trỏ)** là một biến đặc biệt lưu trữ **địa chỉ bộ nhớ** của một biến khác, thay vì lưu giá trị trực tiếp.

### 🎯 Hình ảnh minh họa

**Biến thông thường:**

```
┌─────────────┐
│ name        │ = "John"
│ (0xc000010)│
└─────────────┘
```

**Pointer:**

```
┌─────────────┐         ┌─────────────┐
│ pName       │ ──────→ │ name        │
│ (0xc000020) │         │ (0xc000010) │
│ = 0xc000010 │         │ = "John"    │
└─────────────┘         └─────────────┘
```

### 💡 Ví dụ thực tế

**Giống như địa chỉ nhà:**

```
Biến thông thường:
"Nhà của tôi ở số 123 Nguyễn Huệ"
→ Lưu trực tiếp thông tin

Pointer:
"Đi đến địa chỉ 123 Nguyễn Huệ để tìm nhà của tôi"
→ Lưu địa chỉ, không lưu nhà
```

### 📊 So sánh

|Khía cạnh|Biến thường|Pointer|
|---|---|---|
|**Lưu trữ**|Giá trị|Địa chỉ|
|**Kiểu dữ liệu**|`string`|`*string`|
|**Kích thước**|Tùy giá trị|8 bytes (64-bit)|
|**Truy cập giá trị**|Trực tiếp|Qua `*`|

---

## 2. Cách hoạt động của Pointer

### 2.1. Bộ nhớ và địa chỉ

```go
package main

import "fmt"

func main() {
    name := "Pham Gia Bao"
    
    fmt.Printf("Kiểu dữ liệu: %T\n", name)        // string
    fmt.Printf("Giá trị: %v\n", name)             // Pham Gia Bao
    fmt.Printf("Địa chỉ: %p\n", &name)            // 0xc0000220c0
}
```

**Output:**

```
Kiểu dữ liệu: string
Giá trị: Pham Gia Bao
Địa chỉ: 0xc0000220c0
```

**Giải thích:**

- `name` lưu giá trị `"Pham Gia Bao"`
- Giá trị này được lưu ở địa chỉ `0xc0000220c0` trong RAM
- `&name` cho ta địa chỉ đó

### 2.2. Memory Layout

```
RAM (Bộ nhớ):
┌──────────────┬──────────────────┬─────────────┐
│   Địa chỉ    │     Giá trị      │   Biến      │
├──────────────┼──────────────────┼─────────────┤
│ 0xc0000220c0 │ "Pham Gia Bao"   │ name        │
│ 0xc0000220e0 │ 0xc0000220c0     │ pName       │
│ 0xc000022100 │ 0xc0000220e0     │ pName2      │
└──────────────┴──────────────────┴─────────────┘
```

---

## 3. Lấy địa chỉ và Dereference

### 3.1. Toán tử `&` - Lấy địa chỉ

```go
package main

import "fmt"

func main() {
    age := 25
    
    fmt.Printf("Giá trị age: %v\n", age)      // 25
    fmt.Printf("Địa chỉ age: %p\n", &age)     // 0xc000018098
}
```

**Syntax:**

```go
&variable  // Lấy địa chỉ của biến
```

### 3.2. Toán tử `*` - Dereference

**2 cách dùng `*`:**

**1. Khai báo kiểu pointer:**

```go
var pName *string  // Pointer to string
var pAge *int      // Pointer to int
```

**2. Dereference (truy cập giá trị):**

```go
value := *pointer  // Lấy giá trị tại địa chỉ
```

### 3.3. Ví dụ đầy đủ

```go
package main

import "fmt"

func main() {
    name := "Nguyen Van A"
    
    // Tạo pointer
    pName := &name
    
    fmt.Println("=== Biến gốc ===")
    fmt.Printf("Kiểu: %T\n", name)
    fmt.Printf("Giá trị: %v\n", name)
    fmt.Printf("Địa chỉ: %p\n", &name)
    
    fmt.Println("\n=== Pointer ===")
    fmt.Printf("Kiểu: %T\n", pName)           // *string
    fmt.Printf("Giá trị (địa chỉ): %p\n", pName)
    fmt.Printf("Dereference (*pName): %v\n", *pName)
    
    // Sửa qua pointer
    *pName = "Tran Thi B"
    
    fmt.Println("\n=== Sau khi sửa ===")
    fmt.Printf("name: %v\n", name)             // Đã thay đổi!
    fmt.Printf("*pName: %v\n", *pName)
}
```

**Output:**

```
=== Biến gốc ===
Kiểu: string
Giá trị: Nguyen Van A
Địa chỉ: 0xc0000220c0

=== Pointer ===
Kiểu: *string
Giá trị (địa chỉ): 0xc0000220c0
Dereference (*pName): Nguyen Van A

=== Sau khi sửa ===
name: Tran Thi B
*pName: Tran Thi B
```

### 3.4. Bảng tóm tắt toán tử

|Toán tử|Tên|Ý nghĩa|Ví dụ|
|---|---|---|---|
|`&`|Address-of|Lấy địa chỉ|`&name`|
|`*` (khai báo)|Pointer type|Kiểu pointer|`var p *int`|
|`*` (dereference)|Dereference|Lấy giá trị|`*pName`|

---

## 4. Pass by Value vs Pass by Reference

### 4.1. Pass by Value (Truyền giá trị)

**Go mặc định truyền theo giá trị = copy**

```go
package main

import "fmt"

func updateName(name string) {
    name = "Nguyen Van B"  // Chỉ thay đổi bản copy
    fmt.Printf("Trong hàm - Giá trị: %v, Địa chỉ: %p\n", name, &name)
}

func main() {
    name := "Nguyen Van A"
    fmt.Printf("Trước - Giá trị: %v, Địa chỉ: %p\n", name, &name)
    
    updateName(name)  // Truyền copy
    
    fmt.Printf("Sau  - Giá trị: %v, Địa chỉ: %p\n", name, &name)
}
```

**Output:**

```
Trước - Giá trị: Nguyen Van A, Địa chỉ: 0xc0000220c0
Trong hàm - Giá trị: Nguyen Van B, Địa chỉ: 0xc0000220e0  ← Địa chỉ khác!
Sau  - Giá trị: Nguyen Van A, Địa chỉ: 0xc0000220c0      ← Không đổi
```

**Giải thích:**

```
main()                    updateName()
┌───────────┐            ┌───────────┐
│ name      │  copy →    │ name      │
│ "Van A"   │            │ "Van B"   │
│ 0xc000.c0 │            │ 0xc000.e0 │
└───────────┘            └───────────┘
    ↑                          ↑
Không thay đổi          Chỉ thay đổi copy
```

### 4.2. Pass by Reference (Truyền tham chiếu)

**Dùng pointer để truyền địa chỉ**

```go
package main

import "fmt"

func updateName(pName *string) {
    *pName = "Nguyen Van B"  // Thay đổi giá trị gốc
    fmt.Printf("Trong hàm - Giá trị: %v, Địa chỉ: %p\n", *pName, pName)
}

func main() {
    name := "Nguyen Van A"
    fmt.Printf("Trước - Giá trị: %v, Địa chỉ: %p\n", name, &name)
    
    updateName(&name)  // Truyền địa chỉ
    
    fmt.Printf("Sau  - Giá trị: %v, Địa chỉ: %p\n", name, &name)
}
```

**Output:**

```
Trước - Giá trị: Nguyen Van A, Địa chỉ: 0xc0000220c0
Trong hàm - Giá trị: Nguyen Van B, Địa chỉ: 0xc0000220c0  ← Cùng địa chỉ!
Sau  - Giá trị: Nguyen Van B, Địa chỉ: 0xc0000220c0      ← Đã thay đổi!
```

**Giải thích:**

```
main()                    updateName()
┌───────────┐            ┌───────────┐
│ name      │  địa chỉ → │ pName     │
│ "Van B"   │ ←──────────│ 0xc000.c0 │
│ 0xc000.c0 │            └───────────┘
└───────────┘
    ↑
Thay đổi trực tiếp
```

### 4.3. So sánh

```go
package main

import "fmt"

// Pass by Value
func addByValue(x int) {
    x = x + 10
    fmt.Printf("Trong addByValue: x = %d (địa chỉ: %p)\n", x, &x)
}

// Pass by Reference
func addByReference(px *int) {
    *px = *px + 10
    fmt.Printf("Trong addByReference: *px = %d (địa chỉ: %p)\n", *px, px)
}

func main() {
    num := 5
    fmt.Printf("Ban đầu: num = %d (địa chỉ: %p)\n\n", num, &num)
    
    // By Value
    addByValue(num)
    fmt.Printf("Sau addByValue: num = %d\n\n", num)  // Vẫn 5
    
    // By Reference
    addByReference(&num)
    fmt.Printf("Sau addByReference: num = %d\n", num)  // Đã là 15
}
```

**Output:**

```
Ban đầu: num = 5 (địa chỉ: 0xc000018098)

Trong addByValue: x = 15 (địa chỉ: 0xc0000180b0)
Sau addByValue: num = 5

Trong addByReference: *px = 15 (địa chỉ: 0xc000018098)
Sau addByReference: num = 15
```

---

## 5. Pointer to Pointer

### 5.1. Khái niệm

**Pointer to Pointer** là pointer trỏ đến một pointer khác.

**Syntax:**

```go
var pp **int  // Pointer to pointer to int
```

### 5.2. Ví dụ

```go
package main

import "fmt"

func main() {
    name := "Nguyen Van A"
    
    // Level 1: Pointer to string
    pName := &name
    
    // Level 2: Pointer to pointer to string
    pName2 := &pName
    
    fmt.Println("=== name ===")
    fmt.Printf("Kiểu: %T\n", name)
    fmt.Printf("Giá trị: %v\n", name)
    fmt.Printf("Địa chỉ: %p\n\n", &name)
    
    fmt.Println("=== pName ===")
    fmt.Printf("Kiểu: %T\n", pName)
    fmt.Printf("Giá trị (địa chỉ name): %p\n", pName)
    fmt.Printf("Dereference: %v\n", *pName)
    fmt.Printf("Địa chỉ pName: %p\n\n", &pName)
    
    fmt.Println("=== pName2 ===")
    fmt.Printf("Kiểu: %T\n", pName2)
    fmt.Printf("Giá trị (địa chỉ pName): %p\n", pName2)
    fmt.Printf("*pName2 (địa chỉ name): %p\n", *pName2)
    fmt.Printf("**pName2 (giá trị name): %v\n", **pName2)
}
```

**Output:**

```
=== name ===
Kiểu: string
Giá trị: Nguyen Van A
Địa chỉ: 0xc0000220c0

=== pName ===
Kiểu: *string
Giá trị (địa chỉ name): 0xc0000220c0
Dereference: Nguyen Van A
Địa chỉ pName: 0xc000044040

=== pName2 ===
Kiểu: **string
Giá trị (địa chỉ pName): 0xc000044040
*pName2 (địa chỉ name): 0xc0000220c0
**pName2 (giá trị name): Nguyen Van A
```

### 5.3. Diagram

```
┌──────────────┐
│ name         │
│ "Van A"      │
│ 0xc000...c0  │
└──────────────┘
       ↑
       │
┌──────────────┐
│ pName        │
│ 0xc000...c0  │  ← Trỏ đến name
│ 0xc000...40  │
└──────────────┘
       ↑
       │
┌──────────────┐
│ pName2       │
│ 0xc000...40  │  ← Trỏ đến pName
└──────────────┘
```

### 5.4. Khi nào dùng?

**✅ Dùng Pointer to Pointer khi:**

- Cần thay đổi pointer trong function
- Linked list, tree structures
- Dynamic memory allocation

**Ví dụ: Thay đổi pointer**

```go
package main

import "fmt"

func changePointer(pp **int) {
    newValue := 999
    *pp = &newValue  // Thay đổi pointer gốc
}

func main() {
    x := 10
    y := 20
    
    p := &x
    fmt.Printf("p trỏ đến x: %d\n", *p)  // 10
    
    changePointer(&p)
    fmt.Printf("p trỏ đến giá trị mới: %d\n", *p)  // 999
}
```

---

## 6. Pointer với Struct

### 6.1. Tại sao cần Pointer với Struct?

**Struct thường lớn hơn kiểu dữ liệu cơ bản → Nên dùng pointer**

```go
package main

import "fmt"

type Person struct {
    Name  string
    Age   int
    Email string
}

// Pass by Value - Copy toàn bộ struct
func updateAgeByValue(p Person, newAge int) {
    p.Age = newAge
    fmt.Printf("Trong hàm (value): Age = %d, Địa chỉ: %p\n", p.Age, &p)
}

// Pass by Reference - Chỉ copy địa chỉ (8 bytes)
func updateAgeByPointer(p *Person, newAge int) {
    p.Age = newAge  // Go tự dereference
    fmt.Printf("Trong hàm (pointer): Age = %d, Địa chỉ: %p\n", p.Age, p)
}

func main() {
    person := Person{Name: "John", Age: 25, Email: "john@gmail.com"}
    
    fmt.Printf("Ban đầu: Age = %d, Địa chỉ: %p\n\n", person.Age, &person)
    
    // By Value
    updateAgeByValue(person, 30)
    fmt.Printf("Sau Value: Age = %d\n\n", person.Age)  // Vẫn 25
    
    // By Pointer
    updateAgeByPointer(&person, 30)
    fmt.Printf("Sau Pointer: Age = %d\n", person.Age)  // Đã 30
}
```

**Output:**

```
Ban đầu: Age = 25, Địa chỉ: 0xc000044060

Trong hàm (value): Age = 30, Địa chỉ: 0xc0000440a0  ← Địa chỉ khác
Sau Value: Age = 25  ← Không đổi

Trong hàm (pointer): Age = 30, Địa chỉ: 0xc000044060  ← Cùng địa chỉ
Sau Pointer: Age = 30  ← Đã đổi
```

### 6.2. Auto Dereference với Struct

**Go tự động dereference khi truy cập field**

```go
type Person struct {
    Name string
    Age  int
}

func main() {
    p := &Person{Name: "John", Age: 25}
    
    // Cả hai cách đều OK
    fmt.Println(p.Name)    // Go tự dereference
    fmt.Println((*p).Name) // Dereference thủ công
    
    // Sửa field
    p.Age = 30             // Go tự dereference
    (*p).Age = 30          // Dereference thủ công
}
```

### 6.3. Constructor với Pointer

```go
package main

import "fmt"

type Employee struct {
    ID     int
    Name   string
    Salary float64
}

// Constructor trả về pointer
func NewEmployee(id int, name string, salary float64) *Employee {
    return &Employee{
        ID:     id,
        Name:   name,
        Salary: salary,
    }
}

// Method với pointer receiver
func (e *Employee) IncreaseSalary(amount float64) {
    e.Salary += amount
}

func main() {
    emp := NewEmployee(1, "John Doe", 50000)
    
    fmt.Printf("Trước: %.2f\n", emp.Salary)
    emp.IncreaseSalary(5000)
    fmt.Printf("Sau: %.2f\n", emp.Salary)
}
```

**Output:**

```
Trước: 50000.00
Sau: 55000.00
```

---

## 7. Pointer với Slice, Map

### 7.1. Slice và Map đã là reference type

**⚠️ Lưu ý:** Slice và Map trong Go đã là reference type, không cần pointer!

```go
package main

import "fmt"

// Slice - Không cần pointer
func appendValue(s []int, value int) {
    s = append(s, value)  // Thay đổi nội bộ slice OK
    fmt.Printf("Trong hàm: %v\n", s)
}

func main() {
    slice := []int{1, 2, 3}
    fmt.Printf("Trước: %v\n", slice)
    
    appendValue(slice, 4)
    
    fmt.Printf("Sau: %v\n", slice)  // Vẫn [1 2 3]
}
```

**Tuy nhiên, nếu muốn thay đổi slice header (len, cap), cần pointer:**

```go
func appendValuePointer(s *[]int, value int) {
    *s = append(*s, value)  // Thay đổi slice gốc
}

func main() {
    slice := []int{1, 2, 3}
    appendValuePointer(&slice, 4)
    fmt.Println(slice)  // [1 2 3 4]
}
```

### 7.2. Map - Luôn là reference

```go
package main

import "fmt"

func updateMap(m map[string]int) {
    m["key"] = 100  // Thay đổi map gốc
}

func main() {
    myMap := map[string]int{"key": 1}
    fmt.Println("Trước:", myMap)
    
    updateMap(myMap)
    
    fmt.Println("Sau:", myMap)  // Đã thay đổi
}
```

**Output:**

```
Trước: map[key:1]
Sau: map[key:100]
```

---

## 8. Nil Pointer

### 8.1. Nil Pointer là gì?

**Nil pointer** là pointer không trỏ đến địa chỉ nào.

```go
package main

import "fmt"

func main() {
    var p *int  // Nil pointer
    
    fmt.Printf("p = %v\n", p)         // <nil>
    fmt.Printf("p == nil: %v\n", p == nil)  // true
    
    // fmt.Println(*p)  // PANIC! Cannot dereference nil pointer
}
```

### 8.2. Lỗi Nil Pointer Dereference

```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    var p *Person  // Nil pointer
    
    // ❌ PANIC - nil pointer dereference
    // fmt.Println(p.Name)
    
    // ✅ Kiểm tra nil trước
    if p != nil {
        fmt.Println(p.Name)
    } else {
        fmt.Println("Pointer is nil!")
    }
}
```

### 8.3. Khởi tạo đúng cách

```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    // ❌ Sai - Nil pointer
    var p1 *Person
    // p1.Name = "John"  // PANIC!
    
    // ✅ Đúng - Cách 1: new()
    p2 := new(Person)
    p2.Name = "John"
    fmt.Println(p2.Name)  // OK
    
    // ✅ Đúng - Cách 2: &
    p3 := &Person{Name: "Jane", Age: 25}
    fmt.Println(p3.Name)  // OK
}
```

### 8.4. Safe Pointer Access

```go
package main

import "fmt"

type User struct {
    Name  string
    Email string
}

func getUserEmail(u *User) string {
    if u == nil {
        return "No email"
    }
    return u.Email
}

func main() {
    var user1 *User  // Nil
    user2 := &User{Name: "John", Email: "john@gmail.com"}
    
    fmt.Println(getUserEmail(user1))  // No email
    fmt.Println(getUserEmail(user2))  // john@gmail.com
}
```

---

## 9. Best Practices

### 9.1. Khi nào dùng Pointer?

**✅ Dùng Pointer khi:**

- Struct lớn (nhiều fields)
- Cần sửa dữ liệu gốc
- Tiết kiệm bộ nhớ
- Methods cần thay đổi struct

**❌ Không cần Pointer khi:**

- Kiểu dữ liệu nhỏ (int, bool, string ngắn)
- Chỉ đọc dữ liệu
- Slice, Map (đã là reference)

### 9.2. Luôn kiểm tra Nil

```go
// ✅ Good
func ProcessUser(u *User) error {
    if u == nil {
        return errors.New("user is nil")
    }
    // Process user
    return nil
}

// ❌ Bad
func ProcessUser(u *User) {
    // u.Name  // Có thể panic!
}
```

### 9.3. Return Pointer từ Function

```go
// ✅ Good - Escape to heap
func NewUser(name string) *User {
    return &User{Name: name}  // OK - Go tự quản lý
}

// ✅ Good
func GetUserByID(id int) (*User, error) {
    if id <= 0 {
        return nil, errors.New("invalid ID")
    }
    return &User{ID: id}, nil
}
```

### 9.4. Pointer Receiver vs Value Receiver

```go
type Counter struct {
    count int
}

// Value receiver - Không thay đổi
func (c Counter) GetCount() int {
    return c.count
}

// Pointer receiver - Thay đổi được
func (c *Counter) Increment() {
    c.count++
}
```

---

## 10. Ví dụ thực tế

### 10.1. Linked List

```go
package main

import "fmt"

type Node struct {
    Value int
    Next  *Node  // Pointer to next node
}

type LinkedList struct {
    Head *Node
}

func (ll *LinkedList) Append(value int) {
    newNode := &Node{Value: value}
    
    if ll.Head == nil {
        ll.Head = newNode
        return
    }
    
    current := ll.Head
    for current.Next != nil {
        current = current.Next
    }
    current.Next = newNode
}

func (ll *LinkedList) Print() {
    current := ll.Head
    for current != nil {
        fmt.Printf("%d -> ", current.Value)
        current = current.Next
    }
    fmt.Println("nil")
}

func main() {
    list := &LinkedList{}
    list.Append(1)
    list.Append(2)
    list.Append(3)
    
    list.Print()  // 1 -> 2 -> 3 -> nil
}
```

### 10.2. Swap Function

```go
package main

import "fmt"

func swap(a, b *int) {
    *a, *b = *b, *a
}

func main() {
    x, y := 10, 20
    fmt.Printf("Trước: x=%d, y=%d\n", x, y)
    
    swap(&x, &y)
    
    fmt.Printf("Sau: x=%d, y=%d\n", x, y)
}
```

**Output:**

```
Trước: x=10, y=20
Sau: x=20, y=10
```

### 10.3. Bank Account

```go
package main

import "fmt"

type BankAccount struct {
    Balance float64
}

func (acc *BankAccount) Deposit(amount float64) {
    acc.Balance += amount
}

func (acc *BankAccount) Withdraw(amount float64) error {
    if amount > acc.Balance {
        return fmt.Errorf("insufficient balance")
    }
    acc.Balance -= amount
    return nil
}

func main() {
    account := &BankAccount{Balance: 1000}
    
    account.Deposit(500)
    fmt.Printf("Sau deposit: %.2f\n", account.Balance)
    
    err := account.Withdraw(200)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("Sau withdraw: %.2f\n", account.Balance)
    }
}
```

**Output:**

```
Sau deposit: 1500.00
Sau withdraw: 1300.00
```

---

## 📚 Tổng kết

### Toán tử Pointer

|Toán tử|Ý nghĩa|Ví dụ|
|---|---|---|
|`&`|Lấy địa chỉ|`p := &x`|
|`*` (type)|Kiểu pointer|`var p *int`|
|`*` (value)|Dereference|`value := *p`|

### Pass by Value vs Reference

```go
// Value - Copy
func update(x int) {
    x = 10  // Không ảnh hưởng gốc
}

// Reference - Pointer
func update(px *int) {
    *px = 10  // Thay đổi gốc
}
```

### Khi nào dùng Pointer?

```
✅ Struct lớn
✅ Cần sửa dữ liệu
✅ Tiết kiệm RAM
✅ Methods

❌ Int, bool, string nhỏ
❌ Chỉ đọc
❌ Slice/Map (đã reference)
```

### Nil Check

```go
if ptr != nil {
    // Safe to use
}
```

### Best Practices

1. ✅ Luôn check nil
2. ✅ Return pointer từ constructor
3. ✅ Pointer receiver cho methods
4. ✅ Clear ownership
5. ❌ Không dùng pointer cho small types