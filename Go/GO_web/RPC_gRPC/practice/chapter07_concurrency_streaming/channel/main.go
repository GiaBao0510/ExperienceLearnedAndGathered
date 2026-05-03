package main

import (
	"fmt"
	"time"
)

/*
	gRPC streaming hoạt động rất giống channel trong Go: một bên gửi dữ liệu liên tục, bên kia
nhận và xử lý.
*/

// Producer gửi dữ liệu vào channel — giống gRPC client streaming
func producer(ch chan<- int, n int) {
	defer close(ch) // Đóng sau khi gửi xong
	for i := 1; i <= n; i++ {
		ch <- i // Gửi dữ liệu vào channel
		fmt.Printf("Producer gửi: %d\n", i)
		time.Sleep(100 * time.Millisecond)
	}
}

// Consumer nhận dữ liệu từ channel — giống gRPC server nhận stream
func consumer(nums <-chan int) {
	sum := 0

	// range tự động dừng khi channel đóng
	for v := range nums {
		sum += v
		fmt.Printf("Consumer nhận: %d, tổng hiện tại: %d\n", v, sum)
	}
	fmt.Printf("Consumer tổng cuối cùng: %d\n", sum)
}

func main() {
	ch := make(chan int) // Tạo channel

	go producer(ch, 10)
	consumer(ch) // Consumer sẽ chờ cho đến khi producer gửi xong và đóng channel
}
