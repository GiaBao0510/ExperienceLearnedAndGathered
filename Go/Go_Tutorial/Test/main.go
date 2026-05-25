package main

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	head *Node
}

// Hàm khoi tạo một linked list mới
func NewLinkedList() *LinkedList {
	return &LinkedList{
		head: nil, // Danh sách hoàn toàn trống
	}
}

// Hàm tìm kiếm một node có giá trị cụ thể trong linked list
func (ll *LinkedList) FindNodeByValue(val int) *Node {
	current := ll.head

	for current != nil && current.Val != val {
		current = current.Next
	}

	return current
}

// Hàm tìm kiếm một node tại vị trí cụ thể trong linked list
func (ll *LinkedList) FindNodeByIndex(index int) *Node {
	curent := ll.head
	i := 0

	// Trường hợp nếu i không tới được index mà đã đi hết linked list thì trả về node cuối cùng
	for i != index && curent != nil {
		curent = curent.Next
		i++
	}

	// Nếu curent đã đi hết linked list mà vẫn chưa tới index thì trả về node rỗng
	if curent == nil {
		return &Node{} // Trả về node rỗng nếu index vượt quá độ dài của linked list
	}

	return curent
}

// Hàm thêm một node mới vào cuối linked list
func (ll *LinkedList) addAtEnd(val int) {
	newNode := &Node{Val: val, Next: nil}

	if ll.head == nil {
		ll.head = newNode
		return
	}
	current := ll.head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

// In ra các phần tử của linked list
func (ll *LinkedList) Print() {
	current := ll.head

	if current == nil {
		fmt.Println("Danh sách rỗng")
		return
	}
	for current != nil {
		fmt.Printf("%d ", current.Val)
		current = current.Next
	}
	fmt.Println()
}

// Thêm phần tử mới vào đầu linked list
func (ll *LinkedList) addAtHead(val int) {
	newNode := &Node{Val: val, Next: nil}

	if ll.head == nil {
		ll.head = newNode
		return
	}
	newNode.Next = ll.head // Liên kết node mới với node hiện tại ở đầu
	ll.head = newNode      // Cập nhật head để trỏ đến node mới, giờ nó trở thành node đầu tiên
}

// Thêm phần tử mới vào vị trí cụ thể trong linked list
func (ll *LinkedList) addAtIndex(index int, val int) {
	newNode := &Node{Val: val, Next: nil}

	// Nếu index là 0, thêm node mới vào đầu linked list
	if index == 0 {
		ll.addAtHead(val)
		return
	}

	// Tìm node tại vị trí index
	foundNode := ll.FindNodeByIndex(index)

	// Nếu node tại vị trí index không tồn tại, thêm node mới vào cuối linked list
	if foundNode == nil {
		ll.addAtEnd(val)
		return
	}

	// Thêm node mới vào giữa linked list
	newNode.Next = foundNode.Next // Liên kết node mới với node tiếp theo của node đã tìm thấy
	foundNode.Next = newNode      // Cập nhật liên kết của node đã tìm thấy để trỏ đến node mới
}

// Hàm xóa một node tại vị trí cụ thể trong linked list
func (ll *LinkedList) deleteAtIndex(index int) {

	// Nếu linked list rỗng, không có gì để xóa
	if ll.head == nil {
		fmt.Println("Danh sách rỗng, không có gì để xóa")
		return
	}

	// Nếu index là 0, xóa node đầu tiên
	if index == 0 {
		ll.head = ll.head.Next
		return
	}

	//Tìm node tại vị trí index
	foundNode := ll.FindNodeByIndex(index)

	// Nếu node tại vị trí index không tồn tại, không có gì để xóa
	if foundNode == nil {
		fmt.Printf("Không tìm thấy node tại index %d, không có gì để xóa\n", index)
		return
	}

	// Xóa node bằng cách cập nhật liên kết của node trước đó để bỏ qua node cần xóa
	current := ll.head
	for current.Next != nil && current.Next != foundNode {
		current = current.Next
	}

	// Nếu không tìm thấy
	if current.Next == nil {
		fmt.Printf("Không tìm thấy node tại index %d, không có gì để xóa\n", index)
		return
	}

	current.Next = foundNode.Next // Bỏ qua node cần xóa bằng cách liên kết node trước đó với node tiếp theo của node cần xóa
}

// Hàm sắp xếp mảng sử dụng thuật toán Insertion Sort
func InsertionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		key := arr[i]

		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// Hàm sắp xếp LinkedList sử dụng thuật toán Insertion Sort
func insertionSortList(head *Node) *Node {
	if head == nil {
		fmt.Println("Danh sách rỗng, không có gì để sắp xếp")
		return nil
	}

	sorted := &Node{Next: nil} // Tạo một node giả để giữ phần đầu của danh sách đã sắp xếp
	current := head
	temp_arr := []int{}

	for current != nil {
		temp_arr = append(temp_arr, current.Val) // Thêm giá trị của node hiện tại vào mảng tạm
		current = current.Next
	}

	InsertionSort(temp_arr) // Sắp xếp mảng tạm sử dụng thuật toán Insertion Sort

	current = sorted // Đặt con trỏ current về đầu danh sách đã sắp xếp
	for _, val := range temp_arr {
		current.Next = &Node{Val: val, Next: nil}
		current = current.Next // Di chuyển con trỏ current đến node tiếp theo
	}

	return sorted.Next // Trả về phần đầu của danh sách đã sắp xếp (bỏ qua node giả)
}

func main() {
	ll := NewLinkedList()
	ll.addAtEnd(1)
	ll.addAtEnd(4)
	ll.addAtEnd(8)
	ll.addAtEnd(7)
	ll.addAtEnd(9)
	ll.addAtHead(5)
	ll.addAtIndex(4, 11)

	ll.Print()
	//Xóa node tại index 4
	ll.deleteAtIndex(10)
	ll.Print() // Output: 5 1 4 8

	// Tìm kiếm node có giá trị 4
	node := ll.FindNodeByValue(4)
	fmt.Printf("Node có giá trị 4: %v\n", node.Val) // Output: Node có giá trị 4: &{4 <nil>}

	// Tìm kiếm node tại vị trí index 3
	node = ll.FindNodeByIndex(4)
	fmt.Printf("Node tại index 3: %v\n", node.Val) // Output: Node tại index 3: &{3 <nil>}

	// Sắp xếp linked list sử dụng thuật toán Insertion Sort
	ll.head = insertionSortList(ll.head)
	ll.Print() // Output: 1 4 5 7 8 9 11
}
