N+1 query là kẻ thù của Hiệu năng hệ thống? Nhưng Không...

Tôi tin nhiều bạn dev vẫn tìm cách tránh N+1 query. Thực ra N+1 có phải là kẻ thù của hệ thống, chúng ta cùng phân tích nhé.

![]

##  N + 1 query là gì ?
N+1 Query Problem là vấn đề hiệu năng xảy ra khi application thực hiện một truy vấn để lấy danh sách bản ghi, sau đó tiếp tục thực hiện thêm một truy vấn riêng cho từng bản ghi để lấy dữ liệu liên quan.

Vấn đề thường xuất hiện khi làm việc với quan hệ giữa các bảng, đặc biệt khi backend sử dụng ORM để truy vấn dữ liệu. Tuy nhiên, N+1 không chỉ xuất hiện ở ORM; developer cũng có thể vô tình tạo ra vấn đề này khi viết SQL thủ công và đặt truy vấn bên trong vòng lặp.

Tên gọi N+1 xuất phát từ số lượng truy vấn: một query ban đầu để lấy danh sách, cộng thêm N query để lấy dữ liệu liên quan cho N bản ghi.

Câu chuyện xảy ra khi chúng ta, những coder viết code chưa khéo, sinh ra nhiều truy vấn vào cơ sở dữ liệu làm giảm performance của hệ thống. Khi đồng nghiệp đọc code thường thì sẽ bình luận ngay: fix N + 1

Nếu để trót lọt, sau một thời gian vận hành mà chương trình chậm, điều tra ra nguyên nhân rồi lại câu nói kinh điển:
✌️ Đứa nào code ra cái đống shit này đây ✌️

## Ví dụ về N + 1

Mình viết ví dụ tựa như mã giả, không đặt nặng vấn đề cú pháp:

Giả sử ta có một cơ sở dữ liệu, trong đó table post có khóa ngoại user_id, nói theo kiểu mã giả là một post thuộc về một user

Thực hiện truy vấn vào cơ sở dữ liệu và lấy tất cả User kèm theo các Post của User đó:

```
User.all.each do |user|
  user.posts
end
```

Các câu lệnh SQL sinh ra như sau
```sql
  User Load (0.2ms)  SELECT "users".* FROM "users"
  
  Post Load (0.1ms)  SELECT "posts".* FROM "posts" WHERE "posts"."user_id" = ?  [["user_id", 1]]
  Post Load (0.1ms)  SELECT "posts".* FROM "posts" WHERE "posts"."user_id" = ?  [["user_id", 2]]
  Post Load (0.1ms)  SELECT "posts".* FROM "posts" WHERE "posts"."user_id" = ?  [["user_id", 3]]
```

Lấy máy tính CASIO FX-500 MS ra để đếm thì thấy mình cần dùng 4 câu truy vấn:

- Một truy vấn để lấy ra tất cả users => đây chính là 1 trong "N+1"
- Ba truy vấn để lấy ra các post tương ứng với ba user trong cơ sở dữ liệu => đây chính là N trong "N+1"
Đối với những hệ thống có số lượng bản ghi lớn (cỡ như phải trả về 1000 user thì chúng ta phải thực hiện 1001 truy vấn) hoặc có database với độ trễ cao (thời gian thực thi truy vấn cao) thì ắt hẳn sẽ làm giảm performance của hệ thống.

Quan hệ giữa hai bảng là one-to-many. Backend cần tạo API trả danh sách user kèm các bài viết tương ứng.

Bước đầu tiên, application truy vấn danh sách user:

```
SELECT id, name
FROM users;
```
Giả sử kết quả có 100 user. Tiếp theo, application cần lấy bài viết của từng người:

```
SELECT id, title
FROM posts
WHERE user_id = 1;
```
Sau đó, query tương tự được chạy với user_id = 2, user_id = 3 và tiếp tục cho đến user cuối cùng.
Kết quả là:
- Một query lấy danh sách user.
- Một trăm query lấy bài viết của từng user.

##### **Vì sao N+1 gây chậm?**
Nguyên nhân chính là application thực hiện quá nhiều lần giao tiếp với database cho một yêu cầu logic.

Ngay cả khi mỗi query riêng lẻ chỉ mất vài mili giây, hàng trăm query tuần tự vẫn có thể làm tổng thời gian xử lý tăng đáng kể. Ngoài thời gian database thực thi SQL, application còn phải trả chi phí cho network round trip, xử lý kết quả và quản lý tài nguyên truy vấn.

N+1 cũng làm tăng số lần database phải tiếp nhận và xử lý câu lệnh. Khi nhiều request cùng xảy ra, lượng query phát sinh có thể tạo thêm áp lực lên connection pool và database server.

Vậy làm sao để có thể lấy ra dữ liệu tương đương như vậy nhưng với số lượng truy vấn bé hơn ?

---
### **Bản chất của N+1 Query**
Vấn đề N+1 xảy ra khi chúng ta thực hiện 1 truy vấn ban đầu lấy danh sách (ví dụ: 100 bài viết) và sau đó thực hiện N truy vấn riêng biệt (100 truy vấn) để lấy dữ liệu liên quan cho từng mục 
ví dụ: 
- thông tin tác giả, bình luận, v.v.
- Bạn cần lấy danh sách 100 người dùng và sau đó thực hiện thêm 100 truy vấn riêng biệt để lấy thông tin đơn hàng của họ.
Đây là một vấn đề hiệu năng phổ biến trong các ứng dụng tương tác với cơ sở dữ liệu, đặc biệt khi dùng ORM (Object-Relational Mapping).

Hậu quả: 
- Latency (Độ trễ) tăng vọt.
- Overhead: 101 lần round-trip Network giữa App Server và Database Server.
- Database Load: tăng tải DB
- Scalability: giảm khả năng mở rộng hệ thống trầm trọng.

---
## Cách khắc phục

#### Sử dụng select in ()
Tối ưu câu lệnh SQL ngay và luôn.

```
User Load (0.2ms)  SELECT "users".* FROM "users"
Post Load (0.4ms)  SELECT "posts".* FROM "posts" WHERE "posts"."user_id" IN (1, 2, 3)
```

Chúng ta cần 2 truy vấn:

Truy vấn đầu tiên để load toàn bộ user
Truy vấn thứ 2 để load các post tương ứng với những users đó.


#### Sử dụng joins

- Giải pháp kinh điển: JOIN - Giảm N+1 truy vấn xuống còn 1. 
- Để giải quyết vấn đề trên chúng ta có thể sử dụng JOIN FETCH để lấy toàn bộ data cần thiết trong cùng 1 câu query.
- Ngoài ra có thể sử dụng nativeQuery thay vì sử dụng HQL để lấy dữ liệu


Tuy nhiên, cũng có nhiều hệ lụy
- Cartesian Product: Tạo ra tập kết quả khổng lồ, dư thừa dữ liệu (nếu một bài viết có 100 bình luận, nó sẽ lặp lại dữ liệu bài viết 100 lần), gây tốn bộ nhớ và I/O lớn hơn cho App Server.
- Khó khăn trong việc Caching: Tập kết quả lớn, thay đổi thường xuyên, khó cache toàn bộ.

#### Eager Loading (Preloading)
Eager Loading là cách tải dữ liệu liên quan một cách chủ động ngay khi truy vấn dữ liệu chính, thay vì đợi đến lúc cần từng quan hệ mới truy vấn riêng.
Đây là cách phổ biến nhất để xử lý vấn đề N+1 query giúp lấy tất cả dữ liệu cần thiết chỉ với một vài truy vấn, thay vì phải thực hiện N+1 truy vấn riêng lẻ:
// Lấy users và đồng thời eager loading orders trong 1 lần
var users []User
db.Preload("Orders").Find(&users)
💡 Với cách này, hệ thống sẽ thực hiện 2 truy vấn chính:
- Truy vấn 1: Lấy tất cả users.
- Truy vấn 2: Lấy tất cả orders liên quan đến các users đó.

Giả sử các model được khai báo như sau:

```go
type User struct {
    ID    uint
    Name  string
    Posts []Post `gorm:"foreignKey:UserID"`
}

type Post struct {
    ID     uint
    UserID uint
    Title  string
}
```

Cách dễ phát sinh N+1:
```go
var users []User

if err := db.Find(&users).Error; err != nil {
    return err
}

for i := range users {
    if err := db.
        Where("user_id = ?", users[i].ID).
        Find(&users[i].Posts).Error; err != nil {
        return err
    }
}
```

Ở đây, query đầu tiên lấy danh sách user. Sau đó, mỗi user tạo thêm một query để lấy posts.

Cách cải thiện bằng GORM:
```sql
var users []User

if err := db.Preload("Posts").Find(&users).Error; err != nil {
    return err
}
```

Thay vì truy vấn posts riêng cho từng user, GORM thường thực hiện hai query:
```sql
SELECT * FROM users;

SELECT * FROM posts
WHERE user_id IN (1, 2, 3, 4, ...);
```

Query thứ hai lấy posts của nhóm user và GORM ánh xạ các kết quả về đúng association tương ứng.
Như vậy, nếu cần tải posts cho 100 user, số query cơ bản thường giảm từ 101 xuống còn 2, không tính những query khác mà application có thể thực hiện. Đây là lợi ích chính của eager loading.


#### **Batch Query**
Batch Query là cách lấy dữ liệu liên quan cho nhiều bản ghi trong một query thay vì truy vấn từng bản ghi một.

Đây là ý tưởng cốt lõi mà eager loading thường sử dụng. Tuy nhiên, batch query cũng có thể được triển khai thủ công nếu application dùng SQL thuần hoặc không muốn phụ thuộc vào tính năng ORM.

Thay vì:
```sql
SELECT * FROM posts WHERE user_id = 1;
SELECT * FROM posts WHERE user_id = 2;
SELECT * FROM posts WHERE user_id = 3;
```

application lấy các user ID rồi truy vấn một lần:
```sql
SELECT *
FROM posts
WHERE user_id IN (1, 2, 3);
```
Sau đó, application gom các posts theo user_id trong bộ nhớ và gắn chúng vào từng user.

Cách làm này giảm số query, nhưng cần chú ý các trường hợp không có bài viết, tập ID rỗng và giới hạn số lượng tham số của database driver.

Với danh sách rất lớn, có thể chia ID thành từng batch nhỏ để tránh câu query quá lớn hoặc tiêu thụ quá nhiều bộ nhớ. Kích thước batch nên được chọn theo kích thước dữ liệu, giới hạn driver và hiệu năng thực tế.


#### **Phương án Kết hợp với Item-Level Caching**

Và đây là lúc chúng ta thay đổi tư duy. Khi một hệ thống đã có lớp Caching mạnh mẽ (Redis, Memcached, hoặc thậm chí In-Memory Cache có phân phối), N+1 query lạ là vấn đề nữa, mà trở thành một chiến lược truy cập dữ liệu tốt.

Chiến lược đơn giản như sau: 1 Query (List) + N Query (Cache Lookups)

Hành động:
- Query (1)L Lấy danh sách ID cơ bản (ví dụ: SELECT PostId FROM Posts WHERE...). Query này nhanh, nhẹ, và có thể được Cache toàn bộ Page.
- Query (N) - Truy cập từng Item, thay vì gọi DB, hệ thống gọi Cache cho từng item (ví dụ: Cache.Get<Author>(AuthorId)).
- Hit Cache (hầu hết), độ trễ cực thấp, là Microseconds (thường < 1ms), không có Network Round-trip đến DB.
- Miss Cache (trường hợp lấy lần đầu), chỉ list item đó gọi DB. Sau đó, item này được đưa vào Cache, sẵn sàng cho các lần truy cập tiếp theo.

**Ưu điểm**
- Cache cho N+1 sẽ phù hợp cho bài toán dữ liệu không thay đổi liên tục

**Nhược điểm ở chiến lượt này:**
- Dữ liệu lưu trên Cache phải update liên tục
- Có thể đọc sai thông tin nếu update trễ
- Không phù hợp cho những trường hợp như comment, rating,...

=> Cũng tuỳ cho từng bài toán

---

### Nghệ thuật tối ưu luôn là sự Cân bằng
Trong system design hiện dại, Database luôn là nút thắt cổ chai đắt đỏ nhất. 
Tuy nhiên, không phải 1 chiến lược lúc nào cũng đúng. Tùy từng thời điểm, từng bài toán áp dụng, khi cần JOIN query, chúng ta vẫn dùng, khi cần N+1 query, chúng ta vẫn dùng. Tìm phương án tối ưu cho mỗi bài toán là mục tiêu tối thượng của 1 đứa dev như chúng ta :) 
