package model

import "time"

// Struct API key ánh xạ với bảng trong DB
type APIKey struct {
	ID         string     `json:"id"`
	UserID     int        `json:"user_id"`
	Name       string     `json:"name"`
	KeyHash    string     `json:"key_hash"`               // Không trả về trong JSON
	KeyHint    string     `json:"key_hint"`               // 10 ký tự đầu tiên để nhận dạng
	Scopes     []string   `json:"scopes"`                 // Các quyền của API key, có thể là JSON hoặc CSV
	RateLimit  int        `json:"rate_limit"`             // Số lượng yêu cầu tối đa trong một khoảng thời gian
	IsActive   bool       `json:"is_active"`              // Trạng thái của API key
	LastUsedAt *time.Time `json:"last_used_at,omitempty"` // Thời gian lần cuối sử dụng
	CreatedAt  time.Time  `json:"created_at"`             // Thời gian tạo
	ExpiresAt  *time.Time `json:"expires_at,omitempty"`   // Thời gian hết hạn
}

// Struct Response trả về khi tạo API key mới
type CreateAPIKeyRequest struct {
	UserID    int      `json:"user_id" binding:"required"`
	Name      string   `json:"name" binding:"required"`
	Scopes    []string `json:"scopes" binding:"required"`     // Các quyền của API key, có thể là JSON hoặc CSV
	RateLimit int      `json:"rate_limit" binding:"required"` // Số lượng yêu cầu tối đa trong một khoảng thời gian
}

// Struct Response trả về khi tạo API key mới
type CreateAPIKeyResponse struct {
	APIKey   APIKey `json:"api_key"`
	PlainKey string `json:"plain_key"` // Chỉ trả về 1 lần duy nhất !!!, không lưu trong DB
	Message  string `json:"message"`
}
