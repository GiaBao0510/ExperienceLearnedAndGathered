Ở tài liệu này chúng ta sẽ tập trung ở MongoDB là chính

Có bao giờ trong quá trình tự mình phát triển hoặc truy cập các website các bạn nhận thấy rằng việc di chuyển qua các trang chậm lại khi số trang ngày càng lớn 🧐 hay khi chuyển trang thì item ở trang trước bị lặp lại thêm lần nữa ở trang sau.

Đó là vấn đề mà đa số các website gặp phải khi tiến hành pagination một cách không hiệu quả 😤. Hôm nay chúng ta sẽ cùng tìm hiểu đôi chút về vấn đề này nói chung và cụ thể dùng trong MongoDB nói riêng.

#### limit và skip — phân trang

```
db.users.find().sort({ name: 1 }).limit(10);          // 10 đầu tiên (trang 1)
db.users.find().sort({ name: 1 }).skip(10).limit(10); // bỏ 10, lấy 10 (trang 2)
// trang N: .skip((N-1) * pageSize).limit(pageSize)
```
`limit(n)` trả về tối đa n document; `skip(m)` bỏ qua m document đầu. Kết hợp lại chúng triển khai **phân trang theo offset** (trang N qua skip + limit).

#### Luôn sort khi phân trang

```
// ❌ phân trang không sort → thứ tự không nhất quán (các trang có thể trùng/bỏ sót)
db.users.find().skip(10).limit(10);
// ✅ sort theo một field ổn định để các trang nhất quán
db.users.find().sort({ _id: 1 }).skip(10).limit(10);
```

Không có `sort`, thứ tự document không được đảm bảo, nên các trang phân trang có thể không nhất quán. Luôn sort theo một field ổn định (như `_id`) khi phân trang.

**Phân trang (Pagination) trong MongoDB với Golang** thường được thực hiện qua hai phương pháp chính: **Offset Pagination** (sử dụng `Skip` và `Limit`) và **Cursor-based Pagination** (sử dụng toán tử so sánh toán học như `$gt`).

Dưới đây là chi tiết và cách triển khai từng phương pháp bằng thư viện chính thức `go.mongodb.org/mongo-driver`.

---
# Đặt vấn đề 📜

Pagination là chức năng mà các bạn thường thấy trong hầu hết các ứng dụng, giúp chia số lượng data trả về thành nhiều trang giúp hạn chế payload cũng như tăng trải nhiệm của người dùng.

1️⃣ Nguyên nhân của vấn đề đầu tiên mà mình đề cập ở trên là do khi chúng ta tiến hành dùng `skip` trong MongoDB để di chuyển đến các trang sau, nó sẽ không vì chúng ta nhập vào `skip: 10000, limit: 10` mà bỏ qua 10000 item trước đó và lấy 10 item, thay vào đó nó sẽ lấy ra luôn cả 10000 item đó kèm với 10 item nữa là 10010 item.

Lấy ví dụ ở page Quản lí Flash card của admin mỗi trang chúng ta trả về 20 Flash card. Quá trình chuyển trang diễn ra như sau:

- Ở **trang 1** query sẽ có dạng `find().skip(0).limit(20)`, database sẽ **tìm 20 documents** và **trả về 20 documents** cho chúng ta.
- Khi chuyển qua **trang 2** query sẽ thay đổi thành `find().skip(20).limit(20)`, database sẽ **tìm 40 documents** và **trả về 20 documents** cho chúng ta.
- Tương tự khi chuyển đến **trang 1000** query sẽ là `find().skip(20000).limit(20)`, database sẽ **tìm 20020 documents** và trả về **20 documents**.

✴️ Đó chính là cách `skip` và `limit` hoạt động trong MongoDB (`offset` và `limit` trong SQL cũng sẽ tương tự). Nhìn qua thì các bạn cũng đã hiểu tại sao kết quả ở các trang sau lại tốc độ chậm hơn các trang trước.

2️⃣ Nguyên nhân của vấn đề thứ hai liên quan tới **tính consistency**, ví dụ admin đang query dữ liệu **ở trang 1** như trên và được **sort theo vocabulary** (giả sử vocabulary là unique với tất cả user), trong kết quả trả về **Flash card thứ 20 là Apple**. Tuy nhiên trong lúc đó có một **user** nào đó **thêm vào một Flash card mới** có vocabulary là **Above** và nó làm cho **Apple bị đẩy sang vị trí thứ 21** (vì về mặt thứ tự thì Above sẽ đứng trước Apple), do đó khi admin **chuyển sang trang thứ 2 thì Apple lại xuất hiện thêm một lần nữa**. Tương tự với trường hợp sửa hoặc xóa cũng có thể làm cho dữ liệu hiển thị không được đồng nhất.

---

# Lý thuyết 📚

Chúng ta sẽ cùng điểm sơ qua các cách pagination thông dụng hiện nay.

## 1. Offset Pagination (Sử dụng Skip & Limit)

Đây là cách dùng phổ biến trong đa số các website, dùng `skip`, `limit` (hay`offset`, `limit` trong SQL) để phân trang. Chúng ta vẫn hay thường dùng như bên dưới trong MongoDB:

```
// Cách 1️⃣: Offset, limit pagination
const flash_cards = await this.flash_cards_repository.find()
    .sort({ _id: 1 }).skip(20).limit(20)
// Hoặc
const flash_cards = await this.flash_cards_repository.find({},
    'projection', { sort: { _id: 1 }, skip: 20, limit: 20 })
```

### Ưu điểm ✅

- **Dễ triển khai**: Back-end có thể đọc docs và dùng ngay mà không cần phải chỉnh sửa. Front-end cũng dễ tích hợp vào UI.
- **Dễ đọc hiểu**: khi nhìn vào có thể dễ dàng hiểu ngay.

### Nhược điểm 📛

Thật ra nếu như các bạn đọc tài liệu chính gốc từ MongoDB sẽ thấy có nhắc đến vấn đề này:

> The skip() method requires the server to scan from the beginning of the input results set before beginning to return results. **As the offset increases, skip() will become slower**.

- **Hiệu năng kém**: cơ chế hoạt động `skip` là lấy tất cả các document được trả về từ kết quả của chaining method trước đó (ở trên là `find`) và scan từ đầu cho đến khi đến vị trí `skip` sẽ trả về số lượng documents tương ứng chứ không phải nhập vào 20 là `skip` chỉ scan 20 documents.
- **Dữ liệu không nhất quán**: như đã đề cập ở phần trên, khi admin đang gọi API nếu có user thao tác **add** hoặc **delete** thì dữ liệu trả về có thể bị duplicate (do add) hoặc hiển thị thiếu (do delete).

## 2. Keyset pagination (Cursor-based pagination)

Về mặt khái niệm thì Keyset pagination dùng page trước đó để trả về page tiếp theo, hay nói cách khác là dùng item cuối cùng để trả về page tiếp theo sau item đó.

Tùy theo các property được sort mà chúng ta sẽ có cách dùng tương ứng với các property đó. Ví dụ sort theo `_id` MongoDB thì dùng `_id` của item cuối cùng của trang trước kết hợp với `limit` và `FilterQuery` (hay `where` trong SQL).

> **Có thể các bạn chưa biết**: MongoDB ObjectId là một cấu trúc 12 byte dữ liệu bao gồm:
>
>- 4-byte biểu thị số giây kể từ Unix epoch.
>- 5-byte giá trị ngẫu nhiên được tạo một lần cho mỗi quá trình. Giá trị ngẫu nhiên này là unique cho machine và process.
    - 3-byte machine identifier.
    - 2-byte process id.
>- 3-byte incrementing counter, được khởi tạo thành một giá trị ngẫu nhiên.
>
Do thành phần của nó có liên quan tới timestamp nên chúng ta có thể tiến hành sort ObjectId.

Ví dụ ở page trước chúng ta có item thứ 20 có `_id=64ab87fab9e86239671aded5` thì chúng ta lấy 20 items tiếp theo như sau:

```
const flash_cards = await this.flash_cards_repository
    .find({ _id: { $gt: ObjectId('64ab87fab9e86239671aded5') } })
    .limit(20)
// Hoặc
const flash_cards = await this.flash_cards_repository
    .find({ _id: { $gt: ObjectId('64ab87fab9e86239671aded5') } })
    'projection', { limit: 20 })
```

### Ưu điểm ✅

- **Hiệu tăng cao**: do không còn dùng `skip` nên sẽ không cần scan quá nhiều document.

### Nhược điểm 📛

- **Triển khai phức tạp**: nếu chúng ta sort bằng property khác `_id` hoặc nhiều hơn 2 property thì sẽ khó khăn trong việc triển khai do cần phải tiến hành sort, đến phần ví dụ ở dưới chúng ta sẽ nói rõ hơn.
- **Không thể truy cập ngẫu nhiên vào một trang bất kỳ**: vì page tiếp theo được lấy dữ liệu từ item của page trước đó nên nếu chúng ta muốn vào trang 30 thì phải cần trang 29, tương tự trang 29 cần 28,...