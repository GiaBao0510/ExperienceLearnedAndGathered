package keygen

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

const (
	DefaultPrefix  = "sk_live" // Tiền tố mặc định cho API key
	DefaultByteLen = 32        // Số byte ngẫu nhiên cho phần chính của API key
	HintLen        = 10        // Số ký tự hiển thị trong key hint
)

// Generate tạo một API key mới với tiền tố và độ dài ngẫu nhiên
func Generate(prefix string) (string, error) {

	// ------ Tạo phần ngẫu nhiên của API key ------
	b := make([]byte, DefaultByteLen)

	// Sử dụng crypto/rand để tạo dữ liệu ngẫu nhiên an toàn
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %v", err)
	}

	encoded := base64.RawStdEncoding.EncodeToString(b) // Mã hóa dữ liệu ngẫu nhiên thành chuỗi Base64
	return fmt.Sprintf("%s_%s", prefix, encoded), nil
}

// Hàm băm API key bằng SHA-256
func Hash(apikey string) string {
	h := sha256.Sum256([]byte(apikey)) // Tính toán băm SHA-256 của API key
	return hex.EncodeToString(h[:])    // Trả về chuỗi băm dưới dạng hex
}

// Hàm trả vê N ký tự đầu
func HintKey(apikey string) string {
	if len(apikey) <= HintLen {
		return apikey
	}
	return apikey[:HintLen]
}
