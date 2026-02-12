package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	//Tạo router
	r := gin.Default()

	//Tạo Endpoint GET với trạng thái mã là 200 và kèm theo thông tin trả về
	r.GET("/demo", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	//Endpoint
	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"data": "Danh sach nguoi dung",
		})
	})

	r.GET("/products", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"data": "Danh sach san pham",
		})
	})

	//Path Params
	r.GET(("/user/:id"), func(ctx *gin.Context) {
		id := ctx.Param("id") //Lấy tham số id từ URL
		ctx.JSON(200, gin.H{
			"data": "User: " + id,
		})
	})

	r.GET("/product/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(http.StatusOK, gin.H{
			"data": "Product: " + id,
		})
	})

	//Query Params
	r.GET("/product/name?:p_name,price=5000", func(ctx *gin.Context) {
		p_name := ctx.Param("p_name")
		price := ctx.Query("price")

		ctx.JSON(http.StatusOK, gin.H{
			"inf":   "Thong tin san pham",
			"data":  "Product: " + p_name,
			"price": price,
		})
	})

	r.Run(":8080") //Chạy ở port 8080

}
