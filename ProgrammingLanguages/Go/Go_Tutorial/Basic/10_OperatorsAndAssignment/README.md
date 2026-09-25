# Hướng dẫn đầy đủ về Toán tử trong Golang

## Mục lục

1. [Toán tử số học (Arithmetic Operators)](https://claude.ai/chat/0e8dc34c-40e4-4386-9af8-18e0a97b1ced#1-to%C3%A1n-t%E1%BB%AD-s%E1%BB%91-h%E1%BB%8Dc)
2. [Toán tử gán (Assignment Operators)](https://claude.ai/chat/0e8dc34c-40e4-4386-9af8-18e0a97b1ced#2-to%C3%A1n-t%E1%BB%AD-g%C3%A1n)
3. [Toán tử so sánh (Comparison Operators)](https://claude.ai/chat/0e8dc34c-40e4-4386-9af8-18e0a97b1ced#3-to%C3%A1n-t%E1%BB%AD-so-s%C3%A1nh)
4. [Toán tử logic (Logical Operators)](https://claude.ai/chat/0e8dc34c-40e4-4386-9af8-18e0a97b1ced#4-to%C3%A1n-t%E1%BB%AD-logic)
5. [Toán tử bitwise](https://claude.ai/chat/0e8dc34c-40e4-4386-9af8-18e0a97b1ced#5-to%C3%A1n-t%E1%BB%AD-bitwise)

---

## 1. Toán tử số học (Arithmetic Operators)

Toán tử số học được sử dụng để thực hiện các phép tính toán học cơ bản.

### Bảng tổng hợp

| Toán tử | Tên                       | Ví dụ | Mô tả                                     |
| ------- | ------------------------- | ----- | ----------------------------------------- |
| `+`     | Phép cộng                 | x + y | Cộng hai giá trị                          |
| `-`     | Phép trừ                  | x - y | Trừ giá trị thứ hai từ giá trị thứ nhất   |
| `*`     | Phép nhân                 | x * y | Nhân hai giá trị                          |
| `/`     | Phép chia                 | x / y | Chia giá trị thứ nhất cho giá trị thứ hai |
| `%`     | Phép chia lấy dư (modulo) | x % y | Lấy phần dư của phép chia                 |
| `++`    | Phép tăng (increment)     | x++   | Tăng giá trị lên 1 đơn vị                 |
| `--`    | Phép giảm (decrement)     | x--   | Giảm giá trị xuống 1 đơn vị               |

### Ví dụ thực tế: Tính toán điểm trung bình và xếp loại

```go
package main

import "fmt"

func main() {
    // Kịch bản: Tính điểm trung bình của sinh viên
    diemToan := 8.5
    diemLy := 7.0
    diemHoa := 9.0
    soMonHoc := 3
    
    // Tính điểm trung bình
    tongDiem := diemToan + diemLy + diemHoa
    diemTrungBinh := tongDiem / float64(soMonHoc)
    
    fmt.Printf("Tổng điểm: %.2f\n", tongDiem)
    fmt.Printf("Điểm trung bình: %.2f\n", diemTrungBinh)
    
    // Ví dụ với phép chia lấy dư
    soSinhVien := 47
    soNhom := 5
    
    sinhVienMoiNhom := soSinhVien / soNhom      // 9 sinh viên/nhóm
    sinhVienThua := soSinhVien % soNhom         // 2 sinh viên thừa
    
    fmt.Printf("\nChia %d sinh viên thành %d nhóm:\n", soSinhVien, soNhom)
    fmt.Printf("- Mỗi nhóm có: %d sinh viên\n", sinhVienMoiNhom)
    fmt.Printf("- Số sinh viên thừa: %d\n", sinhVienThua)
    
    // Ví dụ với ++ và --
    soTiet := 10
    soTietDaHoc := 0
    
    fmt.Printf("\nBắt đầu học: %d tiết\n", soTiet)
    
    soTietDaHoc++  // Học xong tiết 1
    soTiet--
    fmt.Printf("Sau tiết 1: Đã học %d tiết, còn %d tiết\n", soTietDaHoc, soTiet)
    
    soTietDaHoc++  // Học xong tiết 2
    soTiet--
    fmt.Printf("Sau tiết 2: Đã học %d tiết, còn %d tiết\n", soTietDaHoc, soTiet)
}
```

### Lưu ý quan trọng về phép chia

```go
// Chia số nguyên
a := 15
b := 6
ketQua1 := a / b  // Kết quả: 2 (không phải 2.5)
fmt.Printf("%d / %d = %d\n", a, b, ketQua1)

// Chia số thực
ketQua2 := float64(a) / float64(b)  // Kết quả: 2.5
fmt.Printf("%d / %d = %.2f\n", a, b, ketQua2)
```

**💡 Ghi nhớ:** Khi chia hai số nguyên, Go sẽ trả về số nguyên (phần nguyên). Muốn kết quả thập phân, phải ép kiểu sang `float64`.

---

## 2. Toán tử gán (Assignment Operators)

Toán tử gán được sử dụng để gán giá trị cho biến và kết hợp với các phép toán số học.

### Bảng tổng hợp

|Toán tử|Ví dụ|Tương đương với|Mô tả|
|---|---|---|---|
|`=`|x = 5|x = 5|Gán giá trị|
|`+=`|x += 3|x = x + 3|Cộng rồi gán|
|`-=`|x -= 3|x = x - 3|Trừ rồi gán|
|`*=`|x *= 3|x = x * 3|Nhân rồi gán|
|`/=`|x /= 3|x = x / 3|Chia rồi gán|
|`%=`|x %= 3|x = x % 3|Chia lấy dư rồi gán|

### Ví dụ thực tế: Quản lý điểm danh và điểm số

```go
package main

import "fmt"

func main() {
    // Kịch bản: Theo dõi điểm danh và điểm tích lũy
    
    // 1. Gán cơ bản
    tongDiem := 0
    soBuoiVang := 0
    
    fmt.Println("=== HỌC KỲ 1 ===")
    fmt.Printf("Điểm khởi đầu: %d\n", tongDiem)
    
    // 2. Sử dụng +=
    tongDiem += 10  // Hoàn thành bài tập 1
    fmt.Printf("Sau bài tập 1: %d điểm\n", tongDiem)
    
    tongDiem += 15  // Hoàn thành bài tập 2
    fmt.Printf("Sau bài tập 2: %d điểm\n", tongDiem)
    
    tongDiem += 20  // Thi giữa kỳ
    fmt.Printf("Sau thi giữa kỳ: %d điểm\n", tongDiem)
    
    // 3. Sử dụng -=
    tongDiem -= 5   // Bị trừ điểm vì nộp trễ
    fmt.Printf("Sau khi bị trừ điểm: %d điểm\n", tongDiem)
    
    // 4. Sử dụng *= (tính điểm nhân hệ số)
    diemThucHanh := 8
    heSo := 2
    diemThucHanh *= heSo  // Thực hành có hệ số 2
    fmt.Printf("Điểm thực hành (x%d): %d\n", heSo, diemThucHanh)
    
    // 5. Sử dụng /= (tính trung bình)
    tongDiemKiemTra := 24
    soLanKiemTra := 3
    tongDiemKiemTra /= soLanKiemTra
    fmt.Printf("Điểm trung bình kiểm tra: %d\n", tongDiemKiemTra)
    
    // 6. Sử dụng %=
    tongBuoiHoc := 47
    soBuoiMotTuan := 5
    tongBuoiHoc %= soBuoiMotTuan  // Số buổi lẻ trong tuần cuối
    fmt.Printf("Số buổi học lẻ trong tuần cuối: %d\n", tongBuoiHoc)
}
```

### So sánh cách viết

```go
// Cách viết dài
diem := 80
diem = diem + 10

// Cách viết ngắn gọn (khuyến khích)
diem := 80
diem += 10
```

---

## 3. Toán tử so sánh (Comparison Operators)

Toán tử so sánh được sử dụng để so sánh hai giá trị và trả về kết quả `true` hoặc `false`.

### Bảng tổng hợp

|Toán tử|Tên|Ví dụ|Mô tả|
|---|---|---|---|
|`==`|Bằng|x == y|Kiểm tra hai giá trị có bằng nhau|
|`!=`|Không bằng|x != y|Kiểm tra hai giá trị có khác nhau|
|`>`|Lớn hơn|x > y|Kiểm tra x có lớn hơn y|
|`<`|Nhỏ hơn|x < y|Kiểm tra x có nhỏ hơn y|
|`>=`|Lớn hơn hoặc bằng|x >= y|Kiểm tra x có lớn hơn hoặc bằng y|
|`<=`|Nhỏ hơn hoặc bằng|x <= y|Kiểm tra x có nhỏ hơn hoặc bằng y|

### Ví dụ thực tế: Hệ thống chấm điểm và xếp loại

```go
package main

import "fmt"

func main() {
    // Kịch bản: Hệ thống xếp loại học lực
    diemSinhVien := 8.5
    diemChuan := 5.0
    diemKha := 7.0
    diemGioi := 8.0
    
    // 1. So sánh bằng (==)
    diemToiDa := 10.0
    if diemSinhVien == diemToiDa {
        fmt.Println("Xuất sắc! Điểm tuyệt đối!")
    }
    
    // 2. So sánh không bằng (!=)
    if diemSinhVien != 0 {
        fmt.Printf("Sinh viên có điểm: %.2f\n", diemSinhVien)
    }
    
    // 3. So sánh lớn hơn (>)
    if diemSinhVien > diemChuan {
        fmt.Println("✓ Đậu (điểm > 5.0)")
    }
    
    // 4. So sánh nhỏ hơn (<)
    if diemSinhVien < diemToiDa {
        diemThieu := diemToiDa - diemSinhVien
        fmt.Printf("Còn thiếu %.2f điểm để đạt điểm tối đa\n", diemThieu)
    }
    
    // 5. So sánh lớn hơn hoặc bằng (>=)
    if diemSinhVien >= diemGioi {
        fmt.Println("Xếp loại: Giỏi (>= 8.0)")
    } else if diemSinhVien >= diemKha {
        fmt.Println("Xếp loại: Khá (>= 7.0)")
    }
    
    // 6. So sánh nhỏ hơn hoặc bằng (<=)
    if diemSinhVien <= diemChuan {
        fmt.Println("✗ Rớt - Cần học lại")
    }
    
    // Ví dụ phức tạp: Kiểm tra điều kiện tốt nghiệp
    fmt.Println("\n=== KIỂM TRA ĐIỀU KIỆN TỐT NGHIỆP ===")
    diemTrungBinh := 7.8
    soTinChi := 150
    soTinChiToiThieu := 140
    
    if diemTrungBinh >= 5.0 && soTinChi >= soTinChiToiThieu {
        fmt.Println("✓ Đủ điều kiện tốt nghiệp")
        fmt.Printf("  - Điểm TB: %.2f (>= 5.0)\n", diemTrungBinh)
        fmt.Printf("  - Tín chỉ: %d (>= %d)\n", soTinChi, soTinChiToiThieu)
    } else {
        fmt.Println("✗ Chưa đủ điều kiện tốt nghiệp")
    }
}
```

### Ví dụ về hệ thống kiểm tra điểm danh

```go
func main() {
    soNgayNghi := 3
    soNgayChoPhep := 3
    
    fmt.Println("=== HỆ THỐNG ĐIỂM DANH ===")
    
    if soNgayNghi > soNgayChoPhep {
        fmt.Printf("⚠ Cảnh báo: Đã nghỉ %d ngày (vượt giới hạn %d ngày)\n", 
                   soNgayNghi, soNgayChoPhep)
        fmt.Println("Có thể bị cấm thi!")
    } else if soNgayNghi == soNgayChoPhep {
        fmt.Printf("⚠ Đã nghỉ %d ngày (đúng giới hạn)\n", soNgayNghi)
        fmt.Println("Không được nghỉ thêm!")
    } else {
        conLai := soNgayChoPhep - soNgayNghi
        fmt.Printf("✓ Đã nghỉ %d ngày, còn được nghỉ %d ngày\n", 
                   soNgayNghi, conLai)
    }
}
```

---

## 4. Toán tử logic (Logical Operators)

Toán tử logic được sử dụng để kết hợp nhiều điều kiện. Kết quả trả về là `true` hoặc `false`.

### Bảng tổng hợp

| Toán tử | Tên            | Mô tả                                            | Ví dụ           |
| ------- | -------------- | ------------------------------------------------ | --------------- |
| `&&`    | AND (Và)       | Trả về `true` nếu **tất cả** điều kiện đều đúng  | x > 5 && x < 10 |
| `\|`    | OR (Hoặc)      | Trả về `true` nếu **ít nhất một** điều kiện đúng | x < 5 \| x > 10 |
| `!`     | NOT (Phủ định) | Đảo ngược kết quả (true → false, false → true)   | !(x > 5)        |

### Bảng chân trị

#### Toán tử AND (&&)

|A|B|A && B|
|---|---|---|
|true|true|true|
|true|false|false|
|false|true|false|
|false|false|false|

#### Toán tử OR (||)

|A|B|A \| B|
|---|---|---|
|true|true|true|
|true|false|true|
|false|true|true|
|false|false|false|

#### Toán tử NOT (!)

|A|!A|
|---|---|
|true|false|
|false|true|

### Ví dụ thực tế 1: Kiểm tra điều kiện nhận học bổng

```go
package main

import "fmt"

func main() {
    fmt.Println("=== HỆ THỐNG XÉT HỌC BỔNG ===\n")
    
    // Thông tin sinh viên
    tenSV := "Nguyễn Văn A"
    diemTrungBinh := 8.5
    diemRenLuyen := 90
    soNgayNghi := 2
    
    // Điều kiện nhận học bổng
    diemTBToiThieu := 8.0
    diemRLToiThieu := 80
    ngayNghiChoPhep := 3
    
    // 1. Toán tử AND (&&) - TẤT CẢ điều kiện phải đúng
    fmt.Println("--- Học bổng Khuyến khích học tập ---")
    if diemTrungBinh >= diemTBToiThieu && diemRenLuyen >= diemRLToiThieu && soNgayNghi <= ngayNghiChoPhep {
        fmt.Printf("✓ %s đủ điều kiện nhận học bổng!\n", tenSV)
        fmt.Printf("  - Điểm TB: %.2f (>= %.2f) ✓\n", diemTrungBinh, diemTBToiThieu)
        fmt.Printf("  - Điểm rèn luyện: %d (>= %d) ✓\n", diemRenLuyen, diemRLToiThieu)
        fmt.Printf("  - Số ngày nghỉ: %d (<= %d) ✓\n", soNgayNghi, ngayNghiChoPhep)
    } else {
        fmt.Printf("✗ %s chưa đủ điều kiện\n", tenSV)
    }
    
    // 2. Toán tử OR (||) - CHỈ CẦN MỘT điều kiện đúng
    fmt.Println("\n--- Học bổng Đặc biệt ---")
    diemOlympic := true
    laNguoiDanToc := false
    laConLietSy := false
    
    if diemOlympic || laNguoiDanToc || laConLietSy {
        fmt.Println("✓ Đủ điều kiện nhận học bổng đặc biệt!")
        if diemOlympic {
            fmt.Println("  - Lý do: Đạt giải Olympic ✓")
        }
        if laNguoiDanToc {
            fmt.Println("  - Lý do: Là người dân tộc ✓")
        }
        if laConLietSy {
            fmt.Println("  - Lý do: Là con liệt sỹ ✓")
        }
    }
    
    // 3. Toán tử NOT (!) - Phủ định
    fmt.Println("\n--- Kiểm tra vi phạm ---")
    coViPham := false
    
    if !coViPham {
        fmt.Println("✓ Sinh viên không có vi phạm")
    } else {
        fmt.Println("✗ Sinh viên có vi phạm - Không được nhận học bổng")
    }
}
```

### Ví dụ thực tế 2: Hệ thống đăng ký môn học

```go
func main() {
    fmt.Println("=== HỆ THỐNG ĐĂNG KÝ MÔN HỌC ===\n")
    
    // Thông tin môn học
    tenMonHoc := "Lập trình nâng cao"
    soSinhVienHienTai := 35
    soChoToiDa := 40
    
    // Thông tin sinh viên
    daNopHocPhi := true
    daDKMonTienQuyet := true
    namHoc := 2
    
    // Kiểm tra điều kiện đăng ký
    fmt.Printf("Môn: %s\n", tenMonHoc)
    fmt.Printf("Số chỗ: %d/%d\n\n", soSinhVienHienTai, soChoToiDa)
    
    // Điều kiện phức tạp: Kết hợp AND và OR
    conCho := soSinhVienHienTai < soChoToiDa
    duDieuKienHocVu := daNopHocPhi && daDKMonTienQuyet
    duNamHoc := namHoc >= 2
    
    if conCho && duDieuKienHocVu && duNamHoc {
        fmt.Println("✓ ĐĂNG KÝ THÀNH CÔNG!")
        fmt.Printf("  - Còn chỗ: %t ✓\n", conCho)
        fmt.Printf("  - Đã nộp học phí: %t ✓\n", daNopHocPhi)
        fmt.Printf("  - Đã ĐK môn tiên quyết: %t ✓\n", daDKMonTienQuyet)
        fmt.Printf("  - Năm học đủ điều kiện: %t ✓\n", duNamHoc)
    } else {
        fmt.Println("✗ KHÔNG THỂ ĐĂNG KÝ!")
        if !conCho {
            fmt.Println("  - Lý do: Hết chỗ")
        }
        if !daNopHocPhi {
            fmt.Println("  - Lý do: Chưa nộp học phí")
        }
        if !daDKMonTienQuyet {
            fmt.Println("  - Lý do: Chưa học môn tiên quyết")
        }
        if !duNamHoc {
            fmt.Println("  - Lý do: Chưa đủ năm học")
        }
    }
}
```

### Ví dụ thực tế 3: Kiểm tra tài khoản thư viện

```go
func main() {
    fmt.Println("=== KIỂM TRA MƯỢN SÁCH THƯ VIỆN ===\n")
    
    tenSV := "Trần Thị B"
    soSachDangMuon := 3
    soSachToiDa := 5
    coSachQuaHan := false
    tienPhat := 0
    
    // Kiểm tra điều kiện mượn sách
    chuaQuaSoSach := soSachDangMuon < soSachToiDa
    khongCoSachQuaHan := !coSachQuaHan
    daDongPhat := tienPhat == 0
    
    fmt.Printf("Sinh viên: %s\n", tenSV)
    fmt.Printf("Đang mượn: %d/%d cuốn\n\n", soSachDangMuon, soSachToiDa)
    
    // Điều kiện: Chưa quá số sách AND không có sách quá hạn AND đã đóng phạt
    if chuaQuaSoSach && khongCoSachQuaHan && daDongPhat {
        soSachConDuocMuon := soSachToiDa - soSachDangMuon
        fmt.Println("✓ ĐƯỢC PHÉP MƯỢN SÁCH")
        fmt.Printf("  - Còn được mượn: %d cuốn\n", soSachConDuocMuon)
    } else {
        fmt.Println("✗ KHÔNG ĐƯỢC MƯỢN SÁCH")
        
        if !chuaQuaSoSach {
            fmt.Println("  - Lý do: Đã mượn đủ số sách cho phép")
        }
        if coSachQuaHan {
            fmt.Println("  - Lý do: Có sách quá hạn chưa trả")
        }
        if !daDongPhat {
            fmt.Printf("  - Lý do: Còn nợ tiền phạt %d VNĐ\n", tienPhat)
        }
    }
}
```

### Mẹo ghi nhớ toán tử logic

- **AND (&&)**: Nghĩ như "cả hai" - Cả hai điều kiện phải đúng
    - Ví dụ: "Để tốt nghiệp cần điểm cao **VÀ** đủ tín chỉ"
- **OR (||)**: Nghĩ như "một trong hai" - Chỉ cần một điều kiện đúng
    - Ví dụ: "Được miễn thi nếu điểm cao **HOẶC** đạt giải Olympic"
- **NOT (!)**: Nghĩ như "ngược lại" - Đảo ngược kết quả
    - Ví dụ: "Nếu **KHÔNG** có vi phạm thì được nhận học bổng"

---

## 5. Toán tử Bitwise (Nâng cao)

Toán tử bitwise làm việc ở mức bit (0 và 1). Thường dùng trong lập trình hệ thống, mã hóa, hoặc tối ưu hóa.

### Bảng tổng hợp

|Toán tử|Tên|Mô tả|Ví dụ|
|---|---|---|---|
|`&`|AND|Bit AND|x & y|
|`\|`|OR|Bit OR|x \| y|
|`^`|XOR|Bit XOR (exclusive OR)|x ^ y|
|`<<`|Left shift|Dịch bit sang trái|x << 2|
|`>>`|Right shift|Dịch bit sang phải|x >> 2|
|`&^`|AND NOT|Bit clear (AND NOT)|x &^ y|

### Ví dụ đơn giản

```go
func main() {
    a := 12  // Binary: 1100
    b := 10  // Binary: 1010
    
    fmt.Printf("a = %d (binary: %b)\n", a, a)
    fmt.Printf("b = %d (binary: %b)\n", b, b)
    fmt.Printf("a & b = %d (binary: %b)\n", a&b, a&b)   // 1000 = 8
    fmt.Printf("a | b = %d (binary: %b)\n", a|b, a|b)   // 1110 = 14
    fmt.Printf("a ^ b = %d (binary: %b)\n", a^b, a^b)   // 0110 = 6
    fmt.Printf("a << 1 = %d (binary: %b)\n", a<<1, a<<1) // 11000 = 24
    fmt.Printf("a >> 1 = %d (binary: %b)\n", a>>1, a>>1) // 110 = 6
}
```

---

## Tổng kết và So sánh

### Thứ tự ưu tiên của toán tử (từ cao đến thấp)

1. `()` - Dấu ngoặc
2. `++`, `--`, `!` - Unary operators
3. `*`, `/`, `%` - Phép nhân, chia
4. `+`, `-` - Phép cộng, trừ
5. `<<`, `>>` - Shift
6. `<`, `<=`, `>`, `>=` - So sánh
7. `==`, `!=` - Bằng/không bằng
8. `&&` - AND logic
9. `||` - OR logic
10. `=`, `+=`, `-=`, ... - Gán

### Ví dụ về thứ tự ưu tiên

```go
ketQua := 5 + 3 * 2        // 11 (không phải 16)
// Vì * có ưu tiên cao hơn +

ketQua = (5 + 3) * 2       // 16
// Sử dụng () để thay đổi thứ tự

ketQua = 10 > 5 && 20 < 30 // true
// So sánh trước, logic sau
```

---

## Bài tập thực hành

### Bài 1: Viết ch