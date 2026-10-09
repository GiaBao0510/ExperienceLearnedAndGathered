# Saga Pattern – Quản lý giao dịch phân tán trong Microservices

**Saga Pattern** là một design pattern (mẫu thiết kế) dùng để quản lý **distributed transaction** (giao dịch phân tán) trong kiến trúc microservices. Nó giúp các service phối hợp với nhau và đạt được **tính nhất quán cuối cùng** (*eventual consistency*), ngay cả khi một service gặp sự cố giữa chừng.

![](https://images.viblo.asia/19213ffb-5146-4b20-ad4b-6ca65e1db742.png)

**Sau bài này bạn sẽ:**

- Hiểu vì sao transaction truyền thống không dùng được trong microservices.
- Giải thích được Saga, compensation (bù đắp) và hai cách triển khai: Choreography và Orchestration.
- Biết các lưu ý quan trọng khi triển khai thật (idempotency, lỗi khi bù đắp, thiếu tính cô lập).

---

## 1. Bối cảnh và vấn đề

### 1.1. Hệ thống monolithic: một transaction là đủ

Trong hệ thống **monolithic** (nguyên khối, toàn bộ chạy trong một ứng dụng và dùng chung một cơ sở dữ liệu), bạn dùng **transaction** để giữ dữ liệu nhất quán:

- Nếu tất cả các bước thành công → **commit** (xác nhận, lưu vĩnh viễn).
- Nếu có bước nào thất bại → **rollback** (hủy, đưa dữ liệu về như trước khi bắt đầu).

Transaction trong cơ sở dữ liệu đảm bảo bốn tính chất **ACID**:

| Chữ cái | Tính chất | Ý nghĩa |
|---|---|---|
| **A** | Atomicity (Nguyên tử) | Hoặc làm hết, hoặc không làm gì cả |
| **C** | Consistency (Nhất quán) | Dữ liệu luôn thỏa mãn các ràng buộc sau transaction |
| **I** | Isolation (Cô lập) | Các transaction chạy đồng thời không "nhìn thấy" dữ liệu dở dang của nhau |
| **D** | Durability (Bền vững) | Đã commit thì không bị mất, kể cả khi hệ thống sập |

**Ví dụ đặt hàng trong hệ thống monolithic:**

- Trừ tiền khách hàng
- Trừ tồn kho sản phẩm
- Ghi nhận đơn hàng

Cả ba thao tác nằm trong một transaction, nên nếu bước nào lỗi → rollback toàn bộ, dữ liệu vẫn nhất quán.

> **Lưu ý:** Việc gửi email xác nhận *không* thể rollback (email đã gửi thì không thu hồi được). Vì vậy thông thường email chỉ được gửi **sau khi** transaction đã commit thành công.

### 1.2. Hệ thống microservices: mỗi service một cơ sở dữ liệu

Với **microservices** (chia ứng dụng thành nhiều service nhỏ, độc lập), mỗi service quản lý cơ sở dữ liệu riêng:

- **Payment Service**: trừ tiền
- **Inventory Service**: trừ tồn kho
- **Notification Service**: gửi email

Không có một transaction ACID nào bao trùm cả ba cơ sở dữ liệu. Nếu một bước thất bại, các bước trước đó **đã commit rồi** và không thể rollback tự động. Tình huống này gọi là **partial failure** (lỗi một phần).

**Ví dụ:** Payment Service đã trừ tiền, nhưng Inventory Service báo hết hàng → khách hàng mất tiền mà không có hàng, dữ liệu giữa các service bị lệch nhau.

### 1.3. Vì sao không dùng Two-Phase Commit (2PC)?

**2PC** là giao thức cho phép nhiều cơ sở dữ liệu cùng commit hoặc cùng rollback. Tuy nhiên, trong microservices nó ít được dùng vì:

- Các service phải **khóa tài nguyên** và chờ nhau, làm giảm hiệu năng và khả năng mở rộng.
- Nếu bộ điều phối (coordinator) hoặc một service bị lỗi, mọi bên có thể bị treo.
- Nhiều hệ thống hiện đại (nhiều NoSQL, message broker) không hỗ trợ 2PC.

Đây là lý do chúng ta cần Saga Pattern: giúp các service phối hợp và duy trì dữ liệu đồng bộ, kể cả khi có lỗi xảy ra.

---

## 2. Saga Pattern là gì?

![](https://assets.200lab.io/cdn-cgi/image/width=960,fit=scale-down,quality=82,format=auto/images/2025/03/04/saga-pattern-la-gi-040419176ba02b381b8c.webp)

Saga Pattern là một design pattern quản lý distributed transaction trong microservices.

Thay vì dùng một transaction lớn bao trùm tất cả (rollback toàn bộ nếu một bước lỗi), **mỗi service tự quản lý transaction cục bộ của mình**. Nếu một bước sau thất bại, hệ thống thực hiện **compensation** (bù đắp) để hoàn tác các bước trước đã commit.

![](https://assets.200lab.io/cdn-cgi/image/width=960,fit=scale-down,quality=82,format=auto/images/2025/03/04/saga-pattern-la-gi-1-b7e857ea25b6cdd0d91b.webp)

**Ví dụ đặt hàng online:**

1. Payment Service: trừ tiền → thành công
2. Inventory Service: trừ tồn kho → lỗi (hết hàng)
3. Notification Service: gửi email → chưa thực hiện

- **Không dùng Saga:** Payment Service đã trừ tiền → khách mất tiền nhưng không có hàng.
- **Dùng Saga:** Inventory Service lỗi → Payment Service hoàn tiền; email chưa gửi nên không gây nhầm lẫn.

Hiểu đơn giản, Saga là **một chuỗi các transaction cục bộ nhỏ và độc lập**:

- Mỗi transaction chạy trong một service riêng, thành công thì kích hoạt (trigger) transaction kế tiếp.
- Nếu một transaction thất bại, Saga chạy các **compensating transaction** (transaction bù đắp) cho những bước đã thành công trước đó, theo thứ tự ngược lại.

> **Ý tưởng cốt lõi:** Mỗi bước tự chịu trách nhiệm và có sẵn cách bù đắp, giúp distributed transaction hoạt động mà không cần khóa toàn hệ thống.

### Hai điểm cần hiểu đúng

**1. Compensation không phải rollback.** Rollback xóa sạch dấu vết như chưa từng xảy ra; compensation là một thao tác nghiệp vụ *mới* để "gỡ" ảnh hưởng của bước cũ. Ví dụ: tiền đã trừ thì ta *hoàn tiền* (vẫn còn lịch sử giao dịch trừ và hoàn), không phải xóa bản ghi trừ tiền. Có những bước không thể bù đắp trọn vẹn (ví dụ email đã gửi), nên cần thiết kế cẩn thận.

**2. Saga chỉ đảm bảo nhất quán cuối cùng.** Trong lúc Saga đang chạy, dữ liệu có thể *tạm thời* chưa nhất quán (ví dụ tiền đã trừ nhưng tồn kho chưa kiểm tra). Khi Saga kết thúc (thành công hoặc đã bù đắp xong), dữ liệu mới nhất quán. Nói theo ACID, Saga có A, C, D ở mức từng service nhưng **không có I (Isolation)**.

---

## 3. Hai cách triển khai Saga Pattern

### 3.1. Choreography Saga (dựa trên sự kiện – Event-Driven)

Không có người chỉ huy. Mỗi service phát ra **event** (sự kiện) khi làm xong hoặc thất bại, các service khác **lắng nghe** event để quyết định làm tiếp hay bù đắp. Event thường được truyền qua một **message broker** (hệ thống trung chuyển thông điệp như Kafka, RabbitMQ).

**Luồng thành công:**

1. Order Service: Tạo đơn hàng → Phát sự kiện `ORDER_CREATED`.
2. Payment Service: Nghe `ORDER_CREATED` → Xử lý thanh toán → Phát `PAYMENT_COMPLETED`.
3. Inventory Service: Nghe `PAYMENT_COMPLETED` → Kiểm tra và giảm tồn kho → Phát `INVENTORY_UPDATED`.
4. Delivery Service: Nghe `INVENTORY_UPDATED` → Lên lịch giao hàng → Phát `DELIVERY_SCHEDULED`.
5. Order Service: Nghe `DELIVERY_SCHEDULED` → Cập nhật trạng thái đơn hàng thành `CONFIRMED`.

**Luồng thất bại (hết hàng):**

1. Order Service: Tạo đơn hàng → Phát `ORDER_CREATED`.
2. Payment Service: Nghe `ORDER_CREATED` → Xử lý thanh toán → Phát `PAYMENT_COMPLETED`.
3. Inventory Service: Nghe `PAYMENT_COMPLETED` → Tồn kho không đủ → Phát `INVENTORY_FAILED`.
4. Payment Service: Nghe `INVENTORY_FAILED` → Hoàn tiền → Phát `PAYMENT_REFUNDED`.
5. Order Service: Nghe `PAYMENT_REFUNDED` → Cập nhật trạng thái đơn hàng thành `CANCELLED`.

**Ưu điểm:**

- Không cần thành phần điều phối trung tâm, các service phối hợp linh hoạt (loose coupling – liên kết lỏng).
- Dễ thêm service mới vào quy trình.

**Nhược điểm:**

- Khó theo dõi trạng thái tổng thể của transaction, vì luồng "rải rác" trong nhiều service.
- Khi có nhiều bước, các service phụ thuộc vòng vào nhau, khó hiểu và khó debug.
- Dễ xảy ra **duplicate event** (event bị gửi lặp) hoặc event đến trễ, cần cơ chế **idempotency** (xem mục 5).

### 3.2. Orchestration Saga (điều phối tập trung)

Một **Orchestrator** (bộ điều phối) đóng vai người chỉ huy: gửi **command** (lệnh) cho từng service, nhận kết quả rồi quyết định bước tiếp theo. Khi có bước thất bại, orchestrator ra lệnh chạy các compensating transaction cho những bước đã làm.

> **Phân biệt:** *Event* nói "chuyện này đã xảy ra" (`PAYMENT_COMPLETED`), còn *command* nói "hãy làm việc này" (`ChargePayment`). Choreography dùng event, Orchestration chủ yếu dùng command.

**Cách hoạt động:**

1. Saga Orchestrator nhận yêu cầu tạo đơn hàng từ người dùng.
2. Orchestrator gọi Order Service tạo đơn hàng với trạng thái `PENDING`.
3. Orchestrator gọi Payment Service xử lý thanh toán.
4. Nếu thanh toán thành công, gọi Inventory Service giảm tồn kho.
5. Nếu giảm tồn kho thành công, gọi Delivery Service lên lịch giao hàng.
6. Cuối cùng, cập nhật trạng thái đơn hàng thành `CONFIRMED`.

Nếu bất kỳ bước nào thất bại (thanh toán không thành công, tồn kho không đủ...), orchestrator chạy các compensating transaction **theo thứ tự ngược lại**.

**Ví dụ:**

- Inventory Service thất bại → Orchestrator gọi Payment Service hoàn tiền.
- Sau đó cập nhật trạng thái đơn hàng từ `PENDING` thành `CANCELLED`.

**Ưu điểm:**

- Dễ quản lý quy trình phức tạp, trạng thái được kiểm soát tập trung.
- Dễ theo dõi, dễ debug, các service không cần biết về nhau.

**Nhược điểm:**

- Orchestrator là điểm tập trung: nếu lỗi hoặc nghẽn sẽ ảnh hưởng nhiều transaction (cần lưu trạng thái Saga bền vững và triển khai nhiều bản sao).
- Cần thêm thành phần trung tâm → tăng độ phức tạp triển khai.
- Vẫn có thể gặp lệnh bị gửi lặp khi retry, nên các service vẫn cần idempotency.

### 3.3. So sánh và cách chọn

| Tiêu chí | Choreography | Orchestration |
|---|---|---|
| Ai điều khiển luồng? | Không ai, các service tự phản ứng với event | Orchestrator trung tâm |
| Cách giao tiếp | Event qua message broker | Command (gọi/gửi lệnh) và phản hồi |
| Độ phức tạp khi ít bước (≤ 3–4) | Thấp, đơn giản | Cao hơn (thêm thành phần) |
| Độ phức tạp khi nhiều bước | Cao, khó theo dõi | Dễ kiểm soát hơn |
| Theo dõi trạng thái | Khó | Dễ |
| Điểm nghẽn/lỗi tập trung | Không | Có (orchestrator) |

**Gợi ý chọn:** Quy trình ngắn, ít service → Choreography. Quy trình dài, nhiều nhánh rẽ, cần theo dõi chặt → Orchestration.

---

## 4. Ví dụ minh họa bằng mã giả

Đoạn mã giả (TypeScript) cho Orchestration Saga: mỗi bước có `action` và `compensate`, nếu lỗi thì bù đắp các bước đã làm theo thứ tự ngược.

```ts
type Step = {
  name: string;
  action: () => Promise<void>;      // Làm việc chính
  compensate: () => Promise<void>;  // Cách "gỡ" nếu bước sau lỗi
};

async function createOrderSaga(orderId: string) {
  const steps: Step[] = [
    { name: "payment",
      action:     () => payment.charge(orderId),
      compensate: () => payment.refund(orderId) },
    { name: "inventory",
      action:     () => inventory.reserve(orderId),
      compensate: () => inventory.release(orderId) },
    { name: "delivery",
      action:     () => delivery.schedule(orderId),
      compensate: () => delivery.cancel(orderId) },
  ];

  const completed: Step[] = [];
  try {
    for (const step of steps) {
      await step.action();
      completed.push(step);          // Ghi nhớ bước đã thành công
    }
    await order.markConfirmed(orderId);       // CONFIRMED
  } catch (err) {
    for (const step of completed.reverse()) { // Bù đắp theo thứ tự ngược
      await retry(() => step.compensate());   // Bù đắp phải thử lại được
    }
    await order.markCancelled(orderId);       // CANCELLED
  }
}
```

Ví dụ **idempotent consumer** (xử lý event lặp mà không sai dữ liệu):

```ts
async function onPaymentCompleted(event: { id: string; orderId: string }) {
  await db.transaction(async (tx) => {
    // Nếu event.id đã xử lý rồi thì bỏ qua
    if (await tx.processedEvents.exists(event.id)) return;
    await tx.inventory.reserve(event.orderId);
    await tx.processedEvents.insert(event.id); // Ghi "đã xử lý" cùng transaction
  });
}
```

> Đây là mã giả để học khái niệm. Mã thật cần xử lý thêm: lưu trạng thái Saga xuống database (để khôi phục khi orchestrator sập), timeout, và giới hạn số lần retry.

---

## 5. Lưu ý quan trọng khi triển khai

1. **Idempotency (tính lặp lại an toàn):** Event/lệnh có thể bị gửi nhiều lần (do retry, lỗi mạng). Mỗi service phải xử lý được việc này mà không sai dữ liệu, thường bằng cách lưu ID của event/lệnh đã xử lý (như ví dụ trên).
2. **Compensation cũng phải idempotent và retry được:** Hoàn tiền hai lần là lỗi nghiêm trọng. Nếu bù đắp thất bại, phải thử lại hoặc chuyển cho con người xử lý (cảnh báo/hàng đợi lỗi).
3. **Thiếu cô lập (isolation):** Do các bước commit từng phần, transaction khác có thể thấy dữ liệu "dở dang" (ví dụ đơn `PENDING`). Cách giảm thiểu thường dùng: đặt trạng thái trung gian như `PENDING` ("semantic lock" – khóa theo nghiệp vụ), và không cho thao tác sửa/hủy khi đơn đang ở trạng thái này.
4. **Sắp xếp thứ tự bước hợp lý:** Đặt các bước **dễ bù đắp** lên trước, các bước **khó hoặc không thể bù đắp** (gửi email, gọi cổng thanh toán ngoài, giao hàng) xuống **cuối**, sau khi các bước còn lại đã thành công.
5. **Ghi log và theo dõi (observability):** Gắn một *correlation ID* (mã định danh chung) cho cả Saga để lần theo dấu qua nhiều service.

---

## 6. Các thuật ngữ kỹ thuật

1. **Transaction:** Chuỗi thao tác trên dữ liệu, đảm bảo ACID. *Ví dụ:* chuyển tiền từ tài khoản A sang B; trừ A thành công nhưng cộng B lỗi → rollback.
2. **Distributed Transaction:** Transaction diễn ra trên nhiều service hoặc cơ sở dữ liệu riêng biệt, cần compensation hoặc eventual consistency. *Ví dụ:* đặt hàng online gồm Payment, Inventory, Notification.
3. **Saga Pattern:** Pattern quản lý distributed transaction bằng chuỗi transaction cục bộ, kèm compensation khi có bước thất bại. *Ví dụ:* Inventory báo hết hàng → Payment hoàn tiền.
4. **Compensation (Bù đắp):** Thao tác nghiệp vụ nhằm gỡ ảnh hưởng của một bước đã commit khi bước khác thất bại. *Ví dụ:* Payment đã trừ tiền nhưng Inventory lỗi → Payment hoàn tiền.
5. **Event (Sự kiện):** Thông điệp bất đồng bộ báo một điều đã xảy ra. *Ví dụ:* Payment gửi `PAYMENT_COMPLETED`, Inventory nghe và trừ tồn kho.
6. **Command (Lệnh):** Thông điệp yêu cầu một service thực hiện việc nào đó. *Ví dụ:* `ChargePayment`.
7. **Orchestrator:** Thành phần trung tâm trong Orchestration Saga, điều phối các bước và kích hoạt compensation khi cần. *Ví dụ:* gửi lệnh trừ tiền → trừ tồn kho → hoàn tiền nếu tồn kho lỗi.
8. **Choreography Saga (Event-Driven Saga):** Cách triển khai Saga mà mỗi service tự phát/lắng nghe event, không có trung tâm. *Ví dụ:* Payment phát `PAYMENT_COMPLETED` → Inventory trừ kho → Inventory phát `INVENTORY_FAILED` → Payment hoàn tiền.
9. **Orchestration Saga:** Cách triển khai Saga với orchestrator trung tâm điều phối. *Ví dụ:* Orchestrator ra lệnh Payment → Inventory → Delivery; bù đắp nếu Inventory lỗi.
10. **Message Broker:** Hệ thống trung chuyển thông điệp giữa các service (Kafka, RabbitMQ...).
11. **Partial Failure (Lỗi một phần):** Một bước thất bại trong khi các bước khác đã commit. *Ví dụ:* trừ tiền xong nhưng hết hàng.
12. **Consistency (Tính nhất quán):** Dữ liệu luôn thỏa các ràng buộc và quy tắc nghiệp vụ. *Ví dụ:* sau khi đặt hàng, số tiền trừ khớp tổng giá đơn, tồn kho giảm đúng số lượng.
13. **Eventual Consistency (Nhất quán cuối cùng):** Dữ liệu có thể tạm thời lệch nhau nhưng cuối cùng sẽ nhất quán. *Ví dụ:* Payment commit trước, Inventory commit sau, rốt cuộc trạng thái tổng thể đúng.
14. **Idempotency (Tính lặp lại an toàn):** Thực hiện một thao tác nhiều lần cho cùng kết quả như một lần. *Ví dụ:* event `PAYMENT_COMPLETED` gửi 2 lần → chỉ trừ tiền 1 lần.

---

## 7. Kết luận

Saga Pattern là giải pháp hiệu quả để quản lý giao dịch phân tán trong microservices. Với hai cách triển khai chính là **Choreography** (dựa trên sự kiện) và **Orchestration** (điều phối tập trung), pattern này giúp xử lý các quy trình nghiệp vụ phức tạp mà vẫn đạt **nhất quán cuối cùng**.

Saga phức tạp hơn transaction ACID truyền thống: phải tự thiết kế compensation, xử lý trùng lặp và chấp nhận thiếu tính cô lập. Đổi lại, hệ thống có khả năng mở rộng tốt, sẵn sàng cao và chịu lỗi hiệu quả hơn so với việc khóa nhiều dịch vụ như 2PC.

---

## 8. Nội dung mở rộng

### 8.1. Phân loại transaction trong Saga

Theo Chris Richardson, các bước trong Saga chia thành ba loại:

- **Compensable transaction:** có thể bù đắp (ví dụ: tạo đơn `PENDING`, giữ chỗ tồn kho).
- **Pivot transaction:** "điểm không quay đầu". Khi bước này thành công thì Saga sẽ chạy đến cùng, không bù đắp ngược nữa (ví dụ: thanh toán đã được cổng thanh toán xác nhận cuối cùng).
- **Retriable transaction:** các bước sau pivot, được đảm bảo thành công cuối cùng nhờ retry (ví dụ: gửi email, cập nhật trạng thái `CONFIRMED`).

### 8.2. Transactional Outbox – tránh mất event

Vấn đề: service vừa cập nhật database vừa gửi event lên broker, nếu một trong hai lỗi thì dữ liệu và event lệch nhau. **Transactional Outbox** giải quyết bằng cách ghi event vào một bảng `outbox` *cùng transaction* với dữ liệu nghiệp vụ, rồi một tiến trình riêng đọc bảng này và đẩy event lên broker (thường kết hợp với CDC – Change Data Capture).

### 8.3. Công cụ và framework hỗ trợ Saga

- **Temporal**, **Camunda**, **Netflix Conductor**: nền tảng workflow/orchestration.
- **Axon Framework** (Java), **MassTransit** và **NServiceBus** (.NET): hỗ trợ saga/state machine.
- **Eventuate Tram Saga**: thư viện Saga cho Java.

### 8.4. Biện pháp giảm thiểu thiếu cô lập

Một số "countermeasure" được mô tả trong tài liệu của Richardson: *semantic lock* (khóa nghiệp vụ bằng trạng thái), *commutative updates* (thao tác cập nhật giao hoán, đổi thứ tự không đổi kết quả), *pessimistic view* (sắp xếp bước để giảm rủi ro đọc dữ liệu sai), *reread value* (đọc lại giá trị trước khi ghi), *version file* và *by value* (chọn cơ chế theo mức rủi ro nghiệp vụ).

### 8.5. Bài tập gợi ý

1. Vẽ sơ đồ Choreography Saga cho nghiệp vụ "đặt vé máy bay + khách sạn + thuê xe", xác định compensation của từng bước.
2. Với nghiệp vụ trên, bước nào nên là pivot transaction? Vì sao?
3. Viết đoạn mã consumer idempotent bằng ngôn ngữ bạn quen thuộc.

---

## Tài liệu tham khảo

1. Garcia-Molina, H., & Salem, K. (1987). *Sagas*. ACM SIGMOD Record, 16(3). Bài báo gốc đề xuất khái niệm Saga.
2. Richardson, C. *Pattern: Saga*. microservices.io: https://microservices.io/patterns/data/saga.html
3. Richardson, C. (2018). *Microservices Patterns*. Manning Publications (Chương 4: Managing transactions with sagas).
4. Microsoft Learn – Azure Architecture Center: *Saga distributed transactions pattern*: https://learn.microsoft.com/en-us/azure/architecture/patterns/saga
5. AWS Prescriptive Guidance: *Saga patterns* (Choreography và Orchestration): https://docs.aws.amazon.com/prescriptive-guidance/latest/modernization-data-persistence/saga-pattern.html
6. Richardson, C. *Pattern: Transactional outbox*: https://microservices.io/patterns/data/transactional-outbox.html
7. Kleppmann, M. (2017). *Designing Data-Intensive Applications*. O'Reilly (phần về distributed transactions và consistency).
8. Nguồn hình ảnh gốc trong bài: Viblo (images.viblo.asia) và 200Lab (assets.200lab.io).
