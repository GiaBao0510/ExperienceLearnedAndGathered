package handler

import (
	"fmt"
	"net/http"
	"router-group/utils"

	"github.com/gin-gonic/gin"
)

// ---- Tạo group quản lý user v1  -----

// Tạo list user
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Constructor để tạo instance của struct User
func NewUser() *User {
	return &User{}
}

// Tạo slice để lưu trữ user
var users []User = []User{
	{ID: 1, Name: "John Doe", Email: "JDoe@gmail.com"},
	{ID: 2, Name: "Jane Smith", Email: "JSmith@gmail.com"},
	{ID: 3, Name: "Bob Johnson", Email: "BJohnson@gmail.com"},
	{ID: 4, Name: "Alice Williams", Email: "AWilliams@gmail.com"},
	{ID: 5, Name: "Charlie Brown", Email: "CBrown@gmail.com"},
	{ID: 6, Name: "Nguyen Van A", Email: "NVAn@gmail.com"},
}

// Lấy danh sách user
func (obj *User) GetUsers(c *gin.Context) {
	fmt.Println("Lấy danh sách người dùng")
	c.JSON(http.StatusOK, gin.H{
		"message": "list user",
		"data":    users,
	})
}

// Tạo struct để kiểm tra đầu vào của GetUserByID
// Điều kiện ID > 0, và nó phải là kiểu int mới sử dụng được. Ngược lại thì không
type GetUserByID_Param struct {
	ID int `uri:"id" binding:"gt=0"`
}

// Lấy thông tin người dùng dựa trên ID
func (obj *User) GetUserByID(ctx *gin.Context) {

	var params GetUserByID_Param
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get user by ID successfully",
		"user_id": params.ID,
	})
}

// Tạo struct để kiểm tra đầu vào của GetUserByID
// Điều kiện ID > 0, và nó phải là kiểu int mới sử dụng được. Ngược lại thì không
type GetUserByUUID_Param struct {
	UUID string `uri:"uuid" binding:"uuid"`
}

// Lấy thông tin người dùng dựa trên UUID
func (obj *User) GetUserBy_UUID(ctx *gin.Context) {

	var params GetUserByUUID_Param
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get user by UUID successfully",
		"UUID":    params.UUID,
	})
}

// Thêm người dùng mới
func (obj *User) CreateUser(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "create user successfully",
	})
}

// Cập nhật thông tin người dùng
func (obj *User) UpdateUser(c *gin.Context) {

	c.JSON(http.StatusNotFound, gin.H{
		"message": "user not found",
	})
}

// Xóa người dùng
func (obj *User) DeleteUser(c *gin.Context) {

	c.JSON(http.StatusNotFound, gin.H{
		"message": "user not found",
		"data":    nil,
	})
}
