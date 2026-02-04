package main

import (
	//"encoding/json"
	"fmt"
	//"golang.org/x/exp/constraints"

	//"go/types"
	//"hello/template2"
	//"hello/services"
	//"golang.org/x/text/message"
	// "hello/template1"
	// "hello/template2"
	//"os"
)

func SumList(list []int) int{
	return SumHelper(list, 0)
}

func SumHelper(list []int, idx int) int{
	if idx >= len(list){
		return 0
	}

	return list[idx] + SumHelper(list, idx + 1)
}

func sumRange(start, end, acc int) int {
	if(start > end){
		return acc
	}

	return sumRange(start +1, end, start + acc)
}

func main() {

	// Tạo mảng và gán giá trị cho từng phần tử
	nums := [...]int{20,40,60,80,100} 		
	fmt.Println(nums)
	fmt.Println("---------------")

	//Mảng chuỗi chua có giá trị
	var chars [10]string
	chars[5] = "Hu"
	fmt.Println(chars)
	fmt.Println("---------------")

	//
}
