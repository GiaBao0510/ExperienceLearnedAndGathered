package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {

	// Lấy thông tin từ context được set bởi middleware
	userID, _ := c.Get("user_id")
	keyID, _ := c.Get("api_key_id")
	scopes, _ := c.Get("scopes")

	c.JSON(http.StatusOK, gin.H{
		"message": "Chào mừng đến với trang profile!",
		"user_id": userID,
		"api_key_id": keyID,
		"scopes": scopes,
		"data": gin.H{
			"name": "John Doe",
			"email": "a@example.com",
		},
	})
}

// Create Data mô phỏng endpoint yêu cầu scope "write"
func CreateData(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "Data created successfully!",
        "by":      c.GetString("user_id"),
	})
}
