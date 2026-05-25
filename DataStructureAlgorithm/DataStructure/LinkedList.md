# Linked List

## Linked List là gì?

Linked List là một cấu trúc dữ liệu dạng tuyến tính, gồm tập hợp các phần tử được gọi là **node**, trong đó mỗi node chứa hai thành phần:

- **Data:** Giá trị của node.
- **Next (pointer):** Con trỏ trỏ đến node tiếp theo trong danh sách.

Các node không được lưu trữ liên tiếp nhau trong bộ nhớ (khác với Array), mà được nối với nhau thông qua các con trỏ.

```
head
 │
 ▼
[1 | next] ──► [2 | next] ──► [3 | next] ──► nil
```

### So sánh Linked List và Array

|Tiêu chí|Array|Linked List|
|---|---|---|
|Bộ nhớ|Liên tiếp (contiguous)|Phân tán (non-contiguous)|
|Truy cập theo index|O(1)|O(n)|
|Chèn / xóa ở đầu|O(n)|O(1)|
|Chèn / xóa ở giữa|O(n)|O(n) — tìm node, sau đó O(1) để nối lại|
|Kích thước|Cố định (ở một số ngôn ngữ) hoặc cần cấp phát lại|Động, không cần cấp phát lại|

Linked List phù hợp khi cần chèn/xóa thường xuyên ở đầu hoặc giữa danh sách. Array phù hợp hơn khi cần truy cập ngẫu nhiên theo index.

### Các loại Linked List phổ biến

- **Singly Linked List:** Mỗi node chỉ có một con trỏ `Next`, trỏ về phía sau. Đây là loại được trình bày trong tài liệu này.
- **Doubly Linked List:** Mỗi node có hai con trỏ `Next` và `Prev`, có thể duyệt theo cả hai chiều.
- **Circular Linked List:** Node cuối cùng trỏ ngược lại node đầu tiên thay vì `nil`.

---

## Các thao tác cơ bản

Linked List hỗ trợ các thao tác chính sau:

- `addAtHead`: Thêm một node vào đầu danh sách — O(1).
- `addAtEnd`: Thêm một node vào cuối danh sách — O(n).
- `addAtIndex`: Thêm một node vào vị trí chỉ định — O(n).
- `deleteAtIndex`: Xóa node tại vị trí chỉ định — O(n).
- `FindNodeByValue`: Tìm node theo giá trị — O(n).
- `FindNodeByIndex`: Tìm node theo vị trí — O(n).
- `Print`: In toàn bộ danh sách — O(n).

---

## Triển khai bằng Go

Tạo file `linked_list.go` với nội dung sau.

### Định nghĩa struct và khởi tạo

```go
type Node struct {
    Val  int
    Next *Node
}

type LinkedList struct {
    head *Node
}

// NewLinkedList khởi tạo một linked list rỗng.
func NewLinkedList() *LinkedList {
    return &LinkedList{head: nil}
}
```

### Thêm node vào đầu danh sách — O(1)

```go
// addAtHead thêm một node mới vào đầu linked list.
func (ll *LinkedList) addAtHead(val int) {
    newNode := &Node{Val: val, Next: ll.head}
    ll.head = newNode
}
```

> Không cần kiểm tra `ll.head == nil` riêng vì khi danh sách rỗng, `ll.head` là `nil`, và `newNode.Next = nil` là hợp lệ.

### Thêm node vào cuối danh sách — O(n)

```go
// addAtEnd thêm một node mới vào cuối linked list.
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
```

### Tìm node theo index — O(n)

```go
// FindNodeByIndex trả về node tại vị trí index (0-based).
// Trả về nil nếu index vượt quá độ dài của danh sách.
func (ll *LinkedList) FindNodeByIndex(index int) *Node {
    current := ll.head
    i := 0

    for current != nil && i != index {
        current = current.Next
        i++
    }

    return current // nil nếu index vượt quá độ dài
}
```

### Tìm node theo giá trị — O(n)

```go
// FindNodeByValue trả về node đầu tiên có giá trị bằng val.
// Trả về nil nếu không tìm thấy.
func (ll *LinkedList) FindNodeByValue(val int) *Node {
    current := ll.head

    for current != nil && current.Val != val {
        current = current.Next
    }

    return current // nil nếu không tìm thấy
}
```

### Thêm node vào vị trí chỉ định — O(n)

```go
// addAtIndex thêm một node mới sau node tại vị trí index.
// Nếu index = 0, thêm vào đầu danh sách.
// Nếu index vượt quá độ dài, thêm vào cuối danh sách.
func (ll *LinkedList) addAtIndex(index int, val int) {
    if index == 0 {
        ll.addAtHead(val)
        return
    }

    prev := ll.FindNodeByIndex(index - 1)
    if prev == nil {
        // index vượt quá độ dài, thêm vào cuối
        ll.addAtEnd(val)
        return
    }

    newNode := &Node{Val: val, Next: prev.Next}
    prev.Next = newNode
}
```


### Xóa node tại vị trí chỉ định — O(n)

```go
// deleteAtIndex xóa node tại vị trí index (0-based).
func (ll *LinkedList) deleteAtIndex(index int) {
    if ll.head == nil {
        fmt.Println("Danh sách rỗng, không có gì để xóa")
        return
    }

    if index == 0 {
        ll.head = ll.head.Next
        return
    }

    prev := ll.FindNodeByIndex(index - 1)
    if prev == nil || prev.Next == nil {
        fmt.Printf("Index %d vượt quá độ dài của danh sách\n", index)
        return
    }

    prev.Next = prev.Next.Next
}
```


### In danh sách — O(n)

```go
// Print in toàn bộ giá trị của linked list ra stdout.
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
```

---

## Sắp xếp Linked List

### Sắp xếp bằng Insertion Sort trực tiếp trên danh sách — O(n²)

Triển khai sắp xếp bằng cách sao chép toàn bộ giá trị ra một slice, sắp xếp slice, rồi tạo lại danh sách mới. Cách này tốn thêm O(n) bộ nhớ và không thực sự sắp xếp trên Linked List. 
Phiên bản dưới đây thực hiện Insertion Sort trực tiếp trên con trỏ, không cần bộ nhớ phụ:

```go
// insertionSortList sắp xếp linked list tăng dần bằng Insertion Sort.
// Trả về head của danh sách đã sắp xếp.
func insertionSortList(head *Node) *Node {
    if head == nil {
        return nil
    }

    // dummy là node giả đứng trước danh sách đã sắp xếp
    dummy := &Node{}
    current := head

    for current != nil {
        next := current.Next // Lưu node tiếp theo trước khi thay đổi liên kết

        // Tìm vị trí chèn trong phần đã sắp xếp
        prev := dummy
        for prev.Next != nil && prev.Next.Val < current.Val {
            prev = prev.Next
        }

        // Chèn current vào vị trí đúng
        current.Next = prev.Next
        prev.Next = current

        current = next
    }

    return dummy.Next
}
```

**Ví dụ hoạt động:**

```
Input:  4 → 2 → 1 → 3
Bước 1: [4]           | xử lý 2 → [2 → 4]
Bước 2: [2 → 4]       | xử lý 1 → [1 → 2 → 4]
Bước 3: [1 → 2 → 4]   | xử lý 3 → [1 → 2 → 3 → 4]
Output: 1 → 2 → 3 → 4
```

---

## Đề xuất cải thiện thêm

### 1. Trả về error thay vì in thẳng ra stdout

Các hàm như `deleteAtIndex` hiện tại in thông báo lỗi trực tiếp bằng `fmt.Println`. Đây là thực hành không tốt trong Go — hàm thư viện không nên tự quyết định cách hiển thị lỗi, mà nên trả về `error` để caller tự xử lý:

```go
func (ll *LinkedList) deleteAtIndex(index int) error {
    if ll.head == nil {
        return fmt.Errorf("linked list is empty")
    }
    // ...
    return nil
}
```

### 2. Dùng `tail` pointer để tối ưu `addAtEnd` về O(1)

Hiện tại `addAtEnd` phải duyệt toàn bộ danh sách để tìm node cuối — O(n). Nếu thêm một con trỏ `tail` vào struct `LinkedList`, thao tác này trở thành O(1):

```go
type LinkedList struct {
    head *Node
    tail *Node
}
```

Cần cập nhật `tail` trong mỗi hàm thêm/xóa node để giữ đồng bộ.

### 3. Theo dõi độ dài danh sách

Thêm trường `size int` vào struct giúp trả về độ dài trong O(1) thay vì O(n), và giúp validate index nhanh hơn trước khi bắt đầu duyệt danh sách.

```go
type LinkedList struct {
    head *Node
    tail *Node
    size int
}
```

### 4. Sắp xếp hiệu quả hơn với Merge Sort — O(n log n)

Insertion Sort có độ phức tạp O(n²), không phù hợp với danh sách lớn. Merge Sort là thuật toán sắp xếp được khuyến nghị cho Linked List vì:

- Độ phức tạp O(n log n) trong mọi trường hợp.
- Không cần truy cập ngẫu nhiên theo index (phù hợp với đặc điểm của Linked List).
- Không cần bộ nhớ phụ O(n) như khi dùng với Array.