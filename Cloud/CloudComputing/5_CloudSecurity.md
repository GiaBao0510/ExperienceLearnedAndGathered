# An Toàn Cho Điện Toán Đám Mây

---

## 1. Các Thuật Ngữ và Khái Niệm Cơ Bản

**An toàn (Security)** là một tập hợp phức tạp các kỹ thuật, công nghệ, quy tắc và hành vi được phối hợp với nhau nhằm đảm bảo **tính toàn vẹn trong việc truy cập** vào các hệ thống máy tính và dữ liệu.

Các biện pháp an toàn CNTT nhằm mục đích **phòng chống các mối đe dọa an ninh**, bao gồm cả các hành động xấu có chủ ý lẫn các lỗi vô ý của người dùng.

---

### Các thuộc tính an toàn cốt lõi

**Tính cơ mật (Confidentiality):** Là đặc tính đảm bảo chỉ những người được cấp phép mới có thể truy cập dữ liệu. Trong môi trường đám mây, tính cơ mật chủ yếu gắn liền với việc **hạn chế truy cập vào dữ liệu truyền tải và lưu trữ**.

**Tính toàn vẹn (Integrity):** Là đặc tính đảm bảo dữ liệu **không bị sửa đổi bởi những người không có thẩm quyền**. Trong môi trường đám mây, thuộc tính này thể hiện ở chỗ dữ liệu được truyền bởi người tiêu dùng có đảm bảo trùng khớp với dữ liệu nhận được bởi dịch vụ đám mây hay không. Tính toàn vẹn mở rộng đến cả cách thức dữ liệu được lưu trữ, xử lý và truy vấn bởi các dịch vụ đám mây.

> **Làm thế nào để biết người dùng nào có hoặc không có thẩm quyền?** Đây là câu hỏi trọng tâm của an ninh thông tin. Hệ thống xác định thẩm quyền thông qua hai bước:
> 
> - **Xác thực (Authentication) – "Bạn là ai?":** Hệ thống xác minh danh tính người dùng, thường qua mật khẩu, mã OTP, sinh trắc học, hoặc chứng chỉ số. Chỉ khi xác thực thành công, người dùng mới được hệ thống "nhận ra".
> - **Phân quyền (Authorization) – "Bạn được phép làm gì?":** Sau khi biết bạn là ai, hệ thống kiểm tra danh sách quyền hạn (Access Control List / Role-Based Access Control) để quyết định bạn được truy cập tài nguyên nào. Ví dụ: nhân viên kế toán chỉ được đọc dữ liệu tài chính, không được xóa; quản trị viên mới có quyền xóa.
> - Tóm lại: **Người có thẩm quyền** là người đã xác thực thành công VÀ có quyền hạn tương ứng với hành động họ muốn thực hiện trên tài nguyên đó.

**Tính xác thực (Authenticity):** Là đặc tính đảm bảo một thực thể (dữ liệu, dịch vụ, người dùng) thực sự đến từ **nguồn đã được ủy quyền**, không bị giả mạo.

**Tính sẵn dùng (Availability):** Là đặc tính đảm bảo hệ thống **có thể truy cập và sử dụng được** trong khoảng thời gian yêu cầu. Trong môi trường đám mây, tính sẵn dùng của dịch vụ có thể là trách nhiệm được chia sẻ giữa **nhà cung cấp đám mây** và **nhà cung cấp truy cập đám mây (cloud carrier)**.

---

### Các khái niệm về rủi ro và biện pháp ứng phó

**Mối đe dọa (Threat):** Là một sự xâm phạm an ninh tiềm tàng có thể vượt qua các tuyến phòng thủ nhằm vi phạm quyền riêng tư hoặc gây tổn hại cho hệ thống.

> **Các mối đe dọa thường gặp đối với dữ liệu lưu trữ trên đám mây:**
> 
> - **Rò rỉ dữ liệu (Data Breach):** Kẻ tấn công xâm nhập và đánh cắp dữ liệu nhạy cảm như thông tin cá nhân, tài khoản ngân hàng, bí mật kinh doanh.
> - **Mất dữ liệu (Data Loss):** Dữ liệu bị xóa hoặc bị mã hóa bởi ransomware (mã độc tống tiền), không thể khôi phục.
> - **Truy cập trái phép (Unauthorized Access):** Kẻ xấu đăng nhập vào tài khoản người dùng thông qua mật khẩu yếu, phishing hoặc lỗ hổng bảo mật.
> - **Tấn công từ chối dịch vụ (DoS/DDoS):** Làm cho hệ thống lưu trữ quá tải, không thể phục vụ người dùng hợp pháp.
> - **Nội gian (Insider Threat):** Nhân viên nội bộ cố tình đánh cắp hoặc phá hoại dữ liệu.
> - **Tấn công chuỗi cung ứng (Supply Chain Attack):** Kẻ tấn công xâm nhập qua phần mềm hoặc thư viện bên thứ ba mà tổ chức đang sử dụng.

**Tính dễ bị tổn thương (Vulnerability):** Là một điểm yếu trong hệ thống có thể bị khai thác vì các biện pháp bảo vệ không đủ mạnh hoặc đã mất tác dụng trước một cuộc tấn công. Nguyên nhân có thể đến từ: cấu hình sai, chính sách an ninh yếu kém, lỗi người dùng, hỏng hóc phần cứng/firmware, lỗi phần mềm, hoặc kiến trúc an ninh không tốt.

**Rủi ro (Risk):** Là khả năng bị mất mát hoặc tổn hại khi thực hiện một hoạt động. Rủi ro thường được đo dựa trên **mức độ nguy hiểm của mối đe dọa** và **số lượng điểm dễ bị tổn thương** đang tồn tại.

> **Công thức đơn giản để hiểu:** `Rủi ro = Mối đe dọa × Điểm dễ tổn thương`. Giảm một trong hai yếu tố này sẽ giảm rủi ro tổng thể.

**Kiểm soát an ninh (Security Control):** Là các biện pháp đối phó để ngăn ngừa hoặc phản ứng lại các mối đe dọa nhằm giảm thiểu rủi ro. Chi tiết thường được quy định trong các **chính sách an ninh** – là tập hợp các luật và quy định mô tả cách thiết lập, vận hành và bảo vệ hệ thống.

**Cơ chế an ninh (Security Mechanisms):** Là các thành phần cụ thể được triển khai trong hệ thống phòng thủ để bảo vệ tài nguyên CNTT, thông tin và dịch vụ. Ví dụ: tường lửa, mã hóa, xác thực hai bước.

**Chính sách an ninh (Security Policies):** Là tập hợp các luật và quy tắc an ninh định nghĩa trước cách thức các luật và quy tắc đó được triển khai và áp dụng trong thực tế.

---

## 2. Các Tác Nhân Đe Dọa (Threat Agents)

**Tác nhân đe dọa (Threat Agent)** là một thực thể (con người hoặc phần mềm) có khả năng phát động một cuộc tấn công, từ đó tạo ra mối đe dọa cho hệ thống.

Các mối đe dọa an ninh đám mây có thể có nguồn gốc từ **bên trong hoặc bên ngoài**, từ **con người hoặc phần mềm**.

---

### Kẻ Tấn Công Ẩn Danh (Anonymous Attacker)

Là người dùng dịch vụ đám mây **không được tin tưởng và không được cấp phép**. Thường tồn tại dưới dạng một chương trình phần mềm bên ngoài khởi động các cuộc tấn công qua mạng công cộng.

Vì có ít thông tin về chính sách bảo mật và hệ thống phòng thủ bên trong, kẻ tấn công ẩn danh thường phải dùng các kỹ thuật như:

- **Chiếm đoạt tài khoản** người dùng hợp pháp.
- **Đánh cắp thông tin xác thực (credentials)** qua phishing, malware.

Bằng cách này, chúng vừa ẩn danh vừa tránh để lại dấu vết có thể bị truy tố.

> **Ví dụ thực tế:** Một bot tự động dò mật khẩu (brute-force attack) từ hàng nghìn địa chỉ IP khác nhau trên Internet để xâm nhập vào tài khoản người dùng đám mây.

---

### Tác Nhân Dịch Vụ Độc Hại (Malicious Service Agent)

Là tác nhân có khả năng **đánh chặn và chuyển tiếp lưu lượng mạng** đang đổ về một đám mây. Có thể là:

- Các tác nhân dịch vụ bị **can thiệp hoặc "nhiễm độc"** từ bên trong.
- Phần mềm bên ngoài đám mây có khả năng **đánh chặn từ xa và làm sai lệch nội dung** các thông điệp truyền qua mạng.

> **Ví dụ thực tế:** Một phần mềm độc hại (malware) được cài vào máy chủ trung gian, âm thầm sao chép mọi dữ liệu đi qua trước khi chuyển tiếp bình thường – người dùng không hay biết.

---

### Kẻ Tấn Công Được Tin Tưởng (Trusted Attacker)

Là kẻ **chia sẻ tài nguyên CNTT trong cùng một môi trường đám mây** với các người tiêu dùng hợp pháp, và cố gắng khai thác các thông tin xác thực hợp lệ để tấn công nhà cung cấp đám mây hoặc các khách hàng khác.

Đặc điểm nguy hiểm: họ tấn công **từ bên trong đường biên tin tưởng (trust boundary)** của đám mây bằng cách:

- Lạm dụng quyền hạn của thông tin xác thực hợp lệ.
- Chiếm đoạt thông tin nhạy cảm hoặc bí mật.
- Khai thác các quy trình xác thực yếu, bẻ khóa, v.v.

> **Ví dụ thực tế:** Một công ty thuê máy chủ ảo trên cùng hạ tầng vật lý với công ty đối thủ. Nếu họ khai thác lỗ hổng trong lớp ảo hóa, họ có thể đọc dữ liệu bộ nhớ của máy ảo bên cạnh – đây là dạng tấn công **VM escape** hay **side-channel attack**.

---

### Nội Gian (Malicious Insider)

Là các **tác nhân độc hại là con người, hành động từ bên trong** tổ chức nhà cung cấp đám mây. Họ thường là:

- Nhân viên **đang làm việc hoặc đã nghỉ việc**.
- Bên thứ ba có quyền truy cập vào tài sản của nhà cung cấp.

Đây là tác nhân đe dọa **cực kỳ nguy hiểm** vì họ đã có quyền truy cập hợp pháp vào hệ thống và hiểu rõ cơ chế bảo mật bên trong.

> **Ví dụ thực tế:** Một nhân viên quản trị hệ thống bất mãn sao chép toàn bộ cơ sở dữ liệu khách hàng trước khi nghỉ việc, hoặc cố tình xóa dữ liệu quan trọng.

---

## 3. Những Mối Đe Dọa Cho An Ninh Đám Mây

### Nghe Trộm Thông Tin (Traffic Eavesdropping)

Xảy ra khi dữ liệu đang truyền tải **bị chặn một cách thụ động** bởi tác nhân dịch vụ độc hại nhằm thu thập thông tin bất hợp pháp. Mục tiêu là đánh cắp dữ liệu bí mật và có thể là thông tin về mối quan hệ giữa người tiêu dùng và nhà cung cấp đám mây.

> **Đặc điểm:** Kiểu tấn công này là **thụ động** – kẻ tấn công không can thiệp vào luồng dữ liệu, chỉ lắng nghe. Điều này khiến nó rất khó phát hiện. **Phòng chống:** Mã hóa dữ liệu khi truyền tải (TLS/HTTPS) là biện pháp chủ yếu – dù bị nghe trộm, dữ liệu vẫn không đọc được.

---

### Trung Gian Độc Hại (Malicious Intermediary)

Xảy ra khi các thông điệp bị **chặn và thay đổi nội dung** bởi tác nhân dịch vụ độc hại, tạo ra vi phạm về **tính cơ mật và tính toàn vẹn** của dữ liệu. Kẻ tấn công cũng có thể **chèn thêm dữ liệu độc hại** vào thông điệp trước khi chuyển tiếp đến nơi nhận.

> **Đặc điểm:** Khác với nghe trộm (thụ động), kiểu tấn công này là **chủ động** – thông điệp bị sửa đổi. **Ví dụ:** Tấn công **Man-in-the-Middle (MitM)** – kẻ tấn công ngồi giữa người dùng và dịch vụ, âm thầm thay đổi nội dung giao dịch (ví dụ: đổi số tài khoản ngân hàng trong lệnh chuyển tiền).

---

### Từ Chối Dịch Vụ (Denial of Service – DoS / DDoS)

Mục đích là **làm quá tải tài nguyên CNTT** để chúng không thể vận hành đúng đắn. Các cách triển khai tấn công:

- **Giả mạo yêu cầu hàng loạt:** Gửi hàng triệu thông điệp giả bắt chước yêu cầu hợp lệ để làm tắc nghẽn hệ thống.
- **Làm quá tải mạng:** Băng thông mạng bị chiếm hết, làm giảm khả năng phản hồi và ảnh hưởng tiêu cực đến hiệu năng.
- **Khai thác tài nguyên hệ thống:** Mỗi yêu cầu được thiết kế để tiêu thụ tối đa bộ nhớ và CPU, làm cạn kiệt tài nguyên nhanh chóng.

> **DDoS (Distributed DoS)** là biến thể nguy hiểm hơn – tấn công đến từ hàng nghìn thiết bị bị nhiễm malware (botnet) cùng lúc, rất khó chặn vì không có một nguồn tấn công duy nhất.

---

### Ủy Quyền Không Phù Hợp (Insufficient Authorization)

Xảy ra khi quyền truy cập được gán **sai hoặc quá rộng** cho một người dùng (hoặc kẻ tấn công), khiến họ có thể truy cập vào các tài nguyên CNTT vốn được bảo vệ.

**Ví dụ 1 – Thiếu kiểm soát truy cập:** Người tiêu dùng dịch vụ đám mây A giành được quyền truy cập trực tiếp vào cơ sở dữ liệu – vốn được thiết kế để chỉ truy cập thông qua một web service có hợp đồng dịch vụ công bố (như người tiêu dùng B). Người tiêu dùng A đã bỏ qua tầng bảo vệ trung gian.

**Ví dụ 2 – Mật khẩu yếu:** Kẻ tấn công bẻ khóa mật khẩu đơn giản của người tiêu dùng dịch vụ A. Sau đó dùng thông tin xác thực đánh cắp được để đăng nhập và truy cập vào máy chủ ảo của người dùng A – bao gồm toàn bộ dữ liệu bên trong.

> **Nguyên nhân phổ biến:** Mật khẩu đặt đơn giản (123456, admin), dùng chung tài khoản cho nhiều người, không áp dụng nguyên tắc **least privilege** (chỉ cấp đúng quyền cần thiết, không hơn).

---

### Tấn Công Ảo Hóa (Virtualization Attack)

Vì nhà cung cấp đám mây gán quyền quản trị cho người tiêu dùng để truy cập vào tài nguyên ảo hóa (như máy chủ ảo), người tiêu dùng đám mây **có thể lạm dụng quyền này để tấn công vào tài nguyên vật lý phía dưới**.

Loại tấn công này khai thác các **điểm yếu trong nền tảng ảo hóa (Hypervisor)** để:

- Thoát khỏi máy ảo và truy cập vào máy chủ vật lý (**VM Escape**).
- Đọc bộ nhớ của máy ảo khác đang chạy trên cùng máy chủ vật lý (**Side-channel Attack**).
- Làm tổn hại đến tính riêng tư, toàn vẹn và khả năng sẵn dùng của toàn bộ nền tảng.

> **Mức độ nguy hiểm cao:** Vì trên một nền tảng ảo hóa có nhiều khách hàng khác nhau, một cuộc tấn công thành công có **tác động lan rộng** – ảnh hưởng đến tất cả người dùng cùng chia sẻ máy chủ vật lý đó.

---

### Chồng Lấn Vùng Tin Tưởng (Overlapping Trust Boundaries)

Trong môi trường đám mây công cộng (public cloud), **nhiều người tiêu dùng chia sẻ cùng một hạ tầng vật lý**. Điều này tạo ra vùng tin tưởng bị chồng lấn giữa những người dùng không liên quan đến nhau.

**Cơ chế tấn công:**

- Người tiêu dùng độc hại **nhắm vào các tài nguyên CNTT trong cùng vùng tin tưởng** mà họ đang chia sẻ.
- Kết quả: một phần hoặc toàn bộ người tiêu dùng khác trong cùng vùng tin tưởng đó bị ảnh hưởng.
- Kẻ tấn công có thể dùng tài nguyên ảo thu được để **tiếp tục tấn công sang các tài nguyên khác** trong cùng vùng tin tưởng (tấn công leo thang – lateral movement).

> **Ví dụ thực tế:** Hình dung một tòa nhà văn phòng (máy chủ vật lý) có nhiều công ty thuê (máy ảo). Nếu một công ty độc hại tìm được cách phá khóa hành lang chung, họ có thể tiếp cận cửa phòng của các công ty khác – dù mỗi công ty có khóa riêng.

> **Phòng chống:** Nhà cung cấp đám mây cần đảm bảo **cô lập mạnh (strong isolation)** giữa các máy ảo, thường xuyên vá lỗi Hypervisor và sử dụng mạng riêng ảo (VPC/VLAN) để phân tách vùng tin tưởng.