# **Relationship (Mối quan hệ giữa các bảng)**

## **One-to-One (1-1)**

1 dòng ở bảng A liên kết với đúng 1 dòng ở bảng B, và ngược lại.

```sql
-- Ví dụ: mỗi user có đúng 1 profile chi tiết (tách riêng để bảng users gọn nhẹ)
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(100)
);

CREATE TABLE user_profiles (
    user_id   INT PRIMARY KEY REFERENCES users(id),  -- PK trùng với FK → đảm bảo 1-1
    full_name VARCHAR(100),
    bio       TEXT
);
```

> Chốt kỹ thuật: đặt `user_id` vừa là `PRIMARY KEY` vừa là `FOREIGN KEY` — đây chính là cách hiện thực hóa quan hệ 1-1 chuẩn trong SQL, vì PK đảm bảo mỗi user_id chỉ xuất hiện tối đa 1 lần ở bảng user_profiles.

### **One-to-Many (1-N)**

1 dòng ở bảng A (bên "1") liên kết với nhiều dòng ở bảng B (bên "N"). Đây là quan hệ phổ biến nhất.

```sql
-- 1 customer có nhiều orders
CREATE TABLE orders (
    id          SERIAL PRIMARY KEY,
    customer_id INT REFERENCES customers(id)  -- FK đặt ở bên "N" (orders)
);
```

Nguyên tắc: FK luôn đặt ở bảng thuộc phía "N" (phía nhiều), trỏ về PK của bảng phía "1".

### **Many-to-Many (M-N)**

1 dòng ở bảng A có thể liên kết với nhiều dòng ở bảng B, và ngược lại. Không thể hiện thực hóa trực tiếp bằng 1 FK đơn giản — bắt buộc phải qua bảng trung gian (junction/bridge table).

```sql
-- 1 sản phẩm có thể thuộc nhiều tag, 1 tag có thể gắn cho nhiều sản phẩm
CREATE TABLE products (id SERIAL PRIMARY KEY, name VARCHAR(100));
CREATE TABLE tags (id SERIAL PRIMARY KEY, name VARCHAR(50));

CREATE TABLE product_tags (
    product_id INT REFERENCES products(id),
    tag_id     INT REFERENCES tags(id),
    PRIMARY KEY (product_id, tag_id)  -- composite PK, đồng thời là bảng trung gian
);
```

Đây chính là kỹ thuật bạn đã áp dụng ở order_items (Weak Entity, Phần 6.3) — về bản chất, bảng trung gian M-N và Weak Entity thường trùng nhau trong thực hành, dù khái niệm lý thuyết xuất phát khác nhau.