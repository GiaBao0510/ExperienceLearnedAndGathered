package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/GiaBao0510/go-apikey-demo/internal/config"
	"github.com/GiaBao0510/go-apikey-demo/internal/db"
	"github.com/GiaBao0510/go-apikey-demo/internal/handler"
	"github.com/GiaBao0510/go-apikey-demo/internal/middewarre"
	"github.com/gin-gonic/gin"
)

func main() {
	//0. Thiết lập thời gian bắt đầu để đo thời gian khởi động
	startTime := time.Now()

	//1. Load cấu hình
	cfg := config.Load()

	fmt.Printf("Starting server on port %s...\n", cfg.AppPort)
	fmt.Printf("Database path: %s\n", cfg.DBPath)

	//2. Khởi tạo Database
	database := db.Init(cfg.DBPath)
	defer database.Close()

	//3. Khởi tạo server và chạy
	r := gin.Default()

	//4. Khởi tạo Handler
	keyHandler := handler.NewAPIKeyHandler(database)

	// --- Roter public không cần xác thực ---
	r.GET("/health", func(c * gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// Quản lý API key
	apiKeys := r.Group("/api-key")
	{
		apiKeys.POST("", keyHandler.CreateAPIKey)
		apiKeys.GET("", keyHandler.ListAPIKeys)
		apiKeys.DELETE("/:id", keyHandler.RevokeAPIKey)
	}

	// --- Roter protected cần xác thực API key ---
	// Group với middleware xác thực scope "read"
	readRoutes := r.Group("/api/v1")
	readRoutes.Use(middewarre.APIKeyAuth(database, "read"))
	{
		readRoutes.GET("/profile", handler.GetProfile)
	}

	// Group với middleware xác thực scope "write"
	writeRoutes := r.Group("/api/v1")
	writeRoutes.Use(middewarre.APIKeyAuth(database, "write"))
	{
		writeRoutes.POST("/data", handler.CreateData)
	}

	// 5. Chạy server - đo thời gian khởi động
	addr := ":" + cfg.AppPort
	log.Printf("Server is running at %s, startup time: %v\n", addr, time.Since(startTime))
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
