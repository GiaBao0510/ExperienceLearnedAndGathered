package main

import (
	"fmt"
)

func MyAppend(slice []int, value int) []int{
	slice = append(slice, value)
	PrintSlice(slice)
	return slice
}

func PrintSlice(slice []int){
	fmt.Printf("Slice: %v\n", slice)
	fmt.Printf("length: %v, cap: %v\n", len(slice), cap(slice))
}

func main() {
	slices := make([]int, 1)

	PrintSlice(slices)

	for i := 1; i <= 5; i++ {
		fmt.Println("\n---------------------")
		slices = MyAppend(slices, i)
	}
}
