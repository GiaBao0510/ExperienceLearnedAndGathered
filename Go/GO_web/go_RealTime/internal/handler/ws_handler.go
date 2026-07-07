package handler

import (
	"net/http"

	"github.com/GiaBao0510/Go-Realtime/internal/hub"
)

// WSHandler xử lý các kết nối WebSocket (2 chiều: Server <-> Client)
type WSHanlder struct {
	hub *hub.Hub
}

func NewWSHanlder(hub *hub.Hub) *WSHanlder {
	return &WSHanlder{hub: hub}
}

// incomingMessage là cấu trúc JSON mà client gửi lên
type incomingMessage struct {
	Content string `json:"content"` // Nội dung message
}

// ServeHTTP xử lý một kết nối WebSocket hoàn chỉnh
//
// Luồng hoạt động:
//  1. Upgrade HTTP → WebSocket (handshake)
//  2. Tạo Client object và đăng ký vào Hub
//  3. Chạy writePump trong goroutine riêng (Hub → Client)
//  4. Vòng lặp readPump ở goroutine hiện tại (Client → Hub)
//  5. Cleanup khi client ngắt kết nối
func (h *WSHanlder) ServerHTTP(w http.ResponseWriter, r *http.Request) {

}
