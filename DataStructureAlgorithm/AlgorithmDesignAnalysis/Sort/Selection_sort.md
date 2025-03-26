Selection Sort là một thuật toán sắp xếp đơn giản dựa trên việc **chọn ra phần tử nhỏ nhất (hoặc lớn nhất)** từ phần tử chưa được sắp xếp và đưa nó về vị trí đúng đang xét.

#### **Nguyên lý hoạt động:**

- **Bước 1:** tìm phần tử nhỏ nhất (hoặc lớn nhất) trong mảng chưa được sắp xếp.
- **Bước 2:** Hoán đổi phẩn tử nhỏ nhất (hoặc lớn nhất) với phần tử đầu tiên trong mảng chưa được sắp xếp.
- **Bước 3:** Lặp lại quá trình này cho đến cuối mảng (bỏ quan phần tử đã được sắp xếp).

---
#### **Độ phức tạp:**
- **Tốt nhất, trung bình và xấu nhất: *O(n^2).***
- **Độ phức tạp không gian: *O(1)*** do việc sắp xếp tại chỗ, không cần thêm bộ nhớ phụ.

---
#### **Ưu, nhược điểm:**

**Ưu điểm:**
- Đơn giản, dễ hiểu và dễ cài đặt
- Không cần phải thêm bộ nhớ phụ (in-place sorting).

**Nhược điểm:**
- Độ phức tạp thời gian ***O(n^2)*** kém hiệu quả với dữ liệu lớn.
- Không ổn định (unstable), có thể thay đổi thứ tự các phần tử bằng nhau.

---
#### **Hình minh họa:**

![](https://images.viblo.asia/2351ff37-da91-406f-b050-d6191e009e29.png)

---
#### **Code:**
```
public static void SelectionSort(int[] arr){

	int n = arr.Length;

	//Sắp xếp từng phần tử của mảng theo thứ tự tăng dần
	for(int i = 0; i < n - 1; i++){

		int minindex = i;
		for(int j = i + 1; j < n; j++ ){
			if(arr[j] < arr[minindex]){
				minindex = j;
			}
		}

		//Hoán vị
		int temp = arr[minindex];
		arr[minindex] = arr[i];
		arr[i] = temp;
	}
}
```



