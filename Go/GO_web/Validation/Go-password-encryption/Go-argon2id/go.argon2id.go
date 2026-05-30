/*
Package goargon2id cung cấp các hàm tiện ích để băm và xác minh mật khẩu
bằng thuật toán Argon2id — một trong những thuật toán băm mật khẩu an toàn nhất hiện nay.

Argon2id là biến thể kết hợp giữa Argon2i (kháng side-channel attack) và
Argon2d (kháng GPU/ASIC attack), được OWASP khuyến nghị cho password hashing.

Cấu hình thông số:
- HashLength: độ dài của hash được tạo ra (thường là 32 byte hoặc 64 byte).
- TimeCost: số lần lặp lại của thuật toán (tiêu chuẩn là 1 hoặc 10). Nếu số lần lặp càng nhiều thì thời gian để tạo hash càng lâu, nhưng cũng làm cho việc tấn công brute-force trở nên khó khăn hơn.
- MemoryCost: lượng bộ nhớ được sử dụng trong quá trình tính toán hash.
- Parallelism: số lượng luồng thực thi song song được sử dụng.

Định dạng chuỗi PHC (Password Hashing Competition) được sử dụng:

	$argon2id$v=<version>$m=<memory>,t=<iterations>,p=<parallelism>$<salt_base64>$<hash_base64>

Ví dụ:

	$argon2id$v=19$m=65536,t=2,p=4$c29tZXNhbHQ$aGFzaHZhbHVl

Trong đó:
- <variant>: loại Argon2 (d, i, hoặc id).
- <version>: phiên bản của thuật toán (thường là 19).
- <memory>: lượng bộ nhớ được sử dụng (tính bằng KiB).
- <iterations>: số lần lặp lại của thuật toán.
- <parallelism>: số lượng luồng thực thi song song.
- <salt>: giá trị salt được mã hóa bằng Base64.
- <hash>: giá trị hash được mã hóa bằng Base64.
*/
package goargon2id

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

// ---------------------------------------------------------------------------
// Sentinel errors — dùng errors.Is() để kiểm tra lỗi cụ thể ở tầng gọi hàm.
// Ví dụ: if errors.Is(err, ErrInvalidHashFormat) { ... }
// ---------------------------------------------------------------------------
var (
	// ErrInvalidHashFormat trả về khi chuỗi hash không đúng định dạng PHC.
	ErrInvalidHashFormat = errors.New("invalid hash format")

	// ErrUnsupportedAlgorithm trả về khi chuỗi hash không phải argon2id.
	ErrUnsupportedAlgorithm = errors.New("unsupported algorithm: expected argon2id")

	// ErrIncompatibleVersion trả về khi phiên bản trong hash khác với thư viện hiện tại.
	ErrIncompatibleVersion = errors.New("incompatible argon2 version")

	// ErrInvalidParams trả về khi parse tham số m/t/p thất bại.
	ErrInvalidParams = errors.New("invalid argon2 parameters")

	// ErrEmptyPassword trả về khi mật khẩu đầu vào rỗng.
	ErrEmptyPassword = errors.New("password must not be empty")
)

// ---------------------------------------------------------------------------
// Hằng số mặc định — tập trung tại một chỗ, dễ điều chỉnh sau này.
// Các giá trị này tuân theo khuyến nghị của OWASP (2023):
//   - TimeCost  : 2 lần lặp
//   - MemoryCost: 64 MiB (64 * 1024 KiB)
//   - Threads   : 1 (OWASP khuyến nghị 1 để tránh phụ thuộc vào số lõi CPU)
//   - KeyLength : 32 byte (256-bit)
//   - SaltLength: 16 byte (128-bit) — đủ để đảm bảo tính ngẫu nhiên
//
// ---------------------------------------------------------------------------
const (
	defaultTimeCost   uint32 = 2
	defaultMemoryCost uint32 = 64 * 1024 // 64 MiB
	defaultThreads    uint8  = 1
	defaultKeyLength  uint32 = 32
	defaultSaltLength uint32 = 16
)

// ---------------------------------------------------------------------------
// Argon2Params chứa các tham số điều chỉnh hành vi của thuật toán Argon2id.
// Đây là "config thuần túy" — không chứa dữ liệu runtime (salt, hash).
// Tách riêng khỏi dữ liệu runtime giúp tái sử dụng config cho nhiều lần băm.
// ---------------------------------------------------------------------------
type Argon2Params struct {
	TimeCost   uint32 // là số vòng lặp. Càng cao càng chậm → khó brute-force hơn.
	MemoryCost uint32 // là số lượng RAM sử dụng (đơn vị KiB). Tăng giá trị này khiến tấn công bằng GPU/ASIC tốn kém hơn
	Threads    uint8  // là số luồng thực thi song song. OWASP khuyến nghị 1 để tránh phụ thuộc vào số lõi CPU.
	KeyLength  uint32 // là độ dài của hash được tạo ra (đơn vị byte). 32 byte (256-bit) là lựa chọn phổ biến.
}

// DefaultArgon2Params trả về bộ tham số mặc định theo khuyến nghị OWASP.
// Dùng hàm này khi không muốn tự cấu hình tham số.
func DefaultArgon2Params() *Argon2Params {
	return &Argon2Params{
		TimeCost:   defaultTimeCost,
		MemoryCost: defaultMemoryCost,
		Threads:    defaultThreads,
		KeyLength:  defaultKeyLength,
	}
}

// ---------------------------------------------------------------------------
// Hasher là đối tượng chính thực hiện băm và xác minh.
// Giữ Argon2Params bên trong để không phải truyền params vào mỗi lần gọi.
// ---------------------------------------------------------------------------
type Hasher struct {
	params Argon2Params
}

// NewHasher tạo một Hasher mới với bộ tham số tuỳ chọn.
// Nếu muốn dùng tham số mặc định, truyền vào DefaultArgon2Params().
//
// Ví dụ:
//
//	h := goargon2id.NewHasher(goargon2id.DefaultArgon2Params())
//	hash, err := h.Hash("mypassword")
func NewHasher(params Argon2Params) *Hasher {
	return &Hasher{
		params: params,
	}
}

// generateSalt tạo một mảng byte ngẫu nhiên có độ dài saltLen dùng làm salt.
func generateSalt(saltLen uint32) ([]byte, error) {
	salt := make([]byte, saltLen) // cấp phát slice byte có độ dài saltLen

	//rand.Read điền vào slice bằng các byte ngẫu nhiên từ OS (CSPRNG).
	if _, err := rand.Read(salt); err != nil {
		return nil, fmt.Errorf("failed to generate salt: %w", err) // nếu có lỗi khi tạo salt, trả về lỗi
	}
	return salt, nil // trả về salt đã tạo
}

// Hash băm mật khẩu bằng thuật toán Argon2id và trả về chuỗi hash theo định dạng PHC.
func (h *Hasher) Hash(password string) (string, error) {
	// Kiểm tra nếu mật khẩu rỗng, trả về lỗi.
	if password == "" {
		return "", ErrEmptyPassword
	}

	// Tạo salt ngẫu nhiên.
	salt, err := generateSalt(defaultSaltLength)
	if err != nil {
		return "", fmt.Errorf("Hash: %w", err) // nếu có lỗi khi tạo salt, trả về lỗi
	}

	// Thực thi Argon2id: kết hợp password + salt + các tham số → raw hash bytes.
	hashRaw := argon2.IDKey(
		[]byte(password),
		salt,
		h.params.TimeCost,
		h.params.MemoryCost,
		h.params.Threads,
		h.params.KeyLength,
	)

	// Mã hóa salt và hash thành Base64 để đưa vào chuỗi PHC.
	saltBase64 := base64.RawStdEncoding.EncodeToString(salt)
	hashBase64 := base64.RawStdEncoding.EncodeToString(hashRaw)

	// Lắp ráp chuỗi PHC theo định dạng:
	encoded := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		h.params.MemoryCost,
		h.params.TimeCost,
		h.params.Threads,
		saltBase64,
		hashBase64,
	)

	return encoded, nil // trả về chuỗi hash đã mã hóa
}

// parsedHash  lưu trữ tham số và dữ liệu hash đã được parse từ chuỗi PHC.
type parsedHash struct {
	params  Argon2Params
	salt    []byte
	hashRaw []byte
}

// Tách và phân tích chuỗi PHC và trả về parseHash
func parseHash(encoded string) (*parsedHash, error) {

	// Chuỗi PHC có định dạng: $argon2id$v=19$m=65536,t=2,p=4$<salt_base64>$<hash_base64>
	parts := strings.Split(encoded, "$")

	// Cấu trúc hợp lệ phải có đúng 6 phần tử (index 0..5).
	// Index 0: "" (rỗng do bắt đầu bằng $)
	// Index 1: "argon2id"
	// Index 2: "v=19"
	// Index 3: "m=65536,t=2,p=1"
	// Index 4: <salt base64>
	// Index 5: <hash base64>
	if len(parts) != 6 {
		return nil, fmt.Errorf("parseHash: %w (got %d parts)", ErrInvalidHashFormat, len(parts))
	}

	// Kiểm tra phần thuật toán phải là "argon2id".
	if parts[1] != "argon2id" {
		return nil, fmt.Errorf("parseHash: %w", ErrUnsupportedAlgorithm)
	}

	// Phần version phải có định dạng.
	var version int
	if n, _ := fmt.Sscanf(parts[2], "v=%d", &version); n != 1 {
		return nil, fmt.Errorf("parseHash: %w (cannot parse version)", ErrInvalidHashFormat)
	}

	if version != argon2.Version {
		return nil, fmt.Errorf("parseHash: %w (hash v%d, library v%d)", ErrIncompatibleVersion, version, argon2.Version)
	}

	//parse 3 tham số m, t, p tư phần thứ 4 (index 3).
	var p Argon2Params
	n, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &p.MemoryCost, &p.TimeCost, &p.Threads)
	if err != nil || n != 3 {
		return nil, fmt.Errorf("parseHash: %w (cannot parse parameters)", ErrInvalidParams)
	}

	// Giải mã salt và hash từ Base64.
	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return nil, fmt.Errorf("parseHash: failed to decode salt: %w", err)
	}

	// Giải mã hash từ Base64.
	hashRaw, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return nil, fmt.Errorf("parseHash: failed to decode hash: %w", err)
	}

	// Keylength suy ra từ độ dài của hash đã giải mã.
	p.KeyLength = uint32(len(hashRaw))

	return &parsedHash{
		params:  p,
		salt:    salt,
		hashRaw: hashRaw,
	}, nil
}

// Verify xác minh mật khẩu plaintext có khớp với chuỗi PHC đã lưu
// Trả về:
// - true nếu mật khẩu hợp lệ (match hash)
// - false nếu mật khẩu không hợp lệ (không match hash)
// - lỗi nếu có vấn đề với định dạng hash hoặc tham số.
func (h *Hasher) Verify(storedHash, providedPassword string) (bool, error) {

	ph, err := parseHash(storedHash)
	if err != nil {
		return false, fmt.Errorf("Verify: %w", err) // nếu có lỗi khi parse hash, trả về lỗi
	}

	// Tính lại hash từ mật khẩu được cung cấp đúng các tham số
	computedHash := argon2.IDKey(
		[]byte(providedPassword),
		ph.salt,
		ph.params.TimeCost,
		ph.params.MemoryCost,
		ph.params.Threads,
		ph.params.KeyLength,
	)

	// subtle.ConstantTimeCompare so sánh hai slice byte trong thời gian cố định
	// (không phụ thuộc vào nội dung) để ngăn chặn timing attack.
	match := subtle.ConstantTimeCompare(ph.hashRaw, computedHash) == 1
	return match, nil // trả về kết quả so sánh
}
