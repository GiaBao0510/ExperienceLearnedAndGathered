package handler

import (
	"fmt"
	"net/http"
	"router-group/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ---- Tạo group quản lý user v1  -----

//Tạo list user
type User struct{
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

//Tạo slice để lưu trữ user
var users []User = []User{}

//Constructor để tạo instance của struct User
func NewUser() *User{

	names := []string{"John Doe", "Jane Smith", "Bob Johnson", "Alice Williams", "Charlie Brown", "Nguyen Van A"}
	emails := []string{"JDoe@gmail.com", "JSmith@gmail.com", "BJohnson@gmail.com", "AWilliams@gmail.com", "CBrown@gmail.com", "NVAn@gmail.com"}

	//Khởi tạo slice user với dữ liệu mẫu
	for i := 0; i < len(names); i++{
		users = append(users, User{
			ID: uuid.New(),		//Tạo ID ngẫu nhiên cho mỗi user
			Name: names[i],
			Email: emails[i],
		})
	}

	return  &User{}
} 

//Lấy danh sách user
func (obj *User) GetUsers(c *gin.Context){
	fmt.Println("Lấy danh sách người dùng")
	c.JSON(http.StatusOK, gin.H{
		"message": "list user",
		"data": users,
	})

}

//Lấy thông tin người dùng dựa trên UUID
func (obj *User) GetUserByUUID(c *gin.Context){
	
	id := c.Param("uuid")
	_, err := utils.ValidationUUID("UUID", id)

	//Kiểm tra định dạng UUID
	if  err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
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
	newUser.ID = uuid.New()

	//Thêm user mới vào slice
	users = append(users, newUser)
	
	c.JSON(http.StatusOK, gin.H{
		"message": "create user successfully",
	})
}

//Cập nhật thông tin người dùng
func (obj *User) UpdateUser(c *gin.Context){

	id := c.Param("uuid")

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

	id := c.Param("uuid")

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
