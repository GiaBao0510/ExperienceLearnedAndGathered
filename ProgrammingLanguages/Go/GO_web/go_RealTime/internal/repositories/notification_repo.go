package repositories

import (
	"context"

	"github.com/GiaBao0510/Go-Realtime/internal/model"
)

type INotificationRepository interface {
	Create(ctx context.Context, notification *model.Notification) error
	CreateUserNotification(ctx context.Context, uid string, notificationID int64) error
	GetUnreadByUser(ctx context.Context, uid string) ([]*model.Notification, error)
	MarkAsRead(ctx context.Context, uid string, notificationID int64) error
	CountUnread(ctx context.Context, uid string) (int, error)
}
