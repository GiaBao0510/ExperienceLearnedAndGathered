
## **Khái niệm cơ bản về CQRS:**

- **CQRS (Command Query Responsibility Segregation)** đây là mô hình kiến trúc phát triển ứng dụng. Tách biệt **đọc dữ liệu (Query)** và **Ghi dữ liệu (Commnad)** thành hai mô hình đọc lập.
- Thay vì sử dụng cùng một model để đọc và ghi dữ liệu như kiến trúc CRUD truyền thống, thì CRQS giúp **tối ưu hiệu suất, khả năng mở rộng và bảo trì hệ thống**
- **Command**: thay đổi trạng thái hệ thống **(POST/PUT/PATCH/DELETE).**
- **Query**: truy vấn dữ liệu mà không thay đổi trạng thái **(GET).**
- **Tại sao cần dùng CQRS?**
	- Tối ưu hiệu năng: Tách biệt đọc/ghi giúp tối ưu hóa từng phần.
	- Dễ bảo trì: Code đọc và ghi độc lập, dễ quản lý.
	- phù hợp cho hệ thống phức tạp: khi yêu cầu đọc và ghi khác biệt lớn.
	- **Dễ mở rộng với Event Sourcing** (nếu kết hớp với Event Sourcing, có thể giúp dễ đang lưu lại và thay đổi theo thời gian).

![](data:image/png;base64,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)

---
## **Kiến trúc CQRS**

- **Command Side:**
	- Xử lý các thao tác ghi dữ liệu ==**(CREATE/UPDATE/DELETE)** / **(POST/PUT/PATCH/DELETE)==.**
	- Thường sử dụng **Domain Model (DDD - Domain-Driven Design)** để bảo đảm tính toàn vẹn dữ liệu.
	- Thường ghi váo **database gốc.**
	- Lưu trữ dữ liệu vào **Write Database (thông thường là SQL).**
- **Query Side:**
	- Xử lý thao tác đọc dữ liệu **(GET)/(READ)** một cách tối ưu
	- Sử dụng Denormalized Data Model để tối ưu truy vấn 
	- Có thể sử dụng **denormalization** (dữ liệu không cần chuẩn hóa để đọc nhanh hơn).
	- Có thể lưu trữ dữ liệu dưới dạng **cache**  hoặc lưu vào **Read-Optimized Database** (thường là NoSQL hoặc SQL với cấu trúc đơn giản). *Ví dụ:* Elasticsearch, Redis, MongoDB.
- **Event Sourcing (tùy chọn):**
	- Lưu trữu các sự kiện *(event)* thay vì trạng thái hiện tại.
	- Giúp tái tạo trạng thái hệ thống từ ccacs sự kiện đã xảy ra.

![](https://tech.cybozu.vn/static/8f93926100f8dd14b0156a7a5254ec40/d00b9/codebase-cqrs-pattern.webp)

💡Có thể đồng bộ dữ liệu giữa Command Model và Query Model bằng Event hoặc Message Queue **(Kafka, RabbitMQ, Azure Service Bus).**

---
## **Ưu nhược điểm của CQRS**

**Ưu điểm:**
- ***Tối ưu hiệu năng:***
	- Tách biệt đọc/ghi giúp tối ưu hóa từng phần riêng biệt.
	- Query Side có thể sử dụng caching hoặc denormalized data để tăng tốc truy vấn
- ***Dễ mở rộng:***
	- Có thể scale riêng biệt Command Side và Query Side
- ***Dễ bảo trì:***
	-  Code đọc và ghi độc lập, giảm sự phức tạp.

**Nhược điểm:**
- ***Độ phức tạp:***
	- Cần quản lý hai mô hnhf dữ liệu riêng biệt
	- Đồng bộ dữ liệu giữa Command Side và Query Side (Nếu sử dụng ở Event Sourcing).
- ***Chi phí phát triển:***
	- Cần thời gina và nỗ lực để triển khai đúng cách.

---
## **Khi nào nên dùng CQRS?**

- **Hệ thống phức tạp:**
	- Khi yêu cầu đọc và ghi khác biệt lớn.
	- Khi tối ưu hiệu năng cho các truy vấn phức tạp
- **Hệ thống cần Scalability:**
	- Khi cần scale riêng biệt đọc và ghi
- **Hệ thống sử dụng Event Sourcing:**
	- Khi cần lưu trữ lịch sử thay đổi trạng thái


---
## **Ví dụ đơn giản về CQRS:**

- **Command Side:**
	- Xử lý tạo mới đơn hàng (CreateOrderCommand).
	- Lưu đơn hàng vào Write Database.
- **Query Side:**
	- truy vấn danh sách đơn hàng (GetOrderQuery)
	- Lấy dữ liệu từ Read database (đã được denormalized).

---
## **Tối ưu hóa Code Base Read Side:**

![](https://tech.cybozu.vn/static/8a097f15e722ab3c613914154fd24988/d00b9/codebase-optimize-2.webp)
![]()

---
## **Sau khi Scale thì sẽ có Deign sau:**

![](https://tech.cybozu.vn/static/28d666991f143611e13b981b3251ac5c/d00b9/cqrs-distributed-database.webp)
![]()

---
## ***Tham khảo:***

1. [CQRS - Thiết kế hệ thống chịu tải lớn và dễ bảo trì](https://tech.cybozu.vn/cqrs-thiet-ke-he-thong-chiu-tai-lon-va-de-bao-tri-99f4b/)
2. [Tìm hiểu Command and Query Responsibility Segregation pattern](https://topdev.vn/blog/cqrs-pattern-la-gi-vi-du-de-hieu-ve-cqrs-pattern/)
3. [CQRS pattern là gì? Ví dụ dễ hiểu về CQRS Pattern](https://lptech.asia/kien-thuc/cqrs-pattern-la-gi-tim-hieu-ve-thiet-ke-cho-cac-kieu-kien-truc)