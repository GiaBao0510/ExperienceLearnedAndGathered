package models

import "time"

type Book struct {
	ID     string
	Title  string
	Author string
	Status bool //true: có sẵn, false: đã được mượn
}

// Cấu trúc người mượn sách
type Borrower struct {
	ID    string
	Name  string
	Email string
}

// Cấu trúc giao dịch mượn sách
type Transaction struct {
	ID         string
	BorrowerID string
	BookID     string
	BorrowDate time.Time	//Mặc định là thời điểm hiện tại khi tạo giao dịch
	ReturnDate time.Time	//Mặc định không có giá trị
}