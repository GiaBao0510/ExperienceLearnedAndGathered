package main

import (
	"golang-mongodb-crud/config"
	"golang-mongodb-crud/routers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//Load biến môi trường
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	//Kết nối MongoDB
	config.ConnectDB()

	// Tạo Gin router
	// gin.Default() tự động thêm Logger và Recovery middleware
    // Logger: in log mỗi request ra terminal
    // Recovery: tự recover khi app bị panic (không crash hoàn toàn)
	router := gin.Default()

	//Đăng ký taatr cả routes
	routers.SetupRouter( router)

	//Chạy server trên cổng 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Mặc định nếu không có biến môi trường PORT
	}

	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to run server: ", err)
	}

	log.Println("Server is running on port " + port)
}