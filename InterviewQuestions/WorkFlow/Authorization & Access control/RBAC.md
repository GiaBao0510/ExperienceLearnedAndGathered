## RBAC là gì?

RBAC là viết tắt của Role-Based Access Control, tức kiểm soát truy cập dựa trên vai trò. Đây là phương pháp quản lý bảo mật hệ thống bằng cách gán người dùng vào các vai trò cụ thể trong một tổ chức. Mỗi vai trò đi kèm một tập quyền riêng, quyết định những hành động mà người dùng thuộc vai trò đó được phép thực hiện. RBAC được sử dụng rộng rãi trong các hệ thống phần mềm và mạng máy tính.

Thay vì cấp quyền trực tiếp cho từng người dùng, quản trị viên gán quyền cho vai trò (ví dụ: quản trị viên, nhà phát triển, nhà phân tích), sau đó gán người dùng vào vai trò. Cách tiếp cận này giúp việc quản lý truy cập trong các tổ chức lớn, nhiều người dùng trở nên đơn giản hơn: khi một nhân sự đổi vị trí, chỉ cần đổi vai trò của họ thay vì cấu hình lại từng quyền.

Ví dụ về các vai trò thường gặp:

- Quản trị viên: có quyền truy cập các chức năng quản trị của hệ thống, bao gồm tạo tài khoản người dùng, quản lý quyền truy cập, cấu hình hệ thống,...
- Nhân viên: có quyền truy cập các chức năng cơ bản như xem thông tin, cập nhật dữ liệu cá nhân, gửi yêu cầu hỗ trợ,...
- Khách hàng: có quyền truy cập các chức năng giới hạn như xem sản phẩm, đặt hàng, thanh toán,...

## Tại sao RBAC quan trọng trong bảo mật

Kiểm soát truy cập là một thành phần cốt lõi của an ninh thông tin. Nếu không có kiểm soát truy cập thích hợp, nhân viên hoặc nhà thầu có thể truy cập thông tin nằm ngoài phạm vi công việc của họ, dẫn đến rò rỉ dữ liệu, mối đe dọa từ bên trong (insider threat), sai sót do cấu hình, hoặc hành vi đánh cắp dữ liệu. Ngay cả khi người dùng không có ý đồ xấu, một tài khoản có quá nhiều quyền cũng khiến hậu quả của một lỗi thao tác hay một vụ chiếm quyền tài khoản trở nên nghiêm trọng hơn.

RBAC hỗ trợ nguyên tắc quyền hạn tối thiểu (principle of least privilege): người dùng chỉ nhận được quyền truy cập cần thiết để hoàn thành công việc. Nguyên tắc này áp dụng cho mọi hệ thống thông tin, không riêng bảo mật ứng dụng web. Lưu ý rằng RBAC chỉ tạo khuôn khổ để áp dụng nguyên tắc này. Nếu các vai trò được thiết kế quá rộng thì nguyên tắc vẫn bị vi phạm.

> - Câu "Đây là một ý tưởng quan trọng trong bảo mật ứng dụng web" đã được sửa vì thu hẹp phạm vi không cần thiết: quyền hạn tối thiểu là nguyên tắc nền tảng của toàn bộ an toàn thông tin.
> - Lưu ý rằng RBAC chỉ hỗ trợ chứ không tự bảo đảm quyền hạn tối thiểu.

## Cách hoạt động của mô hình RBAC

Mô hình RBAC thường bao gồm ba thành phần:

1. Vai trò (role): Chức năng công việc hoặc trách nhiệm được định nghĩa trong một tổ chức, chẳng hạn Quản lý Nhân sự hoặc Quản trị Hệ thống. Một vai trò nhóm lại các quyền cần thiết để thực hiện nhiệm vụ của nó.
2. Quyền hạn (permission): Hành động cụ thể được phép thực hiện trên một tài nguyên, chẳng hạn xóa người dùng, sửa đổi tài liệu, cập nhật cơ sở dữ liệu.
3. Người dùng (user): Cá nhân (hoặc tài khoản dịch vụ) được chỉ định vào một hoặc nhiều vai trò.

Ba thành phần này liên kết với nhau qua hai quan hệ gán: người dùng được gán vào vai trò, và vai trò được gán các quyền. Người dùng không bao giờ được cấp quyền trực tiếp, mà nhận quyền thông qua vai trò.

Ví dụ:

- Vai trò Quản trị viên: có thể quản lý người dùng, cấu hình hệ thống và xem nhật ký.
- Vai trò Nhà phát triển: có thể đẩy mã, chạy bản dựng, nhưng không thể quản lý người dùng.

Cơ chế này giúp quyền truy cập nhất quán giữa những người cùng vị trí và giảm rủi ro so với việc cấu hình quyền riêng cho từng người dùng.

> - Bổ sung hai quan hệ gán (user-role, role-permission) vì đây là cấu trúc cốt lõi của mô hình mà người đọc cần nắm để triển khai, trong khi bản gốc chỉ liệt kê ba thành phần rời rạc.
> - Thêm "tài khoản dịch vụ" vào định nghĩa người dùng vì trong hệ thống backend, các service và pipeline CI/CD cũng cần được gán vai trò.

### Luồng kiểm tra quyền và mô hình dữ liệu điển hình

Khi triển khai trong một ứng dụng backend, RBAC thường được hiện thực như sau:

- Mô hình dữ liệu gồm bốn bảng: `users`, `roles`, `permissions`, cùng hai bảng trung gian `user_roles` và `role_permissions` biểu diễn hai quan hệ gán nhiều-nhiều.
- Luồng xử lý một yêu cầu: xác thực danh tính người dùng (authentication), xác định các vai trò của người dùng, suy ra tập quyền từ các vai trò đó, rồi kiểm tra tập quyền có chứa quyền mà thao tác yêu cầu hay không (authorization). Nếu không có, từ chối yêu cầu.
- Nguyên tắc từ chối mặc định (default deny): thao tác nào không được cấp quyền rõ ràng thì bị từ chối.
- Việc kiểm tra quyền nên đặt ở một lớp thống nhất (ví dụ middleware hoặc tầng usecase) thay vì rải rác các điều kiện `if role == "admin"` trong mã nghiệp vụ. Kiểm tra theo quyền (permission) thay vì theo tên vai trò giúp thêm hoặc đổi vai trò mà không phải sửa mã.

## Ưu - nhược điểm của phương pháp kiểm soát truy cập RBAC là gì?

RBAC đã và đang được sử dụng rộng rãi trong các hệ thống phần mềm và mạng máy tính. Phương pháp này mang lại nhiều lợi ích cho việc bảo mật hệ thống và dữ liệu, tuy nhiên cũng có một số nhược điểm cần cân nhắc trước khi triển khai. Cụ thể:

### Về ưu điểm

- Tăng cường bảo mật: RBAC giúp giảm nguy cơ truy cập trái phép bằng cách chỉ cấp cho người dùng những quyền cần thiết cho vai trò của họ, qua đó bảo vệ dữ liệu nhạy cảm khỏi những người không có thẩm quyền.
- Quản lý dễ dàng: Quản trị viên gán quyền theo vai trò thay vì theo từng cá nhân, tiết kiệm thời gian và công sức, đặc biệt khi số lượng người dùng lớn.
- Thuận tiện khi nhân sự thay đổi: Khi có người mới vào, chuyển vị trí hoặc nghỉ việc, chỉ cần gán, đổi hoặc thu hồi vai trò. Quy trình này nhanh hơn và ít sót quyền hơn so với việc cấu hình quyền thủ công.
- Hỗ trợ kiểm toán và tuân thủ: Vì quyền được gom theo vai trò, việc trả lời câu hỏi "ai có quyền gì" dễ hơn nhiều, và đây là yêu cầu phổ biến của các chuẩn tuân thủ. Để theo dõi được hành động cụ thể của người dùng, hệ thống cần ghi nhật ký kiểm toán (audit log) riêng; bản thân RBAC không tự sinh ra nhật ký này.

### Về nhược điểm

Đi kèm các ưu điểm trên, RBAC cũng có những nhược điểm sau:

- Độ linh hoạt và độ mịn hạn chế: Quyền chỉ phụ thuộc vào các vai trò được định nghĩa trước, nên RBAC khó biểu diễn các quy tắc phụ thuộc vào ngữ cảnh hoặc từng tài nguyên cụ thể (ví dụ "chỉ được sửa bài viết do chính mình tạo", "chỉ được duyệt đơn dưới một mức giá trị", "chỉ truy cập trong giờ làm việc"). Điều này khiến RBAC kém thích ứng trong môi trường động so với ABAC (Attribute-Based Access Control), nơi quyết định truy cập dựa trên thuộc tính của người dùng, tài nguyên và môi trường. So với DAC (Discretionary Access Control), nơi chủ sở hữu tài nguyên tự quyết định chia sẻ, RBAC cũng khó điều chỉnh quyền cho từng cá nhân trong cùng một vai trò.
- Bùng nổ vai trò (role explosion): Xảy ra khi tổ chức tạo ra quá nhiều vai trò rất cụ thể (thường do mã hóa cả phòng ban, khu vực, dự án vào tên vai trò, ví dụ "Quản lý kho - Hà Nội - Dự án A") thay vì các vai trò tổng quát hơn. Khi số lượng vai trò tiến gần hoặc vượt số lượng người dùng, lợi ích của việc gom quyền theo vai trò không còn, và hệ thống trở nên khó duy trì.
- Chi phí bảo trì: Vai trò và quyền phải được rà soát, cập nhật thường xuyên để tránh lạm dụng quyền và bảo đảm người dùng không giữ quyền không cần thiết.
- Quyền hạn chồng chéo: Nhiều vai trò được cấp các quyền giống nhau hoặc tương tự. Điều này gây dư thừa, làm khó việc kiểm toán và dễ gây nhầm lẫn cho quản trị viên.
- Lan tràn quyền hạn (privilege creep): Theo thời gian, tổ chức thay đổi và người dùng tích lũy thêm nhiều vai trò. Nếu các vai trò cũ không được thu hồi khi vị trí hoặc trách nhiệm thay đổi, người dùng sẽ nắm quyền rộng hơn mức cần thiết, vi phạm nguyên tắc quyền hạn tối thiểu.

## Ứng dụng của RBAC là gì?

RBAC được ứng dụng rộng rãi trong nhiều lĩnh vực, có thể kể đến:

- Quản lý quyền truy cập cho người dùng trong ứng dụng web, ứng dụng di động và phần mềm máy tính để bàn.
- Quản lý quyền truy cập vào tài nguyên mạng như bộ định tuyến, máy chủ và thiết bị đầu cuối.
- Quản lý quyền truy cập cho khách hàng, quản trị viên và nhân viên bán hàng trong hệ thống thương mại điện tử.
- Quản lý quyền truy cập cho người dùng trong hệ thống quản lý nội dung (CMS) như WordPress, với các vai trò mặc định như Administrator, Editor, Author, Contributor, Subscriber.
- Quản lý quyền truy cập trong hạ tầng và nền tảng dành cho lập trình viên, ví dụ Kubernetes RBAC (Role, ClusterRole, RoleBinding), Azure RBAC, và hệ thống vai trò của cơ sở dữ liệu như PostgreSQL.

---

## Đề xuất mở rộng

Các mục dưới đây không được áp dụng vào nội dung chính vì chúng là lựa chọn thiết kế phụ thuộc vào quy mô và bối cảnh dự án. Mỗi mục nêu đánh đổi để bạn tự quyết định.

### 1. Chọn mô hình kiểm soát truy cập khi RBAC không đủ

Khi gặp các hạn chế như bùng nổ vai trò hoặc quy tắc phụ thuộc ngữ cảnh, có thể kết hợp hoặc thay thế RBAC bằng mô hình khác.

| Phương án | Ưu điểm | Nhược điểm |
| :--- | :--- | :--- |
| RBAC thuần | Dễ hiểu, dễ kiểm toán, đủ cho đa số hệ thống có vai trò rõ ràng | Khó biểu diễn quy tắc theo ngữ cảnh hoặc từng tài nguyên; dễ bùng nổ vai trò khi yêu cầu phức tạp |
| RBAC kết hợp thuộc tính (mô hình lai) | Giữ sự đơn giản của vai trò, thêm điều kiện theo thuộc tính (phòng ban, chủ sở hữu, giá trị giao dịch) để tránh tạo thêm vai trò | Logic phân quyền phức tạp hơn, khó kiểm toán hơn RBAC thuần |
| ABAC đầy đủ | Rất linh hoạt, quyết định theo thuộc tính người dùng, tài nguyên và môi trường | Tốn công thiết kế và quản lý chính sách; khó trả lời nhanh câu hỏi "ai có quyền gì" |
| Mô hình dựa trên quan hệ (ReBAC, ví dụ theo hướng Google Zanzibar như OpenFGA, SpiceDB) | Phù hợp khi quyền phụ thuộc quan hệ giữa các đối tượng (chia sẻ tài liệu, thư mục lồng nhau, nhóm) | Thêm một thành phần hạ tầng riêng cần vận hành; đường cong học tập cao hơn |

### 2. Cách thực thi chính sách phân quyền trong mã

| Phương án | Ưu điểm | Nhược điểm |
| :--- | :--- | :--- |
| Tự xây dựng (bảng dữ liệu và middleware) | Ít phụ thuộc, dễ hiểu, phù hợp dự án nhỏ hoặc mục đích học tập | Dễ thiếu sót khi mở rộng (phân cấp vai trò, ngoại lệ), phải tự bảo trì |
| Thư viện chính sách nhúng (ví dụ Casbin, có bản cho Go) | Hỗ trợ nhiều mô hình (ACL, RBAC, ABAC) qua tệp cấu hình mô hình, giảm lượng mã phải tự viết | Phải học cú pháp mô hình và chính sách; quyết định phân quyền trở nên khó theo dõi hơn nếu cấu hình phức tạp |
| Engine chính sách riêng (ví dụ Open Policy Agent, ngôn ngữ Rego) | Tách hẳn chính sách khỏi mã ứng dụng, dùng chung cho nhiều dịch vụ | Thêm thành phần vận hành và độ trễ cuộc gọi; cần học ngôn ngữ chính sách mới |

### 3. Nơi lưu thông tin vai trò khi dùng token xác thực (ví dụ JWT)

| Phương án | Ưu điểm | Nhược điểm |
| :--- | :--- | :--- |
| Đưa vai trò vào claim của token | Không cần tra cứu cơ sở dữ liệu ở mỗi yêu cầu, nhanh và dễ mở rộng | Thay đổi hoặc thu hồi vai trò không có hiệu lực ngay, mà chỉ có hiệu lực khi token hết hạn; token lớn hơn khi có nhiều vai trò |
| Tra cứu vai trò ở mỗi yêu cầu (có thể kèm cache ngắn hạn) | Thay đổi vai trò có hiệu lực ngay hoặc gần như ngay | Thêm độ trễ và tải cho cơ sở dữ liệu hoặc cache |
| Kết hợp: token ngắn hạn kèm vai trò, cộng với cơ chế thu hồi cho các thao tác nhạy cảm | Cân bằng giữa hiệu năng và tính kịp thời | Thiết kế và vận hành phức tạp hơn |
