# Go Project Setup - Kiến trúc Dự án Chuẩn

## 📋 Mục lục

1. [Standard Go Project Layout](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#1-standard-go-project-layout)
2. [Chi tiết từng thư mục](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#2-chi-ti%E1%BA%BFt-t%E1%BB%ABng-th%C6%B0-m%E1%BB%A5c)
3. [Kiến trúc phổ biến](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#3-ki%E1%BA%BFn-tr%C3%BAc-ph%E1%BB%95-bi%E1%BA%BFn)
4. [Dự án thực tế](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#4-d%E1%BB%B1-%C3%A1n-th%E1%BB%B1c-t%E1%BA%BF)
5. [Go Blueprint Tool](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#5-go-blueprint-tool)
6. [Best Practices](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#6-best-practices)

---

## 1. Standard Go Project Layout

### 1.1. Cấu trúc chuẩn

Đây là cấu trúc dự án được khuyến nghị bởi cộng đồng Go:

```
project-root/
├── api/              # API definitions (OpenAPI/Swagger, gRPC)
├── assets/           # Static assets (images, logos, fonts)
├── build/            # Build & packaging configs
├── cmd/              # Main applications
├── configs/          # Configuration files
├── deployments/      # Deployment configs (Docker, K8s)
├── docs/             # Documentation
├── examples/         # Example code
├── githooks/         # Git hooks
├── init/             # System init scripts
├── internal/         # Private application code
├── pkg/              # Public library code
├── scripts/          # Build/install/test scripts
├── test/             # Test data & E2E tests
├── third_party/      # Third-party code & tools
├── tools/            # Supporting tools
├── vendor/           # Vendored dependencies
├── web/              # Web application (frontend)
├── .gitignore        # Git ignore rules
├── LICENSE           # License file
├── README.md         # Project description
├── Makefile          # Build automation
└── go.mod            # Go module definition
```

**Link tham khảo:** [golang-standards/project-layout](https://github.com/golang-standards/project-layout)

### 1.2. Tên kiến trúc này là gì?

**❓ Câu hỏi:** "Nhưng mà hiện giờ tôi chưa biết tên kiến trúc thư mục này là gì?"

**💡 Trả lời:**

Kiến trúc này có nhiều tên gọi:

**1. Standard Go Project Layout**

- Tên chính thức nhất
- Được cộng đồng Go chấp nhận rộng rãi
- Tham khảo: `golang-standards/project-layout`

**2. Go Standard Package Layout**

- Tập trung vào cách tổ chức packages

**3. Domain-Driven Design (DDD) Layout**

- Khi kết hợp với DDD principles
- `internal/domain`, `internal/application`, `internal/infrastructure`

**4. Clean Architecture Layout**

- Khi áp dụng Clean Architecture
- Phân tách rõ layers: entities, use cases, interfaces, infrastructure

**5. Hexagonal Architecture (Ports & Adapters)**

- `internal/core` (business logic)
- `internal/ports` (interfaces)
- `internal/adapters` (implementations)

**Tóm lại:** Đây là **Standard Go Project Layout**, có thể kết hợp với các pattern khác tùy dự án.

---

## 2. Chi tiết từng thư mục

### 2.1. `/cmd` - Entry points

**Mục đích:** Chứa các main applications của dự án.

**Cấu trúc:**

```
cmd/
├── api/              # API server
│   └── main.go
├── cli/              # CLI tool
│   └── main.go
├── worker/           # Background worker
│   └── main.go
└── migrator/         # Database migrator
    └── main.go
```

**Ví dụ: `cmd/api/main.go`**

```go
package main

import (
    "log"
    "github.com/myproject/internal/server"
)

func main() {
    srv := server.New()
    if err := srv.Start(); err != nil {
        log.Fatal(err)
    }
}
```

**Best practices:**

- Mỗi app = 1 thư mục
- main.go ngắn gọn, chỉ khởi tạo
- Logic chính ở `internal/`

### 2.2. `/internal` - Private code

**Mục đích:** Code riêng của dự án, không export ra ngoài.

**⚠️ Quan trọng:** Go compiler sẽ **không cho phép** import package từ `/internal` của project khác.

**Cấu trúc:**

```
internal/
├── server/           # Server setup
│   ├── server.go
│   └── routes.go
├── handler/          # HTTP handlers
│   ├── user.go
│   └── product.go
├── service/          # Business logic
│   ├── user.go
│   └── product.go
├── repository/       # Data access
│   ├── user.go
│   └── product.go
├── model/            # Data models
│   ├── user.go
│   └── product.go
└── middleware/       # HTTP middleware
    ├── auth.go
    └── logging.go
```

**Ví dụ: Clean Architecture trong internal/**

```
internal/
├── domain/           # Entities (business objects)
│   ├── user.go
│   └── product.go
├── usecase/          # Use cases (business logic)
│   ├── user/
│   │   ├── create.go
│   │   └── update.go
│   └── product/
│       └── list.go
├── repository/       # Data access interfaces
│   ├── user.go
│   └── product.go
├── delivery/         # Delivery layer (HTTP, gRPC)
│   ├── http/
│   │   ├── handler/
│   │   └── middleware/
│   └── grpc/
└── infrastructure/   # External services
    ├── postgres/
    ├── redis/
    └── s3/
```

### 2.3. `/pkg` - Public libraries

**Mục đích:** Code có thể được **reuse** bởi projects khác.

**Cấu trúc:**

```
pkg/
├── logger/           # Logging utility
│   └── logger.go
├── validator/        # Validation helpers
│   └── validator.go
├── crypto/           # Crypto utilities
│   └── hash.go
└── httpclient/       # HTTP client wrapper
    └── client.go
```

**Ví dụ: `pkg/logger/logger.go`**

```go
package logger

import "go.uber.org/zap"

var log *zap.Logger

func Init() error {
    var err error
    log, err = zap.NewProduction()
    return err
}

func Info(msg string, fields ...zap.Field) {
    log.Info(msg, fields...)
}
```

**⚠️ Khác biệt `/internal` vs `/pkg`:**

|Aspect|`/internal`|`/pkg`|
|---|---|---|
|**Visibility**|Private|Public|
|**Import**|Chỉ trong project|Từ mọi project|
|**Purpose**|Business logic|Reusable utilities|
|**Example**|User service|Logger, validator|

### 2.4. `/api` - API definitions

**Mục đích:** API contracts, specs, schemas.

**Cấu trúc:**

```
api/
├── openapi/          # OpenAPI/Swagger specs
│   └── api.yaml
├── proto/            # Protocol Buffers
│   └── user.proto
├── graphql/          # GraphQL schemas
│   └── schema.graphql
└── http/             # HTTP API docs
    └── endpoints.md
```

**Ví dụ: `api/openapi/api.yaml`**

```yaml
openapi: 3.0.0
info:
  title: User API
  version: 1.0.0
paths:
  /users:
    get:
      summary: List users
      responses:
        '200':
          description: Success
```

### 2.5. `/configs` - Configuration files

**Mục đích:** Config templates và default configs.

**Cấu trúc:**

```
configs/
├── config.yaml       # Main config
├── config.dev.yaml   # Development
├── config.prod.yaml  # Production
└── database.yaml     # Database config
```

**Ví dụ: `configs/config.yaml`**

```yaml
server:
  host: localhost
  port: 8080
  
database:
  host: localhost
  port: 5432
  name: mydb
  
redis:
  host: localhost
  port: 6379
```

### 2.6. `/scripts` - Automation scripts

**Mục đích:** Build, install, test scripts.

**Cấu trúc:**

```
scripts/
├── build.sh          # Build script
├── test.sh           # Test script
├── migrate.sh        # Database migration
└── deploy.sh         # Deployment script
```

**Ví dụ: `scripts/build.sh`**

```bash
#!/bin/bash

echo "Building..."
go build -o bin/api cmd/api/main.go
echo "Build complete!"
```

### 2.7. `/deployments` - Deployment configs

**Mục đích:** Docker, Kubernetes, configs.

**Cấu trúc:**

```
deployments/
├── docker/
│   ├── Dockerfile
│   └── docker-compose.yml
├── kubernetes/
│   ├── deployment.yaml
│   ├── service.yaml
│   └── ingress.yaml
└── terraform/
    └── main.tf
```

### 2.8. `/migrations` - Database migrations

**Cấu trúc:**

```
migrations/
├── 000001_create_users_table.up.sql
├── 000001_create_users_table.down.sql
├── 000002_add_email_to_users.up.sql
└── 000002_add_email_to_users.down.sql
```

**Ví dụ: `000001_create_users_table.up.sql`**

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### 2.9. `/web` - Frontend code

**Mục đích:** Web application frontend.

**Cấu trúc:**

```
web/
├── static/           # Static files (CSS, JS, images)
│   ├── css/
│   ├── js/
│   └── images/
├── templates/        # HTML templates
│   ├── index.html
│   └── user.html
└── dist/             # Built files
```

---

## 3. Kiến trúc phổ biến

### 3.1. Layered Architecture (MVC-style)

```
internal/
├── controller/       # Controllers (HTTP handlers)
├── service/          # Business logic
├── repository/       # Data access
└── model/            # Data models
```

**Flow:**

```
Request → Controller → Service → Repository → Database
                ↓
            Response
```

### 3.2. Clean Architecture

```
internal/
├── domain/           # Entities (core business)
├── usecase/          # Application business rules
├── repository/       # Interface definitions
├── delivery/         # Delivery mechanisms (HTTP, gRPC)
└── infrastructure/   # External concerns (DB, cache)
```

**Dependency rule:** Inner layers không depend on outer layers.

### 3.3. Hexagonal Architecture

```
internal/
├── core/             # Business logic (domain + ports)
│   ├── domain/
│   └── ports/
└── adapters/         # Implementations
    ├── primary/      # Driving adapters (HTTP, CLI)
    └── secondary/    # Driven adapters (DB, external APIs)
```

### 3.4. Feature-based (Domain-driven)

```
internal/
├── user/             # User feature
│   ├── handler.go
│   ├── service.go
│   ├── repository.go
│   └── model.go
├── product/          # Product feature
│   ├── handler.go
│   ├── service.go
│   └── repository.go
└── order/            # Order feature
    └── ...
```

---

## 4. Dự án thực tế

### 4.1. Dự án phức tạp như thế nào?

**❓ Câu hỏi:** "Liệu đối với dự án thực tế thì người ta thiết kế nó phức tạp như thế nào?"

**💡 Trả lời:**

**Small Project (1-3 developers):**

```
myproject/
├── cmd/api/
├── internal/
│   ├── handler/
│   ├── service/
│   ├── repository/
│   └── model/
├── pkg/
├── configs/
└── go.mod
```

**Medium Project (5-10 developers):**

```
myproject/
├── cmd/
│   ├── api/
│   ├── worker/
│   └── cli/
├── internal/
│   ├── user/           # Feature modules
│   │   ├── handler/
│   │   ├── service/
│   │   ├── repository/
│   │   └── model/
│   ├── product/
│   └── order/
├── pkg/
│   ├── logger/
│   ├── validator/
│   └── middleware/
├── api/
├── configs/
├── migrations/
├── deployments/
└── scripts/
```

**Large Project (10+ developers):**

```
myproject/
├── cmd/
│   ├── api-gateway/
│   ├── user-service/
│   ├── product-service/
│   ├── order-service/
│   └── notification-worker/
├── internal/
│   ├── shared/         # Shared code
│   │   ├── middleware/
│   │   └── validator/
│   ├── user/           # Microservice
│   │   ├── domain/
│   │   ├── usecase/
│   │   ├── repository/
│   │   └── delivery/
│   ├── product/
│   └── order/
├── pkg/
│   ├── logger/
│   ├── tracer/
│   ├── metrics/
│   └── httpclient/
├── api/
│   ├── proto/          # gRPC
│   └── openapi/        # REST
├── configs/
├── deployments/
│   ├── docker/
│   └── kubernetes/
├── migrations/
├── scripts/
└── test/
    ├── e2e/
    └── integration/
```

**Enterprise Project (Microservices):**

```
organization/
├── services/
│   ├── user-service/
│   │   ├── cmd/
│   │   ├── internal/
│   │   └── go.mod
│   ├── product-service/
│   │   └── ...
│   └── order-service/
│       └── ...
├── libraries/          # Shared libraries
│   ├── go-logger/
│   ├── go-validator/
│   └── go-auth/
├── api/                # API gateway
└── infrastructure/     # Shared infrastructure
    ├── k8s/
    └── terraform/
```

### 4.2. Làm quen với dự án có sẵn

**💡 Trả lời - 7 bước làm quen:**

**Bước 1: Đọc README.md**

```markdown
# Project Name

## Architecture
- Clean Architecture
- PostgreSQL + Redis
- Docker + K8s

## Getting Started
1. Install dependencies: `go mod download`
2. Setup database: `make migrate-up`
3. Run: `go run cmd/api/main.go`
```

**Bước 2: Hiểu entry point**

```
1. Tìm cmd/api/main.go (hoặc cmd/server/main.go)
2. Đọc hàm main() để hiểu flow khởi tạo
3. Follow các function calls
```

**Bước 3: Vẽ diagram**

```
┌─────────────┐
│   main()    │
└──────┬──────┘
       │
       ├─ Load Config
       ├─ Init Database
       ├─ Init Redis
       ├─ Setup Router
       └─ Start Server
```

**Bước 4: Trace một API request**

```
1. Chọn 1 endpoint đơn giản (GET /users)
2. Trace: router → handler → service → repository
3. Vẽ flow diagram
```

**Bước 5: Đọc tests**

```go
// Test files thường dễ hiểu hơn
func TestCreateUser(t *testing.T) {
    // Setup
    // Action
    // Assert
}
```

**Bước 6: Check dependencies**

```bash
# Xem libraries đang dùng
cat go.mod

# Thường gặp:
- gin-gonic/gin (HTTP framework)
- gorm.io/gorm (ORM)
- go-redis/redis (Redis)
- uber-go/zap (Logger)
```

**Bước 7: Chạy và debug**

```bash
# 1. Setup môi trường
make setup

# 2. Chạy tests
make test

# 3. Chạy local
make run

# 4. Debug với Delve
dlv debug cmd/api/main.go
```

**Checklist làm quen:**

- [ ] Đọc README.md
- [ ] Hiểu cấu trúc thư mục
- [ ] Trace từ main.go
- [ ] Vẽ architecture diagram
- [ ] Trace 1 API request
- [ ] Đọc tests
- [ ] Check dependencies
- [ ] Chạy được local
- [ ] Debug một feature
- [ ] Implement một feature nhỏ

---

## 5. Go Blueprint Tool

### 5.1. Giới thiệu

![Go Blueprint](https://docs.go-blueprint.dev/public/blueprint_1.png)

**Go Blueprint** là tool tự động tạo project structure chuẩn.

**Website:** [go-blueprint.dev](https://go-blueprint.dev/)

**Features:**

- ✅ Chọn framework (Gin, Fiber, Chi, Echo)
- ✅ Chọn database (PostgreSQL, MySQL, SQLite)
- ✅ Docker support
- ✅ CI/CD setup
- ✅ Advanced features (WebSocket, gRPC)

### 5.2. Cài đặt

```bash
go install github.com/melkeydev/go-blueprint@latest
```

### 5.3. Sử dụng

**Cách 1: Interactive mode**

```bash
go-blueprint create
```

**Cách 2: CLI flags**

```bash
go-blueprint create \
  --name myproject \
  --framework gin \
  --driver postgres \
  --advanced \
  --feature websocket \
  --feature docker \
  --git commit
```

**Kết quả:**

```
myproject/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── server/
│   │   ├── routes.go
│   │   └── server.go
│   └── database/
│       └── database.go
├── .air.toml         # Hot reload config
├── .env
├── .gitignore
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

### 5.4. Cấu trúc Blueprint tạo ra

**cmd/api/main.go:**

```go
package main

import (
    "github.com/myproject/internal/server"
)

func main() {
    server := server.New()
    err := server.ListenAndServe()
    if err != nil {
        panic(err)
    }
}
```

**internal/server/server.go:**

```go
package server

import (
    "github.com/gin-gonic/gin"
)

type Server struct {
    router *gin.Engine
}

func New() *Server {
    s := &Server{
        router: gin.Default(),
    }
    s.RegisterRoutes()
    return s
}

func (s *Server) ListenAndServe() error {
    return s.router.Run(":8080")
}
```

**internal/server/routes.go:**

```go
package server

func (s *Server) RegisterRoutes() {
    s.router.GET("/health", s.healthHandler)
}

func (s *Server) healthHandler(c *gin.Context) {
    c.JSON(200, gin.H{"status": "ok"})
}
```

### 5.5. Makefile commands

```makefile
.PHONY: build run test clean

build:
	go build -o bin/api cmd/api/main.go

run:
	go run cmd/api/main.go

test:
	go test -v ./...

clean:
	rm -rf bin/
```

---

## 6. Best Practices

### 6.1. Tổ chức code

**✅ Do:**

- Tách concerns rõ ràng (handler, service, repository)
- Một package một responsibility
- File names match với content (user.go cho User struct)
- Internal cho private code
- Pkg cho reusable utilities

**❌ Don't:**

- Tất cả code trong main package
- Circular dependencies
- Deep nesting (>4 levels)
- God packages (package có quá nhiều code)

### 6.2. Naming conventions

```go
// Package names: short, lowercase, no underscores
package user

// Files: lowercase, underscores OK
user_service.go
user_repository.go

// Interfaces: -er suffix
type UserService interface {}
type UserRepository interface {}

// Implementations
type userService struct {}
type postgresUserRepository struct {}
```

### 6.3. Dependency injection

```go
// ✅ Good - Interface-based
type UserService struct {
    repo UserRepository  // Interface
}

func NewUserService(repo UserRepository) *UserService {
    return &UserService{repo: repo}
}

// ❌ Bad - Concrete dependency
type UserService struct {
    repo *PostgresUserRepository  // Concrete
}
```

### 6.4. Configuration

```go
// ✅ Good - Struct-based config
type Config struct {
    Server   ServerConfig
    Database DatabaseConfig
    Redis    RedisConfig
}

func LoadConfig() (*Config, error) {
    // Load from env, file, etc.
}

// ❌ Bad - Global variables
var DatabaseHost string
var DatabasePort int
```

### 6.5. Error handling

```go
// ✅ Good - Wrap errors with context
func (s *UserService) Create(user *User) error {
    if err := s.repo.Create(user); err != nil {
        return fmt.Errorf("failed to create user: %w", err)
    }
    return nil
}

// ❌ Bad - Swallow errors
func (s *UserService) Create(user *User) error {
    s.repo.Create(user)  // Ignore error
    return nil
}
```

---

## 📚 Tổng kết

### Kiến trúc chuẩn

```
✅ cmd/         - Main applications
✅ internal/    - Private code
✅ pkg/         - Public libraries
✅ api/         - API definitions
✅ configs/     - Config files
✅ scripts/     - Automation
✅ deployments/ - Deploy configs
```

### Patterns phổ biến

1. **Layered** - Controller → Service → Repository
2. **Clean Architecture** - Domain → Usecase → Delivery
3. **Hexagonal** - Core → Ports → Adapters
4. **Feature-based** - Organize by features

### Tools

- **go-blueprint** - Project generator
- **make** - Build automation
- **air** - Hot reload
- **docker** - Containerization

### Làm quen dự án mới

1. README.md
2. Entry point (main.go)
3. Architecture diagram
4. Trace API request
5. Read tests
6. Run local
7. Debug

---

_Good architecture makes development easier and faster!_