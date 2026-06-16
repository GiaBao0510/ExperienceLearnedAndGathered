# Phạm Vi Truy Cập (Scope & Visibility) Trong Go

**Phạm vi (Scope)** của một định danh trong Go xác định vùng mã nguồn mà định danh đó có thể được truy cập. Go sử dụng **phạm vi từ vựng (lexical scoping)** — nghĩa là phạm vi của một định danh được xác định hoàn toàn tại thời điểm biên dịch, dựa trên vị trí khai báo trong mã nguồn, không phụ thuộc vào luồng thực thi.

Go không dùng từ khóa `public` hay `private` như các ngôn ngữ khác. Thay vào đó, **khả năng hiển thị (Visibility)** — tức là định danh có thể truy cập từ package khác hay không — được quyết định bởi **quy tắc viết hoa chữ cái đầu tiên**:

- **Exported (công khai):** Tên bắt đầu bằng **chữ cái viết hoa**. Có thể truy cập từ bất kỳ package nào khác — tương đương `public` trong các ngôn ngữ khác.
- **Unexported (nội bộ):** Tên bắt đầu bằng **chữ cái viết thường**. Chỉ có thể truy cập bên trong package khai báo nó — tương đương `private`.

Quy tắc này áp dụng nhất quán cho **biến, hằng số, hàm, kiểu dữ liệu, và trường của struct**.

---

## Các Cấp Độ Phạm Vi Trong Go

Go có bốn cấp độ phạm vi, được phân biệt bởi vị trí khai báo trong mã nguồn.

### 1. Cấp Độ Package (Biến Toàn Cục)

Định danh được khai báo bên ngoài tất cả các hàm, thường ở đầu file. Phạm vi của nó bao phủ toàn bộ package — tức là có thể truy cập từ bất kỳ file nào thuộc cùng package đó.

- **Viết hoa:** có thể truy cập từ bất kỳ package nào trong dự án (exported).
- **Viết thường:** chỉ truy cập được trong các file thuộc cùng package (unexported).

```go
package main

var SoLuongToiDa int = 1000 // Exported — package khác có thể dùng
var bienNoiBo int = 50      // Unexported — chỉ dùng trong package này
```

### 2. Cấp Độ File (Import)

Các khai báo `import` chỉ có hiệu lực bên trong file hiện tại. Dù hai file thuộc cùng một package, nếu file A đã `import "fmt"` thì file B vẫn phải khai báo `import "fmt"` riêng nếu muốn dùng.

```go
// file_a.go
package main

import "fmt" // chỉ có hiệu lực trong file này

// file_b.go
package main

// import "fmt" // phải khai báo lại nếu muốn dùng fmt ở đây
```

### 3. Cấp Độ Khối / Hàm (Biến Cục Bộ)

Định danh được khai báo bên trong một hàm hoặc khối lệnh (`{}`). Biến chỉ tồn tại và có thể truy cập trong khối đó, không thể dùng từ bên ngoài.

```go
func tinhTong(a, b int) int {
    ketQua := a + b // biến cục bộ — chỉ tồn tại trong hàm này
    return ketQua
}
// fmt.Println(ketQua) // lỗi biên dịch: ketQua không tồn tại ở đây
```

Phạm vi cục bộ còn áp dụng cho các khối lệnh con bên trong hàm:

```go
func kiemTra(n int) {
    if n > 0 {
        thongBao := "dương" // chỉ tồn tại trong khối if
        fmt.Println(thongBao)
    }
    // fmt.Println(thongBao) // lỗi biên dịch: thongBao không tồn tại ở đây
}
```

### 4. Phạm Vi Trong Struct (Trường Exported / Unexported)

Các trường (field) trong struct cũng tuân theo quy tắc viết hoa. Đây là cơ chế đóng gói (encapsulation) chính của Go.

```go
type NhanVien struct {
    Ten        string // Exported — package khác có thể đọc và ghi trực tiếp
    luongCanBan int   // Unexported — chỉ truy cập được trong package khai báo struct này
}
```

---

![](https://st.quantrimang.com/photos/image/2025/02/24/golang-pvi-bien-7.jpg)

---

## Biến Toàn Cục

Biến toàn cục được khai báo bên ngoài tất cả các hàm, có thể truy cập từ bất kỳ hàm nào trong cùng package.

```go
package main

import "fmt"

var soLuong int = 100 // biến toàn cục trong package

func main() {
    fmt.Printf("Số lượng: %d\n", soLuong) // truy cập bình thường
}
```

Biến toàn cục hữu ích khi nhiều hàm cần dùng chung một giá trị, nhưng cần thận trọng vì dễ gây ra lỗi khó theo dõi nếu bị thay đổi ở nhiều nơi.

---

## Ưu Tiên Biến Cục Bộ (Variable Shadowing)

Khi một biến cục bộ có cùng tên với một biến ở phạm vi rộng hơn (toàn cục hoặc khối cha), biến cục bộ sẽ **che khuất (shadow)** biến kia trong phạm vi của nó. Biến ở phạm vi rộng hơn vẫn tồn tại nhưng tạm thời không thể truy cập.

```go
package main

import "fmt"

var soLuong int = 100 // biến toàn cục

func main() {
    var soLuong int = 200 // biến cục bộ — che khuất biến toàn cục
    fmt.Printf("Biến cục bộ được ưu tiên: %d\n", soLuong) // in 200
}
```

Kết quả:

```
Biến cục bộ được ưu tiên: 200
```

Shadowing cũng xảy ra giữa các khối lồng nhau bên trong cùng một hàm:

```go
func demo() {
    x := 10
    fmt.Println(x) // 10

    {
        x := 20    // biến x mới, che khuất x bên ngoài
        fmt.Println(x) // 20
    }

    fmt.Println(x) // 10 — x bên ngoài không bị ảnh hưởng
}
```

---

## Lưu Ý Quan Trọng

**Shadowing vô tình với `:=`:** Toán tử `:=` trong khối con tạo ra biến mới thay vì gán vào biến của khối cha, đây là nguồn gốc của nhiều lỗi logic khó phát hiện:

```go
func layDuLieu() (int, error) {
    gia, err := truyVanDB()
    if err != nil {
        return 0, err
    }

    if gia > 0 {
        gia, err := tinhThue(gia) // TẠO biến gia và err MỚI trong khối if
        _ = err
        fmt.Println(gia)          // dùng biến gia mới trong khối if
    }

    return gia, nil // trả về biến gia BAN ĐẦU — biến trong if không ảnh hưởng
}
```

Cách tránh: dùng `=` thay vì `:=` khi muốn gán vào biến đã tồn tại ở khối cha.

**Exported không có nghĩa là có thể sửa từ bên ngoài:** Trường exported của struct chỉ có thể đọc/ghi từ bên ngoài package nếu bạn có quyền truy cập vào instance của struct đó. Đây là hai khái niệm khác nhau: visibility (ai thấy) và access control (ai được phép thay đổi).

**Không có phạm vi cấp độ file cho biến thông thường:** Không giống một số ngôn ngữ khác, Go không có biến "file-private". Biến khai báo ở cấp package (ngoài hàm) với chữ thường có thể truy cập từ bất kỳ file nào trong cùng package, không bị giới hạn trong một file.

---

## Tổng Kết

| Cấp độ       | Vị trí khai báo         | Phạm vi truy cập                                         |
| ------------ | ----------------------- | -------------------------------------------------------- |
| Package      | Ngoài tất cả hàm        | Toàn bộ package (viết thường) hoặc toàn dự án (viết hoa) |
| File         | Khai báo `import`       | Chỉ trong file hiện tại                                  |
| Hàm / Khối   | Trong hàm hoặc `{}`     | Chỉ trong khối khai báo và các khối con                  |
| Struct field | Trong định nghĩa struct | Theo quy tắc viết hoa — exported hoặc unexported         |
