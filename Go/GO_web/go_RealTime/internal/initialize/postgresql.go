package initialize

import (
	"fmt"
	"strconv"

	"github.com/GiaBao0510/Go-Realtime/internal/config"
	"github.com/GiaBao0510/Go-Realtime/internal/utils"
)

// Kiểm tra xem Config DB đã điền đủ chưa
func Check_ConfigDB(config *config.Config) error { 

	switch {
		case config.DB.Host == "":
			return fmt.Errorf("DB_HOST chưa được điền")
		case config.DB.Port == 0:
			return fmt.Errorf("DB_PORT chưa được điền")
		case config.DB.User == "":
			return fmt.Errorf("DB_USER chưa được điền")
		case config.DB.Password == "":
			return fmt.Errorf("DB_PASSWORD chưa được điền")
		case config.DB.DBName == "":
			return fmt.Errorf("DB_NAME chưa được điền")
	}
	return nil
}

// hàm khởi tạo Postgre
func InitPostgresql() (*config.Config, error) {
	port, err := strconv.Atoi(utils.GetEnv("DB_PORT","5432"))
	if err != nil {
		return nil, fmt.Errorf("DB_PORT không hợp lệ")
	}

	config := &config.Config{
		DB: config.DBConfig{
			Host: utils.GetEnv("DB_HOST","localhost"),
			Po
		},
	}

	return

}
