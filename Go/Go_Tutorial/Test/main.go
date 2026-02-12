package main

import (
	//"encoding/json"
	"fmt"
	"slices"
	//"reflect"
	// "reflect"
	// "slices"
	//"golang.org/x/exp/constraints"
	//"go/types"
	//"hello/template2"
	//"hello/services"
	//"golang.org/x/text/message"
	// "hello/template1"
	// "hello/template2"
	//"os"
)


func main() {

	
	MySlices := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	
	//Clone: Tạo bản sao của slice
	copied := slices.Clone(MySlices)
	fmt.Println(copied)

	//So sánh 2 slices có giống nhau không
	fmt.Println(slices.Equal(MySlices, copied))
	fmt.Println(slices.Equal([]int{1,2,3}, []int{5,6,7,8}))

	//Tìm vị trí đầu tiên của phần tử
	fmt.Println(slices.Index([]int {1,2,3,4,5,7,7,9,5,10,1,7,2}, 7) )

	//Kiểm tra phần tử có nằm trong slice không
	fmt.Println(slices.Contains(MySlices, "f"))

	//Chèn phần tử vào vị trí thứ i
	copied = slices.Insert(copied, 5, "ok")
	fmt.Println(copied)

	//Xóa vị trí thứ i đến j
	copied = slices.Delete(copied, 5, 8)
	fmt.Println(copied)

	//Đảo ngược slice
	slices.Reverse(copied) //Hàm này không có trả về
	fmt.Println(copied)

	//Hàm sort
	slices.Sort(copied)
	fmt.Println(copied)

	//Hàm sort có điều kiện
	Nums := []int {1, 5, 6, 7, 4, 3, 2, 10, 9}
	slices.SortFunc(Nums, func(a, b int) int{
		return b - a
	})
	fmt.Println(Nums)

	//Lấy giá trị lớn nhất
	fmt.Println(slices.Max(Nums))

	//Lấy giá trị nhỏ nhất
	fmt.Println(slices.Min(Nums))
}


