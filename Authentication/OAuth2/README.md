
## **OAuth 2.0 là gì?**

**OAuth 2.0 (Open Authorization)** là một tiêu chuẩn mở giúp các ứng dụng có thể cấp quyền truy cập tài nguyên mà không cần chi sẻ thông tin đăng nhập (username/password). Nó thường được sử dụng để cung cấp quyền cho ứng dụng bên thứ 3 truy cập vào API dịch vụ thay mặt cho người khác.

![](https://docs.oracle.com/cd/E82085_01/160027/JOS%20Implementation%20Guide/Output/img/oauth2-caseflow.png)

---
## **Hoạt động như thể nào?**

---
## **Các thành phần chính trong OAuth 2.0**

##### **Resource Owner (Chủ sở hữu tài nguyên)**
- Là người dùng (hoặc hệ thống) sở hữu tài nguyên cần bảo vệ.
- *Ví dụ*: Bạn là chủ của tài khoản của Facebook của mình.
##### **Client (Ứng dụng khách)**
- Là ứng dụng muốn truy cập vào tài nguyên của người dùng.
- *Ví dụ:* Một ứng dụng bên thứ ba như Google Colab muốn truy cập vào Google Drive của bạn.
##### **Authorization Server (Máy chủ ủy quyền)**
- Xác thực người dùng và cung cấp quyền cho ứng dụng khách.
- Cấp phát mã (token) sau khi quyền truy cập được xác nhận.
- *Ví dụ:* Google OAuth Server.
##### **Resource Server (Máy chủ tài nguyên)**
- Lưu trữ và bảo vệ tài nguyên của người dùng.
- Xác thực và ủy quyền truy cập dựa vào token.
- *Ví dụ:* Google Drive API

---
## **Sơ đồ luồng hoạt động:**

1. **Ứng dụng (Web site hoặc mobile app)** yêu cầu ủy quyền để truy cập vào **Resource Server (Facebook, GitHub, Gmail, ...)** thông qua **User**
2. Nếu **User** ủy quyền cho yêu cầu trên, Ứng dụng sẽ nhận được sự ủy quyền từ phía User (dưới dạng một token string).
3. **Ứng dụng** gửi thông tin định danh (ID) của mình. Kèm theo ủy quyền từ phía **User** tới **Authorizatoin Server**.
4. Nếu thông tin định danh được xác định và ủy quyền hợp lệ thì **Authorization Server** sẻ trả về cho **ứng dụng** `access_token`. Đến đây thì quá trình ủy quyền đã hoàn tất.
5. Để truy cập vào tài nguyên (resource) từ **Resource Server** và lấy thông tin, **Ứng dụng** sẽ phải đưa ra `access_token` để xác thực.
6. Nếu `access_token` hợp lệ, Resource Server sẽ trả về dữ liệu của tài nguyên đã được yêu cầu cho **ứng dụng**.

![](https://topdev.vn/blog/wp-content/uploads/2019/03/mo-hinh-oauth.png)

⏩ **Lý do dùng:** Bảo mật tốt vì Access Token không lộ trên trình duyệt.

---
## **Access Token và Refresh Token**

##### **Access Token**
- Được dùng để xác thực và ủy quyền truy cập tài nguyên.
- Có thời hạn ngắn (từ vài phút đến vài giờ)
- Dùng trong header của HTTP request (`Authorization: Bearer <access_token>`)
##### **Refresh Token**
- Được dùng để lấy Access Token mới khia Access Token hết hạn.
- Có thời hạn dài hơn (vài ngày hoặc vài tuần)
- Không gửi trong mỗi request, mà chỉ cần dùng khi cần refresh Access Token.

⏩ **Lợi ích:** Hạn chế việc người dùng phải đăng nhập lại thường xuyên.

---
## **Cách bảo mật trong OAuth 2.0**

**Dùng HTTPS:** để tránh cuộc tấn công MITM (Man-in-the_Middle)
**Dùng PKCE (Proof Key for Code Exchange):** tăng cường bảo mật cho Authorization Code Flow.
**Hạn chế thời gian sống của Access Token:** Giảm thiểu rủi ro nếu bị đánh cắp
**Bảo mật Refresh Token:** Không lưu trữ trên client-side nếu không cần thiết.
**Dùng JWT (JSON Web Token):** Để mã hóa thông tin trong token.

---
## **Ưu & Nhược điểm:**
##### **Ưu điểm**
- **Bảo mật:** Không cần chia sẻ thông tin đăng nhập của người dùng
- **Linh hoạt:** Hỗ trợ nhiều loại ứng dụng và tình huống khác nhau.
- **Mở rộng:** Có thể mở rộng với các loại token và cơ chế bảo mật bổ sung,
##### **Nhược điểm:**
- **Phức tạp:** Việc triển khai đúng cách có hể phức tạp và dễ mắc lỗi về bảo mật.
- **Phụ thuộc vào token:** Nếu token bị đánh cắp, bên tấn công có thể truy cập vào tài nguyên.

---
## **Khi nào nên dùng OAuth 2.0?**

✅ Khi ứng dụng cần truy cập API của một dịch vụ bên thứ ba.  
✅ Khi ứng dụng cần hỗ trợ đăng nhập một lần (Single Sign-On - SSO).  
✅ Khi muốn cấp quyền truy cập tài nguyên mà không chia sẻ username/password.

---
## **Tài liệu tham khảo:**

1. [OAuth2 là gì? Tìm hiểu về OAuth2](https://topdev.vn/blog/oauth2-la-gi/)
2. [What the Heck is OAuth?](https://developer.okta.com/blog/2017/06/21/what-the-heck-is-oauth)
3. [An Introduction to OAuth 2](https://www.digitalocean.com/community/tutorials/an-introduction-to-oauth-2)