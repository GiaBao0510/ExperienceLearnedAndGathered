### Loki KHÔNG chỉ là công cụ truy vấn — nó vẫn LƯU TRỮ log

Phân vai đúng phải là:

-   **Grafana** = chỉ là lớp hiển thị/trực quan hóa (visualization layer). Bản thân Grafana **không lưu trữ gì cả**, nó chỉ gửi query sang Loki rồi vẽ kết quả lên.
-   **Loki** = vừa **lưu trữ** log (nén thành chunk, ghi xuống filesystem hoặc object storage như S3/MinIO) **vừa** cung cấp API để truy vấn (LogQL) qua Grafana hoặc gọi trực tiếp.

Điểm khác biệt so với Elasticsearch không phải là "có lưu trữ hay không", mà là **lưu trữ + index cái gì**: Loki chỉ index label, còn nội dung log thô vẫn được lưu (nén lại) trong storage — chỉ là khi query nội dung, nó phải quét trực tiếp chunk đó thay vì tra index như Elasticsearch.

Vậy pipeline đầy đủ là:

```
App ghi log → Promtail đọc/đẩy → Loki NHẬN, NÉN, LƯU vào storage → Grafana query Loki để hiển thị
```

---
### Chu kỳ xóa log cũ (retention)

Không có "chu kỳ mặc định" — hoàn toàn tùy vào dự án cấu hình, và tùy dự án. Cấu hình nằm trong `limits_config` của `loki-config.yaml`:

```yml
limits_config:
  retention_period: 744h   # ví dụ 31 ngày
```
Có thể cấu hình chi tiết hơn theo từng tenant hoặc từng stream (label) bằng `retention_stream` — ví dụ log của service quan trọng giữ 90 ngày, log debug chỉ giữ 7 ngày.

---
### Ai thực hiện xóa — background job của Loki ?

Việc xóa là **hoàn toàn tự động, chạy nền, dựa trên thời gian**, do một component riêng của Loki tên là **Compactor** đảm nhiệm — đúng cái đã liệt kê trong ghi chú của mình (Compactor: "gộp index nhỏ + xóa log theo retention policy").

Cách nó hoạt động:

1.  Compactor định kỳ (config bằng `compaction_interval`, mặc định khoảng vài phút/lần) quét toàn bộ chunk trong storage
2.  So sánh timestamp của chunk với `retention_period`
3.  Chunk nào quá hạn → xóa thẳng khỏi storage (filesystem hoặc S3/MinIO)

**Lưu ý quan trọng**: cơ chế này chỉ chạy khi bạn bật `compactor.retention_enabled: true` — mặc định là `false`. Rất nhiều người setup Loki xong thấy disk phình to mãi vì quên bật flag này.

---
### **Ghi log trực tiếp vào Loki hay qua file rồi Promtail đọc — cách nào hiệu quả hơn?**

Đây là câu hỏi thiết kế quan trọng nhất trong 6 câu. Có 2 kiến trúc, không phải 3 như bạn liệt kê — làm rõ để bạn không nhầm:

**Kiến trúc A — Ghi file, Promtail đọc (cách bạn đang làm)**
```
App → ghi file (http.log, error.log, sql.log...) → Promtail tail file → đẩy lên Loki
```
**Kiến trúc B — Push trực tiếp từ Go lên Loki**
```
App → gọi HTTP API /loki/api/v1/push → Loki (không qua file, không qua Promtail)
```

> Không có "ghi vào Promtail" — Promtail không phải nơi lưu trữ, nó chỉ là agent đọc và chuyển tiếp, không có khái niệm "ghi log vào Promtail"

**So sánh thực tế:**

|Tiêu chí|A (File + Promtail)|B (Push trực tiếp)|
|--|--|--|
|Độ tin cậy khi Loki down|**Cao hơn** — log vẫn nằm trong file, Promtail sẽ đẩy bù (retry) khi Loki online lại|**Rủi ro mất log** nếu app không tự buffer/retry — nếu Loki down đúng lúc app gọi API, log đó mất luôn|
|Độ phức tạp code|Thấp — chỉ cần ghi file bình thường|Cao hơn — phải tự viết logic gửi HTTP, retry, buffer, xử lý lỗi mạng|
|Độ trễ (latency)|Có độ trễ nhỏ (Promtail cần thời gian tail + batch gửi)|Nhanh hơn nếu cần gần real-time|
|Debug local|Dễ — mở file `.log` xem trực tiếp bằng mắt/`tail -f`, không cần Loki chạy|Khó hơn — không có Loki thì không xem được gì|
|Chuẩn phổ biến trong production|**Phổ biến hơn**, đặc biệt với Kubernetes (sidecar pattern hoặc DaemonSet Promtail đọc log container)|Thường chỉ dùng khi cần log gần real-time cực cao hoặc không có agent (serverless/lambda)|

>  **Kết luận:** Với quy mô project cá nhân/portfolio , **Kiến trúc A (file + Promtail) là lựa chọn đúng và hiệu quả hơn** — vừa an toàn hơn (log không mất khi Loki tạm down), vừa dễ debug, vừa đúng với pattern chuẩn khi sau này bạn deploy lên Kubernetes (lúc đó Promtail chạy dưới dạng DaemonSet đọc log của mọi container tự động, bạn không cần sửa code Go).

Về việc tách `http.log`, `error.log`, `sql.log` riêng — cách này **có** nhưng cần cân nhắc: chia file theo loại giúp dễ debug thủ công, nhưng lại làm chúng ta phải cấu hình nhiều `scrape_config` hơn trong Promtail và dễ quên đồng bộ label. Cách phổ biến hơn trong thực tế là **gộp chung 1 file `app.log`**, dùng field `component` hoặc `log_type` trong JSON (`{"log_type":"sql", ...}`) để phân biệt, rồi lọc bằng LogQL (`| json | log_type="sql"`) thay vì tách file vật lý. Ít file hơn = ít điểm cấu hình sai hơn.
Phần gộp `file log` này sẽ nói chi tiết ở phần dưới

---
### Tại sao nên gộp chung các File log khác thành một file duy nhất

#### A. Giữ được "Tính liên tục của ngữ cảnh" (Chronological Context) - _Quan trọng nhất_

Hãy tưởng tượng một HTTP Request đi vào hệ thống:

1.  `[HTTP]` Nhận request POST /api/users.
2.  `[LOGIC]` Validate payload.
3.  `[SQL]` Query database kiểm tra email trùng.
4.  `[ERROR]` Panic: Duplicate key.

Nếu em tách ra `http.log`, `sql.log`, `error.log`, khi có bug xảy ra, em sẽ phải mở 3 file log cùng lúc, cố gắng so khớp **Timestamp (thời gian)** để xem dòng lỗi kia bắt nguồn từ request nào. Điều này là một cơn ác mộng khi debug trên Production. 
👉 **Gộp chung thành 1 luồng (Single Stream):** Chúng ta chỉ cần lọc theo `trace_id` hoặc `request_id`, toàn bộ vòng đời của request đó sẽ hiện ra tuần tự từ đầu đến cuối. Đây là sức mạnh của **Distributed Tracing**.

--
#### B. Tuân thủ chuẩn "12-Factor App" và Cloud-Native

Trong môi trường hiện đại (Docker, Kubernetes), ứng dụng không nên tự quản lý file log.

-   Theo chuẩn **12-Factor App**, ứng dụng chỉ nên ghi log ra **Standard Output (`stdout`)** và **Standard Error (`stderr`)**.
-   Các Orchestrator (như K8s) hoặc Log Shipper (Promtail, FluentBit, Grafana Alloy) sẽ tự động bắt các luồng `stdout/stderr` này và đẩy về Loki. Việc ứng dụng tự tạo ra `sql.log`, `http.log` trong container bị coi là một **Anti-pattern** (thói quen xấu).

---
### NHỮNG TRƯỜNG HỢP NGOẠI LỆ (KHI NÀO THÌ NÊN TÁCH FILE?)

Là một Senior, sẽ không bao giờ nói "tuyệt đối". Vẫn có những trường hợp trong các hệ thống Enterprise (Ngân hàng, Tài chính, Thương mại điện tử lớn) người ta **BẮT BUỘC PHẢI TÁCH** luồng log. để có cái nhìn toàn diện:

#### A. Audit Log (Log Kiểm toán / Bảo mật)

-   **Đặc điểm:** Ghi lại hành động của Admin, User (VD: "User A đã xóa User B", "Thay đổi cấu hình hệ thống").
-   **Lý do tách:** Log này mang tính chất pháp lý, compliance (tuân thủ). Nó cần được đẩy về một hệ thống lưu trữ riêng (như AWS S3 Glacier hoặc Splunk) với cơ chế bảo mật nghiêm ngặt, cấm xóa sửa, và Retention (lưu trữ) lên tới 5-10 năm. Nên vậy chúng ta không thể ném nó chung vào `app.log` rồi để Loki tự động xóa sau 7 ngày được.

#### B. Tối ưu chi phí lưu trữ (Cost Routing)

-   **Đặc điểm:** Hệ thống sinh ra hàng TB log mỗi ngày.
-   **Lý do tách:** Người ta sẽ tách **Debug/Info Log** (chiếm 90% dung lượng) đẩy xuống Object Storage giá rẻ (S3) qua Loki. Nhưng **Error/Panic Log** (chỉ chiếm 1%) lại được đẩy thẳng sang các hệ thống Real-time Alerting đắt tiền như Datadog, Sentry hoặc PagerDuty để bắn cuộc gọi báo thức cho Dev lúc 2 giờ sáng.

#### C. PII Data (Dữ liệu cá nhân nhạy cảm)

-   Đôi khi người ta tách log chứa thông tin thẻ tín dụng, CCCD ra một luồng riêng để đi qua một "Masking Service" (Dịch vụ che giấu thông tin) trước khi lưu xuống ổ cứng, tránh vi phạm luật an ninh mạng.

---
### Nếu thư mục log mất, Promtail không đọc được — log cũ trong Loki có mất không?

**Không mất.** Đây là điểm bạn cần khắc sâu: **file log trên đĩa và dữ liệu trong Loki là hai nơi lưu trữ hoàn toàn tách biệt.**

```
[File .log trên host]  --(Promtail chỉ ĐỌC, không sở hữu)-->  [Storage của Loki]
     (nguồn tạm thời)                                          (nơi lưu trữ thật sự)
```

Promtail đóng vai trò giống như một "người đưa thư" — nó đọc từ file rồi gửi đi, chứ **không lưu trữ gì cả** (ngoại trừ 1 file nhỏ `positions.yaml` để nhớ đã đọc tới đâu, tránh gửi trùng). Một khi log đã được Promtail gửi thành công và Loki xác nhận nhận được, dữ liệu đó **đã nằm trong storage của Loki** (filesystem/S3 riêng của Loki, thường mount ở volume `/loki` trong container) — hoàn toàn độc lập với file `.log` gốc.

Vậy nên:

-   Xóa file `http.log`/`error.log` trên host → log **đã gửi lên Loki trước đó vẫn còn nguyên**, xem trong Grafana bình thường.
-   Nhưng nếu container Promtail bị mất luôn cả `positions.yaml` (ví dụ do xóa volume), và file log được tạo lại từ đầu → Promtail sẽ đọc lại từ đầu file mới, có thể gây **trùng log** (đẩy lại những dòng cũ nếu chúng còn trong file mới) hoặc **bỏ sót** (nếu file mới không chứa log cũ). Đây là lý do một số setup production dùng cơ chế `fingerprint`/inode-tracking để Promtail nhận diện đúng file, tránh 2 tình huống trên.

Điều bạn cần nhớ: **volume của Loki (nơi lưu chunk thật sự) mới là chỗ quan trọng cần backup, không phải file log trên app.** Nếu bạn xóa nhầm volume `loki-data` trong `docker-compose.yaml`, đó mới là lúc bạn thực sự mất log vĩnh viễn.

### 7. Với Microservices — mỗi service ghi log riêng hay ghi tập trung về 1 Loki?

**Ghi tập trung về 1 Loki instance duy nhất (hoặc 1 cluster Loki)** — đây gần như là lý do chính Loki tồn tại. Đừng nhầm giữa "ghi tập trung vào 1 Loki" với "mỗi service dùng chung 1 file log" — hai việc khác nhau:

```
Service A (auth-service)     → ghi file log riêng của nó → Promtail (sidecar/DaemonSet riêng)
Service B (order-service)    → ghi file log riêng của nó → Promtail riêng          } → CÙNG 1 Loki
Service C (notification-svc) → ghi file log riêng của nó → Promtail riêng
```

-   Mỗi service **vẫn ghi log riêng của mình** (file/thư mục riêng, không đụng vào nhau) — tách biệt để không service nào ảnh hưởng service khác khi ghi log.
-   Nhưng **tất cả đều đẩy về cùng 1 Loki** (thường Loki chạy như 1 service trung tâm trong hạ tầng, không phải mỗi microservice tự chạy 1 Loki riêng).
-   Điểm mấu chốt để phân biệt log của service nào: dùng **label** khi Promtail scrape, ví dụ `job="auth-service"`, `job="order-service"`. Trong Kubernetes, việc này gần như tự động — Promtail DaemonSet tự gắn label `pod`, `namespace`, `container` cho từng dòng log dựa vào metadata của Kubernetes, bạn không cần cấu hình tay từng service.

**Vì sao phải tập trung về 1 Loki thay vì mỗi service 1 Loki riêng?**

-   Khi debug 1 request đi qua nhiều service (ví dụ: `order-service` gọi `notification-service` gọi `payment-service`), nếu mỗi service có Loki riêng, bạn phải mở 3 Grafana khác nhau để ghép log lại bằng tay — cực kỳ bất tiện.
-   Tập trung 1 chỗ, bạn chỉ cần lọc theo `trace_id` (cái mình gợi ý thêm ở bài tập Giai đoạn 1) là thấy được toàn bộ hành trình của 1 request xuyên suốt các service, trên cùng 1 màn hình Grafana.

Đây cũng chính là lý do bài tập giai đoạn 3 mình đề xuất (gắn `trace_id` vào log) không phải bài tập phụ — nó là nền tảng bắt buộc khi bạn scale lên kiến trúc microservices thật sự sau này với các service auth/order/notification bạn đang thiết kế.

Điều bạn cần nhớ: **volume của Loki (nơi lưu chunk thật sự) mới là chỗ quan trọng cần backup, không phải file log trên app.** Nếu chúng ta xóa nhầm volume `loki-data` trong `docker-compose.yaml`, đó mới là lúc chúng ta thực sự mất log vĩnh viễn.

---
### **Với Microservices — mỗi service ghi log riêng hay ghi tập trung về 1 Loki?**

**Ghi tập trung về 1 Loki instance duy nhất (hoặc 1 cluster Loki)** — đây gần như là lý do chính Loki tồn tại. Đừng nhầm giữa "ghi tập trung vào 1 Loki" với "mỗi service dùng chung 1 file log" — hai việc khác nhau:

```
Service A (auth-service)     → ghi file log riêng của nó → Promtail (sidecar/DaemonSet riêng)
Service B (order-service)    → ghi file log riêng của nó → Promtail riêng          } → CÙNG 1 Loki
Service C (notification-svc) → ghi file log riêng của nó → Promtail riêng
```
-   Mỗi service **vẫn ghi log riêng của mình** (file/thư mục riêng, không đụng vào nhau) — tách biệt để không service nào ảnh hưởng service khác khi ghi log.
-   Nhưng **tất cả đều đẩy về cùng 1 Loki** (thường Loki chạy như 1 service trung tâm trong hạ tầng, không phải mỗi microservice tự chạy 1 Loki riêng).
-   Điểm mấu chốt để phân biệt log của service nào: dùng **label** khi Promtail scrape, ví dụ `job="auth-service"`, `job="order-service"`. Trong Kubernetes, việc này gần như tự động — Promtail DaemonSet tự gắn label `pod`, `namespace`, `container` cho từng dòng log dựa vào metadata của Kubernetes, bạn không cần cấu hình tay từng service.

**Vì sao phải tập trung về 1 Loki thay vì mỗi service 1 Loki riêng?**
-   Khi debug 1 request đi qua nhiều service (ví dụ: `order-service` gọi `notification-service` gọi `payment-service`), nếu mỗi service có Loki riêng, bạn phải mở 3 Grafana khác nhau để ghép log lại bằng tay — cực kỳ bất tiện.
-   Tập trung 1 chỗ, bạn chỉ cần lọc theo `trace_id` (cái mình gợi ý thêm ở bài tập Giai đoạn 1) là thấy được toàn bộ hành trình của 1 request xuyên suốt các service, trên cùng 1 màn hình Grafana.