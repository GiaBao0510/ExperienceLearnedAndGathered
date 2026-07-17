package model

import "time"

type Notification struct {
	NotificationID int64     `json:"notification_id"`
	Type           string    `json:"type"`
	Content        string    `json:"content"`
	CreateAt       time.Time `json:"created_at"`
}

type UserNotification struct {
	Uid            string     `json:"uid"`
	NotificationID int64      `json:"notification_id"`
	IsRead         bool       `json:"is_read"`
	ReadAt         *time.Time `json:"read_at"`
}
