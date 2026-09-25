# Go Testing — Kiến thức & Thực hành Testing trong Golang

## 1. Mục tiêu của thư mục này

Thư mục này tổng hợp kiến thức về **testing/unit test** trong Golang, đi từ cơ bản đến nâng cao, nhằm:

- Nắm vững cách viết unit test đúng chuẩn Go (`testing` package, quy ước `*_test.go`).
- Biết cách viết test dễ đọc, dễ mở rộng bằng **table-driven tests**.
- Sử dụng thành thạo các thư viện hỗ trợ assertion và mocking phổ biến (`testify`, `gomock`).
- Biết cách test các thành phần thường gặp trong backend: HTTP handler, database layer, service layer.
- Hiểu và đo lường **test coverage**, viết **benchmark** để đánh giá hiệu năng.
- Phân biệt rõ Unit Test / Integration Test và biết khi nào dùng loại nào.
- Từng bước hình thành tư duy **TDD (Test-Driven Development)** để áp dụng vào các dự án thực tế.

Đây không chỉ là nơi ghi chú lý thuyết, mà còn là nơi lưu lại các ví dụ thực hành để có thể tái sử dụng khi làm dự án thật.

## 2. Tại sao cần học và áp dụng testing?

Là một Backend Engineer, testing không phải là phần "làm thêm cho có" mà là một phần bắt buộc của quy trình phát triển phần mềm chuyên nghiệp, vì:

- **Phát hiện lỗi sớm**: Bug được phát hiện ở giai đoạn viết code luôn rẻ hơn rất nhiều so với phát hiện ở production.
- **Tự tin khi refactor**: Có test tốt, bạn có thể sửa/tối ưu code mà không sợ làm hỏng chức năng cũ.
- **Tài liệu sống (living documentation)**: Test case mô tả rõ behavior mong đợi của function/module, người đọc sau (kể cả chính mình sau này) hiểu code nhanh hơn.
- **CI/CD an toàn hơn**: Test tự động là điều kiện tiên quyết để build pipeline CI/CD đáng tin cậy, tránh deploy code lỗi.
- **Chuẩn hóa tư duy thiết kế code**: Code dễ test thường là code có kiến trúc tốt (loose coupling, dependency injection, single responsibility) — đây cũng là kỹ năng senior engineer cần có.
- **Thể hiện chuyên môn**: Trong phỏng vấn hoặc code review, khả năng viết test tốt là một tiêu chí đánh giá năng lực rõ ràng.

## 3. Cấu trúc thư mục

```
Go_Testing/
├── 1_Unit_Testing_Basics/
├── 2_Table_Driven_Tests/
├── 3_Assertion_Libraries/
├── 4_Mocking/
├── 5_HTTP_API_Testing/
├── 6_Database_Testing/
├── 7_Benchmark_Testing/
├── 8_Test_Coverage/
├── 9_Integration_Testing/
└── 10_TDD_BestPractices.md
```

### 3.1. `1_Unit_Testing_Basics/`

Nền tảng của testing trong Go: cách đặt tên file `*_test.go`, hàm `TestXxx(t *testing.T)`, cách chạy test bằng `go test`, dùng `t.Run()` để tạo subtest, các hàm `t.Error`, `t.Fatal`, `t.Log`. Đây là điểm khởi đầu bắt buộc trước khi học các phần nâng cao hơn.

### 3.2. `2_Table_Driven_Tests/`

Pattern phổ biến nhất trong cộng đồng Go để viết test gọn, dễ mở rộng: định nghĩa một slice/struct chứa nhiều test case (input, expected output), sau đó loop qua để test. Giúp tránh lặp code và dễ dàng thêm case mới.

### 3.3. `3_Assertion_Libraries/`

Go's `testing` package mặc định khá "thô" (chỉ có `if` + `t.Error`). Thư mục này ghi chú cách dùng `testify/assert` và `testify/require` để viết assertion ngắn gọn, dễ đọc hơn, cùng với sự khác biệt giữa `assert` (tiếp tục chạy) và `require` (dừng ngay khi fail).

### 3.4. `4_Mocking/`

Khi test một function phụ thuộc vào service/interface bên ngoài (DB, API thứ 3...), cần "giả lập" (mock) các dependency đó. Ghi chú cách dùng `testify/mock` và `gomock` (kèm `mockgen`) để tạo mock object, cách thiết kế interface sao cho dễ mock — liên quan trực tiếp đến nguyên tắc Dependency Injection.

### 3.5. `5_HTTP_API_Testing/`

Test các HTTP handler/API endpoint mà không cần chạy server thật, dùng package `net/http/httptest`: `httptest.NewRecorder()`, `httptest.NewServer()`. Phù hợp áp dụng trực tiếp vào các project REST API đã có sẵn (Gin handler, router...).

### 3.6. `6_Database_Testing/`

Hai hướng tiếp cận chính:

- **Mock DB**: dùng `sqlmock` để giả lập câu query SQL, test nhanh, không cần DB thật.
- **Test với DB thật (containerized)**: dùng `testcontainers-go` để spin up Postgres/MongoDB tạm thời trong Docker, phù hợp cho integration test.

### 3.7. `7_Benchmark_Testing/`

Đo hiệu năng function bằng `go test -bench`, viết hàm `BenchmarkXxx(b *testing.B)`, kết hợp `pprof` để phân tích CPU/memory profiling — hữu ích khi tối ưu hiệu năng code Go.

### 3.8. `8_Test_Coverage/`

Cách đo độ bao phủ test bằng `go test -cover`, xuất báo cáo HTML (`go tool cover -html`), và quan trọng hơn: hiểu rằng coverage cao không đồng nghĩa với test tốt — tránh chạy theo con số ảo.

### 3.9. `9_Integration_Testing/`

Phân biệt rõ Unit Test (test 1 đơn vị code, cô lập bằng mock) và Integration Test (test nhiều thành phần phối hợp với nhau, có thể động tới DB/message queue thật). Ghi chú cách dùng **build tags** (`//go:build integration`) để tách 2 loại test khi chạy CI, và cách tận dụng `docker-compose` sẵn có trong các project (`Go_kafka`, `Go_MongoDB`...) làm môi trường test.

### 3.10. `10_TDD_BestPractices.md`

Tổng hợp tư duy Test-Driven Development (Red - Green - Refactor), các nguyên tắc viết test tốt (F.I.R.S.T principles: Fast, Independent, Repeatable, Self-validating, Timely), và các anti-pattern thường gặp khi viết test (test phụ thuộc thứ tự chạy, test phụ thuộc thời gian thực, test quá chi tiết vào implementation...).

## 4. Cách áp dụng vào các project đã có

Ngoài việc học lý thuyết trong `Go_Testing`, nên viết test trực tiếp vào các project thực hành sẵn có trong repo để nhớ lâu và có kinh nghiệm thật:

| Project sẵn có                                    | Loại test nên viết                 | Kỹ thuật áp dụng                        |
| ------------------------------------------------- | ---------------------------------- | --------------------------------------- |
| `GO_connect/GO_PostgreSQL/repository/user.go`     | Unit test cho repository           | `sqlmock`                               |
| `GO_web/go_RESTFUL_API/.../handler/product.go`    | Test HTTP handler                  | `httptest` + table-driven               |
| `problem_solving/error-handling/internal/service` | Unit test cho service layer        | `testify/mock` cho repository interface |
| `GO_connect/Go_kafka`                             | Integration test producer/consumer | `testcontainers` + build tag            |
| `GO_lib/1.SQLC`                                   | Test các query đã generate         | `sqlmock` hoặc DB test thật             |

## 5. Đề xuất bổ sung thêm

Một số hướng có thể mở rộng thêm khi đã vững các phần trên:

- **Fuzz Testing**: Go 1.18+ hỗ trợ native fuzzing (`go test -fuzz`), rất hữu ích để test các hàm xử lý input/parsing.
- **Golden File Testing**: kỹ thuật so sánh output với file mẫu (`.golden`), thường dùng khi test output phức tạp (JSON, HTML render).
- **Contract Testing**: khi làm việc với microservices, tìm hiểu Pact hoặc tương tự để đảm bảo API contract giữa các service không bị phá vỡ.
- **CI Integration**: viết thêm ghi chú về cách chạy `go test` trong GitHub Actions, kèm coverage report tự động.
- **Testing gRPC**: bổ sung riêng vì repo đã có `RPC_gRPC`, dùng `bufconn` để test gRPC service mà không cần mở port thật.

## 6. Ghi chú

Repo này được xây dựng song song với quá trình tự học Go hằng ngày, mục tiêu là hệ thống hóa kiến thức backend một cách có cấu trúc và có thể tái sử dụng khi làm dự án thực tế hoặc chuẩn bị phỏng vấn.
