# Protocol Buffers - Hướng dẫn cho Sinh viên

## 1. Khái niệm cơ bản

### 1.1. Protocol Buffers là gì?

Protocol Buffers (thường gọi tắt là Protobuf) là một ngôn ngữ mô tả cấu trúc dữ liệu được phát triển bởi Google vào năm 2008. Protobuf lưu trữ dữ liệu có cấu trúc dưới dạng nhị phân (binary format), giúp truyền tải dữ liệu qua mạng nhanh hơn và hiệu quả hơn so với JSON hoặc XML.

**Đặc điểm chính:**

- Độc lập với nền tảng và ngôn ngữ lập trình
- Định nghĩa cấu trúc dữ liệu một lần bằng file `.proto`
- Tự động sinh code cho nhiều ngôn ngữ lập trình
- Mã nguồn mở và miễn phí

### 1.2. Protobuf hoạt động như thế nào?

Quy trình làm việc với Protobuf:

1. Định nghĩa cấu trúc dữ liệu trong file `.proto`
2. Sử dụng trình biên dịch `protoc` để sinh code
3. Sử dụng code đã sinh để serialize/deserialize dữ liệu
4. Truyền tải hoặc lưu trữ dữ liệu dưới dạng nhị phân

### 1.3. Protobuf được sử dụng ở đâu?

Protobuf thường được sử dụng trong:

- Giao tiếp giữa các microservice (đặc biệt với gRPC)
- Lưu trữ dữ liệu cấu hình
- Giao tiếp client-server
- Lưu trữ dữ liệu có cấu trúc trong database
- API internal của các hệ thống phân tán

---

## 2. Tại sao nên sử dụng Protocol Buffers?

### 2.1. Ưu điểm

**Kích thước nhỏ gọn:**

- Dữ liệu nhỏ hơn 3-10 lần so với XML hoặc JSON
- Tiết kiệm băng thông mạng và dung lượng lưu trữ

**Tốc độ xử lý nhanh:**

- Encoding và decoding nhanh hơn JSON/XML
- Phân tích cú pháp (parsing) đơn giản hơn

**Type-safe:**

- Định nghĩa rõ ràng kiểu dữ liệu
- Phát hiện lỗi tại thời điểm biên dịch
- Giảm thiểu lỗi runtime

**Khả năng mở rộng:**

- Thêm trường mới mà không ảnh hưởng code cũ
- Tương thích ngược (backward compatibility)
- Tương thích tiến (forward compatibility)

**Hỗ trợ đa ngôn ngữ:**

- C++, C#, Java, Kotlin, Python, Go, Ruby, PHP, JavaScript, và nhiều ngôn ngữ khác

### 2.2. Nhược điểm

**Không thân thiện với con người:**

- Dữ liệu dạng nhị phân, không thể đọc trực tiếp
- Cần công cụ để debug và kiểm tra dữ liệu

**Cần schema:**

- Phải có file `.proto` để deserialize dữ liệu
- Khó khăn khi không biết cấu trúc dữ liệu

**Không phù hợp cho browser:**

- JSON vẫn là lựa chọn tốt hơn cho web API public
- Hỗ trợ hạn chế trong môi trường JavaScript thuần

---

## 3. So sánh Protobuf với các định dạng khác

### 3.1. Protobuf vs JSON

|Tiêu chí|Protobuf|JSON|
|---|---|---|
|Kích thước|Nhỏ (3-10x nhỏ hơn)|Lớn hơn|
|Tốc độ|Nhanh hơn|Chậm hơn|
|Khả năng đọc|Khó (binary)|Dễ (text)|
|Schema|Bắt buộc|Không bắt buộc|
|Hỗ trợ browser|Hạn chế|Rất tốt|
|Use case|Internal API, microservices|Public API, web|

### 3.2. Protobuf vs XML

|Tiêu chí|Protobuf|XML|
|---|---|---|
|Kích thước|Nhỏ hơn nhiều|Lớn và dài dòng|
|Tốc độ|Nhanh hơn nhiều|Chậm|
|Khả năng đọc|Khó|Dễ|
|Validation|Có (schema)|Có (XSD)|
|Công cụ|Ít hơn|Nhiều (XSLT, XPath)|

### 3.3. Khi nào nên dùng gì?

**Nên dùng Protobuf khi:**

- Xây dựng microservices với gRPC
- Cần tối ưu băng thông và tốc độ
- Giao tiếp internal giữa các service
- Lưu trữ dữ liệu có cấu trúc phức tạp

**Nên dùng JSON khi:**

- Xây dựng REST API public
- Cần con người đọc/chỉnh sửa dễ dàng
- Làm việc với browser/JavaScript
- Rapid prototyping

**Nên dùng XML khi:**

- Làm việc với hệ thống legacy (SOAP)
- Cần validation phức tạp
- Yêu cầu của ngành/tổ chức (ví dụ: tài chính, y tế)

---

## 4. Cú pháp Protocol Buffers

### 4.1. Phiên bản Protobuf

Có hai phiên bản chính:

- **proto2**: Phiên bản cũ, vẫn được hỗ trợ
- **proto3**: Phiên bản mới (khuyến nghị sử dụng), đơn giản và dễ sử dụng hơn

Hướng dẫn này sử dụng **proto3**.

### 4.2. Cấu trúc cơ bản của file .proto

```proto
// Khai báo phiên bản
syntax = "proto3";

// Khai báo package (tùy chọn, nhưng nên có)
package customer;

// Định nghĩa message
message Person {
    string name = 1;
    int32 id = 2;
    string email = 3;
}
```

### 4.3. Các kiểu dữ liệu cơ bản

|Protobuf Type|Go Type|Mô tả|
|---|---|---|
|`double`|`float64`|Số thực 64-bit|
|`float`|`float32`|Số thực 32-bit|
|`int32`|`int32`|Số nguyên 32-bit|
|`int64`|`int64`|Số nguyên 64-bit|
|`uint32`|`uint32`|Số nguyên không âm 32-bit|
|`uint64`|`uint64`|Số nguyên không âm 64-bit|
|`bool`|`bool`|Boolean|
|`string`|`string`|Chuỗi UTF-8|
|`bytes`|`[]byte`|Mảng byte|

### 4.4. Field Number (Số thứ tự trường)

Mỗi trường phải có một số thứ tự **duy nhất** từ 1 đến 536,870,911.

```proto
message User {
    string username = 1;  // Field number 1
    string email = 2;     // Field number 2
    int32 age = 3;        // Field number 3
}
```

**Quy tắc quan trọng:**

1. Field number phải duy nhất trong một message
2. Không được dùng số từ 19000 đến 19999 (reserved cho Protobuf)
3. Không được thay đổi field number sau khi đã sử dụng
4. Nên dùng số 1-15 cho các trường hay dùng (chỉ tốn 1 byte để encode)
5. Số 16-2047 tốn 2 bytes để encode

**Tại sao field number quan trọng?**

Field number được dùng để xác định trường khi serialize/deserialize. Nếu bạn thay đổi field number, dữ liệu cũ sẽ không đọc được đúng.

### 4.5. Field Rules (Quy tắc trường)

Trong proto3, có các quy tắc sau:

**Singular (mặc định):**

```proto
message User {
    string name = 1;  // Một giá trị name
}
```

**Optional:**

```proto
message User {
    optional string middle_name = 1;  // Có thể có hoặc không
}
```

**Repeated (mảng):**

```proto
message User {
    repeated string hobbies = 1;  // Danh sách sở thích
}
```

**Map (key-value):**

```proto
message User {
    map<string, int32> scores = 1;  // Điểm số theo môn học
}
```

### 4.6. Message lồng nhau

```proto
message Customer {
    string name = 1;
    int32 id = 2;
    
    // Message Address được định nghĩa bên trong Customer
    message Address {
        string street = 1;
        string city = 2;
        string state = 3;
        string zip = 4;
    }
    
    // Một customer có thể có nhiều địa chỉ
    repeated Address addresses = 3;
}
```

### 4.7. Enum (Kiểu liệt kê)

```proto
message Task {
    string title = 1;
    
    enum Status {
        UNKNOWN = 0;     // Giá trị đầu tiên phải là 0
        PENDING = 1;
        IN_PROGRESS = 2;
        COMPLETED = 3;
        CANCELLED = 4;
    }
    
    Status status = 2;
}
```

**Lưu ý:** Giá trị đầu tiên của enum trong proto3 **phải là 0**.

### 4.8. Default Values (Giá trị mặc định)

Trong proto3, khi một trường không được gán giá trị:

|Kiểu dữ liệu|Giá trị mặc định|
|---|---|
|`string`|`""` (chuỗi rỗng)|
|`bytes`|Empty bytes|
|`bool`|`false`|
|`numeric` (int, float, etc.)|`0`|
|`enum`|Giá trị đầu tiên (phải là 0)|
|`repeated`|Empty list|

### 4.9. Reserved Fields (Trường dành riêng)

Khi xóa một trường, nên đánh dấu field number hoặc tên của nó là `reserved` để tránh xung đột sau này:

```proto
message Customer {
    reserved 2, 15, 9 to 11;  // Reserved field numbers
    reserved "old_field", "deprecated_field";  // Reserved field names
    
    string name = 1;
    // Field 2 đã bị xóa và được reserved
    string email = 3;
}
```

**Tại sao cần reserved?**

Ngăn developer tương lai vô tình sử dụng lại field number hoặc tên đã bị xóa, gây ra lỗi khi deserialize dữ liệu cũ.

### 4.10. Comments (Chú thích)

```proto
syntax = "proto3";

// Đây là comment một dòng

/*
 * Đây là comment
 * nhiều dòng
 */

message Person {
    string name = 1;      // Tên người dùng
    int32 id = 2;         // ID duy nhất
    string email = 3;     // Email liên hệ
}
```

---

## 5. Sử dụng Protocol Buffers với Golang

### 5.1. Cài đặt công cụ

**Bước 1: Cài đặt Protocol Buffer Compiler (protoc)**

```bash
# Trên Linux/macOS
# Download từ https://github.com/protocolbuffers/protobuf/releases
# Hoặc dùng package manager:

# Ubuntu/Debian
sudo apt install -y protobuf-compiler

# macOS
brew install protobuf
```

**Bước 2: Cài đặt protoc-gen-go plugin**

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Đảm bảo $GOPATH/bin nằm trong PATH
export PATH="$PATH:$(go env GOPATH)/bin"
```

### 5.2. Ví dụ hoàn chỉnh

#### Bước 1: Tạo file .proto

Tạo file `user.proto`:

```proto
syntax = "proto3";

package user;

// Chỉ định package Go sẽ được tạo
option go_package = "github.com/yourname/yourproject/pb";

// Định nghĩa message User
message User {
    int32 id = 1;
    string username = 2;
    string email = 3;
    int32 age = 4;
    
    enum Role {
        UNKNOWN = 0;
        ADMIN = 1;
        USER = 2;
        GUEST = 3;
    }
    
    Role role = 5;
    repeated string hobbies = 6;
}

// Định nghĩa message UserList
message UserList {
    repeated User users = 1;
}
```

#### Bước 2: Sinh code Go

```bash
# Tạo thư mục để chứa code được sinh
mkdir -p pb

# Sinh code Go từ file .proto
protoc --go_out=. --go_opt=paths=source_relative user.proto
```

File `user.pb.go` sẽ được tạo ra trong thư mục `pb`.

#### Bước 3: Sử dụng trong code Go

Tạo file `main.go`:

```go
package main

import (
    "fmt"
    "log"
    
    "google.golang.org/protobuf/proto"
    pb "github.com/yourname/yourproject/pb"
)

func main() {
    // Tạo một User object
    user := &pb.User{
        Id:       1,
        Username: "john_doe",
        Email:    "john@example.com",
        Age:      25,
        Role:     pb.User_ADMIN,
        Hobbies:  []string{"reading", "coding", "gaming"},
    }
    
    fmt.Println("Original User:")
    fmt.Printf("ID: %d\n", user.Id)
    fmt.Printf("Username: %s\n", user.Username)
    fmt.Printf("Email: %s\n", user.Email)
    fmt.Printf("Age: %d\n", user.Age)
    fmt.Printf("Role: %s\n", user.Role)
    fmt.Printf("Hobbies: %v\n", user.Hobbies)
    
    // Serialize (Marshal) - chuyển object thành binary
    data, err := proto.Marshal(user)
    if err != nil {
        log.Fatalf("Failed to marshal: %v", err)
    }
    
    fmt.Printf("\nSerialized data size: %d bytes\n", len(data))
    fmt.Printf("Serialized data (hex): %x\n", data)
    
    // Deserialize (Unmarshal) - chuyển binary thành object
    newUser := &pb.User{}
    err = proto.Unmarshal(data, newUser)
    if err != nil {
        log.Fatalf("Failed to unmarshal: %v", err)
    }
    
    fmt.Println("\nDeserialized User:")
    fmt.Printf("ID: %d\n", newUser.Id)
    fmt.Printf("Username: %s\n", newUser.Username)
    fmt.Printf("Email: %s\n", newUser.Email)
    fmt.Printf("Age: %d\n", newUser.Age)
    fmt.Printf("Role: %s\n", newUser.Role)
    fmt.Printf("Hobbies: %v\n", newUser.Hobbies)
    
    // So sánh hai object
    if proto.Equal(user, newUser) {
        fmt.Println("\nTwo users are equal!")
    }
}
```

#### Bước 4: Cài đặt dependencies

```bash
go mod init github.com/yourname/yourproject
go get google.golang.org/protobuf/proto
```

#### Bước 5: Chạy chương trình

```bash
go run main.go
```

**Kết quả mẫu:**

```
Original User:
ID: 1
Username: john_doe
Email: john@example.com
Age: 25
Role: ADMIN
Hobbies: [reading coding gaming]

Serialized data size: 58 bytes
Serialized data (hex): 08011209...

Deserialized User:
ID: 1
Username: john_doe
Email: john@example.com
Age: 25
Role: ADMIN
Hobbies: [reading coding gaming]

Two users are equal!
```

### 5.3. Ví dụ với Message lồng nhau

```proto
syntax = "proto3";

package order;

option go_package = "github.com/yourname/yourproject/pb";

message Order {
    int32 order_id = 1;
    string customer_name = 2;
    
    message Item {
        string product_name = 1;
        int32 quantity = 2;
        float price = 3;
    }
    
    repeated Item items = 3;
    float total_amount = 4;
}
```

Sử dụng trong Go:

```go
package main

import (
    "fmt"
    "log"
    
    "google.golang.org/protobuf/proto"
    pb "github.com/yourname/yourproject/pb"
)

func main() {
    order := &pb.Order{
        OrderId:      12345,
        CustomerName: "Alice",
        Items: []*pb.Order_Item{
            {
                ProductName: "Laptop",
                Quantity:    1,
                Price:       999.99,
            },
            {
                ProductName: "Mouse",
                Quantity:    2,
                Price:       25.50,
            },
        },
        TotalAmount: 1050.99,
    }
    
    // Serialize
    data, err := proto.Marshal(order)
    if err != nil {
        log.Fatalf("Failed to marshal: %v", err)
    }
    
    fmt.Printf("Order serialized: %d bytes\n", len(data))
    
    // Deserialize
    newOrder := &pb.Order{}
    err = proto.Unmarshal(data, newOrder)
    if err != nil {
        log.Fatalf("Failed to unmarshal: %v", err)
    }
    
    fmt.Printf("Order ID: %d\n", newOrder.OrderId)
    fmt.Printf("Customer: %s\n", newOrder.CustomerName)
    fmt.Printf("Total: $%.2f\n", newOrder.TotalAmount)
    
    fmt.Println("\nItems:")
    for i, item := range newOrder.Items {
        fmt.Printf("%d. %s x%d @ $%.2f\n", 
            i+1, item.ProductName, item.Quantity, item.Price)
    }
}
```

### 5.4. Lưu và đọc dữ liệu từ file

```go
package main

import (
    "fmt"
    "log"
    "os"
    
    "google.golang.org/protobuf/proto"
    pb "github.com/yourname/yourproject/pb"
)

func main() {
    user := &pb.User{
        Id:       1,
        Username: "alice",
        Email:    "alice@example.com",
        Age:      30,
    }
    
    // Serialize và lưu vào file
    data, err := proto.Marshal(user)
    if err != nil {
        log.Fatalf("Failed to marshal: %v", err)
    }
    
    err = os.WriteFile("user.bin", data, 0644)
    if err != nil {
        log.Fatalf("Failed to write file: %v", err)
    }
    fmt.Println("User saved to user.bin")
    
    // Đọc từ file và deserialize
    fileData, err := os.ReadFile("user.bin")
    if err != nil {
        log.Fatalf("Failed to read file: %v", err)
    }
    
    loadedUser := &pb.User{}
    err = proto.Unmarshal(fileData, loadedUser)
    if err != nil {
        log.Fatalf("Failed to unmarshal: %v", err)
    }
    
    fmt.Printf("Loaded user: %s (ID: %d)\n", loadedUser.Username, loadedUser.Id)
}
```

---

## 6. Best Practices (Thực hành tốt)

### 6.1. Đặt tên

- Message names: PascalCase (`UserProfile`, `OrderItem`)
- Field names: snake_case (`user_name`, `order_id`)
- Enum names: UPPER_SNAKE_CASE (`PENDING`, `IN_PROGRESS`)
- File names: snake_case (`user_profile.proto`)

### 6.2. Tổ chức file

```
project/
├── proto/
│   ├── user.proto
│   ├── order.proto
│   └── common.proto
├── pb/
│   ├── user.pb.go
│   ├── order.pb.go
│   └── common.pb.go
└── main.go
```

### 6.3. Versioning

Khi cần thay đổi schema:

**Được phép:**

- Thêm trường mới
- Thêm message mới
- Thêm giá trị enum mới

**Không nên:**

- Xóa trường (nên dùng `reserved`)
- Thay đổi field number
- Thay đổi kiểu dữ liệu của trường

### 6.4. Performance tips

1. Dùng field number 1-15 cho các trường hay dùng
2. Tránh message lồng quá sâu (> 3 levels)
3. Sử dụng `bytes` thay vì `string` cho dữ liệu binary
4. Cân nhắc sử dụng `oneof` để tiết kiệm bộ nhớ

### 6.5. Sử dụng oneof

`oneof` cho phép một trường có thể là một trong nhiều kiểu:

```proto
message Payment {
    int32 amount = 1;
    
    oneof payment_method {
        string credit_card = 2;
        string paypal_email = 3;
        string bank_account = 4;
    }
}
```

Trong Go:

```go
payment := &pb.Payment{
    Amount: 100,
    PaymentMethod: &pb.Payment_CreditCard{
        CreditCard: "1234-5678-9012-3456",
    },
}

// Kiểm tra loại payment method
switch x := payment.PaymentMethod.(type) {
case *pb.Payment_CreditCard:
    fmt.Printf("Credit card: %s\n", x.CreditCard)
case *pb.Payment_PaypalEmail:
    fmt.Printf("PayPal: %s\n", x.PaypalEmail)
case *pb.Payment_BankAccount:
    fmt.Printf("Bank: %s\n", x.BankAccount)
}
```


---

## Tài liệu tham khảo

### Official Documentation

- Protocol Buffers Guide: https://protobuf.dev/
- Go Protobuf Tutorial: https://protobuf.dev/getting-started/gotutorial/
- Language Guide (proto3): https://protobuf.dev/programming-guides/proto3/

### Tools

- protoc compiler: https://github.com/protocolbuffers/protobuf
- protoc-gen-go: https://pkg.go.dev/google.golang.org/protobuf/cmd/protoc-gen-go
- Online Protobuf Editor: https://protobuf-decoder.netlify.app/

###  Libraries

- Go Protobuf: https://pkg.go.dev/google.golang.org/protobuf
- gRPC Go: https://grpc.io/docs/languages/go/

---

Tổng kết

Protocol Buffers là một công cụ mạnh mẽ để serialize dữ liệu có cấu trúc. Với những ưu điểm về kích thước và tốc độ, Protobuf đặc biệt phù hợp cho các hệ thống microservices và ứng dụng cần hiệu suất cao.

**Các điểm chính cần nhớ:**

1. Protobuf lưu dữ liệu dưới dạng binary, nhỏ gọn và nhanh
2. Cần định nghĩa schema trong file `.proto`
3. Field number rất quan trọng, không được thay đổi sau khi deploy
4. Tương thích với nhiều ngôn ngữ lập trình
5. Thích hợp cho internal API, không phù hợp cho public web API