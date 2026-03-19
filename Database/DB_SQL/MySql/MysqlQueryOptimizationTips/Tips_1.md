
### **Quá trình MySQL thực thi một câu truy vấn:**

- ***Bước 1:*** MySQL gửi câu truy vấn lên server.
- ***Bước 2:*** MySQL kiểm tra [query cache](https://dev.mysql.com/doc/refman/5.7/en/query-cache.html). Nếu kết quả đã tồn tại thì trả về. Ngược lại thì sẽ thực thi bước tiếp theo. 
- ***Bước 3:*** MySQL thưc thi parse câu truy vấn, để đưa ra plan thực thi tối ưu
- ***Bước 4:*** Plan được thực thi, gọi đến [Storage engines](https://dev.mysql.com/doc/refman/8.0/en/storage-engines.html)
- ***Bước 5:*** [Storage engines](https://dev.mysql.com/doc/refman/8.0/en/storage-engines.html) trả về kết quả là row dữ liệu.

Cách viết câu truy vấn có thể ảnh hưởng đến thời gian thực thi của câu truy vấn:
*có thể làm cho câu truy vấn được thực thi nhanh hoặc chậm đi.*
- 