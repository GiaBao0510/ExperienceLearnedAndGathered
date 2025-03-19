Thuật toán Boyer-Moore Majority Voting là một thuật toán hiệu quả để tìm phần tử chiếm đa số (majority element) trong một mảng. Phần tử chiếm đa số được định nghĩa là phần tử xuất hiện hơn n/2 lần trong mảng có n phần tử.

## Nguyên lý cơ bản

Thuật toán dựa trên một nhận xét quan trọng: nếu một phần tử chiếm đa số trong mảng, thì sau khi "loại bỏ" các cặp phần tử khác nhau, phần tử chiếm đa số vẫn sẽ còn lại.

## Các bước của thuật toán

Thuật toán gồm 2 pha chính:

### Pha 1: Tìm ứng viên (Candidate Finding)

1. Khởi tạo hai biến:
    - `candidate` (ứng viên): giá trị bất kỳ
    - `count` (số lượng): 0
2. Duyệt qua từng phần tử của mảng:
    - Nếu `count` = 0, gán phần tử hiện tại là `candidate` mới và đặt `count` = 1
    - Nếu phần tử hiện tại bằng `candidate`, tăng `count` lên 1
    - Nếu phần tử hiện tại khác `candidate`, giảm `count` đi 1
3. Sau khi duyệt hết mảng, `candidate` là ứng viên có khả năng là phần tử chiếm đa số

### Pha 2: Kiểm tra (Verification)

1. Đếm số lần xuất hiện của `candidate` trong mảng
2. Nếu số lần xuất hiện > n/2, thì `candidate` là phần tử chiếm đa số
3. Nếu không, mảng không có phần tử chiếm đa số

### ***Code:***
```
public static int BoyerMooreMajorityVoting(int[] nums){
	int candidate = nums[0];
	int votes = 1;
	
	//Browse to find the candidate woth the most votes
	for(int i = 0; i < nums.Length; i++){
		
		/* if the number of votes return to 0, then select
		the current candidate */
		if(votes == 0){
			candidate = nums[i];
			votes++;
		}else{
			
			//candidate like elements will increase
			if(nums[i] == candidate)
				votes++;
			
			//Otherwise decrease
			else
				votes--;
		}
	}
	
	//Check if the candidate with most votes > N.2
	votes = 0;
	for(int i = 0; i < nums.Length; i++){
		if(candidate == nums[i])
			votes++;
	}
	
	if(votes > nums.Length/2) return candidate;
	return -1;
}
```

### Độ phức tạp thời gian
- Pha 1: O(n) - duyệt qua mảng một lần
- Pha 2: O(n) - đếm số lần xuất hiện của ứng viên
- Tổng cộng: O(n)

### Độ phức tạp không gian
- O(1) - chỉ sử dụng hai biến phụ: `candidate` và `count`

## Giải thích trực quan

Bạn có thể hiểu thuật toán này như một quá trình "loại bỏ" cặp phần tử khác nhau:

1. Khi gặp một phần tử giống với `candidate`, ta coi như "thêm vào" một phần tử (tăng `count`)
2. Khi gặp một phần tử khác `candidate`, ta coi như "loại bỏ" một cặp gồm phần tử hiện tại và một phần tử `candidate` (giảm `count`)
3. Khi `count` về 0, nghĩa là tất cả các phần tử đã được "ghép cặp và loại bỏ", ta bắt đầu với phần tử mới

Nếu có một phần tử chiếm đa số (>n/2), thì sau khi "loại bỏ" tất cả các cặp phần tử khác nhau, phần tử đó vẫn sẽ còn lại vì nó xuất hiện nhiều hơn tổng số các phần tử còn lại.

## Ứng dụng

Thuật toán Boyer-Moore Majority Voting được ứng dụng trong:
1. Xử lý dữ liệu lớn khi bộ nhớ hạn chế
2. Phân tích dữ liệu và thống kê
3. Hệ thống bỏ phiếu và đồng thuận
4. Kỹ thuật xử lý luồng dữ liệu (data streaming)

## Kết luận

Thuật toán Boyer-Moore Majority Voting là một ví dụ tuyệt vời về thuật toán đơn giản nhưng hiệu quả. Với độ phức tạp thời gian O(n) và không gian O(1), nó là lựa chọn tối ưu để giải quyết bài toán tìm phần tử chiếm đa số trong mảng.