Bubble sort (sắp xếp nổi bọt) đây là một thuật toán sắp xếp đơn giản, hoạt động bằng cách **so sánh và hoán đổi** các phần tử liền kề nếu chúng sai thứ tự.

#### **Nguyên lý hoạt động:**

- **Bước 1:** So sánh 2 phần tử liền kề, nếu phần tử phía trước lớn hơn phần tử phía sau thì hoán đổi giữa chúng
- **Bước 2:** Lặp lại quá trình này cho đến cuối mảng (Tức là mảng đã được sắp xếp)
- **Bước 3:** Mỗi lần lặp lại, phần tử có giá trị lớn nhất sẽ được di chuyển xuống cuối mảng

---
#### **Độ phức tạp:**
- **Tốt *nhất*: O(n)**, chỉ khi mảng đã được sắp xếp
- **Trung bình và xấu nhất: *O(n^2)*** 
- **Độ phức tạp không gian: *O(1)*** do sắp xếp tại chỗ, nên không cần thêm bộ nhớ phụ.


---
#### **Hình minh họa:**

![](https://images.viblo.asia/a0b867f2-069d-48fd-88bd-7a02a8daaa6e.png)

---
#### **Code:**
```
public static void BubbleSort(int[] arr){

	int n = arr.Length;
	int temp; //Biến tạm


	for(int i = 0; i < n-1; i++){
		for(int j = 0; j < n-i-1; j++ ){

			//So sánh 2 phần tử liền kề
			if(arr[j] > arr[j+1]){
				temp = arr[j];
				arr[j] = arr[j+1];
				arr[j+1] = temp;
			}
		}
	}
}
```