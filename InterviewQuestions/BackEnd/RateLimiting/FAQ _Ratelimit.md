
## **Một số câu hỏi phòng vấn và kèm câu trả lời tương ứng**

1. **Rate Limit và Throttling khác nhau như thế nào?**
Rate Limit thường đề cập đến ngưỡng vượt qua là bị từ chối ngay bằng lỗi 429 trong khi Throttling thường chỉ cơ chế làm chậm request (delay) thay vì từ chối hoàn toàn.

2. **Rate Limit có áp dụng khác nhau cho từng endpoint không?**
Có. Hầu hết các API hiện đại đều áp dụng Rate Limit khác nhau cho từng endpoint dựa trên mức độ tiêu tốn tài nguyên của từng thao tác. Những endpoint đơn giản như truy xuất dữ liệu thường có giới hạn cao hơn, trong khi các endpoint phức tạp như tìm kiếm, phân tích dữ liệu, xử lý AI,… thường có giới hạn thấp hơn để bảo vệ hệ thống khỏi quá tải. 

3. **Làm sao tăng Rate Limit nếu hệ thống cần nhiều hơn mức mặc định?**
Hầu hết nhà cung cấp API có cơ chế nâng Rate Limit bằng cách nâng cấp gói dịch vụ (từ Free lên Pro/Enterprise), liên hệ trực tiếp để yêu cầu tăng quota với lý do hợp lý hoặc tự động tăng theo mức độ sử dụng lịch sử,…

