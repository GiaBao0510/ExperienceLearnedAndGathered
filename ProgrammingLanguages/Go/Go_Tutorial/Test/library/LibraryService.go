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

	fmt.Printf("✅ Sách '%s' của tác giả '%s' đã được thêm với ID: %s\n", lib.books[id].Title, lib.books[id].Author, lib.books[id].ID)
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
	if len(lib.books) == 0 {
		fmt.Printf("📚 Thư viện hiện đang trống. Hãy thêm sách để bắt đầu quản lý!")
		return nil
	}

	//Thống kê thông tin sách
	for id, book := range lib.books {
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
	if len(lib.books) == 0 {
		fmt.Printf("📚 Thư viện hiện đang trống. Hãy thêm sách để bắt đầu quản lý!")
		return nil
	}

	//Nhập tiêu đề sách cần tìm
	search := utils.ReadNonEmptyInput("Nhập tiêu đề sách cần tìm: ")

	//Duyệt map để tìm
	for _, book := range lib.books {
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

	// Cập nhập mượn sách
	if err := lib.AddTransactionStore(ID, borrowerID, bookID); err != nil {	
		return err
	}

	fmt.Println("Mượn sách thành công!")

	return nil
}

// trả sách
func ReturnBook(lib *Library) error {
	
	transID := utils.ReadNonEmptyInput("Nhập ID giao dịch mượn sách: ")

	if err := lib.ReturnBook_store(transID); err != nil {
		return err
	}

	fmt.Println("Trả sách thành công!")

	return nil
}

// Lịch sử mượn sách
func HistoryOfBorrowingBooks(lib *Library) error{
	//Kiểm tra độ dài của giao dịch mượn sách
	if len(lib.transactions) == 0{
		fmt.Println("Giao dịch mượn sách trống !!!")
		return nil
	}

	//Thống kê
	for id, trans := range lib.transactions{
		
		//Kiểm tra xem thông tin đã trả sách chưa
		traSach := "Chưa trả sách"

		if lib.books[trans.BookID].Status {
			traSach = "Đã trả sách"
		}
		
		fmt.Printf("ID: %s, BorrowerID: %s, BookID: %s, BorrowDate: %s, Trạng thái trả sách: %s\n", id, trans.BorrowerID, trans.BookID, trans.BorrowDate, traSach)
	}

	return nil
}

// Xem chi tiết lịch sử mượn sách của một người
func DetailHistoryOfBorrowingBooks(lib *Library) error{

	//Kiểm tra độ dài của giao dịch mượn sách
	if len(lib.transactions) == 0{
		return fmt.Errorf("Giao dịch mượn sách trống !!!")
	}

	borrowerID := utils.ReadNonEmptyInput("Nhập ID người mượn sách: ")

	list := lib.DetailHistoryOfBorrowingBooks_Store(borrowerID)

	if list == nil{
		return fmt.Errorf("Không tìm thấy thông tin người mượn sách với ID: %s", borrowerID)
	
	}else{

		for _, element := range list{
			
			//Kiểm tra xem thông tin sách đã trả chưa
			traSach := "Chưa trả sách"
			if !lib.books[element.BookID].Status{
				traSach = element.ReturnDate.Format("01-01-2026")
			}

			fmt.Printf("ID giao dịch: %s, ID người mượn: %s, ID sách: %s, Ngày mượn: %s, Ngày trả sách: %s\n", element.ID, element.BorrowerID, element.BookID, element.BorrowDate.Format("01-01-2026"), traSach)
		}
	}	

	return nil
}


//Tìm kiếm sách dựa trên tiêu đề sách hoặc tên tác giả
func SearchBookByTitleOrAuthor(lib *Library) error{

	query := utils.ReadNonEmptyInput("Nhập tên sách hoặc tên tác giả để tìm kiếm sách: ")

	//Kiểm tra, nếu không nhập thông tin nào thì thông báo lỗi
	if query == ""{
		return fmt.Errorf("Lỗi vì bạn chưa nhập thông tin nào để tìm kiếm sách !!!")
	}

	result := lib.SearchBookByTitleOrAuthor_Store(query)

	for _, book := range result{

		trangThaiSach := "Có sẵn"
		if book.Status == false{
			trangThaiSach = "Đã mượn"
		}

		fmt.Printf("ID: %s, title: %s, author: %s, trạng thái: %s\n", book.ID, book.Title, book.Author, trangThaiSach)
	}
	
	return nil
}