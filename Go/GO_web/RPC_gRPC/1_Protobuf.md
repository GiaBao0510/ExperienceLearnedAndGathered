[Protobuf](https://developers.google.com/protocol-buffers/) hay Protocols Buffer là một ngôn ngữ dùng để mô tả các cấu trúc dữ liệu, chúng ta dùng protoc để biên dịch chúng thành mã nguồn của các ngôn ngữ lập trình khác nhau có chức năng serialize và deserialize các cấu trúc dữ liệu này thành dạng binary stream. So với dạng XML hoặc JSON thì dữ liệu đó nhỏ gọn gấp 3-10 lần và được xử lý rất nhanh.

![](https://zalopay-oss.github.io/go-advanced/images/ch4-2-size.png)  
![](https://zalopay-oss.github.io/go-advanced/images/ch4-2-speed.png)  

_Xem thêm: [Benchmarking Protocol Buffers, JSON and XML in Go](https://medium.com/@shijuvar/benchmarking-protocol-buffers-json-and-xml-in-go-57fa89b8525)_.
Hoặc xem video cách để cài đặt: [Cài đặt](https://www.youtube.com/results?search_query=proto+vscode&sp=EgIIBQ%253D%253D)

Kiểm tra cài đặt
```bash
>protoc --version
libprotoc 33.5
```

## Kết hợp Protobuf với RPC

Đầu tiên chúng ta tạo file `hello.proto` chứa kiểu String được dùng trong RPC HelloService.
**GO_web\RPC_gRPC\practice\protobuf\hello.proto**
```proto
//phiển bản bản của protobuf
syntax = "proto3";

//Tên package của protobuf
package main;

// Quy định đầu ra được đặt ở đâu
option go_package = "./";

//message là một đơn vị dữ liệu trong protobuf
message String {
    //Chuỗi strind được truyền vào message
    string str = 1;
}
```

Để sinh ra mã nguồn Go từ file `hello.proto` ở trên, đầu tiên là cài đặt bộ biên dịch `protoc` qua liên kết [ở đây](https://github.com/google/protobuf/releases), sau đó là cài đặt một plugin cho Go thông qua lệnh:

```
> go get github.com/golang/protobuf/protoc-gen-go
```

Chúng ta sẽ sinh ra mã nguồn Go bằng lệnh sau:
```bash
\Go\GO_web\RPC_gRPC\practice\protobuf> protoc --go_out=. hello.proto
```

Trong đó:
- protoc: chương trình sinh mã nguồn
 - go_out: chỉ cho protoc tải plugin protoc-gen-go, (cũng có java_out, python_out,..)
- --go_out=.: sinh ra mã nguồn tại thư mục hiện tại
- hello.proto: file Protobuf

Sẽ có một file `hello.pb.go` được sinh ra, trong đó cấu trúc String được định nghĩa là:

```go
type String struct {
    Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
    //...
}

func (m *String) Reset()         { *m = String{} }
func (m *String) String() string { return proto.CompactTextString(m) }
func (*String) ProtoMessage()    {}
func (*String) Descriptor() ([]byte, []int) {
    return fileDescriptor_hello_069698f99dd8f029, []int{0}
}
//...
func (m *String) GetValue() string {
    if m != nil {
        return m.Value
    }
    return ""
}
//...
```

Ở [phần trước](Go/GO_web/RPC_gRPC/0.Introdue_RPC_gRPC) chúng ta đã xây dựng một RPC HelloService đơn giản dựa trên thư viện chuẩn [net/rpc](https://godoc.org/net/rpc) có kiểu dữ liệu request, reply do người dùng tự định nghĩa, bây giờ dựa trên kiểu String mới được sinh ra từ Protobuf, chúng ta có thể viết lại RPC HelloService như sau:

**hello.go:**
```go
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
```


Chúng ta vẫn phải tự xây dựng hàm **Hello(request, reply)** bằng cách tự viết. Khi sử dụng Protobuf chúng ta có thể tự định nghĩa luôn service mình có những hàm rpc nào, nhận vào request và trả về reply ra sao. Chúng ta định nghĩa HelloService trong file proto như sau:

**GO_web\RPC_gRPC\practice\protobuf\hello.proto**
```protoc
//Đinh nghĩa một service trong protobuf
service HelloService {
    // định nghĩa lời gọi hàm RPC
    rpc SayHello (String) return (String);
}
```

Chúng ta cần có một plugin để sinh ra mã nguồn service tương ứng với định nghĩa ở trên. Hiện nay Google đã phát triển bộ [gRPC plugin](https://github.com/golang/protobuf/blob/master/protoc-gen-go/grpc/grpc.go) giúp t**ạo ra mã nguồn tương ứng với file proto**. Ở phần dưới sẽ trình bày **cách xây dựng một plugin dựa trên mã nguồn gRPC plugin**, chi tiết về gRPC  sẽ đề cập ở các phần sau.

---
## Viết plugin sinh mã nguồn RPC service

Từ mã nguồn [gRPC plugin](https://github.com/golang/protobuf/blob/master/protoc-gen-go/grpc/grpc.go), chúng ta có thể thấy hàm `generator.RegisterPlugin` được dùng để đăng kí `plugin` đó, Interface của một plugin sẽ như sau: