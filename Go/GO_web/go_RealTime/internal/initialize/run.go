package initialize

import (
	"log"
	"net/http"

	"github.com/GiaBao0510/Go-Realtime/global"
	"github.com/GiaBao0510/Go-Realtime/internal/handler"
	"github.com/GiaBao0510/Go-Realtime/internal/hub"
)

// hàm này để chạy các hàm khởi tạo khác nhau, ví dụ: khởi tạo DB, khởi tạo cache, khởi tạo các service khác
func Run() {
	LoadConfig()     // Load cấu hình từ file config
	InitPostgresql() // Khởi tạo kết nối PostgreSQL

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

	r := InitRouter()                             // Khởi tạo router
	r.Run(":" + global.Config.Server.Server_Port) // Chạy server
}

// Hiển thị thông tin cấu hình đã load
func ShowConfig() {
	log.Println("Cấu hình đã load:")
	log.Printf("DB Host: %s", global.Config.DB.DB_Host)
	log.Printf("DB Port: %d", global.Config.DB.DB_Port)
	log.Printf("DB User: %s", global.Config.DB.DB_User)
	log.Printf("DB Name: %s", global.Config.DB.DB_Name)
	log.Printf("Server Host: %s", global.Config.Server.Server_Host)
	log.Printf("Server Port: %s", global.Config.Server.Server_Port)
}
