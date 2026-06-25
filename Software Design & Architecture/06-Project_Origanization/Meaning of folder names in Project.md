Tài liệu này mô tả ý nghĩa và trách nhiệm của các thư mục phổ biến trong một dự án backend, đặc biệt áp dụng cho các ngôn ngữ như Go, Node.js, hoặc Java. Hiểu rõ cấu trúc thư mục giúp onboard nhanh hơn, giảm nhầm lẫn khi đặt code vào đúng chỗ, và duy trì tính nhất quán trong toàn bộ dự án.

---
## `docs/`

Chứa tài liệu kỹ thuật của dự án: API documentation (Swagger/OpenAPI), ERD (Entity Relationship Diagram), architecture decision records (ADR), và hướng dẫn triển khai. Không chứa code.

---
## `models/` (hoặc `entities/`)

Định nghĩa cấu trúc dữ liệu, ánh xạ các bảng trong cơ sở dữ liệu và các quy tắc ORM/ODM (Object-Relational Mapping / Object-Document Mapping).

```go
type User struct {
    ID        uint      `gorm:"primaryKey"`
    Email     string    `gorm:"uniqueIndex;not null"`
    CreatedAt time.Time
}
```

**Phân biệt với DTO:** Model/Entity ánh xạ với database; DTO (Data Transfer Object) là struct truyền dữ liệu giữa các tầng (ví dụ: request body từ client). Hai loại này không nên dùng chung.

---

## `repositories/` (hoặc `data_access/`, `daos/, repo/`)

Chứa mã nguồn thực hiện các thao tác truy vấn trực tiếp xuống cơ sở dữ liệu: thêm, sửa, xóa, lấy dữ liệu (CRUD). Đây là lớp duy nhất trong hệ thống được phép tương tác trực tiếp với database.

Repository nhận và trả về các entity/model. Tầng Service gọi vào Repository thông qua interface để có thể mock trong unit test.

```go
type UserRepository interface {
    FindByID(ctx context.Context, id uint) (*model.User, error)
    Save(ctx context.Context, user *model.User) error
    Delete(ctx context.Context, id uint) error
}
```

---

## `middlewares/`

Chứa các hàm trung gian xử lý request trước khi đến Controller và response trước khi trả về client. Middleware được áp dụng theo dạng chuỗi (chain).

Các middleware phổ biến:

- **Authentication:** Kiểm tra JWT hoặc API Key.
- **Authorization (RBAC):** Kiểm tra quyền của người dùng với tài nguyên cụ thể.
- **Rate Limiting:** Giới hạn số request trong khoảng thời gian.
- **Request Logging:** Ghi lại method, path, status code, response time.
- **Recovery:** Bắt panic và trả về lỗi 500 thay vì crash server.
- **CORS:** Cấu hình Cross-Origin Resource Sharing.

---
## `tests/` (hoặc `specs/`)

Chứa toàn bộ các file mã nguồn dùng để kiểm thử tự động (Unit Test, Integration Test, E2E Test) nhằm đảm bảo hệ thống chạy đúng theo yêu cầu nghiệp vụ.

---

## `controllers/` (hoặc `handlers/`)

Tầng tiếp nhận HTTP request từ client: parse và validate input, gọi Service layer để xử lý nghiệp vụ, và trả về HTTP response.

Controller không chứa business logic. Nếu logic nghiệp vụ phức tạp xuất hiện trong Controller, đó là dấu hiệu của antipattern **Fat Controller** — cần chuyển xuống Service.

```go
func (h *UserHandler) GetUser(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, response.Error("invalid id"))
        return
    }
    user, err := h.userService.GetByID(c.Request.Context(), uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, response.Error(err.Error()))
        return
    }
    c.JSON(http.StatusOK, response.Success(user))
}
```

---

## `services/` (hoặc `use_cases/`, `business_logic/`)

Chứa toàn bộ business logic của ứng dụng. Đây là tầng xử lý các quy tắc nghiệp vụ: tính toán, xác thực logic phức tạp, điều phối nhiều repository, gọi external service.

Service không biết gì về HTTP request/response. Service nhận vào các kiểu dữ liệu Go thuần, gọi Repository qua interface, và trả về kết quả hoặc error.

```go
type UserService interface {
    GetByID(ctx context.Context, id uint) (*dto.UserResponse, error)
    Register(ctx context.Context, req *dto.RegisterRequest) error
}
```

---

## `scripts/` (hoặc `build/`)

Chứa các đoạn mã tự động hóa (automation scripts) phục vụ cho việc triển khai (deploy), sao lưu dữ liệu (backup), hoặc các tác vụ chạy theo lịch (cron jobs).

---

## `response/`

Chứa các struct và hàm chuẩn hóa format HTTP response trả về cho client. Đảm bảo tất cả API trong dự án trả về cùng một cấu trúc JSON nhất quán.

```go
type Response struct {
    Success bool        `json:"success"`
    Message string      `json:"message,omitempty"`
    Data    interface{} `json:"data,omitempty"`
    Error   string      `json:"error,omitempty"`
}

func OK(data interface{}) Response {
    return Response{Success: true, Data: data}
}

func Err(message string) Response {
    return Response{Success: false, Error: message}
}
```

---

## `database/` (hoặc `migrations/`, `seeds/`)

- **`migrations/`**: 
	Chứa các file quản lý lịch sử thay đổi cấu trúc database (tạo bảng, thêm cột, xóa index...). Mỗi file là một phiên bản thay đổi được đánh số thứ tự.
	Chứa các file SQL hoặc migration script dùng để tạo, sửa đổi và quản lý phiên bản schema database theo thời gian.
	Mỗi migration file thường gồm hai phần: `up` (áp dụng thay đổi) và `down` (hoàn tác thay đổi). Công cụ phổ biến trong Go: **golang-migrate**, **goose**.
Ví dụ:
```file
000001_create_users_table.up.sql
000001_create_users_table.down.sql
000002_add_email_index.up.sql
000002_add_email_index.down.sql
```

- **`seeds/`**: Chứa dữ liệu mẫu (mock data / seed data) để nạp vào database khi khởi tạo hoặc test môi trường.

---
### `pkg/`:

Chứa các package có thể tái sử dụng độc lập với business logic của ứng dụng, hoặc có thể được import bởi code bên ngoài module (nếu là thư viện). Thường bao gồm: wrapper cho thư viện bên thứ ba (logger, cache client, HTTP client), custom middleware framework, utility packages.

Trong Go, `pkg/` nằm ngang cấp với `internal/` và có thể được import từ bên ngoài module. Ngược lại, `internal/` chỉ được import bởi code trong cùng module — đây là cơ chế Go enforce bằng compiler.

---

## Các thư mục khác

| Tên thư mục | Ý nghĩa                                                                                                                                                                                                        |
| ----------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `agency/`   | Thường xuất hiện trong kiến trúc agent-based: chứa các thành phần xử lý tác vụ tự động hoặc điều phối luồng công việc (workflow orchestration). Ít phổ biến, mang tính project-specific.                       |
| `hack/`     | Chứa các đoạn mã giải quyết tạm thời (workaround) hoặc bản vá lỗi nhanh. Thư mục này thường chỉ tồn tại ngắn hạn; code ở đây cần được refactor và chuyển vào đúng tầng về sau.                                 |
| `sql/`      | Chứa các file truy vấn SQL thuần túy (raw SQL), stored procedure, hoặc view definition được viết thủ công, không thông qua ORM.                                                                                |
| `pkg/`      | Phổ biến trong các dự án Go: chứa các package nội bộ có thể tái sử dụng giữa nhiều module trong cùng một repo (monorepo). Khác với `internal/`, code trong `pkg/` được thiết kế để có thể import từ bên ngoài. |
| `utility/`  | Tương đương với `utils/`. Xem mục `utils/` bên dưới.                                                                                                                                                           |

---

## `logs/`

Nơi lưu trữ các file nhật ký hoạt động của hệ thống (error log, access log, audit log), giúp lập trình viên theo dõi và chẩn đoán sự cố khi hệ thống vận hành. Thư mục này thường bị loại trừ khỏi version control (thêm vào `.gitignore`).

---

## `logger/`

Chứa mã nguồn cấu hình và khởi tạo hệ thống logging (ví dụ: wrapper cho `zap`, `logrus`, `winston`). Khác với `log/` là nơi lưu file log đầu ra, `logger/` là nơi định nghĩa _cách_ ghi log.

---

## `utils/` (hoặc `helpers/`, `shared/`)

Chứa các hàm tiện ích thuần túy (pure function), không phụ thuộc vào business logic, có thể tái sử dụng ở nhiều nơi trong dự án.

Ví dụ điển hình: format thời gian, hash mật khẩu, generate random string, parse pagination parameters, validate email format.

**Lưu ý quan trọng:** Không lạm dụng `utils/` như một "thùng rác" chứa code không biết để đâu. Nếu một hàm chỉ liên quan đến một domain cụ thể, hãy đặt nó gần domain đó thay vì trong `utils/`.

---

## `common/`

Tương tự `utils/` nhưng thường chứa các định nghĩa dùng chung ở tầng cao hơn: constants, custom error types, shared interfaces, base structs.

```go
// common/errors.go
var (
    ErrNotFound     = errors.New("resource not found")
    ErrUnauthorized = errors.New("unauthorized")
    ErrForbidden    = errors.New("forbidden")
)
```

**Phân biệt `utils/` và `common/`:**

- `utils/`: Hàm tiện ích thuần túy, xử lý tác vụ cụ thể.
- `common/`: Định nghĩa chung (constants, errors, interfaces) được dùng xuyên suốt nhiều package.

---
## `configs/`

Chứa các file cấu hình ứng dụng (`.yaml`, `.json`, `.env`) và code đọc cấu hình vào struct. Không lưu secret (password, API key) trực tiếp trong file cấu hình được commit vào source control — dùng biến môi trường hoặc secret management system.

```yaml
# config.yaml
server:
  port: 8080
  timeout: 30s

database:
  host: localhost
  port: 5432
  name: myapp
```

---

## `routes/` (hoặc `api/` hoặc `routers/`)

Định nghĩa các API endpoint và ánh xạ chúng tới Controller handler tương ứng. Thường cũng là nơi đăng ký middleware theo nhóm route.

```go
func RegisterRoutes(r *gin.Engine, h *handler.UserHandler, auth middleware.AuthMiddleware) {
    v1 := r.Group("/api/v1")
    {
        v1.POST("/auth/login", h.Login)
        v1.POST("/auth/register", h.Register)

        authorized := v1.Group("/")
        authorized.Use(auth.Authenticate())
        {
            authorized.GET("/users/:id", h.GetUser)
            authorized.PUT("/users/:id", h.UpdateUser)
        }
    }
}
```

---

## `internal/` (hoặc `src/`)

Thư mục chứa toàn bộ mã nguồn cốt lõi của ứng dụng. Trong Go, chứa toàn bộ code nghiệp vụ của ứng dụng. Go compiler **bắt buộc** rằng code bên trong `internal/` chỉ có thể được import bởi code nằm trong cùng module — không package bên ngoài nào có thể import. Đây là cơ chế encapsulation ở cấp module.

| Thư mục con   | Ý nghĩa                                                                                                                                       |
| ------------- | --------------------------------------------------------------------------------------------------------------------------------------------- |
| `config/`     | Đọc và parse file cấu hình, khởi tạo các biến môi trường                                                                                      |
| `consts/`     | Định nghĩa các hằng số (constants) dùng chung toàn ứng dụng                                                                                   |
| `controller/` | Tầng tiếp nhận request từ client                                                                                                              |
| `dao/`        | Data Access Object — thao tác trực tiếp với database                                                                                          |
| `logic/`      | Tầng xử lý nghiệp vụ (tương đương `service/` trong nhiều kiến trúc)                                                                           |
| `di/`         | Dependency Injection — khởi tạo và tiêm phụ thuộc giữa các thành phần                                                                         |
| `domain/`     | Định nghĩa các entity/domain object của nghiệp vụ                                                                                             |
| `initialize/` | Mã khởi động ứng dụng: kết nối DB, cấu hình server, đăng ký route                                                                             |
| `model/`      | Cấu trúc dữ liệu ánh xạ với bảng database (ORM model)                                                                                         |
| `packed/`     | Tài nguyên đã được đóng gói (embedded assets) hoặc file build trung gian                                                                      |
| `service/`    | Tầng xử lý nghiệp vụ chính (Business Logic)                                                                                                   |
| `util/`       | Hàm tiện ích dùng chung                                                                                                                       |
| `middleware/` | Middleware xử lý request/response                                                                                                             |
| `po/`         | Persistent Object — đối tượng ánh xạ trực tiếp với bảng database, thường dùng trong kiến trúc Java/Go phân tách rõ domain object và DB object |
| `repo/`       | Tầng Repository — trừu tượng hóa thao tác database                                                                                            |
| `routers/`    | Định nghĩa và đăng ký các route của ứng dụng                                                                                                  |

**Phân biệt `model/` và `po/` (Persistent Object):**

- `model/`: Struct dùng trong business logic, có thể chứa computed field và methods.
- `po/`: Struct ánh xạ trực tiếp với bảng database (1-to-1 với schema). Tách `po/` khỏi `model/` giúp tránh coupling giữa database schema và business logic.

---
## `manifest/` (hoặc `deploy/`)

Chứa các file khai báo và cấu hình phục vụ cho việc triển khai và vận hành hệ thống. Các thư mục con thường gặp:

| Thư mục con | Ý nghĩa                                                                                   |
| ----------- | ----------------------------------------------------------------------------------------- |
| `config/`   | File cấu hình theo môi trường (development, staging, production)                          |
| `deploy/`   | Manifest triển khai: Kubernetes YAML, Helm chart, Ansible playbook                        |
| `docker/`   | Dockerfile, docker-compose.yml và các file liên quan đến container hóa                    |
| `i18n/`     | File đa ngôn ngữ (internationalization): bản dịch chuỗi giao diện hoặc thông báo hệ thống |
| `protobuf/` | File định nghĩa Protocol Buffer (`.proto`) dùng cho giao tiếp giữa các service (gRPC)     |


---
## `resource/` (hoặc `static/`)

Chứa tài nguyên tĩnh phục vụ ứng dụng. Các thư mục con thường gặp:

| Thư mục con | Ý nghĩa                                                             |
| ----------- | ------------------------------------------------------------------- |
| `public/`   | File tĩnh được phục vụ trực tiếp ra ngoài (CSS, JS, image, font)    |
| `template/` | File template HTML hoặc template email dùng để render nội dung động |

---
## `deployments/` (hoặc `deploy/`)

Chứa toàn bộ cấu hình và script liên quan đến việc triển khai ứng dụng lên các môi trường khác nhau (CI/CD pipeline, infrastructure-as-code, environment config).

---
## Thông tin bổ sung

### 1. Phân biệt `dao/` và `repo/`

Hai thư mục này dễ gây nhầm lẫn vì cùng liên quan đến database:

- **DAO (Data Access Object)**: Pattern xuất phát từ Java EE. Mỗi DAO tương ứng trực tiếp với một bảng hoặc một loại thao tác SQL cụ thể. DAO không có khái niệm trừu tượng hóa nguồn dữ liệu.
- **Repository**: Pattern trong Domain-Driven Design (DDD). Repository trừu tượng hóa nguồn dữ liệu (database, cache, API ngoài) sau một interface, cho phép thay thế implementation mà không ảnh hưởng đến tầng service. Repository thường làm việc với domain object thay vì raw SQL.

Trong thực tế, nhiều dự án dùng lẫn lộn hai khái niệm. Nguyên tắc chung: nếu dự án theo DDD, ưu tiên dùng `repo/`; nếu project đơn giản hoặc dùng active-record pattern, `dao/` là đủ.

### 2. Phân biệt `model/` và `po/` và `domain/`

- **`model/`**: Thường dùng để ánh xạ với bảng database (database model / ORM model).
- **`po/` (Persistent Object)**: Tương đương `model/`, phổ biến trong hệ sinh thái Java (MyBatis, Hibernate). Nhấn mạnh rằng đây là object "sống" trong database, phân biệt với DTO hoặc VO.
- **`domain/`**: Trong DDD, domain object thể hiện khái niệm nghiệp vụ thuần túy, không nhất thiết phải khớp 1-1 với bảng database.

### 3. Chuẩn hóa cấu trúc thư mục theo kiến trúc

| Kiến trúc                | Cấu trúc thư mục đặc trưng                                  |
| ------------------------ | ----------------------------------------------------------- |
| MVC truyền thống         | `controllers/`, `models/`, `views/`, `routes/`              |
| Layered Architecture     | `controllers/`, `services/`, `repositories/`, `models/`     |
| Clean Architecture / DDD | `domain/`, `application/`, `infrastructure/`, `interfaces/` |
| Go standard layout       | `cmd/`, `internal/`, `pkg/`, `api/`, `configs/`, `scripts/` |

---
#### `domain` vs `models` vs `entity` trong Backend

Ba thư mục này **không hoàn toàn giống nhau**, nhưng ranh giới giữa chúng thường bị mờ tuỳ theo kiến trúc và convention của từng team. Hãy phân tích rõ:

#### Ý nghĩa gốc của từng khái niệm

| Thư mục  | Xuất phát từ               | Ý nghĩa cốt lõi                                               |
| -------- | -------------------------- | ------------------------------------------------------------- |
| `entity` | DDD, Clean Architecture    | Object ánh xạ trực tiếp với **database table**                |
| `domain` | DDD (Domain-Driven Design) | **Business logic + rules** của bài toán                       |
| `models` | MVC pattern                | Tầng dữ liệu chung, thường dùng để **transfer hoặc map** data |
#### Phân tích chi tiết

**`entity`** — _Database layer_

```go
// entity/user.go
// Ánh xạ 1-1 với table "users" trong DB
type User struct {
    ID        int64     `gorm:"primaryKey"`
    Email     string    `gorm:"uniqueIndex"`
    Password  string
    CreatedAt time.Time
}
```

**`domain`** — _Business layer_

```go
// domain/user.go
// Chứa business rules, không quan tâm DB lưu thế nào
type User struct {
    ID    int64
    Email Email  // Value Object, có validation riêng
    Role  Role
}

func (u *User) CanAccessResource(r Resource) bool { ... } // business logic
func (u *User) ChangeEmail(newEmail string) error  { ... } // business rule
```

**`models`** — _Thường là DTO / shared struct_

```go
// models/user.go
// Dùng để transfer data giữa các layer, hoặc bind request/response
type CreateUserRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"min=8"`
}

type UserResponse struct {
    ID    int64  `json:"id"`
    Email string `json:"email"`
}
```