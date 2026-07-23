## **Nghĩa 1 — Cardinality trong ERD (Relationship Cardinality)**

Trong cơ sở dữ liệu, Cardinality (độ phân chia hoặc tính đa dạng dữ liệu) ==mô tả mối quan hệ số học giữa các hàng của hai bảng (ví dụ: 1:1, 1:N, N:M) hoặc số lượng các giá trị duy nhất (distinct) trong một cột.== Khái niệm này có ý nghĩa cốt lõi trong việc thiết kế cấu trúc dữ liệu, tối ưu hóa truy vấn SQL và lập chỉ mục (indexing).

#### **Phân loại trong Mối quan hệ giữa các bảng (Data Modeling)**

Cardinality thể hiện số lượng bản ghi ở bảng A có thể liên kết với bao nhiêu bản ghi ở bảng B:
- **Một-Một (One-to-One - 1:1):** Mỗi bản ghi trong bảng A chỉ liên kết với đúng một bản ghi trong bảng B và ngược lại (Ví dụ: Một nhân viên chỉ có một hồ sơ bảo hiểm).
- **Một-Nhiều (One-to-Many - 1:N):** Một bản ghi trong bảng A có thể liên kết với nhiều bản ghi trong bảng B, nhưng bảng B chỉ liên kết với một bản ghi của bảng A (Ví dụ: Một Khách hàng có thể đặt nhiều Đơn hàng, nhưng một Đơn hàng chỉ thuộc về một Khách hàng).
- **Nhiều-Nhiều (Many-to-Many - N:M):** Nhiều bản ghi trong bảng A liên kết với nhiều bản ghi trong bảng B (Ví dụ: Một sinh viên học nhiều môn học, một môn học có nhiều sinh viên). Kiểu này thường cần một bảng trung gian để triển khai trong cơ sở dữ liệu.

---
## **Nghĩa 2 — Cardinality trong Query Optimization (Row Cardinality)**

Là **số lượng giá trị khác nhau (distinct values) **trong 1 cột, so với tổng số dòng của bảng — dùng để đánh giá "độ chọn lọc" (selectivity) của cột đó, ảnh hưởng trực tiếp đến việc query optimizer có chọn dùng Index hay không (liên hệ trực tiếp Phần 8 vừa học).

Phân loại trong Phân tích và Tối ưu cột (Column Cardinality)
- **High Cardinality (Độ đa dạng cao):** Cột chứa phần lớn các giá trị độc lập và duy nhất (ví dụ: cột Mã khách hàng, Số điện thoại). Thường áp dụng hiệu quả để tạo các Index giúp tăng tốc độ tìm kiếm.
- **Low Cardinality (Độ đa dạng thấp):** Cột chứa nhiều giá trị lặp đi lặp lại hoặc cố định (ví dụ: cột Giới tính, Trạng thái đơn hàng)

```sql
-- Cột "status" có Cardinality THẤP: chỉ vài giá trị khác nhau (VD: 3 giá trị) trên hàng triệu dòng
-- → Index trên "status" thường KÉM hiệu quả, vì mỗi giá trị match rất nhiều dòng
SELECT COUNT(DISTINCT status) FROM orders;  -- VD: ra 3

-- Cột "email" có Cardinality CAO: gần như mỗi dòng 1 giá trị khác nhau
-- → Index trên "email" RẤT hiệu quả, vì mỗi giá trị chỉ match 1 (hoặc rất ít) dòng
SELECT COUNT(DISTINCT email) FROM customers;  -- VD: ra gần bằng tổng số dòng
```

**Công thức tính Selectivity:**

> Selectivity = số giá trị distinct / tổng số dòng

Selectivity càng gần 1 (100%) → cột càng "chọn lọc tốt" → index càng hiệu quả. Selectivity càng gần 0 → index càng kém hiệu quả (đây chính là 1 phần lý do "Khi nào Index KHÔNG được dùng" mà bạn đã học ở Phần 8, mục 8.7c — nếu 1 giá trị match >20-30% bảng, optimizer thường bỏ qua index).

> Đây là lý do vì sao khi phỏng vấn, biết phân biệt 2 nghĩa của "Cardinality" thể hiện bạn hiểu sâu cả về thiết kế (ERD) lẫn vận hành/tối ưu (query performance) — 2 khía cạnh khác hẳn nhau của cùng 1 từ.