package main

import (
	v1Handler "router-group/internal/api/v1/handler"
	v2Handler "router-group/internal/api/v2/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//Tạo instance của struct Handler_v1 để gọi các phương thức
	userHandler_v1 := v1Handler.NewUser()
	productHandler_v1 := v1Handler.NewProduct()

	userHandler_v2 := v2Handler.NewUser()

	//Tạo một group quản lý API
	v1 := r.Group("/api/v1")
	{
		//Group luôn user thêm một lần nữa
		user := v1.Group("/users")
		{
			user.GET("/", userHandler_v1.GetUsers)
			user.GET("/:id", userHandler_v1.GetUserByID)
			user.GET("/:uuid", userHandler_v1.GetUserByUUID)
			user.POST("/", userHandler_v1.CreateUser)
			user.PUT("/:id", userHandler_v1.UpdateUser)
			user.DELETE("/:id", userHandler_v1.DeleteUser)
		}

		product := v1.Group("/products")
		{
			product.GET("/", productHandler_v1.GetProducts)
			product.GET("/:id", productHandler_v1.GetProductByID)
			product.POST("/", productHandler_v1.CreateProduct)
			product.PUT("/:id", productHandler_v1.UpdateProduct)
			product.DELETE("/:id", productHandler_v1.DeleteProduct)
		}
		
	}

	v2 := r.Group("/api/v2")
	{
		user := v2.Group("/users")
		{
			user.GET("/", userHandler_v2.GetUsers)
			user.GET("/:id", userHandler_v2.GetUserByID)
			user.POST("/", userHandler_v2.CreateUser)
			user.PUT("/:id", userHandler_v2.UpdateUser)
			user.DELETE("/:id", userHandler_v2.DeleteUser)
		}
	}

	r.Run(":8080") // Lắng nghe trên cổng 8080
}
