package repositories

import (
	"context"

	"github.com/GiaBao0510/Go-Realtime/internal/model"
)

type IMessageRepository interface {
	Create(ctx context.Context, msg *model.DBMessage) error
	GetByConversation(ctx context.Context, conversationID int, limit, offset int) ([]model.DBMessage, error)
	GetLatestByConversation(ctx context.Context, conversationID int) (*model.DBMessage, error)
}
