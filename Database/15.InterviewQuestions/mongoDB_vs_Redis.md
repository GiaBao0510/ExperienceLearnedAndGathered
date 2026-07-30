# So sánh Redis và MongoDB, cách lựa chọn giải pháp phù hợp

## Mục lục

1. [Tổng quan](#tổng-quan)
2. [Redis là gì?](#redis-là-gì)
3. [MongoDB là gì?](#mongodb-là-gì)
4. [Điểm giống nhau](#điểm-giống-nhau)
5. [Cơ chế hoạt động: khác biệt cốt lõi](#cơ-chế-hoạt-động-khác-biệt-cốt-lõi)
6. [Bảng so sánh chi tiết](#bảng-so-sánh-chi-tiết)
7. [Khi nào nên chọn Redis?](#khi-nào-nên-chọn-redis)
8. [Khi nào nên chọn MongoDB?](#khi-nào-nên-chọn-mongodb)
9. [Kết hợp Redis và MongoDB trong cùng hệ thống](#kết-hợp-redis-và-mongodb-trong-cùng-hệ-thống)
10. [Kết luận](#kết-luận)
11. [Mở rộng](#mở-rộng)

---

## Tổng quan

Trong các hệ thống backend hiện đại, việc lựa chọn cơ sở dữ liệu phù hợp ảnh hưởng trực tiếp đến hiệu suất, khả năng mở rộng và trải nghiệm người dùng. Redis và MongoDB đều là các giải pháp NoSQL được sử dụng rộng rãi, nhưng ra đời để giải quyết hai bài toán khác nhau: Redis tối ưu cho **tốc độ**, còn MongoDB tối ưu cho **lưu trữ linh hoạt và bền vững**. Vì vậy, thay vì đặt câu hỏi "Redis hay MongoDB tốt hơn?", câu hỏi đúng nên là "Redis và MongoDB nên đóng vai trò gì trong hệ thống của tôi?".

## Redis là gì?

Redis (Remote Dictionary Server) là cơ sở dữ liệu NoSQL mã nguồn mở, hoạt động chủ yếu trên bộ nhớ RAM (In-Memory Database). Do dữ liệu được xử lý trực tiếp trong RAM thay vì đọc/ghi trên ổ đĩa, Redis đạt tốc độ phản hồi ở mức micro giây đến mili giây.

Redis không chỉ hỗ trợ mô hình Key-Value đơn thuần mà còn cung cấp nhiều kiểu dữ liệu chuyên biệt:

- **String**: lưu giá trị đơn giản, dùng cho cache, đếm số lượt truy cập.
- **Hash**: lưu object có nhiều trường, ví dụ thông tin session người dùng.
- **List**: hàng đợi (queue), lịch sử hoạt động gần nhất.
- **Set / Sorted Set**: bảng xếp hạng (leaderboard), loại bỏ trùng lặp.
- **Bitmap, HyperLogLog**: thống kê số lượng lớn với chi phí bộ nhớ tối thiểu.
- **Stream**: xử lý luồng sự kiện, tương tự message queue nhẹ.

Mặc dù ưu tiên xử lý trên RAM, Redis vẫn hỗ trợ cơ chế ghi xuống ổ đĩa để giảm rủi ro mất dữ liệu khi hệ thống gặp sự cố:

- **RDB (Snapshot)**: chụp lại toàn bộ dữ liệu tại một thời điểm, phù hợp để backup định kỳ nhưng có thể mất một phần dữ liệu giữa hai lần snapshot.
- **AOF (Append Only File)**: ghi lại từng lệnh ghi vào file log, an toàn hơn RDB nhưng tốn dung lượng và ảnh hưởng nhẹ đến hiệu năng.

Nhờ tốc độ vượt trội, Redis thường được dùng làm: bộ nhớ đệm (Cache), lưu Session đăng nhập, hàng đợi thông điệp (Message Queue), bảng xếp hạng thời gian thực, và giới hạn tần suất truy cập (Rate Limiting) - đây cũng là ứng dụng phổ biến trong các hệ thống Golang backend sử dụng thư viện `go-redis`.

## MongoDB là gì?

MongoDB là cơ sở dữ liệu NoSQL hướng tài liệu (Document Database), lưu trữ dữ liệu dưới dạng Document theo định dạng BSON (Binary JSON). Thay vì tổ chức dữ liệu theo bảng - hàng - cột như cơ sở dữ liệu quan hệ, MongoDB nhóm các Document vào trong Collection.

Điểm nổi bật của MongoDB là **schema linh hoạt**: các Document trong cùng một Collection có thể chứa các trường dữ liệu khác nhau mà không cần thay đổi cấu trúc toàn bộ cơ sở dữ liệu. Điều này giúp MongoDB phù hợp với các ứng dụng có yêu cầu dữ liệu thay đổi thường xuyên trong quá trình phát triển.

MongoDB cung cấp các tính năng của một cơ sở dữ liệu trưởng thành, gồm:

- **Index**: tăng tốc truy vấn, hỗ trợ nhiều loại index (single field, compound, text, geospatial).
- **Aggregation Pipeline**: xử lý, biến đổi, thống kê dữ liệu qua nhiều bước xử lý (tương tự GROUP BY nâng cao trong SQL).
- **Replication (Replica Set)**: nhân bản dữ liệu qua nhiều node để đảm bảo tính sẵn sàng khi một node gặp sự cố.
- **Sharding**: phân mảnh dữ liệu theo chiều ngang để mở rộng quy mô lưu trữ và xử lý.
- **Transaction**: hỗ trợ giao dịch đa tài liệu (multi-document ACID transaction) kể từ phiên bản 4.0, giúp MongoDB đáp ứng được một số bài toán trước đây chỉ SQL mới xử lý tốt.

Nhờ đó, MongoDB được dùng phổ biến trong thương mại điện tử, mạng xã hội, hệ thống quản lý nội dung (CMS), ứng dụng IoT và các dịch vụ triển khai trên Cloud.

## Điểm giống nhau

Trước khi đi vào khác biệt, cả Redis và MongoDB đều:

- Thuộc nhóm cơ sở dữ liệu NoSQL, không dùng ngôn ngữ SQL để truy vấn.
- Không yêu cầu schema cố định như cơ sở dữ liệu quan hệ.
- Hỗ trợ mở rộng theo chiều ngang (horizontal scaling) thông qua cluster/sharding.
- Được thiết kế để hoạt động tốt trong kiến trúc phân tán và Microservices.
- Có driver hỗ trợ tốt cho Golang (`go-redis` cho Redis, `mongo-driver` chính thức của MongoDB).

## Cơ chế hoạt động: khác biệt cốt lõi

Sự khác biệt về cơ chế lưu trữ chính là gốc rễ dẫn đến khác biệt về hiệu năng và phạm vi ứng dụng của hai hệ thống:

- **Redis** lưu dữ liệu trực tiếp trong RAM theo mô hình Key-Value. Khi có yêu cầu đọc/ghi, Redis xử lý ngay trên bộ nhớ mà không cần truy cập ổ đĩa, đồng thời Redis là **đơn luồng (single-threaded)** đối với việc xử lý lệnh, giúp tránh race condition mà không cần cơ chế khóa phức tạp - đánh đổi lại là Redis không tận dụng được nhiều nhân CPU cho một lệnh đơn lẻ.
- **MongoDB** lưu dữ liệu dưới dạng Document trên ổ đĩa (có cơ chế cache dữ liệu thường dùng trong RAM để tăng tốc). Khi nhận truy vấn, MongoDB dùng Index để định vị Document phù hợp, sau đó xử lý qua Aggregation Pipeline nếu cần, rồi trả kết quả về. Cơ chế này cho phép truy vấn phức tạp nhưng có độ trễ cao hơn Redis.

Vì hướng tới hai mục tiêu khác nhau - Redis ưu tiên **tốc độ**, MongoDB ưu tiên **khả năng lưu trữ và truy vấn linh hoạt** - nên hai hệ thống này không cạnh tranh trực tiếp mà thường được triển khai bổ trợ cho nhau.

## Bảng so sánh chi tiết

| Tiêu chí | Redis | MongoDB |
|---|---|---|
| Loại cơ sở dữ liệu | NoSQL Key-Value, In-Memory Database | NoSQL hướng tài liệu (Document Database) |
| Nơi lưu trữ chính | RAM (có tùy chọn ghi xuống ổ đĩa) | Ổ đĩa (có cache trong RAM) |
| Định dạng dữ liệu | String, Hash, List, Set, Sorted Set, Bitmap, Stream | Document dạng BSON (Binary JSON) |
| Tốc độ đọc/ghi | Rất cao, độ trễ micro/mili giây | Cao, nhưng thấp hơn Redis do có thao tác I/O với ổ đĩa |
| Khả năng truy vấn | Đơn giản, chủ yếu theo Key | Phong phú: lọc, sắp xếp, phân trang, Aggregation Pipeline |
| Tính bền vững dữ liệu | Có hỗ trợ (RDB/AOF) nhưng không phải mục đích chính | Được thiết kế để lưu trữ dữ liệu bền vững, lâu dài |
| Transaction | Hỗ trợ giới hạn (MULTI/EXEC, không rollback giữa chừng) | Hỗ trợ multi-document ACID transaction từ bản 4.0 |
| Mô hình xử lý | Đơn luồng (single-threaded) cho việc thực thi lệnh | Đa luồng, hỗ trợ song song hóa tốt hơn cho truy vấn phức tạp |
| Chi phí lưu trữ | Cao hơn trên mỗi GB do dùng RAM | Thấp hơn, phù hợp lưu trữ dữ liệu khối lượng lớn |
| Ứng dụng phổ biến | Cache, Session, Message Queue, Leaderboard, Rate Limiting | Quản lý người dùng, đơn hàng, nội dung, hệ thống CMS |

## Khi nào nên chọn Redis?

Nên ưu tiên Redis khi hệ thống cần:

- Tốc độ phản hồi cực nhanh, độ trễ thấp cho các thao tác đọc/ghi lặp lại liên tục.
- Lưu dữ liệu tạm thời, có thời gian sống ngắn (TTL) như session, OTP, token.
- Xây dựng hàng đợi nhẹ (message queue), hệ thống pub/sub thời gian thực.
- Giới hạn tần suất truy cập API (rate limiting), đếm lượt truy cập, bảng xếp hạng.

## Khi nào nên chọn MongoDB?

Nên ưu tiên MongoDB khi hệ thống cần:

- Lưu trữ dữ liệu lâu dài, làm nguồn dữ liệu chính (source of truth) của hệ thống.
- Cấu trúc dữ liệu thay đổi thường xuyên hoặc chưa cố định ngay từ đầu (giai đoạn phát triển sản phẩm nhanh - MVP).
- Truy vấn phức tạp với nhiều điều kiện lọc, sắp xếp, thống kê qua Aggregation Pipeline.
- Dữ liệu có cấu trúc lồng nhau tự nhiên (nested object), ví dụ: một đơn hàng chứa danh sách sản phẩm, thông tin giao hàng, lịch sử trạng thái.

## Kết hợp Redis và MongoDB trong cùng hệ thống

Trong thực tế, phần lớn hệ thống backend không chọn một trong hai mà **kết hợp cả hai** theo mô hình phổ biến gọi là **Cache-Aside Pattern**:

1. Khi ứng dụng cần đọc dữ liệu, trước tiên kiểm tra trong Redis (cache).
2. Nếu có dữ liệu (cache hit) → trả kết quả ngay, không cần truy cập MongoDB.
3. Nếu không có (cache miss) → truy vấn MongoDB, sau đó lưu kết quả vào Redis kèm thời gian sống (TTL) để phục vụ các lần đọc tiếp theo.
4. Khi dữ liệu trong MongoDB thay đổi, hệ thống cần xóa hoặc cập nhật lại cache tương ứng trong Redis (cache invalidation) để tránh dữ liệu cũ.

Mô hình này giúp giảm tải cho MongoDB, tăng tốc độ phản hồi cho các dữ liệu được truy cập thường xuyên, đồng thời vẫn đảm bảo MongoDB đóng vai trò lưu trữ dữ liệu bền vững, chính xác.

## Kết luận

Redis và MongoDB đều là các cơ sở dữ liệu NoSQL mạnh mẽ, nhưng phục vụ hai mục tiêu khác nhau: Redis nổi bật với tốc độ xử lý cực cao nhờ hoạt động trên RAM, còn MongoDB phù hợp cho việc lưu trữ dữ liệu lâu dài và truy vấn linh hoạt. Việc hiểu rõ bản chất và điểm mạnh của từng công nghệ sẽ giúp lựa chọn đúng công cụ cho đúng bài toán, hoặc kết hợp cả hai để xây dựng hệ thống có hiệu năng cao, ổn định và dễ mở rộng.

### Mở rộng

Một số hướng tìm hiểu thêm để nâng cao kiến thức về chủ đề này:

- **Cache Invalidation Strategies**: các chiến lược làm mới cache (TTL, Write-Through, Write-Behind, Cache-Aside) và bài toán kinh điển "there are only two hard things in computer science: cache invalidation and naming things".
- **Redis Cluster và Redis Sentinel**: cơ chế mở rộng và đảm bảo tính sẵn sàng cao (high availability) cho Redis trong môi trường production.
- **MongoDB Replica Set và Read Preference**: cách MongoDB đảm bảo tính sẵn sàng và các mức độ nhất quán khi đọc dữ liệu từ node phụ (secondary).
- **So sánh Redis với Memcached**: một hệ thống in-memory cache khác, để hiểu rõ hơn vì sao Redis được ưa chuộng hơn nhờ hỗ trợ đa dạng kiểu dữ liệu.
- **Thiết kế schema trong MongoDB**: nguyên tắc Embedding vs Referencing khi mô hình hóa dữ liệu có quan hệ.
- **Tích hợp Redis và MongoDB trong Golang**: thực hành xây dựng Cache-Aside Pattern bằng `go-redis` và `mongo-driver`, kết hợp với context timeout và xử lý lỗi khi cache miss hoặc kết nối gián đoạn.