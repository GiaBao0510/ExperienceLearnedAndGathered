**Hướng dẫn biên dịch và chạy ứng dụng Go**

Tài liệu này sẽ giải thích các cách biên dịch và chạy mã nguồn Go, cùng với những lưu ý quan trọng.

### 1. `go run [tên tệp tin cụ thể].go`

- **Mục đích:** Biên dịch và chạy một tệp tin Go cụ thể.
- **Cách hoạt động:** Lệnh này sẽ tìm hàm `main` trong tệp tin được chỉ định, sau đó biên dịch và thực thi ứng dụng.
- **Lưu ý:** Lệnh này _không_ tạo ra tệp thực thi độc lập (ví dụ: `.exe`).

### 2. `go run .`

- **Mục đích:** Biên dịch và chạy ứng dụng từ thư mục hiện hành.
- **Cách hoạt động:** Lệnh này sẽ quét toàn bộ các tệp tin `.go` trong thư mục hiện tại để tìm hàm `main`, sau đó biên dịch và thực thi ứng dụng.
- **Lưu ý:** Tương tự như `go run [tên tệp tin cụ thể].go`, lệnh này cũng _không_ tạo ra tệp thực thi độc lập.

### 3. `go build .`

- **Mục đích:** Biên dịch mã nguồn Go và tạo ra tệp thực thi.
- **Cách hoạt động:** Lệnh này sẽ biên dịch tất cả các tệp `.go` trong thư mục hiện tại và tạo ra một tệp thực thi (ví dụ: `.exe` trên Windows, hoặc một tệp không có phần mở rộng trên Linux/macOS) trong cùng thư mục.
- **Lưu ý:** Mỗi khi có sự thay đổi trong mã nguồn (các tệp `.go`), bạn cần chạy lại lệnh này để cập nhật tệp thực thi.

### Lưu ý quan trọng khi sử dụng `go run .` hoặc `go build .`

Để sử dụng các lệnh `go run .` hoặc `go build .`, dự án Go của bạn cần phải có tệp `go.mod`.

Nếu chưa có, bạn cần chạy lệnh sau trong thư mục gốc của dự án để khởi tạo `go.mod`:

```Bash
go mod init [Domain name/ Module name]
```

_Ví dụ:_ `go mod init example.com/myproject`