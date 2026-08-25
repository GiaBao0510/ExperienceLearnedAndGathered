# Nghiên cứu Tình huống Thực tế (Case Studies)

## Giới thiệu

Các cuộc tấn công mạng xuất hiện hàng ngày, ảnh hưởng đến cá nhân và tổ chức trong cả khu vực công lẫn tư nhân. Module này phân tích năm trường hợp tấn công mạng nổi bật để minh họa mức độ nguy hiểm đang diễn ra trên toàn cầu. Mỗi trường hợp tập trung vào một loại tác nhân đe dọa khác nhau và mang lại bài học thực tiễn riêng biệt.

Năm trường hợp nghiên cứu, theo thứ tự trình bày:

|Trường hợp|Năm|Loại tác nhân đe dọa|Loại tấn công|
|---|---|---|---|
|Stuxnet|2010|Nation-state (tấn công nhà nước)|Vũ khí mạng / ICS attack|
|LAUSD|2022|Criminal gang (băng nhóm tội phạm)|Ransomware|
|NSA / Edward Snowden|2013|Malicious insider (nội gián)|Rò rỉ dữ liệu có chủ đích|
|Cash App|2021–2022|Malicious insider (cựu nhân viên)|Đánh cắp dữ liệu|
|SolarWinds|2020|Nation-state|Supply chain attack|

---

## Stuxnet — Vũ khí mạng thế hệ đầu tiên

**Bối cảnh**

Khi được phát hiện vào năm 2010, Stuxnet là phần mềm độc hại tinh vi và có mục tiêu cụ thể nhất từng được cộng đồng an ninh mạng ghi nhận. Kẻ tấn công dùng Stuxnet để nhắm vào **hệ thống điều khiển công nghiệp (ICS — Industrial Control System)** và sửa đổi các thông số vận hành quan trọng. Theo kết luận của các chuyên gia, mục tiêu thực sự là các **máy ly tâm làm giàu uranium** tại cơ sở hạt nhân Natanz của Iran — thiết bị thiết yếu trong chương trình hạt nhân.

Stuxnet được cho là do cơ quan tình báo Mỹ (NSA/CIA) và Israel (Unit 8200) phối hợp phát triển, mặc dù cả hai chính phủ chưa bao giờ chính thức xác nhận.

![](https://cpcontents.adobe.com/fr/dynamic-protected/4d4b468959834affa3e5c308d0bf4820/protected/account/2135/resources/7868757/7868757/content/scormcontent/assets/Cybersecurity_StuxnetWorm_en.png)

**Điểm đáng chú ý về mặt kỹ thuật**

- **Khai thác bốn lỗ hổng zero-day cùng lúc:** Zero-day là lỗ hổng chưa được nhà phát hành phần mềm biết đến và chưa có bản vá. Stuxnet sử dụng đồng thời bốn lỗ hổng như vậy — mức độ chưa từng có, vì một zero-day đơn lẻ đã được coi là cực kỳ có giá trị.
- **Dùng chứng chỉ số bị đánh cắp:** Stuxnet được ký bằng chứng chỉ số hợp lệ bị đánh cắp từ Realtek Semiconductor và JMicron — khiến hệ thống bảo mật coi đây là phần mềm hợp lệ.
- **Lây lan qua USB — vượt qua air-gap:** Stuxnet lây lan qua ổ USB bị nhiễm, không cần kết nối internet. Điều này phá tan quan niệm sai lầm rằng mạng cục bộ hoàn toàn cô lập (air-gapped network) là bất khả xâm phạm.
- **Nhắm mục tiêu cực kỳ chính xác:** Stuxnet chỉ kích hoạt và tấn công khi phát hiện đúng loại phần cứng mục tiêu — bộ điều khiển Siemens S7-315 kết nối với biến tần tần số cụ thể. Trên các hệ thống khác, nó ở trạng thái ngủ.
- **Kiên trì và tinh chỉnh theo thời gian:** Chiến dịch kéo dài nhiều tháng, với nhiều phiên bản Stuxnet được cập nhật liên tục để tinh chỉnh các công cụ tấn công.

**Tác động**

Stuxnet được ước tính đã phá hủy khoảng 1.000 trong số 5.000 máy ly tâm tại Natanz, làm chậm chương trình hạt nhân của Iran từ 1 đến 2 năm. Đây là lần đầu tiên vũ khí mạng gây ra thiệt hại vật lý thực sự đối với hạ tầng công nghiệp.

**Bài học rút ra**

- Air-gap (cô lập mạng hoàn toàn) không phải biện pháp bảo mật tuyệt đối — con người và thiết bị vật lý (USB) vẫn là vector tấn công.
- Hệ thống ICS/SCADA cần được bảo mật không kém hệ thống IT thông thường.
- Vũ khí mạng do nhà nước phát triển đặt ra câu hỏi chiến lược mới về luật chiến tranh trong không gian mạng.

---

## Los Angeles Unified School District (LAUSD) — Ransomware nhắm vào giáo dục

**Bối cảnh**

Tháng 9 năm 2022, Học khu Thống nhất Los Angeles (LAUSD) — học khu lớn thứ hai Hoa Kỳ với hơn 1.000 trường học và 600.000 học sinh — bị tấn công ransomware nghiêm trọng. Nhóm tội phạm người Nga **Vice Society** đứng sau vụ tấn công.

**Dữ liệu bị ảnh hưởng**

- Số An sinh xã hội (SSN)
- Thông tin đăng nhập hệ thống
- Biểu mẫu thuế
- Tài liệu pháp lý
- Báo cáo tài chính
- Hồ sơ sức khỏe học sinh
- Kết quả kiểm tra lý lịch
- Đánh giá tâm lý học sinh

**Diễn biến và phản ứng**

Ransomware mã hóa và khóa quyền truy cập vào các hệ thống của LAUSD. LAUSD **quyết định không trả tiền chuộc** — theo khuyến nghị của cơ quan thực thi pháp luật Mỹ (FBI) và chuyên gia an ninh mạng, vì trả tiền không đảm bảo dữ liệu được phục hồi và tài trợ trực tiếp cho hoạt động tội phạm.

Đáp lại, Vice Society đã phát tán dữ liệu bị đánh cắp lên **dark web** — gây thiệt hại thứ cấp nghiêm trọng cho các học sinh và nhân viên có dữ liệu bị lộ.

**Bài học rút ra**

- Các cơ sở giáo dục thường có ngân sách an ninh mạng hạn chế nhưng nắm giữ lượng dữ liệu nhạy cảm lớn — làm cho chúng trở thành mục tiêu hấp dẫn.
- Chiến lược sao lưu dữ liệu (backup) và kế hoạch phục hồi sau thảm họa (disaster recovery) là thiết yếu để giảm thiểu tác động của ransomware.
- Quyết định trả hay không trả tiền chuộc là lựa chọn phức tạp, ảnh hưởng đến nhiều bên liên quan và không có đáp án đúng/sai tuyệt đối.

---

## Cơ quan An ninh Quốc gia Mỹ (NSA) — Edward Snowden

**Bối cảnh**

Năm 2013, Edward Snowden — lúc đó là nhà thầu của NSA làm việc cho Booz Allen Hamilton — đã sao chép và rò rỉ ra công chúng một lượng lớn tài liệu tối mật của NSA và các cơ quan tình báo đồng minh. Ông tiếp cận tài liệu thông qua vai trò công việc hợp pháp của mình với quyền truy cập cấp cao.

Snowden chuyển tài liệu cho các **nhà báo điều tra** — cụ thể là Glenn Greenwald của tờ The Guardian, Laura Poitras (nhà làm phim tài liệu), và Barton Gellman của Washington Post. Các tài liệu sau đó được phân tích và công bố dần qua các kênh báo chí.

**Nội dung bị rò rỉ**

Tài liệu bị rò rỉ tiết lộ các chương trình giám sát quy mô lớn của NSA, bao gồm:

- **PRISM:** Chương trình thu thập dữ liệu từ các công ty công nghệ lớn (Google, Facebook, Apple, Microsoft).
- **XKeyscore:** Hệ thống tìm kiếm và phân tích dữ liệu internet quy mô toàn cầu.
- **Boundless Informant:** Hệ thống theo dõi metadata điện thoại và internet.
- Năng lực kỹ thuật và hướng dẫn hoạt động của NSA.

**Tác động**

- Gây khủng hoảng ngoại giao với nhiều đồng minh (đặc biệt Đức khi phát hiện điện thoại Thủ tướng Merkel bị nghe lén).
- Một số thỏa thuận giữa NSA và công ty công nghệ Mỹ bị xem xét lại.
- Đây được coi là vụ rò rỉ thông tin tình báo gây thiệt hại nghiêm trọng nhất trong lịch sử Mỹ.
- Snowden hiện đang sống tị nạn tại Nga.

**Bài học rút ra**

- Insider threat (mối đe dọa từ người nội bộ) nguy hiểm đặc biệt vì người trong hệ thống đã có quyền truy cập hợp pháp.
- Quyền truy cập nên được cấp theo nguyên tắc **least privilege** (quyền tối thiểu cần thiết) — người dùng chỉ có quyền đúng mức cần thiết để hoàn thành nhiệm vụ.
- Cần có cơ chế giám sát (monitoring) và phát hiện bất thường trong hành vi truy cập dữ liệu.

---

## Cash App — Cựu nhân viên đánh cắp dữ liệu 8 triệu người dùng

**Bối cảnh**

Tháng 12 năm 2021, một cựu nhân viên của Block, Inc. (công ty mẹ của Cash App) đã lạm dụng quyền truy cập vào hệ thống **Cash App Investing** để tải xuống trái phép dữ liệu của hơn 8 triệu người dùng. Sự cố được phát hiện và công bố vào tháng 4 năm 2022 — nghĩa là vi phạm đã không bị phát hiện trong **hơn bốn tháng**.

**Dữ liệu bị đánh cắp**

- Tên đầy đủ của khách hàng
- Số tài khoản môi giới
- Danh mục đầu tư chứng khoán
- Giá trị danh mục đầu tư
- Lịch sử giao dịch chứng khoán

**Vấn đề cốt lõi**

Người thực hiện vi phạm là **cựu nhân viên** — tức là đã bị chấm dứt hợp đồng trước khi thực hiện hành vi, nhưng vẫn còn quyền truy cập vào hệ thống. Đây là lỗi quy trình nghiêm trọng: quyền truy cập của nhân viên nghỉ việc **không được thu hồi ngay lập tức**.

**Bài học rút ra**

- **Offboarding process (quy trình bàn giao khi nghỉ việc)** phải bao gồm việc thu hồi ngay lập tức toàn bộ quyền truy cập hệ thống — không phải sau vài ngày hay vài tuần.
- Cần giám sát liên tục hoạt động hệ thống để phát hiện bất thường (ví dụ: tải xuống lượng lớn dữ liệu bất thường sau giờ làm việc).
- Bốn tháng không phát hiện cho thấy thiếu hụt nghiêm trọng trong khả năng giám sát và phát hiện (monitoring và detection).

> Lưu ý: Vụ này liên quan đến **Cash App Investing**, không phải dịch vụ thanh toán Cash App thông thường.

---

## SolarWinds — Tấn công chuỗi cung ứng quy mô lớn

**Bối cảnh**

Tháng 12 năm 2020, các nhà nghiên cứu từ hãng bảo mật FireEye phát hiện phần mềm độc hại (sau được đặt tên là **SUNBURST**) trong các bản cập nhật của **Orion** — nền tảng quản lý CNTT của SolarWinds được sử dụng rộng rãi trong chính phủ và doanh nghiệp. Vụ tấn công được quy kết cho nhóm APT29 (Cozy Bear), được cho là có liên hệ với SVR — cơ quan tình báo đối ngoại của Nga.

**Cơ chế tấn công**

Kẻ tấn công xâm nhập vào môi trường phát triển phần mềm (build environment) của SolarWinds và chèn mã độc vào quy trình build hợp pháp. Kết quả là các bản cập nhật Orion được phát hành chính thức đã **chứa sẵn backdoor**. Khi khách hàng cài đặt bản cập nhật — một hành động bảo mật hoàn toàn được khuyến nghị — họ vô tình cài đặt cả malware.

Đây chính là bản chất của **supply chain attack (tấn công chuỗi cung ứng)**: thay vì tấn công mục tiêu trực tiếp (khó hơn, dễ bị phát hiện hơn), kẻ tấn công nhắm vào một đơn vị có mối quan hệ tin cậy với nhiều mục tiêu.

**Quy mô tác động**

- Khoảng **18.000 tổ chức** đã cài đặt bản cập nhật bị nhiễm.
- Trong đó, kẻ tấn công chọn lọc và tấn công sâu hơn vào khoảng **100 tổ chức** mục tiêu có giá trị cao.
- Nạn nhân bao gồm nhiều cơ quan liên bang Mỹ: Bộ Tài chính, Bộ Ngoại giao, Bộ An ninh Nội địa, NASA và nhiều cơ quan khác.

**Bài học rút ra**

- **Zero trust architecture** — không tin tưởng mặc định bất kỳ phần mềm hoặc cập nhật nào, kể cả từ nhà cung cấp đáng tin cậy.
- Cần xác minh tính toàn vẹn của phần mềm (code signing, hash verification) trước khi cài đặt.
- Các nhà cung cấp phần mềm phải bảo vệ cả build pipeline và môi trường phát triển — không chỉ sản phẩm cuối.
- Bất chấp vụ SolarWinds, chuyên gia an ninh mạng vẫn **tiếp tục khuyến nghị vá lỗi phần mềm thường xuyên** — rủi ro từ phần mềm lỗi thời không được vá vẫn lớn hơn nhiều so với rủi ro cực kỳ hiếm gặp này.

---

## So sánh tổng hợp năm trường hợp

|Trường hợp|Loại tác nhân|Vector tấn công ban đầu|Thời gian không bị phát hiện|Hậu quả chính|
|---|---|---|---|---|
|Stuxnet|Nation-state|USB bị nhiễm|Nhiều tháng–năm|Phá hủy thiết bị vật lý|
|LAUSD|Criminal gang|Không công bố|Không rõ|Gián đoạn dịch vụ, rò rỉ dữ liệu học sinh|
|NSA/Snowden|Insider|Quyền truy cập hợp pháp|Không phát hiện kịp|Rò rỉ tình báo quốc gia|
|Cash App|Insider (cựu NV)|Quyền truy cập chưa bị thu hồi|4+ tháng|Đánh cắp dữ liệu 8 triệu người dùng|
|SolarWinds|Nation-state|Supply chain / bản cập nhật|~14 tháng|Xâm nhập 18.000 tổ chức|

---

## Thông tin bổ sung

### 1. Bổ sung phân tích theo framework MITRE ATT&CK

Mỗi trường hợp nghiên cứu có thể được phân tích theo framework MITRE ATT&CK để hiểu rõ TTP (Tactics, Techniques, Procedures) được sử dụng. Ví dụ với SolarWinds:

- **Initial Access:** T1195.002 (Compromise Software Supply Chain)
- **Persistence:** T1546 (Event Triggered Execution)
- **Defense Evasion:** T1036 (Masquerading — giả dạng phần mềm hợp lệ)

Việc mapping với ATT&CK giúp người học có thể tra cứu biện pháp phòng thủ (mitigations) tương ứng một cách có hệ thống.

### 2. Thêm timeline cho mỗi vụ tấn công

Nhiều trường hợp trong tài liệu thiếu thông tin về thời gian từ lúc xâm phạm đến lúc bị phát hiện (dwell time). Đây là chỉ số quan trọng: theo báo cáo Mandiant M-Trends 2023, **median dwell time toàn cầu là 16 ngày**. Cash App (4+ tháng) và SolarWinds (~14 tháng) vượt xa mức trung bình, cho thấy thiếu hụt nghiêm trọng trong khả năng phát hiện.

### 3. Nguồn tham khảo thêm

- **Cybersecurity and Infrastructure Security Agency (CISA):** [cisa.gov/known-exploited-vulnerabilities-catalog](https://www.cisa.gov/known-exploited-vulnerabilities-catalog)
- **MITRE ATT&CK Groups:** [attack.mitre.org/groups/](https://attack.mitre.org/groups/) — mô tả chi tiết các nhóm tấn công đã biết, bao gồm Vice Society và APT29
- **Verizon Data Breach Investigations Report (DBIR):** Báo cáo hàng năm phân tích xu hướng tấn công mạng toàn cầu — rất có giá trị để hiểu bức tranh tổng thể
- **Krebs on Security:** [krebsonsecurity.com](https://krebsonsecurity.com/) — theo dõi và phân tích sâu các vụ breach lớn