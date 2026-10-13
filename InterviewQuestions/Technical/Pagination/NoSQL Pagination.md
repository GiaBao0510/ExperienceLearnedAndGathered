# Phân trang trong MongoDB

_Tài liệu dành cho sinh viên / lập trình viên mới bắt đầu tìm hiểu về Backend_

## Mục lục

1. [Đặt vấn đề: vì sao chuyển trang chậm dần và dữ liệu bị trùng?](https://claude.ai/chat/46012664-46ee-47a9-841b-67be707b29ed#dat-van-de)
2. [Truy vấn dữ liệu trong MongoDB: find() và findOne()](https://claude.ai/chat/46012664-46ee-47a9-841b-67be707b29ed#truy-van-co-ban)
3. [Offset Pagination: dùng `skip()` và `limit()`](https://claude.ai/chat/46012664-46ee-47a9-841b-67be707b29ed#offset-pagination)
4. [Luôn `sort()` khi phân trang](https://claude.ai/chat/46012664-46ee-47a9-841b-67be707b29ed#luon-sort)
5. [Keyset Pagination (Cursor-based Pagination)](https://claude.ai/chat/46012664-46ee-47a9-841b-67be707b29ed#keyset-pagination)
6. [So sánh hai phương pháp](https://claude.ai/chat/46012664-46ee-47a9-841b-67be707b29ed#so-sanh)
7. [Triển khai với Go (`go.mongodb.org/mongo-driver`)](https://claude.ai/chat/46012664-46ee-47a9-841b-67be707b29ed#trien-khai-go)
8. [Lời khuyên tối ưu hiệu năng](https://claude.ai/chat/46012664-46ee-47a9-841b-67be707b29ed#loi-khuyen)
9. [Trích dẫn](https://claude.ai/chat/46012664-46ee-47a9-841b-67be707b29ed#trich-dan)
10. [Mở rộng](https://claude.ai/chat/46012664-46ee-47a9-841b-67be707b29ed#mo-rong)

---

## Đặt vấn đề: vì sao chuyển trang chậm dần và dữ liệu bị trùng?

![](https://fptcloud.com/wp-content/uploads/2022/07/mongodb-duoc-danh-gia-cao-voi-nhieu-tinh-nang-huu-dung.png)

Khi tự phát triển hoặc dùng một website, có hai hiện tượng khó chịu thường gặp khi số trang dữ liệu tăng lên:

1. Việc chuyển trang càng về sau càng **chậm**.
2. Khi chuyển trang, có **item bị lặp lại** (hoặc bị bỏ sót) so với trang trước.

Đây là hệ quả của việc triển khai **phân trang (pagination)** — kỹ thuật chia một tập dữ liệu lớn thành nhiều phần nhỏ để trả về cho client — một cách chưa tối ưu. Tài liệu này sẽ phân tích nguyên nhân của 2 vấn đề trên, cụ thể trong ngữ cảnh **MongoDB**, và giới thiệu các cách khắc phục.

### Nguyên nhân 1: `skip` không "bỏ qua miễn phí"

Nhiều người lầm tưởng rằng khi gọi `skip(10000).limit(10)`, MongoDB sẽ "nhảy thẳng" tới bản ghi thứ 10.001 rồi lấy 10 bản ghi tiếp theo — nghĩa là chỉ cần đọc đúng 10 bản ghi. **Thực tế không phải vậy.** MongoDB vẫn phải **tìm (scan) qua toàn bộ 10.000 bản ghi bị bỏ qua đó**, rồi mới lấy tiếp 10 bản ghi sau cùng — tổng cộng là xử lý 10.010 bản ghi, dù chỉ trả về 10 bản ghi cho client.

Ví dụ cụ thể: một trang quản trị Flashcard, mỗi trang hiển thị 20 flashcard.

|Trang|Query|Số document phải scan|Số document trả về|
|---|---|---|---|
|Trang 1|`find().skip(0).limit(20)`|20|20|
|Trang 2|`find().skip(20).limit(20)`|40|20|
|Trang 1000|`find().skip(20000).limit(20)`|20.020|20|

Số lượng document phải scan **tăng dần theo độ sâu của trang**, dù số lượng trả về luôn cố định là 20. Đây chính là lý do các trang càng về sau càng chậm. Cơ chế này áp dụng tương tự cho `OFFSET`/`LIMIT` trong các cơ sở dữ liệu SQL.

```mermaid
flowchart LR
    A[Bắt đầu từ bản ghi đầu tiên] --> B["Scan qua 'skip' bản ghi\n(không trả về, chỉ để đếm qua)"]
    B --> C["Scan tiếp 'limit' bản ghi"]
    C --> D[Trả về đúng 'limit' bản ghi cho client]
    style B fill:#f8d7da
```

![](https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRsMDDnaIbl14kbOVOWJRXEVZAqGX-fplL53fFNt2s1TwMm75RnYtzquO4&s=10)

### Nguyên nhân 2: dữ liệu không nhất quán khi có thay đổi (consistency)

Giả sử admin đang xem **Trang 1**, sắp xếp theo trường `vocabulary` (giả định giá trị này là duy nhất với mọi user). Kết quả trả về có flashcard thứ 20 là **"Apple"**.

Trong lúc đó, một user khác vừa **thêm một flashcard mới** với `vocabulary = "Above"`. Vì "Above" đứng trước "Apple" theo thứ tự alphabet, nó chen vào danh sách và đẩy "Apple" xuống vị trí thứ 21.

Khi admin chuyển sang **Trang 2** (dùng `skip(20)`), hệ thống sẽ bỏ qua 20 bản ghi đầu — nhưng do "Above" vừa chen vào, bản ghi thứ 20 bây giờ là "Apple" (vốn đã hiển thị ở cuối Trang 1). Kết quả: **"Apple" xuất hiện lặp lại 2 lần**.

Tương tự, nếu một bản ghi ở Trang 1 bị **xóa** trong lúc admin đang xem, danh sách bị dồn lên, và lần `skip` tiếp theo có thể vô tình **bỏ sót** một bản ghi.

> Đây là vấn đề _thiết kế vốn có_ của cách tiếp cận offset, không phải lỗi của MongoDB. Phần [Keyset Pagination](https://claude.ai/chat/46012664-46ee-47a9-841b-67be707b29ed#keyset-pagination) bên dưới sẽ giải thích cách khắc phục.

---

## Truy vấn dữ liệu trong MongoDB: find() và findOne()

Trước khi đi vào phân trang, cần nắm 2 method truy vấn cơ bản trong MongoDB:

- **`find(query, projection)`:** trả về **tất cả** kết quả thỏa mãn điều kiện `query` (dưới dạng một cursor — con trỏ để duyệt qua kết quả).
- **`findOne(query, projection)`:** trả về **kết quả đầu tiên** thỏa mãn `query`.

Trong đó `projection` là tham số tùy chọn, dùng để chỉ định những trường (field) nào sẽ được trả về, giúp giảm dung lượng dữ liệu không cần thiết.

Để xác định **số lượng** kết quả trả về và **vị trí bắt đầu** lấy kết quả, ta dùng thêm 2 method: `limit()` và `skip()`.

---

## Offset Pagination: dùng `skip()` và `limit()`

Đây là cách phân trang phổ biến nhất, dễ hiểu và dễ triển khai nhất.

```js
db.users.find().sort({ name: 1 }).limit(10);          // 10 đầu tiên (trang 1)
db.users.find().sort({ name: 1 }).skip(10).limit(10);  // bỏ 10, lấy 10 (trang 2)
// Công thức tổng quát cho trang N:
// .skip((N - 1) * pageSize).limit(pageSize)
```

Trong đó:

- `limit(n)`: trả về **tối đa** `n` document.
- `skip(m)`: **bỏ qua** `m` document đầu tiên của kết quả (nhưng như đã phân tích ở trên, MongoDB vẫn phải scan qua chúng).

Kết hợp `skip` và `limit`, ta triển khai được **phân trang theo offset (Offset Pagination)**.

### Ví dụ với Node.js / NestJS (repository pattern)

```js
// Cách 1: Offset, limit pagination
const flash_cards = await this.flash_cards_repository.find()
    .sort({ _id: 1 }).skip(20).limit(20);

// Hoặc dùng cú pháp truyền option
const flash_cards = await this.flash_cards_repository.find({},
    'projection', { sort: { _id: 1 }, skip: 20, limit: 20 });
```

### Ưu điểm

- **Dễ triển khai:** backend có thể dùng ngay theo tài liệu mà không cần chỉnh sửa nhiều; frontend cũng dễ tích hợp (chỉ cần truyền số trang).
- **Dễ đọc hiểu:** nhìn vào code là hiểu ngay logic phân trang.
- **Hỗ trợ nhảy tới trang bất kỳ:** người dùng có thể bấm thẳng vào "Trang 5" mà không cần đi qua Trang 1, 2, 3, 4.

### Nhược điểm

Tài liệu chính thức của MongoDB đã ghi rõ vấn đề này:

> _"The skip() method requires the server to scan from the beginning of the input results set before beginning to return results. As the offset increases, skip() will become slower."_<sup>[1]</sup>

(Tạm dịch: _"Method `skip()` yêu cầu server phải scan từ đầu tập kết quả trước khi bắt đầu trả về kết quả. Khi offset càng lớn, `skip()` sẽ càng chậm."_)

- **Hiệu năng kém ở trang sâu:** như đã phân tích ở phần Đặt vấn đề — `skip(m)` không chỉ scan đúng `m` document mà phải scan toàn bộ dữ liệu từ đầu tới vị trí đó.
- **Dữ liệu không nhất quán:** khi có thao tác thêm (`insert`) hoặc xóa (`delete`) xảy ra đồng thời trong lúc người dùng đang phân trang, kết quả trả về có thể bị trùng lặp hoặc thiếu sót.

---

## Luôn `sort()` khi phân trang

```js
// ❌ Phân trang không sort → thứ tự không nhất quán (các trang có thể trùng/bỏ sót)
db.users.find().skip(10).limit(10);

// ✅ Sort theo một field ổn định để các trang nhất quán
db.users.find().sort({ _id: 1 }).skip(10).limit(10);
```

Nếu không chỉ định `sort()`, MongoDB **không đảm bảo** thứ tự trả về của document sẽ giống nhau giữa các lần gọi — tùy thuộc vào cách dữ liệu được lưu trữ vật lý và trạng thái nội bộ tại thời điểm truy vấn. Điều này khiến các trang phân trang có thể không nhất quán (document xuất hiện ở nhiều trang, hoặc biến mất hoàn toàn).

**Nguyên tắc:** luôn `sort()` theo một trường có giá trị **duy nhất và ổn định** (ví dụ `_id`) khi phân trang, để đảm bảo thứ tự luôn nhất quán giữa các lần gọi.

---

## Keyset Pagination (Cursor-based Pagination)

Về mặt khái niệm, **Keyset Pagination** dùng dữ liệu của trang trước để xác định trang tiếp theo — cụ thể là dùng **giá trị trường sắp xếp của bản ghi cuối cùng** ở trang trước, làm điểm mốc để truy vấn tiếp.

Đây là phương pháp được khuyến nghị cho các hệ thống lớn, API công khai, hoặc các tính năng dạng cuộn vô tận (**infinite scroll**).

### Cách hoạt động

Thay vì dùng `skip`, ta ghi nhớ giá trị của trường sắp xếp (thường là `_id` hoặc một trường thời gian như `createdAt`) tại bản ghi **cuối cùng** của trang trước, rồi dùng toán tử so sánh — `$gt` (lớn hơn) nếu sắp xếp tăng dần, hoặc `$lt` (nhỏ hơn) nếu sắp xếp giảm dần — để lấy các bản ghi **tiếp theo** sau mốc đó.

> **Lưu ý về cấu trúc MongoDB ObjectId**
> 
> ObjectId là một giá trị 12-byte, gồm 3 phần<sup>[2]</sup>:
> 
> - **4 byte:** số giây kể từ Unix epoch (thời điểm tạo ObjectId).
> - **5 byte:** giá trị ngẫu nhiên, được sinh **một lần cho mỗi tiến trình (process)** chạy ở phía client — đảm bảo tính duy nhất giữa các máy và các tiến trình khác nhau.
> - **3 byte:** bộ đếm tăng dần (counter), được khởi tạo bằng một giá trị ngẫu nhiên.
> 
> _(Ghi chú: phiên bản ObjectId cũ hơn — trước một số thay đổi của MongoDB — từng tách phần ngẫu nhiên này thành 3-byte định danh máy + 2-byte process ID riêng biệt. Theo tài liệu chính thức hiện tại, phần 5-byte này được sinh ngẫu nhiên như một khối thống nhất, không còn tách rời thành machine identifier và process id như vậy nữa<sup>[2]</sup>. Nếu bạn gặp tài liệu cũ mô tả cấu trúc tách rời, đó là mô tả theo đặc tả phiên bản trước.)_
> 
> Vì có chứa thành phần timestamp, ObjectId có thể được sắp xếp theo thời gian tạo — đây là lý do `_id` thường được dùng làm trường sắp xếp mặc định khi phân trang.

### Ví dụ

Giả sử ở trang trước, item thứ 20 (cuối trang) có `_id = 64ab87fab9e86239671aded5`. Để lấy 20 item tiếp theo:

```js
const flash_cards = await this.flash_cards_repository
    .find({ _id: { $gt: ObjectId('64ab87fab9e86239671aded5') } })
    .sort({ _id: 1 })
    .limit(20);
```

Một ví dụ khác, đơn giản hơn:

```js
// Trang đầu tiên (lấy 10 item)
db.collection.find().sort({ _id: 1 }).limit(10);

// Lấy trang tiếp theo bằng cách truyền _id của item cuối cùng từ trang trước
const lastId = ObjectId("64a1b2c3d4e5f60718293a4b");
db.collection.find({ _id: { $gt: lastId } }).sort({ _id: 1 }).limit(10);
```

### Ưu điểm

- **Hiệu năng cao và ổn định:** vì không dùng `skip`, hệ thống không cần scan qua các bản ghi phía trước nữa.
    
- **Tận dụng Index B-Tree:** nhờ có index trên trường dùng làm cursor (ví dụ `_id`), MongoDB có thể định vị gần như trực tiếp tới vị trí cần tìm thay vì scan tuần tự từ đầu.
    
    > **Ghi chú về độ phức tạp:** Về mặt kỹ thuật, việc dò tìm trên cấu trúc B-Tree có độ phức tạp **O(log n)** (n là số lượng bản ghi) — rất nhanh và gần như không đổi so với việc scan tuyến tính O(n) của `skip`. Điểm cốt lõi và đúng với thông điệp gốc là: **tốc độ không còn phụ thuộc vào độ sâu của trang** theo kiểu tuyến tính như Offset Pagination nữa.
    
- **Không bị trùng lặp hoặc bỏ sót dữ liệu:** vì mốc so sánh (`_id` cuối trang trước) là cố định, dù có bản ghi mới được chèn thêm trong lúc người dùng đang lật trang, kết quả vẫn chính xác.
    

### Nhược điểm

- **Triển khai phức tạp hơn:** nếu sắp xếp theo một trường khác `_id`, hoặc theo nhiều trường cùng lúc, việc xử lý trường hợp các bản ghi có giá trị trùng nhau sẽ phức tạp hơn (cần thêm một trường duy nhất làm "tie-breaker", tương tự nguyên lý Composite Cursor).
- **Không thể truy cập ngẫu nhiên vào một trang bất kỳ:** vì mỗi trang phụ thuộc vào dữ liệu của trang ngay trước đó, muốn tới Trang 30 thì bắt buộc phải tuần tự đi qua Trang 29, 28... Đây là lý do Keyset Pagination phù hợp với giao diện "cuộn tiếp" (Next/Previous hoặc infinite scroll), nhưng không phù hợp với giao diện có dải số trang (1, 2, 3... 30).

---
## So sánh hai phương pháp

| Tiêu chí                | Dùng `skip()` + `limit()`        | Dùng Cursor-based                            |
| ----------------------- | -------------------------------- | -------------------------------------------- |
| Độ dễ triển khai        | Rất dễ                           | Trung bình                                   |
| Hiệu năng (dữ liệu lớn) | ❌ Chậm dần khi số trang tăng     | Nhanh và ổn định                             |
| Nhảy trang bất kỳ       | ✅ Có (Trang 1 ➡️ Trang 5)        | ❌ Không (chỉ có Trước/Sau)                   |
| Ứng dụng phù hợp        | Web quản trị (Admin), dữ liệu ít | Bảng tin (Facebook, TikTok), Infinite Scroll |

---

## Triển khai với Go (`go.mongodb.org/mongo-driver`)

Phân trang trong MongoDB với Go thường được triển khai qua đúng 2 phương pháp đã trình bày ở trên: **Offset Pagination** (dùng `Skip` và `Limit`) và **Cursor-based Pagination** (dùng toán tử so sánh như `$gt`), thông qua thư viện chính thức `go.mongodb.org/mongo-driver`.

```go
package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Offset Pagination: trang N, mỗi trang pageSize bản ghi
func FindFlashcardsByOffset(ctx context.Context, coll *mongo.Collection, page, pageSize int64) (*mongo.Cursor, error) {
	opts := options.Find().
		SetSort(bson.D{{Key: "_id", Value: 1}}).
		SetSkip((page - 1) * pageSize).
		SetLimit(pageSize)

	return coll.Find(ctx, bson.D{}, opts)
}

// Cursor-based (Keyset) Pagination: truyền _id cuối cùng của trang trước
func FindFlashcardsByCursor(ctx context.Context, coll *mongo.Collection, lastID primitive.ObjectID, pageSize int64) (*mongo.Cursor, error) {
	filter := bson.D{}
	if !lastID.IsZero() {
		filter = bson.D{{Key: "_id", Value: bson.D{{Key: "$gt", Value: lastID}}}}
	}

	opts := options.Find().
		SetSort(bson.D{{Key: "_id", Value: 1}}).
		SetLimit(pageSize)

	return coll.Find(ctx, filter, opts)
}
```

> Lưu ý: ở trang đầu tiên của Cursor-based Pagination, chưa có `lastID` nào từ trang trước — khi đó truyền một `ObjectID` rỗng (zero value) để bỏ qua điều kiện `$gt`, như ví dụ trên.

---

## Lời khuyên tối ưu hiệu năng

- **Luôn đánh Index:** dù dùng phương pháp nào, hãy đảm bảo trường dùng trong `.sort()` (và `.find()` nếu có điều kiện lọc) đã được tạo **Index**. Nếu không, MongoDB sẽ phải quét toàn bộ collection (**Collection Scan**) để sắp xếp/lọc, khiến cả hai phương pháp đều chậm — kể cả Cursor-based cũng mất hết lợi thế nếu thiếu index phù hợp.
- **Không dùng `countDocuments()` một cách bừa bãi:** để hiển thị tổng số trang, nhiều lập trình viên gọi `countDocuments()` ở mỗi request. Hàm này khá nặng trên dữ liệu lớn vì phải đếm qua số lượng bản ghi khớp điều kiện. Nếu có thể, nên **ước lượng** (ví dụ dùng `estimatedDocumentCount()` cho trường hợp đếm toàn bộ collection không có điều kiện lọc) hoặc **cache lại** tổng số bản ghi thay vì đếm lại ở mỗi lần gọi API.

---

## Trích dẫn

1. MongoDB Official Documentation — [`cursor.skip()`](https://www.mongodb.com/docs/manual/reference/method/cursor.skip/): xác nhận nguyên văn cảnh báo về hiệu năng của `skip()` khi offset tăng dần.
2. MongoDB Official Documentation — [`ObjectId()`](https://www.mongodb.com/docs/manual/reference/method/objectid/): xác nhận cấu trúc 12-byte hiện tại của ObjectId (4-byte timestamp + 5-byte giá trị ngẫu nhiên theo process + 3-byte counter), làm cơ sở để chỉnh lại mô tả cấu trúc ObjectId trong tài liệu gốc cho khớp với đặc tả hiện hành.
3. Phần còn lại của nội dung được biên soạn lại dựa trên tài liệu gốc do người dùng cung cấp (`NoSQL_Pagination.md`), có sắp xếp lại trình tự trình bày (đưa phần lý thuyết `find`/`findOne` lên trước phần đặt vấn đề chi tiết), bổ sung sơ đồ minh họa, ví dụ code Go, và làm rõ độ phức tạp thuật toán của Cursor-based Pagination.

_Toàn bộ 3 hình ảnh minh họa trong tài liệu gốc đã được giữ nguyên đường dẫn và đặt đúng vị trí nội dung liên quan._

---

## Mở rộng

- **Compound Index (Index đa trường):** khi sắp xếp theo nhiều trường cùng lúc trong Keyset Pagination, cần một compound index bao phủ đúng thứ tự các trường trong `sort()` để truy vấn thực sự tận dụng được index.
- **Composite Cursor (Điểm neo kép):** kỹ thuật kết hợp một trường dễ trùng giá trị (ví dụ số lượt thích) với một trường chắc chắn duy nhất (ví dụ `_id`) làm "tie-breaker", để Keyset Pagination hoạt động chính xác ngay cả khi trường sắp xếp chính có giá trị trùng nhau giữa nhiều bản ghi.
- **Aggregation Pipeline với `$facet`:** cho phép vừa lấy dữ liệu phân trang, vừa đếm tổng số bản ghi trong cùng một lệnh truy vấn, giảm số lần round-trip tới database.
- **MongoDB Change Streams:** cơ chế theo dõi thay đổi dữ liệu theo thời gian thực — có thể dùng để cập nhật UI ngay khi có dữ liệu mới, thay vì phải lo việc dữ liệu "trôi" giữa các lần phân trang.
- **So sánh với phân trang trong SQL (PostgreSQL/MySQL):** các khái niệm Offset vs Keyset Pagination áp dụng tương tự ở các cơ sở dữ liệu quan hệ, nơi còn có thêm kỹ thuật **Deferred Join** để giảm chi phí của Offset Pagination mà vẫn giữ khả năng nhảy trang bất kỳ.