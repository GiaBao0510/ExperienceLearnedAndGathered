# Git Workflow: Quy trình quản lý source code trong làm việc nhóm

## Mục lục

1. [Git Workflow là gì và vì sao quan trọng?](#git-workflow-là-gì-và-vì-sao-quan-trọng)
2. [Ba mô hình workflow phổ biến](#ba-mô-hình-workflow-phổ-biến)
3. [Quy trình làm việc cơ bản với Git](#quy-trình-làm-việc-cơ-bản-với-git)
4. [Thực hành nâng cao giúp workflow vận hành tốt](#thực-hành-nâng-cao-giúp-workflow-vận-hành-tốt)
5. [Merge Request / Pull Request và Code Review](#merge-request--pull-request-và-code-review)
6. [Hạn chế conflict code khi làm việc nhóm](#hạn-chế-conflict-code-khi-làm-việc-nhóm)
7. [Một workflow thực tế đang được áp dụng hiệu quả](#một-workflow-thực-tế-đang-được-áp-dụng-hiệu-quả)
8. [Công cụ hỗ trợ workflow](#công-cụ-hỗ-trợ-workflow)
9. [Lỗi thường gặp và cách khắc phục](#lỗi-thường-gặp-và-cách-khắc-phục)
10. [Kết luận](#kết-luận)
11. [Mở rộng](#mở-rộng)

---

## Git Workflow là gì và vì sao quan trọng?

**Git Flow** (theo nghĩa tổng quát, không chỉ riêng mô hình cùng tên) là thuật ngữ chỉ vòng đời làm việc với Git trong một dự án - cách tạo nhánh (branch), phát triển tính năng và merge nhánh khi hoàn thành. Khái niệm mô hình **Gitflow** cụ thể được Vincent Driessen giới thiệu năm 2010 nhằm chuẩn hóa quy trình làm việc với Git.

Một workflow tốt giúp nhiều nhóm làm việc song song mà không ảnh hưởng lẫn nhau, đồng thời tách biệt rõ các môi trường development, staging và production - giúp quá trình kiểm thử (QA), phản hồi và xử lý lỗi trở nên gọn gàng, nhất quán hơn.

Điều quan trọng cần hiểu: **Git chỉ là công cụ, workflow mới là quy trình làm việc xung quanh công cụ đó**. Một team có Git nhưng không thống nhất workflow vẫn có thể rơi vào tình trạng hỗn loạn: conflict liên tục, không ai chắc merge thế nào là đúng, lịch sử commit không phản ánh đúng quá trình phát triển.

Trên thực tế, nhiều nghiên cứu về DevOps (như báo cáo *State of DevOps* của DORA) đều cho thấy các đội ngũ có quy trình phát triển phần mềm rõ ràng và nhất quán thường có tần suất triển khai cao hơn và tỷ lệ lỗi khi triển khai thấp hơn đáng kể so với các đội ngũ làm việc tự do, thiếu quy chuẩn. Đây là lý do vì sao việc thống nhất Git workflow ngay từ đầu dự án, dù nhỏ, vẫn rất đáng đầu tư thời gian.

## Ba mô hình workflow phổ biến

### 1. Git Flow

Được Vincent Driessen giới thiệu năm 2010, đây là mô hình có cấu trúc chặt chẽ nhất trong ba mô hình, sử dụng nhiều loại nhánh: `main`/`master`, `develop`, `release`, `hotfix`, `feature`. Mô hình này phù hợp với các dự án phát hành theo chu kỳ cố định (ứng dụng mobile, phần mềm đóng gói theo version), nhưng có thể trở nên nặng nề với các sản phẩm cần cập nhật liên tục hàng ngày.

![Sơ đồ mô hình Git Flow](https://images.viblo.asia/84f47fd1-a009-4beb-8957-26395fe1023d.png)

Các nhánh chính trong Git Flow:

- **master/main**: nhánh tồn tại xuyên suốt vòng đời phần mềm, được tạo mặc định khi khởi tạo repository. Đây là nhánh phản ánh mã nguồn đang chạy trên production.

  ![Nhánh master](https://images.viblo.asia/f71e46bd-452f-48c1-b451-2f8a25fff458.png)

- **develop**: nhánh phát triển chính, tồn tại song song với master, là nơi tích hợp các tính năng đã hoàn thành trước khi chuẩn bị release.

  ![Nhánh develop](https://images.viblo.asia/6e91e85b-3152-4a04-a04d-160aa0bd5135.png)

- **feature**: tách từ develop, dùng để xây dựng từng tính năng riêng biệt mà không ảnh hưởng đến các tính năng khác đang phát triển song song.

  ![Nhánh feature](https://images.viblo.asia/e4f9e958-2d5e-4f9e-98bc-9c77453b5983.png)

- **release**: tách từ develop khi chuẩn bị phát hành, dùng để kiểm tra và sửa lỗi trước khi đưa ra sản phẩm chính thức.

  ![Nhánh release](https://images.viblo.asia/7b05bf3e-e652-4ef5-817d-bef89314ef7c.png)

- **hotfix**: tách từ master, dùng để sửa gấp các lỗi nghiêm trọng phát sinh trên production mà không cần chờ chu kỳ release tiếp theo.

  ![Nhánh hotfix](https://images.viblo.asia/9fab3a45-1282-4b45-9db9-f80dc7143ae5.png)

Trong thực tế triển khai tại nhiều đội ngũ tại Việt Nam, Git Flow đầy đủ thường chỉ được duy trì nghiêm ngặt trong giai đoạn đầu dự án; với các sản phẩm cập nhật liên tục, nhiều team có xu hướng đơn giản hóa dần sang GitHub Flow hoặc Trunk-Based Development bên dưới.

### 2. GitHub Flow

Đơn giản và thực dụng hơn nhiều so với Git Flow: chỉ có `main` và các `feature branch`. Quy trình gồm: tạo feature branch, mở pull request, review, merge, deploy. Mô hình này đặt niềm tin lớn vào CI/CD - để hoạt động tốt, đội ngũ cần có bộ test đáng tin cậy và quy trình deploy tự động.

### 3. Trunk-Based Development

Mô hình hiện đại được các công ty công nghệ lớn (Google, Amazon, Meta) áp dụng rộng rãi. Mọi thành viên push trực tiếp vào `main` (hoặc `trunk`) thường xuyên, những tính năng chưa sẵn sàng được ẩn đi bằng feature flag thay vì giữ trên nhánh riêng dài ngày. Mô hình này đòi hỏi bộ test suite chạy nhanh (thường dưới 10-15 phút) và quy trình code review chặt chẽ; nếu test suite của đội ngũ mất 30-45 phút để chạy, mô hình này sẽ khó khả thi.

## Quy trình làm việc cơ bản với Git

Dù áp dụng mô hình nào, quy trình cơ bản khi làm việc với một tính năng thường theo các bước sau:

1. **Clone hoặc pull** repository để lấy mã nguồn mới nhất.
2. **Tạo nhánh mới** từ nhánh phù hợp (`develop` hoặc `main` tùy mô hình) với tên rõ ràng, ví dụ `feature/user-authentication`.
3. **Phát triển và commit** các thay đổi theo từng phần nhỏ, có ý nghĩa rõ ràng.
4. **Kiểm tra lại thay đổi** (`git diff`, `git status`) trước khi commit để tránh commit nhầm file không liên quan.
5. **Push** nhánh lên remote repository và **tạo merge request/pull request** thay vì merge trực tiếp.
6. **Review, chỉnh sửa theo góp ý**, sau đó merge vào nhánh chính khi đã được phê duyệt và pass CI/CD.

## Thực hành nâng cao giúp workflow vận hành tốt

Đây là những kinh nghiệm thực tế thường không được nêu rõ trong tài liệu Git cơ bản, nhưng ảnh hưởng lớn đến chất lượng workflow của một đội ngũ:

**Chiến lược merge (squash merge)**: phần lớn đội ngũ chuyên nghiệp dùng squash merge cho feature branch. Khi một tính năng hoàn thành với hàng chục commit nhỏ kiểu "fix typo", "thử lại", "quên chưa sửa", việc giữ nguyên toàn bộ các commit này khiến lịch sử Git trở nên khó theo dõi. Squash merge gộp toàn bộ thành một commit duy nhất, giúp nhánh chính luôn có thể revert nguyên vẹn cả tính năng nếu cần, đồng thời lịch sử commit trên `main`/`develop` phản ánh đúng từng tính năng hoặc bản sửa lỗi.

**Quy chuẩn commit message**: một commit message không rõ ràng (ví dụ "update", "fix", "asdf") sẽ gây khó khăn lớn khi cần điều tra nguyên nhân một lỗi phát sinh vài tháng sau đó. Nên áp dụng chuẩn như **Conventional Commits** (`feat:`, `fix:`, `refactor:`, `docs:`...), ví dụ: `feat(auth): thêm middleware xác thực JWT và xử lý refresh token`.

**Branch protection rules**: nếu repository không thiết lập branch protection, bất kỳ ai cũng có thể force push trực tiếp vào `main`, dẫn đến rủi ro mất code hoặc đẩy code lỗi lên production. Khi bật branch protection (yêu cầu review, yêu cầu CI/CD pass trước khi merge), đội ngũ có thêm một lớp bảo vệ quan trọng cho nhánh chính.

## Merge Request / Pull Request và Code Review

**Vì sao nên luôn tạo merge request (MR) thay vì merge trực tiếp:**

- Cho phép trưởng nhóm hoặc thành viên khác review mã nguồn trước khi merge, đảm bảo tính toàn vẹn của mã nguồn - đặc biệt quan trọng khi làm việc nhóm đông người.
- Người review có thể comment trực tiếp lên từng dòng thay đổi, giảm thời gian trao đổi qua lại và tăng hiệu quả làm việc nhóm.
- Lưu lại lịch sử thay đổi có ngữ cảnh rõ ràng: khi phát sinh lỗi hoặc cần đánh giá chất lượng, có thể xem lại toàn bộ thay đổi theo từng merge request thay vì phải lần theo từng commit riêng lẻ.

Thông thường, mọi thay đổi vào các nhánh `develop`, `release`, `main` nên đi qua merge request (ngoại trừ mã nguồn tại thời điểm khởi tạo dự án). Merge request có thể được tạo trực tiếp trên GitHub hoặc GitLab.

**Code review không chỉ để bắt lỗi**: một điều ít được nhắc đến là code review còn giúp lan tỏa kiến thức trong team. Khi reviewer phát hiện một cách tiếp cận chưa tối ưu trong pull request, đây là cơ hội để hướng dẫn lại cho tác giả. Một team thường xuyên đọc code của nhau sẽ cải thiện kỹ năng nhanh hơn nhiều so với một team chỉ push code mà không ai review - vốn dễ dẫn đến mã nguồn rối rắm, khó bảo trì về lâu dài.

Trên GitHub hay GitLab, pull request cũng dần trở thành nơi diễn ra phần lớn các quyết định về kiến trúc: thảo luận thiết kế, nhận góp ý từ reviewer, kết hợp phản hồi tự động từ CI/CD - tất cả hội tụ tại một nơi duy nhất.

## Hạn chế conflict code khi làm việc nhóm

Conflict code là điều khó tránh khỏi khi làm việc nhóm, đặc biệt với các dự án có nhiều thành viên cùng phát triển song song. Một số nguyên tắc giúp giảm thiểu:

- Chia nhỏ mã nguồn thành các module độc lập, hạn chế dồn quá nhiều logic vào một file.
- Thường xuyên đồng bộ nhánh của mình với nhánh gốc (`git pull` hoặc `git rebase`) để đảm bảo đang làm việc trên phiên bản mới nhất.
- Merge (hoặc rebase) nhánh gốc vào nhánh của mình định kỳ trong quá trình phát triển, thay vì chỉ làm việc này một lần duy nhất trước khi tạo merge request.
- Nếu phát sinh conflict, nên xử lý và merge sạch trước khi tạo merge request, tránh để reviewer phải xử lý conflict thay mình.

## Một workflow thực tế đang được áp dụng hiệu quả

Nhiều đội ngũ sản phẩm quy mô vừa (khoảng 20-50 thành viên) áp dụng một mô hình lai (hybrid), kết hợp ưu điểm của cả ba mô hình trên:

- **Nhánh main**: mã nguồn production, được bảo vệ (protected), yêu cầu tối thiểu 2 approval và pass toàn bộ test trước khi merge.
- **Nhánh staging**: gắn với môi trường staging riêng, dùng để kiểm thử trước khi đưa lên main.
- **Feature branch**: tách từ develop (hoặc từ main nếu áp dụng GitHub Flow), đặt tên theo quy ước rõ ràng, ví dụ: `feature/user-auth`, `bugfix/checkout-timeout`.
- **Hotfix branch**: tách từ main để xử lý gấp sự cố production, sau đó merge ngược lại develop để đồng bộ.
- **Quy chuẩn commit message**: theo Conventional Commits (`feat:`, `fix:`, `refactor:`...) hoặc Angular Commit Convention.

Mô hình này không quá phức tạp nhưng đủ chặt chẽ để một đội ngũ 20-50 người vận hành ổn định trong thời gian dài.

## Công cụ hỗ trợ workflow

Một số công cụ nên được đưa vào quy trình để tăng tính nhất quán và giảm lỗi con người:

- **Husky + lint-staged**: tự động chạy linting trước khi commit, tránh việc CI/CD từ chối merge request chỉ vì lỗi định dạng code.
- **commitlint**: kiểm tra định dạng commit message theo quy chuẩn đã thống nhất, giúp lịch sử commit dễ đọc và tra cứu về sau.
- **GitKraken hoặc Fork**: công cụ Git client trực quan, hỗ trợ xem lịch sử nhánh dạng đồ họa (tương đương `git log --graph`), giúp các lập trình viên mới làm quen với Git tự tin hơn khi thao tác.
- **Branch protection trên GitHub/GitLab**: thiết lập các rule bắt buộc (yêu cầu review, yêu cầu pass CI/CD) trước khi cho phép merge vào nhánh quan trọng.

## Lỗi thường gặp và cách khắc phục

| Lỗi thường gặp | Nguyên nhân | Cách khắc phục |
|---|---|---|
| **Merge hell** | Feature branch tồn tại quá lâu, phân kỳ (diverge) nhiều so với nhánh chính | Rebase feature branch với nhánh chính thường xuyên (hàng tuần), hoặc chuyển sang trunk-based development nếu team đủ điều kiện |
| **Mất commit (lost commits)** | Force push đè lên lịch sử chung, thường xảy ra khi ai đó cố "sửa" conflict bằng `git push -f` trên nhánh chung | Không bao giờ force push lên nhánh dùng chung (main/develop); chỉ force push trên nhánh cá nhân khi cần dọn dẹp lịch sử của riêng mình |
| **Dùng Git để đổ lỗi (blame culture)** | Lạm dụng lệnh `git blame` để quy trách nhiệm cá nhân thay vì cải thiện quy trình | Xem một commit lỗi là vấn đề của quy trình review chưa chặt chẽ, không phải lỗi cá nhân của một thành viên |

## Kết luận

Git là công cụ, nhưng workflow - quy trình làm việc xung quanh Git - mới là yếu tố quyết định một đội ngũ có thể ship code ổn định hay liên tục gặp sự cố. Việc chọn đúng mô hình (Git Flow, GitHub Flow, hay Trunk-Based Development) cần dựa trên đặc điểm sản phẩm và năng lực CI/CD của đội ngũ, đồng thời kết hợp các thực hành tốt như squash merge, quy chuẩn commit message, branch protection và code review nghiêm túc. Một workflow tốt không cần phức tạp - điều quan trọng là toàn team hiểu rõ và tuân thủ nhất quán.

### Mở rộng

Một số hướng tìm hiểu thêm để nâng cao kiến thức về chủ đề này:

- **Git rebase vs Git merge**: hiểu rõ khác biệt giữa hai cách tích hợp thay đổi, khi nào nên dùng rebase để giữ lịch sử tuyến tính và rủi ro cần tránh khi rebase nhánh dùng chung.
- **Semantic Versioning (SemVer)** kết hợp với Conventional Commits: cách tự động sinh changelog và version number dựa trên loại commit (`feat`, `fix`, `BREAKING CHANGE`).
- **CI/CD Pipeline cơ bản**: tìm hiểu cách thiết lập pipeline (GitHub Actions, GitLab CI) để tự động chạy test, lint, build khi có pull request - nền tảng bắt buộc để áp dụng GitHub Flow hoặc Trunk-Based Development hiệu quả.
- **Git hooks nâng cao**: ngoài Husky, tìm hiểu cách viết custom Git hook (`pre-commit`, `pre-push`) để tự động hóa các bước kiểm tra trước khi code được đẩy lên remote.
- **Trunk-Based Development và Feature Flags**: nghiên cứu sâu hơn cách các công ty lớn dùng feature flag để tách biệt việc deploy code khỏi việc release tính năng cho người dùng cuối.
- **Monorepo vs Polyrepo**: tìm hiểu cách các tổ chức lớn tổ chức nhiều service (bao gồm cả các service Golang) trong một hoặc nhiều repository, và ảnh hưởng của lựa chọn này đến Git workflow.