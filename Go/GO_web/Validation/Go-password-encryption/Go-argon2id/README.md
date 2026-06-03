# Hướng Dẫn Mã Hóa Mật Khẩu Bằng Argon2id Trong Go

## 1. Giới Thiệu

### 1.1. Mục Tiêu

File `go_argon2id.go` triển khai package `goargon2id` — một wrapper hoàn chỉnh đóng gói thuật toán **Argon2id** để thực hiện hai chức năng cốt lõi trong bảo mật mật khẩu:

- **Băm mật khẩu (Hash):** Chuyển mật khẩu plaintext thành chuỗi hash theo định dạng PHC chuẩn để lưu vào cơ sở dữ liệu.
- **Xác minh mật khẩu (Verify):** So sánh mật khẩu người dùng nhập vào với chuỗi hash đã lưu, trả về kết quả khớp hay không kèm thông tin lỗi nếu có.

So với bcrypt (thuật toán phổ biến trước đây), Argon2id có ưu thế hơn ở chỗ cho phép kiểm soát **cả ba chiều**: thời gian tính toán, lượng RAM sử dụng, và số luồng song song — khiến tấn công bằng GPU/ASIC tốn kém hơn nhiều.

### 1.2. Tại Sao Dùng Argon2id?

Khi lưu mật khẩu vào database, không bao giờ lưu dạng plaintext. Nếu database bị rò rỉ, toàn bộ mật khẩu người dùng bị lộ ngay lập tức. Giải pháp là lưu **hash của mật khẩu** thay vì mật khẩu gốc.

Tuy nhiên, không phải thuật toán hash nào cũng phù hợp để bảo vệ mật khẩu:

|Thuật toán|Thiết kế cho|Phù hợp bảo vệ mật khẩu?|
|---|---|---|
|MD5, SHA-256|Tốc độ cao, kiểm tra toàn vẹn dữ liệu|Không — quá nhanh, dễ brute-force|
|bcrypt|Chậm có chủ đích|Được — nhưng không kiểm soát RAM|
|Argon2id|Chậm + tốn RAM + song song|Tốt nhất hiện tại theo OWASP 2023|

Argon2id là biến thể kết hợp của hai phiên bản:

- **Argon2i:** kháng side-channel attack (tấn công qua kênh phụ như cache timing).
- **Argon2d:** kháng GPU/ASIC attack bằng cách truy cập RAM phụ thuộc vào dữ liệu.

Argon2id thực hiện một số lượt theo kiểu Argon2i ở đầu (chống side-channel) rồi chuyển sang Argon2d (chống GPU), hưởng lợi từ cả hai.

### 1.3. Thư Viện Sử Dụng

#### `golang.org/x/crypto/argon2`

- **Nguồn gốc:** Thuộc nhóm thư viện mở rộng chính thức của Go (`golang.org/x`), được duy trì bởi đội ngũ Go. Không có trong standard library nhưng được coi là thư viện tin cậy cấp chuẩn.
- **Chức năng chính:** Triển khai thuật toán Argon2 (bao gồm cả biến thể Argon2id) để băm mật khẩu với các tham số kiểm soát đầy đủ.
- **Lý do chọn:** Được OWASP khuyến nghị cho password hashing, hỗ trợ kiểm soát ba chiều (time, memory, parallelism), là lựa chọn tiêu chuẩn trong các hệ thống Go hiện đại.

#### Các package standard library được dùng

|Package|Chức năng trong code|
|---|---|
|`crypto/rand`|Tạo salt ngẫu nhiên an toàn từ nguồn entropy của hệ điều hành|
|`crypto/subtle`|So sánh hai slice byte trong thời gian cố định, chống timing attack|
|`encoding/base64`|Mã hóa salt và hash thành chuỗi Base64 để nhúng vào chuỗi PHC|
|`errors`|Định nghĩa sentinel errors để caller kiểm tra lỗi cụ thể|
|`fmt`|Lắp ráp chuỗi PHC và bọc lỗi với ngữ cảnh|
|`strings`|Tách chuỗi PHC thành các phần để parse|

---

## 2. Cài Đặt Môi Trường Và Thư Viện

### 2.1. Yêu Cầu

- Go phiên bản 1.18 trở lên.

### 2.2. Khởi Tạo Module

```bash
mkdir go-argon2id-demo && cd go-argon2id-demo
go mod init go-argon2id-demo
```

### 2.3. Cài Đặt Thư Viện

```bash
go get golang.org/x/crypto/argon2
```

File `go.mod` sau khi cài đặt:

```
module go-argon2id-demo

go 1.21

require golang.org/x/crypto v0.17.0
```

### 2.4. Cấu Trúc Thư Mục Gợi Ý

```
go-argon2id-demo/
├── go.mod
├── go.sum
├── goargon2id/
│   └── go_argon2id.go    <- package chính
└── main.go               <- file sử dụng package
```

---

## 3. Kiến Trúc Tổng Thể Của Package

Trước khi đi vào từng hàm, hãy nắm bức tranh tổng thể của package:

```
goargon2id package
│
├── Argon2Params          -- struct chứa tham số cấu hình (time, memory, threads, keylen)
├── DefaultArgon2Params() -- trả về tham số mặc định theo OWASP
│
├── Hasher                -- struct chính thực hiện băm và xác minh
├── NewHasher(params)     -- constructor tạo Hasher
│
├── Hasher.Hash(password)              -- băm mật khẩu → chuỗi PHC
└── Hasher.Verify(storedHash, password)-- xác minh mật khẩu với chuỗi PHC
         │
         └── parseHash(encoded)        -- hàm nội bộ: tách chuỗi PHC thành các phần
```

Luồng dữ liệu khi băm:

```
plaintext password
       │
       ▼
  generateSalt() → salt ngẫu nhiên (16 byte)
       │
       ▼
  argon2.IDKey(password, salt, timeCost, memoryCost, threads, keyLen)
       │
       ▼
  raw hash bytes (32 byte)
       │
       ▼
  base64(salt) + base64(hash) → chuỗi PHC
       │
       ▼
  "$argon2id$v=19$m=65536,t=2,p=1$<salt_b64>$<hash_b64>"
```

---

## 4. Giải Thích Chi Tiết Code

### 4.1. Định Nghĩa Sentinel Errors

```go
var (
    ErrInvalidHashFormat    = errors.New("invalid hash format")
    ErrUnsupportedAlgorithm = errors.New("unsupported algorithm: expected argon2id")
    ErrIncompatibleVersion  = errors.New("incompatible argon2 version")
    ErrInvalidParams        = errors.New("invalid argon2 parameters")
    ErrEmptyPassword        = errors.New("password must not be empty")
)
```

**Sentinel errors** là các biến lỗi được định nghĩa sẵn ở cấp package, dùng làm "mã lỗi" để caller kiểm tra loại lỗi cụ thể bằng `errors.Is()`.

Ví dụ sử dụng ở tầng gọi hàm:

```go
_, err := hasher.Verify(storedHash, password)
if errors.Is(err, goargon2id.ErrInvalidHashFormat) {
    // xử lý riêng khi hash bị hỏng định dạng
}
if errors.Is(err, goargon2id.ErrIncompatibleVersion) {
    // xử lý riêng khi phiên bản không khớp
}
```

So sánh với cách truyền thống chỉ trả về `error` dạng chuỗi: sentinel errors cho phép code xử lý lỗi **có cấu trúc**, không phải so sánh chuỗi lỗi — vốn dễ sai và khó maintain.

---

### 4.2. Hằng Số Mặc Định

```go
const (
    defaultTimeCost   uint32 = 2          // 2 lần lặp
    defaultMemoryCost uint32 = 64 * 1024  // 64 MiB (đơn vị KiB)
    defaultThreads    uint8  = 1          // 1 luồng
    defaultKeyLength  uint32 = 32         // 32 byte = 256-bit hash
    defaultSaltLength uint32 = 16         // 16 byte = 128-bit salt
)
```

Các giá trị này tuân theo **khuyến nghị OWASP 2023** cho Argon2id:

|Tham số|Giá trị|Lý do|
|---|---|---|
|TimeCost|2|Cân bằng giữa tốc độ và bảo mật|
|MemoryCost|64 MiB|Tốn RAM → GPU/ASIC khó tấn công hàng loạt|
|Threads|1|Tránh phụ thuộc vào số lõi CPU của môi trường triển khai|
|KeyLength|32 byte|256-bit — đủ bảo mật cho hầu hết ứng dụng|
|SaltLength|16 byte|128-bit — đủ entropy để đảm bảo tính duy nhất|

---

### 4.3. Struct `Argon2Params` Và `DefaultArgon2Params`

```go
type Argon2Params struct {
    TimeCost   uint32 // số vòng lặp — càng cao càng chậm
    MemoryCost uint32 // RAM sử dụng (đơn vị KiB)
    Threads    uint8  // số luồng song song
    KeyLength  uint32 // độ dài hash đầu ra (byte)
}

func DefaultArgon2Params() *Argon2Params {
    return &Argon2Params{
        TimeCost:   defaultTimeCost,
        MemoryCost: defaultMemoryCost,
        Threads:    defaultThreads,
        KeyLength:  defaultKeyLength,
    }
}
```

`Argon2Params` là **"config thuần túy"** — chỉ chứa tham số cấu hình, không chứa dữ liệu runtime như salt hay hash. Thiết kế này cho phép tái sử dụng cùng một bộ tham số cho nhiều lần băm khác nhau mà không cần khởi tạo lại.

Lưu ý: `SaltLength` không có trong `Argon2Params` vì salt chỉ cần thiết trong quá trình tạo hash (runtime), không phải là tham số cấu hình thuật toán.

---

### 4.4. Struct `Hasher` Và Constructor `NewHasher`

```go
type Hasher struct {
    params Argon2Params // trường private — chỉ truy cập được trong package
}

func NewHasher(params Argon2Params) *Hasher {
    return &Hasher{
        params: params,
    }
}
```

`Hasher` là đối tượng trung tâm thực hiện mọi thao tác. Trường `params` được khai báo **lowercase (private)** — caller không thể truy cập trực tiếp từ bên ngoài package, chỉ tương tác qua các phương thức `Hash` và `Verify`.

Pattern `NewXxx` là quy ước constructor trong Go. Trả về con trỏ `*Hasher` để tránh sao chép struct mỗi khi truyền qua hàm, và để cho phép phương thức nhận pointer receiver hoạt động đúng.

---

### 4.5. Hàm Tạo Salt `generateSalt`

```go
func generateSalt(saltLen uint32) ([]byte, error) {
    salt := make([]byte, saltLen) // cấp phát slice byte có độ dài saltLen

    // rand.Read điền vào slice bằng các byte ngẫu nhiên từ OS (CSPRNG)
    if _, err := rand.Read(salt); err != nil {
        return nil, fmt.Errorf("failed to generate salt: %w", err)
    }
    return salt, nil
}
```

`crypto/rand.Read` sử dụng **CSPRNG (Cryptographically Secure Pseudo-Random Number Generator)** của hệ điều hành — trên Linux là `/dev/urandom`, trên Windows là `CryptGenRandom`. Đây khác với `math/rand` (random thông thường, không an toàn cho mục đích mật mã).

Salt ngẫu nhiên đảm bảo rằng hai người dùng có cùng mật khẩu sẽ có hash khác nhau, và cùng một người dùng hash hai lần cũng cho kết quả khác nhau — chống tấn công rainbow table.

`fmt.Errorf("failed to generate salt: %w", err)`: ký tự `%w` bọc lỗi gốc vào lỗi mới, giữ nguyên thông tin lỗi gốc để caller có thể dùng `errors.Is()` hoặc `errors.As()` để kiểm tra.

---

### 4.6. Hàm Băm Mật Khẩu `Hash`

```go
func (h *Hasher) Hash(password string) (string, error) {
    // Bước 1: validate đầu vào
    if password == "" {
        return "", ErrEmptyPassword
    }

    // Bước 2: tạo salt ngẫu nhiên 16 byte
    salt, err := generateSalt(defaultSaltLength)
    if err != nil {
        return "", fmt.Errorf("Hash: %w", err)
    }

    // Bước 3: thực thi Argon2id
    // argon2.IDKey là hàm băm Argon2id chính thức
    // Tham số theo thứ tự: password, salt, time, memory, threads, keyLen
    hashRaw := argon2.IDKey(
        []byte(password),        // chuyển string sang []byte
        salt,                    // salt ngẫu nhiên vừa tạo
        h.params.TimeCost,       // số vòng lặp
        h.params.MemoryCost,     // RAM sử dụng (KiB)
        h.params.Threads,        // số luồng
        h.params.KeyLength,      // độ dài hash đầu ra
    )

    // Bước 4: mã hóa Base64 để nhúng vào chuỗi PHC
    // RawStdEncoding: Base64 chuẩn không có padding '='
    saltBase64 := base64.RawStdEncoding.EncodeToString(salt)
    hashBase64 := base64.RawStdEncoding.EncodeToString(hashRaw)

    // Bước 5: lắp ráp chuỗi PHC theo định dạng chuẩn
    encoded := fmt.Sprintf(
        "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
        argon2.Version,          // phiên bản argon2 (hiện tại là 19)
        h.params.MemoryCost,     // m=
        h.params.TimeCost,       // t=
        h.params.Threads,        // p= (parallelism)
        saltBase64,
        hashBase64,
    )

    return encoded, nil
}
```

**Chuỗi PHC (Password Hashing Competition) đầu ra trông như thế này:**

```
$argon2id$v=19$m=65536,t=2,p=1$c29tZXJhbmRvbXNhbHQ$aGFzaGVkb3V0cHV0aGVyZQ
│         │    │               │                     │
│         │    │               └── salt (Base64)     └── hash (Base64)
│         │    └── tham số: memory=65536KiB, time=2, parallelism=1
│         └── phiên bản thuật toán
└── loại thuật toán
```

Chuỗi PHC tự chứa đầy đủ thông tin để xác minh sau này: thuật toán, phiên bản, tham số, salt, và hash — chỉ cần lưu một chuỗi duy nhất vào database.

---

### 4.7. Hàm Nội Bộ `parseHash`

Đây là hàm phức tạp nhất trong package — chịu trách nhiệm tách chuỗi PHC thành các thành phần có thể sử dụng.

```go
func parseHash(encoded string) (*parsedHash, error) {

    // Tách chuỗi theo dấu "$"
    // "$argon2id$v=19$m=65536,t=2,p=1$<salt>$<hash>"
    // → ["", "argon2id", "v=19", "m=65536,t=2,p=1", "<salt>", "<hash>"]
    //    [0]  [1]         [2]     [3]                  [4]       [5]
    parts := strings.Split(encoded, "$")

    // Phải có đúng 6 phần tử (index 0 đến 5)
    // Index 0 luôn rỗng vì chuỗi bắt đầu bằng "$"
    if len(parts) != 6 {
        return nil, fmt.Errorf("parseHash: %w (got %d parts)", ErrInvalidHashFormat, len(parts))
    }

    // Index 1 phải là "argon2id"
    if parts[1] != "argon2id" {
        return nil, fmt.Errorf("parseHash: %w", ErrUnsupportedAlgorithm)
    }

    // Parse phiên bản từ "v=19" → version = 19
    var version int
    if n, _ := fmt.Sscanf(parts[2], "v=%d", &version); n != 1 {
        return nil, fmt.Errorf("parseHash: %w (cannot parse version)", ErrInvalidHashFormat)
    }

    // Kiểm tra phiên bản trong hash có khớp với thư viện hiện tại không
    // argon2.Version là hằng số của thư viện (hiện tại = 19)
    if version != argon2.Version {
        return nil, fmt.Errorf("parseHash: %w (hash v%d, library v%d)",
            ErrIncompatibleVersion, version, argon2.Version)
    }

    // Parse tham số từ "m=65536,t=2,p=1"
    // fmt.Sscanf đọc theo format string, gán vào các biến qua con trỏ
    var p Argon2Params
    n, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d",
        &p.MemoryCost, &p.TimeCost, &p.Threads)
    if err != nil || n != 3 {
        return nil, fmt.Errorf("parseHash: %w (cannot parse parameters)", ErrInvalidParams)
    }

    // Giải mã salt từ Base64 → []byte
    salt, err := base64.RawStdEncoding.DecodeString(parts[4])
    if err != nil {
        return nil, fmt.Errorf("parseHash: failed to decode salt: %w", err)
    }

    // Giải mã hash từ Base64 → []byte
    hashRaw, err := base64.RawStdEncoding.DecodeString(parts[5])
    if err != nil {
        return nil, fmt.Errorf("parseHash: failed to decode hash: %w", err)
    }

    // KeyLength suy ra từ độ dài thực tế của hash đã giải mã
    // (không lưu trong chuỗi PHC, nhưng cần khi gọi argon2.IDKey để verify)
    p.KeyLength = uint32(len(hashRaw))

    return &parsedHash{
        params:  p,
        salt:    salt,
        hashRaw: hashRaw,
    }, nil
}
```

---

### 4.8. Hàm Xác Minh Mật Khẩu `Verify`

```go
func (h *Hasher) Verify(storedHash, providedPassword string) (bool, error) {

    // Bước 1: parse chuỗi PHC đã lưu để lấy salt, params, và hash gốc
    ph, err := parseHash(storedHash)
    if err != nil {
        return false, fmt.Errorf("Verify: %w", err)
    }

    // Bước 2: băm lại mật khẩu được cung cấp
    // Dùng ĐÚNG salt và tham số từ chuỗi PHC đã lưu
    // (không dùng h.params — vì params có thể đã thay đổi kể từ khi hash được tạo)
    computedHash := argon2.IDKey(
        []byte(providedPassword),
        ph.salt,             // salt trích xuất từ hash đã lưu
        ph.params.TimeCost,  // tham số trích xuất từ hash đã lưu
        ph.params.MemoryCost,
        ph.params.Threads,
        ph.params.KeyLength,
    )

    // Bước 3: so sánh hash vừa tính với hash đã lưu
    // subtle.ConstantTimeCompare: so sánh trong thời gian cố định
    // (không phụ thuộc vào nội dung) để chống timing attack
    match := subtle.ConstantTimeCompare(ph.hashRaw, computedHash) == 1
    return match, nil
}
```

**Tại sao dùng `subtle.ConstantTimeCompare` thay vì `==` hay `bytes.Equal`?**

So sánh thông thường (`==`, `bytes.Equal`) dừng lại ngay khi gặp byte đầu tiên khác nhau — thời gian thực thi thay đổi tùy theo mức độ giống nhau của hai chuỗi. Kẻ tấn công có thể đo thời gian phản hồi hàng nghìn lần để suy ra từng byte của hash — đây gọi là **timing attack**.

`subtle.ConstantTimeCompare` luôn so sánh đến hết cả hai slice bất kể chúng có khác nhau hay không, đảm bảo thời gian thực thi **không thay đổi** theo nội dung.

**Tại sao dùng `ph.params` thay vì `h.params` khi tính lại hash?**

Khi xác minh, phải dùng **đúng tham số đã được dùng khi tạo hash ban đầu** — các tham số này được trích xuất từ chuỗi PHC đã lưu. Nếu dùng `h.params` (tham số hiện tại của Hasher), và tham số đó đã thay đổi so với lúc hash được tạo, kết quả xác minh sẽ luôn sai dù mật khẩu đúng.

---

## 5. Ví Dụ Sử Dụng Thực Tế

### 5.1. Sử Dụng Cơ Bản

```go
package main

import (
    "fmt"
    "log"

    "go-argon2id-demo/goargon2id"
)

func main() {
    // Tạo Hasher với tham số mặc định OWASP
    hasher := goargon2id.NewHasher(*goargon2id.DefaultArgon2Params())

    // Băm mật khẩu
    hash, err := hasher.Hash("my_secure_password")
    if err != nil {
        log.Fatal("Lỗi khi băm:", err)
    }
    fmt.Println("Hash:", hash)
    // Ví dụ output:
    // $argon2id$v=19$m=65536,t=2,p=1$<salt_base64>$<hash_base64>

    // Xác minh mật khẩu đúng
    match, err := hasher.Verify(hash, "my_secure_password")
    if err != nil {
        log.Fatal("Lỗi khi xác minh:", err)
    }
    fmt.Println("Mật khẩu đúng:", match) // true

    // Xác minh mật khẩu sai
    match, err = hasher.Verify(hash, "wrong_password")
    if err != nil {
        log.Fatal("Lỗi khi xác minh:", err)
    }
    fmt.Println("Mật khẩu sai:", match) // false
}
```

### 5.2. Tích Hợp Vào Luồng Đăng Ký Và Đăng Nhập

```go
var hasher = goargon2id.NewHasher(*goargon2id.DefaultArgon2Params())

// Luồng đăng ký người dùng
func RegisterUser(username, plainPassword string) error {
    hashedPassword, err := hasher.Hash(plainPassword)
    if err != nil {
        return fmt.Errorf("register: không thể băm mật khẩu: %w", err)
    }
    // Lưu username và hashedPassword vào database
    // Tuyệt đối KHÔNG lưu plainPassword
    return db.SaveUser(username, hashedPassword)
}

// Luồng đăng nhập người dùng
func LoginUser(username, plainPassword string) (bool, error) {
    // Lấy hash đã lưu từ database
    storedHash, err := db.GetUserHash(username)
    if err != nil {
        return false, fmt.Errorf("login: không tìm thấy người dùng: %w", err)
    }

    match, err := hasher.Verify(storedHash, plainPassword)
    if err != nil {
        // Phân biệt lỗi cụ thể để xử lý phù hợp
        if errors.Is(err, goargon2id.ErrInvalidHashFormat) {
            return false, fmt.Errorf("login: hash trong database bị hỏng")
        }
        return false, fmt.Errorf("login: lỗi xác minh: %w", err)
    }
    return match, nil
}
```

### 5.3. Nâng Cấp Cost Cho Hash Cũ

Khi cần tăng độ bảo mật (nâng cost), các hash cũ vẫn xác minh được vì tham số được nhúng trong chuỗi PHC. Có thể tự động nâng cấp hash khi người dùng đăng nhập thành công:

```go
func LoginWithRehash(username, plainPassword string) error {
    storedHash, _ := db.GetUserHash(username)

    newHasher := goargon2id.NewHasher(*goargon2id.DefaultArgon2Params())
    match, err := newHasher.Verify(storedHash, plainPassword)
    if err != nil || !match {
        return fmt.Errorf("mật khẩu không đúng")
    }

    // Kiểm tra hash cũ có dùng tham số cũ hơn không
    // (ví dụ: memory thấp hơn hoặc time ít hơn tiêu chuẩn mới)
    // Nếu có → băm lại với tham số mới và cập nhật database
    newHash, err := newHasher.Hash(plainPassword)
    if err == nil {
        db.UpdateUserHash(username, newHash)
    }
    return nil
}
```

---

## 6. Lưu Ý Quan Trọng

### 6.1. Chọn Tham Số Phù Hợp Với Hệ Thống

Tham số mặc định (64 MiB, 2 lần lặp, 1 luồng) phù hợp với hầu hết hệ thống. Khi điều chỉnh, luôn **benchmark trên môi trường production thực tế**:

```go
// Benchmark thời gian băm với tham số hiện tại
import "time"

start := time.Now()
hasher.Hash("benchmark_password")
elapsed := time.Since(start)
fmt.Printf("Thời gian băm: %v\n", elapsed)
// Mục tiêu: 200ms - 1000ms là phù hợp cho web API
```

Nếu server xử lý nhiều request đồng thời, thời gian băm quá dài sẽ ảnh hưởng đến throughput. Cần cân bằng giữa bảo mật và hiệu năng.

### 6.2. Lưu Trữ Toàn Bộ Chuỗi PHC

Chỉ cần lưu một cột kiểu `TEXT` hoặc `VARCHAR(255)` trong database — chuỗi PHC chứa đầy đủ mọi thứ cần để xác minh sau này. Không cần lưu salt và tham số riêng lẻ.

```sql
CREATE TABLE users (
    id            SERIAL PRIMARY KEY,
    username      VARCHAR(100) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL  -- lưu toàn bộ chuỗi PHC ở đây
);
```

### 6.3. Không Log Mật Khẩu Plaintext

```go
// SAI
log.Printf("Đăng nhập: username=%s, password=%s", username, password)

// ĐÚNG
log.Printf("Đăng nhập: username=%s", username)
```

### 6.4. Giới Hạn Độ Dài Mật Khẩu Đầu Vào

Argon2id không có giới hạn độ dài như bcrypt (72 byte). Tuy nhiên, mật khẩu quá dài (hàng MB) có thể gây DoS vì thuật toán tốn tài nguyên. Nên giới hạn độ dài mật khẩu đầu vào ở tầng validation:

```go
const maxPasswordLength = 1024 // byte

func (h *Hasher) Hash(password string) (string, error) {
    if password == "" {
        return "", ErrEmptyPassword
    }
    if len(password) > maxPasswordLength {
        return "", fmt.Errorf("mật khẩu vượt quá độ dài cho phép")
    }
    // ...
}
```

---

## 7. So Sánh Argon2id Và Bcrypt

|Tiêu chí|Bcrypt|Argon2id|
|---|---|---|
|Năm ra đời|1999|2015 (thắng PHC 2015)|
|Kiểm soát thời gian|Có (cost)|Có (timeCost)|
|Kiểm soát RAM|Không|Có (memoryCost)|
|Kiểm soát luồng|Không|Có (threads)|
|Giới hạn mật khẩu|72 byte|Không giới hạn|
|Kháng GPU/ASIC|Trung bình|Cao (nhờ memory-hard)|
|Khuyến nghị OWASP|Được chấp nhận|Ưu tiên hàng đầu|
|Độ phổ biến trong Go|Cao|Tăng nhanh|

---

## 8. Tổng Kết

|Thành phần|Vai trò|Điểm quan trọng|
|---|---|---|
|`Argon2Params`|Cấu hình thuật toán|Tách biệt khỏi dữ liệu runtime|
|`DefaultArgon2Params()`|Tham số mặc định OWASP|Dùng khi không tự cấu hình|
|`Hasher`|Đối tượng thực thi băm/xác minh|Trường `params` private — truy cập qua phương thức|
|`generateSalt`|Tạo salt ngẫu nhiên|Dùng `crypto/rand`, không dùng `math/rand`|
|`Hash`|Băm mật khẩu → chuỗi PHC|Mỗi lần cho hash khác nhau do salt ngẫu nhiên|
|`parseHash`|Tách chuỗi PHC|Trích xuất params + salt + hash để xác minh|
|`Verify`|Xác minh mật khẩu|Dùng `subtle.ConstantTimeCompare` chống timing attack|
|Sentinel errors|Phân loại lỗi cụ thể|Dùng `errors.Is()` để xử lý có cấu trúc|
|Chuỗi PHC|Định dạng lưu trữ chuẩn|Tự chứa đủ thông tin để xác minh|

> Tham khảo thêm tại:
> 
> - Tài liệu thư viện: https://pkg.go.dev/golang.org/x/crypto/argon2
> - Khuyến nghị OWASP: https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html