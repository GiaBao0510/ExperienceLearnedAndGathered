1. Người dùng gửi request đến server thông qua API (GET /get-otp?email=client@gmail.com)
2. Server nhận được request và kiểm tra xem email trong Database có tồn tại không (nếu không thì báo lỗi)
3. Kiểm tra trong Redis xem có thông tin từ cooldownKey (otp:cooldown:{email}). nếu giá trị là `TRUE` thì phản hồi 429, ngược lại thì đến bước kết tiếp
- cooldownKey: là thời gian đếm ngắn. Kiểm tra xem thời gian gửi mã OTP cho người dùng
4. Tạo mã OTP có 6 chữ số, kiểu chuổi
5. Lưu 2 thông tin tạm thời vào REDIS:
- Key là otp:code:{email}, value sé là mã OTP (ví dụ: "otp:code:client@gmail.com":"123456")  -> giá trị mã OTP, TTL 5 phút    (dùng để verify)
- Key otp:cooldown:{email}, value sẽ là mã 1 (ví dụ: "otp:code:cooldown@gmail.com":1) -> flag đơn giản, TTL 30-60s      (dùng để throttle resend)

6. Gửi mail đến email người dùng có kèm theo mã OTP trong đó
7. Trả 200

---
***Ví dụ minh họa:***

```go
func (s *OTPService) SendOTP(ctx context.Context, email string) error {
    // 1. Kiểm tra email tồn tại trong hệ thống
    exists, err := s.userRepo.ExistsByEmail(ctx, email)
    if err != nil {
        return err
    }
    if !exists {
        return ErrEmailNotFound
    }

    // 2. Atomic check-and-lock cooldown — chỉ cần email, không cần biết OTP
    cooldownKey := "otp:cooldown:" + email
    locked, err := s.redis.SetNX(ctx, cooldownKey, "1", 45*time.Second).Result()
    if err != nil {
        return err
    }
    if !locked {
        ttl, _ := s.redis.TTL(ctx, cooldownKey).Result()
        return &TooManyRequestsError{RetryAfter: ttl}
    }

    // 3-4. Qua được cooldown -> generate OTP mới, lưu riêng key khác, TTL dài hơn
    otpCode := generateOTP()
    codeKey := "otp:code:" + email
    if err := s.redis.Set(ctx, codeKey, otpCode, 5*time.Minute).Err(); err != nil {
        return err
    }

    // 5-6. Gửi email
    return s.mailer.SendOTPEmail(email, otpCode)
}
```