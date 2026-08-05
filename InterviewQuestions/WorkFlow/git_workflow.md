## **Git Workflow: Quy trình quản lý source code**

Vòng đời của một Git được gọi với từ ngữ chuyên ngành là một Git Flow. Git Flow được Vincent Driessen đưa ra nhằm cải thiện quá trình làm việc cùng Git. Thực chất, đấy là cách chia nhánh và merge nhánh vào khi hoàn thành một tập hợp tính năng hoặc fix.

Git Flow đưa ra các quy ước để triển khai công việc. Nó được tổng kết qua quá trình làm việc thực tiễn của nhiều team trên thế giới. Mục đích là các nhóm công việc triển khai song song nhưng không ảnh hưởng tới nhau. Các  môi trường development, staging và production tách biệt giúp quá trình kiểm thử (QA), feedback và xử lý các issue được gọn gàng và thống nhất hơn.

Git hiện nay là chuẩn mực trong ngành, nhưng điều mà ít người nói là: Git chỉ là công cụ. Phần quan trọng là workflow - quy trình làm việc xung quanh Git. Đó là sự khác biệt giữa một team ship code chắc chắn với team rách code suốt ngày.

### **Tại sao workflow lại quan trọng đến vậy?**
Một nghiên cứu của GitHub năm 2023 cho thấy các team có quy trình Git clear và nhất quán có tỷ lệ lỗi production thấp hơn 70% so với team "tự do hóa". Nhưng cái mà thật sự chí mạng là: khi cái git conflict xảy ra và không ai biết merge như thế nào là đúng, đó là lúc history nằm tại chỗ rồi từ đó phát sinh bug như nấm mọc sau mưa.

Mình đã thấy một số team quy mô 50-100 dev ở Việt Nam vẫn đang git push -f để "fix" conflict. Không hiểu là bao nhiêu dòng code đã mất vào chỗ đen tối đó.

---
### **Ba mô hình workflow phổ biến**
- **Git Flow** - cái tên nghe quen quá. Giới thiệu bởi Vincent Driessen vào 2010, nó khá complex: `main`, `develop`, `release`, `hotfix`, `feature branches`. Perfect cho những dự án release theo chu kỳ (mobile app, version-based software). Nhưng jira nếu bạn đang làm product được update every day thì nó khá... nặng nề. Mình thấy nhiều team Việt áp dụng Git Flow mà rồi cứ sau 2 tháng lại back to "dev push directly to main" vì mệt.

![](https://images.viblo.asia/84f47fd1-a009-4beb-8957-26395fe1023d.png)

Các branch chính:

- **master**: là branch tồn tại xuyên suốt quá vòng đời của phần mềm được tạo mặc định trong Git khi ta tạo repository.

![](https://images.viblo.asia/f71e46bd-452f-48c1-b451-2f8a25fff458.png)
    
**develop**: là nơi các develop phát triển chính branch luôn tồn tại song song với master.
![](https://images.viblo.asia/6e91e85b-3152-4a04-a04d-160aa0bd5135.png)

**feature**: là nhánh được tách từ develop nhằm mục đích xây dựng các tính năng riêng mà không phụ thuộc vào nhau.
![](https://images.viblo.asia/e4f9e958-2d5e-4f9e-98bc-9c77453b5983.png)

**release**: là nhánh tách từ develop để kiểm tra và fix bug chuẩn bị cho việc ra mắt sản phẩm.
![](https://images.viblo.asia/7b05bf3e-e652-4ef5-817d-bef89314ef7c.png)

**hotfix**: là nhánh tách từ master để fix gấp những bug còn tồn đọng mà trên release chưa xử lý hết
![](https://images.viblo.asia/9fab3a45-1282-4b45-9db9-f80dc7143ae5.png)


- **GitHub Flow** - simple và practical hơn nhiều. Chỉ có `main` và `feature branches`. Push feature branch, open pull request, review, merge, deploy. Điều hay là nó lạc quan về CI/CD - bạn phải có test solid và automated deployment để cái này hoạt động.

- **Trunk-Based Development** - cực kỳ modern, được Google, Amazon, Facebook sử dụng. Mọi người push to main liên tục, feature flags xử lý những gì chưa ready. Nếu bạn có test suite chạy trong 10 phút và code review process chặt chẽ, đây là con đường ngắn nhất. Nhưng nếu test suite của bạn chạy mất 45 phút thì quên luôn.

---
### **Một vòng đời của Git có gì đặc biệt**
Có lẽ ở phần đầu tiên, chúng ta cũng đã nắm được ý tưởng và cách sẽ thực hiện một Git Flow rồi. Dưới đây sẽ là các step chi tiết để đạt được điều đó.
- Đầu tiên, bạn sẽ clone một repository Git để làm việc
- Sau đó, bạn chỉnh sửa bản sau của mình, phát triển và thay đổi đó
- Kiểm tra lại các thay đổi của mình trước khi commit
- Commit và push các thay đổi của mình tới remote repository
- Sau khi commit và sửa lại các lỗi sai, bạn có thể commit lại và merge chúng vào branch master của mình.

---
### ***Những insight mà dokumentation chưa nói***
**Đầu tiên, merge strategy thay đổi game**. Đa số team dùng squash merge cho feature branches. Tại sao? Vì khi bạn hoàn thành một feature với 15 commits "fix typo", "try again", "omg forgot this", thì giữ lại hết 15 commits đó chỉ làm lịch sử git trở thành một đống rác. Một lần merge squash = một commit clean duy nhất. Chủ yếu là main branch lúc nào cũng có thể revert toàn bộ feature.

**Thứ hai, commit message management**. Một commit message tệ có thể hủy hoại cả workflow. Mình từng thấy team viết commit message kiểu "asdf", "update", "fix". Khi mấy tháng sau cần investigate bug, tìm cái commit nào gây ra nó mà chỉ có 20 commits vô danh, đó là cơn ác mộng. Một commit message tốt: "Add user authentication middleware - handle JWT validation and token refresh logic".

**Thứ ba, branch protection rules là unsung hero**. Nếu repo của bạn không có branch protection, bất kỳ ai cũng có thể force push lên main. Một dev mệt và quên không squash code, một dev khác vô tình push code fail, main branch bị phá nát. Ngược lại, với branch protection (require approvals, pass CI/CD) bạn có lớp bảo vệ solid.

---
## **Một workflow thực tế đang chạy tốt**
Ở các công ty Việt Nam làm product solid mà tôi biết, họ thường dùng hybrid approach:

- **Main branch**: production code, protected, yêu cầu 2 approvals + pass tests
- **Staging branch**: kèm theo staging environment, review ở đây trước pull lên main
- **Feature branches**: từ develop (hoặc có thể từ main nếu dùng GitHub Flow), naming convention rõ ràng: feature/user-auth, bugfix/checkout-timeout
- **Hotfix branches**: từ main, urgent fix được merge ngược lại develop để sync
- **Commit message format**: Angular convention hoặc Conventional Commits (feat:, fix:, refactor:)
Cái này không quá complex nhưng đủ structured để team 20-50 người hoạt động ổn định.

---
## **Code review trong ecosystem Git**
**Một điều ít ai nhắc:** code review không chỉ là để catch bug, mà còn **spread knowledge**trong team. Khi reviewer thấy một approach tệ trong PR, họ có cơ hội teach người khác. Một team đọc code của nhau regularly, kỹ năng sẽ nâng lên exponential. Ngược lại, team chỉ push code ko ai review, 6 tháng sau sẽ thành một đống spaghetti code.

Pull requests đã trở thành nơi mà most architectural decisions được made. Trên GitHub hay GitLab, bạn có thể discuss design, có reviewers comment, có CI/CD feedback - nó là trung tâm của development process.

---
### **Tools tuy rỉ nhưng hiệu quả**
khuyến khích mọi team dùng:

- **Husky + lint-staged:** enforce linting rules trước khi commit. Khỏi bị CI/CD reject vì code style.
- **commitlint:** validate commit message format. Một small thing nhưng rất helpful khi đọc history.
- **GitKraken hoặc Fork:** visual Git client. Có thể git log --graph được, nhưng UI trực quan giúp junior dev vận hành Git tự tin hơn.
GitHub / GitLab branch protection: rule engine, enforce quality gates trước khi merge.

---
### **Lưu ý khi sử dụng Git**
Sử dung merge request: Nhiều người không có thói quen tạo merge request mà merge trực tiếp code vào branchs develop rồi push lên, điều này là không tốt. Thứ nhất, tạo merge request để teamlead hoặc review có thể review mã nguồi trước khi merge để đảm bảo tính toàn vẹn của mã nguồn, điều này là cực kì quan trọng khi phát triển phần mềm với một team nhiều người. Thứ hai: người review sẽ comment trực tiếp cần thay đổi lên merge request để giảm thời gian trao đổi tăng tính hiệu quả khi làm việc nhóm. Thứ ba, tạo merge request để lưu lại lịch sử thay đổi của mã nguồi.Khi có vấn đề về lỗi, chất lượng phần mềm.... chúng ta có thể xem lại tất cả những sự thay đổi trên từ dòng code ( việc này có thể kiếm tra bằng cách kiểm tra trên từng commit nhưng commit thì rất nhiều). Ngoài ra, đây còn là nơi để lưu lại các comment của người review, các lỗi thông thường để các member sao không còn mắc lại lỗi cũ và là nơi để học hỏi code lẫn nhau thông qua việc xem lại sự thay đổi từng dòng code của member khác.

Thông thường thì tất cả các thay đổi về mã nguồi của branchs develop, release, master đều thông qua merge request (trừ mã nguồn lúc khởi tạo dự án).

Để tạo một merge request (pull request) trong GitHub, GitLab.

conflict code: đây là câu chuyện muôn thủa và không thể tránh khỏi khi làm việc nhóm, nhất là với những nhóm dự án đông người. vậy làm thế nào để hạn chết việc conflict code: chia nhỏ code thành từ module độc lập và hạn chế viết quá nhiều code vào một file, thường xuyên merge code ở branchs về để đảm bảo code hiện tại là mới nhất, merge code của branchs trước và sau khi code nếu có conflict thì merge conflict trước khi tạo merge request.

---
### **Lỗi thường gặp và cách tránh**
- **Merge hell** xảy ra khi feature branches quá lâu và diverge từ main. Giải pháp: rebase feature branch lên main hàng tuần, hoặc dùng trunk-based development.

- **Lost commits** hay xảy ra vì force push. Rule đơn giản: không bao giờ force push lên branch chung, chỉ force push lên branch cá nhân khi cần clean up history.

- **Blame** được abuse khi team dùng Git để assign responsibility thay vì collaboration. Một commit tệ không phải lỗi của individual dev, mà là lỗi của team không review kỹ.

