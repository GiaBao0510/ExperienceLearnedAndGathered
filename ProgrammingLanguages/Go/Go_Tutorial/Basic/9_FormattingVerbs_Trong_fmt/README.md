# Hướng dẫn sử dụng Formatting Verbs trong Go

## Formatting Verbs là gì?

**Formatting Verbs** (còn gọi là Printf Verbs hay Format Specifiers) là các ký hiệu đặc biệt trong Go, giúp bạn định dạng và hiển thị dữ liệu theo ý muốn. Chúng hoạt động như các "chỗ trống" trong chuỗi văn bản, nơi bạn có thể chèn giá trị vào và quyết định cách hiển thị giá trị đó.

**Sử dụng trong các hàm:**

- `fmt.Printf()` - In ra màn hình
- `fmt.Sprintf()` - Trả về chuỗi đã định dạng
- `fmt.Fprintf()` - Ghi vào file hoặc output stream

---

## Các Formatting Verbs phổ biến

### 1. `%v` - Hiển thị giá trị mặc định (Universal)

Dùng khi bạn muốn in giá trị của bất kỳ kiểu dữ liệu nào.

```go
name := "Nguyễn Văn A"
age := 20
height := 1.75

fmt.Printf("Tên: %v, Tuổi: %v, Chiều cao: %v\n", name, age, height)
// Output: Tên: Nguyễn Văn A, Tuổi: 20, Chiều cao: 1.75
```

**Biến thể:**

- `%+v` - Hiển thị cả tên field của struct
- `%#v` - Hiển thị dạng syntax của Go

```go
type SinhVien struct {
    Ten   string
    Tuoi  int
    Diem  float64
}

sv := SinhVien{"Trần Thị B", 21, 8.5}

fmt.Printf("%v\n", sv)   // {Trần Thị B 21 8.5}
fmt.Printf("%+v\n", sv)  // {Ten:Trần Thị B Tuoi:21 Diem:8.5}
fmt.Printf("%#v\n", sv)  // main.SinhVien{Ten:"Trần Thị B", Tuoi:21, Diem:8.5}
```

---

### 2. `%T` - Kiểm tra kiểu dữ liệu (Type)

```go
name := "Lê Văn C"
score := 85
gpa := 3.45
isPassed := true

fmt.Printf("Kiểu dữ liệu của name: %T\n", name)       // string
fmt.Printf("Kiểu dữ liệu của score: %T\n", score)     // int
fmt.Printf("Kiểu dữ liệu của gpa: %T\n", gpa)         // float64
fmt.Printf("Kiểu dữ liệu của isPassed: %T\n", isPassed) // bool
```

---

### 3. `%s` - Hiển thị chuỗi (String)

```go
hoTen := "Phạm Minh D"
khoa := "Công nghệ thông tin"

fmt.Printf("Sinh viên %s thuộc khoa %s\n", hoTen, khoa)
// Output: Sinh viên Phạm Minh D thuộc khoa Công nghệ thông tin
```

**Định dạng nâng cao:**

```go
tenMonHoc := "Lập trình Go"

fmt.Printf("|%s|\n", tenMonHoc)      // |Lập trình Go|
fmt.Printf("|%20s|\n", tenMonHoc)    // |       Lập trình Go| (căn phải, rộng 20 ký tự)
fmt.Printf("|%-20s|\n", tenMonHoc)   // |Lập trình Go       | (căn trái, rộng 20 ký tự)
fmt.Printf("|%.5s|\n", tenMonHoc)    // |Lập t| (chỉ lấy 5 ký tự đầu)
```

---

### 4. `%d` - Số nguyên (Decimal Integer)

```go
soSinhVien := 150
soBuoiHoc := 45
diemThi := 87

fmt.Printf("Lớp có %d sinh viên\n", soSinhVien)
fmt.Printf("Môn học có %d buổi, điểm thi: %d/100\n", soBuoiHoc, diemThi)
```

**Định dạng nâng cao:**

```go
soThuTu := 5

fmt.Printf("%d\n", soThuTu)     // 5
fmt.Printf("%5d\n", soThuTu)    //     5 (căn phải, rộng 5 ký tự)
fmt.Printf("%05d\n", soThuTu)   // 00005 (thêm số 0 ở đầu)
fmt.Printf("%-5d\n", soThuTu)   // 5     (căn trái, rộng 5 ký tự)
```

**Các biến thể:**

- `%b` - Binary (nhị phân)
- `%o` - Octal (bát phân)
- `%x` - Hexadecimal (thập lục phân, chữ thường)
- `%X` - Hexadecimal (thập lục phân, chữ hoa)

```go
soHoc := 255

fmt.Printf("Decimal: %d\n", soHoc)    // 255
fmt.Printf("Binary: %b\n", soHoc)     // 11111111
fmt.Printf("Octal: %o\n", soHoc)      // 377
fmt.Printf("Hex: %x\n", soHoc)        // ff
fmt.Printf("Hex: %X\n", soHoc)        // FF
```

---

### 5. `%f` - Số thực (Floating-point)

```go
diemTrungBinh := 8.756
chieuCao := 1.75
canNang := 65.5

fmt.Printf("Điểm TB: %f\n", diemTrungBinh)     // 8.756000 (mặc định 6 chữ số thập phân)
fmt.Printf("Điểm TB: %.2f\n", diemTrungBinh)   // 8.76 (làm tròn 2 chữ số thập phân)
fmt.Printf("Chiều cao: %.2f mét\n", chieuCao)  // 1.75 mét
fmt.Printf("Cân nặng: %.1f kg\n", canNang)     // 65.5 kg
```

**Định dạng nâng cao:**

```go
gia := 25000.5

fmt.Printf("%f\n", gia)        // 25000.500000
fmt.Printf("%.2f\n", gia)      // 25000.50
fmt.Printf("%10.2f\n", gia)    //   25000.50 (rộng 10 ký tự, 2 số thập phân)
fmt.Printf("%010.2f\n", gia)   // 0025000.50 (thêm số 0 ở đầu)
```

---

### 6. `%t` - Boolean

```go
daHoanThanh := true
coViPham := false

fmt.Printf("Đã hoàn thành bài tập: %t\n", daHoanThanh)  // true
fmt.Printf("Có vi phạm nội quy: %t\n", coViPham)        // false
```

---

### 7. `%c` - Ký tự (Character/Rune)

```go
kyTu := 'A'
maASCII := 65

fmt.Printf("Ký tự: %c\n", kyTu)      // A
fmt.Printf("Ký tự: %c\n", maASCII)   // A (từ mã ASCII)
fmt.Printf("Unicode: %c\n", 'Ω')     // Ω
fmt.Printf("Vietnamese: %c\n", 'Ế')  // Ế
```

---

### 8. `%p` - Con trỏ (Pointer)

```go
tuoi := 21
dienTich := 15.5

fmt.Printf("Địa chỉ bộ nhớ của tuoi: %p\n", &tuoi)        // 0xc0000b4008
fmt.Printf("Địa chỉ bộ nhớ của dienTich: %p\n", &dienTich) // 0xc0000b4010
```

---

### 9. `%%` - Hiển thị ký tự %

```go
tiLeHoanThanh := 95.5

fmt.Printf("Tỉ lệ hoàn thành: %.1f%%\n", tiLeHoanThanh)
// Output: Tỉ lệ hoàn thành: 95.5%
```

---

## Ví dụ tổng hợp: Quản lý thông tin sinh viên

```go
package main

import "fmt"

type SinhVien struct {
    MaSV     string
    HoTen    string
    Tuoi     int
    DiemTB   float64
    DaQua    bool
    SoDienThoai string
}

func main() {
    sv := SinhVien{
        MaSV:     "SV001",
        HoTen:    "Nguyễn Văn An",
        Tuoi:     20,
        DiemTB:   8.75,
        DaQua:    true,
        SoDienThoai: "0987654321",
    }
    
    // In thông tin sinh viên
    fmt.Println("========== THÔNG TIN SINH VIÊN ==========")
    fmt.Printf("Mã sinh viên    : %s\n", sv.MaSV)
    fmt.Printf("Họ và tên       : %-20s\n", sv.HoTen)
    fmt.Printf("Tuổi            : %d tuổi\n", sv.Tuoi)
    fmt.Printf("Điểm trung bình : %.2f/10\n", sv.DiemTB)
    fmt.Printf("Trạng thái      : %t (Đã qua môn: %t)\n", sv.DaQua, sv.DaQua)
    fmt.Printf("Số điện thoại   : %s\n", sv.SoDienThoai)
    
    // Kiểm tra kiểu dữ liệu
    fmt.Println("\n========== KIỂM TRA KIỂU DỮ LIỆU ==========")
    fmt.Printf("Kiểu của MaSV: %T\n", sv.MaSV)
    fmt.Printf("Kiểu của DiemTB: %T\n", sv.DiemTB)
    fmt.Printf("Kiểu của DaQua: %T\n", sv.DaQua)
    
    // Hiển thị toàn bộ struct
    fmt.Println("\n========== HIỂN THỊ STRUCT ==========")
    fmt.Printf("Giá trị: %v\n", sv)
    fmt.Printf("Chi tiết: %+v\n", sv)
    fmt.Printf("Go syntax: %#v\n", sv)
    
    // Tính phần trăm điểm
    phanTramDiem := (sv.DiemTB / 10) * 100
    fmt.Printf("\nĐạt %.0f%% tổng điểm\n", phanTramDiem)
}
```

**Output:**

```
========== THÔNG TIN SINH VIÊN ==========
Mã sinh viên    : SV001
Họ và tên       : Nguyễn Văn An       
Tuổi            : 20 tuổi
Điểm trung bình : 8.75/10
Trạng thái      : true (Đã qua môn: true)
Số điện thoại   : 0987654321

========== KIỂM TRA KIỂU DỮ LIỆU ==========
Kiểu của MaSV: string
Kiểu của DiemTB: float64
Kiểu của DaQua: bool

========== HIỂN THỊ STRUCT ==========
Giá trị: {SV001 Nguyễn Văn An 20 8.75 true 0987654321}
Chi tiết: {MaSV:SV001 HoTen:Nguyễn Văn An Tuoi:20 DiemTB:8.75 DaQua:true SoDienThoai:0987654321}
Go syntax: main.SinhVien{MaSV:"SV001", HoTen:"Nguyễn Văn An", Tuoi:20, DiemTB:8.75, DaQua:true, SoDienThoai:"0987654321"}

Đạt 88% tổng điểm
```

---

## Tips khi sử dụng Formatting Verbs

1. **Sử dụng `%v` khi không chắc chắn:** Nếu bạn không chắc nên dùng verb nào, hãy dùng `%v` - nó hoạt động với mọi kiểu dữ liệu.
2. **Debug với `%+v` và `%#v`:** Rất hữu ích khi debug struct phức tạp.
3. **Làm tròn số thực:** Luôn chỉ định số chữ số thập phân với `%.nf` để tránh hiển thị quá nhiều số 0.
4. **Căn chỉnh văn bản:** Sử dụng `%Ns` hoặc `%-Ns` để tạo bảng đẹp mắt.
5. **Nhớ `%%` cho dấu phần trăm:** Khi muốn hiển thị ký tự %, phải dùng `%%`.

---

## Bảng tóm tắt nhanh

|Verb|Mục đích|Ví dụ|
|---|---|---|
|`%v`|Giá trị mặc định|`fmt.Printf("%v", 123)` → `123`|
|`%T`|Kiểu dữ liệu|`fmt.Printf("%T", 123)` → `int`|
|`%s`|Chuỗi|`fmt.Printf("%s", "Go")` → `Go`|
|`%d`|Số nguyên|`fmt.Printf("%d", 123)` → `123`|
|`%f`|Số thực|`fmt.Printf("%.2f", 3.14159)` → `3.14`|
|`%t`|Boolean|`fmt.Printf("%t", true)` → `true`|
|`%c`|Ký tự|`fmt.Printf("%c", 65)` → `A`|
|`%p`|Con trỏ|`fmt.Printf("%p", &x)` → `0xc000...`|
|`%%`|Ký tự %|`fmt.Printf("100%%")` → `100%`|

---

## Kết luận

Formatting Verbs là công cụ mạnh mẽ giúp bạn kiểm soát cách hiển thị dữ liệu trong Go. Việc thành thạo các verb này sẽ giúp code của bạn rõ ràng hơn và dễ debug hơn. Hãy thực hành với các ví dụ trên và thử nghiệm với dữ liệu của riêng bạn!