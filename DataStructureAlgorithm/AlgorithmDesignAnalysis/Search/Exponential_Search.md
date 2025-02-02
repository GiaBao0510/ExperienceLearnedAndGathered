#### **Giới  thiệu về Exponential Search:**

- **==Exponential Search==** là một thuật toán dùng để tìm kiếm nhanh **vị trí ban đầu** của giá trị cần tìm trong **mảng đã sắp xếp**, trước khi áp dụng **Binary Search**.
- Thuật toán này c==hỉ dùng cho mảng được sắp xếp==.
- Tìm phần tử bằng cách tăng chỉ số theo lũy thừa của 2 {1,2,4,8,...}
- **Nhanh hơn thuật toán Binary Search** nếu phần tử cần tìm ở gần đầu mảng.
- **Độ phức tạp trung bình *O(log n)***, tốt cho việc tìm kiếm dữ liệu lớn

---
#### **Ý tưởng thuật toán:**

1. **Tìm khoảng chứa phần tử cần tìm** bằng cách kiểm tra các chỉ số theo lũy thừa của 2 {1,2,4,8,16,...}.
2. **Khi vượt qua vị trí của phần tử cần tìm,** thì thu hẹp phạm vi tìm kiếm bằng ==Binary Search== trong khoảng đã tìm được.

---
#### **Hình ảnh:**

![](data:image/png;base64,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)

---
#### **Code:**
```
 public static int ExponentialSearch(int[] nums, int target) {

	int n = nums.Length;    // Length of the array
	int m = 1;              // Initial range for binary search

	// Base case: If element is present at the first location itself
	if(nums[0] == target)
		return 0;

	// Find range for binary search by repeated doubling
	while(m < n && nums[m] <= target){
		m *= 2;
	}

	// Call binary search for the found range
	int start = m/2;
	int end = Math.Min(m, n-1);

	while(start <= end){
		int mid = start + (end - start)/2;
		if(nums[mid] == target) return mid;
		else if(nums[mid] < target) start = mid + 1;
		else end = mid - 1;
	}
	
	return -1;      // If element is not present in the array
}
```