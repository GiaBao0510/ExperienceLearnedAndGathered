package model

import "time"

type User struct {
	UID          string     `json:"uid"`
	Name         string     `json:"name"`
	Email        string     `json:"email"`
	PasswordHash string     `json:"-"`
	AvatarURL    *string    `json:"avatar_url"`
	IsOnline     bool       `json:"is_online"`
	LastSeenAt   *time.Time `json:"last_seen_at"`
	CreateAt     time.Time  `json:"created_at"`
}

// Create
type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// Update
type UpdateUserRequest struct {
	Name     *string `json:"name"`
	Email    *string `json:"email" `
	Password *string `json:"password"`
}
