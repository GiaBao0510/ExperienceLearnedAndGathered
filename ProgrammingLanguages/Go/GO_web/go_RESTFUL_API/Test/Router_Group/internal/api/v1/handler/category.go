package handler

import (
	"net/http"
	"router-group/utils"

	"github.com/gin-gonic/gin"
)

// Tạo mục hợp lệ cho từng category
var categories = map[string]bool{
	"golang": true,
	"csharp": true,
	"python": true,
}

type Category struct {
}

func NewCategoryHandler() *Category {
	return &Category{}
}

// Lấy Categories theo mục
func (obj *Category) GetCategoryByCategories(c *gin.Context) {

	//Lấy tham số từ param
	category := c.Param("category")

	//Kiểm tra nếu tham số này lấy không hợp lệ thì báo lỗi
	if err := utils.ValidationInList("category", category, categories); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    category,
	})
}
