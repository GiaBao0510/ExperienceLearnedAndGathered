package impl

import(
	"context"
	"database/sql"
	"fmt"
	"time"
	"github.com/GiaBao0510/Go-Realtime/internal/model"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) IUserRepository {
	return &userRepo{db: db}
} 

func(r *userRepo) Create(ctx context.Context, user *model.User) error {
	query := `
		INSERT INTO users(name, email, password_hash, avatar_url)
		VALUES ($1, $2, $3, $4)
		RETURNING uid, created_at
	`

	return r.db.QueryRowContext(ctx, query,
		user.Name, user.Email, user.PasswordHash, user.AvatarURL,
	).Scan(&user.ID, &user.CreatedAt)
}
func(r *userRepo) GetByID(ctx context.Context, id string) (*model.User, error) {

}

func(r *userRepo) Update_Put(ctx context.Context, id string, user *model.User) error
func(r *userRepo) Update_Patch(ctx context.Context, id string, user *model.User) error
func(r *userRepo) UpdateOnlineStatus(ctx context.Context, id string, online bool) error
func(r *userRepo) Delete(ctx context.Context, id string) error
func(r *userRepo) CountOnline(ctx context.Context) (int, error)