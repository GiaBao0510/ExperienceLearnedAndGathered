
**Aggregation Framework** là một công cụ mạnh mẽ trong **MongoDB**, cho phép thực hiện các phép toán phức tạp trên dữ liệu. Nó hoạt động dựa trên khái niệm **pipeline**, trong đó dữ liệu được xử lý thông qua nhiều ==giai đoạn (stages)==. Mỗi giai đoạn thực hiện một phép toán cụ thể và kết quả chuyển tiếp đến giai đoạn kế tiếp.

Tưởng tượng như một dây chuyền sản xuất trong nhà máy: nguyên liệu thô (**dữ liệu ban đầu**) đi qua nhiều công đoạn khác nhau (các **stages**), mỗi công đoạn thực hiện một nhiệm vụ cụ thể, cho đến khi tạo ra sản phẩm cuối cùng (**kết quả mong muốn**).

![](https://miro.medium.com/v2/resize:fit:720/format:webp/1*PgDHQgiLj8JFV1j5m6uTcQ.jpeg)

---
### 1.**Giới thiệu Aggregation Framework:**

##### **Pipeline là gì?**
- Pipeline là một chuỗi các giai đoạn (stages) được áp dụng lên dữ liệu.
- Mỗi giai đoạn thực hiện một phép toán cụ thể, chẳng hạn như: lọc, nhóm, sắp sếp hoặc tính toán.
- **Ưu điểm:**
	- Xử lý dữ liệu theo từng bước có cấu trúc
	- Tối ưu hiệu suất so với việc thực hiện nhiều truy vấn riêng lẻ
	- Giảm lượng dữ liệu cần truyền giữa client và server
	- Cung cấp khả năng thực hiện các phép phân tích phức tạp

##### **Các giai đoạn phổ biến:**
- `$match:` Lọc ra các document phù hợp với điều kiện.
- `$group:` Nhóm các document dựa trên một trường và tính toán giá trị tổng hợp.
- `$sort:` sắp xếp các document.
- `$project:` Chọn các trường cần hiển thị
- `$limit:` giới hạn số lượng document trả về
- `$skip:` Bỏ qua một số lượng document nhất định
- `$unwind`: Tách mảng thành nhiều tài liệu.
- `$lookup`: Thực hiện phép join với collection khác


---
### 2.**Các giai đoạn phổ biến:**

1.**Giai đoạn `$match`:** 
- Lọc ra các document phù hợp với điều kiện.
- ***Ví dụ***: Lọc ra các document có `age` > 25.

```
db.users.aggregate([
	{$match: {age: {$gt:25 } }}
])
```

2.**Giai đoạn` $group`:**
- Nhóm các document dựa trên một trường và tính toán các giá trị tổng hợp
- ***Ví dụ***: Nhóm các document theo trường `city` và đến số lượng document trong mỗi nhóm.

```
db.users.aggregate([
	{$group: {_id_: "_city", total: {$sum:1} }}
])
```

3.**Giai đoạn `$sorrt`:**
- Sắp xếp các document .
- ***Ví dụ***: Sắp xếp các document giảm dần theo `age`.

```
db.users.aggregate([
	{$sort: {age: -1} }
])
```

4.**Giai đoạn `$project`:**
- Chọn các trường cần hiển thị .
- ***Ví dụ***: chỉ hiển thị các trường `name`, `age` bỏ qua trường `_id`.

```
db.users.aggregate([
	{$project: {name: 1, age: 1, _id: 0} }
])
```

5.**Giai đoạn `$limit` và `$skip`:**
- Giới hạn số lượng document và trả về và bỏ qua một lượng document nhất định .
- ***Ví dụ***:  Bỏ qua 10 document đầu tiên và trả về 5 document tiếp theo.
```
db.users.aggregate([
	{$skip: 10},
	{$limit: 5}
])
```

---
### 3.**Ví dụ thực tế:**

#### **Ví dụ 1: Tính toán doanh thu theo từng tháng và sắp xếp theo thứ tự giảm dần**
###### **Cấu trúc collection `orders`**

```
[
  { 
	  _id: 1, product: "A", 
	  quantity: 10, 
	  price: 20, 
	  date: ISODate("2023-10-01") 
  },
  { 
	  _id: 2, product: "B", 
	  quantity: 5, 
	  price: 15, 
	  date: ISODate("2023-10-15") 
  },
  { 
	  _id: 3, 
	  product: "A", 
	  quantity: 7, 
	  price: 20, 
	  date: ISODate("2023-11-05") 
  },
  { 
	  _id: 4, 
	  product: "C", 
	  quantity: 12, 
	  price: 10, 
	  date: ISODate("2023-11-20") 
  },
  { 
	  _id: 5, 
	  product: "B", 
	  quantity: 8, 
	  price: 15, 
	  date: ISODate("2023-12-10") 
  }
]
```

###### **Yêu cầu**
- Tính tổng doanh thu (revenue) cho từng tháng.
- Sắp xếp kết quả theo thứ tự giảm dần của doanh thu.

###### **Câu truy vấn**
```
db.orders.aggregate([
	{
		$group: {
			_id: {year: {$year: "$date"}, month: {$month: "$date"} },
		totalRevenue: { $sum: ${multiply: ["$quantity", "$price"]} }
		},
		{
			$sort: {totalRevenue: -1}
		},
		{
			$poject:{
				_id: 0,
				month: { $dateToString: {format: "%Y-%m", date: {$dateFromParts: {year: "_id.year", month: "_id.month"} } }},
				totalRevenue: 1
			}
		}
	}
])
```

###### **Giải thích**

1. **Giai đoạn `$group`**:
    - Nhóm các document theo năm và tháng (`year` và `month`).
    - Tính tổng doanh thu bằng cách nhân `quantity` với `price` và cộng dồn.
        
2. **Giai đoạn `$sort`**:
    - Sắp xếp kết quả theo `totalRevenue` giảm dần.

3. **Giai đoạn `$project`**:
    - Định dạng lại kết quả để hiển thị tháng dưới dạng `YYYY-MM` và loại bỏ trường `_id`.

###### **Kết quả:**
```
[
  { month: "2023-11", totalRevenue: 260 },
  { month: "2023-10", totalRevenue: 225 },
  { month: "2023-12", totalRevenue: 120 }
]
```

#### **Ví dụ 2: Tìm sản phẩm bán chạy nhất trong mỗi danh mục**

###### **Cấu trúc collection `sales`**
```
[
  { _id: 1, product: "A", category: "Electronics", quantity: 10 },
  { _id: 2, product: "B", category: "Electronics", quantity: 15 },
  { _id: 3, product: "C", category: "Clothing", quantity: 20 },
  { _id: 4, product: "D", category: "Clothing", quantity: 12 },
  { _id: 5, product: "E", category: "Electronics", quantity: 8 }
]
```

###### **Yêu cầu**

- Tìm sản phẩm bán chạy nhất (có `quantity` lớn nhất) trong mỗi danh mục (`category`).

###### **Câu truy vấn**
```
db.sales.aggregate([
	{ $sort: { category: 1 ,quantity: -1}},
	{
		$group: {
			_id: "$category",
			topProduct: {$first: $product},
			maxQuantity: {$max: $quantity}
		}
	},
	{
		$project:{
			_id: 0,
			category: "$_id",
			topProduct: 1,
			maxQuantity: 1
		}
	}
])
```

###### **Giải thích**

1. **Giai đoạn `$sort`**:
    - Sắp xếp các document theo `category` và `quantity` giảm dần.
        
2. **Giai đoạn `$group`**:
    - Nhóm các document theo `category`.
    - Sử dụng `$first` để lấy sản phẩm đầu tiên (sản phẩm có `quantity` lớn nhất) trong mỗi nhóm.
    - Sử dụng `$max` để lấy giá trị `quantity` lớn nhất.
        
3. **Giai đoạn `$project`**:
    - Định dạng lại kết quả để hiển thị thông tin rõ ràng.

###### **Kết quả**
```
[
  { category: "Electronics", topProduct: "B", maxQuantity: 15 },
  { category: "Clothing", topProduct: "C", maxQuantity: 20 }
]
```

### **Ví dụ 3: Tính tỷ lệ phần trăm doanh thu của từng sản phẩm so với tổng doanh thu**

###### **Cấu trúc collection `orders`**

```
[
  { _id: 1, product: "A", quantity: 10, price: 20 },
  { _id: 2, product: "B", quantity: 5, price: 15 },
  { _id: 3, product: "A", quantity: 7, price: 20 },
  { _id: 4, product: "C", quantity: 12, price: 10 },
  { _id: 5, product: "B", quantity: 8, price: 15 }
]
```

#### **Yêu cầu**
- Tính tỷ lệ phần trăm doanh thu của từng sản phẩm so với tổng doanh thu.

###### **Câu truy vấn**
```
db.orders.aggragate:([
	$group: {
		_id: "$product",
		totalRevenue: {$sum: {$multily: [$quantity, $price]} }
	},
	$group: {
		_id: null,
		totalOverallRevenue: {$sum: {"$totalRevenue"} }.
		products: {
			$push {
				product: "$_id", 
				totalRevenue: "$totalRevenue"
			}
		}
	},
	{
		$unwind: "$products"
	},
	$project:{
		_id: 0,
		product: "$products.product",
		totalRevenue: "$products.totalRevenue",
		percentage: {
			$mutity: [
				{$divide: ["$products.totalRevenue", "totalOverallRevenue" ]},
				100
			]
		}
	}
])
```


---
### 4.**Các toán tử tổng hợp phổ biến:**

`$sum`:  Tính tổng các giá trị.
- ***Ví dụ:***
```
{ $sum: "$quantity" }
```

`$avg:` Tính giá trị trung bình
- ***Ví dụ:***
```
{ $avg: "$price" }
```

`$min` và` $max`: Tìm giá trị nhỏ nhất và lớn nhất
- ***Ví dụ:*** 
```
{ $min: "$price" }
{ $max: "$price" }
```

`$push`: Tạo một mảng chứa các giá trị từ các document trong nhóm
- ***Ví dụ:***
```
{ $push: "$product" }
```

`$addToSet`:  Tạo một mảng chứa các giá trị duy nhất từ các document trong nhóm.
- ***Ví dụ:***
```
{ $addToSet: "$product" }
```

---
## **Kết luận:**

Aggregation Framework của MongoDB là một công cụ mạnh mẽ cho phép thực hiện các phép phân tích dữ liệu phức tạp mà không cần phải truy xuất toàn bộ dữ liệu về phía ứng dụng. Bằng cách sử dụng các stages khác nhau trong pipeline, bạn có thể thực hiện nhiều bước xử lý dữ liệu một cách hiệu quả.

Lưu ý khi sử dụng Aggregation Framework:
- Đặt `$match` và `$limit` ở đầu pipeline khi có thể để giảm số lượng tài liệu cần xử lý
- Tránh pipeline quá phức tạp có thể gây áp lực lên server
- Sử dụng index cho các trường được lọc trong `$match` và `$sort`
- Cân nhắc sử dụng tùy chọn `{ allowDiskUse: true }` cho các pipeline xử lý lượng dữ liệu lớn

Với khả năng kết hợp linh hoạt các stages, Aggregation Framework là giải pháp hữu hiệu cho nhiều yêu cầu phân tích dữ liệu trong các ứng dụng thực tế.