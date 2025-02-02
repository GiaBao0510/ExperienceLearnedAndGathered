Thuật toán này được đề xuất bởi Devid E.Ferguson (CACM-1960), thuật toán này dùng để sử dụng tìm kiếm trên một mảng đã được sắp xếp hoặc có cấu trúc đặc biệt.

#### **Nguyên lý hoạt động:**

 - **Mô tả:** Sử dụng dãy số Fibonacci để chia danh sách thành các phần có kích thước gần bằng nhau, sau đó so sánh giá trị với phần tử cần tìm tại vị trí Fibonacci.
 - **Kết quả:** Trả về vị trí của phần tử nếu được tìm thấy, ngược lại thì thông báo vị trí không tồn tại.

---
#### **Độ phức tạp:**

- **Tốt nhất, trung bình và xấu nhất: *O(log n)*.** Fibonacci Search có xu hướng hiệu quả hơn khi danh sách lớn.
- **Độ phức tạp không gian: *O(1)*** (không cần thêm bộ nhớ phụ).


---
#### **Các bước thực hiện Fibonacci Search:**

1. **Khởi tạo dãy Fibonacci:** Tạo ra một dãy Fibonacci sao cho số cuối cúng có giá trị lớn hơn hoặc bằng kích cỡ danh sách.
2. **Chia danh sách:** Sử dụng các số Fibonacci để chia danh sách thành các phần có kích cỡ gần bằng nhau,
3. **So sánh:** So sánh giá trị cần tìm với phần tử tại vị trí Fibonaci
	- Nếu bằng nhau thì trả về vị trí,
	- Nếu giá trị tìm nhỏ hơn, tiếp tụ tìm kiếm trong phần bên trái.
	- Nếu giá trị cần tìm lớn hơn, tiếp tục tìm kiếm trong phần bên phải.
4. **Lặp lại** quá trình này cho đến khi tìm thấy hoặc phạm vi tìm kiếm rỗng.

---
#### **Hình ảnh minh họa:**

![](data:image/jpeg;base64,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)

---
#### **Code:**
```
public static int FibonacciSearch(int[] arr, int x){

	// initialize fibonacci numbers
	int n = arr.Length;     // length of array
	int Fm1 = 1,            // Fm-1
		Fm2 = 0,            // Fm-2
		Fm = Fm1 + Fm2;     // Fm

	// Find the smallest Fibonacci number greater than or equal to n
	while(Fm1 < n){
		Fm2 = Fm1;
		Fm1 = Fm;
		Fm = Fm1 + Fm2;
	}

	// offset
	int offset = -1;

	// search
	while(Fm > 1){

		int i = Math.Min(offset + Fm2, n-1);    // i = min(offset + Fm2, n-1)

		if(arr[i] < x){                         // if x is greater than the value at index i
			Fm = Fm1;
			Fm1 = Fm2;
			Fm2 = Fm - Fm1;
			offset = i;
		}else if(arr[i] > x){                   // if x is less than the value at index i
			Fm = Fm2;
			Fm1 -= Fm2;
			Fm2 = Fm - Fm1;
		}else{                                  // if x is equal to the value at index i
			return i;
		}
	}

	// compare last element with x
	if(Fm1 == 1 && offset + 1 < n && arr[offset + 1 ] == x) return offset + 1; // return

	return -1; // not found
}
```