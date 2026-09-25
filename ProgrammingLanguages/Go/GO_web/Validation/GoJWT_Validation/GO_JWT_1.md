### Access Token Là Gì?

Access Token là một chuỗi ký tự được ký số (thường là JWT) đại diện cho quyền truy cập vào tài nguyên cụ thể trên hệ thống (API, dữ liệu, dịch vụ). Token này được cấp sau khi người dùng hoặc ứng dụng hoàn thành quá trình xác thực, và được sử dụng để thay thế mật khẩu trong các request tiếp theo.

Hiểu đơn giản:

- API = cánh cửa kết nối
- API Key = chìa khóa tổng (không giới hạn thời gian, không phân biệt phạm vi quyền)
- Access Token = vé vào cửa có thời hạn sử dụng và phạm vi quyền cụ thể

Đặc điểm quan trọng:

- Có thời gian sống giới hạn (thường từ 5 phút đến 1 giờ)
- Chứa thông tin về phạm vi quyền (scope) — ví dụ: chỉ được đọc dữ liệu, không được xóa
- Có thể bị thu hồi (revoke) thông qua cơ chế blacklist hoặc bằng cách để token hết hạn
- Không chứa mật khẩu người dùng

---

## Access Token Hoạt Động Như Thế Nào?

### Quy Trình Cấp Và Sử Dụng Token

**Bước 1: Người dùng đăng nhập** Client gửi thông tin đăng nhập (username/password hoặc qua OAuth) đến server.

**Bước 2: Server xác thực** Server kiểm tra thông tin đăng nhập với cơ sở dữ liệu.

**Bước 3: Cấp Access Token và Refresh Token** Nếu thông tin hợp lệ, server tạo và trả về hai token:

- **Access Token**: thời gian sống ngắn (ví dụ: 15 phút), dùng để gọi API.
- **Refresh Token**: thời gian sống dài hơn (ví dụ: 7 ngày), dùng để xin cấp Access Token mới khi Access Token hết hạn.

**Bước 4: Client sử dụng Access Token** Mỗi request API đính kèm Access Token trong header:

```
Authorization: Bearer <access_token>
```

Server kiểm tra tính hợp lệ của token, xác minh phạm vi quyền và trả về dữ liệu tương ứng.

**Ví dụ thực tế:**

Khi đăng nhập vào một ứng dụng quản lý công việc bằng tài khoản Google:

- Google xác thực danh tính người dùng.
- Google cấp Access Token cho ứng dụng với phạm vi quyền giới hạn (ví dụ: chỉ đọc email, tên, ảnh đại diện).
- Ứng dụng dùng token đó để gọi Google API — không cần biết mật khẩu Gmail của người dùng.
- Khi Access Token hết hạn, ứng dụng dùng Refresh Token để xin cấp token mới mà không làm gián đoạn trải nghiệm người dùng.

---

### Refresh Token Là Gì?

Refresh Token là một token tồn tại lâu dài, được sử dụng **duy nhất** cho mục đích: xin cấp Access Token mới khi Access Token cũ đã hết hạn, mà không yêu cầu người dùng đăng nhập lại.

Các đặc điểm chính:

- **Thời gian sống:** Dài hơn Access Token nhiều — thường từ vài giờ đến vài tuần, có thể cấu hình tùy hệ thống.
- **Mục đích duy nhất:** Chỉ được gửi đến endpoint `/auth/refresh` để đổi lấy Access Token mới. Không dùng để gọi API nghiệp vụ.
- **Lưu trữ:** Phải được lưu trữ an toàn — thường trong `HttpOnly cookie` (trình duyệt) hoặc secure storage (ứng dụng di động). Tuyệt đối không lưu trong `localStorage`.
- **Không chia sẻ với bên thứ ba:** Refresh Token chỉ giao tiếp với auth server của chính hệ thống, không bao giờ gửi đến API của bên thứ ba.

---

## Tại Sao Cần Hai Token Thay Vì Một?

Đây là câu hỏi quan trọng mà nhiều tài liệu bỏ qua.

Nếu chỉ dùng một token duy nhất, có hai lựa chọn:

- **Token tồn tại ngắn:** Người dùng phải đăng nhập lại thường xuyên — trải nghiệm kém.
- **Token tồn tại dài:** Nếu bị đánh cắp, kẻ tấn công có thể sử dụng trong thời gian dài — rủi ro bảo mật cao.

Mô hình hai token giải quyết cả hai vấn đề trên bằng cách phân chia trách nhiệm rõ ràng:

|Tiêu chí|Access Token|Refresh Token|
|---|---|---|
|Mục đích|Gọi API, truy cập tài nguyên|Xin cấp Access Token mới|
|Thời gian sống|Ngắn (5 - 60 phút)|Dài (ngày, tuần)|
|Tần suất gửi|Mỗi request API|Chỉ khi Access Token hết hạn|
|Nơi gửi đến|Mọi API endpoint|Chỉ endpoint `/auth/refresh`|
|Rủi ro nếu bị lộ|Thấp — hết hạn nhanh|Cao — cần bảo vệ kỹ|
|Nơi lưu trữ|Bộ nhớ (in-memory)|HttpOnly cookie / secure storage|

Nhờ Access Token có thời gian sống ngắn, dù bị đánh cắp, kẻ tấn công chỉ có thể lợi dụng trong thời gian rất ngắn. Refresh Token sống lâu hơn nhưng được bảo vệ nghiêm ngặt hơn và chỉ giao tiếp với auth server — diện tích tấn công (attack surface) nhỏ hơn nhiều.

---

## Quy Trình Đầy Đủ: Cấp, Sử Dụng Và Làm Mới Token

```
1.  Client gửi thông tin đăng nhập (username/password hoặc social login)
2.  Server xác thực thông tin đăng nhập
3.  Server cấp:
        - Access Token  (ví dụ: hết hạn sau 15 phút)
        - Refresh Token (ví dụ: hết hạn sau 7 ngày)
4.  Client lưu:
        - Access Token  → bộ nhớ tạm (in-memory), không lưu localStorage
        - Refresh Token → HttpOnly cookie hoặc secure storage
5.  Client đính kèm Access Token vào mỗi request API
6.  Access Token hết hạn → Client gửi Refresh Token đến /auth/refresh
7.  Server xác minh Refresh Token (kiểm tra chữ ký, thời hạn, blacklist)
8.  Server cấp Access Token mới (và có thể cấp Refresh Token mới — token rotation)
9.  Lặp lại từ bước 5 cho đến khi Refresh Token hết hạn hoặc bị thu hồi
10. Refresh Token hết hạn hoặc bị thu hồi → Yêu cầu người dùng đăng nhập lại
```

---

## Tại Sao Access Token Quan Trọng?

Một tình huống thực tế:

Bạn dùng ứng dụng quản lý công việc kết nối với Google Calendar. Lúc đầu mọi thứ hoạt động tốt, nhưng sau một giờ, ứng dụng báo lỗi `401 Unauthorized` hoặc `Token expired`.

**Nguyên nhân:**

- Access Token đã hết hạn (thời gian sống thường chỉ 15 - 60 phút).
- Ứng dụng chưa implement cơ chế tự động làm mới token khi nhận được lỗi `401`.

**Giải pháp đúng đắn:**

Developer cần implement logic: khi nhận được lỗi `401`, tự động gọi endpoint `/auth/refresh` để lấy Access Token mới, sau đó thử lại request ban đầu — tất cả diễn ra trong suốt với người dùng.

Đây là lỗi phổ biến ở những developer chưa nắm rõ vòng đời của Access Token và cơ chế refresh.

---

## Access Token Và Refresh Token Được Lưu Ở Đâu?

Câu hỏi này quan trọng vì lưu sai nơi dẫn đến lỗ hổng bảo mật nghiêm trọng.

### Access Token

Được khuyến nghị lưu trong **bộ nhớ tạm của ứng dụng (in-memory)** — ví dụ: biến JavaScript, state của ứng dụng React/Vue. Lý do: in-memory không thể truy cập bởi script của trang khác, giảm nguy cơ XSS.

Hạn chế: khi người dùng tải lại trang, token bị mất — cần dùng Refresh Token để lấy lại.

### Refresh Token

Được khuyến nghị lưu trong **HttpOnly cookie**. Lý do:

- `HttpOnly`: JavaScript không thể đọc cookie này, ngăn chặn tấn công XSS đánh cắp token.
- `Secure`: cookie chỉ được gửi qua HTTPS.
- `SameSite=Strict` hoặc `SameSite=Lax`: ngăn chặn tấn công CSRF.

### Tuyệt Đối Không Lưu Trong `localStorage`

`localStorage` dễ bị đọc bởi bất kỳ JavaScript nào chạy trên trang — kể cả script từ tấn công XSS. Nếu trang web có lỗ hổng XSS, toàn bộ token trong `localStorage` có thể bị đánh cắp.

---

## Refresh Token Có Tồn Tại Mãi Mãi Không?

Không — Refresh Token không tồn tại vĩnh viễn. Có ba trường hợp Refresh Token mất hiệu lực:

**Hết hạn tự nhiên:** Refresh Token có thời gian sống được cấu hình sẵn (ví dụ: 7 ngày). Sau thời điểm đó, token tự động không còn hợp lệ.

**Bị thu hồi (revoke) chủ động:** Server có thể vô hiệu hóa Refresh Token bất cứ lúc nào bằng cơ chế **blacklist** — danh sách các token đã bị thu hồi. Khi nhận được Refresh Token, server kiểm tra danh sách này trước khi cấp Access Token mới. Các tình huống thu hồi thường gặp:

- Người dùng đăng xuất (logout).
- Người dùng đổi mật khẩu.
- Admin phát hiện tài khoản bị xâm phạm và khóa phiên.
- Phát hiện hoạt động bất thường (đăng nhập từ địa điểm lạ).

**Bị thay thế bởi token rotation:** Một số hệ thống áp dụng **Refresh Token Rotation** — mỗi lần dùng Refresh Token để lấy Access Token mới, server cũng cấp Refresh Token mới và vô hiệu hóa Refresh Token cũ. Nếu Refresh Token cũ được sử dụng lại sau khi đã bị thay thế, đây là dấu hiệu token có thể đã bị đánh cắp — server có thể thu hồi toàn bộ phiên đăng nhập.

---

## JWT Bị Lộ Thì Sao?

### Mức Độ Nguy Hiểm Thực Sự

Có — JWT bị lộ là nguy hiểm, nhưng mức độ phụ thuộc vào thiết kế hệ thống.

**Kịch bản nguy hiểm nhất:**

- Token bị đánh cắp qua tấn công XSS, log server, hoặc localStorage leak.
- Token chưa hết hạn.
- Hệ thống không có cơ chế revoke.

Kết quả: kẻ tấn công dùng token đó để gọi API — server vẫn chấp nhận vì JWT là stateless, server không phân biệt được ai đang cầm token.

### Server Có Phát Hiện Token Bị Đánh Cắp Không?

Theo bản chất của JWT stateless: **không thể phát hiện trực tiếp**.

JWT không lưu trạng thái trên server, nên:

- Server không biết token đang ở tay ai.
- Server không biết token có bị sao chép hay chưa.

Đây là đánh đổi (trade-off) cố hữu của JWT: đổi khả năng mở rộng (không cần lưu session) lấy khả năng kiểm soát từng token cụ thể.

### Cách Giảm Thiểu Rủi Ro

- **Đặt thời gian sống ngắn cho Access Token:** dù bị đánh cắp, cửa sổ tấn công nhỏ (15 phút thay vì 24 giờ).
- **Dùng HttpOnly cookie thay vì localStorage** để giảm nguy cơ XSS đánh cắp token.
- **Triển khai Refresh Token Rotation:** phát hiện reuse của Refresh Token cũ là dấu hiệu xâm phạm.
- **Xây dựng hệ thống blacklist** cho Refresh Token để có thể thu hồi phiên khi phát hiện bất thường.
- **Kết hợp thêm fingerprint** (IP, User-Agent) để phát hiện truy cập bất thường — dù không phải giải pháp hoàn hảo.

---

## Cách Sử Dụng Token An Toàn

### Nguyên Tắc Cơ Bản

- **Luôn dùng HTTPS:** đảm bảo token không bị nghe lén trong quá trình truyền tải.
- **Đặt thời gian sống ngắn cho Access Token:** khuyến nghị 15 - 60 phút.
- **Không lưu token nhạy cảm ở nơi có thể bị XSS đọc:** ưu tiên in-memory và HttpOnly cookie.
- **Triển khai Refresh Token đúng cách:** tự động làm mới mà không làm gián đoạn người dùng.

### Best Practice Triển Khai

- Lưu Access Token trong bộ nhớ tạm (in-memory), Refresh Token trong HttpOnly cookie.
- Áp dụng Refresh Token Rotation để phát hiện token bị đánh cắp.
- Giới hạn scope tối thiểu — chỉ cấp đúng quyền cần thiết cho từng use case.
- Ghi log và theo dõi các hoạt động bất thường: nhiều lần refresh liên tiếp, đăng nhập từ địa điểm bất thường.
- Thu hồi toàn bộ Refresh Token của người dùng khi phát hiện xâm phạm hoặc khi người dùng đổi mật khẩu.