package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/GiaBao0510/Go-Realtime/internal/model"
)

// SSEHandler xử lý các kết nối SSE (Server-Sent Events) - 1 chiều từ Server -> Client
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
func (h *SSEHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// >>>>>> Bước 1: Thiết lập SSE header >>>>>>>>
	w.Header().Set("Content-Type", "text/event-stream") // báo browser đây là SSE stream
	w.Header().Set("Cache-Control", "no-cache")         // không cache, luôn nhận data mới nhất
	w.Header().Set("Connection", "keep-alive")          // giữ kết nối HTTP không đóng
	w.Header().Set("X-Accel-Buffering", "no")           // tắt buffer của nginx (nếu có)
	w.Header().Set("Access-Control-Allow-Origin", "*")  // cho phép CORS từ mọi nguồn

	// >>>>>> Bước 2: Lấy Flusher
	// http.Flusher là interface cho phép đẩy dữ liệu xuống client ngay lập tức mà không cần chờ buffer đầy
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Server không hỗ trợ SSE(thiếu Flusher)", http.StatusInternalServerError)
		return
	}

	clientIP := r.RemoteAddr // Lấy IP của client
	log.Printf("[SSEHandler] Client '%s' đã kết nối", clientIP)

	// Bước 3: Gửi event "connected" ngay khi client kết nối
	fmt.Fprintf(w, "event: connected\n")
	fmt.Fprintf(w, "data: {\"message\": \"Kết nối SSE thành công!\", \"time\": \"%s\"}\n\n",
		time.Now().Format("15:04:05"))
	flusher.Flush() // Đẩy dữ liệu xuống client ngay lập tức

	// >>>>>> Bước 4: Gửi event định kỳ mỗi 5 giây >>>>>>>>
	ctx := r.Context() // Lấy context của request để biết khi nào client ngắt kết nối
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop() // Dừng ticker khi kết thúc

	// eventID để client biết là nhận event cụ thể nào
	eventID := 1
	baseOnline := 100 + rand.Intn(50)

	for {
		select {

		// Nếu client ngắt kết nối, context sẽ Done
		case <-ctx.Done():
			log.Printf("[SSE] Client ngắt kết nối: %s", clientIP)
			return

		// Mỗi 5 giây gửi event "update"
		case t := <-ticker.C:
			status := model.LiveStats{
				OnlineUsers: baseOnline + (t.Second() % 20) - 10,
				ServerTime:  t.Format("15:04:05"),
				CPULoad:     20 + rand.Intn(60),
				MemoryUsage: 40 + rand.Intn(40),
			}

			data, err := json.Marshal(status)
			if err != nil {
				log.Printf("[SSE] lỗi marshal Json: %v\n", err)
				continue
			}

			// Ghi SSE event với đầy đủ các field
			fmt.Fprintf(w, "id: %d\n", eventID)                         // ID để reconnect từ đúng chỗ
			fmt.Fprintf(w, "event: %s\n", model.SSE_Event_Stats_Update) // tên event (client lọc theo tại đây)
			fmt.Fprintf(w, "data: %s\n", data)                          // nội dung event (dạng JSON)
			fmt.Fprintf(w, "\n")                                        // kết thúc event
			flusher.Flush()                                             // Đẩy dữ liệu xuống client ngay lập tức

			log.Printf("[SSE] đã push event #%d đến %s\n", eventID, clientIP)
			eventID++
		}
	}
}
