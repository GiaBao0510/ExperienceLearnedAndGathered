## **Danh sách các toán tử trong MongoDB:**

##### **Toán tử so sánh:**

Các toán tử này được sử dụng để so sánh giá trị của các trường truy vấn.

| Toán tử | Mô tả                                                            |
| ------- | ---------------------------------------------------------------- |
| `$eq`   | Khớp với các giá trị bằng giá trị chỉ định                       |
| `$gt`   | Khớp với các giá trị lớn hơn giá trị chỉ định                    |
| `$gte`  | Khớp với các giá trị lớn hơn hoặc bằng giá trị chỉ định          |
| `$in`   | Khớp với bất kỳ giá trị nào được chỉ định trong một mảng.        |
| `$lt`   | Khớp với các giá trị nhỏ hơn giá trị chỉ định                    |
| `$lte`  | Khớp với các giá trị nhỏ hơn hoặc bằng giá trị chỉ định          |
| `$ne`   | Khớp với tát cả các giá trị không bằng với giá trị được chỉ định |
| `$nin`  | Không khớp với bất kỳ giá trị nào được chỉ định trong một mảng   |
***Ví dụ:***  Tìm các trường `age` có giá trị > 25:

```
db.collection.find({age: {$gt: 25}})
```

---
##### **Toán tử Logic:**

Các toán tử này được sử dụng để kết hợp nhiều điều kiện trong truy vấn.

| Toán tử | Mô tả                                                                       |
| ------- | --------------------------------------------------------------------------- |
| `$and`  | Kết hợp các điều kiện bằng logic AND.                                       |
| `$or`   | Kết hợp các điều kiện bằng logic OR.                                        |
| `$not`  | Đảo ngược kết quả của điều kiện.                                            |
| `$nor`  | Kết hợp các điều kiện bằng logic NOR (Không khớp với bất kỳ điều kiện nào). |
***Ví dụ:***  Tìm các trường `age` có giá trị > 23 && < 30:

```
db.collection.find({
	$and: [
		{age: {$gt: 25}}, 
		{age: {$gt:30}} 
	]
})
```

---
##### **Toán tử kiểm tra trường:**

Các toàn tử này dùng để kiểm tra sự tồn tại hoặc kiểu dữ liệu của các trường.

| Toán tử   | Mô tả                                                         |
| --------- | ------------------------------------------------------------- |
| `$exists` | Khớp các document có trường được chỉ định.                    |
| `$type`   | Khớp các document có trường thuộc kiểu dữ liệu được chỉ định. |
***Ví dụ:***  Tìm các document có trường `age`:

```
db.collection.find({
	age: {$exists: true}
})
```

---
##### **Toán tử đánh giá:**

Các toán tử này dùng để thực hiện các phép tính toán phức tạp trong truy vấn

| Toán tử       | Mô tả                                                  |
| ------------- | ------------------------------------------------------ |
| `$expr`       | Cho phép sử dụng các biểu thức tổng hớp trong truy vấn |
| `$jsonSchema` | Xác thực document dựa trên lượt đồ JSON.               |
| `$mod`        | Thực hiện phép toán modulo trên giá trị của trường.    |
| `$regex`      | Khớp các giá trị với biểu thức chính quy               |
| `$text`       | Thực hiện tìm kiếm văn bản                             |
| `$where`      | Khớp các document thỏa mản biểu thức JavaScript.       |
***Ví dụ:***  Tìm các document có `price` > `discount` :

```
db.collection.find({
	$expr: { $gt: ["price","discount"] }
})
```

---
##### **Tóan tử mảng:**

Các toán tử này được sử dụng thể thao tác với các trường có kiểu mảng.

| Toán tử      | Mô tả                                                                   |
| ------------ | ----------------------------------------------------------------------- |
| `$all`       | Khớp các mảng chứa tất cả các phần tử được chỉ định                     |
| `$elemMatch` | Khớp các document có phần tử trong mảng thoản mãn tất cả các điều kiện. |
| `$size`      | Khớp các document có các mảng với kích thước được chỉ định.             |
| `$pop`       | Xóa phần tử đầu/cuối                                                    |
***Ví dụ:***  Tìm các document có 2 tag `mongodb` và `database` :

```
db.collection.find({
	tags: {$all: ["mongodb" , "database"]}
})
```

---
##### **Toán tử cập nhật:**

Các toán tử này được sử dụng để cập nhật các trường document.

| Toán tử        | Mô tả                                                    |
| -------------- | -------------------------------------------------------- |
| `$set`         | Đặt giá trị của một trường                               |
| `$unset`       | Xóa một trường khỏi document                             |
| `$inc`         | Tăng giá trị của một trường theo số lượng được chỉ định. |
| `$push`        | Thêm một phần tử vào mảng                                |
| `$pull`        | Xóa các phần tử khỏi mảng.                               |
| `$addToSet`    | Thêm phần tử vào mảng nếu nó chưa tồn tại.               |
| `$mul`         | Nhân giá trị                                             |
| `$min`         | Cập nhật nếu nhỏ hơn                                     |
| `$max`         | Cập nhật nếu lớn hơn                                     |
| `$rename`      | Đổi tên trường                                           |
| `$currentDate` | Đặt ngày hiện tại                                        |
***Ví dụ:***  Cập nhật tuổi của GiaBao thành 24 :

```
db.collection.updateOne({
	name: "GiaBao", $set: {age: 24}
})
```

---
##### **Toán tử không gian địa lý:**

Các toán tử này được sử dụng để làm việc với dữ liệu địa lý

| Toán tử          | Mô tả                                                      |
| ---------------- | ---------------------------------------------------------- |
| `$geoWithin`     | Khớp các document nằm trong một hình học GeoJSON.          |
| `$geoIntersects` | Khớp các document hình học giao nhau với hình học GeoJSON. |
| `$near`          | Khớp các document gần một điểm                             |
| `$nearSphere`    | Khớp các document gần một điểm trên hình cầu.              |
***Ví dụ:*** Tìm các document gần tọa độ (-73.9667, 4078)

```
db.collection.find({
	location: {
		$near: {
			$geometry: type: "Point", 
			coordinates: [-73.9667, 4078]
		}
	}
})
```

---
##### **Toán tử bit:**

Các toán tử này được sử dụng để làm việc với dữ liệu bit.

| Toán tử        | Mô tả                                                  |
| -------------- | ------------------------------------------------------ |
| `$bitsAllSet`  | khớp các giá trị có tất cả các bit được chỉ định là 1. |
| `$bitAllSet`   | khớp các giá trị có tất cả các bit được chỉ định là 0  |
| `$bitAnySet`   | khớp các giá trị có bất kỳ bit nào được chỉ định là 1  |
| `$bitAnyClear` | khớp các giá trị có bất kỳ bit nào được chỉ định là 0  |
***Ví dụ:*** Tìm các document có bit 1 và 5 được bật

```
db.collection.find({
	flags: {$bitsAllSet: [1,5]}
})
```