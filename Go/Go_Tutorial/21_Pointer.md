# Con Trỏ (Pointer) Trong Go

**Con trỏ** là một biến đặc biệt lưu trữ **địa chỉ bộ nhớ** của một biến khác. Nói cách khác, nó "trỏ" đến vị trí dữ liệu trong bộ nhớ, thay vì lưu giá trị trực tiếp. Điều này giúp tiết kiệm tài nguyên và cho phép thay đổi dữ liệu gốc mà không cần copy.

Ví dụ: Với biến `name := "Pham Gia Bao"`:
- **Kiểu dữ liệu**: `string`
- **Giá trị**: `"Pham Gia Bao"`
- **Địa chỉ bộ nhớ**: Một giá trị hex như `0xc0000220c0` (khác nhau tùy máy).

Biến thông thường lưu giá trị, còn con trỏ lưu địa chỉ của biến đó.

---
### Cách Lấy Địa Chỉ Và Tạo Con Trỏ

- Sử dụng `&` để lấy địa chỉ của biến.
- Sử dụng `*` để khai báo kiểu con trỏ (ví dụ: `*string`) và để truy cập giá trị tại địa chỉ (dereference).

**Ví dụ cơ bản**:
```go
package main
import "fmt"

func main() {
    name := "Nguyen Van A"
    fmt.Printf("Kiểu dữ liệu: %T\n", name)
    fmt.Printf("Giá trị: %v\n", name)
    fmt.Printf("Địa chỉ: %p\n", &name)  // Sử dụng %p cho địa chỉ hex

    // Tạo con trỏ
    pName := &name
    fmt.Printf("\nKiểu dữ liệu pName: %T\n", pName)  // *string
    fmt.Printf("Giá trị pName (địa chỉ): %p\n", pName)
    fmt.Printf("Giá trị tại địa chỉ (*pName): %v\n", *pName)  // Dereference
}
```

**Kết quả mẫu** (địa chỉ có thể khác):
```
Kiểu dữ liệu: string
Giá trị: Nguyen Van A
Địa chỉ: 0xc0000220c0

Kiểu dữ liệu pName: *string
Giá trị pName (địa chỉ): 0xc0000220c0
Giá trị tại địa chỉ (*pName): Nguyen Van A
```

---
### Con Trỏ Đến Con Trỏ (Pointer to Pointer)

Bạn có thể tạo con trỏ trỏ đến con trỏ khác, tạo chuỗi cấp độ.

**Ví dụ**:
```go
package main
import "fmt"

func main() {
    name := "Nguyen Van A"
    pName := &name
    pName2 := &pName  // Con trỏ đến con trỏ

    fmt.Printf("Kiểu dữ liệu pName2: %T\n", pName2)  // **string
    fmt.Printf("Giá trị pName2 (địa chỉ của pName): %p\n", pName2)
    fmt.Printf("Giá trị tại *pName2 (địa chỉ của name): %p\n", *pName2)
    fmt.Printf("Giá trị tại **pName2 (giá trị name): %v\n", **pName2)
}
```

**Kết quả mẫu**:
```
Kiểu dữ liệu pName2: **string
Giá trị pName2 (địa chỉ của pName): 0xc000044040
Giá trị tại *pName2 (địa chỉ của name): 0xc0000220c0
Giá trị tại **pName2 (giá trị name): Nguyen Van A
```

---
### Tại Sao Nên Dùng Con Trỏ?

- **Tránh copy dữ liệu lớn**: Go mặc định truyền tham trị (copy giá trị), tốn bộ nhớ với dữ liệu lớn (như struct, array). Con trỏ truyền địa chỉ, không copy.
- **Cập nhật dữ liệu gốc trực tiếp**: Cho phép hàm thay đổi biến ngoài hàm (pass by reference).

**Ví dụ: Pass by value (không dùng pointer)**:
```go
package main
import "fmt"

func updateName(name string) {
    name = "Nguyen Van B"  // Chỉ thay đổi bản copy
    fmt.Printf("Bên trong hàm - Giá trị: %v, Địa chỉ: %p\n", name, &name)
}

func main() {
    name := "Nguyen Van A"
    fmt.Printf("Trước hàm - Giá trị: %v, Địa chỉ: %p\n", name, &name)

    updateName(name)

    fmt.Printf("Sau hàm - Giá trị: %v, Địa chỉ: %p\n", name, &name)  // Không thay đổi
}
```

**Kết quả mẫu** (thấy địa chỉ khác nhau):
```
Trước hàm - Giá trị: Nguyen Van A, Địa chỉ: 0xc0000220c0
Bên trong hàm - Giá trị: Nguyen Van B, Địa chỉ: 0xc0000220e0
Sau hàm - Giá trị: Nguyen Van A, Địa chỉ: 0xc0000220c0
```

**Ví dụ: Pass by reference (dùng pointer)**:
```go
package main
import "fmt"

func updateName(pName *string) {
    *pName = "Nguyen Van B"  // Thay đổi giá trị gốc qua dereference
    fmt.Printf("Bên trong hàm - Giá trị: %v, Địa chỉ: %p\n", *pName, pName)
}

func main() {
    name := "Nguyen Van A"
    fmt.Printf("Trước hàm - Giá trị: %v, Địa chỉ: %p\n", name, &name)

    updateName(&name)  // Truyền địa chỉ

    fmt.Printf("Sau hàm - Giá trị: %v, Địa chỉ: %p\n", name, &name)  // Đã thay đổi
}
```

**Kết quả mẫu** (địa chỉ giống nhau):
```
Trước hàm - Giá trị: Nguyen Van A, Địa chỉ: 0xc0000220c0
Bên trong hàm - Giá trị: Nguyen Van B, Địa chỉ: 0xc0000220c0
Sau hàm - Giá trị: Nguyen Van B, Địa chỉ: 0xc0000220c0
```

Go Garbage Collector sẽ tự dọn bộ nhớ không dùng, nhưng dùng pointer giúp tránh tạo copy thừa.

---
### Ưu Và Nhược Điểm Của Con Trỏ

**Ưu điểm**:
- **Tiết kiệm bộ nhớ**: Không copy dữ liệu lớn khi truyền hàm.
- **Thay đổi linh hoạt**: Cập nhật dữ liệu gốc dễ dàng, hữu ích cho struct/map/slice.
- **Hiệu suất cao**: Trong cấu trúc dữ liệu phức tạp (linked list, tree), pointer giúp liên kết mà không copy.

**Nhược điểm**:
- **Dễ gây lỗi**: Null pointer (nil) có thể gây panic nếu dereference.
- **Khó hiểu hơn**: Người mới dễ nhầm lẫn giữa giá trị và địa chỉ.
- **Side effects**: Thay đổi không mong muốn nếu lạm dụng, làm code khó debug.
- **Không an toàn**: Có thể dẫn đến dangling pointer (trỏ đến bộ nhớ đã giải phóng).


