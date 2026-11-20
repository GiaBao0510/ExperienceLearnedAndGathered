Khi hệ thống đạt một ngưỡng tải nhất định, hệ thống thường nên được chia nhỏ ra thành các service riêng lẻ.

Các service này đóng một vai trò nhất định: Account Service, Product Service, Order Service,… Mục đích vẫn là chia để trị tốt hơn, tránh dồn hết vào một server, 1 service bị nghẽn sẽ ảnh hưởng tới các toàn bộ service khác.Ở đây các bạn có thể thấy rằng Web Server chỉ còn làm nhiệm vụ của nó là render HTML thôi, các service bên dưới mới thực sự xử lý thông tin. Những service này có thể độc lập hoặc có giao tiếp với nhau mà sau này kiến trúc “tiến hoá” lên Micro Services rất nổi tiếng.


---
### Trích dẫn:

[Web/System Architecture 101 – Kiến trúc web/hệ thống cơ bản · 200Lab](https://200lab.io/blog/kien-truc-web-he-thong-cho-nguoi-moi-bat-dau)