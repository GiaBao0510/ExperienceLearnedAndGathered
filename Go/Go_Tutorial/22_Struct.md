**Struct** là một kiểu dữ liệu tổng hợp (composite type) trong Go, cho phép nhóm các trường (field) có kiểu dữ liệu khác nhau thành một đơn vị duy nhất. Nó giống như một "lớp" đơn giản trong các ngôn ngữ khác, giúp tổ chức dữ liệu rõ ràng, dễ quản lý và tái sử dụng – đặc biệt hữu ích cho việc mô hình hóa đối tượng thực tế như sinh viên, sản phẩm, hoặc giảng viên.

### Khai Báo Và Sử Dụng Struct

- **Cú pháp**: ==type TênStruct struct { field1 kiểu1; field2 kiểu2; ... }==
- Struct có thể chứa bất kỳ kiểu dữ liệu nào (string, int, bool, thậm chí struct khác).
- Khởi tạo: Sử dụng **{field: giá_trị}** hoặc theo thứ tự (nhưng khuyến nghị dùng tên field để rõ ràng).

![](https://selftuts.in/wp-content/uploads/2024/07/image-768x446.png)

Ví dụ thực tế: Mô hình hóa thông tin giảng viên.
```go
package main
import "fmt"

type GiangVien struct {
    Name   string
    Email  string
    Gender int  // 1: Nam, 0: Nữ, -1: Khác
}

func main() {
    gv1 := GiangVien{
        Name:   "Pham Gia Bao",
        Email:  "pgbao123@gmail.com",
        Gender: 1,
    }

    fmt.Printf("Kiểu dữ liệu: %T\n", gv1)
    fmt.Printf("Họ tên: %s\n", gv1.Name)
    fmt.Printf("Email: %s\n", gv1.Email)
    fmt.Printf("Giới tính: %d\n", gv1.Gender)
}
```

Kết quả hiển thi:
```shell
Kiểu dữ liệu: main.GiangVien
Họ tên: Pham Gia Bao
Email: pgbao123@gmail.com
Giới tính: 1
```

### Struct Với Con Trỏ (Pointer)

- Go truyền struct theo giá trị (copy), tốn bộ nhớ nếu struct lớn. Sử dụng pointer (*Struct) để truyền địa chỉ, tránh copy và cho phép sửa dữ liệu gốc.
- Truy cập field qua pointer: Go tự dereference (không cần *), nhưng để sửa thì dùng * nếu cần.

Ví dụ áp dụng con trỏ để tối ưu bộ nhớ
```go
type GiangVien struct {
    name   string
    email  string
    gender int
}
  
// Trong Golang. Chúng ta có thể lượt bỏ luôn dấu "*" và cặp "()", khi in. Vì Goalng hiểu được là chúng ta đang truy cập vào giá trị
func HienThiThongTin2(doiTuong *GiangVien) {
    fmt.Printf("Data type: %T\n", doiTuong)
    fmt.Printf("Ho ten: %s\n", doiTuong.name)
    fmt.Printf("Dia chi email: %s\n", doiTuong.email)
    fmt.Printf("Gioi tinh: %d\n", doiTuong.gender)
}

func main() {
  
  //Khuyến khích không sử dụng kiểu này, mặc dù đúng
    Josuke := GiangVien{
        "Jonathan Joestar",
        "Jojo@gmail.com",
        1,
    }

    HienThiThongTin2(&Josuke)
}
```
**Lợi ích:** Tiết kiệm bộ nhớ, đặc biệt với struct lớn chứa array/slice.

----
## **Reciever trong Struct:**

**- **Receiver** là cách gắn hàm (method) vào struct, biến hàm thành "phương thức" của struct.
- Cú pháp: ==func (receiver TênStruct) TênMethod() { ... }==
- Hai loại:
    - **Value Receiver** (gv GiangVien): Làm việc với bản copy, không sửa struct gốc. Dùng cho đọc dữ liệu.
    - **Pointer Receiver** (gv *GiangVien): Làm việc trực tiếp với gốc, có thể sửa. Dùng cho cập nhật.

![](https://www.practical-go-lessons.com/img/method_anatomy.00332211.png)

***Ví dụ:** Thêm methods để hiển thị và xóa thông tin.*
```go
package main
import "fmt"

type GiangVien struct {
    Name   string
    Email  string
    Gender int
}

// Value Receiver: Không sửa gốc
func (gv GiangVien) HienThiThongTin() {
    fmt.Printf("Họ tên: %s\n", gv.Name)
    fmt.Printf("Email: %s\n", gv.Email)
    fmt.Printf("Giới tính: %d\n", gv.Gender)
}

// Pointer Receiver: Sửa gốc
func (gv *GiangVien) Clear() {
    gv.Name = ""
    gv.Email = ""
    gv.Gender = -1
}

func main() {
    gv := GiangVien{
        Name:   "Pham Gia Bao",
        Email:  "bao123@gmail.com",
        Gender: 1,
    }

    gv.HienThiThongTin()  // Gọi method
    // Kết quả: Họ tên: Pham Gia Bao, Email: bao123@gmail.com, Giới tính: 1

    gv.Clear()            // Sửa gốc
    gv.HienThiThongTin()  // Kết quả: Họ tên: , Email: , Giới tính: -1
}
```

Kết quả:
```bash
Data type: *main.GiangVien
Ho ten: Pham Gia Bao
Dia chi email: bao123@gmail.com
Gioi tinh: 1
Data type: *main.GiangVien
Ho ten:
Dia chi email:
Gioi tinh: -1
```

Vậy **Receiver** có 2 loại:
- **Value Receiver:** Không thay đổi dữ liệu, tạo và làm việc với bản sao
- **Pointer Receiver:** Thay đổi dữ liệu, làm việc trực tiếp với struct gốc

---
## **Tìm hiểu về Tag trong Struct:**

- **Tags** là metadata gắn vào field, dưới dạng ==key:"value"== (dùng backtick ``` ).
- Mục đích: Giúp thư viện bên ngoài (như JSON, XML) xử lý struct, ví dụ chuyển struct thành JSON với tên field tùy chỉnh.
- Phổ biến: ==json:"tên_field"== để encode/decode JSON.

![](https://media2.dev.to/dynamic/image/width=1000,height=420,fit=cover,gravity=auto,format=auto/https%3A%2F%2Fgithub.com%2Fkodelint%2Fblog-assets%2Fraw%2Fmain%2Fimages%2F01-Use-Struct-Tags-in-Golang.jpeg)

Hình ảnh minh hoạ khai báo Tag trong Struct:
![](https://raw.githubusercontent.com/fogio-org/vscode-go-struct-tags-syntax-highlight/master/assets/img/preview-3.png)

Ví dụ: Chuyển struct sang JSON với tên field tiếng Việt.
```go
package main
import (
    "encoding/json"
    "fmt"
    "os"
)

type GiangVien struct {
    Name   string `json:"Ho ten"`
    Email  string `json:"Dia chi email"`
    Gender int    `json:"Gioi tinh"`
}

func main() {
    gv := GiangVien{
        Name:   "Pham Gia Bao",
        Email:  "bao123@gmail.com",
        Gender: 1,
    }

    output, err := json.Marshal(gv)  // Chuyển sang JSON
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    fmt.Println(string(output))  // {"Ho ten":"Pham Gia Bao","Dia chi email":"bao123@gmail.com","Gioi tinh":1}
}
```

**Ứng dụng thực tế**: Khi xây dựng API web, tags giúp trả về JSON dễ đọc cho client (web, app mobile).

### Lưu Ý Quan Trọng

- Field bắt đầu bằng chữ hoa → public (exported); chữ thường → private (unexported).
- Struct có thể lồng nhau: Ví dụ, thêm field DiaChi struct { ThanhPho string }.
- Không có kế thừa như OOP, nhưng có thể nhúng struct (embedding) để "kế thừa" field/method.
- Luôn import cần thiết (fmt, json,...).
- Thực hành: Tạo struct cho "SinhVien" với methods để tính điểm trung bình.