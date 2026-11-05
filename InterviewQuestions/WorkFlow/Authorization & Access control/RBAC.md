## RBAC là gì?

Trên thực tế, RBAC là cụm từ viết tắt của Role Based Access Control, tức kiểm soát truy cập dựa trên vai trò. Đây là phương thức quản lý quyền truy cập được sử dụng rộng rãi trong các hệ thống phần mềm và mạng máy tính. 

Kiểm soát truy cập dựa trên vai trò, hay RBAC, là một phương pháp quản lý bảo mật hệ thống bằng cách gán người dùng vào các vai trò cụ thể trong một tổ chức. Mỗi vai trò đi kèm với một bộ quyền riêng, quyết định những hành động mà người dùng trong vai trò đó được phép thực hiện.

Thay vì cấp quyền cho từng người dùng, bạn có thể gán quyền dựa trên vai trò (ví dụ, quản trị viên, nhà phát triển, nhà phân tích, v.v.).

Cách tiếp cận này giúp dễ dàng hơn trong việc quản lý truy cập trong các tổ chức lớn với nhiều người dùng.

RBAC hoạt động dựa trên nguyên tắc gán quyền truy cập cho người dùng dựa trên vai trò của họ trong tổ chức. Mỗi vai trò được định nghĩa với một tập hợp các quyền riêng biệt, cho phép người dùng thực hiện các hành động cụ thể trong hệ thống. Ví dụ:

- Với vai trò quản trị viên: Người dùng sẽ có quyền truy cập vào tất cả các chức năng của hệ thống, bao gồm tạo tài khoản người dùng, quản lý quyền truy cập, cấu hình hệ thống,...
- Với vai trò nhân viên: Người dùng sẽ có quyền truy cập vào các chức năng cơ bản như xem thông tin, cập nhật dữ liệu cá nhân, gửi yêu cầu hỗ trợ,..
- Với vai trò khách hàng: Người dùng sẽ có quyền truy cập vào các chức năng giới hạn như xem sản phẩm, đặt hàng, thanh toán,...

## Tại Sao RBAC Quan Trọng Trong Bảo Mật
Kiểm soát truy cập là một phần quan trọng của an ninh mạng. Ví dụ, một nhà thầu từng tải xuống 6 GB dữ liệu nhạy cảm vì họ có quá nhiều quyền truy cập. Nếu không có kiểm soát truy cập thích hợp, nhân viên hoặc nhà thầu có thể truy cập thông tin mà họ không nên, điều này có thể dẫn đến rò rỉ dữ liệu, mối đe dọa từ bên trong, cấu hình sai, hoặc thậm chí là trộm cắp.

RBAC hỗ trợ nguyên tắc quyền hạn tối thiểu, nghĩa là người dùng chỉ nhận được quyền truy cập mà họ cần. Đây là một ý tưởng quan trọng trong bảo mật ứng dụng web.

##Cách Hoạt Động Của Mô Hình RBAC

Mô hình RBAC thường bao gồm 3 thành phần:
1. Vai trò: Đây là các chức năng công việc hoặc trách nhiệm được định nghĩa trong một tổ chức, chẳng hạn như Quản lý Nhân sự hoặc Quản trị Hệ thống. Một vai trò nhóm lại các quyền cụ thể cần thiết để thực hiện nhiệm vụ của nó.
2. Quyền hạn - Hành động cụ thể để thực hiện, chẳng hạn như xóa người dùng, sửa đổi tài liệu, cập nhật cơ sở dữ liệu, v.v.
3. Người dùng - Cá nhân được chỉ định vào một hoặc nhiều vai trò

Ví dụ:
- Vai trò Quản trị viên: có thể quản lý người dùng, cấu hình hệ thống, và xem nhật ký
- Vai trò Nhà phát triển: có thể đẩy mã, chạy bản dựng, nhưng không thể quản lý người dùng

Cơ chế này đảm bảo tính nhất quán và giảm rủi ro so với việc quản lý quyền truy cập của từng người dùng riêng lẻ.

## Ưu - nhược điểm của phương pháp kiểm soát truy cập RBAC là gì?
Role Based Access Control (RBAC) đã và đang được sử dụng rộng rãi trong các hệ thống phần mềm và mạng máy tính. Theo đó, phương pháp kiểm soát truy cập này mang lại nhiều lợi ích cho việc bảo mật hệ thống và dữ liệu, tuy nhiên cũng có một số nhược điểm cần được cân nhắc trước khi triển khai. Cụ thể:

### Về ưu điểm
- Tăng cường bảo mật: RBAC giúp giảm thiểu nguy cơ truy cập trái phép bằng cách chỉ cấp quyền cho người dùng thực hiện các hành động cần thiết cho vai trò của họ. Điều này giúp bảo vệ dữ liệu nhạy cảm khỏi những người dùng không được phép truy cập.
- Quản lý dễ dàng: Phương pháp kiểm soát RBAC giúp đơn giản hóa việc quản lý quyền truy cập bằng cách cho phép quản trị viên gán quyền truy cập cho người dùng dựa trên vai trò cụ thể. Qua đó, giúp tiết kiệm thời gian và công sức so với việc quản lý quyền truy cập cho từng cá nhân.
- Tăng hiệu quả: RBAC cũng giúp nhân viên tập trung vào công việc của mình thông qua việc chỉ cung cấp cho họ quyền truy cập vào các chức năng cần thiết, góp phần gián tiếp trong việc tăng năng suất và hiệu quả làm việc của đội ngũ nhân sự tại đơn vị.
- Đáp ứng yêu cầu tra soát: Nhờ việc ghi lại các quyền truy cập của người dùng và các hành động được thực hiện, phương pháp kiểm soát truy cập RBAC giúp dễ dàng kiểm tra, theo dõi hoạt động của người dùng khi cần thiết.

### Về nhược điểm
Đi kèm ưu điểm trên, kiểm soát truy cập RBAC cũng có những nhược điểm như: 
- Ít linh hoạt: RBAC có thể ít linh hoạt hơn các phương thức kiểm soát truy cập khác, chẳng hạn như DAC (Discretionary Access Control). Điều này là do RBAC khó có thể điều chỉnh quyền truy cập cho từng cá nhân trong một vai trò.
- Khó khăn trong việc phân chia đặc quyền: RBAC có thể gặp khó khăn trong việc phân chia đặc quyền một cách chi tiết cho từng trường hợp cụ thể.
- Bùng nổ vai trò xảy ra khi một tổ chức tạo ra quá nhiều vai trò rất cụ thể thay vì các danh mục rộng hơn, khiến việc duy trì trở nên khó khăn. Điều này có thể dẫn đến vấn đề khi số lượng vai trò vượt quá số lượng nhân viên khoảng 20%, vì quản lý quá nhiều vai trò trở nên không thực tế.
- Cấu trúc cứng nhắc: RBAC dựa hoàn toàn vào các vai trò được định nghĩa trước, khiến nó kém linh hoạt trong môi trường động so với ABAC, nơi quyền truy cập có thể thích ứng dựa trên thuộc tính của người dùng, tài nguyên hoặc môi trường.
- Chi phí bảo trì: Các vai trò và quyền hạn phải được xem xét và cập nhật thường xuyên để ngăn chặn lạm dụng quyền và đảm bảo người dùng không có quyền truy cập không cần thiết.
- Quyền hạn chồng chéo: Khi nhiều vai trò được cấp quyền tương tự hoặc giống nhau. Điều này sẽ làm cho việc kiểm toán trở nên khó khăn, tạo ra sự dư thừa và gây nhầm lẫn cho người quản trị.
- Lan tràn quyền hạn: Theo thời gian, có sự thay đổi tổ chức và người dùng tích lũy nhiều vai trò. Nếu vai trò được gán cho người dùng không được cập nhật hoặc thu hồi khi vị trí hoặc trách nhiệm thay đổi, nó sẽ tạo ra quyền truy cập rộng hơn mức cần thiết, vi phạm nguyên tắc quyền hạn tối thiểu.

## Ứng dụng của RBAC là gì?
RBAC đang được ứng dụng rộng rãi trong nhiều lĩnh vực khác nhau, có thể kể đến như:

- Quản lý quyền truy cập cho người dùng trong các ứng dụng web, ứng dụng di động và phần mềm máy tính để bàn.
- Quản lý quyền truy cập vào tài nguyên mạng như bộ định tuyến, máy chủ và thiết bị đầu cuối.
- Quản lý quyền truy cập cho khách hàng, quản trị viên và nhân viên bán hàng tại các hệ thống thương mại điện tử. 
- Quản lý quyền truy cập cho người dùng trong hệ thống quản lý nội dung (CMS) như WordPress,...
