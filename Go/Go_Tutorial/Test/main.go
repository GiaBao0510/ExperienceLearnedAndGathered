package main

import (
	"fmt"
)

func test(n int) int {
	if n <= 2 {
		return n
	}

	idx1, idx2 := 1, 2

	for i := 3; i <= n; i++ {
		temp := idx1 + idx2

		fmt.Printf("- temp: %d, idx1: %d, idx2: %d\n", temp, idx1, idx2)
		idx1 = idx2
		idx2 = temp
		fmt.Printf("--- temp: %d, idx1: %d, idx2: %d\n", temp, idx1, idx2)
	}

	return idx2
}

func UpdateName(name string) {
	name = "Nguyen Van B"
	fmt.Printf("Data type: %T\n", name)
	fmt.Printf("Value: %v\n", name)
	fmt.Printf("Address: %v\n", &name)
}

func main() {

	fmt.Printf("Fibonacci of 10 is: %d\n", test(5))

	// P_name := &name

	// fmt.Printf("\nData type P_name: %T\n", P_name)
	// fmt.Printf("Value P_name: %v\n", P_name)
	// fmt.Printf("Dereference P_name: %v\n", &P_name)
	// fmt.Printf("Value at Dereference P_name: %v\n", *P_name)

	// P_name2 := &P_name
	// fmt.Printf("\nData type P_name2: %T\n", P_name2)
	// fmt.Printf("Value P_name2: %v\n", P_name2)
	// fmt.Printf("Dereference P_name2: %v\n", &P_name2)
	// fmt.Printf("Value at Dereference P_name2: %v\n", *P_name2)
	// fmt.Printf("Value at Dereference Dereference P_name2: %v\n", **P_name2)
}
