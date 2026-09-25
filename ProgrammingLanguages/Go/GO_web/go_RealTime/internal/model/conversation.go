package model

import "time"

type Conversation struct {
	ConversationID int       `json:"conversation_id"`
	Name           *string   `json:"name"`
	Type           string    `json:"type"`
	AvatarURL      *string   `json:"avatar_url"`
	CreatedBy      *string   `json:"created_by"`
	CreateAt       time.Time `json:"created_at"`
}

type Conversation_Member struct {
	ConversationID int       `json:"conversation_id"`
	UID            string    `json:"uid"`
	Role           string    `json:"role"`
	JoinedAt       time.Time `json:"joined_at"`
}
