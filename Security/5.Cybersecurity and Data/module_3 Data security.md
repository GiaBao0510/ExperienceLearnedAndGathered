# Bảo mật dữ liệu (Data Security)

## Mục lục
1. [Giới thiệu](#1-giới-thiệu)
2. [Mô hình CIA Triad](#2-mô-hình-cia-triad)
3. [5 tham số bảo mật cốt lõi](#3-5-tham-số-bảo-mật-cốt-lõi)
4. [Các biện pháp kiểm soát bảo mật (Controls)](#4-các-biện-pháp-kiểm-soát-bảo-mật-controls)
5. [Ngăn ngừa mất mát dữ liệu (DLP)](#5-ngăn-ngừa-mất-mát-dữ-liệu-dlp)
6. [Cơ hội nghề nghiệp](#6-cơ-hội-nghề-nghiệp)
7. [Tổng kết](#7-tổng-kết)
8. [Mở rộng](#mở-rộng)

---

## 1. Giới thiệu

Trong học phần này, bạn sẽ tìm hiểu những kiến thức cơ bản về bảo mật dữ liệu. Bạn sẽ được học về mô hình **CIA** (Tính bảo mật, Tính toàn vẹn và Tính sẵn sàng) — một mô hình thiết yếu trong bảo mật thông tin — và áp dụng các khái niệm này vào những tình huống thực tế. Bạn cũng sẽ khám phá ba loại biện pháp kiểm soát bảo mật dữ liệu, thực hành phân loại các ví dụ về chúng, tìm hiểu về các hệ thống ngăn ngừa mất mát dữ liệu, cũng như cơ hội nghề nghiệp trong lĩnh vực này.

---

## 2. Mô hình CIA Triad

Việc bảo vệ dữ liệu dựa trên ba mục tiêu chính: tính bảo mật, tính toàn vẹn và tính sẵn sàng. Ba mục tiêu này tạo nên mô hình **CIA** (**C**onfidentiality, **I**ntegrity, **A**vailability) — một mô hình bảo mật thông tin nổi tiếng và là nền tảng của an ninh mạng. Mô hình CIA đóng vai trò như một danh mục kiểm tra bảo mật tổng quát: mỗi mục tiêu cung cấp một góc nhìn để đánh giá chương trình bảo mật dữ liệu, giúp chuyên gia an ninh mạng nhận diện lỗ hổng và xác định giải pháp phù hợp.

> **Lưu ý:** Rất khó xác định chính xác nguồn gốc của thuật ngữ "bộ ba CIA" (CIA triad). Một số tài liệu phổ biến cho rằng các khái niệm nền tảng của nó (bảo vệ bí mật, đảm bảo thông tin không bị giả mạo, duy trì khả năng tiếp cận thông tin) đã được các nhà lãnh đạo quân sự vận dụng từ nhiều thế kỷ trước — kể cả trong các chiến dịch của Julius Caesar. Đây là một giai thoại thường được dùng để minh họa cho việc giảng dạy hơn là một sự kiện lịch sử đã được xác thực chặt chẽ, nên bạn có thể xem đây là thông tin tham khảo thú vị hơn là một mốc lịch sử chính xác.

### 2.1. Confidentiality (Tính bảo mật)

Tính bảo mật nghĩa là giữ kín dữ liệu: chỉ những người được ủy quyền mới có thể truy cập hoặc tiết lộ dữ liệu đó. Ví dụ, các công ty phần mềm thường giữ bí mật mã nguồn ứng dụng để duy trì lợi thế cạnh tranh, bằng cách giới hạn quyền truy cập chỉ cho những nhân viên thực sự cần đến nó. Tính bảo mật cũng bao gồm dữ liệu cá nhân của mỗi người — chẳng hạn cơ sở y tế phải đảm bảo dữ liệu thu thập trong quá trình điều trị (chẩn đoán bệnh, đơn thuốc) được giữ kín; trừ một số trường hợp ngoại lệ, chỉ bạn, bác sĩ và nhân viên y tế được ủy quyền mới có quyền truy cập.

Trên thực tế, đảm bảo tính bảo mật đòi hỏi phải triển khai các biện pháp bảo vệ nhằm cấp đúng mức độ truy cập cho đúng nhóm người dùng, vào đúng thời điểm và bằng đúng phương thức phù hợp — đây chính là ý tưởng cốt lõi đằng sau nguyên tắc **Least Privilege** (đã đề cập trong tài liệu về User/Role/Permission).

### 2.2. Integrity (Tính toàn vẹn)

Tính toàn vẹn nghĩa là đảm bảo dữ liệu đáng tin cậy và chính xác, bằng cách bảo vệ dữ liệu khỏi việc bị sửa đổi hoặc phá hủy trái phép. Giả sử bạn chi 10 đô la mua một chiếc pizza — có thể bạn không bận tâm liệu giao dịch đó có được giữ bí mật hay không, nhưng chuyện gì sẽ xảy ra nếu hồ sơ giao dịch bị thay đổi và bạn phải trả tới 10.000 đô la? Tính toàn vẹn của một giao dịch có thể bị xâm phạm do cố ý (tấn công) hoặc vô ý (lỗi kỹ thuật, hoặc lỗi con người — VD: nhập sai số tiền thanh toán).

Để duy trì tính toàn vẹn, bạn cũng cần ngăn chặn những người không có thẩm quyền chỉnh sửa dữ liệu — ở khía cạnh này, tính toàn vẹn và tính bảo mật có sự giao thoa với nhau (một hệ thống kiểm soát quyền truy cập tốt sẽ phục vụ cả hai mục tiêu cùng lúc).

### 2.3. Availability (Tính sẵn sàng)

Tính sẵn sàng nghĩa là đảm bảo việc truy cập và sử dụng dữ liệu diễn ra kịp thời và đáng tin cậy. Ví dụ, bạn mong muốn truy cập được tài khoản ngân hàng trực tuyến 24/7 — để đáp ứng kỳ vọng đó, ngân hàng cần triển khai và duy trì đủ nguồn lực (server, hạ tầng mạng, cơ chế dự phòng) nhằm đảm bảo dịch vụ luôn sẵn sàng và ổn định.

Tuy nhiên, "kịp thời" không nhất thiết đồng nghĩa với "tức thì". Chẳng hạn, khi yêu cầu cấp bảng điểm, bạn có thể phải chờ vài ngày để nhà trường xử lý và gửi tài liệu; nếu cung cấp dưới dạng điện tử, họ còn có thể giới hạn khoảng thời gian bạn được phép truy cập. Dù vậy, miễn là dữ liệu được cung cấp trong một khoảng thời gian hợp lý, đáp ứng đúng nhu cầu sử dụng, thì tính sẵn sàng vẫn được đảm bảo.

Tính bảo mật, tính toàn vẹn và tính sẵn sàng là nền tảng của bảo mật dữ liệu. Bạn càng am hiểu về chúng, bạn càng có nhiều cơ hội thành công trong lĩnh vực an ninh mạng.

### 2.4. Ví dụ áp dụng CIA Triad vào tình huống thực tế

> **Câu hỏi kiểm tra kiến thức:**
> 1. *Bạn đi ăn cùng bạn bè và để lại 20 đô la tiền tip. Nhưng khi kiểm tra tài khoản ngân hàng, bạn phát hiện mình đã bị trừ tới 200 đô la.* → Vi phạm **Integrity** (số tiền ghi nhận sai lệch so với thực tế giao dịch).
> 2. *Một công ty quyết định triển khai xác thực hai yếu tố (2FA) để truy cập vào thông tin độc quyền (proprietary information) của họ.* → Ví dụ về việc bảo vệ **Confidentiality**.
> 3. *Bạn đang làm một dự án cần dùng phần mềm cộng tác trên nền tảng đám mây, toàn bộ thông tin dự án được lưu trên hệ thống đó. Server đột ngột gặp sự cố, khiến bạn không thể tiếp tục công việc.* → Vi phạm **Availability**.
>

---

## 3. 5 tham số bảo mật cốt lõi

Khi xây dựng một chương trình bảo mật dữ liệu, tổ chức cần trả lời rõ 5 câu hỏi sau đây cho **từng loại dữ liệu** mà mình đang nắm giữ. Đây không chỉ là câu hỏi lý thuyết — mỗi câu trả lời sẽ trực tiếp quyết định các biện pháp kiểm soát cụ thể (mục 4) mà tổ chức cần triển khai.

| Tham số | Câu hỏi | Ý nghĩa thực tế |
|---|---|---|
| **Sensitivity** (Độ nhạy cảm) | Dữ liệu nào là dữ liệu nhạy cảm? | Cần phân loại dữ liệu theo mức độ nhạy cảm (VD: Public, Confidential, PII, PHI — đã học ở tài liệu Data Privacy) trước khi quyết định mức bảo vệ cần thiết. Dữ liệu càng nhạy cảm, biện pháp kiểm soát càng phải chặt chẽ. |
| **Storage** (Lưu trữ) | Tổ chức nên lưu trữ dữ liệu nhạy cảm ở đâu? | Liên quan đến việc chọn hạ tầng lưu trữ (on-premises, cloud, database riêng biệt cho dữ liệu nhạy cảm...) và có cần mã hóa dữ liệu tại nơi lưu trữ (encryption at rest) hay không. |
| **Access** (Quyền truy cập) | Ai nên có quyền truy cập vào dữ liệu nhạy cảm? | Chính là câu hỏi đã trả lời chi tiết ở tài liệu User/Role/Permission — áp dụng RBAC và nguyên tắc Least Privilege. |
| **Flow** (Luồng di chuyển) | Dữ liệu nhạy cảm sẽ di chuyển đến đâu và bằng cách nào? | Dữ liệu có được gửi qua API nội bộ, chia sẻ với bên thứ ba, hay xuất ra file không? Mỗi điểm di chuyển là một điểm có nguy cơ rò rỉ cần được kiểm soát (liên quan trực tiếp đến khái niệm **data in motion** ở mục 5). |
| **Monitoring** (Giám sát) | Tổ chức nên giám sát việc truy cập và sử dụng dữ liệu nhạy cảm như thế nào? | Cần có cơ chế ghi log (audit log) mọi hành vi truy cập/chỉnh sửa dữ liệu nhạy cảm, để phát hiện sớm hành vi bất thường và phục vụ điều tra khi có sự cố. |

---

## 4. Các biện pháp kiểm soát bảo mật (Controls)

Sau khi xác định lỗ hổng bảo mật dựa trên mô hình CIA, bạn cần quyết định cách xử lý chúng. Các mối đe dọa đối với dữ liệu đến từ nhiều nguồn khác nhau, do đó thường cần triển khai đa dạng **biện pháp kiểm soát (controls)** — những hành động cụ thể giúp giảm thiểu rủi ro — để ứng phó với chúng. Biện pháp kiểm soát có thể thuộc một trong ba nhóm: hành chính, vật lý, hoặc kỹ thuật.

### 4.1. Administrative controls (Kiểm soát hành chính)

Là những hướng dẫn, chính sách và quy trình được xây dựng nhằm đáp ứng và thực thi mục tiêu bảo mật. Ví dụ: trường học hoặc nơi làm việc của bạn có thể áp dụng chính sách mật khẩu quy định các yêu cầu cụ thể (độ dài tối thiểu, loại ký tự bắt buộc).

### 4.2. Physical controls (Kiểm soát vật lý)

Là những thiết bị hoặc kết cấu được thiết kế để hạn chế quyền truy cập vào khu vực hoặc thiết bị chứa dữ liệu nhạy cảm — ví dụ: tòa nhà, văn phòng, phòng máy chủ. Bao gồm hàng rào, khóa, thẻ từ, camera an ninh, hệ thống báo động và tủ bảo mật.

### 4.3. Technical controls (Kiểm soát kỹ thuật)

Là các thành phần phần cứng hoặc phần mềm giúp bảo vệ dữ liệu hoặc quy trình. Một số ví dụ phổ biến:

- **Phần mềm chống mã độc** (anti-malware): phát hiện, cách ly và tiêu diệt các loại mã độc đe dọa dữ liệu hoặc mạng lưới. Windows Defender, XProtect, Malwarebytes là các ví dụ quen thuộc.
- **Phần mềm mã hóa** (encryption): chuyển đổi dữ liệu sang định dạng mà người không có thẩm quyền không thể đọc được, qua đó đảm bảo tính bảo mật. Ví dụ, nếu kẻ tấn công chặn được một email đã mã hóa, nội dung sẽ chỉ hiện ra dưới dạng các ký tự vô nghĩa.
- **Phần mềm sao lưu** (backup): tạo bản sao bổ sung của dữ liệu, cho phép khôi phục dữ liệu quan trọng bị mất do xâm nhập, lỗi hệ thống hoặc sự cố bảo mật khác.
- **Phần mềm xóa dữ liệu** (data erasure): loại bỏ vĩnh viễn dữ liệu không còn cần thiết. Lưu ý phân biệt: **xóa sạch (erasure)** khác với **xóa thông thường (deletion)** — dữ liệu bị xóa thông thường vẫn còn tồn tại trên thiết bị lưu trữ và có thể khôi phục được (phù hợp khi dữ liệu không nhạy cảm), trong khi dữ liệu bị "erasure" sẽ mất vĩnh viễn (phù hợp cho dữ liệu nhạy cảm) — phần mềm sẽ ghi đè lên dữ liệu cũ bằng chuỗi số 0/1 hoặc ký tự ngẫu nhiên để không ai có thể khôi phục được.

---

## 5. Ngăn ngừa mất mát dữ liệu (DLP)

Dù bạn sử dụng biện pháp kiểm soát nào để bảo vệ tính bảo mật của dữ liệu, mục tiêu vẫn không thay đổi: ngăn chặn mất mát dữ liệu.

**Mất dữ liệu (data loss)** là một cách gọi khác của vi phạm dữ liệu — tình trạng dữ liệu bị đánh cắp hoặc rò rỉ, khiến thông tin mật bị lộ cho người không có thẩm quyền. **Ngăn ngừa mất mát dữ liệu (Data Loss Prevention — DLP)** là khả năng của một tổ chức trong việc phát hiện và ngăn chặn tình trạng này.

### 5.1. Ba trạng thái của dữ liệu

Dữ liệu tồn tại ở ba trạng thái: tĩnh, đang di chuyển, hoặc đang được sử dụng. Một giải pháp DLP hiệu quả cần bảo vệ dữ liệu ở cả ba trạng thái.

![](https://simplyblock.io/assets/images/media/data-at-rest-data-in-transit-data-in-use.webp)

#### Dữ liệu tĩnh (Data at rest)

Là dữ liệu được lưu trữ trên các thiết bị như ổ cứng, ổ USB, đĩa DVD, máy chủ hoặc cơ sở dữ liệu. Khối lượng dữ liệu khổng lồ lưu trữ trên các thiết bị này khiến dữ liệu tĩnh trở thành mục tiêu đặc biệt hấp dẫn với kẻ tấn công.

#### Dữ liệu đang di chuyển (Data in motion)

Còn gọi là dữ liệu đang truyền tải — là dữ liệu đang được truyền đi trên mạng hoặc giữa các hệ thống. Khi bạn nhận email, xem phim trực tuyến hay tải xuống tệp tin, dữ liệu di chuyển qua Internet từ nguồn đến thiết bị của bạn. Nếu thiếu biện pháp bảo mật đầy đủ (VD: mã hóa TLS), kẻ tấn công có thể truy cập hoặc đánh chặn loại dữ liệu này.

#### Dữ liệu đang sử dụng (Data in use)

Là dữ liệu mà một số máy tính hoặc ứng dụng đang tích cực xử lý. Ví dụ: khi bạn mua hàng trên một trang web, hệ thống hậu trường xử lý thông tin thanh toán theo thời gian thực, có thể lưu dữ liệu tạm thời trong bộ nhớ (RAM) trong lúc xử lý. Nếu kẻ tấn công xâm nhập được hệ thống ngay thời điểm đó, chúng có thể truy cập vào dữ liệu này.

### 5.2. Các hệ thống DLP

**Hệ thống ngăn ngừa mất mát dữ liệu** bao gồm các quy trình, thủ tục và công cụ giúp phát hiện và ngăn chặn tình trạng mất mát dữ liệu — phát hiện và ghi lại mọi hoạt động truy cập hoặc chia sẻ đối với dữ liệu nhạy cảm về thời gian. Các hệ thống DLP có nhiều dạng khác nhau:

#### File-level DLP

![](https://www.scaler.com/topics/images/difference-between-file-and-folder-Thumbnail.webp)

Giúp xác định các tệp tin nhạy cảm trong hệ thống tệp. Chính sách bảo mật của tổ chức quy định các thuộc tính mà những tệp này cần có để đảm bảo an toàn — ví dụ: hồ sơ tài chính có độ nhạy cảm cao phải mang thuộc tính xác định rõ "không bao giờ được phép gửi qua email". Hệ thống DLP sẽ nhúng các thuộc tính này vào **metadata** (siêu dữ liệu — dữ liệu mô tả về chính tệp tin đó) của tệp, để đảm bảo chúng luôn đi kèm với tệp khi được di chuyển hoặc sao chép.

#### Network DLP

![](https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSfTASA4I5T8_RNLzfdyNklBvCdJvSykdMWYcd3AQj0_yTIbVas0ElQin4&s=10)

Bảo vệ dữ liệu ở trạng thái tĩnh, đang truyền tải hoặc đang sử dụng trên mạng lưới của tổ chức. Giải pháp này giám sát mọi hoạt động truyền tải dữ liệu (email, tin nhắn tức thời, tải xuống tệp) để phát hiện vi phạm chính sách bảo mật và quyền riêng tư. Ví dụ, Network DLP có thể ghi lại mỗi khi có người tải xuống bản thiết kế của một dịch vụ đang phát triển, thậm chí ngăn chặn việc tải xuống này với tất cả mọi người ngoại trừ một số nhân viên được chỉ định.

#### Cloud DLP

![](https://longvan.net/uploads/pngtree_cloud_computing_technology_cloud_data_center_with_hosting_server_cloud_service_png_image_13246564_87e742214c.png)

Là phần mở rộng của giải pháp DLP mạng, có chức năng bảo vệ dữ liệu được lưu trữ trong các kho lưu trữ đám mây. Giải pháp này phát hiện và mã hóa dữ liệu nhạy cảm trước khi chúng được lưu trữ trên đám mây.

#### Endpoint DLP

![](https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQxbFGaiqJO4_Ke4tJ0EOkXxWAEfSsDHR7pHPXO_UwBgA&s=10)

Giám sát tất cả các **điểm cuối (endpoint)** — thiết bị kết nối với mạng của tổ chức có thực hiện việc sử dụng, di chuyển hoặc lưu trữ dữ liệu (máy chủ, máy tính để bàn, máy tính xách tay) — để phát hiện tình trạng mất mát hoặc rò rỉ dữ liệu.

---

## 6. Cơ hội nghề nghiệp

Bạn có thể xây dựng cả một sự nghiệp xoay quanh việc đánh giá các chương trình bảo mật dữ liệu. Vai trò chính của **chuyên viên phân tích bảo mật dữ liệu (Data Security Analyst)** là bảo vệ khối lượng lớn dữ liệu của tổ chức khỏi các cuộc tấn công từ bên ngoài.

Nhiệm vụ cụ thể sẽ khác nhau tùy theo từng công ty, nhưng thông thường bao gồm:
- Kiểm tra chương trình bảo mật dữ liệu của công ty, sử dụng mô hình CIA làm cơ sở đánh giá.
- Thiết kế và lập tài liệu các chiến lược nhằm tăng cường bảo mật, dựa trên kết quả đánh giá.
- Phổ biến chính sách và quy trình bảo mật dữ liệu cho nhân viên.
- Giám sát việc tuân thủ các quy định bảo mật trong toàn công ty.

---

## 7. Tổng kết

- **Mô hình CIA Triad** (Confidentiality – Integrity – Availability) là nền tảng để đánh giá bất kỳ chương trình bảo mật dữ liệu nào — mỗi lỗ hổng, mỗi sự cố bảo mật đều có thể được nhìn nhận qua việc nó vi phạm mục tiêu nào trong ba mục tiêu này.
- Trước khi triển khai biện pháp bảo vệ, cần trả lời rõ **5 tham số bảo mật cốt lõi**: dữ liệu nào nhạy cảm, lưu ở đâu, ai được truy cập, di chuyển thế nào, và giám sát ra sao.
- Ba nhóm **biện pháp kiểm soát** (hành chính, vật lý, kỹ thuật) cần được kết hợp với nhau — không có nhóm nào là đủ nếu đứng riêng lẻ.
- **DLP** đòi hỏi bảo vệ dữ liệu ở cả ba trạng thái (tĩnh, đang di chuyển, đang sử dụng), thông qua nhiều lớp hệ thống khác nhau: File-level, Network, Cloud, Endpoint DLP.
- Đây là lĩnh vực có lộ trình sự nghiệp rõ ràng, với vai trò trung tâm là Data Security Analyst — người đánh giá và duy trì chương trình bảo mật dữ liệu của tổ chức dựa trên mô hình CIA.

---

### Mở rộng

- **Liên hệ với các tài liệu đã học:** mô hình CIA Triad chính là "khung lý thuyết" bao trùm lên các nội dung bạn đã học trước đó — nguyên tắc Least Privilege (User/Role/Permission) phục vụ trực tiếp cho **Confidentiality**; Transaction/ACID trong database phục vụ **Integrity**; kiến trúc phân tán và Erasure Coding của MinIO phục vụ **Availability**.
- **Encryption in transit vs encryption at rest vs encryption in use**: mở rộng trực tiếp từ khái niệm "ba trạng thái của dữ liệu" ở mục 5 — tìm hiểu thêm về **Confidential Computing** (mã hóa cả dữ liệu đang được xử lý trong bộ nhớ, giải quyết đúng lỗ hổng của "data in use" đã nêu).
- **SIEM (Security Information and Event Management)**: công cụ tổng hợp và phân tích log giám sát bảo mật ở quy mô toàn tổ chức — mở rộng tự nhiên từ tham số "Monitoring" ở mục 3.
- **Zero Trust Architecture**: mô hình kiểm soát truy cập hiện đại, không mặc định tin tưởng bất kỳ ai — liên quan trực tiếp đến tham số "Access" ở mục 3 và các biện pháp kiểm soát kỹ thuật ở mục 4.
- **Áp dụng vào backend Go**: khi thiết kế một service, hãy tự hỏi 5 tham số ở mục 3 cho từng bảng dữ liệu trong schema — đây là bài tập thực hành tốt để chuyển từ lý thuyết CIA Triad sang quyết định thiết kế cụ thể (VD: cột nào cần mã hóa ở tầng ứng dụng, endpoint nào cần thêm rate-limiting để bảo vệ Availability, log truy cập nào cần bật để phục vụ Monitoring).
