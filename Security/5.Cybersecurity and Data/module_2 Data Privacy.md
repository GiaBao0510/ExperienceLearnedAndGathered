# Quyền riêng tư dữ liệu (Data Privacy)

Trong học phần này, bạn sẽ tìm hiểu khái niệm dữ liệu, các loại dữ liệu mà tổ chức thường thu thập, quyền riêng tư dữ liệu và tầm quan trọng của nó, các phương pháp thực hành tốt nhất trong ngành để bảo vệ quyền riêng tư dữ liệu, cũng như những cơ hội nghề nghiệp trong lĩnh vực này.

---

## 1. Dữ liệu là gì và ai đang nắm giữ dữ liệu về bạn?

Bạn làm việc với dữ liệu hàng ngày: văn bản, con số, các phép đo, số liệu thống kê, dữ kiện và hình ảnh chỉ là một vài ví dụ. Dữ liệu này đến từ nhiều nguồn khác nhau: máy tính, điện thoại thông minh, mạng lưới, khảo sát trực tuyến, mạng xã hội và nhiều nguồn khác.

Hãy thử nghĩ về tất cả những dữ liệu hiện có liên quan đến chính bạn — có thể bạn sẽ bất ngờ vì số lượng tổ chức đang nắm giữ thông tin của mình:

### 1.1. Chính phủ

Chính phủ nắm giữ các dữ liệu định danh của bạn, chẳng hạn như số An sinh Xã hội (hoặc số CCCD tại Việt Nam). Khi bạn làm bằng lái xe hay các loại giấy tờ tùy thân khác do cơ quan nhà nước cấp, dữ liệu từ những giấy tờ này sẽ được lưu vào cơ sở dữ liệu liên thông giữa các cơ quan — bao gồm ngày sinh, địa chỉ nhà, ảnh chân dung, cùng các đặc điểm nhận dạng như chiều cao, cân nặng, màu mắt.

### 1.2. Y tế

Các đơn vị cung cấp dịch vụ chăm sóc sức khỏe cũng nắm giữ phần lớn dữ liệu tương tự về bạn, đồng thời lưu trữ thêm dữ liệu về bệnh sử — bao gồm kết quả xét nghiệm, chẩn đoán và các loại thuốc được kê đơn.

### 1.3. Tổ chức tài chính

Ngân hàng, công ty thẻ tín dụng hay các tổ chức tài chính khác theo dõi cách bạn sử dụng tiền: ngày, giờ, địa điểm, đơn vị giao dịch và số tiền chi tiêu cho mỗi lần mua hàng, cũng như mọi khoản tiền nạp vào tài khoản. Họ còn lưu trữ thông tin xác thực tài khoản (số tài khoản thanh toán, mã PIN, tên đăng nhập, mật khẩu) và dữ liệu cá nhân (họ tên, ngày sinh, địa chỉ nhà).

### 1.4. Tổ chức cho vay

Nắm giữ phần lớn thông tin giống như ngân hàng, cộng thêm thông tin về các khoản vay và điểm tín dụng của bạn.

### 1.5. Mạng xã hội

Các công ty mạng xã hội thu thập lượng dữ liệu rất lớn về bạn: thông tin cơ bản (tên, ngày sinh, email, số điện thoại, giới tính), thông tin hồ sơ (vị trí, lịch sử việc làm, quá trình học tập, hình ảnh đã chia sẻ), hoạt động trên ứng dụng (bài đăng, lượt thích, mối kết nối), và với một số nền tảng, cả những trang web bạn truy cập thông qua ứng dụng của họ — bao gồm địa chỉ IP, trình duyệt, hệ điều hành và thông tin phần cứng thiết bị.

---

## 2. Phân loại dữ liệu

Các chuyên gia bảo mật phân loại dữ liệu dựa trên **mức độ nhạy cảm**:

| Loại dữ liệu | Định nghĩa | Ví dụ |
|---|---|---|
| **Public data** (Dữ liệu công khai) | Bất kỳ ai cũng có thể truy cập, sử dụng và phân phối lại mà không bị hạn chế | Tài liệu tiếp thị, thông cáo báo chí, tác phẩm thuộc phạm vi công cộng |
| **Confidential** (Dữ liệu mật) | Bất kỳ dữ liệu nào được tổ chức bảo vệ khỏi truy cập trái phép | Thông tin đặc thù công ty (thông tin độc quyền, mã nguồn), hồ sơ nhân viên, lịch sử mua hàng, PII, PHI |
| **Proprietary** (Dữ liệu độc quyền) | Dữ liệu do tổ chức sở hữu/tạo ra, liên quan đến sản phẩm hoặc hoạt động, cần giữ bí mật | Bí mật thương mại, thông tin R&D, kế hoạch sản phẩm, dữ liệu tài chính công ty, quy trình nội bộ |
| **Private data** (Dữ liệu riêng tư) | Thông tin về một cá nhân và đời tư mà bên khác không được thu thập, sử dụng hoặc tiết lộ nếu không có sự cho phép | Hoạt động tài chính, số thẻ tín dụng, thông tin đăng nhập email, PII, PHI |
| **PII** (Personally Identifiable Information — Thông tin định danh cá nhân) | Dữ liệu riêng tư có thể dùng để xác định danh tính một người cụ thể | Ngày sinh, địa chỉ, số điện thoại, số giấy tờ tùy thân, ảnh chân dung, quá trình làm việc/học tập |
| **PHI** (Protected Health Information — Thông tin sức khỏe được bảo vệ) | Dữ liệu riêng tư trong hồ sơ y tế có thể dùng để nhận diện cá nhân | Chẩn đoán, kết quả xét nghiệm, đơn thuốc, thông tin bảo hiểm y tế |

**Lưu ý quan trọng — các loại dữ liệu này không tách biệt hoàn toàn mà có quan hệ lồng nhau:**

```
Confidential (Dữ liệu mật)
 ├── Proprietary (Dữ liệu độc quyền)
 └── Private data (Dữ liệu riêng tư)
        ├── PII
        └── PHI
```

Ví dụ: hãy xem xét **địa chỉ nhà** của bạn — nó đồng thời là:
- **PII**, vì ai đó có thể dùng nó để xác định danh tính bạn;
- **Dữ liệu riêng tư**, vì công ty chỉ được chia sẻ nó với bên khác khi có sự cho phép;
- **Dữ liệu mật**, vì tổ chức phải ngăn chặn các bên không có thẩm quyền truy cập vào nó — chính vì cùng lúc là PII và dữ liệu riêng tư.

---

## 3. Bảo vệ dữ liệu: Data Security và Data Privacy

Các tổ chức dùng hai thuật ngữ để mô tả việc bảo vệ dữ liệu mật: **bảo mật dữ liệu (data security)** và **quyền riêng tư dữ liệu (data privacy)**. Hai khái niệm này liên quan chặt chẽ nhưng **không đồng nhất**:

- **Data security (Bảo mật dữ liệu):** Là cách một tổ chức **bảo vệ** dữ liệu mật khỏi sự truy cập, tiết lộ hoặc phá hủy trái phép — bao gồm cả thông tin đặc thù công ty (thông tin độc quyền) lẫn dữ liệu cá nhân của khách hàng/nhân viên (địa chỉ, số An sinh Xã hội, số thẻ tín dụng...). Nói cách khác, đây là các **biện pháp kỹ thuật và quy trình** để ngăn chặn truy cập trái phép (mã hóa, tường lửa, kiểm soát quyền truy cập...).

- **Data privacy (Quyền riêng tư dữ liệu):** Là việc đảm bảo dữ liệu cá nhân được **thu thập, xử lý, lưu trữ và sử dụng đúng cách**, tuân thủ pháp luật và đúng với những gì cá nhân đó đã đồng ý. Nói cách khác, data security trả lời câu hỏi *"làm sao ngăn kẻ xấu truy cập dữ liệu?"*, còn data privacy trả lời câu hỏi *"tổ chức có đang sử dụng dữ liệu đúng với những gì đã cam kết và đúng với sự đồng ý của chủ sở hữu dữ liệu hay không?"*.

---
## 4. Vì sao quyền riêng tư dữ liệu quan trọng?

Quyền riêng tư dữ liệu là vấn đề vô cùng quan trọng. Nếu một công ty không triển khai kế hoạch bảo vệ quyền riêng tư dữ liệu, hậu quả có thể vô cùng nghiêm trọng. Vi phạm quyền riêng tư thường gây tổn hại theo ba nhóm hậu quả:

### 4.1. Ba nhóm hậu quả khi rò rỉ dữ liệu

- **Hậu quả tài chính (Financial):** các khoản phạt, tiền phạt, lệ phí, thỏa thuận dàn xếp, phán quyết của tòa án.
- **Hậu quả vận hành (Operational):** công ty cần đưa hệ thống và dữ liệu trở lại trạng thái hoạt động bình thường nhanh nhất có thể — quá trình khôi phục đòi hỏi thời gian, công sức, và do đó cũng kéo theo hậu quả tài chính. Tổ chức có thể phải tạm ngưng một phần hoạt động để điều tra vụ xâm nhập, càng kéo dài thời gian khôi phục.
- **Hậu quả uy tín (Reputational):** bắt nguồn từ việc công chúng biết đến và chỉ trích sự thất bại của tổ chức trong việc bảo vệ dữ liệu cá nhân — có thể làm giảm lòng tin khách hàng, khiến họ chuyển sang dịch vụ khác, thậm chí khiến đối tác chấm dứt hợp tác.

> **Câu hỏi kiểm tra kiến thức:**
> 1. *Tổ chức cần triển khai một quy trình mới, trong đó nhân viên kiểm soát chất lượng phải rà soát code trước khi đưa lên production.* → **Hậu quả vận hành (Operational)** — vì đây là thay đổi trong quy trình làm việc để khắc phục/phòng ngừa sự cố.
> 2. *Sở Dịch vụ Tài chính bang New York phạt First American Financial Corp 487.616 USD.* → **Hậu quả tài chính (Financial)** — khoản phạt trực tiếp bằng tiền.
> 3. *First American Financial Corp đánh mất lòng tin khách hàng, dẫn đến ít giao dịch kinh doanh hơn.* → **Hậu quả uy tín (Reputational)** — thiệt hại đến từ sự sụt giảm lòng tin, không phải khoản phạt hay chi phí khắc phục trực tiếp.

### 4.2. Nghiên cứu tình huống: Vụ rò rỉ dữ liệu Equifax (2017)

Bạn đã bao giờ bị ảnh hưởng bởi một vụ rò rỉ dữ liệu chưa? Nếu ai đó từng cố dùng trộm thẻ tín dụng hoặc xâm nhập tài khoản mạng xã hội của bạn — câu trả lời là có! Điều tương tự cũng có thể xảy ra với các công ty, nhưng ở quy mô lớn hơn rất nhiều.

Để hiểu rõ tác động của một vụ xâm nhập dữ liệu ở quy mô lớn, hãy xem xét vụ tấn công mạng nhắm vào **Equifax** — một công ty xếp hạng tín dụng lớn của Mỹ — vào năm 2017. Sự cố xảy ra sau khi tổ chức này **không kịp thời áp dụng bản vá bảo mật** cho một lỗ hổng đã biết trong hệ thống, tạo điều kiện cho tin tặc xâm nhập vào mạng lưới nội bộ. Mạng lưới này chứa các thông tin xác thực quản trị viên nhưng lại được lưu trữ **không mã hóa** và thiếu các biện pháp kiểm soát truy cập cơ bản. Sau khi chiếm được các thông tin xác thực đó, tin tặc đã giành quyền kiểm soát phần lớn hệ thống và duy trì trạng thái này trong nhiều tháng mà không bị phát hiện.

Vụ xâm nhập đã làm lộ một lượng dữ liệu khổng lồ: thông tin của khoảng **147 triệu người** (bao gồm tên và ngày sinh), khoảng **145,5 triệu số An sinh Xã hội**, cùng khoảng **209.000 số thẻ thanh toán** và ngày hết hạn thẻ.

**Hậu quả đối với Equifax:**

- **Tài chính:** Equifax đồng ý một thỏa thuận dàn xếp toàn cầu trị giá lên tới **700 triệu USD** với Ủy ban Thương mại Liên bang Mỹ (FTC), Cục Bảo vệ Tài chính Người tiêu dùng (CFPB) và các bang — trong đó **tới 425 triệu USD** dành riêng để bồi thường/hỗ trợ người tiêu dùng bị ảnh hưởng, cộng thêm khoản phạt hành chính và chi phí nâng cấp hệ thống bảo mật.
- **Vận hành:** Sau khi phát hiện vi phạm, Equifax mất khoảng một tháng để điều tra nội bộ trước khi công bố ra công chúng, và tiếp tục điều tra trong thời gian dài sau đó. Công ty cũng dành thời gian và nguồn lực đáng kể để nâng cấp bảo mật và sửa đổi chính sách nội bộ.
- **Uy tín:** Equifax hứng chịu làn sóng chỉ trích trên truyền thông, đặc biệt vì **chậm trễ công bố** vụ xâm nhập và triển khai thiếu sót trang web hỗ trợ người tiêu dùng bị ảnh hưởng.

> **Bài học rút ra:** Vụ Equifax cho thấy rõ mối liên hệ giữa các nội dung đã học — đây chính là hậu quả thực tế khi vi phạm nguyên tắc **Least Privilege** và thiếu **mã hóa dữ liệu** (đã đề cập ở tài liệu về User/Role/Permission và MinIO): một hệ thống nội bộ lưu thông tin xác thực quản trị viên mà không mã hóa, không kiểm soát truy cập, chính là "quả bom hẹn giờ" chờ bị khai thác.

---

## 5. Các thực hành tốt nhất về quyền riêng tư (GAPP)

Viện Kế toán Công chứng Hoa Kỳ (AICPA) và Viện Kế toán Công chứng Canada (CICA) đã xây dựng một tiêu chuẩn được công nhận rộng rãi về quyền riêng tư dữ liệu, gọi là **GAPP — Generally Accepted Privacy Principles** (Các Nguyên tắc Quyền riêng tư Được chấp nhận Chung), gồm 10 thực hành tốt nhất dành cho tổ chức:

| # | Nguyên tắc | Ý nghĩa | Ví dụ |
|---|---|---|---|
| 1 | **Management** (Quản lý) | Xác định hình thức lập tài liệu, trao đổi thông tin, thiết lập quy trình và trách nhiệm giải trình rõ ràng | Công ty có một chuyên gia phụ trách tuân thủ (compliance officer) |
| 2 | **Notice** (Thông báo) | Thông báo cho mọi người về quy trình bảo mật thông tin và cách thức sử dụng dữ liệu | Trình hướng dẫn đăng ký tài khoản nêu rõ chính sách bảo mật dữ liệu |
| 3 | **Choice and consent** (Lựa chọn và đồng ý) | Nhận được sự đồng ý trước khi lưu trữ, chia sẻ, sử dụng thông tin cá nhân | Bạn chọn đăng ký (opt-in) nhận email cập nhật hoặc gửi dữ liệu sử dụng phần mềm |
| 4 | **Collection** (Thu thập) | Chỉ thu thập thông tin thực sự cần thiết cho mục đích đã nêu | Dịch vụ streaming cần email, thiết bị, thẻ tín dụng — không cần số CCCD hay số An sinh Xã hội |
| 5 | **Use, retention, and disposal** (Sử dụng, lưu trữ, hủy bỏ) | Chỉ dùng đúng mục đích đã được chấp thuận; không giữ lâu hơn cần thiết; hủy khi hết nhu cầu | Phòng khám lưu hồ sơ bệnh án 6 năm rồi hủy bỏ |
| 6 | **Access** (Quyền truy cập) | Cho phép cá nhân truy cập vào chính thông tin của họ khi có yêu cầu | Cơ quan chính phủ cho phép gửi yêu cầu bằng văn bản để nhận bản sao dữ liệu |
| 7 | **Disclosure to third parties** (Tiết lộ cho bên thứ ba) | Chỉ chia sẻ thông tin với bên thứ ba sau khi đã thông báo và được đồng ý | Khi mua nhà, bạn được hỏi có muốn nhận tài liệu tiếp thị từ bên thứ ba không |
| 8 | **Security** (Bảo mật) | Bảo vệ thông tin cá nhân khỏi truy cập trái phép | Yêu cầu xác thực hai yếu tố (2FA) khi đăng nhập |
| 9 | **Quality** (Chất lượng) | Đảm bảo thông tin lưu trữ đầy đủ và chính xác | Nhân viên xác nhận số điện thoại/số tài khoản đúng định dạng |
| 10 | **Monitoring and enforcement** (Giám sát và thực thi) | Giám sát và thực thi các chính sách bảo mật đã đề ra | Kiểm tra ngẫu nhiên định kỳ, quy trình xử lý khiếu nại về quyền riêng tư |

> **Câu hỏi kiểm tra kiến thức:**
> 1. *Một ngân hàng địa phương thiết lập xác thực hai yếu tố cho khách hàng khi đăng nhập tài khoản.* → Nguyên tắc **Security**.
> 2. *Phòng khám bác sĩ yêu cầu bệnh nhân ký một biểu mẫu nêu rõ quy trình bảo vệ hồ sơ bệnh án.* → Nguyên tắc **Notice**.

---

## 6. Cơ hội nghề nghiệp trong lĩnh vực quyền riêng tư dữ liệu

| Vị trí | Mô tả công việc |
|---|---|
| **Data Privacy Specialist** | Thực hiện các nhiệm vụ đảm bảo tuân thủ luật pháp và quy định về quyền riêng tư — kiểm toán, đánh giá rủi ro, điều tra khiếu nại, đề xuất chính sách/quy trình bảo vệ dữ liệu. |
| **Data Privacy Analyst** | Thực hiện công việc kỹ thuật liên quan đến quyền riêng tư — xây dựng hướng dẫn và cơ chế kiểm soát tự động, giám sát/triển khai chính sách ngăn chặn truy cập trái phép, hỗ trợ điều tra sự cố lộ lọt dữ liệu. |
| **Data Privacy Manager** | Giám sát nhiều nhiệm vụ liên quan đến quyền riêng tư — lập sơ đồ luồng dữ liệu trong tổ chức, xác định rủi ro tiềm ẩn, đề xuất giải pháp, thiết kế chương trình đào tạo và nâng cao nhận thức toàn tổ chức. |

---

## 7. Tổng kết

- Gần như mọi tổ chức bạn từng tương tác — chính phủ, bệnh viện, ngân hàng, mạng xã hội — đều đang nắm giữ một phần dữ liệu về bạn, ở nhiều mức độ nhạy cảm khác nhau (từ Public đến Confidential/PII/PHI).
- **Data security** và **data privacy** là hai khái niệm liên quan nhưng khác nhau: security là *cơ chế bảo vệ* dữ liệu khỏi truy cập trái phép, còn privacy là *nguyên tắc sử dụng đúng cách* dữ liệu cá nhân, dựa trên sự đồng ý của chủ sở hữu dữ liệu.
- Vi phạm quyền riêng tư gây ra ba nhóm hậu quả liên kết chặt chẽ với nhau: tài chính, vận hành và uy tín — vụ Equifax là minh chứng điển hình cho việc thiệt hại có thể lên tới hàng trăm triệu USD chỉ vì một lỗ hổng chưa được vá.
- **GAPP** là khung 10 nguyên tắc thực hành tốt nhất giúp tổ chức xây dựng một chương trình bảo vệ quyền riêng tư dữ liệu toàn diện, từ quản lý, thông báo, thu thập cho đến giám sát thực thi.
- Đây là lĩnh vực có nhu cầu nhân lực lớn, với lộ trình sự nghiệp rõ ràng từ chuyên viên phân tích kỹ thuật đến vị trí quản lý chiến lược.

---

### Mở rộng

- **GDPR (General Data Protection Regulation)** của Liên minh Châu Âu: khung pháp lý về bảo vệ dữ liệu cá nhân có ảnh hưởng lớn nhất thế giới hiện nay, là nguồn cảm hứng cho nhiều luật bảo vệ dữ liệu ở các quốc gia khác — đáng tìm hiểu thêm về các quyền cốt lõi như "quyền được quên" (right to be forgotten) và "quyền truy cập dữ liệu" (data subject access request).
- **CCPA (California Consumer Privacy Act)**: luật bảo vệ quyền riêng tư người tiêu dùng của bang California, Mỹ — một ví dụ khác về cách một khu vực pháp lý cụ thể hóa các nguyên tắc GAPP thành luật.
- **Nghị định 13/2023/NĐ-CP về bảo vệ dữ liệu cá nhân** tại Việt Nam: khung pháp lý gần nhất và trực tiếp áp dụng cho các hệ thống bạn xây dựng nếu phục vụ người dùng Việt Nam — đáng đọc song song với GDPR để so sánh sự tương đồng/khác biệt.
- **Liên hệ trực tiếp đến công việc backend Golang của bạn:** khi thiết kế database cho một hệ thống thực tế, việc phân loại cột dữ liệu nào là PII/PHI ngay từ khi thiết kế schema sẽ quyết định: cột nào cần mã hóa tại tầng ứng dụng hoặc database (encryption at rest), cột nào cần áp dụng Row-Level Security hoặc Column-level Permission (đã học ở tài liệu User/Role/Permission), và cột nào cần có cơ chế xóa/ẩn dữ liệu khi người dùng yêu cầu (right to erasure) — đây chính là điểm giao thoa giữa kiến thức quyền riêng tư dữ liệu và công việc thiết kế hệ thống hàng ngày của một backend engineer.
- **Data anonymization & pseudonymization**: hai kỹ thuật kỹ thuật hóa các nguyên tắc privacy — ẩn danh hóa hoàn toàn dữ liệu (không thể khôi phục danh tính) so với giả danh hóa (có thể khôi phục nếu có khóa riêng) — thường được dùng khi chia sẻ dữ liệu cho mục đích phân tích/thống kê mà vẫn cần tuân thủ GAPP/GDPR.