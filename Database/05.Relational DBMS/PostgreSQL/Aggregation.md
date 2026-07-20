Dưới đây là **bảng tổng hợp các hàm và câu lệnh thường dùng để **gom nhóm dữ liệu (aggregation)** trong **PostgreSQ** — thường được sử dụng với `GROUP BY`, `HAVING` để phân tích và thống kê dữ liệu.

### 📊 **1. Các hàm tổng hợp phổ biến**

|Hàm|Mô tả|Ví dụ|
|---|---|---|
|`SUM()`|Tính tổng|`SELECT SUM(salary) FROM employees;`|
|`AVG()`|Tính trung bình|`SELECT AVG(score) FROM students;`|
|`MIN()`|Giá trị nhỏ nhất|`SELECT MIN(age) FROM users;`|
|`MAX()`|Giá trị lớn nhất|`SELECT MAX(age) FROM users;`|
|`COUNT(*)`|Đếm tất cả dòng|`SELECT COUNT(*) FROM orders;`|
|`COUNT(column)`|Đếm dòng **không NULL** ở một cột|`SELECT COUNT(email) FROM users;`|

---

### 📘 **2. Các hàm tổng hợp mở rộng khác**

|Hàm|Mô tả|Ví dụ|
|---|---|---|
|`STRING_AGG(text, delimiter)`|Ghép chuỗi thành một chuỗi duy nhất|`SELECT STRING_AGG(name, ', ') FROM users;`|
|`ARRAY_AGG(value)`|Gom các giá trị thành mảng|`SELECT ARRAY_AGG(score) FROM students;`|
|`BOOL_AND(condition)`|Trả `true` nếu tất cả điều kiện đều đúng|`SELECT BOOL_AND(active) FROM users;`|
|`BOOL_OR(condition)`|Trả `true` nếu có ít nhất một điều kiện đúng|`SELECT BOOL_OR(active) FROM users;`|
|`VAR_POP()` / `VAR_SAMP()`|Phương sai (population/sample)|`SELECT VAR_SAMP(salary) FROM employees;`|
|`STDDEV_POP()` / `STDDEV_SAMP()`|Độ lệch chuẩn|`SELECT STDDEV_POP(salary) FROM employees;`|
|`PERCENTILE_CONT()`|Tính phần trăm liên tục|Dùng trong `WITHIN GROUP`|
|`PERCENTILE_DISC()`|Tính phần trăm rời rạc|Dùng trong `WITHIN GROUP`|
|`MODE()`|Giá trị xuất hiện nhiều nhất (qua extension)|`SELECT MODE() WITHIN GROUP (ORDER BY col)`|