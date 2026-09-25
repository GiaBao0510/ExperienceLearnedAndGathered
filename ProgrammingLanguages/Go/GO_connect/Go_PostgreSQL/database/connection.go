package database

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	
	// Chuỗi kết nối đến cơ sở dữ liệu PostgreSQL
	constStr := "host=localhost port=5432 user=admin password=admin123 dbname=test sslmode=disable"

	// Mở kết nối đến cơ sở dữ liệu
	pgDB, err := sql.Open("postgres", constStr)

	// Nếu có lỗi khi mở kết nối, log lỗi và dừng chương trình
	if err != nil{
		log.Fatal("Lỗi cấu hình kết nối: ", err)
	}

	//Kiêm tra kết nối thành công
	err = pgDB.Ping()
	if err != nil {
		log.Fatal("Lỗi kết nối đến cơ sở dữ liệu: ", err)
	}

	log.Println("Kết nối PostgreSQL thành công!")
	return pgDB
}