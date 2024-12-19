## **1. Index là gì?**
- Index là một dạng cấu trúc dữ liệu giúp xác định nhanh chóng các records trong table.
- Nếu như không có Index thì SQL sẽ quét toàn bộ bảng để tìm kiếm records cần tìm. Dữ liệu càng lớn thì tốc đọ truy vấn sẽ càng chậm.

## **2. Ưu nhược điểm của Index?**

**_Ưu điểm_:**
- Tăng tốc độ tìm kiếm records theo lệnh where
- Không chỉ giới hạn trong mệnh đề SELECT mà còn UPDATE và DELETE nếu kèm theo câu lệnh WHERE.

**_Nhược điểm_:**
- Khi Sử dụng index thì tốc độ xử lý ghi dữ liệu( UPDATE, DELETE, INSERT) sẽ bị chậm đi. Vì sẽ cập nhật đến index trong table.
- Tốc độc xử lý thì sẽ bị chậm đi cùng tỷ lệ thuận với số lượng index được sử dụng trong bảng.


## **3. Nêu một số kiểu Index trong MySQL cung cấp?**
- MySQL cung cấp 3 kiểu index khác nhau cho data đó là: B-Tree, Hash, R-Tree index

**B-Tree index:**
- thông thường khi chỉ gọi index mà không nêu rõ kiểu thì mặc định sẽ sử dụng loại B-Tree index.