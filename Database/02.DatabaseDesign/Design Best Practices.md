## Design Best Practices (Thực hành thiết kế tốt)

Thiết kế cơ sở dữ liệu (Database Design) vững chắc cần đảm bảo tính nhất quán, tối ưu hiệu năng, và dễ dàng mở rộng. Các quy tắc cốt lõi bao gồm: chuẩn hóa dữ liệu để tránh trùng lặp, chọn kiểu dữ liệu chính xác, đặt tên nhất quán, và lập chỉ mục (indexing) đúng chỗ.

**a) Luôn có Primary Key surrogate, kể cả khi có Natural Key**

Surrogate key là khóa chính nhân tạo (thường là `UUID` hoặc `BIGINT AUTO-INCREMENT`), không mang ý nghĩa nghiệp vụ — khác với natural key (khóa tự nhiên) như số CMND/CCCD, mã số thuế, hay email, vốn là dữ liệu có ý nghĩa thực tế. Nên luôn dùng surrogate key làm PK, **ngay cả khi** bảng đã có một cột tưởng chừng có thể làm natural key duy nhất. Lý do: giá trị nghiệp vụ có thể thay đổi theo thời gian (số CCCD có thể cấp lại, email có thể đổi), trong khi PK **không bao giờ được phép thay đổi** — nếu PK thay đổi, mọi bảng có khóa ngoại tham chiếu đến nó đều phải cập nhật theo, cực kỳ rủi ro và tốn kém ở hệ thống lớn.

**b) Luôn có created_at, updated_at (audit columns)**

```sql
created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
```

Giúp truy vết dữ liệu, debug, và phục vụ các query kiểu "lấy dữ liệu thay đổi gần đây" (đồng bộ, cache invalidation...).

> `TIMESTAMP` (viết tắt của `TIMESTAMP WITHOUT TIME ZONE`). Đây là một trong những lỗi phổ biến và nguy hiểm nhất khi thiết kế schema PostgreSQL: `TIMESTAMP` lưu giờ theo "giờ địa phương" không kèm thông tin múi giờ, dẫn đến sai lệch dữ liệu khi hệ thống mở rộng ra nhiều server ở các múi giờ khác nhau, hoặc khi client và server không đồng nhất timezone. Khuyến nghị gần như tuyệt đối trong cộng đồng PostgreSQL là **luôn dùng `TIMESTAMPTZ`** (`TIMESTAMP WITH TIME ZONE`) — PostgreSQL sẽ tự động lưu trữ nội bộ dưới dạng UTC và quy đổi hiển thị theo timezone của session khi truy vấn, loại bỏ hoàn toàn nhầm lẫn múi giờ. Đã sửa cả đoạn code này và đoạn code Soft Delete bên dưới.

> `DEFAULT` không tự động cập nhật `updated_at`:** Đây là một hiểu lầm rất phổ biến. Mệnh đề `DEFAULT now()` **chỉ áp dụng khi `INSERT`** — nó **không** tự động cập nhật lại giá trị mỗi khi có `UPDATE`. Nếu chỉ khai báo như bản gốc, cột `updated_at` sẽ giữ nguyên giá trị bằng `created_at` mãi mãi, kể cả sau khi dữ liệu đã được sửa nhiều lần — một lỗi âm thầm rất khó phát hiện. Để tự động cập nhật đúng, PostgreSQL cần một **trigger**:
>
> ```sql
> CREATE OR REPLACE FUNCTION set_updated_at()
> RETURNS TRIGGER AS $$
> BEGIN
>     NEW.updated_at = now();
>     RETURN NEW;
> END;
> $$ LANGUAGE plpgsql;
>
> CREATE TRIGGER trg_orders_updated_at
> BEFORE UPDATE ON orders
> FOR EACH ROW
> EXECUTE FUNCTION set_updated_at();
> ```
>
> (Với MySQL, cú pháp đơn giản hơn: `updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP`.)

**c) Cân nhắc Soft Delete thay vì Hard Delete cho dữ liệu quan trọng**

```sql
deleted_at TIMESTAMPTZ NULL  -- NULL = chưa xóa, có giá trị = đã "xóa mềm"

-- Thay vì DELETE FROM orders WHERE id = 1;
UPDATE orders SET deleted_at = now() WHERE id = 1;

-- Query luôn phải nhớ lọc:
SELECT * FROM orders WHERE deleted_at IS NULL;
```

Trade-off cần biết: Soft Delete giữ được lịch sử/audit trail, tránh mất dữ liệu do lỗi thao tác, nhưng làm mọi query đều phải thêm điều kiện lọc `deleted_at IS NULL` (dễ quên sót — rủi ro bug rò rỉ dữ liệu đã xóa), và dữ liệu vẫn chiếm dung lượng lưu trữ mãi mãi trừ khi có job dọn dẹp định kỳ.

> khi kết hợp Soft Delete với ràng buộc `UNIQUE`. Ví dụ: nếu cột `email` có ràng buộc `UNIQUE`, và một user với email `a@x.com` bị xóa mềm (`deleted_at` có giá trị), hệ thống **vẫn không cho phép** tạo user mới với cùng email đó — vì ràng buộc `UNIQUE` mặc định áp dụng cho toàn bộ bảng, không phân biệt bản ghi đã xóa mềm hay chưa. Giải pháp chuẩn trong PostgreSQL là dùng **partial unique index** — chỉ áp dụng ràng buộc duy nhất cho các bản ghi chưa bị xóa:
>
> ```sql
> CREATE UNIQUE INDEX idx_users_email_active
> ON users (email)
> WHERE deleted_at IS NULL;
> ```
>
> Đây là cạm bẫy rất thường gặp trong thực tế khi triển khai Soft Delete mà không lường trước.

**d) Luôn ràng buộc NOT NULL cho cột bắt buộc phải có giá trị**

Đừng để mặc định — mỗi cột nên tự hỏi "cột này có được phép NULL về mặt nghiệp vụ không?" Nếu không → khai báo `NOT NULL` tường minh, giúp DB tự bảo vệ tính đúng đắn thay vì dựa vào tầng application.

**e) Dùng đúng kiểu dữ liệu, tránh lạm dụng VARCHAR/TEXT cho mọi thứ**

```sql
-- Sai: dùng VARCHAR cho tiền tệ (mất chính xác, không so sánh số học được)
price VARCHAR(20)

-- Đúng: dùng NUMERIC cho tiền tệ (chính xác tuyệt đối, không lỗi làm tròn float)
price NUMERIC(10,2)

-- Sai: dùng VARCHAR cho enum-like status, không ràng buộc giá trị
status VARCHAR(20)

-- Đúng: ràng buộc rõ giá trị hợp lệ
status VARCHAR(20) CHECK (status IN ('pending', 'completed', 'cancelled'))
```

**f) Tránh thiết kế "God Table" (bảng ôm quá nhiều trách nhiệm)**

Liên hệ trực tiếp nguyên tắc Chuẩn hóa (Normalization) ở phần dưới — nếu một bảng có quá nhiều cột không liên quan trực tiếp đến nhau (ví dụ: bảng `users` vừa chứa thông tin đăng nhập, vừa chứa địa chỉ giao hàng, vừa chứa lịch sử thanh toán), nên tách thành các bảng riêng theo đúng trách nhiệm (`users`, `addresses`, `payment_methods`).

**g) Luôn đặt Index cho Foreign Key nếu thường xuyên JOIN/filter theo nó**

Đây là chi tiết đặc thù theo hệ quản trị cơ sở dữ liệu cần lưu ý: **PostgreSQL không tự động tạo index cho cột khóa ngoại (Foreign Key)** — khác với Primary Key (luôn tự động có index vì PK ngầm định kèm ràng buộc `UNIQUE`). Nếu bảng con thường xuyên `JOIN` hoặc `WHERE` theo khóa ngoại mà không tạo index thủ công, PostgreSQL sẽ phải quét toàn bộ bảng (sequential scan) cho mỗi thao tác — rất tốn hiệu năng khi bảng lớn dần. Cần tạo index thủ công:

```sql
CREATE INDEX idx_orders_user_id ON orders (user_id);
```

> Lưu ý: MySQL (với storage engine InnoDB) có hành vi khác — **tự động tạo index cho cột FK**. Đây là điểm khác biệt quan trọng giữa hai hệ quản trị mà backend engineer làm việc đa nền tảng cần nhớ, tránh áp dụng nhầm giả định từ hệ này sang hệ khác.

**h) Cân nhắc kỹ trước khi denormalize — chỉ làm khi có bằng chứng hiệu năng cụ thể**

Đừng denormalize (phi chuẩn hóa) chỉ vì "phòng khi cần" — hãy **đo trước bằng công cụ phân tích truy vấn** (ví dụ lệnh `EXPLAIN ANALYZE` trong PostgreSQL) để xác nhận thực sự có vấn đề hiệu năng do phải JOIN nhiều bảng, trước khi đánh đổi lấy sự trùng lặp dữ liệu và rủi ro không đồng bộ (data inconsistency) mà denormalization luôn kèm theo.

---

### Cấu trúc và Chuẩn hóa dữ liệu (Normalization & Denormalization)

- **Chuẩn hóa (Normalized):** Áp dụng chuẩn hóa (thường dừng ở dạng 3NF — Third Normal Form, mức đủ dùng cho phần lớn hệ thống OLTP) để loại bỏ dữ liệu dư thừa. Phân tách bảng dựa trên các thực thể (entities) độc lập.
- **Phi chuẩn hóa (Denormalized) khi cần:** Nếu việc chuẩn hóa tạo ra quá nhiều bảng cần nối (JOIN) và đã được đo đạc xác nhận là bottleneck hiệu năng, hãy cân nhắc phi chuẩn hóa cho các hệ thống đọc là chủ yếu (ví dụ: lưu thêm `user_name` trong bảng `orders` để truy xuất nhanh hơn, chấp nhận đánh đổi phải đồng bộ lại nếu `user_name` gốc thay đổi).
- **Tính nguyên tử (Atomicity):** Mỗi cột chỉ nên chứa một giá trị đơn (ví dụ: chia `fullname` thành `first_name` và `last_name` nếu cần xử lý độc lập, như sắp xếp theo họ hoặc tên).

### Khóa và Ràng buộc (Keys & Constraints)

- **Khóa chính (Primary Key — PK):** Sử dụng khóa đại diện không mang ý nghĩa nghiệp vụ (surrogate key) như `UUID` hoặc `BIGINT AUTO-INCREMENT`. Đảm bảo PK không thay đổi theo thời gian.
- **Khóa ngoại (Foreign Key — FK):** Sử dụng FK để liên kết các bảng, đảm bảo tính toàn vẹn tham chiếu (Referential Integrity) — database sẽ tự động từ chối thao tác tạo bản ghi tham chiếu đến một giá trị không tồn tại ở bảng cha, hoặc ngăn xóa bản ghi cha còn đang được tham chiếu (tuỳ theo chính sách `ON DELETE`/`ON UPDATE` được cấu hình).
- **Ràng buộc (Constraints):** Sử dụng các ràng buộc như `NOT NULL`, `DEFAULT`, `CHECK` để ngăn chặn dữ liệu không hợp lệ ngay từ tầng database — không phụ thuộc hoàn toàn vào validation ở tầng application, vốn có thể bị bỏ sót hoặc bypass (ví dụ khi có nhiều service khác nhau cùng ghi vào một database).

### Kiểu dữ liệu (Data Types)

- **Tiền tệ:** Luôn sử dụng kiểu `DECIMAL` hoặc `NUMERIC` cho dữ liệu tiền tệ hoặc tài chính. Tránh dùng `FLOAT` hoặc `REAL` vì đây là kiểu dấu phẩy động (floating-point) có thể gây sai số làm tròn — không phù hợp cho các phép tính đòi hỏi chính xác tuyệt đối như tiền tệ.

  > quan niệm sai phổ biến về `VARCHAR` vs `TEXT` trong PostgreSQL:** Đây là một trong những hiểu lầm phổ biến nhất mà nhiều lập trình viên mang theo từ kinh nghiệm với các hệ quản trị khác. Theo đúng tài liệu chính thức của PostgreSQL: **không có sự khác biệt về hiệu năng** giữa `VARCHAR(n)`, `VARCHAR` (không giới hạn), và `TEXT` — cả ba đều dùng chung cơ chế lưu trữ nội bộ (`varlena`). Trên thực tế, PostgreSQL khuyến nghị dùng `TEXT` hoặc `VARCHAR` không giới hạn thay vì `VARCHAR(n)`, trừ khi việc giới hạn độ dài là một **ràng buộc nghiệp vụ thực sự cần thiết** (ví dụ: tên đăng nhập tối đa 30 ký tự) — không phải vì lý do hiệu năng hay bộ nhớ đệm. Đây là điểm khác biệt quan trọng so với một số hệ quản trị khác (như SQL Server) nơi sự lựa chọn kiểu có thể ảnh hưởng đến hiệu năng theo cách khác. Đã loại bỏ khuyến nghị sai này.

- **Trạng thái:** Dùng kiểu số nguyên (`INTEGER`) hoặc `ENUM` nếu tập giá trị đã cố định rõ ràng và ít khi thay đổi.

 Bảng dưới đây làm rõ trade-off giữa ba lựa chọn phổ biến:
  >
  > | Cách tiếp cận | Ưu điểm | Nhược điểm |
  > |---|---|---|
  > | `VARCHAR` + `CHECK` | Dễ đọc trực tiếp trong query (`status = 'pending'`), dễ thêm/sửa giá trị hợp lệ (chỉ cần `ALTER TABLE ... DROP/ADD CONSTRAINT`) | Tốn thêm vài byte lưu trữ so với số nguyên (không đáng kể ở đa số hệ thống) |
  > | `INTEGER` + bảng tra cứu (lookup table) | Tiết kiệm không gian lưu trữ tối đa, hiệu năng so sánh/index nhanh nhất | Query kém trực quan hơn (`status_id = 2` thay vì `status = 'completed'`), cần JOIN thêm để hiển thị tên trạng thái |
  > | `ENUM` gốc của database | Ràng buộc kiểu dữ liệu chặt chẽ ở tầng thấp nhất | Trong PostgreSQL, việc **xóa** một giá trị khỏi kiểu ENUM đã tồn tại gần như không thể thực hiện trực tiếp (phải tạo kiểu mới và migrate dữ liệu) — gây khó khăn đáng kể khi nghiệp vụ thay đổi thường xuyên |
  >
  > Với đa số hệ thống backend có nghiệp vụ còn thay đổi (đặc biệt giai đoạn đầu sản phẩm), `VARCHAR + CHECK` thường là lựa chọn cân bằng và linh hoạt nhất.

### Chỉ mục (Indexing)

- **Chỉ mục cho khóa chính:** Primary Key luôn **tự động** được đánh index, vì PK ngầm định kèm theo ràng buộc `UNIQUE` — mọi ràng buộc `UNIQUE` trong PostgreSQL đều tự động tạo index đi kèm.
- **Chỉ mục cho khóa ngoại:** **Không** tự động được tạo trong PostgreSQL (xem giải thích chi tiết ở mục (g) phía trên) — cần tạo thủ công nếu cột FK thường xuyên được dùng để JOIN hoặc filter.
- **Cân nhắc đánh chỉ mục:** Tạo index cho các cột thường xuyên dùng trong mệnh đề `WHERE`, `ORDER BY`, `GROUP BY`, hoặc `JOIN`.
- **Không lạm dụng:** Tránh tạo quá nhiều index không cần thiết — mỗi index bổ sung sẽ làm chậm các câu lệnh ghi dữ liệu (`INSERT`, `UPDATE`, `DELETE`, vì mỗi thao tác ghi đều phải cập nhật lại toàn bộ index liên quan) và chiếm thêm dung lượng ổ đĩa.

### Tài liệu hóa và Tiến hóa cấu trúc (Documentation & Evolution)

- **Tài liệu hóa schema:** Dùng tính năng `COMMENT` trong SQL để mô tả ý nghĩa của bảng và các cột — giúp người mới tham gia dự án hoặc công cụ tự động sinh tài liệu (ví dụ dbdocs, SchemaSpy) hiểu đúng ý nghĩa nghiệp vụ mà không cần hỏi lại.
- **Quản lý migration:** Dùng công cụ migration (ví dụ: `golang-migrate`, `goose` cho hệ sinh thái Go; `Flyway`, `Liquibase` cho đa ngôn ngữ) để kiểm soát việc thay đổi cấu trúc, lưu các tập lệnh tiến hóa (migration script) này vào Git để có lịch sử thay đổi rõ ràng, có thể rollback khi cần.
- **Thay đổi không ngừng (Expand-Contract pattern):** Là kỹ thuật thay đổi schema an toàn khi hệ thống đang chạy production và không thể downtime — chia mỗi thay đổi breaking thành nhiều bước nhỏ không breaking. Khi thêm cột mới, đặt chế độ `NULL` hoặc có `DEFAULT` để không phá vỡ các câu lệnh `INSERT` cũ chưa được cập nhật (giai đoạn "Expand" — mở rộng). Không đổi tên hoặc xóa cột trực tiếp nếu ứng dụng chưa được cập nhật hết để ngừng sử dụng cột đó (giai đoạn "Contract" — thu hẹp, chỉ thực hiện sau khi đã xác nhận không còn code nào phụ thuộc vào cấu trúc cũ), nhằm tránh lỗi gián đoạn dịch vụ giữa các lần deploy không đồng thời.

---

## Bổ sung kiến thức

### 1. Kiểm tra hiệu năng index bằng EXPLAIN ANALYZE

Trước khi quyết định thêm index hay denormalize (theo tinh thần mục h), nên làm quen với cách đọc kết quả `EXPLAIN ANALYZE` trong PostgreSQL:

```sql
EXPLAIN ANALYZE
SELECT * FROM orders WHERE user_id = 123;
```

Nếu kết quả hiển thị `Seq Scan` (quét tuần tự toàn bảng) thay vì `Index Scan`, đó là dấu hiệu cần xem xét thêm index cho cột đang filter — đặc biệt nếu bảng có số lượng bản ghi lớn và câu lệnh này được gọi thường xuyên.

### 2. Composite Index và thứ tự cột

Bản gốc chỉ đề cập index đơn cột. Trong thực tế, nhiều truy vấn cần lọc theo nhiều điều kiện cùng lúc — lúc này **composite index** (index nhiều cột) hiệu quả hơn nhiều so với việc tạo nhiều index đơn cột riêng lẻ:

```sql
-- Nếu thường xuyên query: WHERE user_id = ? AND status = 'pending'
CREATE INDEX idx_orders_user_status ON orders (user_id, status);
```

Lưu ý quan trọng: **thứ tự cột trong composite index có ý nghĩa** — index trên `(user_id, status)` tối ưu cho query lọc theo `user_id` hoặc cả `user_id` + `status`, nhưng **không** tối ưu cho query chỉ lọc riêng theo `status` (do đặc tính của B-tree index, chỉ tận dụng được hiệu quả khi truy vấn dùng cột đứng trước trong index).

### 3. Ví dụ migration theo Expand-Contract pattern

Minh họa cụ thể luồng đổi tên cột `full_name` thành `display_name` an toàn qua nhiều lần deploy:

```
Migration 1 (Expand):  Thêm cột display_name (NULL), backfill dữ liệu từ full_name
Deploy code:           Ứng dụng ghi đồng thời vào cả 2 cột, đọc ưu tiên display_name
Migration 2 (Verify):  Xác nhận không còn code nào đọc/ghi full_name
Migration 3 (Contract): Xóa cột full_name
```

Cách làm này tránh được tình huống hai phiên bản ứng dụng chạy song song (trong lúc rolling deploy) bị lỗi vì một bên còn dùng schema cũ, một bên đã dùng schema mới.