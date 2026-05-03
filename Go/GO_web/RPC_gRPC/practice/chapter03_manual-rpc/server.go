package main

import (
	"encoding/json"
	"fmt"
	"net"
)

// HandleRPC xử lý mỗi kết nối từ client
// Mỗi client kết nối tới server sẽ được xử lý bởi một goroutine riêng của HandleRPC
func HandleRPC(conn net.Conn) {
	defer conn.Close()               // Đảm bảo đống kết nối
	decoder := json.NewDecoder(conn) // Tạo decoder để đọc dữ liệu JSON từ connection, tự động parse JSON từ stream
	encoder := json.NewEncoder(conn) // Tạo encoder để gửi dữ liệu JSON về cho client, tự động chuyển đổi Go objects thành JSON

	var req RPCRequest

	// Cố gắng đọc request từ connection
	if err := decoder.Decode(&req); err != nil {
		return
	}

	// Xử lý yêu cầu và tạo phản hồi
	var resp RPCResponse

	// Dispatch: dựa vào tên hàm (req.Methob), gọi hàm tương ứng
    // Switch case tìm hàm phù hợp và thực hiện nó
	switch req.Methob {

	case "Add": // Nếu yêu cầu là "Add", giải mã tham số và thực hiện phép cộng
		var p AddParams
		json.Unmarshal(req.Params, &p)
		resp.Result = p.A + p.B

	case "Multiply": // Nếu yêu cầu là "Multiply", giải mã tham số và thực hiện phép nhân
		var p MultiplyParams
		json.Unmarshal(req.Params, &p)
		resp.Result = p.A * p.B

	default: // Nếu không tìm thấy hàm phù hợp, trả về lỗi
		resp.Error = "Unknown method" + req.Methob
	}

	// Gửi response (dạng JSON) trở lại cho client qua connection
	encoder.Encode(resp)
}

// RunServer khởi động server và lắng nghe các kết nối từ client
func RunServer() {
	ln, _ := net.Listen("tcp", ":9000")
	fmt.Println("Đang chạy server trên cổng 9000...")

	for {
		// Chấp nhận một kết nối mới từ client
        // conn là connection object để giao tiếp với client này
		conn, _ := ln.Accept()
		go HandleRPC(conn)
	}
}
