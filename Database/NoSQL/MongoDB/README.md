![](https://media2.dev.to/dynamic/image/width=1000,height=420,fit=cover,gravity=auto,format=auto/https%3A%2F%2Fdev-to-uploads.s3.amazonaws.com%2Fuploads%2Farticles%2Fyrikqesprn85cp3rjsw7.png)

**MongoDB** là một hệ quản trị cơ sở dữ liệu **NoSQL (Not Only SQL)** phổ biến, được thiết kế để lưu trữ và quản lý dữ liệu dưới dạng tài liệu ==(document-oriented).== Khác với các hệ quản trị cơ sở dữ liệu quan hệ truyền thống (RDBMS) như MySQL hay PostgreSQL, MongoDB sử dụng mô hình dữ liệu linh hoạt, không yêu cầu schema cố định, giúp phát triển ứng dụng nhanh chóng và dễ dàng mở rộng.

![](https://media.geeksforgeeks.org/wp-content/uploads/20200127193216/mongodb-nosql-working.jpg)

### 1. **Đặc điểm nổi bật của MongoDB**
#### a. **Mô hình dữ liệu dạng tài liệu (Document-Oriented)**

- MongoDB lưu trữ dữ liệu dưới dạng các tài liệu JSON (JavaScript Object Notation), được gọi là BSON (Binary JSON). Mỗi tài liệu là một cấu trúc dữ liệu linh hoạt, có thể chứa các trường và giá trị khác nhau.
    
- Ví dụ một tài liệu trong MongoDB:
```json
{
  "_id": "12345",
  "name": "John Doe",
  "age": 30,
  "address": {
    "city": "New York",
    "zip": "10001"
  },
  "hobbies": ["reading", "traveling"]
}
```

#### b. **Không yêu cầu schema cố định**
- Không giống như các cơ sở dữ liệu quan hệ, MongoDB không yêu cầu schema cố định. Bạn có thể thêm hoặc thay đổi cấu trúc dữ liệu một cách linh hoạt mà không cần phải thay đổi schema trước.
#### c. **Khả năng mở rộng (Scalability)**
- MongoDB hỗ trợ mở rộng ngang (horizontal scaling) thông qua cơ chế **sharding**, cho phép phân chia dữ liệu trên nhiều máy chủ để xử lý khối lượng dữ liệu lớn.
- Ngoài ra, MongoDB cũng hỗ trợ mở rộng dọc (vertical scaling) bằng cách tăng cấu hình phần cứng.

#### d. **Hiệu suất cao**
- MongoDB sử dụng cơ chế lưu trữ dữ liệu trên bộ nhớ đệm **(in-memory storage)** và hỗ trợ chỉ mục **(indexing)** để tăng tốc độ truy vấn.
- Các truy vấn phức tạp có thể được thực hiện nhanh chóng nhờ vào cơ chế **aggregation framework**.

#### e. **Hỗ trợ đa nền tảng**
- MongoDB có thể chạy trên nhiều hệ điều hành như Windows, Linux, macOS và hỗ trợ nhiều ngôn ngữ lập trình như Python, Java, Node.js, C#, v.v.

---
### 2. **Các khái niệm cơ bản trong MongoDB**

#### a. **Database**
- Là một container chứa các collection. Mỗi MongoDB server có thể chứa nhiều database.
#### b. **Collection**
- Tương tự như bảng (table) trong cơ sở dữ liệu quan hệ, nhưng không yêu cầu schema cố định. Một collection chứa nhiều tài liệu (documents).
#### c. **Document**
- Là đơn vị cơ bản của dữ liệu trong MongoDB, được lưu trữ dưới dạng BSON. Mỗi document có thể có cấu trúc khác nhau.
#### d. **Field**
- Là một cặp key-value trong document. Ví dụ: `"name": "John Doe"`.
#### e. **Index**
- Giống như chỉ mục trong cơ sở dữ liệu quan hệ, index giúp tăng tốc độ truy vấn dữ liệu.


---
### 3. **Các tính năng nổi bật**

#### a. **Aggregation Framework**
- Cung cấp các công cụ mạnh mẽ để xử lý và phân tích dữ liệu, bao gồm các phép toán như `$match`, `$group`, `$sort`, v.v.

#### b. **Replication**
- MongoDB hỗ trợ replication để đảm bảo ==tính sẵn sàng cao (high availability)==. Dữ liệu được sao chép trên nhiều máy chủ, giúp đảm bảo không bị mất dữ liệu khi có sự cố.

#### c. **Sharding**
- Cho phép phân chia dữ liệu trên nhiều máy chủ để xử lý khối lượng dữ liệu lớn và tăng hiệu suất.

#### d. **GridFS**
- Là một tính năng cho phép lưu trữ và truy xuất các tệp lớn (như hình ảnh, video) trong MongoDB.

---
### 4. **Ưu điểm của MongoDB**

- **Linh hoạt**: Không yêu cầu schema cố định, phù hợp với các ứng dụng có yêu cầu thay đổi cấu trúc dữ liệu thường xuyên.
- **Hiệu suất cao**: Hỗ trợ chỉ mục và lưu trữ dữ liệu trên bộ nhớ đệm.
- **Dễ mở rộng**: Hỗ trợ sharding và replication để xử lý khối lượng dữ liệu lớn.
- **Cộng đồng lớn**: MongoDB có cộng đồng người dùng và tài liệu hỗ trợ phong phú.

---
### 5. **Nhược điểm của MongoDB**

- **Không phù hợp với các ứng dụng yêu cầu ACID mạnh**: MongoDB không hỗ trợ đầy đủ các tính năng ACID (Atomicity, Consistency, Isolation, Durability) như các cơ sở dữ liệu quan hệ.
- **Tốn bộ nhớ**: Do lưu trữ dữ liệu dưới dạng BSON, MongoDB có thể tốn nhiều bộ nhớ hơn so với các cơ sở dữ liệu quan hệ.

---
## 6.**So sánh NoSQL vs SQL:**

1. **Cấu trúc dữ liệu:**
	- SQL: Schema cố định, quan hệ chặt chẽ giữa các bảng
	- NoSQL: Schema linh hoạt, có thể thay đổi theo thời gian

2. **Khả năng mở rộng:**
	- SQL: Thường scale dọc (tăng cấu hình server)
	- NoSQL: Dễ dàng scale ngang (thêm server)

3. **ACID Properties:**
	- SQL: Đảm bảo tính ACID chặt chẽ
	- NoSQL: Có thể đánh đổi một số tính chất ACID để đạt hiệu năng cao hơn

---
### 7. **Ứng dụng của MongoDB**

MongoDB được sử dụng rộng rãi trong các ứng dụng hiện đại, bao gồm:

- **Ứng dụng web và di động**: Nhờ tính linh hoạt và khả năng mở rộng.
- **Phân tích dữ liệu lớn**: Nhờ aggregation framework và khả năng xử lý dữ liệu phi cấu trúc.
- **Hệ thống quản lý nội dung (CMS)**: Nhờ khả năng lưu trữ dữ liệu đa dạng và linh hoạt.
- **IoT (Internet of Things)**: Nhờ khả năng xử lý lượng lớn dữ liệu từ các thiết bị IoT

---
## **Tài liệu:**
1. [MongoDB Cheat Sheet🌱](https://dev.to/burakboduroglu/mongodb-cheat-sheet-1a6a)
2. [MongoDB Tutorial](https://www.w3schools.com/mongodb/index.php)
3. [Học MongoDB - học NoSQL - học database](https://toidicode.com/hoc-mongodb)