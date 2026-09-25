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
	// ✅ Dereference pointer để lấy giá trị
	bytes, err := bcrypt.GenerateFromPassword([]byte(b.Password), b.Cost)
	return string(bytes), err
}

// Hàm so sánh mật khẩu đã băm với mật khẩu gốc
func (b *Bcrypt) CheckPasswordHash(Password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(Password))
	return err == nil
}
