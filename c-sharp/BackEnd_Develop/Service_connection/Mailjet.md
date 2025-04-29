Chào các bạn sau đây tôi xin hướng dẫn cách gửi mail đơn giản bằng cách áp dụng Mailjet trong .net

Đầu tiên tải thư viện:
```
Mailjet.Api
```

Thiết lập các khóa trong `appsetting.json`
```json
"Mailjet":{
  "APIKEY_PUBLIC":"0xxxxxxxxxxxxxxxx",
  "APIKEY_PRIVATE":"1xxxxxxxxxxxxxxxx"
}
```

Lập ví dụ gửi mã OTP thông qua email người dùng
