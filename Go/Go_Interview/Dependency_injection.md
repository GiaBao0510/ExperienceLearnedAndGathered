## Định Nghĩa

**Dependency Injection (DI)** là một kỹ thuật lập trình giúp giảm sự phụ thuộc trực tiếp giữa các module bằng cách "tiêm" (inject) các dependency — tức các đối tượng hoặc dữ liệu mà một đối tượng cần để thực hiện công việc của nó — từ bên ngoài vào, thay vì để đối tượng tự tạo ra chúng bên trong.

Dependency Injection là một dạng triển khai cụ thể của nguyên lý **Inversion of Control (IoC)** — nguyên lý nói rằng luồng điều khiển nên được đảo ngược, tức là framework hoặc container bên ngoài sẽ quyết định khi nào và cách nào tạo ra các đối tượng, thay vì bản thân đối tượng tự quyết định.

Dependency Injection cũng là cách hiện thực hóa nguyên lý **Dependency Inversion Principle (DIP)** — chữ "D" trong SOLID — với ba ý chính:

- Các module cấp cao không nên phụ thuộc vào các module cấp thấp. Cả hai nên phụ thuộc vào abstraction.
- Abstraction (interface) không nên phụ thuộc vào chi tiết cài đặt, mà ngược lại.
- Các class giao tiếp với nhau thông qua interface, không phải thông qua implementation cụ thể.

> 📚 **Tham khảo**: Martin, R.C. (2003). _Agile Software Development, Principles, Patterns, and Practices_. Prentice Hall. — Chương về Dependency Inversion Principle.

![](https://s3-hfx03.fptcloud.com/codelearnstorage/Media/Default/Users/VuTungMinh/Template/aa.jpeg)

---

## Tại Sao Phải Dùng Dependency Injection?

Hãy xem xét trường hợp **không dùng DI**:

```go
type MyApplication struct{}

func (a *MyApplication) processMessages(message string, receiver string) error {
    // Tạo dependency trực tiếp bên trong — BAD PRACTICE
    smsService := &SMSService{}
    return smsService.SendMessage(message, receiver)
}
```

Vấn đề ở đây:

- `MyApplication` bị "khóa chặt" vào `SMSService`. Nếu muốn dùng `EmailService`, phải sửa code bên trong.
- Không thể viết unit test độc lập vì không thể thay thế `SMSService` bằng một mock object.
- Vi phạm nguyên lý Open/Closed — class không mở rộng được mà không sửa đổi.

**Dependency Injection giải quyết những vấn đề này** bằng cách chuyển trách nhiệm tạo và cung cấp dependency ra bên ngoài.

---

## Các Loại Dependency Injection

Có ba loại DI phổ biến:

### 1. Constructor Injection

Dependency được truyền vào thông qua constructor (hàm khởi tạo). Đây là hình thức được khuyến nghị sử dụng nhất vì dependency rõ ràng, bắt buộc phải cung cấp ngay khi khởi tạo đối tượng.

```go
// Interface định nghĩa hành vi
type IMessageService interface {
    SendMessage(message string, receiver string) error
}

// Struct phụ thuộc vào interface, không phải implementation cụ thể
type MyApplication struct {
    messageService IMessageService
}

// Constructor Injection: dependency được truyền qua hàm khởi tạo
func NewMyApplication(service IMessageService) *MyApplication {
    return &MyApplication{messageService: service}
}

func (a *MyApplication) processMessages(message string, receiver string) error {
    return a.messageService.SendMessage(message, receiver)
}
```

**Sử dụng:**

```go
func main() {
    smsService := &SMSService{}
    app := NewMyApplication(smsService) // inject qua constructor
    app.processMessages("Xin chào!", "0901234567")
}
```

✅ **Ưu điểm**: Dependency bắt buộc và rõ ràng; đối tượng luôn ở trạng thái hợp lệ sau khi khởi tạo.  
⚠️ **Nhược điểm**: Constructor có thể trở nên dài nếu có quá nhiều dependency.

---
### 2. Setter Injection

Dependency được truyền vào sau khi đối tượng đã được khởi tạo, thông qua các phương thức setter. Phù hợp với các dependency tùy chọn (optional).

```go
type MyApplication struct {
    messageService IMessageService
}

// Setter Injection: dependency được truyền qua phương thức setter
func (a *MyApplication) SetMessageService(service IMessageService) {
    a.messageService = service
}

func (a *MyApplication) processMessages(message string, receiver string) error {
    if a.messageService == nil {
        return fmt.Errorf("messageService chưa được khởi tạo")
    }
    return a.messageService.SendMessage(message, receiver)
}
```

**Sử dụng:**

```go
func main() {
    app := &MyApplication{} // khởi tạo trước
    
    smsService := &SMSService{}
    app.SetMessageService(smsService) // inject sau qua setter
    app.processMessages("Xin chào!", "0901234567")

    // Dễ dàng thay thế sang EmailService mà không cần tạo đối tượng mới
    emailService := &EmailService{}
    app.SetMessageService(emailService)
    app.processMessages("Xin chào!", "abc@example.com")
}
```

✅ **Ưu điểm**: Linh hoạt, có thể thay đổi dependency sau khi khởi tạo; phù hợp với dependency tùy chọn.  
⚠️ **Nhược điểm**: Đối tượng có thể ở trạng thái không hợp lệ nếu quên gọi setter; dễ gây lỗi runtime.

---

### 3. Interface Injection

Đây là hình thức ít phổ biến nhất. Dependency tự định nghĩa một interface có chứa phương thức `Inject`. Client (đối tượng cần dependency) phải implement interface đó để nhận dependency.

```go
// Interface do dependency định nghĩa, yêu cầu client phải implement
type IMessageServiceInjector interface {
    InjectMessageService(service IMessageService)
}

// MyApplication implement interface IMessageServiceInjector
type MyApplication struct {
    messageService IMessageService
}

// Phương thức Inject được gọi bởi container/framework từ bên ngoài
func (a *MyApplication) InjectMessageService(service IMessageService) {
    a.messageService = service
}

func (a *MyApplication) processMessages(message string, receiver string) error {
    return a.messageService.SendMessage(message, receiver)
}
```

**Sử dụng (thường do framework/container điều khiển):**

```go
func InjectAll(app IMessageServiceInjector, service IMessageService) {
    app.InjectMessageService(service) // framework gọi phương thức inject
}

func main() {
    app := &MyApplication{}
    smsService := &SMSService{}
    InjectAll(app, smsService) // container/framework thực hiện inject
    app.processMessages("Xin chào!", "0901234567")
}
```

✅ **Ưu điểm**: Framework/container có toàn quyền kiểm soát việc inject.  
⚠️ **Nhược điểm**: Client bị ràng buộc với interface của dependency; phức tạp hơn, ít dùng trong thực tế.

---
### So Sánh Ba Loại

|Tiêu chí|Constructor Injection|Setter Injection|Interface Injection|
|---|---|---|---|
|Thời điểm inject|Khi khởi tạo|Sau khi khởi tạo|Sau khi khởi tạo|
|Dependency bắt buộc?|Có|Không|Không|
|Dễ test?|✅ Rất dễ|✅ Dễ|⚠️ Trung bình|
|Mức độ phổ biến|⭐⭐⭐ Cao nhất|⭐⭐ Trung bình|⭐ Thấp|
|Phù hợp với|Dependency cốt lõi|Dependency tùy chọn|Framework/IoC container|

> 📚 **Tham khảo**: Seemann, M. (2011). _Dependency Injection in .NET_. Manning Publications. — Chương 3: DI Patterns.

---

## Nhiệm Vụ Cụ Thể Của Dependency Injection

Một DI container/framework thực hiện ba nhiệm vụ chính:

1. **Tạo ra các object** (Object creation).
2. **Biết được class nào cần những object đó** (Dependency resolution).
3. **Cung cấp cho những class đó những object chúng cần** (Dependency provision).

---

## Lợi Ích Và Bất Cập Khi Dùng Dependency Injection

### Lợi ích:

- **Dễ test và viết Unit Test**: Khi có thể inject dependency vào một class, ta cũng dễ dàng "tiêm" mock object vào để kiểm thử độc lập từng thành phần.
- **Dễ dàng thấy quan hệ giữa các object**: DI inject các object phụ thuộc vào interface của object bị phụ thuộc, nên ta dễ dàng thấy được toàn bộ dependency của một object.
- **Dễ mở rộng ứng dụng**: Thêm tính năng mới chỉ cần tạo implementation mới mà không cần sửa code hiện tại — tuân thủ Open/Closed Principle.
- **Giảm sự kết dính (coupling)**: Các thành phần ít phụ thuộc lẫn nhau hơn, thay đổi một chỗ ít ảnh hưởng đến chỗ khác.

### Bất lợi:

- **Độ phức tạp tăng**: DI có thể khó học ban đầu và nếu dùng quá mức có thể làm code trở nên khó theo dõi.
- **Lỗi bị đẩy sang runtime**: Nhiều lỗi compile-time có thể bị đẩy sang runtime, khó debug hơn. Vì sử dụng interface nên đôi khi khó biết implementation nào thực sự được truyền vào.
- **Ảnh hưởng IDE**: Có thể làm giảm hiệu quả của tính năng auto-complete hay "find references" trong một số IDE vì dependency bị ẩn sau interface.

---
## Ví Dụ Cụ Thể (Chưa Áp Dụng Thư Viện)

Để minh họa cho khái niệm Dependency Injection, chúng ta sẽ xây dựng một dịch vụ gửi tin nhắn qua Email và SMS.

### Bước 1: Định Nghĩa Interface

Trước tiên, chúng ta định nghĩa một interface `IMessageService` với phương thức `SendMessage`. Interface này không quy định cách thức gửi tin nhắn mà chỉ yêu cầu bất kỳ dịch vụ nào thực hiện interface này phải có khả năng gửi tin nhắn.

```go
// IMessageService định nghĩa hành vi gửi tin nhắn
type IMessageService interface {
    SendMessage(message string, receiver string) error
}
```

### Bước 2: Implement Interface

Tiếp theo, tạo hai struct `SMSService` và `EmailService`, cả hai đều implement `IMessageService`. Mỗi dịch vụ cung cấp phương thức `SendMessage` nhưng với cách thực hiện khác nhau:

```go
// SMSService là một implementation của IMessageService
type SMSService struct{}

func (s *SMSService) SendMessage(message string, receiver string) error {
    fmt.Printf("Đã gửi SMS: %s đến %s\n", message, receiver)
    return nil
}

// EmailService là một implementation khác của IMessageService
type EmailService struct{}

func (e *EmailService) SendMessage(message string, receiver string) error {
    fmt.Printf("Đã gửi Email: %s đến %s\n", message, receiver)
    return nil
}
```

### Bước 3: Tạo MyApplication Struct

Struct `MyApplication` có trường `messageService` kiểu `IMessageService`. Điều này có nghĩa `MyApplication` không cần biết chi tiết về cách gửi tin nhắn; nó chỉ biết rằng nó có một dịch vụ có thể gửi tin nhắn:

```go
// MyApplication sử dụng một service để gửi tin nhắn
type MyApplication struct {
    messageService IMessageService
}

func (a *MyApplication) processMessages(message string, receiver string) error {
    return a.messageService.SendMessage(message, receiver)
}
```

### Bước 4: Thiết Lập Và Sử Dụng

Trong hàm `main`, khởi tạo các dịch vụ và "tiêm" chúng vào `MyApplication`. Đây là ví dụ kết hợp cả Constructor Injection và Setter Injection:

```go
func main() {
    // Constructor Injection
    smsService := &SMSService{}
    app := &MyApplication{messageService: smsService}
    app.processMessages("Xin chào thế giới!", "123")

    // Setter Injection — thay thế service mà không cần tạo app mới
    emailService := &EmailService{}
    app.messageService = emailService
    app.processMessages("Xin chào thế giới!", "abc@example.com")
}
```

---

## Áp Dụng Dependency Injection Với Thư Viện Wire Trong Golang

Trong Go, có các thư viện hỗ trợ DI khá tốt. Một trong số đó là **google/wire** — một công cụ sinh code tự động (code generation) giúp wire các dependency lại với nhau tại compile-time, thay vì dùng reflection tại runtime như một số framework khác.

> 🔗 **Link google/wire**: [https://github.com/google/wire](https://github.com/google/wire)

### Cài Đặt

Chạy lệnh sau trong terminal:

```bash
go install github.com/google/wire/cmd/wire@latest
```

Gõ `wire` để kiểm tra cài đặt thành công.

### Ví Dụ: API Server Sử Dụng Wire

Ở ví dụ này, ta sẽ setup một API server đơn giản dùng `labstack/echo`, có inject `HealthService` thông qua Wire.

**Bước 1 — Định nghĩa service tại `health.go`:**

```go
type IHealthService interface {
    GetHealth() string
}

type healthService struct{}

func (h healthService) GetHealth() string {
    return "Health"
}

// Hàm khởi tạo — Wire sẽ gọi hàm này để tạo dependency
func NewHealthService() IHealthService {
    return &healthService{}
}
```

**Bước 2 — Định nghĩa application tại `main.go`:**

```go
type ApiApplication struct {
    IHealthService service.IHealthService
}

// Constructor Injection — Wire sẽ tự động truyền IHealthService vào đây
func NewApiApplication(IHealthService service.IHealthService) ApiApplication {
    return ApiApplication{IHealthService: IHealthService}
}

func (api ApiApplication) RunApp() {
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        response := api.IHealthService.GetHealth()
        return c.String(http.StatusOK, response)
    })
    e.Logger.Fatal(e.Start(":1323"))
}
```

**Bước 3 — Khai báo Wire providers tại `wire.go`:**

```go
//go:build wireinject

func InitApplication() (ApiApplication, func()) {
    wire.Build(
        NewHealthService,  // provider cho IHealthService
        NewApiApplication, // provider cho ApiApplication
    )
    return ApiApplication{}, func() {}
}
```

**Bước 4 — Chạy lệnh Wire để sinh code:**

```bash
wire ./...
```

Wire sẽ phân tích các provider được khai báo và tự động sinh ra file `wire_gen.go`:

```go
// wire_gen.go — được sinh tự động, KHÔNG chỉnh sửa thủ công
func InitApplication() (ApiApplication, func()) {
    iHealthService := service.NewHealthService()
    apiApplication := NewApiApplication(iHealthService)
    return apiApplication, func() {}
}
```

**Bước 5 — Khởi chạy ứng dụng:**

```go
func main() {
    app, cleanup := InitApplication()
    defer cleanup() // thực hiện dọn dẹp khi app đóng (đóng DB, ghi log, v.v.)
    app.RunApp()
}
```

> 💡 **Lưu ý**: Hàm `cleanup` được trả về để chứa các tác vụ cần thực hiện khi ứng dụng kết thúc, ví dụ: đóng kết nối database, flush log buffer. Khi cần thêm một dependency mới, chỉ cần khai báo provider của nó trong `wire.go` rồi chạy `wire ./...` — Wire sẽ tự động cập nhật `wire_gen.go`.

> 📚 **Tham khảo thêm về Wire**: [Wire User Guide](https://github.com/google/wire/blob/main/docs/guide.md)

---

## Phân Biệt Các Khái Niệm Liên Quan

Nhiều người hay nhầm lẫn giữa các khái niệm sau:

| Khái niệm                                | Ý nghĩa                                                                                                     |
| ---------------------------------------- | ----------------------------------------------------------------------------------------------------------- |
| **Dependency Inversion Principle (DIP)** | Nguyên lý thiết kế: module cấp cao không phụ thuộc vào module cấp thấp; cả hai phụ thuộc vào abstraction.   |
| **Inversion of Control (IoC)**           | Nguyên lý kiến trúc: luồng điều khiển được đảo ngược, framework/container quyết định khi nào tạo đối tượng. |
| **Dependency Injection (DI)**            | Kỹ thuật cụ thể để hiện thực hóa IoC: truyền dependency từ bên ngoài vào thay vì để đối tượng tự tạo.       |

Nói đơn giản: **DIP là nguyên lý → IoC là cách tư duy → DI là cách triển khai cụ thể**.

> 📚 **Tham khảo**: Fowler, M. (2004). _Inversion of Control Containers and the Dependency Injection pattern_. [https://martinfowler.com/articles/injection.html](https://martinfowler.com/articles/injection.html)