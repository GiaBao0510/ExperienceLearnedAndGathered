package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ---- Tạo group quản lý user v1  -----

//Tạo list user
type User struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

//Constructor để tạo instance của struct User
func NewUser() *User{
	return  &User{}
} 

//Tạo slice để lưu trữ user
var users []User = []User{
	{ID: uuid., Name: "John Doe", Email: "JDoe@gmail.com"},
	{ID: 2, Name: "Jane Smith", Email: "JSmith@gmail.com"},
	{ID: 3, Name: "Bob Johnson", Email: "BJohnson@gmail.com"},
	{ID: 4, Name: "Alice Williams", Email: "AWilliams@gmail.com"},
	{ID: 5, Name: "Charlie Brown", Email: "CBrown@gmail.com"},
	{ID: 6, Name: "Nguyen Van A", Email: "NVAn@gmail.com"},
}

//Lấy danh sách user
func (obj *User) GetUsers(c *gin.Context){
	fmt.Println("Lấy danh sách người dùng")
	c.JSON(http.StatusOK, gin.H{
		"message": "list user",
		"data": users,
	})
}

//Lấy thông tin người dùng dựa trên ID
func (obj *User) GetUserByID(c *gin.Context){
	
	id := c.Param("id")
	for _,user := range users{
		if fmt.Sprintf("%v", user.ID) == id{
			c.JSON(http.StatusOK, gin.H{
				"message": "user found",
				"data": user,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "user not found",
		"data": nil,
	})
	return
}

//Lấy thông tin người dùng dựa trên UUID
func (obj *User) GetUserByUUID(c *gin.Context){
	
	id := c.Param("id")

	//Kiểm tra định dạng UUID
	if _, err := uuid.Parse(id); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid uuid format",
			"error": err.Error(),
		})
		return
	}


	for _,user := range users{
		if fmt.Sprintf("%v", user.ID) == id{
			c.JSON(http.StatusOK, gin.H{
				"message": "user found",
				"data": user,
			})
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "user not found",
		"data": nil,
	})
}

//Thêm người dùng mới
func (obj *User) CreateUser(c *gin.Context) {
	fmt.Println("Nhập thông tin người dùng")
	
	//Đọc dữ liệu từ request body
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil{		//Nếu có lỗi khi đọc dữ liệu, trả về lỗi
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
		})
		return
	}

	//Tạo ID mới cho user
	newUser.ID = len(users) + 1

	//Thêm user mới vào slice
	users = append(users, newUser)
	
	c.JSON(http.StatusOK, gin.H{
		"message": "create user successfully",
	})
}


//Cập nhật thông tin người dùng
func (obj *User) UpdateUser(c *gin.Context){

	id := c.Param("id")

	//Đọc dữ liệu từ request body
	var updateUser User
	if err := c.ShouldBindJSON(&updateUser); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
			"error": err.Error(),
		})
		return
	}

	//Tìm ID của User rồi cập nhật thông tin
	for idx, user := range users{
		if fmt.Sprintf("%v",user.ID) == id{
			
			//Cập nhật thông tin user
			users[idx].Name = updateUser.Name
			users[idx].Email = updateUser.Email

			//Trả về thông báo cập nhật thành công
			c.JSON(http.StatusOK, gin.H{
				"message": "update user successfully",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "user not found",
	})
}

//Xóa người dùng
func (obj *User) DeleteUser(c *gin.Context){

	id := c.Param("id")

	for index, user := range users{
		if fmt.Sprintf("%v", user.ID) == id{
			users = append(users[:index], users[index+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"message": "delete user successfully",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "user not found",
		"data": nil,
	})
}
