package repository

import "github/GiaBao0510/error-handling/internal/domain"

// Tạo Interface cho User Repository
type IUserRepository interface {
	GetByID(id int64) (*domain.User, error)
	GetAll() ([]*domain.User, error)
	Create(user *domain.User) (*domain.User, error)
	Update(id int64, user *domain.User) (*domain.User, error)
	Delete(id int64) error
}
