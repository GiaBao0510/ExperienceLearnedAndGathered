package services

//Tạo Interface đại diện cho các hình
type Shape interface {
	Apply() string
	GetInfo() string
}

//Tạo Interface mới
type ShapePlus interface {
	Shape
	PerimeterCalculationFormula() string
}
