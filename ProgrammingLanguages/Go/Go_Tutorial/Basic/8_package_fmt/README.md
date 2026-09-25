**Package fmt** trong ngôn ngữ lập trình Go là một công cụ mạnh mẽ để xử lý việc nhập/xuất dữ liệu và định dạng chuỗi. Nó cung cấp các hàm đơn giản để in ra màn hình, đọc dữ liệu từ người dùng, và tạo chuỗi được định dạng. Dưới đây là hướng dẫn chi tiết về các hàm chính, được phân loại rõ ràng thành các phần: Xuất dữ liệu (Output), Nhập dữ liệu (Input), và Nối chuỗi (String Formatting).
## **Xuất Dữ Liệu (Output)**

Các hàm này dùng để in dữ liệu ra màn hình (stdout). Chúng không trả về giá trị mà chỉ thực hiện việc in trực tiếp.

### fmt.Print()

- **Mô tả**: Hàm này in các đối số ra màn hình mà không tự động thêm dấu xuống dòng (\n) ở cuối, trừ khi bạn chỉ định rõ ràng. Nếu gọi nhiều lần mà không có \n, nội dung sẽ được in liên tục trên cùng một dòng. Đây là hàm cơ bản nhất để xuất dữ liệu đơn giản.
- **Lưu ý**:
    - Các đối số có thể là chuỗi, số, biến, hoặc bất kỳ kiểu dữ liệu nào Go hỗ trợ.
    - Nếu có khoảng trắng giữa các đối số, hàm sẽ tự động thêm khoảng trắng giữa chúng.

ví dụ:
```go
package main
import "fmt"

func main() {
    fmt.Print("Pham ")
    fmt.Print("Gia Bao")
}
```

Kết quả: `Pham Gia Bao (tất cả trên một dòng).`

### fmt.Println()

- **Mô tả**: Tương tự fmt.Print(), nhưng hàm này tự động thêm dấu xuống dòng (\n) ở cuối mỗi lần gọi. Điều này làm cho mỗi lệnh in ra một dòng riêng biệt. Rất hữu ích khi bạn muốn xuất dữ liệu theo dạng danh sách hoặc từng dòng rõ ràng.
- **Lưu ý**:
    - Nếu đối số cuối cùng đã có \n, hàm vẫn thêm một cái nữa, nhưng thường không gây vấn đề.
    - Hàm cũng tự động thêm khoảng trắng giữa các đối số.

Ví dụ:
```go
package main
import "fmt"

func main() {
    fmt.Println("Pham")
    fmt.Println("Gia Bao")
}
```

Kết quả:
```shell
Pham 
Gia Bao
```

Thì lúc này ta thấy mọi thứ điều trên 1 hàng riêng biệt

### fmt.Printf()

- **Mô tả**: Hàm này cho phép định dạng chuỗi trước khi in ra màn hình, sử dụng các placeholder (ký tự đặc biệt như %s cho chuỗi, %d cho số nguyên, %f cho số thực, v.v.). Điều này giúp bạn chèn giá trị biến vào chuỗi một cách linh hoạt. Hàm không tự động thêm dấu xuống dòng ở cuối.
- **Lưu ý**:
    - Cú pháp: fmt.Printf(format, args...), trong đó format là chuỗi định dạng, và args là các giá trị thay thế cho placeholder.
    - Một số placeholder phổ biến: %s (chuỗi), %d (số nguyên), %f (số thực), %v (giá trị mặc định), %t (boolean).
    - Nếu bạn muốn xuống dòng, hãy thêm \n vào chuỗi định dạng.
- **Ví dụ**:
```go
package main
import "fmt"

func main() {
    var firstName = "Gia Bao"
    var lastName = "Pham"
    var age = 23

    fmt.Printf("Hi my name is %s %s. I'm %d years old.\n", lastName, firstName, age)
}
```

Kết quả:
`Hi my name is Pham Gia Bao. I'm 23 years old`

---
## **Nhập**

Các hàm này dùng để đọc dữ liệu từ bàn phím (stdin). Trước khi sử dụng, bạn cần khai báo biến để lưu trữ giá trị nhập vào. Quan trọng: Phải sử dụng toán tử & (địa chỉ) trước biến để hàm có thể ghi dữ liệu vào biến đó (vì Go truyền tham trị, không phải tham chiếu).

### fmt.Scan()

- **Mô tả**: Đọc dữ liệu từ bàn phím và lưu vào biến. Hàm chỉ đọc đến khoảng trắng đầu tiên (space) hoặc dấu xuống dòng, và bỏ qua các khoảng trắng thừa ở đầu. Nếu nhập nhiều từ, nó chỉ lấy từ đầu tiên.
- **Lưu ý**:
    - Phù hợp cho nhập dữ liệu đơn giản như số hoặc từ đơn.
    - Nếu nhập nhiều giá trị, bạn có thể gọi nhiều lần hoặc truyền nhiều biến vào một lệnh.

ví dụ:
```go
package main
import "fmt"

func main() {
    var address string
    fmt.Print("Nhập địa chỉ: ")
    fmt.Scan(&address)  // Chỉ lấy phần trước khoảng trắng đầu tiên

    fmt.Println("Địa chỉ: ", address)
}
```

Kết quả (nếu nhập "Phường An Xuyên"): Địa chỉ: Phường (chỉ lấy "Phường").

### fmt.Scanln()

- **Mô tả**: Tương tự fmt.Scan(), nhưng đọc toàn bộ dòng đến khi gặp dấu xuống dòng (Enter).
- **Lưu ý**:
    - Bỏ qua khoảng trắng thừa ở đầu, nhưng giữ khoảng trắng giữa các từ.
    - Nếu không nhập gì và nhấn Enter, hàm sẽ kết thúc mà không lưu giá trị.

ví dụ:
```go
package main
import "fmt"

func main() {
    var address string
    fmt.Print("Nhập địa chỉ: ")
    fmt.Scanln(&address)  // Chỉ đọc từ đầu tiên sau khoảng trắng

    fmt.Println("Địa chỉ: ", address)
}
```

### fmt.Scanf()

- **Mô tả**: Đọc dữ liệu theo định dạng chỉ định (tương tự fmt.Printf()), sử dụng placeholder như %s, %d, v.v. Hàm đọc đến khi khớp định dạng, và có thể đọc nhiều biến cùng lúc.
- **Lưu ý**:
    - Nếu định dạng không khớp với dữ liệu nhập, hàm sẽ trả về lỗi (có thể kiểm tra bằng giá trị trả về).
    - Không tự động thêm khoảng trắng; bạn phải chỉ định rõ trong định dạng (ví dụ: %s %s cho hai chuỗi cách nhau khoảng trắng).

Ví dụ:
```go
package main
import "fmt"

func main() {
    var address string
    fmt.Print("Nhập địa chỉ: ")
    fmt.Scanf("%s", &address)  // Chỉ đọc một chuỗi

    fmt.Println("Địa chỉ: ", address)
}
```

---
## **Nối Chuỗi (String Formatting)**

### fmt.Sprint()

- **Mô tả**: Nối các đối số thành một chuỗi duy nhất mà không thêm dấu xuống dòng. Tự động thêm khoảng trắng giữa các đối số.
- **Lưu ý**: Trả về chuỗi, không in ra.

- **Ví dụ**:
```go
package main
import "fmt"

func main() {
    msg := fmt.Sprint("Dạo này", " khỏe không?")
    fmt.Println(msg)
}
```

Kết quả: `Dạo này Khỏe không ?`

### fmt.Sprintln()

- **Mô tả**: Tương tự fmt.Sprint(), nhưng tự động thêm dấu xuống dòng (\n) ở cuối chuỗi.
- **Lưu ý**: Nếu bạn in chuỗi này, nó sẽ xuống dòng.
- **Ví dụ**:
```go
msg := fmt.Sprintln("Dạo này", " Khỏe không ?")

fmt.Println(msg)
```

### fmt.Sprintf()

- **Mô tả**: Nối và định dạng chuỗi sử dụng placeholder (tương tự fmt.Printf()), sau đó trả về chuỗi thay vì in ra.
- **Lưu ý**: Rất mạnh mẽ cho việc tạo chuỗi động từ biến.
- **Ví dụ**:

```go 
package main
import "fmt"

func main() {
    var firstName = "Gia Bao"
    var age = 23

    msg := fmt.Sprintf("Tôi tên là %s. Tôi nay được %d tuổi.", firstName, age)
    fmt.Println(msg)
}
```

---
### **Lưu Ý Chung**:

- Package fmt cần được import: import "fmt".
- Xử lý lỗi: Các hàm Scan thường trả về số lượng giá trị đọc thành công và lỗi (error). Trong ví dụ đơn giản, chúng ta bỏ qua, nhưng trong code thực tế, hãy kiểm tra: n, err := fmt.Scan(&var).
- Hiệu suất: Đối với ứng dụng lớn, xem xét sử dụng bufio cho nhập/xuất nhanh hơn.