package __

// RPC struct
type HelloService struct {}

//Định nghĩa hàm RPC, với tham số kiểu string 
func (p *HelloService) Hello(request *String, reply *String) error {
    // các hàm như .GetValue() đã được tạo ra trong file hello.pb.go
    reply.Value = "Hello, " + request.GetValue()
    // trả về nil khi thành công
    return nil
}
