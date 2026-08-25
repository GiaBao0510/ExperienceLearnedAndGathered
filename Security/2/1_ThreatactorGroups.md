# Các Nhóm Tác nhân Đe dọa (Threat Actor Groups)

## Giới thiệu

Trong lĩnh vực an ninh mạng, thuật ngữ **"tác nhân đe dọa" (threat actor)** được dùng phổ biến hơn "kẻ tấn công" (attacker), vì nó bao quát hơn về mặt ngữ nghĩa.

**Kẻ tấn công (attacker / cyber attacker)** là người đang thực sự tiến hành hành vi xâm phạm hệ thống trái phép. **Tác nhân đe dọa (threat actor)** là khái niệm rộng hơn: chỉ bất kỳ thực thể nào — cá nhân, nhóm, hoặc tổ chức — có khả năng và ý định gây ra mối đe dọa an ninh mạng, kể cả khi chưa hành động. Sự phân biệt này quan trọng vì nó cho phép tổ chức chủ động nhận diện và đánh giá mối đe dọa trước khi cuộc tấn công thực sự xảy ra.

Các tác nhân đe dọa được phân loại thành các nhóm dựa trên động cơ, nguồn lực và phương thức tấn công. Module này phân tích năm nhóm chính:

- Script kiddie
- Hacktivist
- Criminal gang
- Nation-state hacker
- Malicious insider

![](https://cpcontents.adobe.com/fr/dynamic-protected/dfcd95bda7864ddd85dc6895ff17a00a/protected/account/2135/resources/7868757/7868757/content/scormcontent/assets/CyberF_AttackerGroups_Eng.png)

---

## Script kiddie

Script kiddie là nhóm có trình độ kỹ thuật thấp nhất trong phân loại các tác nhân đe dọa. Họ sử dụng các công cụ và kỹ thuật tấn công có sẵn — thường do các hacker có kinh nghiệm hơn phát triển — mà không thực sự hiểu cơ chế hoạt động bên trong.

Mặc dù mức độ kỹ thuật thấp, script kiddie vẫn là mối đe dọa thực sự: họ có thể khai thác các lỗ hổng đã biết trên các hệ thống chưa được vá. Số lượng lớn và tính dễ tiếp cận của các công cụ tấn công tự động khiến nhóm này có tần suất hoạt động cao.

![](https://tse4.mm.bing.net/th/id/OIP.Krz1oC2MpRMKUXQ-ZWDexgHaEt?cb=thfc1falcon2&rs=1&pid=ImgDetMain&o=7&rm=3)

### Đặc điểm

- **Đối tượng:** Phần lớn là thanh thiếu niên hoặc người trẻ tuổi, tự học qua diễn đàn, video, và thử nghiệm trực tiếp.
- **Động cơ:** Danh tiếng trong cộng đồng hacker, giải trí, hoặc trả thù cá nhân — hiếm khi có động cơ tài chính có tổ chức.
- **Nguồn lực:** Ngân sách thấp, phụ thuộc vào công cụ miễn phí (Metasploit, SQLmap, exploit kits công khai) và các lỗ hổng đã được công bố (published CVEs).
- **Xu hướng mới:** Ngày càng nhiều script kiddie sử dụng công cụ AI (chatbot, code generation) để tạo script độc hại mà không cần hiểu kỹ thuật nền tảng — làm giảm đáng kể rào cản gia nhập.

> **Penetration testing** (kiểm thử xâm nhập) là loại kiểm thử bảo mật mô phỏng kỹ thuật tấn công thực tế nhằm phát hiện các lỗ hổng mà kẻ tấn công có thể khai thác trước khi chúng bị lợi dụng.

### Phòng thủ

Chiến lược cốt lõi là duy trì một môi trường ít "quả chín trên cây" nhất có thể — tức là đảm bảo các hệ thống không chứa lỗ hổng đã biết và dễ khai thác. Khi tất cả lỗ hổng phổ biến đều được vá, script kiddie sẽ chuyển sang mục tiêu dễ hơn.

- Duy trì lịch vá lỗ hổng (patch management) nghiêm ngặt và kịp thời.
- Triển khai các biện pháp phòng thủ vành đai cơ bản: tường lửa, IDS/IPS, giới hạn bề mặt tấn công.
- Giám sát các CVE mới được công bố liên quan đến hệ thống đang vận hành.

### Hồ sơ tóm tắt

| Tiêu chí         | Mô tả                                                                                    |
| ---------------- | ---------------------------------------------------------------------------------------- |
| Họ là ai?        | Cá nhân tự học, thường là thanh thiếu niên, thiếu nền tảng kỹ thuật chuyên sâu           |
| Mục tiêu của họ? | Danh tiếng, giải trí, hoặc trả thù cá nhân                                               |
| Nguồn lực?       | Ngân sách hạn chế, dùng công cụ tự động miễn phí và lỗ hổng đã công bố                   |
| Phòng chống?     | Patch management hiệu quả và các biện pháp phòng thủ vành đai được cập nhật thường xuyên |

---

## Hacktivist

Thuật ngữ "hacktivist" là sự kết hợp của "hacker" và "activist" (nhà hoạt động xã hội/chính trị). Hacktivist sử dụng tấn công mạng như một phương tiện để tạo áp lực, gây gián đoạn, hoặc gây chú ý cho một mục tiêu chính trị, xã hội, hoặc kinh tế mà họ ủng hộ.

![](https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTiAj1_CloWsODmgO8gzv5wgVTmqJMEPXUlzpxIynqou64csaRN4WLAnrSS&s=10)

### Đặc điểm

- **Động cơ:** Lý tưởng tư tưởng là yếu tố cốt lõi. Mục tiêu tấn công thường liên quan đến các sự kiện chính trị nóng bỏng (xung đột quốc tế, bầu cử, chính sách gây tranh cãi).
- **Thành phần nhóm:** Không đồng nhất — bao gồm cả amateur dễ bị thuyết phục lẫn hacker có kinh nghiệm khi vấn đề gây đủ sức hút. Khi một sự kiện lớn xảy ra (chiến tranh, scandal chính phủ), các hacktivist có trình độ cao hơn thường tham gia, nâng cấp đáng kể năng lực của nhóm.
- **Phương thức:** Chủ yếu dùng DDoS (Distributed Denial of Service) để làm gián đoạn dịch vụ, defacement website (thay đổi giao diện trang web để truyền tải thông điệp), và rò rỉ dữ liệu.
- **Sức mạnh đến từ quy mô:** Một hacktivist đơn lẻ gây mối đe dọa nhỏ, nhưng hàng trăm người phối hợp cùng lúc có thể áp đảo ngay cả hệ thống phòng thủ trung bình.
- **Ví dụ nổi tiếng:** Nhóm **Anonymous** — tổ chức hacktivist quốc tế phi tập trung, đã thực hiện nhiều chiến dịch tấn công nhắm vào các tập đoàn lớn, chính phủ, và tổ chức tôn giáo. Các chiến dịch đáng chú ý bao gồm Operation Payback (2010) chống lại các tổ chức tài chính phong tỏa WikiLeaks, và nhiều chiến dịch liên quan đến xung đột địa chính trị.

### Phòng thủ

Với hacktivist, câu hỏi không chỉ là "làm thế nào để ngăn chặn" mà còn là "làm thế nào để vẫn hoạt động được khi đang bị tấn công". Các cuộc tấn công DDoS quy mô lớn không phải lúc nào cũng có thể chặn hoàn toàn.

- Triển khai biện pháp giảm thiểu DDoS (DDoS mitigation) — có thể dùng dịch vụ CDN như Cloudflare, Akamai.
- Xây dựng kế hoạch ứng phó sự cố (Incident Response Plan) cho kịch bản tấn công kéo dài.
- Nếu tổ chức hoạt động trong lĩnh vực nhạy cảm về mặt chính trị, cần đánh giá liên tục khả năng trở thành mục tiêu và chuẩn bị trước.

### Hồ sơ tóm tắt

| Tiêu chí         | Mô tả                                                                                           |
| ---------------- | ----------------------------------------------------------------------------------------------- |
| Họ là ai?        | Những người theo lý tưởng tư tưởng, hình thành liên minh lỏng lẻo, thường phi tập trung         |
| Mục tiêu của họ? | Tạo ra sự thay đổi chính trị hoặc xã hội thông qua gây gián đoạn và gây chú ý                   |
| Nguồn lực?       | Đa dạng; sức mạnh chính đến từ quy mô phối hợp, không phải kỹ thuật đơn lẻ                      |
| Phòng chống?     | Hệ thống phòng thủ có khả năng chịu đựng tấn công gián đoạn kéo dài; kế hoạch dự phòng vận hành |

---

## Criminal gang (Băng nhóm tội phạm mạng)

Băng nhóm tội phạm mạng là nhóm tác nhân đe dọa đang phát triển nhanh nhất và gây thiệt hại tài chính lớn nhất trong thực tế hiện đại. Internet đã tạo ra một môi trường cho phép tội phạm hoạt động xuyên biên giới với rủi ro bị bắt giữ thấp hơn nhiều so với tội phạm truyền thống.

### Đặc điểm

- **Động cơ:** Hoàn toàn tài chính. Tội phạm mạng là "ngành kinh doanh" với lợi nhuận cao và ngưỡng rủi ro thấp — đặc biệt khi giao dịch qua tiền mã hóa (cryptocurrency) để che giấu dòng tiền.

- **Tổ chức:** Từ nhóm nhỏ vài người đến tổ chức đa quốc gia với hàng trăm thành viên chuyên biệt theo từng vai trò (phát triển malware, triển khai tấn công, rửa tiền, hỗ trợ nạn nhân để thu tiền chuộc). Nhiều nhóm trao đổi công cụ và dịch vụ trên dark web.

- **Phương thức phổ biến:**
  
  - **Ransomware attacks:** Mã hóa dữ liệu nạn nhân và đòi tiền chuộc để giải mã. Các nhóm lớn như LockBit, REvil, BlackCat đã gây thiệt hại hàng tỷ USD.
  - **Data theft:** Đánh cắp dữ liệu khách hàng, thông tin tài chính, bí mật thương mại để bán hoặc tống tiền.
  - **Business Email Compromise (BEC):** Giả mạo email cấp cao trong tổ chức để thuyết phục nhân viên chuyển tiền hoặc chia sẻ thông tin nhạy cảm.
  - **Credential stuffing:** Sử dụng danh sách tài khoản/mật khẩu bị rò rỉ từ các vụ vi phạm trước để đăng nhập hàng loạt vào các dịch vụ khác.

- **Thách thức pháp lý:** Do hoạt động xuyên quốc gia và sử dụng tiền mã hóa, việc bắt giữ và truy tố rất khó khăn. Nhiều băng nhóm hoạt động từ các quốc gia không có hiệp ước dẫn độ với phương Tây.

### Phòng thủ

Với tội phạm mạng có tổ chức, phòng thủ phải phân cấp theo mức độ quan trọng của tài sản:

- Triển khai giải pháp **EDR (Endpoint Detection and Response)** để phát hiện sớm ransomware trước khi lan rộng.
- Duy trì **chiến lược sao lưu 3-2-1**: 3 bản sao, 2 loại phương tiện khác nhau, 1 bản offsite — và kiểm tra khả năng phục hồi định kỳ.
- Đào tạo nhân viên nhận biết phishing và BEC.
- Phân đoạn mạng (network segmentation) để ngăn ransomware lây lan từ endpoint sang production server.

### Hồ sơ tóm tắt

| Tiêu chí         | Mô tả                                                                                                   |
| ---------------- | ------------------------------------------------------------------------------------------------------- |
| Họ là ai?        | Nhóm trong nước và quốc tế, tổ chức theo kiểu doanh nghiệp tội phạm chuyên nghiệp                       |
| Mục tiêu của họ? | Lợi nhuận tài chính                                                                                     |
| Nguồn lực?       | Mua, bán, trao đổi công cụ và dịch vụ trên dark web; sử dụng cryptocurrency để ẩn dòng tiền             |
| Phòng chống?     | Nhân viên được đào tạo, bảo vệ tài sản quan trọng, sao lưu dữ liệu thường xuyên và có khả năng phục hồi |

---

## Nation-state hacker (Hacker được nhà nước bảo trợ)

Đây là nhóm tác nhân đe dọa tinh vi và nguy hiểm nhất. Nhiều quốc gia đã chính thức công nhận không gian mạng là lĩnh vực xung đột thứ năm bên cạnh đất liền, biển, không phận và vũ trụ, và đầu tư vào năng lực tấn công/phòng thủ mạng như một thành phần chiến lược quốc gia.

### Đặc điểm

- **Động cơ:** Phục vụ mục tiêu chiến lược quốc gia — trinh sát và gián điệp (thu thập thông tin tình báo), phá hoại cơ sở hạ tầng đối thủ, thao túng thông tin (information warfare), và lợi thế kinh tế (đánh cắp bí mật thương mại, sở hữu trí tuệ).
- **Năng lực:** Các thành viên là chuyên gia được đào tạo bài bản, làm việc toàn thời gian với đầy đủ nguồn lực. Họ thường tiếp cận và phát triển **zero-day exploits** (lỗ hổng chưa được công bố, không có bản vá). Hoạt động được lên kế hoạch dài hạn, kiên nhẫn, và nhắm vào mục tiêu cụ thể.
- **Nguồn lực:** Ngân sách cấp nhà nước, đội ngũ chuyên biệt (phát triển exploit, triển khai tấn công, phân tích tình báo), cơ sở hạ tầng kỹ thuật riêng.
- **Xu hướng:** Tích cực ứng dụng AI/ML để tự động hóa tấn công và tăng độ tinh vi. Sự bùng nổ của IoT mở rộng bề mặt tấn công đáng kể.
- **Ví dụ thực tế:** Stuxnet (2010) — malware phá hoại máy ly tâm làm giàu uranium của Iran, được cho là do Mỹ và Israel phát triển. APT29 (Cozy Bear) và APT28 (Fancy Bear) của Nga liên quan đến nhiều chiến dịch gián điệp mạng quy mô lớn.

### Phòng thủ

Phòng thủ hiệu quả chống nation-state hacker đòi hỏi cách tiếp cận toàn diện ở cấp độ tổ chức và thường cần sự hỗ trợ từ cơ quan nhà nước (CISA ở Mỹ, NCSC ở Anh, VNCERT tại Việt Nam).

- Áp dụng **Zero Trust Architecture**: không tin tưởng mặc định bất kỳ thực thể nào, kể cả bên trong mạng nội bộ.
- Giám sát liên tục bằng SIEM và threat intelligence feeds.
- Chia sẻ thông tin về mối đe dọa trong ngành (Information Sharing and Analysis Centers — ISACs).
- Kiểm tra bảo mật định kỳ bởi bên thứ ba (red team exercises).

### Hồ sơ tóm tắt

| Tiêu chí         | Mô tả                                                                                     |
| ---------------- | ----------------------------------------------------------------------------------------- |
| Họ là ai?        | Chuyên gia được đào tạo và tài trợ bởi nhà nước, làm việc toàn thời gian                  |
| Mục tiêu của họ? | Mục tiêu chiến lược quốc gia dài hạn: gián điệp, phá hoại, thao túng thông tin            |
| Nguồn lực?       | Ngân sách cấp nhà nước, công cụ tiên tiến bao gồm zero-day exploit, đội ngũ chuyên biệt   |
| Phòng chống?     | Cực kỳ khó; đòi hỏi phòng thủ phối hợp toàn diện và thường cần hỗ trợ từ cơ quan nhà nước |

---

## Malicious insider (Nội gián độc hại)

Nội gián độc hại là nhóm tác nhân đe dọa đặc biệt nguy hiểm vì họ đã có quyền truy cập hợp pháp vào hệ thống và dữ liệu của tổ chức. Không như các tác nhân bên ngoài phải vượt qua các lớp bảo mật, nội gián bắt đầu từ bên trong vành đai phòng thủ.

**Insider threat** bao gồm hai loại chính:

- **Malicious insider:** Người cố ý sử dụng quyền truy cập để gây hại — đây là trọng tâm của phần này.
- **Negligent insider:** Người vô tình gây ra sự cố bảo mật (click vào phishing, cấu hình sai, v.v.) — không có ý định xấu nhưng vẫn là nguồn rủi ro đáng kể.

### Đặc điểm

- **Động cơ:** Đa dạng — lợi ích tài chính (bán dữ liệu, nhận hối lộ), sự cay đắng và trả thù (thường liên quan đến sa thải, thăng tiến bị bỏ qua), danh tiếng, hoặc bị ép buộc/tống tiền bởi bên ngoài.
- **Phương thức:** Thường không cần kỹ năng kỹ thuật đặc biệt — họ dùng chính quyền truy cập và quyền hạn được cấp. Một số dùng kỹ thuật social engineering để mở rộng quyền truy cập.
- **Ví dụ nổi tiếng:** Edward Snowden (2013) — nhà thầu NSA đã sao chép và rò rỉ lượng lớn tài liệu mật của Cơ quan An ninh Quốc gia Mỹ. Đây là một trong những vụ rò rỉ thông tin tình báo lớn nhất trong lịch sử.
- **Dấu hiệu cảnh báo sớm:** Thay đổi hành vi bất thường (làm việc một mình ngoài giờ, truy cập tài nguyên ngoài phạm vi công việc), thể hiện sự bất mãn rõ ràng, sản phẩm làm việc suy giảm, hoặc sao chép lượng lớn dữ liệu ra thiết bị ngoài.

### Phòng thủ

Phòng chống insider threat đòi hỏi kết hợp cả biện pháp kỹ thuật lẫn quản lý con người — không thể chỉ dựa vào một trong hai:

- **Nguyên tắc quyền tối thiểu (Principle of Least Privilege):** Chỉ cấp quyền truy cập đúng mức cần thiết cho công việc. Rà soát và thu hồi quyền không còn cần thiết định kỳ.
- **User and Entity Behavior Analytics (UEBA):** Giám sát và phát hiện bất thường trong hành vi truy cập của người dùng bằng công cụ phân tích tự động.
- **Offboarding process:** Quy trình thu hồi toàn bộ quyền truy cập ngay khi nhân viên nghỉ việc hoặc bị sa thải.
- **Sàng lọc nhân viên:** Kiểm tra lý lịch kỹ khi tuyển dụng, đặc biệt cho vị trí có quyền truy cập dữ liệu nhạy cảm.
- **Văn hóa tổ chức lành mạnh:** Giải quyết sớm các vấn đề nhân sự, tạo kênh phản ánh bất mãn an toàn để giảm nguy cơ hành động trả thù.

### Hồ sơ tóm tắt

| Tiêu chí         | Mô tả                                                                                                            |
| ---------------- | ---------------------------------------------------------------------------------------------------------------- |
| Họ là ai?        | Nhân viên, nhà thầu, hoặc đối tác có quyền truy cập hợp pháp, hành động trái với lợi ích tổ chức                 |
| Mục tiêu của họ? | Trả thù, lợi ích tài chính, hoặc bị ép buộc                                                                      |
| Nguồn lực?       | Không cần ngân sách đặc biệt; sức mạnh đến từ quyền truy cập hợp pháp đã được cấp                                |
| Phòng chống?     | Kết hợp: nguyên tắc quyền tối thiểu, giám sát hành vi (UEBA), quy trình offboarding, và văn hóa tổ chức tích cực |

---

## Lưu ý về phân loại và quy kết (Attribution)

Các phân loại trên không phải ranh giới tuyệt đối trong thực tế:

- Hacktivist có thể tuyển mộ và phối hợp với script kiddie để tăng quy mô tấn công.
- Nation-state hacker đôi khi thuê hoặc hợp tác với criminal gang để tạo khoảng cách phủ nhận hợp lý (plausible deniability).
- Một số tác nhân cố tình giả mạo dấu hiệu kỹ thuật (false flag operation) để che giấu danh tính thực sự.

Những yếu tố này khiến **threat attribution** (quy kết mối đe dọa đúng đối tượng) trở thành một trong những thách thức khó nhất trong lĩnh vực an ninh mạng. Nhiều cuộc điều tra kéo dài nhiều tháng hoặc nhiều năm trước khi có thể quy kết với độ tin cậy cao.

---

## Thông tin bổ sung

### 1. APT — Advanced Persistent Threat

**APT (Advanced Persistent Threat)** là thuật ngữ quan trọng thường liên quan đến nation-state hacker nhưng cũng có thể áp dụng cho criminal gang tinh vi. Đặc điểm của APT:

- **Advanced:** Dùng kỹ thuật tấn công phức tạp, thường bao gồm zero-day exploit và công cụ phát triển riêng.
- **Persistent:** Duy trì hiện diện bí mật trong hệ thống mục tiêu trong thời gian dài (tháng đến năm) mà không bị phát hiện.
- **Threat:** Có mục tiêu cụ thể và rõ ràng, không phải tấn công cơ hội.

Các tổ chức theo dõi và đặt tên các nhóm APT: **MITRE ATT&CK** cung cấp cơ sở dữ liệu toàn diện về chiến thuật, kỹ thuật và thủ tục (TTPs) của các nhóm APT đã biết.

### 2. Threat Intelligence và MITRE ATT&CK Framework

Để phòng thủ hiệu quả, tổ chức cần hiểu cách các nhóm tác nhân đe dọa hoạt động. **MITRE ATT&CK** là framework công khai mô tả chi tiết TTPs của các nhóm tác nhân đe dọa thực tế, giúp tổ chức:

- Mapping phòng thủ hiện tại với các kỹ thuật tấn công đã biết.
- Xác định khoảng trống trong khả năng phát hiện.
- Ưu tiên đầu tư bảo mật dựa trên nhóm tác nhân nguy hiểm nhất với ngành của mình.

### 3. So sánh nhanh các nhóm tác nhân đe dọa

| Nhóm              | Trình độ kỹ thuật | Động cơ chính        | Nguồn lực             | Mối đe dọa điển hình             |
| ----------------- | ----------------- | -------------------- | --------------------- | -------------------------------- |
| Script kiddie     | Thấp              | Danh tiếng, giải trí | Hạn chế               | Khai thác lỗ hổng đã biết        |
| Hacktivist        | Thấp–Trung bình   | Lý tưởng tư tưởng    | Biến động             | DDoS, defacement                 |
| Criminal gang     | Trung bình–Cao    | Tài chính            | Tự túc từ lợi nhuận   | Ransomware, BEC, data theft      |
| Nation-state      | Rất cao           | Chiến lược quốc gia  | Cấp nhà nước          | APT, gián điệp, phá hoại hạ tầng |
| Malicious insider | Biến động         | Tài chính, trả thù   | Quyền truy cập nội bộ | Rò rỉ dữ liệu, phá hoại          |
