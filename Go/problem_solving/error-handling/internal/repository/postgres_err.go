package repository

import (
	"errors"
	"github/GiaBao0510/error-handling/internal/domain"

	"github.com/lib/pq"
)

// HandlePQError xử lý lỗi từ PostgreSQL
func HandlePQError(err error) error {
	var pqErr *pq.Error

	// Sử dụng errors.As để kiểm tra và chuyển đổi lỗi thành *pq.Error
	if errors.As(err, &pqErr) {
		switch pqErr.Code {
		case "23505":
			return domain.NewConflictError("Dữ liệu đã tồn tại")
		case "23503":
			return domain.NewNotFoundError("Tài nguyên không tồn tại")
		case "23502":
			return domain.NewBadRequestError("Thiếu thông tin bắt buộc")
		case "22001":
			return domain.NewBadRequestError("Dữ liệu không hợp lệ")
		case "23514":
			return domain.NewBadRequestError("Ràng buộc dữ liệu bị vi phạm")
		case "22P02":
			return domain.NewBadRequestError("Định dạng dữ liệu không hợp lệ")
		case "40001":
			return domain.NewConflictError("Vi phạm ràng buộc CHECK. Dữ liệu gửi lên không thỏa mãn điều kiện logic của cột")
		case "28P01":
			return domain.NewUnauthorizedError()
		case "08006":
			return domain.NewDatabaseError(errors.New("Kết nối đến cơ sở dữ liệu thất bại"))
		case "42P01":
			return domain.NewDatabaseError(errors.New("Bảng không tồn tại"))
		}
	}
	return domain.NewDatabaseError(err)
}
