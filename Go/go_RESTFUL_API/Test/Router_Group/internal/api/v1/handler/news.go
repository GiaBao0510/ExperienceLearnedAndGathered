package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type NewsHandler struct {
}

func NewsHandlerConstructor() *NewsHandler {
	return &NewsHandler{}
}

func (obj *NewsHandler) GetNews(c *gin.Context){

	slug := c.Param("slug")

	if slug == ""{
		slug = c.DefaultQuery("slug", "thong-tin-moi-cap-nhat")

		c.JSON(http.StatusOK, gin.H{
			"message": slug,
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"message": slug,
		})
	}
}