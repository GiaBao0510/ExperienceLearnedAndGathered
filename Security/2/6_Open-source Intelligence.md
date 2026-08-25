# Tình báo Nguồn mở (Open-Source Intelligence — OSINT)

## Giới thiệu

Trong thập kỷ qua, **tình báo nguồn mở (OSINT)** ngày càng được quan tâm rộng rãi trong cả khu vực chính phủ lẫn tư nhân.

**OSINT** là quá trình thu thập và phân tích thông tin từ các nguồn có thể truy cập công khai — bao gồm trang web công ty, blog, mạng xã hội, bài báo, hồ sơ công khai của chính phủ và nhiều nguồn khác.

Bất kỳ ai có kết nối internet đều có thể thực hiện OSINT — nhà báo điều tra, nhà nghiên cứu bảo mật, chuyên gia cạnh tranh thương mại, hay cả kẻ tấn công mạng. Không cần kỹ thuật tấn công hay công cụ đặc biệt; thông tin đã ở đó, sẵn sàng được thu thập.

Trong module này, bạn sẽ tìm hiểu:

- Lợi ích của OSINT so với các phương pháp thu thập thông tin khác.
- Các nguồn thông tin mở phổ biến mà kẻ tấn công thường khai thác.
- Nguyên tắc cơ bản khi thực hiện nghiên cứu OSINT.
- Rủi ro mà OSINT đặt ra cho cá nhân và tổ chức.
- Cách thực hành OSINT trên chính bản thân để nhận thức rõ hơn về dấu vết kỹ thuật số của mình.

---

## OSINT so với các phương pháp thu thập thông tin khác

Các phương pháp thu thập thông tin truyền thống thường tốn kém, phức tạp về mặt kỹ thuật và nhiều trường hợp là bất hợp pháp. Ví dụ:

- Xâm nhập trái phép vào cơ sở dữ liệu riêng tư.
- Triển khai chiến dịch phishing để đánh cắp thông tin.
- Cài đặt phần mềm gián điệp (spyware) theo dõi hoạt động người dùng.

Ngược lại, OSINT có những ưu điểm rõ ràng:

|Tiêu chí|Phương pháp xâm nhập|OSINT|
|---|---|---|
|Chi phí|Cao|Thấp hoặc miễn phí|
|Tính hợp pháp|Thường bất hợp pháp|Hợp pháp với thông tin công khai|
|Kỹ năng cần thiết|Cao (hacking, lập trình)|Tương đối thấp|
|Nguy cơ bị phát hiện|Cao|Thấp đến không đáng kể|
|Tốc độ thu thập|Chậm|Nhanh|

> **Ví dụ 1 — Theo dõi vị trí không cần phần mềm gián điệp:**
> 
> Giả sử một nhà báo muốn xác định vị trí của một chính trị gia tại một thời điểm cụ thể. Họ có thể cố gắng cài phần mềm gián điệp vào điện thoại của người đó để lấy tọa độ GPS — nhưng cách đó bất hợp pháp, đòi hỏi kỹ thuật cao và dễ bị phát hiện.
> 
> Một cách đơn giản hơn nhiều: theo dõi mạng xã hội của chính trị gia đó. Chỉ cần một trợ lý vô tình đăng ảnh gắn thẻ địa điểm, hoặc một bức ảnh chứa địa danh dễ nhận biết, là đủ để xác định vị trí. Các đơn vị quân đội và cơ quan thực thi pháp luật đã sử dụng kỹ thuật tương tự trong các điều tra thực tế.

> **Ví dụ 2 — Thu thập thông tin ẩn danh và không để lại dấu vết:**
> 
> Giả sử một kẻ tấn công muốn nghiên cứu hệ thống kiểm soát của một nhà máy điện. Nếu chúng quét mạng bên ngoài của nhà máy, hành động đó có thể bị IDS (Intrusion Detection System) phát hiện và ghi log địa chỉ IP của kẻ tấn công.
> 
> Thay vào đó, chúng có thể tìm thấy một kỹ sư của nhà máy đang thảo luận về kế hoạch kỹ thuật trên một diễn đàn chuyên ngành. Nền tảng diễn đàn có thể ghi log IP của người truy cập, nhưng nhà máy điện thì không — chúng đọc thông tin mà không để lại dấu vết đối với mục tiêu.

---

## Các nguồn thông tin mở phổ biến

Dưới đây là các nguồn mà cả kẻ tấn công lẫn chuyên gia bảo mật thường khai thác khi thực hiện OSINT. Danh sách này không đầy đủ — các nguồn mới xuất hiện liên tục theo sự phát triển của internet.

---

### Trang web công ty

Trang web chính thức là điểm xuất phát hiển nhiên nhất. Nó thường tiết lộ:

- Thông tin liên hệ (email, số điện thoại, địa chỉ trụ sở).
- Cấu trúc tổ chức và tên lãnh đạo.
- Liên kết đến các tài khoản mạng xã hội chính thức.
- Công nghệ đang sử dụng (đôi khi qua mã nguồn HTML, cookie header, hay các file cấu hình vô tình bị công khai).

Kẻ tấn công có thể dùng **Google Dorking** (hay còn gọi là Google Hacking) — sử dụng các toán tử tìm kiếm nâng cao của Google để tìm thông tin nhạy cảm mà tổ chức vô tình để lộ. Ví dụ:

```
site:example.com filetype:pdf "confidential"
site:example.com inurl:admin
```

Ngoài ra, công cụ **Wayback Machine** (web.archive.org) cho phép xem lại các phiên bản cũ của trang web — có thể chứa thông tin đã bị xóa nhưng vẫn còn được lưu trữ trong bộ nhớ đệm.

---

### Báo chí và truyền thông

Các bài báo điều tra, báo cáo ngành, thông cáo báo chí và phỏng vấn có thể chứa nhiều thông tin có giá trị về một tổ chức: chiến lược kinh doanh, đối tác, nhân sự chủ chốt, hoặc thậm chí sự cố bảo mật trong quá khứ.

Các nguồn có giá trị cao bao gồm: báo cáo phân tích ngành, đánh giá từ tổ chức xếp hạng tín dụng, phán quyết tòa án được công khai, và hồ sơ IPO/báo cáo tài chính của công ty niêm yết.

---

### Mạng xã hội

Đây thường là nguồn thông tin phong phú và dễ khai thác nhất. Con người tự nguyện chia sẻ lượng thông tin lớn, và các mảnh thông tin rời rạc có thể được ghép lại để tạo ra bức tranh chi tiết về một cá nhân hoặc tổ chức.

Ví dụ về những gì nhân viên vô tình tiết lộ trên mạng xã hội:

- Ảnh chụp màn hình máy tính có chứa thông tin nhạy cảm trong nền.
- Ảnh thẻ nhân viên (tiết lộ định dạng thẻ, logo, thông tin kiểm soát truy cập).
- Bài đăng "selfie tại văn phòng" tiết lộ bố cục không gian làm việc, thiết bị, màn hình.
- Bài đăng than phiền về công việc tiết lộ quy trình nội bộ, phần mềm đang dùng, hoặc vấn đề bảo mật.

**Từ góc độ tấn công:** Kẻ tấn công phát hiện một nhân viên vừa tham dự hội nghị ngành từ ảnh check-in trên LinkedIn. Chúng soạn email spear phishing với nội dung: "Xin chào, tôi cũng tham dự hội nghị tuần trước và rất muốn kết nối với bạn..." — email này có tính thuyết phục cao vì chứa thông tin thực tế.

---

### Hồ sơ chính phủ và hồ sơ công khai

Nhiều quốc gia yêu cầu các loại thông tin nhất định phải được công khai theo luật định. Các nguồn hữu ích bao gồm:

- **Đăng ký doanh nghiệp:** Tên chủ sở hữu, địa chỉ đăng ký, ngày thành lập, cổ đông.
- **Hồ sơ tòa án:** Các vụ kiện tụng, phá sản, vi phạm hợp đồng.
- **Hồ sơ dân sự:** Ngày sinh, địa chỉ (trong một số hệ thống pháp lý).
- **Báo cáo tài chính:** Các công ty niêm yết trên sàn chứng khoán phải công khai báo cáo tài chính định kỳ.
- **Bằng sáng chế:** Tiết lộ công nghệ mà tổ chức đang phát triển hoặc sở hữu.

---

### Diễn đàn và cộng đồng trực tuyến

Reddit, Stack Overflow, Quora, và các diễn đàn kỹ thuật chuyên ngành là nơi các chuyên gia chia sẻ kiến thức — đôi khi quá nhiều. Một kỹ sư đặt câu hỏi "Làm sao để cấu hình X trong môi trường Y của chúng tôi?" có thể vô tình tiết lộ: kiến trúc hệ thống nội bộ, phần mềm đang dùng, hoặc lỗ hổng đang gặp phải.

Đặc biệt, GitHub và các nền tảng lưu trữ code là nguồn OSINT quan trọng. Developer đôi khi vô tình commit: API key, mật khẩu database, private key, hoặc thông tin môi trường production trong file cấu hình.

> Bổ sung: GitHub là nguồn OSINT quan trọng. Công cụ như `truffleHog` hay `gitleaks` được dùng để tự động quét repository tìm secrets bị lộ.

---

### Tin tuyển dụng

Một tin tuyển dụng bình thường có thể chứa lượng thông tin đáng ngạc nhiên cho kẻ tấn công:

- **Phần mềm yêu cầu:** "Yêu cầu kinh nghiệm với SAP ERP, Oracle Database 19c, VMware vSphere" → tiết lộ toàn bộ stack công nghệ.
- **Vai trò và trách nhiệm:** "Quản lý hệ thống bảo mật OT/SCADA" → cho thấy loại hạ tầng đang vận hành.
- **Dự án sắp tới:** "Tham gia vào dự án chuyển đổi đám mây trong năm 2025" → tiết lộ kế hoạch hạ tầng.
- **Cấu trúc tổ chức:** Tiêu đề vị trí, ai báo cáo cho ai, quy mô đội nhóm.

Kẻ tấn công có thể kết hợp thông tin từ nhiều tin tuyển dụng theo thời gian để hiểu rõ hơn về sự thay đổi công nghệ của mục tiêu.

---

### Bản ghi DNS và thông tin tên miền

Hệ thống DNS (Domain Name System) lưu trữ nhiều thông tin có thể khai thác công khai:

- **WHOIS:** Giao thức truy vấn cơ sở dữ liệu tên miền. Cung cấp thông tin về người đăng ký tên miền (tên, email, tổ chức, địa chỉ). Lưu ý: nhiều tổ chức dùng dịch vụ ẩn danh WHOIS để che giấu thông tin này.
- **DNS records:** Bản ghi MX cho thấy nhà cung cấp email; bản ghi SPF/DKIM tiết lộ dịch vụ gửi email; bản ghi CNAME có thể tiết lộ các dịch vụ third-party đang dùng.
- **Subdomain enumeration:** Các subdomain như `vpn.company.com`, `dev.company.com`, `staging.company.com` có thể tiết lộ cấu trúc hạ tầng nội bộ và môi trường không phải production thường ít được bảo vệ hơn.

Các công cụ như `shodan.io` cho phép tìm kiếm các thiết bị và dịch vụ đang kết nối internet theo địa chỉ IP, tên miền hoặc loại phần mềm — cung cấp cái nhìn trực tiếp về bề mặt tấn công của một tổ chức.

---

## Nguyên tắc khi thực hiện OSINT

Dù bạn thực hiện OSINT với mục đích phòng thủ, báo chí điều tra hay kiểm thử bảo mật hợp pháp, ba nguyên tắc sau sẽ giúp bạn làm việc hiệu quả hơn.

---

### 1. Thu thập rộng trước, lọc sau

Ở giai đoạn đầu, hãy thu thập nhiều thông tin nhất có thể mà không bỏ qua bất cứ thứ gì. Bạn chưa biết mảnh thông tin nào sẽ trở thành chìa khóa quan trọng sau này.

Các công cụ phân tích — đặc biệt là các công cụ tìm kiếm mối liên hệ giữa dữ liệu — hoạt động hiệu quả hơn khi có nhiều thông tin đầu vào hơn. Sau khi có đủ dữ liệu, bạn mới tinh chỉnh và xác định những gì thực sự có giá trị.

> Nguyên tắc: Lưu trữ mọi thứ trước, tinh chỉnh sau. Điều mà bạn nghĩ là vô nghĩa hôm nay có thể trở nên quan trọng khi kết hợp với một mảnh thông tin khác bạn tìm thấy sau.

---

### 2. Đa dạng hóa nguồn, tránh phụ thuộc một chiều

Không phải mọi thông tin trực tuyến đều chính xác. Hồ sơ mạng xã hội có thể bị làm giả, dữ liệu có thể đã lỗi thời, và một nguồn thông tin duy nhất có thể bị chính đối tượng kiểm soát nhằm đánh lừa.

Xác minh thông tin từ ít nhất 2-3 nguồn độc lập trước khi coi đó là sự thật. Việc làm giả một tài khoản mạng xã hội tương đối dễ; làm giả đồng thời nhiều nguồn khác nhau thì khó hơn nhiều.

> Điểm thú vị: Nếu bạn phát hiện đối tượng đã xóa hoặc cố ý che giấu một loại thông tin nào đó, bản thân sự vắng mặt của thông tin đó cũng là dữ liệu có giá trị — nó cho thấy đối tượng muốn giấu điều gì.

---

### 3. Kiên nhẫn và linh hoạt khi gặp bế tắc

OSINT không phải lúc nào cũng cho kết quả ngay. Bạn sẽ gặp bế tắc — nguồn cạn, đường dẫn không đi đến đâu, thông tin mâu thuẫn. Đây là phần bình thường của quá trình.

Khi gặp bế tắc: thay đổi góc tiếp cận, khám phá loại nguồn mới, hoặc tạm gác lại và quay lại sau với góc nhìn mới. Các đội điều tra chuyên nghiệp thường cần nhiều tuần cho một điều tra OSINT phức tạp.

> Lưu ý thực tế: Lượng thông tin công khai về một tổ chức hoặc cá nhân rất khác nhau. Một số tổ chức thực hành **OPSEC (Operations Security)** tốt — họ chủ động kiểm soát và giảm thiểu thông tin công khai về mình. Trong những trường hợp đó, OSINT sẽ mang lại ít kết quả hơn.

---

## Tại sao OSINT quan trọng với tất cả mọi người?

### Vấn đề tổng hợp thông tin (Information Aggregation)

Mỗi mẩu thông tin nhỏ bạn chia sẻ trực tuyến có thể vô hại khi đứng riêng lẻ. Nhưng khi nhiều mảnh thông tin nhỏ được kết hợp lại — một quá trình gọi là **tổng hợp thông tin (information aggregation)** — chúng có thể tiết lộ điều gì đó hoàn toàn khác.

Ví dụ: Nơi làm việc của bạn, lịch đi làm thường ngày, con đường bạn hay đi, thói quen buổi tối — mỗi thứ riêng lẻ có vẻ vô hại. Nhưng kết hợp lại, chúng tạo thành một mô hình hành vi chi tiết có thể bị kẻ xấu lợi dụng.

**Ở cấp độ tổ chức:** Hãy tưởng tượng 100 nhân viên, mỗi người vô tình tiết lộ 1% thông tin nhạy cảm về hệ thống nội bộ — qua LinkedIn, GitHub, diễn đàn, hay tin tuyển dụng. Kẻ tấn công có đủ kiên nhẫn để thu thập và ghép những mảnh 1% đó lại có thể tái tạo được bức tranh toàn cảnh về hạ tầng và quy trình nội bộ của tổ chức.

### Thông tin trực tuyến gần như vĩnh viễn

Một điều nhiều người chưa ý thức đầy đủ: thông tin bạn đăng tải trực tuyến hầu như không thể xóa hoàn toàn. Google cache, Wayback Machine, và các dịch vụ lưu trữ khác có thể giữ lại nội dung ngay cả sau khi bạn đã xóa bài đăng gốc.

### Hàm ý cho tổ chức

Khi xây dựng chính sách bảo mật thông tin, tổ chức cần xem xét OSINT từ góc độ của kẻ tấn công:

- Liệt kê những gì đang được công khai (website, mạng xã hội, hồ sơ công ty, tin tuyển dụng).
- Đánh giá xem những thông tin đó tiết lộ gì về hạ tầng, quy trình và con người.
- Tối thiểu hóa thông tin nhạy cảm trong các tài liệu bắt buộc phải công khai.
- Đào tạo nhân viên về những gì không nên chia sẻ trên mạng xã hội — đặc biệt liên quan đến công việc, công nghệ sử dụng và thông tin nội bộ.

---

## Hoạt động thực hành: Tự kiểm tra dấu vết OSINT của bản thân

Cách tốt nhất để hiểu OSINT là trải nghiệm trực tiếp. Hãy thử tìm kiếm thông tin về chính mình để xem người khác có thể thu thập được gì.

### Hướng dẫn từng bước

**Bước 1 — Chuẩn bị môi trường tìm kiếm sạch**

Mở trình duyệt ở chế độ riêng tư (Incognito/Private Mode) để tránh kết quả tìm kiếm bị ảnh hưởng bởi lịch sử duyệt web và cookie cá nhân của bạn. Điều này giúp mô phỏng góc nhìn của người ngoài không có thông tin về bạn.

**Bước 2 — Tìm kiếm trên công cụ tìm kiếm**

Thử các cụm tìm kiếm sau với tên của bạn:

- `"[Họ tên đầy đủ]"`
- `"[Họ tên]" site:linkedin.com`
- `"[Họ tên]" site:github.com`
- `"[Họ tên]" + "[Tên trường/công ty]"`

Ghi lại những gì bạn tìm thấy: ảnh, bài đăng, hồ sơ, thông tin liên lạc.

**Bước 3 — Kiểm tra các mạng xã hội bạn đang dùng**

Tìm kiếm tên của bạn trên: Facebook, Instagram, LinkedIn, X (Twitter), GitHub, YouTube. Chú ý đến:

- Những ảnh nào có thể tiết lộ thông tin vị trí, môi trường làm việc, hoặc người thân.
- Những bình luận hoặc bài đăng cũ nào bạn không còn muốn công khai.
- Có thông tin liên hệ (số điện thoại, email) bị lộ không?

**Bước 4 — Thử từ góc độ người ngoài (tùy chọn)**

Nhờ một người bạn hoặc thành viên gia đình lặp lại các bước trên, không cần bạn hướng dẫn. Xem họ tìm thấy gì mà bạn không chú ý, và họ đã dùng phương pháp gì.

---

### Câu hỏi phản tư sau hoạt động

Sau khi hoàn thành tìm kiếm, hãy suy ngẫm về những câu hỏi sau:

1. Bạn tìm thấy những loại thông tin nào về bản thân? (Ảnh, thông tin liên lạc, lịch sử công việc, địa điểm...)
2. Có thông tin nào khiến bạn ngạc nhiên hoặc lo lắng không? Tại sao?
3. Nếu bạn là kẻ tấn công, bạn sẽ dùng những thông tin đó để tấn công bạn theo cách nào? (Spear phishing? Vishing? Tiếp cận vật lý?)
4. Bạn cần thay đổi thiết lập quyền riêng tư hoặc xóa bỏ thông tin gì sau bài tập này?
5. Nếu 100 đồng nghiệp của bạn làm tương tự, tổ chức của bạn sẽ lộ ra bao nhiêu thông tin khi gộp lại?

---

## Tóm tắt

|Khái niệm|Nội dung chính|
|---|---|
|OSINT là gì|Thu thập thông tin từ các nguồn công khai — hợp pháp và có thể thực hiện bởi bất kỳ ai|
|Ưu điểm so với tấn công kỹ thuật|Chi phí thấp, ít rủi ro bị phát hiện, không cần kỹ năng hacking|
|Nguồn thông tin phổ biến|Website, mạng xã hội, báo chí, hồ sơ công khai, DNS, diễn đàn, tin tuyển dụng, GitHub|
|Nguyên tắc thu thập|Thu thập rộng → đa dạng hóa nguồn → kiên nhẫn với bế tắc|
|Rủi ro cho cá nhân/tổ chức|Tổng hợp thông tin, thông tin không thể xóa hoàn toàn, nhân viên vô tình tiết lộ|
|Ứng dụng hợp pháp|Điều tra báo chí, kiểm thử bảo mật (pentest), nghiên cứu cạnh tranh, phòng thủ bảo mật|
