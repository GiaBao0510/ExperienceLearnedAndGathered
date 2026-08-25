# Các Vai trò trong lĩnh vực An ninh mạng

## Giới thiệu

Module này giới thiệu các vai trò phổ biến trong lĩnh vực an ninh mạng. Bạn sẽ tìm hiểu sự khác biệt giữa: chuyên viên phân tích an ninh mạng (Security Analyst), kỹ sư an ninh mạng (Security Engineer), kiến trúc sư an ninh mạng (Security Architect), lập trình viên phần mềm bảo mật (Security Software Developer), chuyên viên kiểm thử xâm nhập (Penetration Tester), kỹ sư bảo mật ứng dụng (Application Security Engineer), chuyên viên phân tích phần mềm độc hại (Malware Analyst), và điều tra viên pháp y kỹ thuật số (Digital Forensic Investigator).

---

## Các vai trò trong lĩnh vực An ninh mạng

Nhu cầu cao về an ninh mạng đồng nghĩa với nguồn cung việc làm dồi dào trong lĩnh vực này. Theo dữ liệu từ LinkedIn, các vị trí an ninh mạng — bao gồm kỹ sư, quản lý, chuyên viên phân tích và chuyên gia — nằm trong nhóm nghề nghiệp tăng trưởng nhanh nhất tại nhiều quốc gia.

Dưới đây là các vị trí an ninh mạng được săn đón nhiều nhất, cùng với phân tích sự khác biệt giữa các vai trò dễ gây nhầm lẫn với nhau.

---

### Chuyên viên Phân tích An ninh mạng (Security Analyst)

Chuyên viên phân tích an ninh mạng đóng vai trò tuyến đầu — liên tục "đi trước một bước" so với tội phạm mạng để bảo vệ hệ thống và mạng lưới khỏi các mối đe dọa.

**Nhiệm vụ và trách nhiệm:** Phạm vi công việc khá rộng, nhưng trách nhiệm cốt lõi là bảo vệ hệ thống, mạng lưới và dữ liệu của tổ chức. Công việc hàng ngày bao gồm:
- Giám sát hệ thống để phát hiện hoạt động bất thường hoặc dấu hiệu vi phạm an ninh (thường qua công cụ SIEM).
- Điều tra và xử lý cảnh báo bảo mật trước khi chúng gây thiệt hại nghiêm trọng.
- Xây dựng và cập nhật chiến lược phòng thủ.

**Cấp độ:** Tùy cấu trúc tổ chức, có thể là vị trí entry-level, mid-level, hoặc senior. Đây thường là **điểm khởi đầu phổ biến nhất** cho người mới vào ngành an ninh mạng, kể cả những ai chuyển ngành từ IT hoặc phát triển phần mềm.

---

### Kỹ sư An ninh mạng (Security Engineer)

Kỹ sư an ninh mạng bảo vệ hệ thống máy tính và mạng khỏi các mối đe dọa — nhưng khác với Security Analyst ở trọng tâm công việc: **xây dựng** thay vì chỉ giám sát và phản ứng.

**Nhiệm vụ và trách nhiệm:**
- Thiết kế, xây dựng và duy trì kiến trúc bảo mật để bảo vệ thông tin nhạy cảm.
- Giám sát liên tục cơ sở hạ tầng, điều tra vi phạm an ninh khi xảy ra.
- Tiến hành kiểm tra và mô phỏng định kỳ để phát hiện và giảm thiểu lỗ hổng.

**Cấp độ:** Thường là vị trí mid-level hoặc senior.

> **Bổ sung phân biệt quan trọng — Analyst vs. Engineer:** Bản gốc mô tả hai vai trò này khá tương đồng nhau, khiến người đọc khó phân biệt rõ ràng. Điểm khác biệt cốt lõi trong thực tế: **Security Analyst** thiên về vận hành (operations) — giám sát, phát hiện, phản ứng với sự cố hàng ngày, thường làm việc trong SOC. **Security Engineer** thiên về xây dựng (engineering) — thiết kế và triển khai hệ thống, công cụ, và hạ tầng bảo mật mà Analyst sẽ sử dụng để giám sát. Nói cách khác: Engineer xây "hàng rào và camera", Analyst theo dõi màn hình camera và phản ứng khi có báo động. Với background lập trình (như Backend/Golang Developer), vai trò Security Engineer thường là lộ trình chuyển đổi tự nhiên hơn vì đòi hỏi kỹ năng code và thiết kế hệ thống.

---

### Kiến trúc sư An ninh mạng (Security Architect)

Kiến trúc sư an ninh mạng đóng vai trò tương tự kiến trúc sư của một tòa nhà — nhưng đối tượng thiết kế là **mạng máy tính** của tổ chức.

**Nhiệm vụ và trách nhiệm:**
- Thiết kế mạng máy tính an toàn có khả năng chống lại nhiều loại mối đe dọa khác nhau.
- Phát triển và triển khai biện pháp an ninh, cập nhật liên tục khi có mối đe dọa mới.
- Hỗ trợ khôi phục mạng khi xảy ra vi phạm an ninh nghiêm trọng.

**Cấp độ:** Thường là vị trí mid-level hoặc senior — trên thực tế, đây thường là vị trí **cao cấp nhất trong nhóm kỹ thuật thuần túy**, đòi hỏi tầm nhìn tổng thể về toàn bộ hệ thống của tổ chức.

> **Bổ sung phân biệt quan trọng — Engineer vs. Architect:** Cả hai vai trò đều liên quan đến xây dựng hệ thống bảo mật, nhưng khác nhau về **phạm vi và tầm nhìn**. Security Engineer thường tập trung triển khai và vận hành các giải pháp cụ thể (ví dụ: cấu hình firewall, triển khai một công cụ SIEM cụ thể). Security Architect tập trung vào **thiết kế tổng thể** — quyết định chiến lược bảo mật ở cấp độ toàn tổ chức, chọn công nghệ nào phù hợp, và đảm bảo các thành phần khác nhau (do nhiều Engineer xây dựng) hoạt động nhất quán với nhau. Đây là lý do Architect thường là vị trí có kinh nghiệm dày dạn hơn Engineer trong đa số tổ chức.

---

### Lập trình viên Phần mềm Bảo mật (Security Software Developer)

Lập trình viên phần mềm bảo mật thiết kế và phát triển phần mềm an toàn ngay từ đầu để bảo vệ hệ thống khỏi các mối đe dọa tiềm ẩn.

**Nhiệm vụ và trách nhiệm:**
- Tích hợp biện pháp bảo mật vào **mọi giai đoạn** của vòng đời phát triển phần mềm (SDLC — Software Development Life Cycle).
- Cập nhật và vá lỗi phần mềm hiện có để tăng cường khả năng phòng thủ trước mối đe dọa mới.

**Cấp độ:** Thường là vị trí mid-level hoặc senior.

---

### Chuyên viên Kiểm thử Xâm nhập (Penetration Tester)

Chuyên viên kiểm thử xâm nhập — thường gọi tắt là **pen tester** — là một dạng nhà nghiên cứu bảo mật tấn công (offensive security), còn được gọi là **hacker đạo đức (ethical hacker)**.

**Nhiệm vụ và trách nhiệm:**
- Mô phỏng kỹ thuật tấn công thực tế để tìm lỗ hổng trong hệ thống mà kẻ tấn công thực sự có thể khai thác.
- Xác nhận (validate) các lỗ hổng phát hiện qua công cụ quét lỗ hổng tự động có phải là điểm yếu thực sự hay chỉ là false positive.
- Phát hiện thêm các lỗ hổng mà công cụ quét tự động bỏ sót — đặc biệt là lỗ hổng logic nghiệp vụ (business logic flaw) mà công cụ tự động khó phát hiện.
- Báo cáo kết quả chi tiết cho tổ chức để khắc phục trước khi bị kẻ tấn công thực sự khai thác.

**Cấp độ:** Tùy cấu trúc tổ chức, có thể là entry-level, mid-level, hoặc senior.

> **Bổ sung thông tin quan trọng bị thiếu — Yêu cầu ủy quyền hợp pháp:** Bản gốc không đề cập một điểm cực kỳ quan trọng về mặt đạo đức và pháp lý của nghề này: mọi hoạt động kiểm thử xâm nhập **chỉ được thực hiện khi có văn bản ủy quyền rõ ràng (scope of work, rules of engagement)** từ tổ chức sở hữu hệ thống. Đây là ranh giới phân biệt pen tester hợp pháp (white hat) với hành vi xâm nhập trái phép (black hat) — dù kỹ thuật sử dụng có thể giống hệt nhau. Các chứng chỉ phổ biến của nghề này bao gồm **OSCP (Offensive Security Certified Professional)** và **CEH (Certified Ethical Hacker)**.

---

### Kỹ sư Bảo mật Ứng dụng (Application Security Engineer)

Giống Security Software Developer, kỹ sư bảo mật ứng dụng cũng quan tâm đến bảo mật phần mềm. Nhưng khác biệt cốt lõi: Security Software Developer **trực tiếp viết code** phần mềm; Application Security Engineer (thường viết tắt **AppSec Engineer**) **kiểm tra và duy trì** các hệ thống, quy trình được dùng để phát triển và vận hành phần mềm đó.

**Nhiệm vụ và trách nhiệm:**
- Làm việc chặt chẽ với đội ngũ phát triển xuyên suốt SDLC, triển khai biện pháp bảo mật ở mọi giai đoạn.
- Review, quét, và kiểm tra mã nguồn (code review, SAST/DAST scanning) để phát hiện lỗ hổng trước khi triển khai lên production.

**Cấp độ:** Thường là vị trí mid-level.

> **Bổ sung thông tin kỹ thuật quan trọng cho Backend Developer:** Đây là vai trò có mối liên hệ **trực tiếp nhất** với công việc Backend Engineer trong toàn bộ danh sách này — AppSec Engineer thường làm việc cùng đội backend hàng ngày để review pull request, thiết lập pipeline CI/CD tích hợp bảo mật (DevSecOps), và hướng dẫn thực hành **secure coding** (ví dụ: cách phòng chống SQL injection, cách quản lý secret đúng cách). Với nền tảng Golang, việc nắm vững `crypto` package chuẩn, secure coding practice, và công cụ như `gosec` (static analysis cho Go) là bước đệm tốt để chuyển hướng sang vai trò này.

---

### Chuyên viên Phân tích Phần mềm Độc hại (Malware Analyst)

Chuyên viên phân tích phần mềm độc hại được đào tạo chuyên sâu để xử lý các mối đe dọa từ malware.

**Nhiệm vụ và trách nhiệm:**
- Điều tra các tệp hoặc chương trình bị nghi ngờ là phần mềm độc hại.
- Nghiên cứu kỹ lưỡng để hiểu cơ chế hoạt động và tìm cách ngăn chặn.
- Đòi hỏi kiến thức sâu về máy tính, ngôn ngữ lập trình (đặc biệt là Assembly và C), và khả năng giải quyết vấn đề.
- Liên tục cập nhật kiến thức về mối đe dọa và kỹ thuật kiểm soát malware mới nhất.

**Cấp độ:** Thường là vị trí mid-level hoặc senior.

> **Bổ sung kỹ thuật chuyên môn quan trọng bị thiếu:** Bản gốc mô tả công việc khá chung chung mà không nêu **kỹ thuật cốt lõi** mà nghề này sử dụng — đây là thông tin quan trọng để người đọc hình dung đúng bản chất công việc:
> - **Static Analysis (Phân tích tĩnh):** Kiểm tra mã của malware mà **không thực thi** nó — dùng công cụ như disassembler (IDA Pro, Ghidra) để đọc mã máy.
> - **Dynamic Analysis (Phân tích động):** **Thực thi** malware trong môi trường cô lập an toàn (sandbox) để quan sát hành vi thực tế — thay đổi file, kết nối mạng, thay đổi registry.
> - **Reverse Engineering (Kỹ thuật đảo ngược):** Kỹ năng nền tảng để hiểu logic bên trong của một chương trình mà không có mã nguồn gốc — đây là kỹ năng cốt lõi phân biệt Malware Analyst với hầu hết vai trò khác trong danh sách.
>
> Đây là một trong những vai trò **chuyên sâu kỹ thuật nhất** và có rào cản gia nhập cao nhất trong ngành an ninh mạng, thường đòi hỏi nền tảng vững về kiến trúc máy tính và assembly language.

---

### Điều tra viên Pháp y Kỹ thuật số (Digital Forensic Investigator)

Điều tra viên pháp y kỹ thuật số đóng vai trò như một "thám tử kỹ thuật số".

**Nhiệm vụ và trách nhiệm:**
- Thu thập, kiểm tra và phân tích bằng chứng kỹ thuật số phục vụ điều tra tội phạm mạng.
- Công việc thường bắt đầu **sau khi** một sự cố hoặc tội phạm mạng đã xảy ra (khác với các vai trò phòng thủ chủ động khác).
- Sử dụng công cụ pháp y chuyên dụng và hiểu biết về các framework tấn công để kiểm tra dữ liệu, khôi phục tệp bị mất/hỏng, và truy vết dấu vết kỹ thuật số về nguồn gốc tội phạm.

**Cấp độ:** Tùy cấu trúc tổ chức, có thể là entry-level, mid-level, hoặc senior.

> **Bổ sung yêu cầu nghiệp vụ quan trọng bị thiếu:** Bản gốc không đề cập một đặc thù quan trọng của nghề này: bằng chứng thu thập bởi điều tra viên pháp y kỹ thuật số thường có thể được sử dụng **làm bằng chứng pháp lý tại tòa án**. Vì vậy, công việc đòi hỏi tuân thủ nghiêm ngặt quy trình **chuỗi lưu giữ bằng chứng (chain of custody)** — ghi chép đầy đủ ai đã tiếp cận bằng chứng, khi nào, và bằng cách nào — để đảm bảo tính toàn vẹn pháp lý của bằng chứng không bị nghi ngờ. Đây là điểm khác biệt căn bản so với các vai trò kỹ thuật thuần túy khác trong danh sách, vì công việc giao thoa trực tiếp với quy trình pháp lý.

---

## Bảng So sánh Tổng hợp

Để giúp hình dung nhanh sự khác biệt giữa các vai trò, dưới đây là bảng tổng hợp theo hai trục: **tính chất công việc** (phòng thủ chủ động/vận hành/tấn công/điều tra) và **cấp độ kinh nghiệm điển hình**.

| Vai trò | Tính chất công việc | Cấp độ điển hình | Liên hệ gần nhất với Backend Dev |
|---|---|---|---|
| Security Analyst | Vận hành, giám sát, phản ứng | Entry–Senior | Trung bình |
| Security Engineer | Xây dựng hạ tầng bảo mật | Mid–Senior | Cao |
| Security Architect | Thiết kế chiến lược tổng thể | Mid–Senior | Trung bình |
| Security Software Developer | Viết code phần mềm bảo mật | Mid–Senior | Rất cao |
| Penetration Tester | Tấn công mô phỏng (offensive) | Entry–Senior | Trung bình |
| Application Security Engineer | Review, kiểm tra quy trình phát triển | Mid | Rất cao |
| Malware Analyst | Phân tích chuyên sâu kỹ thuật | Mid–Senior | Thấp |
| Digital Forensic Investigator | Điều tra sau sự cố | Entry–Senior | Thấp |

> Bản gốc trình bày các vai trò tuần tự mà không có công cụ đối chiếu tổng hợp — với 8 vai trò dễ nhầm lẫn lẫn nhau, một bảng so sánh trực quan giúp người đọc nắm bắt và ghi nhớ nhanh hơn nhiều so với chỉ đọc tuần tự.

---

## Lưu ý về tính linh hoạt của các vai trò

Có rất nhiều vị trí công việc khác nhau trong lĩnh vực an ninh mạng. Trách nhiệm của các vị trí đôi khi chồng chéo nhau, và chức danh công việc có thể khác nhau tùy quy mô và cấu trúc tổ chức — một công ty nhỏ có thể gộp nhiều vai trò vào một vị trí duy nhất, trong khi tập đoàn lớn có thể chia nhỏ thành nhiều vai trò chuyên biệt hơn. Khi ngành này tiếp tục phát triển, các vị trí mới cũng sẽ xuất hiện, phản ánh bản chất luôn thay đổi của thế giới an ninh mạng.

---

## Bổ sung kiến thức

### 1. Lộ trình chuyển đổi từ Backend Developer sang An ninh mạng

Với nền tảng Backend/Golang, con đường chuyển đổi tự nhiên nhất — theo thứ tự độ khó tăng dần — thường là:

1. **Application Security Engineer** — tận dụng trực tiếp kinh nghiệm code, chỉ cần bổ sung kiến thức về lỗ hổng phổ biến (OWASP Top 10) và công cụ SAST/DAST.
2. **Security Engineer** — mở rộng thêm kiến thức về hạ tầng, network, và các công cụ bảo mật chuyên dụng ngoài phạm vi code.
3. **Penetration Tester** (hướng offensive) — nếu quan tâm tư duy tấn công, cần học thêm về khai thác lỗ hổng chuyên sâu.
4. **Security Architect** — thường là bước phát triển tự nhiên sau nhiều năm kinh nghiệm ở vai trò Engineer.

### 2. Chứng chỉ tương ứng với từng vai trò

Một góc nhìn thực tế bổ sung mà bản gốc chưa đề cập — các chứng chỉ phổ biến giúp định hướng học tập cho từng vai trò:

| Vai trò | Chứng chỉ phổ biến |
|---|---|
| Security Analyst | CompTIA Security+, GSEC |
| Security Engineer | GCIH, CCSP (nếu thiên về cloud) |
| Security Architect | CISSP, SABSA |
| Penetration Tester | OSCP, CEH, eJPT |
| Application Security Engineer | GWEB, CSSLP |
| Malware Analyst | GREM (GIAC Reverse Engineering Malware) |
| Digital Forensic Investigator | GCFE, EnCE |

### 3. Vai trò không nằm trong danh sách gốc nhưng đang tăng trưởng nhanh

Tài liệu gốc liệt kê 8 vai trò truyền thống nhưng bỏ sót một số vai trò mới nổi đang có nhu cầu tăng nhanh, đặc biệt liên quan đến bối cảnh công nghệ hiện tại:

- **Cloud Security Engineer:** Chuyên biệt hóa bảo mật cho hạ tầng cloud (AWS, GCP, Azure) — khác biệt đáng kể so với bảo mật hạ tầng on-premise truyền thống.
- **DevSecOps Engineer:** Tích hợp bảo mật trực tiếp vào pipeline CI/CD, thường yêu cầu kỹ năng tự động hóa và scripting mạnh.
- **AI Security Engineer:** Vai trò mới nổi, tập trung bảo mật cho hệ thống AI/ML (ví dụ: chống prompt injection, data poisoning).

Các vai trò này thường có nhu cầu tuyển dụng tăng trưởng nhanh hơn các vai trò truyền thống trong vài năm gần đây, do gắn liền với xu hướng dịch chuyển hạ tầng công nghệ hiện tại.