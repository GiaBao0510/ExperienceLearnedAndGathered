package square

import (
	"errors"
	"strings"
)

type Square struct {
	Name string `json: "Ten hinh vuong`
}

// Creta a constructor for square
func New(name string) (*Square, error) {

	//đk: loại bỏ khoan trắng thừa
	name = strings.TrimSpace(name)

	//Nếu rỗng thì báo lỗi
	if name == ""{
		return nil, errors.New("Tên không được để trống")
	}

	//Nếu độ dài quá 255 ký tự thì báo lỗi
	if len(name) > 255{
		return nil, errors.New("Số ký tự tối đa là 255.")
	}

	return &Square{
		Name: name,
	}, nil
}

// Hàm lấy giá trị
func (obj *Square) GetInfo() string {
	return obj.Name
}

func (obj *Square) Apply() string {
	return "Tao gach lat vuong"
}