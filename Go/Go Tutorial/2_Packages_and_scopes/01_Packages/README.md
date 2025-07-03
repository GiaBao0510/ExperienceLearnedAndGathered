# Bắt đầu với Go: Làm việc với nhiều file trong cùng Package

Tài liệu này hướng dẫn bạn cách tạo một chương trình Go đơn giản với nhiều file thuộc cùng `package main`. Bạn sẽ học cách tổ chức code, khởi tạo Go module, và chạy chương trình.

## Yêu cầu

Trước tiên, đảm bảo bạn đã cài đặt Go:

1. Tải Go từ golang.org.
2. Chạy lệnh sau để kiểm tra:
```bash
   go version
```

- Nếu thấy phiên bản (ví dụ: `go1.21.0`), bạn đã sẵn sàng!


## Bước 1: Tạo thư mục dự án

Tạo một thư mục có tên `01_Packages`:

```bash
mkdir 01_Packages
cd 01_Packages
```

## Bước 2: Khởi tạo Go Module

Chạy lệnh sau để tạo file `go.mod`, giúp quản lý dự án:

```bash
go mod init example.com/01_Packages
```

- **Kết quả**: File `go.mod` được tạo với nội dung như:
    
    ```go
    module example.com/01_Packages
    
    go 1.21
    ```
    
- **Giải thích**:
    
    - `module example.com/01_Packages`: Tên module (có thể tùy chỉnh).
    - `go 1.21`: Phiên bản Go đang dùng.
    - File này tương tự `package.json` (Node.js) hoặc `pom.xml` (Java).

## Bước 3: Tạo các file Go

Tạo ba file `main.go`, `loading.go`, và `listening.go` trong thư mục `01_Packages`. Bạn có thể dùng bất kỳ editor nào (VS Code, Vim, Notepad, v.v.).

### File `loading.go`

```go
package main

import "fmt"

// loading in ra thông báo khi sự kiện đang tải
func loading() {
    fmt.Println("Loading the event...")
}
```

### File `listening.go`

```go
package main

import "fmt"

// listening in ra thông báo khi sự kiện đang được lắng nghe
func listening() {
    fmt.Println("Listening to the event...")
}
```

### File `main.go`

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")

    // Gọi các hàm từ các file khác trong cùng package
    loading()
    listening()
}
```

**Giải thích**:

- Cả ba file đều thuộc `package main`, nên các hàm `loading()` và `listening()` có thể được gọi trực tiếp từ `main.go`.
- `package main` yêu cầu một hàm `main()` làm điểm bắt đầu của chương trình.

## Bước 4: Build chương trình

Chạy lệnh sau để biên dịch toàn bộ package:

```bash
go build
```

- **Kết quả**: Tạo file thực thi (ví dụ: `01_Packages.exe` trên Windows).
- Chạy file thực thi:

```bash
./01_Packages    # Linux/Mac
  01_Packages.exe  # Windows
```

## Bước 5: Chạy chương trình

Chạy toàn bộ chương trình bằng:

```bash
go run .
```

- **Giải thích**: Lệnh `go run .` biên dịch và chạy tất cả file `.go` trong thư mục.

- **Kết quả**:
```
Hello, World!
Loading the event...
Listening to the event...
```


## Lưu ý quan trọng

- **Không chạy** `go run main.go` **riêng lẻ**: Lệnh này chỉ biên dịch `main.go`, dẫn đến lỗi `undefined: loading` và `undefined: listening` vì thiếu `loading.go` và `listening.go`.
- **Sử dụng** `go run .`: Đảm bảo biên dịch toàn bộ package.
- Nếu gặp lỗi `command not found: go`, hãy kiểm tra lại cài đặt Go.

## Kết luận

Bạn đã tạo một chương trình Go với nhiều file trong cùng `package main`. Tiếp theo, bạn có thể thử:

- Thêm hàm mới vào `loading.go` hoặc `listening.go`.
- Tách code thành nhiều package khác nhau.
- Thêm unit test bằng `go test`.