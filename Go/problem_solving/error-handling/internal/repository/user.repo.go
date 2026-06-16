package repository

import (
	"database/sql"
	"errors"
	"github/GiaBao0510/error-handling/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) IUserRepository {
	return &UserRepository{db: db}
}

// Implement các phương thức của IUserRepository

// GetByID lấy thông tin người dùng theo ID
func (r *UserRepository) GetByID(id int64) (*domain.User, error) {
	query := `SELECT * FROM users WHERE uid = $1`

	user := &domain.User{}
	err := r.db.QueryRow(query, id).Scan(&user.UID, &user.Name, &user.Phone, &user.Address)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.NewNotFoundError("User not found")
		}
		return nil, domain.NewDatabaseError(err)
	}

	return user, nil
}

// GetAll lấy tất cả người dùng
func (r *UserRepository) GetAll() ([]*domain.User, error) {
	query := `SELECT * FROM users`
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, domain.NewDatabaseError(err)
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		user := &domain.User{}
		err := rows.Scan(&user.UID, &user.Name, &user.Phone, &user.Address)
		if err != nil {
			return nil, domain.NewDatabaseError(err)
		}
		users = append(users, user)
	}
	return users, nil
}

// Create tạo mới người dùng
func (r *UserRepository) Create(user *domain.User) (*domain.User, error) {
	query := `INSERT INTO users(name, phone, address) VALUES($1, $2, $3) RETURNING uid`

	err := r.db.QueryRow(query, user.Name, user.Phone, user.Address).Scan(&user.UID)

	if err != nil {
		return nil, HandlePQError(err)
	}

	return user, nil
}

// Update cập nhật thông tin người dùng
func (r *UserRepository) Update(id int64, user *domain.User) (*domain.User, error) {
	query := `UPDATE users SET name = $1, phone = $2, address = $3 WHERE uid = $4 RETURNING *`

	err := r.db.QueryRow(query, user.Name, user.Phone, user.Address, id).Scan(&user.UID, &user.Name, &user.Phone, &user.Address)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.NewNotFoundError("User not found")
		}
		return nil, HandlePQError(err)
	}

	return user, nil
}

// Delete xóa người dùng theo ID
func (r *UserRepository) Delete(id int64) error {
	query := `DELETE FROM users WHERE uid = $1`

	result, err := r.db.Exec(query, id)

	if err != nil {
		return domain.NewDatabaseError(err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return domain.NewNotFoundError("User not found")
	}

	return nil
}
