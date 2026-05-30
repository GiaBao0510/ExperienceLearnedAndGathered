![](https://fostips.com/wp-content/uploads/2021/06/deepin-dark-1600x900.jpg)
## 3.1 Navigating the Linux Desktop

Để trở thành quản trị viên hệ thống Linux, cần phải thành thạo Linux như một hệ điều hành máy tính để bàn và có kỹ năng cơ bản về Công nghệ Thông tin và Truyền thông (CNTT). Sử dụng Linux cho các tác vụ năng suất, thay vì phụ thuộc vào hệ thống Windows hoặc Macintosh, giúp đẩy nhanh quá trình học tập bằng cách làm việc với các công cụ Linux hàng ngày. Quản trị viên hệ thống làm nhiều việc hơn là chỉ quản lý máy chủ; họ thường được yêu cầu hỗ trợ người dùng giải quyết các vấn đề cấu hình, đề xuất phần mềm mới và cập nhật tài liệu, cùng nhiều nhiệm vụ khác.

Hầu hết các bản phân phối Linux cho phép người dùng tải xuống gói cài đặt "máy tính để bàn" có thể được lưu vào USB. Đây là một trong những việc đầu tiên mà những người muốn trở thành quản trị viên hệ thống nên làm; tải xuống một bản phân phối chính và cài đặt nó vào một máy tính cũ. Quá trình này khá đơn giản và có nhiều hướng dẫn trực tuyến. Giao diện máy tính để bàn Linux sẽ quen thuộc với bất kỳ ai đã từng sử dụng PC hoặc Macintosh với các biểu tượng để chọn các chương trình khác nhau và ứng dụng "cài đặt" để cấu hình các thiết lập như tài khoản người dùng, mạng WiFi và thiết bị đầu vào. Sau khi làm quen với Giao diện Người dùng Đồ họa (GUI) của Linux, hay còn gọi là máy tính để bàn, bước tiếp theo là học cách thực hiện các tác vụ từ dòng lệnh.

### 3.1.1 Getting to the Command Line

Giao diện dòng lệnh (CLI) là một hệ thống nhập văn bản đơn giản để nhập bất cứ thứ gì, từ các lệnh một từ đến các kịch bản phức tạp. Hầu hết các hệ điều hành đều có CLI cung cấp cách trực tiếp để truy cập và điều khiển máy tính.

Trên các hệ thống khởi động với giao diện đồ họa người dùng (GUI), có hai cách phổ biến để truy cập dòng lệnh—một thiết bị đầu cuối dựa trên GUI và một thiết bị đầu cuối ảo:

- Thiết bị đầu cuối GUI là một chương trình trong môi trường GUI mô phỏng cửa sổ thiết bị đầu cuối. Có thể truy cập thiết bị đầu cuối GUI thông qua hệ thống menu. Ví dụ, trên máy CentOS, bạn có thể nhấp vào Ứng dụng trên thanh menu, sau đó chọn Công cụ hệ thống > và cuối cùng là Thiết bị đầu cuối. Nếu bạn có công cụ tìm kiếm, bạn có thể tìm kiếm "thiết bị đầu cuối", như được hiển thị ở đây.

![](https://ndg-content-dev.s3.amazonaws.com/media/images/linux-essentials-v2/LEv2_3_1.png)

- Một thiết bị đầu cuối ảo có thể chạy đồng thời với giao diện đồ họa người dùng (GUI) nhưng yêu cầu người dùng đăng nhập thông qua thiết bị đầu cuối ảo trước khi có thể thực thi các lệnh (giống như trước khi truy cập giao diện GUI).

Mỗi bản phân phối Linux dành cho máy tính để bàn có một chút khác biệt, nhưng ứng dụng **terminal** hoặc **x-term** sẽ mở cửa sổ thiết bị đầu cuối từ GUI. Mặc dù có những khác biệt nhỏ giữa các thuật ngữ phiên làm việc trong cửa sổ console và terminal, nhưng từ góc độ quản trị viên, chúng đều giống nhau và yêu cầu cùng một kiến ​​thức về các lệnh để sử dụng.

Các tác vụ dòng lệnh thông thường là khởi chạy chương trình, phân tích cú pháp tập lệnh và chỉnh sửa các tệp văn bản được sử dụng để cấu hình hệ thống hoặc ứng dụng. Hầu hết các máy chủ khởi động trực tiếp vào thiết bị đầu cuối, vì GUI có thể tốn nhiều tài nguyên và thường không cần thiết để thực hiện các thao tác dựa trên máy chủ.

---
## 3.2 Applications

Nhân hệ điều hành giống như người điều khiển không lưu tại sân bay, còn các ứng dụng là những chiếc máy bay dưới sự điều khiển của nó. Nhân hệ điều hành quyết định chương trình nào được cấp phát khối bộ nhớ nào, khởi động và tắt các ứng dụng, và xử lý việc hiển thị văn bản hoặc đồ họa trên màn hình.

Các ứng dụng gửi yêu cầu đến nhân hệ điều hành và đổi lại nhận được các tài nguyên, chẳng hạn như bộ nhớ, CPU và dung lượng ổ đĩa. Nếu hai ứng dụng yêu cầu cùng một tài nguyên, nhân hệ điều hành sẽ quyết định ứng dụng nào được cấp phát, và trong một số trường hợp, sẽ tắt ứng dụng khác để bảo vệ phần còn lại của hệ thống và ngăn ngừa sự cố.

Nhân hệ điều hành cũng trừu tượng hóa một số chi tiết phức tạp khỏi ứng dụng. Ví dụ, ứng dụng không biết liệu một khối lưu trữ trên đĩa nằm trên ổ SSD, ổ cứng cơ học hay thậm chí là một thư mục chia sẻ mạng. Các ứng dụng chỉ cần tuân theo Giao diện Lập trình Ứng dụng (API) của nhân hệ điều hành và do đó không cần phải lo lắng về các chi tiết triển khai. Mỗi ứng dụng hoạt động như thể nó có một khối bộ nhớ lớn trên hệ thống; Nhân hệ điều hành duy trì ảo giác này bằng cách ánh xạ lại các khối bộ nhớ nhỏ hơn, chia sẻ các khối bộ nhớ với các ứng dụng khác, hoặc thậm chí hoán đổi các khối chưa được sử dụng ra đĩa.

Nhân hệ điều hành cũng xử lý việc chuyển đổi giữa các ứng dụng, một quá trình được gọi là đa nhiệm. Một hệ thống máy tính có một số lượng nhỏ các bộ xử lý trung tâm (CPU) và một lượng bộ nhớ hữu hạn. Nhân hệ điều hành đảm nhiệm việc giải phóng một tác vụ và tải một tác vụ mới nếu nhu cầu lớn hơn tài nguyên có sẵn. Khi một tác vụ đã chạy trong một khoảng thời gian nhất định, CPU sẽ tạm dừng nó để một tác vụ khác có thể chạy. Nếu máy tính đang thực hiện nhiều tác vụ cùng một lúc, nhân hệ điều hành sẽ quyết định khi nào chuyển trọng tâm giữa các tác vụ. Với việc các tác vụ chuyển đổi nhanh chóng, có vẻ như máy tính đang thực hiện nhiều việc cùng một lúc.

Khi chúng ta, với tư cách là người dùng, nghĩ về các ứng dụng, chúng ta thường nghĩ đến các trình xử lý văn bản, trình duyệt web và ứng dụng email, tuy nhiên, có rất nhiều loại ứng dụng khác nhau. Nhân hệ điều hành không phân biệt giữa một ứng dụng hướng đến người dùng, một dịch vụ mạng giao tiếp với máy tính từ xa hoặc một tác vụ nội bộ. Từ đó, chúng ta có được một khái niệm trừu tượng gọi là tiến trình. Một tiến trình chỉ là một tác vụ được tải và theo dõi bởi nhân hệ điều hành. Một ứng dụng thậm chí có thể cần nhiều tiến trình để hoạt động, vì vậy nhân hệ điều hành sẽ đảm nhiệm việc chạy các tiến trình, khởi động và dừng chúng theo yêu cầu, và phân bổ tài nguyên hệ thống.









