/*
	Trong gRPC streaming, bạn cần quản lý vòng đời (lifecycle) của stream: khi nào bắt đầu, khi

nào kết thúc, khi nào hủy. Go's context.Context là công cụ cho việc này
*/
package main

import (
	"context"
	"fmt"
	"time"
)

// streamWithContext mô phỏng streaming có timeout
// Trong gRPC: mọi stream đều nhận ctx context.Context
func streamWithContext(ctx context.Context, datachan chan<- int ) {
	defer close(datachan)

	for i := 0; ; i ++ {
		select{
		case <- ctx.Done():
			// Context đã bị hủy hoặc timeout
			fmt.Println("Stream bị hủy: ",ctx.Err())
			return
		case datachan <- i:
			// Gửi dữ liệu vào channel
			fmt.Println("Gửi dữ liệu: ", i)
			time.Sleep(500 * time.Millisecond) // Giả lập thời gian gửi dữ liệu
		}
	}
}

func main() {
	// Tạo context với timeout 1 giây
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel() // Đảm bảo hủy context khi xong

	datachan := make(chan int, 1)
	go streamWithContext(ctx, datachan)

	for v := range datachan {
		fmt.Println("Nhận dữ liệu: ", v)
	}
}