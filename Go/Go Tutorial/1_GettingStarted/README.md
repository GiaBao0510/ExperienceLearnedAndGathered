
## **Bắt đầu với Go - Dự án đầu tiên**
---
Trước khi bắt đầu, hãy đảm bảo bạn đã cài đặt Go trên máy tính. Nếu chưa, tải Go từ [golang.org](https://golang.org/dl/) và chạy lệnh sau để kiểm tra:

```
go version
```
Nếu thấy phiên bản (ví dụ: go1.21.0), bạn đã sẵn sàng!

### **Tạo thư mục dự án**
---

Tạo một thư mục có tên 1_GettingStarted để lưu trữ dự án đầu tiên:
```
mkdir 1_GettingStarted
cd 1_GettingStarted
```

#### **Tạo dự án:**
---

Sau khi tạo xong thì hãy chạy câu lệnh sau tại thư mục `1_GettingStarted`:

```
go mod init helloworld
```

*Giải thích:*
- Lệnh này tạo một **file go.mod,** là nơi quản lý thông tin dự án và các thư viện phụ thuộc.
- Đây là lệnh để khởi tạo một **Go module**. Tương tự như tạo `package.json` trong Node.js hoặc `csproj` trong .NET .Lệnh này sẽ tạo một file `go.mod` ở thư mục hiện tại.
- **📌 Khi nào cần dùng?**
	- Khi **bắt đầu một project mới**
	- Khi muốn quản lý thư viện ngoài (thrid-party) bằng `go get`
	- Khi muốn dùng `import` theo module thay vì `GOPATH`

- 📦 **Sau khi khởi tạo bạn có thể:**
	- Import packages nội bộ hoặc bên ngoài
	- Dùng `go get` để cải thư viện
	- Dùng `go build`, `go run`, `go test` bình thường

*kết quả sau khi chạy:*
```
1_GettingStarted> go mod init helloworld  
go: creating new go.mod: module helloworld
```

### **Tạo file đầu tiên**
---
Tạo file helloworld.go bằng cách:

- Nếu dùng VS Code: `code helloworld.go`.
- Nếu dùng editor khác: Tạo file thủ công (ví dụ: touch helloworld.go trên Linux/Mac, hoặc dùng Notepad trên Windows).

Mở file và thêm code sau:
```go
//Đây là một goi đặt biệt. Nó cho phép Go tạo một tệp thực thi
package main

/*
	- Từ khóa "import" cung cấp một gói khác cho tệp tin ".go" này
	- import "fmt" cho phép bạn có thể truy cập các chức năng từ gói "fmt" tại tệp tin này.
*/
import "fmt"


/*
	- "func main" là đặt biệt.
	- Go cho biết nơi bắt đầu
	- "func main" tạo một điểm bắt đầu cho Go
	- Sau khi biên dịch mã, Go Runtime sẽ chạy hàm này trước tiên
*/
func main() {
    fmt.Println("Gia Bao Xin Chao")
}
```

#### **Xây dựng và chạy chương trình go:**
---

**Build một chương trình Go:**
```
go build main.go
```
- Lệnh này tạo file thực thi (ví dụ: helloworld trên Linux/Mac, helloworld.exe trên Windows).
- Chạy File thực thi:
```
./helloworld   # Linux/Mac
helloworld.exe # Windows
```

**Run một chương trình Go:**
```
go run main.go
```
- **go run**: Biên dịch và chạy file ngay lập tức.

- Ví dụ chạy chương trình `helloworld.go` ở bài trên:
```
1_GettingStarted> go run helloworld.go
Gia Bao Xin Chao
```

- Nếu bạn tạo nhiều file khác và Run tất cả chúng, bạn có thể sử dụng câu lệnh sau:
```
go run .
```