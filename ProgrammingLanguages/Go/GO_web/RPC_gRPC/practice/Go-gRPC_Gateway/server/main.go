package main

import (
	proto "Go-gRPC_Gateway/grpc"
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port",9000, "the port to connect")
)

// struct server dùng để implement các phương thức của service OrderService được định nghĩa trong file proto
type server struct {
	proto.UnimplementedOrderServiceServer
}

func(s *server) NewOrder(ctx context.Context, in *proto.NewOrderRequest) (*proto.NewOrderResponse, error) {
	log.Printf("received order: %v", in.GetOrderRequest())
	return &proto.NewOrderResponse{OrderResponse: "new orderId " + in.GetOrderRequest()}, nil
}

func main(){
	// Lắng nghe trên cổng được chỉ định
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Tạo một gRPC server
	s := grpc.NewServer() // Đăng ký server GRPC mới
	proto.RegisterOrderServiceServer(s, &server{}) // Đăng ký service OrderService với server GRPC
	log.Printf("server listening on port: %v",lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}