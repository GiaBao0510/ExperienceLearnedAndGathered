# **1. Nhóm quản lý kết nối & hệ thống

#### `@@CONECTION` (Số lượng kết nối)
Trong PostgreSQL, chúng ta luôn quan tâm số lượng kết nối hiện tại (active) hơn là lịch sử

 Có thể truy vấn bằng `pg_stat_activity`.
- **`pg_stat_activity`** là một View(bảng ảo) chứa thông tin về các process đang chạy, bao gồm luôn cả user nào nào đang kết nối, đang chạy câu lệnh gì,...
- Giúp chuẩn đoán và khắc phục sự cố (troubleshooting) các vấn đề về hiệu năng truy vấn.

*Lệnh:*
```postgresql
-- Xem số kết nối hiện tại
select count(*) from pg_stat_activity;

-- Xem thông tin chi tiết kết nối
select pid, datname, usename, client_addr ,state, query
from pg_stat_activity;
```

#### `@@MAX_CONNECTIONS` (Số kết nối tối đa cho phép)

Dùng để xem số lượng kết nối đến cơ sở dữ liệu tối đa cho phép
**PostgreSQL** lưu các cấu hình trong **pg_settings**. Sử dụng lệnh `show` sau đây là cách nhanh để lấy thông tin cấu hình.

*Lệnh:*
```postgresql
show max_connections;

-- Hoặc truy vấn bằng cài đặt
select name, setting 
from pg_settings 
where name = 'max_connections';
```


---
# **2. Nhóm xử lý dữ liệu & Auto-increment

