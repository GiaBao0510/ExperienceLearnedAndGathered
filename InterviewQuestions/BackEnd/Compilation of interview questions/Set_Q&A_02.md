# Bộ câu hỏi phỏng vấn Backend 
# NHÓM A: THIẾT KẾ API

## Câu 1: Sự khác nhau giữa Authentication và Authorization là gì?

**Câu trả lời:**

- **Authentication (Xác thực):** Quá trình xác minh danh tính — trả lời câu hỏi "Bạn là ai?". Ví dụ: đăng nhập bằng username/password, OTP, vân tay.
- **Authorization (Phân quyền):** Quá trình xác định quyền hạn sau khi đã biết danh tính — trả lời câu hỏi "Bạn được phép làm gì?". Ví dụ: User thường chỉ được xem đơn hàng của mình, Admin được xem toàn bộ đơn hàng.

Authentication luôn diễn ra trước Authorization. Trên HTTP, hai khái niệm này thể hiện qua hai mã lỗi hay bị nhầm lẫn: `401 Unauthorized` thực chất nghĩa là **chưa xác thực được** (thiếu hoặc sai thông tin đăng nhập/token), còn `403 Forbidden` nghĩa là **đã xác thực thành công nhưng không đủ quyền** thực hiện hành động đó.

**Ví dụ:** Một user đăng nhập thành công (authenticated) vào hệ thống quản trị, nhưng khi cố truy cập trang "Quản lý người dùng" chỉ dành cho Admin, hệ thống trả về `403 Forbidden` vì user đã được xác thực nhưng không có quyền (authorization) truy cập tài nguyên đó.

**Mở rộng:**

- Xem thêm RBAC và ABAC — hai mô hình triển khai Authorization phổ biến.
- Tìm hiểu OAuth 2.0 Scope — cách giới hạn quyền hạn của một access token.

**Nguồn trích dẫn:**

- OWASP – "Authentication Cheat Sheet": https://cheatsheetseries.owasp.org/cheatsheets/Authentication_Cheat_Sheet.html
- RFC 7235 – HTTP/1.1 Authentication

---

## Câu 2: REST API và GraphQL khác nhau như thế nào? Khi nào chọn cái nào?

**Câu trả lời:**

- **REST:** Tổ chức API theo tài nguyên (resource), mỗi resource có một hoặc nhiều endpoint riêng (`/users`, `/users/1/orders`...), dùng các phương thức HTTP chuẩn (GET, POST, PUT, DELETE). Nhược điểm phổ biến là **over-fetching** (client nhận về nhiều field hơn cần dùng) và **under-fetching** (client phải gọi thêm nhiều request phụ để lấy đủ dữ liệu liên quan).
- **GraphQL:** Chỉ có một endpoint duy nhất. Client tự định nghĩa chính xác những field và mối quan hệ dữ liệu mình cần trong một câu query duy nhất, server trả về đúng những gì được yêu cầu — giải quyết triệt để vấn đề over-fetching/under-fetching. Đổi lại, GraphQL có độ phức tạp triển khai cao hơn: khó áp dụng HTTP caching truyền thống, dễ gặp vấn đề N+1 query ở tầng resolver phía server nếu không tối ưu (ví dụ dùng kỹ thuật DataLoader để gom batch), và có đường cong học tập cao hơn.

**Khi nào chọn:**

- Chọn **REST** khi API đơn giản, theo mô hình CRUD rõ ràng, cần tận dụng HTTP caching sẵn có, hoặc xây dựng public API cho nhiều đối tượng dùng chung.
- Chọn **GraphQL** khi client (đặc biệt mobile) cần linh hoạt lấy dữ liệu lồng nhau từ nhiều nguồn khác nhau trong một lần gọi, hoặc khi có nhiều loại client với nhu cầu dữ liệu khác nhau trên cùng một backend.

**Ví dụ:** Một app mobile chỉ cần hiển thị tên và avatar của user kèm 3 bài viết mới nhất. Với REST, có thể phải gọi `GET /users/1` rồi gọi thêm `GET /users/1/posts?limit=3` (under-fetching), hoặc `GET /users/1` trả về toàn bộ thông tin user kể cả những field không dùng tới (over-fetching). Với GraphQL, client chỉ cần gửi một query khai báo đúng field `name`, `avatar`, `posts(limit: 3) { title }` và nhận về chính xác dữ liệu đó trong một lần gọi.

**Mở rộng:**

- Tìm hiểu thêm gRPC — một lựa chọn khác thường dùng cho giao tiếp nội bộ giữa các microservices, tối ưu về hiệu năng hơn REST/GraphQL nhờ Protocol Buffers.

**Nguồn trích dẫn:**

- GraphQL Official Docs: https://graphql.org/learn/
- RESTfulAPI.net – "REST vs GraphQL"

---

## Câu 3: Bạn thiết kế API Versioning như thế nào và tại sao cần làm vậy?

**Câu trả lời:** API Versioning là kỹ thuật quản lý các thay đổi (breaking changes) của API mà không làm hỏng các client cũ đang sử dụng phiên bản trước đó. Khi thay đổi cấu trúc response, xóa/đổi tên field, hoặc thay đổi logic nghiệp vụ quan trọng, cần phát hành một phiên bản mới thay vì sửa trực tiếp API cũ.

**Các cách phổ biến để triển khai versioning:**

- **URI Versioning:** Nhúng version vào đường dẫn, ví dụ `/api/v1/users`, `/api/v2/users`. Đơn giản, dễ nhìn thấy trực tiếp trên URL, dễ test — cách phổ biến nhất trong thực tế dù về mặt lý thuyết không hoàn toàn "thuần REST" (một URI lẽ ra nên đại diện cho resource, không phải version).
- **Header Versioning:** Truyền version qua HTTP header, ví dụ `Accept: application/vnd.myapp.v2+json`. Giữ URL sạch nhưng khó test/debug trực tiếp trên trình duyệt hơn.
- **Query Parameter Versioning:** Ví dụ `/api/users?version=2`. Dễ triển khai nhưng ít được khuyến khích vì dễ bị quên truyền tham số.

**Ví dụ:** `/api/v1/users` trả về field `full_name`, sang `/api/v2/users` đổi thành tách riêng `first_name` và `last_name` — đây là một breaking change bắt buộc phải tăng version để không phá vỡ các client v1 đang chạy.

**Mở rộng:**

- Tìm hiểu về Deprecation Policy và header `Sunset`/`Deprecation` để thông báo trước cho client khi một phiên bản API sắp bị ngừng hỗ trợ.

**Nguồn trích dẫn:**

- Microsoft REST API Guidelines – "Versioning"
- Stripe API Docs – "Versioning": https://stripe.com/docs/api/versioning

---

## Câu 4: Bạn thiết kế API để hỗ trợ Pagination như thế nào?

**Câu trả lời:** Pagination là kỹ thuật chia nhỏ một tập dữ liệu lớn thành nhiều trang để trả về dần, tránh việc client phải tải toàn bộ dữ liệu cùng lúc gây chậm và tốn tài nguyên. Có hai cách tiếp cận chính:

- **Offset-based Pagination:** Dùng `LIMIT` và `OFFSET` (ví dụ `GET /posts?limit=20&offset=40`). Đơn giản, dễ hiểu, cho phép nhảy trực tiếp đến trang bất kỳ, nhưng có hai nhược điểm: hiệu năng giảm dần khi `OFFSET` lớn (database vẫn phải quét qua các bản ghi bị bỏ qua), và dữ liệu có thể bị trùng/thiếu nếu có bản ghi được thêm/xóa giữa các lần gọi trang.
- **Cursor-based Pagination (Keyset Pagination):** Dùng một "con trỏ" (cursor) trỏ đến vị trí bản ghi cuối cùng đã lấy, thường mã hóa từ một cột duy nhất và có thể sắp xếp được (như ID tăng dần hoặc `created_at` kết hợp ID). Request tiếp theo chỉ cần lấy các bản ghi sau cursor đó (ví dụ `GET /posts?limit=20&after=post_123`). Hiệu năng ổn định dù dữ liệu lớn, không bị trùng/thiếu bản ghi, nhưng không hỗ trợ nhảy trực tiếp đến một trang bất kỳ ở giữa.

**Ví dụ:** Facebook News Feed hay Twitter Timeline dùng cursor-based pagination vì dữ liệu liên tục được thêm mới (infinite scroll) — nếu dùng offset-based, người dùng sẽ thấy bài viết bị lặp hoặc bị bỏ sót khi có bài mới chen vào giữa lúc đang cuộn.

**Mở rộng:**

- Tìm hiểu cách các API lớn (Stripe, GitHub, Slack) triển khai cursor-based pagination trong thực tế.

**Nguồn trích dẫn:**

- Slack API Docs – "Pagination": https://api.slack.com/apis/pagination
- Stripe API Docs – "Pagination": https://stripe.com/docs/api/pagination

---

## Câu 5: Làm thế nào để xử lý Rate Limiting trong API?

**Câu trả lời:** Rate Limiting là kỹ thuật giới hạn số lượng request mà một client (user, IP, API key) được phép gọi trong một khoảng thời gian, nhằm bảo vệ hệ thống khỏi bị quá tải hoặc lạm dụng (spam, brute-force, DDoS nhẹ).

**Các thuật toán phổ biến:**

- **Fixed Window Counter:** Đếm số request trong một khung thời gian cố định (ví dụ mỗi phút reset về 0). Đơn giản nhưng có thể bị "burst" gấp đôi giới hạn nếu request dồn vào đúng ranh giới giữa hai khung giờ.
- **Sliding Window:** Cải tiến của Fixed Window, tính toán số request dựa trên cửa sổ thời gian trượt liên tục thay vì cố định, cho kết quả chính xác hơn.
- **Token Bucket:** Mỗi client có một "xô" chứa token, token được nạp vào xô theo tốc độ cố định; mỗi request tiêu tốn 1 token. Cho phép xử lý burst request trong giới hạn số token còn lại trong xô — thuật toán được dùng phổ biến nhất trong thực tế.
- **Leaky Bucket:** Xử lý request theo tốc độ đều đặn như "nước rò rỉ", giúp làm mượt (smooth) traffic thay vì cho phép burst.

Rate Limiting thường được triển khai tập trung tại tầng API Gateway (xem Câu về API Gateway ở Nhóm B) và cần lưu trạng thái đếm dùng chung (thường bằng Redis) nếu hệ thống chạy nhiều instance song song. Khi vượt giới hạn, server trả về mã `429 Too Many Requests` kèm header như `Retry-After` để client biết khi nào nên thử lại.

**Ví dụ:** Giới hạn 100 request/phút cho mỗi API key bằng thuật toán Token Bucket, lưu số token còn lại trong Redis với key là `api_key` và TTL tự làm mới mỗi phút.

**Mở rộng:**

- Tìm hiểu thêm cách Cloudflare, Nginx (`limit_req`) triển khai rate limiting ở tầng hạ tầng.

**Nguồn trích dẫn:**

- Cloudflare Blog – "How we built rate limiting capable of scaling to millions of domains"
- RFC 6585 – Additional HTTP Status Codes (định nghĩa mã 429)

---

## Câu 6: Idempotency trong API là gì và tại sao quan trọng?

**Câu trả lời:** Một thao tác được gọi là **idempotent** (bất biến khi lặp lại) nếu gọi thao tác đó một lần hay nhiều lần với cùng một input đều cho ra cùng một kết quả cuối cùng, không gây thêm tác dụng phụ nào khác. Theo chuẩn REST, các phương thức `GET`, `PUT`, `DELETE` vốn được thiết kế để có tính idempotent, còn `POST` thì không.

**Tại sao quan trọng:** Trong hệ thống thực tế, request có thể bị gọi lặp lại ngoài ý muốn — ví dụ client bị timeout và tự động retry, hoặc người dùng bấm nút "Thanh toán" hai lần do mạng chậm. Nếu API tạo giao dịch (`POST /payments`) không có cơ chế idempotency, hệ thống có thể vô tình trừ tiền hai lần cho cùng một đơn hàng.

**Cách triển khai phổ biến:** Client tự sinh một `Idempotency-Key` (thường là UUID) duy nhất cho mỗi thao tác nghiệp vụ và gửi kèm trong header request. Server lưu lại kết quả xử lý gắn với key đó; nếu nhận được request có cùng key trong một khoảng thời gian nhất định, server trả về ngay kết quả đã lưu trước đó thay vì xử lý lại từ đầu.

**Ví dụ:** Stripe yêu cầu client gửi header `Idempotency-Key` khi gọi API tạo thanh toán — nếu do lỗi mạng client gửi lại đúng request đó với cùng key, Stripe đảm bảo chỉ trừ tiền một lần duy nhất.

**Mở rộng:**

- Tìm hiểu thêm về các mức đảm bảo giao hàng thông điệp trong hệ phân tán: at-most-once, at-least-once, exactly-once delivery.

**Nguồn trích dẫn:**

- Stripe API Docs – "Idempotent Requests": https://stripe.com/docs/api/idempotent_requests
- MDN Web Docs – "HTTP request methods" (mục Idempotent)

---

## Câu 7: Webhook khác gì so với Polling? Khi nào dùng cái nào?

**Câu trả lời:**

- **Polling:** Client chủ động gửi request lặp đi lặp lại theo chu kỳ (ví dụ mỗi 5 giây) để hỏi server "có gì mới không?". Cách triển khai đơn giản nhưng gây lãng phí tài nguyên (phần lớn request trả về "không có gì mới") và có độ trễ tối đa bằng chu kỳ polling.
- **Webhook:** Server chủ động gọi ngược lại (callback) đến một URL do client đăng ký sẵn ngay khi có sự kiện xảy ra. Gần như thời gian thực (real-time) và tiết kiệm tài nguyên hơn nhiều so với polling, nhưng yêu cầu client phải có một endpoint truy cập được từ bên ngoài (public), đồng thời cần xử lý thêm các vấn đề như retry khi gọi thất bại và xác thực chữ ký (signature) để đảm bảo webhook thực sự đến từ nguồn tin cậy.

**Khi nào dùng:** Dùng **Webhook** khi cần cập nhật gần thời gian thực trong giao tiếp server-to-server và bên nhận có thể host một endpoint công khai (ví dụ cổng thanh toán báo kết quả giao dịch). Dùng **Polling** khi bên nhận không thể mở endpoint công khai (ví dụ ứng dụng mobile nằm sau NAT/firewall) hoặc tần suất sự kiện quá thấp, không đáng để đầu tư hạ tầng nhận webhook.

**Ví dụ:** Cổng thanh toán VNPay/Stripe gửi webhook đến server của bạn ngay khi giao dịch được xử lý xong, thay vì để hệ thống của bạn phải liên tục gọi API hỏi "giao dịch X đã xong chưa?" mỗi vài giây.

**Mở rộng:**

- Tìm hiểu thêm WebSocket và Server-Sent Events (SSE) — hai cơ chế khác cho phép server đẩy dữ liệu thời gian thực trực tiếp đến client (thường dùng cho ứng dụng web/mobile hơn là giao tiếp server-to-server).

**Nguồn trích dẫn:**

- Stripe Docs – "Webhooks": https://stripe.com/docs/webhooks

---

## Câu 8: Bạn document API như thế nào và tại sao điều đó quan trọng?

**Câu trả lời:** Tài liệu API (API Documentation) mô tả chi tiết cách một API hoạt động: endpoint, phương thức HTTP, tham số đầu vào, cấu trúc dữ liệu trả về, mã lỗi và cách xác thực — đóng vai trò như một "hợp đồng" giữa bên cung cấp API và bên sử dụng (frontend, mobile, đối tác thứ ba). Tài liệu tốt giúp giảm đáng kể thời gian tích hợp và số lượng câu hỏi hỗ trợ lặp lại.

**Cách triển khai phổ biến:** Sử dụng chuẩn **OpenAPI Specification** (trước đây gọi là Swagger) để mô tả API dưới dạng file YAML/JSON có cấu trúc chuẩn, từ đó có thể tự động sinh ra giao diện tài liệu tương tác (Swagger UI) cho phép người đọc thử gọi API trực tiếp trên trình duyệt. Một tài liệu API tốt cần có: mô tả endpoint và mục đích sử dụng, ví dụ request/response thực tế, danh sách mã lỗi có thể xảy ra, và cách xác thực (API Key, Bearer Token...).

**Ví dụ (rút gọn một endpoint theo chuẩn OpenAPI):**

```yaml
/users/{id}:
  get:
    summary: Lấy thông tin một người dùng theo ID
    parameters:
      - name: id
        in: path
        required: true
        schema:
          type: integer
    responses:
      '200':
        description: Thành công
      '404':
        description: Không tìm thấy user
```

**Mở rộng:**

- Tìm hiểu thêm mô hình **Contract-first** (thiết kế đặc tả OpenAPI trước, sinh code sau) so với **Code-first** (viết code trước rồi tự động sinh tài liệu từ annotation/comment).

**Nguồn trích dẫn:**

- OpenAPI Specification: https://swagger.io/specification/

---

# NHÓM B: KIẾN TRÚC HỆ THỐNG NÂNG CAO

## Câu 9: Phân biệt Cache-aside, Write-through và Write-back

**Câu trả lời:**

- **Cache-aside (Lazy Loading):** Ứng dụng tương tác trực tiếp với cả Cache và DB. Khi cần đọc dữ liệu, ứng dụng tìm trong Cache trước. Nếu không có (Cache Miss), ứng dụng tự đọc từ DB, trả về cho người dùng và chủ động ghi dữ liệu đó vào Cache để phục vụ lần sau. Đây là chiến lược phổ biến nhất.
- **Write-through:** Ứng dụng coi Cache là cổng ghi dữ liệu duy nhất. Khi có dữ liệu mới, ứng dụng ghi vào Cache; hệ thống cache sẽ tự động ghi tiếp dữ liệu đó xuống DB một cách đồng bộ ngay lập tức rồi mới báo thành công cho ứng dụng. Dữ liệu giữa Cache và DB luôn đồng nhất tuyệt đối, nhưng tốc độ ghi bị chậm hơn do phải chờ cả hai bước hoàn tất.
- **Write-back (Write-behind):** Khi có dữ liệu mới, ứng dụng ghi vào Cache và nhận ngay phản hồi thành công. Dữ liệu ghi xuống DB được gom lại và đồng bộ bất đồng bộ theo lô (batch) sau một khoảng thời gian định sẵn. Tốc độ ghi cực nhanh, nhưng có rủi ro mất dữ liệu nếu Cache bị sập đột ngột trước khi kịp đồng bộ xuống DB.

**Ví dụ:** Một trang tin tức dùng Cache-aside cho các bài viết (đọc nhiều, ghi ít). Một hệ thống ví điện tử cần dữ liệu số dư luôn chính xác nên phù hợp với Write-through hơn. Một hệ thống ghi log lượt xem video có thể chấp nhận Write-back để đạt throughput ghi cực cao, chấp nhận rủi ro mất một phần dữ liệu log nếu sự cố xảy ra.

**Mở rộng:**

- Tìm hiểu thêm chiến lược **Cache Invalidation** (khi nào và làm sao xóa/cập nhật cache đã lưu) — một trong những bài toán khó nhất của caching.

**Nguồn trích dẫn:**

- AWS – "Caching Strategies": https://aws.amazon.com/caching/best-practices/

---

## Câu 10: Khi nào nên chuyển đổi hệ thống từ kiến trúc Monolithic sang Microservices?

**Câu trả lời:** Monolithic phù hợp với sản phẩm nhỏ hoặc team chưa quá lớn vì dễ phát triển, dễ test và deploy hơn (toàn bộ code nằm trong một codebase, một lần deploy duy nhất). Tuy nhiên, khi hệ thống mở rộng, nhiều module phụ thuộc chặt vào nhau khiến mỗi lần deploy đều ảnh hưởng đến toàn bộ hệ thống, doanh nghiệp thường cân nhắc chuyển sang Microservices.

**Microservices phù hợp khi cần:**

- Scale từng service độc lập (ví dụ chỉ cần scale service xử lý thanh toán vào giờ cao điểm mà không cần scale toàn bộ hệ thống).
- Deploy riêng từng module mà không ảnh hưởng các phần khác.
- Nhiều team cùng phát triển song song, mỗi team sở hữu một hoặc vài service riêng.
- Hệ thống có traffic lớn và business domain phức tạp, cần tách bạch rõ ràng theo từng nghiệp vụ (Bounded Context).

**Ví dụ:** Một sàn thương mại điện tử ban đầu là một ứng dụng Monolithic duy nhất. Khi lượng đơn hàng tăng vọt vào các đợt sale lớn, team quyết định tách riêng service "Đặt hàng" và service "Thanh toán" thành các microservices độc lập để có thể scale riêng hai service này mà không cần scale cả hệ thống catalog sản phẩm vốn không bị áp lực tải tương tự.

**Mở rộng:**

- Không nên chuyển sang Microservices chỉ vì "công nghệ hot" — chi phí vận hành tăng đáng kể (xem Câu 13). Nhiều đội ngũ nhỏ vẫn vận hành Monolithic rất hiệu quả trong nhiều năm.
- Tìm hiểu mô hình trung gian "Modular Monolith" — giữ một codebase/deploy duy nhất nhưng tổ chức code theo module tách bạch rõ ràng, dễ tách thành microservices sau này nếu cần.

**Nguồn trích dẫn:**

- Martin Fowler – "MonolithFirst": https://martinfowler.com/bliki/MonolithFirst.html

---

## Câu 11: API Gateway Pattern đóng vai trò gì trong hệ thống Microservices?

**Câu trả lời:** API Gateway đóng vai trò như một điểm giao tiếp trung tâm duy nhất (Single Entry Point) đứng trước toàn bộ cụm Microservices phía sau. Thay vì Client (Mobile, Web) phải tự biết địa chỉ và gọi trực tiếp tới hàng chục service nhỏ bên trong, Client chỉ cần biết duy nhất một địa chỉ của API Gateway. API Gateway chịu trách nhiệm:

- **Routing:** Định tuyến request đến đúng service bên trong xử lý.
- **Authentication/Authorization:** Xác thực token tập trung ngay tại cổng, giảm việc lặp lại logic bảo mật ở từng microservice con.
- **Rate Limiting & Throttling:** Giới hạn số lượng request của từng user để bảo vệ hệ thống phía sau.
- **Protocol Translation:** Chuyển đổi giao thức (ví dụ nhận HTTP từ client bên ngoài, chuyển sang gRPC tốc độ cao để các service nội bộ giao tiếp với nhau).

**Ví dụ:** Client mobile gọi `GET /api/orders/123` đến API Gateway; Gateway xác thực token, sau đó định tuyến request nội bộ tới Order Service (qua gRPC), Order Service gọi tiếp Payment Service để lấy trạng thái thanh toán, rồi Gateway gộp kết quả trả về cho client dưới dạng một response JSON duy nhất.

**Mở rộng:**

- Tìm hiểu mẫu **Backend For Frontend (BFF)** — biến thể của API Gateway, mỗi loại client (web, mobile) có một gateway riêng được tối ưu theo đúng nhu cầu dữ liệu của client đó.

**Nguồn trích dẫn:**

- Microsoft Azure Architecture Center – "API Gateway pattern"

---

## Câu 12: Message Queue (Kafka, RabbitMQ) giải quyết những bài toán nào trong kiến trúc Event-Driven Architecture?

**Câu trả lời:** Kafka và RabbitMQ dùng để xử lý giao tiếp bất đồng bộ giữa các service. Thay vì service A gọi trực tiếp service B và phải chờ phản hồi ngay (giao tiếp đồng bộ), service A chỉ cần publish event vào message queue, còn service B sẽ tiêu thụ (consume) và xử lý sau đó. Điều này giúp hệ thống:

- Giảm sự phụ thuộc chặt (coupling) giữa các service — service A không cần biết service B tồn tại hay đang hoạt động bình thường không.
- Tăng khả năng scale — nhiều consumer có thể cùng xử lý song song các message trong queue.
- Chống nghẽn khi traffic tăng đột biến — queue đóng vai trò bộ đệm hấp thụ lượng request tăng vọt thay vì làm sập trực tiếp service xử lý.

Kafka mạnh về streaming dữ liệu khối lượng lớn và throughput cao, thường dùng để lưu trữ và replay lại lịch sử event (event sourcing, log dữ liệu lớn). RabbitMQ phổ biến ở các hệ thống cần queue processing ổn định, có nhiều pattern định tuyến message linh hoạt (exchange/routing key) và dễ triển khai hơn cho các bài toán quy mô vừa và nhỏ.

**Ví dụ:** Khi một đơn hàng được tạo thành công, thay vì service Order gọi trực tiếp và chờ service Email, service Inventory, service Analytics xử lý xong mới trả kết quả cho client, service Order chỉ cần publish event `OrderCreated` vào Kafka/RabbitMQ; ba service kia độc lập lắng nghe và xử lý phần việc của mình mà không làm chậm phản hồi cho người dùng.

**Mở rộng:**

- Xem thêm Câu 13 (Pub/Sub vs Message Queue thông thường) để hiểu cách các service tiêu thụ event theo mô hình nào.

**Nguồn trích dẫn:**

- Apache Kafka Documentation: https://kafka.apache.org/documentation/
- RabbitMQ Documentation: https://www.rabbitmq.com/tutorials

---

## Câu 13: Phân biệt cơ chế Pub/Sub và Message Queue thông thường

**Câu trả lời:**

- **Message Queue thông thường (Point-to-Point):** Hoạt động theo nguyên lý một-đối-một. Một message được đẩy vào queue sẽ chỉ có duy nhất một consumer tiếp nhận và xử lý. Sau khi consumer xử lý thành công, message đó biến mất khỏi queue. Thích hợp để phân chia đều các tác vụ nặng cho nhiều worker xử lý song song (ví dụ hàng đợi xử lý ảnh, mỗi ảnh chỉ cần một worker xử lý).
- **Pub/Sub (Publish/Subscribe):** Hoạt động theo nguyên lý một-đối-nhiều. Publisher phát một message vào một topic. Toàn bộ các subscriber đã đăng ký lắng nghe topic đó đều nhận được một bản sao của message để xử lý độc lập, đồng thời. Thích hợp để phát tán sự kiện (event-driven) cho toàn bộ hệ sinh thái phía sau cùng phản ứng lại một sự kiện.

**Ví dụ:** Khi có 1000 ảnh cần resize, dùng mô hình Queue point-to-point: mỗi ảnh (message) chỉ cần một trong số các worker xử lý — worker nào rảnh sẽ nhận việc, đảm bảo tải được chia đều. Ngược lại, khi có sự kiện `UserRegistered`, dùng mô hình Pub/Sub: cả service gửi email chào mừng, service tạo ví điểm thưởng và service ghi log analytics đều cần nhận và xử lý riêng cùng một sự kiện đó — nếu dùng Queue thông thường, chỉ một trong ba service này nhận được message và hai service còn lại sẽ bị bỏ lỡ.

**Mở rộng:**

- Trong thực tế, Kafka có thể mô phỏng cả hai mô hình tùy cách cấu hình Consumer Group (nhiều consumer trong cùng group chia nhau xử lý message giống Queue; nhiều consumer group khác nhau cùng đọc một topic độc lập giống Pub/Sub).

**Nguồn trích dẫn:**

- Apache Kafka Docs – "Consumer Groups": https://kafka.apache.org/documentation/#intro_consumers

---

## Câu 14: Những khó khăn lớn nhất khi vận hành Microservices là gì?

**Câu trả lời:** Microservices giúp hệ thống linh hoạt hơn nhưng cũng kéo theo nhiều vấn đề vận hành phức tạp:

- **Quản lý tính nhất quán dữ liệu (Distributed Transactions):** Mỗi service quản lý một DB độc lập, nên thực hiện một chuỗi giao dịch đi qua nhiều service mà yêu cầu tính chất ACID truyền thống (transaction xuyên suốt nhiều DB) gần như bất khả thi. Buộc phải áp dụng các mô hình phức tạp hơn như **Saga Pattern** (chuỗi các bước xử lý theo kiểu event-driven, kèm các tác vụ bù trừ — compensating transaction — để hoàn tác khi có bước nào đó thất bại).
- **Giám sát hệ thống (Distributed Tracing & Observability):** Khi request của một user bị lỗi, rất khó xác định nó thất bại ở service nào trong chuỗi có thể lên tới hàng chục service liên kết. Bắt buộc phải triển khai các hệ thống tracing chuyên dụng như Jaeger, Zipkin kết hợp OpenTelemetry, gắn một `Correlation ID` (hay Trace ID) xuyên suốt toàn bộ luồng xử lý của một request.
- **Độ trễ mạng (Network Latency) và độ phức tạp hạ tầng:** Giao tiếp giữa các service qua mạng (REST, gRPC) làm tăng độ trễ tổng thể so với gọi hàm nội bộ trong một Monolithic, đồng thời đòi hỏi hạ tầng DevOps phức tạp hơn nhiều (Kubernetes, Service Mesh, CI/CD tự động hóa cho hàng chục service độc lập).

**Ví dụ:** Trong một luồng đặt hàng đi qua Order Service → Payment Service → Inventory Service, nếu Payment Service trừ tiền thành công nhưng Inventory Service báo hết hàng, hệ thống cần một cơ chế Saga để tự động hoàn tiền (compensating transaction) thay vì để dữ liệu rơi vào trạng thái không nhất quán giữa các service.

**Mở rộng:**

- Tìm hiểu thêm hai kiểu triển khai Saga: **Choreography** (các service tự lắng nghe event của nhau, không có nhạc trưởng) và **Orchestration** (có một service trung tâm điều phối toàn bộ luồng giao dịch).

**Nguồn trích dẫn:**

- microservices.io – "Pattern: Saga": https://microservices.io/patterns/data/saga.html
- OpenTelemetry Docs: https://opentelemetry.io/docs/

---

---

# PHỤ LỤC: BÁO CÁO CHỈNH SỬA (không thuộc nội dung chính)

## 1. Lỗi cấu trúc nghiêm trọng đã sửa

- File gốc thiếu hoàn toàn câu trả lời cho 8 câu hỏi: **REST vs GraphQL, API Versioning, Idempotency, Rate Limiting, Authentication vs Authorization, API Documentation, Webhook vs Polling, Pagination** (chỉ có tiêu đề câu hỏi, không có nội dung). Đã viết bổ sung đầy đủ câu trả lời, ví dụ, phần mở rộng và nguồn trích dẫn cho cả 8 câu này.
- File gốc bị nhảy số: có Câu 8 rồi nhảy thẳng sang Câu 10 (thiếu Câu 9). Đã đánh số lại liên tục từ 1 đến 14 sau khi sắp xếp.

## 2. Sắp xếp lại cấu trúc và thứ tự

Nội dung được chia thành 2 nhóm rõ ràng thay vì trộn lẫn ngẫu nhiên như bản gốc:

- **Nhóm A – Thiết kế API** (8 câu, độ khó tăng dần: Auth vs Authz → REST/GraphQL → Versioning → Pagination → Rate Limiting → Idempotency → Webhook vs Polling → Documentation) — đây là nhóm kiến thức nền tảng, gần với công việc hàng ngày hơn.
- **Nhóm B – Kiến trúc hệ thống nâng cao** (6 câu, độ khó tăng dần: Caching Strategies → Monolithic vs Microservices → API Gateway → Message Queue → Pub/Sub → Khó khăn khi vận hành Microservices) — đặt câu khó nhất (vận hành Microservices, đòi hỏi tổng hợp nhiều kiến thức phân tán) ở cuối cùng.

## 3. Nội dung đã có sẵn (Câu Cache-aside/Write-through/Write-back, Monolithic vs Microservices, khó khăn Microservices, API Gateway, Message Queue, Pub/Sub)

Về cơ bản đã chính xác, chỉ bổ sung thêm phần "Ví dụ" thực tế (bản gốc chưa có ví dụ cụ thể nào) và phần "Mở rộng", không thay đổi nội dung lý thuyết gốc.

## 4. Thuật ngữ được giải thích thêm

Over-fetching/Under-fetching, Idempotency Key, Token Bucket/Leaky Bucket, Cursor-based Pagination, Correlation ID/Trace ID, Compensating Transaction, Bounded Context, Backend For Frontend (BFF).