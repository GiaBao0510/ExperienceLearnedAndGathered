# SQLC - Hướng dẫn học cơ bản

## Mục lục

1. [Giới thiệu về GORM](#1-giới-thiệu-về-gorm)
2. [Giới thiệu về SQLC](#2-giới-thiệu-về-sqlc)
3. [So sánh SQLC và GORM](#3-so-sánh-sqlc-và-gorm)
4. [Cài đặt SQLC trong dự án Golang](#4-cài-đặt-sqlc-trong-dự-án-golang)
5. [Cấu trúc project này](#5-cấu-trúc-project-này)
6. [Giải thích file cấu hình sqlc.yaml](#6-giải-thích-file-cấu-hình-sqlcyaml)
7. [Giải thích các file SQL đầu vào](#7-giải-thích-các-file-sql-đầu-vào)
8. [Giải thích các file Go được sinh ra](#8-giải-thích-các-file-go-được-sinh-ra)
9. [Vị trí hợp lý trong Layered Architecture](#9-vị-trí-hợp-lý-trong-layered-architecture)
10. [Quy trình cập nhật khi schema hoặc query thay đổi](#10-quy-trình-cập-nhật-khi-schema-hoặc-query-thay-đổi)
11. [Lưu ý và best practice](#11-lưu-ý-và-best-practice)
12. [Câu hỏi phỏng vấn thường gặp về SQLC](#12-câu-hỏi-phỏng-vấn-thường-gặp-về-sqlc)

---

## 1. Giới thiệu về GORM

GORM là một thư viện ORM (Object-Relational Mapping) phổ biến cho Golang. Nó cho phép lập trình viên tương tác với cơ sở dữ liệu thông qua các struct và phương thức của Go, thay vì viết câu lệnh SQL thuần.

**Ưu điểm của GORM:**

- Tốc độ phát triển nhanh: không cần viết SQL thủ công, chỉ cần định nghĩa struct và gọi phương thức.
- Hỗ trợ nhiều tính năng sẵn có: auto-migration, hooks, associations (has-one, has-many, many-to-many), soft-delete.
- Tài liệu phong phú, cộng đồng lớn, dễ tìm ví dụ.
- Phù hợp cho prototype, dự án nhỏ hoặc khi team chưa thành thạo SQL.

**Nhược điểm của GORM:**

- Hiệu suất thấp hơn SQL thuần do có lớp trừu tượng trung gian.
- Câu lệnh SQL được sinh ngầm, khó kiểm soát và debug khi có vấn đề về hiệu năng.
- Với các truy vấn phức tạp (join nhiều bảng, subquery, window function...), GORM trở nên cồng kềnh và khó đọc hơn viết SQL trực tiếp.
- Dễ dẫn đến N+1 query nếu lập trình viên không hiểu rõ cách GORM tải dữ liệu (lazy loading vs eager loading).
- Nhiều công ty lớn có chính sách hạn chế hoặc cấm dùng ORM vì lý do kiểm soát hiệu năng và độ minh bạch của câu lệnh SQL.

---

## 2. Giới thiệu về SQLC

SQLC là một công cụ sinh code (code generator), không phải một ORM. Nguyên lý hoạt động của nó hoàn toàn khác:

- Lập trình viên viết SQL thuần (schema và query) theo cú pháp chuẩn.
- SQLC đọc các file SQL đó, phân tích cú pháp, kiểm tra tính hợp lệ, rồi tự động sinh ra code Go tương ứng (type-safe, compile-time checked).
- Kết quả là các file Go chứa struct và method sẵn sàng dùng, không có runtime magic nào xảy ra.

**Ưu điểm của SQLC:**

- Type-safe: mọi tham số và kết quả trả về đều được kiểm tra kiểu tại thời điểm biên dịch, không xảy ra lỗi kiểu dữ liệu lúc runtime.
- SQL minh bạch: lập trình viên kiểm soát hoàn toàn câu lệnh SQL, dễ tối ưu hiệu năng.
- Không có runtime reflection, không có magic: code sinh ra là code Go thuần túy, dễ đọc, dễ debug.
- Tích hợp tốt với transaction và connection pool chuẩn của `database/sql`.
- Phát hiện lỗi SQL sớm: SQLC kiểm tra SQL tại thời điểm sinh code, không phải lúc chạy chương trình.

**Nhược điểm của SQLC:**

- Phải biết viết SQL: phù hợp với lập trình viên đã có nền tảng SQL, không thân thiện với người mới hoàn toàn.
- Mỗi khi thay đổi schema hoặc query, cần chạy lại lệnh `sqlc generate` thủ công (không tự động).
- Không hỗ trợ dynamic query (câu lệnh WHERE thay đổi tuỳ điều kiện) một cách tự nhiên - cần xử lý thủ công hoặc dùng thư viện bổ sung (ví dụ: `squirrel`).
- Chưa hỗ trợ một số dialect SQL hoặc tính năng nâng cao của từng database engine.

**Tại sao các dự án lớn ưa chuộng SQLC hơn GORM?**

Trong môi trường production với lượng dữ liệu lớn, mọi câu lệnh SQL đều phải được tối ưu, reviewed, và có thể explain được. GORM ẩn đi câu SQL thực sự chạy, khiến việc tối ưu trở nên khó khăn. SQLC buộc lập trình viên phải chủ động viết và chịu trách nhiệm về từng câu SQL, phù hợp với quy trình code review nghiêm ngặt và yêu cầu kiểm soát hiệu năng cao.

---

## 3. So sánh SQLC và GORM

| Tiêu chí                  | GORM                         | SQLC                                       |
| ------------------------- | ---------------------------- | ------------------------------------------ |
| Loại công cụ              | ORM (runtime)                | Code generator (compile-time)              |
| Cách viết truy vấn        | Go method chaining           | SQL thuần                                  |
| Type-safety               | Một phần (dùng interface{})  | Hoàn toàn (compile-time)                   |
| Hiệu năng                 | Thấp hơn do overhead         | Ngang với `database/sql` thuần             |
| Kiểm soát SQL             | Thấp (SQL được sinh tự động) | Cao (lập trình viên viết SQL)              |
| Dynamic query             | Hỗ trợ tốt                   | Hạn chế, cần xử lý thêm                    |
| Tốc độ phát triển ban đầu | Nhanh                        | Trung bình                                 |
| Phù hợp với               | Prototype, CRUD đơn giản     | Production, hệ thống yêu cầu hiệu năng cao |
| Học SQL                   | Không bắt buộc               | Bắt buộc                                   |
| Debugging                 | Khó (SQL ẩn)                 | Dễ (SQL tường minh)                        |

---

## 4. Cài đặt SQLC trong dự án Golang

### Bước 1: Cài đặt công cụ SQLC

```powershell
# Dùng Go install (khuyến nghị)
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

Kiểm tra cài đặt thành công:

```powershell
sqlc version
```

### Bước 2: Cài đặt driver database cho Go

Project này dùng PostgreSQL, cài driver `lib/pq`:

```powershell
go get github.com/lib/pq
```

### Bước 3: Tạo file cấu hình `sqlc.yaml`

Xem giải thích chi tiết ở [mục 6](#6-giải-thích-file-cấu-hình-sqlcyaml).

### Bước 4: Viết file SQL và sinh code

```powershell
sqlc generate
```

---

## 5. Cấu trúc project này

```
SQLC/
├── cmd/                          # Entry point của ứng dụng
├── go.mod                        # Go module definition
├── go.sum                        # Checksum của dependencies
├── sqlc.yaml                     # File cấu hình cho SQLC
├── sql/                          # Chứa các file SQL đầu vào cho SQLC
│   ├── schema.sql                # Định nghĩa cấu trúc bảng (DDL)
│   └── queries.sql               # Các câu lệnh truy vấn (DML)
└── internal/
    └── database/                 # Code Go được SQLC sinh ra tự động
        ├── db.go                 # Interface DBTX và struct Queries
        ├── models.go             # Các struct tương ứng với bảng trong DB
        └── queries.sql.go        # Các method thực thi câu lệnh SQL
```

**Lưu ý quan trọng:** Các file trong `internal/database/` được sinh ra tự động bởi SQLC. Không chỉnh sửa trực tiếp các file này. Mọi thay đổi phải được thực hiện ở `sql/schema.sql` hoặc `sql/queries.sql`, sau đó chạy lại `sqlc generate`.

---

## 6. Giải thích file cấu hình `sqlc.yaml`

```yaml
version: "2"
sql:
  - engine: "postgresql"
    queries: "sql/queries.sql"
    schema: "sql/schema.sql"
    gen:
      go:
        package: "database"
        out: "internal/database"
```

| Trường    | Ý nghĩa                                                    |
| --------- | ---------------------------------------------------------- |
| `version` | Phiên bản cú pháp của file cấu hình SQLC                   |
| `engine`  | Loại database engine: `postgresql`, `mysql`, hoặc `sqlite` |
| `queries` | Đường dẫn đến file chứa câu lệnh SQL truy vấn              |
| `schema`  | Đường dẫn đến file định nghĩa cấu trúc bảng                |
| `package` | Tên package Go cho code được sinh ra                       |
| `out`     | Thư mục đầu ra của code Go được sinh ra                    |

---

## 7. Giải thích các file SQL đầu vào

### `sql/schema.sql` - Định nghĩa cấu trúc bảng

File này chứa các câu lệnh DDL (Data Definition Language) để định nghĩa cấu trúc bảng. SQLC đọc file này để hiểu kiểu dữ liệu của từng cột, từ đó ánh xạ sang kiểu dữ liệu Go tương ứng.

```sql
CREATE TABLE users (
    uuid         VARCHAR(255) PRIMARY KEY,
    user_name    VARCHAR(255) NOT NULL,
    email        VARCHAR(255) NOT NULL UNIQUE,
    phone_number VARCHAR(10),
    password     VARCHAR(255) NOT NULL
);

CREATE TABLE orders (
    order_id     VARCHAR(255) PRIMARY KEY,
    uuid         VARCHAR(255) NOT NULL,
    order_date   DATE NOT NULL,
    total_amount DECIMAL(10,2) NOT NULL,
    status       VARCHAR(50) NOT NULL,
    FOREIGN KEY (uuid) REFERENCES users(uuid)
);
```

**Điểm cần chú ý:** Cột `phone_number` không có ràng buộc `NOT NULL`, nên SQLC ánh xạ nó sang `sql.NullString` trong Go thay vì `string` thông thường. Đây là cách SQLC xử lý giá trị nullable một cách type-safe.

### `sql/queries.sql` - Câu lệnh truy vấn

File này chứa các câu lệnh DML (Data Manipulation Language). Mỗi câu lệnh phải có một comment đặc biệt phía trên để SQLC nhận biết:

```sql
-- name: TenHam :KieuTraVe
```

Các kiểu trả về:

| Annotation  | Ý nghĩa                                 | Kiểu trả về Go   |
| ----------- | --------------------------------------- | ---------------- |
| `:one`      | Trả về đúng 1 bản ghi                   | `(T, error)`     |
| `:many`     | Trả về nhiều bản ghi                    | `([]T, error)`   |
| `:exec`     | Thực thi không cần trả về dữ liệu       | `error`          |
| `:execrows` | Thực thi và trả về số hàng bị ảnh hưởng | `(int64, error)` |

**Ví dụ trong project:**

```sql
-- name: CreateUsers :exec
INSERT INTO users (uuid, user_name, email, phone_number, password) VALUES ($1, $2, $3, $4, $5);

-- name: GetUsersByID :one
SELECT * FROM users WHERE uuid = $1;

-- name: GetAllUserss :many
SELECT * FROM users;
```

**Lưu ý về placeholder:** Project này dùng PostgreSQL nên sử dụng `$1, $2, $3...` làm placeholder. Nếu dùng MySQL thì dùng dấu `?`.

---

## 8. Giải thích các file Go được sinh ra

Sau khi chạy `sqlc generate`, SQLC tạo ra 3 file trong thư mục `internal/database/`.

### `db.go` - Tầng kết nối database

```go
// Code generated by sqlc. DO NOT EDIT.

package database

import (
    "context"
    "database/sql"
)

// DBTX là interface trừu tượng hóa kết nối database.
// Cả *sql.DB và *sql.Tx đều implement interface này,
// cho phép dùng cùng một code cho cả truy vấn thường và transaction.
type DBTX interface {
    ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
    PrepareContext(context.Context, string) (*sql.Stmt, error)
    QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
    QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
    return &Queries{db: db}
}

type Queries struct {
    db DBTX
}

// WithTx cho phép tạo một instance Queries mới trong phạm vi một transaction.
func (q *Queries) WithTx(tx *sql.Tx) *Queries {
    return &Queries{db: tx}
}
```

**Tại sao dùng interface DBTX thay vì `*sql.DB` trực tiếp?**

Đây là một thiết kế thông minh. Interface `DBTX` được implement bởi cả `*sql.DB` (kết nối thường) và `*sql.Tx` (transaction). Nhờ vậy, cùng một `Queries`, bạn có thể gọi trong cả hai trường hợp mà không cần viết code riêng. Khi cần transaction, chỉ cần gọi `q.WithTx(tx)`.

### `models.go` - Struct ánh xạ với bảng database

```go
// Code generated by sqlc. DO NOT EDIT.

package database

type Order struct {
    OrderID     string
    Uuid        string
    OrderDate   time.Time
    TotalAmount string
    Status      string
}

type User struct {
    Uuid        string
    UserName    string
    Email       string
    PhoneNumber sql.NullString  // nullable vì không có NOT NULL trong schema
    Password    string
}
```

SQLC tự động ánh xạ kiểu SQL sang kiểu Go:

| Kiểu SQL | Kiểu Go (NOT NULL) | Kiểu Go (nullable) |
|---|---|---|
| VARCHAR, TEXT | `string` | `sql.NullString` |
| INTEGER, BIGINT | `int32`, `int64` | `sql.NullInt32`, `sql.NullInt64` |
| BOOLEAN | `bool` | `sql.NullBool` |
| DATE, TIMESTAMP | `time.Time` | `sql.NullTime` |
| DECIMAL, NUMERIC | `string` | `sql.NullString` |

### `queries.sql.go` - Các method thực thi SQL

File này chứa toàn bộ logic truy vấn. Mỗi câu SQL trong `queries.sql` được sinh thành một method trên struct `Queries`.

**Ví dụ: Query trả về nhiều bản ghi (`:many`)**

```go
func (q *Queries) GetAllOrders(ctx context.Context) ([]Order, error) {
    rows, err := q.db.QueryContext(ctx, getAllOrders)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    var items []Order
    for rows.Next() {
        var i Order
        if err := rows.Scan(
            &i.OrderID,
            &i.Uuid,
            &i.OrderDate,
            &i.TotalAmount,
            &i.Status,
        ); err != nil {
            return nil, err
        }
        items = append(items, i)
    }
    // Kiểm tra lỗi sau khi duyệt hết rows
    if err := rows.Err(); err != nil {
        return nil, err
    }
    return items, nil
}
```

**Ví dụ: Query có tham số đầu vào (`:one`)**

```go
func (q *Queries) GetUsersByID(ctx context.Context, uuid string) (User, error) {
    row := q.db.QueryRowContext(ctx, getUsersByID, uuid)
    var i User
    err := row.Scan(
        &i.Uuid,
        &i.UserName,
        &i.Email,
        &i.PhoneNumber,
        &i.Password,
    )
    return i, err
}
```

**Ví dụ: Query với nhiều tham số (`:exec`)**

```go
// SQLC tạo một struct Params riêng khi query có nhiều tham số đầu vào
type CreateUsersParams struct {
    Uuid        string
    UserName    string
    Email       string
    PhoneNumber sql.NullString
    Password    string
}

func (q *Queries) CreateUsers(ctx context.Context, arg CreateUsersParams) error {
    _, err := q.db.ExecContext(ctx, createUsers,
        arg.Uuid,
        arg.UserName,
        arg.Email,
        arg.PhoneNumber,
        arg.Password,
    )
    return err
}
```

**Tại sao SQLC tạo struct `CreateUsersParams` thay vì truyền từng tham số rời?**

Khi câu lệnh SQL có nhiều tham số (từ 2 trở lên), SQLC nhóm chúng vào một struct để gọi hàm rõ ràng hơn, tránh nhầm thứ tự tham số. Đây là một thiết kế tốt vì nếu truyền nhiều tham số rời `(uuid, userName, email, phoneNumber, password)`, rất dễ nhầm thứ tự khi gọi hàm.

---

## 9. Vị trí hợp lý trong Layered Architecture

### Câu hỏi: Đặt `schema.sql` và `queries.sql` ở đâu?

**Trả lời:** Đặt trong thư mục `sql/` ở root của project (như project này đang làm).

```
├── sql/
│   ├── schema.sql
│   └── queries.sql
```

**Lý do:**

- Hai file này là "nguồn sự thật" (source of truth) cho lớp database, không thuộc về bất kỳ module nghiệp vụ nào cụ thể.
- Chúng cần được version control và review độc lập với code Go.
- Đặt ở root giúp các công cụ khác (migration tool, DBA) dễ dàng tìm thấy và quản lý.
- Một số team đặt chung với migration scripts (ví dụ: `db/migrations/` và `db/queries/`), cũng là cách hợp lý.

Trong dự án Layered Architecture ở câu hỏi, cấu trúc đề xuất:

```
├── db/                  # Hoặc đặt tên là sql/
│   ├── migrations/      # Các file migration (nếu dùng công cụ migration riêng)
│   ├── schema.sql       # Schema cho SQLC
│   └── queries.sql      # Query cho SQLC
```

### Câu hỏi: Đặt `db.go`, `models.go`, `queries.sql.go` ở đâu?

**Trả lời:** Đặt trong `internal/database/` (như project này đang làm) hoặc `internal/repo/database/`.

```
├── internal/
│   ├── modules/
│   │   ├── auth/
│   │   └── user/
│   └── database/        # Code được SQLC sinh ra
│       ├── db.go
│       ├── models.go
│       └── queries.sql.go
```

**Lý do:**

- Thư mục `internal/` đảm bảo code này chỉ được dùng trong nội bộ module, không expose ra ngoài.
- `internal/database/` thuộc về tầng **Infrastructure/Repository** trong kiến trúc phân lớp, đây là nơi xử lý tương tác trực tiếp với database.
- Các module nghiệp vụ như `auth`, `user` sẽ gọi vào `internal/database/` thông qua tầng service hoặc repository, không gọi trực tiếp từ controller.

**Sơ đồ luồng gọi đúng trong Layered Architecture:**

```
Controller (Presentation Layer)
    |
    v
Service / ServiceImpl (Business Layer)
    |
    v
internal/database (Infrastructure Layer - SQLC generated code)
    |
    v
Database (PostgreSQL)
```

---

## 10. Quy trình cập nhật khi schema hoặc query thay đổi

**Trả lời câu hỏi: SQLC có tự động cập nhật không?**

Không. SQLC không tự động cập nhật. Mỗi khi bạn thay đổi `schema.sql` hoặc `queries.sql`, bạn phải chạy lại lệnh sau để tái sinh code Go:

```powershell
sqlc generate
```

Đây là quy trình thủ công có chủ đích. Lý do là để lập trình viên kiểm soát được thời điểm code database layer thay đổi, tránh những cập nhật ngoài ý muốn.

**Quy trình cập nhật được khuyến nghị:**

```
1. Sửa schema.sql và/hoặc queries.sql
         |
         v
2. Chạy: sqlc generate
         |
         v
3. Kiểm tra code sinh ra trong internal/database/
         |
         v
4. Cập nhật code ở tầng service/repository gọi vào SQLC nếu cần
         |
         v
5. Chạy: go build ./...  (để đảm bảo không có lỗi compile)
         |
         v
6. Commit tất cả các thay đổi (cả file SQL lẫn file Go sinh ra) lên Git
```

**Lưu ý quan trọng:** Nên commit cả file `.sql` lẫn file Go được sinh ra lên Git. Điều này giúp code review rõ ràng (reviewer thấy cả SQL thay đổi và code Go tương ứng), và đảm bảo build reproducible mà không cần cài SQLC trên CI/CD.

**Tích hợp vào Makefile (thực hành tốt):**

```makefile
.PHONY: sqlc
sqlc:
	sqlc generate

.PHONY: build
build: sqlc
	go build ./...
```

---

## 11. Lưu ý và best practice

### Đặt tên query đúng convention

Đặt tên query theo quy tắc rõ ràng, nhất quán:

```sql
-- Tốt: Rõ ràng, mô tả đúng hành động
-- name: GetUserByID :one
-- name: ListActiveUsers :many
-- name: CreateUser :exec
-- name: UpdateUserEmail :exec
-- name: DeleteUserByID :exec

-- Tránh: Tên mơ hồ hoặc không nhất quán
-- name: GetAllUserss :many   <-- lỗi chính tả (2 chữ s)
-- name: CreateUsers :exec    <-- sai số (tạo 1 user mà đặt tên số nhiều)
```

### Chọn `SELECT` có chủ đích

Tránh dùng `SELECT *` trong code production. SQLC hỗ trợ tốt khi bạn liệt kê cụ thể tên cột:

```sql
-- Không nên (trong production)
SELECT * FROM users WHERE uuid = $1;

-- Nên: Chỉ lấy những cột cần thiết
SELECT uuid, user_name, email FROM users WHERE uuid = $1;
```

Khi liệt kê tên cột cụ thể, SQLC còn có thể sinh ra struct nhẹ hơn thay vì dùng full model.

### Xử lý transaction đúng cách

```go
// Khởi tạo transaction
tx, err := db.BeginTx(ctx, nil)
if err != nil {
    return err
}
defer tx.Rollback() // Luôn Rollback nếu có lỗi

qtx := queries.WithTx(tx)

// Thực hiện các thao tác trong transaction
if err := qtx.CreateUsers(ctx, createUserParams); err != nil {
    return err // tx.Rollback() sẽ được gọi bởi defer
}

if err := qtx.CreateOrder(ctx, createOrderParams); err != nil {
    return err
}

// Commit khi tất cả thành công
return tx.Commit()
```

### Không chỉnh sửa trực tiếp file được sinh ra

Các file trong `internal/database/` đều có comment đầu file:

```go
// Code generated by sqlc. DO NOT EDIT.
```

Bất kỳ thay đổi nào bạn làm trực tiếp trên các file này sẽ bị ghi đè khi chạy `sqlc generate` lần sau.

---

## 12. Câu hỏi phỏng vấn thường gặp về SQLC

### Câu 1: SQLC hoạt động theo nguyên lý nào? Nó khác ORM ở điểm gì?

**Gợi ý trả lời:**

SQLC là một code generator, không phải runtime library. Nó đọc file SQL định nghĩa schema và query, phân tích cú pháp, rồi sinh ra code Go type-safe tại thời điểm phát triển (development time). Khi chương trình chạy, không có thư viện SQLC nào hoạt động - chỉ có code Go thuần và driver database.

ORM như GORM hoạt động ngược lại: nó tồn tại ở runtime, nhận các method call Go rồi dịch thành SQL ngầm bên trong. Điều này có nghĩa là với SQLC, lỗi SQL được phát hiện sớm (khi chạy `sqlc generate`), còn với GORM, lỗi logic SQL có thể chỉ lộ ra khi chương trình chạy thực tế.

### Câu 2: Tại sao SQLC sinh ra interface `DBTX` thay vì nhận trực tiếp `*sql.DB`?

**Gợi ý trả lời:**

Interface `DBTX` được implement bởi cả `*sql.DB` và `*sql.Tx`. Điều này cho phép cùng một struct `Queries` hoạt động trong cả hai ngữ cảnh: truy vấn thông thường và transaction. Khi cần chạy nhiều thao tác trong một transaction, chỉ cần gọi `queries.WithTx(tx)` để lấy một `Queries` mới gắn với transaction đó, mà không cần viết lại bất kỳ logic SQL nào.

### Câu 3: Làm thế nào để xử lý dynamic query (WHERE điều kiện thay đổi) với SQLC?

**Gợi ý trả lời:**

Đây là điểm hạn chế của SQLC. Vì SQLC yêu cầu SQL tĩnh, có một vài hướng xử lý:

- Viết nhiều query cụ thể cho từng trường hợp (ví dụ: `ListUsersByStatus`, `ListUsersByEmail`).
- Dùng thư viện query builder như `squirrel` hoặc `goqu` để xây dựng câu SQL động, sau đó thực thi bằng `db.QueryContext` thủ công, kết hợp với model từ SQLC.
- Dùng tính năng `sqlc.narg` (nullable argument) để xử lý điều kiện tùy chọn trong một số trường hợp đơn giản.

### Câu 4: Khi thay đổi `schema.sql`, cần làm gì và lưu ý điều gì?

**Gợi ý trả lời:**

Cần chạy lại `sqlc generate` để tái sinh code Go. Tuy nhiên, thay đổi schema thường đi kèm với migration database (thêm/xóa/sửa cột hoặc bảng). SQLC chỉ sinh code, không thực hiện migration. Cần dùng công cụ migration riêng (ví dụ: `golang-migrate`, `goose`, `atlas`) để áp dụng thay đổi schema lên database thực. Sau đó chạy `sqlc generate` và kiểm tra lại toàn bộ code gọi vào tầng database vì các struct và method có thể đã thay đổi tên trường hoặc kiểu dữ liệu.

### Câu 5: Có nên commit file Go được sinh ra bởi SQLC vào Git không?

**Gợi ý trả lời:**

Nên commit. Lý do:

- Đảm bảo build reproducible: bất kỳ ai checkout code đều có thể build ngay mà không cần cài SQLC.
- Code review rõ ràng: reviewer thấy được cả SQL thay đổi lẫn code Go tương ứng trong cùng một pull request.
- CI/CD đơn giản hơn: không cần bước `sqlc generate` trong pipeline build.

Tuy nhiên, cần đảm bảo file SQL và file Go luôn đồng bộ. Có thể thêm bước kiểm tra trong CI: chạy `sqlc generate` rồi `git diff --exit-code` để phát hiện nếu ai đó quên tái sinh code sau khi sửa SQL.