# Tài liệu hướng dẫn: Cronjob trong Golang

**Chủ đề:** Giới thiệu Cronjob cho các tác vụ cần thực hiện dựa trên định kỳ  
**Ngôn ngữ:** Go (Golang)  
**Cấp độ:** Sinh viên năm 3-4, người mới tìm hiểu về lập trình backend

---

## Mục lục

1. [Giới thiệu mục tiêu chương trình](https://claude.ai/chat/888670dd-fadd-40d7-8858-5e6911ff6910#1-gi%E1%BB%9Bi-thi%E1%BB%87u-m%E1%BB%A5c-ti%C3%AAu-ch%C6%B0%C6%A1ng-tr%C3%ACnh)
2. [Các thư viện sử dụng](https://claude.ai/chat/888670dd-fadd-40d7-8858-5e6911ff6910#2-c%C3%A1c-th%C6%B0-vi%E1%BB%87n-s%E1%BB%AD-d%E1%BB%A5ng)
3. [Cài đặt môi trường và thư viện](https://claude.ai/chat/888670dd-fadd-40d7-8858-5e6911ff6910#3-c%C3%A0i-%C4%91%E1%BA%B7t-m%C3%B4i-tr%C6%B0%E1%BB%9Dng-v%C3%A0-th%C6%B0-vi%E1%BB%87n)
4. [Cấu trúc thư mục dự án](https://claude.ai/chat/888670dd-fadd-40d7-8858-5e6911ff6910#4-c%E1%BA%A5u-tr%C3%BAc-th%C6%B0-m%E1%BB%A5c-d%E1%BB%B1-%C3%A1n)
5. [Giải thích chi tiết từng file](https://claude.ai/chat/888670dd-fadd-40d7-8858-5e6911ff6910#5-gi%E1%BA%A3i-th%C3%ADch-chi-ti%E1%BA%BFt-t%E1%BB%ABng-file)
6. [Luồng hoạt động tổng thể](https://claude.ai/chat/888670dd-fadd-40d7-8858-5e6911ff6910#6-lu%E1%BB%93ng-ho%E1%BA%A1t-%C4%91%E1%BB%99ng-t%E1%BB%95ng-th%E1%BB%83)
7. [Các đoạn code cần chú ý đặc biệt](https://claude.ai/chat/888670dd-fadd-40d7-8858-5e6911ff6910#7-c%C3%A1c-%C4%91o%E1%BA%A1n-code-c%E1%BA%A7n-ch%C3%BA-%C3%BD-%C4%91%E1%BA%B7c-bi%E1%BB%87t)
8. [Giải thích từ khóa và câu lệnh nâng cao](https://claude.ai/chat/888670dd-fadd-40d7-8858-5e6911ff6910#8-gi%E1%BA%A3i-th%C3%ADch-t%E1%BB%AB-kh%C3%B3a-v%C3%A0-c%C3%A2u-l%E1%BB%87nh-n%C3%A2ng-cao)
9. [Lỗi thường gặp và cách phòng tránh](https://claude.ai/chat/888670dd-fadd-40d7-8858-5e6911ff6910#9-l%E1%BB%97i-th%C6%B0%E1%BB%9Dng-g%E1%BA%B7p-v%C3%A0-c%C3%A1ch-ph%C3%B2ng-tr%C3%A1nh)

---

## 1. Giới thiệu mục tiêu chương trình

### Cronjob là gì?

Cronjob (hay Crontab) là một cơ chế cho phép lập lịch để thực hiện một tác vụ tự động vào những thời điểm nhất định hoặc theo chu kỳ lặp lại. Ví dụ:

- Cứ mỗi 3 giây, hệ thống tự động gửi email cho người dùng VIP.
- Mỗi ngày lúc 9 giờ tối, hệ thống tự động đăng video đã được lập lịch lên YouTube.
- Trước 10 phút khi một sự kiện webinar bắt đầu, hệ thống quét toàn bộ database và gửi email nhắc nhở tới hàng triệu người đã đăng ký.

### Mục tiêu của chương trình này

Chương trình demo này hướng tới:

- Giới thiệu cách sử dụng thư viện `robfig/cron` để tạo và quản lý các tác vụ định kỳ trong Go.
- Tổ chức code theo kiến trúc rõ ràng, tách biệt: khởi tạo, định nghĩa tác vụ, đăng ký tác vụ.
- Kết hợp logger (`zap`) để theo dõi và ghi lại trạng thái thực thi của từng tác vụ.

### Vấn đề thực tế mà chương trình giải quyết

Một lỗi phổ biến trong thực tế là khi một cronjob đang chạy nhưng chưa hoàn thành, vòng lặp tiếp theo đã kích hoạt và bắt đầu xử lý lại cùng một tập dữ liệu. Hậu quả:

- Email bị gửi hai lần đến cùng một người dùng.
- Dữ liệu bị xử lý trùng lặp (duplicate).
- Hệ thống bị spam, bị ban hoặc bị đánh dấu là nguồn gửi thư rác.

Thư viện `robfig/cron` giải quyết vấn đề này bằng cơ chế kiểm tra: nếu một job đang chạy, job tiếp theo sẽ không được khởi động cho đến khi job hiện tại hoàn thành.

---

## 2. Các thư viện sử dụng

### 2.1. `robfig/cron` (v3)

|Thuộc tính|Chi tiết|
|---|---|
|Tên đầy đủ|`github.com/robfig/cron/v3`|
|Tác giả|Rob Figueiredo|
|Nguồn|https://github.com/robfig/cron|
|Chức năng|Lập lịch và thực thi các tác vụ định kỳ trong Go|

**Chức năng chính:**

- Hỗ trợ cú pháp cron đầy đủ, bao gồm cả trường `giây` (mặc định cron truyền thống chỉ có 5 trường: phút, giờ, ngày, tháng, thứ).
- Cho phép thêm nhiều job vào cùng một bộ lập lịch.
- Phát hiện và ngăn chặn job chạy trùng lặp.
- An toàn với goroutine (goroutine-safe).

**Lý do chọn thư viện này:**

- Ổn định, được sử dụng rộng rãi trong cộng đồng Go.
- Tích hợp nhanh gọn, không cần cấu hình phức tạp.
- Hỗ trợ lập lịch đến đơn vị giây, phù hợp với các tác vụ yêu cầu độ chính xác cao.

---

### 2.2. `go.uber.org/zap`

|Thuộc tính|Chi tiết|
|---|---|
|Tên đầy đủ|`go.uber.org/zap`|
|Tác giả|Uber Technologies|
|Nguồn|https://github.com/uber-go/zap|
|Chức năng|Ghi log hiệu năng cao (high-performance structured logging)|

**Chức năng chính:**

- Ghi log theo cấu trúc (structured logging): mỗi log entry có thể đính kèm các trường dữ liệu bổ sung.
- Hiệu năng vượt trội so với `fmt.Println` hoặc `log` chuẩn của Go.
- Hỗ trợ nhiều mức độ log: `Debug`, `Info`, `Warn`, `Error`, `Fatal`.

**Lý do chọn thư viện này:**

- Phù hợp với môi trường production, nơi cần theo dõi trạng thái của hàng nghìn tác vụ chạy định kỳ.
- Cho phép ghi lại thông tin lỗi chi tiết khi một job thất bại.

---

## 3. Cài đặt môi trường và thư viện

### 3.1. Yêu cầu

- Go phiên bản 1.18 trở lên (khuyến nghị 1.21+).
- Kết nối Internet để tải thư viện.

### 3.2. Khởi tạo module Go

```bash
mkdir scheduler-cron-api
cd scheduler-cron-api
go mod init github.com/GiaBao0510/scheduler-cron-api
```

Lệnh `go mod init` tạo ra file `go.mod`, đây là file quản lý phụ thuộc của dự án Go. Tên module (`github.com/GiaBao0510/scheduler-cron-api`) được dùng làm tiền tố khi import các package nội bộ.

### 3.3. Cài đặt thư viện

```bash
go get github.com/robfig/cron/v3
go get go.uber.org/zap
```

Sau khi chạy, Go sẽ tự động cập nhật file `go.mod` và `go.sum`.

### 3.4. Tải và đồng bộ tất cả thư viện

```bash
go mod tidy
```

Lệnh này dọn dẹp và đảm bảo tất cả thư viện được khai báo trong code đều có mặt trong `go.mod`.

---

## 4. Cấu trúc thư mục dự án

```
problem_solving/
└── Crontab/
    ├── main.go                  # Điểm vào của chương trình
    ├── global/
    │   └── global.go            # Biến toàn cục dùng chung toàn dự án
    ├── initialize/
    │   ├── cronjob.go           # Hàm khởi tạo đối tượng cron
    │   └── logger.go            # Hàm khởi tạo logger
    ├── api/
    │   └── api.go               # Định nghĩa các tác vụ định kỳ
    └── resgistry/
        └── handleApi.go         # Đăng ký các tác vụ vào scheduler
```

**Nguyên tắc tổ chức:**

- `global`: Chứa các biến dùng chung. Tránh khai báo biến toàn cục rải rác trong nhiều file.
- `initialize`: Chứa logic khởi tạo các thành phần hệ thống (cron, logger). Tách riêng để dễ thay thế hoặc mở rộng.
- `api`: Chứa định nghĩa nghiệp vụ. Mỗi hàm đại diện cho một tác vụ cụ thể.
- `resgistry`: Là nơi kết nối mọi thứ lại, đăng ký các tác vụ vào scheduler.

---

## 5. Giải thích chi tiết từng file

### 5.1. `global/global.go` - Biến toàn cục

```go
package global

import (
    "github.com/robfig/cron/v3"
    "go.uber.org/zap"
)

var (
    GO_CRON   *cron.Cron
    GO_LOGGER *zap.Logger
)
```

**Mục đích:**  
File này khai báo hai biến toàn cục:

- `GO_CRON`: Con trỏ đến đối tượng `cron.Cron`. Đây là bộ lập lịch trung tâm, nơi tất cả các job được đăng ký.
- `GO_LOGGER`: Con trỏ đến đối tượng logger của `zap`. Được dùng để ghi log từ bất kỳ đâu trong chương trình.

**Tại sao dùng con trỏ (`*cron.Cron`, `*zap.Logger`)?**  
Vì `cron.Cron` và `zap.Logger` là các struct lớn và có trạng thái nội tại. Sử dụng con trỏ đảm bảo rằng toàn bộ chương trình đều tham chiếu đến **cùng một đối tượng**, thay vì mỗi nơi tạo ra một bản sao riêng.

---

### 5.2. `initialize/logger.go` - Khởi tạo Logger

```go
package initialize

import "go.uber.org/zap"

func InitLogger() *zap.Logger {
    logger, _ := zap.NewProduction()
    return logger
}
```

**Mục đích:**  
Tạo ra một đối tượng logger theo cấu hình production. `zap.NewProduction()` trả về logger với:

- Output dạng JSON (dễ đọc bằng công cụ phân tích log).
- Mức log mặc định là `Info` trở lên.
- Tự động ghi thêm thông tin: timestamp, tên file, số dòng.

**Lưu ý về `_` (blank identifier):**  
`zap.NewProduction()` trả về hai giá trị: `(*zap.Logger, error)`. Ký hiệu `_` nghĩa là bỏ qua giá trị lỗi. Trong code production thực tế, bạn nên xử lý lỗi này thay vì bỏ qua.

---

### 5.3. `initialize/cronjob.go` - Khởi tạo Cron

File này chứa hàm khởi tạo đối tượng `cron.Cron`:

```go
package resgistry

import (
    "github.com/GiaBao0510/scheduler-cron-api/api"
    "github.com/GiaBao0510/scheduler-cron-api/global"
)

func RegisApiRunCronjob() {
    api.SendEmailForVipUsersEvery3Seconds(global.GO_CRON)
    api.GetInformationForVipUsersEvery5Seconds(global.GO_CRON)

    global.GO_CRON.Start()
}
```

Hàm `initialize.InitCron()` (được gọi trong `main.go`) cần khởi tạo `cron.Cron` với tùy chọn hỗ trợ trường giây:

```go
// Cách khởi tạo đúng để hỗ trợ trường giây
c := cron.New(cron.WithSeconds())
```

**Tại sao cần `cron.WithSeconds()`?**  
Mặc định, `robfig/cron` sử dụng 5 trường (phút, giờ, ngày, tháng, thứ), giống cron truyền thống trên Linux. Khi thêm `cron.WithSeconds()`, cú pháp mở rộng thành 6 trường: `giây phút giờ ngày tháng thứ`, cho phép lập lịch chính xác đến giây.

---

### 5.4. `api/api.go` - Định nghĩa các tác vụ định kỳ

Đây là file quan trọng nhất, định nghĩa hai tác vụ:

#### Tác vụ 1: Gửi email mỗi 3 giây

```go
func SendEmailForVipUsersEvery3Seconds(cr *cron.Cron) {
    fmt.Println("... Gửi email cho mỗi người dùng Vip mỗi 3 giây ...")

    _, err := cr.AddFunc("*/3 * * * * *", func() {
        log.Println("Gửi email cho người dùng Vip")
    })

    if err != nil {
        global.GO_LOGGER.Error("Lỗi khi thêm cron job: ", zap.Error(err))
    }
}
```

**Giải thích cú pháp cron `"*/3 * * * * *"`:**

```
*/3  *  *  *  *  *
 |   |  |  |  |  |
 |   |  |  |  |  +-- Thứ trong tuần (0-6, 0 = Chủ nhật)
 |   |  |  |  +----- Tháng (1-12)
 |   |  |  +-------- Ngày trong tháng (1-31)
 |   |  +----------- Giờ (0-23)
 |   +-------------- Phút (0-59)
 +------------------ Giây (0-59), */3 nghĩa là mỗi 3 giây
```

`*/3` là cú pháp "step value": thực thi tại mọi giá trị chia hết cho 3 trong khoảng (0, 3, 6, 9, ..., 57).

#### Tác vụ 2: Lấy thông tin từ API ngoài mỗi 5 giây

```go
func GetInformationForVipUsersEvery5Seconds(cr *cron.Cron) {
    fmt.Println("... Lấy thông tin cho mỗi người dùng Vip mỗi 5 giây ...")

    _, err := cr.AddFunc("*/5 * * * * *", func() {
        log.Println("Lấy thông tin cho người dùng Vip")
    })

    rs, err := http.Get("https://httpbin.org/get")
    if err != nil {
        global.GO_LOGGER.Error("Lỗi khi gửi yêu cầu HTTP: ", zap.Error(err))
        return
    }

    body, err := io.ReadAll(rs.Body)
    if err != nil {
        global.GO_LOGGER.Error("Lỗi khi đọc phản hồi HTTP: ", zap.Error(err))
        return
    }

    log.Println("Get information for Vip users:", string(body))

    if err != nil {
        global.GO_LOGGER.Error("Lỗi khi thêm cron job: ", zap.Error(err))
    }
}
```

---

### 5.5. `resgistry/handleApi.go` - Đăng ký tác vụ

```go
func RegisApiRunCronjob() {
    api.SendEmailForVipUsersEvery3Seconds(global.GO_CRON)
    api.GetInformationForVipUsersEvery5Seconds(global.GO_CRON)

    global.GO_CRON.Start()
}
```

**Vai trò:**  
File này là nơi "lắp ghép" tất cả các tác vụ đã định nghĩa ở `api` vào bộ lập lịch `GO_CRON`. Sau khi đăng ký xong, lệnh `global.GO_CRON.Start()` kích hoạt toàn bộ scheduler.

**Lưu ý:** `Start()` chạy cron trong một goroutine riêng, không chặn luồng chính. Đây là lý do tại sao `main.go` cần `select{}` ở cuối.

---

### 5.6. `main.go` - Điểm vào chương trình

```go
func main() {
    // Bước 1: Khởi tạo Logger
    global.GO_LOGGER = initialize.InitLogger()

    // Bước 2: Khởi tạo cron job
    global.GO_CRON = initialize.InitCron()

    // Bước 3: Đăng ký và chạy cron job
    resgistry.RegisApiRunCronjob()

    select{} // Giữ cho ứng dụng chạy mãi mãi
}
```

**Thứ tự khởi tạo quan trọng:**

1. Logger phải được khởi tạo trước, vì các bước sau có thể cần ghi log lỗi.
2. Cron scheduler phải được khởi tạo trước khi đăng ký tác vụ.
3. Đăng ký tác vụ và gọi `Start()` sau cùng.

---

## 6. Luồng hoạt động tổng thể

```
main()
  |
  +-- InitLogger()       --> GO_LOGGER (zap.Logger)
  |
  +-- InitCron()         --> GO_CRON (cron.Cron với WithSeconds)
  |
  +-- RegisApiRunCronjob()
        |
        +-- SendEmailForVipUsersEvery3Seconds(GO_CRON)
        |     --> cr.AddFunc("*/3 * * * * *", handler)
        |
        +-- GetInformationForVipUsersEvery5Seconds(GO_CRON)
        |     --> cr.AddFunc("*/5 * * * * *", handler)
        |
        +-- GO_CRON.Start()   --> Chạy scheduler trong goroutine nền
  |
  +-- select{}           --> Giữ chương trình chạy vô hạn
```

---

## 7. Các đoạn code cần chú ý đặc biệt

### 7.1. Lỗi thiết kế trong `GetInformationForVipUsersEvery5Seconds`

Đây là đoạn code có vấn đề thiết kế quan trọng cần nhận ra:

```go
func GetInformationForVipUsersEvery5Seconds(cr *cron.Cron) {
    // Dòng này chạy 1 lần khi đăng ký job (DUNG)
    fmt.Println("... Lấy thông tin cho mỗi người dùng Vip mỗi 5 giây ...")

    // Đăng ký job, handler bên trong sẽ chạy mỗi 5 giây
    _, err := cr.AddFunc("*/5 * * * * *", func() {
        log.Println("Lấy thông tin cho người dùng Vip")
        // Chú ý: lệnh gọi HTTP ở NGOÀI closure này, không phải bên trong!
    })

    // Phần code bên dưới chạy NGAY LẬP TỨC khi hàm được gọi,
    // KHÔNG phải mỗi 5 giây. Đây là lỗi thiết kế trong code demo.
    rs, err := http.Get("https://httpbin.org/get")
    // ...
}
```

**Vấn đề:** Lệnh `http.Get(...)` và xử lý response đặt bên ngoài `func(){}` (closure), nên nó chỉ chạy một lần khi `GetInformationForVipUsersEvery5Seconds` được gọi lúc khởi động, không phải mỗi 5 giây.

**Cách sửa đúng:** Đưa toàn bộ logic vào bên trong closure:

```go
_, err := cr.AddFunc("*/5 * * * * *", func() {
    log.Println("Lấy thông tin cho người dùng Vip")

    rs, err := http.Get("https://httpbin.org/get")
    if err != nil {
        global.GO_LOGGER.Error("Lỗi khi gửi yêu cầu HTTP: ", zap.Error(err))
        return
    }
    defer rs.Body.Close() // Quan trọng: luôn đóng Body sau khi dùng

    body, err := io.ReadAll(rs.Body)
    if err != nil {
        global.GO_LOGGER.Error("Lỗi khi đọc phản hồi HTTP: ", zap.Error(err))
        return
    }

    log.Println("Kết quả:", string(body))
})
```

---

### 7.2. Câu lệnh `select{}`

```go
select{} // Giữ cho ứng dụng chạy mãi mãi
```

**Giải thích:**  
`select{}` là một câu lệnh `select` không có case nào. Trong Go, `select` dùng để chờ trên nhiều channel. Khi không có case nào, goroutine hiện tại sẽ bị chặn vô thời hạn (block forever) mà không tiêu tốn CPU.

Cách này được dùng để giữ cho tiến trình `main` không kết thúc, trong khi cron scheduler vẫn chạy trong nền (background goroutine). Nếu `main` kết thúc, toàn bộ chương trình Go sẽ dừng lại.

**Các cách thay thế phổ biến:**

```go
// Cách 1: Dùng channel, cho phép graceful shutdown
quit := make(chan os.Signal, 1)
signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
<-quit
fmt.Println("Dừng chương trình...")
global.GO_CRON.Stop()

// Cách 2: Dùng WaitGroup (ít phổ biến hơn cho trường hợp này)
var wg sync.WaitGroup
wg.Add(1)
wg.Wait()
```

Cách 1 được khuyến nghị cho production vì cho phép dừng chương trình một cách an toàn (graceful shutdown) khi nhận tín hiệu từ hệ điều hành.

---

### 7.3. Xử lý giá trị trả về của `cr.AddFunc`

```go
_, err := cr.AddFunc("*/3 * * * * *", func() {
    log.Println("Gửi email cho người dùng Vip")
})
```

`cr.AddFunc` trả về hai giá trị: `(cron.EntryID, error)`.

- `cron.EntryID` là một định danh số nguyên của job vừa được thêm vào. Có thể dùng để xóa job sau này bằng `cr.Remove(id)`.
- `_` nghĩa là bỏ qua `EntryID`. Trong demo này không cần dùng đến.

**Ví dụ nếu cần xóa job sau:**

```go
jobID, err := cr.AddFunc("*/3 * * * * *", func() {
    log.Println("Gửi email")
})
if err == nil {
    // Sau 10 giây, xóa job này
    time.AfterFunc(10*time.Second, func() {
        cr.Remove(jobID)
    })
}
```

---

## 8. Giải thích từ khóa và câu lệnh nâng cao

### 8.1. Goroutine và tính đồng thời

Khi `GO_CRON.Start()` được gọi, thư viện `robfig/cron` khởi động một goroutine nền. Goroutine là "luồng nhẹ" (lightweight thread) trong Go, được quản lý bởi runtime của Go, không phải hệ điều hành.

```go
// Minh họa khái niệm goroutine
go func() {
    // Đây là một goroutine chạy song song với goroutine chính
    for {
        time.Sleep(3 * time.Second)
        fmt.Println("Tác vụ định kỳ chạy...")
    }
}()
```

`robfig/cron` về bản chất hoạt động tương tự như vòng lặp trên, nhưng với cú pháp lịch biểu linh hoạt hơn và cơ chế kiểm soát tốt hơn.

---

### 8.2. Closure (Hàm đóng)

Trong Go, hàm ẩn danh `func() { ... }` được truyền vào `cr.AddFunc` là một closure. Closure có thể "bắt" (capture) các biến từ phạm vi bên ngoài nó.

```go
// Ví dụ về closure capture biến
prefix := "VIP"
_, _ = cr.AddFunc("*/3 * * * * *", func() {
    // Biến 'prefix' được capture từ phạm vi bên ngoài
    log.Printf("Gửi email cho nhóm: %s\n", prefix)
})
```

**Cảnh báo với closure trong vòng lặp:**

```go
// Code SAI - tất cả job đều dùng cùng một giá trị cuối của i
for i := 0; i < 3; i++ {
    cr.AddFunc("*/5 * * * * *", func() {
        fmt.Println(i) // Luôn in ra 3 (giá trị sau cùng)
    })
}

// Code ĐÚNG - tạo biến cục bộ trong từng vòng lặp
for i := 0; i < 3; i++ {
    idx := i // Tạo bản sao cục bộ
    cr.AddFunc("*/5 * * * * *", func() {
        fmt.Println(idx) // In ra 0, 1, 2 đúng như kỳ vọng
    })
}
```

---

### 8.3. Con trỏ (Pointer) và tham chiếu

```go
var GO_CRON *cron.Cron
```

Dấu `*` trước kiểu dữ liệu khai báo đây là **con trỏ** (pointer). Con trỏ lưu địa chỉ bộ nhớ của một giá trị, không phải bản sao giá trị đó.

```go
// Không dùng con trỏ: mỗi hàm nhận bản sao, thay đổi không ảnh hưởng bên ngoài
func doSomething(c cron.Cron) { ... }

// Dùng con trỏ: mỗi hàm nhận cùng đối tượng, thay đổi có hiệu lực toàn cục
func doSomething(c *cron.Cron) { ... }
```

---

### 8.4. Cú pháp cron mở rộng với giây

|Biểu thức|Ý nghĩa|
|---|---|
|`* * * * * *`|Mỗi giây|
|`*/3 * * * * *`|Mỗi 3 giây|
|`0 */5 * * * *`|Đầu mỗi khoảng 5 phút (giây = 0)|
|`0 0 9 * * *`|Mỗi ngày lúc 9:00:00 sáng|
|`0 0 9 * * 1`|Mỗi thứ Hai lúc 9:00:00 sáng|
|`0 30 10 1 * *`|Ngày đầu mỗi tháng lúc 10:30:00|

---

### 8.5. `defer rs.Body.Close()`

```go
rs, err := http.Get("https://httpbin.org/get")
if err != nil {
    return
}
defer rs.Body.Close() // Luôn viết ngay sau khi kiểm tra lỗi
```

`defer` đảm bảo rằng `rs.Body.Close()` được gọi **ngay trước khi hàm hiện tại trả về**, dù hàm kết thúc bình thường hay do lỗi. Nếu không đóng `Body`, kết nối TCP sẽ không được giải phóng, dẫn đến rò rỉ kết nối (connection leak) sau nhiều lần chạy.

---

## 9. Lỗi thường gặp và cách phòng tránh

### 9.1. Không dùng `cron.WithSeconds()` khi khởi tạo

**Triệu chứng:** Cú pháp `*/3 * * * * *` (6 trường) gây ra lỗi parse hoặc hành vi không như mong đợi.

**Nguyên nhân:** Mặc định `cron.New()` chỉ chấp nhận 5 trường.

**Cách sửa:**

```go
c := cron.New(cron.WithSeconds()) // Bắt buộc khi dùng cú pháp 6 trường
```

---

### 9.2. Đặt logic ngoài closure

Như đã phân tích ở mục 7.1, nếu đặt code xử lý bên ngoài `func() { ... }`, code đó chỉ chạy một lần khi đăng ký, không chạy theo lịch.

**Nguyên tắc:** Mọi tác vụ cần chạy định kỳ phải nằm bên trong closure được truyền cho `AddFunc`.

---

### 9.3. Không xử lý lỗi từ `AddFunc`

```go
// Thiếu kiểm tra lỗi - nguy hiểm
cr.AddFunc("cú pháp sai", func() { ... })

// Đúng - luôn kiểm tra lỗi
_, err := cr.AddFunc("*/3 * * * * *", func() { ... })
if err != nil {
    log.Fatal("Không thể thêm cronjob: ", err)
}
```

---

### 9.4. Không đóng Body của HTTP response

```go
// Sai: Body không được đóng
rs, _ := http.Get(url)
body, _ := io.ReadAll(rs.Body)

// Đúng: luôn dùng defer để đóng Body
rs, err := http.Get(url)
if err != nil { return }
defer rs.Body.Close()
body, _ := io.ReadAll(rs.Body)
```

---

### 9.5. Job chạy trùng lặp

Khi một job mất nhiều thời gian hơn chu kỳ lập lịch, vòng tiếp theo sẽ kích hoạt trong khi vòng trước chưa xong. Để ngăn điều này, dùng `cron.SkipIfStillRunning`:

```go
c := cron.New(
    cron.WithSeconds(),
    cron.WithChain(
        cron.SkipIfStillRunning(cron.DefaultLogger), // Bỏ qua nếu job đang chạy
    ),
)
```

---

## 10. Tóm tắt và bài tập thực hành

### Tóm tắt kiến thức

|Khái niệm|Mô tả ngắn|
|---|---|
|Cronjob|Tác vụ được lập lịch chạy định kỳ|
|`robfig/cron`|Thư viện Go để quản lý cronjob, hỗ trợ cú pháp 6 trường|
|`zap`|Thư viện ghi log hiệu năng cao của Uber|
|`cron.WithSeconds()`|Tùy chọn bật hỗ trợ trường giây trong biểu thức cron|
|`AddFunc(spec, handler)`|Đăng ký một tác vụ định kỳ với biểu thức cron và hàm xử lý|
|`Start()`|Khởi động scheduler trong goroutine nền|
|`select{}`|Chặn goroutine chính vô hạn để chương trình không thoát|
|Closure|Hàm ẩn danh có thể capture biến từ phạm vi bên ngoài|
|`defer`|Trì hoãn thực thi một lệnh cho đến khi hàm hiện tại kết thúc|

---

_Tài liệu này được biên soạn dựa trên mã nguồn demo của dự án `scheduler-cron-api`. Mọi đoạn code trong tài liệu đều có thể chạy được với Go 1.18 trở lên và các thư viện đã nêu._