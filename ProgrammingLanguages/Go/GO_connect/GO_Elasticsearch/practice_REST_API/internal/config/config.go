package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// struct Config chứa toàn bộ cấu hình của ứng dụng
type Config struct {
	ES_Address string	// Địa chỉ đầy đủ của Elasticsearch, ví dụ: http://localhost:9200
	ES_Port string		
	APP_Port string
	ES_INDEX_PRODUCTS string
}

// Hàm này để đọc file .env
func Load() (*Config, error) {
	// Đọc và nạp file .env vào os.Getenv
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("Lỗi, không đọc được .env: %v\n", err)
	}

	return &Config{
		ES_Address: os.Getenv("ES_HOST") +":" + os.Getenv("ES_PORT"),
		ES_Port: os.Getenv("ES_PORT"),
		APP_Port: os.Getenv("APP_PORT"),
		ES_INDEX_PRODUCTS: os.Getenv("ES_INDEX_PRODUCTS"),
	}, nil
}
