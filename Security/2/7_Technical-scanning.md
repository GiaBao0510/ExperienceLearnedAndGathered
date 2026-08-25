# Quét Kỹ thuật (Technical Scanning)

## Giới thiệu

Trong module này, bạn sẽ tìm hiểu về các kỹ thuật quét kỹ thuật và lý do kẻ tấn công sử dụng chúng. Cụ thể, bạn sẽ khám phá cách các tác nhân đe dọa (threat actor) triển khai quét trong giai đoạn **trinh sát (reconnaissance)** — giai đoạn thu thập thông tin đầu tiên của một cuộc tấn công mạng.

---

## Tại sao cần quét kỹ thuật?

Các kỹ thuật quét kỹ thuật đóng vai trò thiết yếu trong quản trị và phân tích mạng của một tổ chức. Từ góc độ tấn công, khi điều tra một thiết bị mục tiêu trên mạng, kẻ tấn công muốn trả lời các câu hỏi sau:

- Hệ điều hành nào đang được sử dụng?
- Những dịch vụ nào đang chạy trên thiết bị?
- Có dịch vụ nào tồn tại lỗ hổng bảo mật đã biết không?

Phần dưới đây sẽ khám phá từng kỹ thuật quét kỹ thuật điển hình.

---

## Ping Test

### Cơ chế hoạt động

Kiểm tra ping đo thời gian cần thiết để một gói dữ liệu di chuyển từ thiết bị này đến thiết bị khác và quay trở lại. Gói dữ liệu này nhỏ và được định dạng chuẩn — tương tự như một bưu thiếp kỹ thuật số.

![](https://support.ipvanish.com/hc/article_attachments/1260802822289)

Trong một bài kiểm tra ping, thiết bị quét gửi một gói tin **ICMP Echo Request** (yêu cầu phản hồi) đến địa chỉ IP của thiết bị mục tiêu. Nếu thiết bị mục tiêu đang hoạt động và không bị chặn bởi tường lửa, nó sẽ phản hồi bằng một gói tin **ICMP Echo Reply**.

> **Lưu ý kỹ thuật:** Nhiều tổ chức cấu hình tường lửa để chặn ICMP nhằm ẩn thiết bị khỏi công cụ quét ping. Do đó, việc không nhận được phản hồi ping không có nghĩa chắc chắn là thiết bị không tồn tại hoặc không hoạt động.

### Thông tin thu được

Một bài kiểm tra ping cho biết:
- Thiết bị có phản hồi không (tức là có đang hoạt động và kết nối mạng không).
- Khi thực hiện trên dải địa chỉ IP, xác định được bao nhiêu thiết bị đang hoạt động trên mạng.
- Khoảng cách thiết bị trong mạng thông qua thuộc tính **TTL (Time To Live)** của gói tin.

**Giải thích TTL:** Mỗi gói tin được khởi tạo với một giá trị TTL nhất định (ví dụ: Windows mặc định TTL = 128, Linux = 64). Mỗi router chuyển tiếp gói tin sẽ giảm TTL đi 1. Khi TTL về 0, gói tin bị hủy và thiết bị gửi lại thông báo lỗi.

> **Ví dụ:** Gói tin bắt đầu với TTL = 120 và đến đích với TTL = 108. Điều này có nghĩa gói tin đã đi qua **12 router** (12 hop). Thông tin này giúp ước tính khoảng cách mạng đến mục tiêu.

> Bản gốc ghi "TTL còn lại là 108, đã trải qua 12 giai đoạn" — đúng về số học (120 - 108 = 12) nhưng thiếu giải thích cơ chế TTL và giá trị TTL mặc định theo hệ điều hành, thông tin quan trọng để diễn giải kết quả ping đúng.

**Lệnh trên các hệ điều hành:**

```bash
# Windows
ping target_name

# Linux / macOS
ping target_name
ping -c 4 target_name    # Giới hạn 4 gói tin
```

---

## Traceroute

### Cơ chế hoạt động

Traceroute là công cụ chẩn đoán mạng cho phép lập bản đồ đường đi của gói tin từ thiết bị quét đến đích. Công cụ này khai thác cơ chế TTL: gửi các gói tin với giá trị TTL **tăng dần** (bắt đầu từ 1, sau đó 2, 3...).

Khi TTL của một gói tin giảm về 0 tại một router trung gian, router đó sẽ hủy gói tin và gửi lại thông báo lỗi ICMP cho thiết bị quét — đồng thời tiết lộ địa chỉ IP của chính nó. Bằng cách tăng dần TTL, traceroute "khám phá" từng router trên đường đi đến đích.

![](https://www.wikihow.com/images_en/thumb/9/9b/Traceroute-Step-6-Version-3.jpg/v4-460px-Traceroute-Step-6-Version-3.jpg.webp)

> Bản gốc mô tả TTL "tăng dần hoặc giảm dần" — không chính xác. Traceroute sử dụng TTL **tăng dần** (từ 1 trở lên), không giảm dần. Đã sửa.

### Thông tin thu được

Traceroute cho phép:
- Lập bản đồ toàn bộ đường đi mạng từ thiết bị quét đến đích, bao gồm địa chỉ IP của tất cả router trung gian.
- Xác định số lượng router và switch nằm giữa hai điểm.
- Phát hiện điểm nghẽn cổ chai (bottleneck) hoặc điểm thất bại trong mạng.

**Router** là thiết bị phần cứng kết nối các mạng khác nhau với nhau (và với internet). **Switch** tích hợp các thiết bị trong cùng một mạng LAN, cho phép chúng giao tiếp nội bộ.

> **Ví dụ:** Mục tiêu cách 12 hop. Khi gửi gói tin với TTL = 1, router đầu tiên sẽ hủy gói và trả về địa chỉ IP của nó. Khi TTL = 2, router thứ hai phản hồi. Cứ tiếp tục như vậy cho đến khi gói tin đến được đích (TTL = 12). Kết quả là một danh sách đầy đủ 12 địa chỉ IP của các router trên đường đi.

**Lệnh trên các hệ điều hành:**

```bash
# Windows
tracert target_name

# Linux / macOS
traceroute target_name
```

---

## Quét cổng (Port Scanning)

### Cơ chế hoạt động

Trong mạng máy tính, ứng dụng giao tiếp với bên ngoài thông qua các **cổng kỹ thuật số (port)**. Cổng là điểm đầu cuối logic để gửi và nhận dữ liệu cho một dịch vụ mạng cụ thể.

**Hệ thống địa chỉ:** Địa chỉ IP xác định thiết bị trên mạng, còn số cổng xác định dịch vụ cụ thể trên thiết bị đó. Hãy hình dung địa chỉ IP như địa chỉ tòa nhà, còn số cổng như số căn phòng trong tòa nhà.

TCP (Transmission Control Protocol) cung cấp tổng cộng **65.536 cổng** (từ 0 đến 65.535), chia làm ba nhóm:

| Nhóm | Phạm vi | Đặc điểm |
|---|---|---|
| Well-known ports | 0 – 1023 | Được IANA gán cho các dịch vụ tiêu chuẩn |
| Registered ports | 1024 – 49151 | Có thể đăng ký cho ứng dụng cụ thể |
| Dynamic/Private ports | 49152 – 65535 | Dùng tạm thời, không đăng ký |

**Trạng thái cổng:** Khi quét, mỗi cổng có thể ở một trong ba trạng thái:
- **Open (mở):** Cổng đang lắng nghe và chấp nhận kết nối — dịch vụ đang chạy.
- **Closed (đóng):** Thiết bị phản hồi nhưng không có dịch vụ nào lắng nghe trên cổng này.
- **Filtered (lọc):** Thiết bị không phản hồi — thường do tường lửa hoặc bộ lọc gói (packet filter) chặn.

![](https://www.paloaltonetworks.com/content/dam/pan/en_US/images/cyberpedia/port-scanning.png?imwidth=720)

### Thông tin thu được

Bằng cách xem xét các cổng đang mở của một thiết bị, người quét có thể suy luận mục đích sử dụng của thiết bị đó. Dưới đây là một số cổng quan trọng cần biết:

| Cổng | Giao thức/Dịch vụ | Ý nghĩa bảo mật |
|---|---|---|
| 22 | SSH | Truy cập từ xa an toàn; nếu mở, kiểm tra xác thực |
| 80 | HTTP | Web server; kiểm tra lỗ hổng web application |
| 443 | HTTPS | Web server với TLS; kiểm tra cấu hình chứng chỉ |
| 445 | SMB (Windows File Sharing) | Mục tiêu của WannaCry, EternalBlue; cần vá ngay |
| 3306 | MySQL | Database; không nên để mở ra internet |
| 3389 | RDP (Remote Desktop Protocol) | Truy cập desktop từ xa; mục tiêu brute-force phổ biến |
| 8080 | HTTP thay thế | Thường dùng cho ứng dụng web, API server |

> **Ví dụ thực tế — WannaCry và cổng 445:** Ransomware WannaCry (2017) khai thác lỗ hổng EternalBlue trên giao thức SMB qua cổng TCP 445. Cuộc tấn công lây lan toàn cầu và gây thiệt hại ước tính hàng tỷ USD. Đây là minh chứng rõ ràng về nguy cơ của cổng mở không được bảo vệ.
>
> **Ví dụ thực tế — RDP (cổng 3389):** Cổng RDP mở ra internet là mục tiêu thường xuyên của các cuộc tấn công brute-force và credential stuffing. Nhiều vụ tấn công ransomware bắt đầu bằng cách khai thác RDP không được bảo vệ.

---

## Quét lỗ hổng (Vulnerability Scanning)

### Cơ chế hoạt động

Quét lỗ hổng là bước nâng cao hơn so với quét cổng — không chỉ xác định dịch vụ đang chạy mà còn chủ động tìm kiếm các điểm yếu có thể khai thác.

![](https://siegecyber.com.au/app/uploads/2025/04/Vulnerability-Scanning-Best-Practices-for-Accurate-Detection5-1-768x403.jpg)

Hai kỹ thuật cơ bản:

- **Version detection (Phát hiện phiên bản):** Xác định phiên bản cụ thể của phần mềm đang chạy (ví dụ: Apache 2.4.49). Kẻ tấn công đối chiếu thông tin này với cơ sở dữ liệu CVE (Common Vulnerabilities and Exposures) để tìm lỗ hổng đã biết cho phiên bản đó.
- **OS detection (Phát hiện hệ điều hành):** Xác định hệ điều hành và phiên bản của thiết bị (ví dụ: Windows Server 2019, Ubuntu 22.04). Thông tin này giúp thu hẹp danh sách exploit tiềm năng.

Ngoài ra còn có **quét động (dynamic scanning)** — mô phỏng các kỹ thuật tấn công thực tế như SQL injection, XSS để kiểm tra xem hệ thống có bị tổn thương không.

### Thông tin thu được

Quét lỗ hổng là công cụ hai lưỡi:
- **Với tổ chức:** Xác định điểm yếu trong mạng để ưu tiên vá lỗi trước khi bị tấn công.
- **Với kẻ tấn công:** Tìm kiếm nạn nhân tiềm năng có lỗ hổng dễ khai thác.

> **Ví dụ:** Công cụ quét kết nối đến máy chủ web và phát hiện đang chạy Apache 2.4.49 — phiên bản có lỗ hổng Path Traversal nghiêm trọng (CVE-2021-41773, CVSS score 7.5). Công cụ quét có thể tự động thử khai thác để xác nhận lỗ hổng tồn tại và báo cáo kết quả.

> **Lưu ý pháp lý — Cực kỳ quan trọng:** Quét lỗ hổng có thể tự động thực hiện các hành động bị coi là tấn công mạng theo pháp luật nhiều quốc gia. **Chỉ thực hiện quét trên hệ thống mà bạn được phép kiểm thử hoặc trên hệ thống do bạn sở hữu.** Việc quét trái phép vào hệ thống của người khác — dù không gây hại — có thể vi phạm luật tội phạm mạng và dẫn đến hậu quả pháp lý nghiêm trọng.

---

## Shodan — Công cụ tìm kiếm cho thiết bị kết nối internet

**Shodan** ([shodan.io](https://www.shodan.io/)) là công cụ tìm kiếm chuyên biệt, thường được mô tả là "Google dành cho IoT và thiết bị kết nối internet".

![](https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSnwbtPaxTUF90hcGqaOpzJiCC07yhehoHd4akG2yML2ynLlRf3be5GXms&s=10)

Khác với Google index nội dung web, Shodan liên tục quét toàn bộ không gian địa chỉ IPv4 và lập chỉ mục thông tin về các thiết bị và dịch vụ đang lắng nghe trên internet — bao gồm: banner thông tin dịch vụ, phiên bản phần mềm, chứng chỉ SSL, thông tin địa lý, và nhiều hơn nữa.

Shodan thu hút cả nhà nghiên cứu bảo mật lẫn kẻ tấn công vì khả năng tìm kiếm theo nhiều tiêu chí:

```
# Ví dụ tìm kiếm trên Shodan
apache 2.4.49          # Tìm máy chủ chạy Apache phiên bản dễ bị tấn công
port:3389 country:VN   # Tìm RDP mở ở Việt Nam
"default password"     # Thiết bị sử dụng mật khẩu mặc định
product:MongoDB        # Tìm MongoDB server (nhiều cái không có xác thực)
```

> Shodan là ví dụ điển hình về sức mạnh của OSINT kỹ thuật: thông tin về hàng tỷ thiết bị kết nối internet đã được thu thập sẵn, không cần tấn công hay quét trực tiếp vào mục tiêu.

---

## Quét mạng với Nmap

Nmap ([nmap.org](https://nmap.org/)) — viết tắt của **Network Mapper** — là công cụ quét mạng mã nguồn mở phổ biến nhất, hỗ trợ Windows, macOS và Linux.

Mặc dù quét cổng là tính năng cốt lõi, Nmap còn cung cấp nhiều khả năng quan trọng khác:

![](https://cpcontents.adobe.com/fr/dynamic-protected/4d4b468959834affa3e5c308d0bf4820/protected/account/2135/resources/7868757/7868757/content/scormcontent/assets/NetworkScanning_Path.png)

Lập bản đồ đường dẫn mạng — hiển thị tất cả router/host trên đường đi đến đích.

![](https://cpcontents.adobe.com/fr/dynamic-protected/4d4b468959834affa3e5c308d0bf4820/protected/account/2135/resources/7868757/7868757/content/scormcontent/assets/NetworkScanning_Detection.png)

Phát hiện phiên bản dịch vụ và hệ điều hành.

![](https://cpcontents.adobe.com/fr/dynamic-protected/4d4b468959834affa3e5c308d0bf4820/protected/account/2135/resources/7868757/7868757/content/scormcontent/assets/NetworkScanning_Firewall.png)

Phát hiện tường lửa và bộ lọc gói tin.

### Các lệnh Nmap cơ bản

```bash
# Quét ping — kiểm tra host nào đang hoạt động
nmap -sn 192.168.1.0/24

# Quét cổng cơ bản (1000 cổng phổ biến nhất)
nmap 192.168.1.1

# Quét toàn bộ 65.535 cổng
nmap -p- 192.168.1.1

# Quét với phát hiện phiên bản dịch vụ và OS
nmap -sV -O 192.168.1.1

# Quét nhanh và toàn diện (phổ biến trong pentest)
nmap -A 192.168.1.1

# Quét chậm để tránh bị IDS phát hiện
nmap -sS -T2 192.168.1.1
```

> Bản gốc chỉ đề cập Nmap ở mức mô tả chung. Đã bổ sung bảng lệnh cơ bản vì đây là kỹ năng thực hành thiết yếu cho bất kỳ ai học về bảo mật mạng.

### Zenmap — Giao diện đồ họa cho Nmap

**Zenmap** ([nmap.org/zenmap](https://nmap.org/zenmap/)) là giao diện đồ họa (GUI) chính thức của Nmap. Zenmap phù hợp với người mới bắt đầu vì:
- Cho phép chọn loại quét qua giao diện trực quan thay vì nhớ cú pháp lệnh.
- Hiển thị lệnh Nmap tương ứng với cấu hình quét đã chọn — giúp người dùng học dần cú pháp lệnh thực tế.
- Trực quan hóa kết quả quét dưới dạng đồ thị topology mạng.

---

## AI trong quét kỹ thuật

Trí tuệ nhân tạo (AI) và học máy (ML) đang được tích hợp vào các công cụ quét mạng hiện đại, mang lại khả năng vượt trội so với quét truyền thống:

**Phân tích tự động:** Thay vì chờ chuyên viên xem xét từng kết quả, công cụ AI có thể tự động phân tích, phân loại lỗ hổng, và ưu tiên chúng theo mức độ nghiêm trọng (dựa trên CVSS score, khả năng khai thác thực tế, và giá trị tài sản bị ảnh hưởng).

**Học từ lịch sử:** Mô hình ML được huấn luyện trên dữ liệu quét lịch sử có thể nhận diện các mẫu tấn công (attack patterns) và phát hiện bất thường so với baseline thông thường của mạng.

**Giảm false positive:** Một trong những thách thức lớn của quét lỗ hổng là lượng lớn cảnh báo giả (false positive). AI giúp lọc và xác nhận các lỗ hổng thực sự đáng lo ngại, để chuyên viên tập trung vào những vấn đề quan trọng.

**Ví dụ:** **IBM QRadar Advisor with Watson** tích hợp AI để phân tích sự kiện bảo mật, điều tra vi phạm tiềm tàng, và cung cấp thông tin hành động được (actionable insights) cho đội ngũ bảo mật. Thay vì xem xét thủ công hàng nghìn cảnh báo mỗi ngày, chuyên viên chỉ cần tập trung vào những mối đe dọa đã được AI xác nhận là nghiêm trọng.

---

## Đề xuất cải thiện thêm

### 1. Bổ sung phân biệt quét chủ động vs. thụ động

Tài liệu hiện tại chỉ đề cập quét chủ động (active scanning) — gửi gói tin đến mục tiêu và nhận phản hồi. Cần bổ sung khái niệm **quét thụ động (passive scanning)**:

- **Active scanning:** Gửi probe packet đến mục tiêu, dễ bị phát hiện bởi IDS/IPS.
- **Passive scanning:** Thu thập thông tin từ traffic mạng sẵn có mà không gửi gói tin đến mục tiêu — không để lại dấu vết trực tiếp. Ví dụ: lắng nghe ARP broadcast trong mạng LAN để lập bản đồ thiết bị.

### 2. Giới thiệu CVE và CVSS

Tài liệu đề cập lỗ hổng bảo mật nhưng không giải thích hệ thống định danh lỗ hổng tiêu chuẩn:

- **CVE (Common Vulnerabilities and Exposures):** Hệ thống định danh duy nhất cho lỗ hổng bảo mật đã biết. Mỗi lỗ hổng được gán một mã CVE-YYYY-NNNNN (ví dụ: CVE-2021-41773). Cơ sở dữ liệu tại [nvd.nist.gov](https://nvd.nist.gov/).
- **CVSS (Common Vulnerability Scoring System):** Hệ thống cho điểm mức độ nghiêm trọng từ 0.0 đến 10.0. Điểm này được tính dựa trên nhiều yếu tố: vector tấn công, độ phức tạp, yêu cầu xác thực, và tác động đến confidentiality/integrity/availability.

### 3. Giới thiệu các công cụ quét lỗ hổng phổ biến

Ngoài Nmap, một số công cụ thường được dùng trong đánh giá bảo mật thực tế:

- **Nessus / Tenable:** Công cụ quét lỗ hổng thương mại phổ biến nhất trong doanh nghiệp. Có phiên bản Essentials miễn phí.
- **OpenVAS (Greenbone):** Thay thế mã nguồn mở đầy đủ tính năng cho Nessus.
- **Metasploit Framework:** Nền tảng kiểm thử xâm nhập có tích hợp quét và khai thác lỗ hổng.
- **Masscan:** Quét cổng cực nhanh, có thể quét toàn bộ không gian IPv4 trong vài phút.

### 4. Liên kết với MITRE ATT&CK

Các kỹ thuật quét được mô tả trong tài liệu này tương ứng với chiến thuật **Reconnaissance (TA0043)** trong MITRE ATT&CK:

- Ping sweep → T1595.001 (Active Scanning: Scanning IP Blocks)
- Port scanning → T1046 (Network Service Discovery)
- OS/Version detection → T1592 (Gather Victim Host Information)

Việc liên kết với MITRE ATT&CK giúp người học hiểu vị trí của các kỹ thuật này trong bức tranh tấn công tổng thể và tra cứu biện pháp phòng thủ tương ứng.
