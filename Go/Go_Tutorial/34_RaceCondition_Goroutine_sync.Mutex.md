# Race Condition và Mutex trong Go - Xử lý đồng bộ an toàn

## 📋 Mục lục

1. [Race Condition là gì?](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#1-race-condition-l%C3%A0-g%C3%AC)
2. [Tại sao Race Condition nguy hiểm?](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#2-t%E1%BA%A1i-sao-race-condition-nguy-hi%E1%BB%83m)
3. [Phát hiện Race Condition](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#3-ph%C3%A1t-hi%E1%BB%87n-race-condition)
4. [Mutex - Giải pháp cơ bản](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#4-mutex---gi%E1%BA%A3i-ph%C3%A1p-c%C6%A1-b%E1%BA%A3n)
5. [RWMutex - Tối ưu cho Read nhiều](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#5-rwmutex---t%E1%BB%91i-%C6%B0u-cho-read-nhi%E1%BB%81u)
6. [Atomic Operations](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#6-atomic-operations)
7. [So sánh các giải pháp](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#7-so-s%C3%A1nh-c%C3%A1c-gi%E1%BA%A3i-ph%C3%A1p)
8. [Best Practices](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#8-best-practices)
9. [Các lỗi thường gặp](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#9-c%C3%A1c-l%E1%BB%97i-th%C6%B0%E1%BB%9Dng-g%E1%BA%B7p)

---

## 1. Race Condition là gì?

**Race Condition** (điều kiện tranh chấp) xảy ra khi:

- Nhiều **Goroutines** truy cập **cùng một tài nguyên**
- **Ít nhất 1 goroutine** đang **ghi** (write)
- **Không có** cơ chế **đồng bộ hóa**

### 🎯 Hình ảnh minh họa

![Race Condition](https://coffeebytes.dev/en/go/go-race-conditions-on-goroutines-and-mutexes/images/race-conditions-go.png)

### 📊 Ví dụ thực tế

**Tương tự:**

```
2 người cùng rút tiền từ 1 tài khoản:
- Số dư ban đầu: 1000$
- Người A rút 600$
- Người B rút 500$

Race Condition:
1. A đọc: 1000$ → OK
2. B đọc: 1000$ → OK
3. A ghi: 1000 - 600 = 400$
4. B ghi: 1000 - 500 = 500$

Kết quả: 500$ (sai!)
Đúng phải là: -100$ hoặc 400$ hoặc 500$
```

### 💡 Ví dụ code - Race Condition

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0
	var wg sync.WaitGroup

	// 1000 goroutines cùng tăng counter
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			counter++ // Race condition ở đây!
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
}
```

**Output (chạy nhiều lần):**

```bash
$ go run .
Counter: 986

$ go run .
Counter: 988

$ go run .
Counter: 990

$ go run .
Counter: 979
```

> ❌ **Vấn đề:** Mỗi lần chạy ra kết quả khác nhau, không bao giờ đúng 1000!

### 🔍 Tại sao lại sai?

**Phân tích `counter++`:**

```go
counter++ 
// Thực chất gồm 3 bước:
// 1. Đọc giá trị counter hiện tại
// 2. Tăng lên 1
// 3. Ghi lại vào counter
```

**Timeline Race Condition:**

```
Time    Goroutine A          Goroutine B          Counter
─────────────────────────────────────────────────────────
t1      Đọc counter = 5                           5
t2                           Đọc counter = 5       5
t3      Tính 5 + 1 = 6                           5
t4                           Tính 5 + 1 = 6       5
t5      Ghi counter = 6                           6
t6                           Ghi counter = 6       6
─────────────────────────────────────────────────────────
Kết quả: 6 (sai! Phải là 7)
```

---

## 2. Tại sao Race Condition nguy hiểm?

### 2.1. Kết quả không xác định

```go
// Chạy lần 1: 986
// Chạy lần 2: 991
// Chạy lần 3: 978
// → Không thể tin tưởng được!
```

### 2.2. Lỗi khó phát hiện

```go
// Có thể chạy đúng 99 lần
// Lần thứ 100 mới lỗi
// → Rất khó debug!
```

### 2.3. Hậu quả nghiêm trọng

**Ví dụ thực tế:**

```go
// Banking System
balance := 1000
go withdraw(500)  // Goroutine A
go withdraw(600)  // Goroutine B
// → Có thể rút quá số dư!

// E-commerce
inventory := 1
go purchase()  // Goroutine A
go purchase()  // Goroutine B
// → Bán quá hàng tồn kho!

// Ticket booking
seats := 1
go book()  // Goroutine A
go book()  // Goroutine B
// → Bán trùng ghế!
```

### 2.4. Ví dụ cụ thể: Hệ thống booking

```go
package main

import (
	"fmt"
	"sync"
)

var availableSeats = 10

func bookSeat(customer string, wg *sync.WaitGroup) {
	defer wg.Done()

	if availableSeats > 0 {
		// Giả lập delay
		// time.Sleep(10 * time.Millisecond)
		
		availableSeats--
		fmt.Printf("✅ %s đặt ghế thành công. Còn %d ghế\n", 
			customer, availableSeats)
	} else {
		fmt.Printf("❌ %s không đặt được ghế\n", customer)
	}
}

func main() {
	var wg sync.WaitGroup

	// 20 khách cùng đặt 10 ghế
	for i := 1; i <= 20; i++ {
		wg.Add(1)
		customerName := fmt.Sprintf("Customer-%d", i)
		go bookSeat(customerName, &wg)
	}

	wg.Wait()
	fmt.Printf("\n🎫 Tổng ghế còn lại: %d\n", availableSeats)
}
```

**Output (sai!):**

```
✅ Customer-1 đặt ghế thành công. Còn 9 ghế
✅ Customer-2 đặt ghế thành công. Còn 8 ghế
...
✅ Customer-15 đặt ghế thành công. Còn -5 ghế  ← SAI!

🎫 Tổng ghế còn lại: -5  ← Âm?!
```

---

## 3. Phát hiện Race Condition

### 3.1. Race Detector - Công cụ của Go

Go có công cụ **built-in** để phát hiện race condition!

**Cách dùng:**

```bash
go run -race main.go
go build -race
go test -race
```

**Ví dụ:**

```go
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
			counter++
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
}
```

**Chạy với -race:**

```bash
$ go run -race main.go
==================
WARNING: DATA RACE
Write at 0x00c000018090 by goroutine 7:
  main.main.func1()
      /path/to/main.go:14 +0x3e

Previous read at 0x00c000018090 by goroutine 6:
  main.main.func1()
      /path/to/main.go:14 +0x2a
==================
Counter: 987
Found 1 data race(s)
```

> ✅ **Luôn dùng `-race` khi test!**

### 3.2. Các dấu hiệu Race Condition

**Dấu hiệu 1: Kết quả không nhất quán**

```go
// Chạy 10 lần, 10 kết quả khác nhau
```

**Dấu hiệu 2: Nhiều goroutines truy cập cùng biến**

```go
// Nhiều goroutines đọc/ghi cùng biến
// Không có lock/unlock
```

**Dấu hiệu 3: Test thỉnh thoảng fail**

```go
// Test pass 99 lần
// Lần thứ 100 fail không rõ lý do
```

---

## 4. Mutex - Giải pháp cơ bản

### 4.1. Mutex là gì?

**Mutex** (Mutual Exclusion - Loại trừ lẫn nhau) là cơ chế đảm bảo:

- Chỉ **1 goroutine** tại một thời điểm có thể truy cập tài nguyên
- Các goroutines khác phải **đợi**

### 🎯 Hình ảnh minh họa

![Mutex Lock](https://coffeebytes.dev/en/go/go-race-conditions-on-goroutines-and-mutexes/images/mutex-lock-en-go.png)

### 4.2. Cách hoạt động

```
Goroutine A                 Mutex                 Goroutine B
     │                        │                        │
     ├─── Lock() ────────────→│                        │
     │                    Khóa lại                      │
     │                        │                        │
     │  Truy cập counter      │                        │
     │                        │    ┌─── Lock() ────────┤
     │                        │    │              Phải đợi!
     │                        │    │                   │
     ├─── Unlock() ──────────→│    │                   │
     │                    Mở khóa  │                   │
     │                        │←───┘                   │
     │                        │  Được vào              │
     │                        │                        │
```

### 4.3. Ví dụ cơ bản với Mutex

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0
	var wg sync.WaitGroup
	var mu sync.Mutex // Tạo Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()   // Khóa
			counter++   // Critical section
			mu.Unlock() // Mở khóa
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
}
```

**Output (luôn đúng):**

```bash
$ go run .
Counter: 1000

$ go run .
Counter: 1000

$ go run .
Counter: 1000
```

> ✅ **Kết quả:** Luôn đúng 1000!

### 4.4. Mutex với defer

**✅ Best Practice:**

```go
func increment(counter *int, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock() // Đảm bảo luôn unlock, kể cả khi panic
	
	*counter++
	// Nếu panic ở đây, defer vẫn chạy unlock()
}
```

### 4.5. Ví dụ thực tế: Bank Account

```go
package main

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	balance int
	mu      sync.Mutex
}

func (acc *BankAccount) Deposit(amount int) {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	
	acc.balance += amount
	fmt.Printf("💰 Gửi %d. Số dư: %d\n", amount, acc.balance)
}

func (acc *BankAccount) Withdraw(amount int) bool {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	
	if acc.balance >= amount {
		acc.balance -= amount
		fmt.Printf("💸 Rút %d. Số dư: %d\n", amount, acc.balance)
		return true
	}
	
	fmt.Printf("❌ Không đủ tiền rút %d. Số dư: %d\n", amount, acc.balance)
	return false
}

func (acc *BankAccount) GetBalance() int {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	
	return acc.balance
}

func main() {
	account := &BankAccount{balance: 1000}
	var wg sync.WaitGroup

	// 5 người cùng gửi tiền
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			account.Deposit(100)
		}(i)
	}

	// 3 người cùng rút tiền
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			account.Withdraw(200)
		}(i)
	}

	wg.Wait()
	fmt.Printf("\n✅ Số dư cuối: %d\n", account.GetBalance())
}
```

**Output:**

```
💰 Gửi 100. Số dư: 1100
💰 Gửi 100. Số dư: 1200
💸 Rút 200. Số dư: 1000
💰 Gửi 100. Số dư: 1100
💸 Rút 200. Số dư: 900
💰 Gửi 100. Số dư: 1000
💸 Rút 200. Số dư: 800
💰 Gửi 100. Số dư: 900

✅ Số dư cuối: 900
```

---

## 5. RWMutex - Tối ưu cho Read nhiều

### 5.1. Vấn đề với Mutex thông thường

```go
// Với sync.Mutex
// Read và Write đều phải chờ nhau
// → Chậm nếu có nhiều reads!

mu.Lock()
data := myData // Chỉ đọc thôi, nhưng vẫn phải lock!
mu.Unlock()
```

### 5.2. RWMutex - Read-Write Mutex

**RWMutex** cho phép:

- **Nhiều goroutines** đọc **cùng lúc**
- **Chỉ 1 goroutine** ghi, và **block tất cả** reads/writes khác

```go
var mu sync.RWMutex

// Nhiều readers cùng lúc
mu.RLock()   // Read lock
data := myData
mu.RUnlock()

// Chỉ 1 writer
mu.Lock()    // Write lock (độc quyền)
myData = newData
mu.Unlock()
```

### 5.3. Ví dụ: Cache

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]string),
	}
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock() // Read lock
	defer c.mu.RUnlock()
	
	value, ok := c.data[key]
	return value, ok
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock() // Write lock
	defer c.mu.Unlock()
	
	c.data[key] = value
}

func main() {
	cache := NewCache()
	var wg sync.WaitGroup

	// 1 writer
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			key := fmt.Sprintf("key-%d", i)
			cache.Set(key, fmt.Sprintf("value-%d", i))
			fmt.Printf("📝 Set: %s\n", key)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// 10 readers
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(50 * time.Millisecond)
			for j := 1; j <= 3; j++ {
				key := fmt.Sprintf("key-%d", j)
				if value, ok := cache.Get(key); ok {
					fmt.Printf("📖 Reader %d: %s = %s\n", id, key, value)
				}
			}
		}(i)
	}

	wg.Wait()
}
```

### 5.4. Khi nào dùng RWMutex?

**✅ Dùng RWMutex khi:**

- Reads **nhiều hơn** writes (tỷ lệ 10:1, 100:1)
- Critical section nhỏ
- Muốn tối ưu performance

**❌ Không dùng RWMutex khi:**

- Writes nhiều
- Read và write cân bằng
- Critical section lớn (overhead của RWMutex cao hơn Mutex)

---

## 6. Atomic Operations

### 6.1. Khái niệm

**Atomic operations** là các thao tác được đảm bảo:

- **Không thể gián đoạn** (indivisible)
- **Không cần** Mutex
- **Nhanh hơn** Mutex

### 6.2. Package sync/atomic

```go
import "sync/atomic"

// Các hàm atomic
atomic.AddInt64(&counter, 1)     // Tăng
atomic.LoadInt64(&counter)        // Đọc
atomic.StoreInt64(&counter, 100)  // Ghi
atomic.SwapInt64(&counter, 200)   // Hoán đổi
atomic.CompareAndSwapInt64(...)   // So sánh và hoán đổi
```

### 6.3. Ví dụ với Atomic

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64 // Phải dùng int64
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&counter, 1) // Atomic increment
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter) // Luôn 1000
}
```

### 6.4. Atomic vs Mutex

```go
// Với Mutex (chậm hơn)
mu.Lock()
counter++
mu.Unlock()

// Với Atomic (nhanh hơn)
atomic.AddInt64(&counter, 1)
```

### 6.5. Benchmark so sánh

```go
// Mutex:     ~100 ns/op
// RWMutex:   ~80 ns/op (read)
// Atomic:    ~5 ns/op
```

### 6.6. Khi nào dùng Atomic?

**✅ Dùng Atomic cho:**

- Counters đơn giản
- Flags (bool)
- Integers, pointers
- Không có logic phức tạp

**❌ Không dùng Atomic cho:**

- Nhiều biến cùng lúc
- Logic phức tạp
- Structs

---

## 7. So sánh các giải pháp

### 7.1. Bảng so sánh

|Giải pháp|Tốc độ|Sử dụng|Use Case|
|---|---|---|---|
|**Mutex**|Trung bình|Dễ|Mọi trường hợp|
|**RWMutex**|Nhanh (nhiều reads)|Trung bình|Read >> Write|
|**Atomic**|Rất nhanh|Khó|Counter, flag đơn giản|
|**Channel**|Chậm|Dễ|Communication|

### 7.2. Decision Tree

```
Cần đồng bộ?
    │
    ├─ Counter đơn giản? → Atomic
    │
    ├─ Nhiều reads? → RWMutex
    │
    ├─ Cần giao tiếp? → Channel
    │
    └─ Còn lại → Mutex
```

### 7.3. Ví dụ so sánh code

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func withMutex() {
	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup

	start := time.Now()
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			counter++
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("Mutex: %v\n", time.Since(start))
}

func withAtomic() {
	var counter int64
	var wg sync.WaitGroup

	start := time.Now()
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&counter, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("Atomic: %v\n", time.Since(start))
}

func main() {
	withMutex()  // ~300ms
	withAtomic() // ~50ms
}
```

---

## 8. Best Practices

### 8.1. Luôn dùng defer với Unlock

```go
// ✅ ĐÚNG
func goodExample(mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock() // Luôn unlock, kể cả panic
	
	// Critical section
}

// ❌ SAI
func badExample(mu *sync.Mutex) {
	mu.Lock()
	// Critical section
	mu.Unlock() // Nếu panic ở trên, unlock không chạy!
}
```

### 8.2. Giữ Critical Section nhỏ

```go
// ❌ SAI - Critical section quá lớn
mu.Lock()
doExpensiveWork1() // 1s
doExpensiveWork2() // 2s
counter++
doExpensiveWork3() // 1s
mu.Unlock()

// ✅ ĐÚNG - Chỉ lock khi cần
doExpensiveWork1()
doExpensiveWork2()
mu.Lock()
counter++
mu.Unlock()
doExpensiveWork3()
```

### 8.3. Embed Mutex trong struct

```go
// ✅ ĐÚNG
type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}
```

### 8.4. Không copy Mutex

```go
type Counter struct {
	mu    sync.Mutex
	value int
}

// ❌ SAI - Copy Mutex
func bad(c Counter) { // Copy by value!
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

// ✅ ĐÚNG - Dùng pointer
func good(c *Counter) {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}
```

### 8.5. Avoid Deadlock

```go
// ❌ Deadlock - Lock 2 lần
mu.Lock()
someFunc() // someFunc cũng gọi mu.Lock()!
mu.Unlock()

// ✅ Không lock lồng nhau
```

### 8.6. Document locking strategy

```go
type Cache struct {
	mu   sync.RWMutex
	data map[string]string // protected by mu
}
```

---

## 9. Các lỗi thường gặp

### 9.1. Quên Unlock

**❌ Lỗi:**

```go
func bad() {
	mu.Lock()
	if condition {
		return // Quên unlock!
	}
	mu.Unlock()
}
```

**✅ Sửa:**

```go
func good() {
	mu.Lock()
	defer mu.Unlock() // Luôn unlock
	
	if condition {
		return
	}
}
```

### 9.2. Deadlock - Lock 2 lần

**❌ Lỗi:**

```go
func outer() {
	mu.Lock()
	defer mu.Unlock()
	inner() // Deadlock!
}

func inner() {
	mu.Lock() // Lock lần 2!
	defer mu.Unlock()
	// ...
}
```

**✅ Sửa:**

```go
func outer() {
	mu.Lock()
	defer mu.Unlock()
	innerUnlocked() // Không lock
}

func innerUnlocked() {
	// Không lock, assume đã lock rồi
}
```

### 9.3. Copy Mutex

**❌ Lỗi:**

```go
type Counter struct {
	mu    sync.Mutex
	value int
}

c1 := Counter{}
c2 := c1 // Copy Mutex - sai!
```

**✅ Sửa:**

```go
c1 := &Counter{}
c2 := c1 // Dùng pointer
```

### 9.4. Lock không đúng scope

**❌ Lỗi:**

```go
mu.Lock()
for i := 0; i < 1000; i++ {
	counter++ // Lock cả vòng lặp!
}
mu.Unlock()
```

**✅ Sửa:**

```go
for i := 0; i < 1000; i++ {
	mu.Lock()
	counter++
	mu.Unlock()
}
```

---

## 📚 Tổng kết

### Race Condition

- Xảy ra khi nhiều goroutines truy cập cùng tài nguyên
- Kết quả không xác định
- Dùng `go run -race` để phát hiện

### Mutex

- Đảm bảo chỉ 1 goroutine truy cập tại một thời điểm
- `Lock()` / `Unlock()`
- Luôn dùng `defer mu.Unlock()`

### RWMutex

- Tối ưu cho trường hợp read nhiều
- `RLock()` / `RUnlock()` cho read
- `Lock()` / `Unlock()` cho write

### Atomic

- Nhanh nhất
- Chỉ cho operations đơn giản
- Counter, flag, integer

### Best Practices

1. ✅ Luôn dùng `defer mu.Unlock()`
2. ✅ Giữ critical section nhỏ
3. ✅ Dùng pointer với struct chứa Mutex
4. ✅ Document locking strategy
5. ✅ Test với `-race`

---

## 📚 Tài liệu tham khảo

- [Go Race Detector](https://go.dev/doc/articles/race_detector)
- [sync Package](https://pkg.go.dev/sync)
- [atomic Package](https://pkg.go.dev/sync/atomic)

---
_Race conditions là nguồn gốc của nhiều bugs khó debug. Hãy luôn cẩn thận!_