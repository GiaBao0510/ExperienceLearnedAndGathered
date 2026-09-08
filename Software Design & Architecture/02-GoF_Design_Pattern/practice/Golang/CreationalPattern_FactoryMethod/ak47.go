package main

// Thuộc thành phần Concrete Product
type Ak47 struct {
	Gun
}

func newAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name: "Ak47 gun",
			power: 100,
		},
	}
}