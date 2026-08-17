package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GinLoggerMiddleware trả về một middleware Gin, tự động log mọi request đi qua
func GinLoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Cho request đi tiếp qua các handler khác
		c.Next()

		// Sau khi handler xử lý xong, tính thời gian đã trôi qua
		latency := time.Since(start)
		status := c.Writer.Status()

		fields := []zap.Field{
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", status),
			zap.Duration("latency_ms", latency),
			zap.String("client_ip", c.ClientIP()),
		}

		// Log level tự động thay đổi theo status code — đây là thực hành chuẩn
		switch {
		case status >= 500:
			logger.Error("request completed", fields...)
		case status >= 400:
			logger.Warn("request completed", fields...)
		default:
			logger.Info("request completed", fields...)
		}
	}
}