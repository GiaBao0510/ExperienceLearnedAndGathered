# So Sánh Lập Trình Tuần Tự, Đồng Thời và Song Song Trong Go

## 1. Tổng Quan

Thời gian đầu, CPU chỉ có một nhân duy nhất, các ngôn ngữ lập trình khi đó theo mô hình **tuần tự (sequential)**, điển hình là ngôn ngữ C. Ngày nay, với sự phát triển của công nghệ đa xử lý, các mô hình **đồng thời (concurrent)** và **song song (parallel)** ra đời để tận dụng tối đa sức mạnh của CPU nhiều nhân.

Go được thiết kế với khả năng xử lý đồng thời rất hiệu quả thông qua khái niệm **Goroutines** — nhẹ hơn và linh hoạt hơn so với thread truyền thống.

|Lập trình tuần tự|Lập trình song song|
|---|---|
|![](https://zalopay-oss.github.io/go-advanced/images/ch1-5-sequence-programming.png)|![](https://zalopay-oss.github.io/go-advanced/images/ch1-5-parallelprograming.png)|

---

## 2. Ba Mô Hình Xử Lý

### 2.1. Sequential — Tuần Tự

- **Mô tả:** Một lõi CPU xử lý từng tác vụ một theo thứ tự. Task 1 hoàn thành mới chạy Task 2.
- **Đặc điểm:** Đơn giản, dễ lập trình, dễ debug. Nhưng không tận dụng được tài nguyên CPU khi có nhiều tác vụ chờ đợi (ví dụ: chờ I/O, chờ mạng).

![](https://media.beehiiv.com/cdn-cgi/image/fit=scale-down,format=auto,onerror=redirect,quality=80/uploads/asset/file/ac4b1636-2b73-4b14-a51b-d92c0031c096/image.png?t=1720731877)

**Ví dụ thực tế:** Trình duyệt web theo mô hình tuần tự sẽ phải: nghe nhạc xong → mở tài liệu đọc → tải xong mới dừng. Người dùng phải chờ từng việc một.

---

### 2.2. Concurrent — Đồng Thời

> _"Concurrency is about **dealing** with lots of things at once." — Rob Pike_

- **Mô tả:** Khả năng **phân chia và điều phối** nhiều tác vụ trong cùng một khoảng thời gian. Tuy nhiên, **tại một thời điểm cụ thể**, chỉ có **một tác vụ** được CPU xử lý.
- **Đặc điểm:** CPU không chờ một tác vụ hoàn thành mới làm việc khác — nó chia tác vụ lớn thành các tác vụ nhỏ, **xen kẽ** chúng với nhau. Người dùng cảm giác mọi việc xảy ra "đồng thời", nhưng bản chất CPU vẫn chỉ làm một việc tại một thời điểm.

> **Chú ý:** Concurrent **khác** với Parallel! Concurrent là về cách **tổ chức** công việc, không phải thực sự làm nhiều việc cùng lúc.

**Ví dụ thực tế:** Bạn vừa nghe nhạc, vừa đọc tài liệu, vừa tải file — cả ba việc "diễn ra đồng thời" dù CPU chỉ có một nhân.

![](https://zalopay-oss.github.io/go-advanced/images/ch1.6-concurrent-process.png)

**Cơ chế hoạt động — Context Switch:** Khi nhân CPU chuyển từ xử lý tác vụ này sang tác vụ khác, thao tác đó gọi là **context switch**. CPU lưu lại trạng thái của tác vụ đang làm (để sau này tiếp tục) rồi chuyển sang tác vụ tiếp theo. Việc này xảy ra cực nhanh — hàng nghìn lần mỗi giây — nên người dùng không cảm nhận được sự gián đoạn.

---

### 2.3. Parallel — Song Song

> _"Parallelism is about **doing** lots of things at once." — Rob Pike_

- **Mô tả:** Nhiều tác vụ được thực thi **đồng thời thật sự** tại cùng một thời điểm, mỗi tác vụ chạy trên một nhân CPU riêng biệt.
- **Yêu cầu:** Máy tính phải có **nhiều hơn 1 nhân CPU**. Các tác vụ hoàn toàn độc lập nhau.

**Ví dụ thực tế:** Với CPU 3 nhân, nghe nhạc chạy trên nhân 1, đọc tài liệu trên nhân 2, tải file trên nhân 3 — cả ba thực sự xảy ra **cùng một lúc**.

![](https://zalopay-oss.github.io/go-advanced/images/ch1.6-parallelism-process.png) _Mô hình xử lý song song các tác vụ cùng một thời điểm_

Trong thực tế, trên mỗi nhân CPU vẫn xảy ra xử lý đồng thời (concurrent) — miễn là không có hai nhân cùng xử lý một tác vụ:

![](https://zalopay-oss.github.io/go-advanced/images/ch1.6-parallelism-process1.png) _Mô hình song song kết hợp đồng thời_

---

### 2.4. Concurrent + Parallel — Đồng Thời Và Song Song

- **Mô tả:** Kết hợp cả hai mô hình — nhiều nhân CPU xử lý song song, đồng thời mỗi nhân cũng xử lý đồng thời nhiều tác vụ nhỏ.
- **Đây là hình thức mạnh nhất**, tận dụng toàn bộ sức mạnh phần cứng hiện đại.
- Ứng dụng: Hệ thống web server xử lý hàng nghìn request đồng thời, game engine, xử lý video,...

![](https://jenkov.com/images/java-concurrency/concurrency-vs-parallelism-3.png)

---

### 2.5. Bảng So Sánh Tổng Hợp

|                       | Sequential      | Concurrent        | Parallel     | Concurrent + Parallel |
| --------------------- | --------------- | ----------------- | ------------ | --------------------- |
| **Số nhân CPU**       | 1               | 1+                | 2+           | 2+                    |
| **Cùng lúc thật sự?** | ❌               | ❌                 | ✅            | ✅                     |
| **Nhiều tác vụ?**     | ❌               | ✅                 | ✅            | ✅                     |
| **Độ phức tạp**       | Thấp            | Trung bình        | Cao          | Rất cao               |
| **Ví dụ**             | Script đơn giản | Web server 1 nhân | Render video | Web server hiện đại   |
-
> Bạn có thể xem thêm bài diễn thuyết của Rob Pike phân biệt hai mô hình tại [đây](https://blog.golang.org/concurrency-is-not-parallelism).

---

## 3. Nền Tảng: Process và Thread

Trước khi tìm hiểu Goroutine, cần nắm vững hai khái niệm nền tảng của hệ điều hành.

### 3.1. Process — Tiến Trình

**Process** là một chương trình đang được thực thi trong máy tính. Khi bạn mở trình duyệt web, đó là một process. Khi bạn chạy một chương trình Go, hệ điều hành sẽ:

- Cấp cho nó một **vùng nhớ riêng**
- Gán một **PID (Process ID)** để quản lý
- Tạo ít nhất một **luồng chính (main thread)** để chạy chương trình

> Khi main thread ngừng hoạt động → chương trình kết thúc.

### 3.2. Thread — Luồng

**Thread (tiểu trình)** là đơn vị thực thi bên trong một process. Một process có thể có nhiều thread chạy song song, cùng chia sẻ tài nguyên của process đó.

![](https://zalopay-oss.github.io/go-advanced/images/ch1.6-process.png) _Mô hình process và các thread bên trong_

**Đặc điểm bộ nhớ của Thread:**

|Vùng nhớ|Mô tả|Chia sẻ?|
|---|---|---|
|**Stack**|Lưu biến cục bộ, tham số hàm, địa chỉ trả về|❌ Riêng mỗi thread (~1–2 MB cố định)|
|**Heap**|Lưu dữ liệu động (cấp phát bằng `new`, `malloc`)|✅ Chung toàn bộ process|

**Vấn đề khi dùng nhiều thread:**

- **Stack overflow:** Tạo quá nhiều thread → hết bộ nhớ stack.
- **Race condition:** Nhiều thread cùng đọc/ghi một vùng nhớ chung → dữ liệu bị sai.
- **Chi phí cao:** Khởi tạo và chuyển đổi context giữa các thread tốn tài nguyên.

> **Kinh nghiệm thực tế:** Số thread tối ưu thường là `số nhân CPU × 2`. Tạo quá nhiều thread không giúp chương trình nhanh hơn mà còn gây lỗi.

---

## 4. Goroutine — Giải Pháp Của Go

### 4.1. Goroutine Là Gì?

**Goroutine** là đơn vị xử lý đồng thời của Go — nhẹ hơn thread hệ điều hành rất nhiều. Khởi tạo một goroutine chỉ cần từ khóa `go`:

```go
go tenHam() // Chạy tenHam() trong một goroutine mới
```

### 4.2. So Sánh Goroutine và Thread

|Tiêu chí|Thread (OS)|Goroutine (Go)|
|---|---|---|
|**Kích thước stack ban đầu**|~1–2 MB (cố định)|~2–4 KB (động, tăng theo nhu cầu)|
|**Stack tối đa**|~8 MB|~1 GB|
|**Quản lý bởi**|Hệ điều hành|Go Runtime|
|**Chi phí khởi tạo**|Cao|Rất thấp|
|**Số lượng khả thi**|Hàng nghìn|Hàng triệu|
|**Context switch**|Kernel space (chậm)|User space (nhanh)|

![](https://zalopay-oss.github.io/go-advanced/images/ch1.6-compare-thread-goroutine.png)

### 4.3. Cơ Chế Lập Lịch — M:N Scheduler

Go Runtime sử dụng mô hình **M:N** — ghép **M Goroutines** lên **N thread** của hệ điều hành:

```
Goroutines (M)   →   Go Scheduler   →   OS Threads (N)   →   CPU Cores
  (hàng triệu)                           (số ít)              (vài nhân)
```

Biến `runtime.GOMAXPROCS` quy định số lượng OS thread chạy song song (mặc định = số nhân CPU):

```go
import "runtime"

func main() {
    // Đặt số lượng thread tối đa = số nhân CPU
    runtime.GOMAXPROCS(runtime.NumCPU())
}
```

---

## 5. Ví Dụ Goroutine

### Ví dụ 1: Goroutine Cơ Bản

```go
package main

import "fmt"

func main() {
    go fmt.Println("Xin chào goroutine")
    fmt.Println("Xin chào main goroutine")
}
```

_Output:_
```shell
> go run .
Xin chào main goroutine
```

**Vấn đề:** Chương trình trên có thể chỉ in ra một dòng, hoặc in cả hai nhưng không theo thứ tự nhất định. Nguyên nhân: khi `main()` kết thúc, **toàn bộ chương trình dừng lại** — kể cả các goroutine đang chạy dở.

**Giải pháp tạm thời** — dùng `time.Sleep`:

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    go fmt.Println("Xin chào từ goroutine khác")
    fmt.Println("Xin chào từ main goroutine")

    // Chờ 1 giây để goroutine kia có thời gian chạy
    time.Sleep(time.Second)
}
```

_Output:_
```shell
> go run .
Xin chào main goroutine
Xin chào goroutine
```

> **Lưu ý:** Dùng `time.Sleep` để đồng bộ goroutine là cách **không tốt** trong thực tế vì không đảm bảo chắc chắn. Cách đúng là dùng `sync.WaitGroup` hoặc **channel** (sẽ được trình bày ở bài sau).

---

### Ví dụ 2: Goroutine Với Anonymous Function

```go
package main

import (
    "fmt"
    "time"
)

func MyPrintln(id int, delay time.Duration) {
    go func() {
        time.Sleep(delay)
        fmt.Println("Xin chào, tôi là goroutine:", id)
    }()
}

func main() {
    for i := 0; i < 10; i++ {
        MyPrintln(i, 1*time.Second)
    }

    time.Sleep(3 * time.Second)
    fmt.Println("Chương trình kết thúc")
}
```

**Kết quả (thứ tự có thể khác nhau mỗi lần chạy):**

```
Xin chào, tôi là goroutine: 3
Xin chào, tôi là goroutine: 0
Xin chào, tôi là goroutine: 7
...
Chương trình kết thúc
```

**Giải thích:**

- Mỗi lần gọi `MyPrintln`, một goroutine mới được tạo ra chạy trong nền.
- Tất cả goroutine ngủ 1 giây rồi in ra — nhưng thứ tự in **không đảm bảo** vì chúng chạy đồng thời.
- `time.Sleep(3 * time.Second)` ở `main()` để chờ các goroutine hoàn thành.

---
### Ví dụ 3: Goroutine Với `sync.WaitGroup` (Cách Đúng)

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup

    for i := 0; i < 5; i++ {
        wg.Add(1) // Báo WaitGroup có thêm 1 goroutine cần chờ
        go func(id int) {
            defer wg.Done() // Báo goroutine này hoàn thành (Giải phóng goutine đã hoàn thành)
            fmt.Println("Goroutine:", id)
        }(i)
    }

    wg.Wait() // Chờ tất cả goroutine hoàn thành
    fmt.Println("Tất cả goroutine đã xong!")
}
```

**Ưu điểm:** Không cần đoán thời gian chờ, đảm bảo tất cả goroutine hoàn thành trước khi `main()` kết thúc.

---

## 6. Các Vấn Đề Thường Gặp Với Goroutine

### 6.1. Race Condition

Xảy ra khi nhiều goroutine cùng đọc/ghi một biến chia sẻ mà không có cơ chế bảo vệ:

```go
// ❌ Sai — race condition
package main

import (
    "fmt"
    "sync"
)

func main() {
    counter := 0
    var wg sync.WaitGroup

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter++ // Nhiều goroutine cùng ghi → kết quả sai
        }()
    }

    wg.Wait()
    fmt.Println("Counter:", counter) // Không phải 1000!
}
```

**Giải pháp** — dùng `sync.Mutex`:

```go
// ✅ Đúng — dùng Mutex để bảo vệ vùng nhớ chia sẻ
var mu sync.Mutex

go func() {
    defer wg.Done()
    mu.Lock()
    counter++
    mu.Unlock()
}()
```

> **Tip:** Chạy `go run -race main.go` để Go tự động phát hiện race condition.

### 6.2. Goroutine Leak

Goroutine không bao giờ kết thúc → rò rỉ bộ nhớ. Thường xảy ra khi goroutine chờ channel nhưng không ai gửi dữ liệu:

```go
// ❌ Goroutine bị "treo" mãi mãi
go func() {
    val := <-ch // Nếu không ai gửi vào ch, goroutine này sẽ không bao giờ kết thúc
    fmt.Println(val)
}()
```

---

## 7. Tóm Tắt

```
Sequential  →  Một tác vụ tại một thời điểm, theo thứ tự
Concurrent  →  Nhiều tác vụ luân phiên nhau trên 1 nhân CPU (context switch)
Parallel    →  Nhiều tác vụ thực sự cùng lúc trên nhiều nhân CPU
Cả hai      →  Kết hợp: nhiều nhân + luân phiên trong mỗi nhân

Goroutine   →  Nhẹ hơn thread, Go Runtime quản lý, khởi tạo bằng từ khóa "go"
WaitGroup   →  Đồng bộ goroutine đúng cách (thay cho time.Sleep)
Mutex       →  Bảo vệ dữ liệu chia sẻ khỏi race condition
```
