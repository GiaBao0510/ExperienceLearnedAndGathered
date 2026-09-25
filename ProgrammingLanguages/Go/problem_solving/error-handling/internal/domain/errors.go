package domain

import (
	"errors"
	"net/http"
)

// Định nghĩa các lỗi phổ biến trong ứng dụng
var (
	// 400 - Bad Request
	ErrInvalidInput   = errors.New("Invalid input")
	ErrMissingField   = errors.New("Missing required field")
	ErrInvalidFormat  = errors.New("Invalid format")
	ErrInvalidValue   = errors.New("Invalid value")
	ErrPhoneDuplicate = errors.New("Phone number already exists")
	ErrEmailDuplicate = errors.New("Email already exists")

	// 401 - Unauthorized
	ErrUnauthorized = errors.New("Unauthorized")
	ErrTokenExpired = errors.New("Token expired")
	ErrTokenInvalid = errors.New("Invalid token")
	ErrTokenMissing = errors.New("Token missing")

	// 403 - Forbidden
	ErrForbidden    = errors.New("Forbidden")
	ErrAccessDenied = errors.New("Access denied")

	// 404 - Not Found
	ErrNotFound     = errors.New("Resource not found")
	ErrUserNotFound = errors.New("User not found")

	// 409 - Conflict
	ErrConflict = errors.New("Conflict")

	// 500 - Internal Server Error
	ErrInternalServer  = errors.New("Internal server error")
	ErrDatabase        = errors.New("Database error")
	ErrDatabaseTimeout = errors.New("Database timeout")
)

/* AppError bọc lỗi với mã lỗi và thông điệp chi tiết. ErrorCatalog là bảng tra cứu lỗi */
type AppError struct {
	Code    int    `json:"code"`    // HTTP status code
	ErrKey  error  `json:"-"`       // sentinal error - dùng cho error.Is()
	Message string `json:"message"` // chi tiết lỗi
	Status  string `json:"status"`  // thông điệp dành cho nhà phát triển
}

func (e *AppError) Error() string { return e.Message }
func (e *AppError) Unwrap() error { return e.ErrKey }

// ══════════════════════════════════════════════════════════
// CONSTRUCTOR FUNCTIONS – Tạo lỗi có cấu trúc
// ══════════════════════════════════════════════════════════
// -------- 400 Bad Request --------
func NewBadRequestError(message string) *AppError {
	return &AppError{
		Code:    http.StatusBadRequest,
		Message: message,
		ErrKey:  ErrInvalidInput,
		Status:  "Bad Request",
	}
}

func NewMissingFieldError(field string) *AppError {
	return &AppError{
		Code:    http.StatusBadRequest,
		Message: "Thiếu trường bắt buộc: " + field,
		ErrKey:  ErrMissingField,
		Status:  "Bad Request",
	}
}

func NewInvalidFormatError(field string) *AppError {
	return &AppError{
		Code:    http.StatusBadRequest,
		Message: "Định dạng không hợp lệ cho trường: " + field,
		ErrKey:  ErrInvalidFormat,
		Status:  "Bad Request",
	}
}

func NewInvalidValueError(field string) *AppError {
	return &AppError{
		Code:    http.StatusBadRequest,
		Message: "Giá trị không hợp lệ cho trường: " + field,
		ErrKey:  ErrInvalidValue,
		Status:  "Bad Request",
	}
}

func NewPhoneDuplicateError() *AppError {
	return &AppError{
		Code:    http.StatusBadRequest,
		Message: "Số điện thoại đã tồn tại",
		ErrKey:  ErrPhoneDuplicate,
		Status:  "Bad Request",
	}
}

func NewEmailDuplicateError() *AppError {
	return &AppError{
		Code:    http.StatusBadRequest,
		Message: "Email đã tồn tại",
		ErrKey:  ErrEmailDuplicate,
		Status:  "Bad Request",
	}
}

// -------- 401 Unauthorized --------
func NewUnauthorizedError() *AppError {
	return &AppError{
		Code:    http.StatusUnauthorized,
		Message: "Yêu cầu xác thực",
		ErrKey:  ErrUnauthorized,
		Status:  "Unauthorized",
	}
}

func NewTokenExpiredError() *AppError {
	return &AppError{
		Code:    http.StatusUnauthorized,
		Message: "Token đã hết hạn",
		ErrKey:  ErrTokenExpired,
		Status:  "Unauthorized",
	}
}

func NewTokenInvalidError() *AppError {
	return &AppError{
		Code:    http.StatusUnauthorized,
		Message: "Token không hợp lệ",
		ErrKey:  ErrTokenInvalid,
		Status:  "Unauthorized",
	}
}

func NewTokenMissingError() *AppError {
	return &AppError{
		Code:    http.StatusUnauthorized,
		Message: "Token bị thiếu",
		ErrKey:  ErrTokenMissing,
		Status:  "Unauthorized",
	}
}

// -------- 403 Forbidden --------
func NewForbiddenError() *AppError {
	return &AppError{
		Code:    http.StatusForbidden,
		Message: "Không có quyền truy cập",
		ErrKey:  ErrForbidden,
		Status:  "Forbidden",
	}
}

func NewAccessDeniedError() *AppError {
	return &AppError{
		Code:    http.StatusForbidden,
		Message: "Truy cập bị từ chối",
		ErrKey:  ErrAccessDenied,
		Status:  "Forbidden",
	}
}

// -------- 404 Not Found --------
func NewNotFoundError(resource string) *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: resource + " không tìm thấy",
		ErrKey:  ErrNotFound,
		Status:  "Not Found",
	}
}

func NewUserNotFoundError() *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: "Người dùng không tìm thấy",
		ErrKey:  ErrUserNotFound,
		Status:  "Not Found",
	}
}

// -------- 409 Conflict --------
func NewConflictError(message string) *AppError {
	return &AppError{
		Code:    http.StatusConflict,
		Message: message,
		ErrKey:  ErrConflict,
		Status:  "Conflict",
	}
}

// -------- 500 Internal Server Error --------
func NewInternalServerError(err error) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: "Lỗi máy chủ nội bộ",
		ErrKey:  err,
		Status:  "Internal Server Error",
	}
}

func NewDatabaseError(err error) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: "Lỗi cơ sở dữ liệu",
		ErrKey:  err,
		Status:  "Internal Server Error",
	}
}

func NewDatabaseTimeoutError(err error) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: "Kết nối cơ sở dữ liệu bị timeout",
		ErrKey:  err,
		Status:  "Internal Server Error",
	}
}
