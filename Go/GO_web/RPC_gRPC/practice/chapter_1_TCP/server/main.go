package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {

	// Lắng nghe kết nối trên cổng 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Không thể mở port: ", err)
	}
	defer listener.Close()

	fmt.Println("Server đang lắng nghe trên port 8080...")

	// Vòng lặp để chấp nhận kết nối từ client
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Lỗi khi chấp nhận kết nối: ", err)
			continue
		}

		// Mỗi client được xử lý trong một goroutine riêng
		// Nhờ đó server có thể xử lý nhiều client cùng lúc
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Client đã kết nối: ", conn.RemoteAddr())

	// Đọc dữ liệu từ client
	reader := bufio.NewReader(conn)

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Println("Client đã đóng kết nối: ", conn.RemoteAddr())
				return
			}
			log.Println("Lỗi khi đọc dữ liệu: ", err)
			return
		}

		fmt.Println("Tin nhắn từ client: ", msg)

		// Gửi phản hồi lại cho client
		res := "Server đã nhận được: " + msg
		_, err = conn.Write([]byte(res))
		if err != nil {
			log.Println("Lỗi khi gửi phản hồi: ", err)
			return
		}
	}
	
}
