package library

import (
	"fmt"
	"hello/models"
)

//Tệp tin này tạp trung chủ yếu để xử lý dữ liệu
type Library struct{
	book map[string]models.Book
	borrowers map[string]models.Borrower
	transaction map[string]models.Transaction
}

//Hàm khởi tạo trả về con trỏ đến một đối tượng mới
func NewLibrary() *Library{			
	return &Library{
		book: make(map[string]models.Book),
		borrowers: make(map[string]models.Borrower),
		transaction: make(map[string]models.Transaction) ,
	}
}

//Tạo hàm thêm sách
func (lib *Library) AddBookStore(id, title, author string) error{
	
	//Kiếm tra ID đã tồn tại chưa
	if _, exist := lib.book[id]; exist{
		return fmt.Errorf("ID: %s. sách đã trùng. Xin vui lòng nhập lại\n", id)
	}
	
	lib.book[id] = models.Book{
		ID: id,
		Title: title,
		Author: author,
		Status: true,
	}

	return nil
}

//Tạo hàm thêm người mượn
func (lib *Library) AddBorrowerStore(id, name, email string) error{
	
	//Kiếm tra ID đã tồn tại chưa
	if _, exist := lib.borrowers[id]; exist{
		return fmt.Errorf("Lỗi vì ID người mượn sách đã bị trùng: %s\n", id)
	}

	lib.borrowers[id] = models.Borrower{
		ID: id,
		Name: name,
		Email: email,
	}

	return nil
}

//Kiểm tra ID người mượn sách đã tồn tại chưa
func findBorrowerByID(lib *Library, id string) *models.Borrower{
	if borrower, exist := lib.borrowers[id]; exist{
		return &borrower
	}
	return nil
} 


//Tạo hàm thêm giao dịch mượn sách
func (lib *Library) AddTransactionStore(id, borrowerID, bookID string) error{
	
	//Kiếm tra ID đã tồn tại chưa
	if _,exist := lib.transaction[id]; exist{
		return fmt.Errorf("Lỗi vì ID giao dịch đã bị trùng: %s\n", id)
	}

	lib.transaction[id] = models.Transaction{
		ID: id,
		BorrowerID: borrowerID,
		BookID: bookID,
	}

	return nil
}