package main

import (
	"fmt"
	"time"
)

type StreamResult struct {
	Data  string
	Error error
}

// serverStreamSimulator mô phỏng server streaming trong gRPC
// Trong gRPC thật: server gọi stream.Send() để gửi từng item

func serverStreamSimulator(query string) <-chan StreamResult {
	ch := make(chan StreamResult)

	go func() {
		defer close(ch)
		// Giả lập tìm kiếm và trả về kết quả từng phần
		items := []string{"result1", "result2", "result3"}
		for _, item := range items {
			time.Sleep(200 * time.Millisecond) // Giả lập delay xử lý
			ch <- StreamResult{Data: fmt.Sprintf("{%s} %s", query, item)}
		}

	}()
	return ch
}

func main() {
	fmt.Println("Bắt đầu nhận stream từ server...")

	// Client nhận từng item ngay khi server gửi
	// Trong gRPC thật: client gọi stream.Recv() trong vòng lặp
	for result := range serverStreamSimulator("Go gRPC") {
		if result.Error != nil {
			fmt.Printf("Lỗi nhận stream: %v\n", result.Error)
			break
		}
		fmt.Printf("Nhận từ server: %s\n", result.Data)
	}
	fmt.Println("Đã nhận hết stream từ server.")
}
