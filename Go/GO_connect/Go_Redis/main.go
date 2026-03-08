package main

import (
	"bufio"
	"fmt"
	//"log"
	"os"
	//"time"
)

//Tạo reader toàn cục để dùng chung
var reader = bufio.NewReader(os.Stdin)

func main(){

	//Khởi tạo Redis Client
	InitRedis()

	for{
		//Chọn kiểu trong Redis để thao tác
		fmt.Println("Chọn kiểu dữ liệu Redis để thao tác:")
		fmt.Println("1. String")
		fmt.Println("2. Hash")
		fmt.Println("3. List")
		fmt.Println("4. Set")
		fmt.Println("5. Sorted Set")
		fmt.Println("0. Để thoát ")
		fmt.Print("Lựa chọn của bạn: ")
		
		choice,_ := reader.ReadByte()
		reader.ReadString('\n') // Đọc bỏ ký tự newline sau khi đọc byte

		switch choice {
			case '1':
				Menu_STRING()
			case '0':
				return


			default:
				fmt.Println("Lựa chọn không hợp lệ, vui lòng chọn lại.")
		}
	}

	
}