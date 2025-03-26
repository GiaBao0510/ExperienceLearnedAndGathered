Quick Sort là một thuật toán sắp xếp hiệu quả dựa trên phương pháp **chia để trị (Divide and Conquer)**

![](https://www.w3resource.com/w3r_images/quick-sort-part-2.png)

![](https://www.w3resource.com/w3r_images/Sorting_quicksort_anim.gif)

---
### **Nguyên lý hoạt động:**

- **Bước 1:** Chọn một phần tử làm **pivot** (chốt) thường là phần tử dùng để làm **pivot** ở vị trí đầu tiên, hoặc ở giữa hoặc ở cuối mảng.
- **Bước 2 Phân hoạch (Partition):** Sắp xếp mảng sao cho các phần tử nhỏ hơn pivot ở bên trái và các phần tử lớn hơn pivot nằm ở bên phải
- **Bước 3 Đệ quy:** áp dụng Quick Sort cho 2 mảng con bên trái và bên phải pivot

---
### **Độ phức tạp thời gian:**

- **Tốt nhất và trung bình:** ***O(n log n)***
- **Xấu nhất:** ***O(n^2)*** .Khi pivot được chọn không tốt, ví dụ như mảng đã được sắp xếp và pivot là phần tử đầu tiên hoặc cuối cùng. 
- **Độ phức tạp không gian:**  ***O (log n)*** do sử dụng đệ quy.

---
### **Code:**

```
public static void QuickSort(int[] arr, int low, int high){

	if(low < high){
		int pivot = partition(arr, low, high);  // pivot
		
		QuickSort(arr, low, pivot - 1);         //Before pivot
		QuickSort(arr, pivot + 1, high);        //After pivot
	}
}


public static int partition(int[] arr, int low, int high){

	int pivot = arr[high];  //here default take the pivot at the end of the array
	int i = low - 1;        //index of smaller element
	int temp;               //temporary variable

	for(int j = low; j < high; j++){
		
		// if the value under consideration is less than the value
		//at the pivot, then swap the value at the index i 
		//with the value at the index j

		if(arr[j] < pivot){            
			i++;
			temp = arr[i];
			arr[i] = arr[j];
			arr[j] = temp;
		}
	}

	// swap the value at the pivot index with the value at the index i + 1
	temp = arr[i + 1];
	arr[i + 1] = arr[high];
	arr[high] = temp;
	
	// return the pivot index
	return i + 1;
}
```