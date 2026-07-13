# Ứng phó Sự cố An ninh mạng (Respond to Attacks)

## Giới thiệu

Ngay cả hệ thống phòng thủ và phát hiện tốt nhất cũng không thể ngăn chặn tuyệt đối mọi cuộc tấn công. Mọi tổ chức, không sớm thì muộn, đều phải đối mặt với một sự cố an ninh mạng thực sự. Một phần thiết yếu của chiến lược an ninh là **chuẩn bị sẵn kế hoạch ứng phó** — xác định rõ cần làm gì, ai chịu trách nhiệm, và theo trình tự nào khi sự cố xảy ra.

Module này giới thiệu framework ứng phó sự cố (Incident Response) gồm sáu giai đoạn, được minh họa qua một tình huống thực tế.

---

## Ứng phó Sự cố (Incident Response) — Sáu giai đoạn

> **. Đây là mô hình dựa trên **SANS Incident Response process** (đôi khi gọi tắt là **PICERL**: Preparation, Identification, Containment, Eradication, Recovery, Lessons Learned) — một trong những framework ứng phó sự cố được sử dụng rộng rãi nhất trong ngành, cùng với NIST SP 800-61. Biết tên chính thức giúp người học tra cứu thêm tài liệu chuyên sâu khi cần.

### Giai đoạn 1: Chuẩn bị (Preparation)

Đây là giai đoạn diễn ra **trước khi** sự cố xảy ra — và là giai đoạn quan trọng nhất vì chất lượng chuẩn bị quyết định trực tiếp tốc độ và hiệu quả của toàn bộ quá trình ứng phó sau này.

Các hoạt động điển hình:

- Xây dựng và tài liệu hóa kế hoạch ứng phó sự cố (Incident Response Plan — IRP).
- Chuẩn bị nguồn lực: công cụ, quyền truy cập khẩn cấp, kênh liên lạc dự phòng.
- Thử nghiệm quy trình thông qua diễn tập (tabletop exercise) hoặc mô phỏng tấn công thực tế.
- Phân công vai trò và trách nhiệm rõ ràng cho từng thành viên đội ứng phó (Incident Response Team — IRT).
- Thiết lập sẵn công cụ giám sát, logging, và kênh cảnh báo (đã đề cập trong module Phát hiện Tấn công).

### Giai đoạn 2: Nhận diện (Identification)

Nhóm bảo mật phát hiện một sự kiện bất thường (thường qua SIEM, logging, hoặc network monitoring) và cần xác định: **đây có thực sự là sự cố bảo mật hay chỉ là báo động giả (false positive)?**

Khi xác nhận đây là sự cố thực sự, nhóm chính thức **khởi động quy trình ứng phó sự cố** và bắt đầu điều tra sâu hơn: phạm vi ảnh hưởng, mức độ nghiêm trọng, và hệ thống nào bị tác động.

> **Phân loại mức độ nghiêm trọng (Severity Classification):** Hầu hết tổ chức trưởng thành sẽ gán một mức độ nghiêm trọng (ví dụ: SEV1 – Critical, SEV2 – High, SEV3 – Medium, SEV4 – Low) ngay khi xác nhận sự cố, vì mức độ này quyết định quy trình leo thang (escalation), ai cần được thông báo ngay lập tức, và SLA thời gian phản hồi tương ứng.

### Giai đoạn 3: Ngăn chặn (Containment)

Mục tiêu của giai đoạn này là **ngăn tình hình trở nên tồi tệ hơn** — không nhất thiết phải loại bỏ hoàn toàn mối đe dọa ngay lập tức, mà là giới hạn phạm vi thiệt hại có thể lan rộng.

Các biện pháp điển hình: phân đoạn mạng (network segmentation), tắt các tuyến truy cập hoặc hệ thống cụ thể, cô lập thiết bị bị nghi ngờ.

> **Short-term vs. Long-term Containment:** Trong thực hành chuyên nghiệp, giai đoạn này thường chia thành hai bước:
> 
> - **Ngăn chặn ngắn hạn (short-term containment):** Hành động tức thời để chặn đứng thiệt hại đang diễn ra — ví dụ rút cáp mạng của máy chủ bị xâm nhập. Có thể làm gián đoạn dịch vụ nhưng cần thiết để ngăn lan rộng.
> - **Ngăn chặn dài hạn (long-term containment):** Giải pháp tạm thời nhưng ổn định hơn, cho phép hệ thống tiếp tục hoạt động ở mức tối thiểu trong khi chờ diệt trừ hoàn toàn — ví dụ chuyển traffic sang hệ thống dự phòng sạch trong khi điều tra hệ thống gốc.
> 
> Phân biệt này giúp đội ứng phó cân bằng giữa tốc độ phản ứng và duy trì hoạt động kinh doanh.

### Giai đoạn 4: Diệt trừ (Eradication)

Nhóm bảo mật loại bỏ **hoàn toàn** sự hiện diện và tác động của kẻ tấn công hoặc phần mềm độc hại khỏi môi trường — ví dụ: xóa sạch và khôi phục thiết bị về trạng thái sạch đã biết (known-good state).

**Lưu ý quan trọng:** Diệt trừ không triệt để là nguyên nhân phổ biến khiến phần mềm độc hại **tái xuất hiện** sau khi tổ chức tưởng rằng đã xử lý xong. Vì vậy, nỗ lực diệt trừ phải được thực hiện kỹ lưỡng và xác minh đầy đủ trước khi chuyển sang giai đoạn phục hồi.

> **Xác định Root Cause trước khi diệt trừ:** Nếu chỉ xóa malware đã phát hiện mà không tìm ra cách kẻ tấn công xâm nhập ban đầu (initial access vector) hoặc các backdoor/persistence mechanism khác mà chúng đã cài đặt, tổ chức có nguy cơ bị tấn công lại ngay lập tức qua cùng lỗ hổng hoặc qua backdoor còn sót lại. Đây là lý do tại sao quá trình diệt trừ chuyên nghiệp luôn đi kèm với digital forensics (điều tra số) để trả lời đầy đủ: kẻ tấn công vào bằng cách nào, đã làm gì, và để lại gì.

### Giai đoạn 5: Phục hồi (Recovery)

Tổ chức đưa hệ thống trở lại hoạt động bình thường: loại bỏ các giải pháp tạm thời (ví dụ: containment tạm thời ở giai đoạn 3), khôi phục dịch vụ theo thứ tự ưu tiên.

**Nguyên tắc khôi phục theo mức độ ưu tiên:** Không phải mọi hệ thống đều cần được khôi phục cùng lúc. Tổ chức cần xác định trước (thường trong giai đoạn Chuẩn bị) hệ thống nào là **critical** (cần khôi phục đầu tiên) dựa trên tác động kinh doanh — khái niệm này liên quan đến **RTO (Recovery Time Objective)** và **RPO (Recovery Point Objective)** trong kế hoạch Business Continuity.



### Giai đoạn 6: Rút ra Bài học Kinh nghiệm (Lessons Learned)

Sau khi hoạt động trở lại bình thường, đội ứng phó và các bên liên quan cùng nhìn lại: nguyên nhân gốc rễ của sự cố là gì, quy trình ứng phó đã hiệu quả đến đâu, và cần thay đổi gì để phòng tránh sự cố tương tự.

> Đây là quan sát chính xác và đáng giữ lại. Bổ sung thêm: kết quả của giai đoạn này thường được ghi thành văn bản chính thức gọi là **Post-Incident Report (PIR)** hoặc **Post-Mortem Report** — tài liệu này không chỉ để lưu trữ mà cần được **hành động hóa** thành các action item cụ thể, có người phụ trách và deadline rõ ràng, nếu không giai đoạn này chỉ là hình thức.

> **Lưu ý về tính linh hoạt của framework:** Một số loại sự cố có thể yêu cầu mở rộng một giai đoạn cụ thể — ví dụ, sự cố rò rỉ dữ liệu do mất thiết bị vật lý có thể có giai đoạn Diệt trừ (Eradication) ngắn gọn (không có malware cần loại bỏ), nhưng giai đoạn Phục hồi lại kéo dài và phức tạp hơn nhiều do liên quan đến nghĩa vụ pháp lý thông báo cho khách hàng, cơ quan quản lý, và các bên liên quan khác.

---

## Tình huống thực hành: HealthyOnline

Hãy xem xét tình huống của **HealthyOnline** — một nhà cung cấp dịch vụ y tế trực tuyến — khi đối mặt với sự cố rò rỉ dữ liệu.

### Giai đoạn 1 — Chuẩn bị

Trước khi sự cố xảy ra, HealthyOnline đã:

- Thường xuyên tiến hành kiểm thử xâm nhập (penetration testing) để chủ động phát hiện lỗ hổng.
- Chuẩn bị sẵn kế hoạch ứng phó sự cố chi tiết.
- Sao lưu toàn bộ dữ liệu quan trọng.
- Đào tạo tất cả thành viên đội ứng phó về vai trò và trách nhiệm cụ thể của họ.

### Giai đoạn 2 — Nhận diện

Nhóm bảo mật phát hiện **lượng dữ liệu bất thường** được truyền từ máy chủ đến một địa chỉ IP không xác định — dấu hiệu điển hình của **data exfiltration** (đánh cắp dữ liệu ra ngoài). HealthyOnline nhanh chóng kích hoạt quy trình ứng phó sự cố và xác nhận khả năng cao đây là một vụ vi phạm dữ liệu thực sự.

### Giai đoạn 3 — Ngăn chặn

Nhóm bảo mật thực hiện ngay các biện pháp ngăn chặn:

- **Cách ly máy chủ bị xâm nhập** khỏi mạng — ngăn kẻ tấn công tiếp tục truyền dữ liệu ra ngoài hoặc lan rộng sang hệ thống khác.
- **Đăng xuất toàn bộ phiên người dùng đang hoạt động** — giảm thiểu khả năng kẻ tấn công đang sử dụng phiên đăng nhập bị đánh cắp.

### Giai đoạn 4 — Diệt trừ

Nhóm bảo mật:

- Xác định phần mềm độc hại là nguyên nhân gây xâm nhập và loại bỏ hoàn toàn khỏi hệ thống.
- Thay thế máy chủ bị xâm nhập bằng **bản sao lưu sạch (clean backup)** đã biết là an toàn.
- Tăng cường quy tắc tường lửa liên quan để ngăn chặn cùng phương thức tấn công tái diễn.

### Giai đoạn 5 — Phục hồi

Sau khi xác nhận mối đe dọa đã được loại bỏ hoàn toàn:

- Kết nối lại máy chủ với mạng.
- Khôi phục hoạt động **theo thứ tự ưu tiên**, bắt đầu từ các dịch vụ quan trọng nhất.
- **Đặt lại mật khẩu cho toàn bộ nhân viên** — biện pháp phòng ngừa trong trường hợp thông tin xác thực đã bị đánh cắp.
- **Thông báo cho khách hàng** về sự cố và khuyến nghị họ cập nhật mật khẩu.

> Với một tổ chức y tế như HealthyOnline, việc "thông báo cho khách hàng" không chỉ là thực hành tốt mà thường là **nghĩa vụ pháp lý bắt buộc**. Tại Mỹ, luật HIPAA (Health Insurance Portability and Accountability Act) yêu cầu thông báo vi phạm dữ liệu y tế trong thời hạn cụ thể (thường trong vòng 60 ngày). Tại EU, GDPR yêu cầu thông báo cho cơ quan giám sát trong vòng 72 giờ kể từ khi phát hiện vi phạm. Việc bỏ sót nghĩa vụ pháp lý và thời hạn thông báo trong tình huống này là một khoảng trống quan trọng — tổ chức xử lý dữ liệu y tế/nhạy cảm cần tích hợp bước thông báo pháp lý này ngay trong giai đoạn Phục hồi, không phải là hành động tùy chọn.

### Giai đoạn 6 — Rút ra Bài học Kinh nghiệm

Sau khi hoạt động ổn định trở lại, HealthyOnline tổng hợp phát hiện thành báo cáo chính thức. Báo cáo xác định **nguyên nhân gốc rễ**: bản vá phần mềm được khuyến nghị đã không được áp dụng kịp thời, tạo ra lỗ hổng mà kẻ tấn công khai thác.

**Hành động khắc phục cụ thể:** Tổ chức quyết định **tự động hóa quy trình vá lỗi (automated patch management)** để loại bỏ yếu tố con người — nguyên nhân trực tiếp dẫn đến sự cố lần này.

### Kết luận tình huống

HealthyOnline đã hoàn thành đầy đủ sáu giai đoạn của framework ứng phó sự cố một cách nhanh chóng và có hệ thống. Quan trọng hơn, bằng cách chuyển bài học kinh nghiệm thành **hành động khắc phục cụ thể** (tự động hóa patch management), tổ chức thực sự trở nên kiên cường hơn — thay vì chỉ dừng lại ở việc viết báo cáo.

---

## Bổ sung kiến thức

### 1. So sánh PICERL (SANS) với NIST SP 800-61

Tài liệu này trình bày framework sáu giai đoạn (PICERL), nhưng một framework phổ biến khác cũng đáng biết là **NIST SP 800-61** — có bốn giai đoạn:

1. Preparation (Chuẩn bị)
2. Detection & Analysis (Phát hiện & Phân tích) — gộp chung Identification
3. Containment, Eradication & Recovery (Ngăn chặn, Diệt trừ & Phục hồi) — gộp ba giai đoạn thành một nhóm lặp lại theo chu kỳ
4. Post-Incident Activity (Hoạt động sau sự cố) — tương đương Lessons Learned

Về bản chất, hai framework tương đương nhau — khác biệt chủ yếu ở cách phân chia và nhóm giai đoạn. NIST nhấn mạnh rằng Containment, Eradication và Recovery thường **lặp lại theo chu kỳ** (không phải tuyến tính hoàn toàn) trong các sự cố phức tạp — ví dụ có thể cần quay lại containment nếu phát hiện thêm hệ thống bị xâm nhập trong quá trình diệt trừ.

### 2. Vai trò của Incident Response Plan (IRP) và Runbook

Một kế hoạch ứng phó sự cố hiệu quả trong thực tế thường bao gồm:

- **IRP (Incident Response Plan):** Tài liệu cấp cao mô tả chính sách, vai trò, quy trình leo thang tổng thể.
- **Runbook/Playbook:** Tài liệu chi tiết, từng bước cụ thể cho **từng loại sự cố** (ví dụ: runbook riêng cho ransomware, runbook riêng cho data breach, runbook riêng cho DDoS). Runbook giúp đội ứng phó hành động nhanh và nhất quán ngay cả dưới áp lực cao, không cần "nghĩ lại từ đầu" trong lúc khủng hoảng.

Với vai trò Backend Engineer, việc tham gia xây dựng runbook cho các sự cố liên quan đến hệ thống mình phụ trách (ví dụ: quy trình rollback khi phát hiện lỗi bảo mật trong một service cụ thể) là đóng góp thực tế và giá trị.

### 3. Vai trò của Backend Engineer trong Incident Response

Mặc dù Incident Response thường được xem là trách nhiệm của đội bảo mật/SOC, backend engineer đóng vai trò quan trọng và thường được huy động trong nhiều giai đoạn:

- **Giai đoạn Nhận diện:** Cung cấp context kỹ thuật về hệ thống — kiến trúc, luồng dữ liệu, dependency — giúp đội bảo mật đánh giá đúng phạm vi ảnh hưởng nhanh hơn.
- **Giai đoạn Ngăn chặn:** Thực hiện các thay đổi kỹ thuật khẩn cấp — vô hiệu hóa endpoint, revoke API key, chặn traffic ở tầng ứng dụng.
- **Giai đoạn Diệt trừ và Phục hồi:** Triển khai bản vá, rollback service về phiên bản sạch, khôi phục từ backup.
- **Giai đoạn Bài học Kinh nghiệm:** Đóng góp góc nhìn kỹ thuật về nguyên nhân gốc rễ và đề xuất cải thiện quy trình CI/CD, code review, hoặc kiến trúc hệ thống để tránh lặp lại lỗ hổng tương tự.

Vì lý do này, hiểu biết cơ bản về Incident Response là kỹ năng có giá trị với bất kỳ backend engineer nào, không chỉ riêng đội bảo mật.
