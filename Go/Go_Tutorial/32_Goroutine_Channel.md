**Goroutine** là một "luồng nhẹ" (lightweight thread) mà Go tạo ra để chạy một tác vụ nào đó **đồng thời (concurrent) với chương trình chính**.

Thường thì trước giờ chúng ta  thường làm theo kiểu **Sequential Approach (Tuần tự)**. Nghĩa là tác vụ A hoàn thành xong mới đến tác vụ B và tác vụ B xong thì mới đến tác vụ C,...

Còn **Concurrent Approach (Đồng thời)**, ví dụ như là trong thời gian tác vụ A đang thực thi, thì vẫn có thể chạy được tác vụ B, trong tác vụ B, B đang chạy thì vẫn có thể chạy tác vụ C. (Ví dụ đơn giản dễ hiểu ví dụ 1 CPU là 1 người nấu ăn,. Tuy nhiên tại thời điểm này người này có thể vừa nấu ăn, vừa nghe điện thoại vừa rửa rau cùng một lúc).

Còn **Parallelism (song song)**, cái phần này nó hơi dễ bị nhẫm lẫn so với song song. Để lấy ví dụ dễ hiểu có 3 CPU như là 3 người, mỗi người thực hiện một công việc khác nhau. Chẳng hạn như một người nấu ăn, một người nghe điện thoại, một người rửa rau

Như vậy khi chúng ta có những công việc ngắn, nhỏ, nhẹ thì sử dụng **Concurrent Approach**. Còn khi những công việc năng, lâu khó thì nên sử dụng **Parallelism**.

{Mặc dù tôi hiểu về ví dụ Concurrent Approach với Parallelism, nhung về mặc định nghĩa thì tôi cảm thấy hơi thiếu thông tin về 2 phần này (Nếu được bạn hãy bổ sung thêm ví dụ thực tế khi nào áp dụng cái nào) }

---
## **Cách thức hoạt động của Goroutine**

Ví dụ cách thức hoạt động của Sequential
```go
package main

import (
	"fmt"
	"time"
)

//Mỗi task sẽ mất 1 giây để thực hiện
func Task(id int){
	fmt.Printf("Task %d bắt đầu \n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Task %d kết thúc \n", id)
}

func main() {
	start := time.Now()

	for i := 1; i<= 5; i++{
		Task(i)
	}

	time.Sleep(2 * time.Second)
	fmt.Printf("==> Thời gian thực hiện: ", time.Since(start))
}
```

Kết quả: Ta thấy đầu ra nó thực hiện tuần tự
```bash
go run .
Task 1 bắt đầu 
Task 1 kết thúc 
Task 2 bắt đầu
Task 2 kết thúc 
Task 3 bắt đầu
Task 3 kết thúc 
Task 4 bắt đầu
Task 4 kết thúc 
Task 5 bắt đầu
Task 5 kết thúc 
==> Thời gian thực hiện:  7.0117659s
```

Ví dụ cách thức hoạt động của Goroutine
```go
package main

import (
	"fmt"
	"time"
)

//Mỗi task sẽ mất 1 giây để thực hiện
func Task(id int){
	fmt.Printf("Task %d bắt đầu \n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Task %d kết thúc \n", id)
}

func main() {
	start := time.Now()		// Lưu thời gian bắt đầu

	for i := 1; i<= 5; i++{
		go Task(i)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("==> Thời gian thực hiện: ", time.Since(start))	//In tổng thời gian hoàn thành
}
```

Kết quả: như đã thấy thì ta thấy thời gian bắt đầu và kết thúc được phân chia rõ ràng
```bash
go run .
Task 5 bắt đầu
Task 2 bắt đầu
Task 4 bắt đầu
Task 4 bắt đầu
Task 3 bắt đầu
Task 1 bắt đầu
Task 1 kết thúc
Task 2 kết thúc
Task 5 kết thúc
Task 3 kết thúc
Task 4 kết thúc
==> Thời gian thực hiện:  2.0011272s
```

Goroutine chạy đồng thời chung với main (Tức là nó chạy riêng biệt với main). Nó không phụ thuộc vào main

- Nếu hàm main() kết thúc, tất cả Goroutine cũng bị dừng, ngay cả khi chúng chưa hoàn thành
- trong ví dụ trên dùng time.sleep() là cách tạm thời để chờ các task hoàn thành, nhưng chưa phải là giải pháp tốt

---
### **Đồng bộ hóa với sync.Waitgroup trong Goroutine**

Áp dụng Waitgroup để giải quyết vấn đề này.
Đầu tiên tạo `var wg sync.WaitGroup`
Thứ hai là mỗi khi trước hàm nào đó có từ khóa `go` thì tạo đằng trước dòng lệnh `wg.Add(1)` để tạo ra một vùng đếm (là 1) và bên trong hàm có từ khóa `go` thì thêm câu lệnh `defer wg.Done()` để giảm bộ đếm (là 1). Trước khi kết thúc (tức sau tất cả Goroutine) thì thêm câu lệnh ở cuối là `wg.Wait()` để chờ tất cả Goroutine hoàn thành thì mới kết thúc.

--Giải thích:
- **wg.Add(1):** tăng bộ đếm cho mỗi Goroutine
- **wg.Done():** Giảm bộ đếm khi Goroutine hoàn thành
- **wg.Wait():** Chặn luồng chính cho đến khi bộ đếm về 0

--Giải thích thêm về từ khóa `defer` từ khóa này nhằm làm cho câu lệnh nào đó đã xác định từ trước là nó sẽ luôn chạy cuối cùng, mặc dù nằm ở bất kỳ đâu. Ví dụ:
```go
func main() {
	defer fmt.Println("=> Kết thúc")
	fmt.Println("Bắt đầu thực hiện các task...")
	fmt.Println("Running....1")
	fmt.Println("Running....2")
	fmt.Println("Running....3")
	fmt.Println("Running....4")
}
```

Kết quả:
```bash
> go run .
Bắt đầu thực hiện các task...
Running....1
Running....2
Running....3
Running....4
=> Kết thúc
```

Ví dụ quay lại bài toán trên về đòng bộ hóa sync.WaitGroup:

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

//Mỗi task sẽ mất 1 giây để thực hiện
func Task(id int, wg *sync.WaitGroup){

	defer wg.Done()		// Đảm bảo gọi Done() khi công việc hoàn thành, giảm số lượng công việc cần chờ

	fmt.Printf("Task %d bắt đầu \n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Task %d kết thúc \n", id)
}


func main() {
	start := time.Now()		// Lưu thời gian bắt đầu
	var wg sync.WaitGroup

	for i := 1; i<= 5; i++{
		wg.Add(1)			// Tăng số lượng công việc cần chờ
		go Task(i, &wg)
	}

	wg.Wait()		// Chờ tất cả các goroutine hoàn thành

	fmt.Println("==> Thời gian thực hiện: ", time.Since(start))	//In tổng thời gian hoàn thành
}
```

Kết quả:
```bash
> go run .
Task 2 bắt đầu 
Task 3 bắt đầu
Task 1 bắt đầu
Task 4 bắt đầu
Task 5 bắt đầu
Task 4 kết thúc 
Task 3 kết thúc
Task 1 kết thúc
Task 2 kết thúc
Task 5 kết thúc
==> Thời gian thực hiện:  1.0015085s
```

---
## **Tìm hiểu về Channel khi dùng Goroutine**

Chanel trong Golang là một ống dẫn (pipeline) cho phép các Goroutine giao tiếp và truyền dữ liệu với nhau một cách an toàn. Chúng ta có thể hình dung như là một băng truyền nhà máy: một Goroutine đặt dữ liệu lên băng truyền (gửi) và một Goroutine khác lấy dữ liệu từ băng chuyền (nhận)

![](https://miro.medium.com/1*ULwliSnRHWNyMIFufxjnxg.gif)

![](https://media2.dev.to/dynamic/image/width=800%2Cheight=%2Cfit=scale-down%2Cgravity=auto%2Cformat=auto/https%3A%2F%2Fdev-to-uploads.s3.amazonaws.com%2Fuploads%2Farticles%2Ffbxksmntkjou2qq5zwa7.gif)

### **Các loại Channel**

#### **Channel không buffer (Unbuffered channel):**
Chỉ chứa 1 giá trị tại một thời điểm
 - Gửi (ch <- value) sẽ bị chặn lại cho đến khi có Goroutine nhận (ch)
 - Nhận (<-ch) sẽ bị chặn lại cho đến khi có dữ liệu gửi
 - Bắt buộc phải có ít nhất 1 cặp sender goroutine và receiver Goroutine
***Ví dụ:*** Hộp thư chỉ chứa được 1 lá thư. Người gửi phải đợi người nhận lấy thư trước khi gửi lá thư tiếp theo

Ví dụ nếu chỉ  gửi không:
```go
func main() {
	//tạo channel chỉ nhận kiểu Int
	ch := make(chan int)   //Bị block vì không thất bất kf ai nhận dữ liệu

	// Giá trị gửi vào Channel này
	ch <- 1
}
```

Kết quả sau khi chạy:
```bash
> go run .
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
```
Qua đó có thể thấy nếu chỉ bên gửi mà không có bên nhận. Nên vì thế nó bị lỗi "deadlock!" vì không có ai nhận dữ liệu

Để muốn nhận được ta quan sát ví dụ sau:
```go
func main() {
	//tạo channel chỉ nhận kiểu Int
	ch := make(chan int)

	// Channel này gửi dữ liệu có giá trị đi
	go func(){
		ch <- 1			//Block 1
		ch <- 10		//Block 2 (Nếu block 2 tạm)
		ch <- 100		//Block 3
		fmt.Println("Gửi: ")
	}()

	// Goroutine Anonymous, để nhận dữ liệu từ channel
	go func(){
		fmt.Println("Received value:", <-ch)	// Unblock 1
		fmt.Println("Received value:", <-ch)	// Unblock 2 (Nếu dừng tại Unblock 2 thì Block 3 và dòng lệnh bên dưới sẽ không in ra vì bị dừng ở tại nhánh 2 nên không thể tiếp tục thực hiện)
		fmt.Println("Nhận ")
	}()

	time.Sleep(2 * time.Second)
}
```

Kết quả:
```bash
> go run .
Received value: 1
Received value: 10
Nhận
```

Trong Channel có khái niệm là `close`.Nghĩa là sau khi toàn bộ gửi đi rồi thì nó có thể `close`  channel đi được. Lưu ý cách này chỉ sử dụng khi chúng ta **close channel**

Ví dụ:
```go
func main() {
	//tạo channel chỉ nhận kiểu Int
	ch := make(chan int)

	// Channel này gửi dữ liệu có giá trị đi
	go func(){
		defer close(ch)	// Đảm bảo đóng channel sau khi gửi xong dữ liệu
		ch <- 1			//Block 1
		ch <- 10		//Block 2
		ch <- 100		//Block 3
	}()

	for value := range ch{	// Sử dụng range để nhận dữ liệu từ channel, sẽ tự động dừng khi channel bị đóng
		fmt.Printf("Nhận giá trị: %d\n", value)
	}

	time.Sleep(2 * time.Second)
}
```

Kết quả:
```bash
> go run .
Nhận giá trị: 1
Nhận giá trị: 10
Nhận giá trị: 100
```

Nếu chúng ta để Close bên ngoài thì nó sẽ báo lỗi
ví dụ:
```go
func main() {
	//tạo channel chỉ nhận kiểu Int
	ch := make(chan int)

	// Channel này gửi dữ liệu có giá trị đi
	go func(){
		ch <- 1			//Block 1
		ch <- 10		//Block 2
		ch <- 100		//Block 3
	}()

	close(ch)

	for value := range ch{	// Sử dụng range để nhận dữ liệu từ channel, sẽ tự động dừng khi channel bị đóng
		fmt.Printf("Nhận giá trị: %d\n", value)
	}

	time.Sleep(2 * time.Second)
}
```

Kết quả bị lỗi:
```bash
> go run .
panic: send on closed channel
```

Bị lỗi này là do hàm main() chạy nhanh hơn hàm  Anonimus bên trong. Vì vừa vào nó chạy **close** liền. Nên dẫn đễn lỗi. Vì vậy muốn **close** được thì phải send hết rồi mới ***close***

#### **Channel có buffer (Buffered channel)**
Có thể chứa nhiều giá trị (Kích cỡ buffer do chúng ta dự định nghĩa)
- Gửi (ch <- value) chỉ bị chặn khi buffer đầy
- Nhận (<-ch) chỉ bị chặn khi buffer rỗng
*Ví dụ:* Hộp thư chứa được 5 lá thư. Người gửi có thể gửi 5 thư trước khi phải đợi

Ví dụ: khi  gửi ngoài phạm vi đặt trước thì nó báo lỗi như sau:
```go
func main() {
	//tạo channel chỉ nhận kiểu Int
	ch := make(chan int, 1)	// Tạo channel có buffer với kích thước 1

	ch <- 1			//Block 1
	ch <- 10		//Block 2
	ch <- 100		//Block 3
	close(ch)

	time.Sleep(2 * time.Second)
}
```

kết quả:
```bash
> go run .
fatal error: all goroutines are asleep - deadlock!
```

Ví dụ khi  gửi trong phạm vi:
```go
func main() {
	//tạo channel chỉ nhận kiểu Int
	ch := make(chan int, 5)	// Tạo channel có buffer với kích thước 5

	ch <- 1			//Block 1
	ch <- 10		//Block 2
	ch <- 100		//Block 3
	close(ch)

	for value := range ch{	// Sử dụng range để nhận dữ liệu từ channel, sẽ tự động dừng khi channel bị đóng
		fmt.Printf("Nhận giá trị: %d\n", value)
	}

	time.Sleep(2 * time.Second)
}
```

Kết quả:
```bash
> go run .
Nhận giá trị: 1
Nhận giá trị: 10
Nhận giá trị: 100
```

Channel có Buffer bị chặn khi đầy. Vì đây sử dụng Goroutine. Khi sử dụng Goroutine thì nó gửi từng cái một
Ví dụ:
```go
func main() {
	//tạo channel chỉ nhận kiểu Int
	ch := make(chan int, 5)	// Tạo channel có buffer với kích thước 10

	// Channel có Buffer bị chặn khi đầy. Vì đây sử dụng Goroutine. Khi sử dụng Goroutine thì nó gửi từng cái một
	go func(){
		defer close(ch)	// Đảm bảo rằng channel sẽ được đóng sau khi gửi xong tất cả dữ liệu
		ch <- 1			//Block 1
		ch <- 10		//Block 2
		ch <- 100		//Block 3
	}()



	for value := range ch{	// Sử dụng range để nhận dữ liệu từ channel, sẽ tự động dừng khi channel bị đóng
		fmt.Printf("Nhận giá trị: %d\n", value)
	}

	time.Sleep(2 * time.Second)
}
```

Kết quả:
```bash
> go run .
Nhận giá trị: 1
Nhận giá trị: 10
Nhận giá trị: 100
```

Qua đó nhận xét ta thấy buffer giúp chúng ta đỡ cực hơn một chút, nhưng điểm yếu là chúng ta không thể kiểm soát được tin nhắn đầu vào

---
### **Kết hợp sync.WaitGroup và Channel khi làm Goroutine**

Ví dụ:
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

//Mỗi task sẽ mất 1 giây để thực hiện
func Task(id int,ch chan<- string, wg *sync.WaitGroup){

	defer wg.Done()		// Đảm bảo gọi Done() khi công việc hoàn thành, giảm số lượng công việc cần chờ

	fmt.Printf("Task %d bắt đầu \n", id)
	time.Sleep(1 * time.Second)
	ch <- fmt.Sprintf("Task %d đã hoàn thành", id) // Gửi thông báo hoàn thành qua channel
	ch <- fmt.Sprintf("Task %d has been completed", id)
}


func main() {
	start := time.Now()		// Lưu thời gian bắt đầu
	var wg sync.WaitGroup
	ch := make(chan string) //Tạo một channel để giao tiếp với dạng string

	for i := 1; i<= 5; i++{
		wg.Add(1)			// Tăng số lượng công việc cần chờ
		go Task(i, ch, &wg)
	}

	//Dùng Go routine để bắt lấy điều này và chờ dữ liệu
	//Nếu không dùng Go routine, chương trình sẽ bị chặn tại đây và không thể nhận được thông báo hoàn thành từ channel
	// Nếu không đặt trong Go routine thì nó đang nằm ngay main thì nó sẽ chạy trước và dừng ngay tại đó
	go func(){
		wg.Wait()
		close(ch)		// Đóng channel sau khi tất cả công việc đã hoàn thành
	}()

	for value := range ch {
		fmt.Println(value) // Nhận thông báo hoàn thành từ channel
	}

	fmt.Println("==> Thời gian thực hiện: ", time.Since(start))	//In tổng thời gian hoàn thành
}
```