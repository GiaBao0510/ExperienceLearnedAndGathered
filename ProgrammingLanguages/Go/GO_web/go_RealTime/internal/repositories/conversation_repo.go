package repositories

import (
	"context"

	"github.com/GiaBao0510/Go-Realtime/internal/model"
)

type IConversationRepository interface {
	Create(ctx context.Context, conversation *model.Conversation) error
	GetByID(ctx context.Context, id int) (*model.Conversation, error)
	//Update_Put(ctx context.Context, id int, conversation *model.Conversation) error
	//Update_Patch(ctx context.Context, id int, conversation *model.Conversation) error
	//Delete(ctx context.Context, id int) error
	GetByUserID(ctx context.Context, userID string) ([]*model.Conversation, error) // Lấy danh sách cuộc trò chuyện dựa trên ID User
	AddMember(ctx context.Context, member *model.Conversation_Member) error
	RemoveMember(ctx context.Context, conversationID int, userID string) error
	GetMembers(ctx context.Context, conversationID int) ([]*model.Conversation_Member, error)
}
