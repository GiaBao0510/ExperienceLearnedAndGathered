#### 📘 **1. Toán tử so sánh cơ bản**
|Toán tử|Mô tả|Ví dụ|Kết quả|
|---|---|---|---|
|`=`|So sánh bằng|`SELECT 1 = 1;`|`true`|
|`!=` hoặc `<>`|So sánh khác|`SELECT 1 != 2;`|`true`|
|`>`|Lớn hơn|`SELECT 5 > 3;`|`true`|
|`<`|Nhỏ hơn|`SELECT 3 < 5;`|`true`|
|`>=`|Lớn hơn hoặc bằng|`SELECT 5 >= 5;`|`true`|
|`<=`|Nhỏ hơn hoặc bằng|`SELECT 4 <= 5;`|`true`|

#### 📘 **2. Toán tử logic**
|Toán tử|Mô tả|Ví dụ|Kết quả|
|---|---|---|---|
|`AND`|Và|`SELECT TRUE AND FALSE;`|`false`|
|`OR`|Hoặc|`SELECT TRUE OR FALSE;`|`true`|
|`NOT`|Phủ định|`SELECT NOT TRUE;`|`false`|

#### 📘 **3. Toán tử điều kiện mở rộng**
|Toán tử|Mô tả|Ví dụ|Ghi chú|
|---|---|---|---|
|`BETWEEN`|Trong khoảng từ ... đến ...|`SELECT 5 BETWEEN 1 AND 10;`|Bao gồm cả 1 và 10|
|`IN`|Trong tập giá trị|`SELECT 'apple' IN ('apple', 'banana');`||
|`NOT IN`|Không nằm trong tập|`SELECT 3 NOT IN (1,2,4);`||
|`IS NULL`|Kiểm tra giá trị NULL|`SELECT NULL IS NULL;`||
|`IS NOT NULL`|Không phải NULL|`SELECT 5 IS NOT NULL;`||
|`LIKE`|So khớp chuỗi|`SELECT 'abc' LIKE 'a%';`|`%` đại diện cho nhiều ký tự|
|`ILIKE`|Giống LIKE nhưng không phân biệt hoa thường|`SELECT 'ABC' ILIKE 'a%';`|PostgreSQL-specific|
|`SIMILAR TO`|So khớp theo biểu thức chính quy SQL|`SELECT 'abc' SIMILAR TO 'a(b|c)%';`|
|`~`|So khớp regex (case-sensitive)|`SELECT 'abc' ~ 'a.*';`||
|`~*`|So khớp regex (không phân biệt hoa thường)|`SELECT 'ABC' ~* 'a.*';`||
|`!~`|Không khớp regex|`SELECT 'abc' !~ 'd.*';`||
|`!~*`|Không khớp regex (không phân biệt)|`SELECT 'abc' !~* 'D.*';`||

#### 📘 **4. Một số toán tử khác (nâng cao)**
| Toán tử                | Mô tả                                  | Ví dụ                                     |
| ---------------------- | -------------------------------------- | ----------------------------------------- |
| `IS DISTINCT FROM`     | So sánh kể cả khi có NULL              | `NULL IS DISTINCT FROM NULL` → `false`    |
| `IS NOT DISTINCT FROM` | Giống như `=` nhưng hoạt động với NULL | `NULL IS NOT DISTINCT FROM NULL` → `true` |
| `EXISTS`               | Kiểm tra tồn tại (dùng với subquery)   | `SELECT EXISTS (SELECT 1 FROM users);`    |

