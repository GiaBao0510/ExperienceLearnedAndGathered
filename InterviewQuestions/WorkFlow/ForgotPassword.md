1. Người dùng Gửi API (POST /forgot-password) quên mật khẩu (kèm theo số điện thoại hoặc email)
2. Xác nhận xem email hoặc số điện thoại có tồn tại trong database không để tìm người dùng (nếu không thì báo lỗi)
3. Nếu có thì soạn mã OTP lưu vào redis với nội dung:
- SETNX: 
 gửi vào