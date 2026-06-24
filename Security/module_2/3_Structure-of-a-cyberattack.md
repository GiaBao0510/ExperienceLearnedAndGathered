# Cấu trúc của một cuộc Tấn công Mạng

## Giới thiệu

Khi hệ thống máy tính thay đổi, các phương thức tấn công cũng thay đổi theo. Ví dụ, một loại phần mềm độc hại cụ thể có thể chỉ khai thác được trên máy tính chạy phiên bản trình duyệt web lỗi thời. Sau khi bản vá bảo mật được áp dụng, kẻ tấn công không thể tiếp tục dùng cùng phương thức đó.

Mặc dù kỹ thuật tấn công mạng không ngừng phát triển — kể cả việc tích hợp AI — cấu trúc tổng thể của một cuộc tấn công điển hình vẫn tuân theo một khuôn mẫu có thể phân tích và dự đoán được. Tài liệu này giới thiệu hai framework được các chuyên gia an ninh mạng sử dụng rộng rãi để phân tích cấu trúc đó:

- **Cyber Kill Chain®** của Lockheed Martin
- **MITRE ATT&CK Matrix**

Mỗi framework tiếp cận vấn đề từ một góc độ khác nhau và phục vụ các mục đích bổ sung cho nhau — không phải thay thế lẫn nhau.

---

## Cyber Kill Chain® Framework — Lockheed Martin

Lockheed Martin là tập đoàn quốc phòng và công nghệ toàn cầu của Mỹ. Các nhà nghiên cứu tại đây nhận thấy sự tương đồng giữa khái niệm "kill chain" trong học thuyết quân sự (mô tả chuỗi bước cần thiết để tiêu diệt một mục tiêu quân sự) và quy trình xâm nhập vào hệ thống kỹ thuật số. Họ phát triển Cyber Kill Chain® để mô tả hóa và phân tích quy trình đó.

Từ "chain" (chuỗi) mang ý nghĩa quan trọng: **các bước phải xảy ra theo thứ tự tuần tự**, vì mỗi bước phụ thuộc vào kết quả của bước trước. Điều này có hệ quả chiến lược quan trọng cho người phòng thủ: **ngăn chặn bất kỳ bước nào trong chuỗi cũng đủ để phá vỡ toàn bộ cuộc tấn công**. Không nhất thiết phải chờ đến bước cuối mới hành động.

Framework này giúp chuyên gia bảo mật chuyển từ tư duy "ứng phó sau sự kiện" sang tư duy "phân tích và ngăn chặn theo từng giai đoạn".

![](https://cpcontents.adobe.com/fr/dynamic-protected/f270aed172ef4f82b6946a64f3e126cf/protected/account/2135/resources/7868757/7868757/content/scormcontent/assets/CyberF_LM_CyberKillChain_Eng.jpg)

---

### Bảy bước của Cyber Kill Chain®

#### Bước 1: Trinh sát (Reconnaissance)

Kẻ tấn công thu thập thông tin về mục tiêu trước khi hành động. Giai đoạn này chia thành hai loại:

- **Passive reconnaissance:** Thu thập thông tin mà không tương tác trực tiếp với hệ thống mục tiêu — đọc tin tức công khai, phân tích social media, tra cứu DNS, tìm kiếm thông tin trên LinkedIn, GitHub, và các nguồn OSINT (Open Source Intelligence).
- **Active reconnaissance:** Tương tác trực tiếp với hệ thống mục tiêu để thu thập thông tin kỹ thuật — quét cổng mạng (port scanning), phát hiện dịch vụ đang chạy (service enumeration), fingerprinting hệ điều hành và ứng dụng.

**Ý nghĩa với người phòng thủ:** Giảm thiểu thông tin công khai về hạ tầng kỹ thuật (không expose service banner, không để port scan dễ dàng từ internet, quản lý thông tin nhân viên trên mạng xã hội).

---

#### Bước 2: Vũ khí hóa (Weaponization)

Sau khi xác định lỗ hổng cụ thể qua trinh sát, kẻ tấn công tạo ra công cụ tấn công phù hợp. Đây là giai đoạn kẻ tấn công **không tương tác với mục tiêu** — toàn bộ hoạt động diễn ra phía kẻ tấn công.

Các hình thức vũ khí hóa:

- Mua công cụ sẵn có từ thị trường ngầm (exploit kit, malware-as-a-service).
- Tùy chỉnh malware hiện có để né tránh phát hiện của hệ thống bảo mật mục tiêu cụ thể.
- Phát triển zero-day exploit cho lỗ hổng chưa được công bố — thường chỉ xuất hiện ở nhóm tấn công nhà nước (nation-state APT).
- Kết hợp exploit với một dropper (mã chịu trách nhiệm vận chuyển và thực thi payload).

**Ý nghĩa với người phòng thủ:** Không thể quan sát trực tiếp giai đoạn này. Phòng thủ gián tiếp thông qua: vá lỗ hổng kịp thời (giảm bề mặt có thể vũ khí hóa), threat intelligence về các công cụ tấn công mới.

---

#### Bước 3: Phân phối (Delivery)

Kẻ tấn công truyền tải công cụ tấn công đến môi trường của mục tiêu. Đây là **điểm tiếp xúc đầu tiên giữa kẻ tấn công và hạ tầng mục tiêu** — và cũng là một trong những cơ hội phòng thủ tốt nhất.

Các vector phân phối phổ biến:

- **Email** (phổ biến nhất): tệp đính kèm độc hại, link đến trang web giả mạo, phishing/spear phishing.
- **Web:** Drive-by download (người dùng truy cập trang web bị compromised và malware tự download), malicious ad (malvertising).
- **Thiết bị vật lý:** USB bị nhiễm, thiết bị ngoại vi bị thay thế.
- **Supply chain:** Compromised third-party software hoặc update (ví dụ: vụ SolarWinds 2020).

---

#### Bước 4: Khai thác (Exploitation)

Sau khi payload đến được hệ thống mục tiêu, kẻ tấn công cần **kích hoạt** nó. Giai đoạn này là khoảnh khắc lỗ hổng thực sự bị khai thác.

Kích hoạt có thể xảy ra qua:

- Hành động của người dùng: mở tệp đính kèm, nhấp vào link, cắm USB.
- Tự động: zero-click exploit (khai thác lỗ hổng không cần tương tác người dùng — phổ biến trong tấn công cấp cao).

**Exploit** là mã khai thác — đoạn code lợi dụng một lỗ hổng cụ thể để thực thi code trái phép. Exploit thường nhắm vào: lỗ hổng trong ứng dụng (trình duyệt, Office), lỗ hổng hệ điều hành, lỗ hổng trong thư viện/dependencies.

**Ý nghĩa với người phòng thủ:** Patch management kịp thời là biện pháp quan trọng nhất ở giai đoạn này. Sandboxing và application whitelisting có thể ngăn chặn kể cả khi exploit thành công.

---

#### Bước 5: Cài đặt (Installation)

Sau khi khai thác thành công, malware thiết lập **cơ chế duy trì hiện diện (persistence)** trong hệ thống mục tiêu — đảm bảo kẻ tấn công vẫn có quyền truy cập ngay cả khi hệ thống bị khởi động lại hoặc lỗ hổng ban đầu được vá.

Các cơ chế persistence phổ biến:

- Tạo tài khoản người dùng mới (thường với quyền admin ẩn).
- Cài đặt Remote Access Tool (RAT) hoặc backdoor.
- Thêm entry vào registry (Windows) hoặc cron job (Linux) để tự khởi động.
- Đưa lỗ hổng mới vào hệ thống (để có nhiều vector xâm nhập dự phòng).
- Rootkit để che giấu sự hiện diện khỏi phần mềm bảo mật.

> Điểm quan trọng: Ngay cả khi tổ chức phát hiện và vá lỗ hổng ban đầu, kẻ tấn công **vẫn có thể duy trì quyền truy cập** thông qua backdoor đã cài đặt. Đây là lý do tại sao incident response không thể chỉ là "vá lỗ hổng" — cần tìm kiếm và loại bỏ toàn bộ persistence mechanism.

---

#### Bước 6: Điều khiển và Kiểm soát — C2 (Command & Control)

Kẻ tấn công thiết lập kênh liên lạc hai chiều với hệ thống bị xâm nhập để **gửi lệnh và nhận dữ liệu** theo thời gian thực. Đây là kênh điều phối toàn bộ hoạt động tiếp theo.

Các kênh C2 phổ biến:

- **HTTP/HTTPS:** Phổ biến nhất vì traffic hòa lẫn với web traffic thông thường, khó phát hiện.
- **DNS tunneling:** Mã hóa lệnh trong DNS query/response để vượt qua tường lửa.
- **Mạng xã hội và dịch vụ đám mây:** Sử dụng Twitter, GitHub, Google Drive, Dropbox làm kênh C2 — khó block vì là dịch vụ hợp pháp.
- **Peer-to-peer (P2P):** Không có C2 server tập trung, khó takedown hơn.

**Ý nghĩa với người phòng thủ:** Network monitoring và DNS analysis có thể phát hiện C2 traffic bất thường ngay cả khi được che giấu. Threat intelligence về C2 infrastructure của các nhóm tấn công đã biết giúp block proactively.

---

#### Bước 7: Hành động trên Mục tiêu (Actions on Objectives)

Sau khi hoàn thành toàn bộ chuỗi, kẻ tấn công thực hiện mục tiêu cuối cùng. Mục tiêu này khác nhau tùy theo động cơ của kẻ tấn công:

- **Đánh cắp dữ liệu (data exfiltration):** Sao chép thông tin nhạy cảm ra ngoài (hồ sơ khách hàng, bí mật thương mại, thông tin tình báo).
- **Phá hủy/gián đoạn:** Xóa dữ liệu, phá hủy cơ sở hạ tầng, tắt dịch vụ quan trọng.
- **Ransomware deployment:** Mã hóa dữ liệu để tống tiền.
- **Thiết lập bàn đạp (pivot):** Dùng hệ thống bị xâm phạm làm bước đệm tấn công các hệ thống khác trong network (lateral movement).
- **Thao túng dữ liệu:** Sửa đổi dữ liệu theo cách có hại nhưng không dễ phát hiện.

---

### Hạn chế của Cyber Kill Chain®

Cyber Kill Chain® được thiết kế ban đầu để mô tả các cuộc tấn công từ bên ngoài (external attacker). Framework này có một số hạn chế cần biết:

- **Ít phù hợp với insider threat:** Kẻ tấn công nội bộ bỏ qua nhiều bước đầu (không cần trinh sát, không cần phân phối malware).
- **Không phản ánh đủ tấn công web:** Các cuộc tấn công trực tiếp vào web application (SQL injection, XSS) có thể có chuỗi bước rất khác.
- **Quá tuyến tính:** Các cuộc tấn công thực tế thường không tuần tự hoàn toàn — kẻ tấn công có thể lặp lại một số bước, nhảy bước, hoặc thực hiện song song.

Đây là lý do MITRE ATT&CK được phát triển như một bổ sung và mở rộng cho Kill Chain.

---

## MITRE ATT&CK Matrix

**MITRE** là tổ chức phi lợi nhuận của Mỹ, chuyên giải quyết các vấn đề phức tạp trong lĩnh vực quốc phòng, an toàn thông tin và nhiều lĩnh vực khác thông qua các giải pháp dựa trên dữ liệu thực tế.

**ATT&CK** là viết tắt của **Adversarial Tactics, Techniques, and Common Knowledge** — phát âm là "attack". Đây là cơ sở dữ liệu công khai, miễn phí, liên tục cập nhật, tổng hợp các chiến thuật, kỹ thuật và thủ tục (TTP — Tactics, Techniques, and Procedures) mà các nhóm tấn công thực tế đã sử dụng, được rút ra từ quan sát trong thực tế.

![](https://cpcontents.adobe.com/fr/dynamic-protected/d4bb20aaf6cc40fa991f557b89800f4b/protected/account/2135/resources/7868757/7868757/content/scormcontent/assets/CyberF_ATT%26CK_Matrix_Eng.jpg)

Truy cập toàn bộ matrix tại: [https://attack.mitre.org](https://attack.mitre.org/)

---

### Cách đọc ma trận ATT&CK

Ma trận được tổ chức theo hai chiều:

- **Cột (Tactic):** Mỗi cột đại diện cho một **chiến thuật** — mục tiêu tổng quát mà kẻ tấn công muốn đạt được ở giai đoạn đó. Ví dụ: Initial Access (Tiếp cận ban đầu), Persistence (Duy trì hiện diện), Lateral Movement (Di chuyển ngang), Exfiltration (Lọc dữ liệu ra ngoài).
- **Ô trong cột (Technique):** Mỗi ô liệt kê một **kỹ thuật cụ thể** để đạt được chiến thuật đó. Nhiều kỹ thuật còn có **sub-technique** (kỹ thuật con) chi tiết hơn. Nhấp vào từng kỹ thuật để xem: mô tả chi tiết, ví dụ thực tế từ các vụ tấn công đã xảy ra, công cụ được dùng, và biện pháp phòng chống.

**Điểm khác biệt cốt lõi so với Cyber Kill Chain®:** ATT&CK không yêu cầu tuyến tính. Kẻ tấn công có thể dùng bất kỳ kỹ thuật nào ở bất kỳ giai đoạn nào, và thường thực hiện nhiều chiến thuật song song. Ma trận phản ánh sự phức tạp thực tế này tốt hơn.

---

### Ví dụ: Brute Force Attack trong ATT&CK

Kẻ tấn công muốn đạt được **Initial Access** (tiếp cận ban đầu) vào hệ thống. Sau khi xác định hệ thống không có cơ chế khóa tài khoản (account lockout), kẻ tấn công sử dụng kỹ thuật **Brute Force** — thử tự động hàng triệu tổ hợp username/password cho đến khi tìm ra combination hợp lệ.

Trong ATT&CK, Brute Force được phân loại dưới chiến thuật **Credential Access (TA0006)**, với các sub-technique:

- T1110.001: Password Guessing
- T1110.002: Password Cracking
- T1110.003: Password Spraying (thử một mật khẩu phổ biến trên nhiều tài khoản — tránh lockout)
- T1110.004: Credential Stuffing (dùng credential từ breach khác)

Nếu kỹ thuật này thất bại (ví dụ: hệ thống đã bật lockout sau 5 lần thử sai), kẻ tấn công chuyển sang kỹ thuật khác trong cùng chiến thuật — ví dụ: Password Spraying để tránh trigger lockout.

---

## Tầm quan trọng của việc hiểu cấu trúc tấn công

Kẻ tấn công có tính kiên trì cao. Một cuộc tấn công bị gián đoạn hiếm khi bị từ bỏ hoàn toàn — thường kẻ tấn công điều chỉnh kỹ thuật và thử lại. Nhiều chiến dịch APT kéo dài hàng tháng đến hàng năm. Người phòng thủ giỏi không chờ đợi bị tấn công mà chủ động dự đoán bước đi tiếp theo dựa trên hiểu biết về TTP của đối thủ.

---

### Tình huống thực tế: Janina điều tra xâm nhập

Tình huống sau minh họa cách một chuyên gia bảo mật sử dụng ATT&CK trong điều tra thực tế.

**Bối cảnh:** Hệ thống của công ty đã bị xâm nhập. Janina — chuyên gia ứng phó sự cố (incident responder) — bắt đầu điều tra.

---

**Bước 1 — Xác định điểm tiếp cận ban đầu (Initial Access)**

Hai chiến thuật đầu tiên trong ATT&CK (Reconnaissance và Resource Development) đã diễn ra trước khi hệ thống bị xâm nhập. Janina bắt đầu từ chiến thuật **Initial Access** — tìm kiếm điểm xâm nhập đầu tiên vào môi trường của công ty.

---

**Bước 2 — Xác định kỹ thuật: Spear Phishing**

Janina điều tra các log email và phát hiện kẻ tấn công đã sử dụng kỹ thuật **Spear Phishing** (T1566.002 — Spear Phishing Link) nhắm vào một giám đốc điều hành cấp cao. Email giả mạo thuyết phục người này nhấp vào link và cung cấp thông tin đăng nhập trên một trang đăng nhập giả mạo.

---

**Bước 3 — Dự đoán chiến thuật tiếp theo: Execution**

Dựa trên ATT&CK, chiến thuật thường xuất hiện ngay sau Initial Access là **Execution** (Thực thi). Janina biết rằng sau khi có credential của giám đốc, kẻ tấn công có thể dùng chúng để đăng nhập và thực thi mã độc.

Với kinh nghiệm thực tế, Janina ưu tiên kiểm tra các kỹ thuật scripting phổ biến sau spear phishing — đặc biệt là **PowerShell** (T1059.001) và **Windows Command Shell** (T1059.003).

---

**Bước 4 — Xác nhận: Phát hiện PowerShell Script độc hại**

Kiểm tra Windows Event Log và PowerShell transcript log, Janina tìm thấy bằng chứng về PowerShell script được thực thi với quyền của tài khoản giám đốc. Script này tải xuống công cụ tấn công bổ sung từ C2 server.

---

**Bước 5 — Chuẩn bị phòng thủ: Persistence và Privilege Escalation**

Với thông tin hiện có, Janina dự đoán kẻ tấn công đang hoặc sẽ cố gắng thiết lập **Persistence** (Duy trì hiện diện) và thực hiện **Privilege Escalation** (Leo thang đặc quyền). Cô tra cứu ATT&CK để xem các kỹ thuật phổ biến trong hai chiến thuật này và áp dụng các countermeasure tương ứng: kiểm tra các tài khoản mới được tạo, review scheduled task và registry run key, kiểm tra các service bất thường.

---

**Kết luận**

Bằng cách sử dụng ATT&CK như framework điều tra và dự đoán, Janina không chỉ xác định được những gì đã xảy ra mà còn chủ động ngăn chặn các bước tiếp theo của cuộc tấn công — hạn chế thiệt hại và khôi phục hệ thống nhanh hơn.

---

## So sánh: Cyber Kill Chain® vs. MITRE ATT&CK

| Tiêu chí         | Cyber Kill Chain®                                                    | MITRE ATT&CK                                                    |
| ---------------- | -------------------------------------------------------------------- | --------------------------------------------------------------- |
| Nguồn gốc        | Lockheed Martin (2011)                                               | MITRE Corporation (2013, cập nhật liên tục)                     |
| Cách tiếp cận    | Tuyến tính — 7 bước theo thứ tự                                      | Phi tuyến — ma trận chiến thuật và kỹ thuật                     |
| Mức độ chi tiết  | Tổng quát — mô tả giai đoạn                                          | Chi tiết — kỹ thuật cụ thể và sub-technique                     |
| Dữ liệu thực tế  | Không gắn với vụ tấn công cụ thể                                     | Liên kết với TTPs của các nhóm tấn công thực tế                 |
| Dùng cho         | Hiểu tổng quan chuỗi tấn công, lập kế hoạch phòng thủ theo giai đoạn | Threat hunting, incident response, đánh giá gap trong phòng thủ |
| Phù hợp nhất cho | Người mới bắt đầu, lập kế hoạch chiến lược                           | Chuyên gia thực hành, phân tích kỹ thuật                        |
| Cập nhật         | Tương đối ổn định                                                    | Cập nhật thường xuyên theo mối đe dọa mới                       |

---

## Bổ sung

### 1. Mở rộng sang Unified Kill Chain

**Unified Kill Chain** (2017) là sự kết hợp và mở rộng của Cyber Kill Chain® và MITRE ATT&CK, bao gồm 18 phase thay vì 7, phản ánh đầy đủ hơn vòng đời tấn công thực tế — đặc biệt là các giai đoạn lateral movement và persistence bên trong mạng. Phù hợp cho tổ chức muốn framework phân tích toàn diện hơn.

### 2. Diamond Model of Intrusion Analysis

**Diamond Model** là framework bổ sung, phân tích mỗi sự kiện tấn công qua bốn chiều: Adversary (kẻ tấn công), Capability (công cụ/kỹ thuật), Infrastructure (hạ tầng C2), và Victim (nạn nhân). Diamond Model đặc biệt hữu ích cho **threat attribution** (quy kết ai đứng sau cuộc tấn công) và tích hợp tốt với ATT&CK.

### 3. Ứng dụng ATT&CK Navigator

**ATT&CK Navigator** (https://mitre-attack.github.io/attack-navigator/) là công cụ visualization cho phép:

- Đánh dấu các kỹ thuật mà hệ thống phòng thủ hiện tại đã cover.
- Highlight các kỹ thuật được dùng bởi một nhóm APT cụ thể để đánh giá mức độ rủi ro.
- So sánh coverage phòng thủ giữa các giai đoạn khác nhau.

Đây là công cụ thực tế quan trọng mà mọi tổ chức nên biết khi làm việc với ATT&CK.

### 4. Tích hợp với SIEM và Threat Detection

Trong thực tế, các rule phát hiện trong SIEM (như Splunk, Microsoft Sentinel, Elastic SIEM) thường được tag với ATT&CK technique ID (ví dụ: T1110.003). Khi một alert được trigger, security analyst ngay lập tức biết được giai đoạn nào trong kill chain đang xảy ra và có thể tra cứu ATT&CK để hiểu bối cảnh và các bước có thể xảy ra tiếp theo.

Framework **Sigma** cung cấp kho rule phát hiện mã nguồn mở được map với ATT&CK, có thể chuyển đổi sang định dạng của nhiều SIEM khác nhau.
