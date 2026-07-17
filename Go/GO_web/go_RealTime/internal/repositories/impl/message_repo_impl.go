package impl

import (
	"database/sql"

	"github.com/GiaBao0510/Go-Realtime/internal/repositories"
)

type messageRepo struct {
	db *sql.DB
}

func NewMessageRepo(db *sql.DB) repositories.IMessageRepository {
	return &messageRepo{db: db}
}
