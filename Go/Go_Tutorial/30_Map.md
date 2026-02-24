# Map trong Go - Hướng dẫn Chi tiết

## 📋 Mục lục

1. [Map là gì?](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#1-map-l%C3%A0-g%C3%AC)
2. [Khai báo và khởi tạo Map](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#2-khai-b%C3%A1o-v%C3%A0-kh%E1%BB%9Fi-t%E1%BA%A1o-map)
3. [Các thao tác cơ bản](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#3-c%C3%A1c-thao-t%C3%A1c-c%C6%A1-b%E1%BA%A3n)
4. [Duyệt Map](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#4-duy%E1%BB%87t-map)
5. [Map kết hợp với Struct](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#5-map-k%E1%BA%BFt-h%E1%BB%A3p-v%E1%BB%9Bi-struct)
6. [Map kết hợp với Slice](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#6-map-k%E1%BA%BFt-h%E1%BB%A3p-v%E1%BB%9Bi-slice)
7. [Map lồng nhau](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#7-map-l%E1%BB%93ng-nhau)
8. [So sánh Map với Slice và Array](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#8-so-s%C3%A1nh-map-v%E1%BB%9Bi-slice-v%C3%A0-array)
9. [Các lưu ý quan trọng](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#9-c%C3%A1c-l%C6%B0u-%C3%BD-quan-tr%E1%BB%8Dng)

---

## 1. Map là gì?

**Map** là một cấu trúc dữ liệu lưu trữ các cặp **Key - Value** (Khóa - Giá trị). Mỗi key là duy nhất và được dùng để truy xuất value tương ứng.

> 💡 Map trong Go tương tự như:
> 
> - **Dictionary** trong Python và C#
> - **HashMap** trong Java
> - **Object/Map** trong JavaScript

### 📊 Minh họa cấu trúc Map

```
map[string]int

  KEY (string)       VALUE (int)
┌──────────────┬──────────────────┐
│  "coffee"    │      12000       │
├──────────────┼──────────────────┤
│  "tea"       │      15000       │
├──────────────┼──────────────────┤
│ "apple juice"│      20000       │
└──────────────┴──────────────────┘
```

### 🔑 Đặc điểm chính của Map

| Đặc điểm            | Mô tả                                                    |
| ------------------- | -------------------------------------------------------- |
| **Key duy nhất**    | Không có 2 key giống nhau trong cùng 1 map               |
| **Không có thứ tự** | Thứ tự khi duyệt không cố định                           |
| **Kiểu tham chiếu** | Map là reference type (như Slice)                        |
| **Tìm kiếm nhanh**  | Độ phức tạp O(1) trung bình                              |
| **Key hợp lệ**      | Key phải là kiểu **comparable** (int, string, bool, ...) |

### ✅ Kiểu dữ liệu hợp lệ làm Key

```go
map[string]int       // ✅ Key là string
map[int]string       // ✅ Key là int
map[bool]float64     // ✅ Key là bool
map[float64]string   // ✅ Key là float64

// ❌ KHÔNG hợp lệ làm Key:
// map[[]int]string  // Slice không thể làm key
// map[map[string]int]string // Map không thể làm key
```

---

## 2. Khai báo và khởi tạo Map

Có **3 cách** để khai báo và khởi tạo Map trong Go:

### Cách 1: Khai báo và khởi tạo trực tiếp (Phổ biến nhất)

```go
func main() {
    // Cú pháp: map[KeyType]ValueType{ key: value, ... }
    drink := map[string]int{
        "coffee":      12000,
        "tea":         15000,
        "apple juice": 20000,
    }

    fmt.Println(drink)
    // Output: map[apple juice:20000 coffee:12000 tea:15000]
}
```

> 💡 **Lưu ý:** Dấu phẩy `,` sau phần tử cuối cùng là **bắt buộc** trong Go khi viết nhiều dòng.

### Cách 2: Dùng hàm `make()` (Khi chưa biết dữ liệu trước)

```go
func main() {
    // Cú pháp: make(map[KeyType]ValueType)
    menu := make(map[string]float64)

    // Gán giá trị sau
    menu["Coffee"]  = 12.000
    menu["Tea"]     = 10.000
    menu["Cơm tấm"] = 25.000

    fmt.Println(menu)
    // Output: map[Coffee:12 Cơm tấm:25 Tea:10]
}
```

### Cách 3: Khai báo rỗng rồi cấp phát sau (Ít dùng)

```go
func main() {
    // Bước 1: Khai báo biến map (chưa dùng được)
    var menu map[string]float64 // Lúc này menu = nil

    // Bước 2: Cấp phát bộ nhớ
    menu = make(map[string]float64)

    // Bước 3: Gán giá trị
    menu["Coffee"]  = 12.000
    menu["Tea"]     = 10.000
    menu["Cơm tấm"] = 25.000

    fmt.Println(menu)
}
```

> ⚠️ **Cảnh báo:** Nếu quên `make()` ở Cách 3 và gán giá trị vào map `nil`, chương trình sẽ **panic** (crash):

```go
var menu map[string]float64
menu["Coffee"] = 12.000 // ❌ panic: assignment to entry in nil map
```

### 📝 Khi nào dùng cách nào?

|Tình huống|Nên dùng|
|---|---|
|Biết trước dữ liệu|Cách 1 (khai báo trực tiếp)|
|Đọc dữ liệu từ input hoặc file|Cách 2 (`make()`)|
|Khai báo ở global scope|Cách 2 hoặc 3|

---

## 3. Các thao tác cơ bản

### 3.1. Thêm và cập nhật phần tử

```go
func main() {
    scores := make(map[string]int)

    // Thêm phần tử mới
    scores["Alice"] = 95
    scores["Bob"]   = 87
    scores["Carol"] = 91

    fmt.Println(scores) // map[Alice:95 Bob:87 Carol:91]

    // Cập nhật giá trị (ghi đè nếu key đã tồn tại)
    scores["Bob"] = 90 // Bob từ 87 → 90

    fmt.Println(scores) // map[Alice:95 Bob:90 Carol:91]
}
```

> 💡 Cú pháp **thêm** và **cập nhật** đều giống nhau: `map[key] = value`
> 
> - Nếu key **chưa có** → Thêm mới
> - Nếu key **đã có** → Ghi đè (cập nhật)

### 3.2. Đọc giá trị

```go
func main() {
    scores := map[string]int{
        "Alice": 95,
        "Bob":   87,
    }

    // Đọc thông thường
    fmt.Println(scores["Alice"]) // 95

    // Đọc key không tồn tại → Trả về zero value (0 với int)
    fmt.Println(scores["David"]) // 0 (không báo lỗi!)
}
```

### 3.3. Kiểm tra Key có tồn tại không (Quan trọng!)

Vì khi đọc key không tồn tại trả về **zero value** (0, "", false...) chứ không báo lỗi, nên cần kiểm tra:

```go
func main() {
    menu := map[string]float64{
        "Coffee":  12.000,
        "Tea":     10.000,
        "Cơm tấm": 25.000,
    }

    // Cú pháp: value, ok := map[key]
    // ok = true  nếu key tồn tại
    // ok = false nếu key không tồn tại
    
    // Trường hợp 1: Key tồn tại
    price, exists := menu["Coffee"]
    if exists {
        fmt.Printf("Giá Coffee: %.3f\n", price) // Giá Coffee: 12.000
    }

    // Trường hợp 2: Key không tồn tại
    price, exists = menu["Phở bò"]
    if exists {
        fmt.Printf("Giá Phở bò: %.3f\n", price)
    } else {
        fmt.Println("❌ Không tìm thấy món này trong menu") // Dòng này được in
    }

    // Trường hợp 3: Chỉ cần kiểm tra tồn tại, không cần value
    if _, exists := menu["Tea"]; exists {
        fmt.Println("✅ Có món Tea trong menu")
    }
}
```

**Output:**

```
Giá Coffee: 12.000
❌ Không tìm thấy món này trong menu
✅ Có món Tea trong menu
```

### 3.4. Xóa phần tử

```go
func main() {
    scores := map[string]int{
        "Alice": 95,
        "Bob":   87,
        "Carol": 91,
    }

    fmt.Println("Trước khi xóa:", scores)

    // Xóa phần tử theo key
    delete(scores, "Bob")

    fmt.Println("Sau khi xóa Bob:", scores)

    // Xóa key không tồn tại → Không có lỗi
    delete(scores, "David") // Hoàn toàn OK
    fmt.Println("Sau khi xóa David (không tồn tại):", scores)
}
```

**Output:**

```
Trước khi xóa: map[Alice:95 Bob:87 Carol:91]
Sau khi xóa Bob: map[Alice:95 Carol:91]
Sau khi xóa David (không tồn tại): map[Alice:95 Carol:91]
```

### 3.5. Đếm số phần tử

```go
func main() {
    menu := map[string]float64{
        "Coffee":  12.000,
        "Tea":     10.000,
        "Cơm tấm": 25.000,
    }

    fmt.Println("Số món trong menu:", len(menu)) // 3
}
```

### 📋 Tổng hợp các thao tác cơ bản

```go
m := make(map[string]int)

m["key"] = 100        // Thêm / Cập nhật
value := m["key"]     // Đọc giá trị
v, ok := m["key"]     // Đọc + kiểm tra tồn tại
delete(m, "key")      // Xóa phần tử
length := len(m)      // Đếm số phần tử
```

---

## 4. Duyệt Map

### 4.1. Duyệt key và value

```go
func main() {
    menu := map[string]float64{
        "Coffee":  12.000,
        "Tea":     10.000,
        "Cơm tấm": 25.000,
    }

    for key, value := range menu {
        fmt.Printf("%-15s: %.3f VND\n", key, value)
    }
}
```

**Output (thứ tự có thể khác mỗi lần chạy):**

```
Coffee         : 12.000 VND
Tea            : 10.000 VND
Cơm tấm        : 25.000 VND
```

> ⚠️ **Lưu ý quan trọng:** Khi duyệt map, **thứ tự không cố định**. Mỗi lần chạy có thể ra thứ tự khác nhau. Đây là thiết kế có chủ đích của Go nhằm tối ưu hiệu năng.

### 4.2. Duyệt chỉ key

```go
for key := range menu {
    fmt.Println("Món:", key)
}
```

### 4.3. Duyệt chỉ value

```go
for _, value := range menu {
    fmt.Printf("Giá: %.3f\n", value)
}
```

### 4.4. Duyệt theo thứ tự (Cần sort)

Nếu cần duyệt theo thứ tự nhất định, hãy lấy keys ra, sắp xếp rồi duyệt:

```go
import (
    "fmt"
    "sort"
)

func main() {
    menu := map[string]float64{
        "Coffee":  12.000,
        "Tea":     10.000,
        "Cơm tấm": 25.000,
        "Bánh mì": 8.000,
    }

    // Lấy tất cả keys
    keys := make([]string, 0, len(menu))
    for k := range menu {
        keys = append(keys, k)
    }

    // Sắp xếp keys
    sort.Strings(keys)

    // Duyệt theo thứ tự đã sắp xếp
    fmt.Println("=== Menu (theo thứ tự A-Z) ===")
    for _, k := range keys {
        fmt.Printf("%-15s: %.3f VND\n", k, menu[k])
    }
}
```

**Output (luôn cố định):**

```
=== Menu (theo thứ tự A-Z) ===
Bánh mì        : 8.000 VND
Coffee         : 12.000 VND
Cơm tấm        : 25.000 VND
Tea            : 10.000 VND
```

---

## 5. Map kết hợp với Struct

Kết hợp Map và Struct giúp lưu trữ thông tin phức tạp hơn, trong đó **Struct luôn là Value** của Map.

```
map[KeyType]StructType
      Key → { field1, field2, field3, ... }
```

### Ví dụ cơ bản:

```go
type Employee struct {
    Id   int
    Name string
    Role string
}

func main() {
    employees := map[string]Employee{
        "emp1": {Id: 1, Name: "John Doe",      Role: "Software Engineer"},
        "emp2": {Id: 2, Name: "Jane Smith",     Role: "Product Manager"},
        "emp3": {Id: 3, Name: "Alice Johnson",  Role: "UX Designer"},
    }

    // In tất cả thông tin
    fmt.Println("=== Danh sách nhân viên ===")
    for key, emp := range employees {
        fmt.Printf("[%s] ID: %d | Tên: %-15s | Vị trí: %s\n",
            key, emp.Id, emp.Name, emp.Role)
    }

    // Truy cập một trường cụ thể
    fmt.Println("\nVị trí của emp1:", employees["emp1"].Role)

    // Kiểm tra nhân viên có tồn tại không
    if emp, exists := employees["emp5"]; exists {
        fmt.Println("Tìm thấy:", emp.Name)
    } else {
        fmt.Println("\n❌ Không tìm thấy emp5")
    }
}
```

**Output:**

```
=== Danh sách nhân viên ===
[emp1] ID: 1 | Tên: John Doe        | Vị trí: Software Engineer
[emp2] ID: 2 | Tên: Jane Smith      | Vị trí: Product Manager
[emp3] ID: 3 | Tên: Alice Johnson   | Vị trí: UX Designer

Vị trí của emp1: Software Engineer

❌ Không tìm thấy emp5
```

### ⚠️ Cập nhật trường trong Struct value

Đây là điểm **dễ gây nhầm lẫn** — không thể sửa trực tiếp trường của struct trong map:

```go
type Employee struct {
    Id   int
    Name string
    Role string
}

func main() {
    employees := map[string]Employee{
        "emp1": {Id: 1, Name: "John Doe", Role: "Junior Dev"},
    }

    // ❌ KHÔNG làm được - Compile error!
    // employees["emp1"].Role = "Senior Dev"

    // ✅ CÁCH ĐÚNG: Lấy ra, sửa, gán lại
    emp := employees["emp1"]  // Lấy ra (copy)
    emp.Role = "Senior Dev"   // Sửa bản copy
    employees["emp1"] = emp   // Gán lại vào map

    fmt.Println(employees["emp1"].Role) // Senior Dev
    
    // ✅ CÁCH KHÁC: Dùng pointer trong map
    // map[string]*Employee sẽ cho phép sửa trực tiếp
}
```

### Dùng Pointer Struct để sửa trực tiếp:

```go
type Employee struct {
    Id   int
    Name string
    Role string
}

func main() {
    // Dùng *Employee (pointer) thay vì Employee
    employees := map[string]*Employee{
        "emp1": {Id: 1, Name: "John Doe", Role: "Junior Dev"},
        "emp2": {Id: 2, Name: "Jane Smith", Role: "Product Manager"},
    }

    // ✅ Sửa trực tiếp được vì dùng pointer
    employees["emp1"].Role = "Senior Dev"

    fmt.Println(employees["emp1"].Role) // Senior Dev
}
```

---

## 6. Map kết hợp với Slice

Mỗi Key sẽ trỏ đến một **danh sách giá trị** (Slice), hữu ích khi cần nhóm dữ liệu.

```
map[KeyType][]ValueType
      Key → [value1, value2, value3, ...]
```

### Ví dụ cơ bản:

```go
func main() {
    menus := map[string][]string{
        "menu1": {"Cơm xào bò", "Mì xào hải sản", "Lẩu thái"},
        "menu2": {"Bún bò Huế", "Phở bò", "Bánh cuốn"},
        "menu3": {"Bánh mì", "Bánh xèo"},
    }

    // In tất cả
    for key, dishes := range menus {
        fmt.Printf("%s: %v\n", key, dishes)
    }

    // Truy cập một món cụ thể
    fmt.Println("\nMón đầu tiên trong menu1:", menus["menu1"][0])

    // Duyệt món trong một menu cụ thể
    fmt.Println("\nCác món trong menu2:")
    for i, dish := range menus["menu2"] {
        fmt.Printf("  %d. %s\n", i+1, dish)
    }
}
```

**Output:**

```
menu1: [Cơm xào bò Mì xào hải sản Lẩu thái]
menu2: [Bún bò Huế Phở bò Bánh cuốn]
menu3: [Bánh mì Bánh xèo]

Món đầu tiên trong menu1: Cơm xào bò

Các món trong menu2:
  1. Bún bò Huế
  2. Phở bò
  3. Bánh cuốn
```

### Thêm phần tử vào Slice trong Map:

```go
func main() {
    // Map lưu danh sách học sinh theo lớp
    classes := make(map[string][]string)

    // Thêm học sinh vào lớp
    classes["10A"] = append(classes["10A"], "Nguyễn Văn An")
    classes["10A"] = append(classes["10A"], "Trần Thị Bình")
    classes["10B"] = append(classes["10B"], "Lê Văn Cường")
    classes["10A"] = append(classes["10A"], "Phạm Thị Dung")

    // In danh sách theo lớp
    for class, students := range classes {
        fmt.Printf("\nLớp %s (%d học sinh):\n", class, len(students))
        for i, s := range students {
            fmt.Printf("  %d. %s\n", i+1, s)
        }
    }
}
```

**Output:**

```
Lớp 10A (3 học sinh):
  1. Nguyễn Văn An
  2. Trần Thị Bình
  3. Phạm Thị Dung

Lớp 10B (1 học sinh):
  1. Lê Văn Cường
```

---

## 7. Map lồng nhau

Map có thể chứa Map khác làm Value, tạo thành **Map lồng nhau** (Nested Map):

```
map[KeyType]map[KeyType]ValueType
```

### Ví dụ ứng dụng:

```go
func main() {
    // Cấu trúc: Trường → Khối → Danh sách lớp
    school := map[string]map[string][]string{
        "THCS": {
            "Khối 6": {"6A1", "6A2", "6A3"},
            "Khối 7": {"7A1", "7A2"},
            "Khối 8": {"8A1", "8A2", "8A3", "8A4"},
        },
        "THPT": {
            "Khối 10": {"10A1", "10A2", "10A3"},
            "Khối 11": {"11A1", "11A2"},
            "Khối 12": {"12A1", "12A2", "12A3"},
        },
    }

    // Duyệt toàn bộ
    for level, grades := range school {
        fmt.Printf("\n📚 Bậc: %s\n", level)
        for grade, classes := range grades {
            fmt.Printf("  📖 %s: %v\n", grade, classes)
        }
    }

    // Truy cập lớp cụ thể
    fmt.Println("\nCác lớp khối 10 THPT:", school["THPT"]["Khối 10"])

    // Thêm khối mới vào THPT
    school["THPT"]["Khối 9"] = []string{"9A1", "9A2"}
}
```

**Output:**

```
📚 Bậc: THCS
  📖 Khối 6: [6A1 6A2 6A3]
  📖 Khối 7: [7A1 7A2]
  📖 Khối 8: [8A1 8A2 8A3 8A4]

📚 Bậc: THPT
  📖 Khối 10: [10A1 10A2 10A3]
  📖 Khối 11: [11A1 11A2]
  📖 Khối 12: [12A1 12A2 12A3]

Các lớp khối 10 THPT: [10A1 10A2 10A3]
```

### ⚠️ Khởi tạo Map lồng nhau cẩn thận:

```go
func main() {
    // ❌ SAI - Map bên trong chưa được khởi tạo
    outer := make(map[string]map[string]int)
    // outer["a"]["b"] = 1  // panic: assignment to entry in nil map

    // ✅ ĐÚNG - Phải khởi tạo map bên trong trước
    outer = make(map[string]map[string]int)
    outer["a"] = make(map[string]int) // Khởi tạo map bên trong
    outer["a"]["b"] = 1
    fmt.Println(outer) // map[a:map[b:1]]
}
```

---

## 8. So sánh Map với Slice và Array

|Tiêu chí|Array|Slice|Map|
|---|---|---|---|
|**Truy cập**|Index số (0, 1, 2...)|Index số (0, 1, 2...)|Key tùy chỉnh|
|**Kích thước**|Cố định|Động|Động|
|**Thứ tự**|Có|Có|Không cố định|
|**Tìm kiếm**|O(n)|O(n)|O(1) trung bình|
|**Key**|Chỉ là số nguyên|Chỉ là số nguyên|Tùy chỉnh|
|**Use Case**|Dữ liệu cố định|Danh sách|Lookup table, nhóm dữ liệu|

### Khi nào dùng Map?

```
Dùng Map khi:
✅ Cần truy xuất dữ liệu theo "tên" thay vì số thứ tự
✅ Cần kiểm tra sự tồn tại của một phần tử nhanh
✅ Cần nhóm dữ liệu theo một tiêu chí
✅ Dữ liệu dạng cặp "khóa - giá trị"
✅ Cần đếm tần suất xuất hiện

Không dùng Map khi:
❌ Cần duyệt theo thứ tự nhất định
❌ Dữ liệu là danh sách tuần tự đơn giản
```

### Ví dụ thực tế - Đếm tần suất từ:

```go
func main() {
    words := []string{"go", "is", "fast", "go", "is", "great", "go"}

    // Dùng Map để đếm số lần xuất hiện
    count := make(map[string]int)
    for _, word := range words {
        count[word]++ // Tự động tăng, kể cả key chưa tồn tại (zero value = 0)
    }

    fmt.Println("Tần suất xuất hiện:")
    for word, freq := range count {
        fmt.Printf("  %-10s: %d lần\n", word, freq)
    }
}
```

**Output:**

```
Tần suất xuất hiện:
  go        : 3 lần
  is        : 2 lần
  fast      : 1 lần
  great     : 1 lần
```

---

## 9. Các lưu ý quan trọng

### ✅ Best Practices

**1. Luôn kiểm tra key tồn tại trước khi dùng:**

```go
value, ok := m[key]
if !ok {
    // xử lý khi không tìm thấy
}
```

**2. Dùng `make()` khi tạo map rỗng:**

```go
// ✅ ĐÚNG
m := make(map[string]int)

// ❌ SAI (map nil, không thể gán)
var m map[string]int
```

**3. Map là reference type — cẩn thận khi copy:**

```go
original := map[string]int{"a": 1, "b": 2}

// Đây KHÔNG phải bản sao - cùng trỏ đến một map
copied := original
copied["a"] = 999

fmt.Println(original["a"]) // 999 - original cũng bị thay đổi!

// ✅ Tạo bản sao thật sự
trueCopy := make(map[string]int)
for k, v := range original {
    trueCopy[k] = v
}
```

**4. Không thể so sánh 2 map bằng `==`:**

```go
m1 := map[string]int{"a": 1}
m2 := map[string]int{"a": 1}
// fmt.Println(m1 == m2) // ❌ Compile error

// ✅ Dùng reflect hoặc so sánh thủ công
import "reflect"
fmt.Println(reflect.DeepEqual(m1, m2)) // true
```

**5. Khai thác zero value để viết code ngắn gọn:**

```go
// ✅ Đếm không cần kiểm tra tồn tại trước
counter := make(map[string]int)
counter["key"]++ // OK vì zero value của int là 0

// ✅ Tạo danh sách nhóm không cần kiểm tra tồn tại
groups := make(map[string][]string)
groups["A"] = append(groups["A"], "item") // OK vì zero value của slice là nil
```

### ⚠️ Các lỗi thường gặp

```go
// Lỗi 1: Gán vào nil map
var m map[string]int
m["key"] = 1 // ❌ panic: assignment to entry in nil map

// Lỗi 2: Không kiểm tra tồn tại dẫn đến dùng zero value nhầm
config := map[string]int{"timeout": 30}
timeout := config["TIMEOUT"] // timeout = 0 (không phải 30!)
// Phải viết: "timeout" không phải "TIMEOUT"

// Lỗi 3: Sửa trực tiếp trường của Struct value trong Map
type Point struct{ X, Y int }
points := map[string]Point{"A": {1, 2}}
// points["A"].X = 10 // ❌ Compile error
p := points["A"]     // ✅ Lấy ra
p.X = 10             // ✅ Sửa
points["A"] = p      // ✅ Gán lại
```

---

## 📚 Tài liệu tham khảo

- [Go Tour - Maps](https://go.dev/tour/moretypes/19)
- [Go by Example - Maps](https://gobyexample.com/maps)
- [Effective Go - Maps](https://go.dev/doc/effective_go#maps)
- [Go Documentation - builtin](https://pkg.go.dev/builtin#make)