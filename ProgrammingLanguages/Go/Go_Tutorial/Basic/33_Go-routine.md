# Goroutine  trong Go - Hướng dẫn từ Cơ bản đến Nâng cao

## 📋 Mục lục

1. [Goroutine là gì?](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#1-goroutine-l%C3%A0-g%C3%AC)
2. [Sequential vs Concurrent - So sánh đơn giản](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#2-sequential-vs-concurrent---so-s%C3%A1nh-%C4%91%C6%A1n-gi%E1%BA%A3n)
3. [Concurrency vs Parallelism - Hiểu đúng về hai khái niệm](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#3-concurrency-vs-parallelism---hi%E1%BB%83u-%C4%91%C3%BAng-v%E1%BB%81-hai-kh%C3%A1i-ni%E1%BB%87m)
4. [Goroutine cơ bản - Bắt đầu với ví dụ](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#4-goroutine-c%C6%A1-b%E1%BA%A3n---b%E1%BA%AFt-%C4%91%E1%BA%A7u-v%E1%BB%9Bi-v%C3%AD-d%E1%BB%A5)
5. [Vấn đề đồng bộ hóa - Tại sao cần WaitGroup](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#5-v%E1%BA%A5n-%C4%91%E1%BB%81-%C4%91%E1%BB%93ng-b%E1%BB%99-h%C3%B3a---t%E1%BA%A1i-sao-c%E1%BA%A7n-waitgroup)
6. [sync.WaitGroup - Giải pháp đồng bộ](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#6-syncwaitgroup---gi%E1%BA%A3i-ph%C3%A1p-%C4%91%E1%BB%93ng-b%E1%BB%99)
7. [Channel - Giao tiếp giữa Goroutine](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#7-channel---giao-ti%E1%BA%BFp-gi%E1%BB%AFa-goroutine)
8. [Select - Xử lý nhiều Channel](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#8-select---x%E1%BB%AD-l%C3%BD-nhi%E1%BB%81u-channel)
9. [Patterns thực tế](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#9-patterns-th%E1%BB%B1c-t%E1%BA%BF)
10. [Các lỗi thường gặp và cách tránh](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#10-c%C3%A1c-l%E1%BB%97i-th%C6%B0%E1%BB%9Dng-g%E1%BA%B7p-v%C3%A0-c%C3%A1ch-tr%C3%A1nh)

---

## 1. Goroutine là gì?

**Goroutine** là một "luồng nhẹ" (lightweight thread) trong Go, cho phép chạy nhiều tác vụ đồng thời.

### 🔑 Đặc điểm:

|Đặc điểm|Goroutine|Thread truyền thống|
|---|---|---|
|**Bộ nhớ**|~2KB|~1-2MB|
|**Số lượng**|Hàng triệu|Hàng nghìn|
|**Tạo/Hủy**|Rất nhanh|Chậm|
|**Quản lý**|Go runtime|OS|

### 💡 Cách tạo Goroutine:

```go
// Hàm thông thường
sayHello()

// Goroutine - Thêm "go" phía trước
go sayHello()
```

> **Ghi nhớ:** Chỉ cần thêm từ khóa `go` trước lời gọi hàm!

---

## 2. Sequential vs Concurrent - So sánh đơn giản

### 2.1. Sequential (Tuần tự)

Các tác vụ chạy **lần lượt**, tác vụ sau phải chờ tác vụ trước hoàn thành.

```go
package main

import (
	"fmt"
	"time"
)

func Task(id int) {
	fmt.Printf("Task %d bắt đầu\n", id)
	time.Sleep(1 * time.Second) // Giả lập công việc mất 1 giây
	fmt.Printf("Task %d kết thúc\n", id)
}

func main() {
	start := time.Now()

	// Chạy tuần tự
	Task(1)
	Task(2)
	Task(3)

	fmt.Printf("⏱️  Tổng thời gian: %v\n", time.Since(start))
}
```

**Output:**

```
Task 1 bắt đầu
Task 1 kết thúc
Task 2 bắt đầu
Task 2 kết thúc
Task 3 bắt đầu
Task 3 kết thúc
⏱️  Tổng thời gian: 3s
```

**Minh họa:**

```
Timeline:
0s ────────────────────> 3s
   [Task 1][Task 2][Task 3]
```

### 2.2. Concurrent (Đồng thời)

Các tác vụ chạy **đồng thời** (overlapping), không phải chờ nhau.

```go
package main

import (
	"fmt"
	"time"
)

func Task(id int) {
	fmt.Printf("Task %d bắt đầu\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Task %d kết thúc\n", id)
}

func main() {
	start := time.Now()

	// Chạy đồng thời với goroutine
	go Task(1) // Thêm "go"
	go Task(2) // Thêm "go"
	go Task(3) // Thêm "go"

	time.Sleep(2 * time.Second) // Chờ goroutines hoàn thành
	fmt.Printf("⏱️  Tổng thời gian: %v\n", time.Since(start))
}
```

**Output:**

```
Task 1 bắt đầu
Task 3 bắt đầu
Task 2 bắt đầu
Task 1 kết thúc
Task 2 kết thúc
Task 3 kết thúc
⏱️  Tổng thời gian: 2s
```

**Minh họa:**

```
Timeline:
0s ────────────────────> 1s
   [Task 1]
   [Task 2]
   [Task 3]
   ↑
Chạy cùng lúc!
```

### 📊 So sánh hiệu quả:

```
Sequential: 1s + 1s + 1s = 3s
Concurrent: max(1s, 1s, 1s) = 1s

→ Nhanh gấp 3 lần!
```

---

## 3. Concurrency vs Parallelism - Hiểu đúng về hai khái niệm

> ⚠️ **Chú ý:** Đây là hai khái niệm khác nhau nhưng dễ nhầm lẫn!

### 3.1. Concurrency (Đồng thời)

**Định nghĩa:** Khả năng xử lý nhiều tác vụ **bằng cách chuyển đổi nhanh** giữa chúng.

**Đặc điểm:**

- Chỉ cần **1 CPU core**
- Tác vụ **không chạy cùng thời điểm**, nhưng **tiến triển đan xen**
- CPU chuyển đổi nhanh giữa các tác vụ (context switching)

**Ví dụ thực tế:**

```
Một đầu bếp làm 3 món ăn:
1. Cho gạo vào nồi cơm → Bật lửa → Chuyển sang việc khác
2. Thái rau → Chuyển sang việc khác
3. Ướp thịt → Quay lại kiểm tra cơm
4. Xào rau → Quay lại nướng thịt
...

→ Một người (1 CPU) làm nhiều việc bằng cách chuyển đổi
```

**Minh họa:**

```
1 CPU core:
Time: 0  1  2  3  4  5  6  7  8  9
CPU:  A  B  C  A  B  C  A  B  C  A
      ↓  ↓  ↓
Chuyển đổi nhanh giữa A, B, C
```

### 3.2. Parallelism (Song song)

**Định nghĩa:** Khả năng thực thi nhiều tác vụ **đúng nghĩa cùng một lúc**.

**Đặc điểm:**

- Yêu cầu **nhiều CPU cores**
- Các tác vụ **thực sự chạy đồng thời**
- Không có context switching

**Ví dụ thực tế:**

```
Ba đầu bếp làm 3 món ăn:
Đầu bếp 1: Nấu cơm      [████████████]
Đầu bếp 2: Xào rau      [████████████]
Đầu bếp 3: Nướng thịt   [████████████]

→ Ba người (3 CPUs) làm ba việc cùng lúc
```

**Minh họa:**

```
3 CPU cores:
Time:   0  1  2  3  4  5  6  7  8  9
Core 1: A  A  A  A  A  A  A  A  A  A
Core 2: B  B  B  B  B  B  B  B  B  B
Core 3: C  C  C  C  C  C  C  C  C  C
        ↑
Chạy thật sự đồng thời
```

### 3.3. So sánh trực quan

|Khía cạnh|Concurrency|Parallelism|
|---|---|---|
|**CPU cần**|1 core đủ|Nhiều cores|
|**Cách chạy**|Chuyển đổi nhanh|Chạy đồng thời|
|**Thời điểm**|Không cùng lúc|Cùng thời điểm|
|**Ví dụ cuộc sống**|1 đầu bếp - 3 món|3 đầu bếp - 3 món|
|**Phù hợp**|I/O tasks|CPU-intensive tasks|

### 3.4. Khi nào dùng gì?

**Dùng Concurrency khi:**

- ✅ Tác vụ **chờ đợi nhiều** (I/O bound)
- ✅ Web server xử lý requests
- ✅ Download nhiều files
- ✅ Database queries
- ✅ API calls

**Ví dụ:** Web server có thể xử lý 1000 requests trên 1 CPU vì hầu hết thời gian là **chờ** database, network.

**Dùng Parallelism khi:**

- ✅ Tác vụ **tính toán nặng** (CPU bound)
- ✅ Xử lý video/image
- ✅ Machine learning
- ✅ Data analysis
- ✅ Rendering 3D

**Ví dụ:** Render video cần **tính toán thật sự** trên nhiều cores.

### 3.5. Go hỗ trợ cả hai!

```go
import "runtime"

func main() {
    // Thiết lập số CPU cores để dùng
    runtime.GOMAXPROCS(4) // Dùng 4 cores

    // Tạo nhiều goroutines
    for i := 0; i < 1000; i++ {
        go doWork(i)
    }
    
    // Go runtime tự động:
    // - Phân phối 1000 goroutines lên 4 cores (Parallelism)
    // - Mỗi core chạy nhiều goroutines concurrent
}
```

### 📝 Tóm tắt dễ nhớ:

```
Concurrency = Dealing with many things at once
             (Xử lý nhiều việc cùng lúc)

Parallelism = Doing many things at once
             (Làm nhiều việc cùng lúc)
```

---

## 4. Goroutine cơ bản - Bắt đầu với ví dụ

### 4.1. Ví dụ đơn giản

```go
package main

import (
	"fmt"
	"time"
)

func sayHello(name string) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("Hello %s - lần %d\n", name, i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	// Chạy bình thường
	sayHello("An")
	sayHello("Bình")
}
```

**Output (Sequential):**

```
Hello An - lần 1
Hello An - lần 2
Hello An - lần 3
Hello Bình - lần 1
Hello Bình - lần 2
Hello Bình - lần 3
```

### 4.2. Thêm Goroutine

```go
func main() {
	// Chạy với goroutine
	go sayHello("An")    // Thêm "go"
	go sayHello("Bình")  // Thêm "go"

	time.Sleep(2 * time.Second) // Chờ goroutines
}
```

**Output (Concurrent):**

```
Hello An - lần 1
Hello Bình - lần 1
Hello An - lần 2
Hello Bình - lần 2
Hello An - lần 3
Hello Bình - lần 3
```

> 💡 **Chú ý:** Output có thể khác nhau mỗi lần chạy vì goroutines chạy đồng thời!

---

## 5. Vấn đề đồng bộ hóa - Tại sao cần WaitGroup

### 5.1. Vấn đề 1: Main goroutine kết thúc sớm

```go
package main

import (
	"fmt"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d bắt đầu\n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("Worker %d kết thúc\n", id)
}

func main() {
	go worker(1)
	go worker(2)
	go worker(3)
	
	// main() kết thúc ngay lập tức!
	fmt.Println("Main kết thúc")
}
```

**Output:**

```
Main kết thúc
```

> ❌ **Vấn đề:** Workers không chạy vì main() đã thoát!

### 5.2. Giải pháp tạm thời: `time.Sleep()`

```go
func main() {
	go worker(1)
	go worker(2)
	go worker(3)
	
	time.Sleep(3 * time.Second) // Chờ workers
	fmt.Println("Main kết thúc")
}
```

**Output:**

```
Worker 1 bắt đầu
Worker 2 bắt đầu
Worker 3 bắt đầu
Worker 1 kết thúc
Worker 2 kết thúc
Worker 3 kết thúc
Main kết thúc
```

> ⚠️ **Vấn đề với `time.Sleep()`:**
> 
> - Không biết chính xác workers mất bao lâu
> - Nếu workers nhanh → Lãng phí thời gian
> - Nếu workers chậm → Bị cắt ngang

### 5.3. Cần giải pháp tốt hơn

**Chúng ta cần:**

1. ✅ Đợi **tất cả** goroutines hoàn thành
2. ✅ Không **đoán mò** thời gian
3. ✅ **Chính xác** và **hiệu quả**

> 💡 **Giải pháp:** `sync.WaitGroup`

---

## 6. sync.WaitGroup - Giải pháp đồng bộ

### 6.1. WaitGroup là gì?

**WaitGroup** là một bộ đếm giúp chờ một nhóm goroutines hoàn thành.

**Cơ chế hoạt động:**

```
Bước 1: wg.Add(1)    → Tăng bộ đếm +1 (có 1 goroutine cần chờ)
Bước 2: wg.Done()    → Giảm bộ đếm -1 (1 goroutine đã xong)
Bước 3: wg.Wait()    → Chờ đến khi bộ đếm = 0 (tất cả xong)
```

### 6.2. Trước khi học WaitGroup - Hiểu `defer`

**`defer` là gì?**

Từ khóa `defer` đảm bảo một câu lệnh **luôn chạy cuối cùng** khi function kết thúc.

```go
package main

import "fmt"

func example() {
	defer fmt.Println("=> 3. Kết thúc") // defer - chạy cuối
	fmt.Println("1. Bắt đầu")
	fmt.Println("2. Xử lý")
}

func main() {
	example()
}
```

**Output:**

```
1. Bắt đầu
2. Xử lý
=> 3. Kết thúc  ← Chạy cuối cùng dù khai báo đầu tiên!
```

**Tại sao cần `defer` với WaitGroup?**

```go
func worker(wg *sync.WaitGroup) {
	defer wg.Done() // ✅ Đảm bảo luôn gọi Done()
	
	// Nếu có lỗi ở đây
	// panic("lỗi!")
	
	// Done() vẫn được gọi!
}
```

> 💡 **Lợi ích:** `defer wg.Done()` đảm bảo Done() luôn được gọi, kể cả khi có lỗi!

### 6.3. Cách sử dụng WaitGroup

**Ba bước quan trọng:**

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Bước 2: Giảm bộ đếm khi xong

	fmt.Printf("Worker %d bắt đầu\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Worker %d kết thúc\n", id)
}

func main() {
	var wg sync.WaitGroup // Tạo WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)         // Bước 1: Tăng bộ đếm
		go worker(i, &wg) // Truyền pointer của wg
	}

	wg.Wait() // Bước 3: Chờ bộ đếm về 0
	fmt.Println("Tất cả workers đã xong!")
}
```

**Output:**

```
Worker 1 bắt đầu
Worker 2 bắt đầu
Worker 3 bắt đầu
Worker 1 kết thúc
Worker 2 kết thúc
Worker 3 kết thúc
Tất cả workers đã xong!
```

### 6.4. Luồng hoạt động chi tiết

```
Main Goroutine                Worker Goroutines
      │
      ├─ var wg WaitGroup
      │
      ├─ wg.Add(1) ─────────→ Counter = 1
      ├─ go worker(1, &wg)    Worker 1 đang chạy
      │
      ├─ wg.Add(1) ─────────→ Counter = 2
      ├─ go worker(2, &wg)    Worker 2 đang chạy
      │
      ├─ wg.Add(1) ─────────→ Counter = 3
      ├─ go worker(3, &wg)    Worker 3 đang chạy
      │
      ├─ wg.Wait()
      │    (BLOCKED)
      │    Chờ counter = 0
      │
      │                       Worker 1 Done() → Counter = 2
      │                       Worker 2 Done() → Counter = 1
      │                       Worker 3 Done() → Counter = 0
      │
      │    (UNBLOCKED)
      ├─ Continue
      └─ fmt.Println("Done!")
```

### 6.5. Lỗi thường gặp

**❌ Lỗi 1: Không dùng pointer**

```go
// SAI
func worker(id int, wg sync.WaitGroup) {
	defer wg.Done() // Giảm bộ đếm của BẢN SAO!
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go worker(1, wg) // Truyền bản sao
	wg.Wait()        // Chờ mãi vì bộ đếm gốc không giảm!
}
```

**✅ Đúng: Dùng pointer**

```go
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Giảm bộ đếm GỐC
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go worker(1, &wg) // Truyền pointer
	wg.Wait()
}
```

**❌ Lỗi 2: Quên `defer`**

```go
func worker(wg *sync.WaitGroup) {
	fmt.Println("Bắt đầu")
	panic("Lỗi!") // Crash tại đây
	wg.Done()      // Không bao giờ chạy được!
}
```

**✅ Đúng: Dùng `defer`**

```go
func worker(wg *sync.WaitGroup) {
	defer wg.Done() // Luôn chạy, kể cả khi panic
	fmt.Println("Bắt đầu")
	panic("Lỗi!")
}
```

### 6.6. Ví dụ thực tế: Download nhiều files

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func downloadFile(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("⬇️  Đang tải: %s\n", url)
	time.Sleep(2 * time.Second) // Giả lập download
	fmt.Printf("✅ Hoàn thành: %s\n", url)
}

func main() {
	urls := []string{
		"file1.zip",
		"file2.zip",
		"file3.zip",
		"file4.zip",
		"file5.zip",
	}

	var wg sync.WaitGroup
	start := time.Now()

	for _, url := range urls {
		wg.Add(1)
		go downloadFile(url, &wg)
	}

	wg.Wait()
	fmt.Printf("\n🎉 Tất cả files đã tải xong!\n")
	fmt.Printf("⏱️  Thời gian: %v\n", time.Since(start))
}
```

**Output:**

```
⬇️  Đang tải: file1.zip
⬇️  Đang tải: file2.zip
⬇️  Đang tải: file3.zip
⬇️  Đang tải: file4.zip
⬇️  Đang tải: file5.zip
✅ Hoàn thành: file1.zip
✅ Hoàn thành: file2.zip
✅ Hoàn thành: file3.zip
✅ Hoàn thành: file4.zip
✅ Hoàn thành: file5.zip

🎉 Tất cả files đã tải xong!
⏱️  Thời gian: 2s

(Nếu không dùng goroutine: 10s!)
```

---
