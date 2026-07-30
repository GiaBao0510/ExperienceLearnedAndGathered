# PostgreSQL và MySQL: Khác biệt, hiệu năng thực tế và cách lựa chọn

MySQL và PostgreSQL là hai hệ quản trị cơ sở dữ liệu quan hệ (RDBMS) mã nguồn mở phổ biến nhất hiện nay. Tài liệu này không chỉ so sánh về mặt tính năng mà còn đi sâu vào kiến trúc xử lý bên trong, giúp hiểu rõ **vì sao** hai hệ thống có sự khác biệt về hiệu năng trong từng loại truy vấn cụ thể - đây là kiến thức quan trọng khi thiết kế backend, không chỉ đơn thuần là chọn theo "cảm giác quen thuộc".

![So sánh MySQL và PostgreSQL](https://viettelidc.com.vn/uploadimage/Root/root/mysql-postgresql-1.png)

## Mục lục

1. [Điểm giống nhau giữa MySQL và PostgreSQL](#điểm-giống-nhau-giữa-mysql-và-postgresql)
2. [Bảng so sánh khái niệm](#bảng-so-sánh-khái-niệm)
3. [Kiến trúc xử lý kết nối](#kiến-trúc-xử-lý-kết-nối)
4. [Cách xử lý Point Query và Range Query](#cách-xử-lý-point-query-và-range-query)
5. [Benchmark hiệu năng thực tế](#benchmark-hiệu-năng-thực-tế)
6. [Khi nào nên chọn MySQL? Khi nào nên chọn PostgreSQL?](#khi-nào-nên-chọn-mysql-khi-nào-nên-chọn-postgresql)
7. [Kết luận](#kết-luận)
8. [Mở rộng](#mở-rộng)

---

## Điểm giống nhau giữa MySQL và PostgreSQL

Vì đều là hệ quản trị cơ sở dữ liệu quan hệ mã nguồn mở, MySQL và PostgreSQL có nhiều điểm chung quan trọng:

- **RDBMS**: cả hai đều lưu trữ dữ liệu có cấu trúc dạng bảng, hỗ trợ truy vấn phức tạp và duy trì mối quan hệ giữa các bảng thông qua khóa ngoại.
- **Tuân thủ chuẩn SQL**: cung cấp cú pháp truy vấn nhất quán theo chuẩn SQL, giúp dễ dàng chuyển đổi kiến thức giữa hai hệ thống.
- **Hỗ trợ ACID**: đều đảm bảo tính nguyên tử, nhất quán, cô lập và bền vững cho giao dịch (mức độ tuân thủ đầy đủ có khác nhau tùy storage engine - xem bảng so sánh bên dưới).
- **Mã nguồn mở, miễn phí**: có thể tùy chỉnh, triển khai tự do và không phụ thuộc vào một nhà cung cấp duy nhất.
- **Hỗ trợ JSON**: cả hai đều có kiểu dữ liệu JSON, phù hợp khi cần lưu trữ dữ liệu bán cấu trúc trong hệ thống quan hệ.
- **Đa nền tảng**: hoạt động tốt trên Windows, Linux, macOS và các hệ thống Unix khác.
- **Bảo mật**: hỗ trợ mã hóa kết nối qua SSL/TLS, phân quyền người dùng chi tiết theo vai trò (role-based access control); việc bổ sung xác thực hai yếu tố thường được thực hiện thông qua công cụ hoặc middleware bên ngoài chứ không phải tính năng gốc của database.
- **Sao lưu và phục hồi**: đều có công cụ backup/restore trưởng thành (`mysqldump`, `mysqlpump` với MySQL; `pg_dump`, `pg_basebackup` với PostgreSQL).
- **Cộng đồng lớn**: tài liệu phong phú, hệ sinh thái công cụ và driver hỗ trợ đầy đủ cho hầu hết ngôn ngữ lập trình, bao gồm Golang (`database/sql` + driver `go-sql-driver/mysql` hoặc `pgx`/`lib/pq`).

## Bảng so sánh khái niệm

| Đặc điểm | MySQL | PostgreSQL |
|---|---|---|
| Tuân thủ ACID | Tuân thủ đầy đủ khi dùng storage engine InnoDB (mặc định từ MySQL 5.5); engine cũ MyISAM không hỗ trợ transaction | Tuân thủ ACID đầy đủ trong mọi cấu hình mặc định |
| Kiểm soát đồng thời (MVCC) | InnoDB **có hỗ trợ MVCC** và row-level locking; chỉ engine MyISAM (đã lỗi thời) mới dùng table-level locking và không có MVCC | Hỗ trợ MVCC ở toàn bộ hệ thống, cho phép đọc và ghi đồng thời mà không chặn lẫn nhau |
| Chỉ mục | B-tree (mặc định) và R-tree/spatial index | B-tree, Hash, GiST, SP-GiST, GIN, BRIN - đa dạng hơn, phù hợp với nhiều loại truy vấn đặc thù (full-text search, dữ liệu địa lý, mảng...) |
| Kiểu dữ liệu | Cơ sở dữ liệu quan hệ thuần túy, không hỗ trợ kiểu dữ liệu đối tượng phức tạp | Cơ sở dữ liệu quan hệ - đối tượng (object-relational), hỗ trợ mảng (array), kiểu dữ liệu tùy biến (custom type), XML, JSON/JSONB |
| Chế độ xem (View) | Hỗ trợ view cơ bản | Hỗ trợ view nâng cao, bao gồm materialized view giúp cải thiện hiệu năng cho truy vấn phức tạp lặp lại |
| Stored Procedure | Viết bằng SQL (theo chuẩn SQL/PSM) | Có thể viết bằng SQL hoặc nhiều ngôn ngữ khác (PL/pgSQL, PL/Python, PL/Perl...) |
| Trigger | Hỗ trợ BEFORE và AFTER cho INSERT, UPDATE, DELETE | Hỗ trợ thêm INSTEAD OF (chủ yếu dùng cho view), linh hoạt hơn khi cần can thiệp logic phức tạp |

## Kiến trúc xử lý kết nối

Khác biệt về kiến trúc xử lý kết nối là một trong những nguyên nhân cốt lõi dẫn đến khác biệt hiệu năng giữa hai hệ thống.

![Kiến trúc xử lý kết nối của PostgreSQL](https://statics.cdn.200lab.io/2024/10/postgresql-vs-mysql-pg-arch-diagram.png?width=800)

**PostgreSQL - Process-Based Model**: mỗi kết nối mới sẽ tạo ra một process (tiến trình) riêng biệt ở tầng hệ điều hành. Nếu một kết nối gặp sự cố (ví dụ: lỗi truy cập bộ nhớ không hợp lệ), chỉ tiến trình đó bị ảnh hưởng, các kết nối khác vẫn hoạt động bình thường - đây là ưu điểm về khả năng cô lập (isolation) và độ ổn định. Tuy nhiên, mô hình process-per-connection khá tốn tài nguyên: với 1.000 kết nối đồng thời, hệ thống sẽ tạo ra 1.000 tiến trình, tiêu tốn đáng kể RAM và CPU. Vì vậy, trong môi trường production, PostgreSQL gần như bắt buộc phải dùng kèm connection pooler như **PgBouncer** hoặc **pgcat** để giới hạn và tái sử dụng kết nối hiệu quả.

**MySQL - Thread-Based Model**: MySQL sử dụng một process duy nhất, xử lý mỗi kết nối bằng một thread riêng (thread-per-connection). Do thread nhẹ hơn process rất nhiều về mặt tài nguyên hệ điều hành, MySQL có thể xử lý số lượng kết nối đồng thời lớn với chi phí bộ nhớ thấp hơn PostgreSQL, phù hợp tốt với các ứng dụng thiên về đọc dữ liệu (read-heavy) có lượng truy cập lớn.

**Cách tổ chức dữ liệu vật lý** cũng khác nhau đáng kể:

- **PostgreSQL - Heap-Organized Table**: dữ liệu được lưu theo thứ tự thêm vào (insertion order), không sắp xếp theo primary key hay bất kỳ cột nào. Không có mối liên kết vật lý giữa vị trí lưu trữ và giá trị index.
- **MySQL (InnoDB) - Clustered Index**: dữ liệu của bảng được lưu trữ vật lý ngay trong cấu trúc B-tree của primary key. Nói cách khác, primary key index và dữ liệu thực tế là một, giúp truy vấn theo primary key rất nhanh vì không cần thêm bước tra cứu.

## Cách xử lý Point Query và Range Query

Hai loại truy vấn phổ biến nhất khi đánh giá hiệu năng đọc:

- **Point query**: trả về một hoặc một nhóm nhỏ bản ghi, điều kiện thường là khóa chính hoặc cột đã đánh index. Ví dụ: `SELECT * FROM khach_hang WHERE id = 2`.
- **Range query**: trả về nhiều bản ghi nằm trong một khoảng giá trị, thường dùng trên cột đã có secondary index. Ví dụ: `SELECT * FROM khach_hang WHERE tuoi BETWEEN 25 AND 40`.

**Với point query:**

- *PostgreSQL*: tìm vị trí row thông qua primary key index (index chỉ lưu vị trí, không lưu dữ liệu thực tế), sau đó truy cập vào bảng heap để lấy dữ liệu đầy đủ → cần **2 bước tra cứu**.
- *MySQL (InnoDB)*: vì dùng clustered index, dữ liệu đã nằm sẵn cùng với chỉ mục primary key → chỉ cần **1 bước tra cứu**, trả kết quả ngay lập tức.

→ Đây là lý do vì sao MySQL thường nhanh hơn PostgreSQL trong các point query đơn giản.

**Với range query:**

- *MySQL (InnoDB)*: tìm kiếm B-tree trên secondary index theo cột điều kiện (ví dụ `order_date`) để lấy `order_id`, sau đó tra cứu tiếp trong clustered index bằng `order_id` để lấy đầy đủ dữ liệu → cần **2 lần tra cứu** (secondary index + clustered index).
- *PostgreSQL*: secondary index lưu trực tiếp giá trị cột được đánh index kèm theo TID (Tuple Identifier - con trỏ đến vị trí thực tế trong heap). Vì vậy chỉ cần tìm trên secondary index để lấy TID, rồi truy cập thẳng vào heap → chỉ cần **1 lần tra cứu chính**, không phải đi qua primary key trung gian.

→ Đây là lý do vì sao PostgreSQL thường có lợi thế hơn trong range query và các truy vấn ghi dữ liệu lớn.

## Benchmark hiệu năng thực tế

Kết quả benchmark dưới đây minh họa cho các phân tích kiến trúc ở trên, được thực hiện trên 3 kịch bản dữ liệu khác nhau:

- **Cached**: toàn bộ dữ liệu nằm vừa trong bộ nhớ đệm (RAM).
- **Less IO-bound**: dữ liệu lớn hơn dung lượng RAM, khoảng 64 triệu dòng mỗi bảng.
- **More IO-bound**: dữ liệu vượt xa dung lượng RAM, khoảng 200 triệu dòng mỗi bảng.

| Kịch bản | PostgreSQL | MySQL |
|---|---|---|
| Cached | Nhanh hơn ở tác vụ ghi dữ liệu lớn (write-heavy) và range query; point query tương đương MySQL | Point query tương đương PostgreSQL |
| Less IO-bound | Nhanh hơn ở write-heavy và range query | Nhanh hơn ở point query - MySQL chỉ cần khoảng 0.1 lần đọc đĩa mỗi truy vấn nhờ clustered index, trong khi PostgreSQL cần khoảng 3 lần đọc do phải qua cả secondary index lẫn heap |
| More IO-bound | Nhanh hơn ở write-heavy và range query, dù phải thực hiện số lần đọc I/O gấp khoảng 2 lần MySQL cho mỗi point query | Nhanh hơn ở point query, nhưng cần nhiều thao tác I/O hơn cho mỗi range query |

Có thể tóm gọn: **MySQL (InnoDB) mạnh về point query nhờ clustered index**, còn **PostgreSQL mạnh về range query và ghi dữ liệu lớn nhờ MVCC và cấu trúc heap không phụ thuộc thứ tự index**. Khoảng cách hiệu năng này càng rõ rệt khi dữ liệu vượt quá dung lượng RAM (IO-bound).

## Khi nào nên chọn MySQL? Khi nào nên chọn PostgreSQL?

**Nên chọn MySQL khi:**

- Ứng dụng thiên về đọc dữ liệu (read-heavy) với lượng truy cập lớn, ví dụ: trang blog, tin tức, hệ thống hiển thị nội dung.
- Truy vấn chủ yếu là point query đơn giản theo khóa chính hoặc index.
- Đội ngũ mới bắt đầu, cần thiết lập nhanh, ít cấu hình phức tạp ban đầu.
- Hệ thống cần xử lý số lượng kết nối đồng thời lớn với chi phí tài nguyên thấp (nhờ mô hình thread-based).

**Nên chọn PostgreSQL khi:**

- Ứng dụng cần xử lý ghi dữ liệu thường xuyên (write-heavy) với nhiều thao tác đồng thời.
- Cần thực hiện range query, aggregation phức tạp, hoặc truy vấn đòi hỏi tính toàn vẹn dữ liệu cao.
- Dữ liệu có cấu trúc phức tạp: cần kiểu dữ liệu mảng, JSON/JSONB nâng cao, custom type, hoặc full-text search với nhiều loại index (GIN, GiST).
- Đội ngũ đã có kinh nghiệm vận hành, sẵn sàng đầu tư cấu hình hạ tầng (connection pooler, tuning) để khai thác tối đa hiệu năng.

## Kết luận

MySQL và PostgreSQL đều là những hệ quản trị cơ sở dữ liệu quan hệ trưởng thành, nhưng có triết lý thiết kế khác nhau ở tầng kiến trúc: MySQL (InnoDB) tối ưu cho các truy vấn đọc đơn giản nhờ clustered index và mô hình thread nhẹ, trong khi PostgreSQL tối ưu cho các hệ thống cần ghi dữ liệu nhiều, truy vấn phức tạp và tính toàn vẹn cao nhờ MVCC toàn diện và hệ thống index đa dạng. Việc lựa chọn đúng công nghệ nên dựa trên đặc điểm tải truy vấn thực tế (read-heavy hay write-heavy) của hệ thống, thay vì chỉ dựa vào sự quen thuộc hoặc xu hướng.

### Mở rộng

Một số hướng tìm hiểu thêm để nâng cao kiến thức về chủ đề này:

- **Storage Engine trong MySQL**: tìm hiểu sâu hơn về InnoDB so với các engine khác như MyISAM, Memory, để hiểu rõ vì sao lựa chọn storage engine ảnh hưởng trực tiếp đến tính ACID và MVCC.
- **Connection Pooling**: cách hoạt động của PgBouncer, pgcat (cho PostgreSQL) và connection pool tích hợp sẵn trong driver Golang (`database/sql` với `SetMaxOpenConns`, `SetMaxIdleConns`).
- **EXPLAIN ANALYZE**: công cụ phân tích execution plan của cả MySQL và PostgreSQL, giúp xác định query có đang tận dụng index hiệu quả hay không.
- **Vacuum trong PostgreSQL**: cơ chế dọn dẹp các phiên bản dữ liệu cũ (dead tuple) sinh ra từ MVCC, và vì sao việc cấu hình `autovacuum` đúng cách ảnh hưởng lớn đến hiệu năng lâu dài.
- **Read Replica và Horizontal Scaling**: cách cả hai hệ thống mở rộng khả năng đọc thông qua replica, và giới hạn của việc mở rộng ghi trong kiến trúc RDBMS truyền thống (tiền đề dẫn đến NewSQL như CockroachDB, Google Spanner).
- **Thực hành với Golang**: so sánh trải nghiệm dùng `pgx` (driver hiệu năng cao cho PostgreSQL, hỗ trợ tốt kiểu dữ liệu nâng cao) với `go-sql-driver/mysql`, đặc biệt trong bối cảnh dùng `sqlc` để sinh code Go từ SQL query.