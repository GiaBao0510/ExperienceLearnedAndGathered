# Struct trong Go - Hướng dẫn Chi tiết

## 📋 Mục lục

1. [Struct là gì?](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#1-struct-l%C3%A0-g%C3%AC)
2. [Khai báo và Khởi tạo Struct](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#2-khai-b%C3%A1o-v%C3%A0-kh%E1%BB%9Fi-t%E1%BA%A1o-struct)
3. [Struct với Pointer](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#3-struct-v%E1%BB%9Bi-pointer)
4. [Constructor Function](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#4-constructor-function)
5. [Methods và Receivers](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#5-methods-v%C3%A0-receivers)
6. [Struct Tags](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#6-struct-tags)
7. [Struct Embedding](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#7-struct-embedding)
8. [Best Practices](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#8-best-practices)
9. [Ví dụ thực tế](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#9-v%C3%AD-d%E1%BB%A5-th%E1%BB%B1c-t%E1%BA%BF)

---

## 1. Struct là gì?

**Struct** là một kiểu dữ liệu tổng hợp (composite type) trong Go, cho phép nhóm các trường (field) có kiểu dữ liệu khác nhau thành một đơn vị duy nhất.

### 🎯 So sánh với thực tế

**Giống như:**

- Một **tờ khai thông tin** có nhiều mục khác nhau
- Một **hồ sơ nhân viên** gồm: tên, tuổi, email, lương...
- Một **sản phẩm** có: ID, tên, giá, mô tả...

**Trong lập trình:**

```go
// Struct Student (Sinh viên)
type Student struct {
    ID      int
    Name    string
    Age     int
    Email   string
    GPA     float64
}
```

### 💡 Tại sao cần Struct?

**Không dùng Struct (khó quản lý):**

```go
// Thông tin sinh viên 1
name1 := "Nguyen Van A"
age1 := 20
email1 := "nva@gmail.com"
gpa1 := 3.5

// Thông tin sinh viên 2
name2 := "Tran Thi B"
age2 := 21
email2 := "ttb@gmail.com"
gpa2 := 3.8

// Rất khó quản lý khi có nhiều sinh viên!
```

**Dùng Struct (dễ quản lý):**

```go
type Student struct {
    Name  string
    Age   int
    Email string
    GPA   float64
}

// Tạo nhiều sinh viên dễ dàng
student1 := Student{"Nguyen Van A", 20, "nva@gmail.com", 3.5}
student2 := Student{"Tran Thi B", 21, "ttb@gmail.com", 3.8}
student3 := Student{"Le Van C", 19, "lvc@gmail.com", 3.2}
```

### 📊 Struct trong Go vs Class trong OOP

![Struct Illustration](https://selftuts.in/wp-content/uploads/2024/07/image-768x446.png)

|Feature|Go Struct|OOP Class|
|---|---|---|
|**Dữ liệu**|✅ Fields|✅ Properties|
|**Hành vi**|✅ Methods|✅ Methods|
|**Kế thừa**|❌ Không có|✅ Có|
|**Embedding**|✅ Có|-|
|**Interface**|✅ Implicit|✅ Explicit|

---

## 2. Khai báo và Khởi tạo Struct

### 2.1. Khai báo Struct

**Cú pháp:**

```go
type TênStruct struct {
    field1 kiểu1
    field2 kiểu2
    field3 kiểu3
}
```

**Ví dụ:**

```go
type GiangVien struct {
    Name   string
    Email  string
    Gender int  // 1: Nam, 0: Nữ, -1: Khác
}
```

### 2.2. Các cách khởi tạo Struct

#### **Cách 1: Chỉ định tên field (✅ Khuyến nghị)**

```go
gv1 := GiangVien{
    Name:   "Pham Gia Bao",
    Email:  "pgbao123@gmail.com",
    Gender: 1,
}
```

**Ưu điểm:**

- ✅ Rõ ràng, dễ đọc
- ✅ Không phụ thuộc thứ tự
- ✅ Có thể bỏ qua field (sẽ có giá trị zero value)

#### **Cách 2: Theo thứ tự (⚠️ Không khuyến nghị)**

```go
gv2 := GiangVien{
    "Nguyen Van A",
    "nva@gmail.com",
    1,
}
```

**Nhược điểm:**

- ❌ Phải đúng thứ tự
- ❌ Dễ nhầm lẫn
- ❌ Khó đọc

#### **Cách 3: Khởi tạo rỗng**

```go
// Cách 3a: Zero values
var gv3 GiangVien
// Name = "", Email = "", Gender = 0

// Cách 3b: new()
gv4 := new(GiangVien)  // Trả về pointer
// *gv4 = GiangVien{"", "", 0}
```

#### **Cách 4: Khởi tạo một số field**

```go
gv5 := GiangVien{
    Name: "Le Thi C",
    // Email và Gender sẽ có zero value
}
// gv5 = GiangVien{Name: "Le Thi C", Email: "", Gender: 0}
```

### 2.3. Truy cập và sửa field

```go
package main

import "fmt"

type GiangVien struct {
    Name   string
    Email  string
    Gender int
}

func main() {
    gv := GiangVien{
        Name:   "Pham Gia Bao",
        Email:  "pgbao123@gmail.com",
        Gender: 1,
    }

    // Đọc field
    fmt.Printf("Kiểu dữ liệu: %T\n", gv)
    fmt.Printf("Họ tên: %s\n", gv.Name)
    fmt.Printf("Email: %s\n", gv.Email)
    fmt.Printf("Giới tính: %d\n", gv.Gender)

    // Sửa field
    gv.Email = "newmail@gmail.com"
    gv.Gender = 0

    fmt.Printf("\nSau khi sửa:\n")
    fmt.Printf("Email: %s\n", gv.Email)
    fmt.Printf("Giới tính: %d\n", gv.Gender)
}
```

**Output:**

```
Kiểu dữ liệu: main.GiangVien
Họ tên: Pham Gia Bao
Email: pgbao123@gmail.com
Giới tính: 1

Sau khi sửa:
Email: newmail@gmail.com
Giới tính: 0
```

---

## 3. Struct với Pointer

### 3.1. Tại sao cần Pointer?

**Vấn đề: Go truyền struct theo giá trị (copy)**

```go
package main

import "fmt"

type Student struct {
    Name string
    Age  int
}

// Value receiver - Nhận bản copy
func UpdateAge(s Student, newAge int) {
    s.Age = newAge
    fmt.Println("Trong hàm:", s.Age)
}

func main() {
    student := Student{Name: "John", Age: 20}
    
    fmt.Println("Trước khi update:", student.Age)
    UpdateAge(student, 25)
    fmt.Println("Sau khi update:", student.Age)  // Vẫn 20!
}
```

**Output:**

```
Trước khi update: 20
Trong hàm: 25
Sau khi update: 20  ← Không thay đổi!
```

**Giải pháp: Dùng Pointer**

```go
// Pointer receiver - Nhận địa chỉ
func UpdateAgePointer(s *Student, newAge int) {
    s.Age = newAge  // Go tự dereference
    fmt.Println("Trong hàm:", s.Age)
}

func main() {
    student := Student{Name: "John", Age: 20}
    
    fmt.Println("Trước khi update:", student.Age)
    UpdateAgePointer(&student, 25)  // Truyền địa chỉ
    fmt.Println("Sau khi update:", student.Age)  // Đã thay đổi!
}
```

**Output:**

```
Trước khi update: 20
Trong hàm: 25
Sau khi update: 25  ← Đã thay đổi!
```

### 3.2. Truy cập field qua Pointer

**Go tự động dereference:**

```go
type GiangVien struct {
    Name   string
    Email  string
    Gender int
}

func HienThiThongTin(gv *GiangVien) {
    // Không cần (*gv).Name
    // Go tự hiểu gv.Name
    fmt.Printf("Họ tên: %s\n", gv.Name)
    fmt.Printf("Email: %s\n", gv.Email)
    fmt.Printf("Giới tính: %d\n", gv.Gender)
}

func main() {
    gv := GiangVien{
        Name:   "Pham Gia Bao",
        Email:  "pgbao@gmail.com",
        Gender: 1,
    }

    HienThiThongTin(&gv)  // Truyền địa chỉ
}
```

### 3.3. Khi nào dùng Value vs Pointer?

**Dùng Value khi:**

- ✅ Struct nhỏ (vài fields)
- ✅ Không cần sửa dữ liệu gốc
- ✅ Chỉ đọc dữ liệu

**Dùng Pointer khi:**

- ✅ Struct lớn (nhiều fields, chứa slice/map)
- ✅ Cần sửa dữ liệu gốc
- ✅ Tiết kiệm bộ nhớ

**Ví dụ so sánh:**

```go
// Struct nhỏ - OK dùng value
type Point struct {
    X, Y int
}

// Struct lớn - Nên dùng pointer
type Employee struct {
    ID        int
    Name      string
    Email     string
    Phone     string
    Address   string
    Skills    []string     // Slice - lớn
    Projects  []string     // Slice - lớn
    Metadata  map[string]string  // Map - lớn
}
```

---

## 4. Constructor Function

### 4.1. Constructor là gì?

**Constructor Function** là hàm tạo đối tượng struct với giá trị khởi tạo và validation.

### 4.2. Tại sao cần Constructor?

Constructor Function giúp:

**1. Validation (Kiểm tra dữ liệu đầu vào)**

```go
func NewGiangVien(name, email string, gender int) (*GiangVien, error) {
    // Kiểm tra name không rỗng
    if name == "" {
        return nil, errors.New("name cannot be empty")
    }
    
    // Kiểm tra email hợp lệ
    if !strings.Contains(email, "@") {
        return nil, errors.New("invalid email")
    }
    
    // Kiểm tra gender trong range 0-1
    if gender < 0 || gender > 1 {
        return nil, errors.New("gender must be 0 or 1")
    }
    
    return &GiangVien{Name: name, Email: email, Gender: gender}, nil
}
```

**2. Default Values (Thiết lập giá trị mặc định)**

```go
func NewProduct(name string, price float64) *Product {
    return &Product{
        ID:        generateUniqueID(),  // Tự sinh ID
        Name:      name,
        Price:     price,
        CreatedAt: time.Now(),          // Tự động timestamp
        Status:    "active",            // Giá trị mặc định
        Views:     0,                   // Khởi tạo counter
    }
}
```

**3. Complex Initialization (Khởi tạo phức tạp)**

```go
func NewDatabase(host, port string) (*Database, error) {
    db := &Database{
        Host: host,
        Port: port,
    }
    
    // Kết nối database
    if err := db.Connect(); err != nil {
        return nil, fmt.Errorf("failed to connect: %w", err)
    }
    
    // Khởi tạo connection pool
    db.Pool = db.createConnectionPool()
    
    // Setup migrations
    if err := db.runMigrations(); err != nil {
        return nil, err
    }
    
    return db, nil
}
```

**4. Encapsulation (Đóng gói logic)**

```go
func NewEmployee(name string, salary float64) *Employee {
    emp := &Employee{
        ID:     generateEmployeeID(),
        Name:   name,
        Salary: salary,
    }
    
    // Logic tính toán phức tạp
    emp.calculateBenefits()
    emp.assignDefaultPermissions()
    emp.sendWelcomeEmail()
    
    return emp
}
```

**5. Consistency (Đảm bảo tính nhất quán)**

```go
// Không có constructor - dễ sai
user1 := User{Name: "John", Email: "invalid"}  // Email sai!
user2 := User{Name: "", Email: "valid@gmail.com"}  // Name rỗng!

// Có constructor - luôn đúng
user1, err := NewUser("John", "john@gmail.com")  // Kiểm tra trước khi tạo
if err != nil {
    // Xử lý lỗi
}
```

### 4.3. Ví dụ đầy đủ

```go
package main

import (
    "errors"
    "fmt"
    "strings"
)

type GiangVien struct {
    Name   string
    Email  string
    Gender int
}

// Constructor với validation
func NewGiangVien(name, email string, gender int) (*GiangVien, error) {
    // Validation
    if name == "" {
        return nil, errors.New("name cannot be empty")
    }
    
    if !strings.Contains(email, "@") {
        return nil, errors.New("invalid email")
    }
    
    if gender < 0 || gender > 1 {
        return nil, errors.New("gender must be 0 or 1")
    }
    
    // Tạo và trả về pointer
    return &GiangVien{
        Name:   name,
        Email:  email,
        Gender: gender,
    }, nil
}

func main() {
    // ✅ Tạo thành công
    gv1, err := NewGiangVien("Pham Gia Bao", "pgbao@gmail.com", 1)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Printf("✅ Created: %+v\n", gv1)
    
    // ❌ Tạo thất bại - Email không hợp lệ
    gv2, err := NewGiangVien("John", "invalid-email", 1)
    if err != nil {
        fmt.Println("❌ Error:", err)
    }
    
    // ❌ Tạo thất bại - Gender không hợp lệ
    gv3, err := NewGiangVien("Jane", "jane@gmail.com", 5)
    if err != nil {
        fmt.Println("❌ Error:", err)
    }
}
```

**Output:**

```
✅ Created: &{Name:Pham Gia Bao Email:pgbao@gmail.com Gender:1}
❌ Error: invalid email
❌ Error: gender must be 0 or 1
```

---

## 5. Methods và Receivers

### 5.1. Method là gì?

**Method** là hàm gắn với struct, cho phép struct có "hành vi".

**Cú pháp:**

```go
func (receiver TênStruct) TênMethod() kiểu_trả_về {
    // Code
}
```

![Method Anatomy](https://www.practical-go-lessons.com/img/method_anatomy.00332211.png)

### 5.2. Value Receiver vs Pointer Receiver

#### **Value Receiver (Nhận bản copy)**

```go
package main

import "fmt"

type GiangVien struct {
    Name   string
    Email  string
    Gender int
}

// Value Receiver - Không sửa gốc
func (gv GiangVien) HienThiThongTin() {
    fmt.Printf("Họ tên: %s\n", gv.Name)
    fmt.Printf("Email: %s\n", gv.Email)
    fmt.Printf("Giới tính: %d\n", gv.Gender)
}

// Value Receiver - Thử sửa (không thành công)
func (gv GiangVien) UpdateEmailWrong(newEmail string) {
    gv.Email = newEmail  // Chỉ sửa bản copy!
    fmt.Println("Trong method:", gv.Email)
}

func main() {
    gv := GiangVien{
        Name:   "Pham Gia Bao",
        Email:  "old@gmail.com",
        Gender: 1,
    }

    gv.HienThiThongTin()
    
    fmt.Println("\nTrước update:", gv.Email)
    gv.UpdateEmailWrong("new@gmail.com")
    fmt.Println("Sau update:", gv.Email)  // Vẫn old@gmail.com
}
```

**Output:**

```
Họ tên: Pham Gia Bao
Email: old@gmail.com
Giới tính: 1

Trước update: old@gmail.com
Trong method: new@gmail.com
Sau update: old@gmail.com  ← Không thay đổi!
```

#### **Pointer Receiver (Nhận địa chỉ)**

```go
// Pointer Receiver - Sửa được gốc
func (gv *GiangVien) UpdateEmail(newEmail string) {
    gv.Email = newEmail  // Sửa struct gốc
}

// Pointer Receiver - Clear data
func (gv *GiangVien) Clear() {
    gv.Name = ""
    gv.Email = ""
    gv.Gender = -1
}

func main() {
    gv := GiangVien{
        Name:   "Pham Gia Bao",
        Email:  "old@gmail.com",
        Gender: 1,
    }

    fmt.Println("Trước update:", gv.Email)
    gv.UpdateEmail("new@gmail.com")  // Go tự động lấy địa chỉ
    fmt.Println("Sau update:", gv.Email)  // Đã thay đổi!
    
    gv.Clear()
    gv.HienThiThongTin()  // Tất cả fields đã bị clear
}
```

**Output:**

```
Trước update: old@gmail.com
Sau update: new@gmail.com  ← Đã thay đổi!
Họ tên: 
Email: 
Giới tính: -1
```

### 5.3. Khi nào dùng loại nào?

**Value Receiver:**

```go
// ✅ Dùng khi chỉ đọc dữ liệu
func (s Student) GetName() string {
    return s.Name
}

// ✅ Struct nhỏ
func (p Point) Distance() float64 {
    return math.Sqrt(p.X*p.X + p.Y*p.Y)
}
```

**Pointer Receiver:**

```go
// ✅ Cần sửa dữ liệu
func (s *Student) SetGPA(gpa float64) {
    s.GPA = gpa
}

// ✅ Struct lớn (tiết kiệm bộ nhớ)
func (e *Employee) Promote() {
    e.Level++
    e.Salary *= 1.1
}
```

---

## 6. Struct Tags

### 6.1. Tag là gì?

**Tags** là metadata gắn vào field, giúp thư viện bên ngoài xử lý struct.

**Cú pháp:**

```go
type StructName struct {
    FieldName type `key:"value"`
}
```

![Struct Tags](https://media2.dev.to/dynamic/image/width=1000,height=420,fit=cover,gravity=auto,format=auto/https%3A%2F%2Fgithub.com%2Fkodelint%2Fblog-assets%2Fraw%2Fmain%2Fimages%2F01-Use-Struct-Tags-in-Golang.jpeg)

![Struct Tags Syntax](https://raw.githubusercontent.com/fogio-org/vscode-go-struct-tags-syntax-highlight/master/assets/img/preview-3.png)

### 6.2. JSON Tags

```go
package main

import (
    "encoding/json"
    "fmt"
)

type GiangVien struct {
    Name   string `json:"Ho ten"`
    Email  string `json:"Dia chi email"`
    Gender int    `json:"Gioi tinh"`
}

func main() {
    gv := GiangVien{
        Name:   "Pham Gia Bao",
        Email:  "bao123@gmail.com",
        Gender: 1,
    }

    // Struct → JSON
    jsonData, err := json.Marshal(gv)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println(string(jsonData))
}
```

**Output:**

```json
{"Ho ten":"Pham Gia Bao","Dia chi email":"bao123@gmail.com","Gioi tinh":1}
```

---

## 7. Struct Embedding

### 7.1. Embedding là gì?

**Embedding** (nhúng) cho phép một struct chứa struct khác, "kế thừa" fields và methods.

```go
package main

import "fmt"

// Base struct
type Person struct {
    Name string
    Age  int
}

// Methods của Person
func (p Person) Introduce() {
    fmt.Printf("Xin chào, tôi là %s, %d tuổi\n", p.Name, p.Age)
}

// Student nhúng Person
type Student struct {
    Person         // Embedding
    StudentID string
    GPA       float64
}

func main() {
    student := Student{
        Person: Person{
            Name: "Nguyen Van A",
            Age:  20,
        },
        StudentID: "SV001",
        GPA:       3.5,
    }

    // Truy cập field của Person
    fmt.Println("Tên:", student.Name)  // Không cần student.Person.Name
    
    // Gọi method của Person
    student.Introduce()
}
```

---

## 📚 Tổng kết

### Struct Basics

```go
type Student struct {
    Name string
    Age  int
}

s := Student{Name: "John", Age: 20}
```

### Pointer vs Value

```go
// Value - Không sửa gốc
func (s Student) Method1() { }

// Pointer - Sửa được gốc
func (s *Student) Method2() { }
```

### Constructor - TẠI SAO CẦN?

```go
// ✅ Validation
// ✅ Default values
// ✅ Complex initialization
// ✅ Encapsulation
// ✅ Consistency

func NewStudent(name string) (*Student, error) {
    if name == "" {
        return nil, errors.New("invalid name")
    }
    return &Student{Name: name}, nil
}
```

### Tags

```go
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name,omitempty"`
}
```