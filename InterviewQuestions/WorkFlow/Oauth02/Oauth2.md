OAuth 2.0, viết tắt của **“Open Authorization”**, là một tiêu chuẩn được thiết kế để cho phép một trang web hoặc ứng dụng thay mặt người dùng truy cập các tài nguyên được lưu trữ bởi các ứng dụng web khác

Trước đây thì có OAuth 1.0, nhưng nó có nhiều vấn đề nên đã được thay thế bằng OAuth 2.0.

---
**Các bước trong workflow đăng nhập OAuth 2.0**

**Bước 1: Người dùng yêu cầu đăng nhập**
-  Người dùng nhấn nút "Đăng nhập bằng Google/Facebook/GitHub" trên ứng dụng của bạn (Client). [[1](https://duthanhduoc.com/blog/p5-giai-ngo-authentication-OAuth2)]

**Bước 2: Chuyển hướng (Redirect) đến trang xác thực**
- Ứng dụng chuyển hướng trình duyệt của người dùng đến máy chủ ủy quyền (Authorization Server) kèm theo `client_id`, `redirect_uri` và `scope` (quyền hạn yêu cầu)

**Bước 3: Người dùng cấp quyền**
- Người dùng đăng nhập vào nhà cung cấp (ví dụ: nhập email và password từ tài khoản Google) và đồng ý cho phép ứng dụng truy cập thông tin

**Bước 4: Nhận mã ủy quyền (Authorization Code)**
- Máy chủ ủy quyền chuyển hướng người dùng trở lại ứng dụng của bạn kèm theo một đoạn mã tạm thời gọi là `authorization_code` thông qua `redirect_uri`.

**Bước 5: Đổi mã lấy Access Token**
- Ứng dụng gửi ngầm (server-to-server) mã `authorization_code` cùng với `client_secret` trực tiếp lên máy chủ ủy quyền

**Bước 6: Trả về Access Token**
- Kiểm tra xem email này đã tồn tại hay chưa:
	- Nếu người dùng chưa tồn tại thì lưu thông tin mới này vào trong Database
	- Nếu đã tồn tại thì cập nhật thông tin đăng nhập (như thời điểm, IP,...)
- Máy chủ xác thực kiểm tra thông tin. Nếu hợp lệ, nó trả về `access_token` (và có thể có `refresh_token`)

**Bước 7: Truy cập tài nguyên**
- Ứng dụng dùng `access_token` này để gọi API lấy thông tin người dùng (tên, email, ảnh đại diện) từ máy chủ tài nguyên (Resource Server) và hoàn tất quá trình đăng nhập

---
### **Tài liệu tham khảo:**
1. [[P5] Giải ngố authentication: OAuth 2.0](https://duthanhduoc.com/blog/p5-giai-ngo-authentication-OAuth2)
2. [Tìm hiểu đôi chút về OAuth2](https://viblo.asia/p/tim-hieu-doi-chut-ve-oauth2-eW65GvMLlDO)