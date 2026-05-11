# Logout trên một hoặc nhiều thiết bị

## Tổng quan

Tài liệu này mô tả cơ chế xử lý logout ở phía Backend, sử dụng kết hợp **Whitelist** và **Blacklist** lưu trên Redis để kiểm soát trạng thái token một cách an toàn và hiệu quả.

**Điều kiện tiên quyết:**

- Hệ thống sử dụng JWT (JSON Web Token) làm cơ chế xác thực.
- Mỗi JWT được cấp phát có một trường `jti` (JWT ID) định danh duy nhất cho từng token.
- Redis được sử dụng làm storage cho Whitelist và Blacklist, tận dụng cơ chế TTL (Time To Live) tích hợp sẵn.

---

## Các khái niệm cốt lõi

### JWT ID (`jti`)

Mỗi JWT được cấp phát kèm theo một claim `jti` — một chuỗi định danh duy nhất (thường là UUID). Khi người dùng đăng nhập trên nhiều thiết bị, mỗi phiên đăng nhập sẽ nhận được một JWT riêng biệt với `jti` khác nhau. Đây là cơ sở để phân biệt token giữa các thiết bị.

### Whitelist

Lưu trữ tất cả các token hợp lệ đang hoạt động. Mỗi entry có dạng:

```
Key:   whitelist:{jti}
Value: <user_id hoặc metadata tùy thiết kế>
TTL:   Thời gian sống còn lại của token
```

### Blacklist

Lưu trữ các token đã bị vô hiệu hóa (do logout). Mỗi entry có dạng:

```
Key:   blacklist:{jti}
Value: <user_id hoặc "revoked">
TTL:   Thời gian sống còn lại của token (đảm bảo entry tự xóa sau khi token hết hạn)
```

> **Lý do đặt TTL cho Blacklist bằng thời gian còn lại của token:** Sau khi token hết hạn tự nhiên, nó không còn có thể được sử dụng nữa, nên việc giữ entry trong Blacklist là không cần thiết. Cơ chế TTL giúp Redis tự động dọn dẹp, tránh tích lũy dữ liệu thừa.

---

## Luồng xử lý Logout trên một thiết bị

Đây là trường hợp người dùng chỉ muốn đăng xuất khỏi thiết bị hiện tại, các phiên trên thiết bị khác vẫn còn hiệu lực.

### Sơ đồ luồng

```
Client gửi request logout (kèm token trong header)
    │
    ▼
Backend giải mã và xác thực token
    │
    ├─ Token không hợp lệ / đã hết hạn → Trả về lỗi 401
    │
    ▼
Trích xuất jti từ payload của token
    │
    ▼
Ghi jti vào Blacklist (với TTL = thời gian còn lại của token)
    │
    ▼
Xóa entry whitelist:{jti} khỏi Whitelist
    │
    ▼
Trả về response thành công cho client
```

### Chi tiết từng bước

**Bước 1 — Nhận request logout:**

Client gửi HTTP request đến endpoint logout, kèm JWT trong header:

```
Authorization: Bearer <token>
```

**Bước 2 — Xác thực và giải mã token:**

Backend kiểm tra chữ ký và thời hạn của token. Nếu token không hợp lệ hoặc đã hết hạn, trả về lỗi `401 Unauthorized`.

**Bước 3 — Trích xuất `jti`:**

Sau khi xác thực thành công, lấy giá trị `jti` từ payload của JWT.

**Bước 4 — Cập nhật Redis:**

```
# Tính TTL còn lại của token
remaining_ttl = token.exp - current_time

# Thêm vào Blacklist
SET blacklist:{jti} "revoked" EX {remaining_ttl}

# Xóa khỏi Whitelist
DEL whitelist:{jti}
```

**Bước 5 — Trả về response:**

```json
HTTP 200 OK
{
  "message": "Logout successful"
}
```

---

## Luồng xử lý Logout trên nhiều thiết bị

Đây là trường hợp người dùng muốn đăng xuất khỏi tất cả các thiết bị đang đăng nhập cùng lúc.

### Sự khác biệt so với logout một thiết bị

Thay vì sử dụng `jti` (định danh một token cụ thể) để vô hiệu hóa, hệ thống sử dụng `user_id` (định danh người dùng) làm khóa cho Blacklist. Vì `user_id` là duy nhất cho mỗi tài khoản, một entry duy nhất trong Blacklist sẽ đủ để vô hiệu hóa tất cả các token thuộc về người dùng đó.

### Cập nhật cơ chế kiểm tra token

Khi nhận mỗi request sau khi logout, middleware xác thực phải kiểm tra thêm:

```
# Kiểm tra blacklist theo user_id (cho logout toàn bộ thiết bị)
EXISTS blacklist:user:{user_id}
```

Nếu key này tồn tại, tất cả các token của người dùng đều bị coi là không hợp lệ, bất kể `jti` là gì.

### Chi tiết cập nhật Redis khi logout nhiều thiết bị

```
# Tính TTL phù hợp — thường lấy max TTL có thể của access token
# hoặc một khoảng thời gian đủ dài để bao phủ tất cả các token đang active

SET blacklist:user:{user_id} "revoked" EX {max_token_ttl}

# Xóa toàn bộ Whitelist của user (nếu có lưu theo user_id)
DEL whitelist:user:{user_id}
```

> **Lưu ý về TTL cho Blacklist theo `user_id`:** Không thể xác định chính xác TTL tối ưu vì các token khác nhau có thời điểm cấp phát khác nhau. Một cách tiếp cận thực tế là đặt TTL bằng thời gian sống tối đa của access token trong hệ thống (ví dụ: nếu access token sống 24 giờ, đặt TTL là 24 giờ).

---

## So sánh hai phương pháp

|Tiêu chí|Logout một thiết bị|Logout nhiều thiết bị|
|---|---|---|
|Khóa Blacklist|`blacklist:{jti}`|`blacklist:user:{user_id}`|
|Phạm vi ảnh hưởng|Chỉ token hiện tại|Tất cả token của user|
|TTL|Thời gian còn lại của token|TTL tối đa của access token|
|Độ phức tạp|Thấp|Trung bình|
|Use case|Logout thông thường|Đổi mật khẩu, phát hiện xâm nhập|

---

## Luồng xác thực token (Middleware)

Mỗi request API cần được kiểm tra theo thứ tự sau:

```
1. Giải mã và xác thực chữ ký JWT
       │
       ▼
2. Kiểm tra token có trong Blacklist theo jti không?
   EXISTS blacklist:{jti}
       │ Có → Từ chối (401)
       │
       ▼
3. Kiểm tra user có trong Blacklist không? (cho logout toàn bộ thiết bị)
   EXISTS blacklist:user:{user_id}
       │ Có → Từ chối (401)
       │
       ▼
4. Kiểm tra token có trong Whitelist không? (tùy chọn, tăng bảo mật)
   EXISTS whitelist:{jti}
       │ Không → Từ chối (401)
       │
       ▼
5. Cho phép request tiếp tục xử lý
```

---

## Đề xuất cải thiện thêm

### 1. Xem xét lại sự cần thiết của Whitelist

Trong nhiều hệ thống, Whitelist không phải lúc nào cũng cần thiết và có thể gây overhead không đáng có:

- Nếu mục tiêu chính là vô hiệu hóa token khi logout, **chỉ cần Blacklist là đủ**. Khi một `jti` không có mặt trong Blacklist, token được coi là hợp lệ (kết hợp với việc kiểm tra chữ ký và TTL của JWT).
- Whitelist chỉ thực sự cần thiết khi hệ thống yêu cầu khả năng **chủ động kiểm soát từng phiên** — ví dụ: chỉ cho phép tối đa N thiết bị đăng nhập đồng thời, hoặc hiển thị danh sách "phiên đang hoạt động" cho người dùng quản lý.

### 2. Sử dụng Redis Sets để quản lý phiên theo user

Thay vì chỉ lưu một key duy nhất cho logout nhiều thiết bị, có thể dùng Redis Set để lưu tất cả các `jti` đang hoạt động của một user:

```
Key:     sessions:{user_id}     (Redis Set)
Members: {jti_1}, {jti_2}, ...
```

Cách này cho phép:

- Hiển thị danh sách thiết bị đang đăng nhập.
- Logout thiết bị cụ thể mà không ảnh hưởng thiết bị khác.
- Logout toàn bộ bằng cách xóa toàn bộ Set và blacklist tất cả các `jti` trong đó.

### 3. Cân nhắc dùng Refresh Token thay vì kiểm soát Access Token

Một kiến trúc phổ biến hơn là:

- **Access Token** có TTL ngắn (5–15 phút), không lưu trong Redis.
- **Refresh Token** có TTL dài, được lưu và quản lý trong database hoặc Redis.

Khi logout, chỉ cần xóa Refresh Token. Access Token sẽ tự hết hạn sau khoảng thời gian ngắn. Cách này giảm áp lực lên Redis vì không cần Blacklist cho Access Token.

### 4. Ghi log sự kiện logout

Nên ghi lại các sự kiện logout (thời gian, user, thiết bị, IP) để phục vụ audit và phát hiện bất thường bảo mật.