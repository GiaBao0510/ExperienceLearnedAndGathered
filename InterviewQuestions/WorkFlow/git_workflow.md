## **Git Workflow: Quy trình quản lý source code**

Git hiện nay là chuẩn mực trong ngành, nhưng điều mà ít người nói là: Git chỉ là công cụ. Phần quan trọng là workflow - quy trình làm việc xung quanh Git. Đó là sự khác biệt giữa một team ship code chắc chắn với team rách code suốt ngày.

### **Tại sao workflow lại quan trọng đến vậy?**
Một nghiên cứu của GitHub năm 2023 cho thấy các team có quy trình Git clear và nhất quán có tỷ lệ lỗi production thấp hơn 70% so với team "tự do hóa". Nhưng cái mà thật sự chí mạng là: khi cái git conflict xảy ra và không ai biết merge như thế nào là đúng, đó là lúc history nằm tại chỗ rồi từ đó phát sinh bug như nấm mọc sau mưa.

Mình đã thấy một số team quy mô 50-100 dev ở Việt Nam vẫn đang git push -f để "fix" conflict. Không hiểu là bao nhiêu dòng code đã mất vào chỗ đen tối đó.

---
### **Ba mô hình workflow phổ biến**
- **Git Flow** - cái tên nghe quen quá. Giới thiệu bởi Vincent Driessen vào 2010, nó khá complex: `main`, `develop`, `release`, `hotfix`, `feature branches`. Perfect cho những dự án release theo chu kỳ (mobile app, version-based software). Nhưng jira nếu bạn đang làm product được update every day thì nó khá... nặng nề. Mình thấy nhiều team Việt áp dụng Git Flow mà rồi cứ sau 2 tháng lại back to "dev push directly to main" vì mệt.

- **GitHub Flow** - simple và practical hơn nhiều. Chỉ có `main` và `feature branches`. Push feature branch, open pull request, review, merge, deploy. Điều hay là nó lạc quan về CI/CD - bạn phải có test solid và automated deployment để cái này hoạt động.

- **Trunk-Based Development** - cực kỳ modern, được Google, Amazon, Facebook sử dụng. Mọi người push to main liên tục, feature flags xử lý những gì chưa ready. Nếu bạn có test suite chạy trong 10 phút và code review process chặt chẽ, đây là con đường ngắn nhất. Nhưng nếu test suite của bạn chạy mất 45 phút thì quên luôn.

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


