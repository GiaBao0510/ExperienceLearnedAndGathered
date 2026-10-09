aga Pattern – một design pattern quản lý distributed transaction trong microservices. Nó giúp các service phối hợp nhịp nhàng, giữ dữ liệu đồng bộ và duy trì tính nhất quán cuối cùng (eventual consistency) ngay cả khi một service gặp sự cố.

![](https://images.viblo.asia/19213ffb-5146-4b20-ad4b-6ca65e1db742.png)

---
## **Bối cảnh và vấn đề**
Trong các hệ thống truyền thống (monolithic), bạn có thể dùng transaction để đảm bảo dữ liệu luôn nhất quán:

Nếu tất cả các bước thành công → commit
Nếu có bước nào thất bại → rollback

**Ví dụ đặt hàng trong hệ thống monolithic:**

- Trừ tiền khách hàng
- Trừ tồn kho sản phẩm
- Gửi email xác nhận
Tất cả nằm trong một transaction, nên nếu bước nào lỗi → rollback toàn bộ, dữ liệu vẫn nhất quán.

Tuy nhiên, với microservices, mỗi bước thường do service riêng quản lý, với cơ sở dữ liệu riêng:

- Payment Service: trừ tiền
- Inventory Service: trừ tồn kho
- Notification Service: gửi email
Nếu một bước thất bại, các bước trước có thể đã commit, dẫn đến dữ liệu không đồng bộ.

Ví dụ: khách hàng bị trừ tiền nhưng hàng không còn, hoặc email xác nhận chưa gửi.

Đây là vấn đề mà Saga Pattern giải quyết: giúp các service trong microservices phối hợp nhịp nhàng và duy trì dữ liệu đồng bộ ngay cả khi có lỗi xảy ra.

---
## **Saga Pattern là gì?**

![](https://assets.200lab.io/cdn-cgi/image/width=960,fit=scale-down,quality=82,format=auto/images/2025/03/04/saga-pattern-la-gi-040419176ba02b381b8c.webp)

Saga Pattern là một design pattern quản lý distributed transaction trong microservices.

Thay vì dùng transaction truyền thống (rollback toàn bộ nếu một bước lỗi), mỗi service tự quản lý transaction riêng, và nếu bước sau thất bại, hệ thống sẽ thực hiện bù đắp (compensation) cho các bước trước.

Trong hệ thống microservices, mỗi service thường có cơ sở dữ liệu riêng. Vì vậy, việc thực hiện một transaction xuyên suốt tất cả các service không thể sử dụng cách làm thông thường như với một cơ sở dữ liệu duy nhất (transaction ACID). Đây là lý do tại sao chúng ta cần Saga Pattern.

![](https://assets.200lab.io/cdn-cgi/image/width=960,fit=scale-down,quality=82,format=auto/images/2025/03/04/saga-pattern-la-gi-1-b7e857ea25b6cdd0d91b.webp)

Ví dụ đặt hàng online:

1. Payment Service: trừ tiền → thành công
2. Inventory Service: trừ tồn kho → lỗi (hết hàng)
3. Notification Service: gửi email → chưa thực hiện
- Không dùng Saga Pattern: Payment Service đã trừ tiền → khách hàng mất tiền nhưng không có hàng
- Dùng Saga Pattern: Inventory Service lỗi → Payment Service hoàn tiền, Email chưa gửi → tránh nhầm lẫn

Hiểu một cách đơn giản, Saga là một chuỗi các transaction nhỏ và độc lập:

- Mỗi transaction sẽ được thực hiện trong một service riêng lẻ, sau đó trigger (kích hoạt) transaction tiếp theo.
- Nếu một transaction nào đó gặp lỗi, Saga sẽ thực hiện transaction hoàn tác (compensating transaction) để khôi phục lại dữ liệu về trạng thái ban đầu, đảm bảo hệ thống không bị sai lệch.
Saga Pattern giúp giữ cho dữ liệu trong hệ thống luôn nhất quán, ngay cả khi xảy ra lỗi ở một hoặc một số service.

> Ý tưởng cốt lõi: Mỗi bước tự chịu trách nhiệm và có cơ chế compensation, giúp các thao tác trong distributed transaction phối hợp mà không phá vỡ toàn bộ hệ thống.

---
## **Hai cách triển khai Saga Pattern**

#### **Event-Driven Saga (Dựa trên sự kiện)**
- Mỗi bước gửi thông điệp (event) khi hoàn thành hoặc thất bại
- Bước tiếp theo lắng nghe event để quyết định thực hiện hay bù đắp

Ví dụ về luồng thành công:

1. Order Service: Tạo đơn hàng → Phát sự kiện ORDER_CREATED.
2. Payment Service: Lắng nghe sự kiện ORDER_CREATED → Xử lý thanh toán → Phát sự kiện PAYMENT_COMPLETED.
3. Inventory Service: Lắng nghe sự kiện PAYMENT_COMPLETED → Kiểm tra và giảm tồn kho → Phát sự kiện INVENTORY_UPDATED.
4. Delivery Service: Lắng nghe sự kiện INVENTORY_UPDATED → Lên lịch giao hàng → Phát sự kiện DELIVERY_SCHEDULED.
5. Order Service: Lắng nghe sự kiện DELIVERY_SCHEDULED → Cập nhật trạng thái đơn hàng thành CONFIRMED.

Ví dụ về luồng thất bại:

1. Order Service: Tạo đơn hàng → Phát sự kiện ORDER_CREATED.
2. Payment Service: Lắng nghe sự kiện ORDER_CREATED → Xử lý thanh toán → Phát sự kiện PAYMENT_COMPLETED.
3. Inventory Service: Lắng nghe sự kiện PAYMENT_COMPLETED → Kiểm tra tồn kho không đủ → Phát sự kiện INVENTORY_FAILED.
4. Payment Service: Lắng nghe sự kiện INVENTORY_FAILED → Hoàn tiền → Phát sự kiện PAYMENT_REFUNDED.
5. Order Service: Lắng nghe sự kiện PAYMENT_REFUNDED → Cập nhật trạng thái đơn hàng thành CANCELLED.

**Ưu điểm:**

- Không cần orchestrator tập trung, các service tự quản lý và phối hợp linh hoạt
- Dễ mở rộng khi thêm service mới vào quy trình


**Nhược điểm:**
- Khó theo dõi trạng thái tổng thể của transaction
- Dễ xảy ra duplicate event hoặc trễ event, cần cơ chế idempotency

#### **Orchestration Saga (Điều phối tập trung)**
- Một orchestrator trung tâm điều phối các bước
- Khi bước nào thất bại, orchestrator ra lệnh rollback các bước trước đó

Cách hoạt động:

1. Saga Orchestrator nhận yêu cầu tạo đơn hàng từ phía người dùng.
2. Orchestrator gọi Order Service để tạo đơn hàng với trạng thái ban đầu là PENDING.
3. Sau khi đơn hàng được tạo, Orchestrator gọi Payment Service để xử lý thanh toán.
4. Nếu thanh toán thành công, Orchestrator tiếp tục gọi Inventory Service để giảm tồn kho.
5. Nếu việc giảm tồn kho thành công, Orchestrator sẽ gọi Delivery Service để lên lịch giao hàng.
6. Cuối cùng, Orchestrator cập nhật trạng thái đơn hàng thành CONFIRMED.
Nếu bất kỳ bước nào trong các service thất bại (ví dụ: thanh toán không thành công hoặc tồn kho không đủ), Orchestrator sẽ thực hiện các bước bù trừ (compensating transactions) theo thứ tự ngược lại.

Ví dụ:
- Nếu Inventory Service thất bại, Orchestrator sẽ gọi Payment Service để hoàn tiền.
- Sau đó, Orchestrator sẽ cập nhật trạng thái đơn hàng từ PENDING thành CANCELLED.

**Ưu điểm:**

- Dễ quản lý các quy trình phức tạp, kiểm soát trạng thái tập trung
- Theo dõi dễ dàng và giảm rủi ro duplicate hoặc missing event

**Nhược điểm:**

- Orchestrator trở thành điểm tập trung, nếu lỗi hoặc nghẽn → ảnh hưởng toàn bộ transaction
- Cần thêm thành phần trung tâm → tăng độ phức tạp triển khai

> Saga Pattern giúp từng thao tác trong distributed transaction độc lập nhưng vẫn phối hợp hiệu quả, đảm bảo dữ liệu đồng bộ và trải nghiệm người dùng tốt.

---
**Các thuật ngữ kỹ thuật**

1. Transaction: Chuỗi thao tác trên dữ liệu, đảm bảo ACID (Atomicity, Consistency, Isolation, Durability). Ví dụ: Chuyển tiền từ tài khoản A sang B trong ngân hàng; nếu trừ A thành công nhưng cộng B lỗi → rollback.

2. Distributed Transaction: Transaction diễn ra trên nhiều service hoặc cơ sở dữ liệu riêng biệt, cần compensation hoặc eventual consistency. Ví dụ: Đặt hàng online: Payment Service trừ tiền, Inventory Service trừ tồn kho, Notification Service gửi email.

3. Saga Pattern: Design pattern quản lý distributed transaction bằng cách thực hiện compensation khi bước tiếp theo thất bại. Ví dụ: Inventory Service báo hết hàng → Payment Service hoàn tiền.

4. Compensation (Bù đắp): Hoàn tác một bước đã commit nếu bước khác thất bại. Ví dụ: Payment Service đã trừ tiền nhưng Inventory Service lỗi → Payment Service hoàn tiền.

5. Event (Sự kiện): Thông điệp bất đồng bộ giữa các service, báo trạng thái transaction. Ví dụ: Payment Service gửi “PaymentSuccess”, Inventory Service lắng nghe và trừ tồn kho.

6. Orchestrator: Thành phần trung tâm trong Orchestration Saga, điều phối các bước và rollback khi cần. Ví dụ: Orchestrator gửi lệnh trừ tiền → trừ tồn kho → rollback nếu cần.

7. Partial Failure (Lỗi một phần): Một bước trong distributed transaction thất bại, các bước khác đã commit. Ví dụ: Payment Service trừ tiền thành công nhưng Inventory Service báo hết hàng.

8. Consistency (Tính nhất quán): Dữ liệu luôn thỏa mãn ràng buộc và quy tắc nghiệp vụ sau transaction. Ví dụ: Sau khi đặt hàng, tổng tiền khách trừ = tổng giá đơn, tồn kho giảm đúng số lượng.

9. Eventual Consistency (Nhất quán cuối cùng): Hệ thống sẽ trở nên nhất quán theo thời gian, không cần ngay lập tức. Ví dụ: Payment Service commit trước, Inventory Service commit sau, cuối cùng trạng thái tổng thể đúng.

10. Idempotency (Tính lặp lại an toàn): Thực hiện thao tác nhiều lần mà không gây sai lệch dữ liệu, tránh duplicate event. Ví dụ: Event “PaymentSuccess” gửi 2 lần → Payment Service chỉ trừ tiền 1 lần.

11. Orchestration Saga: Triển khai Saga Pattern với orchestrator trung tâm điều phối các bước. Ví dụ: Orchestrator ra lệnh Payment → Inventory → Notification; rollback nếu Inventory lỗi.

12. Event-Driven Saga: Triển khai Saga Pattern, mỗi service tự quản lý transaction, phát/lắng nghe event, không cần trung tâm. Ví dụ: Payment gửi “PaymentSuccess” → Inventory trừ tồn kho → Inventory gửi “InventoryFailed” → Payment hoàn tiền.

---
**Kết luận**
Saga Pattern là một giải pháp hiệu quả để quản lý các giao dịch phân tán trong kiến trúc microservices. Với hai cách triển khai chính - Choreography và Orchestration, pattern này mang đến sự linh hoạt trong việc xử lý các quy trình nghiệp vụ phức tạp, đồng thời đảm bảo dữ liệu luôn nhất quán.

Mặc dù phức tạp hơn so với các giao dịch ACID truyền thống, Saga Pattern lại giúp hệ thống phân tán hiện đại trở nên mạnh mẽ hơn nhờ khả năng mở rộng tốt, sẵn sàng cao, và chống chịu lỗi hiệu quả.
