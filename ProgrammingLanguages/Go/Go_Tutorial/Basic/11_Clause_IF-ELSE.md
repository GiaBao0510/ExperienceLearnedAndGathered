# Câu lệnh điều kiện IF-ELSE trong Go

## Giới thiệu

Câu lệnh điều kiện giúp chương trình ra quyết định dựa trên điều kiện. Ví dụ: "**Nếu** điểm >= 5 **thì** đậu, **nếu không** thì rớt".

---

## 1. IF cơ bản

### Cú pháp

```go
if điều_kiện {
    // Code chạy khi điều kiện đúng
}
```

**Lưu ý:** Dấu `{}` là bắt buộc, không cần `()` bao quanh điều kiện.

### Ví dụ

```go
diem := 7.5

if diem >= 5.0 {
    fmt.Println("✓ Đậu")
}
```

---

## 2. IF-ELSE

### Cú pháp

```go
if điều_kiện {
    // Code khi điều kiện đúng
} else {
    // Code khi điều kiện sai
}
```

**Quan trọng:** `else` phải cùng dòng với `}` của `if`.

### Ví dụ: Kiểm tra chẵn/lẻ

```go
so := 15

if so%2 == 0 {
    fmt.Printf("%d là số CHẴN\n", so)
} else {
    fmt.Printf("%d là số LẺ\n", so)
}
// Output: 15 là số LẺ
```

---

## 3. IF-ELSE IF-ELSE (Nhiều điều kiện)

### Cú pháp

```go
if điều_kiện_1 {
    // Code cho điều kiện 1
} else if điều_kiện_2 {
    // Code cho điều kiện 2
} else {
    // Code khi tất cả điều kiện sai
}
```

**Lưu ý:** Chỉ có **MỘT** khối code được thực thi (kiểm tra từ trên xuống).

### Ví dụ 1: Xếp loại học lực

```go
diemTB := 8.5

if diemTB >= 9.0 {
    fmt.Println("Xuất sắc 🏆")
} else if diemTB >= 8.0 {
    fmt.Println("Giỏi ⭐")
} else if diemTB >= 7.0 {
    fmt.Println("Khá ✓")
} else if diemTB >= 5.0 {
    fmt.Println("Trung bình")
} else {
    fmt.Println("Yếu ✗")
}
// Output: Giỏi ⭐
```

### Ví dụ 2: Tính phí ship

```go
khoangCach := 8.5 // km

if khoangCach <= 3 {
    fmt.Println("Phí ship: 15,000đ")
} else if khoangCach <= 7 {
    fmt.Println("Phí ship: 25,000đ")
} else if khoangCach <= 15 {
    fmt.Println("Phí ship: 40,000đ")
} else {
    fmt.Println("Phí ship: 70,000đ")
}
// Output: Phí ship: 40,000đ
```

---

## 4. IF với Short Statement

Khai báo biến ngay trong `if`. Biến chỉ tồn tại trong khối `if-else`.

### Cú pháp

```go
if biến := giá_trị; điều_kiện {
    // Dùng biến ở đây
}
// Biến không tồn tại ở đây
```

### Ví dụ: Tính điểm tổng kết

```go
diemGiuaKy := 7.5
diemCuoiKy := 8.0

if diemTK := (diemGiuaKy*0.4 + diemCuoiKy*0.6); diemTK >= 5.0 {
    fmt.Printf("✓ Đậu - Điểm: %.2f\n", diemTK)
} else {
    fmt.Printf("✗ Rớt - Điểm: %.2f\n", diemTK)
}
// Output: ✓ Đậu - Điểm: 7.80
```

---

## 5. IF lồng nhau

Đặt `if` bên trong `if` khác (không nên lồng quá 3 tầng).

### Ví dụ: Xét điều kiện học bổng

```go
diemTB := 8.5
diemRenLuyen := 85
soNgayNghi := 1

if diemTB >= 8.0 {
    if diemRenLuyen >= 80 && soNgayNghi <= 2 {
        fmt.Println("✓ Được nhận học bổng")
    } else {
        fmt.Println("✗ Chưa đủ điểm rèn luyện hoặc nghỉ nhiều")
    }
} else {
    fmt.Println("✗ Chưa đủ điểm TB")
}
```

---

## 6. Ví dụ tổng hợp

### Xét học bổng đầy đủ

```go
package main

import "fmt"

func main() {
    tenSV := "Nguyễn Văn A"
    diemTB := 8.7
    diemRL := 88
    ngayNghi := 1
    
    fmt.Printf("Sinh viên: %s\n", tenSV)
    fmt.Printf("Điểm TB: %.1f | RL: %d | Nghỉ: %d\n\n", diemTB, diemRL, ngayNghi)
    
    // Xét học bổng
    if diemTB >= 9.0 && diemRL >= 90 && ngayNghi == 0 {
        fmt.Println("🏆 Học bổng XUẤT SẮC: 5 triệu")
    } else if diemTB >= 8.5 && diemRL >= 85 && ngayNghi <= 1 {
        fmt.Println("⭐ Học bổng GIỎI: 3 triệu")
    } else if diemTB >= 8.0 && diemRL >= 80 && ngayNghi <= 2 {
        fmt.Println("✓ Học bổng KHÁ: 2 triệu")
    } else if diemTB >= 7.0 && diemRL >= 70 && ngayNghi <= 3 {
        fmt.Println("• Học bổng KHUYẾN KHÍCH: 1 triệu")
    } else {
        fmt.Println("✗ Chưa đủ điều kiện")
    }
}
```

**Output:**

```
Sinh viên: Nguyễn Văn A
Điểm TB: 8.7 | RL: 88 | Nghỉ: 1

⭐ Học bổng GIỎI: 3 triệu
```

---

## 7. Lỗi thường gặp

### ❌ Lỗi 1: Quên dấu `{}`

```go
// SAI
if diem >= 5.0
    fmt.Println("Đậu")

// ĐÚNG
if diem >= 5.0 {
    fmt.Println("Đậu")
}
```

### ❌ Lỗi 2: `else` xuống dòng mới

```go
// SAI
if diem >= 5.0 {
    fmt.Println("Đậu")
}
else {
    fmt.Println("Rớt")
}

// ĐÚNG
if diem >= 5.0 {
    fmt.Println("Đậu")
} else {
    fmt.Println("Rớt")
}
```

### ❌ Lỗi 3: Dùng `=` thay vì `==`

```go
// SAI - = là gán
if diem = 5.0 {
    fmt.Println("Điểm là 5")
}

// ĐÚNG - == là so sánh
if diem == 5.0 {
    fmt.Println("Điểm là 5")
}
```

---

## 8. Tips quan trọng

✅ **Ưu tiên điều kiện đơn giản trước**

```go
// TỐT - Kiểm tra điều kiện dừng trước
if diem < 0 {
    fmt.Println("Điểm không hợp lệ")
    return
}
// Tiếp tục xử lý...
```

✅ **Đặt tên biến boolean rõ ràng**

```go
daDongHocPhi := true
duDiemDauVao := false

if daDongHocPhi && duDiemDauVao {
    fmt.Println("Được đăng ký học")
}
```

✅ **Tránh IF lồng quá sâu**

```go
// KHÔNG TỐT
if a {
    if b {
        if c {
            // code
        }
    }
}

// TỐT HƠN - Early return
if !a {
    return
}
if !b {
    return
}
if !c {
    return
}
// code
```


---

## Tóm tắt

|Loại|Khi nào dùng|Ví dụ|
|---|---|---|
|`if`|Kiểm tra 1 điều kiện|`if diem >= 5 { }`|
|`if-else`|Có 2 lựa chọn|`if diem >= 5 { } else { }`|
|`if-else if-else`|Nhiều lựa chọn|Xếp loại học lực|
|Short statement|Cần biến tạm|`if x := a+b; x > 10 { }`|
|Nested if|Điều kiện lồng nhau|Kiểm tra học bổng|

**Nhớ:** `else` luôn cùng dòng với `}`, dùng `==` để so sánh, dùng `{}` cho mọi khối code!