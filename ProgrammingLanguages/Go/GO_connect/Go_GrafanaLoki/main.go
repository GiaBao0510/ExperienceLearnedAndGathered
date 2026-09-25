package main

import (
	"go-grafana-loki/internal/logger"
	"go-grafana-loki/internal/middleware"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	zapLogger, err := logger.NewLogger("./internal/logs/http.log")
	if err != nil {
		log.Fatalf("Không thể khởi tạo logger: %v", err)
	}
	defer zapLogger.Sync() // Đảm bảo rằng tất cả log được ghi ra trước khi thoát

	r := gin.New()
	r.Use(middleware.GinLoggerMiddleware(zapLogger)) // Sử dụng middleware logger
	r.Use(gin.Recovery()) // Middleware để phục hồi từ panic và ghi log lỗi

	r.GET("health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.GET("/error-demo", func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "This is a demo error"})
	})

	zapLogger.Info("server starting",  zap.String("port", "8080"))

	if err := r.Run(":8080"); err != nil {
		zapLogger.Fatal("server failed to start", zap.String("error", err.Error()))
	}
}