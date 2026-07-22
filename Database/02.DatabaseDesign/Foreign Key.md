## **Foreign Key (Khóa ngoại)**

**Định nghĩa:** Là 1 cột (hoặc tổ hợp cột) ở 1 bảng, *có giá trị tham chiếu đến Primary Key (hoặc Unique key) của 1 bảng khác* — dùng để liên kết dữ liệu giữa các bảng và đảm bảo Referential Integrity.

```sql
CREATE TABLE customers (
    id   SERIAL PRIMARY KEY,
    name VARCHAR(100)
);

CREATE TABLE orders (
    id          SERIAL PRIMARY KEY,
    customer_id INT REFERENCES customers(id),  -- Foreign Key
    total       NUMERIC(10,2)
);
```

## **Các hành vi khi bảng cha bị xóa/sửa — `ON DELETE` / `ON UPDATE`**

*ví dụ:*

```sql
-- 1. RESTRICT (mặc định ở nhiều DB): không cho xóa customer nếu còn orders liên quan
customer_id INT REFERENCES customers(id) ON DELETE RESTRICT

-- 2. CASCADE: xóa customer thì tự động xóa luôn mọi orders của khách đó
customer_id INT REFERENCES customers(id) ON DELETE CASCADE

-- 3. SET NULL: xóa customer thì orders.customer_id tự chuyển thành NULL (đơn hàng vẫn còn, mất liên kết)
customer_id INT REFERENCES customers(id) ON DELETE SET NULL

-- 4. SET DEFAULT: xóa customer thì customer_id chuyển về giá trị DEFAULT đã khai báo
customer_id INT REFERENCES customers(id) DEFAULT 0 ON DELETE SET DEFAULT

-- 5. NO ACTION: gần giống RESTRICT, nhưng kiểm tra ràng buộc trễ hơn (cuối transaction thay vì ngay lập tức) — mặc định thực tế của Postgres nếu không khai báo gì
customer_id INT REFERENCES customers(id)  -- không ghi ON DELETE gì → mặc định NO ACTION
```

### **Cách hoạt động của khóa ngoại trong cơ sở dữ liệu**
#### **Tính toàn vẹn tham chiếu**
Tính toàn vẹn tham chiếu là khái niệm cốt lõi trong quản trị cơ sở dữ liệu mà khóa ngoại có nhiệm vụ phải thực thi. Nguyên tắc này đảm bảo rằng giá trị của khóa ngoại trong cơ sở dữ liệu luôn luôn trỏ đến một bản ghi có thực và hợp lệ trong bảng cha.

Điều này có nghĩa là bất kỳ nỗ lực nào nhằm thêm mới, cập nhật hoặc xóa dữ liệu trong bảng con đều phải tuân thủ nghiêm ngặt các ràng buộc do khóa ngoại đặt ra. 

Ví dụ: Nếu bạn cố gắng xóa một khách hàng trong khi họ vẫn còn các đơn hàng chưa xử lý, cơ sở dữ liệu sẽ ngay lập tức ngăn chặn hành động này để bảo vệ tính nhất quán của dữ liệu.

#### **Bảng Cha và bảng Con**
Trong bối cảnh của khóa ngoại, mối quan hệ giữa các bảng được phân vai rất rõ ràng:
- Bảng Cha (Parent Table): Là bảng chứa Khóa chính (Primary Key).
- Bảng Con (Child Table): Là bảng chứa Khóa ngoại (Foreign Key).

Khóa ngoại trong bảng con sẽ tham chiếu trực tiếp đến khóa chính trong bảng cha, từ đó thiết lập một sợi dây liên kết bền vững giữa hai bên.

Hãy hình dung ví dụ sau để dễ hiểu hơn:
- Bảng Cha (Customers): Chứa thông tin khách hàng, với CustomerID là khóa chính.
- Bảng Con (Orders): Chứa thông tin đơn hàng, với CustomerID là khóa ngoại để liên kết ngược lại với bảng Customers.

Cấu trúc này đảm bảo rằng mỗi đơn hàng được tạo ra đều phải gắn liền với một khách hàng hợp lệ, qua đó duy trì sự chính xác tuyệt đối cho dữ liệu.

**Ví dụ thực hành chạy được để thấy khác biệt CASCADE vs RESTRICT:**
```sql
CREATE TABLE customers_test (id SERIAL PRIMARY KEY, name VARCHAR(50));
CREATE TABLE orders_cascade (id SERIAL PRIMARY KEY, customer_id INT REFERENCES customers_test(id) ON DELETE CASCADE);
CREATE TABLE orders_restrict (id SERIAL PRIMARY KEY, customer_id INT REFERENCES customers_test(id) ON DELETE RESTRICT);

INSERT INTO customers_test (name) VALUES ('An');  -- id = 1
INSERT INTO orders_cascade (customer_id) VALUES (1);
INSERT INTO orders_restrict (customer_id) VALUES (1);

DELETE FROM customers_test WHERE id = 1;
-- ❌ ERROR: update or delete on table "customers_test" violates foreign key constraint
-- (bị chặn bởi orders_restrict, dù orders_cascade cho phép — Postgres kiểm tra TẤT CẢ FK trỏ tới, chỉ cần 1 cái RESTRICT là chặn toàn bộ thao tác)
```

*OutPut*:
```
Failed to run sql query: ERROR:  23503: update or delete on table "customers_test" violates foreign key constraint "orders_restrict_customer_id_fkey" on table "orders_restrict"
DETAIL:  Key (id)=(1) is still referenced from table "orders_restrict".
```

> Lưu ý quan trọng: nếu 1 bảng cha có nhiều bảng con tham chiếu tới với các quy tắc ON DELETE khác nhau, chỉ cần 1 trong số đó là RESTRICT/NO ACTION và có dữ liệu liên quan, thao tác DELETE sẽ bị chặn hoàn toàn — Postgres không "xóa 1 phần".

## **Lợi ích khóa ngoại trong cơ sở dữ liệu**
#### **Đảm bảo toàn vẹn dữ liệu**
Khóa ngoại trong cơ sở dữ liệu hoạt động như một người gác cổng nghiêm ngặt, thực thi tính toàn vẹn tham chiếu để đảm bảo mối quan hệ giữa các bảng luôn nhất quán. 

Cơ chế này ngăn chặn triệt để sự xuất hiện của các bản ghi "mồ côi" – tức là những dữ liệu ở bảng con tồn tại lơ lửng mà không gắn liền với bất kỳ đối tượng thực tế nào ở bảng cha. Nhờ đó, độ chính xác của toàn bộ cơ sở dữ liệu luôn được duy trì ở mức cao nhất.

#### **Đơn giản hóa quản lý dữ liệu**
Việc thiết lập khóa ngoại trong cơ sở dữ liệu giúp định hình rõ ràng cấu trúc liên kết giữa các bảng, biến một kho dữ liệu phức tạp trở nên ngăn nắp và dễ quản lý hơn. Khi các mối quan hệ đã được xác định minh bạch, việc thực hiện các truy vấn dữ liệu hay trích xuất thông tin liên quan trở nên trực quan và nhanh chóng, góp phần nâng cao hiệu suất vận hành chung của hệ thống.

#### **Tự động hóa hành động lan truyền **
Một tính năng kỹ thuật mạnh mẽ của khóa ngoại là khả năng cấu hình các hành động lan truyền như ON DELETE CASCADE hoặc ON UPDATE CASCADE. Thay vì phải cập nhật thủ công từng bảng, tính năng này cho phép mọi thay đổi ở bảng cha (như xóa hoặc sửa đổi dữ liệu) sẽ tự động được áp dụng đồng bộ xuống các bảng con.

Điều này giúp giảm thiểu đáng kể khối lượng công việc thủ công, đồng thời loại bỏ rủi ro sai lệch dữ liệu do con người quên cập nhật ở các bảng liên quan.

#### **Gánh nặng về hiệu năng**
Việc triển khai khóa ngoại trong cơ sở dữ liệu có thể vô tình tạo ra áp lực lên hiệu suất của hệ thống, đặc biệt là đối với các cơ sở dữ liệu quy mô lớn có tần suất cập nhật và xóa dữ liệu liên tục. 

Nguyên nhân là do hệ thống buộc phải thực hiện thêm các bước kiểm tra tính toàn vẹn tham chiếu mỗi khi có thay đổi xảy ra. Những thao tác kiểm tra ngầm này tích tụ lại có thể làm chậm tốc độ xử lý chung của các tác vụ.

**Foreign Key có bắt buộc phải có Index không?**

Câu hỏi rất hay bị bỏ sót: PostgreSQL KHÔNG tự động tạo index trên cột Foreign Key (khác với Primary Key — luôn tự động có index). Nếu bạn thường xuyên JOIN hoặc query theo cột FK mà không có index, hiệu năng sẽ kém.

```sql
-- Nên tạo thêm index thủ công cho FK nếu hay JOIN/filter theo nó
CREATE INDEX idx_orders_customer_id ON orders(customer_id);
```