# Goroutine và Context

![](https://www.digitalocean.com/api/static-content/v1/images?src=https%3A%2F%2Fcommunity-cdn-digitalocean-com.global.ssl.fastly.net%2Frdy4YmqgHCKE2BTGp6VR96rk&width=828)
## 1. Context Là Gì?

**Context** là một cơ chế trong Go giúp ==quản lý vòng đời của các tác vụ đang chạy, đặc biệt khi làm việc với goroutine==. Context cho phép:

- **Hủy (cancel)** một goroutine từ bên ngoài khi không còn cần thiết
- **Đặt thời hạn (timeout/deadline)** để tự động dừng tác vụ nếu chạy quá lâu
- **Truyền dữ liệu (metadata)** qua chuỗi các hàm và goroutine mà không cần biến toàn cục

Hình dung đơn giản: context giống như một **hợp đồng công việc** được ký giữa người giao việc (goroutine cha) và người nhận việc (goroutine con). Hợp đồng ghi rõ: "Công việc này phải hoàn thành trong 2 giây, nếu không sẽ bị hủy." Cả hai bên đều biết điều kiện này và hành xử theo đó.

Điểm quan trọng cần nắm ngay từ đầu: **context không tự dừng goroutine**. Nó chỉ phát tín hiệu. Goroutine phải chủ động lắng nghe tín hiệu đó qua `ctx.Done()` và tự dừng lại.

---
## 2. Các Loại Context

Package `context` cung cấp bốn hàm tạo context chính. Hiểu rõ từng loại giúp chọn đúng công cụ cho từng tình huống.

### 2.1. `context.Background()` và `context.TODO()`

Đây là hai context gốc — không có cha, không có thời hạn, không bao giờ bị hủy.

```go
ctx := context.Background() // dùng làm context gốc ở entry point (main, handler HTTP...)
ctx := context.TODO()       // dùng tạm khi chưa biết dùng context nào — đánh dấu cần xem lại
```

Mọi context khác đều được tạo ra từ hai context này thông qua các hàm `With...`.

### 2.2. `context.WithCancel` — Hủy Thủ Công

Trả về một context con và một hàm `cancel`. Khi gọi `cancel()`, context bị hủy và `ctx.Done()` được đóng.

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel() // luôn gọi cancel để giải phóng tài nguyên
```

Dùng khi: muốn hủy tác vụ theo quyết định của chương trình (ví dụ: người dùng bấm nút hủy).

### 2.3. `context.WithTimeout` — Hủy Sau Một Khoảng Thời Gian

Tự động hủy context sau một khoảng thời gian tính từ thời điểm tạo.

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()
// context tự hủy sau 2 giây, hoặc sớm hơn nếu cancel() được gọi
```

### 2.4. `context.WithDeadline` — Hủy Tại Một Thời Điểm Cụ Thể

Tự động hủy context tại một thời điểm tuyệt đối trong tương lai.

```go
deadline := time.Now().Add(2 * time.Second)
ctx, cancel := context.WithDeadline(context.Background(), deadline)
defer cancel()
```

`WithTimeout(ctx, d)` thực chất là cú pháp ngắn gọn của `WithDeadline(ctx, time.Now().Add(d))`.

### 2.5. `context.WithValue` — Truyền Dữ Liệu Qua Context

Nhúng một cặp key-value vào context để truyền metadata qua chuỗi hàm.

```go
// Quan trọng: key phải là kiểu tự định nghĩa, không dùng string thô
type contextKey string
const priorityKey contextKey = "priority"

ctx = context.WithValue(ctx, priorityKey, "high")
priority := ctx.Value(priorityKey).(string) // đọc lại giá trị
```

Lý do không dùng string thô làm key: nếu hai package khác nhau đều dùng `"priority"` làm key, chúng sẽ ghi đè lên nhau. Dùng kiểu tự định nghĩa đảm bảo key là duy nhất theo namespace của package.

---
## 3. Tại Sao Luôn Phải `defer cancel()`?

Mỗi context `WithCancel`, `WithTimeout`, `WithDeadline` đều giữ một số tài nguyên nội bộ (goroutine theo dõi timer, entry trong cây context cha-con). Nếu không gọi `cancel()`, các tài nguyên này tồn tại cho đến khi context cha bị hủy — gây **context leak**.

```go
// SAI — context leak nếu hàm kết thúc sớm
func doWork() {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    // quên cancel — tài nguyên bị giữ suốt 5 giây dù hàm đã xong
    fetchData(ctx)
}

// ĐÚNG — defer đảm bảo cancel luôn được gọi dù hàm kết thúc bằng cách nào
func doWork() {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel() // gọi ngay sau khi tạo context
    fetchData(ctx)
}
```

---

## 4. Cơ Chế Hoạt Động: Context Tree

Các context được tổ chức theo cấu trúc cây cha-con. Khi context cha bị hủy, **toàn bộ context con cháu đều bị hủy theo**. Chiều ngược lại không đúng: hủy context con không ảnh hưởng đến context cha.

```
context.Background()
        │
        ├── WithTimeout(5s)          ← hủy ctx này
        │       │
        │       ├── WithCancel       ← tự động bị hủy theo
        │       │
        │       └── WithValue(...)   ← tự động bị hủy theo
        │
        └── WithTimeout(10s)         ← không bị ảnh hưởng
```

Đây là cơ chế cho phép hủy một nhóm goroutine liên quan chỉ bằng một lệnh `cancel()` duy nhất.

---

## 5. Ví Dụ 1 — Context Timeout Với Goroutine

Ví dụ minh họa manager giao việc cho nhân viên với thời hạn 2 giây. Nhân viên tự kiểm tra xem còn thời gian không trước mỗi lần làm việc.

```go
package main

import (
    "context"
    "fmt"
    "time"
)

// Định nghĩa kiểu key riêng để tránh xung đột với package khác
type contextKey string

const priorityKey contextKey = "priority"

func main() {
    // Tạo context với timeout 2 giây
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel() // luôn gọi cancel để tránh context leak

    // Nhúng metadata vào context — dùng key kiểu tự định nghĩa, không dùng string thô
    ctx = context.WithValue(ctx, priorityKey, "high")

    // Giao task cho goroutine con
    go doWork(ctx)

    // Goroutine main chờ đủ 3 giây để quan sát goroutine con bị hủy sau 2 giây
    time.Sleep(3 * time.Second)
    fmt.Println("Main: kết thúc chương trình")
}

func doWork(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            // ctx.Err() cho biết lý do hủy:
            // - context.DeadlineExceeded: hết timeout/deadline
            // - context.Canceled: cancel() được gọi thủ công
            fmt.Println("Goroutine nhận tín hiệu hủy:", ctx.Err())
            return

        default:
            // Đọc metadata từ context — cần type assertion vì Value trả về interface{}
            priority, ok := ctx.Value(priorityKey).(string)
            if !ok {
                priority = "unknown"
            }
            fmt.Printf("Đang làm việc... Độ ưu tiên: %s\n", priority)
            time.Sleep(500 * time.Millisecond)
        }
    }
}
```

Output:

```
Đang làm việc... Độ ưu tiên: high
Đang làm việc... Độ ưu tiên: high
Đang làm việc... Độ ưu tiên: high
Đang làm việc... Độ ưu tiên: high
Goroutine nhận tín hiệu hủy: context deadline exceeded
Main: kết thúc chương trình
```

Luồng thực thi:

```
t=0s:    tạo context timeout 2s, khởi động goroutine doWork
t=0.5s:  doWork in lần 1
t=1.0s:  doWork in lần 2
t=1.5s:  doWork in lần 3
t=2.0s:  doWork in lần 4
t=2.0s:  context hết hạn, ctx.Done() được đóng
t=2.0s:  vòng lặp tiếp theo: select chọn case <-ctx.Done(), goroutine thoát
t=3.0s:  main thoát
```

---

## 6. Ví Dụ 2 — Kết Hợp Context, Channel, Và Select

Ví dụ này minh họa tình huống thực tế hơn: ba tác vụ nấu ăn chạy song song, mỗi tác vụ mất thời gian khác nhau. Chương trình nhận kết quả từ tác vụ nào xong trước, và tự động hủy các tác vụ còn lại khi hết timeout.

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    // Ba channel nhận kết quả từ ba goroutine
    phoCh    := make(chan string, 1) // buffer=1 để goroutine không bị block khi gửi
    pizzaCh  := make(chan string, 1)
    banhCuonCh := make(chan string, 1)

    // Context timeout 1500ms — đủ để phở (1000ms) và bánh cuốn (1499ms) xong,
    // nhưng không đủ cho pizza (2000ms)
    ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
    defer cancel()

    // Khởi động ba goroutine song song
    go cookPho(ctx, phoCh)
    go cookPizza(ctx, pizzaCh)
    go cookBanhCuon(ctx, banhCuonCh)

    // Đếm số tác vụ đã hoàn thành hoặc bị hủy
    completed := 0
    total := 3

    for completed < total {
        select {
        case result := <-phoCh:
            fmt.Println("Nhận được:", result)
            completed++

        case result := <-pizzaCh:
            fmt.Println("Nhận được:", result)
            completed++

        case result := <-banhCuonCh:
            fmt.Println("Nhận được:", result)
            completed++

        case <-ctx.Done():
            // Context hết hạn — các goroutine chưa xong sẽ tự hủy
            fmt.Println("Hết thời gian:", ctx.Err())
            // Gọi cancel() để chắc chắn tất cả goroutine nhận tín hiệu hủy
            cancel()
            return
        }
    }

    fmt.Println("Tất cả tác vụ hoàn thành")
}

// cookPho hoàn thành sau 1000ms — trong thời hạn context
func cookPho(ctx context.Context, ch chan<- string) {
    fmt.Println("Bắt đầu nấu phở...")
    select {
    case <-time.After(1000 * time.Millisecond):
        ch <- "Phở đã xong"
    case <-ctx.Done():
        fmt.Println("Hủy nấu phở:", ctx.Err())
    }
}

// cookPizza hoàn thành sau 2000ms — vượt quá timeout, sẽ bị hủy
func cookPizza(ctx context.Context, ch chan<- string) {
    fmt.Println("Bắt đầu nướng pizza...")
    select {
    case <-time.After(2000 * time.Millisecond):
        ch <- "Pizza đã xong"
    case <-ctx.Done():
        fmt.Println("Hủy nướng pizza:", ctx.Err())
    }
}

// cookBanhCuon hoàn thành sau 1499ms — sát ngưỡng timeout
func cookBanhCuon(ctx context.Context, ch chan<- string) {
    fmt.Println("Bắt đầu nấu bánh cuốn...")
    select {
    case <-time.After(1499 * time.Millisecond):
        ch <- "Bánh cuốn đã xong"
    case <-ctx.Done():
        fmt.Println("Hủy nấu bánh cuốn:", ctx.Err())
    }
}
```

Output (thứ tự phở và bánh cuốn có thể thay đổi tùy lần chạy):

```
Bắt đầu nấu phở...
Bắt đầu nướng pizza...
Bắt đầu nấu bánh cuốn...
Nhận được: Phở đã xong
Nhận được: Bánh cuốn đã xong
Hết thời gian: context deadline exceeded
Hủy nướng pizza: context deadline exceeded
```

---
## 7. Context Propagation — Truyền Context Qua Chuỗi Hàm

Trong thực tế, một request thường đi qua nhiều lớp hàm. Context cần được truyền xuyên suốt để tín hiệu hủy lan xuống đến tận cùng.

```go
// Quy ước: ctx luôn là tham số đầu tiên, tên luôn là ctx
func handleRequest(ctx context.Context, userID int) error {
    // Truyền context xuống lớp service
    user, err := userService.GetUser(ctx, userID)
    if err != nil {
        return err
    }

    // Truyền tiếp xuống lớp repository
    orders, err := orderRepo.GetOrders(ctx, user.ID)
    if err != nil {
        return err
    }

    return nil
}

func (s *UserService) GetUser(ctx context.Context, id int) (*User, error) {
    // Truyền tiếp xuống database layer
    return s.db.QueryRowContext(ctx, "SELECT * FROM users WHERE id=$1", id)
}
```

Nếu bất kỳ lớp nào gọi `cancel()` hoặc context hết hạn, **tất cả các lớp phía dưới** đều nhận được tín hiệu hủy và dừng lại — không cần cơ chế thông báo riêng giữa các lớp.

---

## 8. Context Trong HTTP Request — Ứng Dụng Thực Tế Phổ Biến Nhất

Mỗi HTTP request trong Go tự động có một context gắn liền. Context này bị hủy khi client ngắt kết nối. Đây là use case thực tế quan trọng nhất của context.

```go
package main

import (
    "context"
    "database/sql"
    "fmt"
    "net/http"
    "time"
)

func userHandler(w http.ResponseWriter, r *http.Request) {
    // Lấy context từ request — tự động bị hủy khi client ngắt kết nối
    ctx := r.Context()

    // Thêm timeout cho toàn bộ xử lý request: không quá 3 giây
    ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
    defer cancel()

    // Truyền context xuống các lớp bên dưới
    result, err := fetchUserFromDB(ctx, 42)
    if err != nil {
        if ctx.Err() != nil {
            // Client đã ngắt kết nối hoặc timeout — không cần trả response
            fmt.Println("Request bị hủy:", ctx.Err())
            return
        }
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    fmt.Fprintln(w, result)
}

func fetchUserFromDB(ctx context.Context, userID int) (string, error) {
    // QueryContext tự động dừng query nếu context bị hủy
    // Điều này giải phóng tài nguyên database ngay lập tức thay vì chờ query xong
    row := db.QueryRowContext(ctx, "SELECT name FROM users WHERE id = $1", userID)

    var name string
    if err := row.Scan(&name); err != nil {
        return "", err
    }
    return name, nil
}

var db *sql.DB
```

Lợi ích cụ thể: nếu client ngắt kết nối giữa chừng, `ctx` bị hủy và `QueryRowContext` dừng query đang chạy trên database — tiết kiệm tài nguyên database thay vì để query chạy đến hết dù không ai cần kết quả.

---

## 9. Hủy Nhiều Goroutine Cùng Lúc

`WithCancel` kết hợp với `close(done)` pattern cho phép hủy một nhóm goroutine chỉ bằng một lệnh duy nhất.

```go
package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
    defer wg.Done()
    for {
        select {
        case <-ctx.Done():
            fmt.Printf("Worker %d dừng: %v\n", id, ctx.Err())
            return
        default:
            fmt.Printf("Worker %d đang chạy...\n", id)
            time.Sleep(500 * time.Millisecond)
        }
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())

    var wg sync.WaitGroup
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(ctx, i, &wg)
    }

    // Cho các worker chạy 1.5 giây
    time.Sleep(1500 * time.Millisecond)

    // Một lệnh cancel() hủy tất cả 3 goroutine cùng lúc
    fmt.Println("Main: gửi tín hiệu hủy tất cả worker")
    cancel()

    // Chờ tất cả goroutine thực sự thoát trước khi main kết thúc
    wg.Wait()
    fmt.Println("Main: tất cả worker đã dừng")
}
```

---

## 10. Lưu Ý Quan Trọng

**Luôn truyền context là tham số đầu tiên:**

```go
// Đúng — ctx là tham số đầu tiên theo quy ước Go
func DoSomething(ctx context.Context, arg string) error

// Sai — ctx không phải tham số đầu tiên
func DoSomething(arg string, ctx context.Context) error
```

**Không lưu context vào struct:**

```go
// Sai — context gắn với một request/operation cụ thể, không phải với đối tượng
type Server struct {
    ctx context.Context // không làm vậy
}

// Đúng — truyền context qua tham số mỗi khi gọi phương thức
func (s *Server) HandleRequest(ctx context.Context) error { ... }
```

**Không truyền `nil` context:**

```go
// Sai — nil context gây panic khi các hàm bên dưới gọi ctx.Done() hoặc ctx.Value()
doWork(nil)

// Đúng — nếu không có context phù hợp, dùng context.Background() hoặc context.TODO()
doWork(context.Background())
```

**`ctx.Value()` chỉ dùng cho metadata, không dùng để truyền tham số bắt buộc:**

```go
// Sai — userID là tham số nghiệp vụ, không phải metadata
ctx = context.WithValue(ctx, "userID", 123)
userID := ctx.Value("userID").(int)

// Đúng — truyền userID qua tham số hàm thông thường
func GetOrders(ctx context.Context, userID int) ([]Order, error)
```

**Kiểm tra `ctx.Err()` để phân biệt lý do hủy:**

```go
if err := ctx.Err(); err != nil {
    switch err {
    case context.DeadlineExceeded:
        // hết timeout hoặc deadline
        log.Println("Tác vụ hết thời gian")
    case context.Canceled:
        // cancel() được gọi thủ công
        log.Println("Tác vụ bị hủy chủ động")
    }
}
```

---

## 11. Tổng Kết

|Loại context|Hàm tạo|Khi nào hủy|Dùng khi|
|---|---|---|---|
|Gốc|`Background()`, `TODO()`|Không bao giờ|Entry point, placeholder|
|Hủy thủ công|`WithCancel`|Khi gọi `cancel()`|Hủy theo logic nghiệp vụ|
|Timeout|`WithTimeout`|Sau khoảng thời gian|Giới hạn thời gian xử lý|
|Deadline|`WithDeadline`|Tại thời điểm cụ thể|Deadline tuyệt đối|
|Truyền dữ liệu|`WithValue`|Theo context cha|Truyền metadata (request-id, user-id...)|

Bốn nguyên tắc cốt lõi khi làm việc với context:

- Context truyền tín hiệu, goroutine tự chịu trách nhiệm lắng nghe và dừng.
- Luôn `defer cancel()` ngay sau khi tạo context có cancel function.
- Truyền context qua tham số đầu tiên, không lưu vào struct.
- Dùng kiểu tự định nghĩa làm key cho `WithValue`, không dùng string thô.

> **Phần tiếp theo:** Chúng ta sẽ tìm hiểu về `sync.Mutex` và `sync.RWMutex` — cơ chế đồng bộ hóa khi nhiều goroutine cùng truy cập một vùng dữ liệu dùng chung mà context không phải lựa chọn phù hợp.