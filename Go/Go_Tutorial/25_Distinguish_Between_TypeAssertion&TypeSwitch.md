Type Assertion được sử dụng để truy xuất giá trị cụ thể từ Empty Interface (Interface{})
Nó giúp cho chúng ta có thể chuyển đổi một giá trị từ kiểu Interface{} sang kiểu dữ liệu

Ví dụ:
```go
func main() {

	var i interface{} = "Hello, world!"

	s, ok := i.(string)
	if ok {
		fmt.Println("Giá trị là kiểu String: ", s)
	} else{
		fmt.Println("Giá trị không phải kiểu String")
	}

	f, ok := i.(float32)
	if ok {
		fmt.Println("Giá trị là kiểu Float32: ", f)
	} else{
		fmt.Println("Giá trị không phải kiểu Float32")
	}

}
```

Kết quả: 
```shell
Giá trị là kiểu String:  Hello, world!
Giá trị không phải kiểu Float32
```

---
Type Switch là một cấu trúc điều khiển cho phép kiểm tra kiểu dữ liệu của một Empty Interface (interface{})
Nó giống như switch thông thường nhưng thay vì so sánh giá trị, nó so sánh kiểu dữ liệu