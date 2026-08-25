# Đánh giá Cuối khoá — Câu hỏi và Giải thích (Final Assessment)

---

#### Câu 1. Bạn truy cập một trang web bằng cùng URL quen thuộc. Tuy nhiên, trang web trông khác đi và URL trong thanh địa chỉ có vẻ hơi khác so với bình thường.

Tình huống này thuộc loại tấn công mạng nào?

- Tấn công từ chối dịch vụ (DoS)
- Tấn công chèn SQL (SQL Injection)
- **Tấn công hệ thống tên miền (DNS Attack)** [Đúng]
- Tấn công người trung gian (MitM)

**Giải thích:** Tình huống này mô tả một cuộc tấn công DNS — cụ thể là **DNS spoofing (giả mạo DNS)** hoặc **DNS hijacking (chiếm đoạt DNS)**. Kẻ tấn công thao túng quá trình phân giải tên miền, chuyển hướng người dùng đến một trang web giả mạo có giao diện tương tự trang gốc, trong khi URL trong thanh địa chỉ hiển thị khác biệt nhỏ (thường là tên miền giả mạo gần giống tên miền thật).

> **Tại sao không phải MitM?** Tấn công MitM (Man-in-the-Middle) cũng có thể chặn và chuyển hướng traffic, nhưng trong tình huống này, dấu hiệu rõ ràng nhất là **URL thay đổi** — đặc trưng của DNS attack. Trong MitM thuần túy, URL thường không thay đổi vì kẻ tấn công chỉ đứng giữa giao tiếp, không thay đổi tên miền được hiển thị.

---

#### Câu 2. Gracie là chuyên gia an ninh mạng tại công ty. Một trong những trách nhiệm của cô là định kỳ chạy quét phát hiện phiên bản (version detection scan) trên các thiết bị của công ty.

Tại sao Gracie lại thực hiện quét phát hiện phiên bản?

- **Để phát hiện phần mềm lỗi thời đang chạy trên các thiết bị.** [Đúng]
- Để lập danh mục tất cả các thiết bị mạng.
- Để chặn liên lạc giữa các máy chủ.
- Để tìm thông tin đăng nhập quản trị viên được mã hóa trong phần mềm.

**Giải thích:** Quét phát hiện phiên bản xác định **số phiên bản cụ thể của phần mềm và dịch vụ** đang chạy trên các thiết bị (ví dụ: Apache 2.4.49, OpenSSH 7.2, Windows Server 2016). Thông tin này cho phép đối chiếu với cơ sở dữ liệu CVE (Common Vulnerabilities and Exposures) để xác định phiên bản nào có lỗ hổng đã biết. Gracie thực hiện quét này định kỳ để phát hiện và đánh dấu phần mềm lỗi thời cần cập nhật bản vá trước khi kẻ tấn công có thể khai thác.

---

#### Câu 3. Giả sử bạn đang sử dụng Zenmap để quét một máy chủ.

Bạn cần thông tin gì về máy chủ để bắt đầu quá trình quét?

- Server software
- **Web address (địa chỉ web / địa chỉ IP)** [Đúng]
- Operating system
- Port number

**Giải thích:** Để quét một máy chủ bằng Zenmap (giao diện đồ họa của Nmap), thông tin tối thiểu bắt buộc là **địa chỉ mục tiêu** — có thể là địa chỉ IP (ví dụ: `192.168.1.1`) hoặc tên miền (ví dụ: `example.com`). Zenmap sẽ dùng địa chỉ này làm điểm khởi đầu để thực hiện toàn bộ quy trình quét; các thông tin khác như hệ điều hành, dịch vụ, và số cổng là **kết quả** của quá trình quét, không phải đầu vào.

---

#### Câu 4. Krysia là chuyên gia đánh giá lỗ hổng bảo mật. Cô thực hiện quét cổng trên một thiết bị mạng và phát hiện cổng TCP 80 đang mở.

Krysia nên kết luận điều gì từ phát hiện này?

- Chức năng chia sẻ tệp Windows (Windows File Sharing) đang hoạt động trên thiết bị.
- Ai đó đang điều khiển thiết bị từ xa.
- Mạng dễ bị tấn công.
- **Có thể đang sử dụng một ứng dụng web hoặc máy chủ HTTP.** [Đúng]

**Giải thích:** Cổng TCP 80 là cổng tiêu chuẩn dành riêng cho giao thức **HTTP (HyperText Transfer Protocol)**. Khi cổng này đang mở, có khả năng cao thiết bị đang chạy một **web server hoặc ứng dụng web**. Đây là kết luận hợp lý nhất dựa trên thông tin duy nhất có được — một cổng mở không tự động có nghĩa là mạng dễ bị tấn công; cần điều tra thêm về phần mềm và phiên bản đang chạy.

> **Tham chiếu:** Cổng 445 (SMB/Windows File Sharing), cổng 3389 (RDP/điều khiển từ xa), và cổng 80 (HTTP) là các cổng có ý nghĩa khác nhau hoàn toàn.

---

#### Câu 5. Bạn là thành viên nhóm an ninh mạng và nhận thấy một số lần đăng nhập không thành công vào máy chủ web.

Bạn nên làm gì để hiểu rõ hơn về những lần đăng nhập này?

- Thực hiện một số bài kiểm tra ping.
- **Thực hiện quét lỗ hổng bảo mật (vulnerability scanning).** [Đúng]
- Chạy lệnh tracert.
- Sử dụng công cụ tìm kiếm Shodan.

**Giải thích:** Các lần đăng nhập thất bại liên tiếp là dấu hiệu của cuộc tấn công **brute force** hoặc **credential stuffing** đang được thực hiện nhằm khai thác điểm yếu trong cơ chế xác thực. **Quét lỗ hổng bảo mật** giúp xác định các điểm yếu trong hệ thống mà kẻ tấn công đang cố gắng khai thác — từ đó có thể vá lỗ hổng trước khi cuộc tấn công thành công.

> **Tại sao không phải Shodan?** Shodan là công cụ OSINT thu thập thông tin về thiết bị từ bên ngoài. Trong tình huống này, bạn cần **điều tra từ bên trong hệ thống của mình**, không phải nhìn từ ngoài vào.

---

#### Câu 6. Sabine là nhà báo điều tra đang lên kế hoạch viết bài về một chính trị gia địa phương. Cô dự định dùng công cụ phân tích tìm kiếm mối liên hệ giữa các tập dữ liệu thu thập được.

Hướng dẫn nào về thu thập thông tin nguồn mở (OSINT) sẽ giúp công cụ phân tích của Sabine hoạt động tốt hơn?

- Trước tiên hãy xem xét các cuộc thảo luận trên diễn đàn có liên quan.
- **Thu thập một lượng lớn thông tin.** [Đúng]
- Chỉ dựa vào thông tin có thể xác minh.
- Chỉ sử dụng các nguồn trên internet khi có thể.

**Giải thích:** Các công cụ phân tích tìm kiếm mối liên hệ (link analysis tools) hoạt động hiệu quả hơn khi có **nhiều dữ liệu đầu vào hơn**, vì thuật toán có nhiều điểm kết nối tiềm năng để phân tích. Trong nghiên cứu OSINT, nguyên tắc "thu thập rộng trước, lọc sau" được ưu tiên — bởi vì người thu thập không biết trước mảnh thông tin nào sẽ trở thành chìa khóa quan trọng, đặc biệt khi thông tin được ghép lại từ nhiều nguồn không rõ ràng.

---

#### Câu 7. Mustafa — chuyên gia tư vấn an ninh mạng — đang phân tích vụ tấn công vào DataVista Technologies bằng khung Cyber Kill Chain®. Kẻ tấn công đã: (1) dò tìm lỗ hổng, (2) chọn phần mềm độc hại, (3) gửi malware qua email cho nhân viên. Ở bước thứ tư, phần mềm độc hại được kích hoạt.

Sự kiện nào tương ứng với bước thứ tư — Khai thác (Exploitation)?

- Kẻ tấn công dùng máy chủ ẩn để liên lạc với hệ thống bị xâm phạm.
- Kẻ tấn công làm gián đoạn hệ thống quản lý kho hàng của công ty.
- Malware tạo backdoor để duy trì quyền truy cập ngay cả khi lỗ hổng ban đầu được vá.
- **Một nhân viên nhấp vào liên kết đáng ngờ trong email, cho phép malware truy cập vào hệ thống.** [Đúng]

**Giải thích:** Trong Cyber Kill Chain®, bảy bước lần lượt là: (1) Trinh sát, (2) Vũ khí hóa, (3) Phân phối, **(4) Khai thác**, (5) Cài đặt, (6) C2, (7) Hành động trên mục tiêu.

Bước **Khai thác (Exploitation)** là khoảnh khắc **mục tiêu kích hoạt** phần mềm độc hại — thường bằng cách mở tệp đính kèm hoặc nhấp vào liên kết độc hại trong email. Đây chính xác là sự kiện được mô tả. Các sự kiện khác tương ứng với bước 5 (backdoor = Cài đặt), bước 6 (máy chủ ẩn = C2), và bước 7 (gián đoạn kho hàng = Hành động trên mục tiêu).

> **Lưu ý:** Bản gốc viết "khung Cyber Kill" — tên đầy đủ và chính xác là **Cyber Kill Chain®** (nhãn hiệu của Lockheed Martin). Đã sửa.

---

#### Câu 8. Bạn đang giúp công ty phát triển chính sách an ninh mạng mới và tóm tắt vụ vi phạm dữ liệu Cash App Investing năm 2021–2022.

Bài học chính nào từ vụ vi phạm này cần được nhấn mạnh?

- **Sự cần thiết của chính sách kiểm soát truy cập nghiêm ngặt đối với nhân viên bị sa thải.** [Đúng]
- Sự cần thiết phải cập nhật phần mềm hệ thống và mạng thường xuyên.
- Sự cần thiết phải chống lại việc trả tiền chuộc trong các vụ tấn công ransomware.
- Sự cần thiết phải hạn chế khả năng của bên thứ ba trong chuỗi cung ứng.

**Giải thích:** Trong vụ vi phạm này, một **cựu nhân viên** (đã bị chấm dứt hợp đồng) đã lợi dụng quyền truy cập chưa bị thu hồi để tải xuống dữ liệu của hơn 8 triệu người dùng — và hành vi này không bị phát hiện trong **hơn bốn tháng**. Bài học cốt lõi: quyền truy cập của nhân viên nghỉ việc phải được **thu hồi ngay lập tức** trong quy trình offboarding, không phải sau vài ngày hay vài tuần. Các lựa chọn còn lại phù hợp hơn với các vụ tấn công khác: cập nhật phần mềm (liên quan đến exploit lỗ hổng kỹ thuật), từ chối trả tiền chuộc (liên quan đến ransomware như LAUSD), và kiểm soát chuỗi cung ứng (liên quan đến SolarWinds).

---

#### Câu 9. Một người lạ mặc đồng phục công sở đang cố vào tòa nhà bạn làm việc. Họ nói ai đó từ trụ sở chính cử đến nhưng quên mang thẻ ra vào.

Bạn nên làm gì?

- Từ chối cho họ vào và gọi cảnh sát địa phương.
- **Yêu cầu họ đợi bên ngoài trong khi nhân viên bảo vệ xác minh câu chuyện.** [Đúng]
- Phớt lờ họ và tiếp tục công việc.
- Cho họ vào vì đồng phục trông hợp lệ và họ đang vội.

**Giải thích:** Tình huống này mô tả kỹ thuật **pretexting** trong social engineering — kẻ tấn công dựng lên một kịch bản ("được cử từ trụ sở chính") để khiến nạn nhân tin tưởng và hành động mà không cần kiểm tra. Đồng phục và cảm giác khẩn cấp ("đang vội") là những yếu tố được sử dụng để giảm sự nghi ngờ.

Hành động đúng là **yêu cầu họ đợi và để bảo vệ xác minh** — không tự mình cho vào và cũng không gọi cảnh sát ngay (vì có thể là nhầm lẫn thực sự). Việc xác minh là biện pháp cân bằng giữa lịch sự và bảo mật.

---

#### Câu 10. Khách hàng hỏi bạn — chuyên gia tư vấn an ninh mạng — tại sao nhiều tội phạm mạng sử dụng tiền điện tử (cryptocurrency) cho các hoạt động bất hợp pháp.

Bạn sẽ trả lời như thế nào?

- Các tổ chức quốc tế chấp nhận tiền điện tử.
- Tội phạm có thể đánh cắp tiền điện tử dễ dàng hơn.
- Người dùng có thể dễ dàng theo dõi các giao dịch tiền điện tử.
- **Tiền điện tử cung cấp tính ẩn danh và khó truy vết.** [Đúng]

**Giải thích:** Tiền điện tử — đặc biệt là các **privacy coin** như Monero — cung cấp mức độ ẩn danh cao và khó truy vết hơn so với hệ thống ngân hàng truyền thống. Các giao dịch không cần gắn với danh tính thực tế (không có KYC bắt buộc trong nhiều trường hợp), và các kỹ thuật như **cryptocurrency mixing/tumbling** (trộn tiền mã hóa) càng làm phức tạp thêm việc truy vết. Đây là lý do tội phạm mạng — đặc biệt là nhóm ransomware — yêu cầu thanh toán bằng tiền điện tử.

> **Lưu ý kỹ thuật:** Không phải tất cả tiền điện tử đều ẩn danh như nhau. Bitcoin có blockchain công khai có thể phân tích — các cơ quan thực thi pháp luật đã thu hồi nhiều khoản tiền Bitcoin. Monero có cơ chế ẩn danh mạnh hơn đáng kể.

---

#### Câu 11. Anders lãnh đạo nhóm an ninh mạng của SecureNest Investments. Anh nhận được cảnh báo tình báo rằng một nhóm hacktivist có thể tấn công tổ chức để trả đũa phát ngôn chính trị của CEO.

Chiến lược phù hợp nhất Anders nên đề xuất là gì?

- Thành lập nhóm đặc nhiệm xem xét chính sách mã hóa dữ liệu của công ty.
- Triển khai các biện pháp phòng thủ phối hợp toàn diện trên mọi khía cạnh.
- **Xác nhận rằng hệ thống phòng thủ có thể đối phó với một cuộc tấn công gián đoạn kéo dài.** [Đúng]
- Xem xét hồ sơ mạng xã hội nhân viên để tìm liên kết với nhóm hacktivist.

**Giải thích:** Hacktivist thường phối hợp **hàng trăm thành viên** để phát động các cuộc tấn công **DDoS (Distributed Denial of Service)** kéo dài và song song, nhằm làm gián đoạn dịch vụ và tạo sức ép. Mối đe dọa chính từ hacktivist không phải là xâm nhập dữ liệu mà là **gián đoạn hoạt động kéo dài**. Do đó, biện pháp ưu tiên là đảm bảo hệ thống có khả năng chịu đựng (resilience) trước tấn công liên tục — bao gồm DDoS mitigation, redundancy, và kế hoạch business continuity.

> **Giải thích phần tiếng Anh trong bản gốc:** Phần giải thích câu 11 trong bản gốc được viết hoàn toàn bằng tiếng Anh trong khi toàn bộ tài liệu dùng tiếng Việt — đây là lỗi nhất quán ngôn ngữ. Đã dịch và tích hợp vào phần giải thích tiếng Việt.

---

#### Câu 12. Hãy tưởng tượng ransomware lây nhiễm vào máy chủ của công ty bạn.

Kẻ tấn công có thể sẽ yêu cầu gì để khôi phục lại các tệp?

- Lời xin lỗi công khai về hành vi sai trái bị cáo buộc của CEO.
- **Một khoản tiền chuộc cụ thể (thường bằng tiền điện tử).** [Đúng]
- Quyền truy cập không hạn chế vào cơ sở dữ liệu của công ty.
- Tài sản vật chất thuộc về công ty.

**Giải thích:** Ransomware mã hóa dữ liệu của nạn nhân và yêu cầu **thanh toán tiền chuộc** — thường bằng Bitcoin hoặc Monero — để đổi lấy khóa giải mã. Đây là mô hình kinh doanh cốt lõi của ransomware. Kẻ tấn công không yêu cầu quyền truy cập thêm (vì họ đã xâm nhập rồi) hay tài sản vật chất (khó thu thập và dễ truy vết).

> Lưu ý: Nhiều nhóm ransomware hiện áp dụng **double extortion** — vừa mã hóa dữ liệu vừa đe dọa công bố dữ liệu đánh cắp nếu không trả tiền. Ví dụ: vụ LAUSD (câu hỏi 8 trong bài học trước).

---

#### Câu 13. Rhonda — chuyên gia an ninh mạng — nhận cảnh báo về xâm nhập trái phép và đang dùng ma trận MITRE ATT&CK để dự đoán kỹ thuật tiếp theo của kẻ tấn công.

Trong số các mục được liệt kê trong ma trận MITRE ATT&CK dưới đây, mục nào là **kỹ thuật** mà kẻ tấn công có thể sử dụng?

_(Chọn tất cả đáp án đúng)_

![](https://cpcontents.adobe.com/fr/dynamic-protected/4ff69c5a0e4e44388dedf2467d23d464/protected/account/2135/resources/7868757/7868757/content/scormcontent/assets/Quiz_ATT%26CK%20Matrix%20Excerpt.png)

- **Thao túng tài khoản (Account Manipulation)** [Đúng — đây là kỹ thuật]
- **Né tránh trình gỡ lỗi (Debugger Evasion)** [Đúng — đây là kỹ thuật]
- Né tránh phòng thủ (Defense Evasion) [Sai — đây là **chiến thuật**, không phải kỹ thuật]
- Leo thang đặc quyền (Privilege Escalation) [Sai — đây là **chiến thuật**, không phải kỹ thuật]

**Giải thích:** Trong ma trận MITRE ATT&CK, cấu trúc phân cấp gồm hai cấp độ:

- **Chiến thuật (Tactics):** Mục tiêu tổng quát của kẻ tấn công ở từng giai đoạn — thể hiện dưới dạng **tiêu đề cột** trong ma trận. Ví dụ trong hình: _Privilege Escalation_ và _Defense Evasion_.
- **Kỹ thuật (Techniques):** Phương pháp cụ thể để đạt chiến thuật đó — thể hiện dưới dạng **các mục liệt kê trong cột**. Ví dụ: _Account Manipulation_ là kỹ thuật thuộc chiến thuật _Privilege Escalation_; _Debugger Evasion_ là kỹ thuật thuộc chiến thuật _Defense Evasion_.

---

#### Câu 14. Trong một ngày làm việc bận rộn, bạn nhận được email bất ngờ từ đồng nghiệp, yêu cầu giúp đỡ khẩn cấp và chia sẻ thông tin mà bạn thường không chia sẻ.

Bạn nên làm gì?

- Trả lời email để xác nhận danh tính người gửi.
- Cung cấp thông tin vì đồng nghiệp đang gặp khó khăn.
- Báo cáo với cấp trên về việc đồng nghiệp yêu cầu thông tin nhạy cảm.
- **Xác minh người gửi qua kênh khác trước khi trả lời.** [Đúng]

**Giải thích:** Tình huống này mô tả dấu hiệu điển hình của **spear phishing** hoặc **Business Email Compromise (BEC)** — email giả mạo đồng nghiệp tạo cảm giác khẩn cấp để khiến nạn nhân bỏ qua bước xác minh.

Hành động đúng là **xác minh qua kênh độc lập** — gọi điện trực tiếp cho đồng nghiệp theo số điện thoại đã lưu, không phải số trong email. Lý do: email có thể bị spoofed (giả mạo địa chỉ gửi) hoặc tài khoản đồng nghiệp có thể đã bị xâm phạm. Trả lời email là **không đủ** để xác minh danh tính vì kẻ tấn công đang kiểm soát hòm thư đó.

> **Tại sao không phải "báo cáo với cấp trên"?** Đây không phải bước đầu tiên cần thiết — bước đầu tiên là xác minh danh tính trực tiếp. Nếu sau khi xác minh xác nhận đây là giả mạo, thì mới báo cáo.

---

#### Câu 15. Bạn là chuyên gia tư vấn an ninh mạng được thuê bởi một tổ chức chính quyền địa phương vừa bị tấn công bởi script kiddie. Tổ chức muốn bảo vệ mình khỏi các cuộc tấn công tương tự trong tương lai.

Bạn nên đề xuất gì?

- Xác minh rằng hệ thống phòng thủ có thể đối phó với tấn công DoS kéo dài.
- Giám sát nhân viên cẩn thận và xây dựng văn hóa tổ chức tích cực.
- **Đảm bảo lịch trình vá lỗi (patch management) hiệu quả và hệ thống phòng thủ biên (perimeter defense) được cập nhật thường xuyên.** [Đúng]
- Tập hợp lực lượng lao động được đào tạo đầy đủ với các biện pháp bảo vệ tài sản quan trọng và sao lưu dữ liệu.

**Giải thích:** **Script kiddie** là nhóm tác nhân đe dọa có **trình độ kỹ thuật thấp nhất** — họ dùng công cụ có sẵn và khai thác các lỗ hổng đã được công bố, không tự phát triển exploit. Vì vậy, biện pháp phòng thủ cơ bản nhưng nhất quán là đủ hiệu quả:

- **Patch management:** Cập nhật bản vá kịp thời loại bỏ các lỗ hổng đã biết mà script kiddie dựa vào.
- **Perimeter defense:** Tường lửa, IDS/IPS, và kiểm soát truy cập cơ bản ngăn chặn các công cụ tự động của họ.

Các lựa chọn khác phù hợp hơn với mối đe dọa cấp cao hơn: phòng thủ DDoS kéo dài (cho hacktivist), giám sát nhân viên (cho insider threat), và bảo vệ tài sản quan trọng + backup (cho criminal gang/ransomware).

---

## Đề xuất cải thiện thêm

### 1. Bổ sung mapping với MITRE ATT&CK và Cyber Kill Chain®

Nhiều câu hỏi trong bài đánh giá này liên quan đến các kỹ thuật và giai đoạn tấn công cụ thể. Việc bổ sung tham chiếu chéo sang MITRE ATT&CK technique ID và bước Kill Chain tương ứng giúp người học củng cố kết nối giữa kiến thức lý thuyết và thực tế:

|Câu|Kỹ thuật MITRE ATT&CK|Bước Kill Chain|
|---|---|---|
|1 (DNS attack)|T1584.002 — DNS Server Compromise|Delivery / Exploitation|
|4 (Port 80)|T1046 — Network Service Discovery|Reconnaissance|
|7 (Employee clicks link)|T1204.001 — User Execution: Malicious Link|Exploitation (bước 4)|
|9 (Tailgating/Pretexting)|T1566 — Phishing (Social Engineering)|Delivery|
|13 (Account Manipulation)|T1098 — Account Manipulation|Privilege Escalation|

### 2. Gợi ý câu hỏi phản tư cho từng tình huống

Thay vì chỉ xem đáp án đúng/sai, người học sẽ hiểu sâu hơn nếu được hỏi thêm:

- Câu 5: Ngoài quét lỗ hổng, bạn còn nên xem log đăng nhập và triển khai account lockout policy như thế nào?
- Câu 9: Nếu người lạ phản ứng tức giận khi bạn yêu cầu họ đợi, điều đó nói lên điều gì?
- Câu 15: Script kiddie ngày càng sử dụng AI/LLM để tạo công cụ tấn công — điều này thay đổi mức độ nguy hiểm của nhóm này như thế nào?

### 3. Bổ sung tình huống về AI-assisted attack

Với xu hướng hiện tại (2024–2025), các câu hỏi đánh giá nên bổ sung kịch bản liên quan đến: deepfake trong vishing attacks, LLM-generated phishing email không có lỗi ngữ pháp, và AI-powered vulnerability scanning từ phía kẻ tấn công.