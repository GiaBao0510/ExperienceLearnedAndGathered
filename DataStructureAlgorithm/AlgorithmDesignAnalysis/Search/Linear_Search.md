**Liner search(Tìm kiếm tuyến tính)** là một thuật toán tìm kiếm đơn giản nhất.
**Ý tưởng:** Duyệt từng phần tử trong mảng từ trái sang phải cho đến khi tìm thấy phần tử cần tìm, trường hợp xấu nhất là duyệt hết mảng.
**Ứng dụng:** Tìm kiếm trong danh sách nhỏ hoặc khi dữ liệu chưa được sắp xếp.

---
##### **Đặc điểm:**
**Độ phức tạp thời gian:**
- **Trường hợp tốt nhất:** ***O(1)*** (chỉ khi phần tử tìm kiếm ở vị trí đầu).
- T**rường hợp trung bình, xấu nhất:** ***O(n)*** (Phần tử tìm kiếm có thể ở giữa , cuối hoặc không có trong mảng)
**Không yêu cầu sắp xếp trước** (Có thể tìm trên danh sách chưa sắp xếp).
Dễ hiểu và triển khai, nhưng không hiệu quả với danh sách lớn.

---
##### **Hình minh họa:**

![](https://cafedev.vn/wp-content/uploads/2020/10/cafedev_line_search.png)

---
##### **Code:**
```
public static int LinearSearch(int[] arr, int x){

	// Linearly search for x in arr[]
	for(int i =0; i < arr.Length - 1; i++ ){

		if(arr[i] == x) return i;
	}

	return -1;  // if element is not present in the array
}
```