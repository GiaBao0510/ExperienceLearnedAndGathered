# Hướng Dẫn Xuất Dữ Liệu Từ PostgreSQL Sang Excel Với Go

## 1. Giới Thiệu

### 1.1. Mục Tiêu Chương Trình

Chương trình này thực hiện một tác vụ phổ biến trong thực tế: **xuất toàn bộ dữ liệu từ một bảng trong PostgreSQL ra file Excel (.xlsx)**, với khả năng xử lý số lượng lớn bản ghi (khoảng 1 triệu dòng) mà không gây tràn bộ nhớ.

Thách thức chính khi xuất dữ liệu lớn:

- Nếu tải toàn bộ 1 triệu dòng vào RAM rồi mới ghi ra file, chương trình sẽ tiêu tốn hàng GB bộ nhớ và dễ bị crash.
- Cần một cơ chế đọc từng dòng từ database và ghi ngay vào file theo kiểu "streaming" — xử lý từng phần thay vì tất cả cùng lúc.

Chương trình giải quyết bài toán này bằng cách kết hợp:

- `sql.Rows` để đọc dữ liệu từ PostgreSQL theo từng dòng (streaming từ database).
- `StreamWriter` của thư viện `excelize` để ghi từng dòng vào Excel mà không giữ toàn bộ nội dung trong RAM.

---

### 1.2. Các Thư Viện Sử Dụng

#### `database/sql` (thư viện chuẩn của Go)

- Nguồn gốc: có sẵn trong Go standard library, không cần cài thêm.
- Chức năng: cung cấp interface chung để làm việc với các loại cơ sở dữ liệu quan hệ (PostgreSQL, MySQL, SQLite...). Bản thân `database/sql` chỉ định nghĩa interface — cần thêm driver cụ thể để kết nối từng loại database.
- Lý do chọn: là chuẩn của Go, ổn định, được hỗ trợ lâu dài.

#### `github.com/lib/pq` (PostgreSQL driver)

- Nguồn gốc: thư viện bên thứ ba, tải tại `github.com/lib/pq`.
- Chức năng: driver cho phép `database/sql` giao tiếp với PostgreSQL. Thư viện này được import với dấu `_` (blank import) vì chỉ cần nó tự đăng ký driver với `database/sql`, không cần gọi trực tiếp hàm nào từ nó.
- Lý do chọn: driver PostgreSQL phổ biến nhất và ổn định nhất cho Go.

#### `github.com/xuri/excelize/v2` (thao tác file Excel)

- Nguồn gốc: thư viện bên thứ ba, tải tại `github.com/xuri/excelize/v2`.
- Chức năng: tạo, đọc, và chỉnh sửa file Excel (.xlsx). Điểm nổi bật là hỗ trợ **StreamWriter** — cho phép ghi dữ liệu vào Excel theo từng dòng mà không cần giữ toàn bộ nội dung trong bộ nhớ.
- Lý do chọn: hỗ trợ streaming, phù hợp với bài toán xuất dữ liệu lớn. Nếu dùng thư viện khác không có streaming, chương trình sẽ tốn rất nhiều RAM.

#### `time` và `log` (thư viện chuẩn của Go)

- `time`: đo thời gian thực thi chương trình.
- `log`: in thông báo lỗi kèm timestamp, dừng chương trình với `log.Fatal` khi gặp lỗi nghiêm trọng.

---

## 2. Cài Đặt Môi Trường Và Thư Viện

### 2.1. Yêu Cầu

- Go phiên bản 1.18 trở lên.
- PostgreSQL đang chạy và có bảng `users` với cấu trúc: `id (int)`, `name (varchar)`, `phone (varchar)`, `address (varchar)`.

### 2.2. Khởi Tạo Module

```bash
mkdir export-excel && cd export-excel
go mod init export-excel
```

### 2.3. Cài Đặt Thư Viện

```bash
go get github.com/lib/pq
go get github.com/xuri/excelize/v2
```

Sau khi chạy hai lệnh trên, Go sẽ tự động cập nhật file `go.mod` và tải thư viện về thư mục cache.

### 2.4. Kiểm Tra `go.mod`

File `go.mod` sau khi cài đặt sẽ có dạng:

```
module export-excel

go 1.21

require (
    github.com/lib/pq v1.10.9
    github.com/xuri/excelize/v2 v2.8.0
)
```

### 2.5. Chuẩn Bị Dữ Liệu Mẫu Trong PostgreSQL

```sql
CREATE TABLE users (
    id      SERIAL PRIMARY KEY,
    name    VARCHAR(100),
    phone   VARCHAR(20),
    address VARCHAR(255)
);

-- Tạo 1 triệu bản ghi mẫu
INSERT INTO users (name, phone, address)
SELECT
    'User_' || i,
    '09' || LPAD(i::text, 8, '0'),
    'Address_' || i
FROM generate_series(1, 1000000) AS i;
```

---

## 3. Giải Thích Chi Tiết Code

### 3.1. Toàn Bộ Code

```go
package main

import (
    "database/sql"
    "log"
    "time"

    "github.com/xuri/excelize/v2"
    _ "github.com/lib/pq"
)

func main() {

    // Bắt đầu đo thời gian thực thi
    startTime := time.Now()

    db, err := sql.Open("postgres", "host=localhost port=5432 user=admin password=admin123 dbname=test sslmode=disable")
    if err != nil {
        log.Fatal("Lỗi cấu hình kết nối: ", err)
    }
    defer db.Close()

    rows, err := db.Query("SELECT * FROM users;")
    if err != nil {
        log.Fatal("Lỗi truy vấn dữ liệu: ", err)
    }
    defer rows.Close()

    elapsed := time.Since(startTime)
    log.Printf("Thời gian thực thi: %s", elapsed)

    f := excelize.NewFile()

    streamWriter, err := f.NewStreamWriter("Sheet1")
    if err != nil {
        log.Fatal("Lỗi tạo StreamWriter: ", err)
    }

    rowIndex := 1
    for rows.Next() {
        var id int
        var name string
        var phone string
        var address string

        err := rows.Scan(&id, &name, &phone, &address)
        if err != nil {
            log.Fatal("Lỗi scan dữ liệu: ", err)
        }

        cell, _ := excelize.CoordinatesToCellName(1, rowIndex)
        err = streamWriter.SetRow(cell, []interface{}{id, name, phone, address})
        if err != nil {
            log.Fatal("Lỗi ghi dữ liệu vào Excel: ", err)
        }

        rowIndex++
    }

    if err := streamWriter.Flush(); err != nil {
        log.Fatal("Lỗi khi flush dữ liệu vào Excel: ", err)
    }

    if err := f.SaveAs("output2.xlsx"); err != nil {
        log.Fatal("Lỗi khi lưu file Excel: ", err)
    }

    endElapsed := time.Since(startTime)
    log.Printf("Thời gian thực thi: %s", endElapsed)
}
```

---

### 3.2. Bước 1 — Kết Nối PostgreSQL

```go
db, err := sql.Open("postgres", "host=localhost port=5432 user=admin password=admin123 dbname=test sslmode=disable")
if err != nil {
    log.Fatal("Lỗi cấu hình kết nối: ", err)
}
defer db.Close()
```

**Giải thích từng phần:**

`sql.Open("postgres", connectionString)`:

- Tham số đầu tiên `"postgres"` là tên driver đã được thư viện `lib/pq` đăng ký.
- Tham số thứ hai là **connection string** — chuỗi chứa thông tin kết nối.
- Lưu ý quan trọng: `sql.Open` **không thực sự kết nối ngay** đến database. Nó chỉ tạo đối tượng `*sql.DB` và kiểm tra cú pháp connection string. Kết nối thực sự chỉ xảy ra khi có truy vấn đầu tiên.

Cấu trúc connection string:

```
host=localhost        → địa chỉ máy chủ PostgreSQL
port=5432             → cổng mặc định của PostgreSQL
user=admin            → tên người dùng
password=admin123     → mật khẩu
dbname=test           → tên database
sslmode=disable       → tắt SSL (dùng trong môi trường dev/local)
```

`defer db.Close()`:

- Đăng ký việc đóng kết nối để thực thi khi `main()` kết thúc.
- Đảm bảo kết nối luôn được giải phóng dù hàm kết thúc bình thường hay gặp lỗi ở bước sau.

---

### 3.3. Bước 2 — Truy Vấn Dữ Liệu (Streaming Từ Database)

```go
rows, err := db.Query("SELECT * FROM users;")
if err != nil {
    log.Fatal("Lỗi truy vấn dữ liệu: ", err)
}
defer rows.Close()
```

**Điểm quan trọng nhất:** `db.Query` **không tải toàn bộ 1 triệu dòng vào RAM ngay lập tức**.

Thay vào đó, PostgreSQL mở một **cursor** phía server và `rows` là một con trỏ trỏ đến cursor đó. Dữ liệu chỉ được truyền về client (chương trình Go) từng phần nhỏ khi bạn gọi `rows.Next()` — đây là cơ chế streaming từ database, giúp tiết kiệm bộ nhớ đáng kể.

`defer rows.Close()`: đảm bảo cursor được đóng và tài nguyên phía server được giải phóng sau khi hoàn thành.

---

### 3.4. Bước 3 — Tạo File Excel Với StreamWriter

```go
f := excelize.NewFile()

streamWriter, err := f.NewStreamWriter("Sheet1")
if err != nil {
    log.Fatal("Lỗi tạo StreamWriter: ", err)
}
```

`excelize.NewFile()`: tạo một workbook Excel mới trong bộ nhớ.

`f.NewStreamWriter("Sheet1")`: tạo một `StreamWriter` gắn với sheet có tên "Sheet1". StreamWriter hoạt động theo cơ chế **append-only** — chỉ ghi tiến về phía trước, không thể quay lại sửa dòng đã ghi. Đổi lại, nó không cần giữ toàn bộ nội dung sheet trong RAM.

So sánh hai cách ghi Excel:

| Cách ghi                          | RAM sử dụng                       | Có thể sửa dòng cũ | Phù hợp khi                                |
| --------------------------------- | --------------------------------- | ------------------ | ------------------------------------------ |
| API thông thường (`SetCellValue`) | Cao — giữ toàn bộ sheet trong RAM | Có                 | Dữ liệu nhỏ (< 10.000 dòng)                |
| StreamWriter                      | Thấp — ghi thẳng ra file          | Không              | Dữ liệu lớn (hàng trăm nghìn dòng trở lên) |

---

### 3.5. Bước 4 — Vòng Lặp Đọc Và Ghi Dữ Liệu

Đây là phần trung tâm của chương trình. Mỗi vòng lặp xử lý **đúng một dòng**: đọc từ database rồi ghi ngay vào Excel.

```go
rowIndex := 1
for rows.Next() {
    // Khai báo biến để nhận dữ liệu từng cột
    var id int
    var name string
    var phone string
    var address string

    // Scan: ánh xạ dữ liệu cột trong hàng hiện tại vào các biến Go
    err := rows.Scan(&id, &name, &phone, &address)
    if err != nil {
        log.Fatal("Lỗi scan dữ liệu: ", err)
    }

    // Chuyển (cột=1, hàng=rowIndex) thành tên ô Excel, ví dụ: (1,1) → "A1"
    cell, _ := excelize.CoordinatesToCellName(1, rowIndex)

    // Ghi một hàng dữ liệu vào Excel bắt đầu từ ô 'cell'
    err = streamWriter.SetRow(cell, []interface{}{id, name, phone, address})
    if err != nil {
        log.Fatal("Lỗi ghi dữ liệu vào Excel: ", err)
    }

    rowIndex++ // chuyển sang hàng tiếp theo
}
```

**Giải thích chi tiết từng phần quan trọng:**

`rows.Next()`:

- Trả về `true` nếu còn dòng tiếp theo, đồng thời di chuyển con trỏ đến dòng đó.
- Trả về `false` khi đã đọc hết tất cả các dòng.
- Mỗi lần gọi `rows.Next()`, một dòng dữ liệu được truyền từ PostgreSQL về chương trình.

`rows.Scan(&id, &name, &phone, &address)`:

- Đọc giá trị từng cột của dòng hiện tại và gán vào các biến Go tương ứng.
- Truyền vào **con trỏ** (dấu `&`) để `Scan` có thể ghi giá trị vào đúng vùng nhớ.
- Số lượng và thứ tự tham số phải khớp chính xác với số cột và thứ tự cột trong câu `SELECT`.

`excelize.CoordinatesToCellName(1, rowIndex)`:

- Chuyển tọa độ dạng số `(cột, hàng)` thành tên ô Excel dạng chữ-số.
- Ví dụ: `(1, 1)` → `"A1"`, `(1, 2)` → `"A2"`, `(1, 1000)` → `"A1000"`.
- Tham số đầu tiên `1` là cột bắt đầu — nghĩa là dữ liệu của mỗi hàng bắt đầu từ cột A.

`streamWriter.SetRow(cell, []interface{}{id, name, phone, address})`:

- Ghi một hàng dữ liệu vào Excel bắt đầu từ ô `cell`.
- `[]interface{}` là slice chứa giá trị của các ô trong hàng đó, theo thứ tự từ trái sang phải.
- Với `cell = "A2"` và slice 4 phần tử, dữ liệu sẽ được ghi vào A2, B2, C2, D2.

---

### 3.6. Bước 5 — Flush Và Lưu File

```go
// Flush: ghi toàn bộ dữ liệu còn trong buffer của StreamWriter ra file
if err := streamWriter.Flush(); err != nil {
    log.Fatal("Lỗi khi flush dữ liệu vào Excel: ", err)
}

// SaveAs: lưu workbook ra đĩa với tên file chỉ định
if err := f.SaveAs("output2.xlsx"); err != nil {
    log.Fatal("Lỗi khi lưu file Excel: ", err)
}
```

`streamWriter.Flush()`:

- Bắt buộc phải gọi sau khi ghi xong tất cả dữ liệu.
- StreamWriter có thể giữ một phần dữ liệu trong buffer nội bộ để tối ưu tốc độ ghi. `Flush()` đảm bảo toàn bộ buffer được xả ra và sheet được đánh dấu hoàn chỉnh.
- Nếu bỏ qua bước này, file Excel có thể bị thiếu dữ liệu ở cuối hoặc bị hỏng.

`f.SaveAs("output2.xlsx")`:

- Ghi workbook ra đĩa thành file `output2.xlsx` trong thư mục hiện tại.
- Nếu muốn lưu vào đường dẫn cụ thể: `f.SaveAs("/data/exports/output2.xlsx")`.

---

### 3.7. Xử Lý Lỗi Và Đóng Tài Nguyên

Chương trình sử dụng hai cơ chế xử lý lỗi và đóng tài nguyên:

**`defer` để đóng tài nguyên tự động:**

```go
defer db.Close()    // đóng connection pool đến PostgreSQL
defer rows.Close()  // đóng cursor truy vấn phía server
```

`defer` đảm bảo hai lời gọi này luôn được thực thi khi `main()` kết thúc — dù kết thúc bình thường hay do `log.Fatal`. Đây là pattern chuẩn trong Go để tránh resource leak.

**`log.Fatal` để dừng ngay khi gặp lỗi nghiêm trọng:**

```go
if err != nil {
    log.Fatal("Mô tả lỗi: ", err)
}
```

`log.Fatal` thực hiện hai việc: in thông báo lỗi kèm timestamp ra stderr, sau đó gọi `os.Exit(1)` để dừng chương trình ngay lập tức. Các hàm `defer` đã đăng ký **không được gọi** khi dùng `log.Fatal` — đây là điểm cần lưu ý nếu cần dọn dẹp tài nguyên trước khi thoát.

---

### 3.8. Lỗi Chính Tả Trong Code Gốc

Trong code gốc có một lỗi chính tả nhỏ:

```go
// SAI — tên biến bị viết sai
streeamWriter, err := f.NewStreamWriter("Sheet1")

// ĐÚNG
streamWriter, err := f.NewStreamWriter("Sheet1")
```

Lỗi này không ảnh hưởng đến chức năng (Go chấp nhận bất kỳ tên biến hợp lệ nào) nhưng làm giảm khả năng đọc code.

---

## 4. Lưu Ý Về Hiệu Năng Khi Xuất 1 Triệu Dòng

### 4.1. Điểm Mạnh Của Thiết Kế Hiện Tại

- **Streaming hai đầu:** dữ liệu chảy từ PostgreSQL qua Go sang Excel theo từng dòng — RAM không tăng tỉ lệ thuận với số dòng dữ liệu.
- **Connection pooling:** `database/sql` tự động quản lý pool kết nối, không mở kết nối mới cho mỗi truy vấn.

### 4.2. Hạn Chế Và Cách Cải Thiện

**Hạn chế 1: Không có header cho file Excel**

Code hiện tại bắt đầu ghi dữ liệu từ hàng 1 mà không có dòng tiêu đề cột. Cách thêm header:

```go
// Ghi header trước khi vào vòng lặp
headerCell, _ := excelize.CoordinatesToCellName(1, 1)
streamWriter.SetRow(headerCell, []interface{}{"ID", "Name", "Phone", "Address"})
rowIndex = 2 // bắt đầu dữ liệu từ hàng 2
```

**Hạn chế 2: `SELECT *` không rõ ràng**

`SELECT *` phụ thuộc vào thứ tự cột trong database. Nếu cấu trúc bảng thay đổi, chương trình có thể Scan sai cột mà không báo lỗi. Nên dùng:

```go
rows, err := db.Query("SELECT id, name, phone, address FROM users;")
```

**Hạn chế 3: Không kiểm tra lỗi sau vòng lặp `rows.Next()`**

`rows.Next()` có thể dừng sớm do lỗi mạng hoặc lỗi server, không chỉ vì hết dữ liệu. Cần kiểm tra:

```go
for rows.Next() {
    // ... xử lý dữ liệu
}

// Kiểm tra lỗi xảy ra trong quá trình lặp
if err := rows.Err(); err != nil {
    log.Fatal("Lỗi trong quá trình đọc dữ liệu: ", err)
}
```

**Hạn chế 4: Truy vấn không có giới hạn thời gian**

Với 1 triệu dòng, truy vấn có thể chạy rất lâu. Nếu kết nối bị đứt giữa chừng, chương trình sẽ treo. Nên dùng `context` để đặt timeout:

```go
import "context"

ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
defer cancel()

rows, err := db.QueryContext(ctx, "SELECT id, name, phone, address FROM users;")
```

**Hạn chế 5: Giới hạn của file Excel**

Excel có giới hạn tối đa **1.048.576 hàng** mỗi sheet. Nếu dữ liệu vượt quá giới hạn này, `StreamWriter` sẽ báo lỗi. Cần xử lý tách file hoặc tách sheet nếu dữ liệu có thể vượt giới hạn.

### 4.3. Ước Tính Hiệu Năng

Trên máy tính thông thường với kết nối local đến PostgreSQL:

|Số dòng|Thời gian ước tính|RAM tiêu thụ|
|---|---|---|
|100.000 dòng|5 - 15 giây|50 - 100 MB|
|1.000.000 dòng|60 - 180 giây|100 - 300 MB|

Thời gian thực tế phụ thuộc vào: tốc độ đĩa, tốc độ mạng đến database, kích thước dữ liệu mỗi cột, và cấu hình phần cứng.

---

## 5. Các Khái Niệm Và Pattern Quan Trọng

### 5.1. `defer` — Trì Hoãn Thực Thi

`defer` đăng ký một lời gọi hàm để thực thi **ngay trước khi hàm chứa nó kết thúc**, bất kể hàm kết thúc bình thường hay do lỗi `panic`.

```go
func readFile() {
    f, err := os.Open("data.txt")
    if err != nil {
        return // defer vẫn chạy khi hàm kết thúc ở đây
    }
    defer f.Close() // luôn được gọi dù hàm kết thúc ở đâu

    // ... xử lý file
} // f.Close() được gọi tại đây
```

Nhiều `defer` trong cùng một hàm thực thi theo thứ tự **LIFO** (Last-In, First-Out):

```go
defer fmt.Println("1") // chạy thứ ba
defer fmt.Println("2") // chạy thứ hai
defer fmt.Println("3") // chạy đầu tiên
// Output: 3, 2, 1
```

### 5.2. `context` — Kiểm Soát Thời Gian Và Hủy Tác Vụ

`context` là cơ chế trong Go để truyền **tín hiệu hủy** và **thời hạn (deadline/timeout)** qua các hàm và goroutine. Đây là pattern bắt buộc trong code backend thực tế.

```go
import (
    "context"
    "time"
)

func main() {
    // Tạo context với timeout 5 phút
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
    defer cancel() // luôn gọi cancel để giải phóng tài nguyên context

    // Truyền context vào truy vấn
    rows, err := db.QueryContext(ctx, "SELECT id, name FROM users")
    // Nếu truy vấn chưa xong sau 5 phút, ctx sẽ tự động hủy truy vấn
}
```

Các loại context thường dùng:

|Loại|Hàm tạo|Khi nào dùng|
|---|---|---|
|Timeout|`context.WithTimeout(parent, duration)`|Giới hạn thời gian tối đa|
|Deadline|`context.WithDeadline(parent, time)`|Đặt thời điểm kết thúc cụ thể|
|Cancel|`context.WithCancel(parent)`|Hủy thủ công từ bên ngoài|

### 5.3. Error Handling — Xử Lý Lỗi Trong Go

Go không dùng try-catch. Thay vào đó, hàm trả về lỗi như một giá trị thông thường, và lập trình viên phải kiểm tra tường minh:

```go
result, err := someFunction()
if err != nil {
    // xử lý lỗi
    return fmt.Errorf("mô tả ngữ cảnh: %w", err) // %w để wrap lỗi gốc
}
```

Phân biệt các cách xử lý lỗi:

|Cách|Hành vi|Dùng khi|
|---|---|---|
|`return err`|Trả lỗi lên caller|Hàm thông thường|
|`log.Fatal(err)`|In lỗi và thoát ngay|Lỗi không thể phục hồi trong `main`|
|`log.Println(err)`|In lỗi và tiếp tục|Lỗi không nghiêm trọng, cần log lại|
|`panic(err)`|Dừng goroutine hiện tại|Lỗi lập trình (không nên dùng cho lỗi runtime)|

### 5.4. Goroutine và Channel — Xử Lý Đồng Thời

Nếu muốn tăng tốc độ xuất dữ liệu, có thể tách công việc đọc và ghi thành hai goroutine riêng biệt giao tiếp qua channel. Đây là **Pipeline Pattern**:

```go
// Goroutine 1: đọc từ database, gửi dữ liệu vào channel
func readFromDB(db *sql.DB, rowCh chan<- []interface{}) {
    defer close(rowCh) // đóng channel khi đọc xong

    rows, err := db.Query("SELECT id, name, phone, address FROM users")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
        var id int
        var name, phone, address string
        rows.Scan(&id, &name, &phone, &address)
        rowCh <- []interface{}{id, name, phone, address} // gửi sang goroutine ghi
    }
}

// Goroutine 2: nhận dữ liệu từ channel, ghi vào Excel
func writeToExcel(rowCh <-chan []interface{}, filename string) {
    f := excelize.NewFile()
    sw, _ := f.NewStreamWriter("Sheet1")

    rowIndex := 1
    for row := range rowCh { // nhận cho đến khi channel đóng
        cell, _ := excelize.CoordinatesToCellName(1, rowIndex)
        sw.SetRow(cell, row)
        rowIndex++
    }

    sw.Flush()
    f.SaveAs(filename)
}

func main() {
    rowCh := make(chan []interface{}, 1000) // buffer 1000 dòng

    go readFromDB(db, rowCh)   // chạy song song
    writeToExcel(rowCh, "output.xlsx") // main goroutine xử lý ghi
}
```

Lợi ích: trong khi goroutine đọc đang chờ dữ liệu từ PostgreSQL (I/O), goroutine ghi vẫn có thể tiếp tục ghi dữ liệu đã có trong channel — tận dụng thời gian chờ I/O.

### 5.5. Batch Processing — Xử Lý Theo Lô

Thay vì lấy toàn bộ 1 triệu dòng trong một truy vấn, có thể chia thành nhiều lô nhỏ (batch) bằng `LIMIT` và `OFFSET`. Cách này giúp kiểm soát bộ nhớ tốt hơn và dễ xử lý lỗi từng phần:

```go
const batchSize = 10000

for offset := 0; ; offset += batchSize {
    rows, err := db.QueryContext(ctx,
        "SELECT id, name, phone, address FROM users ORDER BY id LIMIT $1 OFFSET $2",
        batchSize, offset,
    )
    if err != nil {
        log.Fatal(err)
    }

    count := 0
    for rows.Next() {
        // xử lý từng dòng trong lô
        count++
    }
    rows.Close()

    if count < batchSize {
        break // đã xử lý hết dữ liệu
    }
}
```

Lưu ý: với dữ liệu lớn, `OFFSET` cao sẽ làm PostgreSQL chậm vì phải bỏ qua nhiều dòng. Dùng **keyset pagination** (`WHERE id > lastID`) sẽ hiệu quả hơn:

```go
lastID := 0
for {
    rows, err := db.QueryContext(ctx,
        "SELECT id, name, phone, address FROM users WHERE id > $1 ORDER BY id LIMIT $2",
        lastID, batchSize,
    )
    // ...
    // cập nhật lastID sau mỗi lô
}
```

---

## 6. Tổng Kết

| Thành phần                | Vai trò                   | Điểm quan trọng                                 |
| ------------------------- | ------------------------- | ----------------------------------------------- |
| `sql.Open`                | Cấu hình kết nối          | Chưa kết nối thực sự, chỉ tạo đối tượng         |
| `db.Query`                | Mở cursor streaming từ DB | Không tải toàn bộ dữ liệu vào RAM               |
| `rows.Next` + `rows.Scan` | Đọc từng dòng             | Thứ tự tham số Scan phải khớp SELECT            |
| `rows.Err()`              | Kiểm tra lỗi sau vòng lặp | Hay bị bỏ sót, cần kiểm tra                     |
| `StreamWriter`            | Ghi Excel không tốn RAM   | Append-only, phải Flush sau khi xong            |
| `defer`                   | Đóng tài nguyên tự động   | Chạy khi hàm kết thúc, không chạy sau `os.Exit` |
| `context`                 | Kiểm soát timeout         | Nên dùng `QueryContext` thay vì `Query`         |

> Để nâng cao hơn, bạn có thể tìm hiểu thêm về: `pgx` (driver PostgreSQL hiệu năng cao hơn `lib/pq`), `sync.WaitGroup` để đồng bộ nhiều goroutine, và `errgroup` để xử lý lỗi trong môi trường đa goroutine.