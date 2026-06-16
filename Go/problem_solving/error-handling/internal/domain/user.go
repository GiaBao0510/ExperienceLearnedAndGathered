package domain

type User struct {
	UID int64 `json:"uid"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Address string `json:"address"`
}

// CreateUserRequest đại diện cho payload khi tạo người dùng mới
type CreateUserRequest struct {
	Name  string `json:"name" binding:"required,min=1,max=100"`
	Phone string `json:"phone" binding:"required,min=10,max=15"`
	Address string `json:"address" binding:"required"`
}

// UpdateUserRequest đại diện cho payload khi cập nhật thông tin người dùng
type UpdateUserRequest struct {
	Name  string `json:"name" binding:"omitempty,min=1,max=100"`
	Phone string `json:"phone" binding:"omitempty,min=10,max=15"`
	Address string `json:"address" binding:"omitempty"`
}