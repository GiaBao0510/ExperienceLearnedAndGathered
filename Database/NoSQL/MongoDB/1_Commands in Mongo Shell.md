#### **Khởi động MongoDB Shell:**

```
mongosh
```
- Lệnh này khởi tạo MongoDB Shell, cho phép tương tác với MongoDB Server.

---
#### **HIển thị các cơ sở dữ liệu (databases):**
```
show dbs
```
- Lệnh này giúp hiển thị danh sách tất cả các cơ sở dữ liệu (databases) trên MongoDB Server.

---
#### **Chuyển đổi sang hoặc tạo cơ sở dữ liệu mới:**
```
use [database_name]
```
- Chuyển đến cơ sở dữ liệu đã được chỉ định.
- Nếu cơ sở dữ liệu không tồn tại thì sẽ tạo mới

---
#### **Hiển thị các collections trong cơ sở dữ liệu:**
```
show collections
```
- Hiển thị danh sách tất cả các collection đang tồn tại trong cơ sở dữ liệu
---
#### **Tạo mới collection:**
```
db.createCollection("collection_name")
```
- Tạo mới collection trong cơ sở dữ liệu hiện tại
- ***Ví dụ:***
	```
	db.createCollection("Users")
	```

---
#### **Tham chiếu đến collection:**
```
db.collection
```
- Lệnh trên tham chiếu đến collection cụ thể
- ***Ví dụ:***
	```
	db.Users
	```

---
#### **Truy vấn dữ liệu trong collection:**
```
db.collection.find(query, projection)
```
- Truy vấn và hiển thị các ==documents== muốn lấy dựa trên điều kiện ==query==.
- **query**: Điều kiện tìm kiếm (Ví dụ: `{ job: "developers" }` ).
- **projection**: Các trường cần hiển thị (Ví dụ: `{name: "bao", age: 23 }` ).
- ***Ví dụ:***
	- Lấy tất cả document trong collection:
	```
	db.users.find()
	```
	- Lấy các document có `age =  23`:
	```
	db.users.find({age: 23})
	```
---
#### **Truy vấn một document duy nhất:**
```
db.collection.findOne(query, projection)
```
- Truy vấn và hiển thị document đầu tiên mà nó tìm thấy.
- ***Ví dụ:*** 
```
db.users.findOne({age: 23})
```

---
#### **Thêm document vào collection:**
```
db.collection.insertOne(document)
```
- Thêm một document vào collection.
- MongoDB sẽ tự động tạo và một **`_id`** duy nhất cho document.
- ***Ví dụ:***
```
	db.users.insertOne(
		{name:"Pham Gia Bao", age: 23}
	)
```

---
#### **Thêm nhiều documents vào collection:**
```
db.collection.insertMany([doc1, doc2,...])
```
- Thêm nhiều document vào collection cùng lúc
- ***Ví dụ:***
```
db.users.insertMany(
		{name:"Pham Gia Bao", age: 23},
		{name:"Nguyen Van A", age: 24},
		{name:"Nguyen Van B", age: 25}
	)
```

---
#### **Xóa document:**

- Xóa một **document**, được tìm thấy đầu tiên dựa trên điều kiện **`query`**.
```
db.collection.deleteOne(query)
```
- Xóa nhiều **document**, được tìm thấy dựa trên điều kiện **`query`**.
```
db.collection.deleteMany(query)
```
- Xóa tất cả các document trong collection:
```
db.collection.deleteMany({})
```

---
#### **Xóa collection:**
```
db.collection.drop()
```
- Xóa toàn bộ ==collections== trong ==CSDL==.

***Ví dụ:***
```
db.user.drop()
```

---
#### **Cập nhật document:**

- Cập nhật một document đầu tiên được tìm thấy dựa trên `query`.
```
db.collection.updateOne(query, update)
```
- Cập nhật nhiều document dựa trên `query`.
```
db.collection.update.updateMany(query, update)
```

***Ví dụ:*** Cập nhật người dùng có tên "GiaBao"
```
db.user.updateOne({name: "GiaBao"}, {$set:{age: 25} })
```

---
#### **Tạo chỉ mục (index):**
```
db.collection.create({field: 1})
```
- Tạo chỉ mục trên một trường để tăng tốc độ truy vấn

***Ví dụ:*** 
```
db.user.createIndex({name: 1})
```

---
#### **Đếm số lượng documents:**
```
db.collection.countDocuments(query)
```
- Đếm số lượng document phù hợp với điều kiện `query`.

***Ví dụ:*** 
```
db.users.countDocuments({age: {$gt:25} })
```

---
#### **Sắp xếp và giới hạn kết quả hiển thị:**
```
db.collection.find().sort({field: 1}).limit(10)
```
- Sắp xếp kết quả dựa theo trường `field` (1: tăng dần, -1: giảm dần) và giới hạn số lượng trả về

***Ví dụ:*** 
```
db.user.find().sort({age: -1}).limit(15)
```

---
#### **Sử dụng Aggregation Framework:**
```
db.collection.aggregate([pipeline])
```
- thực hiện phép tính toán phức tạp trên dữ liệu.

***Ví dụ:*** 
```
db.users.aggregate([
	{$match: {age: {$gt: 25}} },
	{$group: {_id: "$city", total: {$sum:1} }}
])
```

---
## **Lấy danh sách giá trị duy nhất trong một trường:**
```

db.collection.distinct("field")
```
