package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

func main() {

	//Lắng nghe trên cổng 1234, chờ kết nối từ server
	client, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("ListenTCP error: ", err)
	}

	clientChan := make(chan *rpc.Client)

	go func() {
		for {
			conn, err := client.Accept()
			if err != nil {
				log.Fatal("Accept error: ", err)
			}

			/*
					Khi mỗi đường link được thiết lập, đối tượng RPC Client
				được khởi tạo trên link đó và được gửi vào channel clientChan để xử lý tiếp
			*/
			clientChan <- rpc.NewClient(conn)
		}
	}()

	doClientWork(clientChan)
}

func doClientWork(clientChan <-chan *rpc.Client) {
	//Nhận vào đối tượng RPC Client từ channel
	client := <-clientChan

	//ĐÓng kết nối với client trước khi hàm exit
	defer client.Close()

	var reply string

	//Thực hiện lời gọi bình thường
	err := client.Call("GreetingService.Hello", " Gia Bao", &reply)
	if err != nil {
		log.Fatal("Call error: ", err)
	}

	fmt.Println("Response: ", reply)
}
