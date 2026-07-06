package handler

// SSEHandler là handler cho Server-Sent Events (SSE)
type SSEHandler struct{}

// NewSSEHandler tạo một instance mới của SSEHandler
func NewSSEHandler() *SSEHandler {
	return &SSEHandler{}
}

// ServeHTTP là entry point khi client kết nối vào /sse
//
// Luồng hoạt động:
//  1. Set header "Content-Type: text/event-stream" để browser hiểu đây là SSE
//  2. Lấy http.Flusher để có thể đẩy dữ liệu xuống ngay lập tức
//  3. Gửi event định kỳ cho đến khi client ngắt kết nối
//func (h *SSEHandler) ServerHTTP(w http.Re
