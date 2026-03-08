
- Race Condition xảy ra khi nhiều Goroutine truy cập và sửa đổi cùng một lúc trên một tài nguyên (ví dụ: biến, thuộc tính trong bảng,...) mà không được đồng bộ hóa
- Kết quả có thể không xác định
- Có thể dẫn đến tình trạng duplicate
*Ảnh minh họa:*
![](https://coffeebytes.dev/en/go/go-race-conditions-on-goroutines-and-mutexes/images/race-conditions-go.png)

Ví dụ trường hợp bị Race condition
```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	token := 0
	var wg sync.WaitGroup

	//Cho 1 ngàng go routine cùng cập nhật token lên 1
	for i:= 0; i < 1000; i++{
		wg.Add(1)
		go func(){
			token++
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Token: ", token)
}

```

Kết quả:
```bash
PS D:\HocTap\CuuAmChanKinh\Go\Go_Tutorial\Test> go run .
Token:  986
PS D:\HocTap\CuuAmChanKinh\Go\Go_Tutorial\Test> go run .
Token:  988
PS D:\HocTap\CuuAmChanKinh\Go\Go_Tutorial\Test> go run .
Token:  990
PS D:\HocTap\CuuAmChanKinh\Go\Go_Tutorial\Test> go run .
Token:  979
```
Có thể thấy mỗi lần chạy thì nó trả về kết quả khác nhau. Vì khi các sự kiện xảy ra trên cùng một thao tác gần như là đồng thời, thì Go Scheduler sẽ chọn ngẫu nhiên kết quả

Để giải quyết tình trạng này chúng ta có thể sử dụng MUTEX để giải quyết vấn đề Condition.
**Mutex (Mutual Exclusion)** là một cơ chế đồng bộ hóa trong lập trình đa luồng (mutithreading), dùng để đảm bảo rằng chỉ có một luồng (thread) tại một thời điểm có thể truy cập vào vùng tài nguyên dùng chung (shared resource), chẳng hạn như: biến toàn cục, bộ nhớ hoặc tệp tin

*Ảnh minh họa:*
![](https://coffeebytes.dev/en/go/go-race-conditions-on-goroutines-and-mutexes/images/mutex-lock-en-go.png)

Ví dụ:
```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	token := 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	//Cho 1 ngàng go routine cùng cập nhật token lên 1
	for i:= 0; i < 1000; i++{
		wg.Add(1)
		go func(){
			mu.Lock()
			token++
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Token: ", token)
}
```