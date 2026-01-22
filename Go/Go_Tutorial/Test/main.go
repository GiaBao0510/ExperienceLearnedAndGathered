package main

import (
	//"encoding/json"
	"fmt"
	//"hello/template2"
	"hello/services"
	// "hello/template1"
	// "hello/template2"
	//"os"
)

// func test(num1 string, num2 string) string {

// }

// Tạo hàm áp dụng thực tế của các hình
func PracticalApplication(x services.Shape) {
	fmt.Printf("%s\n", x.GetInfo())
	fmt.Printf("%s\n", x.Apply())
}

// Tạo hàm áp dụng thực tế của các hình
func PracticalApplicationPlus(x services.ShapePlus) {
	fmt.Printf("%s\n", x.GetInfo())
	fmt.Printf("%s\n", x.Apply())
	fmt.Printf("perimeter caculation formula: %s\n", x.PerimeterCalculationFormula())
}

func PrintInfo(v interface{}) {
	fmt.Println(v)
}

func main() {

	var i interface{} = "Hello, world!"

	s, ok := i.(string)
	if ok {
		fmt.Println("Giá trị là kiểu String: ", s)
	} else{
		fmt.Println("Giá trị không phải kiểu String")
	}

	f, ok := i.(float32)
	if ok {
		fmt.Println("Giá trị là kiểu Float32: ", f)
	} else{
		fmt.Println("Giá trị không phải kiểu Float32")
	}

}
