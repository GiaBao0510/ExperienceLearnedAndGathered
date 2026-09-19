package main

import (
	"log"
	"net"
	"net/rpc"
)

type GeetingService struct{}

func (p *GeetingService) SayHello(req, res *string) error {

	*res = "hello " + *req
	return nil
}

func main() {
	// Đăng ký service với tên là "Geeting" có kiểu là *GeetingService
	rpc.RegisterName("Geeting", new(GeetingService))

	// Tạo server lắng nghe trên cổng 1234
	listener, err := net.Listen("tcp",":1234")
	if err != nil {
		log.Fatal("Listen error:", err)
	}

	log.Println("Server is listening on port 1234...")

	for{
		// Chấp nhận kết nối từ client
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Accept error:", err)
			continue
		}

		// phục vụ kết nối RPC
		go rpc.ServeConn(conn)
	}
}