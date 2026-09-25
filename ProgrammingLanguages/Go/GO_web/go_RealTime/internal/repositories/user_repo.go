package repositories

import (
	"context"

	"github.com/GiaBao0510/Go-Realtime/internal/model"
)

type IUserRepository interface {
	Create(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, id string) (*model.User, error)
	Update_Put(ctx context.Context, id string, user *model.User) error
	Update_Patch(ctx context.Context, id string, user *model.User) error
	UpdateOnlineStatus(ctx context.Context, id string, online bool) error
	Delete(ctx context.Context, id string) error
	CountOnline(ctx context.Context) (int, error)
}
