# Bài Tập Thực Hành: Hệ thống Quản lý Thư viện

## 📋 Mục lục

1. [Giới thiệu](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#1-gi%E1%BB%9Bi-thi%E1%BB%87u)
2. [Kiến trúc hệ thống](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#2-ki%E1%BA%BFn-tr%C3%BAc-h%E1%BB%87-th%E1%BB%91ng)
3. [Chuẩn bị và cài đặt](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#3-chu%E1%BA%A9n-b%E1%BB%8B-v%C3%A0-c%C3%A0i-%C4%91%E1%BA%B7t)
4. [Hướng dẫn từng bước](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#4-h%C6%B0%E1%BB%9Bng-d%E1%BA%ABn-t%E1%BB%ABng-b%C6%B0%E1%BB%9Bc)
5. [Giải thích chi tiết](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#5-gi%E1%BA%A3i-th%C3%ADch-chi-ti%E1%BA%BFt)
6. [Các pattern được sử dụng](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#6-c%C3%A1c-pattern-%C4%91%C6%B0%E1%BB%A3c-s%E1%BB%AD-d%E1%BB%A5ng)
7. [Testing và Debug](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#7-testing-v%C3%A0-debug)
8. [Mở rộng hệ thống](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#8-m%E1%BB%9F-r%E1%BB%99ng-h%E1%BB%87-th%E1%BB%91ng)

---

## 1. Giới thiệu

### 🎯 Mục tiêu bài tập

Xây dựng **hệ thống quản lý thư viện** hoàn chỉnh sử dụng **Map và Struct** trong Go. Hệ thống cho phép:

✅ **Quản lý sách**: Thêm, xóa, sửa, tìm kiếm sách  
✅ **Quản lý người mượn**: Đăng ký, xem danh sách người mượn  
✅ **Quản lý giao dịch**: Mượn sách, trả sách, xem lịch sử  
✅ **Tìm kiếm nâng cao**: Tìm theo tiêu đề hoặc tác giả

### 🏗️ Kiến thức áp dụng

- **Map**: Lưu trữ và truy xuất dữ liệu nhanh
- **Struct**: Định nghĩa các đối tượng nghiệp vụ
- **Pointer**: Quản lý bộ nhớ hiệu quả
- **Error Handling**: Xử lý lỗi chuyên nghiệp
- **Package**: Tổ chức code theo module
- **Constructor Pattern**: Khởi tạo đối tượng
- **Time**: Xử lý ngày giờ

### 📊 Yêu cầu nghiệp vụ

**Quản lý sách:**

- Mỗi sách có: ID (tự sinh), Tiêu đề, Tác giả, Trạng thái (có sẵn/đã mượn)
- Tìm kiếm sách theo tiêu đề hoặc tác giả
- Không được trùng ID sách

**Quản lý người mượn:**

- Mỗi người mượn có: ID (tự sinh), Tên, Email
- Không được trùng ID người mượn

**Quản lý giao dịch:**

- Mỗi giao dịch có: ID, ID người mượn, ID sách, Ngày mượn, Ngày trả
- Kiểm tra sách có sẵn trước khi cho mượn
- Cập nhật trạng thái sách khi mượn/trả
- Xem lịch sử mượn của một người cụ thể

---

## 2. Kiến trúc hệ thống

### 📁 Cấu trúc thư mục

```
Go_Tutorial/Test/
├── main.go                    # Entry point
├── go.mod                     # Module definition
├── models/                    # Data models
│   └── models.go             # Book, Borrower, Transaction
├── library/                   # Business logic
│   ├── LibraryStore.go       # Data storage & operations
│   └── LibraryService.go     # User interface & validation
└── utils/                     # Utility functions
    └── util.go               # Input, validation, ID generation
```

### 🔄 Luồng xử lý dữ liệu

```
┌─────────────┐
│   main.go   │  ← Entry Point
└──────┬──────┘
       │
       ▼
┌──────────────────────┐
│  LibraryService.go   │  ← UI Layer (Input/Output)
└──────────┬───────────┘
           │ Validate & Format
           ▼
┌──────────────────────┐
│   LibraryStore.go    │  ← Business Logic Layer
└──────────┬───────────┘
           │ CRUD Operations
           ▼
┌──────────────────────┐
│   Map[string]Struct  │  ← Data Layer
└──────────────────────┘
```

### 🎨 Thiết kế kiến trúc

**Phân tách trách nhiệm (Separation of Concerns):**

|Layer|File|Trách nhiệm|
|---|---|---|
|**Presentation**|`main.go`|Menu, điều hướng|
|**Service**|`LibraryService.go`|Nhập liệu, validation, format output|
|**Business**|`LibraryStore.go`|Logic nghiệp vụ, xử lý dữ liệu|
|**Data**|`models.go`|Định nghĩa cấu trúc dữ liệu|
|**Utility**|`util.go`|Hàm tiện ích dùng chung|

---

## 3. Chuẩn bị và cài đặt

### Bước 1: Khởi tạo Project

```bash
# Tạo thư mục dự án
mkdir -p Go_Tutorial/Test
cd Go_Tutorial/Test

# Khởi tạo Go module
go mod init hello

# Cài đặt package UUID (để tạo ID tự động)
go get github.com/google/uuid
```

### Bước 2: Kiểm tra go.mod

File `go.mod` sẽ có nội dung:

```
module hello

go 1.21  // hoặc version bạn đang dùng

require github.com/google/uuid v1.6.0
```

### Bước 3: Tạo cấu trúc thư mục

```bash
# Tạo các thư mục
mkdir models library utils

# Tạo các file
touch models/models.go
touch library/LibraryStore.go
touch library/LibraryService.go
touch utils/util.go
touch main.go
```

---

## 4. Hướng dẫn từng bước

### Bước 1: Định nghĩa Data Models

**File:** `models/models.go`

```go
package models

import "time"

// Book: Cấu trúc dữ liệu sách
type Book struct {
	ID     string // UUID tự sinh
	Title  string // Tiêu đề sách
	Author string // Tác giả
	Status bool   // true: có sẵn, false: đã được mượn
}

// Borrower: Cấu trúc người mượn sách
type Borrower struct {
	ID    string // UUID tự sinh
	Name  string // Tên người mượn
	Email string // Email liên hệ
}

// Transaction: Cấu trúc giao dịch mượn sách
type Transaction struct {
	ID         string    // UUID tự sinh
	BorrowerID string    // ID người mượn (Foreign Key)
	BookID     string    // ID sách (Foreign Key)
	BorrowDate time.Time // Thời điểm mượn (tự động gán)
	ReturnDate time.Time // Thời điểm trả (mặc định zero value)
}
```

**📝 Giải thích:**

- **Book.Status**: `true` = có sẵn để mượn, `false` = đã được mượn
- **Transaction.BorrowDate**: Tự động gán `time.Now()` khi tạo giao dịch
- **Transaction.ReturnDate**: Mặc định là zero value, cập nhật khi trả sách
- Các ID đều dùng **UUID** để đảm bảo tính duy nhất

---

### Bước 2: Xây dựng Data Layer (LibraryStore)

**File:** `library/LibraryStore.go`

```go
package library

import (
	"fmt"
	"hello/models"
	"strings"
	"time"
)

// Library: Struct chứa toàn bộ dữ liệu của thư viện
type Library struct {
	books        map[string]models.Book        // Key: BookID
	borrowers    map[string]models.Borrower    // Key: BorrowerID
	transactions map[string]models.Transaction // Key: TransactionID
}

// NewLibrary: Constructor - Khởi tạo Library mới
func NewLibrary() *Library {
	return &Library{
		books:        make(map[string]models.Book),
		borrowers:    make(map[string]models.Borrower),
		transactions: make(map[string]models.Transaction),
	}
}

// ═══════════════════════════════════════════════════════════
//                    QUẢN LÝ SÁCH
// ═══════════════════════════════════════════════════════════

// AddBookStore: Thêm sách vào thư viện
func (lib *Library) AddBookStore(id, title, author string) error {
	// Kiểm tra ID đã tồn tại chưa
	if _, exist := lib.books[id]; exist {
		return fmt.Errorf("ID sách đã trùng: %s", id)
	}

	// Thêm sách vào map
	lib.books[id] = models.Book{
		ID:     id,
		Title:  title,
		Author: author,
		Status: true, // Mặc định có sẵn
	}

	return nil
}

// UpdateBookStore: Cập nhật thông tin sách
func (lib *Library) UpdateBookStore(id, title, author string) error {
	// Kiểm tra sách có tồn tại không
	book, exist := lib.books[id]
	if !exist {
		return fmt.Errorf("không tìm thấy sách với ID: %s", id)
	}

	// Cập nhật thông tin
	book.Title = title
	book.Author = author
	lib.books[id] = book

	return nil
}

// DeleteBookStore: Xóa sách khỏi thư viện
func (lib *Library) DeleteBookStore(id string) error {
	// Kiểm tra sách có tồn tại không
	if _, exist := lib.books[id]; !exist {
		return fmt.Errorf("không tìm thấy sách với ID: %s", id)
	}

	// Kiểm tra sách có đang được mượn không
	if !lib.books[id].Status {
		return fmt.Errorf("không thể xóa sách đang được mượn: %s", id)
	}

	// Xóa sách
	delete(lib.books, id)
	return nil
}

// SearchBookByTitleOrAuthor_Store: Tìm kiếm sách theo tiêu đề hoặc tác giả
func (lib *Library) SearchBookByTitleOrAuthor_Store(query string) []models.Book {
	// Chuyển query sang chữ thường để tìm kiếm không phân biệt hoa/thường
	query = strings.ToLower(query)

	var result []models.Book

	// Duyệt tất cả sách
	for _, book := range lib.books {
		titleLower := strings.ToLower(book.Title)
		authorLower := strings.ToLower(book.Author)

		// Kiểm tra có chứa query không
		if strings.Contains(titleLower, query) ||
			strings.Contains(authorLower, query) {
			result = append(result, book)
		}
	}

	return result
}

// ═══════════════════════════════════════════════════════════
//                 QUẢN LÝ NGƯỜI MƯỢN
// ═══════════════════════════════════════════════════════════

// AddBorrowerStore: Thêm người mượn sách
func (lib *Library) AddBorrowerStore(id, name, email string) error {
	// Kiểm tra ID đã tồn tại chưa
	if _, exist := lib.borrowers[id]; exist {
		return fmt.Errorf("ID người mượn đã trùng: %s", id)
	}

	// Thêm người mượn
	lib.borrowers[id] = models.Borrower{
		ID:    id,
		Name:  name,
		Email: email,
	}

	return nil
}

// FindBorrowerByID: Tìm người mượn theo ID
func (lib *Library) FindBorrowerByID(id string) *models.Borrower {
	if borrower, exist := lib.borrowers[id]; exist {
		return &borrower
	}
	return nil
}

// ═══════════════════════════════════════════════════════════
//                 QUẢN LÝ GIAO DỊCH
// ═══════════════════════════════════════════════════════════

// AddTransactionStore: Tạo giao dịch mượn sách
func (lib *Library) AddTransactionStore(id, borrowerID, bookID string) error {
	// 1. Kiểm tra sách có tồn tại không
	book, bookExist := lib.books[bookID]
	if !bookExist {
		return fmt.Errorf("sách không tồn tại: %s", bookID)
	}

	// 2. Kiểm tra sách có sẵn không
	if !book.Status {
		return fmt.Errorf("sách đã được mượn: %s", bookID)
	}

	// 3. Kiểm tra người mượn có tồn tại không
	if _, exist := lib.borrowers[borrowerID]; !exist {
		return fmt.Errorf("người mượn không tồn tại: %s", borrowerID)
	}

	// 4. Cập nhật trạng thái sách
	book.Status = false
	lib.books[bookID] = book

	// 5. Tạo giao dịch mượn sách
	lib.transactions[id] = models.Transaction{
		ID:         id,
		BorrowerID: borrowerID,
		BookID:     bookID,
		BorrowDate: time.Now(),
		// ReturnDate để zero value (chưa trả)
	}

	return nil
}

// ReturnBook_store: Xử lý trả sách
func (lib *Library) ReturnBook_store(transactionID string) error {
	// 1. Kiểm tra giao dịch có tồn tại không
	trans, exist := lib.transactions[transactionID]
	if !exist {
		return fmt.Errorf("giao dịch không tồn tại: %s", transactionID)
	}

	// 2. Kiểm tra sách đã được trả chưa
	book := lib.books[trans.BookID]
	if book.Status {
		return fmt.Errorf("sách đã được trả rồi")
	}

	// 3. Cập nhật trạng thái sách
	book.Status = true
	lib.books[trans.BookID] = book

	// 4. Cập nhật ngày trả sách
	trans.ReturnDate = time.Now()
	lib.transactions[transactionID] = trans

	return nil
}

// DetailHistoryOfBorrowingBooks_Store: Lấy lịch sử mượn sách của một người
func (lib *Library) DetailHistoryOfBorrowingBooks_Store(borrowerID string) []models.Transaction {
	// Kiểm tra người mượn có tồn tại không
	if _, exist := lib.borrowers[borrowerID]; !exist {
		return nil
	}

	// Lọc các giao dịch của người này
	var result []models.Transaction
	for _, trans := range lib.transactions {
		if trans.BorrowerID == borrowerID {
			result = append(result, trans)
		}
	}

	return result
}
```

**📝 Giải thích quan trọng:**

**1. Constructor Pattern:**

```go
func NewLibrary() *Library {
    return &Library{
        books: make(map[string]models.Book),
        // ...
    }
}
```

- Trả về **pointer** để tránh copy dữ liệu
- Khởi tạo tất cả map trước khi dùng (tránh nil map panic)

**2. Method với Pointer Receiver:**

```go
func (lib *Library) AddBookStore(...) error {
    lib.books[id] = book // Thay đổi trực tiếp
}
```

- Dùng `*Library` thay vì `Library`
- Cho phép thay đổi dữ liệu gốc

**3. Validation nhiều lớp:**

```go
// Kiểm tra tồn tại
if _, exist := lib.books[id]; exist {
    return error
}

// Kiểm tra trạng thái
if !book.Status {
    return error
}
```

---

### Bước 3: Xây dựng Service Layer

**File:** `library/LibraryService.go`

```go
package library

import (
	"fmt"
	"hello/utils"
)

// ═══════════════════════════════════════════════════════════
//                    QUẢN LÝ SÁCH
// ═══════════════════════════════════════════════════════════

// AddBook: Thêm sách (UI Layer)
func AddBook(lib *Library) error {
	fmt.Println("\n╔════════════════════════════════╗")
	fmt.Println("║       THÊM SÁCH MỚI           ║")
	fmt.Println("╚════════════════════════════════╝")

	// Nhập thông tin
	id := utils.GenerateID()
	title := utils.ReadNonEmptyInput("📖 Nhập tiêu đề sách: ")
	author := utils.ReadNonEmptyInput("✍️  Nhập tác giả: ")

	// Gọi Store để lưu
	if err := lib.AddBookStore(id, title, author); err != nil {
		return err
	}

	// Lấy sách vừa thêm để hiển thị
	book := lib.books[id]
	fmt.Printf("\n✅ Đã thêm sách thành công!\n")
	fmt.Printf("   📚 Tiêu đề: %s\n", book.Title)
	fmt.Printf("   ✍️  Tác giả: %s\n", book.Author)
	fmt.Printf("   🆔 ID: %s\n", book.ID)

	return nil
}

// EditBook: Sửa thông tin sách
func EditBook(lib *Library) error {
	fmt.Println("\n╔════════════════════════════════╗")
	fmt.Println("║      CHỈNH SỬA SÁCH           ║")
	fmt.Println("╚════════════════════════════════╝")

	// Kiểm tra có sách không
	if len(lib.books) == 0 {
		return fmt.Errorf("thư viện chưa có sách nào")
	}

	// Hiển thị danh sách sách
	fmt.Println("\n📚 Danh sách sách hiện có:")
	for id, book := range lib.books {
		status := "✅ Có sẵn"
		if !book.Status {
			status = "❌ Đã mượn"
		}
		fmt.Printf("  [%s] %s - %s (%s)\n", id, book.Title, book.Author, status)
	}

	// Nhập ID sách cần sửa
	id := utils.ReadNonEmptyInput("\n🆔 Nhập ID sách cần sửa: ")

	// Kiểm tra sách tồn tại
	book, exist := lib.books[id]
	if !exist {
		return fmt.Errorf("không tìm thấy sách với ID: %s", id)
	}

	// Hiển thị thông tin hiện tại
	fmt.Printf("\n📋 Thông tin hiện tại:\n")
	fmt.Printf("   Tiêu đề: %s\n", book.Title)
	fmt.Printf("   Tác giả: %s\n", book.Author)

	// Nhập thông tin mới
	fmt.Println("\n✏️  Nhập thông tin mới (Enter để giữ nguyên):")
	newTitle := utils.ReadInput("📖 Tiêu đề mới: ")
	if newTitle == "" {
		newTitle = book.Title
	}

	newAuthor := utils.ReadInput("✍️  Tác giả mới: ")
	if newAuthor == "" {
		newAuthor = book.Author
	}

	// Cập nhật
	if err := lib.UpdateBookStore(id, newTitle, newAuthor); err != nil {
		return err
	}

	fmt.Println("\n✅ Cập nhật sách thành công!")
	return nil
}

// DeleteBook: Xóa sách
func DeleteBook(lib *Library) error {
	fmt.Println("\n╔════════════════════════════════╗")
	fmt.Println("║         XÓA SÁCH              ║")
	fmt.Println("╚════════════════════════════════╝")

	// Kiểm tra có sách không
	if len(lib.books) == 0 {
		return fmt.Errorf("thư viện chưa có sách nào")
	}

	// Hiển thị danh sách sách
	fmt.Println("\n📚 Danh sách sách:")
	for id, book := range lib.books {
		status := "✅ Có sẵn"
		if !book.Status {
			status = "❌ Đang mượn"
		}
		fmt.Printf("  [%s] %s - %s (%s)\n", id, book.Title, book.Author, status)
	}

	// Nhập ID sách cần xóa
	id := utils.ReadNonEmptyInput("\n🆔 Nhập ID sách cần xóa: ")

	// Xác nhận
	confirm := utils.ReadInput("⚠️  Bạn có chắc chắn muốn xóa? (y/n): ")
	if confirm != "y" && confirm != "Y" {
		fmt.Println("❌ Đã hủy thao tác xóa")
		return nil
	}

	// Xóa sách
	if err := lib.DeleteBookStore(id); err != nil {
		return err
	}

	fmt.Println("\n✅ Xóa sách thành công!")
	return nil
}

// ListBooks: Hiển thị danh sách sách
func ListBooks(lib *Library) error {
	fmt.Println("\n╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║              DANH SÁCH TẤT CẢ SÁCH                        ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")

	if len(lib.books) == 0 {
		fmt.Println("📚 Thư viện hiện đang trống. Hãy thêm sách để bắt đầu!")
		return nil
	}

	// Header
	fmt.Printf("\n%-38s %-30s %-20s %-12s\n",
		"ID", "Tiêu đề", "Tác giả", "Trạng thái")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	// Data
	for id, book := range lib.books {
		status := "✅ Có sẵn"
		if !book.Status {
			status = "❌ Đã mượn"
		}
		fmt.Printf("%-38s %-30s %-20s %-12s\n",
			id, book.Title, book.Author, status)
	}

	fmt.Printf("\n📊 Tổng số sách: %d\n", len(lib.books))
	return nil
}

// SearchBook: Tìm kiếm sách theo tiêu đề chính xác
func SearchBook(lib *Library) error {
	if len(lib.books) == 0 {
		return fmt.Errorf("thư viện chưa có sách nào")
	}

	search := utils.ReadNonEmptyInput("🔍 Nhập tiêu đề sách cần tìm: ")

	// Tìm kiếm chính xác
	for _, book := range lib.books {
		if book.Title == search {
			fmt.Printf("\n✅ Tìm thấy sách:\n")
			fmt.Printf("   🆔 ID: %s\n", book.ID)
			fmt.Printf("   📖 Tiêu đề: %s\n", book.Title)
			fmt.Printf("   ✍️  Tác giả: %s\n", book.Author)
			status := "Có sẵn"
			if !book.Status {
				status = "Đã mượn"
			}
			fmt.Printf("   📊 Trạng thái: %s\n", status)
			return nil
		}
	}

	return fmt.Errorf("không tìm thấy sách: %s", search)
}

// SearchBookByTitleOrAuthor: Tìm kiếm nâng cao
func SearchBookByTitleOrAuthor(lib *Library) error {
	fmt.Println("\n╔════════════════════════════════╗")
	fmt.Println("║      TÌM KIẾM NÂNG CAO        ║")
	fmt.Println("╚════════════════════════════════╝")

	query := utils.ReadNonEmptyInput("🔍 Nhập từ khóa (tiêu đề hoặc tác giả): ")

	result := lib.SearchBookByTitleOrAuthor_Store(query)

	if len(result) == 0 {
		fmt.Printf("\n❌ Không tìm thấy sách nào với từ khóa: %s\n", query)
		return nil
	}

	// Hiển thị kết quả
	fmt.Printf("\n✅ Tìm thấy %d sách:\n", len(result))
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	for i, book := range result {
		status := "Có sẵn"
		if !book.Status {
			status = "Đã mượn"
		}
		fmt.Printf("%d. %s - %s (%s)\n", i+1, book.Title, book.Author, status)
	}

	return nil
}

// ═══════════════════════════════════════════════════════════
//                 QUẢN LÝ NGƯỜI MƯỢN
// ═══════════════════════════════════════════════════════════

// AddBorrower: Thêm người mượn sách
func AddBorrower(lib *Library) error {
	fmt.Println("\n╔════════════════════════════════╗")
	fmt.Println("║     THÊM NGƯỜI MƯỢN SÁCH      ║")
	fmt.Println("╚════════════════════════════════╝")

	id := utils.GenerateID()
	name := utils.ReadNonEmptyInput("👤 Nhập tên người mượn: ")
	email := utils.ReadNonEmptyInput("📧 Nhập email: ")

	if err := lib.AddBorrowerStore(id, name, email); err != nil {
		return err
	}

	borrower := lib.borrowers[id]
	fmt.Printf("\n✅ Đã thêm người mượn thành công!\n")
	fmt.Printf("   👤 Tên: %s\n", borrower.Name)
	fmt.Printf("   📧 Email: %s\n", borrower.Email)
	fmt.Printf("   🆔 ID: %s\n", borrower.ID)

	return nil
}

// ListBorrowers: Hiển thị danh sách người mượn
func ListBorrowers(lib *Library) error {
	fmt.Println("\n╔════════════════════════════════════════════════════╗")
	fmt.Println("║         DANH SÁCH NGƯỜI MƯỢN SÁCH                 ║")
	fmt.Println("╚════════════════════════════════════════════════════╝")

	if len(lib.borrowers) == 0 {
		fmt.Println("👥 Chưa có người mượn nào đăng ký!")
		return nil
	}

	fmt.Printf("\n%-38s %-25s %-30s\n", "ID", "Tên", "Email")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	for _, borrower := range lib.borrowers {
		fmt.Printf("%-38s %-25s %-30s\n",
			borrower.ID, borrower.Name, borrower.Email)
	}

	fmt.Printf("\n📊 Tổng số người mượn: %d\n", len(lib.borrowers))
	return nil
}

// ═══════════════════════════════════════════════════════════
//                 QUẢN LÝ GIAO DỊCH
// ═══════════════════════════════════════════════════════════

// BorrowBook: Mượn sách
func BorrowBook(lib *Library) error {
	fmt.Println("\n╔════════════════════════════════╗")
	fmt.Println("║          MƯỢN SÁCH            ║")
	fmt.Println("╚════════════════════════════════╝")

	borrowerID := utils.ReadNonEmptyInput("👤 Nhập ID người mượn: ")
	bookID := utils.ReadNonEmptyInput("📚 Nhập ID sách: ")

	transID := utils.GenerateID()

	if err := lib.AddTransactionStore(transID, borrowerID, bookID); err != nil {
		return err
	}

	fmt.Println("\n✅ Mượn sách thành công!")
	fmt.Printf("   🆔 Mã giao dịch: %s\n", transID)
	return nil
}

// ReturnBook: Trả sách
func ReturnBook(lib *Library) error {
	fmt.Println("\n╔════════════════════════════════╗")
	fmt.Println("║           TRẢ SÁCH            ║")
	fmt.Println("╚════════════════════════════════╝")

	transID := utils.ReadNonEmptyInput("🆔 Nhập ID giao dịch: ")

	if err := lib.ReturnBook_store(transID); err != nil {
		return err
	}

	fmt.Println("\n✅ Trả sách thành công!")
	return nil
}

// HistoryOfBorrowingBooks: Xem tất cả lịch sử mượn sách
func HistoryOfBorrowingBooks(lib *Library) error {
	fmt.Println("\n╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║           LỊCH SỬ MƯỢN SÁCH (TẤT CẢ)                     ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")

	if len(lib.transactions) == 0 {
		fmt.Println("📋 Chưa có giao dịch nào!")
		return nil
	}

	fmt.Printf("\n%-38s %-38s %-38s %-20s %-12s\n",
		"ID Giao dịch", "ID Người mượn", "ID Sách", "Ngày mượn", "Trạng thái")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	for _, trans := range lib.transactions {
		status := "❌ Chưa trả"
		book := lib.books[trans.BookID]
		if book.Status {
			status = "✅ Đã trả"
		}

		fmt.Printf("%-38s %-38s %-38s %-20s %-12s\n",
			trans.ID,
			trans.BorrowerID,
			trans.BookID,
			trans.BorrowDate.Format("02/01/2006 15:04"),
			status)
	}

	fmt.Printf("\n📊 Tổng số giao dịch: %d\n", len(lib.transactions))
	return nil
}

// DetailHistoryOfBorrowingBooks: Xem lịch sử mượn của một người
func DetailHistoryOfBorrowingBooks(lib *Library) error {
	fmt.Println("\n╔════════════════════════════════╗")
	fmt.Println("║    LỊCH SỬ CÁ NHÂN           ║")
	fmt.Println("╚════════════════════════════════╝")

	if len(lib.transactions) == 0 {
		return fmt.Errorf("chưa có giao dịch nào")
	}

	borrowerID := utils.ReadNonEmptyInput("👤 Nhập ID người mượn: ")

	list := lib.DetailHistoryOfBorrowingBooks_Store(borrowerID)
	if list == nil || len(list) == 0 {
		return fmt.Errorf("không tìm thấy lịch sử mượn sách của người này")
	}

	// Hiển thị thông tin người mượn
	borrower := lib.FindBorrowerByID(borrowerID)
	if borrower != nil {
		fmt.Printf("\n👤 Người mượn: %s\n", borrower.Name)
		fmt.Printf("📧 Email: %s\n", borrower.Email)
	}

	fmt.Println("\n📚 Lịch sử mượn sách:")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	for i, trans := range list {
		book := lib.books[trans.BookID]
		status := "❌ Chưa trả"
		returnDate := "N/A"

		if book.Status {
			status = "✅ Đã trả"
			returnDate = trans.ReturnDate.Format("02/01/2006 15:04")
		}

		fmt.Printf("\n%d. Sách: %s\n", i+1, book.Title)
		fmt.Printf("   Mượn ngày: %s\n", trans.BorrowDate.Format("02/01/2006 15:04"))
		fmt.Printf("   Trả ngày: %s\n", returnDate)
		fmt.Printf("   Trạng thái: %s\n", status)
	}

	return nil
}
```

---

### Bước 4: Xây dựng Utility Layer

**File:** `utils/util.go`

```go
package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

// ReadInput: Đọc input từ bàn phím
func ReadInput(prompt string) string {
	fmt.Printf("%s ", prompt)
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')
	return strings.TrimSpace(value)
}

// ReadNonEmptyInput: Đọc input không được rỗng
func ReadNonEmptyInput(prompt string) string {
	for {
		value := ReadInput(prompt)

		if !IsEmpty(value) {
			return value
		}

		fmt.Println("❌ Giá trị không được để trống")
	}
}

// IsEmpty: Kiểm tra chuỗi rỗng
func IsEmpty(value string) bool {
	return value == "" || len(value) == 0
}

// GetConvertedInt: Đọc và chuyển đổi sang số nguyên
func GetConvertedInt(prompt string) int {
	for {
		input := ReadInput(prompt)

		if IsEmpty(input) {
			fmt.Println("❌ Giá trị không được để trống")
			continue
		}

		value, err := strconv.Atoi(input)

		if err == nil && value > -1 {
			return value
		}

		fmt.Println("❌ Giá trị không hợp lệ hoặc nhỏ hơn 0")
	}
}

// GetConvertedFloat: Đọc và chuyển đổi sang số thực
func GetConvertedFloat(prompt string) float64 {
	for {
		input := ReadInput(prompt)

		if IsEmpty(input) {
			fmt.Println("❌ Giá trị không được để trống")
			continue
		}

		value, err := strconv.ParseFloat(input, 64)

		if err == nil && value > -1 {
			return value
		}

		fmt.Println("❌ Giá trị không hợp lệ hoặc nhỏ hơn 0")
	}
}

// ClearScreen: Xóa màn hình console
func ClearScreen() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println("Error clearing screen:", err)
	}
}

// GenerateID: Tạo UUID ngẫu nhiên
func GenerateID() string {
	return uuid.New().String()
}
```

---

### Bước 5: Xây dựng Main Program

**File:** `main.go`

```go
package main

import (
	"fmt"
	"hello/library"
	"hello/utils"
)

func main() {
	// Khởi tạo thư viện
	lib := library.NewLibrary()

	// Main loop
	for {
		utils.ClearScreen()

		// Menu
		fmt.Println("\n╔════════════════════════════════════════════════╗")
		fmt.Println("║     HỆ THỐNG QUẢN LÝ THƯ VIỆN                ║")
		fmt.Println("╚════════════════════════════════════════════════╝")
		fmt.Println("")
		fmt.Println("📚 QUẢN LÝ SÁCH:")
		fmt.Println("   1. Thêm sách")
		fmt.Println("   2. Xóa sách")
		fmt.Println("   3. Chỉnh sửa sách")
		fmt.Println("   4. Xem danh sách sách")
		fmt.Println("   5. Tìm kiếm sách (chính xác)")
		fmt.Println("   6. Tìm kiếm sách (nâng cao)")
		fmt.Println("")
		fmt.Println("👥 QUẢN LÝ NGƯỜI MƯỢN:")
		fmt.Println("   7. Thêm người mượn")
		fmt.Println("   8. Xem danh sách người mượn")
		fmt.Println("")
		fmt.Println("📋 QUẢN LÝ GIAO DỊCH:")
		fmt.Println("   9. Mượn sách")
		fmt.Println("  10. Trả sách")
		fmt.Println("  11. Xem lịch sử mượn sách (tất cả)")
		fmt.Println("  12. Xem lịch sử cá nhân")
		fmt.Println("")
		fmt.Println("  0. 🚪 Thoát")
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

		choice := utils.GetConvertedInt("👉 Chọn chức năng: ")

		var err error

		switch choice {
		case 1:
			err = library.AddBook(lib)
		case 2:
			err = library.DeleteBook(lib)
		case 3:
			err = library.EditBook(lib)
		case 4:
			err = library.ListBooks(lib)
		case 5:
			err = library.SearchBook(lib)
		case 6:
			err = library.SearchBookByTitleOrAuthor(lib)
		case 7:
			err = library.AddBorrower(lib)
		case 8:
			err = library.ListBorrowers(lib)
		case 9:
			err = library.BorrowBook(lib)
		case 10:
			err = library.ReturnBook(lib)
		case 11:
			err = library.HistoryOfBorrowingBooks(lib)
		case 12:
			err = library.DetailHistoryOfBorrowingBooks(lib)
		case 0:
			fmt.Println("\n👋 Cảm ơn bạn đã sử dụng hệ thống!")
			fmt.Println("🔚 Thoát chương trình...")
			return
		default:
			fmt.Println("❌ Lựa chọn không hợp lệ. Vui lòng chọn lại.")
		}

		// Hiển thị lỗi nếu có
		if err != nil {
			fmt.Printf("\n❌ Lỗi: %v\n", err)
		}

		utils.ReadInput("\n⏸️  Nhấn Enter để tiếp tục...")
	}
}
```

---

## 5. Giải thích chi tiết

### 5.1. Tại sao dùng Map?

**So sánh với Slice:**

|Thao tác|Slice|Map|
|---|---|---|
|Tìm theo ID|O(n) - phải duyệt|O(1) - truy cập trực tiếp|
|Thêm|O(1)|O(1)|
|Xóa|O(n) - phải dịch chuyển|O(1)|
|Kiểm tra tồn tại|O(n)|O(1)|

**Ví dụ:**

```go
// Với Slice - phải duyệt
func FindBookByID_Slice(id string, books []Book) *Book {
    for i := range books {
        if books[i].ID == id {
            return &books[i]
        }
    }
    return nil  // O(n) - chậm
}

// Với Map - truy cập trực tiếp
func FindBookByID_Map(id string, books map[string]Book) *Book {
    if book, ok := books[id]; ok {
        return &book
    }
    return nil  // O(1) - nhanh
}
```

### 5.2. Tại sao dùng Pointer Receiver?

```go
// ❌ SAI - Không thay đổi được dữ liệu gốc
func (lib Library) AddBook(book Book) {
    lib.books[book.ID] = book  // Chỉ thay đổi bản copy
}

// ✅ ĐÚNG - Thay đổi dữ liệu gốc
func (lib *Library) AddBook(book Book) {
    lib.books[book.ID] = book  // Thay đổi trực tiếp
}
```

### 5.3. Error Handling Pattern

```go
func (lib *Library) AddBookStore(...) error {
    // Validation
    if _, exist := lib.books[id]; exist {
        return fmt.Errorf("ID đã tồn tại")
    }
    
    // Business logic
    lib.books[id] = book
    
    // Success
    return nil
}

// Sử dụng
if err := lib.AddBookStore(...); err != nil {
    fmt.Printf("Lỗi: %v\n", err)
}
```

### 5.4. Cập nhật Struct trong Map

**Vấn đề:**

```go
// ❌ KHÔNG hoạt động
lib.books[id].Status = false  // Compile error!
```

**Giải pháp:**

```go
// ✅ Cách 1: Lấy ra, sửa, gán lại
book := lib.books[id]
book.Status = false
lib.books[id] = book

// ✅ Cách 2: Dùng pointer trong map
// map[string]*Book thay vì map[string]Book
```

---

## 6. Các Pattern được sử dụng

### 6.1. Constructor Pattern

```go
func NewLibrary() *Library {
    return &Library{
        books: make(map[string]models.Book),
        // ...
    }
}
```

**Lợi ích:**

- Đảm bảo khởi tạo đúng
- Tránh nil map panic
- Dễ mở rộng

### 6.2. Repository Pattern

```go
// LibraryStore.go = Repository (Data Access Layer)
func (lib *Library) AddBookStore(...) error { }
func (lib *Library) FindBookByID(...) *Book { }

// LibraryService.go = Service Layer (Business Logic)
func AddBook(lib *Library) error { }
```

**Lợi ích:**

- Tách biệt data và UI
- Dễ test
- Dễ maintain

### 6.3. Error Handling Pattern

```go
// Trả về error thay vì panic
func DoSomething() error {
    if /* có lỗi */ {
        return fmt.Errorf("mô tả lỗi")
    }
    return nil
}
```

---

## 7. Testing và Debug

### 7.1. Test Case cơ bản

```go
package library

import "testing"

func TestAddBook(t *testing.T) {
    lib := NewLibrary()
    
    err := lib.AddBookStore("id1", "Go Programming", "John Doe")
    if err != nil {
        t.Errorf("AddBook failed: %v", err)
    }
    
    // Kiểm tra sách đã thêm
    if len(lib.books) != 1 {
        t.Errorf("Expected 1 book, got %d", len(lib.books))
    }
}

func TestBorrowBook(t *testing.T) {
    lib := NewLibrary()
    
    // Setup
    lib.AddBookStore("book1", "Test Book", "Author")
    lib.AddBorrowerStore("borrower1", "John", "john@email.com")
    
    // Test
    err := lib.AddTransactionStore("trans1", "borrower1", "book1")
    if err != nil {
        t.Errorf("BorrowBook failed: %v", err)
    }
    
    // Verify
    if lib.books["book1"].Status != false {
        t.Error("Book status should be false after borrowing")
    }
}
```

### 7.2. Debug Tips

```go
// Thêm log để debug
import "log"

func (lib *Library) AddBookStore(...) error {
    log.Printf("Adding book: %s - %s", title, author)
    
    // ...
    
    log.Printf("Book added successfully: %s", id)
    return nil
}
```

---

## 8. Mở rộng hệ thống

### 8.1. Thêm tính năng Phạt trễ hạn

```go
type Transaction struct {
    ID         string
    BorrowerID string
    BookID     string
    BorrowDate time.Time
    ReturnDate time.Time
    DueDate    time.Time    // Thêm: Hạn trả
    Fine       float64      // Thêm: Tiền phạt
}

func CalculateFine(trans Transaction) float64 {
    if trans.ReturnDate.IsZero() {
        return 0
    }
    
    overdueDays := int(trans.ReturnDate.Sub(trans.DueDate).Hours() / 24)
    if overdueDays > 0 {
        return float64(overdueDays) * 5000  // 5000 VND/ngày
    }
    return 0
}
```

### 8.2. Lưu dữ liệu vào File (JSON)

```go
import (
    "encoding/json"
    "os"
)

func (lib *Library) SaveToFile(filename string) error {
    data := map[string]interface{}{
        "books":        lib.books,
        "borrowers":    lib.borrowers,
        "transactions": lib.transactions,
    }
    
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    
    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ")
    return encoder.Encode(data)
}

func (lib *Library) LoadFromFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    
    var data map[string]interface{}
    decoder := json.NewDecoder(file)
    return decoder.Decode(&data)
}
```

### 8.3. Thêm Database (SQLite)

```go
import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, error) {
    db, err := sql.Open("sqlite3", "./library.db")
    if err != nil {
        return nil, err
    }
    
    // Tạo bảng
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS books (
            id TEXT PRIMARY KEY,
            title TEXT,
            author TEXT,
            status BOOLEAN
        )
    `)
    
    return db, err
}
```

### 8.4. Thêm API REST

```go
import (
    "encoding/json"
    "net/http"
)

func (lib *Library) HandleGetBooks(w http.ResponseWriter, r *http.Request) {
    books := []models.Book{}
    for _, book := range lib.books {
        books = append(books, book)
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}

func main() {
    lib := NewLibrary()
    
    http.HandleFunc("/api/books", lib.HandleGetBooks)
    http.ListenAndServe(":8080", nil)
}
```

---

## 📚 Tài liệu tham khảo

- [Go Maps](https://go.dev/blog/maps)
- [Go Structs](https://go.dev/tour/moretypes/2)
- [Error Handling](https://go.dev/blog/error-handling-and-go)
- [UUID Package](https://github.com/google/uuid)