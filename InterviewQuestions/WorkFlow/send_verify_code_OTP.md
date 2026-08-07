# OTP: Khái niệm, quy trình gửi - xác thực và các lưu ý bảo mật

## Mục lục

1. [OTP là gì?](#otp-là-gì)
2. [OTP API là gì?](#otp-api-là-gì)
3. [Các kênh gửi OTP phổ biến](#các-kênh-gửi-otp-phổ-biến)
4. [Quy trình gửi mã OTP (Send OTP)](#quy-trình-gửi-mã-otp-send-otp)
5. [Quy trình xác thực mã OTP (Verify OTP)](#quy-trình-xác-thực-mã-otp-verify-otp)
6. [Những lưu ý bảo mật quan trọng](#những-lưu-ý-bảo-mật-quan-trọng)
7. [Kết luận](#kết-luận)
8. [Mở rộng](#mở-rộng)

---

## OTP là gì?

OTP (One Time Password) là mã dùng một lần, được sử dụng phổ biến như một lớp xác thực bổ sung (thường gọi là xác thực hai yếu tố - 2FA) để xác minh danh tính người dùng. Nguyên lý hoạt động khá đơn giản: hệ thống gửi mã OTP tới một địa chỉ đã được xác định trước đó của người dùng (email, số điện thoại...), mã này chỉ có hiệu lực một lần duy nhất và trong một khoảng thời gian ngắn xác định, thường không quá 5 phút.

Dù nguyên lý đơn giản, việc triển khai OTP đúng cách lại liên quan đến nhiều khía cạnh bảo mật quan trọng mà cả pentester lẫn developer cần nắm rõ - những vấn đề này sẽ được trình bày chi tiết ở phần [Những lưu ý bảo mật quan trọng](#những-lưu-ý-bảo-mật-quan-trọng) phía sau, bởi đây thường là phần bị bỏ sót nhiều nhất khi triển khai OTP trong thực tế.

## OTP API là gì?

OTP API là giao diện lập trình cho phép hệ thống (website/app) tích hợp chức năng gửi và xác thực OTP, có thể tự xây dựng hoặc tích hợp qua dịch vụ bên thứ ba (ví dụ: Twilio, ESP email, hoặc SMS gateway trong nước).

Luồng hoạt động điển hình: người dùng đăng ký tài khoản → hệ thống gọi OTP API để tạo và gửi mã OTP qua SMS/email/app → người dùng nhập mã nhận được → hệ thống xác thực mã đó có đúng và còn hiệu lực hay không.

## Các kênh gửi OTP phổ biến

OTP có thể được gửi qua nhiều kênh khác nhau, tùy vào nhu cầu và ngân sách của từng hệ thống:

| Kênh | Ưu điểm | Hạn chế |
|---|---|---|
| **SMS OTP** | Độ phủ sóng rộng, không phụ thuộc Internet, phổ biến nhất tại Việt Nam | Chi phí cao, có thể bị delay khi mạng viễn thông quá tải |
| **Email OTP** | Chi phí thấp, dễ triển khai | Phụ thuộc Internet, dễ bị lọc vào mục spam |
| **Voice OTP** | Phù hợp với người dùng điện thoại không thông minh (feature phone) | Chi phí cao hơn cả SMS OTP |
| **App OTP (Authenticator)** | Bảo mật cao, hoạt động được cả khi không có mạng (dựa trên TOTP - thuật toán sinh mã theo thời gian) | Yêu cầu người dùng cài đặt và làm quen với ứng dụng riêng |
| **OTT OTP (Zalo, WhatsApp, Telegram...)** | Chi phí thấp, tỷ lệ mở tin nhắn cao | Chỉ hoạt động với người dùng đã cài ứng dụng nhắn tin tương ứng |

## Quy trình gửi mã OTP (Send OTP)

Quy trình gửi OTP tiêu chuẩn gồm các bước sau:

1. **Nhận request**: người dùng gửi yêu cầu lấy OTP đến server, ví dụ `POST /api/otp/send` kèm email.
2. **Validate định dạng đầu vào**: kiểm tra email có đúng định dạng hợp lệ hay không (thường thực hiện ở tầng handler/DTO, trước khi vào business logic).
3. **Kiểm tra email tồn tại trong hệ thống**: xác minh email đã được đăng ký trong database hay chưa.
4. **Kiểm tra cooldown trong Redis**: dùng key dạng `otp:cooldown:{email}` để kiểm tra người dùng có đang trong thời gian chờ giữa hai lần gửi hay không. Nếu key đã tồn tại (còn hiệu lực) → phản hồi mã lỗi `429 Too Many Requests`; nếu chưa → tiếp tục bước kế tiếp. Cooldown thường đặt trong khoảng 30-60 giây để chống spam.
5. **Sinh mã OTP**: tạo chuỗi số ngẫu nhiên gồm 6 chữ số.
6. **Lưu thông tin tạm thời vào Redis**, gồm 2 key riêng biệt với TTL khác nhau:
   - `otp:code:{email}` → lưu giá trị mã OTP, TTL khoảng 5 phút, dùng để đối chiếu ở bước xác thực.
   - `otp:cooldown:{email}` → lưu cờ đánh dấu (ví dụ giá trị `"1"`), TTL 30-60 giây, dùng để giới hạn tần suất gửi lại (resend throttling).
7. **Gửi mã OTP** đến người dùng qua kênh đã chọn (email, SMS...).
8. **Phản hồi kết quả** về client, ví dụ `200 OK` kèm thông báo đã gửi mã thành công.

### Ví dụ minh họa: Send OTP

```go
func (s *OTPService) SendOTP(ctx context.Context, email string) error {
    // 1. Kiểm tra email tồn tại trong hệ thống
    // Lưu ý: xem thêm mục "Những lưu ý bảo mật quan trọng" bên dưới
    // về rủi ro account enumeration khi trả lỗi này trực tiếp cho client.
    exists, err := s.userRepo.ExistsByEmail(ctx, email)
    if err != nil {
        return err
    }
    if !exists {
        return ErrEmailNotFound
    }

    // 2. Atomic check-and-lock cooldown bằng SetNX — đảm bảo không có
    // race condition khi nhiều request gửi đồng thời cho cùng một email
    cooldownKey := "otp:cooldown:" + email
    locked, err := s.redis.SetNX(ctx, cooldownKey, "1", 45*time.Second).Result()
    if err != nil {
        return err
    }
    if !locked {
        ttl, _ := s.redis.TTL(ctx, cooldownKey).Result()
        return &TooManyRequestsError{RetryAfter: ttl}
    }

    // 3. Qua được cooldown -> sinh OTP mới, lưu vào key riêng, TTL dài hơn cooldown
    otpCode := generateOTP()
    codeKey := "otp:code:" + email
    if err := s.redis.Set(ctx, codeKey, otpCode, 5*time.Minute).Err(); err != nil {
        return err
    }

    // 4. Gửi email chứa mã OTP
    return s.mailer.SendOTPEmail(email, otpCode)
}
```

## Quy trình xác thực mã OTP (Verify OTP)

Đây là phần thường bị bỏ sót khi tài liệu chỉ tập trung vào việc gửi OTP, trong khi bước xác thực mới là nơi tiềm ẩn nhiều rủi ro bảo mật nhất - đặc biệt là tấn công brute-force (thử toàn bộ 1.000.000 tổ hợp mã 6 chữ số). Quy trình xác thực chuẩn gồm:

1. **Nhận request xác thực**: client gửi email và mã OTP người dùng đã nhập, ví dụ `POST /api/otp/verify`.
2. **Kiểm tra và tăng số lần thử (attempt counter)**: dùng key dạng `otp:attempts:{email}` trong Redis, tăng giá trị mỗi lần xác thực. Nếu vượt quá giới hạn cho phép (ví dụ 5 lần trong 5 phút) → từ chối luôn, không cần kiểm tra mã OTP nữa.
3. **Lấy mã OTP đã lưu**: đọc giá trị từ key `otp:code:{email}`. Nếu key không tồn tại (đã hết hạn hoặc chưa từng gửi) → trả lỗi mã đã hết hạn hoặc không tồn tại.
4. **So sánh mã**: đối chiếu mã người dùng nhập với mã đã lưu trong Redis.
5. **Xử lý kết quả**:
   - Nếu đúng: xóa ngay key `otp:code:{email}` và `otp:attempts:{email}` để đảm bảo mã OTP không thể dùng lại lần thứ hai (chống replay attack), sau đó cho phép người dùng tiếp tục luồng nghiệp vụ (đăng ký, đăng nhập, đổi mật khẩu...).
   - Nếu sai: trả lỗi mã không hợp lệ, giữ nguyên attempt counter để tính vào giới hạn số lần thử.

### Ví dụ minh họa: Verify OTP

```go
func (s *OTPService) VerifyOTP(ctx context.Context, email, code string) error {
    attemptsKey := "otp:attempts:" + email
    codeKey := "otp:code:" + email

    // 1. Giới hạn số lần thử để chống brute-force
    attempts, err := s.redis.Incr(ctx, attemptsKey).Result()
    if err != nil {
        return err
    }
    if attempts == 1 {
        s.redis.Expire(ctx, attemptsKey, 5*time.Minute)
    }
    if attempts > 5 {
        return ErrTooManyAttempts
    }

    // 2. Lấy mã OTP đã lưu
    storedCode, err := s.redis.Get(ctx, codeKey).Result()
    if errors.Is(err, redis.Nil) {
        return ErrOTPExpiredOrNotFound
    }
    if err != nil {
        return err
    }

    // 3. So sánh mã
    if storedCode != code {
        return ErrInvalidOTP
    }

    // 4. Xác thực thành công -> xóa ngay để tránh dùng lại (replay attack)
    s.redis.Del(ctx, codeKey, attemptsKey)
    return nil
}
```

## Những lưu ý bảo mật quan trọng

Đây là nội dung cốt lõi mà tài liệu gốc chỉ nhắc tên mà chưa đi vào chi tiết. Khi triển khai OTP trong thực tế, cần đặc biệt lưu ý các điểm sau:

- **Chống brute-force ở bước xác thực**: mã OTP 6 chữ số chỉ có 1.000.000 khả năng - nếu không giới hạn số lần thử (như `otp:attempts:{email}` ở trên), kẻ tấn công hoàn toàn có thể dò được mã trong thời gian hiệu lực 5 phút bằng script tự động. Giới hạn số lần thử là bắt buộc, không phải tùy chọn.
- **Tránh account enumeration**: đoạn code ở bước gửi OTP trả về lỗi `ErrEmailNotFound` khi email chưa đăng ký - nếu lỗi này được trả trực tiếp cho client với thông báo rõ ràng, kẻ tấn công có thể dò ra được danh sách email nào đã đăng ký tài khoản trong hệ thống, chỉ bằng cách thử gửi OTP hàng loạt và quan sát phản hồi khác nhau. Cách khắc phục: luôn trả về một thông điệp chung chung như "Nếu email tồn tại, mã OTP đã được gửi" cho cả hai trường hợp tồn tại hoặc không tồn tại, đồng thời áp dụng rate limit theo IP để hạn chế dò quét hàng loạt.
- **Xóa OTP ngay sau khi xác thực thành công**: nếu không xóa, cùng một mã OTP có thể bị dùng lại nhiều lần trong thời gian còn hiệu lực (replay attack).
- **Không log mã OTP dưới dạng rõ (plaintext)** trong log hệ thống hoặc công cụ giám sát (Grafana/Loki), vì log thường được nhiều người trong tổ chức truy cập được.
- **Giới hạn tần suất gửi (cooldown) song song với giới hạn số lần xác thực (attempt limit)**: đây là hai lớp bảo vệ độc lập, riêng cooldown chỉ chống spam gửi mã, không chống được brute-force ở bước xác thực.
- **Cân nhắc dùng `crypto/rand` thay vì `math/rand`** khi sinh mã OTP trong Go, để đảm bảo mã được sinh ra có tính ngẫu nhiên đủ mạnh về mặt mật mã học (cryptographically secure), tránh trường hợp mã có thể bị dự đoán trước.
- **Luôn dùng HTTPS** cho toàn bộ luồng gửi/xác thực OTP để tránh mã bị đánh cắp qua kênh truyền trung gian (man-in-the-middle).

## Kết luận

OTP là một lớp bảo mật bổ sung hiệu quả và dễ triển khai về mặt khái niệm, nhưng độ an toàn thực tế phụ thuộc rất nhiều vào chi tiết triển khai: giới hạn số lần thử, tránh rò rỉ thông tin qua thông báo lỗi, xử lý TTL và xóa dữ liệu đúng cách trong Redis. Một hệ thống OTP chỉ dừng ở việc "gửi mã và so sánh chuỗi" mà bỏ qua các lớp bảo vệ này vẫn tiềm ẩn rủi ro bị khai thác nghiêm trọng, dù về mặt chức năng vẫn hoạt động bình thường.

### Mở rộng

Một số hướng tìm hiểu thêm để nâng cao kiến thức về chủ đề này:

- **TOTP (Time-based One Time Password)**: thuật toán đứng sau các ứng dụng như Google Authenticator, cho phép sinh mã OTP mà không cần server gửi qua mạng - tìm hiểu chuẩn RFC 6238 để hiểu cách hoạt động.
- **Constant-time comparison**: tìm hiểu package `crypto/subtle` trong Go (`subtle.ConstantTimeCompare`) để so sánh chuỗi một cách an toàn hơn trước tấn công timing attack, dù với OTP 6 chữ số kèm attempt limit thì mức độ ưu tiên của kỹ thuật này thấp hơn so với hệ thống so sánh token/password.
- **Rate Limiting nâng cao theo IP và theo thiết bị**: kết hợp với tài liệu Rate Limiting đã tìm hiểu trước đó (thuật toán GCRA qua `redis_rate`), áp dụng thêm một lớp giới hạn theo IP cho endpoint gửi/xác thực OTP để chống tấn công diện rộng.
- **Idempotency Key**: khi thiết kế API gửi OTP, tìm hiểu thêm về idempotency key để tránh trường hợp client gửi trùng request do lỗi mạng làm hệ thống gửi nhiều mã liên tiếp ngoài ý muốn.
- **So sánh OTP với WebAuthn/Passkey**: xu hướng xác thực không mật khẩu (passwordless) hiện đại, khắc phục nhiều điểm yếu bảo mật vốn có của OTP qua SMS/email.
- **Testing luồng OTP**: cách viết unit test cho service `SendOTP`/`VerifyOTP` bằng `go-redis/redismock` hoặc `miniredis`, tránh phải kết nối Redis thật khi chạy test.