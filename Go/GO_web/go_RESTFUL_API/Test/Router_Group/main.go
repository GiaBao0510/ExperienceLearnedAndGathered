package main

import (
	v1Handler "router-group/internal/api/v1/handler"
	v2Handler "router-group/internal/api/v2/handler"
	"router-group/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	if err := utils.RegisterValidationError(); err != nil {
		panic(err)
	}

	//Tạo instance của struct Handler_v1 để gọi các phương thức
	userHandler_v1 := v1Handler.NewUser()
	productHandler_v1 := v1Handler.NewProduct()
	categoryHandler_v1 := v1Handler.NewCategoryHandler()
	newsHandler_v1 := v1Handler.NewsHandlerConstructor()

	userHandler_v2 := v2Handler.NewUser()
	productHandler_v2 := v2Handler.NewProduct()
	categoryHandler_v2 := v2Handler.NewCategoryHandler()

	//Tạo một group quản lý API
	v1 := r.Group("/api/v1")
	{
		//Group luôn user thêm một lần nữa
		user := v1.Group("/users")
		{
			user.GET("/", userHandler_v1.GetUsers)
			user.GET("/:uuid", userHandler_v1.GetUserByUUID)
			user.POST("/", userHandler_v1.CreateUser)
			user.PUT("/:uuid", userHandler_v1.UpdateUser)
			user.DELETE("/:uuid", userHandler_v1.DeleteUser)
		}

		product := v1.Group("/products")
		{
			product.GET("/", productHandler_v1.GetProducts)
			product.GET("/:slug", productHandler_v1.GetProductBySlug)
			product.GET("", productHandler_v1.SearchProducts)
			product.POST("/", productHandler_v1.CreateProduct)
			product.PUT("/:id", productHandler_v1.UpdateProduct)
			product.DELETE("/:id", productHandler_v1.DeleteProduct)
		}

		category := v1.Group("/categories")
		{
			category.GET("/:category", categoryHandler_v1.GetCategoryByCategories)
		}

		news := v1.Group("/news")
		{
			news.GET("/", newsHandler_v1.GetNews)
			news.GET("/:slug", newsHandler_v1.GetNews)
		}

	}

	v2 := r.Group("/api/v2")
	{
		user := v2.Group("/users")
		{
			user.GET("/", userHandler_v2.GetUsers)
			user.GET("/:id", userHandler_v2.GetUserByID)
			user.GET("/uuid/:uuid", userHandler_v2.GetUserBy_UUID)	// Add
			user.POST("/", userHandler_v2.CreateUser)
			user.PUT("/:id", userHandler_v2.UpdateUser)
			user.DELETE("/:id", userHandler_v2.DeleteUser)
		}

		product := v2.Group("/products")
		{
			product.GET("/", productHandler_v2.GetProducts)
			product.GET("/:slug", productHandler_v2.GetProductBySlug)
			product.POST("/", productHandler_v2.CreateProduct)
			product.PUT("/:id", productHandler_v2.UpdateProduct)
			product.DELETE("/:id", productHandler_v2.DeleteProduct)
		}

		category := v2.Group("/categories")
		{
			category.GET("/:category", categoryHandler_v2.GetCategoryByCategories)
		}
	}

	r.Run(":8080") // Lắng nghe trên cổng 8080
}
