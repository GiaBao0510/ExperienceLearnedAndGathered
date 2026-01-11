# Hiểu về Module và Package trong Go

Tài liệu này giới thiệu các khái niệm cơ bản về **module** và **package** trong Go, giúp bạn hiểu cách tổ chức code trong một dự án Go.

## Module là gì?

- **Module** là một tập hợp các **package** liên quan, được quản lý bởi file `go.mod`.
- Một module thường đại diện cho một dự án hoặc thư viện, chứa một hoặc nhiều package.
- File `go.mod` định nghĩa:
    - Tên module (thường theo dạng `github.com/username/project`).
    - Phiên bản Go.
    - Các dependency (nếu có).

Ví dụ nội dung file `go.mod`:

```go
module github.com/yourname/myproject

go 1.21
```

**Cách tạo module**:

```bash
go mod init github.com/yourname/myproject
```

- **Lưu ý**: Tên module nên là đường dẫn duy nhất (như URL repository) để tránh xung đột khi chia sẻ hoặc sử dụng ở dự án khác.

## Package là gì?

- **Package** là một thư mục chứa nhiều file `.go` thuộc cùng một nhóm chức năng.
- Có hai loại package chính:
    - **`package main`**: Tạo file thực thi (`.exe` trên Windows hoặc file nhị phân trên Linux/Mac). Yêu cầu có hàm `main()`.
    - **Package khác** (ví dụ: `util`, `database`): Dùng làm thư viện, không tạo file thực thi.
    
- **Quy tắc đặt tên package**:
    - Tên ngắn gọn, chữ thường (ví dụ: `util`, không phải `Util`).
    - Phải khớp với tên thư mục chứa file `.go`.

- **Scope trong package**:
    - Tên hàm/biến bắt đầu bằng **chữ hoa** (ví dụ: `MyFunction`) là **public**, có thể truy cập từ package khác.
    - Tên bắt đầu bằng **chữ thường** (ví dụ: `myFunction`) là **private**, chỉ dùng trong cùng package.

## Minh họa cấu trúc Module và Package

Hình dưới đây (từ [go.dev](https://go.dev/doc/modules/images/source-hierarchy.png)) cho thấy cách tổ chức code trong Go:

![Go Module Hierarchy](https://go.dev/doc/modules/images/source-hierarchy.png)

- Một module chứa nhiều package (thư mục).
- Mỗi package chứa các file `.go` liên quan.

## Ví dụ thực tế

Dưới đây là ví dụ về một module với nhiều package:

### Cấu trúc thư mục

```
myproject/
├── go.mod
├── main.go
├── util/
│   ├── helper.go
```

### File `go.mod`

```go
module github.com/yourname/myproject

go 1.21
```

### File `main.go`

```go
package main

import (
    "fmt"
    "github.com/yourname/myproject/util"
)

func main() {
    fmt.Println("Hello, World!")
    util.PrintMessage() // Gọi hàm từ package util
}
```

### File `util/helper.go`

```go
package util

import "fmt"

// PrintMessage là hàm public
func PrintMessage() {
    fmt.Println("Message from util package")
}
```

### Chạy chương trình

1. Khởi tạo module:
    ```bash
    go mod init github.com/yourname/myproject
    ```

2. Chạy chương trình:
```bash
go run .
```

**Kết quả**:

```
Hello, World!
Message from util package
```


---
## Package `fmt`

- Package `fmt` là một phần của thư viện chuẩn Go, dùng để định dạng và in dữ liệu.
    
- Ví dụ sử dụng:
    
    ```go
    package main
    
    import "fmt"
    
    func main() {
        fmt.Println("Hello from fmt!")
    }
    ```
    
---
## Lưu ý khi làm việc với Module và Package

- **Khởi tạo module**: Luôn chạy `go mod init` trước khi làm việc với dự án.
- **Import package**:
    - Package trong cùng module: Dùng đường dẫn tương đối (ví dụ: `import "./util"`).
    - Package bên ngoài: Dùng `go get github.com/author/package`.
- **Lỗi thường gặp**:
    - Chạy `go run main.go` thay vì `go run .` khi có nhiều file trong `package main` → Lỗi `undefined`.
    - Giải pháp: Dùng `go run .` để biên dịch toàn bộ package.
- **Quản lý dependency**: Chạy `go mod tidy` để dọn dẹp file `go.mod`.

---
## Kết luận
- **Module** giúp quản lý dự án và các package.
- **Package** tổ chức code thành các nhóm chức năng.
- Tiếp theo, bạn có thể thử:
    - Tạo module với nhiều package.
    - Thêm dependency bên ngoài bằng `go get`.
    - Viết unit test cho package bằng `go test`.