
Tại đây sẽ tìm hiểu về các kỹ thuật truy vấn nâng cao trong MongoDB, bao gồm với mảng, sử dụng biểu thức chính quy (regex), phân trang và sắp kết quả

![](https://www.hostinger.com/tutorials/wp-content/uploads/sites/2/2021/08/what-is-a-query.jpg)

---
### **Truy vấn với mảng:**

MongoDb có hỗ trỡ lưu trữ và truy vấn với trường có kiểu mảng. Sau đây là các toán tử phổ biến theo tác với mảng:

##### 1.**Toán tử `$all`:**
- Khớp các document có mảng và chứa tất cả các phần tử được chỉ định.
- ***Ví dụ:*** tìm hai giá trị `"mongodb"` & `"database"` trong trường `tags`

```
db.users.find(
	{tags: {$all: ["mongodb", "database"]}}
)
```

##### 2.**Toán tử `$elemMatch`:**
- Khớp với các document có ít nhất một phần tử trong mảng thỏa mãn tất cả các điều kiện được chỉ định.
- ***Ví dụ:*** tìm các document có trường score chứa ít nhất một phần tử > 80 & < 90
```
db.user.find({
	scores: {$elemMatch: {$gt:80 ,$lt:90 }}
})
```

##### 3.**Toán tử `$size`:**
- Khớp với các document có mảng với kích thước được chỉ định
- ***Ví dụ:*** tìm các document có trường `tags` chứa 3 phần tử

```
db.users.find({
	tags: {$size: 3}
})
```

##### 4.**Truy vấn lồng nhau:**
- MongoDB hỗi trợ truy vẫn các trường lồng nhay trong mảng hoặc document.
- ***Ví dụ:*** tìm các document có trường `address.city` bằng `New York`.

```
db.users.find({
	"address.city": "New York"
})
```

---
### **Truy vấn với regex:**

MongoDb hộ trợ sử dụng biểu thức chính quy (regex) để tìm kiếm dữ liệu dự trên chuỗi.

##### 1.**Cú pháp cơ bản:**
- Sử dụng toán tử `$regex` để khớp các giá trị với biểu thức chính quy
-  ***Ví dụ:*** tìm các document có trường `name` bắt đầu từ chứ `A` (Không phân biệt hoa thường)

```
db.users.find({
	name: {$regex: /^A/, $option:"i" }
})
```

##### 2.**Các tùy chọn (options):**
- `i `: Không phân biệt hoa thường
- `m`: Khớp nhiều dòng
- `x`: Bỏ qua khoảng trắng và các ký tự comment.
---
### **Phân trang và sắp xếp:**

Khi làm việc với lượng dữ liệu lớn, việc phân trang và sắp xếp kết quả là rất quan trọng. Trong MongoDB cung cấp các phương thức để thực hiện điều này.

##### 1.**Phân trang với `limit()` và `skip()`:**
- `limit(n)`: giới hạn số lượng kết quả trả về.
- `skip(n)`: bỏ qua `n` kết quả đầu tiên.
- ***Ví dụ:*** Bỏ qua 10 document đầu tiên và trả về 15 document tiếp theo

```
db.users.find().skip(5).limit(15)
```

##### 2.**Sắp xếp kết với `sort()`:**
- `sort({ field: 1})`: sắp xếp tăng dần theo trường `field`.
- `sort({ field: -1})`: sắp xếp giảm dần theo trường `field`.
-  ***Ví dụ:*** trả về 5 document có tuổi lớn nhất, và sắp xếp giảm dần theo trường `age`

```
db.users.find().sort({age: -1}).limit(5)
```

##### 3.**Kết hợp phân trang và sắp xếp:**

- **Ví dụ**: Trả về 10 document tiếp theo, sắp xếp tăng dần theo trường `age`, sau khi bỏ qua 20 document đầu tiên.

```
db.users.find().sort({ age: 1 }).skip(20).limit(10)
```
    

---
### ***Ví dụ tổng hợp:***

Giả sử dữ liệu có trong colletion `usser` với các document như sau:

```
[
  { name: "Alice", age: 25, tags: ["mongodb", "database"], scores: [85, 90, 78] },
  { name: "Bob", age: 30, tags: ["programming", "database"], scores: [88, 92, 80] },
  { name: "Charlie", age: 35, tags: ["mongodb", "programming"], scores: [75, 82, 88] },
  { name: "David", age: 40, tags: ["database", "nosql"], scores: [90, 85, 95] }
]
```

Ví dụ: **Truy vấn với mảng**:
```
db.users.find({ tags: { $all: ["mongodb", "database"] } })
```

Ví dụ: **Truy vấn với regex**:
```
db.users.find({ name: { $regex: /^A/, $options: "i" } })
```

Ví dụ: **Phân trang và sắp xếp**:
```
db.users.find().sort({ age: -1 }).skip(1).limit(2)
```