package main

import (
	"net"
	"net/rpc"
	"time"
)

//Định nghĩa service struct
type GreetingService struct{}

/*Định nghĩa hàm service theo nguyên tắc:
1. Hàm service phải public (bắt đầu bằng chữ hoa)
2. Hàm service phải có 2 tham số, trong đó tham số thứ nhất là request, tham số thứ hai response kiểu pointer
3. phải trả về kiểu error
*/
func (p *GreetingService) Hello(req string, res *string) error {
	*res = "Hello " + req

	return nil
}

func main() {

	// Đăng ký service với tên "GreetingService" có kiểu là GreetingService
	rpc.Register(new(GreetingService)) 

	for {
		//Chủ động gọi đến client
		conn, _ := net.Dial("tcp", "localhost:1234")

		if conn == nil {
			time.Sleep(time.Second)
			continue
		}

		rpc.ServeConn(conn)
		conn.Close()
	}
}