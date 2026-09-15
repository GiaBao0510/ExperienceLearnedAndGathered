# Bộ Câu Hỏi Phỏng Vấn Microservices 

## Câu 1: Các công ty nổi tiếng sử dụng kiến trúc Microservices?

- Twitter
- Netflix
- Amazon

**Ví dụ**: Netflix là một trong những ví dụ được nhắc tới nhiều nhất — họ chuyển từ kiến trúc monolithic sang hàng trăm microservice để phục vụ hàng triệu người dùng streaming đồng thời trên toàn cầu, đồng thời cũng là nơi phát triển nhiều công cụ nổi tiếng cho hệ sinh thái microservices (ví dụ Hystrix, Eureka).

### Mở rộng

- Ngoài ba cái tên trên, các công ty lớn khác cũng được biết đến với kiến trúc microservices: Uber, Spotify, Airbnb, Shopify.

### Nguồn trích dẫn

- Chris Richardson — [microservices.io](https://microservices.io/patterns/microservices.html)

---

## Câu 2: Giải thích cách microservice giao tiếp với nhau?

Giao tiếp giữa các microservice thường thực hiện theo một trong các cách sau:

- **HTTP/REST** với JSON, hoặc một giao thức nhị phân (ví dụ Protocol Buffers trong gRPC) — dùng cho kiểu giao tiếp đồng bộ (request/response), nơi bên gọi chờ phản hồi ngay.
- **WebSocket** — dùng cho giao tiếp dạng streaming, khi cần truyền dữ liệu liên tục hai chiều giữa client và server.
- **Message broker** (bất đồng bộ) — bên gửi đẩy message vào broker, bên nhận xử lý khi sẵn sàng, không cần chờ phản hồi ngay lập tức.

**RabbitMQ**, **Kafka**... là các message broker phổ biến, mỗi loại được thiết kế phù hợp với những nhu cầu xử lý message khác nhau (ví dụ Kafka phù hợp cho luồng dữ liệu lớn, cần lưu trữ lại log sự kiện; RabbitMQ phù hợp cho hàng đợi tác vụ có độ trễ thấp).

**Ví dụ**: `order-service` gọi đồng bộ tới `payment-service` qua REST để xác nhận thanh toán ngay lập tức, nhưng sau khi tạo đơn hàng thành công, nó publish một event `OrderCreated` bất đồng bộ qua Kafka để `notification-service` xử lý gửi email mà không làm chậm phản hồi cho người dùng.

### Mở rộng

- gRPC — framework giao tiếp đồng bộ hiệu năng cao, ngày càng phổ biến giữa các microservice nội bộ.
- Sự khác biệt giữa giao tiếp đồng bộ và bất đồng bộ, và ảnh hưởng của mỗi loại tới độ trễ, độ tin cậy của hệ thống.

### Nguồn trích dẫn

- Chris Richardson — [microservices.io — Pattern: Messaging](https://microservices.io/patterns/communication-style/messaging.html)

---

## Câu 3: Domain-Driven Design (DDD) là gì?

**DDD (Domain-Driven Design)** là một phương pháp tiếp cận trong phân tích và phát triển phần mềm, dùng để giải quyết các bài toán nghiệp vụ phức tạp. Ý tưởng cốt lõi là xây dựng kết nối chặt chẽ giữa thiết kế phần mềm và mô hình nghiệp vụ (domain model) trong suốt vòng đời phát triển sản phẩm. Để đạt được điều đó, DDD nhấn mạnh ba nguyên tắc cơ bản:

- Trọng tâm của dự án là các nguyên tắc và logic nghiệp vụ (domain logic), không phải công nghệ.
- Thiết kế phần mềm cần phản ánh chính xác mô hình nghiệp vụ thực tế.
- Cần sự cộng tác liên tục giữa kỹ sư phần mềm và chuyên gia nghiệp vụ (domain expert).

![Domain-Driven Design](https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRUK5aWX7k1Xbuik65V3Y7ZhDbWswDe6t8uoA94G0Z4uzhNu6IZSI31r84&s=10)

Khi kết hợp DDD với kiến trúc Microservices, kết quả phân tích domain theo DDD giúp xác định cách tổ chức và phát triển các microservice sao cho mỗi service giải quyết đúng một phần nghiệp vụ tương ứng. Việc kết hợp này thường theo một quy trình gồm các bước:

1. Phân tích nghiệp vụ, xây dựng domain model.
2. Định nghĩa ranh giới ngữ cảnh (**Bounded Context**).
3. Định nghĩa các đối tượng nghiệp vụ: entity, aggregate, và domain service bên trong mỗi context.
4. Từ các Bounded Context đã xác định, quyết định microservice nào cần xây dựng.

![Quy trình kết hợp DDD và Microservices](https://img.magnific.com/premium-vector/process-workflow-timeline-modern-design-template-infographics-with-8-steps_77494-816.jpg?semt=ais_hybrid&w=740&q=80)

**Ví dụ**: Với một hệ thống thương mại điện tử, sau khi phân tích domain, đội ngũ xác định được các Bounded Context: Đặt hàng, Thanh toán, Kho. Mỗi Bounded Context này trở thành ranh giới hợp lý để tách thành một microservice riêng: `order-service`, `payment-service`, `inventory-service`.

### Mở rộng

- Bounded Context, Ubiquitous Language — các khái niệm nền tảng khác trong DDD.
- Strategic Design vs Tactical Design trong DDD.

### Nguồn trích dẫn

- Eric Evans — _Domain-Driven Design: Tackling Complexity in the Heart of Software_
- Martin Fowler — [BoundedContext](https://martinfowler.com/bliki/BoundedContext.html)

---

## Câu 4: Client Certificates là gì?

**Client certificate** là một loại chứng chỉ số (digital certificate), thường được dùng để cho phép hệ thống client tự xác thực danh tính của mình với server từ xa. Đây là thành phần quan trọng trong các thiết kế **mutual TLS (mTLS)** — mô hình xác thực hai chiều, nơi cả client và server đều phải chứng minh danh tính với nhau bằng chứng chỉ, thay vì chỉ server xác thực với client như trong HTTPS thông thường.

**Ví dụ**: Trong một hệ thống microservices nội bộ, `order-service` khi gọi tới `payment-service` phải đính kèm client certificate của mình. `payment-service` xác minh certificate này hợp lệ trước khi xử lý request — đảm bảo chỉ những service được cấp phép mới gọi được tới nó, ngay cả khi request đi trong cùng mạng nội bộ.

### Mở rộng

- Mutual TLS (mTLS) trong Service Mesh (ví dụ Istio) — cách tự động hóa việc cấp và xoay vòng (rotate) client certificate giữa các service.
- Public Key Infrastructure (PKI) — hạ tầng quản lý và phát hành chứng chỉ số.

### Nguồn trích dẫn

- Cloudflare — [What is mutual TLS (mTLS)?](https://www.cloudflare.com/learning/access-management/what-is-mutual-tls/)

---

## Câu 5: Giải thích về OAuth và OAuth 2.0?

**OAuth** ("Open" + "Auth", viết tắt của **Open Authorization**) là một giao thức **ủy quyền (authorization)** — cho phép một ứng dụng bên thứ ba được người dùng ủy quyền để truy cập tài nguyên của họ trên một dịch vụ khác, mà **không cần** ứng dụng đó biết mật khẩu của người dùng.

> **Ghi chú sửa lỗi kiến thức**: Bản gốc mô tả tên gọi OAuth gắn với cả "Authentication" (xác thực) lẫn "Authorization" (cấp quyền). Trên thực tế, **OAuth về bản chất là một giao thức Authorization, không phải Authentication** — nó chỉ trả lời câu hỏi "ứng dụng này có được phép truy cập tài nguyên X hay không", chứ không tự nó xác minh "người dùng này là ai". Việc xác thực danh tính người dùng (authentication) dựa trên nền OAuth 2.0 thuộc về một giao thức riêng gọi là **OpenID Connect (OIDC)**, xây dựng thêm một lớp trên OAuth 2.0.

**OAuth 2.0** là phiên bản kế thừa OAuth 1.0, nhưng cần lưu ý: OAuth 2.0 là một **thiết kế lại gần như hoàn toàn**, không tương thích ngược với OAuth 1.0 (khác chữ ký request, khác luồng cấp quyền), chứ không đơn thuần là một bản "nâng cấp" nhỏ. OAuth 2.0 cho phép các ứng dụng chia sẻ một phần tài nguyên với nhau mà không cần người dùng phải nhập lại username/password ở từng nơi, giúp giảm phiền toái khi phải quản lý quá nhiều tài khoản.

Trong OAuth 2.0, có 4 vai trò chính:

- **Resource Owner**: người dùng sở hữu tài nguyên, có quyền cấp phép truy cập.
- **Resource Server**: nơi lưu trữ tài nguyên, xử lý các request truy cập tài nguyên được bảo vệ.
- **Client**: ứng dụng bên thứ ba muốn truy cập tài nguyên thay mặt resource owner, và phải được ủy quyền trước khi truy cập.
- **Authorization Server**: xác thực thông tin người dùng gửi lên, rồi cấp quyền truy cập cho client bằng cách sinh ra access token. Đôi khi Authorization Server và Resource Server là cùng một hệ thống.

![Luồng OAuth 2.0](https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRPzh4OpeM1d0kA_4t4Sh8wB_1HebCYxstRfqsxWiKvPbg0W37Rv16ngVY&s=10)

**Ví dụ**: Khi bạn đăng nhập vào một ứng dụng bên thứ ba bằng nút "Đăng nhập với Google", ứng dụng đó (Client) được Google (Authorization Server) cấp một access token sau khi bạn (Resource Owner) đồng ý — token này cho phép ứng dụng đọc một phần thông tin hồ sơ Google của bạn (Resource Server), mà ứng dụng không bao giờ nhìn thấy mật khẩu Google của bạn.

### Mở rộng

- OpenID Connect (OIDC) — lớp authentication xây dựng trên nền OAuth 2.0.
- Các luồng cấp quyền (grant type) trong OAuth 2.0: Authorization Code, Client Credentials, Refresh Token.

### Nguồn trích dẫn

- OAuth.net — [OAuth 2.0](https://oauth.net/2/)
- Internet Engineering Task Force — [RFC 6749: The OAuth 2.0 Authorization Framework](https://datatracker.ietf.org/doc/html/rfc6749)

---

## Câu 6: Giải thích về CDC (Consumer-Driven Contract)?

**CDC (Consumer-Driven Contract)** là một mẫu kiểm thử (testing pattern), trong đó **bên tiêu thụ (consumer)** của một API định nghĩa trước những gì mình mong đợi ở API đó — bao gồm định dạng dữ liệu, các trường bắt buộc, mã trạng thái... Thỏa thuận này được gọi là **hợp đồng (contract)**. Bên cung cấp dịch vụ (provider) sau đó dùng chính hợp đồng này để kiểm tra xem thay đổi trong service của mình có làm hỏng những gì consumer đang mong đợi hay không, trước khi thay đổi đó được deploy.

Nói cách khác, CDC giúp phát hiện sớm các thay đổi API gây phá vỡ (breaking change) giữa các service, mà không cần phải chạy toàn bộ hệ thống end-to-end để kiểm tra.

**Ví dụ**: `order-service` (consumer) định nghĩa hợp đồng: "khi gọi `GET /products/{id}`, tôi mong đợi response luôn có trường `id`, `name`, và `price` kiểu số". Hợp đồng này được lưu lại và dùng để kiểm tra `product-service` (provider) mỗi khi team `product-service` thay đổi API — nếu ai đó lỡ đổi tên trường `price` thành `cost`, bài test dựa trên hợp đồng sẽ báo lỗi ngay, trước khi ảnh hưởng tới `order-service` thật sự.

### Mở rộng

- PACT — công cụ phổ biến nhất để triển khai Consumer-Driven Contract testing (xem câu tiếp theo).
- Contract testing vs Integration testing — CDC không thay thế hoàn toàn integration test, mà bổ sung một lớp kiểm tra nhanh hơn, ít phụ thuộc môi trường hơn.

### Nguồn trích dẫn

- Martin Fowler — [Consumer-Driven Contracts](https://martinfowler.com/articles/consumerDrivenContracts.html)

---

## Câu 7: PACT là gì?

![Cách PACT hoạt động](https://docs.pact.io/img/how-pact-works/summary.png)

**PACT** là một công cụ mã nguồn mở, dùng để triển khai kiểm thử theo mô hình **Consumer-Driven Contract** (xem câu hỏi trước): cho phép bên cung cấp dịch vụ (provider) và bên tiêu thụ (consumer) kiểm tra các tương tác giữa họ một cách tách biệt, dựa trên hợp đồng (contract) đã thống nhất — qua đó tăng độ tin cậy khi tích hợp giữa các microservice mà không cần dựng toàn bộ hệ thống để test. PACT hỗ trợ nhiều ngôn ngữ, ví dụ Ruby, Java, Scala, .NET, JavaScript, Swift/Objective-C, và cả Go.

**Ví dụ**: Đội phát triển `order-service` (consumer) viết test bằng PACT để tự động sinh ra file hợp đồng mô tả API họ cần từ `product-service`. File hợp đồng này được chia sẻ (thường qua Pact Broker) cho đội `product-service` (provider) để họ chạy lại test xác minh API của mình vẫn thỏa mãn đúng những gì `order-service` mong đợi.

### Mở rộng

- Pact Broker — nơi lưu trữ và chia sẻ hợp đồng giữa các team consumer/provider.
- Các công cụ contract testing khác: Spring Cloud Contract (hệ sinh thái Java/Spring).

### Nguồn trích dẫn

- Pact — [Tài liệu chính thức](https://docs.pact.io/)

---

## Câu 8: Semantic Monitoring là gì?

**Semantic monitoring** (còn gọi là giám sát tổng hợp — _synthetic monitoring_) sử dụng các bài kiểm tra tự động, chạy định kỳ, mô phỏng hành vi người dùng thật để phát hiện lỗi trong quy trình nghiệp vụ, thay vì chỉ xem các chỉ số hạ tầng đơn thuần (CPU, RAM...). Kỹ thuật này giúp có cái nhìn sâu hơn về hiệu suất giao dịch, tính khả dụng của dịch vụ, và hiệu suất tổng thể của ứng dụng — từ đó phát hiện sớm các vấn đề về hiệu suất của microservices, lỗi trong giao dịch, và duy trì mức hiệu suất tổng thể cao hơn.

**Ví dụ**: Một bài kiểm tra semantic monitoring có thể tự động thực hiện toàn bộ luồng "đăng nhập → thêm sản phẩm vào giỏ → thanh toán" mỗi 5 phút trên môi trường production, để phát hiện ngay khi một bước trong luồng nghiệp vụ này bị lỗi — dù các chỉ số hạ tầng (CPU, RAM của từng service) vẫn hoàn toàn bình thường.

### Mở rộng

- Synthetic monitoring vs Real User Monitoring (RUM) — hai cách tiếp cận bổ trợ nhau khi giám sát trải nghiệm người dùng.

### Nguồn trích dẫn

- Microsoft Learn — [Monitoring and diagnostics guidance](https://learn.microsoft.com/en-us/azure/architecture/best-practices/monitoring)

---

## Câu 9: Continuous Monitoring là gì?

> **Ghi chú sửa lỗi kiến thức**: Bản gốc định nghĩa Continuous Monitoring theo nghĩa **giám sát tuân thủ tài chính/rủi ro doanh nghiệp** — đây là định nghĩa đúng trong lĩnh vực quản trị doanh nghiệp/audit, nhưng **không phải ý nghĩa được hỏi tới trong ngữ cảnh microservices**. Trong kỹ thuật phần mềm, Continuous Monitoring mang nghĩa khác, được trình bày lại dưới đây.

**Continuous Monitoring (giám sát liên tục)**, trong ngữ cảnh microservices, là việc **liên tục thu thập và theo dõi** các chỉ số kỹ thuật của hệ thống đang chạy — như logs, metrics (CPU, RAM, latency, error rate), health check, và traces — nhằm phát hiện sớm sự cố, thay vì chỉ kiểm tra khi có báo lỗi từ người dùng. Đây là một phần quan trọng của **Observability** trong hệ thống microservices, vì với nhiều service chạy độc lập, việc phát hiện sớm dịch vụ nào đang gặp vấn đề là rất quan trọng để tránh ảnh hưởng dây chuyền.

**Ví dụ**: Một hệ thống dùng Prometheus để liên tục thu thập metric (số request/giây, tỉ lệ lỗi) từ mỗi microservice, kết hợp Grafana để hiển thị dashboard, và cấu hình cảnh báo (alert) tự động gửi tới Slack khi tỉ lệ lỗi của một service vượt ngưỡng 5% trong 5 phút liên tiếp — giúp đội vận hành phát hiện sự cố gần như ngay lập tức thay vì chờ người dùng báo lỗi.

### Mở rộng

- Ba trụ cột của Observability: Logs, Metrics, Traces.
- Prometheus, Grafana, OpenTelemetry — bộ công cụ phổ biến cho continuous monitoring trong hệ thống microservices.

### Nguồn trích dẫn

- Microsoft Learn — [Monitoring and diagnostics guidance](https://learn.microsoft.com/en-us/azure/architecture/best-practices/monitoring)

---

## Câu 10: Distributed Transactions là gì?

**Giao dịch phân tán (distributed transaction)** là một giao dịch mà các bước xử lý được thực hiện trải dài qua nhiều service/database khác nhau, cần hoàn thành theo trình tự để đảm bảo dữ liệu nhất quán trên toàn bộ các phần liên quan.

> **Ghi chú làm rõ**: Bản gốc mô tả distributed transaction là "một cách tiếp cận lỗi thời". Để chính xác hơn: **bản thân nhu cầu xử lý giao dịch trải dài qua nhiều service vẫn tồn tại và rất phổ biến** trong microservices (ví dụ: đặt hàng cần trừ kho, tạo thanh toán, gửi thông báo — tất cả phải nhất quán). Điều thực sự "lỗi thời" và ít được khuyến khích trong microservices hiện đại là dùng kỹ thuật **2PC (Two-Phase Commit)** truyền thống của giao dịch ACID phân tán để giải quyết bài toán này — vì 2PC yêu cầu khóa tài nguyên trên nhiều service cùng lúc, gây nghẽn và làm giảm khả năng mở rộng (scalability), đồng thời không phù hợp khi các service dùng database khác nhau và cần deploy độc lập.

Thay vì 2PC, microservices hiện đại thường giải quyết bài toán giao dịch trải dài qua nhiều service bằng **Saga Pattern** — chia giao dịch lớn thành một chuỗi các giao dịch cục bộ nhỏ hơn tại từng service, kèm theo bước "bù trừ" (compensating action) để hoàn tác nếu một bước ở giữa chuỗi thất bại, thay vì khóa toàn bộ tài nguyên chờ tất cả các bước hoàn tất.

![Vấn đề nhất quán trong giao dịch phân tán](https://images.ctfassets.net/00voh0j35590/2TOAHdVeW829evfKYX0GYB/d754b3487de811b2452838766b7e4380/distributed-transactions-and-the-consistency-problem.jpg)

**Ví dụ**: Khi đặt hàng, thay vì mở một giao dịch 2PC khóa cả `order-service`, `inventory-service`, `payment-service` cùng lúc, hệ thống dùng Saga: `order-service` tạo đơn hàng ở trạng thái "chờ", `inventory-service` trừ kho, `payment-service` xử lý thanh toán; nếu bước thanh toán thất bại, một hành động bù trừ được gọi để hoàn lại kho và hủy đơn hàng, thay vì rollback đồng thời như một transaction ACID truyền thống.

### Mở rộng

- Saga Pattern (Choreography-based vs Orchestration-based).
- Eventual Consistency — mô hình nhất quán dữ liệu thường được chấp nhận thay cho nhất quán tức thời (strong consistency) trong hệ thống phân tán.

### Nguồn trích dẫn

- Chris Richardson — [microservices.io — Pattern: Saga](https://microservices.io/patterns/data/saga.html)