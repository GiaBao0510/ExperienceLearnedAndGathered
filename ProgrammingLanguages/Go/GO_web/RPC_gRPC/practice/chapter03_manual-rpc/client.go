package main

import (
	"encoding/json"
	"fmt"
	"net"
)

// CallRPC là hàm để gọi một RPC trên server từ xa
// Tham số: methob = tên hàm cần gọi, params = tham số của hàm
// Return: kết quả từ server (interface{} = có thể là bất kỳ kiểu nào), error nếu có lỗi
func CallRPC(methob string, params interface{}) (interface{}, error) {
	
	// Kết nối tới server qua TCP trên localhost cổng 9000
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil { return  nil, err}
	defer conn.Close()

	// Chuyển đổi tham số (Go object) thành JSON bytes
	parambytes,_ := json.Marshal(params)
	req := RPCRequest{Methob: methob, Params: parambytes}

	encoder := json.NewEncoder(conn)	// Tạo encoder để gửi dữ liệu JSON tới server
	decoder := json.NewDecoder(conn)	// Tạo decoder để nhận dữ liệu JSON từ server

	// Gửi request tới server (encoder sẽ tự động chuyển request thành JSON)
	encoder.Encode(req)

	var resp RPCResponse
	decoder.Decode(&resp)

	if resp.Error != "" { return  nil, fmt.Errorf(resp.Error) }

	// Nếu thành công, trả về kết quả từ server
	return resp.Result, nil
}