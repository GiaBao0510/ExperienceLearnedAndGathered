# Dependency Injection trong Go

## Mục lục

1. [Định nghĩa](https://claude.ai/chat/d6f7164f-e12c-439e-8152-336e1aa0d9c5#1-%C4%91%E1%BB%8Bnh-ngh%C4%A9a)
2. [Ví dụ: dịch vụ gửi tin nhắn qua Email và SMS](https://claude.ai/chat/d6f7164f-e12c-439e-8152-336e1aa0d9c5#2-v%C3%AD-d%E1%BB%A5-d%E1%BB%8Bch-v%E1%BB%A5-g%E1%BB%ADi-tin-nh%E1%BA%AFn-qua-email-v%C3%A0-sms)
3. [Constructor injection kiểu Go: hàm `NewXxx()`](https://claude.ai/chat/d6f7164f-e12c-439e-8152-336e1aa0d9c5#3-constructor-injection-ki%E1%BB%83u-go-h%C3%A0m-newxxx)
4. [Vì sao cách làm này giúp Unit Test dễ hơn?](https://claude.ai/chat/d6f7164f-e12c-439e-8152-336e1aa0d9c5#4-v%C3%AC-sao-c%C3%A1ch-l%C3%A0m-n%C3%A0y-gi%C3%BAp-unit-test-d%E1%BB%85-h%C6%A1n)
5. [Tổng kết](https://claude.ai/chat/d6f7164f-e12c-439e-8152-336e1aa0d9c5#5-t%E1%BB%95ng-k%E1%BA%BFt)
6. [Thực hành: Áp dụng Google Wire để tự động hóa DI](https://claude.ai/chat/390b62f6-1be1-462b-92a7-dff70a15edea#6-th%E1%BB%B1c-h%C3%A0nh-%C3%A1p-d%E1%BB%A5ng-google-wire-%C4%91%E1%BB%83-t%E1%BB%B1-%C4%91%E1%BB%99ng-h%C3%B3a-di)
7. [Mở rộng](https://claude.ai/chat/d6f7164f-e12c-439e-8152-336e1aa0d9c5#m%E1%BB%9F-r%E1%BB%99ng)

> **Phần lý thuyết:** Nếu bạn chưa nắm khái niệm Dependency Injection, DIP, IoC, hãy đọc trước tài liệu `DI_Introduction.md` trước khi xem ví dụ Go dưới đây.

---

## 1. Định nghĩa

Dependency Injection là kỹ thuật giúp giảm sự phụ thuộc giữa các module code, bằng cách "tiêm" các dependency (các đối tượng hoặc dữ liệu mà một đối tượng cần để thực hiện công việc của nó) từ bên ngoài, thay vì để đối tượng tự tạo ra chúng.

Trong Go, điều này thường được thực hiện thông qua **interface**: một đối tượng chỉ cần định nghĩa những gì nó cần (thông qua một interface), và các đối tượng khác có thể cung cấp đúng thứ nó cần, miễn là chúng tuân thủ interface đó — không cần biết chi tiết bên trong cách triển khai.

---

## 2. Ví dụ: dịch vụ gửi tin nhắn qua Email và SMS

Trong ví dụ này, chúng ta sẽ xây dựng một dịch vụ gửi tin nhắn, có thể gửi qua Email hoặc SMS mà không cần sửa code ở nơi sử dụng dịch vụ.

**Bước 1 — Định nghĩa interface `MessageService`:**

Interface này không quy định _cách_ gửi tin nhắn, chỉ yêu cầu bất kỳ "dịch vụ tin nhắn" nào cũng phải có khả năng gửi tin nhắn.

```go
package messaging

// MessageService định nghĩa hành vi mà mọi dịch vụ gửi tin nhắn phải có.
type MessageService interface {
    SendMessage(message string, receiver string) error
}
```

**Bước 2 — Cài đặt hai implementation cụ thể:**

`SMSService` và `EmailService` cùng implement `MessageService`, nhưng cách gửi tin nhắn thực tế của mỗi loại khác nhau.

```go
package messaging

import "fmt"

// SMSService là một cách triển khai của MessageService, gửi qua SMS.
type SMSService struct{}

func (s *SMSService) SendMessage(message string, receiver string) error {
    fmt.Printf("Send SMS: %s to %s\n", message, receiver)
    return nil
}

// EmailService là một cách triển khai khác của MessageService, gửi qua Email.
type EmailService struct{}

func (e *EmailService) SendMessage(message string, receiver string) error {
    fmt.Printf("Send Email: %s to %s\n", message, receiver)
    return nil
}
```

**Bước 3 — `MyApplication` chỉ phụ thuộc vào interface, không phụ thuộc vào implementation cụ thể:**

```go
package messaging

// MyApplication sử dụng một MessageService để gửi tin nhắn.
type MyApplication struct {
    messageService MessageService
}

func (a *MyApplication) ProcessMessages(message string, receiver string) error {
    return a.messageService.SendMessage(message, receiver)
}
```

`MyApplication` không cần biết chi tiết _cách_ gửi tin nhắn diễn ra — nó chỉ cần biết rằng mình có một `messageService` có thể gửi tin nhắn. Đây chính là Dependency Injection: `MyApplication` **không tự tạo ra** một `MessageService`, mà **được cung cấp** một `MessageService` từ bên ngoài.

---

## 3. Constructor injection kiểu Go: hàm `NewXxx()`

Go không có khái niệm "constructor" theo đúng nghĩa như Java/C# (không có từ khóa `constructor`). Thay vào đó, quy ước phổ biến trong Go là viết một hàm khởi tạo tên `NewXxx()`, trả về một con trỏ tới struct đã được khởi tạo đầy đủ — đây chính là cách Go hiện thực hóa **Constructor injection** đã nhắc đến ở tài liệu lý thuyết.

```go
// NewMyApplication là "constructor" theo quy ước của Go — nhận vào
// MessageService từ bên ngoài (Constructor injection), thay vì để
// MyApplication tự khởi tạo SMSService/EmailService bên trong nó.
func NewMyApplication(ms MessageService) *MyApplication {
    return &MyApplication{messageService: ms}
}

// SetMessageService minh họa Setter injection — kiểu DI thứ hai đã nêu
// ở tài liệu lý thuyết, cho phép đổi dependency sau khi đã khởi tạo.
func (a *MyApplication) SetMessageService(ms MessageService) {
    a.messageService = ms
}
```

**`main.go` — nơi thực hiện việc "tiêm" (inject) dependency:**

```go
package main

import "yourmodule/messaging"

func main() {
    // Constructor injection: truyền SMSService vào ngay lúc khởi tạo MyApplication
    smsService := &messaging.SMSService{}
    app := messaging.NewMyApplication(smsService)
    app.ProcessMessages("Hello World", "0987654321")

    // Setter injection: gán trực tiếp một implementation khác (EmailService)
    // vào field messageService sau khi app đã được khởi tạo
    emailService := &messaging.EmailService{}
    app.SetMessageService(emailService)
    app.ProcessMessages("Hello World", "abc@example.com")
}
```

> **Lưu ý:** Ví dụ ở đây tách rõ 2 kiểu injection để bạn dễ đối chiếu với lý thuyết — trong dự án thực tế, bạn thường chỉ cần chọn **một** cách (thường là Constructor injection) để tránh object ở trạng thái không nhất quán giữa lúc khởi tạo và lúc bị đổi dependency giữa chừng.

---

## 4. Vì sao cách làm này giúp Unit Test dễ hơn?

Đây là lợi ích lớn nhất của DI đã nêu ở tài liệu lý thuyết (mục 8), và có thể chứng minh trực tiếp bằng code. Vì `MyApplication` chỉ phụ thuộc vào interface `MessageService`, ta có thể tạo một **mock implementation** giả lập việc gửi tin nhắn — không cần gửi SMS/Email thật khi chạy test:

```go
package messaging

import "testing"

// mockMessageService là một implementation giả lập của MessageService,
// chỉ dùng cho mục đích test — không thực sự gửi tin nhắn.
type mockMessageService struct {
    called   bool
    lastMsg  string
    lastRecv string
}

func (m *mockMessageService) SendMessage(message string, receiver string) error {
    m.called = true
    m.lastMsg = message
    m.lastRecv = receiver
    return nil
}

func TestMyApplication_ProcessMessages(t *testing.T) {
    mock := &mockMessageService{}
    app := NewMyApplication(mock) // tiêm mock thay vì SMSService/EmailService thật

    err := app.ProcessMessages("Hello Test", "tester@example.com")

    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }
    if !mock.called {
        t.Fatal("expected SendMessage to be called")
    }
    if mock.lastMsg != "Hello Test" || mock.lastRecv != "tester@example.com" {
        t.Fatalf("unexpected message/receiver: %s / %s", mock.lastMsg, mock.lastRecv)
    }
}
```

Nếu `MyApplication` tự khởi tạo `SMSService` bên trong nó (không dùng DI), bạn sẽ **không có cách nào** thay thế nó bằng `mockMessageService` khi test — buộc phải gửi SMS/Email thật mỗi lần chạy test, hoặc phải sửa code sản xuất chỉ để phục vụ việc test. Đây chính là lý do DI và khả năng kiểm thử (testability) luôn đi cùng nhau trong thực tế.

---

## 5. Tổng kết

- Trong Go, DI được hiện thực hóa chủ yếu qua **interface** — không cần framework hay annotation đặc biệt như Java.
- **Constructor injection** kiểu Go: dùng hàm `NewXxx()` nhận dependency qua tham số, đây là cách được khuyến nghị dùng trong phần lớn trường hợp.
- **Setter injection** vẫn khả thi (gán trực tiếp vào field hoặc qua setter method), nhưng nên cân nhắc kỹ vì có thể khiến object rơi vào trạng thái không nhất quán giữa các lần đổi dependency.
- Lợi ích rõ ràng nhất khi áp dụng DI vào code Go: viết được unit test dùng **mock**, không cần phụ thuộc vào dịch vụ thật (SMS gateway, SMTP server...) khi chạy test.

---

## 6. Thực hành: Áp dụng Google Wire để tự động hóa DI

Ở mục 3, chúng ta tự tay viết hàm `NewMyApplication(ms MessageService)` và tự tay gọi nó trong `main.go`. Với 1-2 dependency thì việc này rất nhẹ nhàng. Nhưng hãy tưởng tượng `MyApplication` cần 10 service, mỗi service lại cần vài dependency con bên trong (ví dụ `EmailService` cần thêm `SMTPConfig`, `Logger`...) — lúc đó `main.go` sẽ chứa hàng chục dòng khởi tạo thủ công, rất dễ khởi tạo sai thứ tự hoặc quên một dependency.

**Google Wire** giải quyết vấn đề này bằng cách: bạn chỉ khai báo "tôi có những provider nào" và "tôi muốn khởi tạo cái gì", Wire sẽ **tự sinh code** (giống hệt code bạn tự viết tay ở mục 3) tại compile-time. Vì code được sinh ra là code Go bình thường (không dùng reflection lúc runtime), nên hiệu năng không bị ảnh hưởng và lỗi thiếu/sai dependency được phát hiện ngay lúc build, không phải lúc chạy chương trình.

### 6.1. Cài đặt Wire

```bash
go install github.com/google/wire/cmd/wire@latest
```

Sau khi cài, lệnh `wire` sẽ có sẵn trong `$GOPATH/bin` (nhớ đảm bảo thư mục này nằm trong biến môi trường `PATH`).

Thêm package `wire` vào project để dùng annotation `wire.Build`:

```bash
go get github.com/google/wire
```

### 6.2. Khái niệm cốt lõi của Wire

|Khái niệm|Ý nghĩa|
|---|---|
|**Provider**|Một hàm biết cách tạo ra một giá trị, ví dụ `NewMyApplication` — chính là các hàm `NewXxx()` bạn đã quen ở mục 3|
|**Provider Set**|Một nhóm các provider được gom lại bằng `wire.NewSet(...)`, để dùng lại nhiều lần|
|**Injector**|Hàm "khung" do bạn khai báo (thân hàm gọi `wire.Build(...)`), Wire sẽ đọc hàm này và **sinh ra code thật** thay thế nó|
|**`wire_gen.go`**|File do Wire tự sinh ra, chứa code khởi tạo dependency thật — đây là file được compile và chạy|

### 6.3. Áp dụng vào ví dụ `MessageService`

Giả sử cấu trúc thư mục project như sau:

```
yourmodule/
├── messaging/
│   ├── message_service.go   // interface + SMSService/EmailService (mục 2)
│   └── application.go       // MyApplication + NewMyApplication (mục 3)
└── cmd/
    └── app/
        ├── wire.go           // khai báo provider set + injector (mục 6.3)
        ├── wire_gen.go        // do lệnh `wire` tự sinh ra, KHÔNG tự sửa tay
        └── main.go            // chỉ gọi injector, không tự khởi tạo dependency
```

**Bước 1 — Viết provider cho từng implementation cụ thể**

Wire cần một hàm "provider" trả về đúng kiểu `MessageService` (interface), bọc quanh implementation cụ thể muốn dùng. Ta thêm hàm này ngay trong package `messaging`:

```go
// messaging/message_service.go

package messaging

// NewEmailMessageService là provider cho Wire: trả về MessageService
// dưới dạng interface, nhưng bên trong thực chất khởi tạo EmailService.
// Đây chính là nơi bạn "chọn" implementation nào sẽ được dùng thật sự.
func NewEmailMessageService() MessageService {
    return &EmailService{}
}

// Nếu muốn dùng SMS thay vì Email, chỉ cần viết thêm provider tương tự
// rồi đổi provider được liệt kê trong wire.Build (xem bước 3) — không
// cần sửa bất kỳ dòng nào trong MyApplication hay main.go.
func NewSMSMessageService() MessageService {
    return &SMSService{}
}
```

**Bước 2 — File `wire.go`: khai báo Provider Set và Injector**

File này có build tag `//go:build wireinject` — nghĩa là nó **chỉ tồn tại để Wire đọc**, không bao giờ được compile vào binary thật (Wire sẽ thay nó bằng `wire_gen.go`).

```go
//go:build wireinject
// +build wireinject

// wire.go — "bản thiết kế" cho Wire, không chứa logic thật.
// Lệnh `wire` sẽ đọc file này để sinh ra wire_gen.go.
package main

import (
    "github.com/google/wire"
    "yourmodule/messaging"
)

// ProviderSet gom các provider liên quan đến messaging lại một chỗ,
// để có thể tái sử dụng ở nhiều injector khác nhau nếu project lớn dần.
var MessagingProviderSet = wire.NewSet(
    messaging.NewEmailMessageService, // chọn Email làm implementation mặc định
    messaging.NewMyApplication,       // constructor injection ở mục 3
)

// InitializeApp là injector: khai báo "tôi muốn có một *MyApplication".
// Thân hàm chỉ gọi wire.Build — Wire sẽ phân tích kiểu trả về
// (*messaging.MyApplication) và tự nối các provider ở trên lại với nhau
// theo đúng thứ tự dependency, y hệt như bạn tự viết tay ở mục 3.
func InitializeApp() *messaging.MyApplication {
    wire.Build(MessagingProviderSet)
    return nil // dòng này không bao giờ chạy thật, Wire yêu cầu có để hàm hợp lệ về mặt cú pháp
}
```

**Bước 3 — Sinh code bằng lệnh `wire`**

Chạy trong thư mục chứa `wire.go`:

```bash
cd cmd/app
wire
```

Wire sẽ đọc `wire.go`, phân tích đồ thị dependency, và tạo ra file `wire_gen.go` — ví dụ Wire sẽ sinh ra nội dung tương đương thế này (bạn không cần gõ tay, chỉ cần hiểu để debug khi cần):

```go
// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import "yourmodule/messaging"

// InitializeApp là bản triển khai THẬT, do Wire tự sinh dựa trên
// wire.Build ở wire.go — logic hoàn toàn giống code bạn tự viết tay
// ở mục 3, chỉ khác là được sinh tự động và luôn đồng bộ với provider set.
func InitializeApp() *messaging.MyApplication {
    messageService := messaging.NewEmailMessageService()
    myApplication := messaging.NewMyApplication(messageService)
    return myApplication
}
```

**Bước 4 — `main.go` chỉ còn việc gọi injector**

```go
// cmd/app/main.go
package main

func main() {
    // main.go không còn biết (và không cần biết) EmailService hay SMSService
    // được tạo ra như thế nào — mọi việc khởi tạo dependency đã được
    // Wire lo tại compile-time thông qua InitializeApp().
    app := InitializeApp()
    app.ProcessMessages("Hello World", "abc@example.com")
}
```

### 6.4. Vì sao cách này vẫn giữ được lợi ích testability ở mục 4?

`wire_gen.go` chỉ ảnh hưởng đến **cách `main.go` khởi tạo dependency thật khi chạy chương trình**. Trong file test (`application_test.go`), bạn hoàn toàn không đụng đến Wire — vẫn gọi thẳng `NewMyApplication(mock)` như mục 4, vì `NewMyApplication` vẫn nhận `MessageService` qua tham số như bình thường. Wire không thay đổi bản chất của constructor injection, nó chỉ tự động hóa việc _gọi_ các constructor đó theo đúng thứ tự.

### 6.5. Khi nào nên dùng Wire, khi nào không cần?

|Tình huống|Nên dùng Wire?|
|---|---|
|Project nhỏ, 1-3 dependency, ít thay đổi|**Không cần** — tự viết tay như mục 3 vẫn rõ ràng và dễ đọc hơn|
|Dependency graph lớn (10+ service, nhiều tầng phụ thuộc lồng nhau)|**Nên dùng** — tránh boilerplate và lỗi khởi tạo sai thứ tự|
|Muốn dễ dàng đổi qua lại giữa nhiều implementation (Email ⇄ SMS, thật ⇄ mock) theo môi trường (dev/staging/prod)|**Nên dùng** — chỉ cần đổi provider trong `wire.Build`, không sửa `main.go`|
|Muốn phát hiện lỗi thiếu dependency ngay lúc build thay vì lúc chạy|**Nên dùng** — đây là ưu điểm cốt lõi của Wire so với các DI container dùng reflection runtime|

---

### Mở rộng

- **Google Wire**: đã trình bày chi tiết ở [mục 6](https://claude.ai/chat/390b62f6-1be1-462b-92a7-dff70a15edea#6-th%E1%BB%B1c-h%C3%A0nh-%C3%A1p-d%E1%BB%A5ng-google-wire-%C4%91%E1%BB%83-t%E1%BB%B1-%C4%91%E1%BB%99ng-h%C3%B3a-di) — công cụ sinh code DI cho Go lúc compile-time, giúp giảm boilerplate khi đồ thị dependency phức tạp với hàng chục service.
- **`gomock`/`testify/mock`**: thay vì tự viết `mockMessageService` thủ công như mục 4, các thư viện này giúp tự động sinh mock từ interface — hữu ích khi interface có nhiều phương thức hoặc dự án có nhiều interface cần mock.
- **Context-based dependency (dùng `context.Context`)**: một cách "tiêm" dữ liệu khác trong Go, thường dùng cho các giá trị theo từng request (request-scoped) như `trace_id`, `user_id` — khác với DI ở cấp độ struct/service đã học trong tài liệu này, nhưng cùng chung mục tiêu giảm truyền tham số thủ công qua nhiều lớp gọi hàm.
- **Áp dụng vào dự án thực tế:** thử áp dụng DI vào chính dự án e-commerce Go của bạn — ví dụ tách `PaymentService` thành interface, viết thêm `MockPaymentService` để test luồng đặt hàng mà không cần gọi cổng thanh toán thật.