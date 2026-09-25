# Hướng Dẫn Mã Hóa Mật Khẩu Bằng Bcrypt Trong Go

## 1. Giới Thiệu

### 1.1. Mục Tiêu

File `go_bcrypt.go` triển khai một **wrapper (lớp bọc)** đóng gói thư viện `bcrypt` của Go thành một struct tiện dụng, cung cấp hai chức năng cốt lõi:

- **Băm mật khẩu (hash):** Chuyển mật khẩu dạng plaintext thành chuỗi hash an toàn để lưu vào cơ sở dữ liệu.
- **Xác minh mật khẩu (verify):** So sánh mật khẩu người dùng nhập vào với chuỗi hash đã lưu, trả về kết quả khớp hay không.

Đây là thao tác bắt buộc trong bất kỳ hệ thống nào có tính năng đăng ký và đăng nhập người dùng.

### 1.2. Tại Sao Không Lưu Mật Khẩu Trực Tiếp?

Lưu mật khẩu dạng plaintext (văn bản thô) vào database là lỗi bảo mật nghiêm trọng. Nếu database bị rò rỉ, toàn bộ mật khẩu người dùng bị lộ ngay lập tức.

Giải pháp là lưu **hash của mật khẩu** thay vì mật khẩu gốc. Khi đăng nhập, hệ thống băm mật khẩu người dùng nhập vào và so sánh với hash đã lưu — không bao giờ so sánh trực tiếp với mật khẩu gốc vì mật khẩu gốc không được lưu ở bất kỳ đâu.

### 1.3. Tại Sao Dùng Bcrypt Thay Vì MD5 Hay SHA?

MD5 và SHA (SHA-1, SHA-256...) là các thuật toán băm **tốc độ cao** — thiết kế để xử lý nhanh dữ liệu lớn. Đặc tính này lại là điểm yếu khi dùng để bảo vệ mật khẩu: kẻ tấn công có thể thử hàng tỷ mật khẩu mỗi giây (brute force).

Bcrypt được thiết kế đặc biệt để **chậm có chủ đích** thông qua tham số `cost` — mỗi lần tăng cost lên 1, thời gian tính toán tăng gấp đôi. Điều này khiến brute force trở nên không khả thi trong thực tế.

Ngoài ra, bcrypt tự động thêm **salt ngẫu nhiên** vào mỗi lần băm — đảm bảo cùng một mật khẩu sẽ cho ra các hash khác nhau mỗi lần, chống lại tấn công rainbow table (bảng tra cứu hash có sẵn).

---

### 1.4. Thư Viện Sử Dụng

#### `golang.org/x/crypto/bcrypt`

- **Nguồn gốc:** Thuộc nhóm thư viện mở rộng chính thức của Go (`golang.org/x`), được duy trì bởi đội ngũ Go. Không có trong standard library nhưng được coi là thư viện tin cậy cấp chuẩn.
- **Chức năng chính:** Triển khai thuật toán bcrypt để băm và xác minh mật khẩu. Tự động xử lý việc tạo salt, nhúng salt vào chuỗi hash, và kiểm soát độ phức tạp qua tham số cost.
- **Lý do chọn:** Đây là lựa chọn tiêu chuẩn trong cộng đồng Go để bảo vệ mật khẩu, được sử dụng rộng rãi trong môi trường production.

---

## 2. Cài Đặt Môi Trường Và Thư Viện

### 2.1. Yêu Cầu

- Go phiên bản 1.16 trở lên.

### 2.2. Khởi Tạo Module

```bash
mkdir go-bcrypt-demo && cd go-bcrypt-demo
go mod init go-bcrypt-demo
```

### 2.3. Cài Đặt Thư Viện

```bash
go get golang.org/x/crypto/bcrypt
```

Sau khi chạy lệnh trên, file `go.mod` sẽ được cập nhật tự động:

```
module go-bcrypt-demo

go 1.21

require golang.org/x/crypto v0.17.0
```

### 2.4. Cấu Trúc Thư Mục Gợi Ý

```
go-bcrypt-demo/
├── go.mod
├── go.sum
├── gobcrypt/
│   └── go_bcrypt.go    ← file source chính
└── main.go             ← file sử dụng package gobcrypt
```

---

## 3. Giải Thích Chi Tiết Code

### 3.1. Toàn Bộ Code

```go
package gobcrypt

import "golang.org/x/crypto/bcrypt"

type Bcrypt struct {
    Password string
    Cost     int
}

// Hàm khởi tạo
func NewBcrypt(obj Bcrypt) *Bcrypt {
    Cost := 14

    // Phạm vi của cost là từ 4 đến 31, nếu không hợp lệ sẽ sử dụng giá trị mặc định
    if obj.Cost < 4 || obj.Cost > 31 {
        obj.Cost = Cost
    }

    return &Bcrypt{
        Password: obj.Password,
        Cost:     obj.Cost,
    }
}

// Hàm băm mật khẩu
func (b *Bcrypt) HashPassword() (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(b.Password), b.Cost)
    return string(bytes), err
}

// Hàm so sánh mật khẩu đã băm với mật khẩu gốc
func (b *Bcrypt) CheckPasswordHash(Password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(Password))
    return err == nil
}
```

---

### 3.2. Khai Báo Package Và Struct

```go
package gobcrypt   // đây là package có thể import vào project khác

import "golang.org/x/crypto/bcrypt"   // import thư viện bcrypt

// Bcrypt là struct đóng gói hai thông tin cần thiết để băm mật khẩu
type Bcrypt struct {
    Password string   // mật khẩu plaintext cần băm
    Cost     int      // độ phức tạp của thuật toán (4 đến 31)
}
```

Struct `Bcrypt` đóng vai trò như một "hộp công cụ" — chứa đủ thông tin và cung cấp các phương thức để thao tác với mật khẩu.

---

### 3.3. Hàm Khởi Tạo `NewBcrypt`

```go
func NewBcrypt(obj Bcrypt) *Bcrypt {
    Cost := 14   // giá trị cost mặc định nếu người dùng truyền vào không hợp lệ

    // Validate cost: bcrypt chỉ chấp nhận cost từ 4 đến 31
    // Nếu nằm ngoài phạm vi, gán về giá trị mặc định 14
    if obj.Cost < 4 || obj.Cost > 31 {
        obj.Cost = Cost
    }

    // Trả về con trỏ đến Bcrypt struct đã được khởi tạo hợp lệ
    return &Bcrypt{
        Password: obj.Password,
        Cost:     obj.Cost,
    }
}
```

**Mục đích của hàm constructor `NewBcrypt`:**

Go không có constructor tích hợp sẵn như các ngôn ngữ OOP. Pattern `NewXxx` là quy ước trong Go để tạo một hàm khởi tạo kiểm soát quá trình tạo đối tượng — đảm bảo struct luôn ở trạng thái hợp lệ trước khi sử dụng.

Nếu không có hàm này, người dùng có thể vô tình truyền `Cost = 0` hoặc `Cost = 100` — các giá trị mà thư viện bcrypt không chấp nhận và sẽ trả về lỗi.

**Tại sao trả về con trỏ (`*Bcrypt`) thay vì giá trị (`Bcrypt`)?**

Trả về con trỏ giúp tránh sao chép toàn bộ struct khi truyền qua các hàm, đặc biệt quan trọng khi struct lớn. Đây cũng là quy ước phổ biến với constructor trong Go.

**Ý nghĩa của tham số `cost`:**

|Cost|Thời gian băm (ước tính)|Khuyến nghị|
|---|---|---|
|4|Dưới 1ms|Chỉ dùng trong unit test|
|10|Khoảng 100ms|Mức tối thiểu cho production|
|12|Khoảng 400ms|Khuyến nghị cho hầu hết hệ thống|
|14|Khoảng 1.5 giây|Mặc định trong code này — bảo mật cao|
|16|Khoảng 6 giây|Cho hệ thống yêu cầu bảo mật rất cao|

Mỗi lần tăng cost thêm 1, thời gian tính toán tăng gấp đôi. Kẻ tấn công phải đối mặt với chi phí tính toán tương tự khi thử từng mật khẩu.

---

### 3.4. Hàm Băm Mật Khẩu `HashPassword`

```go
func (b *Bcrypt) HashPassword() (string, error) {
    // bcrypt.GenerateFromPassword nhận vào:
    //   - []byte(b.Password): mật khẩu dạng byte slice
    //   - b.Cost: độ phức tạp
    // Hàm tự động tạo salt ngẫu nhiên và nhúng vào chuỗi hash kết quả
    bytes, err := bcrypt.GenerateFromPassword([]byte(b.Password), b.Cost)

    // Chuyển kết quả từ []byte sang string để dễ lưu trữ (database, file...)
    return string(bytes), err
}
```

**Cấu trúc chuỗi hash bcrypt trả về:**

```
$2a$14$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy
 │   │  │                            │
 │   │  └── salt (22 ký tự)          └── hash thực sự (31 ký tự)
 │   └── cost = 14
 └── phiên bản bcrypt ($2a$ là phiên bản phổ biến nhất)
```

Toàn bộ thông tin cần thiết để xác minh (phiên bản, cost, salt) được nhúng ngay trong chuỗi hash. Vì vậy bạn chỉ cần lưu một chuỗi duy nhất này vào database.

**Lưu ý quan trọng:** Cùng một mật khẩu, mỗi lần gọi `HashPassword` sẽ cho ra chuỗi hash **khác nhau** do salt được tạo ngẫu nhiên mỗi lần. Đây là hành vi đúng và được thiết kế có chủ đích.

---

### 3.5. Hàm Xác Minh Mật Khẩu `CheckPasswordHash`

```go
// Password: mật khẩu plaintext người dùng vừa nhập (chưa băm)
// hash:     chuỗi hash đã lưu trong database
func (b *Bcrypt) CheckPasswordHash(Password, hash string) bool {
    // bcrypt.CompareHashAndPassword:
    //   - Tự động trích xuất salt và cost từ chuỗi hash
    //   - Băm lại Password với salt và cost đó
    //   - So sánh kết quả với phần hash trong chuỗi
    //   - Trả về nil nếu khớp, trả về error nếu không khớp
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(Password))

    // Chuyển kết quả thành bool: nil = khớp = true, error = không khớp = false
    return err == nil
}
```

**Tại sao không tự băm rồi so sánh chuỗi?**

Nhiều người nhầm tưởng có thể làm như sau:

```go
// SAI — không bao giờ làm vậy
newHash, _ := HashPassword(inputPassword)
return newHash == storedHash   // luôn trả về false dù mật khẩu đúng!
```

Cách này sai vì mỗi lần băm tạo salt mới → hash mới hoàn toàn khác với hash đã lưu. Phải dùng `CompareHashAndPassword` để thư viện tự trích xuất salt từ hash đã lưu và băm lại đúng cách.

**Lưu ý thứ tự tham số:**

```go
bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
//                              ^               ^
//                        hash trước       password sau
```

Đảo ngược thứ tự sẽ gây lỗi runtime.

---

## 4. Ví Dụ Sử Dụng Thực Tế

### 4.1. Sử Dụng Package Trong `main.go`

```go
package main

import (
    "fmt"
    "log"

    "go-bcrypt-demo/gobcrypt"
)

func main() {
    // Khởi tạo với mật khẩu và cost hợp lệ
    b := gobcrypt.NewBcrypt(gobcrypt.Bcrypt{
        Password: "my_secure_password",
        Cost:     12,
    })

    // Băm mật khẩu
    hash, err := b.HashPassword()
    if err != nil {
        log.Fatal("Lỗi khi băm mật khẩu:", err)
    }
    fmt.Println("Hash:", hash)
    // Ví dụ output: $2a$12$...

    // Xác minh mật khẩu đúng
    isValid := b.CheckPasswordHash("my_secure_password", hash)
    fmt.Println("Mật khẩu đúng:", isValid)   // true

    // Xác minh mật khẩu sai
    isValid = b.CheckPasswordHash("wrong_password", hash)
    fmt.Println("Mật khẩu sai:", isValid)    // false
}
```

### 4.2. Tích Hợp Vào Luồng Đăng Ký Và Đăng Nhập

```go
// Luồng đăng ký người dùng
func RegisterUser(username, plainPassword string) error {
    b := gobcrypt.NewBcrypt(gobcrypt.Bcrypt{
        Password: plainPassword,
        Cost:     12,
    })

    hashedPassword, err := b.HashPassword()
    if err != nil {
        return fmt.Errorf("lỗi băm mật khẩu: %w", err)
    }

    // Lưu username và hashedPassword vào database
    // Tuyệt đối KHÔNG lưu plainPassword
    err = db.SaveUser(username, hashedPassword)
    return err
}

// Luồng đăng nhập người dùng
func LoginUser(username, plainPassword string) (bool, error) {
    // Lấy hash đã lưu từ database
    storedHash, err := db.GetUserHash(username)
    if err != nil {
        return false, fmt.Errorf("người dùng không tồn tại: %w", err)
    }

    // Khởi tạo Bcrypt (Password ở đây không dùng cho verify,
    // chỉ cần để tạo struct — Cost cũng không quan trọng khi verify)
    b := gobcrypt.NewBcrypt(gobcrypt.Bcrypt{Cost: 12})

    // Xác minh mật khẩu người dùng nhập vào với hash trong database
    isValid := b.CheckPasswordHash(plainPassword, storedHash)
    return isValid, nil
}
```

---

## 5. Lưu Ý Quan Trọng

### 5.1. Chọn Giá Trị Cost Phù Hợp

- **Cost mặc định trong code là 14** — đây là mức bảo mật cao, phù hợp cho hệ thống xử lý dữ liệu nhạy cảm.
- Với hầu hết ứng dụng web thông thường, **cost 12** là mức cân bằng tốt giữa bảo mật và hiệu năng (khoảng 300-400ms mỗi lần băm).
- **Không dùng cost dưới 10** trong môi trường production.
- Nên đặt cost vào biến cấu hình (config) để dễ điều chỉnh theo từng môi trường mà không cần sửa code.

### 5.2. Độ Dài Mật Khẩu Và Giới Hạn Của Bcrypt

Bcrypt chỉ xử lý tối đa **72 byte** đầu tiên của mật khẩu. Phần vượt quá 72 byte bị bỏ qua hoàn toàn. Điều này có nghĩa là hai mật khẩu chỉ khác nhau ở ký tự thứ 73 trở đi sẽ cho ra hash giống nhau và đều xác minh thành công.

Nếu hệ thống cho phép mật khẩu rất dài, cần lưu ý hoặc kết hợp pre-hashing (băm trước bằng SHA-256) trước khi đưa vào bcrypt.

### 5.3. Cost Lưu Trong Hash — Dễ Nâng Cấp

Vì cost được nhúng vào chuỗi hash, bạn có thể nâng cost lên theo thời gian mà không làm mất hiệu lực các hash cũ:

- Hash cũ vẫn xác minh được vì `CompareHashAndPassword` tự đọc cost từ chuỗi hash.
- Khi người dùng đăng nhập thành công, có thể kiểm tra cost hiện tại của hash và băm lại với cost mới nếu cần.

```go
import "golang.org/x/crypto/bcrypt"

// Kiểm tra hash cũ có cần nâng cấp cost không
func needsRehash(hash string, targetCost int) bool {
    cost, err := bcrypt.Cost([]byte(hash))
    if err != nil {
        return false
    }
    return cost < targetCost
}
```

### 5.4. Xử Lý Lỗi Đúng Cách

```go
hash, err := b.HashPassword()
if err != nil {
    // Đây thường là lỗi hệ thống (ví dụ: không đủ entropy để tạo salt)
    // Không nên tiếp tục — trả lỗi cho caller xử lý
    return fmt.Errorf("không thể băm mật khẩu: %w", err)
}
```

Hàm `CheckPasswordHash` trong code hiện tại trả về `bool` và bỏ qua error. Trong production, nên cân nhắc trả về cả error để phân biệt "mật khẩu sai" với "hash bị hỏng":

```go
// Phiên bản nâng cao hơn
func (b *Bcrypt) CheckPasswordHash(password, hash string) (bool, error) {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    if err == bcrypt.ErrMismatchedHashAndPassword {
        return false, nil   // mật khẩu sai — không phải lỗi hệ thống
    }
    if err != nil {
        return false, err   // lỗi thực sự (hash bị hỏng, định dạng sai...)
    }
    return true, nil
}
```

### 5.5. Không Lưu Mật Khẩu Gốc Bất Kỳ Đâu

- Không log mật khẩu plaintext.
- Không lưu vào database dù chỉ tạm thời.
- Không truyền mật khẩu qua URL (query string).
- Xóa mật khẩu khỏi bộ nhớ sau khi băm xong nếu có thể.

---

## 6. Tổng Kết

|Thành phần|Vai trò|Điểm quan trọng|
|---|---|---|
|`Bcrypt` struct|Đóng gói mật khẩu và cost|Tách biệt cấu hình khỏi logic|
|`NewBcrypt`|Constructor, validate cost|Đảm bảo cost luôn trong phạm vi 4-31|
|`HashPassword`|Băm mật khẩu|Mỗi lần cho hash khác nhau do salt ngẫu nhiên|
|`CheckPasswordHash`|Xác minh mật khẩu|Không so sánh chuỗi trực tiếp — dùng `CompareHashAndPassword`|
|`cost`|Độ phức tạp tính toán|Tăng 1 = chậm gấp đôi; khuyến nghị 12-14 cho production|
|Salt|Tự động thêm vào mỗi hash|Chống rainbow table, đảm bảo hash duy nhất|

> Để tìm hiểu thêm, tham khảo tài liệu chính thức tại: https://pkg.go.dev/golang.org/x/crypto/bcrypt