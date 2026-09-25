package controller

import (
	"github/GiaBao0510/error-handling/internal/domain"
	"github/GiaBao0510/error-handling/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
	Controller bây giờ cực kỳ đơn giản: parse request → gọi service → trả JSON thành công.
Khi có lỗi: chỉ cần return err. Build() sẽ lo phần còn lại
*/

type UserController struct{ userService service.IUserService }

func NewUserController(obj service.IUserService) *UserController {
	return &UserController{userService: obj}
}

// GetByID - GET api/v1/users/:id
func (c *UserController) GetByID(ctx *gin.Context) error {
	id, err := strconv.ParseInt(ctx.Param("uid"), 10, 64)
	if err != nil {
		return domain.NewBadRequestError("ID phải là số nguyên")
	}

	user, err := c.userService.GetByID(id)
	if err != nil {
		return err // Trả lỗi về cho Builder xử lý
	}

	ctx.JSON(http.StatusOK, domain.SuccessResponse{
		Code:    http.StatusOK,
		Data:    user,
		Message: "Lấy thông tin người dùng thành công",
	})
	return nil
}

// GetAll - GET api/v1/users
func (c *UserController) GetAll(ctx *gin.Context) error {
	users, err := c.userService.GetAll()
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, domain.SuccessResponse{
		Code:    http.StatusOK,
		Data:    users,
		Message: "Lấy danh sách người dùng thành công",
	})
	return nil
}

// Create - POST api/v1/users
func (c *UserController) Create(ctx *gin.Context) error {
	var req domain.CreateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		return domain.NewBadRequestError("Dữ liệu không hợp lệ: " + err.Error())
	}

	user, err := c.userService.Create(&req)
	if err != nil {
		return err
	}
 
	ctx.JSON(http.StatusCreated, domain.SuccessResponse{
		Code:    http.StatusCreated,
		Data:    user,
		Message: "Tạo người dùng thành công",
	})
	return nil
}

func (c *UserController) Update(ctx *gin.Context) error {

	id, err := strconv.ParseInt(ctx.Param("uid"), 10, 64)
	if err != nil {
		return domain.NewBadRequestError("ID phải là số nguyên")
	}

	var req domain.UpdateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		return domain.NewBadRequestError("Dữ liệu không hợp lệ: " + err.Error())
	}

	user, err := c.userService.Update(id, &req)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, domain.SuccessResponse{
		Code:    http.StatusOK,
		Data:    user,
		Message: "Cập nhật người dùng thành công",
	})
	return nil
}

func (c *UserController) Delete(ctx *gin.Context) error {
	id, err := strconv.ParseInt(ctx.Param("uid"), 10, 64)
	if err != nil {
		return domain.NewBadRequestError("ID phải là số nguyên")
	}

	if err := c.userService.Delete(id); err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, domain.SuccessResponse{
		Code:    http.StatusOK,
		Message: "Xóa người dùng thành công",
	})

	return nil
}
