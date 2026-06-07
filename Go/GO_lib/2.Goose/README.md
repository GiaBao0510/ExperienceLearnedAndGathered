# Goose - Hướng dẫn học cơ bản

## Mục lục

1. [Bối cảnh và vấn đề cần giải quyết](#1-bối-cảnh-và-vấn-đề-cần-giải-quyết)
2. [Giới thiệu về Goose](#2-giới-thiệu-về-goose)
3. [So sánh Goose và SQLC](#3-so-sánh-goose-và-sqlc)
4. [Cài đặt Goose](#4-cài-đặt-goose)
5. [Cấu trúc project này](#5-cấu-trúc-project-này)
6. [Cú pháp annotation của Goose](#6-cú-pháp-annotation-của-goose)
7. [Các lệnh cơ bản của Goose](#7-các-lệnh-cơ-bản-của-goose)
8. [Hướng dẫn thực hành từng bước](#8-hướng-dẫn-thực-hành-từng-bước)
9. [Kết hợp Goose và SQLC trong một dự án](#9-kết-hợp-goose-và-sqlc-trong-một-dự-án)
10. [Tích hợp Makefile](#10-tích-hợp-makefile)
11. [Lưu ý và best practice](#11-lưu-ý-và-best-practice)
12. [Câu hỏi phỏng vấn thường gặp](#12-câu-hỏi-phỏng-vấn-thường-gặp)

---

## 1. Bối cảnh và vấn đề cần giải quyết

Ở phần trước, chúng ta đã tìm hiểu về **SQLC** - công cụ sinh code Go type-safe từ các file SQL định nghĩa sẵn. Tuy nhiên, SQLC có một hạn chế quan trọng:

> **SQLC chỉ sinh code Go, không thực hiện việc tạo hoặc thay đổi bảng trên database thực tế.**

Điều này có nghĩa là: khi schema thay đổi (thêm cột, đổi kiểu dữ liệu, thêm bảng mới...), lập trình viên phải tự tay áp dụng thay đổi đó lên database. Ở quy mô nhỏ, điều này có thể thực hiện thủ công. Nhưng trong môi trường production, với nhiều môi trường (development, staging, production) và nhiều thành viên trong team, cách tiếp cận thủ công dẫn đến các vấn đề nghiêm trọng:

- Không biết database ở môi trường này đang ở phiên bản nào.
- Một thành viên áp dụng thay đổi nhưng thành viên khác thì chưa.
- Không thể rollback khi schema có lỗi.
- Không có lịch sử thay đổi rõ ràng cho quá trình code review.

**Đây chính là bài toán mà Goose giải quyết: quản lý phiên bản schema database một cách có kiểm soát.**

---

## 2. Giới thiệu về Goose

Goose là một công cụ **database migration** viết bằng Go. Nó cho phép lập trình viên:

- Định nghĩa các thay đổi schema dưới dạng file SQL (hoặc Go) có đánh số phiên bản.
- Áp dụng (migrate lên) hoặc hoàn tác (rollback) các thay đổi theo thứ tự.
- Theo dõi trạng thái hiện tại của database qua một bảng quản lý nội bộ (`goose_db_version`).

**Nguyên lý hoạt động:**

1. Lập trình viên tạo file migration (có định dạng `<timestamp>_<tên>.sql`).
2. Goose đọc các file đó, so sánh với bảng `goose_db_version` trong database để biết những file nào đã được áp dụng.
3. Khi chạy lệnh `up`, Goose chỉ áp dụng những file migration chưa được chạy, theo đúng thứ tự timestamp.
4. Khi chạy lệnh `down`, Goose hoàn tác migration cuối cùng đã được áp dụng.

**Ưu điểm của Goose:**

- Quản lý lịch sử thay đổi schema rõ ràng, có thể review như code bình thường.
- Hỗ trợ cả file SQL thuần và file Go (cho các migration phức tạp cần logic xử lý dữ liệu).
- Hỗ trợ nhiều database: PostgreSQL, MySQL, SQLite, Microsoft SQL Server, CockroachDB.
- Có thể nhúng trực tiếp vào ứng dụng Go (embedded migration) hoặc dùng như CLI độc lập.
- Đơn giản, không phụ thuộc vào framework cụ thể nào.

**Nhược điểm của Goose:**
- Không tự động phát hiện sự khác biệt giữa schema hiện tại và schema mong muốn (không có "diff" như một số công cụ khác như Atlas).
- Lập trình viên phải tự viết cả phần `Up` và `Down` cho mỗi migration.
- Không có giao diện đồ họa quản lý migration.

---

## 3. So sánh Goose và SQLC

Đây là câu hỏi thường gặp vì cả hai đều liên quan đến database. Tuy nhiên, chúng giải quyết hai bài toán hoàn toàn khác nhau và thường được dùng **kết hợp** với nhau:

| Tiêu chí             | SQLC                                  | Goose                                      |
| -------------------- | ------------------------------------- | ------------------------------------------ |
| Mục đích             | Sinh code Go từ SQL query             | Quản lý phiên bản schema database          |
| Đầu vào              | File `schema.sql` và `queries.sql`    | Các file migration có đánh số thứ tự       |
| Đầu ra               | Code Go type-safe (`.go` files)       | Thay đổi trực tiếp trên database           |
| Thực thi khi nào     | Lúc phát triển (development time)     | Lúc triển khai (deployment time)           |
| Tương tác với DB     | Không (chỉ đọc file SQL để sinh code) | Có (kết nối và thực thi SQL trên database) |
| Rollback             | Không áp dụng                         | Có (lệnh `down`)                           |
| Thường dùng cùng với | Goose (để migration), pgx/lib pq      | SQLC (để sinh code), hoặc độc lập          |

**Tóm lại:** Goose lo việc *schema của database thực tế ở phiên bản nào*. SQLC lo việc *code Go tương tác với schema đó trông như thế nào*. Hai công việc độc lập, bổ trợ cho nhau.

---

## 4. Cài đặt Goose

### Cài đặt qua Go install (khuyến nghị)

```shell
go install github.com/pressly/goose/v3/cmd/goose@latest
```

### Kiểm tra cài đặt thành công

```shell
goose -version
```

Kết quả mong đợi:

```
goose version: v3.x.x
```

### Yêu cầu

- Go 1.21 trở lên.
- Đã cài đặt và có thể kết nối tới database (project này dùng PostgreSQL).

---

## 5. Cấu trúc project này

```
2.Goose/
├── makefile                      # Các lệnh tắt để chạy Goose
└── sql/                          # Thư mục chứa các file migration
    └── 20260607125048_order.sql  # File migration: tạo bảng orders
```

**Giải thích quy ước đặt tên file migration:**

```
20260607125048_order.sql
│              │
│              └─── Tên mô tả nội dung migration (viết thường, dùng dấu _)
└────────────────── Timestamp (YYYYMMDDHHmmss) - đây là version ID
```

Goose sử dụng phần timestamp ở đầu tên file để xác định thứ tự áp dụng migration. Timestamp càng nhỏ thì được áp dụng trước. Đây là lý do tại sao **tuyệt đối không được đổi tên file migration sau khi đã áp dụng lên database**.

---

## 6. Cú pháp annotation của Goose

Mỗi file migration SQL của Goose phải chứa hai annotation đặc biệt dưới dạng comment SQL:

```sql
-- +goose Up
-- Viết câu lệnh SQL cho migration lên (thêm bảng, thêm cột, v.v.)

-- +goose Down
-- Viết câu lệnh SQL để hoàn tác migration trên (xóa bảng, xóa cột, v.v.)
```

**Ví dụ thực tế từ project này (`20260607125048_order.sql`):**

```sql
-- +goose Up
CREATE TABLE orders (
    orderId INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    userId  INT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS orders;
```

**Lưu ý quan trọng về phần `Down`:**

- Phần `Down` phải là thao tác **nghịch đảo** chính xác của phần `Up`. Nếu `Up` tạo bảng, thì `Down` phải xóa bảng đó. Nếu `Up` thêm cột, thì `Down` phải xóa cột đó.
- Trong trường hợp migration liên quan đến dữ liệu (data migration), cần cân nhắc kỹ khả năng mất dữ liệu khi rollback.
- Không phải mọi migration đều có thể rollback hoàn toàn (ví dụ: xóa cột có dữ liệu). Trong trường hợp đó, cần ghi chú rõ ràng hoặc để phần `Down` trống với comment giải thích.

---

## 7. Các lệnh cơ bản của Goose

Goose hỗ trợ hai cú pháp gọi lệnh:

**Cú pháp cũ (dùng trong project này):**
```shell
goose <driver> "<connection_string>" <command>
```

**Cú pháp mới (khuyến nghị từ v3):**
```shell
goose -driver=<driver> -dbstring="<connection_string>" -dir=<migration_dir> <command>
```

### Bảng tổng hợp các lệnh

| Lệnh       | Mô tả                                                                 |
| ---------- | --------------------------------------------------------------------- |
| `up`       | Áp dụng tất cả migration chưa được chạy (theo thứ tự timestamp)       |
| `up-by-one`| Chỉ áp dụng một migration tiếp theo                                   |
| `up-to VERSION` | Áp dụng migration tới phiên bản cụ thể                          |
| `down`     | Hoàn tác (rollback) một migration gần nhất đã được áp dụng            |
| `down-to VERSION` | Hoàn tác về phiên bản cụ thể                                  |
| `reset`    | Hoàn tác toàn bộ migration (rollback về trạng thái ban đầu)           |
| `status`   | Hiển thị danh sách migration và trạng thái (applied / pending)        |
| `version`  | Hiển thị phiên bản migration hiện tại của database                    |
| `create <name> sql` | Tạo file migration SQL mới với timestamp tự động             |
| `create <name> go`  | Tạo file migration Go mới                                    |

---

## 8. Hướng dẫn thực hành từng bước

Phần này trình bày toàn bộ luồng làm việc thực tế với Goose, từ khi tạo migration đến khi kiểm tra kết quả.

### Bước 1: Tạo file migration mới

```shell
# Chạy lệnh từ trong thư mục chứa thư mục sql/
goose -dir=sql create order sql
```

Goose sẽ tạo ra file có tên dạng `<timestamp>_order.sql` trong thư mục `sql/`:

```
2026/06/07 19:50:48 Created new file: sql/20260607125048_order.sql
```

File được tạo ra có nội dung mẫu:

```sql
-- +goose Up
SELECT 'up SQL query';

-- +goose Down
SELECT 'down SQL query';
```

### Bước 2: Chỉnh sửa nội dung migration

Thay thế nội dung mẫu bằng câu lệnh DDL thực tế:

```sql
-- +goose Up
CREATE TABLE orders (
    orderId INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    userId  INT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS orders;
```

### Bước 3: Áp dụng migration lên database

```shell
goose postgres "host=localhost port=5432 user=admin password=admin123 dbname=test sslmode=disable" up
```

Kết quả thành công:

```
2026/06/07 20:48:46 OK   20260607125048_order.sql (29.2ms)
2026/06/07 20:48:46 goose: successfully migrated database to version: 20260607125048
```

### Bước 4: Kiểm tra kết quả trên database

**Kiểm tra bảng đã được tạo chưa:**

```sql
SELECT table_name
FROM information_schema.tables
WHERE table_schema = 'public'
  AND table_type = 'BASE TABLE';
```

**Kiểm tra lịch sử migration qua bảng quản lý của Goose:**

```sql
SELECT * FROM goose_db_version;
```

Kết quả:

| id | version_id        | is_applied | tstamp                      |
|----|-------------------|------------|-----------------------------|
| 1  | 0                 | true       | 2026-06-07 13:43:59.181957  |
| 2  | 20260607125048    | true       | 2026-06-07 13:48:46.234182  |

> Bản ghi đầu tiên (`version_id = 0`) là bản ghi khởi tạo mà Goose tự tạo khi lần đầu kết nối. Các bản ghi tiếp theo tương ứng với mỗi file migration đã được áp dụng thành công.

### Bước 5: Chạy lại lệnh `up` khi đã cập nhật

Nếu tất cả migration đã được áp dụng, Goose sẽ thông báo và không làm gì thêm:

```
2026/06/07 21:06:46 goose: no migrations to run. current version: 20260607125048
```

Đây là hành vi **idempotent** - an toàn khi chạy nhiều lần.

### Bước 6: Rollback migration

```shell
goose postgres "host=localhost port=5432 user=admin password=admin123 dbname=test sslmode=disable" down
```

```
2026/06/07 21:08:11 OK   20260607125048_order.sql (53.5ms)
```

Lệnh `down` chỉ hoàn tác **một** migration gần nhất. Bảng `orders` sẽ bị xóa vì phần `Down` trong file migration có câu lệnh `DROP TABLE`.

Để áp dụng lại, chỉ cần chạy `up`:

```shell
goose postgres "host=localhost port=5432 user=admin password=admin123 dbname=test sslmode=disable" up
```

```
2026/06/07 21:09:14 OK   20260607125048_order.sql (33.71ms)
2026/06/07 21:09:14 goose: successfully migrated database to version: 20260607125048
```

---

## 9. Kết hợp Goose và SQLC trong một dự án

Đây là pattern phổ biến nhất khi xây dựng backend bằng Go. Goose và SQLC phục vụ hai mục đích khác nhau nên có thể dùng song song mà không xung đột.

### Cấu trúc thư mục đề xuất

```
project/
├── db/
│   ├── migrations/          # Các file migration của Goose (lịch sử thay đổi schema)
│   │   ├── 20260601000001_create_users.sql
│   │   └── 20260607125048_create_orders.sql
│   └── queries/             # Các file query của SQLC
│       └── users.sql
├── internal/
│   └── database/            # Code Go được SQLC sinh ra (không chỉnh sửa trực tiếp)
│       ├── db.go
│       ├── models.go
│       └── queries.sql.go
└── sqlc.yaml
```

### Quy trình phối hợp khi thay đổi schema

```
1. Viết file migration mới với Goose
         |
         v
2. Chạy: make up  (áp dụng migration lên database development)
         |
         v
3. Cập nhật file schema.sql dùng cho SQLC (đồng bộ với migration mới nhất)
         |
         v
4. Chạy: sqlc generate  (tái sinh code Go)
         |
         v
5. Cập nhật code tầng service/repository nếu cần
         |
         v
6. Chạy: go build ./...  (kiểm tra không có lỗi compile)
         |
         v
7. Commit tất cả thay đổi (file migration, file SQL, file Go sinh ra) lên Git
```

**Câu hỏi thực tế: Dự án đã có database sẵn (Database First), làm thế nào để dùng Goose và SQLC?**

Đây là tình huống hay gặp khi tiếp nhận dự án cũ. Quy trình xử lý như sau:

1. **Tạo migration khởi điểm (baseline migration):** Xuất schema hiện tại của database thành file SQL (dùng `pg_dump --schema-only` với PostgreSQL hoặc công cụ tương đương). Đặt file này làm migration đầu tiên của Goose.
2. **Đánh dấu migration đã được áp dụng:** Vì database đã có sẵn schema, cần báo cho Goose biết migration đầu tiên này không cần chạy lại. Dùng lệnh `goose mark-applied` (hoặc `goose up-by-one` theo từng bước trong môi trường mới).
3. **Dùng SQLC với schema hiện tại:** Cung cấp file schema (từ `pg_dump`) cho SQLC để sinh code.
4. **Các thay đổi tiếp theo** đều được quản lý bình thường qua Goose và SQLC như một dự án mới.

Cả hai thư viện đều **không hỗ trợ tự động reverse-engineering** (tức là tự động tạo code hoặc migration từ database có sẵn). Công cụ phù hợp hơn cho reverse-engineering là **Atlas** (`ariga.io/atlas`), vốn có tính năng so sánh schema hiện tại với schema mong muốn và tự động tạo ra migration cần thiết.

---

## 10. Tích hợp Makefile

Dùng Makefile để rút gọn các lệnh dài, tránh lỗi gõ tay và chuẩn hóa quy trình cho cả team.

```makefile
# Chỉ định shell phù hợp trên Windows
SHELL := cmd.exe
.SHELLFLAGS := /C

GOOSE_DRIVER       ?= postgres
GOOSE_DBSTRING      = "host=localhost port=5432 user=admin password=admin123 dbname=test sslmode=disable"
GOOSE_MIGRATION_DIR ?= sql

up:
	goose -driver=$(GOOSE_DRIVER) -dbstring=$(GOOSE_DBSTRING) -dir=$(GOOSE_MIGRATION_DIR) up

down:
	goose -driver=$(GOOSE_DRIVER) -dbstring=$(GOOSE_DBSTRING) -dir=$(GOOSE_MIGRATION_DIR) down

reset:
	goose -driver=$(GOOSE_DRIVER) -dbstring=$(GOOSE_DBSTRING) -dir=$(GOOSE_MIGRATION_DIR) reset

status:
	goose -driver=$(GOOSE_DRIVER) -dbstring=$(GOOSE_DBSTRING) -dir=$(GOOSE_MIGRATION_DIR) status

.PHONY: up down reset status
```

Sử dụng:

```shell
make up      # Áp dụng tất cả migration còn thiếu
make down    # Rollback migration cuối
make reset   # Rollback toàn bộ về trạng thái ban đầu
make status  # Xem trạng thái các migration
```

**Lưu ý bảo mật:** Không nên hardcode thông tin kết nối database (mật khẩu, tên người dùng) trực tiếp trong Makefile. Trong dự án thực tế, nên dùng biến môi trường hoặc file `.env` (và thêm `.env` vào `.gitignore`).

---

## 11. Lưu ý và best practice

### Không bao giờ chỉnh sửa file migration đã được áp dụng

Sau khi một file migration đã được chạy trên bất kỳ môi trường nào (kể cả development), **tuyệt đối không sửa nội dung file đó**. Goose xác định trạng thái dựa trên `version_id` (phần timestamp trong tên file), không phải nội dung file. Nếu nội dung thay đổi nhưng version_id giữ nguyên, database sẽ ở trạng thái không khớp với code mà không có cảnh báo nào.

Nếu cần sửa một thay đổi đã migration, hãy **tạo file migration mới** để sửa chữa.

### Viết phần `Down` cẩn thận

Phần `Down` phải nghịch đảo chính xác phần `Up`. Luôn kiểm tra bằng cách chạy `up` rồi `down` rồi `up` lại trên môi trường development để xác nhận rollback hoạt động đúng.

### Commit file migration vào Git

Tương tự như code, file migration phải được commit và review qua pull request. Đây là lịch sử thay đổi của database - tài sản quan trọng của dự án.

### Đặt tên file migration mô tả rõ nội dung

```
# Tốt: rõ ràng, dễ hiểu khi đọc danh sách migration
20260601000001_create_users_table.sql
20260607125048_create_orders_table.sql
20260610080000_add_status_column_to_orders.sql

# Tránh: mơ hồ, không biết migration làm gì
20260601000001_update.sql
20260607125048_fix.sql
```

### Tách biệt schema migration và data migration

- **Schema migration:** Thay đổi cấu trúc bảng (DDL). An toàn để rollback.
- **Data migration:** Thay đổi hoặc chuyển đổi dữ liệu hiện có. Cần cẩn thận vì rollback có thể gây mất dữ liệu.

Khi cần data migration phức tạp, hãy cân nhắc dùng **file migration Go** thay vì SQL thuần, vì Go cho phép xử lý logic phức tạp hơn.

---

## 12. Câu hỏi phỏng vấn thường gặp

### Câu 1: Goose là gì và tại sao cần công cụ migration trong dự án backend?

**Gợi ý trả lời:**

Goose là công cụ quản lý phiên bản schema database. Trong dự án thực tế, schema database thay đổi liên tục theo yêu cầu nghiệp vụ. Nếu không có công cụ migration, lập trình viên phải tự tay áp dụng thay đổi trên từng môi trường - dễ sai sót, không có lịch sử, không rollback được. Goose giải quyết vấn đề này bằng cách lưu lịch sử thay đổi dưới dạng file có thứ tự, theo dõi trạng thái qua bảng `goose_db_version`, và hỗ trợ áp dụng hoặc hoàn tác từng bước.

### Câu 2: Sự khác biệt giữa Goose và SQLC là gì? Chúng có thể dùng cùng nhau không?

**Gợi ý trả lời:**

Đây là hai công cụ giải quyết hai bài toán khác nhau. Goose quản lý phiên bản schema của database thực tế (thực thi SQL lên database). SQLC sinh code Go type-safe từ file SQL query (không tương tác với database). Chúng thường được dùng **kết hợp**: Goose lo việc schema database luôn ở đúng phiên bản, SQLC lo việc code Go tương tác với schema đó một cách an toàn về kiểu dữ liệu.

### Câu 3: Điều gì xảy ra nếu chỉnh sửa nội dung một file migration đã được áp dụng?

**Gợi ý trả lời:**

Goose theo dõi migration qua `version_id` (timestamp trong tên file), không phải nội dung file. Vì vậy, nếu nội dung thay đổi nhưng tên file giữ nguyên, Goose vẫn coi migration đó đã được áp dụng và sẽ không chạy lại. Kết quả là database thực tế sẽ ở trạng thái không khớp với file migration mà không có bất kỳ cảnh báo nào. Đây là lỗi nghiêm trọng, đặc biệt khi xảy ra trên production. Nguyên tắc bất di bất dịch: **không bao giờ sửa file migration đã được commit và áp dụng**.

### Câu 4: Làm thế nào để tích hợp Goose vào quy trình CI/CD?

**Gợi ý trả lời:**

Trong pipeline CI/CD, bước migration thường được đặt trước bước khởi động ứng dụng. Cách đơn giản nhất là gọi lệnh `goose up` với connection string từ biến môi trường CI. Ngoài ra, Goose hỗ trợ **embedded migration**: nhúng Goose trực tiếp vào ứng dụng Go và tự động chạy migration khi ứng dụng khởi động. Phương pháp này phù hợp với môi trường container hóa (Docker/Kubernetes), đảm bảo schema luôn được cập nhật trước khi ứng dụng bắt đầu nhận request.

### Câu 5: Với dự án đã có database sẵn (Database First), nên dùng Goose như thế nào?

**Gợi ý trả lời:**

Goose không hỗ trợ tự động reverse-engineering từ database có sẵn. Cách tiếp cận thực tế là: xuất schema hiện tại thành file SQL (dùng công cụ của database như `pg_dump`), đặt làm migration đầu tiên (baseline), rồi dùng lệnh `goose mark-applied` để đánh dấu rằng migration này đã được áp dụng mà không cần chạy lại. Từ đó về sau, mọi thay đổi schema đều được quản lý qua file migration mới. Nếu cần công cụ tự động phát hiện sự khác biệt giữa schema hiện tại và schema mong muốn, nên xem xét **Atlas** thay vì Goose.