
### **1️⃣ Run-Length Encoding là gì?**

**Run-Length Encoding (RLE)** là một thuật toán nén dữ liệu không mất mát **(lossless compression)** bằng cách thay thế dãy ký tự lặp lại liên tục bằng một cặp *(Số lần lặp, ký tự)*.

***Ví dụ***:

![](https://cdn.prod.website-files.com/6544bdfdf184393c08205ae6/67092df871db3a47522abe04_66f685de0eb9f9e353276eec_66f67c2d90b36c7925e55307_How%252520generative%252520AI%252520works-1.jpeg)

---
### **2️⃣ Khi nào nên sử dụng RLE?**

✅**Tốt khi dữ liệu có nhiều ký tự lặp lại** (Ví dụ: ảnh trắng đen, văn bản chứa nhiều khoảng trắng).
✅Hiệu quả cho định dạng ảnh như BMP, TIF, PCX hoặc chuỗi DNA trong sinh học.
❌**Không hiệu quả nếu dữ liệu không có nhiều phần lặp** (Ví dụ: văn bản ngẫu nhiên).

---
### **3️⃣ Độ phức tạp của thuật toán**

| Trường hợp     | Độ phức tạp                              |
| -------------- | ---------------------------------------- |
| **Tốt nhất**   | *O(n)* - nếu có nhiều phần tử giống nhau |
| **trung bình** | *O(n)*                                   |
| **Xấu nhất**   | *O(n)* - nếu không có ký tự lặp lại.     |
❗**Lưu ý:** nếu dữ liệu không có ký tự lặp, kích thước sau khi nén có thể lớn hơn ban đầu (Do thêm lần xuất hiện vào trước mỗi ký tự).

---

## **Code:** 

```
public static string RLE(string s){

	if(s.Length == 0)
		return "";

	//StringBuilder
	var result = new System.Text.StringBuilder(s.Length * 2);

	//Initial count and character
	int count = 1;
	char prevChar = s[0];

	//Go through each character and count it instantly
	for(int i = 1; i < s.Length; i++){
		/* if the charater being viewed is the same as the previous character
		then count */
		if(s[i] == prevChar)
			count++;

		/*if the character being viewed is different from the previous character
		then append the count and the character to the result */

		else{
			result.Append(count);
			result.Append(prevChar);
			prevChar = s[i];
			count = 1;
		}
	}



	//Append the last character and count
	result.Append(count);
	result.Append(prevChar);

	return result.ToString();
}
```