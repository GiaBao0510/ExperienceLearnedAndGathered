package controller

import (
	"errors"
	"github/GiaBao0510/error-handling/internal/domain"

	"github.com/gin-gonic/gin"
	"net/http"
)

/*
1️⃣  Nhận một AppHandler (hàm controller trả về error)
2️⃣  Thực thi handler – nếu không có lỗi → kết thúc bình thường
3️⃣  Nếu có lỗi → switch-case kiểm tra kiểu lỗi:
     • errors.As(*AppError) → dùng Code và Status từ struct lỗi
     • errors.Is(ErrNotFound) → trả 404
     • errors.Is(ErrUnauthorized) → trả 401
     • default → trả 500 Internal Server Error
4️⃣  Controller KHÔNG BAO GIỜ tự gọi ctx.JSON khi có lỗi
*/

// AppHandler là một hàm controller trả về lỗi thay vì trả về void
type AppHandler func(ctx *gin.Context) error

// Builder - wrapper bắt lỗi tập trung. Đây là ĐIỂM XỬ LÝ DUY NHẤT
func Build(handler AppHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 1. Thực thi logic controller
		err := handler(ctx)
		if err == nil {
			return // Không có lỗi, kết thúc bình thường
		}

		// 2.Kiểm tra xem có phải AppError có cấu trúc không
		var appErr *domain.AppError
		if errors.As(err, &appErr) {
			ctx.JSON(appErr.Code, domain.ErrorResponse{
				Code:    appErr.Code,
				Status:  appErr.Status,
				Message: appErr.Message,
			})
			return
		}

		// 3. Fallback cho các lỗi khác
		switch {
		case errors.Is(err, domain.ErrUnauthorized):
			ctx.JSON(http.StatusUnauthorized, domain.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Status:  "Unauthorized",
				Message: "Không có quyền truy cập - detail: " + err.Error(),
			})
		case errors.Is(err, domain.ErrForbidden):
			ctx.JSON(http.StatusForbidden, domain.ErrorResponse{
				Code:    http.StatusForbidden,
				Status:  "Forbidden",
				Message: "Hành động bị cấm - detail: " + err.Error(),
			})
		case errors.Is(err, domain.ErrNotFound):
			ctx.JSON(http.StatusNotFound, domain.ErrorResponse{
				Code:    http.StatusNotFound,
				Status:  "Not Found",
				Message: "Không tìm thấy tài nguyên - detail: " + err.Error(),
			})
		case errors.Is(err, domain.ErrInvalidFormat):
			ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{
				Code:    http.StatusBadRequest,
				Status:  "Bad Request",
				Message: "Định dạng không hợp lệ - detail: " + err.Error(),
			})
		case errors.Is(err, domain.ErrPhoneDuplicate):
			ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{
				Code:    http.StatusBadRequest,
				Status:  "Bad Request",
				Message: "Số điện thoại đã tồn tại - detail: " + err.Error(),
			})
		default:
			ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Status:  "Internal Server Error",
				Message: "Lỗi máy chủ - detail: " + err.Error(),
			})
		}
	}
}
