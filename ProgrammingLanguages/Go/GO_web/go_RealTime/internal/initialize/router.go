package initialize

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	var r *gin.Engine

	r = gin.Default()

	return r
}