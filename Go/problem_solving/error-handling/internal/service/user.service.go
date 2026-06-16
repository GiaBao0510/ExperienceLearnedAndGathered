package service

import (
	"github/GiaBao0510/error-handling/internal/domain"
	"github/GiaBao0510/error-handling/internal/repository"
)

type UserService struct {
	repo repository.IUserRepository
}

func NewUserSerivce(repo repository.IUserRepository) IUserService {
	return &UserService{repo: repo}
}

// Triển khai các phương thức của IUserService
// Tại đây các lỗi từ repository sẽ được xử lý và chuyển đổi thành lỗi phù hợp với tầng service
func (s *UserService) GetByID(id int64) (*domain.User, error) {
	if id <= 0 {
		return nil, domain.NewBadRequestError("ID phải là số dương")
	}
	return s.repo.GetByID(id)
}

func (s *UserService) GetAll() ([]*domain.User, error) {
	return s.repo.GetAll()
}
func (s *UserService) Create(user *domain.CreateUserRequest) (*domain.User, error) {
	if user.Name == "" {
		return nil, domain.NewBadRequestError("Tên không được để trống")
	}
	if user.Phone == "" || len(user.Phone) < 10 {
		return nil, domain.NewBadRequestError("Số điện thoại không hợp lệ")
	}
	if user.Address == "" {
		return nil, domain.NewBadRequestError("Địa chỉ không được để trống")
	}

	userEntity := &domain.User{
		Name:    user.Name,
		Phone:   user.Phone,
		Address: user.Address,
	}

	return s.repo.Create(userEntity)
}

func (s *UserService) Update(id int64, user *domain.UpdateUserRequest) (*domain.User, error) {
	if id <= 0 {
		return nil, domain.NewBadRequestError("ID phải là số dương")
	}
	return s.repo.Update(id, &domain.User{
		Name:    user.Name,
		Phone:   user.Phone,
		Address: user.Address,
	})
}

// Delete xóa người dùng theo ID
func (s *UserService) Delete(id int64) error {
	if id <= 0 {
		return domain.NewBadRequestError("ID phải là số dương")
	}
	return s.repo.Delete(id)
}
