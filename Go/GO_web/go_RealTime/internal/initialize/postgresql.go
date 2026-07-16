package initialize

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/GiaBao0510/Go-Realtime/global"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// hàm kiểm tra lỗi
func CheckErrorPanic(err error, message string) {
	if err != nil {
		panic(fmt.Sprintf("%s: %v", message, err))
	}
}

// Kiểm tra kết nối
func CheckConnection(db *sql.DB) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		CheckErrorPanic(err, "không thể ping DB: %w")
	}
}

// hàm khởi tạo Postgre
func InitPostgresql() {

	m := global.Config.DB

	// Chuỗi kết nối PostgreSQL
	var stringConn = fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		m.DB_Host, m.DB_Port, m.DB_User, m.DB_Password, m.DB_Name, m.DB_SSLMode,
	)

	fmt.Println("Chuỗi kết nối PostgreSQL:", stringConn)

	db, err := sql.Open("pgx", stringConn)
	CheckErrorPanic(err, "Khởi tạo kết nối Postgresql: Kết nối thất bại")
	CheckConnection(db)

	// Kết nối thành công, gán giá trị cho biến toàn cục PostgreSQL
	global.PostgreSQL = db
	log.Println("Kết nối PostgreSQL thành công")

	setPool(db) // Cấu hình pool kết nối
}

func setPool(db *sql.DB) {
	db.SetMaxOpenConns(25)                 // Số lượng kết nối tối đa mở
	db.SetMaxIdleConns(10)                 // Số lượng kết nối tối đa nhàn rỗi
	db.SetConnMaxIdleTime(5 * time.Minute) // Thời gian tối đa một kết nối có thể nhàn rỗi
}
