package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string
	DBPath  string
}

// Load đọc file .env và trả về cấu hình
func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("[config] No .env file found, using system env")
	}
	return &Config{
		AppPort: getEnv("APP_PORT", "8080"),
		DBPath: getEnv("DB_PATH", "host=localhost;port=5432;user=admin;password=admin123;dbname=test;sslmode=disable"),
	}
}

// getEnv lấy giá trị từ biến môi trường hoặc trả về giá trị mặc định nếu biến không tồn tại
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}