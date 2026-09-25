package router

import (
	"github/GiaBao0510/error-handling/internal/controller"
	"github/GiaBao0510/error-handling/internal/middleware"

	"github.com/gin-gonic/gin"
)

/*
Nơi đây sẽ kết nối tất cả Controller
*/

func SetupRouterForUser(url *controller.UserController) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		// Các route cho User không cần auth
		users := api.Group("/users")
		{
			users.GET("", controller.Build(url.GetAll))
			users.GET("/:uid", controller.Build(url.GetByID))
			users.POST("", controller.Build(url.Create))
			users.PUT("/:uid", controller.Build(url.Update))
			users.DELETE("/:uid", controller.Build(url.Delete))
		}

		// Router kiểm nghiệm auth
		api.GET(
			"/auth-test",
			middleware.Auth(),
			func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Bạn đã xác thực thành công!"})
			},
		)
	}

	return r
}
