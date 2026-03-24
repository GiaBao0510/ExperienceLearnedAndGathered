package handler

import (
	"net/http"
	"router-group/utils"

	"github.com/gin-gonic/gin"
)

type Category struct {
}

func NewCategoryHandler() *Category {
	return &Category{}
}

type GetCategoryByCategoriest_Param struct {
	Category string `uri:"category" binding:"oneof=php python golang java"`
}

// Lấy Categories theo mục
func (obj *Category) GetCategoryByCategories(c *gin.Context) {

	var params GetCategoryByCategoriest_Param
	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"course":  params.Category,
	})
}
