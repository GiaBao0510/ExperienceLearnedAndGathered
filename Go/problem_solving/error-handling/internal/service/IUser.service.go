package service

import "github/GiaBao0510/error-handling/internal/domain"

type IUserService interface {
	GetByID(id int64) (*domain.User, error)
	GetAll() ([]*domain.User, error)
	Create(user *domain.CreateUserRequest) (*domain.User, error)
	Update(id int64, user *domain.UpdateUserRequest) (*domain.User, error)
	Delete(id int64) error
}