# SQL và NoSQL: So sánh, ứng dụng và cách lựa chọn hệ quản trị phù hợp

## Mục lục

1. [SQL là gì?](#sql-là-gì)
2. [NoSQL là gì?](#nosql-là-gì)
3. [Điểm giống nhau giữa SQL và NoSQL](#điểm-giống-nhau-giữa-sql-và-nosql)
4. [Bảng so sánh chi tiết](#bảng-so-sánh-chi-tiết)
5. [ACID và BASE: vì sao SQL và NoSQL có triết lý khác nhau](#acid-và-base-vì-sao-sql-và-nosql-có-triết-lý-khác-nhau)
6. [Khi nào nên chọn SQL?](#khi-nào-nên-chọn-sql)
7. [Khi nào nên chọn NoSQL?](#khi-nào-nên-chọn-nosql)
8. [Ví dụ thực tế](#ví-dụ-thực-tế)
9. [Kết luận](#kết-luận)
10. [Mở rộng](#mở-rộng)

---

## SQL là gì?

**SQL (Structured Query Language)** là ngôn ngữ truy vấn có cấu trúc, dùng để thao tác với các hệ quản trị cơ sở dữ liệu quan hệ (Relational Database Management System - RDBMS) như Oracle, MySQL, SQL Server, PostgreSQL.

SQL không chỉ dùng để truy xuất, thêm, sửa, xóa dữ liệu (DML - Data Manipulation Language), mà còn cho phép:

- Định nghĩa và thay đổi cấu trúc cơ sở dữ liệu: tạo bảng, khóa chính, khóa ngoại, ràng buộc (DDL - Data Definition Language).
- Kiểm soát quyền truy cập dữ liệu: cấp quyền, thu hồi quyền cho từng người dùng hoặc vai trò (DCL - Data Control Language).
- Quản lý giao dịch (transaction): đảm bảo một nhóm thao tác được thực hiện trọn vẹn hoặc không thực hiện gì cả (TCL - Transaction Control Language).

Một hệ quản trị SQL thường bao gồm nhiều thành phần hỗ trợ khác nhau, ví dụ như Database Engine (xử lý truy vấn và lưu trữ), Integration Services (tích hợp và di chuyển dữ liệu), Analysis Services (phân tích dữ liệu đa chiều), Reporting Services (tạo báo cáo). Nhờ khả năng bảo mật và truy vấn mạnh mẽ, SQL được ứng dụng rộng rãi trong quản lý hệ thống, phân tích dữ liệu, và phát triển phần mềm nói chung.

## NoSQL là gì?

**NoSQL (Not Only SQL)** là thuật ngữ chỉ nhóm các hệ quản trị cơ sở dữ liệu không sử dụng mô hình quan hệ dạng bảng truyền thống. Thay vào đó, NoSQL lưu trữ dữ liệu theo nhiều mô hình khác nhau, phổ biến nhất là:

- **Key-Value**: lưu dữ liệu dưới dạng cặp khóa - giá trị (ví dụ: Redis, DynamoDB).
- **Document**: lưu dữ liệu dưới dạng tài liệu JSON/BSON linh hoạt (ví dụ: MongoDB, Couchbase).
- **Wide-Column**: lưu dữ liệu theo cột với schema linh hoạt cho từng dòng (ví dụ: Cassandra, HBase).
- **Graph**: lưu dữ liệu dưới dạng đỉnh (node) và cạnh (edge), phù hợp với dữ liệu có nhiều mối quan hệ phức tạp (ví dụ: Neo4j).

Nhờ không bị ràng buộc bởi một schema cố định, NoSQL cho phép mở rộng quy mô dễ dàng và quản lý đa dạng loại dữ liệu, kể cả dữ liệu phi cấu trúc hoặc bán cấu trúc. Bằng cách phân tán dữ liệu trên nhiều node/máy chủ (sharding, replication), NoSQL giúp hệ thống duy trì khả năng hoạt động liên tục ngay cả khi một phần hạ tầng gặp sự cố. Đây là lựa chọn phù hợp cho các hệ thống cần khả năng mở rộng nhanh và độ trễ thấp.

## Điểm giống nhau giữa SQL và NoSQL

Trước khi đi vào khác biệt, cần lưu ý cả hai loại đều:

- Cho phép lưu trữ, truy xuất, cập nhật và xóa dữ liệu (các thao tác CRUD).
- Hỗ trợ xây dựng chỉ mục (index) để tăng tốc độ truy vấn.
- Có thể triển khai ở quy mô lớn nếu được thiết kế và vận hành đúng cách (SQL mở rộng dọc là chủ yếu, nhưng các RDBMS hiện đại như PostgreSQL, CockroachDB cũng hỗ trợ mở rộng ngang ở mức độ nhất định).
- Đều cần được thiết kế schema/mô hình dữ liệu hợp lý ngay từ đầu để tránh nợ kỹ thuật (technical debt) về sau, dù NoSQL linh hoạt hơn về mặt này.

## Bảng so sánh chi tiết

| Tiêu chí | SQL | NoSQL |
|---|---|---|
| Ngôn ngữ truy vấn | Sử dụng ngôn ngữ SQL chuẩn hóa | Sử dụng API hoặc ngôn ngữ truy vấn riêng của từng hệ thống (ví dụ: MongoDB Query Language) |
| Mô hình dữ liệu | Dạng bảng (hàng và cột), có schema cố định | Linh hoạt: Document, Key-Value, Wide-Column, Graph |
| Khả năng mở rộng | Chủ yếu mở rộng theo chiều dọc (nâng cấp CPU, RAM của server) | Mở rộng theo chiều ngang (thêm node vào cluster) |
| Tính nhất quán | Tuân thủ ACID, ưu tiên tính toàn vẹn dữ liệu | Thường theo mô hình BASE, ưu tiên tính sẵn sàng, chấp nhận nhất quán cuối cùng (eventual consistency) |
| Quan hệ dữ liệu | Hỗ trợ tốt các mối quan hệ phức tạp qua khóa ngoại, JOIN | Hạn chế hoặc không hỗ trợ JOIN; dữ liệu liên quan thường được nhúng (embed) hoặc denormalize |
| Hiệu năng | Tối ưu cho truy vấn phức tạp, nhiều bảng liên kết | Tối ưu cho đọc/ghi nhanh với khối lượng dữ liệu lớn, cấu trúc đơn giản |
| Chi phí vận hành ban đầu | Cần thiết kế schema kỹ trước khi triển khai | Triển khai nhanh, schema có thể thay đổi linh hoạt theo thời gian |
| Ứng dụng phù hợp | Hệ thống tài chính, ngân hàng, ERP, CRM - nơi cần tính chính xác cao | Ứng dụng web/mobile quy mô lớn, IoT, hệ thống cần độ trễ thấp và khả năng mở rộng nhanh |
| Ví dụ | Oracle, MySQL, MariaDB, SQL Server, PostgreSQL | MongoDB, Couchbase, Cassandra, Redis, Neo4j |

## ACID và BASE: vì sao SQL và NoSQL có triết lý khác nhau

Đây là phần cốt lõi để hiểu bản chất khác biệt giữa hai loại cơ sở dữ liệu, thay vì chỉ học thuộc bảng so sánh.

**ACID** (áp dụng cho SQL):

- **Atomicity (Tính nguyên tử)**: một giao dịch hoặc thực hiện trọn vẹn, hoặc không thực hiện gì cả.
- **Consistency (Tính nhất quán)**: dữ liệu luôn ở trạng thái hợp lệ trước và sau giao dịch.
- **Isolation (Tính cô lập)**: các giao dịch chạy song song không ảnh hưởng lẫn nhau.
- **Durability (Tính bền vững)**: dữ liệu đã commit sẽ không bị mất, kể cả khi hệ thống gặp sự cố.

**BASE** (thường gặp ở NoSQL):

- **Basically Available**: hệ thống luôn phản hồi, ưu tiên tính sẵn sàng hơn là chờ dữ liệu nhất quán tuyệt đối.
- **Soft State**: trạng thái dữ liệu có thể thay đổi theo thời gian ngay cả khi không có input mới, do quá trình đồng bộ giữa các node.
- **Eventually Consistent**: dữ liệu sẽ nhất quán trên toàn hệ thống sau một khoảng thời gian, chứ không nhất thiết ngay lập tức.

Sự khác biệt này bắt nguồn từ định lý **CAP** (Consistency, Availability, Partition Tolerance): trong hệ thống phân tán, không thể đồng thời đảm bảo cả ba yếu tố. SQL truyền thống thường thiên về Consistency, còn nhiều hệ NoSQL thiên về Availability khi xảy ra phân vùng mạng (network partition). Đây là lý do vì sao việc chọn SQL hay NoSQL không chỉ là vấn đề công nghệ mà còn là bài toán đánh đổi (trade-off) dựa trên yêu cầu nghiệp vụ thực tế.

## Khi nào nên chọn SQL?

Nên ưu tiên SQL khi hệ thống cần:

- Tính toàn vẹn và chính xác dữ liệu tuyệt đối, ví dụ: hệ thống giao dịch tài chính, ngân hàng.
- Quản lý dữ liệu có mối quan hệ phức tạp giữa các thực thể (ví dụ: đơn hàng, khách hàng, sản phẩm liên kết chặt chẽ với nhau).
- Thực hiện các truy vấn phân tích, báo cáo phức tạp với nhiều điều kiện JOIN, GROUP BY.
- Đội ngũ phát triển đã quen thuộc với mô hình quan hệ và cần công cụ trưởng thành, hệ sinh thái ổn định.

## Khi nào nên chọn NoSQL?

Nên cân nhắc NoSQL khi hệ thống cần:

- Lưu trữ dữ liệu phi cấu trúc hoặc có cấu trúc thay đổi liên tục, ví dụ: dữ liệu cảm biến IoT, log hệ thống, nội dung do người dùng tạo.
- Khả năng mở rộng ngang nhanh chóng để đáp ứng lượng truy cập hoặc dữ liệu tăng đột biến (ví dụ: mạng xã hội, sàn thương mại điện tử vào mùa cao điểm).
- Độ trễ đọc/ghi thấp, ưu tiên tốc độ phản hồi hơn là tính nhất quán tức thời.
- Mô hình dữ liệu linh hoạt, thay đổi thường xuyên trong giai đoạn phát triển sản phẩm (đặc biệt phù hợp với các dự án MVP cần lặp nhanh).

## Ví dụ thực tế

Để dễ hình dung, dưới đây là một số tình huống áp dụng cụ thể:

- Một hệ thống ngân hàng cần đảm bảo số dư tài khoản chính xác tuyệt đối sau mỗi giao dịch chuyển tiền → nên chọn **SQL** (PostgreSQL, MySQL) để tận dụng ACID.
- Một nền tảng mạng xã hội cần lưu hàng tỷ bài đăng, bình luận, lượt thích với tốc độ ghi cực nhanh → nên chọn **NoSQL** dạng Document hoặc Wide-Column (MongoDB, Cassandra).
- Một hệ thống gợi ý sản phẩm dựa trên mối quan hệ giữa người dùng, ví dụ "bạn bè của bạn cũng mua sản phẩm này" → nên chọn **NoSQL** dạng Graph (Neo4j).
- Một hệ thống thương mại điện tử thực tế thường **kết hợp cả hai**: dùng SQL để quản lý đơn hàng, thanh toán (yêu cầu chính xác), và dùng NoSQL (Redis) để cache dữ liệu, quản lý giỏ hàng tạm thời (yêu cầu tốc độ).

## Kết luận

Không có một lựa chọn cơ sở dữ liệu nào phù hợp cho mọi bài toán. SQL phù hợp với các hệ thống cần tính chính xác, quan hệ dữ liệu phức tạp và giao dịch nghiêm ngặt. NoSQL phù hợp với các hệ thống cần mở rộng nhanh, xử lý khối lượng dữ liệu lớn với độ trễ thấp và mô hình dữ liệu linh hoạt.

Trong thực tế, nhiều hệ thống backend hiện đại áp dụng chiến lược **polyglot persistence** - sử dụng nhiều loại cơ sở dữ liệu khác nhau cho từng thành phần trong cùng một hệ thống, tùy theo đặc điểm dữ liệu và yêu cầu của từng module. Việc lựa chọn đúng công nghệ cần dựa trên yêu cầu nghiệp vụ cụ thể, chứ không nên chọn theo xu hướng hoặc thói quen.

### Mở rộng

Một số hướng tìm hiểu thêm để nâng cao kiến thức về chủ đề này:

- **NewSQL**: nhóm cơ sở dữ liệu như CockroachDB, Google Spanner, TiDB - kết hợp khả năng mở rộng ngang của NoSQL với tính ACID của SQL truyền thống.
- **Định lý CAP** và **định lý PACELC**: tìm hiểu sâu hơn về các đánh đổi trong thiết kế hệ thống phân tán.
- **Chiến lược Sharding và Replication**: cách các hệ NoSQL phân tán dữ liệu trên nhiều node để đảm bảo hiệu năng và tính sẵn sàng.
- **Denormalization trong thiết kế NoSQL**: kỹ thuật thiết kế dữ liệu trùng lặp có chủ đích để tối ưu tốc độ đọc, khác với nguyên tắc chuẩn hóa (normalization) trong SQL.
- **Polyglot Persistence**: mô hình kiến trúc sử dụng nhiều loại database khác nhau trong cùng một hệ thống microservices.
- **So sánh cụ thể PostgreSQL vs MongoDB vs Redis** trong bối cảnh xây dựng backend bằng Golang, đặc biệt liên quan đến các thư viện như `sqlc`, `pgx`, `go-redis`.