package main

import (
	"log"
	"net/http"

	"github.com/GiaBao0510/Go-Realtime/internal/handler"
	"github.com/GiaBao0510/Go-Realtime/internal/hub"
)

func main() {
	// Khởi tạo Hub và chạy event loop của Hub trong goroutine riêng
	// chạy Hub tước khi chạy bất kỳ client kết nối nào
	wsHub := hub.NewHub()
	go wsHub.Run() // Goroutin quản lý tất cả websocket client (register/unregister/broadcast)

	// Đăng ký routes
	mux := http.NewServeMux()

	// Route 1: SSE
	mux.Handle("/sse", handler.NewSSEHandler())

	// Route 2: WebSocket
	mux.Handle("/ws", handler.NewWSHandler(wsHub))

	// Route 3: HTTP API (RESTful)
	mux.Handle("/", http.FileServer(http.Dir("./static")))

	// Chạy server
	addr := ":8080"
	log.Printf("Server đang chạy tại http://localhost%s", addr)
	log.Printf("SSE endpoint: http://localhost%s/sse_demo.html", addr)
	log.Printf("WebSocket endpoint: http://localhost%s/ws_demo.html", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal("Lỗi khởi động server:", err)
	}
}
