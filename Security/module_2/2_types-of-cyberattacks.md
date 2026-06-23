# Các Loại Tấn công Mạng (Types of Cyberattacks)

---

## Denial-of-Service (DoS) — Tấn công từ chối dịch vụ

Tấn công DoS là hành vi cố ý gây ra sự cố hoàn toàn hoặc một phần đối với một hệ thống, dịch vụ, hoặc mạng bằng cách áp đảo tài nguyên của nó, khiến người dùng hợp lệ không thể truy cập.

Có hai cơ chế chính:

- **Flood-based (Lũ lụt lưu lượng):** Kẻ tấn công gửi lượng lớn request hoặc dữ liệu đến mục tiêu, làm cạn kiệt băng thông hoặc khả năng xử lý. Tương tự tắc đường: quá nhiều xe khiến giao thông ngừng trệ hoàn toàn.
- **Resource exhaustion (Cạn kiệt tài nguyên):** Kẻ tấn công gửi các request được thiết kế để tiêu thụ tài nguyên hệ thống (CPU, RAM, kết nối) một cách không cân xứng với kích thước của chúng, khiến hệ thống chậm lại hoặc sập.

> **Ví dụ — XML Bomb (Billion Laughs Attack):** Kẻ tấn công tạo một tài liệu XML nhỏ chứa các entity lồng nhau. Khi bộ xử lý XML mở rộng các entity này, kích thước dữ liệu tăng theo hàm mũ — một tài liệu vài kilobyte có thể mở rộng thành gigabyte dữ liệu trong bộ nhớ, làm cạn kiệt RAM và sập hệ thống. Kỹ thuật này nhắm vào bộ xử lý XML, không phải băng thông mạng.

> 

---

## Distributed Denial-of-Service (DDoS) — Tấn công từ chối dịch vụ phân tán

DDoS là biến thể quy mô lớn của DoS: thay vì một nguồn duy nhất, cuộc tấn công phát sinh từ hàng trăm đến hàng triệu thiết bị đồng thời — khiến việc chặn theo địa chỉ IP trở nên vô hiệu.

**Cơ chế — Botnet:**

Kẻ tấn công xây dựng **botnet** (mạng lưới bot) bằng cách lây nhiễm phần mềm độc hại vào các thiết bị kết nối internet của người dùng thông thường. Các thiết bị bị nhiễm này (còn gọi là **zombie**) nhận lệnh từ kẻ tấn công thông qua một máy chủ điều khiển trung tâm (**C2 — Command and Control server**). Khi ra lệnh tấn công, toàn bộ botnet đồng loạt gửi lưu lượng đến mục tiêu.

Botnet hiện đại có thể bao gồm hàng triệu thiết bị, bao gồm cả các thiết bị IoT (camera an ninh, router gia đình, thiết bị thông minh) — thường có bảo mật yếu. **Mirai botnet** (2016) là ví dụ nổi tiếng, tập hợp hàng trăm nghìn thiết bị IoT để thực hiện DDoS lớn nhất từng được ghi nhận tại thời điểm đó (620 Gbps).

AI đang được ứng dụng vào DDoS để tối ưu hóa số lượng bot cần thiết, điều chỉnh cường độ tấn công theo thời gian thực, và giám sát hiệu quả từng nguồn tấn công — làm tăng đáng kể hiệu quả trong khi giảm tài nguyên cần thiết.

> **Ví dụ:** Kẻ tấn công ra lệnh cho botnet gồm 500.000 thiết bị đồng loạt gửi request đến một máy chủ thương mại điện tử. Ngay cả khi mỗi thiết bị chỉ gửi một vài request mỗi giây, tổng lưu lượng có thể đạt hàng trăm gigabits, áp đảo hoàn toàn hạ tầng mục tiêu.

---

## Phishing — Lừa đảo trực tuyến

Phishing là kỹ thuật tấn công kết hợp **social engineering** (thao túng tâm lý) và thủ thuật kỹ thuật để lừa người dùng tiết lộ thông tin nhạy cảm hoặc thực hiện hành động gây hại — bằng cách giả mạo một nguồn tin cậy.

Kẻ tấn công thường giả mạo: ngân hàng, nhà cung cấp dịch vụ email, tổ chức chính phủ, hoặc đồng nghiệp. Sau khi người dùng tương tác với nội dung độc hại, hậu quả có thể là: cài đặt malware, bật kết nối điều khiển từ xa (RAT), hoặc tiết lộ thông tin xác thực.

**Phân loại phishing theo kênh:**

| Loại           | Kênh       | Mô tả                                |
| -------------- | ---------- | ------------------------------------ |
| Email phishing | Email      | Phổ biến nhất, gửi hàng loạt         |
| Smishing       | SMS        | Lừa đảo qua tin nhắn văn bản         |
| Vishing        | Điện thoại | Lừa đảo qua cuộc gọi thoại           |
| Quishing       | QR Code    | Chèn QR Code độc hại vào tài liệu in |

> **Ví dụ:** Kẻ tấn công gửi email giả mạo thông báo của ngân hàng, yêu cầu người nhận "xác minh tài khoản" bằng cách nhấp vào liên kết. Liên kết dẫn đến trang đăng nhập giả mạo có giao diện y hệt trang ngân hàng thật — khi người dùng nhập thông tin, thông tin xác thực được gửi thẳng đến kẻ tấn công.

> **Ví dụ:** Kẻ tấn công gửi email giả mạo thông báo của ngân hàng, yêu cầu người nhận "xác minh tài khoản" bằng cách nhấp vào liên kết. Liên kết dẫn đến trang đăng nhập giả mạo có giao diện y hệt trang ngân hàng thật — khi người dùng nhập thông tin, thông tin xác thực được gửi thẳng đến kẻ tấn công.

---

## Spear Phishing — Lừa đảo có chủ đích

Spear phishing là phiên bản nâng cao của phishing: thay vì tấn công hàng loạt không phân biệt mục tiêu, kẻ tấn công nhắm vào **một cá nhân, nhóm, hoặc tổ chức cụ thể** và đầu tư công sức nghiên cứu trước để cá nhân hóa nội dung tấn công.

**Điểm khác biệt then chốt với phishing thông thường:**

Phishing thông thường dùng nội dung chung chung ("Tài khoản của bạn bị khóa"). Spear phishing dùng thông tin cụ thể về nạn nhân — tên thật, chức vụ, tên đồng nghiệp, dự án đang làm việc, giao dịch gần đây — khiến nạn nhân khó nghi ngờ hơn nhiều.

Kẻ tấn công thu thập thông tin từ: mạng xã hội (LinkedIn, Facebook), website công ty, báo cáo thường niên, và thậm chí từ các vụ rò rỉ dữ liệu trước đó. AI hiện được dùng để tự động hóa quá trình thu thập và tổng hợp thông tin này, sau đó sinh ra nội dung email được cá nhân hóa cao.

**Whaling** là dạng spear phishing nhắm vào lãnh đạo cấp cao (C-suite): CEO, CFO, CISO — những người có quyền phê duyệt giao dịch tài chính lớn hoặc truy cập vào dữ liệu chiến lược nhất.

> **Ví dụ:** Kẻ tấn công nghiên cứu LinkedIn để biết rằng nạn nhân là kế toán trưởng tại một công ty xây dựng, và CEO công ty đang ở nước ngoài. Kẻ tấn công gửi email giả danh CEO, yêu cầu chuyển tiền khẩn cấp để thanh toán hợp đồng nhà cung cấp — thường gọi là **Business Email Compromise (BEC)**. Vụ tấn công này thuyết phục vì sử dụng đúng tên, ngữ cảnh công việc, và tạo cảm giác khẩn cấp.

---

## Malware — Phần mềm độc hại

Malware (viết tắt của **Malicious Software**) là thuật ngữ tổng quát chỉ mọi phần mềm được thiết kế để gây hại, gián đoạn, hoặc truy cập trái phép vào hệ thống mà không có sự đồng ý của chủ sở hữu.

Malware thường được cài vào hệ thống qua các vector phổ biến: tệp đính kèm email, liên kết độc hại, phần mềm crack/keygen, USB lây nhiễm, hoặc khai thác lỗ hổng phần mềm chưa được vá.

**Các loại malware phổ biến:**

| Loại       | Cơ chế                                  | Mục tiêu                              |
| ---------- | --------------------------------------- | ------------------------------------- |
| Virus      | Lây lan bằng cách chèn vào file hợp lệ  | Phá hủy dữ liệu, lây lan              |
| Worm       | Tự lây lan qua mạng không cần file host | Lây lan nhanh, tấn công hàng loạt     |
| Trojan     | Giả mạo phần mềm hợp lệ                 | Mở cửa hậu, đánh cắp dữ liệu          |
| Ransomware | Mã hóa dữ liệu của nạn nhân             | Tống tiền                             |
| Spyware    | Giám sát và báo cáo hoạt động           | Đánh cắp thông tin                    |
| Keylogger  | Ghi lại phím bấm                        | Đánh cắp mật khẩu, thông tin thẻ      |
| Rootkit    | Ẩn sâu trong hệ điều hành               | Duy trì quyền truy cập bí mật dài hạn |
| Adware     | Hiển thị quảng cáo trái phép            | Doanh thu quảng cáo, thu thập hành vi |

> AI đang được dùng để phát triển malware có khả năng **polymorphic** (thay đổi mã nguồn sau mỗi lần lây nhiễm) hoặc **metamorphic** (tái cấu trúc hoàn toàn), giúp né tránh phát hiện dựa trên chữ ký (signature-based detection).

AI đang được dùng để phát triển malware có khả năng **polymorphic** (thay đổi mã nguồn sau mỗi lần lây nhiễm) hoặc **metamorphic** (tái cấu trúc hoàn toàn), giúp né tránh phát hiện dựa trên chữ ký (signature-based detection).

> **Ví dụ:** LockBit ransomware mã hóa toàn bộ file trên hệ thống của nạn nhân, sau đó yêu cầu thanh toán bằng cryptocurrency để nhận key giải mã. Trong một số chiến dịch, nhóm LockBit còn đe dọa công bố dữ liệu đánh cắp nếu nạn nhân không trả tiền — kỹ thuật này gọi là **double extortion**.

---

## Man-in-the-Middle (MitM) — Tấn công trung gian

Tấn công MitM xảy ra khi kẻ tấn công bí mật xen vào giữa kênh truyền thông giữa hai bên — thường là client và server — mà cả hai bên không hay biết. Kẻ tấn công có thể **nghe lén** (eavesdropping) hoặc **sửa đổi** dữ liệu đang truyền.

**Các kỹ thuật MitM phổ biến:**

- **ARP Spoofing:** Kẻ tấn công gửi các gói ARP giả mạo trong mạng nội bộ, làm cho traffic của nạn nhân đi qua máy của kẻ tấn công.
- **Rogue Wi-Fi Access Point:** Tạo điểm truy cập WiFi giả mạo điểm đáng tin cậy (như WiFi khách sạn).
- **SSL Stripping:** Hạ cấp kết nối từ HTTPS xuống HTTP để đọc được dữ liệu plaintext.
- **DNS Spoofing:** Giả mạo phản hồi DNS để chuyển hướng nạn nhân đến máy chủ độc hại.

> **Ví dụ:** Tại một quán cà phê, kẻ tấn công tạo điểm truy cập WiFi tên "CafeGuest" — y hệt tên WiFi thật. Người dùng kết nối vào mạng giả này. Kẻ tấn công dùng SSL stripping để hạ cấp kết nối HTTPS xuống HTTP khi người dùng đăng nhập vào email, đọc được thông tin xác thực dưới dạng plaintext. Việc triển khai **HSTS (HTTP Strict Transport Security)** phía server có thể ngăn chặn kỹ thuật này.

---

## Domain Name System (DNS) Attack — Tấn công hệ thống tên miền

DNS là "danh bạ điện thoại" của internet: chuyển đổi tên miền (ví dụ: `google.com`) thành địa chỉ IP tương ứng (ví dụ: `142.250.80.46`) để máy tính có thể thiết lập kết nối. Khi DNS bị tấn công, người dùng có thể bị chuyển hướng đến các trang web giả mạo mà hoàn toàn không hay biết.

**Các loại tấn công DNS phổ biến:**

- **DNS Cache Poisoning:** Kẻ tấn công chèn bản ghi DNS giả mạo vào bộ nhớ đệm (cache) của DNS resolver, khiến người dùng bị chuyển hướng đến địa chỉ IP độc hại ngay cả khi nhập đúng tên miền.
- **DNS Hijacking:** Kẻ tấn công chiếm quyền kiểm soát cấu hình DNS (thông qua tấn công router, tấn công nhà đăng ký tên miền, hoặc malware), thay đổi bản ghi DNS để chuyển hướng toàn bộ traffic.
- **DNS Tunneling:** Dùng giao thức DNS để mã hóa và truyền dữ liệu ra ngoài mạng, thường nhằm lọc dữ liệu hoặc duy trì kênh C2 vượt qua tường lửa.
- **NXDOMAIN Attack:** Flood DNS server bằng các truy vấn cho các tên miền không tồn tại, làm cạn kiệt tài nguyên của server.

> **Ví dụ:** Nhóm **Roaming Mantis** xâm nhập vào các router không dây gia đình bằng cách khai thác mật khẩu mặc định hoặc lỗ hổng firmware, sau đó thay đổi cấu hình DNS của router. Kết quả là toàn bộ thiết bị trong mạng gia đình bị chuyển hướng đến trang web giả mạo phát tán malware **Wroba** — một banking trojan nhắm vào thiết bị Android. Các thiết bị bị nhiễm sau đó được dùng làm bot để lây lan đến các router khác.

---

## SQL Injection (SQLi) — Tấn công chèn SQL

SQL (Structured Query Language) là ngôn ngữ chuẩn để tương tác với cơ sở dữ liệu quan hệ. SQL injection xảy ra khi ứng dụng web nhúng trực tiếp dữ liệu đầu vào của người dùng vào câu truy vấn SQL mà không xác thực hoặc làm sạch (sanitize) đúng cách — cho phép kẻ tấn công can thiệp vào logic truy vấn.

SQL injection liên tục xuất hiện trong danh sách **OWASP Top 10** — danh sách 10 rủi ro bảo mật ứng dụng web nguy hiểm nhất.

**Ví dụ cơ chế tấn công:**

```sql
-- Truy vấn gốc của ứng dụng (dễ bị tấn công):
SELECT * FROM users WHERE username = '[INPUT]' AND password = '[INPUT]'

-- Kẻ tấn công nhập vào ô username:
' OR '1'='1' --

-- Câu truy vấn sau khi chèn:
SELECT * FROM users WHERE username = '' OR '1'='1' --' AND password = '...'
-- Điều kiện '1'='1' luôn đúng → bỏ qua xác thực hoàn toàn
```

**Các loại SQL injection:**

- **In-band SQLi (Classic):** Kết quả trả về trực tiếp trong response của ứng dụng.
- **Blind SQLi:** Ứng dụng không trả về dữ liệu trực tiếp, kẻ tấn công suy luận thông qua các câu hỏi đúng/sai hoặc thời gian phản hồi.
- **Out-of-band SQLi:** Kết quả được gửi qua kênh khác (DNS, HTTP request đến server bên ngoài).

**Phòng thủ — đặc biệt quan trọng với Backend Engineer:**

- **Prepared Statements / Parameterized Queries:** Tách biệt mã SQL và dữ liệu — phương pháp phòng thủ hiệu quả nhất.
- **Stored Procedures:** Truy vấn được định nghĩa trước, tham số hóa.
- **Input validation và escaping:** Lọc và làm sạch dữ liệu đầu vào.
- **ORM (Object-Relational Mapping):** Hầu hết ORM hiện đại mặc định dùng parameterized query.
- **Principle of Least Privilege:** Tài khoản database của ứng dụng chỉ có đúng quyền cần thiết, không có quyền DROP TABLE.

> **Ví dụ thực tế:** Năm 2015, hai thiếu niên sử dụng SQL injection để xâm nhập cơ sở dữ liệu của **TalkTalk** (nhà cung cấp viễn thông Anh Quốc), đánh cắp thông tin cá nhân và tài chính của khoảng 157.000 khách hàng. Vụ việc dẫn đến phạt tiền 400.000 GBP từ cơ quan quản lý và thiệt hại ước tính 60 triệu GBP — phần lớn do TalkTalk không áp dụng các biện pháp bảo mật cơ bản như parameterized query.



---

## Ứng dụng AI trong Tấn công Mạng

Kẻ tấn công ngày càng tích hợp AI vào quy trình tấn công để tăng tốc độ, quy mô, và độ tinh vi — trong khi giảm yêu cầu về kỹ năng kỹ thuật. Đây là thay đổi cơ bản trong bức tranh mối đe dọa, không chỉ là sự cải tiến tăng thêm.

### Task Automation — Tự động hóa tác vụ

AI có thể tự động hóa các tác vụ lặp lại theo quy mô không thể đạt được bằng phương pháp thủ công: tạo hàng triệu email phishing được cá nhân hóa, quét lỗ hổng trên hàng nghìn mục tiêu đồng thời, và tự động điều chỉnh payload tấn công dựa trên phản hồi của hệ thống mục tiêu.

### Detection Evasion — Né tránh phát hiện

AI hỗ trợ phát triển **polymorphic malware** và **metamorphic malware** — các biến thể tự thay đổi cấu trúc mã sau mỗi lần lây nhiễm để né tránh phát hiện dựa trên chữ ký (signature detection). Các mô hình học máy (ML) được huấn luyện để phân tích phản hồi của phần mềm bảo mật và điều chỉnh chiến thuật né tránh theo thời gian thực.

### Target Identification — Nhận diện mục tiêu

AI phân tích lượng lớn dữ liệu thu thập được từ nhiều nguồn (OSINT, dark web, dữ liệu rò rỉ từ các vụ breach trước đó) để xác định mục tiêu tối ưu: hệ thống nào đang chạy phần mềm lỗi thời, tổ chức nào có lịch sử bảo mật yếu, cá nhân nào có quyền truy cập cấp cao nhưng ít được bảo vệ.

### Social Engineering — Kỹ thuật xã hội tăng cường bởi AI

**Deepfake** là công nghệ AI tạo ra hình ảnh, video, hoặc âm thanh giả mạo người thật với độ chân thực cao. Kẻ tấn công dùng deepfake âm thanh để mạo danh giọng nói của CEO hoặc người thân trong các cuộc gọi lừa đảo — một dạng vishing nâng cao. Ví dụ thực tế: năm 2020, kẻ tấn công dùng deepfake giọng nói để lừa một giám đốc ngân hàng chuyển 35 triệu USD vào tài khoản giả.

LLM (Large Language Models) cũng được khai thác để tạo nội dung phishing không còn mắc các lỗi ngữ pháp — vốn là dấu hiệu nhận biết phổ biến của phishing truyền thống.

---

## Thông tin bổ sung

### 1. OWASP Top 10 — Tài nguyên không thể thiếu với Backend Engineer

OWASP (Open Web Application Security Project) công bố danh sách 10 rủi ro bảo mật ứng dụng web phổ biến nhất, cập nhật định kỳ. SQL injection thuộc nhóm **Injection** (A03:2021). Danh sách đầy đủ bao gồm: Broken Access Control, Cryptographic Failures, Injection, Insecure Design, Security Misconfiguration, Vulnerable and Outdated Components, Identification and Authentication Failures, Software and Data Integrity Failures, Security Logging and Monitoring Failures, Server-Side Request Forgery.

Website: [owasp.org/www-project-top-ten](https://owasp.org/www-project-top-ten/)

### 2. Phân biệt DoS và DDoS từ góc độ phòng thủ

| Tiêu chí        | DoS            | DDoS                                              |
| --------------- | -------------- | ------------------------------------------------- |
| Nguồn tấn công  | Một địa chỉ IP | Hàng nghìn đến hàng triệu IP                      |
| Phòng thủ       | Chặn IP nguồn  | Cần giải pháp chuyên biệt (CDN, scrubbing center) |
| Khó phát hiện   | Thấp           | Cao (traffic trộn lẫn với traffic hợp lệ)         |
| Yêu cầu kỹ năng | Thấp hơn       | Cao hơn (cần quản lý botnet)                      |

### 3. Backend Security Checklist liên quan đến các loại tấn công trong tài liệu

Với Backend Engineer, các biện pháp phòng thủ ưu tiên:

- **Chống SQL injection:** Luôn dùng parameterized query hoặc ORM. Không bao giờ nối chuỗi SQL trực tiếp với input người dùng.
- **Chống phishing trên API:** Implement rate limiting, CAPTCHA, và xác thực đa yếu tố (MFA) cho các endpoint nhạy cảm.
- **Chống MitM:** Enforce HTTPS với HSTS, triển khai certificate pinning cho ứng dụng mobile.
- **Chống DDoS:** Triển khai rate limiting tại API gateway, dùng CDN với tính năng DDoS protection.
- **Chống DNS attack:** Dùng DNSSEC để xác thực phản hồi DNS, giám sát thay đổi bản ghi DNS.
