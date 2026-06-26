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
