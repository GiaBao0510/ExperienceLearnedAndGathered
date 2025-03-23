## **CI/CD là gì?**

**CI/CD** được viết tắt từ **Continuous Integration** và **Continuous Delivery/ Continuous Deyploment**, có nghĩa là ==tích hợp liên tục và chuyển giao/triển khai liên tục==. CI/CD là quá trình phát triển phần mềm và chuyển giao sản phẩm có tích hợp sẳn các phương phức tự động hóa, giúp cho doanh nghiệp rút ngắn thời gian triển khai, hoàn thiện và cung cấp sản phẩm đến tay người đung một cách nhanh chóng.

Quy trình triển khai CI/CD bao gồm: Đẩy code lên server (Developer commit code), CI/CD tự động chạy build, chạy test deploy sản phầm, chuyển giao sản phẩm đến tay người dùng.

![](https://statics.cdn.200lab.io/2024/10/ci-cd-la-gi-ci-cd-1.png?width=1200)

---
## **CI là gì?**

#### 1.**Định nghĩa CI:**

**CI (Continuous Integration)** là quá trình liên tục kiếm tra và tích hợp code mới vào chính code của một dự án. Mục tiêu chính của **CI** là đảm bảo rằng bất kỳ sự thay đổi nào trong **code** (*Chẳng hạn như sửa lỗi , thêm tính năng mới*) đều được kiểm tra ngay lập túc để phát hiện lỗi trước khi nó gây ản hưởng xấu lên hệ thống.

![](https://statics.cdn.200lab.io/2024/10/ci-cd-la-gi-ci-1.png?width=1200)

**CI** hoạt động ==như một hệ thống kiểm tra tự động==. Mỗi khi thay đổi code và commit lên hệ thống, các bài kiểm tra tự động sẽ được chạy để kiểm tra xem code có hoạt động đúng và gây ra lỗi hay không
#### 2.**Những bước cơ bản trong CI:**

- **Functional tests:** Kiểm tra xem ==phần mềm có hoạt động như mong đợi không==. Đây thường là các bài kiểm tra tự động mà đội ngũ **QA (Quality Assurance)** phát triển để đảm bảo cho tính năng hoạt động đúng như mong đợi. *VD:* Người dùng có thể thực hiện được việc thêm sản phẩm vào giỏ hàng và thực hiện thanh toán thành công hay không.
- **Security scans:** Kiểm tra code có lỗ hổng bảo mật nào hay không, chẳng hạn như khả năng bị [SQL Injection](https://www.w3schools.com/sql/sql_injection.asp) hay  [XSS](https://200lab.io/blog/giai-phap-bao-mat-giao-dien-web-ngan-chan-tan-cong-xss-va-csrf-hieu-qua/) (Cross-Site Scripting).
- **Code quality scans:** Đảm bảo code tuân thủ các tiêu chuẩn. Ví dụ như độ dài hàm, cách sử dụng khoảng trắng và các quy tắc coding style. ***VD:*** Tiêu chuẩn đặt tên biến.
- **Performance tests:** Kiểm tra xem code có đáp ứng được yêu cầu hiệu suất không, chẳng hạn như thời gian xử lý yêu cầu hoặc khả năng chịu tải.
- **License scanning:** Kiểm tra xem tất cả các thư viện hoặc công cụ sử dụng có giấy phép phù hợp hay không, để tránh các vấn đề pháp lý
- **Fuzz testing:** Gửi các dữ liệu để kiểm tra xem nó có bị crash hay gặp lỗi bất ngờ không.

#### 3.**Lợi ích của việc thiết lập CI:**

CI giúp phát hiện sớm các lỗi tiềm ẩn, tiết kiệm rất nhiều thời gian và công sức. Vậy việc sửa một lỗi ngay khi phát hiện sẽ dễ dàng hơn nhiều so với khi lỗi đó đã tích hợp vào nhiều phần khác của hệ thống.

***Ví dụ:*** Giả sử đang phát triển một tính năng mới cho ứng dụng. Thay vì đợi đến cuối quá trình phát triển để kiểm, CI cho phép thực hiện các bài kiểm tra ngay khi viết code. Nhờ đó, nếu có lỗi xuất hiện, thì có thể phát hiện ngay lập tức và sửa chữa ==khi vấn đề còn nhỏ==, thay vì đợi đến khi tính năng đã hoàn thành và **tích hợp** với các phần khác.

![](https://statics.cdn.200lab.io/2024/10/ci-cd-la-gi-devops.jpg?width=1200)

**CI** giúp toàn bộ đội nhóm đều có thể thấy tình trạng hiện tại của phần mềm, tất cả các thành viên từ lập trình viên, quản lý, đến **QA** (Kiểm thử chất lượng), **UX** (Trải nghiệm người dùng), và thậm chí là bộ phận bảo mật có thể theo dõi tiến trình và trạng thái của dự án. Từ đó, mọi người có thể điều chỉnh công việc của mình để phù hợp với tình trạng thực tế.

***Ví dụ:*** Khi các bài kiểm tra bảo mật được thực hiện thường xuyên, cả nhóm sẽ nhận biết được dự án có gặp phải các vấn đề bảo mật mới không. Nếu CI phát hiện rằng một đoạn code mới làm lộ lỗ hỏng bảo mật, thì không chỉ các lập trình viên mà các quản lý dự án cũng biết ngay lập tức. Nhóm sẽ có thể phân bổ nguồn lực để sửa lỗi hoặc điều chỉnh kế hoạch phát triển. 

---
## **CD là gì?**

##### **Định nghĩa CD:**

CD là quá trình đưa code từ noi bạn viết lên một nơi mà người dùng có thể sử dụng được. Như là Website hoặc mobileapp. Có 2 loại CD chính:

- ***Continuous Delivery (Phân phối liên tục):*** Đảm bảo code luốn sẳn sàng để được triển khai bất cứ lúc nào. ==Sau khi code đã vượt qua các kiểm tra CI==, nó được đặt ở trạng thái sẳn sàng để triển khai. Tuy nhiên, việc triển khai có thể được thực hiện thủ công.
- ***Continuous Deployment (Triển khai liên tục):*** Tự động triển khai code ngay khi nó vượt qua tất cả các kiểm tra mà không cần sự can thiệp thủ công, hệ thống sẽ tự động đưa code vào môi trường production.

![](https://statics.cdn.200lab.io/2024/10/ci-cd-la-gi-cd.png?width=1200)

##### **Những bước cơ bản trong CD:**

- **Build:** Trước khi triển khai code, đôi khi chúng ta ==cần phải thực hiện quá trình build== cho các ngôn ngữ biên dịch (compiled language) như C, C++, Java,...
- **Deploy:** Triển khai code: Đẩy Docker image lên repository, dùng dòng lệnh đưa code lên AWS,... Triển khai có thể tự động hoặc thủ công .Tùy thuộc vào cách mà mình thiết lập.

***Ví dụ:*** Khi bạn hoàn thành tính năng mới như giỏ hàng, CD sẽ tự động kiểm tra và đóng gói tính năng này rồi triển khai lên website thật để khách hàng có thể sử dụng ngay lập tức.

##### **Lợi ích của việc thiết lập CD:**

**Mục tiêu của CD** là giúp việc triển khai và phát hành (release) phần mềm trở nên ==dễ dàng và có thể diễn ra thường xuyên== mà không gặp nhiều rủi ro. Khi **CD** được thiết lập, mỗi lần bạn commit code, code có thể tự động được triển khai vào các môi trường như: review enviroment, Staging eviroment hay product enviroment

Các môi trường khác ngoài production (như review và staging) hoạt động giống như những phép thử, nhằm đảm bảo code chạy ổn định trước khi release chính thức cho người dùng.

---
## **Sự khác biệt của hệ thống có cài đặt?**

Thống thường thì khi, các **Dev** ==sau khi code xong sẽ phải build trực tiếp trên máy tính==. Sau đó, tải bản build lên máy chủ và khởi động lại ứng dụng để chấp nhận bản build mới. Tiếp đến, ==Dev cần thông báo cho **QA/QC**== để tiến hành kiểm thử sản phẩm.

Nếu xảy ra bất kỳ sự sai xót nào, quy trình trên sẽ được thực hiện lại từ đầu. Như vậy các **Dev** và **QA/QC** phải ==mất rất nhiều thời gian để triển khai dự án==, ảnh hưởng đến tiến độ và hiệu suất làm việc. Ngoài ra, việc **Dev** phát triển tính năng mới cũng có thể làm hỏng tính năng cũ trên ứng dụng, hoặc phát sinh một số lỗi mà chỉ khi thực hiện deply tính năng mới có thể nhận biết được.

Bằng cách áp dụng **CI/CD**, ==quy trình phát triển phần mềm trở nên tiện lợi và nhanh== gọn hơn bao giờ hết, Các **Developer** chỉ cần **commit code**, các bước như chạy **build**, **test** và **deploy** sẽ được thực hiện một cách tự động. Kết hợp **automation test** giúp cho quá trình trở nên chặt chẽ và hạn chế các trường hợp phát sinh lỗi.

| Tiêu chí          | Không có CI/CD                      | có CI/CD                         |
| ----------------- | ----------------------------------- | -------------------------------- |
| Kiểm tra lỗi      | Kiểm tra thủ công, dễ bỏ sót lỗi    | Tự động kiểm tra ngay khi commit |
| Tốc độ triển khai | Chậm do cần nhiều thao tác thủ công | Nhanh chóng, tự động hóa         |
| Chất lượng code   | Dễ gặp lỗi trong quá trình merge    | Code được kiểm soát tốt hơn      |

---
## **Cách thức hoạt động của CI/CD?**

Khi **Developer** thực hiện thay đổi trên các **Repository** có chứa mã nguồn, các **repositories** sẽ phát thông báo đến hệ thống ***CI/CD***. Lúc này, ***CI/CD*** sẽ tự động thực hiện các thao tác đã cài đặt từ trước dựa trên các nhánh của **Repository**.

Sau khi hoàn tất các thao tác trên, **CI**/CD sẽ cập nhật kết quả. **Developers** cần kiểm tra lại các bước thực hiện trên **CI/CD**, hoặc bản **source** có chạy ổn định hay không. **Reviewer** nên dựa vào **pipeline** đã build để kiểm tra tính khả dụng của phần mềm, đồng thời đánh giá xem những thay đổi mới trên sản phẩm có đạt hiệu quả không

---
## **Ưu/Nhược điểm?**

##### **Ưu điểm:**
- **Tránh được những lỗi không đáng có:** Chẳng hạn như lỗi compile (Khi đẩy code lên) hoặc các lỗi phát sinh liên quan đến môi trường build sản phẩm. ***Ví dụ**: Khi làm thủ công, cùng 1 source code nhưng sẽ có sự khác biệt khi bạn A build trên máy bạn A, bạn B build trên máy bạn B.*
- **Đảm bảo logic:**  (Vì quy trình CI/CD có phần mềm automation test), khi Developer xây dựng tính năng mới sẽ ==không gây ảnh hưởng đến tính năng cũ==.
- **Giúp tập trung vào công việc:** Bởi quy trình CI/CD mang tính tự động cao nên Developer ==không cần phải thực hiện việc build và deploy== phần mềm/ứng dụng trên máy cá nhân.
- **Nâng cao chất lượng code:** Thông qua quy trình, Developer có thể **cài đặt những ràng buộc ngay từ đầu**. ***Ví dụ:** pull request khi được tạo ra thì không được quá lớn, không được vượt quá X thay đổi,... Điều này góp phần giúp chất lượng pull request ngày càng tốt hơn.* 
- **Phát triển kỹ năng unit test cho Develoopers:** Thông qua các chỉ số ràng buộc về code coverage (% code đã được cover) được cài đặt trong quy trình CI/CD. Nghĩa là khi phát triển tính năng mới, để không làm giảm chỉ số code coverage. Developer phải ý thức được ==tầm quan trọng của unit test== và chủ động học hỏi, nâng cao kỹ năng liên quan.
- **Tối ưu tốc độ phát triển sản phẩm:** thông qua việc ==theo dõi thời gian build pipeline== (các bước chạy test, build, chạy static code analytics (lint check)).

##### **Nhược điểm:**
- Trong một dự án nếu có quá nhiều **Developer** cũng tham gia phát triển sản phẩm, sẽ có thời điểm ==phát sinh nhiều pull request== cần được **merge** vào **brach**. Lúc này, các thành viên phải ==chờ pull request của người trước== được merge hoàn tất, sau đó thực hiện update (cập nhật) lại sourrce (trong trường hợp có thống báo ==conflict từ Git repository==) và phải trải qua các bước test lại từ đầu. Hệ quả sẽ làm gián đoạn thời gian phát triển sản phẩm.
- Vì sử dụng dịch vụ CI/CD của bên service thứ 3 nên nếu service đó gặp vấn đề và bị crash (khai tử), thì ==những dự án liên quan đến việc áp dụng CI/CD cũng bị ảnh hưởng nghiêm trọng.==  

---
## **Khi nào nên áp dụng và khi nào không nên áp dụng quy trình CI/CD?**

Các tổ chức nên áp dụng quy trình **CI/CD** vào việc tích hợp nên thực hiện càng sớm càng tốt. Khi quy trình tốt thì chất lượng công việc **Developer** cũng tối ưu hơn. Kể cả ==khi dự án theo cá nhân thì vẫn nên tích hợp **CI/CD**== (có thể lựa chọn những service miễn phí) để tận dụng những ưu điểm đã được liệt kê.

==Tuy nhiên==, trong một số trường hợp như: ==tổ chức không có người có đủ khả năng vận hành quy trình **CI/CD**==, Developer chưa làm chủ và chưa nắm rõ về tool hoặc không biết làm sao để đảm bảo quy trình **CI/CD**,... thì có thể nên cân nhắc việc chưa sử dụng. Vì nếu có vấn đề không mong muốn xảy ra nhưng lại đưa cho người không có đủ kiến thức chuyển môn để xử lý thì sẽ ==mất nhiều thời gina và gây ra những gián đoạn không cần thiết.==

---
## **Các nguyên tắc khi triển khai quy trình CI/CD:**

1. Xác định team phù hợp để triển khai quy trình **CI/CD** đầu tiên.
2. Bắt đầu triển khai càng sớm càng tốt.
3. Nên test nhiều service trước khi đưa ra quyết định
4. Nên lựa chọn service phù hợp với nhiều team để tối ưu chi phí.

---
## **Tiêu chí để lựa chọn Service CI/CD tốt nhất?**

1. Phải đáp ứng được nhu cầu mà mình cần.
2. Nếu nhân sự không quá chuyên về CI/CD thì tool service cung cấp yêu cầu phải dễ sử dụng.
3. **Có nhiều lựa chọn cấu hình vì** cấu hình liên quan đến build time - yếu tố quan trong trong quy trình, build pineline phải càn nhanh càng tốt.
4. **Lựa chọn service càng phổ biến** càng tốt vì nhiều người biết cánh sử dụng. Chẳng hạn: ==Circle CI, Bitrise, Gitlab, TeamCity, Github Actions, TravisCI,...==

| Công cụ CI/CD      | Ưu điểm                               | Nhược điểm                         |
| ------------------ | ------------------------------------- | ---------------------------------- |
| **GitHub Actions** | Tích hợp sẵn trong GitHub, dễ sử dụng | Giới hạn miễn phí cho private repo |
| **GitLab CI/CD**   | Mạnh mẽ, hỗ trợ nhiều ngôn ngữ        | Cấu hình phức tạp hơn              |
| **Jenkins**        | Open-source, tùy chỉnh cao            | Cần nhiều cấu hình ban đầu         |

5. **Chi phí phù hợp với ngân sách của tổ chức**: (Ở đây Amanotes, chi phí sử dụng service CI/CD quan trong nhưng không phải yếu tố tiên quyết).

---
## **Kết luận**

Việc thành thạo CI/CD không chỉ giúp các nhóm DevOps tối ưu hóa quy trình phát triển, triển khai phần mềm, mà còn tạo ra sự linh hoạt trong việc đưa sản phẩm ra thị trường nhanh chóng và an toàn.

---
## **Tham khảo:**

1. [CI/CD là gì? Lợi ích của việc thành thạo CI/CD trong DevOps](https://200lab.io/blog/ci-cd-la-gi?srsltid=AfmBOorQLBn8E6zhJH19zpa6-Q4gx8UOEmgVJ1WMC34FCXWUENfGYXPf)
2. [CI/CD là gì? Vai trò và các nguyên tắc triển khai CI/CD](https://viettelidc.com.vn/tin-tuc/ci-cd-la-gi)
3. [CI/CD là gì? Lợi ích và các nguyên tắc triển khai CI/CD vào quy trình phát triển phần mềm](https://itviec.com/blog/ci-cd-la-gi/)