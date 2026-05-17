# OAuth 2.0

## OAuth 2.0 là gì?

**OAuth 2.0 (Open Authorization)** là một tiêu chuẩn mở (open standard) cho phép các ứng dụng được cấp quyền truy cập có giới hạn vào tài nguyên của người dùng trên một dịch vụ khác, mà không cần người dùng chia sẻ thông tin đăng nhập (username/password).

OAuth 2.0 giải quyết bài toán **ủy quyền (authorization)**, không phải xác thực (authentication). Đây là điểm thường bị nhầm lẫn: OAuth 2.0 trả lời câu hỏi _"Ứng dụng này được phép làm gì?"_, không phải _"Người dùng này là ai?"_. Nếu cần xác thực danh tính người dùng, cần dùng thêm **OpenID Connect (OIDC)** — một lớp xây dựng bên trên OAuth 2.0.

**Ví dụ thực tế:** Khi bạn dùng tính năng "Đăng nhập bằng Google" trên một ứng dụng bên thứ ba, ứng dụng đó không bao giờ nhìn thấy mật khẩu Google của bạn. Thay vào đó, Google cấp cho ứng dụng một token với phạm vi quyền truy cập giới hạn (ví dụ: chỉ đọc thông tin profile).

---

## Các thành phần chính trong OAuth 2.0

**Resource Owner (Chủ sở hữu tài nguyên)**

- Là người dùng (hoặc hệ thống) sở hữu tài nguyên cần bảo vệ.
- _Ví dụ:_ Bạn là chủ của tài khoản Google Drive của mình.

**Client (Ứng dụng khách)**

- Là ứng dụng muốn truy cập vào tài nguyên của người dùng. Client phải đăng ký với Authorization Server trước để nhận `client_id` và `client_secret`.
- _Ví dụ:_ Google Colab muốn truy cập vào Google Drive của bạn.

**Authorization Server (Máy chủ ủy quyền)**

- Xác thực người dùng, nhận sự đồng ý của người dùng, và cấp phát token cho Client.
- _Ví dụ:_ Google Authorization Server (`accounts.google.com`).

**Resource Server (Máy chủ tài nguyên)**

- Lưu trữ và bảo vệ tài nguyên của người dùng. Nhận request từ Client kèm `access_token`, xác thực token, và trả về dữ liệu nếu hợp lệ.
- Authorization Server và Resource Server có thể là cùng một hệ thống hoặc tách biệt nhau.
- _Ví dụ:_ Google Drive API.

---

## Luồng hoạt động (Authorization Code Flow)

OAuth 2.0 định nghĩa nhiều luồng (flow) khác nhau tùy theo loại ứng dụng. Luồng phổ biến và được khuyến nghị nhất là **Authorization Code Flow**, thường dùng cho ứng dụng web có backend.

```
Resource Owner (User)
        │
        │  (1) User truy cập Client, chọn "Đăng nhập bằng Google"
        ▼
   Client (App)
        │
        │  (2) Client chuyển hướng User đến Authorization Server
        │      kèm: client_id, redirect_uri, scope, state
        ▼
Authorization Server
        │
        │  (3) User đăng nhập và đồng ý cấp quyền
        │
        │  (4) Authorization Server chuyển hướng về redirect_uri
        │      kèm: authorization_code, state
        ▼
   Client (App)
        │
        │  (5) Client gửi authorization_code + client_secret
        │      đến Authorization Server (server-to-server, không qua browser)
        ▼
Authorization Server
        │
        │  (6) Xác thực, trả về access_token (và refresh_token)
        ▼
   Client (App)
        │
        │  (7) Client dùng access_token gọi Resource Server
        │      Header: Authorization: Bearer <access_token>
        ▼
 Resource Server
        │
        │  (8) Xác thực token, trả về dữ liệu được yêu cầu
        ▼
   Client (App)
```

**Lý do dùng `authorization_code` thay vì trả thẳng `access_token`:** Bước trao đổi `authorization_code` → `access_token` diễn ra qua kênh server-to-server (back channel), không đi qua trình duyệt. Điều này đảm bảo `access_token` không bao giờ bị lộ trong URL của trình duyệt hay lịch sử duyệt web.

**Tham số `state`:** Một chuỗi ngẫu nhiên do Client tạo ra, gửi đi ở bước (2) và nhận lại ở bước (4). Dùng để chống tấn công CSRF.

**Tham số `scope`:** Xác định phạm vi quyền truy cập mà Client yêu cầu. Ví dụ: `scope=read:profile read:email`. Resource Owner sẽ thấy danh sách quyền này khi đồng ý ủy quyền.

---

## Các Grant Type trong OAuth 2.0

OAuth 2.0 định nghĩa nhiều luồng cấp quyền (grant type) cho các tình huống khác nhau:

|Grant Type|Dùng khi|Mức độ khuyến nghị|
|---|---|---|
|Authorization Code|Web app có backend, mobile app|Khuyến nghị cao nhất|
|Authorization Code + PKCE|SPA, mobile app (không có backend bảo mật)|Khuyến nghị cho public client|
|Client Credentials|Giao tiếp server-to-server, không có user|Phù hợp cho M2M|
|Device Code|Thiết bị không có trình duyệt (Smart TV, CLI)|Dùng khi cần thiết|
|Implicit|(Đã lỗi thời)|Không nên dùng|
|Resource Owner Password|(Không khuyến nghị)|Tránh dùng|

---

## Access Token và Refresh Token

**Access Token**

- Được dùng để xác thực và ủy quyền truy cập tài nguyên trên Resource Server.
- Có thời hạn ngắn (thường từ vài phút đến 1 giờ) để giảm thiểu rủi ro nếu bị đánh cắp.
- Gửi trong header của mỗi HTTP request: `Authorization: Bearer <access_token>`.
- Định dạng phổ biến: JWT (có thể tự xác thực) hoặc opaque token (cần introspection endpoint).

**Refresh Token**

- Được dùng để lấy Access Token mới khi Access Token hết hạn, mà không cần người dùng đăng nhập lại.
- Có thời hạn dài hơn nhiều (từ vài ngày đến vài tuần, tùy cấu hình).
- Chỉ gửi đến Authorization Server khi cần refresh, không đính kèm trong mỗi request API thông thường.
- Cần được bảo vệ cẩn thận: nếu Refresh Token bị lộ, kẻ tấn công có thể lấy Access Token mới liên tục.

> **Lưu ý:** Không phải tất cả các flow đều cấp phát Refresh Token. Ví dụ, Client Credentials Flow thường không có Refresh Token vì Client có thể tự lấy Access Token mới bất cứ lúc nào.

---

## Bảo mật trong OAuth 2.0

**Dùng HTTPS bắt buộc:** Toàn bộ giao tiếp trong OAuth 2.0 phải đi qua HTTPS để chống tấn công Man-in-the-Middle (MITM). Không có ngoại lệ.

**Dùng PKCE (Proof Key for Code Exchange):** Bắt buộc với public client (SPA, mobile app) — những ứng dụng không thể bảo mật `client_secret`. PKCE sử dụng một cặp `code_verifier` / `code_challenge` để đảm bảo chỉ Client đã khởi tạo request mới có thể đổi `authorization_code` lấy token.

**Validate tham số `state`:** Luôn kiểm tra giá trị `state` trả về khớp với giá trị đã gửi, để chống CSRF.

**Giới hạn `scope`:** Chỉ yêu cầu các quyền thực sự cần thiết (principle of least privilege).

**Hạn chế thời gian sống của Access Token:** Token có TTL ngắn giúp giảm thiểu cửa sổ tấn công nếu token bị lộ.

**Bảo mật Refresh Token:**

- Lưu phía server (trong database/Redis), không lưu trong `localStorage` của trình duyệt.
- Với web app, dùng `HttpOnly` cookie để tránh bị đọc bởi JavaScript.
- Triển khai Refresh Token Rotation: mỗi lần dùng Refresh Token để lấy Access Token mới, cấp phát luôn Refresh Token mới và vô hiệu hóa cái cũ. Nếu phát hiện một Refresh Token cũ được dùng lại, đây là dấu hiệu token đã bị đánh cắp.

**Validate `redirect_uri`:** Authorization Server phải chỉ chấp nhận các `redirect_uri` đã được đăng ký trước, để tránh tấn công open redirect.

---

## Ưu và nhược điểm

**Ưu điểm**

- Người dùng không cần chia sẻ thông tin đăng nhập với ứng dụng bên thứ ba.
- Có thể cấp quyền truy cập có giới hạn (theo scope) thay vì toàn quyền.
- Hỗ trợ nhiều loại ứng dụng và tình huống qua các grant type khác nhau.
- Cho phép thu hồi quyền truy cập bất cứ lúc nào mà không cần đổi mật khẩu.

**Nhược điểm**

- Triển khai đúng cách tương đối phức tạp, dễ mắc lỗi bảo mật nếu không hiểu rõ từng bước.
- OAuth 2.0 chỉ xử lý ủy quyền, không phải xác thực — cần thêm OpenID Connect nếu cần định danh người dùng.
- Nếu token bị đánh cắp trước khi hết hạn, kẻ tấn công có thể truy cập tài nguyên trong khoảng thời gian đó.

---

## Khi nào nên dùng OAuth 2.0?

- Khi ứng dụng cần truy cập API của một dịch vụ bên thứ ba thay mặt người dùng (ví dụ: đọc Google Calendar, đăng bài lên Twitter).
- Khi muốn hỗ trợ đăng nhập một lần (Single Sign-On — SSO) qua nhiều ứng dụng.
- Khi muốn cấp quyền truy cập tài nguyên mà không chia sẻ username/password.
- Khi cần giao tiếp an toàn giữa các service nội bộ (dùng Client Credentials Flow).

---

## Đề xuất cải thiện thêm

### 1. Phân biệt rõ OAuth 2.0 và OpenID Connect (OIDC)

Đây là điểm gây nhầm lẫn phổ biến nhất với OAuth 2.0. Tính năng "Đăng nhập bằng Google" mà nhiều người hay nhắc đến thực chất là **OpenID Connect**, không phải OAuth 2.0 thuần túy.

- **OAuth 2.0:** Chỉ xử lý ủy quyền (authorization). Trả về `access_token`.
- **OpenID Connect:** Lớp xác thực (authentication) xây dựng trên OAuth 2.0. Trả về thêm `id_token` (JWT chứa thông tin định danh người dùng như `email`, `name`, `sub`).

Nếu ứng dụng chỉ cần truy cập API thay mặt user → OAuth 2.0 là đủ. Nếu cần biết user là ai → cần OIDC.

### 2. Lựa chọn grant type theo loại ứng dụng

Một điểm quan trọng mà tài liệu gốc bỏ qua là không phải lúc nào cũng dùng Authorization Code Flow:

- **Web app có backend:** Authorization Code Flow (bảo mật nhất, giữ `client_secret` phía server).
- **SPA hoặc mobile app:** Authorization Code Flow + PKCE (không có `client_secret`, bù lại bằng PKCE).
- **Microservice gọi nhau, không có user:** Client Credentials Flow.
- **Tránh hoàn toàn:** Implicit Flow (đã bị RFC 9700 khuyến nghị không dùng) và Resource Owner Password Flow (yêu cầu user nhập password trực tiếp vào Client, vi phạm nguyên tắc cốt lõi của OAuth).

### 3. Token Introspection và Token Revocation

Khi dùng opaque token (không phải JWT), Resource Server cần gọi **Token Introspection endpoint** (RFC 7662) để kiểm tra tính hợp lệ của token với Authorization Server.

Để vô hiệu hóa token trước khi hết hạn (ví dụ khi user logout), dùng **Token Revocation endpoint** (RFC 7009).

### 4. Phiên bản tiếp theo: OAuth 2.1

OAuth 2.1 (đang trong quá trình chuẩn hóa) hợp nhất các best practice từ nhiều RFC vào một tài liệu duy nhất, bao gồm: bắt buộc PKCE cho tất cả Authorization Code Flow, loại bỏ Implicit Flow và Resource Owner Password Flow. Nên theo dõi nếu đang thiết kế hệ thống mới.

---

## Tài liệu tham khảo

1. [RFC 6749 — The OAuth 2.0 Authorization Framework](https://datatracker.ietf.org/doc/html/rfc6749)
2. [RFC 7636 — PKCE for OAuth Public Clients](https://datatracker.ietf.org/doc/html/rfc7636)
3. [RFC 7009 — OAuth 2.0 Token Revocation](https://datatracker.ietf.org/doc/html/rfc7009)
4. [OAuth 2.0 Security Best Current Practice](https://datatracker.ietf.org/doc/html/draft-ietf-oauth-security-topics)
5. [OAuth2 là gì? Tìm hiểu về OAuth2 — TopDev](https://topdev.vn/blog/oauth2-la-gi/)
6. [What the Heck is OAuth? — Okta Developer](https://developer.okta.com/blog/2017/06/21/what-the-heck-is-oauth)
7. [An Introduction to OAuth 2 — DigitalOcean](https://www.digitalocean.com/community/tutorials/an-introduction-to-oauth-2)