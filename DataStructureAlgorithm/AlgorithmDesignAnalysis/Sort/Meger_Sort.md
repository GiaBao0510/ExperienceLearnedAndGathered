**Sắp xếp trộn (Meger Sort)** là một giải thuật sắp xếp dựa trên giải thuật ***Chia để trị (Divide and Conquer)***. Với độ phức tạp thời gian trường hợp xấu nhất là O(n log n).

#### **1. Nguyên lý hoạt động:**

- Chia (Divide): chia mảng thành 2 nửa bằng nhau (hoặc gần bằng nhau) cho đến khi chỉ còn 1 phần tử.
- Trị(Conquer): Sắp xếp từng phần tử nhỏ.
- Kết hợp (Combine): trộn (meger) các thành phần đã sắp xếp lại với nhau lại thành mảng hoàn chỉnh.

---
#### **2. Độ phức tạp thời gian:**

- **Tốt nhất, trung bình và xấu nhất: O (n log n).**
	- Luôn chia mảng thành 2 nửa và trộn lại với độ phức tạp là O(n).
- Độ phức tạp không gian O(n) do cần bộ nhớ để lưu trữ mảng tạm.

---
### **Ý tưởng thuật toán:**
1. Nếu độ dài mảng không lớn hơn 1 phần tử thì coi như xong, không thực hiện gì thêm
2. Nếu độ dài mảng nhiều hơn 1 phần tử, thì chia đôi mảng ra thành hai thành phần bằng nhau hoặc gần bằng nhau
3. Sắp xếp 2 dãy con này
4. Gộp 2 dãy có thứ tự này thành dãy có kích thước ban đầu

---
### **3. Các bước thực hiện Meger Sort:**

#### 1. Chia mảng: thực hiện chia mảng thành 2 nửa.

Việc này coi như dễ, cho rằng
- Dãy thứ nhất bắt đầu từ `left` đến `mid`
- Dãy thứ 2 bắt đầu từ `mid + 1` đến `right`
Ở đây thì `mid` và vị trí ngay giữa mảng: `mid = (left + right) / 2`

#### 2. Sắp xếp 2 dãy con này.
 Chỉ cần gọi đệ quy là được
 Thì biết rằng, nguyên tắt để có thể dùng đệ quy thì ta cần tuân thủ theo một số điều kiện như sau:
 - Có tồn tại trường hợp thoát đệ quy
 - Có cách tổng hợp lại kết quả từ bài toán lớn, từ các bài toán con đệ quy
Thì ở đây, điều kiện thứ (1) đã có "độ dài mảng không lớn hơn 1 phần tử"
Còn điều kiện thứ 2 dẽ nói ở bước tiếp theo
#### 3. Trộn 2 nửa: Trộn 2 nữa đã sắp xếp lại với nhau

Bài toán cho 2 dãy đã được sắp xếp tăng dần, hẫy gộp chúng lại thành một dãy mới. Dãy mới phải đảm bảo cũng được sắp xếp tăng dần

Cách giải quyết đơn giản là ta chỉ chần giải thuật O (M + N) là có thể giải quyết được (với N và M lần lượt là kích thước của 2 dãy con).

---
#### **4. Ưu, nhược điểm:**
**Ưu điểm:**
- Luôn có độ phức tạp ***O (n log n)*** trong mọi trường hợp.
- Ổn định (stable), không đổi thứ tự các phần tử bằng nhau.
**Nhược điểm:**
- Cần thêm bộ nhớ phụ ***O ( n )*** để lưu trữ mảng tạm.
- Không phải là thuật toán tại chỗ (in-place).

---
#### **5. Hình minh họa:**
![](https://freetuts.net/upload/tut_post/images/2021/09/17/5502/bai23-02.png)


---
#### **6. Code:**
```csharp
//Hàm trộn 2 mảng con
public static void Merge(int[] arr, int left, int mid, int right){

	int n1 = mid - left + 1,        //Lấy kích cỡ mảng con bên trái
		n2 = right - mid;           //Lấy kích cỡ mảng con bên phải

	//Tạo mảng tạm
	int[] L = new int[n1];
	int[] R = new int[n2];

	//Sao chép dữ liệu vào mảng tạm
	for(int i = 0; i < n1; i++)
		L[i] = arr[left + i];
	for(int j = 0; j < n2; j++)
		R[j] = arr[mid + 1 + j];

	//Trộn 2 mảng con
	int k = left;       //Chỉ số bắt đầu của mảng kết quả
	int x = 0, y = 0;   //Chỉ số mảng con bên trái và bên phải

	while( x < n1 && y < n2){

		if(L[x] <= R[y]){
			arr[k] = L[x];
			x++;
		}else{
			arr[k] = R[y];
			y++;
		}
		k++;
	}

	//Sao chép các phần tử còn lại của mảng con bên trái (nếu có)
	while(x < n1){
		arr[k] = L[x];
		x++;
		k++;
	}

	//Sao chép các phần tử còn lại của mảng con bên phải (nếu có)
	while(y < n2){
		arr[k] = R[y];
		y++;
		k++;
	}
}

//Hàm sắp xếp Merger sort
public static void Sort(int[] arr, int left, int right){

	if(left < right){
		int mid = left + (right -left)/2;   //Tìm điểm giữa

		//Gọi đệ quy để thực hiện phân tách và sắp xếp
		Sort(arr, left, mid);
		Sort(arr, mid+1, right);

		//Trộn 2 mảng đã sắp xếp lại với nhau
		Merge(arr, left, mid, right);
	}
}
```

```go
// Hàm gộp hai phần đã sắp xếp
func Sort(arr []int, left, right int) {
	if left < right {
		mid := left + (right-left)/2

		// Gọi đệ quy để thực hiện phân tách và sắp xếp
		Sort(arr, left, mid)
		Sort(arr, mid+1, right)

		// Gộp hai phần đã sắp xếp
		Merge(arr, left, mid, right)

	}
}

// Hàm gộp hai phần đã sắp xếp
func Merge(arr []int, left, mid, right int) {

	// Tính kích thước của hai phần đã sắp xếp
	n1 := mid - left + 1 // Lấy độ dài của phần bên trái
	n2 := right - mid    // Lấy độ dài của phần bên phải

	// Tạo mảng tạm để lưu trữ phần đã sắp xếp
	L := make([]int, n1)
	R := make([]int, n2)

	// Sao chép dữ liệu vào mảng tạm
	for i := 0; i < n1; i++ {
		L[i] = arr[left+i]
	}
	for j := 0; j < n2; j++ {
		R[j] = arr[mid+1+j]
	}

	// Trộn 2 mảng con lại thành một mảng đã sắp xếp
	k := left
	x, y := 0, 0
	for x < n1 && y < n2 {
		if L[x] <= R[y] {
			arr[k] = L[x]
			x++
		} else{
			arr[k] = R[y]
			y++
		}
		k++
	}

	// Sao chép phần còn lại của mảng L nếu có
	for x < n1 {
		arr[k] = L[x]
		x++
		k++
	}

	// Sao chép phần còn lại của mảng R nếu có
	for y < n2 {
		arr[k] = R[y]
		y++
		k++
	}

}
```

----
### **So sánh thuật Quick sort:**

1. Tính ổn định: Ổn định hơn Quick Sort
2. Không gian lưu trữ: Tốn bộ nhớ gấp đôi Quick Sort (do phải lưu mảng tạm)
3. Cài đặt: Ý tưởng đơn giản, nhưng cài đặt dài dòng hơn Quick Sort.

---
### Kết luận:
