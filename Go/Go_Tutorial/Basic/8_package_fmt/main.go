package main

import "fmt"

func main() {

	var address string
	fmt.Print("Nhập địa chỉ: ")
	fmt.Scanln(&address) // Đọc toàn bộ dòng

	fmt.Println("Địa chỉ: ", address)
}
