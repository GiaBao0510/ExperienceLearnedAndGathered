**Brute Force (hay Naive Search)** là một thuật toán tìm kiếm chuỗi con đơn giản, dựa trên việc so sánh từng ký tự của chuỗi con với từng ký tự của văn bản.

![](https://resources.stdio.vn/content/article/5ef629fcfc765a6e805ebc6e/resources/res-1596467893-1596467893337.png)

---
## **Nguyên lý hoạt động**

- Mô tả: duyệt và từng vị trí trong văn bản để so sánh từng ký tự của chuỗi con với văn bản bắt đầu từ vị trí đang xét.
- Kết quả: Trả về vị trí đang xét trong văn bản nếu tìm thấy, ngược lại thì trả về `-1`

---
## **Độ phức tạp thời gian:**

- **Tốt nhất:** ***O(n)*** chuỗi con nằm ở phía đầu văn bản.
- **Trung bình và xấu nhất:** ***O(n x m)*** với n độ dài của văn bản, còn m là độ dài của chuỗi con.
- **Độ phức tạp không gian:** ***O(1)*** không cần thêm bộ nhớ phụ

---
## **Code:**
```
public static int NaiveSearch(string target, string text){

	//Get the length of the text and the target
	int n = text.Length;
	int m = target.Length;

	//Loop through the text
	for(int i = 0; i < n; i++){

		// compare each character from the substring in the text
		int j;

		for(j = 0; j < m; j++){

			//If the characters do not match, break
			if(text[i + j] != target[j])
				break;
		}

		//If the inner loop completed, it means the substring was found
		if(j == m)
			return i;
	}

	return -1;  //If the substring was not found
}
```