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
#### **3. Các bước thực hiện Meger Sort:**

1. Chia mảng: thực hiện chia mảng thành 2 nửa.
2. Gọi đệ quy: Gọi Meger Sort cho từng nửa.
3. Trộn 2 nửa: Trộn 2 nữa đã sắp xếp lại với nhau

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
```
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