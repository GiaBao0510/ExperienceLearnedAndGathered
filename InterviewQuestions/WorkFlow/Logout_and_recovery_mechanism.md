
# Cơ chế Đăng xuất (Logout) & Thu hồi Token

## Mục lục

1.  [Tổng quan](https://claude.ai/chat/d6f7164f-e12c-439e-8152-336e1aa0d9c5#1-t%E1%BB%95ng-quan)
2.  [Các khái niệm cốt lõi](https://claude.ai/chat/d6f7164f-e12c-439e-8152-336e1aa0d9c5#2-c%C3%A1c-kh%C3%A1i-ni%E1%BB%87m-c%E1%BB%91t-l%C3%B5i)
3.  [Vì sao dùng cả Whitelist lẫn Blacklist?](https://claude.ai/chat/d6f7164f-e12c-439e-8152-336e1aa0d9c5#3-v%C3%AC-sao-d%C3%B9ng-c%E1%BA%A3-whitelist-l%E1%BA%ABn-blacklist)
4.  [Ba kịch bản quản lý token](https://claude.ai/chat/d6f7164f-e12c-439e-8152-336e1aa0d9c5#4-ba-k%E1%BB%8Bch-b%E1%BA%A3n-qu%E1%BA%A3n-l%C3%BD-token)
5.  [Access Token vs Refresh Token ](https://claude.ai/chat/d6f7164f-e12c-439e-8152-336e1aa0d9c5#5-access-token-vs-refresh-token--kh%C3%A1i-ni%E1%BB%87m-c%C3%B2n-thi%E1%BA%BFu)
6.  [Luồng xử lý đăng xuất](https://claude.ai/chat/d6f7164f-e12c-439e-8152-336e1aa0d9c5#6-lu%E1%BB%93ng-x%E1%BB%AD-l%C3%BD-%C4%91%C4%83ng-xu%E1%BA%A5t)
7.  [Luồng thu hồi token (Case 3)](https://claude.ai/chat/d6f7164f-e12c-439e-8152-336e1aa0d9c5#7-lu%E1%BB%93ng-thu-h%E1%BB%93i-token-case-3)
8.  [Luồng xác thực token ở Middleware](https://claude.ai/chat/d6f7164f-e12c-439e-8152-336e1aa0d9c5#8-lu%E1%BB%93ng-x%C3%A1c-th%E1%BB%B1c-token-%E1%BB%9F-middleware)
9.  [Ví dụ code Go minh họa](https://claude.ai/chat/d6f7164f-e12c-439e-8152-336e1aa0d9c5#9-v%C3%AD-d%E1%BB%A5-code-go-minh-h%E1%BB%8Da)
10.  [Tổng kết](https://claude.ai/chat/d6f7164f-e12c-439e-8152-336e1aa0d9c5#10-t%E1%BB%95ng-k%E1%BA%BFt)
11.  [Mở rộng](https://claude.ai/chat/d6f7164f-e12c-439e-8152-336e1aa0d9c5#m%E1%BB%9F-r%E1%BB%99ng)

----------

## 1. Tổng quan

Tài liệu này mô tả cơ chế xử lý logout ở phía Backend, sử dụng kết hợp **Whitelist** và **Blacklist** lưu trên Redis để kiểm soát trạng thái token một cách an toàn và hiệu quả.

**Điều kiện tiên quyết:**

-   Hệ thống sử dụng JWT (JSON Web Token) làm cơ chế xác thực.
-   Mỗi JWT được cấp phát có một trường `jti` (JWT ID) — định danh duy nhất cho từng token.
-   Redis được dùng làm storage cho Whitelist và Blacklist, tận dụng cơ chế TTL (Time To Live) tích hợp sẵn.

Mục tiêu của cơ chế này là:

- Cho phép người dùng đăng xuất khỏi một thiết bị.
- Cho phép thu hồi token theo từng phiên đăng nhập.
- Cho phép thu hồi toàn bộ tài khoản khi cần.
- Hạn chế việc token bị dùng lại sau khi đã logout.
- Giữ thời gian kiểm tra token nhanh, vì Redis thường được truy cập rất nhanh.

----------

## 2. Các khái niệm cốt lõi

#### **`JWT`**
JWT là chuỗi  token chứa  xác thực thường dùng trong API. gồm ba phần:
`header.payload.signature`

Ví dụ trong HTTP header:
`Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...`

#### **`jti`** (JWT ID)
jti là viết tắt của JWT ID. Mã định danh **duy nhất cho từng token** được cấp phát — nằm trong payload của JWT.

Ví dụ payload JWT:

```json
{
  "sub": "user_123",
  "jti": "0b8f3a1c-6a2e-4f6b-9d1a-7c0f0d9b2a11",
  "iat": 1730000000,
  "exp": 1730000900
}
```
Nếu JWT không có jti, chúng ta sẽ khó quản lý logout/thu hồi theo từng token.

#### **`user_id`**

Mã định danh **người dùng**, không đổi qua các lần đăng nhập. Giá trị thường được định dạng dưới dạng UUID, nhưng bản chất khái niệm là "ID người dùng" chứ không phải "UUID" nói chung — tài liệu này gọi thống nhất là `user_id` để không gây nhầm với khái niệm định danh token (`jti`).

#### **`Redis`**
Redis được dùng để lưu trạng thái token. Các lệnh thường dùng:

- SET
- GET
- DEL
- EXISTS
- EXPIRE
- TTL

#### **Whitelist**
Danh sách các token **đang hợp lệ, được phép sử dụng**. Một token chỉ được middleware chấp nhận nếu nó có mặt trong Whitelist (tùy chọn, tăng bảo mật — xem mục 3).

Nguyên tắc:
Nếu token không có trong whitelist thì token bị từ chối.

#### **Blacklist**
Danh sách các token/người dùng **đã bị vô hiệu hóa**, không còn được phép sử dụng dù chữ ký JWT vẫn hợp lệ và chưa hết hạn tự nhiên.

Nguyên tắc:
Nếu token có trong blacklist thì token bị từ chối.

#### **`Middleware`**
Middleware là tầng xử lý nằm giữa request và handler.

Ví dụ: `Request → Auth Middleware → Handler`

Auth middleware thường làm các việc:

- Đọc token từ header.
- Xác thực chữ ký JWT.
- Kiểm tra token có bị thu hồi không.
- Gắn thông tin user vào context.
- Từ chối request nếu token không hợp lệ.

#### **TTL** (Time To Live)
TTL là Time To Live,  tức thời gian sống của một key trong Redis.
Thời gian một entry tồn tại trong Redis trước khi tự động bị xóa — Redis hỗ trợ sẵn cơ chế này qua lệnh `EXPIRE`, giúp ta không cần tự viết job dọn dẹp định kỳ.

Ví dụ: `SET blacklist:token:token_abc_123 1 EX 900`


Mỗi người dùng ứng với một `user_id` duy nhất. Một người dùng có thể sở hữu **nhiều `jti`** cùng lúc — mỗi lần đăng nhập trên một thiết bị sẽ sinh ra một token mới, tức một `jti` mới. Đây chính là điểm mấu chốt cần nắm trước khi đọc tiếp mục 4: hệ thống cần phân biệt rõ "thu hồi 1 token cụ thể" (theo `jti`) và "thu hồi toàn bộ token của 1 người dùng" (theo `user_id`).

----------

## 3. Vì sao dùng cả Whitelist lẫn Blacklist?

Đa số hệ thống chỉ cần **một trong hai** cơ chế là đủ (phổ biến nhất là chỉ dùng Blacklist, vì đơn giản hơn). Việc kết hợp cả hai như tài liệu gốc đề xuất mang lại lợi ích **fail-safe** (an toàn khi có lỗi), đổi lại chi phí vận hành phức tạp hơn:

-   **Blacklist trả lời câu hỏi:** "Token này đã bị thu hồi chưa?" — mặc định mọi token còn hạn đều hợp lệ, trừ khi có trong Blacklist.
-   **Whitelist trả lời câu hỏi ngược lại:** "Token này có thực sự được hệ thống cấp phát và còn hiệu lực không?" — mặc định mọi token đều **không** hợp lệ, trừ khi có trong Whitelist.

Kết hợp cả hai giúp hệ thống của bạn chống chịu tốt hơn với các tình huống bất thường — ví dụ nếu entry Blacklist vô tình bị xóa nhầm (do lỗi thao tác hoặc sự cố Redis), Whitelist vẫn có thể là lớp chặn thứ hai nếu entry whitelist tương ứng cũng đã bị xóa khi logout. Đây là đánh đổi giữa **độ an toàn** và **độ phức tạp/chi phí vận hành** (nhiều lệnh Redis hơn cho mỗi lần đăng nhập/đăng xuất) — cần cân nhắc theo mức độ nhạy cảm của hệ thống bạn đang xây dựng.

----------

## 4. Thiết kế khóa Redis theo từng trường hợp

### 4.1. Case 1 — Đăng xuất chỉ trên 1 thiết bị - Mỗi thiết bị có phiên độc lập

Yêu cầu:

- Người dùng có thể đăng nhập trên nhiều thiết bị.
- Đăng xuất trên thiết bị nào chỉ thu hồi token của thiết bị đó.
- Các thiết bị khác vẫn hoạt động.

Khi đăng nhập trên thiết bị A, thao tác đăng xuất trên thiết bị A **chỉ** vô hiệu hóa đúng token của thiết bị A, không ảnh hưởng đến các thiết bị khác đang đăng nhập.

**Cấu trúc lưu trên Redis:**

-   Khi đăng nhập: `SET whitelist:{jti} {user_id} EX <thời gian sống của token>`
-   Khi đăng xuất: `SET blacklist:{jti} 1 EX <thời gian còn lại của token>`, đồng thời `DEL whitelist:{jti}`

Ví dụ: người dùng đăng nhập trên điện thoại và laptop sẽ có 2 key riêng biệt: `whitelist:jti_phone_01` và `whitelist:jti_laptop_01`. Đăng xuất trên điện thoại chỉ tạo `blacklist:jti_phone_01` — token trên laptop không bị ảnh hưởng.

**Middleware kiểm tra**
1. Xác thực JWT
2. Kiểm tra blacklist:token:{jti}
3. Nếu dùng whitelist, kiểm tra whitelist:token:{jti}

### 4.2. Case 2 — Đăng nhập thiết bị mới cùng loại sẽ đăng xuất thiết bị cũ cùng loại - Chỉ cho phép một phiên trên mỗi loại thiết bị

Yêu cầu:

- Người dùng đăng nhập trên điện thoại mới.
- Điện thoại cũ cùng loại bị đăng xuất.
- Máy tính hoặc web có thể là loại thiết bị khác, tùy nghiệp vụ

Kịch bản: người dùng đăng nhập vào một điện thoại mới → điện thoại cũ (nếu đang đăng nhập) sẽ tự động bị đăng xuất, nhưng phiên làm việc trên web/desktop thì **không** bị ảnh hưởng.

> thêm `device_type` vào key, và lưu **giá trị là chính `jti` hiện hành** thay vì chỉ lưu `"1"`, để middleware có thể so sánh:

```
Khi đăng nhập:
SET whitelist:{user_id}:{device_type} {jti}  EX <thời gian sống token>

Khi middleware xác thực 1 request:
GET whitelist:{user_id}:{device_type}
Nếu giá trị trả về KHÁC jti trong token hiện tại → từ chối (401)
 (nghĩa là đã có 1 lần đăng nhập MỚI hơn trên cùng loại thiết bị, ghi đè giá trị cũ)
```

Ví dụ: `whitelist:u01:mobile` và `whitelist:u01:web` là hai key tách biệt — đăng nhập điện thoại mới chỉ ghi đè `whitelist:u01:mobile`, không đụng đến `whitelist:u01:web`.

### 4.3. Case 3 — Khóa/thu hồi toàn bộ tài khoản

Yêu cầu:

- Khi khóa tài khoản hoặc người dùng đổi mật khẩu, toàn bộ token của user đó phải bị thu hồi.
- Không quan tâm thiết bị nào.
- Không quan tâm token nào.

Khi quản trị viên khóa tài khoản, hoặc hệ thống phát hiện dấu hiệu bị xâm nhập, **mọi** token của người dùng đó (trên mọi thiết bị) phải lập tức mất hiệu lực.

**Cấu trúc lưu trên Redis:**

```
SET blacklist:{user_id} 1 EX <thời gian còn lại của token có hạn dài nhất>
```

Ví dụ:
```
blacklist:user:user_123 = 1
```

Khi middleware kiểm tra:
```
EXISTS blacklist:user:user_123
```

Nếu tồn tại, từ chối toàn bộ request của user này.

----------

## 5. Access Token vs Refresh Token — khái niệm còn thiếu

Loại token

Vòng đời điển hình

Whitelist/Blacklist có cần áp dụng không?

**Access Token**

Ngắn (5–15 phút)

Thường **không** cần check Redis mỗi request — chấp nhận độ trễ thu hồi tối đa bằng thời gian sống của access token, đổi lại giữ được tính "stateless" (không cần tra Redis) vốn là lợi thế cốt lõi của JWT, giúp giảm tải cho Redis ở quy mô lớn.

**Refresh Token**

Dài (vài ngày đến vài tuần)

**Có** — đây chính là loại token nên áp dụng cơ chế Whitelist/Blacklist trong tài liệu này, vì nó tồn tại đủ lâu để việc thu hồi sớm thực sự có ý nghĩa.

**Vì sao điều này quan trọng?** Nếu áp dụng whitelist/blacklist cho **cả access token**, mỗi request API đều phải tra Redis — điều này làm mất đi lợi ích "không cần tra database mỗi request" vốn là lý do chính người ta chọn JWT thay vì session truyền thống. Cách tiếp cận phổ biến trong thực tế: chấp nhận access token có thể "sống sót" tối đa vài phút sau khi bị thu hồi (rủi ro chấp nhận được với hầu hết hệ thống), còn refresh token — thứ quyết định việc có cấp access token mới hay không — mới thực sự cần kiểm tra whitelist/blacklist qua Redis ở bước làm mới token (`/refresh`).

----------

## 6. Luồng xử lý đăng xuất

Sơ đồ dưới đây áp dụng cho **Case 1** (đăng xuất 1 thiết bị)  khi người dùng chỉ đăng xuất 1 thiết bị:

```
Client gửi request logout
    │
    ▼
Backend nhận token từ Authorization header
    │
    ▼
Backend xác thực JWT
    │
    ├─ Token sai chữ ký / sai định dạng → 401
    │
    ▼
Trích xuất jti, user_id, exp từ token
    │
    ▼
Tính TTL còn lại = exp - thời gian hiện tại
    │
    ├─ Token đã hết hạn → có thể trả về 200/204 nếu muốn idempotent
    │
    ▼
Nếu TTL còn lại > 0:
    SET blacklist:token:{jti} 1 EX <ttl_còn_lại>
    │
    ▼
DEL whitelist:token:{jti}
    │
    ▼
Trả về 200 hoặc 204 cho client
```

----------
## 7. Luồng thu hồi token (Case 3)

```
Admin (hoặc hệ thống phát hiện xâm nhập) yêu cầu khóa tài khoản {user_id}
    │
    ▼
Backend xác định thời gian còn lại dài nhất trong số các token
đang hoạt động của user_id (hoặc dùng TTL mặc định an toàn nếu không biết chính xác) trong White list
    │
    ▼
Ghi blacklist:{user_id} = 1, TTL = thời gian còn lại đó
    │
    ▼
(Tùy chọn) Xóa toàn bộ entry whitelist:{user_id}:* liên quan,
nếu hệ thống đang dùng whitelist theo device_type như mục 4.2
    │
    ▼
Từ thời điểm này, MỌI request kèm token của user_id sẽ bị middleware
từ chối ở bước kiểm tra blacklist:{user_id} (xem mục 8) — bất kể jti nào

```

----------

## 8. Luồng xác thực token ở Middleware

Mỗi request API cần được kiểm tra theo thứ tự sau:

```
1. Xác thực chữ ký JWT và thời hạn
2. Trích xuất jti, user_id, device_type nếu có
3. Kiểm tra blacklist:token:{jti}
4. Kiểm tra blacklist:user:{user_id}
5. Kiểm tra whitelist hoặc session hiện tại nếu có
6. Cho phép request đi tiếp
```

Mô tả cụ thể:

```
1. Giải mã và xác thực chữ ký JWT
       │
       ▼
2. Kiểm tra token có bị thu hồi không?
   EXISTS blacklist:{jti}          → Có → Từ chối (401), token cụ thể đã bị logout
   EXISTS blacklist:{user_id}      → Có → Từ chối (401), toàn bộ tài khoản đã bị khóa/thu hồi
       │
       ▼
3. (Tùy chọn, tăng bảo mật) Kiểm tra token có trong Whitelist không?
   EXISTS whitelist:{jti}   (nếu dùng chiến lược whitelist theo Case 1)
   hoặc GET whitelist:{user_id}:{device_type} và so sánh với jti  (nếu dùng chiến lược theo Case 2)
       │ Không khớp → Từ chối (401)
       │
       ▼
4. Cho phép request tiếp tục xử lý
```

----------

## 9. Ví dụ code Go minh họa

Ví dụ dưới đây minh họa middleware xác thực token theo đúng luồng đã mô tả ở mục 8, dùng thư viện `go-redis`:

```go
func AuthMiddleware(rdb *redis.Client, jwtSecret []byte) gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx := c.Request.Context()
        tokenStr := extractTokenFromHeader(c)

        // 1. Giải mã và xác thực chữ ký JWT
        claims, err := parseAndValidateJWT(tokenStr, jwtSecret)
        if err != nil {
            c.AbortWithStatusJSON(401, gin.H{"error": "invalid or expired token"})
            return
        }

        // 2. Kiểm tra blacklist theo jti VÀ theo user_id
        blacklistedByJTI, _ := rdb.Exists(ctx, "blacklist:"+claims.JTI).Result()
        blacklistedByUser, _ := rdb.Exists(ctx, "blacklist:"+claims.UserID).Result()
        if blacklistedByJTI > 0 || blacklistedByUser > 0 {
            c.AbortWithStatusJSON(401, gin.H{"error": "token revoked"})
            return
        }

        // 3. (Tùy chọn) Kiểm tra whitelist theo jti
        exists, _ := rdb.Exists(ctx, "whitelist:"+claims.JTI).Result()
        if exists == 0 {
            c.AbortWithStatusJSON(401, gin.H{"error": "token not recognized"})
            return
        }

        c.Set("user_id", claims.UserID)
        c.Next()
    }
}

```

```go
func LogoutHandler(rdb *redis.Client) gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx := c.Request.Context()
        claims := c.MustGet("claims").(*JWTClaims)

        remainingTTL := time.Until(claims.ExpiresAt.Time)

        // Chỉ blacklist đúng token hiện tại (Case 1) — KHÔNG blacklist user_id ở đây
        rdb.Set(ctx, "blacklist:"+claims.JTI, 1, remainingTTL)
        rdb.Del(ctx, "whitelist:"+claims.JTI)

        c.JSON(200, gin.H{"message": "logged out"})
    }
}

```

----------

## 10. Tổng kết

-   Cơ chế kết hợp Whitelist + Blacklist mang lại tính fail-safe cao hơn, đổi lại phức tạp và tốn chi phí vận hành Redis hơn — cần cân nhắc theo mức độ nhạy cảm của hệ thống.
-   Case 2 (giới hạn 1 phiên/loại thiết bị) trong thiết kế gốc có lỗ hổng logic — cần key theo `{user_id}:{device_type}` và lưu giá trị là `jti` để middleware so sánh được, thay vì chỉ lưu theo `user_id`.
-   Luồng logout thông thường (Case 1) **không nên** blacklist theo `user_id` — chỉ Case 3 (khóa toàn tài khoản) mới cần việc đó. Đây là lỗi logic quan trọng nhất đã sửa trong tài liệu gốc.
-   Middleware cần dùng **đúng 1 quy ước đặt tên key** nhất quán giữa bước ghi (login/logout) và bước đọc (mỗi request) — sai lệch tên key khiến toàn bộ cơ chế thu hồi không hoạt động trong thực tế dù thiết kế trên giấy có vẻ đúng.
-   Trong thực tế, cơ chế whitelist/blacklist nên áp dụng chủ yếu cho **Refresh Token** (vòng đời dài) thay vì Access Token (vòng đời ngắn), để giữ được lợi thế "stateless" của JWT ở phần lớn các request.

----------

### Mở rộng

-   **Redis Cluster & tính sẵn sàng cao (High Availability):** vì cơ chế xác thực giờ phụ thuộc vào Redis, Redis trở thành một điểm có thể gây lỗi toàn hệ thống (single point of failure) nếu không được thiết kế HA — đáng tìm hiểu Redis Sentinel/Cluster để giảm rủi ro này.
-   **Token rotation cho Refresh Token:** kỹ thuật cấp một refresh token mới mỗi khi client dùng refresh token để lấy access token mới, đồng thời blacklist refresh token cũ ngay lập tức — giúp phát hiện sớm nếu refresh token bị đánh cắp và tái sử dụng trái phép (reuse detection).
-   **So sánh JWT stateless thuần túy vs JWT có Redis check:** đây chính là đánh đổi cốt lõi được nhắc đến ở mục 5 — đáng đọc thêm về mô hình "hybrid" này và khi nào nên quay lại dùng session truyền thống (lưu toàn bộ state ở server) thay vì cố "vá" JWT bằng Redis.
-   **Liên hệ với tài liệu Bảo mật dữ liệu đã học:** toàn bộ cơ chế whitelist/blacklist trong tài liệu này chính là một ví dụ cụ thể của "Technical controls" (đã học ở Data Security) áp dụng cho mục tiêu **Confidentiality** — ngăn chặn việc dùng lại token của người dùng đã đăng xuất hoặc bị xâm nhập.
-   **Rate limiting cho endpoint `/logout` và `/refresh`:** hai endpoint này là mục tiêu hấp dẫn cho tấn công vét cạn hoặc lạm dụng nếu không giới hạn tần suất gọi — nên tìm hiểu thêm cách áp dụng rate limiting (VD: dựa trên `user_id` hoặc IP) cho riêng hai endpoint này.