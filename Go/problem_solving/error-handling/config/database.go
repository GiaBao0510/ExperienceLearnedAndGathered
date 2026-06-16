package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Kết nối Database
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
	Sslmode  string
}

// DNS kết nối đến cơ sở dữ liệu PostgreSQL
func (c *DBConfig) DNS() string {
	return fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
        c.Host, c.Port, c.User, c.Password, c.Dbname, c.Sslmode,
    )
}

// NewDBConfig tạo một cấu hình mới từ biến môi trường
func NewDatabase(dns string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dns)
	fmt.Println("Đang kết nối đến cơ sở dữ liệu: ", dns)
	if err != nil {
		return nil, fmt.Errorf("Lỗi khi mở cơ sở dữ liệu: %v", err)
	}

	// Kiểm tra kết nối
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Không thể kết nối đến cơ sở dữ liệu: %v", err)
	}

	// Cấu hình
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)

	log.Println("✅Kết nối đến cơ sở dữ liệu thành công!")
	return db, nil
}

// LoadEnv tải biến môi trường từ file .env
func LoadEnv() error {
	return godotenv.Load()
}
