Mặc định Goroutine sẽ tự động chạy Parrallelism kết hợp với Concurrent nếu ứng dụng chúng ta có nhiều Concurrent, cũng như sẽ tự động phân phối lượng CPU cho phù hợp để đạt được hiệu năng tốt nhất cho ứng dụng. Tuy nhiên, trong một số trường hợp điều này làm chiếm dụng hết toàn bộ CPU trên hệ thống khiến các chương trình khác bị treo. Vì thế chúng ta cần quản lý số lượng CPU phù hợp cho ứng dụng.

Ví dụ:

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	numCPU := runtime.NumCPU() // Lấy số lượng CPU có sẵn
	fmt.Printf("Số lượng CPU: %d\n", numCPU)

	// Thiết lập số lượng tối đa CPU sử dụng cho chương trình
	runtime.GOMAXPROCS(numCPU)

	//Tạo wait group để chờ các goroutine hoàn thành
	var wg sync.WaitGroup

	for i := 0; i < 10; i++{
		wg.Add(1)
		go heavyTask(&wg)
	}


	wg.Wait()	// Chờ goroutine hoàn thành

	fmt.Println("Tổng thời gian:  ", time.Since(start))

}

// Hàm nyaf thực hiện một tác vụ tính tổng nặng
func heavyTask(wg *sync.WaitGroup){
	defer wg.Done()

	sum := 0
	for i := 0; i < 100e8; i++{
		sum += i
	}
}
```