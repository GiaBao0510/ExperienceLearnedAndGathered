# Chiến lược An ninh mạng (Security Strategy)

## Giới thiệu

Để chống lại các cuộc tấn công mạng, các tổ chức cần xây dựng và thực thi một chiến lược an ninh có hệ thống. Câu hỏi cốt lõi là: làm thế nào để giảm thiểu mối đe dọa và nâng cao khả năng ứng phó khi sự cố xảy ra?

Module này giới thiệu các thành phần chính của một chiến lược an ninh toàn diện, mô hình đánh giá mức độ trưởng thành an ninh (C2M2), và 10 bước thực hành được khuyến nghị bởi Trung tâm An ninh mạng Quốc gia Vương quốc Anh (NCSC).

---

## Chiến lược An ninh là gì?

Chiến lược an ninh là kế hoạch tổng thể và có cơ sở của một tổ chức nhằm bảo vệ tài sản kỹ thuật số và vật lý trước các mối đe dọa. Một chiến lược an ninh toàn diện bao gồm năm thành phần cốt lõi:

**Đánh giá rủi ro (Risk Assessment)** Xác định và hiểu rõ các mối đe dọa, lỗ hổng tiềm ẩn, và tác động tiềm tàng lên tài sản của tổ chức. Đây là nền tảng để ưu tiên nguồn lực bảo mật.

**Chính sách và quy trình an ninh (Security Policies and Procedures)** Thiết lập các quy tắc, hướng dẫn và quy trình rõ ràng để duy trì an ninh — bao gồm chính sách mật khẩu, kiểm soát truy cập, phân loại dữ liệu và quy trình xử lý sự cố.

**Nâng cao nhận thức và đào tạo (Security Awareness and Training)** Trang bị cho nhân viên kiến thức và kỹ năng để nhận diện và ngăn chặn các vi phạm an ninh — đặc biệt là các mối đe dọa từ social engineering và phishing.

**Ứng phó sự cố (Incident Response — IR)** Tập hợp các hành động mà tổ chức thực hiện để **chuẩn bị, phát hiện, ngăn chặn và phục hồi** sau các cuộc tấn công mạng. IR cần được lên kế hoạch và thực hành trước, không phải ứng phó ngẫu nhiên.

**Kiểm toán và thử nghiệm (Audit and Testing)** Đảm bảo các biện pháp an ninh hoạt động hiệu quả như thiết kế — thông qua penetration testing, vulnerability assessment, và kiểm toán định kỳ.

Trách nhiệm xây dựng chiến lược thường thuộc về nhóm bảo mật CNTT nội bộ hoặc chuyên gia tư vấn an ninh mạng bên ngoài. Chiến lược nên được thiết lập từ khi tổ chức thành lập — nhưng không bao giờ là quá muộn để bắt đầu. Quan trọng hơn, chiến lược cần được **xem xét và cập nhật ít nhất hàng năm** (hoặc thường xuyên hơn khi có thay đổi lớn về công nghệ, cấu trúc tổ chức, hoặc bối cảnh mối đe dọa).

---

## Mức độ Trưởng thành An ninh (Security Maturity)

Mức độ trưởng thành an ninh phản ánh **khả năng thực tế** của một tổ chức trong việc bảo vệ tài sản và ứng phó với mối đe dọa. Chiến lược an ninh càng được xác định rõ ràng và thực thi nhất quán, mức độ trưởng thành càng cao.

Có nhiều mô hình đánh giá mức độ trưởng thành an ninh. Module này tập trung vào **C2M2 — Cybersecurity Capability Maturity Model** (Mô hình Mức độ Trưởng thành Năng lực An ninh mạng).

### Giới thiệu C2M2

C2M2 được phát triển bởi sự hợp tác giữa Bộ Năng lượng Hoa Kỳ (DOE), Bộ An ninh Nội địa (DHS) và các tổ chức thuộc khu vực tư nhân. Mô hình này giúp tổ chức **đánh giá trạng thái hiện tại** của năng lực an ninh mạng và **xác định lộ trình cải thiện** — từ mức chưa trưởng thành (phản ứng, thiếu cấu trúc) sang mức trưởng thành cao (chủ động, có hệ thống, tích hợp vào văn hóa tổ chức).

> Lưu ý quan trọng: Một tổ chức có thể đạt mức độ trưởng thành cao trong một lĩnh vực an ninh nhưng lại ở mức thấp trong lĩnh vực khác. Đánh giá C2M2 là đánh giá theo từng lĩnh vực (domain), không phải một điểm số tổng thể duy nhất.

### Thang đo C2M2 (MIL0 đến MIL3)

C2M2 định nghĩa bốn mức độ trưởng thành triển khai (MIL — Maturity Implementation Level):

| Mức      | Đặc điểm tổng quát                                                                                                           |
| -------- | ---------------------------------------------------------------------------------------------------------------------------- |
| **MIL0** | Các hoạt động an ninh mạng **không được thực hiện**                                                                          |
| **MIL1** | Các hoạt động được thực hiện nhưng **ngẫu hứng, không có quy trình chuẩn**                                                   |
| **MIL2** | Các hoạt động được **ghi chép, có nguồn lực, và thực hiện nhất quán** hơn so với MIL1                                        |
| **MIL3** | An ninh mạng được **tích hợp vào chiến lược tổ chức**, có phân công trách nhiệm rõ ràng, được đánh giá và cải thiện liên tục |

**MIL0 — Không hoạt động:** Tổ chức không thực hiện bất kỳ hoạt động an ninh mạng có hệ thống nào.

**MIL1 — Khởi đầu:** Các hoạt động an ninh tồn tại nhưng mang tính tự phát và phản ứng — được thực hiện không nhất quán, thiếu quy trình và tài liệu chính thức.

**MIL2 — Quản lý cơ bản:**

- Đặc điểm quản lý: Các quy trình được ghi chép lại; nguồn lực đầy đủ được cung cấp.
- Đặc điểm tiếp cận: Hoạt động an ninh hoàn thiện và nhất quán hơn so với MIL1.

**MIL3 — Tích hợp chiến lược:**

- Đặc điểm quản lý: Hoạt động được dẫn dắt bởi chính sách tổ chức; trách nhiệm, nghĩa vụ và quyền hạn được phân công rõ ràng; nhân viên có kỹ năng và kiến thức đầy đủ; hiệu quả được theo dõi và đánh giá định kỳ.
- Đặc điểm tiếp cận: Hoạt động an ninh hoàn thiện và tiên tiến hơn so với MIL2.

### Ứng dụng C2M2 trong thực tế

Để minh họa cách C2M2 được áp dụng, dưới đây là ví dụ từ Bộ Năng lượng Hoa Kỳ về **giảm thiểu lỗ hổng bảo mật** theo từng mức độ trưởng thành:

**MIL1 — Hoạt động tùy ý:**

- Xác định nguồn dữ liệu để hỗ trợ phát hiện lỗ hổng một cách tùy ý (ad hoc).
- Thu thập và phân tích dữ liệu lỗ hổng một cách tùy ý.
- Thực hiện đánh giá và khắc phục lỗ hổng một cách tùy ý.

**MIL2 — Hoạt động có kế hoạch:**

- Giám sát các nguồn dữ liệu lỗ hổng liên quan đến **tài sản có mức độ ưu tiên cao**.
- Thực hiện đánh giá lỗ hổng **thường xuyên** và khi có sự kiện quan trọng (ví dụ: thay đổi cấu hình mạng).
- Phân tích và **xếp hạng lỗ hổng theo mức độ nghiêm trọng**, sau đó xử lý theo thứ tự ưu tiên.
- Đánh giá tác động vận hành trước khi áp dụng biện pháp vá lỗi.
- Chia sẻ thông tin lỗ hổng với các bên liên quan.

**MIL3 — Hoạt động tích hợp chiến lược:**

- Giám sát các nguồn dữ liệu lỗ hổng liên quan đến **tất cả tài sản IT và OT** (Operational Technology) có liên quan.
- Đảm bảo đánh giá lỗ hổng được thực hiện bởi **người độc lập** với hoạt động đang được đánh giá.
- Xem xét **hiệu quả** của các biện pháp giảm thiểu lỗ hổng đã áp dụng.
- Thiết lập và duy trì kênh tiếp nhận **báo cáo lỗ hổng từ bên ngoài** (ví dụ: qua trang web công khai — vulnerability disclosure program).

> **OT (Operational Technology)** là các hệ thống phần cứng và phần mềm giám sát và kiểm soát thiết bị vật lý, quy trình và cơ sở hạ tầng — ví dụ: hệ thống SCADA trong nhà máy điện, hệ thống điều khiển trong sản xuất công nghiệp. Khác với IT (Information Technology), OT tương tác trực tiếp với thế giới vật lý

---

## Điểm khởi đầu cho tổ chức — 10 Bước An ninh mạng của NCSC

Khi bắt đầu xây dựng chiến lược an ninh mạng, nhiều tổ chức không biết nên bắt đầu từ đâu. Một hướng dẫn thực hành được khuyến nghị là **10 Bước An ninh mạng** do [Trung tâm An ninh mạng Quốc gia Vương quốc Anh (NCSC)](https://www.ncsc.gov.uk/) công bố tại [ncsc.gov.uk/collection/10-steps](https://www.ncsc.gov.uk/collection/10-steps).

Framework này phân chia nhiệm vụ bảo vệ tổ chức thành 10 bước cụ thể, có thể triển khai theo từng giai đoạn:

| Bước | Tên                                                                  | Mục tiêu                                                                                            |
| ---- | -------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| 1    | Quản lý rủi ro (Risk Management)                                     | Áp dụng phương pháp dựa trên rủi ro để ưu tiên bảo vệ dữ liệu và hệ thống                           |
| 2    | Thu hút và đào tạo (Engage and Train People)                         | Xây dựng văn hóa an ninh — đảm bảo mọi người trong tổ chức đều hiểu vai trò và trách nhiệm của mình |
| 3    | Quản lý tài sản (Asset Management)                                   | Hiểu rõ tổ chức đang có những dữ liệu và hệ thống nào, chúng hỗ trợ nhu cầu kinh doanh nào          |
| 4    | Kiến trúc và cấu hình (Architecture and Configuration)               | Thiết kế, xây dựng và vận hành hệ thống một cách an toàn từ đầu                                     |
| 5    | Quản lý lỗ hổng (Vulnerability Management)                           | Chủ động phát hiện và khắc phục lỗ hổng trong suốt vòng đời hệ thống                                |
| 6    | Quản lý danh tính và quyền truy cập (Identity and Access Management) | Kiểm soát chặt chẽ ai và cái gì có thể truy cập vào hệ thống và dữ liệu                             |
| 7    | Bảo mật dữ liệu (Data Security)                                      | Bảo vệ dữ liệu tại các điểm dễ bị tổn thương — cả khi lưu trữ (at rest) và truyền tải (in transit)  |
| 8    | Ghi nhật ký và giám sát (Logging and Monitoring)                     | Thiết kế hệ thống để có thể phát hiện, điều tra và ứng phó với sự cố                                |
| 9    | Quản lý sự cố (Incident Management)                                  | Chuẩn bị và thực hành kế hoạch ứng phó sự cố **trước khi** sự cố xảy ra                             |
| 10   | An ninh chuỗi cung ứng (Supply Chain Security)                       | Quản lý rủi ro đến từ nhà cung cấp và đối tác trong chuỗi cung ứng                                  |

Bằng cách triển khai tuần tự 10 bước này, các tổ chức có thể giảm thiểu đáng kể khả năng xảy ra tấn công và giảm thiểu tác động khi sự cố không thể tránh khỏi xảy ra.

---

## Hệ sinh thái sản phẩm và dịch vụ an ninh mạng

Khi xây dựng chiến lược an ninh, hầu hết tổ chức không cần và không nên tự phát triển mọi công cụ từ đầu. Một thị trường rộng lớn các sản phẩm và dịch vụ an ninh mạng đang tồn tại để hỗ trợ.

Hầu hết các tổ chức lớn sử dụng sản phẩm từ **nhiều nhà cung cấp khác nhau** — ví dụ: hệ thống DLP (Data Loss Prevention) từ nhà cung cấp A và tường lửa từ nhà cung cấp B. Hệ sinh thái này được hỗ trợ bởi nhiều thành phần:

- **Nhà cung cấp thương mại:** Công ty bán sản phẩm và dịch vụ an ninh mạng (Palo Alto Networks, CrowdStrike, Splunk, v.v.).
- **Tổ chức tiêu chuẩn và quy định:** NIST, ISO, PCI DSS, GDPR đặt ra yêu cầu và hướng dẫn.
- **Cơ quan chính phủ:** CISA (Mỹ), NCSC (Anh), ANSSI (Pháp) cung cấp advisory, framework và hỗ trợ miễn phí.
- **Cộng đồng mã nguồn mở:** OWASP, Metasploit, Snort và nhiều dự án khác cung cấp công cụ miễn phí.

---

## Bổ sung kiến thức

### 1. Bổ sung các framework an ninh quan trọng khác

Tài liệu đề cập C2M2 và NCSC 10 Steps — đây là hai framework hữu ích, nhưng người học cần biết thêm về các framework phổ biến hơn trong thực tế doanh nghiệp và đặc biệt có giá trị với Backend Engineer:

- **NIST Cybersecurity Framework (CSF):** Framework phổ biến nhất tại Mỹ, gồm 5 chức năng: Identify → Protect → Detect → Respond → Recover. Phiên bản 2.0 (2024) bổ sung thêm Govern. Rất phù hợp để tổ chức đánh giá và lên kế hoạch cải thiện.
- **ISO/IEC 27001:** Tiêu chuẩn quốc tế về quản lý an ninh thông tin (ISMS). Nhiều doanh nghiệp yêu cầu nhà cung cấp/đối tác có chứng nhận ISO 27001.

- **SOC 2 (Service Organization Control 2):** Đặc biệt quan trọng với SaaS company và cloud service provider. Nhiều khách hàng doanh nghiệp yêu cầu nhà cung cấp có báo cáo SOC 2.
- **Zero Trust Architecture:** Mô hình an ninh hiện đại "không tin tưởng mặc định, luôn xác minh" — đặc biệt phù hợp với môi trường cloud-native và microservices mà Backend Engineer thường làm việc.

### 2. Liên hệ thực tế cho Backend Engineer

Chiến lược an ninh không chỉ là trách nhiệm của team bảo mật — Backend Engineer đóng vai trò quan trọng trong nhiều bước:

- **Bước 4 (Kiến trúc và cấu hình):** Secure coding practices, không hardcode secrets, dùng secret manager (HashiCorp Vault, AWS Secrets Manager).
- **Bước 5 (Quản lý lỗ hổng):** Dependency scanning (Snyk, Dependabot), SAST (Static Application Security Testing) trong CI/CD pipeline.
- **Bước 6 (IAM):** Thiết kế API với xác thực và phân quyền đúng đắn (OAuth 2.0, RBAC, Principle of Least Privilege).
- **Bước 7 (Bảo mật dữ liệu):** Mã hóa dữ liệu nhạy cảm at rest và in transit, không log thông tin nhạy cảm.
- **Bước 8 (Logging và monitoring):** Structured logging với đủ context để điều tra sự cố, tích hợp với SIEM.

### 3. Security by Design vs. Security as an Afterthought

Một điểm quan trọng không được đề cập trong tài liệu: chiến lược an ninh hiệu quả nhất là khi an ninh được tích hợp vào thiết kế từ đầu (**Security by Design**), thay vì thêm vào sau như một lớp bọc bên ngoài (**Security as an Afterthought**). Điều này đặc biệt đúng với phát triển phần mềm: việc vá lỗi bảo mật sau khi hệ thống đã triển khai tốn kém gấp nhiều lần so với thiết kế đúng từ đầu — nguyên tắc này được gọi là **"shift left security"** (đưa bảo mật về sớm hơn trong vòng đời phát triển).
