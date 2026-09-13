# Bộ câu hỏi phỏng vấn Backend 
---

## Câu 1: Một vòng đời của HTTP Request diễn ra như thế nào?

**Câu trả lời:** Khi người dùng truy cập website hoặc gửi request từ trình duyệt/ứng dụng, request sẽ đi qua các bước chính sau:

1. **Phân giải DNS:** Tên miền (domain) được phân giải thành địa chỉ IP của server.
2. **Thiết lập kết nối:** Trình duyệt thực hiện bắt tay ba bước TCP (TCP three-way handshake); nếu là HTTPS thì có thêm bước bắt tay TLS để mã hóa đường truyền.
3. **Gửi request:** Trình duyệt gửi HTTP request đến server qua giao thức HTTP/HTTPS.
4. **Xử lý tại server:** Web server tiếp nhận request, tiến hành xác thực (Authentication – "bạn là ai") và phân quyền (Authorization – "bạn được làm gì").
5. **Xử lý nghiệp vụ:** Nếu hợp lệ, server xử lý logic nghiệp vụ, truy vấn database nếu cần.
6. **Trả response:** Server trả kết quả về client dưới dạng JSON, XML hoặc định dạng dữ liệu khác; trình duyệt nhận và hiển thị (render) kết quả.

**Ví dụ:** Khi bạn gõ `https://example.com` và nhấn Enter, trình duyệt tra DNS ra IP, bắt tay TCP/TLS, gửi `GET /` tới server, server trả về file HTML, trình duyệt tiếp tục tải CSS/JS/hình ảnh liên quan rồi render trang.

**Mở rộng:**

- Phân biệt rõ **Authentication** (xác thực danh tính) và **Authorization** (phân quyền hành động) – đây là hai khái niệm hay bị nhầm lẫn.
- Tìm hiểu thêm về TCP three-way handshake và TLS handshake để hiểu vì sao HTTPS có độ trễ ban đầu cao hơn HTTP.
- Tìm hiểu về CDN và DNS caching giúp rút ngắn bước 1 và 2.

**Nguồn trích dẫn:**

- MDN Web Docs – "An overview of HTTP": https://developer.mozilla.org/en-US/docs/Web/HTTP/Overview
- RFC 7230/7231 – HTTP/1.1 Message Syntax and Semantics

---

## Câu 2: Web Server và Application Server khác nhau như thế nào?

**Câu trả lời:**

- **Web Server:** Tiếp nhận HTTP request từ client và trả về tài nguyên tĩnh (static content) như HTML, CSS, JavaScript, hình ảnh (ví dụ: Nginx, Apache). Đóng vai trò như một "người gác cổng", xử lý lượng kết nối đồng thời tốt, hỗ trợ nén dữ liệu và cấu hình Reverse Proxy.
- **Application Server:** Thực thi logic nghiệp vụ động (dynamic content), kết nối trực tiếp với database và các service nội bộ. Nhận request đã được Web Server chuyển tiếp, xử lý rồi trả kết quả (JSON/HTML) ngược lại.

**Ví dụ:** Trong một hệ thống thực tế, Nginx đóng vai trò Web Server đứng trước, làm reverse proxy và phục vụ file tĩnh; phía sau là một Application Server viết bằng Go, Node.js hoặc Java (chạy trên Tomcat) xử lý logic nghiệp vụ và truy vấn database.

**Mở rộng:**

- Trong kiến trúc hiện đại, ranh giới này không còn tuyệt đối: Node.js/Express hay Go net/http có thể tự đóng cả hai vai trò cùng lúc (vừa nhận request, vừa xử lý logic) mà không cần Nginx đứng trước, dù trong production vẫn thường đặt reverse proxy phía trước để cân bằng tải, cache và bảo mật.
- Tìm hiểu thêm về Reverse Proxy và Load Balancer.

**Nguồn trích dẫn:**

- Nginx Docs – "What is a Web Server?": https://www.nginx.com/resources/glossary/web-server/

---

## Câu 3: Phân biệt sự khác nhau giữa Process và Thread?

**Câu trả lời:**

- **Process (Tiến trình):** Là một chương trình đang được thực thi, được hệ điều hành cấp phát một khối tài nguyên và vùng bộ nhớ độc lập. Các Process không thể can thiệp trực tiếp vào bộ nhớ của nhau; giao tiếp giữa các Process (IPC – Inter-Process Communication) tốn nhiều chi phí hệ thống hơn.
- **Thread (Luồng):** Là đơn vị thực thi nhỏ nhất bên trong một Process. Các Thread trong cùng một Process dùng chung không gian bộ nhớ và tài nguyên hệ thống. Giao tiếp giữa các Thread nhanh hơn nhiều, nhưng lập trình viên phải xử lý các vấn đề xung đột dữ liệu (race condition).

**Ví dụ:** Trình duyệt Chrome chạy mỗi tab như một Process riêng biệt (nếu một tab crash, các tab khác không bị ảnh hưởng). Trong khi đó, một web server viết bằng Java có thể dùng nhiều Thread trong cùng một Process để xử lý song song nhiều request.

**Mở rộng:**

- Xem thêm khái niệm Race Condition ở Câu 15.
- Go không dùng trực tiếp OS Thread cho mỗi request mà dùng Goroutine – một đơn vị thực thi nhẹ hơn nhiều do Go runtime tự quản lý.

**Nguồn trích dẫn:**

- Operating System Concepts – Silberschatz, Galvin, Gagne, chương "Process and Threads"

---

## Câu 4: Nguyên lý DRY và KISS trong Backend là gì?

**Câu trả lời:**

- **DRY (Don't Repeat Yourself):** Không viết lặp lại mã nguồn. Mọi logic nghiệp vụ chỉ nên được định nghĩa duy nhất một nơi trong hệ thống. Tuân thủ DRY giúp dễ bảo trì — khi cần sửa logic chỉ cần sửa tại một nơi tập trung thay vì dò tìm khắp dự án.
- **KISS (Keep It Simple, Stupid):** Giữ cho mã nguồn và kiến trúc hệ thống luôn đơn giản nhất có thể. Đừng cố phức tạp hóa vấn đề bằng những thiết kế thừa thãi (over-engineering). Code đơn giản giúp đội ngũ dễ đọc hiểu, dễ viết test và giảm thiểu bug ngầm.

**Ví dụ:**

- Vi phạm DRY: Copy cùng một đoạn logic tính thuế VAT vào 5 chỗ khác nhau trong code — khi thuế thay đổi phải sửa cả 5 nơi, rất dễ bỏ sót.
- Vi phạm KISS: Dùng một Design Pattern phức tạp (ví dụ Abstract Factory nhiều tầng) cho một tính năng chỉ cần một hàm đơn giản để giải quyết.

**Mở rộng:**

- DRY và KISS thường đi cùng với nguyên lý YAGNI (You Aren't Gonna Need It) — không xây dựng tính năng khi chưa thực sự cần.

**Nguồn trích dẫn:**

- The Pragmatic Programmer – Andrew Hunt, David Thomas (nguồn gốc của DRY)

---

## Câu 5: CORS lỗi là do đâu? Cách cấu hình Backend để giải quyết?

**Câu trả lời:** CORS (Cross-Origin Resource Sharing) xảy ra khi frontend gọi API từ một domain (origin) khác nhưng server chưa cho phép truy cập cross-origin. Ví dụ: frontend chạy ở `localhost:3000` gọi API ở `api.example.com`, trình duyệt sẽ chặn response nếu server không cấu hình CORS phù hợp — đây là cơ chế bảo mật của trình duyệt, không phải lỗi của server hay mạng.

**Backend xử lý bằng cách cấu hình các header:**

- **Access-Control-Allow-Origin:** Khai báo chính xác domain của frontend được phép gọi API (hạn chế dùng dấu `*` trong môi trường Production).
- **Access-Control-Allow-Methods:** Khai báo danh sách phương thức được phép (GET, POST, PUT, DELETE...).
- **Access-Control-Allow-Headers:** Khai báo các header được phép gửi lên (như Content-Type, Authorization).

**Ví dụ:** Với các request không đơn giản (ví dụ có header `Authorization` hoặc method `PUT`), trình duyệt sẽ tự động gửi trước một request `OPTIONS` gọi là "preflight request" để hỏi server xem request thật có được phép hay không, trước khi gửi request chính thức.

**Mở rộng:**

- Tìm hiểu sự khác nhau giữa Simple Request và Preflight Request trong CORS.

**Nguồn trích dẫn:**

- MDN Web Docs – "Cross-Origin Resource Sharing (CORS)": https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS

---

## Câu 6: Phân biệt cơ chế xác thực Session-based Authentication và Token-based Authentication (JWT)?

**Câu trả lời:**

- **Session-based:** Khi đăng nhập thành công, Server tạo một phiên làm việc (Session) lưu trên bộ nhớ server và trả về Client một `SessionID` qua Cookie. Ở các request sau, Client gửi kèm Cookie này, Server đối chiếu với bộ nhớ để xác thực.
    - _Nhược điểm:_ Tốn tài nguyên RAM khi lượng user lớn; khó mở rộng hệ thống sang nhiều server độc lập (tính chất Stateful) nếu không dùng thêm session store dùng chung (như Redis).
- **Token-based (JWT):** Server không lưu thông tin phiên làm việc (Stateless). Sau khi đăng nhập, server ký số tạo một chuỗi JSON Web Token (JWT) chứa thông tin user rồi trả về cho Client tự lưu (LocalStorage hoặc Cookie). Mỗi request sau, Client gửi kèm JWT trên Header; server chỉ cần dùng Secret Key để giải mã và xác thực mà không cần tra bộ nhớ.
    - _Nhược điểm:_ JWT khó thu hồi (revoke) trước khi hết hạn — nếu token bị lộ hoặc user bị khóa tài khoản, hệ thống vẫn phải chấp nhận token đó cho đến khi hết hạn, trừ khi triển khai thêm cơ chế blacklist hoặc thời gian sống (TTL) ngắn kết hợp refresh token.

**Ví dụ:** Một hệ thống microservices với nhiều server độc lập thường ưu tiên JWT vì không cần chia sẻ session store giữa các service; ngược lại, một website nội bộ quy mô nhỏ có thể dùng Session vì dễ kiểm soát và thu hồi ngay lập tức khi cần.

**Mở rộng:**

- Tìm hiểu cơ chế Refresh Token đi kèm Access Token để cân bằng giữa bảo mật và trải nghiệm người dùng.
- Tìm hiểu về Redis Session Store để giải quyết nhược điểm Stateful của Session khi scale ngang.

**Nguồn trích dẫn:**

- RFC 7519 – JSON Web Token (JWT)
- OWASP – "Session Management Cheat Sheet": https://cheatsheetseries.owasp.org/cheatsheets/Session_Management_Cheat_Sheet.html

---

## Câu 7: Làm thế nào để lưu trữ mật khẩu người dùng vào Database một cách an toàn nhất?

**Câu trả lời:** Mật khẩu tuyệt đối không nên lưu dạng plain text trong database. Backend cần hash mật khẩu trước khi lưu bằng các thuật toán chuyên dụng cho mật khẩu như bcrypt, scrypt hoặc Argon2 (không dùng các hàm hash nhanh và tổng quát như MD5 hay SHA-256 thuần vì chúng dễ bị tấn công brute-force do tốc độ tính toán quá nhanh). Ngoài ra, các thuật toán trên đã tự động sinh và lưu kèm salt (một chuỗi ngẫu nhiên) để hai người dùng có cùng mật khẩu vẫn cho ra chuỗi hash khác nhau, hạn chế nguy cơ tấn công rainbow table. Khi user đăng nhập, hệ thống hash lại mật khẩu vừa nhập rồi so sánh với chuỗi hash đã lưu.

**Ví dụ:** Một chuỗi hash bcrypt trông như sau: `$2b$12$KIXQ...` — trong đó đã bao gồm cả thông tin về số vòng lặp (cost factor) và salt, không cần lưu salt ở cột riêng.

**Mở rộng:**

- Tìm hiểu thêm khái niệm "Pepper" (một khóa bí mật bổ sung lưu ở phía server, tách biệt với database) như một lớp bảo vệ nâng cao.

**Nguồn trích dẫn:**

- OWASP – "Password Storage Cheat Sheet": https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html

---

## Câu 8: Lỗ hổng SQL Injection là gì? Làm thế nào để phòng tránh?

**Câu trả lời:** SQL Injection là lỗ hổng bảo mật xảy ra khi kẻ tấn công chèn câu lệnh SQL độc hại vào input của người dùng để truy cập hoặc thao túng database trái phép. Nguyên nhân gốc rễ là việc nối chuỗi (string concatenation) trực tiếp giá trị input của người dùng vào câu lệnh SQL mà không qua xử lý.

Cách phòng tránh phổ biến nhất là dùng **parameterized query** (hay còn gọi là **prepared statement**): thay vì nối chuỗi SQL với input, câu lệnh SQL được biên dịch trước với các placeholder (dấu `?` hoặc `$1`), sau đó giá trị input được truyền vào riêng biệt như dữ liệu thuần túy — database sẽ không bao giờ diễn giải phần dữ liệu này thành câu lệnh SQL, dù nó có chứa ký tự đặc biệt. Ngoài ra, cần validate dữ liệu đầu vào và giới hạn quyền truy cập database theo nguyên tắc đặc quyền tối thiểu (Principle of Least Privilege).

**Ví dụ:**

```sql
-- KHÔNG AN TOÀN: nối chuỗi trực tiếp
query := "SELECT * FROM users WHERE username = '" + input + "'"
-- Nếu input = "' OR '1'='1", câu lệnh thực thi trở thành:
-- SELECT * FROM users WHERE username = '' OR '1'='1'  → trả về toàn bộ user

-- AN TOÀN: dùng parameterized query
query := "SELECT * FROM users WHERE username = ?"
db.Query(query, input)  // input luôn được coi là dữ liệu, không phải mã lệnh
```

**Mở rộng:**

- Tìm hiểu thêm về ORM (Hibernate, Prisma, GORM...) — hầu hết ORM hiện đại đã tự động dùng parameterized query bên dưới, giúp giảm rủi ro SQL Injection nếu dùng đúng cách (tránh viết raw query bằng nối chuỗi).
- Đây là một trong những lỗi bảo mật cơ bản nhưng vẫn xuất hiện khá nhiều ở các hệ thống thiếu kiểm soát code review.

**Nguồn trích dẫn:**

- OWASP – "SQL Injection Prevention Cheat Sheet": https://cheatsheetseries.owasp.org/cheatsheets/SQL_Injection_Prevention_Cheat_Sheet.html

---

## Câu 9: Cơ chế phân quyền RBAC và ABAC khác nhau thế nào?

**Câu trả lời:**

- **RBAC (Role-Based Access Control):** Phân quyền dựa trên Vai trò. Hệ thống định nghĩa các vai trò cố định (Admin, Manager, User...), mỗi vai trò có một tập quyền hạn cụ thể. User được gán vào vai trò nào sẽ có toàn bộ quyền của vai trò đó. Cấu trúc đơn giản, dễ triển khai nhưng thiếu linh hoạt khi hệ thống phình to.
- **ABAC (Attribute-Based Access Control):** Phân quyền dựa trên Thuộc tính. Hệ thống quyết định cho phép truy cập dựa trên sự kết hợp của nhiều thuộc tính: thuộc tính của User (phòng ban, cấp bậc), thuộc tính của tài nguyên (loại tài liệu, độ mật), thuộc tính của môi trường (thời gian truy cập, địa chỉ IP). Rất linh hoạt, chi tiết nhưng triển khai phức tạp hơn nhiều.

**Ví dụ:** RBAC phù hợp cho một hệ thống nội bộ đơn giản: nhân viên phòng Kế toán được gán role "Accountant" thì mặc nhiên có quyền xem toàn bộ hóa đơn. ABAC phù hợp khi cần luật chi tiết hơn, ví dụ: "Chỉ nhân viên phòng Kế toán, đang truy cập trong giờ hành chính, từ mạng nội bộ công ty mới được xem hóa đơn có giá trị trên 100 triệu".

**Mở rộng:**

- Trong thực tế, nhiều hệ thống dùng mô hình lai (Hybrid) kết hợp RBAC làm nền và bổ sung một số luật ABAC cho các trường hợp đặc biệt.

**Nguồn trích dẫn:**

- NIST SP 800-162 – "Guide to Attribute Based Access Control (ABAC)"

---

## Câu 10: Khi nào nên chọn SQL và NoSQL? Nêu các trade-off cụ thể

**Câu trả lời:**

- **Chọn SQL (Relational DB – MySQL, PostgreSQL):** Khi dữ liệu có cấu trúc rõ ràng, mối quan hệ phức tạp (1-n, n-n) và hệ thống yêu cầu tính toàn vẹn dữ liệu tuyệt đối, tuân thủ nghiêm ngặt tính chất ACID (như hệ thống giao dịch tài chính, ví điện tử, quản lý đơn hàng).
- **Chọn NoSQL (Non-relational DB – MongoDB, Redis, Cassandra):** Khi dữ liệu không có cấu trúc cố định (schema-less), dạng Key-Value hoặc Document, hệ thống cần đọc/ghi tốc độ cao và khả năng mở rộng theo chiều ngang (horizontal scaling) linh hoạt trên nhiều máy chủ (như hệ thống chat, log tracking, newsfeed mạng xã hội).

**Trade-off:** Chọn SQL, bạn có sự nhất quán và toàn vẹn dữ liệu cao nhưng scale ngang khó khăn và tốn chi phí. Chọn NoSQL, bạn đạt hiệu năng cao và mở rộng dễ dàng nhưng phải chấp nhận đánh đổi tính nhất quán tức thời — hệ thống chỉ đạt tính nhất quán cuối cùng (Eventual Consistency).

**Ví dụ:** Hệ thống ví điện tử (MoMo, ZaloPay) cần SQL để đảm bảo số dư tài khoản luôn chính xác tuyệt đối sau mỗi giao dịch. Ngược lại, hệ thống lưu tin nhắn chat hoặc log truy cập website phù hợp với NoSQL vì lượng ghi rất lớn và không yêu cầu quan hệ phức tạp giữa các bản ghi.

**Mở rộng:**

- Tìm hiểu thêm về mô hình lai NewSQL (như CockroachDB, Google Spanner) cố gắng kết hợp ưu điểm của cả hai.
- Xem thêm CAP Theorem ở Câu 16 để hiểu rõ hơn về Eventual Consistency.

**Nguồn trích dẫn:**

- MongoDB Docs – "SQL to MongoDB Mapping Chart"

---

## Câu 11: Quy trình chuẩn hóa dữ liệu (Database Normalization từ 1NF đến 3NF) dùng để làm gì? Khi nào cần chống chuẩn hóa (Denormalization)?

**Câu trả lời:** Chuẩn hóa dữ liệu là quy trình tách các bảng dữ liệu lớn thành các bảng nhỏ hơn nhằm loại bỏ lưu trữ dữ liệu dư thừa, tránh các lỗi bất thường (anomalies) khi thêm/sửa/xóa dữ liệu và đảm bảo tính toàn vẹn.

- **1NF:** Mỗi ô dữ liệu phải chứa giá trị nguyên tố (không chia nhỏ được nữa), không chứa mảng hoặc danh sách giá trị trong một ô.
- **2NF:** Đạt 1NF, và toàn bộ thuộc tính không phải khóa phải phụ thuộc hoàn toàn vào khóa chính (loại bỏ phụ thuộc một phần — chỉ áp dụng khi bảng có khóa chính ghép nhiều cột).
- **3NF:** Đạt 2NF, và loại bỏ các phụ thuộc bắc cầu giữa các thuộc tính không phải khóa (một thuộc tính không khóa không được phụ thuộc vào một thuộc tính không khóa khác).

**Khi nào cần chống chuẩn hóa:** Khi hệ thống có lượng truy vấn đọc (Read) rất lớn và việc tuân thủ 3NF buộc phải thực hiện quá nhiều câu lệnh JOIN phức tạp giữa nhiều bảng, gây sụt giảm hiệu năng. Lúc này, ta chủ động chấp nhận dư thừa dữ liệu, gộp bớt bảng lại để tăng tốc độ đọc.

**Ví dụ:** Một bảng `Orders` lưu cả `customer_name` và `customer_address` lặp lại ở mỗi đơn hàng là vi phạm chuẩn hóa — nếu khách đổi địa chỉ, phải sửa ở tất cả các đơn hàng cũ. Chuẩn hóa đúng là tách ra bảng `Customers` riêng và `Orders` chỉ lưu `customer_id` tham chiếu tới. Ngược lại, một hệ thống báo cáo (dashboard) đọc dữ liệu liên tục có thể cố ý lưu thêm cột `customer_name` ngay trong bảng `Orders` (denormalize) để tránh phải JOIN mỗi lần hiển thị danh sách đơn hàng.

**Mở rộng:**

- Tìm hiểu thêm về BCNF (Boyce-Codd Normal Form) — dạng chuẩn chặt hơn 3NF.

**Nguồn trích dẫn:**

- Database System Concepts – Silberschatz, Korth, Sudarshan, chương "Relational Database Design"

---

## Câu 12: Database Index hoạt động thế nào bên dưới? Tại sao tạo quá nhiều Index lại làm chậm hệ thống?

**Câu trả lời:** Cơ chế hoạt động của Database Index giống mục lục của một cuốn sách. Thay vì quét toàn bộ bảng từ đầu đến cuối (Full Table Scan) với độ phức tạp O(N), DB sử dụng cấu trúc cây cân bằng B-Tree (hoặc B+ Tree) để lưu các giá trị index đã sắp xếp thứ tự, giúp thu hẹp phạm vi tìm kiếm xuống độ phức tạp O(log N), tăng tốc độ truy vấn đọc lên nhiều lần.

**Vì sao tạo quá nhiều Index lại làm chậm hệ thống:**

1. **Chậm ghi dữ liệu:** Mỗi khi thực hiện INSERT, UPDATE hoặc DELETE, database không chỉ cập nhật dữ liệu ở bảng gốc mà còn phải cập nhật lại toàn bộ các cấu trúc B-Tree của mọi index liên quan đến các cột bị thay đổi. Càng nhiều index, chi phí ghi càng lớn.
2. **Tốn dung lượng lưu trữ:** Mỗi index là một cấu trúc dữ liệu riêng, chiếm thêm không gian đĩa, đôi khi gần bằng kích thước bảng gốc.
3. **Query planner khó chọn đúng:** Khi có quá nhiều index trên cùng một bảng, trình tối ưu truy vấn (query planner) mất nhiều thời gian hơn để phân tích và đôi khi chọn nhầm index không tối ưu cho một câu truy vấn cụ thể.

Vì vậy, việc tạo index cần cân bằng giữa tốc độ đọc và chi phí ghi, chỉ nên đánh index trên các cột thường xuyên dùng để lọc (`WHERE`), sắp xếp (`ORDER BY`) hoặc join.

**Ví dụ:** Một bảng `orders` có 2 triệu bản ghi và 10 index khác nhau — mỗi lần INSERT một đơn hàng mới, hệ thống phải ghi vào bảng gốc và cập nhật cả 10 cấu trúc B-Tree, khiến tốc độ ghi chậm đi rõ rệt so với bảng chỉ có 1-2 index thiết yếu.

**Mở rộng:**

- Tìm hiểu Composite Index (index trên nhiều cột) và thứ tự cột ảnh hưởng thế nào đến hiệu quả sử dụng.
- Tìm hiểu lệnh `EXPLAIN`/`EXPLAIN ANALYZE` để xem query planner đang chọn index nào.

**Nguồn trích dẫn:**

- PostgreSQL Official Documentation – "Indexes": https://www.postgresql.org/docs/current/indexes.html
- Use The Index, Luke! – https://use-the-index-luke.com/

---

## Câu 13: Lỗi N+1 Query Problem là gì? Cách phát hiện và khắc phục?

**Câu trả lời:**

- **Bản chất lỗi:** Xảy ra khi dùng ORM (Hibernate, Sequelize, Prisma, GORM...) để lấy một danh sách dữ liệu có quan hệ với bảng khác. Thay vì thực hiện 1 câu lệnh duy nhất, ORM chạy 1 truy vấn để lấy danh sách cha (N bản ghi), sau đó tự động chạy thêm N câu truy vấn con để lấy dữ liệu liên quan của từng bản ghi cha — tổng cộng N+1 câu lệnh xuống DB, gây nghẽn kết nối.
- **Cách phát hiện:** Bật log SQL thực tế của ORM ở môi trường Dev; nếu thấy một chuỗi câu lệnh SELECT gần giống hệt nhau lặp lại liên tục, hệ thống đang dính N+1.
- **Cách khắc phục:** Dùng Eager Loading thay vì Lazy Loading, bằng các câu lệnh liên kết tường minh như `JOIN FETCH` (Java/Hibernate) hoặc các hàm nạp dữ liệu đồng thời có sẵn của thư viện (như `.include()`, `.populate()`, `Preload()` trong GORM) để gom dữ liệu vào chỉ 1-2 câu lệnh truy vấn.

**Ví dụ:**

```go
// LỖI N+1: lặp 1 query cha + N query con trong vòng lặp
var users []User
db.Find(&users)          // 1 query
for _, u := range users {
    db.Model(&u).Related(&u.Orders) // N query, mỗi user 1 query riêng
}

// KHẮC PHỤC: dùng Eager Loading — gom về 1-2 query
db.Preload("Orders").Find(&users)
```

**Mở rộng:**

- N+1 không chỉ xảy ra với ORM — cũng có thể xảy ra khi gọi tuần tự nhiều API nội bộ trong vòng lặp thay vì gọi gộp (batch API).

**Nguồn trích dẫn:**

- GORM Official Docs – "Preloading (Eager Loading)": https://gorm.io/docs/preload.html

---

## Câu 14: Cơ chế Blocking I/O và Non-blocking I/O ảnh hưởng thế nào đến hiệu năng server?

**Câu trả lời:**

- **Blocking I/O:** Thread thực hiện thao tác I/O (đọc file, gọi database, gọi network...) sẽ bị "đứng chờ" (block) cho đến khi thao tác đó hoàn tất mới được xử lý tiếp việc khác. Khi traffic lớn, mỗi thread bị chiếm giữ lâu, hệ thống cần cấp phát rất nhiều thread để phục vụ đồng thời nhiều request, gây tốn tài nguyên.
- **Non-blocking I/O:** Thread gửi yêu cầu I/O rồi tiếp tục xử lý công việc khác ngay, không chờ đợi; khi thao tác I/O hoàn tất, hệ thống sẽ được thông báo (thường qua cơ chế event loop hoặc callback) để xử lý kết quả. Nhờ đó một số lượng thread nhỏ vẫn có thể phục vụ hàng chục nghìn kết nối đồng thời.

**Ví dụ:** Node.js dùng một Event Loop đơn luồng (single-threaded) kết hợp Non-blocking I/O để xử lý hàng nghìn kết nối đồng thời mà không cần tạo nhiều thread. Go dùng Goroutine kết hợp một cơ chế Non-blocking I/O ẩn bên dưới (netpoller) để hàng chục nghìn goroutine có thể cùng chờ I/O mà không chiếm hẳn một OS thread mỗi goroutine.

**Mở rộng:**

- Tìm hiểu mô hình Reactor Pattern — nền tảng của hầu hết các runtime bất đồng bộ (Node.js, Nginx, Redis).

**Nguồn trích dẫn:**

- Node.js Official Docs – "The Node.js Event Loop": https://nodejs.org/en/learn/asynchronous-work/event-loop-timers-and-nexttick

---

## Câu 15: Race Condition là gì? So sánh Optimistic Locking và Pessimistic Locking

**Câu trả lời:** Race Condition là hiện tượng xung đột dữ liệu xảy ra khi có từ hai luồng (Thread hoặc Process) trở lên cùng đọc, tính toán và ghi đè giá trị vào cùng một vùng dữ liệu dùng chung tại gần như cùng một thời điểm, khiến kết quả cuối cùng sai lệch so với logic kỳ vọng.

**So sánh hai giải pháp khóa:**

- **Pessimistic Locking (khóa bi quan):** Giả định xung đột chắc chắn sẽ xảy ra. Khi một tiến trình vào đọc dữ liệu, nó khóa ngay bản ghi đó lại (ví dụ `SELECT ... FOR UPDATE`). Các tiến trình khác muốn đọc/ghi phải chờ đến khi tiến trình đầu tiên hoàn thành và giải phóng khóa. _Nhược điểm:_ dễ gây nghẽn cổ chai và tăng nguy cơ deadlock.
- **Optimistic Locking (khóa lạc quan):** Giả định xung đột hiếm khi xảy ra, không khóa khi đọc dữ liệu. Bảng dữ liệu được bổ sung thêm một cột `version` (hoặc timestamp). Khi ghi dữ liệu, hệ thống kiểm tra giá trị `version` hiện tại trong DB có còn trùng với `version` lúc vừa đọc ra hay không. Nếu trùng, ghi thành công và tăng `version` lên 1; nếu không trùng (đã có người khác sửa trước), hệ thống từ chối và yêu cầu thử lại. Phù hợp cho hệ thống có tỷ lệ đọc lớn hơn nhiều so với tỷ lệ ghi.

**Ví dụ:** Hai người dùng cùng cố mua nốt 1 vé xem phim cuối cùng cùng lúc. Với Pessimistic Locking, người đến trước sẽ khóa bản ghi vé đó lại, người sau phải chờ; nếu vé đã hết, hệ thống báo lỗi ngay khi họ được xử lý. Với Optimistic Locking, cả hai đều đọc thấy còn vé, nhưng khi ghi, chỉ người có `version` khớp mới đặt vé thành công, người còn lại nhận lỗi và phải thử lại.

**Mở rộng:**

- Tìm hiểu thêm về Deadlock và các chiến lược phòng tránh (timeout, wait-die, wound-wait).

**Nguồn trích dẫn:**

- Designing Data-Intensive Applications – Martin Kleppmann, chương "Weak Isolation Levels"

---

## Câu 16: Định lý CAP (CAP Theorem) trong hệ thống phân tán là gì? Tại sao không thể đạt cả 3 yếu tố cùng lúc?

**Câu trả lời:** Định lý CAP khẳng định rằng trong một hệ thống phân tán, khi xảy ra sự cố mạng (network partition), ta không thể đồng thời đảm bảo cả 3 yếu tố sau ở mức tối đa cùng lúc:

- **C (Consistency):** Mọi nút trong hệ thống đều trả về cùng một kết quả dữ liệu tại cùng một thời điểm truy vấn.
- **A (Availability):** Mọi request gửi lên hệ thống đều nhận được phản hồi (thành công hoặc thất bại), không bị treo hay timeout, kể cả khi có nút đang gặp sự cố.
- **P (Partition Tolerance):** Hệ thống vẫn tiếp tục hoạt động ngay cả khi kết nối mạng giữa các nút nội bộ bị đứt gãy, chia cắt.

Trong thực tế, Partition Tolerance gần như là điều kiện bắt buộc đối với mọi hệ phân tán thật (mạng luôn có khả năng lỗi), nên lựa chọn thực sự chỉ còn giữa Consistency và Availability khi partition xảy ra: hệ thống hoặc từ chối phục vụ để đảm bảo dữ liệu nhất quán tuyệt đối (CP), hoặc tiếp tục phục vụ và chấp nhận dữ liệu có thể tạm thời không đồng bộ giữa các nút (AP).

**Ví dụ:** Các hệ thống như Cassandra, DynamoDB thường ưu tiên AP (Availability) để đảm bảo dịch vụ không gián đoạn, chấp nhận eventual consistency. Ngược lại, các hệ thống ngân hàng dùng cơ sở dữ liệu quan hệ với cấu hình replication đồng bộ thường ưu tiên CP (Consistency) — thà tạm ngừng phục vụ còn hơn trả về số dư tài khoản sai.

**Mở rộng:**

- Tìm hiểu thêm mô hình PACELC — mở rộng của CAP, xét thêm trade-off giữa Latency và Consistency ngay cả khi không có partition.

**Nguồn trích dẫn:**

- Eric Brewer – "CAP Twelve Years Later: How the 'Rules' Have Changed" (2012)

---

---

# PHỤ LỤC: BÁO CÁO CHỈNH SỬA (không thuộc nội dung chính)

## 1. Sắp xếp lại thứ tự câu hỏi (dễ → khó)

Thứ tự cũ khá ngẫu nhiên (xen kẽ nền tảng web, bảo mật, database, hệ điều hành). Đã sắp xếp lại theo 4 nhóm chủ đề tăng dần độ khó:

1. Nền tảng Web/OS: HTTP lifecycle, Web Server vs App Server, Process vs Thread, DRY/KISS, CORS
2. Bảo mật & Xác thực: Session vs JWT, lưu mật khẩu, SQL Injection, RBAC vs ABAC
3. Cơ sở dữ liệu: SQL vs NoSQL, Normalization, Index, N+1 Query
4. Xử lý đồng thời & Hệ phân tán: Blocking/Non-blocking I/O, Race Condition & Locking, CAP Theorem

## 2. Lỗi chính tả / diễn đạt đã sửa

- "reuquest" → "request", "reponse" → "response", "phần quyền" → "phân quyền" (Câu HTTP lifecycle).
- Chuẩn hóa lại một số câu văn lủng củng, thiếu dấu chấm câu.

## 3. Nội dung thiếu chính xác / chưa đầy đủ đã bổ sung

- **Database Index:** phần placeholder "{Ở đây thiếu phần giải thích tại sao}" đã được thay bằng giải thích 3 lý do cụ thể (chi phí ghi tăng, tốn dung lượng, query planner khó chọn đúng index).
- **SQL Injection:** phần placeholder về parameterized query/prepared statement đã được giải thích rõ cơ chế và bổ sung ví dụ code minh họa (đoạn không an toàn vs đoạn an toàn).
- **Session vs JWT:** bổ sung nhược điểm quan trọng còn thiếu của JWT — khó thu hồi (revoke) trước khi hết hạn — để câu trả lời không thiên vị một chiều.
- **Lưu mật khẩu:** làm rõ vì sao không dùng MD5/SHA-256 thuần (tốc độ tính toán nhanh dễ bị brute-force), bổ sung ví dụ định dạng chuỗi hash bcrypt.
- **HTTP Request lifecycle:** bổ sung ngắn gọn bước bắt tay TCP/TLS bị thiếu trong luồng xử lý gốc, và làm rõ khác biệt Authentication/Authorization.

## 4. Thuật ngữ được giải thích thêm

- Authentication vs Authorization, IPC, Preflight Request (CORS), Pepper (bảo mật mật khẩu), PACELC (mở rộng của CAP).

## 5. Ví dụ minh họa đã bổ sung

Tất cả 16 câu đều đã có phần "Ví dụ" thực tế (trước đó phần lớn câu không có ví dụ cụ thể, chỉ có lý thuyết).

## 6. Nội dung giữ nguyên không thay đổi

Các câu về DRY/KISS, RBAC/ABAC, Process/Thread, Normalization, N+1 Query, Race Condition, CAP Theorem về cơ bản đã chính xác — chỉ bổ sung ví dụ và phần "Mở rộng", không sửa nội dung lý thuyết gốc.