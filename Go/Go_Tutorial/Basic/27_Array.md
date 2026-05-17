# Array trong Golang

_Hướng dẫn toàn diện về mảng và cách sử dụng hiệu quả_

---

## Mục lục

1. [Giới thiệu về Array](https://claude.ai/chat/60e30e6a-7f00-4c17-af2e-f7923264b7b8#1-gi%E1%BB%9Bi-thi%E1%BB%87u-v%E1%BB%81-array)
2. [Khái niệm cơ bản](https://claude.ai/chat/60e30e6a-7f00-4c17-af2e-f7923264b7b8#2-kh%C3%A1i-ni%E1%BB%87m-c%C6%A1-b%E1%BA%A3n)
3. [Cách khai báo và khởi tạo Array](https://claude.ai/chat/60e30e6a-7f00-4c17-af2e-f7923264b7b8#3-c%C3%A1ch-khai-b%C3%A1o-v%C3%A0-kh%E1%BB%9Fi-t%E1%BA%A1o-array)
4. [Truy cập và thay đổi phần tử](https://claude.ai/chat/60e30e6a-7f00-4c17-af2e-f7923264b7b8#4-truy-c%E1%BA%ADp-v%C3%A0-thay-%C4%91%E1%BB%95i-ph%E1%BA%A7n-t%E1%BB%AD)
5. [Array đa chiều](https://claude.ai/chat/60e30e6a-7f00-4c17-af2e-f7923264b7b8#5-array-%C4%91a-chi%E1%BB%81u)
6. [Duyệt Array với for-range](https://claude.ai/chat/60e30e6a-7f00-4c17-af2e-f7923264b7b8#6-duy%E1%BB%87t-array-v%E1%BB%9Bi-for-range)
7. [Array vs Slice - Sự khác biệt quan trọng](https://claude.ai/chat/60e30e6a-7f00-4c17-af2e-f7923264b7b8#7-array-vs-slice---s%E1%BB%B1-kh%C3%A1c-bi%E1%BB%87t-quan-tr%E1%BB%8Dng)
8. [Kết hợp Array với Struct](https://claude.ai/chat/60e30e6a-7f00-4c17-af2e-f7923264b7b8#8-k%E1%BA%BFt-h%E1%BB%A3p-array-v%E1%BB%9Bi-struct)
9. [Các thao tác phổ biến với Array](https://claude.ai/chat/60e30e6a-7f00-4c17-af2e-f7923264b7b8#9-c%C3%A1c-thao-t%C3%A1c-ph%E1%BB%95-bi%E1%BA%BFn-v%E1%BB%9Bi-array)
10. [Lưu ý quan trọng khi làm việc với Array](https://claude.ai/chat/60e30e6a-7f00-4c17-af2e-f7923264b7b8#10-l%C6%B0u-%C3%BD-quan-tr%E1%BB%8Dng-khi-l%C3%A0m-vi%E1%BB%87c-v%E1%BB%9Bi-array)

---

## 1. Giới thiệu về Array

**Array (Mảng)** là một cấu trúc dữ liệu cơ bản trong lập trình, cho phép lưu trữ một tập hợp các phần tử có cùng kiểu dữ liệu. Trong Golang, `Array` có những đặc điểm riêng biệt mà bạn cần nắm vững.

![](https://zalopay-oss.github.io/go-advanced/images/ch1-1-array-and-array-index-representation.png)

### Tại sao cần học về Array?

- Quản lý nhiều giá trị cùng loại một cách có tổ chức
- Truy cập nhanh đến phần tử thông qua chỉ số (index)
- Nền tảng để hiểu các cấu trúc dữ liệu phức tạp hơn như Slice, Map
- Được sử dụng rộng rãi trong thuật toán và xử lý dữ liệu

---

## 2. Khái niệm cơ bản

### Định nghĩa

**Array** là một tập hợp **có kích thước cố định** chứa các phần tử có **cùng kiểu dữ liệu**.

### Đặc điểm quan trọng

| Đặc điểm                   | Mô tả                                                                           |
| -------------------------- | ------------------------------------------------------------------------------- |
| **Kích thước cố định**     | Sau khi khai báo, không thể thay đổi số lượng phần tử                           |
| **Kiểu dữ liệu đồng nhất** | Tất cả phần tử phải cùng kiểu (int, string, float64...)                         |
| **Chỉ số bắt đầu từ 0**    | Phần tử đầu tiên có index = 0                                                   |
| **Vùng nhớ liên tục**      | Các phần tử được lưu liền kề nhau trong bộ nhớ                                  |
| **Giá trị mặc định**       | Phần tử chưa gán có giá trị zero value (0 cho số, "" cho chuỗi, false cho bool) |

### Cú pháp khai báo

```go
var tênMảng [kíchThước]kiểuDữLiệu
```

**Ví dụ:**

```go
var numbers [5]int        // Mảng 5 số nguyên
var names [10]string      // Mảng 10 chuỗi
var flags [3]bool         // Mảng 3 giá trị boolean
```

---

## 3. Cách khai báo và khởi tạo Array

### 3.1 Khai báo không gán giá trị

Khi khai báo mà không gán giá trị, các phần tử sẽ có giá trị mặc định (zero value).

```go
package main

import "fmt"

func main() {
    // Mảng số nguyên - giá trị mặc định là 0
    var nums [5]int
    fmt.Println("Mảng số nguyên:", nums)
    // Output: [0 0 0 0 0]

    // Mảng chuỗi - giá trị mặc định là ""
    var names [3]string
    fmt.Println("Mảng chuỗi:", names)
    // Output: [  ]

    // Mảng boolean - giá trị mặc định là false
    var flags [2]bool
    fmt.Println("Mảng boolean:", flags)
    // Output: [false false]
}
```

### 3.2 Khai báo và gán giá trị từng phần tử

```go
package main

import "fmt"

func main() {
    var scores [5]int
    
    // Gán giá trị cho từng phần tử
    scores[0] = 85
    scores[1] = 90
    scores[2] = 78
    scores[3] = 92
    scores[4] = 88
    
    fmt.Println("Điểm các môn học:", scores)
    // Output: [85 90 78 92 88]
}
```

### 3.3 Khởi tạo trực tiếp với giá trị

**Cách 1: Chỉ định kích thước**

```go
package main

import "fmt"

func main() {
    // Khai báo ngắn gọn với :=
    temperatures := [7]float64{25.5, 26.0, 24.8, 27.2, 26.5, 25.9, 26.3}
    fmt.Println("Nhiệt độ 7 ngày:", temperatures)
    
    // Hoặc dùng var
    var prices [4]int = [4]int{100, 200, 150, 300}
    fmt.Println("Giá sản phẩm:", prices)
}
```

**Cách 2: Để compiler tự tính kích thước**

```go
package main

import "fmt"

func main() {
    // Dùng ... để Go tự đếm số phần tử
    weekdays := [...]string{"T2", "T3", "T4", "T5", "T6", "T7", "CN"}
    fmt.Println("Các ngày trong tuần:", weekdays)
    fmt.Println("Số ngày:", len(weekdays)) // 7
}
```

### 3.4 Khởi tạo một phần

Bạn có thể chỉ gán giá trị cho một số vị trí cụ thể, các vị trí còn lại sẽ nhận giá trị mặc định.

```go
package main

import "fmt"

func main() {
    // Chỉ gán giá trị tại index 0 và 4
    nums := [5]int{0: 10, 4: 50}
    fmt.Println(nums)
    // Output: [10 0 0 0 50]
    
    // Ví dụ khác
    months := [12]string{0: "Jan", 11: "Dec"}
    fmt.Println(months)
    // Output: [Jan           Dec]
}
```

---

## 4. Truy cập và thay đổi phần tử

### 4.1 Truy cập phần tử qua index

```go
package main

import "fmt"

func main() {
    fruits := [5]string{"Apple", "Banana", "Orange", "Mango", "Grape"}
    
    // Truy cập phần tử
    fmt.Println("Phần tử đầu tiên:", fruits[0])   // Apple
    fmt.Println("Phần tử thứ 3:", fruits[2])      // Orange
    fmt.Println("Phần tử cuối:", fruits[4])       // Grape
    
    // Hoặc dùng len() - 1 để lấy phần tử cuối
    fmt.Println("Phần tử cuối:", fruits[len(fruits)-1]) // Grape
}
```

### 4.2 Thay đổi giá trị phần tử

```go
package main

import "fmt"

func main() {
    prices := [4]float64{99.99, 149.50, 79.99, 199.00}
    
    fmt.Println("Giá ban đầu:", prices)
    
    // Giảm giá sản phẩm thứ 2 (index 1)
    prices[1] = 129.50
    
    // Tăng giá sản phẩm cuối
    prices[3] = 229.00
    
    fmt.Println("Giá sau khi điều chỉnh:", prices)
}
```

### 4.3 Lỗi thường gặp - Index ngoài phạm vi

```go
package main

import "fmt"

func main() {
    nums := [3]int{10, 20, 30}
    
    // OK - index hợp lệ (0, 1, 2)
    fmt.Println(nums[0]) // 10
    fmt.Println(nums[2]) // 30
    
    // LỖI - index ngoài phạm vi
    // fmt.Println(nums[3]) // panic: runtime error: index out of range
    // fmt.Println(nums[-1]) // compile error: invalid array index
}
```

> **Lưu ý:** Golang kiểm tra index trong quá trình compile (nếu dùng hằng số) hoặc runtime (nếu dùng biến). Truy cập index ngoài phạm vi sẽ gây panic.

---

## 5. Array đa chiều

### 5.1 Array 2 chiều (Ma trận)

Array 2 chiều thường được dùng để biểu diễn bảng, ma trận, hoặc lưới.

```go
package main

import "fmt"

func main() {
    // Ma trận 3x3
    var matrix [3][3]int = [3][3]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    
    // In toàn bộ ma trận
    fmt.Println("Ma trận 3x3:")
    fmt.Println(matrix)
    
    // Truy cập phần tử tại hàng 2, cột 3 (index [1][2])
    fmt.Println("\nPhần tử [1][2] =", matrix[1][2]) // 6
    
    // Thay đổi giá trị
    matrix[2][2] = 100
    fmt.Println("\nSau khi thay đổi [2][2]:")
    fmt.Println(matrix)
}
```

### 5.2 Kích thước ma trận

```go
package main

import "fmt"

func main() {
    matrix := [3][4]int{
        {1, 2, 3, 4},
        {5, 6, 7, 8},
        {9, 10, 11, 12},
    }
    
    rows := len(matrix)           // Số hàng
    cols := len(matrix[0])        // Số cột
    
    fmt.Printf("Ma trận có kích thước: %d x %d\n", rows, cols)
    // Output: Ma trận có kích thước: 3 x 4
}
```

### 5.3 Ví dụ thực tế - Bảng điểm sinh viên

```go
package main

import "fmt"

func main() {
    // 4 sinh viên, mỗi sinh viên có 3 môn học
    scores := [4][3]float64{
        {8.5, 7.0, 9.0},  // Sinh viên 1
        {7.5, 8.0, 7.5},  // Sinh viên 2
        {9.0, 9.5, 8.5},  // Sinh viên 3
        {6.5, 7.0, 8.0},  // Sinh viên 4
    }
    
    // Tính điểm trung bình của từng sinh viên
    for i := 0; i < len(scores); i++ {
        sum := 0.0
        for j := 0; j < len(scores[i]); j++ {
            sum += scores[i][j]
        }
        average := sum / float64(len(scores[i]))
        fmt.Printf("Sinh viên %d - Điểm TB: %.2f\n", i+1, average)
    }
}
```

**Output:**

```
Sinh viên 1 - Điểm TB: 8.17
Sinh viên 2 - Điểm TB: 7.67
Sinh viên 3 - Điểm TB: 9.00
Sinh viên 4 - Điểm TB: 7.17
```

### 5.4 Các thao tác với ma trận

**In đường chéo chính:**

```go
package main

import "fmt"

func main() {
    matrix := [3][3]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    
    fmt.Println("Đường chéo chính:")
    for i := 0; i < len(matrix); i++ {
        fmt.Printf("%d ", matrix[i][i])
    }
    fmt.Println()
    // Output: 1 5 9
}
```

**In đường chéo phụ:**

```go
package main

import "fmt"

func main() {
    matrix := [3][3]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    
    fmt.Println("Đường chéo phụ:")
    for i := 0; i < len(matrix); i++ {
        fmt.Printf("%d ", matrix[i][len(matrix)-i-1])
    }
    fmt.Println()
    // Output: 3 5 7
}
```

**Chuyển vị ma trận:**

```go
package main

import "fmt"

func main() {
    matrix := [3][2]int{
        {1, 2},
        {3, 4},
        {5, 6},
    }
    
    // Ma trận chuyển vị sẽ có kích thước 2x3
    var transpose [2][3]int
    
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[i]); j++ {
            transpose[j][i] = matrix[i][j]
        }
    }
    
    fmt.Println("Ma trận gốc:")
    for _, row := range matrix {
        fmt.Println(row)
    }
    
    fmt.Println("\nMa trận chuyển vị:")
    for _, row := range transpose {
        fmt.Println(row)
    }
}
```

---

## 6. Duyệt Array với for-range

### 6.1 For loop truyền thống

```go
package main

import "fmt"

func main() {
    numbers := [5]int{10, 20, 30, 40, 50}
    
    // Duyệt bằng for truyền thống
    for i := 0; i < len(numbers); i++ {
        fmt.Printf("numbers[%d] = %d\n", i, numbers[i])
    }
}
```

### 6.2 For-range (Được khuyến nghị)

For-range giúp code ngắn gọn và an toàn hơn.

```go
package main

import "fmt"

func main() {
    fruits := [4]string{"Apple", "Banana", "Orange", "Grape"}
    
    // index: chỉ số, value: giá trị
    for index, value := range fruits {
        fmt.Printf("[%d] = %s\n", index, value)
    }
}
```

**Output:**

```
[0] = Apple
[1] = Banana
[2] = Orange
[3] = Grape
```

### 6.3 Bỏ qua index hoặc value

**Chỉ cần value (không cần index):**

```go
package main

import "fmt"

func main() {
    prices := [3]float64{99.99, 149.50, 79.99}
    
    total := 0.0
    for _, price := range prices { // Dùng _ để bỏ qua index
        total += price
    }
    
    fmt.Printf("Tổng tiền: %.2f\n", total)
}
```

**Chỉ cần index (không cần value):**

```go
package main

import "fmt"

func main() {
    data := [5]int{10, 20, 30, 40, 50}
    
    for index := range data { // Không cần khai báo value
        fmt.Printf("Index: %d\n", index)
    }
}
```

### 6.4 Duyệt array 2 chiều

```go
package main

import "fmt"

func main() {
    matrix := [3][3]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    
    // Cách 1: Duyệt từng hàng
    fmt.Println("In từng hàng:")
    for _, row := range matrix {
        fmt.Println(row)
    }
    
    fmt.Println("\nIn chi tiết từng phần tử:")
    // Cách 2: Duyệt chi tiết từng phần tử
    for i, row := range matrix {
        for j, value := range row {
            fmt.Printf("[%d][%d] = %d\n", i, j, value)
        }
    }
}
```

---

## 7. Array vs Slice - Sự khác biệt quan trọng

### 7.1 So sánh Array và Slice

|Tiêu chí|Array|Slice|
|---|---|---|
|**Kích thước**|Cố định, không đổi|Động, có thể thay đổi|
|**Khai báo**|`[n]Type`|`[]Type`|
|**Giá trị mặc định**|Zero value của kiểu|`nil`|
|**So sánh**|Có thể so sánh với `==`|Không so sánh được trực tiếp|
|**Truyền hàm**|Truyền theo giá trị (copy)|Truyền tham chiếu|
|**Linh hoạt**|Thấp|Cao|
|**Sử dụng**|Ít phổ biến|Rất phổ biến|

### 7.2 Ví dụ minh họa

```go
package main

import "fmt"

func main() {
    // ARRAY - Kích thước cố định
    var arr [3]int = [3]int{1, 2, 3}
    fmt.Printf("Array: %v, Type: %T\n", arr, arr)
    // Output: Array: [1 2 3], Type: [3]int
    
    // SLICE - Kích thước động
    var slice []int = []int{1, 2, 3}
    fmt.Printf("Slice: %v, Type: %T\n", slice, slice)
    // Output: Slice: [1 2 3], Type: []int
    
    // Thêm phần tử vào slice (KHÔNG THỂ với array)
    slice = append(slice, 4, 5)
    fmt.Println("Slice sau khi append:", slice)
    // Output: [1 2 3 4 5]
}
```

### 7.3 Khi nào dùng Array?

- Khi biết chắc số lượng phần tử và không thay đổi
- Cần hiệu suất cao (không cần cấp phát thêm bộ nhớ)
- Ví dụ: Lưu tọa độ 3D `[3]float64{x, y, z}`, ma trận cố định, RGB color `[3]uint8{r, g, b}`

### 7.4 Khi nào dùng Slice?

- Không biết trước số lượng phần tử
- Cần thêm/xóa phần tử linh hoạt
- Hầu hết các trường hợp trong thực tế

> **Khuyến nghị:** Trong Go, Slice được sử dụng phổ biến hơn Array rất nhiều. Tuy nhiên, hiểu Array là nền tảng để làm việc tốt với Slice.

---

## 8. Kết hợp Array với Struct

### 8.1 Array chứa Struct

Đây là cách phổ biến để quản lý danh sách các đối tượng có cấu trúc.

```go
package main

import "fmt"

// Định nghĩa struct Student
type Student struct {
    ID    string
    Name  string
    Age   int
    Grade float64
}

func main() {
    // Tạo array chứa 3 sinh viên
    students := [3]Student{
        {ID: "SV001", Name: "Nguyễn Văn A", Age: 20, Grade: 8.5},
        {ID: "SV002", Name: "Trần Thị B", Age: 19, Grade: 9.0},
        {ID: "SV003", Name: "Lê Văn C", Age: 21, Grade: 7.5},
    }
    
    // In danh sách sinh viên
    fmt.Println("Danh sách sinh viên:")
    for i, student := range students {
        fmt.Printf("%d. ID: %s | Tên: %s | Tuổi: %d | Điểm: %.1f\n",
            i+1, student.ID, student.Name, student.Age, student.Grade)
    }
}
```

**Output:**

```
Danh sách sinh viên:
1. ID: SV001 | Tên: Nguyễn Văn A | Tuổi: 20 | Điểm: 8.5
2. ID: SV002 | Tên: Trần Thị B | Tuổi: 19 | Điểm: 9.0
3. ID: SV003 | Tên: Lê Văn C | Tuổi: 21 | Điểm: 7.5
```

### 8.2 Ví dụ thực tế - Quản lý sản phẩm

```go
package main

import "fmt"

type Product struct {
    Code     string
    Name     string
    Price    float64
    Quantity int
}

func main() {
    // Danh sách sản phẩm trong kho
    products := [4]Product{
        {"P001", "Laptop Dell", 15000000, 10},
        {"P002", "iPhone 15", 25000000, 5},
        {"P003", "iPad Pro", 20000000, 8},
        {"P004", "AirPods", 5000000, 15},
    }
    
    // Tính tổng giá trị hàng trong kho
    totalValue := 0.0
    for _, product := range products {
        value := product.Price * float64(product.Quantity)
        totalValue += value
        fmt.Printf("%s - %s: %.0f VND x %d = %.0f VND\n",
            product.Code, product.Name, product.Price, 
            product.Quantity, value)
    }
    
    fmt.Printf("\nTổng giá trị kho: %.0f VND\n", totalValue)
}
```

### 8.3 Struct chứa Array

```go
package main

import "fmt"

type Classroom struct {
    Name      string
    Students  [30]string // Tối đa 30 sinh viên
    Count     int        // Số sinh viên thực tế
}

func main() {
    class := Classroom{
        Name:  "Lớp 12A1",
        Count: 3,
    }
    
    // Thêm sinh viên
    class.Students[0] = "Nguyễn Văn A"
    class.Students[1] = "Trần Thị B"
    class.Students[2] = "Lê Văn C"
    
    fmt.Printf("Lớp: %s\n", class.Name)
    fmt.Printf("Sĩ số: %d/%d\n", class.Count, len(class.Students))
    fmt.Println("Danh sách:")
    for i := 0; i < class.Count; i++ {
        fmt.Printf("%d. %s\n", i+1, class.Students[i])
    }
}
```

---

## 9. Các thao tác phổ biến với Array

### 9.1 Tìm giá trị lớn nhất / nhỏ nhất

```go
package main

import "fmt"

func main() {
    numbers := [7]int{45, 23, 67, 12, 89, 34, 56}
    
    max := numbers[0]
    min := numbers[0]
    
    for _, num := range numbers {
        if num > max {
            max = num
        }
        if num < min {
            min = num
        }
    }
    
    fmt.Println("Mảng:", numbers)
    fmt.Println("Giá trị lớn nhất:", max)  // 89
    fmt.Println("Giá trị nhỏ nhất:", min)  // 12
}
```

### 9.2 Tính tổng và trung bình

```go
package main

import "fmt"

func main() {
    scores := [5]float64{8.5, 7.0, 9.0, 6.5, 8.0}
    
    sum := 0.0
    for _, score := range scores {
        sum += score
    }
    
    average := sum / float64(len(scores))
    
    fmt.Printf("Tổng điểm: %.2f\n", sum)
    fmt.Printf("Điểm trung bình: %.2f\n", average)
}
```

### 9.3 Tìm kiếm phần tử

```go
package main

import "fmt"

func main() {
    fruits := [5]string{"Apple", "Banana", "Orange", "Mango", "Grape"}
    search := "Orange"
    
    found := false
    index := -1
    
    for i, fruit := range fruits {
        if fruit == search {
            found = true
            index = i
            break
        }
    }
    
    if found {
        fmt.Printf("Tìm thấy '%s' tại vị trí %d\n", search, index)
    } else {
        fmt.Printf("Không tìm thấy '%s'\n", search)
    }
}
```

### 9.4 Đếm số lần xuất hiện

```go
package main

import "fmt"

func main() {
    numbers := [10]int{1, 2, 3, 2, 4, 2, 5, 3, 2, 1}
    target := 2
    
    count := 0
    for _, num := range numbers {
        if num == target {
            count++
        }
    }
    
    fmt.Printf("Số %d xuất hiện %d lần\n", target, count)
    // Output: Số 2 xuất hiện 4 lần
}
```

### 9.5 Đảo ngược mảng

```go
package main

import "fmt"

func main() {
    arr := [5]int{10, 20, 30, 40, 50}
    
    fmt.Println("Mảng ban đầu:", arr)
    
    // Đảo ngược bằng cách swap các phần tử
    for i := 0; i < len(arr)/2; i++ {
        j := len(arr) - 1 - i
        arr[i], arr[j] = arr[j], arr[i]
    }
    
    fmt.Println("Mảng sau khi đảo:", arr)
    // Output: [50 40 30 20 10]
}
```

### 9.6 Sao chép Array

```go
package main

import "fmt"

func main() {
    original := [4]int{10, 20, 30, 40}
    
    // Cách 1: Gán trực tiếp (copy toàn bộ)
    copy1 := original
    
    // Cách 2: Sao chép thủ công
    var copy2 [4]int
    for i, v := range original {
        copy2[i] = v
    }
    
    // Thay đổi bản sao không ảnh hưởng bản gốc
    copy1[0] = 999
    
    fmt.Println("Original:", original) // [10 20 30 40]
    fmt.Println("Copy1:", copy1)       // [999 20 30 40]
    fmt.Println("Copy2:", copy2)       // [10 20 30 40]
}
```

### 9.7 So sánh hai Array

```go
package main

import "fmt"

func main() {
    arr1 := [3]int{1, 2, 3}
    arr2 := [3]int{1, 2, 3}
    arr3 := [3]int{1, 2, 4}
    
    // Array có thể so sánh trực tiếp
    fmt.Println("arr1 == arr2:", arr1 == arr2) // true
    fmt.Println("arr1 == arr3:", arr1 == arr3) // false
}
```

---

## 10. Lưu ý quan trọng khi làm việc với Array

### 10.1 Array là value type (kiểu giá trị)

Khi truyền array vào hàm, Go sẽ **copy toàn bộ** array, không phải tham chiếu.

```go
package main

import "fmt"

func modifyArray(arr [3]int) {
    arr[0] = 999
    fmt.Println("Trong hàm:", arr)
}

func main() {
    numbers := [3]int{1, 2, 3}
    
    modifyArray(numbers)
    fmt.Println("Ngoài hàm:", numbers)
}
```

**Output:**

```
Trong hàm: [999 2 3]
Ngoài hàm: [1 2 3]   ← Không thay đổi!
```

**Giải pháp:** Dùng pointer hoặc slice nếu muốn thay đổi array gốc.

```go
package main

import "fmt"

func modifyArray(arr *[3]int) {
    arr[0] = 999
}

func main() {
    numbers := [3]int{1, 2, 3}
    
    modifyArray(&numbers)  // Truyền địa chỉ
    fmt.Println("Sau khi modify:", numbers)
    // Output: [999 2 3]
}
```

### 10.2 Kích thước là một phần của kiểu dữ liệu

```go
package main

func main() {
    var arr1 [3]int
    var arr2 [5]int
    
    // arr1 = arr2  // LỖI: cannot use arr2 (type [5]int) as type [3]int
    
    // [3]int và [5]int là HAI KIỂU DỮ LIỆU KHÁC NHAU
}
```

### 10.3 Không thể thay đổi kích thước

```go
package main

func main() {
    arr := [3]int{1, 2, 3}
    
    // KHÔNG THỂ thêm phần tử vào array
    // arr = append(arr, 4)  // LỖI: first argument to append must be slice
    
    // Nếu cần thêm phần tử, dùng Slice
    slice := []int{1, 2, 3}
    slice = append(slice, 4)  // OK với slice
}
```

### 10.4 Index phải trong phạm vi hợp lệ

```go
package main

import "fmt"

func main() {
    arr := [3]int{10, 20, 30}
    
    // Hợp lệ: index từ 0 đến 2
    fmt.Println(arr[0])  // OK
    fmt.Println(arr[2])  // OK
    
    // Lỗi runtime
    // fmt.Println(arr[3])  // panic: index out of range
    
    // An toàn hơn: kiểm tra trước
    index := 5
    if index >= 0 && index < len(arr) {
        fmt.Println(arr[index])
    } else {
        fmt.Println("Index không hợp lệ")
    }
}
```

### 10.5 Hiệu suất với Array lớn

Với array lớn, việc copy có thể tốn bộ nhớ và thời gian. Hãy cân nhắc dùng slice hoặc pointer.

```go
package main

import "fmt"

func processLargeArray(arr [1000000]int) {
    // Copy 1 triệu phần tử → tốn bộ nhớ
    fmt.Println(arr[0])
}

func processLargeSlice(slice []int) {
    // Chỉ truyền tham chiếu → hiệu quả hơn
    fmt.Println(slice[0])
}

func main() {
    var largeArr [1000000]int
    processLargeArray(largeArr)  // Copy toàn bộ
    
    largeSlice := make([]int, 1000000)
    processLargeSlice(largeSlice)  // Chỉ truyền con trỏ
}
```

---

## Tóm tắt

### Những điều cần nhớ về Array

1. **Kích thước cố định** - Không thể thay đổi sau khi khai báo
2. **Kiểu đồng nhất** - Tất cả phần tử cùng kiểu dữ liệu
3. **Index từ 0** - Phần tử đầu tiên có index = 0
4. **Value type** - Truyền vào hàm sẽ copy toàn bộ
5. **Ít dùng hơn Slice** - Trong thực tế, Slice phổ biến hơn

### Khi nào nên dùng Array?

- Số lượng phần tử cố định và biết trước
- Cần hiệu suất cao với kích thước nhỏ
- Dữ liệu có cấu trúc cố định (RGB, tọa độ 3D...)

### Khi nào nên dùng Slice?

- Số lượng phần tử động, không biết trước
- Cần thêm/xóa phần tử linh hoạt
- Hầu hết các trường hợp trong thực tế