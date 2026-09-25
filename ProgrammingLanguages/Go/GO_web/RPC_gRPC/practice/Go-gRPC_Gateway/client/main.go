package main

import (
	proto "Go-gRPC_Gateway/grpc"
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	//
	address := "localhost:9000"
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))	// Tạo kết nối đến server gRPC mà không sử dụng TLS (insecure)

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := proto.NewOrderServiceClient(conn)	// Tạo một client gRPC từ kết nối đã tạo, client này sẽ được sử dụng để gọi các phương thức của service OrderService

	ticker := time.NewTicker(2 * time.Second)	// Tạo một ticker để gửi yêu cầu mỗi 2 giây
	defer ticker.Stop()

	orderId := 1001
	for range ticker.C { // Đây là cách tiện lợi để tạo vòng lặp vô hạn, mỗi 2 giây sẽ gửi một yêu cầu mới
		 
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		r, err := c.NewOrder(ctx, &proto.NewOrderRequest{OrderRequest: fmt.Sprintf("%d", orderId)}) // Gửi yêu cầu NewOrder với orderId hiện tại
		orderId++ // Tăng orderId lên 1 cho lần gửi tiếp theo

		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Order Response: %s", r.GetOrderResponse())
		cancel() // Hủy context sau khi sử dụng xong để tránh rò rỉ bộ nhớ
	}
}