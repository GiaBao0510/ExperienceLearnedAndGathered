### **Giới thiệu:**

![](https://avatao.com/media/JWT-blogpost.png)

JSON Web Token (JWT) là một phương tiện đại diện cho các yêu cầu chuyển giao giữa 2 bên **Client-Server**, nhằm ==xác minh thông tin an toàn giữa các bên Client-Server dưới dạng JSON object==. Thông tin này có thể được xác minh và tin cậy vì nó được ký điện tử - **digitally signed**. JWT có thể được ký bằng cách sử dụng một **secret (với thuật toán HMAC)** hoặc cặp **public/private key** dùng chuẩn RSA hoặc ECDSA.

**Signed tokens** có thể xác minh tính toàn vẹn các **claim** có trong đó, trong khi **encrypted** ==ẩn đi các claim từ các bên khác==. Khi token được đăng ký bởi cặp khóa **public/private keys**, signature cũng xác nhận rằng chỉ có bên giữa **private key** là nơi đã đăng ký nó.

---
### **Tại sao phải sử dụng JSON Web Token (JWT) để bảo mật API**

**1. Tính bảo mật:** JWT giúp xác thực danh tính người dùng một cách an toàn. Token được ký số (thường bằng HMAC hoặc RSA) nên không thể bị giả mạo. Server có thể xác minh token mà không cần truy cập cơ sở dữ liệu mỗi lần.

**2. Không trạng thái (Stateless):** ==Server không cần lưu thông tin phiên làm việc (session)==. Tất cả thông tin xác thực được chứa trong token, giúp hệ thống dễ dàng mở rộng theo chiều ngang (scalable).

**3. Gọn nhẹ và nhanh chóng:** JWT là chuỗi string đơn giản (base64), dễ truyền qua HTTP headers, cookies hoặc query string, nên rất thích hợp cho giao tiếp giữa client và API.

**4. Dễ tích hợp phân quyền:** Bạn có thể nhúng các quyền (roles, permissions) vào payload của JWT để kiểm tra phân quyền ở bất kỳ đâu mà không cần truy vấn lại cơ sở dữ liệu.

**5. Đa nền tảng:** JWT hoạt động tốt với các frontend hiện đại (React, Angular, Mobile...) vì client có thể dễ dàng lưu trữ và gửi token khi gọi API.

### 🧠 Ví dụ:
> Giống như một vé vào cửa sự kiện: bạn chỉ cần đưa vé (JWT) cho bảo vệ (API), nếu vé hợp lệ và chưa hết hạn, bạn được vào. Không cần hỏi lại hệ thống mỗi lần.

---
### **Khi nào nên dùng JSON Web Tokens?**

- **Ủy quyền - Authorization:** Đây là trường hợp nên sử dụng JWT. Khi người dùng đã đăng nhập, ==mỗi request tiếp theo gửi từ Client sẽ bao gồm JWT==, cho phép ==người dùng được access vào routes, service and resources== được phép với token đó. **Single Sign ON** là tính năng sử dụng JWT rộng rãi hiện nay, vì chi chi phí thấp dễ dàng sử dụng các Domains khác nhau
- **Trao đổi thông tin - Information Exchange:** JSON Web Tokens là một cách tốt để truyền thông tin an toàn giữa các bên Client và Server. Vì JWT có thể **signed**. Ví dụ , sử dụng các cặp public/private key, bạn có thể biết chắc người gửi. Ngoài ra, vì **signature** được xác định dựa vào **header** và **payload**, bạn cũng có thể xác minh rằng nội dung chưa bị giả mạo

---
### **Cấu trúc JSON Web Tokens**

![](https://topdev.vn/blog/wp-content/uploads/2017/12/jwt-la-gi.jpeg)

JSON Web Token bao gồm 3 phần được phân tách bằng **dấu chấm** **`"."`**

1. Header 
2. Payload 
3. Signature (Chữ ký)

Do đó, JWT thường trông như sau: `xxxxx.yyyyy.zzzzz`

##### 1.**Header**
Trong header gồm có 2 phần, đó là: loại mã token (là JWT); và thuật toán sử dụng, chẳng hạn SHA256, HMAC hay RSA.

***Ví dụ:***
```json
{
	"alg":"HS256",
	"typ": "JWT"
}
```

- **"typ" (Type)** chỉ ra rằng đối tượng là một JWT
- **"alg" (alogorithm)** xác định thuật toán mã hóa cho chuỗi là HS256

Sau đó, JSON này được mã hóa **Base64Url** để tạo thành phần đầu tiên của JWT.

##### 2.**Payload**
Phần thứ 2 của Token là payload, nó chứa các **the claims**, **Claims** thường chứa các thuộc tính như: **typically**, thông tin ***user*** và các dữ liệu bổ sung. Có 3 loại claims: **registered**, **public** và **private claims**.

**2.1.Registered claims:** Đây là một ==tập hợp các claims được xác định trước== không bắt buộc (được khuyến nghị), để cung cấp một tập hợp các claims hữu ích, có thể tương tác, Thường là: iss (nhà phát hành), exp (thời gian hết hạn), sub (chủ đề), aud (audience và những thứ khác).

Một số registered claims bao gồm:
- **`iss` (Issuer):** tổ chức, đơn vị cung cấp, phát hành JWT
- **`sub` (Subject):** Chủ thể JWT, xác định rằng đây là người sở hữu hoặc có quyền truy cập các resource (tài nguyên).
- **`aud` (Audience):** Được hiểu là người nhận thông tin, và có thể xác thực tính hợp lệ của JWT.
- **`exp` (Expiration time):** Thời hạn của JWT, vượt quá thời gian này, JWT coi là không hợp lệ.

> Lưu ý là claims names thường chỉ chứa 3 ký tự.

**2.2.Public claims:** Chúng có thể được ==xác định theo ý muốn của những người sử dụng JWT==. Nhưng để tránh xung đột, chúng ta phải được xác định trong **IANA JSON Web Token Registry** hoặc được định nghĩa là URI chứa namespace chống xung đột.

|Claim Name|Claim Description|
|---|---|
|iss|Issuer|
|sub|Subject|
|aud|Audience|
|exp|Expiration Time|
|nbf|Not Before|
|iat|Issued At|
|jti|JWT ID|
|name|Full name|
|given_name|Given name(s) or first name(s)|
|family_name|Surname(s) or last name(s)|
|middle_name|Middle name(s)|
|nickname|Casual name|
|preferred_username|Shorthand name by which the End-User wishes to be referred to|
|profile|Profile page URL|
|picture|Profile picture URL|
Một số _public claims_ điển hình :
- `name, given_name, family_name, middle_name`: Thông tin tên nói chung của user
- `email`: Thông tin email của user.
- `locale` : Địa chỉ của user.
- `profile, picture` : Thông tin của trang web gửi đến.

**2.3.Private claims:** Đây là ==các claims tùy chỉnh được tạo để chia sẻ thông tin giữa các bên đồng ý sử dụng== chúng và không phải là các registered hay public claims.

Ví dụ payload: 
```json
{
	"sub": "123456789",
	"name": "Gia Bao",
	"role": "admin"
}
```

Payload sẽ được mã hóa Base64URl để tạo thành phần thứ 2 của JSON Web Token.

> ***⚠️Lưu ý:** đối với các signed tokens, thông tin này mặc dù được bảo vệ để chống giả mạo, nhưng mọi người đều có thể đọc được, ==không nên đưa thông tin bảo mật vào các phần tử payload hoặc header==. Trừ khi được mã hóa.*

##### 3.**Signature**
Để tạo **signature** ==bạn phải lấy header được mã hóa, payload được mã hóa, một secret, thuật toán được chỉ định trong header và sign==. Ví dụ bạn dùng thuật toán HMAC SHA256, signature sẽ được tạo ra như sau:
```
HMACSHA256(
	baser64UrlEncode(header)+ "." +
	baser64UrlEncode(payload),
	secret
)
```
- **base64UrlEncoder**: thuật toán mã hóa **header** và **payload**

Đoạn code trên sau khi mã hóa **header** và **payload** bằng thuật toán `base64UrlEncode` ta sẽ được chuỗi như sau:
```
// header
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9
// payload
eyJhdWQiOlsidGVzdGp3dHJlc291cmNlaWQiXSwidXNlcl9uYW1lIjoiYWRtaW4iLCJzY29wZSI6WyJyZWFkIiwid3JpdGUiXSwiZXhwIjoxNTEzNzE
```

Sau khi mã hóa 2 chuỗi trên kèm theo **secret (khóa bí mật)** bằng thuật toán HS256 ta sẽ có chuỗi **signature** như sau:
```
9nRhBWiRoryc8fV5xRpTmw9iyJ6EM7WTGTjvCM1e36Q
```

Cuối cùng, kết hợp 3 chuỗi trên ta sẽ được một chuỗi **JWT** hoàn chỉnh.
```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsidGVzdGp3dHJlc291cmNlaWQiXSwidXNlcl9uYW1lIjoiYWRtaW4iLCJzY29wZSI6WyJyZWFkIiwid3JpdGUiXSwiZXhwIjoxNTEzNzE.9nRhBWiRoryc8fV5xRpTmw9iyJ6EM7WTGTjvCM1e36Q
```

**Signature** được sử dụng để xác minh tin nhắn không bị thay đổi trên đường truyền và trong trường hợp token được ký bằng **private key**, nó cũng có thể **xác minh người gửi jwt**.

---
### **JSON Web Tokens hoạt động như thế nào?**

![](https://media2.dev.to/dynamic/image/width=800%2Cheight=%2Cfit=scale-down%2Cgravity=auto%2Cformat=auto/https%3A%2F%2Fdev-to-uploads.s3.amazonaws.com%2Fuploads%2Farticles%2Figvye1880i9491jbz525.png)

Trong xác thực, khi người dùng đăng nhập thành công bằng thông tin đăng nhập của họ. JSON Web Token sẽ được trả về. Vì token là thông tin xác thực, cần phải cần thận về vấn đề bảo mật. Không nên giữ token lâu hơn yêu cầu.

Không nên lưu dữ liệu nhạy cảm trên **session** trong bộ nhớ trình duyệt do thiếu bảo mật.

Bất cứ khi nào người dùng muốn truy cập **route** hoặc **resource** được bảo vệ, tác nhân người dùng nên gửi JWT, thêm **Authorization** trong header với nội dung là **Bearer + token.** Nội dung của header trong như sau: `Authorization: Bearer <token>`

Máy chủ nhận **server** sẽ kiểm tra tính hợp lệ JWT trong header mỗi khi nhận **request**, nếu hợp lệ người dùng sẽ được phép truy cập các `resource` được bảo vệ. Nếu JWT chứa dữ liệu cần thiết, nhu cầu truy vấn cơ sở dữ liệu cho các hoạt động nhất định thì sẽ có thể bị giảm, mặc dù điều này có thể không phải luôn luôn như vậy.

Nếu **token** được gửi trong **Authorization header**, Chia sẻ tài nguyên nguồn góc chéo (Cross-Origin Resource Sharing - CORS) sẽ không thành vấn đề vì nó không sử dụng Cookie.

Sau đây là sơ đồ minh họa cách **JWT** được lấy và sử dụng để truy cập API và Resource:
![](https://images.viblo.asia/36f83306-2f90-4b34-ae0e-18b6e2956b37.png)

1. **Application** hoặc **client request authorization** đến **Authorization server.** Điều này được thực hiện thông qua một trong các luồng **authorization** khác nhau. *Ví dụ:* một ứng dụng web tuân thủ OpenID Connect điển hình sẽ đi qua /oauth/ ủy quyền đến điểm cuối bằng cách sử dụng luông mã authorization
2. Khi **authorization** được cấp, **authorization server** sẽ trả lại access token cho application
3. Application sẽ sử dụng access token để truy cập vào resource (Api).

---
### **Ưu/nhược điểm:**

##### **Ưu điểm:**
- **Gọn nhẹ:** JWT nhỏ gọn, chi phí truyền tải thấp giúp tăng hiệu suaart các ứng dụng
- **Bảo mật:** JWT sử dụng các mật mã khóa để tiến hành xác thực danh tính người dùng. Ngoài ra, cấu trúc của JWT cho phép chống giả mạt nên thông tin được đảm bảo an toàn trong quá trình trao đổi.
- **Phổ thông:** JWT được sử dụng dựa trên JSON, là một dạng dữ liệu phổ biến, có thể sử dụng hầu hết ở các ngôn ngữ lập trình. Ngoài ra triển khai JWT tương đối dễ dàng và tích hợp nhiều với thiết bị, vì JWT đã tương đối phổ biến.

##### **Nhược điểm:**
- **Kích thước:** Nhưng do kích thước được truyển trên **HTTP Header**, vì thế. JWT có giới hạn tương đương với **HTTP Header** (khoảng 8KB).
- **Rủi ro bảo mật:** Khi sử dụng JWT không đúng cách, ví dụ: như không kiểm tra tính hợp lệ của signature, không kiểm tra expire time, kẻ tấn công có thể lợi dụng kẻ hở để truy cập các thông tin trái phép. (Ngoài ra việc để ra thời gian hết hạn của JWT quá dài cũng có thể tạo ra kẻ hở tương tự).

---
### **Một số ứng dụng JWT:**

- **Single Sign-On (SSO):** JWT có thể được sử dụng cung cấp single sign-on cho người dùng. Điều này cho phép họ đăng nhập vào nhiều ứng dụng với một tài khoản duy nhất.
- **API Authorization:** JWT thường được sử dụng để  phân quyền cho người dùng đến những tài nguyên cụ thể, từ những claims chứa trong JWT đó.
- **User Authentication:** JWT cung cấp khả năng xác thực người dùng và cấp quyền cho họ truy cập vào các tài nguyên mong muốn trong hệ thống 
- **Microservices Communication:** _JWT_ còn có thể sử dụng cho việc giao tiếp giữa các _service_ nhỏ trong hệ thống _microservices_.