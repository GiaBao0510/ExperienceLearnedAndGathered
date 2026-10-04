# Bộ câu hỏi phỏng vấn Backend 
---

## Câu 1: Cơ sở dữ liệu (Database) là gì?

**Câu trả lời:** Cơ sở dữ liệu (Database) là một tập hợp dữ liệu có tổ chức, được lưu trữ và quản lý một cách có hệ thống để có thể truy cập, cập nhật và xử lý một cách hiệu quả. Việc lưu trữ, truy vấn và quản lý dữ liệu trong database thường không do con người thao tác trực tiếp lên file, mà được thực hiện thông qua một phần mềm trung gian gọi là **DBMS (Database Management System)** — ví dụ MySQL, PostgreSQL, MongoDB, SQL Server. DBMS chịu trách nhiệm đảm bảo dữ liệu được lưu đúng cấu trúc, hỗ trợ truy vấn nhanh, kiểm soát truy cập đồng thời và bảo vệ tính toàn vẹn dữ liệu.

**Ví dụ:** Một hệ thống quản lý bán hàng cần lưu thông tin khách hàng, sản phẩm, đơn hàng — tất cả được tổ chức thành các bảng (table) có quan hệ với nhau trong một database quan hệ như MySQL, thay vì lưu rời rạc trong các file Excel không có liên kết.

**Mở rộng:**
- Database được chia thành hai nhóm lớn: **Relational Database (SQL)** và **Non-relational Database (NoSQL)** — xem chi tiết ở Câu 5.

**Nguồn trích dẫn:**
- Oracle – "What is a Database?": https://www.oracle.com/database/what-is-database/

---

## Câu 2: Phân biệt giữa Session và Cookie

**Câu trả lời:**

- **Cookie:** Là một đoạn dữ liệu nhỏ (thường dưới 4KB) được lưu trực tiếp phía trình duyệt (client). Server gửi Cookie về client qua header `Set-Cookie`; sau đó trình duyệt sẽ tự động đính kèm Cookie này vào mọi request tiếp theo gửi tới cùng domain, cho đến khi Cookie hết hạn (expire) hoặc bị xóa. Cookie có thể lưu bất kỳ loại dữ liệu nhỏ nào — không nhất thiết phải liên quan đến đăng nhập.
- **Session:** Là cơ chế lưu trạng thái của người dùng ở **phía server** (trong RAM, Redis hoặc database). Khi user đăng nhập, server tạo một `Session ID` duy nhất đại diện cho phiên làm việc đó, lưu toàn bộ dữ liệu chi tiết (thông tin user, quyền hạn...) ở phía server, rồi chỉ gửi `Session ID` (thường thông qua Cookie) về cho client giữ.

**Mối quan hệ:** Cookie chỉ là một trong những "phương tiện vận chuyển" phổ biến nhất để mang `Session ID` giữa client và server, nhưng hai khái niệm không đồng nhất — Cookie hoàn toàn có thể dùng độc lập mà không cần Session (ví dụ lưu theme giao diện người dùng đã chọn), và ngược lại, Session ID cũng có thể được truyền qua cách khác như URL hoặc header tùy chỉnh.

**Ví dụ:** Khi đăng nhập vào một website, server tạo Session lưu thông tin "user_id = 5, role = admin" trong Redis, rồi gửi về trình duyệt một Cookie chứa `session_id=abc123`. Mỗi request sau đó, trình duyệt tự gửi kèm Cookie này; server tra `abc123` trong Redis để biết đây là admin, không cần client tự gửi lại thông tin nhạy cảm.

**Mở rộng:**

- Tìm hiểu thêm về cơ chế Token-based Authentication (JWT) — một cách tiếp cận khác không cần lưu Session ở phía server (Stateless).
- Tìm hiểu các thuộc tính bảo mật quan trọng của Cookie: `HttpOnly` (chặn JavaScript đọc cookie, chống XSS), `Secure` (chỉ gửi qua HTTPS), `SameSite` (hạn chế gửi cookie cross-site, chống CSRF).

**Nguồn trích dẫn:**

- MDN Web Docs – "Using HTTP cookies": https://developer.mozilla.org/en-US/docs/Web/HTTP/Cookies

---

## Câu 3: CORS (Cross-Origin Resource Sharing) là gì?

**Câu trả lời:** CORS là một cơ chế bảo mật của trình duyệt, cho phép (hoặc chặn) một trang web gọi tài nguyên từ một domain (origin) khác với domain đang hiển thị trang đó. Mặc định, trình duyệt áp dụng chính sách **Same-Origin Policy** — chặn mọi request cross-origin trừ khi server đích chủ động cho phép thông qua các header CORS. Đây là cơ chế bảo vệ người dùng, ngăn một trang web độc hại âm thầm gọi API tới một trang khác bằng chính phiên đăng nhập của nạn nhân.

**Backend cấu hình CORS bằng các header:**

- `Access-Control-Allow-Origin`: khai báo chính xác domain được phép gọi API (hạn chế dùng `*` ở môi trường Production).
- `Access-Control-Allow-Methods`: danh sách phương thức HTTP được phép (GET, POST, PUT, DELETE...).
- `Access-Control-Allow-Headers`: danh sách header client được phép gửi lên (ví dụ `Authorization`, `Content-Type`).

**Ví dụ:** Frontend chạy ở `https://app.example.com` gọi API ở `https://api.example.com`. Nếu `api.example.com` không cấu hình `Access-Control-Allow-Origin: https://app.example.com`, trình duyệt sẽ chặn response dù request đã thực sự được server xử lý xong — đây là lỗi phổ biến khiến nhiều người nhầm tưởng do server hoặc mạng.

**Mở rộng:**

- Với các request "không đơn giản" (ví dụ có header `Authorization` hoặc method `PUT`/`DELETE`), trình duyệt sẽ tự động gửi trước một **Preflight Request** (`OPTIONS`) để hỏi server xem request thật có được phép hay không.

**Nguồn trích dẫn:**

- MDN Web Docs – "Cross-Origin Resource Sharing (CORS)": https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS

---

## Câu 4: RBAC (Role-Based Access Control) là gì?

**Câu trả lời:** RBAC là mô hình phân quyền dựa trên **Vai trò (Role)**. Thay vì gán quyền hạn trực tiếp cho từng user riêng lẻ, hệ thống định nghĩa sẵn một số vai trò cố định (ví dụ: Admin, Manager, User), mỗi vai trò gắn với một tập quyền hạn cụ thể. Khi cần cấp quyền cho một user, quản trị viên chỉ cần gán user đó vào vai trò phù hợp; user sẽ tự động có toàn bộ quyền hạn của vai trò đó. Cách tiếp cận này giúp việc quản lý quyền hạn trở nên đơn giản và tập trung, đặc biệt hiệu quả khi số lượng user lớn nhưng số loại vai trò giới hạn.

**Ví dụ:** Trong một hệ thống quản lý trường học, vai trò "Teacher" có quyền nhập điểm và xem danh sách lớp mình dạy, vai trò "Student" chỉ có quyền xem điểm của chính mình, còn vai trò "Admin" có toàn quyền quản lý tài khoản. Khi có giáo viên mới, quản trị viên chỉ cần gán tài khoản đó vào vai trò "Teacher" thay vì cấu hình lại từng quyền một theo cách thủ công.

**Mở rộng:**

- Khi hệ thống cần các luật phân quyền phức tạp hơn vai trò cố định (ví dụ: "chỉ cho phép xem dữ liệu trong giờ hành chính, từ mạng nội bộ"), có thể cân nhắc mô hình **ABAC (Attribute-Based Access Control)** — phân quyền dựa trên nhiều thuộc tính (của user, tài nguyên, môi trường) thay vì chỉ dựa trên vai trò cố định. RBAC đơn giản, dễ triển khai nhưng kém linh hoạt hơn khi yêu cầu phân quyền phức tạp; ABAC linh hoạt, chi tiết hơn nhưng phức tạp hơn khi thiết kế và vận hành — tùy quy mô và độ phức tạp nghiệp vụ mà lựa chọn mô hình phù hợp, hoặc kết hợp cả hai.

**Nguồn trích dẫn:**

- NIST – "Role Based Access Control": https://csrc.nist.gov/projects/role-based-access-control

---

## Câu 5: Sự khác biệt giữa SQL và NoSQL là gì? Trường hợp sử dụng từng loại như thế nào?

**Câu trả lời:**

- **SQL (Relational Database – MySQL, PostgreSQL, SQL Server):** Lưu dữ liệu theo cấu trúc bảng (table) với schema cố định được định nghĩa trước, các bảng có quan hệ với nhau qua khóa ngoại (foreign key). Tuân thủ nghiêm ngặt tính chất ACID (xem Câu 7), đảm bảo tính toàn vẹn và nhất quán dữ liệu rất cao.
- **NoSQL (Non-relational Database – MongoDB, Redis, Cassandra):** Lưu dữ liệu linh hoạt hơn, không yêu cầu schema cố định (schema-less), theo nhiều mô hình khác nhau: Document (MongoDB), Key-Value (Redis), Column-family (Cassandra), Graph (Neo4j). Ưu tiên tốc độ đọc/ghi cao và khả năng mở rộng theo chiều ngang (horizontal scaling) dễ dàng trên nhiều máy chủ, nhưng thường chỉ đạt tính nhất quán cuối cùng (Eventual Consistency) thay vì nhất quán tức thời.

**Khi nào dùng loại nào:** Chọn **SQL** khi dữ liệu có cấu trúc rõ ràng, nhiều mối quan hệ phức tạp và hệ thống yêu cầu tính toàn vẹn tuyệt đối (hệ thống ngân hàng, quản lý đơn hàng). Chọn **NoSQL** khi dữ liệu không có cấu trúc cố định, cần tốc độ đọc/ghi cực cao và khả năng mở rộng linh hoạt (hệ thống chat, lưu log, newsfeed mạng xã hội).

**Ví dụ:** Hệ thống quản lý tài khoản ngân hàng cần SQL để đảm bảo số dư luôn chính xác sau mỗi giao dịch (không chấp nhận dữ liệu tạm thời sai lệch). Hệ thống lưu tin nhắn chat hoặc bộ đếm lượt xem video phù hợp với NoSQL vì khối lượng ghi cực lớn và không đòi hỏi quan hệ phức tạp giữa các bản ghi.

**Mở rộng:**

- Tìm hiểu thêm CAP Theorem để hiểu rõ hơn trade-off giữa Consistency và Availability mà nhiều hệ thống NoSQL phân tán phải lựa chọn.

**Nguồn trích dẫn:**

- MongoDB Docs – "SQL to MongoDB Mapping Chart": https://www.mongodb.com/docs/manual/reference/sql-comparison/

---

## Câu 6: Trong SQL có những loại Join nào?

**Câu trả lời:** Join là thao tác kết hợp dữ liệu từ hai hay nhiều bảng dựa trên một điều kiện liên kết (thường là khóa chính – khóa ngoại). Các loại Join phổ biến:

- **INNER JOIN:** Chỉ trả về các bản ghi có dữ liệu khớp ở **cả hai** bảng. Đây là loại Join phổ biến nhất.
- **LEFT JOIN (LEFT OUTER JOIN):** Trả về toàn bộ bản ghi ở bảng bên trái, kèm dữ liệu khớp từ bảng bên phải (nếu không khớp, các cột bên phải sẽ là `NULL`).
- **RIGHT JOIN (RIGHT OUTER JOIN):** Ngược lại với LEFT JOIN — trả về toàn bộ bản ghi ở bảng bên phải, kèm dữ liệu khớp từ bảng bên trái (nếu có).
- **FULL JOIN (FULL OUTER JOIN):** Trả về toàn bộ bản ghi từ cả hai bảng, bất kể có khớp hay không; những phần không khớp sẽ nhận `NULL` ở cột tương ứng.
- **CROSS JOIN:** Trả về tích Descartes — kết hợp mọi bản ghi của bảng này với mọi bản ghi của bảng kia, không cần điều kiện liên kết (cẩn trọng vì số bản ghi kết quả tăng rất nhanh: m bản ghi × n bản ghi).
- **SELF JOIN:** Một bảng tự join với chính nó, thường dùng để xử lý dữ liệu phân cấp (ví dụ bảng nhân viên có cột `manager_id` tham chiếu ngược lại chính bảng nhân viên).

**Ví dụ:** Với bảng `Customers` và `Orders`, `INNER JOIN` chỉ trả về khách hàng đã từng đặt hàng. `LEFT JOIN` trả về toàn bộ khách hàng, kể cả những người chưa từng đặt đơn nào (cột thông tin đơn hàng sẽ là `NULL`) — hữu ích khi cần báo cáo "khách hàng nào chưa phát sinh đơn hàng".

**Mở rộng:**

- MySQL không hỗ trợ trực tiếp `FULL JOIN` — thường phải mô phỏng bằng cách kết hợp `LEFT JOIN UNION RIGHT JOIN`.

**Nguồn trích dẫn:**

- PostgreSQL Official Docs – "Joins Between Tables": https://www.postgresql.org/docs/current/tutorial-join.html

---

## Câu 7: Nêu 4 tính chất ACID

**Câu trả lời:** ACID là 4 tính chất đảm bảo một giao dịch (transaction) trong database được xử lý đáng tin cậy:

- **A – Atomicity (Tính nguyên tử):** Toàn bộ các thao tác trong một transaction được xem là một khối thống nhất — hoặc tất cả cùng thành công, hoặc nếu có bất kỳ bước nào lỗi thì toàn bộ transaction phải được rollback (hoàn tác), không để lại trạng thái dữ liệu nửa vời.
- **C – Consistency (Tính nhất quán):** Transaction chỉ được phép đưa database từ một trạng thái hợp lệ sang một trạng thái hợp lệ khác, không vi phạm bất kỳ ràng buộc (constraint), quy tắc toàn vẹn nào đã định nghĩa trên database.
- **I – Isolation (Tính cô lập):** Nhiều transaction chạy đồng thời không được phép ảnh hưởng lẫn nhau theo cách gây ra kết quả sai lệch — kết quả cuối cùng phải tương đương như thể các transaction được chạy tuần tự, lần lượt từng cái một.
- **D – Durability (Tính bền vững):** Sau khi một transaction đã commit thành công, dữ liệu phải được lưu trữ vĩnh viễn, không bị mất ngay cả khi hệ thống gặp sự cố (mất điện, crash) ngay sau đó.

**Ví dụ:** Giao dịch chuyển khoản 1 triệu đồng từ tài khoản A sang tài khoản B gồm 2 bước: trừ tiền A, cộng tiền B. Nhờ **Atomicity**, nếu bước cộng tiền B thất bại (ví dụ do lỗi hệ thống), bước trừ tiền A ở trên cũng phải được rollback lại — không thể để xảy ra tình huống tiền bị trừ ở A nhưng không được cộng vào B.

**Mở rộng:**

- Tìm hiểu các mức độ cô lập (Isolation Level) trong SQL: Read Uncommitted, Read Committed, Repeatable Read, Serializable — mỗi mức đánh đổi giữa hiệu năng và mức độ an toàn khi nhiều transaction chạy song song.

**Nguồn trích dẫn:**

- Database System Concepts – Silberschatz, Korth, Sudarshan, chương "Transactions"

---

## Câu 8: Index là gì? Áp dụng vào trường hợp nào? Tại sao?

**Câu trả lời:** Index là một cấu trúc dữ liệu bổ sung, được database xây dựng trên một hoặc nhiều cột của bảng để tăng tốc độ truy vấn đọc, tương tự như mục lục của một cuốn sách giúp người đọc tìm nhanh một chương mà không cần lật qua từng trang. Thay vì quét toàn bộ bảng (Full Table Scan) với độ phức tạp O(N), database dùng cấu trúc cây cân bằng (thường là B-Tree/B+ Tree) để thu hẹp phạm vi tìm kiếm xuống còn O(log N).

**Nên áp dụng Index khi nào:** Trên các cột thường xuyên xuất hiện trong mệnh đề `WHERE`, `JOIN`, `ORDER BY`, hoặc các cột có tính chọn lọc cao (nhiều giá trị khác nhau, ví dụ `email`, `user_id`). **Tại sao không nên tạo quá nhiều Index:** Mỗi lần `INSERT`/`UPDATE`/`DELETE`, database phải cập nhật lại toàn bộ cấu trúc Index liên quan, làm chậm tốc độ ghi; đồng thời Index chiếm thêm dung lượng lưu trữ và có thể khiến trình tối ưu truy vấn (query planner) mất nhiều thời gian hơn để chọn đúng Index phù hợp cho một câu truy vấn cụ thể. Vì vậy, việc đánh Index cần cân bằng giữa lợi ích tăng tốc đọc và chi phí ghi, không nên đánh index tràn lan trên mọi cột.

**Ví dụ:** Một bảng `users` với 5 triệu bản ghi, nếu không có Index trên cột `email`, câu lệnh `SELECT * FROM users WHERE email = 'a@gmail.com'` phải quét qua toàn bộ 5 triệu dòng. Nếu đánh Index trên cột `email`, database chỉ cần duyệt qua cấu trúc B-Tree với độ sâu rất nhỏ (thường chỉ vài bước) để tìm ra đúng bản ghi, nhanh hơn rất nhiều lần.

**Mở rộng:**
- Tìm hiểu Composite Index (Index trên nhiều cột) và thứ tự khai báo cột ảnh hưởng thế nào đến khả năng Index được sử dụng.
- Dùng lệnh `EXPLAIN`/`EXPLAIN ANALYZE` để kiểm tra query planner có thực sự dùng Index như kỳ vọng hay không.

**Nguồn trích dẫn:**

- PostgreSQL Official Documentation – "Indexes": https://www.postgresql.org/docs/current/indexes.html
- Use The Index, Luke!: https://use-the-index-luke.com/

---

## Câu 9: Sự khác biệt giữa Kafka và Redis Pub/Sub là gì?

**Câu trả lời:** Cả Kafka và Redis Pub/Sub đều cho phép một Publisher phát message đến nhiều Subscriber, nhưng khác nhau căn bản về độ tin cậy và mục đích sử dụng:

- **Redis Pub/Sub:** Hoạt động theo cơ chế "fire-and-forget" (gửi rồi quên) — message **không được lưu trữ** lại. Nếu tại thời điểm publish không có Subscriber nào đang kết nối và lắng nghe, message đó sẽ **mất vĩnh viễn**, không thể phát lại (replay). Ưu điểm là cực kỳ đơn giản, độ trễ thấp, phù hợp cho các thông báo tức thời mà việc mất một vài message không ảnh hưởng nghiêm trọng.
- **Kafka:** Là một nền tảng event streaming phân tán thực sự. Message được ghi xuống đĩa theo các log có thứ tự (partition) và được giữ lại trong một khoảng thời gian cấu hình được (retention period), bất kể có Consumer đang lắng nghe hay không. Consumer có thể tham gia muộn và vẫn đọc lại toàn bộ lịch sử message từ bất kỳ vị trí (offset) nào. Kafka cũng hỗ trợ Consumer Group để chia tải xử lý và đảm bảo throughput cực cao cho dữ liệu lớn.

**Khi nào dùng loại nào:** Dùng **Redis Pub/Sub** cho các trường hợp đơn giản, chấp nhận mất message, cần độ trễ cực thấp (ví dụ: hiển thị trạng thái "đang gõ..." trong chat, cập nhật tỉ số trực tiếp mà chỉ cần giá trị mới nhất). Dùng **Kafka** khi cần đảm bảo không được mất dữ liệu, cần khả năng phát lại lịch sử (replay) để xử lý lỗi hoặc phân tích sau này, hoặc cần xử lý khối lượng event cực lớn (hệ thống log, event sourcing, xử lý dữ liệu streaming quy mô lớn).

**Ví dụ:** Một ứng dụng chat dùng Redis Pub/Sub để phát sự kiện "User A đang gõ..." — nếu một thiết bị đang offline bỏ lỡ sự kiện này cũng không sao, vì trạng thái "đang gõ" sẽ nhanh chóng được cập nhật lại ở lần gõ tiếp theo. Ngược lại, một hệ thống ghi nhận giao dịch tài chính cần dùng Kafka để đảm bảo mọi event `TransactionCreated` đều được lưu lại và xử lý đầy đủ, kể cả khi service xử lý tạm thời gặp sự cố và cần đọc lại từ đầu sau khi khôi phục.

**Mở rộng:**

- Tìm hiểu thêm Redis Streams — một tính năng mới hơn của Redis (khác với Pub/Sub truyền thống) có hỗ trợ lưu trữ và replay message, thu hẹp phần nào khoảng cách với Kafka cho các bài toán quy mô vừa.

**Nguồn trích dẫn:**

- Redis Docs – "Pub/Sub": https://redis.io/docs/latest/develop/interact/pubsub/
- Apache Kafka Documentation: https://kafka.apache.org/documentation/

---