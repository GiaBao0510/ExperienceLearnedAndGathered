## **Naming Convention (Quy ước đặt tên)**

Đây là phần rất thực tế, ít lý thuyết nhưng ảnh hưởng lớn đến khả năng bảo trì lâu dài — và cũng là điểm dễ bị đánh giá thấp khi review code/schema trong phỏng vấn hoặc code review thực tế.

**Quy ước phổ biến nhất (PostgreSQL convention — snake_case)**

||||
|---|---|---|
|Tên bảng|	số nhiều, snake_case|customers, order_items|
|Tên cột|snake_case, số ít|first_name, created_at|
|Primary Key|luôn đặt tên id|id SERIAL PRIMARY KEY|
|Foreign Key|<tên_bảng_số_ít>_id|customer_id, product_id|
|Bảng trung gian (M-N)|ghép tên 2 bảng, số ít, theo alphabet|product_tags, student_courses|
|Boolean|	tiền tố is_/has_|is_active, has_verified_email|
|Timestamp|hậu tố _at|created_at, updated_at, deleted_at|
|Date (không giờ)|hậu tố _date|birth_date, due_date|
|Index|idx_<bảng>_<cột>|idx_orders_customer_id|
|Foreign Key Constraint|fk_<bảng>_<bảng_tham_chiếu>	|fk_orders_customers|
|Unique Constraint|uq_<bảng>_<cột>|uq_customers_email|
|Check Constraint|chk_<bảng>_<mô_tả>|chk_products_price_positive|

**Vì sao dùng số nhiều cho tên bảng, số ít cho tên cột?**

Đây là câu hỏi hay bị hỏi: quy ước phổ biến nhất (và Postgres/Rails/Laravel convention) là bảng số nhiều (vì 1 bảng chứa nhiều bản ghi — customers chứa nhiều "customer"), còn cột số ít (vì 1 cột trong 1 dòng chỉ chứa 1 giá trị — first_name, không phải first_names). Một số team/framework khác (VD: Oracle truyền thống) lại dùng số ít cho cả bảng — không có đúng/sai tuyệt đối, quan trọng là nhất quán trong toàn bộ dự án.

```sql
CREATE TABLE customers (
    id         SERIAL PRIMARY KEY,
    email      VARCHAR(100) NOT NULL,
    is_active  BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),

    CONSTRAINT uq_customers_email UNIQUE (email)
);

CREATE TABLE orders (
    id          SERIAL PRIMARY KEY,
    customer_id INT NOT NULL,
    total       NUMERIC(10,2),
    order_date  DATE DEFAULT CURRENT_DATE,

    CONSTRAINT fk_orders_customers FOREIGN KEY (customer_id) REFERENCES customers(id)
);

CREATE INDEX idx_orders_customer_id ON orders(customer_id);
```

**Vì sao đặt tên rõ ràng cho constraint** (fk_orders_customers thay vì để Postgres tự sinh tên)? Vì khi có lỗi vi phạm constraint, thông báo lỗi sẽ hiện đúng tên bạn đặt (dễ debug hơn nhiều so với tên tự sinh dạng orders_customer_id_fkey), và khi cần DROP CONSTRAINT để sửa sau này, bạn biết chính xác tên cần gõ mà không cần query lại system catalog.