### 4. **Làm thế nào phát hiện `Data race` trong Go? và cách giải quyết chúng như thế nào?**

---

#### 🔍 Data Race là gì?

**Data race** là một lỗi đồng thời (concurrency bug) xảy ra khi **hai hay nhiều goroutine cùng truy cập vào một vùng nhớ (biến) tại cùng một thời điểm**, và **ít nhất một trong số đó đang thực hiện thao tác ghi (write)** — mà không có cơ chế đồng bộ hóa nào kiểm soát thứ tự truy cập.

> 💡 **Hình dung đơn giản:** Tưởng tượng hai người cùng lúc chỉnh sửa một tài liệu Google Docs mà không có "khóa chỉnh sửa" — kết quả cuối cùng sẽ bị sai hoặc mất dữ liệu.

![](https://www.mathworks.com/products/polyspace/static-analysis-notes/what-data-races-how-avoid-during-software-development/_jcr_content/mainParsys/band_367826542_copy_/mainParsys/columns/1692d66f-2a59-4c32-aed9-8aaa10235f4a/image_copy_copy_copy.adapt.full.medium.gif/1637843648656.gif)

**Hậu quả của Data Race:**

- Kết quả tính toán **không nhất quán** (mỗi lần chạy cho ra kết quả khác nhau).
- Chương trình có thể **bị crash** hoặc **hoạt động sai** một cách khó đoán.
- Lỗi này **rất khó tái hiện** vì nó phụ thuộc vào thời điểm các goroutine chạy.

---

#### ❌ Ví dụ: Data Race xuất hiện ở đâu?

```go
package main

import (
	"fmt"
	"sync"
)

var counter int // Biến dùng chung giữa các goroutine

func increment() {
	counter++ // ⚠️ Đây là nơi xảy ra Data Race!
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
}
```

**Tại sao `counter++` lại gây ra Data Race?**

Dù trông có vẻ đơn giản, câu lệnh `counter++` thực ra được CPU thực hiện qua **3 bước riêng biệt**:

```
1. Đọc giá trị hiện tại của counter  →  READ
2. Cộng thêm 1                        →  ADD
3. Ghi giá trị mới vào counter        →  WRITE
```

Khi có **1000 goroutine** chạy đồng thời, chúng có thể cùng thực hiện bước READ trước khi bất kỳ goroutine nào kịp WRITE → nhiều goroutine cùng đọc **cùng một giá trị cũ** → cùng ghi lại **cùng một kết quả** → các lần tăng bị **mất đi**.

**Kết quả thực tế khi chạy 5 lần (kết quả luôn < 1000 và không nhất quán):**

```bash
Test> go run .
Counter: 964
Test> go run .
Counter: 990
Test> go run .
Counter: 966
Test> go run .
Counter: 967
Test> go run .
Counter: 980
```

---

#### 🛠️ Cách phát hiện Data Race

Go cung cấp sẵn **Race Detector** — một công cụ tích hợp để tự động phát hiện Data Race khi chạy chương trình. Chỉ cần thêm flag `-race`:

```bash
go run -race .
```

Kết quả sẽ cảnh báo chi tiết nơi xảy ra xung đột:

```bash
PS D:\HocTap\CuuAmChanKinh\Go\Go_Tutorial\Test> go run -race .
==================
WARNING: DATA RACE
Read at 0x00014025a7b0 by goroutine 8:
  main.increment()
      D:/HocTap/CuuAmChanKinh/Go/Go_Tutorial/Test/main.go:11 +0x75
  main.main.func1()
      D:/HocTap/CuuAmChanKinh/Go/Go_Tutorial/Test/main.go:21 +0x69

Previous write at 0x00014025a7b0 by goroutine 7:
  main.increment()
      D:/HocTap/CuuAmChanKinh/Go/Go_Tutorial/Test/main.go:11 +0x8d
  main.main.func1()
      D:/HocTap/CuuAmChanKinh/Go/Go_Tutorial/Test/main.go:21 +0x69

Goroutine 8 (running) created at:
  main.main()
      D:/HocTap/CuuAmChanKinh/Go/Go_Tutorial/Test/main.go:19 +0x56

Goroutine 7 (finished) created at:
  main.main()
      D:/HocTap/CuuAmChanKinh/Go/Go_Tutorial/Test/main.go:19 +0x56
==================
==================
WARNING: DATA RACE
Write at 0x00014025a7b0 by goroutine 9:
  main.increment()
      D:/HocTap/CuuAmChanKinh/Go/Go_Tutorial/Test/main.go:11 +0x8d
  main.main.func1()
      D:/HocTap/CuuAmChanKinh/Go/Go_Tutorial/Test/main.go:21 +0x69

Previous write at 0x00014025a7b0 by goroutine 18:
  main.increment()
      D:/HocTap/CuuAmChanKinh/Go/Go_Tutorial/Test/main.go:11 +0x8d
  main.main.func1()
      D:/HocTap/CuuAmChanKinh/Go/Go_Tutorial/Test/main.go:21 +0x69

Goroutine 9 (running) created at:
  main.main()
      D:/HocTap/CuuAmChanKinh/Go/Go_Tutorial/Test/main.go:19 +0x56

Goroutine 18 (finished) created at:
  main.main()
      D:/HocTap/CuuAmChanKinh/Go/Go_Tutorial/Test/main.go:19 +0x56
==================
Counter: 999
Found 2 data race(s)
exit status 66
```

> ⚠️ **Lưu ý:** Race Detector làm chương trình chạy chậm hơn (~2-20x) và tốn nhiều RAM hơn. Chỉ dùng trong môi trường **development/testing**, không dùng trên production.

---

#### ✅ Cách khắc phục Data Race

Go cung cấp nhiều cách để giải quyết Data Race. Dưới đây là **3 cách phổ biến nhất:**

---

##### Cách 1: Dùng `sync.Mutex` (Mutual Exclusion Lock)

`Mutex` hoạt động như một **chiếc khóa** — chỉ cho phép **1 goroutine** vào vùng critical section tại một thời điểm. Các goroutine khác phải **chờ** đến khi khóa được mở.

```go
package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex
var counter int

func increment() {
	mutex.Lock()         // 🔒 Khóa lại — chỉ goroutine này được vào
	defer mutex.Unlock() // 🔓 Tự động mở khóa khi hàm kết thúc
	counter++
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter) // Luôn ra: Counter: 1000
}
```

**Kết quả:** Luôn đúng `Counter: 1000`.

---

##### Cách 2: Dùng `sync/atomic` (Atomic Operations)

Với các thao tác số học đơn giản (cộng, trừ, so sánh...), Go cung cấp package `sync/atomic` — thực hiện thao tác ở mức **CPU instruction** nên **nhanh hơn Mutex**.

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64 // Phải dùng int32 hoặc int64

func increment() {
	atomic.AddInt64(&counter, 1) // ⚛️ Tăng counter một cách an toàn (thread-safe)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter) // Luôn ra: Counter: 1000
}
```

> 💡 **Khi nào dùng atomic thay vì Mutex?**
> 
> - Dùng `atomic` khi chỉ cần thao tác đơn giản trên **một biến số** (int, bool...).
> - Dùng `Mutex` khi cần bảo vệ **một đoạn code phức tạp** hoặc nhiều biến cùng lúc.

---

##### Cách 3: Dùng Channel (Go-idiomatic approach)

Triết lý của Go là: **"Don't communicate by sharing memory; share memory by communicating."** (Đừng chia sẻ bộ nhớ để giao tiếp; hãy giao tiếp để chia sẻ bộ nhớ.)

Channel cho phép các goroutine gửi dữ liệu cho nhau thay vì cùng truy cập một biến:

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 1000) // Buffered channel chứa 1000 giá trị

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- 1 // Gửi giá trị 1 vào channel thay vì chỉnh biến trực tiếp
		}()
	}

	// Đợi tất cả goroutine xong rồi đóng channel
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Goroutine chính tổng hợp kết quả từ channel
	counter := 0
	for v := range ch {
		counter += v
	}

	fmt.Println("Counter:", counter) // Luôn ra: Counter: 1000
}
```

---

#### 📊 Tổng kết — So sánh các cách giải quyết

|Phương pháp|Ưu điểm|Nhược điểm|Khi nào dùng|
|---|---|---|---|
|`sync.Mutex`|Linh hoạt, dễ hiểu|Có thể gây deadlock nếu dùng sai|Bảo vệ nhiều biến hoặc đoạn code phức tạp|
|`sync/atomic`|Nhanh, đơn giản|Chỉ dùng được với kiểu số nguyên|Thao tác đơn giản trên một biến số|
|Channel|Go-idiomatic, tránh shared state|Phức tạp hơn với logic đơn giản|Truyền dữ liệu giữa các goroutine|

> ✅ **Lời khuyên thực tế:** Trong phỏng vấn, hãy đề cập cả 3 cách và giải thích trade-off. Điều này cho thấy bạn hiểu sâu về concurrency trong Go, không chỉ biết "dùng Mutex là xong".