**Insertion sort (sắp xếp chèn)** là một thuật toán sắp xếp đơn giản, hoạt động bằng cách **chèn từ phần tử vào đúng vị trí** trong mảng đã được sắp xếp.

#### **Nguyên lý hoạt động:**

- **Bước 1:** Coi phần tử đầu tiên của mảng là một mảng đã được sắp xếp
- **Bước 2:** So sánh từng phần tử tiền nhiệm (trước đó của nó), chèn nó vào đúng vị trí trong mảng được sắp xếp
- **Bước 3:** Lặp lại quá trình cho đến khi toàn bộ mảng được sắp xếp.

---
#### **Độ phức tạp thời gian:**
- **Tốt nhất: *O(n)*** khi mảng đã được sắp xếp.
- **Trung bình và xấu nhất là O(n^2).**
- Độ phức tạp không gian là O(1) do sắp xếp tại chố, không cần thêm bộ nhớ phụ.

---
#### **Ưu,nhược điểm**

**Ưu điểm:**
- Đơn giản, dễ hiểu, dễ cài đặt
- Hiệu quả với dữ liệu nhỏ hoặc gần như đã được sắp xếp.
- Không cần thêm bộ nhớ phụ (in-place sorting).
- Ổn định (stable), không thay đổi thứ tự các phần tử bằng nhau.

**Nhược điểm:**
- Độ phức tạp thời gian ***O(n^2)*** kém hiệu quả với dữ liệu lớn.

---
#### **Hình minh họa:**

![](https://media.geeksforgeeks.org/wp-content/uploads/20240802210251/Insertion-sorting.png)

---
#### **Code:**
```c
public static void InsertionSort(int[] arr){

	int n = arr.Length;

	for(int i = 1; i < n; i++){
		int j = i - 1;
		int key = arr[i];

		//Di chuyển cấc phần tử lớn hơn key về sau một vị trí so với vị trí ban đầu của nó
		while(j >=0 && arr[j] > key ){
			arr[j+1] = arr[j];
			j--;
		}

		//Chèn key vào vị trí đã đúng
		arr[j+1] = key;        
	}
}
```

```go
// Hàm sắp xếp mảng sử dụng thuật toán Insertion Sort
func InsertionSort(arr []int) {
	for i := 0; i < len(arr); i++{
		key := arr[i]

		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}
```