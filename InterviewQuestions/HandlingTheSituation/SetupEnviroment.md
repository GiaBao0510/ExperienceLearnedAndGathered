
Khi triển khai hệ thống để phục vụ người dùng thì sẽ có các yêu cầu về phần cứng, phần mềm và môi trường triển khai phải được thiết kế một cách cẩn thận để đảm bảo **tính sẵn sàng cao (high availability), Khả năng mở rộng(scalabitity) và hiệu suất(performance).**

### 1. **Yêu cầu cơ bản cần Setup**

#### *Phần cứng:*
   **Máy chủ (server):**
   - **Load Balancer:** phân phối lưu lượng truy cập giữa các máy chủ ứng dụng.
   - **Application Servers:** Xử lý logic ứng dụng, chịu tải người dùng.
   - **Database Server:** lưu trữ dữ liệu, hỗ trợ trực tuyến.
   - **Cache server:** Tăng tốc độ xử lý dữ liệu tạm thời (Redis, Memcached).
   - **Storage Server:** Lưu trữ tệp lớn (media files, logs).
   - **Search Server:** Xử lý tìm kiếm nhanh (ElasicSearch).
   - **Message Queue Server**: Xử lý luồng công việc không đồ bộ (RabbitMQ, Kafka).
   
   **Cấu hình máy chủ (ước lượng ban đầu):**
   **Network:**
**Request:**
	- server (Window/ Linux/ ).
	- Ram: 8/16/32/64 GB,
	- core i5/i6/i7/...

---
**Install:**
	- SqlServer => Enable Sql Server Authentication, open port 1433 => Remote

Install IIS: