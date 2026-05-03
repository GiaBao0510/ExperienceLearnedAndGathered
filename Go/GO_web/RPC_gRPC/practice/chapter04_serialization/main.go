package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"time"
)

// User là một cấu trúc dữ liệu mẫu để thử nghiệm serialization
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// benchmarkJson so sánh tốc độ và kích thước của JSON serialization
// u = User object cần benchmark
// n = số lần lặp để benchmark
// Return: (kích thước dữ liệu JSON, thời gian thực hiện)
func benchmarkJson(u User, n int) (int, time.Duration) {
	start := time.Now()	// Ghi lại thời gian bắt đầu
	var size int

	for i := 0; i < n; i++ {
	
		data, _ := json.Marshal(u) // Serialize cấu trúc dữ liệu User thành JSON
		size = len(data)           // Lấy kích thước của dữ liệu JSON

		// Chuyển đổi JSON bytes trở lại thành User struct
        // Unmarshal: JSON bytes → Go object
		var u2 User
		json.Unmarshal(data, &u2)
	}

	return size, time.Since(start)
}

// benchmarkGob so sánh tốc độ và kích thước của GOB serialization
// GOB là định dạng nhị phân của Go, nhỏ gọn và nhanh hơn JSON
// u = User object cần benchmark
// n = số lần lặp để benchmark
// Return: (kích thước dữ liệu GOB, thời gian thực hiện)
func benchmarkGob(u User, n int) (int, time.Duration) {
	start := time.Now()	// Ghi lại thời gian bắt đầu

	var size int
	for i := 0; i < n; i++ {
		// Tạo buffer để lưu dữ liệu GOB
		var buf bytes.Buffer

		gob.NewEncoder(&buf).Encode(u) // Serialize cấu trúc dữ liệu User thành GOB
		size = buf.Len()               // Lấy kích thước của dữ liệu GOB

		// ⚠️ BUG FIX: Reset buffer để chuẩn bị cho vòng lặp tiếp theo
        // Nếu không reset, vòng lặp tiếp theo sẽ thêm dữ liệu vào thay vì thay thế
		buf.Reset()

		// Chuyển đổi GOB bytes trở lại thành User struct
        // Gob decoding: GOB bytes → Go object
		var u2 User
		gob.NewDecoder(&buf).Decode(&u2) // Deserialize dữ liệu GOB thành cấu trúc dữ liệu User
	}

	return size, time.Since(start)
}

// simulateProtobuf mô phỏng Protobuf (một định dạng serialization khác)
// Protobuf là định dạng nhị phân được phát triển bởi Google, rất hiệu quả
// Trong ví dụ này, ta chỉ mô phỏng kích thước, không mô phỏng thực tế quá trình encoding
//
// Các định dạng serialization:
// - JSON: dễ đọc, dễ debug, nhưng size lớn (text-based)
// - GOB: nhị phân, nhanh, size nhỏ hơn, nhưng chỉ dùng trong Go
// - Protobuf: nhị phân, nhanh, size nhỏ nhất, được hỗ trợ bởi nhiều ngôn ngữ
func simulateProtobuf(u User) []byte {
	// Protobuf dùng field number + varint encoding
	// Ở đây ta chỉ mô phỏng kích thước nhỏ hơn
	data,_ := json.Marshal(u) // Sử dụng JSON để mô phỏng dữ liệu nhỏ hơn

	// Protobuf thực tế sẽ nhỏ hơn JSON, ta giả sử nó chỉ chiếm 60% kích thước của JSON
	return data[:len(data)*6/10]	
}

func main(){
	User := User{
		ID: 1,
		Name: "John Doe",
		Email: "john.doe@example.com",
		Phone: "123-456-7890",
	}

	// Số lần lặp để benchmark (càng nhiều càng chính xác, nhưng chậm hơn)
	n := 100000

	// Chạy benchmark cho cả 3 phương pháp
	jsonSize, jsonDur := benchmarkJson(User, n)
	gobSize, gobDur := benchmarkGob(User, n)
	protobufData := simulateProtobuf(User)

	println("JSON Size:", jsonSize, "bytes, Time:", jsonDur.String())
	println("GOB Size:", gobSize, "bytes, Time:", gobDur.String())
	println("Simulated Protobuf Size:", len(protobufData), "bytes")
}