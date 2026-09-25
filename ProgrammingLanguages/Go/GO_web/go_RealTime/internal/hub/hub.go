package hub

import (
	"context"
	"log"
	"sync"

	"nhooyr.io/websocket"
)

// Client đại diện cho một kết nối Websocket đang hoạt động
type Client struct {
	ID   string          // ID duy nhất của client
	Conn *websocket.Conn // Kết nối Websocket
	Send chan []byte     // Channel buffer để gửi message (tránh block)
}

// Hub là trung tâm quản lý tất cả client đang kết nối.
//
// Cách hoạt động:
//   - Các goroutine gửi yêu cầu qua channel (register/unregister/broadcast)
//   - Chỉ một goroutine duy nhất (Run) xử lý map clients → không cần lock
//   - Đây là pattern "share memory by communicating" đặc trưng của Go
type Hub struct {
	// Client lưu toàn bộ client đang kết nối đến: key = clientID
	clients map[string]*Client

	// Channel để nhận message cần gửi đến tất cả client
	broadcast chan []byte

	// register nhận yêu cầu đăng ký client mới
	register chan *Client

	// unregister nhận yêu cầu hủy đăng ký client
	unregister chan *Client

	// mu bảo vệ client map khi truy cập từ nhiều goroutine (chỉ cần khi đọc/ghi trực tiếp vào map)
	mu sync.RWMutex
}

// Khởi tạo Hub
func NewHub() *Hub {
	return &Hub{
		clients:    make(map[string]*Client),
		broadcast:  make(chan []byte, 256),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

// Broadcast gửi message đến tất cả client đang kết nối
// Thread-safe: có thể gọi từ bất kỳ goroutine nào
func (h *Hub) Broadcast(msg []byte) {
	h.broadcast <- msg
}

// Register đăng ký một client mới vào hub
func (h *Hub) Register(client *Client) {
	h.register <- client
}

// Unregister hủy đăng ký một client khỏi hub
func (h *Hub) Unregister(client *Client) {
	h.unregister <- client
}

// Trả về số lượng kết nối hiện tại
func (h *Hub) Count() int {
	h.mu.RLock()         // Lock để đọc map clients
	defer h.mu.RUnlock() // Unlock sau khi đọc xong
	return len(h.clients)
}

// Gửi message đến client cụ thể dựa trên ID. Trả về false nếu client không tồn tại.
func (h *Hub) SendToClient(clientID string, msg []byte) bool {
	h.mu.RLock() // Lock để đọc map clients
	client, exists := h.clients[clientID]
	h.mu.RUnlock() // Unlock sau khi đọc xong

	if !exists {
		return false
	}

	// Dùng select để tránh block nếu buffer đầy
	select {
	case client.Send <- msg:
		return true
	default:
		log.Printf("⚠️  [Hub] Không thể gửi đến '%s': buffer đầy", clientID)
		return false
	}
}

// Run là vòng lặp sự kiện chính của hub
// Phải chạy trong goroutine riêng: go hub.Run()
//
// Tại sao không dùng mutex cho clients map?
//   - Chỉ một goroutine duy nhất (Run) thao tác trực tiếp với map clients → không có race condition
//   - Các goroutine khác chỉ gửi yêu cầu qua channel (register/unregister/broadcast) → thread-safe
func (h *Hub) Run() {
	for {
		select {

		// Có client kết nối
		case client := <-h.register:
			h.mu.Lock()                   // Lock để ghi vào map clients
			h.clients[client.ID] = client // Thêm client vào map
			h.mu.Unlock()                 // Unlock sau khi ghi xong
			log.Printf("[Hub] Client '%s' đã kết nối. Tổng kết nối: %d", client.ID, h.Count())

		// Có client ngắt kết nối
		case client := <-h.unregister:
			h.mu.Lock()
			if _, exists := h.clients[client.ID]; exists {
				delete(h.clients, client.ID) // Xóa client khỏi map
				close(client.Send)           // Đóng channel Send để goroutine gửi message của client có thể thoát
			}
			h.mu.Unlock()
			log.Printf("[Hub] Client '%s' đã ngắt kết nối. Tổng kết nối: %d", client.ID, h.Count())

		case message := <-h.broadcast:
			h.mu.RLock()
			for id, client := range h.clients {
				select {
				case client.Send <- message:
					// Gửi thành công vào buffer
				default:
					// Buffer đầy => client bị lag nghiêm trọng, ngắt kết nối
					log.Printf("⚠️  [Hub] Không thể gửi đến '%s': buffer đầy, ngắt kết nối", id)
					close(client.Send)
					delete(h.clients, id)
				}
			}
			h.mu.RUnlock()
		}
	}
}

// WritePump đọc từ channel Send và ghi xuống websocket connection
// chạy trong goroutine riêng cho mỗi client. Nếu channel Send bị đóng, goroutine sẽ thoát.
func WritePump(ctx context.Context, client *Client, cancel context.CancelFunc) {

	defer cancel() // Đảm bảo hủy context khi goroutine kết thúc

	// Vòng lặp gửi message xuống websocket
	for {
		select {
		case msg, ok := <-client.Send:
			if !ok {
				// Channel đã bị đóng (Hub kick hoặc unregister)
				_ = client.Conn.Close(websocket.StatusNormalClosure, "Hub đã ngắt kết nối")
				return
			}

			// Ghi message xuống client với timeout 5 giây
			writeCtx, writeCancel := context.WithTimeout(ctx, 5e9)
			err := client.Conn.Write(writeCtx, websocket.MessageText, msg) // Ghi message xuống websocket
			writeCancel()                                                  // Hủy context sau khi ghi xong

			// Nếu có lỗi khi gửi, log ra và tiếp tục vòng lặp
			if err != nil {
				log.Printf("⚠️  [WritePump] Lỗi khi gửi đến client '%s': %v", client.ID, err)
				return
			}

		case <-ctx.Done():
			// Context bị hủy (có thể do server shutdown hoặc client disconnect)
			return
		}
	}
}
