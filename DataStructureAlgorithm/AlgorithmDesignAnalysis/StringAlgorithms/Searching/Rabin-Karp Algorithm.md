![](https://images.shiksha.com/mediadata/ugcDocuments/images/wordpressImages/2023_01_MicrosoftTeams-image-252.jpg)

Rabin-Karp Alogorithm là một thuật toán tìm kiếm chuỗi con hiệu quả, sử dụng hashing để so sánh chuỗi con với các đoạn trong văn bản. Đặc biệt thuật toán hữu ích khi cần tìm nhiều chuỗi con cùng lúc hoặc khi chuỗi con có độ dài cố định

---
## **Nguyên lý hoạt động:**

- **Mô tả:** Thuật toán sử dụng hàm băm (hash funtion) để tính giá trị băm của chuỗi con và so sánh nó với giá trị băm của các đoạn có cùng độ dài trong chuỗi lớn
- **Kết quả:** Nếu giá trị băm khớp, thuật toán sẽ so sánh từng ký tự kết quả
- **Mục tiêu:** Giảm độ phức tạp thời gian trung bình từ ***O(n x m)* (Brute Force)** xuống ***O(n + m)***

---
## **Các bước thực hiện Rabin-Karp:**

**1. Tính giá trị băm của chuỗi con (pattern hash):**
- Sử dụng hàm băm để tính giá trị băm của chuỗi con
**2. Tính giá trị băm của các đoạn trong văn bản (text hash):**
- Tính giá trị băm của đoạn đầu tiên trong chuỗi có độ dài bằng chuỗi con.
- Sử dụng ***Rolling hash*** để tính giá trị băm của đoạn tiếp theo một cách hiệu quả.
**3. So sánh giá trị băm**
- Nếu giá trị băm của đoạn trong văn bản khớp với giá trị băm của chuỗi đơn so sánh từng ký tự để xác nhận 

---
## **Rolling Hash:**

- **Mô tả:** Rolling hash là kỹ thuật tính toán giá trị băm của đoạn tiếp theo dựa trên giá trị băm của đoạn trước đó, giúp giảm độ phức tạp thời gian.
    
- **Công thức:**
    
    **hashnew = ( hashold − text[i] * base^m−1) × base + text[i+m]**
    
    Trong đó:
    - hashold​: Giá trị băm của đoạn cũ.
    - hashnew​: Giá trị băm của đoạn mới.
    - base: Cơ số (thường là số nguyên tố lớn).
    - m: Độ dài chuỗi con.

---
## **Độ phức tạp:**

- **Tốt nhất và trung bình: *O(n + m)***
- **Xấu nhất: *O(n * m)***

---
## **Code:**
```csharp
// Hàm băm cho một chuỗi
static long ComputeHash(string str, int length, int prime, int baseValue){

	long hash = 0;
	for(int i = 0; i < length; i++){
		hash = (hash * baseValue + str[i]) % prime;
	}

	return hash;
}

//Hàm tính rolling hash
static long RecomputeHash(long hash, char oldChar, char newChar, int baseValue, int prime, int power){
	hash = (hash - oldChar * power) % prime;
	if(hash < 0) hash += prime;
	hash = (hash * baseValue + newChar) % prime;
	
	return hash;
}

// Hàm tìm kiếm Rabin-Karp
static int RabinKarpSearch(string text, string pattern)
{
	int prime = 101;        //Số nguyên tố lớn
	int baseValue = 256;    //Số ký tự trong bảng ASCII
	int n = text.Length;
	int m = pattern.Length;

	//Tính base^(m-1) % prime
	int power = 1;
	for(int i = 0; i < m-1; i++)
		power = (power * baseValue) % prime;

	//Lấy giá trị băm tại pattern và đoạn đầu của văn bản
	long patternHash = ComputeHash(pattern, m, prime, baseValue);
	long textHash = ComputeHash(text, m, prime, baseValue);

	//Duyệt qua text
	for(int i = 0; i<= n-m; i++){
		if(patternHash == textHash){

			bool match = true;
			for(int j = 0; j < m; j++){
				if(pattern[j] != text[i+j]){
					match = false;
					break;
				}
			}

			if(match) return i;
		}

		//Cập nhật lại giá trị băm cho đoạn tiếp theo
		if(i < n - m){
			textHash = RecomputeHash(textHash, text[i], text[i+m], baseValue, prime, power);
		}
	}

	return -1;
}
```

---
## **Ưu điểm và nhược điểm:**

**Ưu điểm:**
- Hiệu quả với chuỗi lớn và chuỗi con dài.
- Dễ dàng mở rộng để tìm nhiều chuỗi con cùng lúc
- Sử dụng rolling hash để tính nhanh giá trị băm

**Nhược điểm:**
- Có thể xảy ra xung đột băm (hash collision), cần so sánh từng ký tự để xác nhận
- Phức tạp hơn so với Brute Force việc cài đặt

---
## **Tài liệu:**
1. [Thuật toán Rabin-Karp ứng dụng trong đối sánh mẫu](https://tek4.vn/khoa-hoc/cau-truc-du-lieu-va-thuat-toan/thuat-toan-rabin-karp-ung-dung-trong-doi-sanh-mau)
2. [Thuật toán tìm kiếm Rabin Karp](https://stackjava.com/algorithm/thuat-toan-tim-kiem-rabin_karp.html)
3. [Thuật toán Rabin-Karp](https://dothanhspyb.com/thuat-toan-rabin-karp/)