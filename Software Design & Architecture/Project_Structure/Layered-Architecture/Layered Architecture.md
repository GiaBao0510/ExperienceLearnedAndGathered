# Layered Architecture — Kiến trúc phân tầng

# I. Khái niệm

**Layered Architecture** (còn gọi là **N-tier Architecture**) là kiểu kiến trúc tổ chức ứng dụng thành nhiều tầng xếp chồng theo chiều dọc, mỗi tầng có trách nhiệm rõ ràng và chỉ giao tiếp với tầng liền kề. Số lượng tầng và chức năng của từng tầng có thể khác nhau tùy theo độ phức tạp của dự án, nhưng các tầng đều phải tuân thủ các nguyên tắc sau:

**Technical Partitioning**

Kiến trúc được chia thành các nhóm chức năng (tầng). Mỗi component trong một tầng chỉ thực hiện chức năng đã được định nghĩa cho tầng đó. Ví dụ: tầng Persistence chỉ quản lý transaction và truy vấn database; tầng Application (Business) xử lý logic nghiệp vụ; tầng Presentation xử lý UI.

**Separation of Concerns**

Mỗi tầng có vai trò riêng biệt và độc lập, đảm bảo logic không bị chồng chéo. Component trong tầng Presentation chỉ xử lý giao diện và tương tác người dùng, không được truy vấn trực tiếp database — đó là trách nhiệm của tầng Persistence. Nhờ đó, một thay đổi ở tầng Presentation (ví dụ: mở rộng sang mobile) không ảnh hưởng đến Business hay Persistence layer.

**Giao tiếp giữa các tầng**

Tầng trên sử dụng các chức năng mà tầng dưới cung cấp thông qua interface đã được định nghĩa sẵn. Tầng dưới thực thi yêu cầu từ tầng trên và không có phụ thuộc ngược lại (không gọi ngược lên tầng trên).

---

#### Mục đích sử dụng

- Giảm độ phức tạp của hệ thống bằng cách chia thành các lớp logic độc lập.
- Tăng khả năng bảo trì và nâng cấp khi mỗi lớp hoạt động độc lập.
- Tăng khả năng tái sử dụng thông qua tách biệt chức năng.
- Tạo nền tảng dễ hiểu cho team mới tiếp cận dự án.

#### Các bước áp dụng

Bối cảnh: Một doanh nghiệp xây dựng hệ thống quản lý khách hàng (CRM).

- Bước 1: Xác định các tầng cốt lõi (thường là Presentation — Business — Persistence — Database).
- Bước 2: Định nghĩa trách nhiệm và ranh giới rõ ràng của từng tầng.
- Bước 3: Thiết kế interface giữa các tầng để giảm phụ thuộc trực tiếp (loose coupling).
- Bước 4: Tách logic nghiệp vụ khỏi giao diện để dễ mở rộng và kiểm thử độc lập.
- Bước 5: Viết unit test cho từng tầng để đảm bảo chất lượng từng phần.

Ví dụ: Tầng Business thay đổi quy tắc tính điểm khách hàng mà không ảnh hưởng đến giao diện người dùng ở tầng Presentation.

#### Lưu ý thực tiễn

- Tránh để các tầng phụ thuộc ngược (Presentation gọi trực tiếp Database).
- Không nhồi quá nhiều logic vào một tầng, dễ tạo ra "God Layer" — tầng biết và làm quá nhiều thứ, phá vỡ nguyên tắc Separation of Concerns.
- Áp dụng nguyên tắc SOLID và Dependency Injection để giảm phụ thuộc cứng giữa các tầng.

---

# II. Kiến trúc cơ bản

Kiến trúc phân tầng không quy định cố định số lượng tầng, nhưng cấu hình phổ biến nhất gồm 4 tầng:

![](https://images.viblo.asia/122e6427-8edd-4c01-ae50-1616c68cae9b.png)

**Presentation**

Xử lý tương tác với người dùng cuối: validate input, render UI dựa trên dữ liệu trả về từ tầng Application, và chuyển tiếp yêu cầu của người dùng xuống tầng Application để xử lý. Tầng này không chứa business logic.

**Application (Business)**

Một số tài liệu gọi là Business Layer. Chứa toàn bộ logic nghiệp vụ của hệ thống. Định nghĩa và cung cấp interface cho tầng Presentation gọi vào, đồng thời gọi xuống tầng Persistence để lấy hoặc cập nhật dữ liệu.

**Persistence**

Chịu trách nhiệm tương tác với nguồn lưu trữ dữ liệu. Không chỉ giới hạn ở database — mục đích chính của tầng này là duy trì trạng thái của các entity trong suốt vòng đời của ứng dụng. Entity có thể được lưu trên database quan hệ, NoSQL, file system, hoặc bất kỳ hệ thống lưu trữ nào khác. Tầng này quản lý connection pool và thực thi các câu truy vấn theo yêu cầu từ tầng trên.

**Database**

Là các hệ thống lưu trữ thực tế như MySQL, PostgreSQL, MongoDB, Redis... Tầng này không thuộc phần code của ứng dụng mà là hạ tầng bên ngoài. Tầng Persistence là thành phần duy nhất trong ứng dụng được phép tương tác trực tiếp với tầng này.

---

# III. Open Layer và Closed Layer

**Closed Layer**

Bắt buộc mỗi request phải đi qua tất cả các tầng theo thứ tự từ trên xuống, không được bỏ qua tầng nào:

```
Presentation → Business → Persistence → Database
```

![](https://images.viblo.asia/e94ad7e5-a502-4f25-81a8-bdc682249579.png)

Cách tiếp cận này áp dụng khái niệm **Layers of Isolation**: mỗi tầng chỉ biết và phụ thuộc vào tầng liền kề phía dưới. Khi một tầng thay đổi implementation nội bộ (ví dụ: đổi ORM, đổi cache strategy), các tầng khác không bị ảnh hưởng miễn là interface không thay đổi.

Ứng dụng thực tế trong API backend theo mô hình Controller — Service — Repository:

![](https://statics.cdn.200lab.io/2024/01/Screenshot-2024-01-19-at-23.37.08.png?width=800)

1. **Controller** (Presentation Layer): Parse và validate HTTP request từ client, trả về HTTP response.
2. **Service** (Business Layer): Chứa toàn bộ logic nghiệp vụ.
3. **Repository** (Persistence Layer): Đóng gói toàn bộ logic truy vấn và mapping dữ liệu với database. Đây là nơi dùng DAO (Data Access Object) hoặc trả về DTO (Data Transfer Object) để tầng Service xử lý, không phải nơi dùng DTO để truyền dữ liệu giữa tầng Service và Repository theo nghĩa gốc.
4. **Database** (Database Layer): Hệ thống lưu trữ thực tế.

![](https://statics.cdn.200lab.io/2024/01/Screenshot-2024-01-19-at-23.45.30.png?width=800)

Closed layer giúp hệ thống tường minh, dễ hiểu và dễ bảo trì, giảm thiểu phụ thuộc chéo (interdependencies). Phù hợp với hệ thống yêu cầu tính ổn định và dễ bảo trì.

**Open Layer**

Ngược lại với closed layer, open layer cho phép một tầng gọi trực tiếp đến các tầng bên dưới nó, không nhất thiết phải đi qua tầng liền kề. Lưu ý: chiều ngược lại vẫn không được phép — tầng dưới không được gọi lên tầng trên.

Ví dụ với kiến trúc 5 tầng:

```
Presentation → Business → Services → Persistence → Database
```

Với open layer, một request có thể bỏ qua tầng `Services` nếu tầng này chỉ forward yêu cầu mà không thêm logic:

```
Presentation → Business → Persistence → Database
```

![](https://images.viblo.asia/f38fd72f-083d-407a-b357-a8d3e16bdde8.png)

Open layer tăng hiệu suất bằng cách loại bỏ các pass-through layer (tầng chỉ làm nhiệm vụ chuyển tiếp, không xử lý thêm logic — còn gọi là Layer Bloat). Nhược điểm là làm giảm tính mô-đun, khó kiểm thử và bảo trì hơn, đồng thời cần tài liệu rõ ràng để giải thích lý do bỏ qua tầng nào trong từng trường hợp.

Phù hợp với hệ thống yêu cầu hiệu suất cao và có pass-through layer rõ ràng không cần thiết.

---

# IV. Tính chất và Antipattern

Layered Architecture là kiến trúc phổ biến nhất vì sự đơn giản và dễ áp dụng. Tuy nhiên, để sử dụng hiệu quả cần liên tục review và maintain kiến trúc. Một bẫy thường gặp là **Architecture Sinkhole Anti-pattern**.

**Architecture Sinkhole Anti-pattern**

Xảy ra khi phần lớn các request đi qua tất cả các tầng nhưng mỗi tầng chỉ forward hoặc thêm một lượng logic không đáng kể. Điều này làm giảm hiệu suất toàn hệ thống và tăng độ phức tạp không cần thiết.

Ngưỡng thực tế để đánh giá: nếu hơn 80% request rơi vào tình trạng này, cần cân nhắc tái cấu trúc bằng các cách sau:

- Thêm hoặc xóa tầng khi cần thiết.
- Chuyển tầng từ closed sang open để cho phép bỏ qua pass-through layer.
- Gộp hai tầng có chức năng chồng lặp thành một tầng mới.

---
#### Ví dụ cấu trúc thư mục dự án Go theo Layered Architecture

```
.
├── Dockerfile
├── Makefile
├── README.md
├── build.sh
├── cmd/
│   └── drunk/
│       └── main.go
├── environment/
├── global/
├── go.mod
├── go.sum
├── init.sql
├── internal/
│   ├── common/
│   ├── initialize/
│   │   ├── mysql.go
│   │   ├── router.go
│   │   ├── run.go
│   │   └── service.go
│   ├── middleware/
│   │   ├── cors.go
│   │   └── validation.go
│   ├── modules/
│   │   ├── auth/
│   │   │   ├── controller.go       # Presentation layer
│   │   │   ├── dto/
│   │   │   │   └── account.go      # Data Transfer Objects
│   │   │   ├── entity/
│   │   │   │   └── account.go      # Domain entities
│   │   │   ├── service.go          # Business layer (interface)
│   │   │   └── serviceimpl.go      # Business layer (implementation)
│   │   ├── rbac/
│   │   └── user/
│   └── router/
│       ├── auth/
│       │   ├── auth.router.go
│       │   ├── enter.go
│       │   └── rbac.router.go
│       ├── enter.go
│       └── user/
│           └── enter.go
├── pkg/
│   └── response/
│       ├── codeErr.go
│       └── response.go
├── scripts/
├── tests/
└── utils/
    ├── cache.go
    └── validator.go
```

---

# V. Nên dùng khi nào

**Hệ thống nhỏ và cần triển khai nhanh**

Layered Architecture là lựa chọn phù hợp cho các hệ thống nhỏ, chi phí thấp, hoặc có deadline ngắn vì:

- Là kiến trúc monolithic — không mang sự phức tạp của kiến trúc phân tán (distributed systems).
- Được sử dụng rộng rãi nên hầu hết developer đều quen thuộc, giúp onboard team mới nhanh hơn.
- Phân chia team theo technical specialty (FE, BE, Database) rất tự nhiên: FE phụ trách Presentation, BE phụ trách Business và Persistence, DBA phụ trách Database layer.

Với các dự án kiểu này, không cần đầu tư thời gian phân tích và chia nhỏ theo domain như các kiến trúc phức tạp hơn. Điều quan trọng là đánh giá scope dự án khách quan và chính xác trước khi quyết định kiến trúc.

**Thay đổi thiên về kỹ thuật (Technical-Oriented)**

Layered Architecture phù hợp với các dự án mà yêu cầu nghiệp vụ tương đối ổn định, và các thay đổi chủ yếu ở tầng kỹ thuật: đổi database, thêm hỗ trợ platform mới mà không thay đổi business logic. Ví dụ: hệ thống quản lý ban đầu triển khai trên web, sau đó mở rộng sang mobile — chỉ cần mở rộng tầng Presentation; Business và Persistence layer không bị ảnh hưởng.

---
# VI. Không nên dùng khi nào

**Hệ thống yêu cầu scalability và fault tolerance cao**

Layered Architecture thường đi kèm với monolithic deployment. Toàn bộ các tầng chạy trong một khối duy nhất. Điều này có lợi thế (giảm overhead giao tiếp, đơn giản hóa transaction) nhưng dẫn đến hai hạn chế lớn:

- **Scale issue:** Khi một tầng cần scale (ví dụ: Business layer đang bị bottleneck), phải scale toàn bộ monolith, kể cả các tầng khác không cần thêm tài nguyên. Điều này vừa tốn kém vừa không hiệu quả.

- **Fault tolerance:** Khi một phần của hệ thống gặp sự cố, toàn bộ ứng dụng có thể bị ảnh hưởng (single point of failure). Giải pháp Active-Standby giúp giảm thiểu downtime nhưng không giải quyết được vấn đề scale căn bản.


**Hệ thống thay đổi thường xuyên theo domain**

Vì Layered Architecture là kỹ thuật phân vùng theo chiều kỹ thuật (horizontal slicing), một thay đổi domain nhỏ (ví dụ: thêm trường `expire_date` cho sản phẩm) sẽ lan rộng qua tất cả các tầng: Database phải thêm cột, Persistence phải cập nhật query, Business phải xử lý logic mới, Presentation phải hiển thị thêm trường. Càng nhiều tầng, chi phí cho mỗi thay đổi domain càng cao, và khả năng bảo trì càng giảm theo thời gian.

---

# VII. Câu hỏi thường gặp (FAQ)

**Q1: Layered Architecture có phù hợp cho hệ thống lớn không?**

Có thể, nhưng cần đánh giá cẩn thận. Nếu hệ thống lớn nhưng thay đổi chủ yếu về mặt kỹ thuật và không yêu cầu scalability cao, Layered Architecture vẫn khả thi. Khi business domain phức tạp và thay đổi thường xuyên, nên cân nhắc Domain-Driven Design hoặc kiến trúc phân tán (microservices).

**Q2: Có thể bỏ qua tầng Business không?**

Không nên. Bỏ qua Business layer dẫn đến business logic rò rỉ vào Presentation hoặc Persistence layer, vi phạm Separation of Concerns và khiến code khó kiểm thử, khó bảo trì.

**Q3: Layered Architecture khác gì Microservices?**

Layered Architecture là mẫu tổ chức bên trong một ứng dụng (intra-service). Microservices là mẫu chia hệ thống thành nhiều dịch vụ độc lập (inter-service). Hai khái niệm không loại trừ nhau: mỗi microservice có thể áp dụng Layered Architecture bên trong.

**Q4: Có nên dùng chung tầng Data Access cho nhiều ứng dụng không?**

Có thể, nhưng cần cẩn thận. Chia sẻ trực tiếp tầng Data Access giữa nhiều ứng dụng tạo ra tight coupling ở tầng database. Cách tiếp cận an toàn hơn là cung cấp API trung gian (internal service hoặc shared library với interface rõ ràng) để tránh phụ thuộc chéo và cho phép thay đổi implementation độc lập.

**Q5: Layered Architecture có giúp tăng bảo mật không?**

Có ở mức độ nhất định. Mỗi tầng có thể kiểm soát truy cập và validation riêng (ví dụ: Presentation validate input format, Business validate business rules, Persistence validate data integrity). Tuy nhiên, Layered Architecture không phải giải pháp bảo mật toàn diện — vẫn cần các biện pháp bảo mật riêng biệt ở tầng network, infrastructure và application.

**Q6: Khi nào nên chuyển từ Layered sang kiến trúc khác?**

Các dấu hiệu cho thấy cần xem xét thay đổi kiến trúc: hệ thống thường xuyên rơi vào Architecture Sinkhole Anti-pattern; một tầng trở thành bottleneck nhưng không thể scale độc lập; business domain phân tán và thay đổi theo nhiều hướng khác nhau; team lớn và việc làm việc trên cùng một monolith gây conflict thường xuyên.

---

# VIII. Kết luận

Layered Architecture là lựa chọn lý tưởng cho những dự án cần sự đơn giản, dễ hiểu và triển khai nhanh. Kiến trúc này đặc biệt phù hợp khi dự án có quy mô vừa và nhỏ, team quen thuộc với cách tổ chức code theo technical layer, và các thay đổi thiên về kỹ thuật hơn là domain.

Tuy nhiên, kiến trúc này có những hạn chế rõ ràng khi đối mặt với yêu cầu scalability cao, fault tolerance, hoặc domain thay đổi thường xuyên. Quan trọng hơn, hiệu quả của Layered Architecture không đến từ việc áp dụng cơ học mà từ việc thực sự hiểu nguyên tắc và liên tục review để tránh rơi vào Architecture Sinkhole Anti-pattern.

---
## Tài liệu tham khảo

1. Fundamentals of Software Architecture — Mark Richards & Neal Ford
2. Software Architecture Patterns — Mark Richards
3. [Software Architecture Monday — Mark Richards](https://www.youtube.com/playlist?list=PLdsOZAx8I5umhnn5LLTNJbFgwA3xbycar)
4. [Microservices — Martin Fowler](https://martinfowler.com/articles/microservices.html#CharacteristicsOfAMicroserviceArchitecture)

> Bản gốc ghi tác giả là "Martin Flower" — tên đúng là **Martin Fowler**. Đã sửa.

---
## Thông tin bổ sung

### 1. So sánh Layered Architecture với các kiến trúc thay thế phổ biến

Khi Layered Architecture bắt đầu bộc lộ hạn chế, các kiến trúc sau thường được cân nhắc:

|Kiến trúc|Phân vùng theo|Phù hợp khi|
|---|---|---|
|Layered Architecture|Kỹ thuật (technical)|Dự án nhỏ-vừa, team kỹ thuật chuyên biệt|
|Modular Monolith|Domain (trong một monolith)|Muốn tách domain nhưng chưa cần microservices|
|Hexagonal (Ports & Adapters)|Domain + Infrastructure tách biệt|Cần kiểm thử domain logic độc lập hoàn toàn|
|Clean Architecture|Domain ở trung tâm, dependency hướng vào trong|Hệ thống phức tạp, nhiều external dependencies|
|Microservices|Domain services độc lập|Hệ thống lớn, cần scale và deploy độc lập từng service|

### 2. Layered Architecture trong context của Go

Go không có framework MVC/Layered bắt buộc, nhưng cộng đồng đã hình thành convention phổ biến:

- Dùng `internal/` để đảm bảo các package không bị import bởi code bên ngoài module.
- Tổ chức theo module/feature (ví dụ `internal/modules/auth/`) thay vì theo technical layer (`controller/`, `service/`, `repository/` ở cùng cấp) — cách này giúp tìm kiếm code liên quan đến một feature dễ hơn khi dự án lớn lên.
- Interface được định nghĩa phía consumer (tầng trên) thay vì phía implementer (tầng dưới) — đây là Go idiomatic style và hỗ trợ tốt cho Dependency Inversion.

### 3. Dependency Injection trong Layered Architecture với Go

Để các tầng không phụ thuộc trực tiếp vào implementation của tầng dưới, Go thường dùng interface + constructor injection:

```go
// Business layer định nghĩa interface nó cần từ Persistence
type UserRepository interface {
    FindByID(ctx context.Context, id int64) (*User, error)
    Save(ctx context.Context, user *User) error
}

// Business layer nhận dependency qua constructor
type UserService struct {
    repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
    return &UserService{repo: repo}
}
```

Pattern này cho phép mock Repository trong unit test của Service mà không cần database thật, đúng với nguyên tắc Layers of Isolation.