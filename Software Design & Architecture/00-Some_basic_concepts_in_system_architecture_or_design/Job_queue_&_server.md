Trên thực tế, không phải tất cả request đều cần hệ thống thực hiện đầy đủ rồi trả về. Rất nhiều trường hợp, hệ thống cần đưa vào hàng đợi (Queue), dưới dạng FIFO (First-in-first-out), rồi trả về client nhanh nhất có thể. Hàng đợi này sẽ được một xử lý thông qua các Job/Server khác để tránh overload các service chính.

VD: Đơn hàng tạo thành công, hệ thống cần gởi email thông báo, gởi notification cho các bên có liên quan và một loạt các nghiệp vụ khác. Nếu để client phải đợi hệ thống xong hết các công việc này sẽ không hợp lý, resource tại servers tiếp nhận request cũng tiêu tốn rất lớn.


---
### Trích dẫn:

[Web/System Architecture 101 – Kiến trúc web/hệ thống cơ bản · 200Lab](https://200lab.io/blog/kien-truc-web-he-thong-cho-nguoi-moi-bat-dau)