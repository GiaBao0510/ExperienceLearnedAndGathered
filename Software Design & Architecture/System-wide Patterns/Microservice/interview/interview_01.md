# Bộ Câu Hỏi Phỏng Vấn Microservices

**Microservice** là một kỹ thuật phát triển phần mềm — một biến thể của kiến trúc hướng dịch vụ (Service-Oriented Architecture) — cấu trúc một ứng dụng thành tập hợp các dịch vụ nhỏ, được kết nối lỏng lẻo (loosely coupled) với nhau. Trong kiến trúc microservice, mỗi dịch vụ được thiết kế gọn, tập trung vào một chức năng cụ thể, và giao tiếp với nhau qua các giao thức nhẹ (như REST, gRPC).

Việc mô hình hoá ứng dụng dưới dạng các microservice mang lại một số lợi ích:

- Dễ bảo trì và kiểm thử hơn (mỗi service nhỏ, phạm vi rõ ràng).
- Kết nối lỏng lẻo giữa các phần của hệ thống.
- Có thể triển khai (deploy) độc lập từng service.
- Được thiết kế và tổ chức xoay quanh năng lực nghiệp vụ (business capability) của doanh nghiệp.
- Mỗi service có thể được quản lý bởi một team nhỏ.

Danh sách câu hỏi bên dưới được sắp xếp theo mạch từ khái niệm nền tảng đến các chủ đề chuyên sâu hơn, phù hợp để ôn tập trước phỏng vấn.

---

## Câu 1: Sự khác biệt giữa Monolithic, SOA và Microservices?

![So sánh Monolithic, SOA và Microservices](https://media.licdn.com/dms/image/v2/D4D12AQFz0mZqf7Fknw/article-cover_image-shrink_600_2000/article-cover_image-shrink_600_2000/0/1687247100768?e=2147483647&v=beta&t=MBRhqCb5B7_WyX0xBFzUiGPghdro5i7-t8lyMYU6Eis)

**Monolithic**: Giống như một khối lớn duy nhất, trong đó toàn bộ thành phần của ứng dụng (giao diện, logic nghiệp vụ, truy cập dữ liệu) được đóng gói cùng nhau, chạy trong một tiến trình, dùng chung một codebase.

**SOA (Service-Oriented Architecture — Kiến trúc hướng dịch vụ)**: Là một nhóm các dịch vụ tương tác, giao tiếp với nhau. Tùy bản chất giao tiếp, đó có thể là trao đổi dữ liệu đơn giản, hoặc nhiều dịch vụ phối hợp để thực hiện một nghiệp vụ phức tạp hơn. SOA thường dùng một tầng trung gian tập trung gọi là **Enterprise Service Bus (ESB)** để định tuyến và biến đổi thông điệp giữa các dịch vụ.

**Microservices**: Là cách cấu trúc một ứng dụng thành một cụm (cluster) các dịch vụ nhỏ, tự trị, được mô hình hóa xoay quanh một miền nghiệp vụ (domain) cụ thể. Mỗi service có thể được triển khai độc lập, mở rộng độc lập, nhằm phục vụ một mục tiêu nghiệp vụ cụ thể, và giao tiếp với các service khác qua giao thức tiêu chuẩn — thường là các API nhẹ (REST, gRPC) thay vì một ESB tập trung.

**Ví dụ**: Một ứng dụng bán hàng viết theo Monolithic sẽ có một chương trình Go duy nhất xử lý cả đăng ký, giỏ hàng, thanh toán trong một process. Cùng ứng dụng đó viết theo Microservices sẽ tách thành `user-service`, `cart-service`, `payment-service` — mỗi service là một chương trình Go riêng, deploy độc lập.

### Mở rộng

- Modular Monolith — mô hình lai giữa Monolith và Microservices, module hóa rõ ràng bên trong một codebase duy nhất.
- Nguyên tắc "Smart endpoints, dumb pipes" của Martin Fowler — lý do microservices thường tránh dùng ESB tập trung.

### Nguồn trích dẫn

- Martin Fowler — [Microservices Resource Guide](https://martinfowler.com/microservices/)
- Chris Richardson — [microservices.io — Pattern: Microservice architecture](https://microservices.io/patterns/microservices.html)

---

## Câu 2: Các tính năng chính của Microservices?

![Các tính năng của Microservices](https://www.edureka.co/blog/wp-content/uploads/2018/02/Microservices-Features-What-Is-Microservices-edureka.png)

- **Phân tách (Componentization qua service)**: Các dịch vụ trong hệ thống được phân tách rõ ràng. Nhờ đó, toàn bộ ứng dụng dễ xây dựng, thay đổi, và mở rộng hơn.
- **Thành phần hóa**: Mỗi microservice được xem như một thành phần độc lập, có thể dễ dàng thay thế hoặc nâng cấp mà không ảnh hưởng phần còn lại.
- **Quy mô theo nghiệp vụ**: Mỗi microservice tương đối đơn giản, chỉ tập trung phục vụ một năng lực nghiệp vụ cụ thể.
- **Quyền tự chủ của nhóm**: Các nhà phát triển của từng service có thể làm việc độc lập với các team khác, giúp tiến độ dự án nhanh hơn.
- **Phân phối liên tục (Continuous Delivery)**: Cho phép phát hành phần mềm thường xuyên, nhờ hệ thống tự động hóa việc build, kiểm thử và triển khai.
- **Trách nhiệm với sản phẩm, không phải dự án**: Team không xem ứng dụng như một dự án làm xong rồi bàn giao, mà xem nó như một sản phẩm mình chịu trách nhiệm vận hành lâu dài.
- **Quản trị phi tập trung**: Mỗi team có thể tự chọn công cụ, công nghệ phù hợp nhất để giải quyết vấn đề của service mình, thay vì bị ép dùng một stack chung cho toàn hệ thống.
- **Kết hợp tốt với Agile**: Do mỗi service nhỏ và độc lập, có thể tạo tính năng mới nhanh chóng, thử nghiệm, và loại bỏ nếu không hiệu quả mà ít ảnh hưởng tới phần còn lại của hệ thống.

### Mở rộng

- Nine characteristics of microservices theo Martin Fowler (bao gồm cả các đặc điểm ở Câu 3 bên dưới).
- DevOps và Continuous Delivery — nền tảng giúp các đặc điểm trên khả thi trong thực tế.

### Nguồn trích dẫn

- Martin Fowler & James Lewis — [Microservices — a definition of this new architectural term](https://martinfowler.com/articles/microservices.html)

---

## Câu 3: Các đặc điểm cơ bản của thiết kế Microservice?

- **Dựa trên năng lực nghiệp vụ (Organized around Business Capabilities)**: Các dịch vụ được phân chia và tổ chức xoay quanh những gì doanh nghiệp _làm_ (ví dụ: quản lý đơn hàng, thanh toán), thay vì theo tầng kỹ thuật.
- **Sản phẩm, không phải dự án (Products not Projects)**: Một service nên thuộc về một team cụ thể, chịu trách nhiệm cho nó trong suốt vòng đời — từ phát triển đến vận hành — thay vì bàn giao cho team khác sau khi "hoàn thành dự án".
- **Smart endpoints, dumb pipes (Giao tiếp thông minh ở hai đầu, đường truyền đơn giản)**: Thay vì đặt logic xử lý message vào một tầng trung gian tập trung (như ESB trong SOA truyền thống), microservices đặt phần logic ("thông minh") ngay tại chính các service, còn kênh giao tiếp giữa chúng ("đường ống") được giữ đơn giản — thường chỉ là HTTP/REST hoặc một message broker nhẹ, không xử lý logic nghiệp vụ.

**Ví dụ**: Thay vì dựng một ESB để "dịch" và định tuyến message giữa `order-service` và `payment-service`, hai service này tự gọi thẳng nhau qua một REST API đơn giản (`POST /pay`) — phần xử lý nghiệp vụ (validate, tính phí...) nằm trong chính từng service, không nằm ở tầng trung gian.

### Mở rộng

- So sánh chi tiết ESB (SOA) với giao tiếp trực tiếp qua REST/event trong Microservices.
- Nguyên tắc Single Responsibility áp dụng ở cấp độ service.

### Nguồn trích dẫn

- Martin Fowler & James Lewis — [Microservices — a definition of this new architectural term](https://martinfowler.com/articles/microservices.html)

---

## Câu 4: Cohesion và Coupling là gì?

**Coupling (Tính liên kết/phụ thuộc giữa các module)**: Là mức độ một module phụ thuộc vào hoặc phải biết chi tiết bên trong của module khác. Coupling thường được chia theo mức độ:

- **Tightly coupled** (liên kết chặt): các module phụ thuộc nhiều vào nhau, thay đổi ở module này dễ kéo theo phải sửa module kia.
- **Loosely coupled** (liên kết lỏng): các module phụ thuộc rất ít vào nhau, thường thông qua interface hoặc API rõ ràng.
- **Decoupled** (tách rời hoàn toàn): các module gần như không phụ thuộc nhau.

Mục tiêu khi thiết kế hệ thống (đặc biệt là microservices) là đạt **loose coupling** — các service phụ thuộc vào nhau càng ít càng tốt, thường đạt được bằng cách giao tiếp qua interface/API đã định nghĩa rõ, thay vì gọi thẳng vào chi tiết triển khai của nhau.

**Cohesion (Tính gắn kết bên trong một module)**: Là mức độ các phần tử _bên trong cùng một module_ liên quan chặt chẽ với nhau, cùng phục vụ một mục đích rõ ràng. Một module có **cohesion cao** thường thực hiện tốt một chức năng cụ thể mà không cần "vay mượn" logic từ module khác — điều này giúp module dễ hiểu, dễ kiểm thử, và dễ bảo trì hơn.

Nguyên tắc vàng khi thiết kế (bao gồm cả khi chia microservices) là: **high cohesion, loose coupling** — gắn kết cao bên trong mỗi module/service, nhưng phụ thuộc lỏng giữa các module/service với nhau.

![Minh họa Cohesion và Coupling](https://media.geeksforgeeks.org/wp-content/uploads/20200615000022/Untitled273.png)

**Ví dụ**: Một `payment-service` chỉ chứa logic liên quan tới thanh toán (tạo giao dịch, xử lý hoàn tiền) — đây là cohesion cao. Nếu `payment-service` gọi thẳng vào hàm nội bộ của `order-service` để đọc dữ liệu đơn hàng thay vì gọi qua API công khai, đó là dấu hiệu của tight coupling — nên tránh.

### Mở rộng

- Nguyên tắc SOLID, đặc biệt Single Responsibility Principle — liên quan trực tiếp tới cohesion.
- High cohesion & loose coupling khi quyết định ranh giới (boundary) của một microservice — xem thêm khái niệm Bounded Context ở câu hỏi tiếp theo.

### Nguồn trích dẫn

- GeeksforGeeks — [Software Engineering | Coupling and Cohesion](https://www.geeksforgeeks.org/software-engineering-coupling-and-cohesion/)

---

## Câu 5: Bounded Context là gì?

**Bounded Context** là một mẫu (pattern) trung tâm trong **DDD (Domain-Driven Design)**. Đây là một ranh giới rõ ràng, trong đó một domain model cụ thể và ngôn ngữ chung của nhóm (Ubiquitous Language) có ý nghĩa nhất quán. DDD chia một domain model lớn, phức tạp thành nhiều Bounded Context nhỏ hơn để dễ quản lý — mỗi context có model và ngôn ngữ riêng, chỉ có ý nghĩa trong phạm vi của nó, và mối quan hệ giữa các context được xác định rõ ràng qua ranh giới.

Bounded Context khuyến khích cách tiếp cận hướng đối tượng (object-oriented) khi phát triển service: mỗi service gắn liền với một model dữ liệu cụ thể, và cũng là nơi chịu trách nhiệm đảm bảo tính toàn vẹn (integrity) và khả năng thay đổi của model đó theo thời gian.

![Minh họa Bounded Context](https://martinfowler.com/bliki/images/boundedContext/sketch.png)

**Ví dụ**: Trong một hệ thống thương mại điện tử, "Customer" ở Bounded Context Bán hàng (gắn với cơ hội bán hàng, hợp đồng) có thể khác hẳn "Customer" ở Bounded Context Hỗ trợ khách hàng (gắn với ticket, lịch sử liên hệ) — dù cùng một từ, nhưng model và ý nghĩa khác nhau tùy theo ranh giới.

### Mở rộng

- Ubiquitous Language, Context Mapping — các khái niệm đi kèm trong Strategic Design của DDD.
- Mối quan hệ giữa Bounded Context và ranh giới của một microservice.

### Nguồn trích dẫn

- Martin Fowler — [BoundedContext](https://martinfowler.com/bliki/BoundedContext.html)
- Eric Evans — _Domain-Driven Design: Tackling Complexity in the Heart of Software_ (sách gốc giới thiệu khái niệm này)

---

## Câu 6: Kiến trúc tham chiếu của một hệ thống Microservices thường gồm những thành phần nào?


![Kiến trúc tham chiếu Microservices](https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSUUxoOdCRGrJZYceWAw_J9JVM9y_-mm2BS7M5fN3my1gUQTkNckvypLZg&s=10)

- **Clients**: Người dùng khác nhau gửi request từ nhiều loại thiết bị khác nhau (web, mobile, IoT...).
- **Identity Provider**: Xác thực định danh người dùng hoặc client, và cấp token bảo mật cho các request tiếp theo.
- **API Gateway**: Tiếp nhận và xử lý request từ client, đóng vai trò điểm vào duy nhất trước khi chuyển tới các service phía sau.
- **Static Content**: Thành phần phục vụ các nội dung tĩnh của hệ thống (hình ảnh, CSS, JS, file tải về...).
- **Management**: Thành phần giám sát, cân bằng tải giữa các node chạy service, đồng thời phát hiện khi có node gặp lỗi.
- **Service Discovery**: Giúp các service tự động tìm ra địa chỉ của nhau để giao tiếp, thay vì phải cấu hình cứng.
- **CDN (Content Delivery Network)**: Mạng lưới các proxy server và trung tâm dữ liệu đặt gần người dùng, giúp phân phối nội dung (đặc biệt là static content) nhanh hơn.
- **Remote Service**: Cung cấp khả năng truy cập từ xa vào dữ liệu hoặc chức năng nằm trên các máy chủ/thiết bị khác trong mạng.

**Ví dụ**: Khi người dùng mở app mobile, request trước tiên đi qua CDN để tải các nội dung tĩnh, sau đó gọi API Gateway kèm token đã được Identity Provider cấp; Gateway tra cứu Service Discovery để biết địa chỉ instance của service cần gọi, rồi chuyển tiếp request tới đó.

### Mở rộng

- So sánh kiến trúc tham chiếu này với các nhóm "thành phần hạ tầng" ở câu hỏi tiếp theo (Container, IaC, Cloud Infrastructure...) — đây là hai góc nhìn khác nhau: một là sơ đồ luồng xử lý request, một là nhóm công nghệ nền tảng để xây dựng hệ thống.
- API Gateway pattern và Backend for Frontend (BFF) pattern.

### Nguồn trích dẫn

- AWS — [Microservices architecture](https://aws.amazon.com/microservices/)
- Microsoft Learn — [Azure Architecture Center — Microservices](https://learn.microsoft.com/en-us/azure/architecture/microservices/)

---

## Câu 7: Các thành phần/công nghệ nền tảng thường dùng khi xây dựng Microservices?

- **Containers, Clustering và Orchestration**: Đóng gói mỗi service thành container (ví dụ Docker), rồi dùng công cụ orchestration (ví dụ Kubernetes) để quản lý việc chạy, scale, và phục hồi các container đó trên một cluster.
- **IaC (Infrastructure as Code)**: Quản lý hạ tầng (server, network, database...) bằng code/file cấu hình thay vì thao tác thủ công, giúp hạ tầng có thể versioning và tái tạo lại dễ dàng.
- **Cloud Infrastructure**: Hạ tầng cloud (AWS, GCP, Azure...) cung cấp khả năng scale linh hoạt, phù hợp với đặc thù cần mở rộng độc lập của microservices.
- **API Gateway**: Điểm vào duy nhất, xử lý routing, authentication, rate limiting cho các request tới hệ thống (xem chi tiết ở Câu 6).
- **Service Registry / Service Discovery**: Thành phần lưu và cập nhật danh sách địa chỉ các service instance đang hoạt động.

### Mở rộng

- Message Broker (Kafka, RabbitMQ) như lựa chọn giao tiếp bất đồng bộ phổ biến hơn ESB trong microservices hiện đại.
- Service Mesh (Istio, Linkerd) — lớp hạ tầng đảm nhiệm discovery, load balancing, retry giữa các service.

### Nguồn trích dẫn

- Martin Fowler & James Lewis — [Microservices — a definition of this new architectural term](https://martinfowler.com/articles/microservices.html) (phần "Smart endpoints and dumb pipes")

---

## Câu 8: Các công cụ thường dùng cho Microservices?

- **Docker**: Đóng gói mỗi service cùng toàn bộ dependency thành một container, đảm bảo chạy nhất quán giữa các môi trường.
- **WireMock**: Công cụ giả lập (mock) API của các service khác trong lúc test, giúp kiểm thử một service độc lập mà không cần các service phụ thuộc thật sự chạy.
- **Hystrix**: Thư viện của Netflix, dùng để triển khai circuit breaker — cô lập lỗi khi gọi tới một service khác, tránh lỗi lan rộng ra toàn hệ thống.

> **Hystrix hiện đã ngừng phát triển tích cực và chuyển sang chế độ maintenance mode** (theo thông báo chính thức từ Netflix). Các dự án mới thường được khuyến nghị dùng **resilience4j** — thư viện circuit breaker/rate limiter/retry hiện đại hơn, được cộng đồng duy trì tích cực — thay vì Hystrix.

**Ví dụ cấu hình circuit breaker bằng Go** (dùng thư viện `sony/gobreaker`, một lựa chọn phổ biến trong hệ sinh thái Go, tương đương ý tưởng của Hystrix/resilience4j):

```go
package main

import (
	"fmt"
	"time"

	"github.com/sony/gobreaker"
)

func main() {
	settings := gobreaker.Settings{
		Name:        "payment-service-breaker",
		MaxRequests: 3,
		Timeout:     10 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			// Ngắt mạch nếu tỉ lệ lỗi vượt quá 50% trong các lần gọi gần nhất
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return counts.Requests >= 5 && failureRatio >= 0.5
		},
	}

	cb := gobreaker.NewCircuitBreaker(settings)

	_, err := cb.Execute(func() (interface{}, error) {
		// Gọi payment-service ở đây; nếu liên tục lỗi,
		// circuit breaker sẽ tự "mở mạch" và từ chối gọi tiếp
		// trong một khoảng thời gian, thay vì tiếp tục dồn lỗi.
		return callPaymentService()
	})
	if err != nil {
		fmt.Println("Gọi payment-service thất bại hoặc mạch đang mở:", err)
	}
}
```

### Mở rộng

- resilience4j (Java), gobreaker/hystrix-go (Go) — các thư viện circuit breaker theo từng ngôn ngữ.
- Contract Testing (ví dụ Pact) — bổ sung cho WireMock khi cần đảm bảo API giữa các service luôn tương thích.

### Nguồn trích dẫn

- Netflix — [Hystrix — GitHub repository (Status: Maintenance mode)](https://github.com/Netflix/Hystrix)
- Resilience4j — [Tài liệu chính thức](https://resilience4j.readme.io/)
- WireMock — [Trang chủ dự án](https://wiremock.org/)

---

## Câu 9: Ưu nhược điểm của Microservices?

|Ưu điểm|Nhược điểm|
|---|---|
|Module triển khai khép kín và độc lập.|Kiểm thử và giám sát khó khăn hơn do độ phức tạp của kiến trúc phân tán.|
|Các dịch vụ được quản lý độc lập.|Đòi hỏi văn hóa làm việc phù hợp (team tự chủ, chịu trách nhiệm sản phẩm) — nếu tổ chức chưa sẵn sàng, kiến trúc khó phát huy hiệu quả.|
|Có thể triển khai dịch vụ trên nhiều server để cải thiện hiệu suất.|Cần lập kế hoạch kỹ trước khi triển khai (xác định ranh giới service, chiến lược dữ liệu...).|
|Dễ kiểm thử hơn ở cấp từng service, ít phụ thuộc hơn so với monolith.|Độ phức tạp khi phát triển tăng lên (giao tiếp mạng, nhất quán dữ liệu...).|
|Khả năng mở rộng (scale) linh hoạt hơn.|Đòi hỏi một sự chuyển dịch về văn hóa làm việc và tổ chức team.|
|Dễ khoanh vùng khi debug và bảo trì, vì lỗi thường giới hạn trong một service.|Chi phí hạ tầng và vận hành thường cao hơn so với monolith.|
|Giao tiếp rõ ràng hơn giữa nhà phát triển và người làm nghiệp vụ (mỗi service gắn với một năng lực nghiệp vụ cụ thể).|Có thêm rủi ro bảo mật do nhiều điểm giao tiếp qua mạng giữa các service.|
|Mỗi team phát triển có quy mô nhỏ hơn, dễ quản lý.|Vận hành và giám sát hệ thống mạng phức tạp hơn.|

### Mở rộng

- Chi phí vận hành (operational overhead) của microservices so với monolith — nên cân nhắc khi hệ thống còn nhỏ.
- Observability (logging, metrics, tracing) — công cụ giúp giảm bớt khó khăn khi giám sát hệ thống phân tán.

### Nguồn trích dẫn

- Chris Richardson — [microservices.io — Pattern: Microservice architecture](https://microservices.io/patterns/microservices.html)

---

## Câu 10: Các thách thức khi sử dụng Microservices?

Những thách thức khi dùng microservices có thể chia thành hai nhóm: chức năng (functional) và kỹ thuật (technical).

**Thách thức về mặt tổ chức/chức năng:**

- Yêu cầu thiết lập cơ sở hạ tầng khá nặng (container, orchestration, service registry, monitoring...).
- Cần đầu tư đáng kể về công cụ, hạ tầng và nhân sự trước khi triển khai.
- Đòi hỏi lập kế hoạch kỹ lưỡng để kiểm soát chi phí vận hành nhiều thành phần.

**Thách thức về mặt kỹ thuật:**

- Các microservice thường phụ thuộc lẫn nhau để hoàn thành một nghiệp vụ, nên phải giao tiếp với nhau qua mạng — kéo theo độ trễ và khả năng lỗi mạng.
- Đây là một hệ phân tán, nên vốn dĩ có nhiều điểm kết nối (liên kết) hơn so với một ứng dụng chạy trong một tiến trình duy nhất.
- Cần chuẩn bị cho chi phí vận hành cao hơn — nhiều service đồng nghĩa nhiều thứ cần giám sát, backup, nâng cấp.
- Do các microservice thường được viết bằng nhiều công nghệ khác nhau (polyglot), cần đội ngũ có chuyên môn đa dạng để hỗ trợ.
- Số lượng thành phần trong hệ thống nhiều hơn hẳn so với monolith, khiến việc tự động hóa (CI/CD, giám sát...) phức tạp hơn — mỗi thành phần đều cần được build, triển khai và giám sát riêng biệt.
- Quản lý cấu hình nhất quán trên nhiều môi trường (dev, staging, production) cho tất cả các thành phần là một bài toán khó.
- Việc triển khai, gỡ lỗi (debug) và kiểm thử (test) toàn hệ thống end-to-end khó hơn nhiều so với một ứng dụng nguyên khối.

### Mở rộng

- Distributed Tracing (ví dụ OpenTelemetry, Jaeger) — công cụ giúp debug một request đi qua nhiều service.
- Saga Pattern — giải quyết bài toán giao dịch (transaction) trải dài qua nhiều service, một trong những thách thức kỹ thuật lớn nhất.

### Nguồn trích dẫn

- Chris Richardson — [microservices.io — Introduction to Microservices](https://microservices.io/patterns/microservices.html)