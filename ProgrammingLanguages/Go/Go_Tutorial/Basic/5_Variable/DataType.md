Trong Go, **kiểu dữ liệu** xác định loại giá trị mà một biến có thể lưu trữ và các thao tác được phép thực hiện trên biến đó. Go là ngôn ngữ **statically typed** (kiểu tĩnh), nghĩa là kiểu dữ liệu của biến được xác định tại thời điểm biên dịch và không thay đổi khi chạy chương trình.

Go chia kiểu dữ liệu thành bốn nhóm chính:
1. **[Basic type (Kiểu dữ liệu cơ bản)](obsidian://open?vault=CuuAmChanKinh&file=Go%2FGo%20Tutorial%2F5_Variable%2FBasic-type):** Bao gồm số, chuỗi, và boolean.
2. **[Aggregate type (Kiểu tổng hợp)](obsidian://open?vault=CuuAmChanKinh&file=Go%2FGo%20Tutorial%2F5_Variable%2FAggregate-type):** Bao gồm mảng và struct.
3. **[Reference type (Kiểu tham chiếu)](obsidian://open?vault=CuuAmChanKinh&file=Go%2FGo%20Tutorial%2F5_Variable%2FReference-type):** Bao gồm con trỏ, slice, map, hàm, và channel.
4. **[Interface type (Kiểu giao diện)](obsidian://open?vault=CuuAmChanKinh&file=Go%2FGo%20Tutorial%2F5_Variable%2FInterface-type)**: Định nghĩa hành vi cho các kiểu dữ liệu.

---
### **Dưới đây là giải thích sơ lượt về từng nhóm, kèm ví dụ minh họa**

Các kiểu cơ bản là những kiểu dữ liệu đơn giản, dùng để lưu trữ số, chuỗi, hoặc giá trị logic.

- **Numbers (Số)**:
    
    - **Integer** (int, int8, int16, int32, int64, uint, v.v.): Lưu số nguyên.
    - **Float** (float32, float64): Lưu số thực.
    - **Complex** (complex64, complex128): Lưu số phức (ít dùng).
        
- **String** (string): Lưu chuỗi ký tự.
- **Boolean** (bool): Lưu giá trị true hoặc false