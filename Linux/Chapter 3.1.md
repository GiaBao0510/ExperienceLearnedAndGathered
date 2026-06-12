![](https://fostips.com/wp-content/uploads/2021/06/deepin-dark-1600x900.jpg)

## 3.1 Điều hướng trên Linux Desktop

Để trở thành quản trị viên hệ thống Linux, cần phải thành thạo Linux như một hệ điều hành máy tính để bàn và có kỹ năng cơ bản về Công nghệ Thông tin và Truyền thông (CNTT). Sử dụng Linux cho các tác vụ hàng ngày, thay vì phụ thuộc vào Windows hoặc macOS, giúp đẩy nhanh quá trình học tập thông qua việc tiếp xúc liên tục với các công cụ Linux. Quản trị viên hệ thống không chỉ quản lý máy chủ; họ còn thường xuyên hỗ trợ người dùng giải quyết vấn đề cấu hình, đề xuất phần mềm mới và cập nhật tài liệu kỹ thuật.

Hầu hết các bản phân phối Linux cung cấp gói cài đặt dạng "live USB" — có thể chạy trực tiếp từ USB mà không cần cài đặt, hoặc cài đặt vào máy thật. Đây là bước thực hành đầu tiên được khuyến nghị: tải về một bản phân phối phổ biến (Ubuntu, Fedora, hoặc Debian) và cài đặt vào một máy tính cũ. Giao diện desktop Linux có điểm tương đồng với Windows và macOS — biểu tượng ứng dụng, trình quản lý tệp, cài đặt WiFi và tài khoản người dùng — giúp người dùng quen thuộc làm quen nhanh chóng. Sau khi nắm được GUI, bước tiếp theo quan trọng hơn là học tương tác qua dòng lệnh (CLI).

---
### 3.1.1 Truy cập dòng lệnh

Giao diện dòng lệnh (CLI) là môi trường nhập văn bản cho phép người dùng thực thi lệnh, chạy script và điều khiển hệ thống ở mức chi tiết mà GUI không cung cấp. Trên các hệ thống khởi động với GUI, có hai cách phổ biến để truy cập dòng lệnh:

**Terminal Emulator (Terminal GUI)**

Là ứng dụng chạy bên trong môi trường desktop, mô phỏng cửa sổ terminal. Có thể mở qua menu ứng dụng hoặc công cụ tìm kiếm. Các terminal emulator phổ biến bao gồm `gnome-terminal`, `xterm`, `Konsole` (KDE), `Alacritty`.

![](https://ndg-content-dev.s3.amazonaws.com/media/images/linux-essentials-v2/LEv2_3_1.png)

**Virtual Terminal (TTY)**

Là giao diện dòng lệnh độc lập, chạy song song với GUI nhưng không phụ thuộc vào nó. Truy cập bằng tổ hợp phím `Ctrl + Alt + F2` đến `F6` (tùy bản phân phối). Khi truy cập virtual terminal lần đầu, hệ thống yêu cầu đăng nhập bằng username và password trước khi cho phép thực thi lệnh — khác với terminal emulator trong GUI vốn đã thừa hưởng phiên đăng nhập hiện tại.

Mặc dù có khác biệt nhỏ về thuật ngữ giữa "console", "terminal" và "shell", từ góc độ quản trị viên, tất cả đều phục vụ cùng mục đích và yêu cầu cùng kiến thức về lệnh.

Các tác vụ CLI thông thường bao gồm: khởi chạy chương trình, chạy script tự động hóa và chỉnh sửa file cấu hình hệ thống. Hầu hết máy chủ Linux khởi động trực tiếp vào chế độ CLI (không có GUI) vì GUI tiêu tốn tài nguyên không cần thiết trong môi trường server.

---
## 3.2 Ứng dụng

Nhân hệ điều hành (kernel) đóng vai trò điều phối trung tâm: quyết định chương trình nào được cấp phát bộ nhớ, khởi động và dừng ứng dụng, và xử lý việc hiển thị đầu ra lên màn hình.

Các ứng dụng giao tiếp với kernel thông qua **system call** — interface chuẩn hóa cho phép ứng dụng yêu cầu tài nguyên như bộ nhớ, CPU và dung lượng lưu trữ một cách an toàn. Nếu hai ứng dụng cạnh tranh cùng một tài nguyên, kernel sẽ phân xử; trong trường hợp nghiêm trọng, kernel có thể kết thúc một tiến trình để bảo vệ sự ổn định của hệ thống.

Kernel cũng cung cấp lớp trừu tượng phần cứng cho ứng dụng: một ứng dụng không cần biết dữ liệu đang nằm trên SSD, HDD hay network file share — nó chỉ cần gọi đúng system call. Kernel quản lý bộ nhớ ảo (virtual memory), cho phép mỗi tiến trình hoạt động như thể nó có không gian bộ nhớ riêng biệt lớn, trong khi thực tế kernel đang ánh xạ và chia sẻ các trang bộ nhớ vật lý, đồng thời hoán đổi (swap) các trang ít dùng ra đĩa khi cần.

Kernel xử lý đa nhiệm (multitasking) thông qua **bộ lập lịch tiến trình (process scheduler)**. Trên hệ thống có ít CPU hơn số tiến trình đang chạy, scheduler phân bổ thời gian CPU cho từng tiến trình theo chu kỳ rất ngắn (time slice), tạo ra ảo giác tất cả chương trình chạy đồng thời.

Từ góc độ kernel, mọi đơn vị công việc đều là một **tiến trình (process)** — dù là ứng dụng người dùng, network service, hay tác vụ nội bộ. Mỗi tiến trình được kernel theo dõi với một PID (Process ID) duy nhất. Một ứng dụng có thể tạo ra nhiều tiến trình con để hoạt động song song; kernel chịu trách nhiệm quản lý toàn bộ vòng đời của chúng.

---
## 3.2.1 Phân loại ứng dụng chính

Linux có thể hoạt động đồng thời ở nhiều vai trò: máy chủ cung cấp dịch vụ, máy trạm làm việc, hoặc môi trường phát triển và kiểm thử. Một lợi thế quan trọng là Linux có thể tái hiện môi trường production trên phần cứng giá rẻ — cho phép phát triển, kiểm thử và xác minh mà không cần phần cứng chuyên dụng hay giấy phép phần mềm đắt tiền.

Phần mềm Linux được phân thành ba loại:

**Ứng dụng máy chủ (Server Applications)**

Phần mềm không tương tác trực tiếp với người dùng qua màn hình và bàn phím. Mục đích chính là cung cấp dịch vụ cho các máy khác (clients) qua mạng, hoặc xử lý dữ liệu tự động trong nền.

**Ứng dụng máy tính để bàn (Desktop Applications)**

Phần mềm người dùng tương tác trực tiếp: trình duyệt web, trình soạn thảo văn bản, ứng dụng email, trình phát nhạc. Đây là phía "client" trong mô hình client/server.

**Công cụ hệ thống (Utilities/Tools)**

Phần mềm hỗ trợ quản trị và vận hành hệ thống: công cụ cấu hình hệ thống, shell (Bash, Zsh), trình biên dịch (GCC), trình gỡ lỗi (GDB), và các tiện ích dòng lệnh. Đây là nhóm công cụ mà quản trị viên sử dụng thường xuyên nhất.

Tính khả dụng của ứng dụng thay đổi theo bản phân phối vì các distro khác nhau có các phiên bản thư viện hệ thống khác nhau. Tuy nhiên, các ứng dụng phổ biến như Firefox và LibreOffice có sẵn trên hầu hết mọi bản phân phối Linux chính.

---
## 3.2.2 Ứng dụng máy chủ

Linux chiếm ưu thế trong thị trường máy chủ nhờ độ tin cậy, hiệu năng và khả năng tối ưu hóa: quản trị viên có thể cài đặt chỉ đúng những gì cần thiết, loại bỏ hoàn toàn GUI và các thành phần không dùng đến.

---
### 3.2.2.1 Web Servers

Máy chủ web lưu trữ và phục vụ nội dung trang web thông qua giao thức HTTP (cổng 80) hoặc HTTPS (cổng 443). Nội dung có thể là tĩnh (HTML, CSS, JS, hình ảnh được trả về nguyên vẹn) hoặc động (máy chủ web chuyển tiếp request đến ứng dụng backend để tạo nội dung theo yêu cầu).

**Apache HTTP Server (httpd)** là một trong những máy chủ web lâu đời và phổ biến nhất. Sau khi dự án Apache phát triển, nhóm sáng lập thành lập **Apache Software Foundation (ASF)** — tổ chức phi lợi nhuận hiện duy trì hơn 300 dự án phần mềm mã nguồn mở, bao gồm Kafka, Hadoop, Tomcat và nhiều dự án khác.

**NGINX** là máy chủ web tập trung vào hiệu năng cao, được thiết kế với kiến trúc event-driven không đồng bộ (asynchronous event-driven architecture), khác với mô hình thread-per-connection của Apache. NGINX đặc biệt hiệu quả khi phục vụ nội dung tĩnh đồng thời với lượng lớn kết nối. Hiện nay, NGINX và Apache cùng nhau chiếm phần lớn thị phần máy chủ web toàn cầu.

---
### 3.2.2.2 Private Cloud Servers

Khi nhu cầu kiểm soát dữ liệu và tuân thủ quy định tăng lên, nhiều tổ chức lựa chọn triển khai hạ tầng đám mây riêng (private cloud) thay vì phụ thuộc hoàn toàn vào cloud công cộng.

**ownCloud** được Frank Karlitschek khởi động năm 2010 như một giải pháp lưu trữ, đồng bộ và chia sẻ file tự lưu trữ (self-hosted), tương tự Dropbox nhưng hoàn toàn do tổ chức kiểm soát. Có hai phiên bản: mã nguồn mở (GNU AGPLv3) và doanh nghiệp (commercial license).

**Nextcloud** được Karlitschek tách ra (fork) từ ownCloud năm 2016, phát triển theo hướng minh bạch và cộng đồng nhiều hơn. Nextcloud hiện được triển khai rộng rãi hơn ownCloud và cung cấp thêm nhiều tính năng cộng tác: calendar, video call, document editing tích hợp.

Cả hai dự án đều phù hợp cho tổ chức có yêu cầu về bảo mật dữ liệu, quyền riêng tư và tuân thủ quy định (GDPR, HIPAA...) mà không muốn lưu dữ liệu trên hạ tầng của bên thứ ba.

---
### 3.2.2.3 Database Servers

Máy chủ cơ sở dữ liệu là thành phần nền tảng của hầu hết ứng dụng web và dịch vụ trực tuyến. Ứng dụng backend đọc và ghi dữ liệu thông qua database server; người dùng không tương tác trực tiếp với database.

**MariaDB** là một nhánh (fork) mã nguồn mở do cộng đồng phát triển của MySQL, được tạo ra sau khi Oracle mua lại Sun Microsystems (và MySQL) năm 2010. MariaDB duy trì tính tương thích cao với MySQL trong khi cung cấp các cải tiến về hiệu năng và tính năng.

Các hệ quản trị cơ sở dữ liệu quan hệ phổ biến khác trên Linux:

- **PostgreSQL:** Được biết đến với tính tuân thủ chuẩn SQL, hỗ trợ tính năng nâng cao (JSON, full-text search, extensibility), phù hợp cho ứng dụng doanh nghiệp phức tạp.
- **SQLite:** Cơ sở dữ liệu nhúng (embedded), không cần server riêng, thường dùng trong ứng dụng mobile và desktop.

Dữ liệu được truy vấn và thao tác thông qua **SQL (Structured Query Language)** — ngôn ngữ chuẩn hóa để đọc, thêm, sửa và xóa dữ liệu trong cơ sở dữ liệu quan hệ.

---
### 3.2.2.4 Email Servers

Hạ tầng email Linux bao gồm ba thành phần với trách nhiệm tách biệt:

**MTA — Mail Transfer Agent**

Chịu trách nhiệm chuyển tiếp email giữa các mail server qua giao thức SMTP. Hai MTA phổ biến nhất:

- **Sendmail:** Một trong những MTA lâu đời nhất, cấu hình phức tạp.
- **Postfix:** Được thiết kế với mục tiêu đơn giản hóa và bảo mật tốt hơn Sendmail; hiện là lựa chọn phổ biến nhất trong triển khai mới.

**MDA — Mail Delivery Agent**

Còn gọi là **Local Delivery Agent (LDA)**. Nhận email từ MTA và lưu vào hộp thư (mailbox) của người dùng trên máy chủ cục bộ. Thường được MTA gọi tự động ở bước cuối cùng trong chuỗi chuyển tiếp.

**POP/IMAP Server**

Cho phép email client (Outlook, Thunderbird...) truy xuất email từ mailbox trên server:

- **POP3 (Post Office Protocol v3):** Tải email về client và thường xóa khỏi server. Phù hợp khi dùng một thiết bị duy nhất.
- **IMAP (Internet Message Access Protocol):** Đồng bộ email giữa server và nhiều client. Email được lưu trên server, phù hợp với người dùng nhiều thiết bị.

**Dovecot** là POP/IMAP server phổ biến nhất, được biết đến nhờ tính dễ cấu hình, hiệu năng cao và bảo mật tốt. **Cyrus IMAP** là một lựa chọn khác, thường dùng trong môi trường doanh nghiệp lớn.

Điểm khác biệt của hệ sinh thái email mã nguồn mở so với Microsoft Exchange là tính mô-đun: mỗi thành phần (MTA, MDA, IMAP server, spam filter, webmail) có thể được chọn và kết hợp độc lập. Microsoft Exchange là giải pháp tích hợp khép kín từ một nhà cung cấp; trong thế giới mã nguồn mở, các thành phần riêng lẻ từ nhiều dự án khác nhau được kết hợp lại để tạo thành một hệ thống email hoàn chỉnh.

---
### 3.2.2.5 File Sharing và Network Infrastructure

**File Sharing**

**Samba** là giải pháp chia sẻ file giữa Linux và Windows, triển khai giao thức **SMB/CIFS** của Microsoft. Samba cho phép máy Linux xuất hiện như một máy Windows trên mạng, chia sẻ thư mục và máy in, thậm chí đóng vai trò Domain Controller trong môi trường Active Directory.

**Netatalk** cho phép máy Linux đóng vai trò file server cho macOS, sử dụng giao thức **AFP (Apple Filing Protocol)**. Tuy nhiên, từ macOS 10.9+, Apple ưu tiên SMB thay cho AFP, nên Netatalk ít được triển khai mới hơn.

**NFS (Network File System)** là giao thức chia sẻ file gốc của UNIX/Linux. NFS được tích hợp sẵn vào kernel, cho phép mount filesystem từ xa như một đĩa cục bộ — thao tác hoàn toàn trong suốt với các ứng dụng. NFS phổ biến trong môi trường datacenter thuần Linux.

**Network Infrastructure Services**

**DNS (Domain Name System)** chuyển đổi tên miền (ví dụ `www.example.com`) thành địa chỉ IP (ví dụ `93.184.216.34`). DNS cũng lưu trữ các bản ghi khác như MX record (xác định mail server của một domain). **BIND** (Berkeley Internet Name Domain), được duy trì bởi Internet Systems Consortium (ISC), là DNS server mã nguồn mở phổ biến nhất.

**LDAP (Lightweight Directory Access Protocol)** là giao thức truy vấn thư mục người dùng và tài nguyên mạng. LDAP lưu trữ đối tượng (người dùng, nhóm, thiết bị) trong cấu trúc cây phân cấp (DIT — Directory Information Tree). **OpenLDAP** là triển khai mã nguồn mở phổ biến nhất, thường dùng làm backend xác thực tập trung trong môi trường doanh nghiệp Linux. Microsoft Active Directory cũng hỗ trợ LDAP.

**DHCP (Dynamic Host Configuration Protocol)** tự động cấp phát địa chỉ IP và thông tin cấu hình mạng (subnet mask, default gateway, DNS server) cho các máy khi kết nối vào mạng. Máy chủ DHCP phổ biến nhất trong môi trường Linux là **ISC DHCP** và **Kea** (thay thế hiện đại hơn, cũng do ISC phát triển).

---
## 3.2.3 Desktop Applications

_(Nội dung phần này sẽ được bổ sung trong phần tiếp theo.)_

---