## **1. Làm việc với Array:**

Array là tập hợp các giá trị (phần tử) có cùng kiểu dữ liệu với kích thước cố định
Khi một Array được khai báo, kích thước của nó không thể thay đổi

Ví dụ:
```go
func main() {
	// Tạo mảng chưa gán giá trị
	var nums [10]int
	nums[5] = 100 		//Gán giá trị vào một vị trí cụ thể
	fmt.Println(nums)
	fmt.Println("---------------")

	//Mảng chuỗi chua có giá trị
	var chars [10]string
	fmt.Println(chars)
	fmt.Println("---------------")
}
```

kết quả:
```bassh
> go run .
[0 0 0 0 0 100 0 0 0 0]
---------------
[         ]
---------------
```

Lưu ý: Nếu gán giá trị cho mảng tại một vị trí nào đó mà vượt qua giới hạn độ dài của mảng thì nó lập tức báo lỗi

## **Khởi tạo Array**

Cách 1: khởi tạo từng phần tử

Ví dụ:
```go
// Tạo mảng và gán giá trị cho từng phần tử
var nums [5]int
nums[0] = 20
nums[1] = 40
nums[2] = 60
nums[3] = 80
nums[4] = 100 		
fmt.Println(nums)
```

Kết quả:
```bash
> go run .
[20 40 60 80 100]
```

Cách 2: Khởi tạo trực tiếp

ví dụ:
```go
// Tạo mảng và gán giá trị cho từng phần tử
nums := [5]int{20,40,60,80,100}        
fmt.Println(nums)
```

Cách 3: Khởi tạo không biết kích thước

Ví dụ:
```go
nums := [...]int{20,40,60,80,100}      
fmt.Println(nums)
```

---
## **2. Hiểu và Áp dụng mảng đa chiều

