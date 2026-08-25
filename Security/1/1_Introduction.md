# Giới thiệu tổng quan về An ninh mạng và Mô hình CIA

## Bối cảnh và tầm quan trọng

Chi phí trung bình toàn cầu của một vụ vi phạm dữ liệu đã lên tới **4,35 triệu USD**, mức cao nhất từ trước đến nay (theo báo cáo IBM Cost of a Data Breach 2022). Tội phạm mạng được dự báo gây thiệt hại **8 nghìn tỷ USD** cho toàn thế giới vào năm 2023 — nếu xét theo quy mô nền kinh tế quốc gia, con số này sẽ đưa tội phạm mạng trở thành nền kinh tế lớn thứ ba thế giới, chỉ sau Mỹ và Trung Quốc.

> Bản gốc trích dẫn số liệu mà không có nguồn tham chiếu, gây khó kiểm chứng. Đã bổ sung nguồn (IBM Cost of a Data Breach Report 2022) để tăng tính chính xác và khả năng tra cứu cho người đọc.

Ngày nay, công nghệ kỹ thuật số ảnh hưởng đến hầu hết mọi khía cạnh trong cuộc sống: ngân hàng, mua sắm, y tế, giáo dục và tương tác xã hội đều phụ thuộc sâu sắc vào hạ tầng số. Thế giới ngày càng kết nối hơn thông qua mạng lưới thiết bị rộng lớn được gọi là **Internet of Things (IoT)** — Internet vạn vật. Những thiết bị này tạo ra và truyền tải một lượng dữ liệu khổng lồ, làm thay đổi căn bản cách chúng ta giao tiếp, làm việc và sinh sống.

Tuy nhiên, cùng với sự tiện lợi đó là những rủi ro tương xứng. Dấu vết kỹ thuật số của mỗi người ngày càng mở rộng, và dữ liệu cá nhân lẫn dữ liệu nghề nghiệp nếu không được bảo vệ đúng cách sẽ trở thành mục tiêu của các cuộc tấn công. An ninh mạng chính là lĩnh vực giải quyết những thách thức này.

> *"The art of war teaches us to rely not on the likelihood of the enemy's not coming, but on our own readiness to receive him; not on the chance of his not attacking, but rather on the fact that we have made our position unassailable."*
>
> — Sun Tzu

---

## Cybersecurity — An ninh mạng là gì?

Khi hầu hết mọi người nghĩ về an ninh mạng, họ chỉ liên tưởng đến công nghệ: tường lửa, mã hóa, phần mềm diệt virus. Tuy nhiên, đó chỉ là một phần của bức tranh. Hãy xem xét hai tình huống sau:

- Một kẻ lừa đảo gửi email giả mạo nhân viên ngân hàng, yêu cầu nạn nhân cung cấp mã PIN. Đây có phải là mối lo ngại về an ninh mạng không?
- Một thám tử tư gọi điện thuyết phục nhân viên nội bộ in tài liệu mật và để ở phòng thư tín để người ngoài lấy đi. Đây có phải là mối lo ngại về an ninh mạng không?

Cả hai đều là: mối đe dọa an ninh mạng không chỉ đến từ phần mềm độc hại hay hacker kỹ thuật cao — chúng còn đến từ hành vi con người và các lỗ hổng vật lý. Vì vậy, an ninh mạng thực sự bao gồm ba thành phần:

---

### Ba thành phần của An ninh mạng

**Digital Security (Bảo mật kỹ thuật số)**

Bảo vệ dữ liệu và hệ thống khỏi các mối đe dọa số: phần mềm độc hại (malware), virus, ransomware, tấn công xâm nhập hệ thống và đánh cắp thông tin nhạy cảm. Các biện pháp điển hình bao gồm tường lửa (firewall), phần mềm mã hóa, hệ thống phát hiện xâm nhập (IDS/IPS).

**Human Security (Bảo mật con người)**

Bảo vệ dữ liệu khỏi các rủi ro phát sinh từ hành vi của con người — cả vô ý lẫn cố ý. Ví dụ: nhân viên vô tình nhấp vào liên kết lừa đảo (phishing), hoặc nội gián chủ động làm rò rỉ thông tin. Các biện pháp điển hình bao gồm đào tạo nhận thức bảo mật, chính sách mật khẩu mạnh, và xác thực đa yếu tố (MFA).

**Physical Security (Bảo mật vật lý)**

Bảo vệ các tài sản hữu hình hỗ trợ hạ tầng kỹ thuật số: máy trạm, phòng máy chủ, trung tâm dữ liệu. Mối đe dọa vật lý bao gồm trộm cắp, phá hoại thiết bị, hoặc thiên tai. Các biện pháp điển hình bao gồm hệ thống camera giám sát (CCTV), kiểm soát truy cập vật lý (thẻ từ, sinh trắc học), và kế hoạch phục hồi sau thảm họa (DRP).

**Tóm lại:** An ninh mạng là việc **bảo vệ và phục hồi** dữ liệu, mạng lưới, thiết bị và chương trình khỏi các mối đe dọa — trong đó mối đe dọa có thể mang thành phần kỹ thuật số, con người hoặc vật lý, và thường là kết hợp của nhiều thành phần cùng lúc.

---

## Mô hình CIA Triad

An ninh mạng hiệu quả cần đạt được ba mục tiêu cốt lõi, được tổng hợp trong mô hình **CIA Triad** — một trong những khung tư duy nền tảng và được sử dụng rộng rãi nhất trong lĩnh vực an ninh thông tin:

- **C**onfidentiality — Bảo mật thông tin
- **I**ntegrity — Toàn vẹn dữ liệu
- **A**vailability — Khả dụng

---

### Confidentiality (Bảo mật thông tin)

Confidentiality đảm bảo rằng dữ liệu chỉ được truy cập hoặc tiết lộ bởi những người được ủy quyền hợp lệ.

Ví dụ trong thực tế:

- Các công ty phần mềm hạn chế quyền truy cập vào mã nguồn chỉ cho những kỹ sư thực sự cần thiết, nhằm bảo vệ lợi thế cạnh tranh và ngăn rò rỉ sở hữu trí tuệ.
- Nhà cung cấp dịch vụ y tế phải đảm bảo hồ sơ bệnh nhân — chẩn đoán, đơn thuốc, kết quả xét nghiệm — chỉ được truy cập bởi bệnh nhân, bác sĩ điều trị và nhân viên y tế được ủy quyền.

Trên thực tế, triển khai Confidentiality đòi hỏi xác định **đúng người, đúng quyền, đúng thời điểm, đúng phương thức** — thông qua kiểm soát truy cập (access control), mã hóa dữ liệu, và nguyên tắc quyền tối thiểu (Principle of Least Privilege).

---

### Integrity (Toàn vẹn dữ liệu)

Integrity đảm bảo dữ liệu chính xác, đáng tin cậy và không bị sửa đổi hoặc phá hủy trái phép.

Ví dụ minh họa: Bạn thanh toán 10 USD cho một chiếc pizza. Nếu một lỗi nào đó — dù do hệ thống hay do con người nhập sai — khiến số tiền giao dịch thay đổi thành 10.000 USD, tính toàn vẹn của giao dịch đã bị vi phạm. Đáng chú ý là vi phạm toàn vẹn có thể xảy ra **cả cố ý lẫn vô ý**, và nguyên nhân có thể là kỹ thuật (lỗi phần mềm) hoặc con người (nhập liệu sai).

![](https://cpcontents.adobe.com/fr/dynamic-protected/2831aaf30b994e7b848c780a9efb7871/protected/account/2135/resources/7868746/7868746/content/scormcontent/assets/Cybersecurity_AuditTrail%20%281%29.png)

Để bảo toàn Integrity, cần kết hợp: kiểm soát truy cập (ngăn người không được phép chỉnh sửa), checksum và hàm băm (phát hiện thay đổi dữ liệu), audit trail (ghi lại mọi thao tác thay đổi để truy vết), và sao lưu dữ liệu định kỳ.

> Lưu ý: Integrity và Confidentiality có vùng chồng lấp — một người không được phép truy cập dữ liệu đương nhiên cũng không nên được phép chỉnh sửa nó. Tuy nhiên, hai mục tiêu này giải quyết các mối lo ngại khác nhau: Confidentiality chống lộ lọt thông tin, còn Integrity chống sửa đổi trái phép.

---

### Availability (Khả dụng)

Availability đảm bảo dữ liệu và hệ thống có thể được truy cập và sử dụng một cách kịp thời và đáng tin cậy bởi những người có quyền.

![](https://cpcontents.adobe.com/fr/dynamic-protected/2831aaf30b994e7b848c780a9efb7871/protected/account/2135/resources/7868746/7868746/content/scormcontent/assets/Cybersecurity_Availability%20%281%29.png)

Ví dụ: Người dùng ngân hàng trực tuyến kỳ vọng truy cập tài khoản 24/7. Để đáp ứng điều đó, ngân hàng phải duy trì hạ tầng dự phòng, cân bằng tải (load balancing), và kế hoạch phục hồi sau sự cố (disaster recovery).

Tuy nhiên, "kịp thời" không đồng nghĩa với "ngay lập tức trong mọi trường hợp". Ví dụ: yêu cầu bản sao học bạ có thể mất vài ngày xử lý, và bản sao điện tử có thể bị giới hạn thời gian truy cập. Dữ liệu vẫn được coi là "có sẵn" nếu được cung cấp trong khoảng thời gian hợp lý và nhất quán với cam kết dịch vụ.

Các mối đe dọa chính đối với Availability bao gồm: tấn công từ chối dịch vụ (DoS/DDoS), lỗi phần cứng, thảm họa tự nhiên, và cấu hình sai hệ thống.

---

## Ứng dụng thực tế của CIA Triad

Mô hình CIA Triad cung cấp khung tư duy để đánh giá bất kỳ quyết định nào liên quan đến bảo mật thông tin. Tùy theo đặc thù hoạt động, các tổ chức sẽ ưu tiên khác nhau:

| Loại tổ chức | Mục tiêu ưu tiên | Lý do |
|---|---|---|
| Bệnh viện, y tế | Availability > Confidentiality | Dữ liệu bệnh nhân phải luôn truy cập được trong cấp cứu |
| Ngân hàng, tài chính | Integrity > Confidentiality | Sai lệch giao dịch gây thiệt hại trực tiếp |
| Cơ quan chính phủ, quốc phòng | Confidentiality > Availability | Rò rỉ thông tin mật gây hậu quả nghiêm trọng |
| E-commerce | Availability > Integrity | Hệ thống ngừng hoạt động = mất doanh thu tức thì |

> Lưu ý quan trọng: Trong thực tế, ba mục tiêu của CIA Triad đôi khi **xung đột với nhau**. Tăng cường Confidentiality (thêm lớp xác thực, mã hóa) có thể làm giảm Availability (người dùng mất nhiều thao tác hơn để truy cập). Nhiệm vụ của chuyên gia bảo mật là tìm điểm cân bằng phù hợp với nhu cầu tổ chức.

---

## Thông tin bổ sung

### 1. Mô hình mở rộng: CIA Triad và các biến thể

CIA Triad là điểm khởi đầu, nhưng một số tổ chức và framework mở rộng thêm các yếu tố khác:

- **Non-repudiation (Chống chối bỏ):** Đảm bảo người thực hiện hành động không thể phủ nhận đã làm điều đó. Ví dụ: chữ ký số trong hợp đồng điện tử.
- **Authentication (Xác thực):** Đảm bảo danh tính của người dùng trước khi cấp quyền truy cập.
- **Authorization (Ủy quyền):** Xác định người dùng được phép làm gì sau khi đã xác thực.

Mô hình **Parkerian Hexad** (1998) mở rộng CIA thêm ba yếu tố: Possession/Control (kiểm soát vật lý dữ liệu), Authenticity (tính xác thực nguồn gốc), và Utility (tính hữu dụng).

### 2. Social Engineering — Mối đe dọa thường bị đánh giá thấp

Hai ví dụ trong tài liệu (email giả mạo ngân hàng, cuộc gọi thám tử tư) đều là **Social Engineering** — kỹ thuật tấn công khai thác yếu tố tâm lý con người thay vì lỗ hổng kỹ thuật. Đây là một trong những vector tấn công hiệu quả và phổ biến nhất hiện nay, bao gồm:

- **Phishing:** Email/tin nhắn giả mạo để đánh cắp thông tin.
- **Vishing:** Lừa đảo qua điện thoại (voice phishing).
- **Pretexting:** Dựng kịch bản giả để thuyết phục nạn nhân cung cấp thông tin hoặc thực hiện hành động.
- **Tailgating/Piggybacking:** Theo vào khu vực hạn chế bằng cách lợi dụng người có thẩm quyền mở cửa.

### 3. Các framework bảo mật thực tiễn tham chiếu CIA Triad

Người học nên biết các framework sau đây đều xây dựng trên nền tảng CIA Triad:

- **NIST Cybersecurity Framework (CSF):** Framework phổ biến nhất tại Mỹ, gồm 5 chức năng: Identify, Protect, Detect, Respond, Recover.
- **ISO/IEC 27001:** Tiêu chuẩn quốc tế về quản lý an ninh thông tin (ISMS).
- **OWASP Top 10:** Danh sách 10 rủi ro bảo mật web phổ biến nhất, áp dụng trực tiếp ba mục tiêu CIA vào ứng dụng web.
