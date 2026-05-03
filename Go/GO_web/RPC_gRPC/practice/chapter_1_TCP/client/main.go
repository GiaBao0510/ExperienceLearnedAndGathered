package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	// Thực hiện kết nối đến server TCP tại địa chỉ localhost:8080
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Không thể kết nối đến server: ", err)
	}
	defer conn.Close()

	// Gửi một mảng tin nhắn đến server, phải có ký tự \n để server có thể nhận biết kết thúc tin nhắn
	message := []string{
		"Xin chào, Server!\n",
		"Đây là client gửi tin nhắn.\n",
		"Xin vui lòng xử lý yêu cầu của tôi.\n",
		"Cảm ơn bạn đã lắng nghe!\n",
	}

	reader := bufio.NewReader(conn)

	for _, msg := range message {
		// Gửi tin nhắn đến server
		_, err := conn.Write([]byte(msg))
		if err != nil {
			log.Println("Lỗi khi gửi tin nhắn: ", err)
			return
		}

		// Đọc phản hồi từ server
		response, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Println("Server đã đóng kết nối.")
				return
			}
			log.Println("Lỗi khi đọc phản hồi: ", err)
			return
		}

		fmt.Printf("Phản hồi từ server: %s", response)
	}
}
