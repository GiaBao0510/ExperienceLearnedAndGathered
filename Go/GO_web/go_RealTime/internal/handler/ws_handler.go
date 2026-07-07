package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/GiaBao0510/Go-Realtime/internal/hub"
	"github.com/GiaBao0510/Go-Realtime/internal/model"
	"nhooyr.io/websocket"
)

// WSHandler xử lý các kết nối WebSocket (2 chiều: Server <-> Client)
type WSHandler struct {
	hub *hub.Hub
}

func NewWSHandler(hub *hub.Hub) *WSHandler {
	return &WSHandler{hub: hub}
}

// incomingMessage là cấu trúc JSON mà client gửi lên
type incomingMessage struct {
	Content string `json:"content"` // Nội dung message
}

// broadcastSystemMsg gửi thông báo lên hệ thống (join/leave) đến tất cả client
func(h *WSHandler) broadcastSystemMsg(content string) {
	msg := model.Message{
		Type: model.TypeSystem,
		Content: content,
		Sender: "system",
		Timestamp: time.Now(),
	}
	msgBytes, _ := json.Marshal(msg)
	h.hub.Broadcast(msgBytes)
}

// ServeHTTP xử lý một kết nối WebSocket hoàn chỉnh
//
// Luồng hoạt động:
//  1. Upgrade HTTP → WebSocket (handshake)
//  2. Tạo Client object và đăng ký vào Hub
//  3. Chạy writePump trong goroutine riêng (Hub → Client)
//  4. Vòng lặp readPump ở goroutine hiện tại (Client → Hub)
//  5. Cleanup khi client ngắt kết nối
func (h *WSHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Bước 1: Upgrade HTTP → WebSocket
	// - websocket.Accept() sẽ thực hiện handshake và trả về *websocket.Conn
	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		// InsecureSkipVerify: Bỏ qua việc kiểm tra chứng chỉ TLS (chỉ nên dùng trong môi trường dev)
		InsecureSkipVerify: true, 
	})
	if err != nil {
		log.Println("[WS] WebSocket Accept error:", err)
		return
	}

	// CloseNow đảm bảo rằng kết nối WebSocket sẽ được đóng khi hàm này kết thúc
	// Dùng closeNow() thay vì Close() để tránh block nếu có goroutine đang đọc/ghi
	defer conn.CloseNow()

	// Bước 2: lấy thông tin client từ query params
	userName := r.URL.Query().Get("username")
	if userName == "" {
		userName = "Anonymous"
	}

	// Tạo ID duy nhất cho client (có thể dùng UUID hoặc timestamp)
	clientID := time.Now().Format("150405") // Dùng timestamp làm ID tạm thời

	// Bước 3: Tạo Client object và đăng ký vào Hub
	client := &hub.Client{
		ID: clientID,
		Conn: conn,
		Send: make(chan []byte,256),
	}
	h.hub.Register(client)

	// Đảm bảo unregister client khi kết thúc
	defer func() {
		h.hub.Unregister(client)
		h.broadcastSystemMsg(userName + "đã rời phòng chat")
	}()

	// Bước 4: Gửi thông báo hệ thống khi client mới tham gia
	h.broadcastSystemMsg(userName + " đã tham gia phòng chat")

	// Bước 5: tạo context để điều phối 2 goroutine readPump và writePump
	// - khi readPump hoặc writePump gặp lỗi -> cancel context -> goroutine còn lại cũng dừng
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	// Bước 6: Chạy writePump trong goroutine riêng (Hub → Client)
	// - writePump sẽ lắng nghe channel client.Send và gửi dữ liệu đến client
	go hub.WritePump(ctx, client, cancel)

	// Bước 7: vòng lặp readPump ở goroutine hiện tại (Client → Hub)
	// Đây là vòng lặp chặn (blocking loop), chạy ở goroutine hiện tại
	// Mỗi khi client gửi message → parse → broadcast lên Hub
	for {
		// Đọc message từ client
		_, rawData, err := conn.Read(ctx)
		if err != nil {
			// Lỗi này có thể do client đóng tab, mất kết nối, server cancel ctx
			log.Printf("[WS] '%s' ngắt kết nối: %v", client.ID, err)
			break
		}

		// Parse Json message từ client
		var incoming incomingMessage
		if err := json.Unmarshal(rawData, &incoming); err != nil {
			log.Printf("[WS] '%s' gửi message không hợp lệ: %v", client.ID, err)
			continue // bỏ qua message này, tiếp tục vòng lặp
		}

		// Bỏ qua message rỗng
		if incoming.Content == "" {
			continue
		}

		log.Printf("[WS] '%s' gửi: %s", userName, incoming.Content)

		// Tạo message chuẩn để broadcast
		msg := model.Message{
			Type: model.TypeChat,
			Content: incoming.Content,
			Sender: userName,
			Timestamp: time.Now(),
		}

		msgBytes, err := json.Marshal(msg)
		if err != nil {
			log.Printf("[WS] Lỗi marshal message: %v", err)
			continue
		}

		h.hub.Broadcast(msgBytes)
	}




}
