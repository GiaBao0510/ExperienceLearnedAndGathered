# Bộ câu hỏi phỏng vấn Backend
---

# NHÓM A: NỀN TẢNG HỆ THỐNG BACKEND

## Câu 1: HTTP và HTTPS có gì khác nhau?

**Câu trả lời:**

- **HTTP (HyperText Transfer Protocol):** Giao thức truyền tải dữ liệu giữa client và server dưới dạng văn bản thuần (plain text), không mã hóa. Dữ liệu truyền qua HTTP có thể bị bên thứ ba nghe lén hoặc chỉnh sửa giữa đường (tấn công Man-in-the-Middle) mà client/server không hề hay biết. Cổng mặc định: 80.
- **HTTPS (HTTP Secure):** Là HTTP được bọc thêm một lớp mã hóa TLS/SSL (Transport Layer Security). Trước khi truyền dữ liệu thực sự, client và server thực hiện một bước "bắt tay" (TLS Handshake) để thống nhất khóa mã hóa và xác minh server thông qua chứng chỉ số (SSL Certificate) do một tổ chức uy tín (Certificate Authority) cấp. Nhờ đó, HTTPS đảm bảo 3 yếu tố: dữ liệu được mã hóa (Confidentiality), không bị chỉnh sửa giữa đường (Integrity), và xác minh đúng server mình đang kết nối (Authentication). Cổng mặc định: 443.

**Ví dụ:** Khi truy cập một trang ngân hàng trực tuyến, trình duyệt hiển thị biểu tượng ổ khóa cho biết kết nối dùng HTTPS — thông tin đăng nhập và số thẻ được mã hóa trong quá trình truyền, kẻ tấn công chặn được gói tin cũng không đọc được nội dung. Ngược lại, nếu trang web chỉ dùng HTTP, mật khẩu gõ vào có thể bị đọc trực tiếp dưới dạng văn bản nếu bị chặn trên cùng mạng Wi-Fi công cộng.

**Mở rộng:**

- Tìm hiểu thêm về HSTS (HTTP Strict Transport Security) — cơ chế buộc trình duyệt luôn dùng HTTPS cho một domain, không cho phép quay lại HTTP kể cả khi bị tấn công downgrade.

**Nguồn trích dẫn:**

- MDN Web Docs – "An overview of HTTP": https://developer.mozilla.org/en-US/docs/Web/HTTP/Overview
- Cloudflare Learning Center – "What is HTTPS?": https://www.cloudflare.com/learning/ssl/what-is-https/

---

## Câu 2: Bạn hiểu như thế nào về Performance và Scalability?

**Câu trả lời:**

- **Performance (Hiệu năng):** Đo lường mức độ "nhanh" của hệ thống ở một điều kiện tải cụ thể, thường qua các chỉ số như **Latency** (thời gian phản hồi một request) và **Throughput** (số lượng request xử lý được trong một đơn vị thời gian).
- **Scalability (Khả năng mở rộng):** Đo lường khả năng hệ thống **duy trì** hiệu năng tốt (hoặc suy giảm ở mức chấp nhận được) khi tải tăng lên nhiều lần, thông qua việc bổ sung thêm tài nguyên. Có hai hướng mở rộng chính: **Vertical Scaling (Scale Up)** — nâng cấp tài nguyên của một máy chủ duy nhất (thêm CPU, RAM); và **Horizontal Scaling (Scale Out)** — thêm nhiều máy chủ chạy song song và phân phối tải giữa chúng (thường cần kết hợp Load Balancer — xem Câu 4).

Hai khái niệm này độc lập với nhau: một hệ thống có thể có hiệu năng (Performance) rất tốt ở tải thấp nhưng lại kém khả năng mở rộng (Scalability) nếu kiến trúc không cho phép dễ dàng thêm tài nguyên khi tải tăng vọt.

**Ví dụ:** Một API xử lý 100 request/giây với độ trễ trung bình 50ms (Performance tốt). Tuy nhiên, nếu kiến trúc hệ thống lưu trạng thái (session) chỉ trên một server duy nhất (không chia sẻ được), khi traffic tăng lên 10.000 request/giây, hệ thống không thể đơn giản thêm server để chia tải — đây là dấu hiệu của Scalability kém, dù Performance ban đầu không hề tệ.

**Mở rộng:**

- Tìm hiểu thêm các chỉ số đo hiệu năng quan trọng khác: P95/P99 Latency (độ trễ của 95%/99% request, phản ánh trải nghiệm thực tế tốt hơn độ trễ trung bình).

**Nguồn trích dẫn:**

- AWS – "What is Scalability?": https://aws.amazon.com/what-is/scalability/

---

## Câu 3: Caching là gì? Tầm quan trọng của Caching trong lập trình Backend?

**Câu trả lời:** Caching là kỹ thuật lưu tạm kết quả của một phép tính toán hoặc một truy vấn dữ liệu tốn kém (chậm) vào một nơi lưu trữ có tốc độ truy xuất nhanh hơn nhiều (thường là bộ nhớ RAM, ví dụ qua Redis/Memcached), để những lần truy cập sau có thể lấy ngay kết quả đã lưu thay vì phải tính toán hoặc truy vấn lại từ đầu.

**Tầm quan trọng:**

- **Giảm độ trễ phản hồi:** Đọc dữ liệu từ RAM nhanh hơn đọc từ Database (đĩa cứng/SSD) hàng chục đến hàng trăm lần.
- **Giảm tải cho Database:** Những dữ liệu được truy cập lặp lại nhiều lần (ví dụ thông tin sản phẩm phổ biến) không cần phải truy vấn DB mỗi lần, giúp DB "rảnh" để xử lý các truy vấn quan trọng khác.
- **Tăng khả năng chịu tải (Scalability):** Hệ thống có thể phục vụ nhiều người dùng hơn với cùng một lượng tài nguyên DB.

**Ví dụ:** Trang chủ một sàn thương mại điện tử hiển thị danh sách "Sản phẩm bán chạy" — dữ liệu này không thay đổi liên tục trong ngày, nên thay vì truy vấn DB mỗi lần có người truy cập trang chủ, hệ thống tính toán danh sách này một lần, lưu vào Redis với thời gian sống (TTL) 10 phút; mọi request trong 10 phút đó đều được phục vụ trực tiếp từ Redis.

**Mở rộng:**

- Tìm hiểu các chiến lược caching phổ biến: Cache-aside, Write-through, Write-back, và bài toán Cache Invalidation (biết khi nào cần xóa/cập nhật cache cũ).

**Nguồn trích dẫn:**

- AWS – "What is Caching?": https://aws.amazon.com/caching/

---

## Câu 4: Load Balancer là gì? Khi nào nên sử dụng Load Balancer?

**Câu trả lời:** Load Balancer (Bộ cân bằng tải) là một thành phần đứng trước nhiều server backend, có nhiệm vụ phân phối lưu lượng truy cập (traffic) đến đồng đều giữa các server đó, thay vì để toàn bộ traffic dồn vào một server duy nhất. Nhờ vậy, hệ thống đạt được hai lợi ích chính: **tăng hiệu năng tổng thể** (traffic được chia đều, không server nào bị quá tải) và **tăng tính sẵn sàng cao (High Availability)** — nếu một server gặp sự cố, Load Balancer tự động phát hiện và ngừng chuyển traffic đến server đó, chuyển hướng sang các server còn hoạt động.

**Các thuật toán phân phối phổ biến:** Round Robin (lần lượt xoay vòng từng server), Least Connections (ưu tiên server đang có ít kết nối nhất), IP Hash (cùng một client luôn được route đến cùng một server, hữu ích khi cần duy trì session).

**Khi nào nên sử dụng:** Khi hệ thống đã mở rộng theo chiều ngang (Horizontal Scaling — xem Câu 2), tức có từ hai server backend trở lên chạy song song, và cần một cơ chế điều phối traffic hợp lý giữa chúng.

**Ví dụ:** Một hệ thống API chạy 3 instance giống nhau trên 3 server khác nhau. Load Balancer (như Nginx, AWS ELB) đứng trước, nhận toàn bộ request từ client và phân chia đều cho 3 server này; nếu server thứ 2 bị crash, Load Balancer tự động ngừng gửi request đến server đó cho tới khi nó khôi phục trở lại, người dùng gần như không nhận ra sự cố đã xảy ra.

**Mở rộng:**

- Tìm hiểu sự khác nhau giữa Load Balancer tầng L4 (hoạt động ở tầng TCP/UDP, chỉ dựa vào IP/Port, tốc độ xử lý rất nhanh) và tầng L7 (hoạt động ở tầng ứng dụng HTTP, có thể định tuyến dựa trên nội dung request như URL hay header, linh hoạt hơn nhưng tốn tài nguyên xử lý hơn).

**Nguồn trích dẫn:**

- Cloudflare Learning Center – "What is load balancing?": https://www.cloudflare.com/learning/performance/what-is-load-balancing/

---

## Câu 5: Dependency Injection là gì?

**Câu trả lời:** Dependency Injection (DI) là một kỹ thuật thiết kế trong đó một đối tượng (object) **nhận các đối tượng phụ thuộc (dependency)** của nó **từ bên ngoài truyền vào** (thường qua constructor, hàm khởi tạo, hoặc một tham số), thay vì tự mình khởi tạo (new/khởi tạo trực tiếp) các dependency đó bên trong chính nó. Đây là cách áp dụng cụ thể của nguyên lý Dependency Inversion (chữ "D" cuối cùng trong SOLID) — các module cấp cao không nên phụ thuộc trực tiếp vào chi tiết triển khai của module cấp thấp, mà nên phụ thuộc vào abstraction (interface).

**Lợi ích chính:**

- **Giảm coupling (sự phụ thuộc chặt):** Đối tượng không cần biết cách khởi tạo dependency của mình, chỉ cần biết dependency đó tuân theo interface nào.
- **Dễ viết Unit Test:** Có thể dễ dàng thay thế dependency thật bằng một phiên bản giả lập (mock/stub) khi test, vì dependency được truyền từ bên ngoài vào.
- **Dễ thay đổi, mở rộng:** Muốn đổi một thành phần (ví dụ đổi từ MySQL sang PostgreSQL), chỉ cần thay đối tượng được "tiêm" vào mà không cần sửa code bên trong class sử dụng nó.

**Ví dụ (minh họa bằng Go):**

```go
// KHÔNG dùng DI: UserService tự khởi tạo kết nối DB bên trong,
// khó thay thế hoặc mock khi viết test
type UserService struct{}
func NewUserService() *UserService {
    db := ConnectToMySQL() // phụ thuộc cứng vào MySQL thật
    return &UserService{db: db}
}

// DÙNG DI: dependency (db) được truyền từ bên ngoài vào qua constructor
type UserService struct {
    db Database // là một interface, không phải implementation cụ thể
}
func NewUserService(db Database) *UserService {
    return &UserService{db: db}
}
// Khi test, có thể truyền vào một mock Database thay vì kết nối MySQL thật
```

**Mở rộng:**

- Tìm hiểu về DI Container/Framework (ví dụ Wire, Google Fx cho Go; Spring cho Java) — công cụ tự động hóa việc khởi tạo và "tiêm" dependency, giúp giảm code soạn sẵn (boilerplate) khi hệ thống có nhiều tầng dependency lồng nhau.

**Nguồn trích dẫn:**

- Martin Fowler – "Inversion of Control Containers and the Dependency Injection pattern": https://martinfowler.com/articles/injection.html

---

# NHÓM B: DOCKER & TRIỂN KHAI

## Câu 6: Cho biết các tính năng chính của Docker

**Câu trả lời:** Docker là một nền tảng containerization, cho phép đóng gói ứng dụng cùng toàn bộ môi trường cần thiết (thư viện, biến môi trường, runtime) thành một đơn vị độc lập gọi là container. Các tính năng chính:

- **Containerization:** Đóng gói ứng dụng và dependency vào một container duy nhất, đảm bảo ứng dụng chạy nhất quán trên mọi môi trường (máy dev, staging, production) — giải quyết triệt để vấn đề kinh điển "chạy được trên máy tôi mà không chạy được trên máy khác".
- **Isolation (Cô lập) nhẹ:** Các container chạy cô lập với nhau (namespace, cgroup riêng) nhưng vẫn dùng chung nhân hệ điều hành (OS kernel) của máy host, giúp container khởi động gần như tức thời và tiêu tốn ít tài nguyên hơn nhiều so với máy ảo (Virtual Machine) truyền thống — vốn phải mô phỏng toàn bộ một hệ điều hành riêng.
- **Portability (Tính di động):** Một Image Docker đã build có thể chạy giống hệt nhau trên bất kỳ máy nào đã cài Docker, không phụ thuộc hệ điều hành host.
- **Layered Filesystem:** Mỗi dòng lệnh trong Dockerfile tạo ra một layer riêng; Docker tận dụng cache các layer không thay đổi giữa các lần build để tăng tốc độ build đáng kể.
- **Hệ sinh thái chia sẻ (Docker Hub):** Cho phép lưu trữ, tìm kiếm và chia sẻ Image công khai hoặc riêng tư.

**Ví dụ:** Một backend viết bằng Go yêu cầu đúng Go 1.22 và một số thư viện hệ thống cụ thể. Thay vì yêu cầu mọi máy dev và server production phải cài đặt thủ công đúng phiên bản, team chỉ cần đóng gói toàn bộ vào một Docker Image — ai cũng chạy `docker run` là có môi trường giống hệt nhau.

**Mở rộng:**

- Tìm hiểu sự khác nhau cụ thể giữa Container và Virtual Machine (VM) về kiến trúc (VM có Hypervisor và Guest OS riêng cho mỗi máy ảo, trong khi Container chia sẻ chung OS Kernel của host).

**Nguồn trích dẫn:**

- Docker Official Docs – "What is a Container?": https://www.docker.com/resources/what-container/

---

## Câu 7: Docker Image là gì?

**Câu trả lời:** Docker Image là một bản snapshot (ảnh chụp) **bất biến (immutable)**, đóng gói đầy đủ mọi thứ cần thiết để chạy một ứng dụng: mã nguồn, runtime, thư viện phụ thuộc, biến môi trường và các file cấu hình. Image được xây dựng theo nhiều **layer** xếp chồng lên nhau, mỗi layer tương ứng với một dòng lệnh (instruction) trong file `Dockerfile` (ví dụ `FROM`, `COPY`, `RUN`); các layer không thay đổi giữa các lần build sẽ được Docker tái sử dụng từ cache, giúp build nhanh hơn đáng kể.

**Phân biệt Image và Container:** Image giống như một "bản thiết kế" hay "class" tĩnh, còn **Container** là một **instance đang chạy** được khởi tạo từ Image đó — giống mối quan hệ giữa `class` và `object` trong lập trình hướng đối tượng. Từ một Image duy nhất, có thể khởi chạy nhiều Container độc lập cùng lúc.

**Ví dụ:** File `Dockerfile` chỉ định `FROM golang:1.22`, sau đó `COPY` mã nguồn vào, rồi `RUN go build`. Khi chạy `docker build`, Docker tạo ra một Image hoàn chỉnh chứa binary đã build sẵn. Chạy `docker run <image>` nhiều lần sẽ tạo ra nhiều Container độc lập, mỗi cái là một tiến trình đang chạy riêng biệt từ cùng một Image gốc.

**Mở rộng:**

- Tìm hiểu kỹ thuật **Multi-stage Build** trong Dockerfile — dùng một stage để build ứng dụng (cần đầy đủ công cụ biên dịch) rồi chỉ copy kết quả binary sang một stage cuối cùng gọn nhẹ hơn, giúp giảm đáng kể kích thước Image production.

**Nguồn trích dẫn:**

- Docker Official Docs – "Docker Images": https://docs.docker.com/get-started/docker-concepts/the-basics/what-is-an-image/

---

## Câu 8: Docker Compose là gì?

**Câu trả lời:** Docker Compose là công cụ cho phép định nghĩa và chạy đồng thời nhiều container liên quan đến nhau (multi-container application) chỉ bằng một file cấu hình duy nhất ở định dạng YAML, thường đặt tên là `docker-compose.yml`. Thay vì phải gõ lần lượt nhiều câu lệnh `docker run` phức tạp (kèm theo cấu hình network, volume, biến môi trường cho từng container), Docker Compose cho phép khai báo toàn bộ các **service** (mỗi service tương ứng một container), **network** (mạng nội bộ để các container giao tiếp với nhau) và **volume** (lưu trữ dữ liệu bền vững) trong một file duy nhất, sau đó chỉ cần chạy một lệnh `docker compose up` để khởi động toàn bộ hệ thống cùng lúc.

**Ví dụ:** Một ứng dụng backend cần 3 thành phần chạy cùng lúc: API server, PostgreSQL, và Redis. Thay vì chạy riêng 3 lệnh `docker run` với các tham số kết nối phức tạp, file `docker-compose.yml` định nghĩa cả 3 service này cùng với thứ tự khởi động (`depends_on`), chỉ cần `docker compose up` là toàn bộ hệ thống (API + DB + Cache) được khởi chạy và tự kết nối với nhau qua một network nội bộ do Compose tự tạo.

**Mở rộng:**

- Docker Compose phù hợp cho môi trường phát triển (local) và các hệ thống quy mô vừa/nhỏ; ở quy mô lớn hơn với nhiều server, thường cần các công cụ orchestration mạnh hơn như Kubernetes.

**Nguồn trích dẫn:**

- Docker Official Docs – "Docker Compose Overview": https://docs.docker.com/compose/

---

## Câu 9: .dockerignore là gì? Dùng khi nào?

**Câu trả lời:** `.dockerignore` là một file cấu hình (tương tự `.gitignore`) dùng để liệt kê các file và thư mục **không nên** được đưa vào quá trình build Docker Image, dù chúng vẫn tồn tại trong thư mục dự án.

**Dùng khi nào và tại sao cần thiết:**

- **Giảm kích thước Build Context:** Khi chạy `docker build`, Docker gửi toàn bộ nội dung thư mục dự án (build context) đến Docker daemon trước khi build. Nếu không loại trừ các thư mục nặng và không cần thiết (như `node_modules`, `.git`, các file log), quá trình gửi dữ liệu này sẽ chậm hơn đáng kể.
- **Tránh rò rỉ thông tin nhạy cảm:** Loại trừ các file chứa secret như `.env`, file chứa API key hoặc mật khẩu — tránh vô tình đóng gói các thông tin này vào Image (Image có thể bị đẩy lên registry công khai hoặc được người khác truy cập).
- **Tăng tốc build nhờ tận dụng cache layer tốt hơn:** Loại bỏ các file thay đổi thường xuyên nhưng không ảnh hưởng đến ứng dụng (ví dụ file log, file tạm) giúp Docker tái sử dụng cache của các layer trước đó hiệu quả hơn.

**Ví dụ:** Một file `.dockerignore` điển hình cho dự án Go:

```
.git
.env
*.log
node_modules
tmp/
```

Nhờ vậy, khi build Image cho ứng dụng Go, Docker sẽ không copy thư mục `.git` (thường rất nặng do lưu lịch sử commit) hay file `.env` chứa thông tin kết nối database nhạy cảm vào trong Image cuối cùng.

**Mở rộng:**

- Luôn kiểm tra kỹ `.dockerignore` trước khi build Image dự định public lên Docker Hub, tránh vô tình để lộ secret — đây là lỗi bảo mật khá phổ biến trong thực tế.

**Nguồn trích dẫn:**

- Docker Official Docs – "Build context" (mục .dockerignore): https://docs.docker.com/build/concepts/context/#dockerignore-files

---