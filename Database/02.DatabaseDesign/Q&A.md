**1. Natural Key và Surrogate Key khác nhau thế nào? Vì sao thực tế thường ưu tiên Surrogate Key?**

Natural Key dùng thuộc tính có sẵn trong dữ liệu thực tế (VD: email, mã số) làm khóa chính; Surrogate Key là giá trị nhân tạo do DB sinh ra (VD: số tự tăng, UUID), không mang ý nghĩa nghiệp vụ. Thực tế ưu tiên Surrogate Key vì Natural Key có rủi ro thay đổi theo thời gian (VD: user đổi email), khi đó mọi Foreign Key tham chiếu tới đều phải cập nhật theo — cực kỳ tốn kém và dễ lỗi. Surrogate Key ổn định vĩnh viễn, không bao giờ cần đổi.

**2. Sự khác nhau giữa ON DELETE CASCADE, SET NULL, RESTRICT, NO ACTION?**

CASCADE: xóa dòng cha thì tự động xóa các dòng con liên quan. SET NULL: xóa dòng cha thì FK ở con chuyển thành NULL, dòng con vẫn tồn tại. RESTRICT: chặn việc xóa dòng cha ngay lập tức nếu còn dòng con tham chiếu. NO ACTION: về bản chất gần giống RESTRICT nhưng thời điểm kiểm tra ràng buộc trễ hơn (có thể đến cuối transaction thay vì ngay lập tức) — đây cũng là hành vi mặc định của Postgres nếu không khai báo ON DELETE gì.

**3. Foreign Key có tự động được đánh Index không? Vì sao cần quan tâm điều này?**

Không. Khác với Primary Key (luôn tự động có Unique Index), PostgreSQL không tự động tạo index trên cột Foreign Key. Nếu ứng dụng thường xuyên JOIN hoặc filter theo cột FK mà không tạo index thủ công, các thao tác đó sẽ chậm hơn đáng kể, đặc biệt với bảng lớn — cần chủ động CREATE INDEX cho các cột FK hay dùng.

**4. Làm sao hiện thực hóa quan hệ Many-to-Many trong SQL? Vì sao không thể dùng 1 FK đơn giản?**

Không thể dùng 1 FK đơn giản vì FK chỉ cho phép 1 dòng ở bảng con trỏ đến đúng 1 dòng ở bảng cha — không biểu diễn được việc 1 dòng liên kết với nhiều dòng ở cả 2 phía cùng lúc. Giải pháp là tạo bảng trung gian (junction table) chứa 2 cột FK trỏ tới cả 2 bảng gốc, thường dùng composite Primary Key trên cả 2 cột đó.

**5. Giải thích 2 nghĩa khác nhau của "Cardinality" trong CSDL.**

Nghĩa 1 (trong ERD/thiết kế): mô tả số lượng quan hệ tối đa giữa 2 entity — 1-1, 1-N, hay M-N. Nghĩa 2 (trong tối ưu query): là số lượng giá trị khác nhau (distinct values) trong 1 cột so với tổng số dòng, dùng để đo độ chọn lọc (selectivity) — Cardinality càng cao (nhiều giá trị khác nhau) thì Index trên cột đó càng hiệu quả.

**6. Vì sao nên đặt tên constraint tường minh (VD: fk_orders_customers) thay vì để DB tự sinh tên?**

Vì khi có lỗi vi phạm ràng buộc, thông báo lỗi sẽ hiển thị đúng tên bạn đặt, giúp debug nhanh và dễ hiểu hơn nhiều so với tên tự sinh khó đọc. Ngoài ra khi cần sửa/xóa constraint sau này (ALTER TABLE ... DROP CONSTRAINT), bạn biết chính xác tên cần dùng mà không phải tra cứu lại hệ thống catalog của DB.

**7. Soft Delete là gì? Trade-off khi áp dụng?**

Soft Delete là kỹ thuật đánh dấu 1 dòng "đã xóa" bằng cột deleted_at (hoặc is_deleted) thay vì xóa vật lý khỏi bảng bằng DELETE. Ưu điểm: giữ được lịch sử/audit trail, có thể khôi phục dữ liệu, tránh mất mát do thao tác nhầm. Nhược điểm: mọi câu query đọc dữ liệu đều phải nhớ thêm điều kiện lọc deleted_at IS NULL (dễ bị quên sót, gây lộ dữ liệu đã "xóa"), và dữ liệu vẫn tồn tại chiếm dung lượng mãi mãi nếu không có cơ chế dọn dẹp định kỳ.

**8. Cho tình huống: bạn thiết kế bảng users chứa cả thông tin đăng nhập, địa chỉ, và lịch sử thanh toán trong 1 bảng duy nhất với hơn 40 cột. Đánh giá thiết kế này và đề xuất cải thiện.**

Đây là dạng "God Table" — vi phạm nguyên tắc tách trách nhiệm rõ ràng (liên hệ Normalization Phần 6): các nhóm thông tin (đăng nhập, địa chỉ, thanh toán) có chu kỳ thay đổi và mục đích sử dụng khác nhau nên nên tách thành các bảng riêng (users, addresses, payment_methods) liên kết qua Foreign Key. Lợi ích: giảm dư thừa nếu 1 user có nhiều địa chỉ/phương thức thanh toán (quan hệ 1-N thay vì gò ép vào 1 dòng), dễ maintain, giảm rủi ro khóa (lock) tranh chấp khi nhiều phần khác nhau của hệ thống cùng update các nhóm thông tin khác nhau của cùng 1 user.