package main

// Thuộc thành phần Product Interface
type IGun interface {
	setName(name string)
	getName() string
	setPower(power int)
	getPower() int
}