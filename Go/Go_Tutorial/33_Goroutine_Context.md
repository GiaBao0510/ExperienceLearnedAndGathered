**Context** trong Golang là một cơ chế ==giúp quản lý và kiểm soát thời gian sống cũng như hành vi của các tác vụ== trong chương trình, đặc biệt khi làm việc với Goroutine. Chúng ta có thể hình dung Context như một bộ điều khiển từ xa cho phép:
- Hủy (cancel) các Goroutine khi không cần thiết
- Đặt thời gian sống (timeout) để dừng tác vụ sau một khoản thời gian nào đó
- Truyền dữ liệu (Metadata) giữa các Goroutine mà không cần biến toàn cục
Vậy tóm tại Context giúp quản lý kiểm soát thời gian sống  và truyền dữ liệu giữa các **Goroutine**. 
Ví dụ: Manager giao một Task cho nhân việ trong 2 giây, có thể hủy bất cứ lúc nào. Sau 2 giây sẽ hủy. Mặc dù có hoàn thành hay chưa

Ví dụ:
```go
package main

import (
	"fmt"
	"context"
	"time"
)

func main() {
	//Tạo context trong 2 giây
	ctx, cancel := context.WithTimeout(
		context.Background(),				// Đang chạy ở Background
		2 * time.Second,					//Thời gian sống của chương trình chỉ 2 giây
	)
	defer cancel()	//Đảm bảo hủy context sau khi sử dụng xong, để tránh leak. Nếu để bị leak thì chương trình bị gãy

	//Gừi thêm ghi chú vào context (ví dụ độ ưu tiên)
	ctx = context.WithValue(ctx, "priority", "high")

	//Giao task cho nhân viên
	go assignTask(ctx)

	//Đợi 3s để xem task bị hủy tự động
	time.Sleep(3 * time.Second)
}

func assignTask(ctx context.Context){
	
	//Tạo vòng lặp vô tận
	for{

		//Dùng select để
		select{
			//Nhận tín hiệu hủy từ context
			case <- ctx.Done():
				fmt.Println("Công việc bị hủy: ", ctx.Err())
				return
			default:
				//Tiếp tục làm việc
				priority := ctx.Value("priority")
				fmt.Printf("Đang làm việc ... Ưu tiên: %s.\n",priority)
				time.Sleep(500 * time.Millisecond) // Giả lập công việc mất thời gian
		}
	}
}
```

Kết quả:
```bash
> go run .
Đang làm việc ... Ưu tiên: high.
Đang làm việc ... Ưu tiên: high.
Đang làm việc ... Ưu tiên: high.
Đang làm việc ... Ưu tiên: high.
Công việc bị hủy:  context deadline exceeded
```

---
### Áp dụng thực tế Context và Channel khi làm với Goroutine

```go
package main

import (
	"fmt"
	"context"
	"time"
)



func main() {
	// Tạo 2 channel
	phoCh := make(chan string)
	pizzaCh := make(chan string)
	BanCuonCh := make(chan string)

	//Tạo context với timeout 1500ms
	ctx, cancel := context.WithTimeout(context.Background(), 1500 * time.Millisecond)
	defer cancel() 	// Để đảm bảo rằng cancel sẽ được gọi khi main function kết thúc, giúp giải phóng tài nguyên liên quan đến context.

	//Khởi tạo bằng Goroutine
	go CookPho(ctx, phoCh)
	go CookPizza(ctx, pizzaCh)
	go CookBanCuon(ctx, BanCuonCh)

	//Nhận dữ liệu với select
	for i := 0; i< 3; i++{
		select{
			case phoResult := <- phoCh:
				fmt.Println("Nhận được: ",phoResult)
			case pizzaResult := <- pizzaCh:
				fmt.Println("Nhận được: ", pizzaResult)
			case banhCuonResult := <- BanCuonCh:
				fmt.Println("Nhận được: ", banhCuonResult)
			case <- ctx.Done():
				fmt.Println("Hết thời gian chờ, hủy các món ăn còn lại")
		}	
	}
}

func CookPho(ctx context.Context, ch chan string){
	fmt.Println("Bắt đầu nấu phở ....")
	select{
		case <- time.After(1 * time.Second):	// Giả sử nấu phở mất 1 giây
			ch <- "Phở đã nấu xong"

		case <- ctx.Done():	// Nếu context bị hủy trước khi nấu xong, sẽ nhận được tín hiệu từ ctx.Done()
			fmt.Println("Hủy nấu phở")
			return
	}
}

func CookPizza(ctx context.Context, ch chan string){
	fmt.Println("Bắt đầu nướng Pizza ....")
	select {
		case <- time.After(2 * time.Second):	// Giả sử nướng Pizza mất 2 giây
			ch <- "Pizza đã nướng xong"

		case <-ctx.Done(): // Nếu context bị hủy trước khi nướng xong, sẽ nhận được tín hiệu từ ctx.Done()
			fmt.Println("Hủy nướng Pizza")
			return
	}
}

func CookBanCuon(ctx context.Context, ch chan string){
	fmt.Println("Bắt đầu nấu bánh cuốn ....")
	select {
		case <- time.After(1499 * time.Millisecond): 	// Giả sử nấu bánh cuốn mất 1200ms
			ch <- "Bánh cuốn đã nấu xong"
		case <- ctx.Done():
			fmt.Println("Hủy nấu bánh cuốn")
			return
	}
}
```

