package main

import (
	"log"
	"net/rpc"
)

func main() {
	// Tạo một client RPC kết nối đến server
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Dial error:", err)
	}

	var reply string

	err = client.Call("Geeting.SayHello", "World", &reply)
	if err != nil {
		log.Fatal("Call error:", err)
	}

	log.Println("Response from server:", reply)
}