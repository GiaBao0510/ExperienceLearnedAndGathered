## **Primary Key (Khóa chính)**

Định nghĩa: Là 1 cột (hoặc tổ hợp nhiều cột) có **giá trị duy nhất cho mỗi dòng và không bao giờ NULL**, dùng để định danh chính xác 1 bản ghi trong bảng. Đây chính là thành phần đảm bảo Entity Integrity mà bạn đã học ở Phần 5.

Tổng độ dài của các cột trong khóa chính không được vượt quá 900 byte và số lượng cột tối đa là 16, vì vậy trong thực tế thường ưu tiên chọn các cột kiểu số, ngắn để tối ưu hiệu năng.

**Những chức năng chính của Primary Key là gì?**
Khóa chính (Primary key) trong SQL Server giữ vai trò trung tâm trong thiết kế và vận hành cơ sở dữ liệu, thể hiện qua các khía cạnh sau:

- Lưu trữ giá trị duy nhất cho từng bản ghi: Mỗi dòng dữ liệu trong bảng được gắn với một giá trị duy nhất tại cột (hoặc nhóm cột) khóa chính, giúp xác định và truy xuất chính xác từng record mà không bị nhầm lẫn.
- Tạo liên kết giữa các bảng: Giá trị khóa chính thường được dùng làm khóa ngoại ở các bảng khác, nhờ đó các bảng trong cùng cơ sở dữ liệu có thể liên kết với nhau và đảm bảo dữ liệu khớp chính xác.
- Mỗi bảng chỉ có một khóa chính: Trong một table, chỉ tồn tại duy nhất một primary key, nhưng khóa này có thể được nhiều bảng khác tham chiếu thông qua foreign key, hình thành mạng lưới quan hệ logic trong toàn bộ hệ thống dữ liệu.
- Hỗ trợ công cụ như Access tự động thiết lập: Khi tạo bảng mới trong Microsoft Access, nếu người dùng không định nghĩa, hệ thống có thể tự động sinh một trường ID kiểu AutoNumber làm khóa chính, giúp đơn giản hóa thao tác thiết kế bảng cho người dùng.

> Lời khuyên từ chuyên gia: Dù các công cụ có thể tự động tạo, nhưng trong môi trường doanh nghiệp, việc thiết kế khóa chính cần được tính toán kỹ từ đầu để tránh việc phải thay đổi cấu trúc bảng sau này—một thao tác rất tốn kém tài nguyên và dễ gây lỗi hệ thống.

## **Phạm vi áp dụng đối với Primary Key**
Khóa chính (Primary Key) là thành phần bắt buộc trong thiết kế cơ sở dữ liệu quan hệ, áp dụng xuyên suốt từ các ứng dụng quản lý thông tin cơ bản đến những hệ thống phân tán phức tạp. Trong quy trình phát triển, khóa chính được sử dụng để thiết lập cấu trúc bảng, xây dựng biểu mẫu nhập liệu và định nghĩa các ràng buộc toàn vẹn tham chiếu khi thiết lập quan hệ giữa các thực thể dữ liệu.

Việc triển khai khóa chính đòi hỏi sự chính xác ngay từ giai đoạn thiết kế logic. Quá trình này bao gồm thao tác khai báo kiểu dữ liệu tối ưu, cấu hình thuộc tính tự tăng (auto-increment) và thiết lập chỉ mục cụm (clustered index). Ứng dụng đúng nguyên lý khóa chính giúp hệ thống loại bỏ dữ liệu trùng lặp, tối ưu hóa tốc độ truy vấn thông qua cấu trúc cây chỉ mục và đảm bảo tính nhất quán dữ liệu trong các thao tác cập nhật hoặc truy xuất liên bảng.

## **Các loại Primary Key**

**a) Natural Key (Khóa tự nhiên)** — dùng 1 thuộc tính có sẵn trong dữ liệu thực tế làm khóa:

```sql
CREATE TABLE countries (
    country_code CHAR(2) PRIMARY KEY,  -- VD: 'VN', 'US' — mã quốc gia có sẵn, duy nhất tự nhiên
    country_name VARCHAR(100)
);
```

**b) Surrogate Key (Khóa thay thế/nhân tạo)** — DB tự sinh giá trị (thường là số tự tăng hoặc UUID), không mang ý nghĩa nghiệp vụ:

```sql
CREATE TABLE customers (
    id SERIAL PRIMARY KEY,   -- surrogate key: 1, 2, 3... không có ý nghĩa thực tế
    email VARCHAR(100) UNIQUE
);

-- Hoặc dùng UUID (phổ biến trong hệ thống phân tán, microservices)
CREATE TABLE orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    customer_id INT
);
```

**c) Composite Key (Khóa kết hợp)** — dùng nhiều cột cùng lúc làm khóa chính:

```sql
CREATE TABLE order_items (
    order_id   INT,
    product_id INT,
    quantity   INT,
    PRIMARY KEY (order_id, product_id)  -- composite key
);
```

#### **Natural Key vs Surrogate Key — nên dùng loại nào?**
	
||**Natural Key**|**Surrogate Key**|
|---|---|---|
|Ý nghĩa nghiệp vụ|Có (VD: mã số CMND, email)|Không (chỉ là số định danh nội bộ)|
|Rủi ro|Dữ liệu nghiệp vụ có thể đổi (VD: đổi email, đổi mã SKU) → phá vỡ mọi FK tham chiếu|Không bao giờ đổi, ổn định tuyệt đối|
|Hiệu năng JOIN|Có thể chậm hơn nếu là chuỗi dài|Nhanh (thường là số nguyên)|
|Khuyến nghị thực tế|Chỉ dùng khi chắc chắn giá trị không đổi (mã ISO quốc gia, mã ISO tiền tệ)|Mặc định nên dùng cho hầu hết bảng nghiệp vụ, kể cả khi đã có unique field khác (email) — vẫn nên thêm UNIQUE constraint riêng cho email, giữ id làm PK|

⚠️ Bẫy thực tế hay gặp: nhiều bạn mới dùng email làm Primary Key cho bảng users. Vấn đề: nếu sau này user đổi email, toàn bộ Foreign Key ở các bảng khác tham chiếu đến email đó đều phải cập nhật theo (cực kỳ tốn kém, dễ lỗi). Giải pháp chuẩn: id SERIAL PRIMARY KEY + email VARCHAR UNIQUE riêng biệt.