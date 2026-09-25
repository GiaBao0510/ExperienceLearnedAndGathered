package routers

import (
	"golang-mongodb-crud/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {

	//Tạo instance của StudentHandler
	studentHandler := handlers.NewStudentHandler()

	// Định nghĩa các route cho student
	api := router.Group("/api/v1")
	{
		students := api.Group("students")
		{
			students.POST("", studentHandler.CreateStudent)
			students.GET("", studentHandler.GetAll)
			students.GET("/:id", studentHandler.GetByID)
			students.DELETE("/:id", studentHandler.DeleteStudent)
			students.PUT("/:id", studentHandler.UpdateStudent)
		}
	}
}