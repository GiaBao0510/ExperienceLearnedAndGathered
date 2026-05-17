# Bài Tập Thực Hành: Quản Lý Nhân Viên và Bộ Phận

## 📋 Mục lục

1. [Giới thiệu](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#1-gi%E1%BB%9Bi-thi%E1%BB%87u)
2. [Cấu trúc dự án](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#2-c%E1%BA%A5u-tr%C3%BAc-d%E1%BB%B1-%C3%A1n)
3. [Kiến thức cần nắm](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#3-ki%E1%BA%BFn-th%E1%BB%A9c-c%E1%BA%A7n-n%E1%BA%AFm)
4. [Hướng dẫn từng bước](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#4-h%C6%B0%E1%BB%9Bng-d%E1%BA%ABn-t%E1%BB%ABng-b%C6%B0%E1%BB%9Bc)
5. [Giải thích thuật toán](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#5-gi%E1%BA%A3i-th%C3%ADch-thu%E1%BA%ADt-to%C3%A1n)
6. [Mở rộng và cải tiến](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#6-m%E1%BB%9F-r%E1%BB%99ng-v%C3%A0-c%E1%BA%A3i-ti%E1%BA%BFn)

---

## 1. Giới thiệu

### 🎯 Mục tiêu bài tập

Xây dựng chương trình **quản lý nhân viên và bộ phận** của công ty sử dụng **Array/Slice** trong Go. Bài tập giúp sinh viên:

- ✅ Hiểu và áp dụng **Slice** để lưu trữ dữ liệu động
- ✅ Thực hành **CRUD** (Create, Read, Update, Delete)
- ✅ Làm việc với **Struct** và **Method**
- ✅ Tổ chức code theo **Package** (module hóa)
- ✅ Xử lý **validation** và **tìm kiếm** dữ liệu
- ✅ Hiểu thuật toán **Binary Search** và **Modified Binary Search**

### 🏢 Yêu cầu nghiệp vụ

**Quản lý Bộ phận:**

- Thêm, xóa, sửa, xem danh sách bộ phận
- Mỗi bộ phận có: ID, Tên, Ngày tạo
- Tìm kiếm bộ phận theo ID

**Quản lý Nhân viên:**

- Thêm, xóa, sửa, xem danh sách nhân viên
- Mỗi nhân viên có: ID, Tên, Ngày sinh, Lương, ID Bộ phận
- Kiểm tra ID bộ phận có tồn tại trước khi thêm nhân viên
- Tính lương trung bình của tất cả nhân viên

---

## 2. Cấu trúc dự án

### 📁 Cây thư mục

```
Test/
├── main.go                      # File chính, điểm khởi đầu
├── department/                  # Package quản lý bộ phận
│   ├── department.go           # Model và logic tìm kiếm
│   └── department_service.go   # Các chức năng CRUD
├── staff/                      # Package quản lý nhân viên
│   ├── staff.go                # Model và logic tìm kiếm
│   └── staff_service.go        # Các chức năng CRUD
└── utils/                      # Package tiện ích chung
    └── util.go                 # Hàm đọc input, validate, clear screen
```

### 🔗 Quan hệ giữa các module

```
main.go
   │
   ├─── department (Package)
   │      ├─── department.go
   │      └─── department_service.go
   │
   ├─── staff (Package)
   │      ├─── staff.go
   │      └─── staff_service.go
   │
   └─── utils (Package)
          └─── util.go
```

### 📦 Module Go (go.mod)

Trước khi bắt đầu, khởi tạo module Go:

```bash
cd Test
go mod init hello
```

File `go.mod` sẽ có nội dung:

```
module hello

go 1.21  // hoặc phiên bản Go bạn đang dùng
```

**Lưu ý:** Tên module `hello` được sử dụng trong import:

```go
import "hello/department"
import "hello/staff"
import "hello/utils"
```

---

## 3. Kiến thức cần nắm

### 3.1. Struct (Cấu trúc)

**Struct** là kiểu dữ liệu tự định nghĩa gồm nhiều trường (field):

```go
type Department struct {
    id           int
    name         string
    creationDate string
}
```

### 3.2. Method (Phương thức)

**Method** là hàm gắn liền với Struct:

```go
func (obj Department) GetInfo() string {
    return fmt.Sprintf("id: %d; name: %s", obj.id, obj.name)
}
```

Sử dụng:

```go
dept := Department{id: 1, name: "IT"}
fmt.Println(dept.GetInfo())  // "id: 1; name: IT"
```

### 3.3. Slice - Mảng động

Slice lưu trữ danh sách các đối tượng:

```go
var Departments []Department  // Slice rỗng

// Thêm phần tử
Departments = append(Departments, dept)

// Duyệt slice
for _, dept := range Departments {
    fmt.Println(dept.GetInfo())
}
```

### 3.4. Pointer (Con trỏ)

**Con trỏ** cho phép thay đổi giá trị gốc:

```go
// Trả về pointer
func FindByID(id int) (*Department, bool) {
    return &Departments[0], true
}

// Sử dụng
dept, found := FindByID(1)
if found {
    dept.name = "New Name"  // Thay đổi giá trị gốc trong Departments
}
```

### 3.5. Package (Gói)

**Package** giúp tổ chức code thành các module:

```go
package department  // Khai báo tên package

import "hello/utils"  // Import package khác
```

**Quy tắc:**

- Tên package = tên thư mục
- Biến/hàm viết **hoa chữ cái đầu** → Public (truy cập được từ bên ngoài)
- Biến/hàm viết **thường chữ cái đầu** → Private (chỉ dùng trong package)

Ví dụ:

```go
var Departments []Department  // Public - dùng được từ main.go
var departments []Department  // Private - chỉ dùng trong package
```

---

## 4. Hướng dẫn từng bước

### Bước 1: Tạo Package Utils (Tiện ích)

**File:** `Test/utils/util.go`

```go
package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

// ReadInput: Đọc dữ liệu từ bàn phím
func ReadInput(prompt string) string {
	fmt.Printf("%s ", prompt)
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')  // Đọc đến khi gặp Enter
	return strings.TrimSpace(value)       // Loại bỏ khoảng trắng thừa
}

// ReadNonEmptyInput: Đọc chuỗi không được để trống
func ReadNonEmptyInput(prompt string) string {
	for {
		value := ReadInput(prompt)
		
		if !IsEmpty(value) {
			return value
		}
		
		fmt.Println("❌ Giá trị không được để trống")
	}
}

// IsEmpty: Kiểm tra chuỗi rỗng
func IsEmpty(value string) bool {
	return value == "" || len(value) == 0
}

// GetConvertedInt: Đọc và chuyển đổi sang số nguyên
func GetConvertedInt(prompt string) int {
	for {
		input := ReadInput(prompt)
		
		// Kiểm tra rỗng
		if IsEmpty(input) {
			fmt.Println("❌ Giá trị không được để trống")
			continue
		}
		
		// Chuyển đổi sang int
		value, err := strconv.Atoi(input)
		
		// Kiểm tra hợp lệ và không âm
		if err == nil && value > -1 {
			return value
		}
		
		fmt.Println("❌ Giá trị không hợp lệ hoặc nhỏ hơn 0")
	}
}

// GetConvertedFloat: Đọc và chuyển đổi sang số thực
func GetConvertedFloat(prompt string) float64 {
	for {
		input := ReadInput(prompt)
		
		if IsEmpty(input) {
			fmt.Println("❌ Giá trị không được để trống")
			continue
		}
		
		// Chuyển đổi sang float64
		value, err := strconv.ParseFloat(input, 64)
		
		if err == nil && value > -1 {
			return value
		}
		
		fmt.Println("❌ Giá trị không hợp lệ hoặc nhỏ hơn 0")
	}
}

// ClearScreen: Xóa màn hình console dựa trên hệ điều hành
func ClearScreen() {
	var cmd *exec.Cmd
	
	// Phát hiện hệ điều hành
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:  // Linux, macOS
		cmd = exec.Command("clear")
	}
	
	// Thực thi lệnh
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println("Error clearing screen:", err)
	}
}
```

**📝 Giải thích:**

- `ReadInput()`: Đọc một dòng text từ bàn phím
- `ReadNonEmptyInput()`: Đọc input và bắt buộc không được rỗng
- `GetConvertedInt()`: Đọc input và chuyển thành số nguyên (có validation)
- `GetConvertedFloat()`: Đọc input và chuyển thành số thực
- `ClearScreen()`: Xóa màn hình console (tương thích Windows và Linux/macOS)

---

### Bước 2: Tạo Package Department (Bộ phận)

#### File 1: `Test/department/department.go` (Model & Logic tìm kiếm)

```go
package department

import (
	"fmt"
	"sort"
)

// Department: Cấu trúc dữ liệu bộ phận
type Department struct {
	id           int
	name         string
	creationDate string
}

// Departments: Slice lưu trữ danh sách bộ phận (Public)
var Departments []Department

// GetInfo: Method trả về thông tin bộ phận dạng chuỗi
func (obj Department) GetInfo() string {
	return fmt.Sprintf("id: %d; name: %s; creationDate: %s", 
		obj.id, obj.name, obj.creationDate)
}

// FindByID_Orderly: Tìm kiếm bộ phận theo ID (Binary Search - yêu cầu đã sắp xếp)
// Độ phức tạp: O(log n)
func FindByID_Orderly(id int) (*Department, bool) {
	left, right := 0, len(Departments)-1
	
	// Binary Search
	for left <= right {
		mid := left + (right-left)/2
		
		if Departments[mid].id == id {
			return &Departments[mid], true
		} else if Departments[mid].id < id {
			left = mid + 1  // Tìm bên phải
		} else {
			right = mid - 1  // Tìm bên trái
		}
	}
	
	return nil, false
}

// FindByID_Unordered: Tìm kiếm bộ phận theo ID (Không yêu cầu sắp xếp)
// Độ phức tạp: O(n/2) trong trường hợp tốt, O(n) trong trường hợp xấu
func FindByID_Unordered(id int) (*Department, int, bool) {
	left, right := 0, len(Departments)-1
	
	// Tìm kiếm từ 2 đầu vào giữa
	for left <= right {
		mid := (left + right) / 2
		
		// Kiểm tra 3 vị trí: giữa, trái, phải
		if Departments[mid].id == id {
			return &Departments[mid], mid, true
		} else if Departments[left].id == id {
			return &Departments[left], left, true
		} else if Departments[right].id == id {
			return &Departments[right], right, true
		} else {
			left++   // Thu hẹp từ trái
			right--  // Thu hẹp từ phải
		}
	}
	
	return nil, -1, false
}

// SortDepartment: Sắp xếp bộ phận theo ID tăng dần
func SortDepartment() {
	sort.Slice(Departments, func(i, j int) bool {
		return Departments[i].id < Departments[j].id
	})
}
```

**📝 Giải thích:**

**1. Struct Department:**

- `id`: Mã bộ phận (duy nhất)
- `name`: Tên bộ phận
- `creationDate`: Ngày tạo (định dạng dd/mm/yyyy)

**2. FindByID_Orderly() - Binary Search:**

- Yêu cầu: Slice **phải được sắp xếp** trước
- Độ phức tạp: **O(log n)** - rất nhanh
- Cách hoạt động: Chia đôi dữ liệu mỗi lần tìm kiếm

**3. FindByID_Unordered() - Modified Binary Search:**

- Không yêu cầu sắp xếp
- Kiểm tra 3 vị trí mỗi vòng lặp: giữa, trái, phải
- Thu hẹp từ 2 đầu vào giữa
- Độ phức tạp: **O(n/2)** trung bình

**4. Khi nào dùng hàm nào?**

- `FindByID_Orderly()`: Khi dữ liệu đã sắp xếp, cần tìm kiếm nhanh
- `FindByID_Unordered()`: Khi dữ liệu chưa sắp xếp, cần index để xóa/sửa

---

#### File 2: `Test/department/department_service.go` (CRUD Operations)

```go
package department

import (
	"fmt"
	"hello/utils"
	"time"
)

// AddDepartment: Thêm bộ phận mới
func AddDepartment() {
	fmt.Println("\n### 1. Thêm bộ phận")
	
	// Nhập thông tin
	id := utils.GetConvertedInt("Nhập id: ")
	name := utils.ReadNonEmptyInput("Nhập tên bộ phận: ")
	creationDate := utils.ReadNonEmptyInput("Nhập ngày tạo (dd/mm/yyyy): ")
	
	// Tạo đối tượng Department
	department := Department{
		id:           id,
		name:         name,
		creationDate: creationDate,
	}
	
	// Thêm vào slice
	Departments = append(Departments, department)
	fmt.Printf("✅ Đã thêm bộ phận thành công: %v\n", department.GetInfo())
}

// DeleteDepartment: Xóa bộ phận theo ID
func DeleteDepartment() {
	fmt.Println("\n### 2. Xóa bộ phận")
	id := utils.GetConvertedInt("Nhập mã bộ phận cần xóa: ")
	
	// Tìm bộ phận theo ID
	_, idx, found := FindByID_Unordered(id)
	
	if found {
		// Kỹ thuật xóa: Hoán đổi với phần tử cuối, sau đó cắt slice
		Departments[idx] = Departments[len(Departments)-1]
		Departments = Departments[:len(Departments)-1]
		
		fmt.Println("✅ Đã xóa thành công bộ phận có mã:", id)
	} else {
		fmt.Println("❌ Không tìm thấy bộ phận có mã:", id)
	}
}

// EditDepartment: Chỉnh sửa thông tin bộ phận
func EditDepartment() {
	fmt.Println("\n### 3. Chỉnh sửa bộ phận")
	id := utils.GetConvertedInt("Nhập mã bộ phận cần chỉnh sửa: ")
	
	// Tìm bộ phận
	depart, idx, found := FindByID_Unordered(id)
	
	if found {
		// Cập nhật thông tin
		depart.name = utils.ReadInput("Nhập lại tên bộ phận: ")
		depart.creationDate = time.Now().Local().Format("02/01/2006")
		
		// Ghi lại vào slice
		Departments[idx] = *depart
		
		fmt.Printf("✅ Đã chỉnh sửa thành công bộ phận có mã %d: %s\n", 
			id, depart.GetInfo())
	} else {
		fmt.Printf("❌ Không tìm thấy bộ phận có mã %d\n", id)
	}
}

// SearchDepartment: Tìm kiếm bộ phận theo ID
func SearchDepartment() {
	fmt.Println("\n### 5. Tìm kiếm bộ phận")
	id := utils.GetConvertedInt("Nhập mã bộ phận: ")
	
	department, _, found := FindByID_Unordered(id)
	
	if found {
		fmt.Printf("✅ Bộ phận có id %d: %s\n", id, department.GetInfo())
	} else {
		fmt.Printf("❌ Không tìm thấy bộ phận có id %d\n", id)
	}
}

// CheckDuplicateID: Kiểm tra ID bộ phận có tồn tại
func CheckDuplicateID(id int) bool {
	_, _, found := FindByID_Unordered(id)
	return found
}

// ListDepartment: Hiển thị danh sách tất cả bộ phận
func ListDepartment() {
	fmt.Println("\n### 4. Hiển thị danh sách bộ phận")
	
	if len(Departments) == 0 {
		fmt.Println("❌ Không có bộ phận nào trong danh sách")
	} else {
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		for i, department := range Departments {
			fmt.Printf("[%d] %s\n", i+1, department.GetInfo())
		}
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	}
}

// DepartmentManagement: Menu quản lý bộ phận
func DepartmentManagement() {
	for {
		utils.ClearScreen()
		
		// Menu
		fmt.Println("\n╔════════════════════════════════╗")
		fmt.Println("║    QUẢN LÝ BỘ PHẬN            ║")
		fmt.Println("╚════════════════════════════════╝")
		fmt.Println("1. Thêm bộ phận")
		fmt.Println("2. Xóa bộ phận")
		fmt.Println("3. Sửa bộ phận")
		fmt.Println("4. Hiển thị bộ phận")
		fmt.Println("5. Tìm kiếm bộ phận")
		fmt.Println("0. Quay lại menu chính")
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		
		choice := utils.GetConvertedInt("👉 Chọn chức năng: ")
		
		switch choice {
		case 1:
			AddDepartment()
		case 2:
			DeleteDepartment()
		case 3:
			EditDepartment()
		case 4:
			ListDepartment()
		case 5:
			SearchDepartment()
		case 0:
			fmt.Println("⬅️  Quay lại menu chính")
			return
		default:
			fmt.Println("❌ Lựa chọn không hợp lệ")
		}
		
		utils.ReadInput("\n⏸️  Nhấn Enter để tiếp tục...")
	}
}
```

**📝 Giải thích:**

**1. Kỹ thuật xóa phần tử trong Slice:**

```go
// Bước 1: Hoán đổi với phần tử cuối
Departments[idx] = Departments[len(Departments)-1]

// Bước 2: Cắt bỏ phần tử cuối
Departments = Departments[:len(Departments)-1]
```

Tại sao không dùng `append()` hay `copy()`?

- Cách này **nhanh nhất** (O(1)) nhưng **không giữ thứ tự**
- Phù hợp khi không cần duy trì thứ tự sắp xếp

**2. Format ngày tháng trong Go:**

```go
time.Now().Local().Format("02/01/2006")
```

- `02`: Ngày (đặt 02 để định dạng DD)
- `01`: Tháng (đặt 01 để định dạng MM)
- `2006`: Năm (đặt 2006 để định dạng YYYY)

**3. Pointer trong hàm Edit:**

```go
depart, idx, found := FindByID_Unordered(id)
// depart là *Department (pointer)

depart.name = "New Name"  // Thay đổi giá trị
Departments[idx] = *depart  // Gán lại vào slice
```

---

### Bước 3: Tạo Package Staff (Nhân viên)

#### File 1: `Test/staff/staff.go` (Model & Logic)

```go
package staff

import (
	"fmt"
	"sort"
)

// Staff: Cấu trúc dữ liệu nhân viên
type Staff struct {
	id            int
	name          string
	DayOfBirth    string
	salary        float64
	id_department int
}

// Staffs: Slice lưu trữ danh sách nhân viên
var Staffs []Staff

// GetInfo: Method trả về thông tin nhân viên
func (obj Staff) GetInfo() string {
	return fmt.Sprintf("id: %d; Name: %s; DayOfBirth: %s; Salary: %.2f; id_department: %d",
		obj.id, obj.name, obj.DayOfBirth, obj.salary, obj.id_department)
}

// FindByID_Orderly: Tìm kiếm nhân viên theo ID (Binary Search)
func FindByID_Orderly(id int) (*Staff, bool) {
	left, right := 0, len(Staffs)-1
	
	for left <= right {
		mid := left + (right-left)/2
		
		if Staffs[mid].id == id {
			return &Staffs[mid], true
		} else if Staffs[mid].id < id {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	
	return nil, false
}

// FindByID_Unordered: Tìm kiếm nhân viên theo ID (Không yêu cầu sắp xếp)
func FindByID_Unordered(id int) (*Staff, int, bool) {
	left, right := 0, len(Staffs)-1
	
	for left <= right {
		mid := (left + right) / 2
		
		if Staffs[mid].id == id {
			return &Staffs[mid], mid, true
		} else if Staffs[left].id == id {
			return &Staffs[left], left, true
		} else if Staffs[right].id == id {
			return &Staffs[right], right, true
		} else {
			left++
			right--
		}
	}
	
	return nil, -1, false
}

// SortStaff: Sắp xếp nhân viên theo ID
func SortStaff() {
	sort.Slice(Staffs, func(i, j int) bool {
		return Staffs[i].id < Staffs[j].id
	})
}

// AverageEmployeeSalary: Tính lương trung bình của nhân viên
func AverageEmployeeSalary() float64 {
	if len(Staffs) == 0 {
		fmt.Println("❌ Hiện tại danh sách trống")
		return -1
	}
	
	var total float64 = 0
	for _, e := range Staffs {
		total += e.salary
	}
	
	return total / float64(len(Staffs))
}
```

---

#### File 2: `Test/staff/staff_service.go` (CRUD Operations)

```go
package staff

import (
	"fmt"
	"hello/department"
	"hello/utils"
)

// AddStaff: Thêm nhân viên mới
func AddStaff() {
	fmt.Println("\n### 1. Thêm nhân viên")
	
	id := utils.GetConvertedInt("Nhập id: ")
	name := utils.ReadNonEmptyInput("Nhập tên nhân viên: ")
	DayOfBirth := utils.ReadNonEmptyInput("Nhập ngày sinh (dd/mm/yyyy): ")
	Salary := utils.GetConvertedFloat("Nhập thu nhập của nhân viên: ")
	ID_department := utils.GetConvertedInt("Nhập id bộ phận: ")
	
	// Kiểm tra xem ID bộ phận có tồn tại hay không
	if !department.CheckDuplicateID(ID_department) {
		fmt.Println("❌ ID bộ phận không tồn tại")
		return
	}
	
	// Tạo đối tượng Staff
	emp := Staff{
		id:            id,
		name:          name,
		DayOfBirth:    DayOfBirth,
		salary:        Salary,
		id_department: ID_department,
	}
	
	Staffs = append(Staffs, emp)
	fmt.Printf("✅ Đã thêm nhân viên thành công: %v\n", emp.GetInfo())
}

// DeleteStaff: Xóa nhân viên theo ID
func DeleteStaff() {
	fmt.Println("\n### 2. Xóa nhân viên")
	id := utils.GetConvertedInt("Nhập mã nhân viên cần xóa: ")
	
	_, idx, found := FindByID_Unordered(id)
	
	if found {
		Staffs[idx] = Staffs[len(Staffs)-1]
		Staffs = Staffs[:len(Staffs)-1]
		
		fmt.Println("✅ Đã xóa thành công nhân viên có mã:", id)
	} else {
		fmt.Println("❌ Không tìm thấy nhân viên có mã:", id)
	}
}

// EditStaff: Chỉnh sửa thông tin nhân viên
func EditStaff() {
	fmt.Println("\n### 3. Chỉnh sửa nhân viên")
	id := utils.GetConvertedInt("Nhập mã nhân viên cần chỉnh sửa: ")
	
	obj, idx, found := FindByID_Unordered(id)
	
	if found {
		obj.name = utils.ReadInput("Nhập lại tên nhân viên: ")
		obj.DayOfBirth = utils.ReadInput("Nhập lại ngày sinh (dd/mm/yyyy): ")
		obj.salary = utils.GetConvertedFloat("Nhập lại thu nhập của nhân viên: ")
		obj.id_department = utils.GetConvertedInt("Nhập lại id bộ phận: ")
		
		Staffs[idx] = *obj
		
		fmt.Printf("✅ Đã chỉnh sửa thành công nhân viên có mã %d: %s\n",
			id, obj.GetInfo())
	} else {
		fmt.Printf("❌ Không tìm thấy nhân viên có mã %d\n", id)
	}
}

// SearchStaff: Tìm kiếm nhân viên theo ID
func SearchStaff() {
	fmt.Println("\n### 5. Tìm kiếm nhân viên")
	id := utils.GetConvertedInt("Nhập mã nhân viên: ")
	
	staff, _, found := FindByID_Unordered(id)
	
	if found {
		fmt.Printf("✅ Nhân viên có id %d: %s\n", id, staff.GetInfo())
	} else {
		fmt.Printf("❌ Không tìm thấy nhân viên có id %d\n", id)
	}
}

// ListStaff: Hiển thị danh sách tất cả nhân viên
func ListStaff() {
	fmt.Println("\n### 4. Hiển thị danh sách nhân viên")
	
	if len(Staffs) == 0 {
		fmt.Println("❌ Không có nhân viên nào trong danh sách")
	} else {
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		for i, staff := range Staffs {
			fmt.Printf("[%d] %s\n", i+1, staff.GetInfo())
		}
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	}
}

// CheckDuplicateID: Kiểm tra ID nhân viên có tồn tại
func CheckDuplicateID(id int) bool {
	_, _, found := FindByID_Unordered(id)
	return found
}

// StaffManagement: Menu quản lý nhân viên
func StaffManagement() {
	for {
		utils.ClearScreen()
		
		// Menu
		fmt.Println("\n╔════════════════════════════════╗")
		fmt.Println("║    QUẢN LÝ NHÂN VIÊN          ║")
		fmt.Println("╚════════════════════════════════╝")
		fmt.Println("1. Thêm nhân viên")
		fmt.Println("2. Xóa nhân viên")
		fmt.Println("3. Sửa nhân viên")
		fmt.Println("4. Hiển thị nhân viên")
		fmt.Println("5. Tìm kiếm nhân viên")
		fmt.Println("6. Tính lương trung bình của nhân viên")
		fmt.Println("0. Quay lại menu chính")
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		
		choice := utils.GetConvertedInt("👉 Chọn chức năng: ")
		
		switch choice {
		case 1:
			AddStaff()
		case 2:
			DeleteStaff()
		case 3:
			EditStaff()
		case 4:
			ListStaff()
		case 5:
			SearchStaff()
		case 6:
			avg := AverageEmployeeSalary()
			if avg >= 0 {
				fmt.Printf("💰 Lương trung bình của nhân viên: %.2f VND\n", avg)
			}
		case 0:
			fmt.Println("⬅️  Quay lại menu chính")
			return
		default:
			fmt.Println("❌ Lựa chọn không hợp lệ")
		}
		
		utils.ReadInput("\n⏸️  Nhấn Enter để tiếp tục...")
	}
}
```

**📝 Giải thích:**

**Validation ID bộ phận:**

```go
if !department.CheckDuplicateID(ID_department) {
    fmt.Println("❌ ID bộ phận không tồn tại")
    return
}
```

- Kiểm tra ràng buộc: Nhân viên phải thuộc một bộ phận tồn tại
- Đây là ví dụ về **Foreign Key** trong database

---

### Bước 4: Tạo File Main

**File:** `Test/main.go`

```go
package main

import (
	"fmt"
	"hello/department"
	"hello/staff"
	"hello/utils"
)

func main() {
	for {
		utils.ClearScreen()
		
		// Menu chính
		fmt.Println("\n╔════════════════════════════════════════╗")
		fmt.Println("║  CHƯƠNG TRÌNH QUẢN LÝ NHÂN VIÊN       ║")
		fmt.Println("╚════════════════════════════════════════╝")
		fmt.Println("1. 👥 Quản lý nhân viên")
		fmt.Println("2. 🏢 Quản lý bộ phận")
		fmt.Println("0. 🚪 Thoát")
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		
		choice := utils.GetConvertedInt("👉 Chọn chức năng: ")
		
		switch choice {
		case 1:
			staff.StaffManagement()
		case 2:
			department.DepartmentManagement()
		case 0:
			fmt.Println("\n👋 Cảm ơn bạn đã sử dụng chương trình!")
			fmt.Println("🔚 Thoát chương trình...")
			return
		default:
			fmt.Println("❌ Lựa chọn không hợp lệ")
			utils.ReadInput("\n⏸️  Nhấn Enter để tiếp tục...")
		}
	}
}
```

---

### Bước 5: Chạy chương trình

```bash
# Bước 1: Di chuyển vào thư mục dự án
cd Test

# Bước 2: Khởi tạo module (nếu chưa có)
go mod init hello

# Bước 3: Chạy chương trình
go run main.go

# Hoặc build ra file exe/binary
go build -o employee_management
./employee_management  # Linux/macOS
employee_management.exe  # Windows
```

---

## 5. Giải thích thuật toán

### 5.1. Binary Search (Tìm kiếm nhị phân)

**Điều kiện:** Dữ liệu **phải được sắp xếp** trước.

**Cách hoạt động:**

```
Bước 1: Tìm phần tử giữa
Bước 2: So sánh với giá trị cần tìm
        - Bằng nhau → Tìm thấy
        - Nhỏ hơn → Tìm bên phải
        - Lớn hơn → Tìm bên trái
Bước 3: Lặp lại cho đến khi tìm thấy hoặc hết phần tử
```

**Minh họa:**

Tìm số `7` trong mảng `[1, 3, 5, 7, 9, 11, 13]`:

```
Vòng 1:  [1, 3, 5, |7|, 9, 11, 13]
         left=0, right=6, mid=3
         arr[3] = 7 → Tìm thấy! ✅

Tìm số 11:
Vòng 1:  [1, 3, 5, |7|, 9, 11, 13]
         arr[3]=7 < 11 → Tìm bên phải
         
Vòng 2:  [9, |11|, 13]
         left=4, right=6, mid=5
         arr[5] = 11 → Tìm thấy! ✅
```

**Code:**

```go
func FindByID_Orderly(id int) (*Department, bool) {
    left, right := 0, len(Departments)-1
    
    for left <= right {
        mid := left + (right-left)/2  // Tránh overflow
        
        if Departments[mid].id == id {
            return &Departments[mid], true  // Tìm thấy
        } else if Departments[mid].id < id {
            left = mid + 1  // Tìm bên phải
        } else {
            right = mid - 1  // Tìm bên trái
        }
    }
    
    return nil, false  // Không tìm thấy
}
```

**Độ phức tạp:**

- **O(log n)**: Rất nhanh, mỗi vòng lặp giảm một nửa dữ liệu
- Ví dụ: 1,000,000 phần tử chỉ cần ~20 vòng lặp

---

### 5.2. Modified Binary Search (Không cần sắp xếp)

**Ưu điểm:** Không yêu cầu sắp xếp. **Nhược điểm:** Chậm hơn Binary Search thông thường.

**Cách hoạt động:**

```
Thu hẹp từ 2 đầu vào giữa, kiểm tra 3 vị trí mỗi lần:
- Vị trí giữa (mid)
- Vị trí trái (left)
- Vị trí phải (right)
```

**Minh họa:**

Tìm số `7` trong mảng không sắp xếp `[5, 1, 9, 3, 7, 11, 2]`:

```
Vòng 1:  [|5|, 1, 9, |3|, 7, 11, |2|]
         left=0, mid=3, right=6
         Kiểm tra: 5 ≠ 7, 3 ≠ 7, 2 ≠ 7
         left++, right--
         
Vòng 2:  [|1|, 9, |7|, 11]
         left=1, mid=2, right=5
         arr[2] = 7 → Tìm thấy! ✅
```

**Code:**

```go
func FindByID_Unordered(id int) (*Department, int, bool) {
    left, right := 0, len(Departments)-1
    
    for left <= right {
        mid := (left + right) / 2
        
        // Kiểm tra 3 vị trí
        if Departments[mid].id == id {
            return &Departments[mid], mid, true
        } else if Departments[left].id == id {
            return &Departments[left], left, true
        } else if Departments[right].id == id {
            return &Departments[right], right, true
        } else {
            left++   // Thu hẹp từ trái
            right--  // Thu hẹp từ phải
        }
    }
    
    return nil, -1, false
}
```

**Độ phức tạp:**

- **O(n/2)** trung bình
- **O(n)** trong trường hợp xấu nhất
- Nhanh hơn Linear Search (O(n)) nhưng chậm hơn Binary Search (O(log n))

---

### 5.3. So sánh các thuật toán tìm kiếm

|Thuật toán|Yêu cầu sắp xếp|Độ phức tạp|Tốc độ|Use Case|
|---|---|---|---|---|
|**Linear Search**|Không|O(n)|Chậm|Dữ liệu nhỏ, không sắp xếp|
|**Binary Search**|Có|O(log n)|Rất nhanh|Dữ liệu đã sắp xếp, chỉ tìm kiếm|
|**Modified Binary**|Không|O(n/2)|Trung bình|Dữ liệu không sắp xếp, cần index|

**Ví dụ so sánh với 1,000,000 phần tử:**

|Thuật toán|Số vòng lặp (trung bình)|
|---|---|
|Linear Search|500,000|
|Modified Binary|250,000|
|Binary Search|20|

---

## 6. Mở rộng và cải tiến

### 6.1. Cải tiến hiện tại

**1. Kiểm tra ID trùng lặp khi thêm:**

```go
func AddDepartment() {
    // ... nhập thông tin ...
    
    // Kiểm tra ID đã tồn tại
    if CheckDuplicateID(id) {
        fmt.Println("❌ ID đã tồn tại, vui lòng nhập ID khác")
        return
    }
    
    // Tiếp tục thêm...
}
```

**2. Lưu dữ liệu vào file (Persistence):**

```go
import (
    "encoding/json"
    "os"
)

// Lưu dữ liệu
func SaveDepartments(filename string) error {
    data, err := json.MarshalIndent(Departments, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile(filename, data, 0644)
}

// Đọc dữ liệu
func LoadDepartments(filename string) error {
    data, err := os.ReadFile(filename)
    if err != nil {
        return err
    }
    return json.Unmarshal(data, &Departments)
}
```

**3. Thêm chức năng sắp xếp:**

```go
// Sắp xếp theo tên
func SortByName() {
    sort.Slice(Departments, func(i, j int) bool {
        return Departments[i].name < Departments[j].name
    })
}

// Sắp xếp theo ngày tạo
func SortByDate() {
    sort.Slice(Departments, func(i, j int) bool {
        return Departments[i].creationDate < Departments[j].creationDate
    })
}
```

**4. Thêm chức năng lọc và thống kê:**

```go
// Lọc nhân viên theo bộ phận
func FilterByDepartment(deptID int) []Staff {
    var result []Staff
    for _, s := range Staffs {
        if s.id_department == deptID {
            result = append(result, s)
        }
    }
    return result
}

// Thống kê số nhân viên theo bộ phận
func CountStaffByDepartment() map[int]int {
    count := make(map[int]int)
    for _, s := range Staffs {
        count[s.id_department]++
    }
    return count
}

// Tìm nhân viên có lương cao nhất
func FindHighestSalary() *Staff {
    if len(Staffs) == 0 {
        return nil
    }
    
    maxStaff := &Staffs[0]
    for i := range Staffs {
        if Staffs[i].salary > maxStaff.salary {
            maxStaff = &Staffs[i]
        }
    }
    return maxStaff
}
```

---
## 7. Kiến thức bổ sung

### 7.1. Slice vs Array - Chi tiết hơn

```go
// Array - Kích thước cố định
var arr [5]int = [5]int{1, 2, 3, 4, 5}
// arr[5] = 6  // ❌ Compile error

// Slice - Kích thước động
var slice []int = []int{1, 2, 3, 4, 5}
slice = append(slice, 6)  // ✅ OK
```

**Slice internally:**

```
Slice = Pointer + Length + Capacity

slice := []int{1, 2, 3}
┌─────────┬────────┬──────────┐
│ Pointer │ Len: 3 │ Cap: 3   │
└────┬────┴────────┴──────────┘
     │
     v
┌───┬───┬───┐
│ 1 │ 2 │ 3 │
└───┴───┴───┘
```

### 7.2. Pointer - Khi nào dùng?

**Dùng pointer khi:**

1. Muốn thay đổi giá trị gốc
2. Tránh copy dữ liệu lớn (tiết kiệm bộ nhớ)
3. Trả về reference đến phần tử trong slice

```go
// ❌ Không dùng pointer - Copy toàn bộ
func GetDepartment(id int) Department {
    return Departments[0]  // Copy struct
}

// ✅ Dùng pointer - Chỉ copy địa chỉ
func GetDepartment(id int) *Department {
    return &Departments[0]  // Trả về địa chỉ
}
```

### 7.3. Visibility trong Go

```go
// Public - Viết hoa chữ cái đầu
var Departments []Department  // Truy cập được từ bên ngoài
func AddDepartment() {}       // Truy cập được từ bên ngoài

// Private - Viết thường chữ cái đầu
var departments []Department  // Chỉ dùng trong package
func addDepartment() {}       // Chỉ dùng trong package
```

---

## 8. Tổng kết

### ✅ Những gì đã học

1. **Slice**: Mảng động, thêm/xóa linh hoạt
2. **Struct**: Định nghĩa kiểu dữ liệu riêng
3. **Method**: Hàm gắn với struct
4. **Package**: Tổ chức code theo module
5. **Pointer**: Truyền tham chiếu, tránh copy
6. **Binary Search**: Thuật toán tìm kiếm nhanh
7. **CRUD**: Create, Read, Update, Delete
8. **Validation**: Kiểm tra dữ liệu đầu vào

### 🎯 Kỹ năng đạt được

- ✅ Xây dựng ứng dụng console hoàn chỉnh
- ✅ Quản lý dữ liệu với Slice
- ✅ Tổ chức code theo module
- ✅ Xử lý logic nghiệp vụ phức tạp
- ✅ Áp dụng thuật toán tìm kiếm hiệu quả

### 📚 Tài liệu tham khảo

- [Go Tour](https://go.dev/tour/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Documentation](https://pkg.go.dev/)