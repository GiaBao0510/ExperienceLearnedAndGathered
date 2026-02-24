package library

//Tệp tin này chủ yếu xử lý logic code

import (
	"fmt"
	"hello/utils"
)

// Thêm sách
func AddBook(lib *Library) error {
	id := utils.GenerateID()
	title := utils.ReadNonEmptyInput("Nhập tiêu đề sách: ")
	author := utils.ReadNonEmptyInput("Nhập tác giả sách: ")

	if err := lib.AddBookStore(id, title, author); err != nil {
		return err
	}

	fmt.Printf("✅ Sách '%s' của tác giả '%s' đã được thêm với ID: %s\n", lib.book[id].Title, lib.book[id].Author, lib.book[id].ID)
	return nil
}

// Sửa sách
func EditBook() {

}

// Xóa sách
func DeleteBook() {

}

// Danh sách sách
func ListBooks(lib *Library) error {
	//Kiểm tra độ dài của map sách
	if len(lib.book) == 0 {
		fmt.Printf("📚 Thư viện hiện đang trống. Hãy thêm sách để bắt đầu quản lý!")
		return nil
	}

	//Thống kê thông tin sách
	for id, book := range lib.book {
		trangThaiMuon := "Có sẵn"
		if !book.Status {
			trangThaiMuon = "Đã mượn"
		}
		fmt.Printf("ID: %s, title: %s, author: %s, trạng thái: %s\n", id, book.Title, book.Author, trangThaiMuon)
	}

	return nil
}

// Tìm kiếm sách
func SearchBook(lib *Library) error {
	//Kiểm tra độ dài của map sách
	if len(lib.book) == 0 {
		fmt.Printf("📚 Thư viện hiện đang trống. Hãy thêm sách để bắt đầu quản lý!")
		return nil
	}

	//Nhập tiêu đề sách cần tìm
	search := utils.ReadNonEmptyInput("Nhập tiêu đề sách cần tìm: ")

	//Duyệt map để tìm
	for _, book := range lib.book {
		if book.Title == search {
			fmt.Printf("✅ Tìm thấy sách: ID: %s, title: %s, author: %s\n", book.ID, book.Title, book.Author)
			return nil
		}
	}

	return fmt.Errorf("Không tìm thấy thông tin sách cần tìm: %s", search)
}

// Thêm người mượn sách
func AddBorrower(lib *Library) error{
	id := utils.GenerateID()
	name := utils.ReadNonEmptyInput("Nhập tên người mượn: ")
	email := utils.ReadNonEmptyInput("Nhập email người mượn: ")

	if err := lib.AddBorrowerStore(id, name, email); err != nil {
		return err
	}

	fmt.Printf("✅ Người mượn '%s' đã được thêm với ID: %s\n", lib.borrowers[id].Name, lib.borrowers[id].ID)
	return nil
}

// Danh sách người mượn
func ListBorrowers(lib *Library) error {
	//Kiểm tra độ dài của map người mượn
	if len(lib.borrowers) == 0 {
		fmt.Printf("📚 Không có người mượn nào trong thư viện!")
		return nil
	}

	//Thống kê thông tin sách
	for id, borrower := range lib.borrowers {
		fmt.Printf("ID: %s, name: %s, email: %s\n", id, borrower.Name, borrower.Email)
	}

	return nil
}

// Mượn sách
func BorrowBook(lib *Library) error{
	
	//Nhập thông tin đầu vào
	borrowerID := utils.ReadNonEmptyInput("Nhập ID người mượn: ")
	bookID := utils.ReadNonEmptyInput("Nhập ID sách cần mượn: ")
	ID := utils.GenerateID()

	//Kiểm tra ID của người dùng đã tồn tại chưa
	if _, exist := lib.borrowers[borrowerID]; !exist{
		return fmt.Errorf("Lỗi vì ID người mượn sách không tồn tại: %s\n", borrowerID)
	}

	//Kiểm tra ID sách đã tồn tại chưa
	if _, exist := lib.book[bookID]; !exist{
		return fmt.Errorf("Lỗi vì ID sách không tồn tại: %s\n", bookID)
	}

	//Kiểm tra xem sách này đã có sẵn chưa
	if !lib.book[bookID].Status{
		return fmt.Errorf("Lỗi vì sách đã được mượn: %s\n", bookID)
	}

	// Cập nhập mượn sách
	lib.AddTransactionStore(ID, borrowerID, bookID)

	return nil
}

// trả sách
func ReturnBook() {

}

// Lịch sử mượn sách
func HistoryOfBorrowingBooks() {}

// Xem chi tiết lịch sử mượn sách của một người
func DetailHistoryOfBorrowingBooks() {}
