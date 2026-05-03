package main

import (
	"encoding/json"
	"fmt"
)

// RPCRequest là cấu trúc dữ liệu để gửi yêu cầu từ client đến server
// Nó chứa tên hàm cần gọi và các tham số
type RPCRequest struct {
	Methob string          `json:"method"` // Tên hàm cần gọi
	Params json.RawMessage `json:"params"` // Tham số của hàm, được mã hóa dưới dạng JSON
	// json.RawMessage cho phép lưu JSON dạng thô, giải mã sau khi biết loại dữ liệu
}

// RPCResponse là cấu trúc response: kết quả trả về sau khi gọi hàm
type RPCResponse struct {
	Result interface{} `json:"result"` // Kết quả trả về sau khi gọi hàm, có thể là bất kỳ kiểu dữ liệu nào
	Error  string      `json:"error"`  // Thông báo lỗi nếu có lỗi xảy ra trong quá trình gọi hàm
}

// Tham số của các hàm
type AddParams struct{ A, B int }
type MultiplyParams struct{ A, B int }


func main() {
	// Chạy server trong một goroutine riêng (chạy song song với main)
    // Nếu không có "go", server sẽ chặn thread chính và phần code dưới không chạy được
	go RunServer()

	// Gọi hàm "Add" từ xa với tham số A=10, B=15
    // CallRPC sẽ: 1) Kết nối tới server, 2) Gửi yêu cầu, 3) Nhận kết quả trả về
	result,_ := CallRPC("Add", AddParams{A: 10, B: 15})
	fmt.Println("Kết quả của Add:", result)

	// Gọi hàm Multiply từ xa
	result,_ = CallRPC("Multiply", MultiplyParams{A: 500, B: 300})
	fmt.Println("Kết quả của Multiply:", result)
}
