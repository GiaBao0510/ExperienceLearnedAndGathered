package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/xuri/excelize/v2"
	_ "github.com/lib/pq"
)

// Tại đây sẽ thực hiện việc xuất dữ liệu từ PostgreSQL sang Excel với số lượng lớn (hàng triệu bản ghi).

func main() {

	// Bắt đầu đo thời gian thực thi
	startTime := time.Now()

	db, err := sql.Open("postgres", "host=localhost port=5432 user=admin password=admin123 dbname=test sslmode=disable")
	if err != nil {
		log.Fatal("Lỗi cấu hình kết nối: ", err)
	}
	defer db.Close() // Đảm bảo đóng kết nối sau khi hoàn thành

	rows, err := db.Query("SELECT * FROM users;")
	if err != nil {
		log.Fatal("Lỗi truy vấn dữ liệu: ", err)
	}
	defer rows.Close() // Đảm bảo đóng kết quả truy vấn sau khi hoàn thành

	// Đo thời gian thực thi sau khi truy vấn dữ liệu
	elapsed := time.Since(startTime)
	log.Printf("Thời gian thực thi: %s", elapsed)

	// Tạo một file Excel mới
	f := excelize.NewFile()

	// Áp dụng stream để ghi dữ liệu vào Excel một cách hiệu quả, tránh việc tải toàn bộ dữ liệu vào bộ nhớ
	streeamWriter, err := f.NewStreamWriter("Sheet1")
	if err != nil {
		log.Fatal("Lỗi tạo StreamWriter: ", err)
	}

	rowIndex := 1
	for rows.Next() {
		var id int
		var name string
		var phone string
		var address string

		// Thực hiện scan dữ liệu từ hàng hiện tại vào các biến tương ứng
		err := rows.Scan(&id, &name, &phone, &address)
		if err != nil {
			log.Fatal("Lỗi scan dữ liệu: ", err)
		}

		// Ghi dữ liệu vào Excel sử dụng StreamWriter
		cell, _ := excelize.CoordinatesToCellName(1, rowIndex)
		err = streeamWriter.SetRow(cell, []interface{}{id, name, phone, address})
		if err != nil {
			log.Fatal("Lỗi ghi dữ liệu vào Excel: ", err)
		}

		rowIndex++ // Tăng chỉ số hàng để ghi dữ liệu vào hàng tiếp theo
	}

	// Đảm bảo flush dữ liệu vào Excel sau khi ghi xong
	if err := streeamWriter.Flush(); err != nil {
		log.Fatal("Lỗi khi flush dữ liệu vào Excel: ", err)
	}

	// Lưu file Excel sau khi ghi xong
	if err := f.SaveAs("output2.xlsx"); err != nil {
		log.Fatal("Lỗi khi lưu file Excel: ", err)
	}

	endElapsed := time.Since(startTime)
	log.Printf("Thời gian thực thi: %s", endElapsed)
}