### **Biến**
![](https://miro.medium.com/v2/resize:fit:720/format:webp/1*9TfM27i8XsSnTHfAEA92Ww.png)
Biến trong Golang là gì? 
	Biến là một bộ nhớ để lưu trữ dữ liệu.
Cấu trúc gồm có: [Tên biến] = [Giá trị]

Kiểu dữ liệu trong Golang (int, float, string, bool, v.v.)

Để ứng dụng chạy nhanh hơn thì cần phải chọn kiểu dữ liệu phù hợp

---
### Khai báo biến và hằng số:

Quy tắt đặt tên biến trong Golang:

- Tên biến khuyên là nên được biết theo định dạng `camelCase`.
- Có thể khai báo nhiều tên biến cùng lúc

Cách để khai báo biến. gồm có 2 cách:

Cách 1:
```go
var variableName type = value
```
- Lưu ý: Có thể không cần phải khai báo kiểu dữ liệu hoặc giá trị

Cách 2:
```go
variableName := value
```
- Lưu ý:
	- Kiểu dữ liệu của biến được suy ra từ giá trị (Có nghĩa là trình biến dịch quyết định kiểu dữ liệu của biến dựa trên giá trị).
	- Luôn phải khai báo giá trị.

Khi nào nên dùng `cost` thay vì `var`
