package library

import (
	"fmt"
	"hello/models"
	"strings"
	"time"
)

//Tệp tin này tạp trung chủ yếu để xử lý dữ liệu
type Library struct{
	books map[string]models.Book
	borrowers map[string]models.Borrower
	transactions map[string]models.Transaction
}

//Hàm khởi tạo trả về con trỏ đến một đối tượng mới
func NewLibrary() *Library{			
	return &Library{
		books: make(map[string]models.Book),
		borrowers: make(map[string]models.Borrower),
		transactions: make(map[string]models.Transaction) ,
	}
}

//Tạo hàm thêm sách
func (lib *Library) AddBookStore(id, title, author string) error{
	
	//Kiếm tra ID đã tồn tại chưa
	if _, exist := lib.books[id]; exist{
		return fmt.Errorf("ID: %s. sách đã trùng. Xin vui lòng nhập lại\n", id)
	}
	
	lib.books[id] = models.Book{
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
	
	//Kiểm tra ID sách có tồn tại chưa
	book, bookExist := lib.books[bookID] 
	
	if bookExist == false{
		return fmt.Errorf("Lỗi vì ID sách không tồn tại: %s\n", bookID)
	}

	//Kiểm tra xem sách đã có sẵn chưa
	if !book.Status {
		return fmt.Errorf("Lỗi vì sách đã được mượn: %s\n", bookID)
	}

	//Kiểm tra ID của người dùng đã tồn tại chưa
	if _, exist := lib.borrowers[borrowerID]; !exist{
		return fmt.Errorf("Lỗi vì ID người mượn sách không tồn tại: %s\n", borrowerID)
	}

	//Cập nhật trạng thái sách đã mượn
	book.Status = false
	lib.books[bookID] = book

	//Tạo thông tin giao dịch mượn sách
	lib.transactions[id] = models.Transaction{
		ID: id,
		BorrowerID: borrowerID,
		BookID: bookID,
		BorrowDate: time.Now(),
	}

	return nil
}

// Danh sách giao dịch mượn sách của 1 người cụ thể
func (lib *Library) DetailHistoryOfBorrowingBooks_Store(borrowerID string) []models.Transaction{

	//Kiểm tra ID người mượn có tồn tại không
	if _, exist := lib.borrowers[borrowerID]; !exist{
		return nil
	}

	//Duyệt map để tìm thông tin giao dịch của người dùng này
	var result []models.Transaction
	for _, trans := range lib.transactions{
		if trans.BorrowerID == borrowerID{
			result = append(result, trans)
		}
	}

	return result
}

//Thực hiện trả sách
func (lib *Library) ReturnBook_store(transactionID string) error{
	
	//Kiểm tra ID giao dịch có tồn tại không
	if _, exist := lib.transactions[transactionID]; !exist{
		return fmt.Errorf("Lỗi vì ID giao dịch không tồn tại: %s\n", transactionID)
	}

	//Kiểm tra xem sách đã được trả chưa, Nếu trả rồi thì thông báo
	if lib.books[lib.transactions[transactionID].BookID].Status{
		return fmt.Errorf("Lỗi vì sách đã được trả rồi: %s\n", lib.transactions[transactionID].BookID)
	}

	//Cập nhật lại trạng thái sách và ngày trả sách
	book := lib.books[lib.transactions[transactionID].BookID]
	book.Status = true
	lib.books[lib.transactions[transactionID].BookID] = book

	transaction := lib.transactions[transactionID]
	transaction.ReturnDate = time.Now()
	lib.transactions[transactionID] = transaction

	return nil
}

// Tìm kiêm sách dựa trên tên sách hoặc tên tác giả
func (lib *Library) SearchBookByTitleOrAuthor_Store(query string) []models.Book{
	
	//Chuyển chuyển tìm kiếm thành chữ thường hết
	query = strings.ToLower(query)

	var result []models.Book

	//Duyệt
	for _, book := range lib.books{
		if strings.Contains(strings.ToLower(book.Title), query) ||
			strings.Contains(strings.ToLower(book.Author), query){
				result = append(result, book)
			}
	}

	return result
}