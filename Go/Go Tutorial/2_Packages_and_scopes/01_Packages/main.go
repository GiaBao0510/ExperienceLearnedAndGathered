package main

import "fmt"

func main() {

	fmt.Println("Hello, World!")

	/*
		- Bạn có thể truy cập các chức năng từ các tệp tin khác nhau
		,cái mà nằm trong cùng gói
		- Tại đây `main()` có thể truy cập đến `listening()` và `loading()`
		- bởi vì cả 3 tệp tin này nằm trong cùng một gói
	*/

	loading()
	listening()
}
