## **1. Index là gì?**
- Index là một dạng cấu trúc dữ liệu giúp xác định nhanh chóng các records trong table.
- Nếu như không có Index thì SQL sẽ quét toàn bộ bảng để tìm kiếm records cần tìm. Dữ liệu càng lớn thì tốc đọ truy vấn sẽ càng chậm.

## **2. Ưu nhược điểm của Index?**

**_Ưu điểm_:**
- Tăng tốc độ tìm kiếm records theo lệnh where
- Không chỉ giới hạn trong mệnh đề SELECT mà còn UPDATE và DELETE nếu kèm theo câu lệnh WHERE.

**_Nhược điểm_:**
- Khi Sử dụng index thì tốc độ xử lý ghi dữ liệu( UPDATE, DELETE, INSERT) sẽ bị chậm đi. Vì sẽ cập nhật đến index trong table.
- Tốc độc xử lý thì sẽ bị chậm đi cùng tỷ lệ thuận với số lượng index được sử dụng trong bảng.


## **3. Nêu một số kiểu Index trong MySQL cung cấp?**

- MySQL cung cấp 3 kiểu index khác nhau cho data đó là: B-Tree, Hash, R-Tree index

**B-Tree index:**
- thông thường khi chỉ gọi index mà không nêu rõ kiểu thì mặc định sẽ sử dụng loại B-Tree index.
- _Cú pháp_:
```
// Tạo index
CREATE INDEX id_index ON TableName (column_name[,column_name...]) USING BTREE;

//Điều chỉnh thêm index
ALTER TABLE TableName ADD INDEX id_index (column_name[,column_name...]);

//Xóa index
DROP INDEX index_name ON table_name;
```

- **Đặc điểm của B-TREE Index:**
	- dữ liệu được tổ chức và lưu trữ theo dạng tree. Tức là sẽ có root, branch and leaf.
	![BTree](https://d3hi6wehcrq5by.cloudfront.net/itnavi-blog/b-plus-tree-1.png)
	- Giá trị của từng node được tổ chức tăng dần từ trái sang phải.
	- B-Tree index được sử dụng trong các biểu thức so sánh dạng =, > , <, >=, <=, BETWEEN và LIKE. Có thể tối ưu cho câu lệnh **ORDER BY**.
	- khi thực hiện nhiệm vụ truy vấn dữ liệu thì sẽ không quét toàn bộ bảng dể tìm mà việc tìm kiếm trong B-Tree là một quá trình đề quy, bắt đầu từ root node và tìm kiếm tới branch và leaf, đến khi trả về các dữ liệu thỏa mãn với điều kiện.

**Đặc điểm của HASH Index:**
- HashIndex dựa trên giải thuật **Hash Function(hàm băm)**. Tương ứng với mỗi khối dữ liệu (index) sẽ sinh ra một bucket key(giá trị băm) để phân biệt
- _Cú pháp:_
```
//Tạo index
CREATE INDEX id_index ON TableName(column_name[,column_name...]) USING HASH;

//Hoặc

ALTER TABLE TableName ADD INDEX id_index (column_name[,column_name...]) USING HASH;
```

- _**Các đặc điểm của Hash Index:**_
	![HashIndex](https://d3hi6wehcrq5by.cloudfront.net/itnavi-blog/hU4Tc.png)
	- Khác với B-Tree, thì Hash Index nên chỉ được sử dụng trong các biểu thức toán tử như = và <>. Không sử dụng cho toán tử tìm kiếm khoảng giá trị như > hoặc <.
	- Không có thể tối ưu hóa toán tử **ORDER BY** bằng việc sử dụng Hash index vì nó không tìm kiếm được phần tử tiếp theo trong Order.
	- Hash có tốc độ nhanh hơn kiểu B-Tree

## **4. Các kiểu Index tương ứng với Storage Engine?**
- Việc chọn index theo kiểu B-Tree hay Hash thì ngoài yếu tố mục đích sử dụng index thì nó còn phụ thuộc vào việc Storage Engine có hỗ trợ loại đó hay không

| Storage Engine | Các Kiểu Index được hỗ trợ |
| -------------- | -------------------------- |
| InnoDB         | BTREE                      |
| MyISAM         | BTREE                      |
| MEMORY/HEAP    | HASH, BTREE                |
| NDB            | HASH, BTREE                |
## **5. Giải thích Primary Key Index?**

_**Mô tả:**_
- Đây là loại chỉ mục đặt biệt, được tạo ra tự động khi định nghĩa **Primary Key** trên một hoặc một nhóm cột trong bảng.
- Giá trị của cột được đánh dấu **Primary Key** phải là **duy nhất** và **không được null.**
**Đặc điểm:**
- Mỗi bảng chỉ có một **Primary Key**.
- Dùng để xác nhận duy nhất một bản ghi trong bảng.
**Ví dụ:**
```
CREATE TABLE employee(
	emplotee_id PRIMARY KEY,
	name VARCHAR(50)
);
```

## **6.Giải thích Unique Index?**
_**Mô tả:**_
- Đảm bảo các giá trị trong cột hoặc nhóm cột trong bảng là có giá trị duy nhất. Điểm khác với **Primary Key** là nó có thể chứa **NULL.**
- Giá trị của cột được đánh dấu **Primary Key** phải là **duy nhất** và **không được null.**
**Đặc điểm:**
- Mỗi bảng có thể có nhiều **Unique index**.
- Thường để ngăn sự trùng lặp trong bảng ghi.
**Ví dụ:**
```
CREATE TABLE employee(
	email VARCHAR(100) UNIQUE KEY,
	phone_number VARCHAR(10) UNIQUE KEY
);
```

## **7. Giải thích về Regular Index (Index thông thường)?**
_**Mô tả:**_
- là một loại chỉ mục cơ bản nhất, dùng để tăng tốc tìm kiếm trên một hoặc nhiều cột.
**Đặc điểm:**
- Có thể áp dụng trong các cột thường xuyên sử dụng trong mệnh đề ==WHERE==, ==ORDER BY== hoặc ==GROUP BY==.
**Ví dụ:**
```
CREATE INDEX idx_name ON employee(name);
```

## **8. Giải thích về Full-Text Index?**
_**Mô tả:**_
- Dùng cho các tìm kếm văn bản toàn văn, thường áp dụng lên các cột có kiểu ==CHAR, VARCHAR or TEXT.==
- Hỗ trợ cho việc tìm kiếm phức tạp như cụm từ hoặc các từ trong văn bản.
**Đặc điểm:**
- Thích hợp cho các cột chứa dữ liệu văn bản lớn.
- Không thể áp dụng trên tất cả các strorage engine(Chỉ hỗ trợ InnoDB và MyISAM).
**Ví dụ:**
```
CRETAE TABLE articles(
	id INT PRIMARY KEY,
	title VARCHAR(255),
	content TEXT,
	FULLTEXT(title, content)
);

SELECT * FROM articles
WHERE MATCH(title, content) AGAINST('MySQL indexing');
```

## **9. Giải thích về Composite Index?**
_**Mô tả:**_
- Là chỉ mục được tạo trên nhiều cột trong bảng thay vì 1 cột.
- HIệu quả khi các truy vấn sử dụng mệnh đề ==WHERE== hoặc sắp xếp theo thứ tự trong giống như thứ tự cột **Composite Index.** 
**Đặc điểm:**
- Cần sắp xếp thứ tự các cột trong chỉ mục cẩn thận. Vì thứ tự sẽ ảnh hưởng đến hiệu suất.
**Ví dụ:**
```
CREATE TABLE order(
	order_id INT PRIMARY KEY,
	customer_id INT,
	product_id INT,
	order_date DATE,
	INDEX idx_customer_order_product_date(customer_id,product_id,order_date)
);
```
## **10. Giải thích Foreign Key Index?**
_**Mô tả:**_
- Tự động tạo khi định nghĩa một Forgein Key trên 1 cột
- Đảm bảo tính toàn vẹn tham chiếu giữa bảng chính và bảng phụ
**Đặc điểm:**
- Dùng để nhanh chóng xác thực và liên kết dữ liệu giữa các bảng.
**Ví dụ:**

```
CREATE TABLE orders ( 
	order_id INT PRIMARY KEY, 
	customer_id INT, 
	FOREIGN KEY (customer_id) REFERENCES customers(customer_id)
);
```
