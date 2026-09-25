package circle

import (
	"errors"
	"strings"
)

type Circle struct {
	Name string `json: "Ten hinh tron`
}

// Create a constructor for circle
func New(name string) (*Circle, error) {

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

	return &Circle{
		Name: name,
	}, nil
}

// Hàm lấy giá trị
func (obj *Circle) GetInfo() string {
	return obj.Name
}

func (obj *Circle) Apply() string {
	return "Tao banh rang cua"
}

func (obj *Circle) PerimeterCalculationFormula() string{
	return "C = 2 * π * r"
}