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