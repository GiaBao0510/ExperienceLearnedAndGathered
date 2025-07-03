package main

import (
	"fmt"
)

// Được khai báo bên ngoài hàm, thì bên trong hàm khác có thể được gọi
var address = "Q.Go Vap, tpHCM"

// Khai báo nhiều biến toàn cục cùng lúc
var (
	monHoc string
	diemSo int
)

const Pi = "3.14" // Toàn cục, type inference: string

func main() {
	const Pi float32 = 3.14                  // Cục bộ, che giấu hằng số toàn cục
	fmt.Println("Hằng số Pi trong hàm:", Pi) // In: 3.14
	fmt.Println("Hằng số Pi toàn cục:", Pi)  // Lỗi nếu cố truy cập trực tiếp
}
