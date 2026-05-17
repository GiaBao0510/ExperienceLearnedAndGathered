# Slice trong Go - Hướng dẫn Chi tiết

## 📌 Mục lục

1. [Slice là gì?](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#1-slice-l%C3%A0-g%C3%AC)
2. [So sánh Array vs Slice](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#2-so-s%C3%A1nh-array-vs-slice)
3. [Các cách khởi tạo Slice](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#3-c%C3%A1c-c%C3%A1ch-kh%E1%BB%9Fi-t%E1%BA%A1o-slice)
4. [Thao tác cơ bản với Slice](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#4-thao-t%C3%A1c-c%C6%A1-b%E1%BA%A3n-v%E1%BB%9Bi-slice)
5. [Sub-slice (Slice con)](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#5-sub-slice-slice-con)
6. [Hiểu về Length và Capacity](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#6-hi%E1%BB%83u-v%E1%BB%81-length-v%C3%A0-capacity)
7. [Các hàm xử lý Slice](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#7-c%C3%A1c-h%C3%A0m-x%E1%BB%AD-l%C3%BD-slice)
8. [Lưu ý quan trọng](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#8-l%C6%B0u-%C3%BD-quan-tr%E1%BB%8Dng)

---

## 1. Slice là gì?

**Slice** là một cấu trúc dữ liệu trong Go cho phép lưu trữ một tập hợp các phần tử có cùng kiểu dữ liệu với **kích thước có thể thay đổi linh hoạt**.

### 🔑 Đặc điểm chính:

- ✅ Kích thước **động** (có thể thay đổi khi runtime)
- ✅ Được xây dựng dựa trên **Array** nhưng linh hoạt hơn
- ✅ Tự động **mở rộng** khi cần thêm phần tử
- ✅ Là **kiểu tham chiếu** (reference type)

### 📊 Minh họa cấu trúc Slice:

```
Slice:  [pointer] [length] [capacity]
           |
           v
Array:   [1] [2] [3] [4] [5] [_] [_] [_]
         └──────┬──────┘
            length=5
         └──────────┬──────────┘
              capacity=8
```

![](https://zalopay-oss.github.io/go-advanced/images/1-3-golang-slices-length-capacity.png)

`Slices` thì phức tạp hơn, cấu trúc của chúng cũng như `string`, tuy nhiên việc giới hạn chỉ-đọc như string được lược bỏ.

Cấu trúc của slice là `reflect.SliceHeader`
```go
type  SliceHeader  struct {
    Data  uintptr 
    Len   int 
    Cap   int 
}
```

Slice được xem là fat pointer *(mọi người có thể đọc thêm ở bài viết sau để hiểu hơn về fat pointer trong [Go](https://nullprogram.com/blog/2019/06/30/))* Cấu trúc slice bao gồm:

- `Data`: chứa đỉa chỉ trong bộ nhớ của con trỏ trỏ tới underlying array của slice.
- `Len`: độ dài của slice.
- `Cap`: kích thước tối đa mà vùng nhớ trỏ tới slice được cấp phát.

Hình bên dưới sẽ miêu tả slice `x := []int{2,3,5,7,11}` và slice `y := x[1:3]`:
![](https://zalopay-oss.github.io/go-advanced/images/ch1-10-slice-1.ditaa.png)

---

## 2. So sánh Array vs Slice

|Tiêu chí|Array|Slice|
|---|---|---|
|**Kích thước**|Cố định, xác định lúc compile|Động, có thể thay đổi|
|**Khai báo**|`[5]int{1,2,3,4,5}`|`[]int{1,2,3,4,5}`|
|**Kiểu dữ liệu**|Value type|Reference type|
|**Linh hoạt**|Thấp|Cao|
|**Sử dụng**|Ít phổ biến|Rất phổ biến|

### Ví dụ minh họa:

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Khai báo Slice (không có số trong [] || nil slice)
	var list []int
	fmt.Println("Slice rỗng:", list) // Output: []
	
	// empty slice, khác với nil
    b = []int{}

	// Khởi tạo Slice với giá trị
	list = []int{1, 2, 3, 4, 5} // Kích thước linh hoạt
	
	// Khởi tạo Array với giá trị
	arr := [5]int{1, 2, 3, 4, 5} // Kích thước cố định
	
	// Phân biệt kiểu dữ liệu
	fmt.Println("Kiểu của Array:", reflect.TypeOf(arr).Kind()) // Output: array
	fmt.Println("Kiểu của Slice:", reflect.TypeOf(list).Kind()) // Output: slice
	
	// Thêm phần tử vào Slice (OK)
	list = append(list, 6, 7, 8)
	fmt.Println("Slice sau khi thêm:", list) // Output: [1 2 3 4 5 6 7 8]
	
	// Thêm phần tử vào Array (KHÔNG THỂ)
	// arr = append(arr, 6) // ❌ Compile error!
}
```

---

## 3. Các cách khởi tạo Slice

### 3.1. Khởi tạo từ Array

```go
func main() {
	// Tạo Array trước
	arr := [5]int{11, 12, 13, 14, 15}
	
	// Tạo Slice từ Array (lấy từ index 1 đến 3)
	slices := arr[1:4] // Lấy: 12, 13, 14
	
	fmt.Printf("Kiểu: %v\n", reflect.TypeOf(slices).Kind()) // slice
	fmt.Println("Giá trị:", slices) // [12 13 14]
}
```

**Lưu ý:** `arr[1:4]` nghĩa là lấy từ index 1 (bao gồm) đến index 4 (không bao gồm)

### 3.2. Khởi tạo trực tiếp với `[]`

```go
func main() {
	// Cách đơn giản nhất - khuyên dùng
	slices := []int{1, 2, 3, 4, 5}
	
	fmt.Printf("Kiểu: %v\n", reflect.TypeOf(slices).Kind()) // slice
	fmt.Println("Giá trị:", slices) // [1 2 3 4 5]
}
```

### 3.3. Khởi tạo với hàm `make()`

```go
func main() {
	/*
	   Cú pháp: make([]Type, length, capacity)
	   - Type: Kiểu dữ liệu của phần tử
	   - length: Số phần tử khởi tạo ban đầu (giá trị mặc định là 0)
	   - capacity: Dung lượng tối đa có thể chứa (tùy chọn)
	*/
	
	slices := make([]int, 3, 5)
	// Tạo slice có 3 phần tử (giá trị 0), capacity = 5
	
	fmt.Printf("Slice: %v\n", slices)                    // [0 0 0]
	fmt.Printf("Length (độ dài): %d\n", len(slices))     // 3
	fmt.Printf("Capacity (dung lượng): %d\n", cap(slices)) // 5
	
	// Thêm phần tử
	slices = append(slices, 5, 8, 9)
	fmt.Printf("Sau khi thêm: %v\n", slices)             // [0 0 0 5 8 9]
	fmt.Printf("Length mới: %d\n", len(slices))          // 6
	fmt.Printf("Capacity mới: %d\n", cap(slices))        // 10 (tự động tăng gấp đôi)
}
```

### 📝 Khi nào dùng `make()`?

- Khi biết trước **kích thước ước lượng** để tối ưu hiệu năng
- Khi muốn **tránh reallocate** nhiều lần
- Khi tạo slice để **đọc dữ liệu** vào (buffering)

```go
// Tốt: Biết trước cần 1000 phần tử
data := make([]int, 0, 1000)

// Không tốt: Phải reallocate nhiều lần
data := []int{}
for i := 0; i < 1000; i++ {
    data = append(data, i) // Mỗi lần append có thể phải copy lại toàn bộ
}
```

---

## 4. Thao tác cơ bản với Slice

### 4.1. Duyệt qua Slice

```go
func main() {
	slices := []string{"a", "b", "c", "d", "e"}
	
	// Cách 1: Dùng for truyền thống
	fmt.Println("=== Cách 1 ===")
	for i := 0; i < len(slices); i++ {
		fmt.Printf("Index %d: %s\n", i, slices[i])
	}
	
	// Cách 2: Dùng range (khuyên dùng)
	fmt.Println("\n=== Cách 2 ===")
	for index, value := range slices {
		fmt.Printf("Index %d: %s\n", index, value)
	}
	
	// Cách 3: Chỉ lấy value, bỏ qua index
	fmt.Println("\n=== Cách 3 ===")
	for _, value := range slices {
		fmt.Println(value)
	}
	
	// Cách 4: Chỉ lấy index, bỏ qua value
	fmt.Println("\n=== Cách 4 ===")
	for index := range slices {
		fmt.Println("Index:", index)
	}
}
```

### 4.2. Thêm phần tử với `append()`

```go
func main() {
	// Thêm một phần tử
	slices := []int{1, 2, 3}
	slices = append(slices, 4)
	fmt.Println(slices) // [1 2 3 4]
	
	// Thêm nhiều phần tử
	slices = append(slices, 5, 6, 7)
	fmt.Println(slices) // [1 2 3 4 5 6 7]
	
	// Thêm một slice vào slice khác (dùng ...)
	slice2 := []int{8, 9, 10}
	slices = append(slices, slice2...)
	fmt.Println(slices) // [1 2 3 4 5 6 7 8 9 10]
}
```

**⚠️ Lưu ý quan trọng:** Luôn gán lại kết quả của `append()` vì nó có thể trả về slice mới!

```go
// ❌ SAI: Không gán lại
slices := []int{1, 2, 3}
append(slices, 4) // Không có tác dụng!
fmt.Println(slices) // [1 2 3]

// ✅ ĐÚNG: Gán lại kết quả
slices = append(slices, 4)
fmt.Println(slices) // [1 2 3 4]
```

Bên cạnh thêm phần tử vào cuối slice, chúng ta cũng có thể thêm phần tử vào đầu slice như sau

```go
var a = []int{1,2,3}
// thêm phần tử 0 vào đầu slice a
a = append([]int{0}, a...)
// thêm các phần tử -3, -2, -1 vào đầu slice a
a = append([]int{-3,-2,-1}, a...)
```

Thêm phần tử vào đầu slice sẽ gây ra việc cấp phát lại vùng nhớ và làm những phần tử đang tồn tại trong slice sẽ được sao chép một lần nữa. Do đó, hiệu suất của việc thêm phần tử vào đầu slice sẽ thấp hơn thêm phần tử vào cuối slice.

### 4.3. Slice lồng nhau (2D Slice)

```go
func main() {
	// Slice 2 chiều (giống mảng 2 chiều)
	school := [][]string{
		{"6/1", "6/2", "6/3"},
		{"7/1", "7/2", "7/3"},
		{"8/1", "8/2", "8/3"},
	}
	
	fmt.Println("Trường học ban đầu:")
	fmt.Println(school)
	
	// Thêm khối 9
	school = append(school, []string{"9/1", "9/2", "9/3", "9/4"})
	
	// Thêm nhiều khối cùng lúc
	school2 := [][]string{
		{"10/1", "10/2", "10/3"},
		{"11/1", "11/2", "11/3"},
	}
	school = append(school, school2...)
	
	fmt.Println("\nTrường học sau khi thêm:")
	for khoi, lop := range school {
		fmt.Printf("Khối %d: %v\n", khoi+6, lop)
	}
}
```

**Output:**

```
Khối 6: [6/1 6/2 6/3]
Khối 7: [7/1 7/2 7/3]
Khối 8: [8/1 8/2 8/3]
Khối 9: [9/1 9/2 9/3 9/4]
Khối 10: [10/1 10/2 10/3]
Khối 11: [11/1 11/2 11/3]
```

---

## 5. Sub-slice (Slice con)

**Sub-slice** là việc tạo một slice mới từ một slice đã có sẵn bằng cách cắt (slicing).

### Cú pháp:

```
slice[start:end]
slice[:end]    // Từ đầu đến end
slice[start:]  // Từ start đến cuối
slice[:]       // Toàn bộ (copy tham chiếu)
```

### Ví dụ:

```go
func main() {
	slices := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	fmt.Printf("Slice gốc: %v\n", slices)
	
	// Lấy từ index 2 đến 8 (không bao gồm 9)
	subslice1 := slices[2:9]
	fmt.Printf("slices[2:9]: %v\n", subslice1) // [c d e f g h i]
	
	// Lấy từ đầu đến index 4 (không bao gồm 5)
	subslice2 := slices[:5]
	fmt.Printf("slices[:5]: %v\n", subslice2) // [a b c d e]
	
	// Lấy từ index 5 đến cuối
	subslice3 := slices[5:]
	fmt.Printf("slices[5:]: %v\n", subslice3) // [f g h i j k]
	
	// Lấy toàn bộ
	subslice4 := slices[:]
	fmt.Printf("slices[:]: %v\n", subslice4) // [a b c d e f g h i j k]
}
```

### ⚠️ Lưu ý: Sub-slice chia sẻ bộ nhớ với slice gốc!

```go
func main() {
	original := []int{1, 2, 3, 4, 5}
	sub := original[1:4] // [2 3 4]
	
	fmt.Println("Ban đầu:")
	fmt.Println("Original:", original) // [1 2 3 4 5]
	fmt.Println("Sub:", sub)           // [2 3 4]
	
	// Thay đổi sub-slice
	sub[0] = 999
	
	fmt.Println("\nSau khi sửa sub[0] = 999:")
	fmt.Println("Original:", original) // [1 999 3 4 5] ← Bị thay đổi!
	fmt.Println("Sub:", sub)           // [999 3 4]
}
```

**Giải pháp:** Nếu muốn tạo bản sao độc lập, dùng `copy()` hoặc `slices.Clone()`:

```go
import "slices"

original := []int{1, 2, 3, 4, 5}
sub := slices.Clone(original[1:4]) // Tạo bản sao độc lập
sub[0] = 999
// original không bị thay đổi
```

---

## 6. Hiểu về Length và Capacity

### Khái niệm:

- **Length (len)**: Số phần tử **hiện tại** trong slice
- **Capacity (cap)**: Số phần tử **tối đa** mà slice có thể chứa mà không cần cấp phát lại bộ nhớ

### Quy tắc:

```
0 ≤ len(slice) ≤ cap(slice)
```

### Ví dụ minh họa:

```go
func main() {
	// Slice gốc
	slices := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	fmt.Printf("Slice gốc: %v\n", slices)
	fmt.Printf("Length: %d, Capacity: %d\n\n", len(slices), cap(slices))
	// Length: 11, Capacity: 11
	
	// Tạo sub-slice
	subslice := slices[2:9] // Lấy từ index 2 đến 8
	fmt.Printf("Sub-slice: %v\n", subslice)
	fmt.Printf("Length: %d, Capacity: %d\n", len(subslice), cap(subslice))
	// Length: 7, Capacity: 9
}
```

**Output:**

```
Slice gốc: [a b c d e f g h i j k]
Length: 11, Capacity: 11

Sub-slice: [c d e f g h i]
Length: 7, Capacity: 9
```

### 🤔 Tại sao Capacity = 9?

Minh họa:

```
Slice gốc:  [a] [b] [c] [d] [e] [f] [g] [h] [i] [j] [k]
Index:       0   1   2   3   4   5   6   7   8   9   10

Sub-slice bắt đầu từ index 2:
            [c] [d] [e] [f] [g] [h] [i] [j] [k]
             └────────┬────────┘
               length = 7 (từ c đến i)
             └──────────────┬──────────────┘
               capacity = 9 (từ c đến k)
```

**Capacity của sub-slice** = Số phần tử từ điểm bắt đầu của sub-slice đến **cuối slice gốc**.

### Ví dụ khác:

```go
s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

s1 := s[2:5]  // [2 3 4]
fmt.Printf("len=%d cap=%d %v\n", len(s1), cap(s1), s1)
// len=3 cap=8 [2 3 4]
// Capacity = 8 vì từ index 2 đến cuối có 8 phần tử

s2 := s[5:8]  // [5 6 7]
fmt.Printf("len=%d cap=%d %v\n", len(s2), cap(s2), s2)
// len=3 cap=5 [5 6 7]
// Capacity = 5 vì từ index 5 đến cuối có 5 phần tử
```

Trong trường hợp slice ban đầu không đủ sức chứa khi thêm vào phần tử, hàm append sẽ hiện thực cấp phát lại vùng nhớ có kích thước:

- Nếu kích thước cũ (cap) < 1024: cấp phát gấp đôi (x2) vùng nhớ cũ.
- Nếu kích thước cũ >= 1024: cấp phát 1.25x vùng nhớ cũ.

Sau đó, dữ liệu cũ sẽ được sao chép sang.Mọi người có thể xem đoạn mã nguồn về việc cấp pháp lại vùng nhớ cho slice [ở đây](https://golang.org/src/runtime/slice.go?fbclid=IwAR0xgVnf7SFJu_Kai8zo_5PZXolsuEL3JgfKejj7Ww0CpO1G82rbXbcWosQ#L66).

![](https://zalopay-oss.github.io/go-advanced/images/recapacity-slice.png)

Ví dụ bên dưới cho thấy giá trị **cap tăng gấp 2** khi thực thi hàm append vượt quá kích thước ban đầu (< 1024)

```go
package main

import (
	"fmt"
)

func MyAppend(slice []int, value int) []int{
	slice = append(slice, value)
	PrintSlice(slice)
	return slice
}

func PrintSlice(slice []int){
	fmt.Printf("Slice: %v\n", slice)
	fmt.Printf("length: %v, cap: %v\n", len(slice), cap(slice))
}

func main() {
	slices := make([]int, 1)

	PrintSlice(slices)

	for i := 1; i <= 10; i++ {
		fmt.Println("\n---------------------")
		slices = MyAppend(slices, i)
	}
}
```

kết quả:
```bash
Go\Go_Tutorial\Test> go run .
Slice: [0]
length: 1, cap: 1

---------------------
Slice: [0 1]
length: 2, cap: 2

---------------------
Slice: [0 1 2]
length: 3, cap: 4

---------------------
Slice: [0 1 2 3]
length: 4, cap: 4

---------------------
Slice: [0 1 2 3 4]
length: 5, cap: 8

---------------------
Slice: [0 1 2 3 4 5]
length: 6, cap: 8
```

---
## 7. Các hàm xử lý Slice

**⚠️ Yêu cầu:** Package `slices` chỉ có từ **Go 1.21** trở lên.

```go
import "slices"
```

### 7.1. Clone - Sao chép Slice

```go
func main() {
	original := []string{"a", "b", "c", "d", "e"}
	
	// Tạo bản sao độc lập
	copied := slices.Clone(original)
	
	// Thay đổi bản sao
	copied[0] = "Z"
	
	fmt.Println("Original:", original) // [a b c d e]
	fmt.Println("Copied:", copied)     // [Z b c d e]
}
```

### 7.2. Equal - So sánh 2 Slice

```go
func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{1, 2, 4}
	
	fmt.Println(slices.Equal(s1, s2)) // true
	fmt.Println(slices.Equal(s1, s3)) // false
	
	// So sánh với toán tử == không hoạt động!
	// fmt.Println(s1 == s2) // ❌ Compile error
}
```

### 7.3. Contains - Kiểm tra phần tử có tồn tại

```go
func main() {
	fruits := []string{"apple", "banana", "orange"}
	
	fmt.Println(slices.Contains(fruits, "banana")) // true
	fmt.Println(slices.Contains(fruits, "grape"))  // false
}
```

### 7.4. Index - Tìm vị trí đầu tiên

```go
func main() {
	numbers := []int{1, 2, 3, 4, 5, 7, 7, 9, 5, 10, 1, 7, 2}
	
	// Tìm vị trí đầu tiên của số 7
	pos := slices.Index(numbers, 7)
	fmt.Println("Vị trí của 7:", pos) // 5
	
	// Không tìm thấy sẽ trả về -1
	pos = slices.Index(numbers, 100)
	fmt.Println("Vị trí của 100:", pos) // -1
}
```

### 7.5. Insert - Chèn phần tử

```go
func main() {
	s := []string{"a", "b", "c", "d", "e"}
	
	// Chèn "X" vào vị trí index 2
	s = slices.Insert(s, 2, "X")
	fmt.Println(s) // [a b X c d e]
	
	// Chèn nhiều phần tử
	s = slices.Insert(s, 5, "Y", "Z")
	fmt.Println(s) // [a b X c d Y Z e]
}
```

### 7.6. Delete - Xóa phần tử

```go
func main() {
	s := []string{"a", "b", "c", "d", "e", "f", "g"}
	
	// Xóa từ index 2 đến 4 (không bao gồm 5)
	s = slices.Delete(s, 2, 5)
	fmt.Println(s) // [a b f g]
	
	// Xóa một phần tử tại index 1
	s = slices.Delete(s, 1, 2)
	fmt.Println(s) // [a f g]
}
```

### 7.7. Reverse - Đảo ngược Slice

```go
func main() {
	s := []int{1, 2, 3, 4, 5}
	
	slices.Reverse(s) // Hàm này không có return value
	fmt.Println(s) // [5 4 3 2 1]
}
```

### 7.8. Sort - Sắp xếp

```go
func main() {
	// Sắp xếp số
	numbers := []int{5, 2, 8, 1, 9, 3}
	slices.Sort(numbers)
	fmt.Println(numbers) // [1 2 3 5 8 9]
	
	// Sắp xếp chuỗi
	words := []string{"dog", "cat", "bird", "ant"}
	slices.Sort(words)
	fmt.Println(words) // [ant bird cat dog]
}
```

### 7.9. SortFunc - Sắp xếp tùy chỉnh

```go
func main() {
	nums := []int{1, 5, 6, 7, 4, 3, 2, 10, 9}
	
	// Sắp xếp giảm dần (từ lớn đến nhỏ)
	slices.SortFunc(nums, func(a, b int) int {
		return b - a // Nếu b > a thì b sẽ đứng trước
	})
	
	fmt.Println(nums) // [10 9 7 6 5 4 3 2 1]
}
```

**Giải thích hàm compare:**

```go
func(a, b int) int {
    return b - a
}
```

- Nếu trả về **số âm**: `a` đứng trước `b`
- Nếu trả về **số dương**: `b` đứng trước `a`
- Nếu trả về **0**: Giữ nguyên vị trí

### 7.10. Min và Max

```go
func main() {
	nums := []int{5, 2, 8, 1, 9, 3}
	
	fmt.Println("Min:", slices.Min(nums)) // 1
	fmt.Println("Max:", slices.Max(nums)) // 9
}
```

### 📋 Bảng tổng hợp các hàm

|Hàm|Mô tả|Ví dụ|
|---|---|---|
|`Clone(s)`|Tạo bản sao|`slices.Clone(s)`|
|`Equal(s1, s2)`|So sánh 2 slice|`slices.Equal(s1, s2)`|
|`Contains(s, v)`|Kiểm tra tồn tại|`slices.Contains(s, "x")`|
|`Index(s, v)`|Tìm vị trí|`slices.Index(s, 5)`|
|`Insert(s, i, v)`|Chèn phần tử|`slices.Insert(s, 2, "x")`|
|`Delete(s, i, j)`|Xóa từ i đến j|`slices.Delete(s, 2, 5)`|
|`Reverse(s)`|Đảo ngược|`slices.Reverse(s)`|
|`Sort(s)`|Sắp xếp|`slices.Sort(s)`|
|`SortFunc(s, cmp)`|Sắp xếp tùy chỉnh|`slices.SortFunc(s, cmp)`|
|`Min(s)`|Giá trị nhỏ nhất|`slices.Min(s)`|
|`Max(s)`|Giá trị lớn nhất|`slices.Max(s)`|

---
## 8.  **Tránh gây ra memory leak trên slice:**

#### 8.1. Vấn đề "Mảng ẩn" (The Hidden Giant)

**Hiện tượng:** Khi bạn tạo một slice nhỏ từ một slice/mảng rất lớn, slice nhỏ này vẫn giữ tham chiếu đến **toàn bộ** mảng lớn ban đầu thông qua trường `Cap`. Chừng nào slice nhỏ còn tồn tại, Garbage Collector (Bộ thu gom rác - GC) sẽ không thể giải phóng mảng lớn, dù bạn chỉ dùng 1% diện tích của nó.

#### Ví dụ chưa tối ưu:

```go
func FindPhoneNumber(filename string) []byte {
    // Đọc toàn bộ file vào bộ nhớ (ví dụ file nặng 100MB)
    b, _ := os.ReadFile(filename) 
    
    // Tìm số điện thoại (chỉ chiếm khoảng vài chục byte)
    // Kết quả trả về thực chất vẫn trỏ vào mảng 100MB ban đầu
    return regexp.MustCompile("[0-9]+").Find(b)
}
```

**Hệ quả:** Bạn trả về 20 byte kết quả, nhưng thực tế bạn đang "găm" 100MB trong RAM cho đến khi kết quả đó không còn được sử dụng nữa.

#### Cách khắc phục: Sao chép sang vùng nhớ mới

Hãy tạo một slice mới hoàn toàn và copy dữ liệu cần thiết sang đó. Điều này giúp ngắt kết nối với mảng khổng lồ ban đầu.

```Go
func FindPhoneNumber(filename string) []byte {
    b, _ := os.ReadFile(filename)
    res := regexp.MustCompile("[0-9]+").Find(b)
    
    if res == nil {
        return nil
    }

    // Cách 1: Sử dụng append (ngắn gọn)
    // return append([]byte{}, res...)

    // Cách 2: Sử dụng copy (rõ ràng, khuyến khích cho sinh viên)
    newRes := make([]byte, len(res))
    copy(newRes, res)
    return newRes
}
```

#### 8.2. Vấn đề "Phần tử bóng ma" (The Lingering Ghost)

**Hiện tượng:** Khi làm việc với một **slice chứa con trỏ** (hoặc struct chứa con trỏ), việc bạn cắt ngắn slice (`a = a[:len(a)-1]`) chỉ đơn giản là giảm chỉ số `Len`. Phần tử vừa bị loại bỏ vẫn còn nằm trong mảng bên dưới và vẫn trỏ đến vùng nhớ của đối tượng.

Vì GC thấy vẫn còn một tham chiếu (ẩn) nằm trong mảng, nó sẽ không giải phóng đối tượng đó.

#### Ví dụ gây lãng phí bộ nhớ:

```Go
var a []*Student // Giả sử slice chứa 1000 sinh viên
// ... thêm dữ liệu ...

// Xóa phần tử cuối bằng cách cắt slice
a = a[:len(a)-1] 

// LÚC NÀY: Phần tử cuối vẫn nằm trong mảng ẩn bên dưới và trỏ tới Student.
// Student đó sẽ không bao giờ được giải phóng cho đến khi toàn bộ slice 'a' bị hủy.
```

#### Cách khắc phục: Gán Nil trước khi cắt

Để GC biết rằng đối tượng đó không còn được dùng nữa, ta cần "ngắt kết nối" thủ công.
```Go
var a []*Student
// ...

// 1. Gán nil cho phần tử cần xóa để xóa tham chiếu
a[len(a)-1] = nil

// 2. Sau đó mới cắt slice
a = a[:len(a)-1]
```


---
## 9. Lưu ý quan trọng

### ✅ Best Practice

1. **Luôn gán lại kết quả của `append()`**
    
    ```go
    // ✅ ĐÚNG
    s = append(s, 1)
    
    // ❌ SAI
    append(s, 1)
    ```
    
2. **Sử dụng `make()` khi biết trước kích thước**
    
    ```go
    // ✅ Tốt - tránh reallocate
    s := make([]int, 0, 1000)
    for i := 0; i < 1000; i++ {
        s = append(s, i)
    }
    
    // ⚠️ Kém - phải reallocate nhiều lần
    s := []int{}
    for i := 0; i < 1000; i++ {
        s = append(s, i)
    }
    ```
    
3. **Cẩn thận với sub-slice chia sẻ bộ nhớ**
    
    ```go
    // Dùng Clone nếu cần bản sao độc lập
    sub := slices.Clone(original[1:4])
    ```
    
4. **Dùng `range` để duyệt slice**
    
    ```go
    // ✅ Đơn giản, rõ ràng
    for _, v := range s {
        fmt.Println(v)
    }
    ```
    

### ⚠️ Các lỗi thường gặp

1. **Quên gán lại kết quả `append()`**
    
    ```go
    s := []int{1, 2, 3}
    append(s, 4) // ❌ Không có tác dụng!
    ```
    
2. **So sánh slice bằng `==`**
    
    ```go
    s1 := []int{1, 2, 3}
    s2 := []int{1, 2, 3}
    // if s1 == s2 {} // ❌ Compile error
    if slices.Equal(s1, s2) {} // ✅ Đúng
    ```
    
3. **Index out of range**
    
    ```go
    s := []int{1, 2, 3}
    fmt.Println(s[5]) // ❌ Panic: runtime error
    ```
    
4. **Không hiểu sub-slice chia sẻ bộ nhớ**
    
    ```go
    original := []int{1, 2, 3, 4, 5}
    sub := original[1:4]
    sub[0] = 999
    // original cũng bị thay đổi!
    ```
    

### 🎯 Khi nào dùng Slice vs Array?

**Dùng Slice khi:**

- Không biết trước kích thước
- Cần thêm/xóa phần tử động
- Làm việc với dữ liệu từ API, file, database
- Hầu hết các trường hợp trong Go

**Dùng Array khi:**

- Biết chính xác kích thước cố định
- Cần hiệu năng tối ưu (ít khi)
- Tạo buffer có kích thước nhỏ

**💡 Khuyên dùng:** Trong Go, **Slice được ưu tiên sử dụng** hơn Array trong hầu hết các trường hợp!