package main

import (
	proto "Go-gRPC_Gateway/grpc/task"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	//
	address := "localhost:50051" // Địa chỉ của server gRPC mà client sẽ kết nối đến
	conn, err := grpc.NewClient(address, 
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)	// Tạo kết nối đến server gRPC mà không sử dụng TLS (insecure)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close() // Đảm bảo đóng kết nối khi hàm main kết thúc

	client := proto.NewTaskServiceClient(conn) // Tạo một client gRPC từ kết nối đã tạo
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel() // Đảm bảo hủy context khi hàm main kết thúc

	task := []struct{
		title string
		description string
	}{
		{"Task 1", "Description for Task 1"},
		{"Task 2", "Description for Task 2"},
		{"Task 3", "Description for Task 3"},
	}

	for _, t := range task {
		resp, err := client.CreateTask(ctx, &proto.CreateTaskRequest{
			Title: t.title,
			Description: t.description,
		})

		if err != nil {
			log.Fatalf("could not create task: %v", err)
		}
		log.Printf("Created task: %s - %s", resp.Id, resp.Title)
	}
}