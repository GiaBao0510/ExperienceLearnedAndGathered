package model

import "time"

// MessageType định nghĩa các loại tin nhắn message trong hệ thống
type MessageType string

const (
	TypeChat   MessageType = "chat"   // Tin nhắn chat giữa các client
	TypeSystem MessageType = "system" // Tin nhắn hệ thống, ví dụ thông báo kết nối/disconnect
	TypeStatus MessageType = "status" // Tin nhắn trạng thái, ví dụ thông báo trạng thái người dùng
)

// Message đây là cấu trúc dữ liệu chung cho WebSocket & SSE
type Message struct {
	Type      MessageType `json:"type"`      // Loại tin nhắn
	Content   string      `json:"content"`   // Nội dung tin nhắn
	Sender    string      `json:"sender"`    // Người gửi tin nhắn
	Timestamp time.Time   `json:"timestamp"` // Thời gian gửi tin nhắn
}

// LiveStats là dữ liệu dùng cho SSE demo (dashboard thời gian thực)
type LiveStats struct {
	OnlineUsers int    `json:"online_users"` // Số lượng người dùng đang online
	ServerTime  string `json:"server_time"`  // Thời gian hiện tại của server
	CPULoad     int    `json:"cpu_load"`     // Tải CPU hiện tại của server
	MemoryUsage int    `json:"memory_usage"` // Sử dụng bộ nhớ hiện tại của server
}

// DBMessage — entity tương ứng với bảng `message` trong DB
type DBMessage struct {
	MessageID      int64       `json:"message_id"`      // ID của tin nhắn trong cơ sở dữ liệu
	ConversationID int         `json:"conversation_id"` // ID của cuộc trò chuyện
	SenderUID      *string     `json:"sender"`          // UID của người gửi
	Content        string      `json:"content"`         // Nội dung tin nhắn
	MessageType    MessageType `json:"message_type"`    // Loại tin nhắn
	IsEdited       bool        `json:"is_edited"`       // Trạng thái chỉnh sửa của tin nhắn
	CreatedAt      time.Time   `json:"created_at"`      // Thời gian tạo tin nhắn
}
