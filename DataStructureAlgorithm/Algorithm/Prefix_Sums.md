Giới thiệu mảng cộng dồn (prefix sum) trên mảng 1 chiều (1D)

Ví dụ về mảng Prefix sum:
****Input:**** arr[] = [10, 20, 10, 5, 15]  
****Output:**** [10, 30, 40, 45, 60]  
****Explanation:**** For each index i, add all the elements from 0 to i:  
prefixSum[0] = 10,   
prefixSum[1] = 10 + 20 = 30,   
prefixSum[2] = 10 + 20 + 10 = 40 and so on.

**Input:** arr[] = [30, 10, 10, 5, 50]  
**Output:** [30, 40, 50, 55, 105]  
***Explanation:** For each index i, add all the elements from 0 to i:  
prefixSum[0] = 30,   
prefixSum[1] = 30 + 10 = 40,  
prefixSum[2] = 30 + 10+ 10 = 50 and so on.

Ý tưởng để tạo một mảng PrefixSum:
- S0 = c, với c là một hằng số thực
- cho một mảng A có n phần tử được đánh số từ 0 đến n-1, ta sử dụng mảng S(A) theo quy tắc sau:  `S[i] = S[i-1] + A[i]`
- Mảng S(A) được gọi là mảng cộng dồn (tiền tố) theo c của A, gọi cách khác là prefix sum của A. Từ mảng A, ta có thể sinh ra vô hạn mảng S(A) bằng cách chọn số thực c tùy ý; trên thực tế chúng ta sẽ mặc định c = 0

![](https://i.imgur.com/lzBYJ89.gif)

Ví dụ:
```go
func BuildPrefixSum(nums []int, c float64) []int{
	ps := make([]int, len(nums)+1) //Khởi tạo mảng có độ dài hơn mảng đầu khoảng 1 đơn vị

	ps[0] = int(c)		// cho phần tử đầu tiên là giá trị của c

	// Tính toán giá trị prefix sum cho các phần tử tiếp theo
	for i := 0 ;i < len(nums); i++{
		ps[i + 1] = ps[i] + nums[i]
	}

	return ps
}

func prefixSum(nums []int, left, right int) int{
	return nums[right + 1] - nums[left]
}

func main(){
	nums := []int{-2, 0, 3, -5, 2, -1}
	result := BuildPrefixSum(nums, 0)
	
	fmt.Println(result) // Output: [0, -2, -2, 1, -4, -2, -3]
	fmt.Println(prefixSum(result, 2, 5)) // Output: -1
}
```

---
**Mảng hiệu (difference array)**

