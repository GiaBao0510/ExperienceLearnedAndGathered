package main

// Thuộc thành phần Concrete Product
type musket struct {
	Gun
}

func newMusket() IGun {
	return &musket{
		Gun: Gun{
			name: "musket gun",
			power: 50,
		},
	}
}