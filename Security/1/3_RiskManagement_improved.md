# Module 3: Quản lý Rủi ro trong An ninh mạng

## Giới thiệu

**Rủi ro (Risk)** là khả năng xảy ra một sự kiện không mong muốn với hậu quả tiêu cực. Rủi ro là một phần không thể tránh khỏi trong mọi hoạt động kinh doanh và cuộc sống hàng ngày.

Quản lý rủi ro là cốt lõi của nhiều ngành công nghiệp. Điển hình là ngành bảo hiểm: một công ty bảo hiểm đánh giá rủi ro liên quan đến việc bảo hiểm cho một cá nhân hoặc tổ chức bằng cách xác định xác suất xảy ra sự cố và khả năng yêu cầu bồi thường. Từ đó, họ đặt mức phí bảo hiểm phù hợp — về bản chất là chuyển giao rủi ro tài chính từ khách hàng sang công ty bảo hiểm để đổi lấy các khoản phí định kỳ.

Trong an ninh mạng, quản lý rủi ro giúp tổ chức xác định và ưu tiên bảo vệ tài sản thông tin một cách có hệ thống, thay vì phản ứng ngẫu nhiên với từng sự cố.

---

## Định lượng rủi ro (Risk Valuation)

Không phải mọi rủi ro đều quan trọng như nhau. Để phân bổ nguồn lực bảo mật hợp lý, cần định lượng mức độ nghiêm trọng của từng rủi ro. Công thức cơ bản:

![](https://images.postaffiliatepro.com/images/processed/blog/0x968b32cafbce83b5.webp)

> **Risk Value = Consequence × Likelihood**
>
> (Giá trị rủi ro = Hậu quả × Xác suất xảy ra)

- **Hậu quả (Consequence):** Mức độ tác động và thiệt hại nếu rủi ro xảy ra.
- **Xác suất (Likelihood):** Tần suất hoặc khả năng rủi ro đó xảy ra trong thực tế.

Khi cả hai yếu tố đều cao, giá trị rủi ro sẽ lớn và cần được ưu tiên giảm thiểu ngay. Khi một trong hai thấp, rủi ro có thể chờ xử lý sau.

### Ví dụ: Tính giá trị rủi ro thủng lốp xe

Giả sử cứ 10 chiếc xe thì có 1 chiếc bị thủng lốp trong một năm, và hậu quả là người lái mất một ngày làm việc.

- **Hậu quả:** Trung bình (mất một ngày làm việc, gây khó chịu nhưng không nghiêm trọng).
- **Xác suất:** Thấp (chỉ 1 trên 10 xe mỗi năm, tương đương 10%).
- **Giá trị rủi ro:** Thấp — không cần phản ứng khẩn cấp.

---

## Công thức mở rộng: Tính xác suất trong An ninh mạng

Trong lĩnh vực an ninh mạng, xác suất bị tấn công khó đo lường trực tiếp hơn các lĩnh vực khác, do công nghệ liên tục thay đổi và sự can thiệp của kẻ tấn công bên ngoài là yếu tố khó dự đoán. Vì vậy, xác suất thường được ước tính qua ba thành phần:

> **Likelihood = Adversary Capability × Adversary Motivation × Vulnerability Severity**
>
> (Xác suất = Khả năng của kẻ tấn công × Động cơ của kẻ tấn công × Mức độ nghiêm trọng của lỗ hổng)

### Adversary (Kẻ thù / Tác nhân đe dọa)

Kẻ thù (adversary) hay **tác nhân đe dọa (threat actor)** là thuật ngữ chung chỉ bất kỳ thực thể nào có ý định xâm phạm hệ thống thông tin — bao gồm cá nhân, nhóm tội phạm, tổ chức tin tặc được nhà nước bảo trợ, hoặc nội gián trong tổ chức.

> Bản gốc chỉ dùng thuật ngữ "kẻ thù" mà không đề cập đến "threat actor" — thuật ngữ kỹ thuật chuẩn trong ngành an ninh mạng. Đã bổ sung để người đọc làm quen với ngôn ngữ chuyên ngành.

![](https://fortune.com/img-assets/wp-content/uploads/2025/02/GettyImages-1923603027-e1739738963311.jpg?format=webp&w=1000&q=100)

### Adversary Capability (Khả năng của kẻ tấn công)

Bao gồm nguồn lực tài chính, công cụ tấn công, khả năng tiếp cận công nghệ hiện đại, và trình độ kỹ thuật. Kẻ tấn công có khả năng cao có thể phát triển công cụ riêng (zero-day exploit), trong khi kẻ tấn công có khả năng thấp thường chỉ sử dụng công cụ sẵn có (script kiddie).

### Adversary Motivation (Động cơ của kẻ tấn công)

Lý do thúc đẩy kẻ tấn công thực hiện hành vi xâm phạm. Động cơ phổ biến bao gồm:

- **Tài chính:** Đánh cắp dữ liệu để bán, tống tiền ransomware, gian lận tài chính.
- **Gián điệp:** Thu thập thông tin tình báo, sở hữu trí tuệ của đối thủ.
- **Chính trị / Hacktivism:** Gây rối hệ thống của tổ chức hoặc chính phủ vì lý do tư tưởng.
- **Trả thù hoặc phá hoại:** Nhân viên cũ bất mãn.
- **Thách thức kỹ thuật:** Tấn công vì muốn chứng minh khả năng.

> Bản gốc chỉ liệt kê "lợi ích tài chính, thông tin, ảnh hưởng chính trị, hoặc cảm giác hồi hộp" mà không đề cập đến ransomware, hacktivism hay insider threat — các vector tấn công quan trọng trong thực tế hiện đại. Đã bổ sung đầy đủ hơn.

### Vulnerability Severity (Mức độ nghiêm trọng của lỗ hổng)

**Lỗ hổng bảo mật (vulnerability)** là điểm yếu trong hệ thống có thể bị khai thác để xâm phạm tính bảo mật, toàn vẹn hoặc khả dụng. Lỗ hổng có thể nằm trong phần mềm, cấu hình hệ thống, quy trình vận hành, hoặc hành vi con người.

Hệ thống tiêu chuẩn để đánh giá mức độ nghiêm trọng của lỗ hổng là **CVSS (Common Vulnerability Scoring System)** — cho điểm từ 0.0 đến 10.0, trong đó 9.0–10.0 là mức Critical.

> Bản gốc không đề cập đến CVSS — một hệ thống đánh giá lỗ hổng được sử dụng rộng rãi trong ngành. Đây là kiến thức quan trọng mà người học bảo mật cần biết để thực hành đánh giá rủi ro đúng cách.

### Ví dụ ứng dụng công thức xác suất

Một nhóm tin tặc mới thành lập nhắm mục tiêu vào ngân hàng để đánh cắp thông tin đăng nhập của người dùng:

| Yếu tố | Đánh giá | Lý do |
|---|---|---|
| Khả năng của kẻ tấn công | Thấp | Nhóm mới, thiếu công cụ chuyên biệt và kinh nghiệm |
| Động cơ | Cao | Lợi ích tài chính rõ ràng, có thể thực hiện nhiều cuộc tấn công liên tiếp |
| Mức độ nghiêm trọng lỗ hổng | Cao | Lỗ hổng đã được công bố trực tuyến, dễ sao chép tấn công |
| **Xác suất tổng thể** | **Trung bình-Cao** | Động cơ mạnh và lỗ hổng dễ khai thác bù đắp cho khả năng hạn chế |

**Lưu ý về phân tích định tính vs. định lượng:** Trong thực tế, dữ liệu thống kê chính xác về tần suất tấn công và xác suất thường rất khó thu thập. Vì vậy, phân tích rủi ro thường dùng thang định tính (Thấp / Trung bình / Cao) thay vì con số cụ thể. Đây là bình thường — điều quan trọng là tính nhất quán trong phương pháp đánh giá, không phải độ chính xác tuyệt đối.

---

## Thực hành: Đánh giá rủi ro cho Phòng khám tư nhân

Phòng khám bác sĩ tư nhân có hồ sơ hệ thống như sau:

1. Hệ điều hành lỗi thời (không còn được nhà cung cấp hỗ trợ)
2. Xác thực một yếu tố (single-factor authentication) cho tài khoản email
3. Phần mềm chống phần mềm độc hại không được cập nhật thường xuyên
4. Các tệp quan trọng không được sao lưu
5. Phần cứng lỗi thời

### Phân tích từng rủi ro

**Rủi ro 1: Hệ điều hành lỗi thời**

Máy tính của phòng khám sử dụng hệ điều hành đã hết vòng đời hỗ trợ (End of Life — EOL). Nhà cung cấp không còn phát hành bản vá bảo mật cho phiên bản này.

*Câu hỏi đánh giá:* Hệ điều hành lỗi thời này có thể gây ra những lỗ hổng bảo mật nào, và hậu quả nếu bị khai thác là gì?

*Phân tích:* Các hệ điều hành EOL tích lũy nhiều lỗ hổng đã biết mà không bao giờ được vá. Kẻ tấn công có thể dễ dàng tìm kiếm công khai các lỗ hổng này (CVE database) và khai thác mà không cần kỹ năng cao. Giá trị rủi ro: **Cao**.

---

**Rủi ro 2: Xác thực một yếu tố cho tài khoản email**

Hệ thống email chỉ yêu cầu mật khẩu để đăng nhập. Tiêu chuẩn ngành hiện tại là **xác thực đa yếu tố (MFA — Multi-Factor Authentication)**, yêu cầu ít nhất hai loại thông tin xác thực khác nhau (ví dụ: mật khẩu + mã OTP từ ứng dụng xác thực).

*Câu hỏi đánh giá:* Xác thực một yếu tố tạo ra lỗ hổng gì, và kẻ tấn công có thể khai thác chúng như thế nào?

*Phân tích:* Mật khẩu đơn thuần dễ bị đánh cắp qua phishing, credential stuffing (thử mật khẩu bị rò rỉ từ vụ vi phạm khác), hoặc brute force. Với MFA, dù kẻ tấn công có mật khẩu cũng không đăng nhập được nếu không có yếu tố xác thực thứ hai. Lỗ hổng của xác thực một yếu tố càng dễ khai thác thì xác suất tấn công thành công càng cao, dẫn đến giá trị rủi ro lớn hơn.

---

**Rủi ro 3: Phần mềm chống phần mềm độc hại không được cập nhật thường xuyên**

Phần mềm antimalware không được cấu hình để tự động cập nhật. Người quản trị hệ thống cập nhật thủ công không đều đặn.

*Câu hỏi đánh giá:* Xác suất hệ thống bị phần mềm độc hại tấn công thành công trong điều kiện này là bao nhiêu?

*Phân tích:* Phần mềm chống malware hoạt động dựa trên cơ sở dữ liệu chữ ký (signature database). Nếu không được cập nhật, nó mù quáng với các biến thể malware mới. Trong thực tế, hàng trăm nghìn biến thể malware mới xuất hiện mỗi ngày. Cơ sở dữ liệu chữ ký lỗi thời trực tiếp làm tăng xác suất tấn công thành công — do đó làm tăng giá trị rủi ro.

---

**Rủi ro 4: Các tệp quan trọng không được sao lưu**

Phòng khám không có quy trình sao lưu định kỳ cho dữ liệu quan trọng.

*Câu hỏi đánh giá:* Việc mất dữ liệu bệnh nhân có thể dẫn đến những hậu quả nào?

*Phân tích:* Hậu quả tiềm tàng bao gồm:
- **Vận hành:** Không thể khôi phục hồ sơ bệnh nhân sau sự cố phần cứng hoặc tấn công ransomware, gián đoạn chăm sóc y tế.
- **Pháp lý và tài chính:** Vi phạm quy định bảo vệ dữ liệu y tế — ở Mỹ là HIPAA (Health Insurance Portability and Accountability Act), ở Việt Nam có thể là Luật An toàn thông tin mạng — dẫn đến phạt tiền và hình phạt pháp lý.
- **Uy tín:** Tổn hại lâu dài đến niềm tin của bệnh nhân.

Hậu quả nghiêm trọng và đa chiều này đẩy giá trị rủi ro lên mức **Cao**.

---

**Rủi ro 5: Phần cứng lỗi thời**

Một số máy tính cũ không hỗ trợ phần mềm hoặc bản cập nhật mới nhất. Nhân viên chỉ dùng chúng cho tác vụ cơ bản và không lưu trữ dữ liệu nhạy cảm trực tiếp trên đó.

*Câu hỏi đánh giá:* Các máy tính cũ này có kết nối với mạng nội bộ của phòng khám không?

*Phân tích:* Mặc dù bản thân các máy tính này không chứa dữ liệu nhạy cảm, nếu chúng kết nối với cùng mạng nội bộ với các thiết bị khác, kẻ tấn công có thể sử dụng chúng như **điểm xâm nhập ban đầu (initial access point)** — sau đó di chuyển ngang (lateral movement) qua mạng để tấn công các máy tính có dữ liệu nhạy cảm. Kỹ thuật này được gọi là **pivoting**. Hậu quả thực sự của rủi ro này phụ thuộc vào cấu trúc mạng và mức độ phân đoạn mạng (network segmentation) của phòng khám.

---

## Bốn phương án ứng phó rủi ro (Risk Response)

Sau khi đánh giá toàn bộ rủi ro, tổ chức cần quyết định cách xử lý từng rủi ro. Có bốn phương án cơ bản:

### 1. Chấp nhận (Acceptance)

Tổ chức thừa nhận rủi ro và chấp nhận hậu quả tiềm tàng nếu nó xảy ra. Đây là lựa chọn hợp lý khi chi phí giảm thiểu rủi ro cao hơn giá trị tài sản cần bảo vệ, hoặc khi rủi ro được đánh giá là có xác suất thấp và hậu quả có thể quản lý được.

Quyết định chấp nhận rủi ro phải được một người có thẩm quyền cấp cao — **người chịu trách nhiệm rủi ro (risk owner)** — phê duyệt chính thức và ghi lại trong tài liệu.

### 2. Giảm thiểu (Reduction / Mitigation)

Tổ chức triển khai các biện pháp để giảm xác suất xảy ra hoặc hậu quả của rủi ro. Đây là phương án phổ biến nhất trong an ninh mạng, bao gồm triển khai kiểm soát bảo mật, vá lỗ hổng, hoặc cải thiện quy trình.

### 3. Chuyển giao (Transference)

Tổ chức chuyển một phần hoặc toàn bộ rủi ro sang bên thứ ba. Các hình thức phổ biến:
- **Bảo hiểm mạng (Cyber Insurance):** Chuyển rủi ro tài chính sang công ty bảo hiểm.
- **Thuê ngoài (Outsourcing):** Giao cho nhà cung cấp dịch vụ bảo mật quản lý (MSSP) chịu trách nhiệm vận hành.

Lưu ý: Chuyển giao rủi ro không loại bỏ rủi ro — nó chỉ phân phối lại trách nhiệm tài chính hoặc vận hành.

### 4. Tránh né / Từ chối (Avoidance / Rejection)

Tổ chức quyết định không tham gia vào hoạt động tạo ra rủi ro. Đây là phương án cực đoan nhất và thường ảnh hưởng lớn đến hoạt động kinh doanh: đóng cửa một dòng sản phẩm, rút khỏi một thị trường, hoặc không triển khai một công nghệ nhất định.

> Bản gốc dùng thuật ngữ "Rejection" (từ chối) thay vì "Avoidance" — thuật ngữ chuẩn trong các framework quản lý rủi ro như ISO 31000 và NIST. Đã bổ sung thuật ngữ chuẩn để người đọc nhận biết khi gặp trong tài liệu chuyên ngành.

### Ví dụ minh họa: Rủi ro hỏa hoạn trong kinh doanh bánh tại nhà

| Phương án | Hành động | Ý nghĩa |
|---|---|---|
| Chấp nhận | Tin tưởng vào kỹ năng làm bánh, chấp nhận rủi ro nhỏ và chuẩn bị sẵn tiền sửa chữa | Rủi ro được ước tính là có thể quản lý được |
| Giảm thiểu | Lắp máy dò khói (giảm xác suất) và hệ thống chữa cháy tự động (giảm hậu quả) | Chi phí nhỏ nhưng giảm đáng kể cả xác suất và hậu quả |
| Chuyển giao | Mua bảo hiểm bao gồm cháy nổ liên quan đến nấu ăn | Rủi ro tài chính chuyển sang công ty bảo hiểm |
| Từ chối | Thay đổi công thức không cần lò nướng, hoặc không bắt đầu kinh doanh | Loại bỏ hoàn toàn hoạt động tạo ra rủi ro |

---

## Khẩu vị rủi ro (Risk Appetite)

**Khẩu vị rủi ro (Risk Appetite)** là mức độ rủi ro mà một tổ chức sẵn sàng chấp nhận trong quá trình theo đuổi các mục tiêu chiến lược. Đây là một quyết định quản trị quan trọng, không phải quyết định kỹ thuật.

**Tổ chức có khẩu vị rủi ro cao** sẵn sàng áp dụng công nghệ mới nhất, chấp nhận rủi ro bảo mật tiềm tàng để đổi lấy lợi thế cạnh tranh và tốc độ đổi mới. Họ đầu tư vào khả năng phục hồi (resilience) và ứng phó sự cố (incident response) để xử lý khi vi phạm xảy ra.

**Tổ chức có khẩu vị rủi ro thấp** ưu tiên sự ổn định và tuân thủ hơn tốc độ. Họ đầu tư mạnh vào phòng ngừa: tường lửa, mã hóa, kiểm soát truy cập nghiêm ngặt, và cập nhật hệ thống thường xuyên. Điều này có thể khiến họ chậm hơn trong việc áp dụng công nghệ mới.

Không có khẩu vị rủi ro nào là đúng hay sai cho mọi tổ chức. Quyết định này phụ thuộc vào:
- **Lĩnh vực hoạt động:** Ngân hàng và y tế thường có khẩu vị rủi ro thấp; startup công nghệ thường có khẩu vị rủi ro cao hơn.
- **Quy định pháp lý:** Các ngành chịu sự giám sát chặt chẽ (tài chính, y tế, năng lượng) bị ràng buộc bởi mức khẩu vị rủi ro tối thiểu theo luật.
- **Nguồn lực:** Tổ chức lớn có thể đầu tư nhiều hơn vào cả phòng ngừa lẫn phục hồi.
- **Danh tiếng:** Tổ chức phụ thuộc nhiều vào niềm tin của khách hàng (ví dụ: ngân hàng, bệnh viện) thường cần khẩu vị rủi ro thấp hơn.

---

## Thông tin bổ sung

### 1. Các framework Quản lý rủi ro An ninh mạng phổ biến

Các framework sau đây cung cấp phương pháp luận chuẩn hóa để thực hiện quản lý rủi ro an ninh mạng:

- **NIST RMF (Risk Management Framework):** Framework toàn diện gồm 7 bước: Prepare → Categorize → Select → Implement → Assess → Authorize → Monitor. Phổ biến trong cơ quan chính phủ Mỹ và tổ chức tuân thủ FISMA.
- **ISO 31000:** Tiêu chuẩn quốc tế về quản lý rủi ro, áp dụng cho mọi loại rủi ro, không chỉ an ninh mạng.
- **ISO/IEC 27005:** Tiêu chuẩn chuyên biệt về quản lý rủi ro an toàn thông tin.
- **FAIR (Factor Analysis of Information Risk):** Framework định lượng rủi ro an ninh mạng bằng con số tài chính cụ thể, phù hợp để trình bày với lãnh đạo cấp cao.

### 2. Risk Register — Công cụ theo dõi rủi ro

Trong thực tế, các tổ chức duy trì một **Risk Register** (Sổ đăng ký rủi ro) — tài liệu liệt kê toàn bộ rủi ro đã xác định, kèm theo:
- Mô tả rủi ro
- Đánh giá hậu quả và xác suất
- Giá trị rủi ro
- Phương án ứng phó đã chọn
- Người chịu trách nhiệm (risk owner)
- Trạng thái hiện tại (open / mitigated / accepted / transferred)
- Ngày đánh giá lại tiếp theo

Risk Register là tài liệu sống — cần được cập nhật định kỳ (thường mỗi quý hoặc khi có thay đổi đáng kể trong môi trường kỹ thuật).

### 3. Residual Risk — Rủi ro còn lại

Sau khi triển khai các biện pháp kiểm soát, rủi ro thường không được loại bỏ hoàn toàn mà chỉ được giảm xuống mức chấp nhận được. Phần rủi ro còn lại sau khi đã áp dụng kiểm soát gọi là **Residual Risk** (Rủi ro còn lại). Tổ chức cần xác định mức residual risk chấp nhận được và đảm bảo các kiểm soát duy trì rủi ro dưới ngưỡng đó liên tục.
