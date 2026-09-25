### JWT Là Gì?

JWT (JSON Web Token) là một tiêu chuẩn mở (RFC 7519) định nghĩa cách truyền thông tin giữa các bên dưới dạng JSON object được ký số (digitally signed). JWT được sử dụng phổ biến để xác thực (authentication) và phân quyền (authorization) trong các hệ thống web hiện đại.

![](https://substackcdn.com/image/fetch/$s_!Ht0r!,w_848,c_limit,f_webp,q_auto:good,fl_progressive:steep/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2F7316c124-1739-45ef-897a-5a11657b914f_1818x1274.png)

Hiểu một cách đơn giản, JWT là một chuỗi ký tự được mã hóa, chứa thông tin xác thực và dữ liệu người dùng, cho phép server nhận diện và phân quyền người dùng mà không cần lưu trạng thái phiên (session) trên server.

JWT giống như một "hộ chiếu kỹ thuật số": khi đăng nhập thành công, server cấp cho client một token chứa thông tin xác thực. Client dùng token này để chứng minh danh tính trong mỗi request tiếp theo, server chỉ cần kiểm tra tính hợp lệ của token mà không cần tra cứu thêm.

Hai đặc điểm nổi bật của JWT:

- **Compact (Nhỏ gọn):** Token có kích thước nhỏ, dễ dàng truyền qua URL, POST body, hoặc HTTP header.
- **Self-contained (Tự chứa):** Token mang theo đầy đủ thông tin cần thiết về người dùng (claims), trong nhiều trường hợp server không cần truy vấn thêm cơ sở dữ liệu để xác minh.

---

### Cấu Trúc JWT

Mỗi JWT bao gồm ba phần, phân tách nhau bằng dấu chấm (`.`):

```
xxxxx.yyyyy.zzzzz
```

Trong đó `xxxxx` là Header, `yyyyy` là Payload, và `zzzzz` là Signature.

![](https://supertokens.com/static/b0172cabbcd583dd4ed222bdb83fc51a/9af93/what-is-jwt.png)

#### Header

Header chứa hai thông tin:

- `typ`: loại token, thường là `"JWT"`.
- `alg`: thuật toán ký được sử dụng, ví dụ `HS256` (HMAC SHA-256) hoặc `RS256` (RSA SHA-256).

Ví dụ Header trước khi mã hóa:

```json
{
  "alg": "HS256",
  "typ": "JWT"
}
```

Header được mã hóa bằng **Base64Url** để tạo thành phần đầu tiên của JWT.

> Lưu ý: Base64Url **không phải là mã hóa bảo mật** — bất kỳ ai cũng có thể giải mã và đọc nội dung Header và Payload. Tính bảo mật của JWT đến từ Signature, không phải từ việc mã hóa nội dung.

#### Payload

Payload chứa các **claims** — tập hợp các thông tin về người dùng hoặc dữ liệu liên quan. Có ba loại claims:

- **Registered claims (Claims chuẩn):** được định nghĩa sẵn trong RFC 7519, không bắt buộc nhưng được khuyến nghị sử dụng. Ví dụ:
    
    - `iss` (issuer): bên phát hành token.
    - `sub` (subject): định danh người dùng.
    - `exp` (expiration time): thời điểm token hết hạn (Unix timestamp).
    - `iat` (issued at): thời điểm token được tạo.
    - `nbf` (not before): token chưa hợp lệ trước thời điểm này.
- **Public claims:** do người dùng tự định nghĩa, nên đăng ký tại IANA JSON Web Token Registry để tránh xung đột.
    
- **Private claims:** thông tin tùy chỉnh do hai bên (client và server) thỏa thuận sử dụng. Ví dụ: `"role": "admin"`, `"user_id": 123`.
    

Ví dụ Payload trước khi mã hóa:

```json
{
  "sub": "1234567890",
  "name": "Nguyen Van A",
  "role": "admin",
  "iat": 1516239022,
  "exp": 1516242622
}
```

Payload cũng được mã hóa bằng **Base64Url** để tạo thành phần thứ hai của JWT.

#### Signature

Signature được tạo bằng cách ký kết hợp Header đã mã hóa và Payload đã mã hóa với một **khóa bí mật (secret key)**, sử dụng thuật toán đã khai báo trong Header.

Ví dụ với thuật toán HS256:

```
HMACSHA256(
  base64UrlEncode(header) + "." + base64UrlEncode(payload),
  secret_key
)
```

Signature có vai trò then chốt: nếu Header hoặc Payload bị thay đổi dù chỉ một ký tự, Signature sẽ không còn khớp và server sẽ từ chối token. Nhờ đó, server có thể tin tưởng rằng thông tin trong token là nguyên vẹn và đến từ đúng nguồn phát hành.

---

### JWT Hoạt Động Như Thế Nào?

#### Quy Trình Tạo JWT (Đăng Nhập)

1. Client gửi thông tin đăng nhập (username/email và password) đến server.
2. Server xác thực thông tin với cơ sở dữ liệu.
3. Nếu hợp lệ, server tạo JWT bằng cách:
    - Tạo Header với thông tin thuật toán.
    - Tạo Payload với thông tin người dùng, quyền hạn, thời gian hết hạn.
    - Ký Header và Payload bằng khóa bí mật để tạo Signature.
4. Server trả JWT về cho client.
5. Client lưu token để sử dụng cho các request tiếp theo.

> Lưu ý bảo mật: Không nên lưu JWT trong `localStorage` nếu ứng dụng có nguy cơ bị tấn công XSS (Cross-Site Scripting). Cân nhắc dùng `HttpOnly cookie` để bảo vệ token tốt hơn.

#### Quy Trình Xác Thực JWT (Mỗi Request Tiếp Theo)

1. Client đính kèm JWT vào mỗi request, thường trong HTTP header:
    
    ```
    Authorization: Bearer <token>
    ```
    
2. Server trích xuất JWT từ header.
3. Server xác thực token bằng cách:
    - Tái tạo Signature từ Header và Payload nhận được, so sánh với Signature trong token — nếu không khớp, token đã bị giả mạo hoặc sửa đổi.
    - Kiểm tra claim `exp` để đảm bảo token chưa hết hạn.
    - Kiểm tra các claims khác nếu cần (ví dụ: `iss`, `aud`).
4. Nếu token hợp lệ, server đọc thông tin từ Payload để xác định quyền truy cập và xử lý request.
5. Server trả về response tương ứng.

Vì server không lưu trạng thái session, JWT hoạt động theo mô hình **stateless** — mỗi request độc lập và tự chứa đủ thông tin xác thực. Điều này giúp hệ thống dễ mở rộng theo chiều ngang (horizontal scaling).

---

### JWT Dùng Để Làm Gì?

#### Xác Thực Người Dùng (Authentication)

Đây là ứng dụng phổ biến nhất. Với cách tiếp cận session truyền thống, server phải lưu thông tin session của từng người dùng (trong bộ nhớ hoặc database), tạo ra áp lực lớn khi hệ thống cần scale.

JWT giải quyết vấn đề này bằng cách chuyển trách nhiệm lưu trữ sang client. Server chỉ cần xác minh tính hợp lệ của token dựa trên khóa bí mật, không cần truy vấn thêm, giúp tăng hiệu suất và khả năng mở rộng.

#### Phân Quyền Truy Cập (Authorization)

Payload của JWT có thể chứa thông tin về vai trò và quyền hạn của người dùng. Server kiểm tra các claims này để quyết định cho phép hay từ chối truy cập vào từng tài nguyên cụ thể mà không cần tra cứu database.

Ví dụ: token chứa `"role": "admin"` cho phép truy cập vào các endpoint quản trị; token chứa `"role": "user"` chỉ truy cập được các endpoint thông thường.

#### Các Ứng Dụng Phổ Biến Khác

- **RESTful API:** JWT là phương pháp xác thực tiêu chuẩn cho API RESTful, bảo vệ endpoint và kiểm soát quyền truy cập.
- **Single Sign-On (SSO):** Người dùng đăng nhập một lần và có thể truy cập nhiều ứng dụng khác nhau mà không cần đăng nhập lại — JWT đóng vai trò là "vé thông hành" chung giữa các hệ thống.
- **Microservices:** Trong kiến trúc microservices, JWT giúp truyền thông tin xác thực giữa các service một cách an toàn và không cần mỗi service phải kết nối tập trung vào một auth server.
- **Mobile Applications:** JWT phù hợp với ứng dụng di động vì không yêu cầu duy trì kết nối liên tục và hoạt động tốt trong môi trường mạng không ổn định.
- **Serverless Functions:** Trong môi trường serverless, JWT cung cấp xác thực stateless, phù hợp với đặc điểm vòng đời ngắn của các function.

---

### Lưu Ý Quan Trọng Khi Sử Dụng JWT

**Không lưu thông tin nhạy cảm trong Payload:** Payload chỉ được mã hóa Base64Url, không được mã hóa bảo mật. Bất kỳ ai có token đều có thể đọc nội dung Payload. Tuyệt đối không lưu mật khẩu, thông tin thẻ tín dụng, hoặc dữ liệu nhạy cảm trong Payload.

**Luôn đặt thời gian hết hạn (exp):** Token không có thời gian hết hạn sẽ có hiệu lực vĩnh viễn — rất nguy hiểm nếu bị lộ. Thực tế thường dùng access token ngắn hạn (15 phút đến 1 giờ) kết hợp với refresh token dài hạn hơn.

**Bảo vệ khóa bí mật:** Khóa bí mật dùng để ký token phải được bảo mật tuyệt đối. Nếu lộ khóa, kẻ tấn công có thể tự tạo token hợp lệ với bất kỳ thông tin nào.

**JWT không thể thu hồi ngay lập tức:** Do server không lưu trạng thái, khi cần vô hiệu hóa một token trước khi hết hạn (ví dụ: người dùng đăng xuất hoặc tài khoản bị khóa), cần có cơ chế bổ sung như **token blacklist** hoặc đặt thời gian hết hạn ngắn và dùng refresh token.

**Chọn thuật toán phù hợp:**

- `HS256` (HMAC): dùng chung một khóa bí mật cho cả ký và xác minh — phù hợp khi chỉ có một bên cần xác minh token.
- `RS256` (RSA): dùng cặp khóa công khai/riêng tư — phù hợp với hệ thống phân tán hoặc SSO, nơi nhiều bên cần xác minh token mà không cần biết khóa riêng tư.

---

### Ghi Chú Chỉnh Sửa

|Hạng mục|Nội dung thay đổi|Lý do|
|---|---|---|
|Định nghĩa JWT|Thêm cụm "được ký số (digitally signed)" vào định nghĩa|Bản gốc mô tả JWT là "chuỗi được mã hóa" — không chính xác; JWT được ký chứ không mã hóa nội dung|
|Phần Header|Thêm ví dụ JSON và lưu ý Base64Url không phải mã hóa bảo mật|Bản gốc thiếu ví dụ minh họa; nhiều người nhầm Base64 với mã hóa bảo mật|
|Phần Payload|Bổ sung ba loại claims (registered, public, private) và bảng claims chuẩn (iss, sub, exp, iat, nbf)|Đây là kiến thức cơ bản quan trọng của JWT bị bỏ sót hoàn toàn|
|Phần Payload|Thêm ví dụ JSON minh họa|Giúp sinh viên hình dung rõ cấu trúc thực tế|
|Phần Signature|Thêm công thức tính Signature với HS256|Bản gốc mô tả mơ hồ, thiếu công thức cụ thể|
|Quy trình xác thực|Bổ sung cơ chế xác minh Signature (tái tạo và so sánh)|Bản gốc chỉ nói "xác thực signature" mà không giải thích cách server làm điều đó|
|Quy trình xác thực|Thêm định dạng header `Authorization: Bearer <token>`|Đây là quy ước thực tế quan trọng bị bỏ sót|
|Lưu ý bảo mật|Thêm cảnh báo về localStorage và XSS|Bản gốc khuyến nghị localStorage mà không đề cập rủi ro bảo mật|
|Mục "Lưu ý quan trọng"|Thêm mục mới hoàn toàn|Bản gốc không có mục này; đây là phần thiết yếu khi dạy JWT cho sinh viên|
|Lưu ý không thể thu hồi|Giải thích vấn đề token blacklist|Đây là hạn chế quan trọng nhất của JWT thường bị bỏ qua|
|Lưu ý chọn thuật toán|Phân biệt HS256 và RS256|Sinh viên cần biết khi nào dùng loại nào|
|Chỉnh sửa nhỏ|Sửa lỗi đánh số danh sách (bản gốc nhảy từ 3 sang 5 ở cả hai quy trình)|Lỗi định dạng|