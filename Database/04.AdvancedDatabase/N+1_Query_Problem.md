##  N + 1 query là gì ?
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
```
  User Load (0.2ms)  SELECT "users".* FROM "users"
  
  Post Load (0.1ms)  SELECT "posts".* FROM "posts" WHERE "posts"."user_id" = ?  [["user_id", 1]]
  Post Load (0.1ms)  SELECT "posts".* FROM "posts" WHERE "posts"."user_id" = ?  [["user_id", 2]]
  Post Load (0.1ms)  SELECT "posts".* FROM "posts" WHERE "posts"."user_id" = ?  [["user_id", 3]]
```

Lấy máy tính CASIO FX-500 MS ra để đếm thì thấy mình cần dùng 4 câu truy vấn:

- Một truy vấn để lấy ra tất cả users => đây chính là 1 trong "N+1"
- Ba truy vấn để lấy ra các post tương ứng với ba user trong cơ sở dữ liệu => đây chính là N trong "N+1"
Đối với những hệ thống có số lượng bản ghi lớn (cỡ như phải trả về 1000 user thì chúng ta phải thực hiện 1001 truy vấn) hoặc có database với độ trễ cao (thời gian thực thi truy vấn cao) thì ắt hẳn sẽ làm giảm performance của hệ thống.

Vậy làm sao để có thể lấy ra dữ liệu tương đương như vậy nhưng với số lượng truy vấn bé hơn ?



```
```

```
```

```
```
