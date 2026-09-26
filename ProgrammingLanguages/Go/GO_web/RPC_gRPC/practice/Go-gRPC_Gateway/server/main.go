package main

import (
	proto "Go-gRPC_Gateway/grpc/task"
	"context"
	"flag"
	"fmt"
	"net"
	"log"
	"sync"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port",50051, "the port to connect")
)

// struct server dùng để implement các phương thức của service OrderService được định nghĩa trong file proto
type taskServer struct {
	proto.UnimplementedTaskServiceServer
	mu sync.Mutex							
	tasks []*proto.Task
	counter int
}

// Hàm này để tạo một task mới, nhận vào context và request từ client, trả về response và error nếu có
func(s *taskServer) CreateTask(ctx context.Context, req *proto.CreateTaskRequest) (*proto.CreateTaskResponse, error) {
	s.mu.Lock() // Khóa mutex để đảm bảo an toàn khi truy cập vào slice tasks
	defer s.mu.Unlock() // Giải phóng mutex khi hàm kết thúc
	s.counter++

	id := fmt.Sprintf("task-%d",s.counter)

	task := &proto.Task{
		Id: id,
		Title: req.Title,
		Description: req.Description,
	}

	log.Printf("Created task: %v", task)

	return &proto.CreateTaskResponse{
		Id: id,
		Title: req.Title,
		Description: req.Description,
	}, nil
}

// Hàm này sẽ trả về danh sách tất cả các task hiện có, nhận vào context và request từ client, trả về response và error nếu có
func(s *taskServer) ListTasks(req *proto.ListTaskRequest, stream proto.TaskService_ListTasksServer) error {
	s.mu.Lock() // Khóa mutex để đảm bảo an toàn khi truy cập vào slice tasks
	defer s.mu.Unlock() // Giải phóng mutex khi hàm kết thúc

	for _, task := range s.tasks{
		if err := stream.Send(task); err != nil {
			return err
		}
	}

	return nil
}

// Hàm main
func main(){

	//Lắng nghe 
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Tạo một instance của server
	s := grpc.NewServer()
	proto.RegisterTaskServiceServer(s, &taskServer{})

	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

