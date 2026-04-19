# Kiến trúc Monolithic (Monolithic Architecture)

## **Monolithic là gì?**

**Monolithic** (Nguyên khối) là một phương pháp **thiết kế** kiến trúc theo kiểu **truyền thống và thống nhất**. Trong đó, tất cả các chức năng hoạt động như một **đơn vị duy nhất, không thể tách rời**.

Một ứng dụng Monolithic bao gồm:

- **Giao diện người dùng (UI)** - User Interface
- **Logic nghiệp vụ** - Business Logic điều khiển cốt lõi
- **Lớp truy cập dữ liệu** - Data Access Layer chịu **trách nhiệm** giao tiếp với cơ sở dữ liệu

### Sơ đồ kiến trúc 3 tầng:

```
┌─────────────────────────────────────────┐
│     MONOLITHIC APPLICATION              │
│  ┌───────────────────────────────────┐  │
│  │     User Interface (UI)           │  │
│  │   - Web Pages / API Endpoints     │  │
│  └───────────────────────────────────┘  │
│                  ↓                      │
│  ┌───────────────────────────────────┐  │
│  │     Business Logic Layer          │  │
│  │  • Authentication                 │  │
│  │  • Business Rules                 │  │
│  │  • Data Validation                │  │
│  │  • Core Processing                │  │
│  └───────────────────────────────────┘  │
│                  ↓                      │
│  ┌───────────────────────────────────┐  │
│  │     Data Access Layer             │  │
│  │   - Repository Pattern            │  │
│  │   - Database Queries              │  │
│  └───────────────────────────────────┘  │
│                  ↓                      │
│  ┌───────────────────────────────────┐  │
│  │         Database                  │  │
│  │   (PostgreSQL, MySQL, MongoDB)    │  │
│  └───────────────────────────────────┘  │
└─────────────────────────────────────────┘
      Deploy as ONE unit (single binary)
```

Trong mô hình này, các module hoạt động một cách **liên kết chặt chẽ**, phụ thuộc vào nhau, và mọi sự thay đổi hay bảo trì đều được thực hiện trên **cùng một đơn vị**.

---

## **Các đặc điểm của Monolithic**

### 1. **Tự cung - Tự cấp (Self-contained)**

- Các ứng dụng Monolithic được thiết kế để **hoạt động độc lập**
- Thường giảm thiểu sự phụ thuộc với các hệ thống bên ngoài
- Tất cả tính năng tích hợp trong một ứng dụng duy nhất

### 2. **Liên kết chặt chẽ (Tightly Coupled)**

- Các thành phần bên trong được kết nối với nhau một cách phức tạp
- Thay đổi ở một module có thể gây ra **hiệu ứng dây chuyền** trên toàn bộ hệ thống
- Khó tách rời các thành phần để phát triển độc lập

### 3. **Mã nguồn duy nhất (Single Codebase)**

- Toàn bộ mã nguồn của ứng dụng được **tập trung hóa** trong một repository
- Cho phép phát triển cộng tác trong một môi trường **duy nhất**, dùng chung
- Đây là **đặc điểm chính** của kiến trúc phần mềm Monolithic

### 4. **Triển khai đơn giản (Single Deployment Unit)**

- Toàn bộ ứng dụng được build và deploy như **một đơn vị duy nhất**
- Chỉ cần một quy trình deployment cho toàn bộ hệ thống
- Một binary file hoặc một container duy nhất

### 5. **Chia sẻ tài nguyên**

- Tất cả modules sử dụng chung: Memory, CPU, Database connection pool, Framework

---

## **Ứng dụng Monolithic hoạt động như thế nào?**

### **Luồng xử lý request:**

```
User → UI → Business Logic → Data Access → Database
  ↑                                            │
  └────────────── Response ───────────────────┘
```

### **Chi tiết từng tầng:**

### **1. User Interface (UI)**

Khi người dùng tương tác với ứng dụng (nhấn nút, điền form), yêu cầu sẽ được:

- **Đóng gói** thành HTTP request
- **Route** đến handler tương ứng
- **Chuyển tiếp** đến Business Logic để xử lý

**Ví dụ:** Người dùng gửi POST request để đăng nhập

```
POST /api/auth/login
Body: {
  "email": "user@example.com",
  "password": "123456"
}
```

### **2. Logic nghiệp vụ (Business Logic Layer)**

Lớp này là **"bộ não"** của kiến trúc Monolithic, chứa:

- Quy tắc nghiệp vụ phức tạp
- Các phép tính và thuật toán
- Quy trình đưa ra quyết định

**Các thao tác quan trọng:**

#### **a) Xác thực dữ liệu đầu vào (Input Validation)**

- Đảm bảo dữ liệu phù hợp với yêu cầu
- Kiểm tra format, độ dài, kiểu dữ liệu
- **Ví dụ:** Email đúng format? Password đủ mạnh?

#### **b) Thực hiện các phép tính (Processing)**

- Xử lý dựa trên yêu cầu của người dùng
- Áp dụng business rules
- **Ví dụ:** Tính tổng giá trị đơn hàng, áp dụng giảm giá

#### **c) Triển khai Logic rẽ nhánh (Conditional Logic)**

- Đưa ra quyết định dựa trên điều kiện
- **Ví dụ:** Nếu user là VIP → giảm 20%, thường → giảm 10%

#### **d) Phối hợp với Lớp Dữ liệu**

- Gửi và nhận thông tin từ Data Access Layer
- **Ví dụ:** Lấy thông tin user, kiểm tra password, cập nhật last_login

### **3. Lớp truy cập dữ liệu (Data Access Layer)**

Lớp này chịu trách nhiệm:

- **Tương tác với cơ sở dữ liệu**
- **Thực hiện các thao tác CRUD:**
    - **C**reate (INSERT) - Tạo mới dữ liệu
    - **R**ead (SELECT) - Đọc/truy vấn dữ liệu
    - **U**pdate - Cập nhật dữ liệu
    - **D**elete - Xóa dữ liệu

**Vai trò:**

- Trừu tượng hóa database (Business Logic không cần biết chi tiết SQL)
- Quản lý kết nối (Connection pooling)
- Mapping dữ liệu (ORM - Object-Relational Mapping)

---

## **Ví dụ thực tế với Golang**

### **Cấu trúc thư mục Monolithic App:**

```
ecommerce-monolith/
├── cmd/
│   └── server/
│       └── main.go              # Entry point
│
├── internal/
│   ├── handlers/                # HTTP Handlers (UI Layer)
│   │   ├── auth_handler.go
│   │   ├── product_handler.go
│   │   └── order_handler.go
│   │
│   ├── services/                # Business Logic
│   │   ├── auth_service.go
│   │   ├── product_service.go
│   │   └── order_service.go
│   │
│   ├── repositories/            # Data Access Layer
│   │   ├── user_repository.go
│   │   ├── product_repository.go
│   │   └── order_repository.go
│   │
│   └── models/                  # Domain Models
│       ├── user.go
│       ├── product.go
│       └── order.go
│
├── pkg/
│   └── database/
│       └── postgres.go
│
├── go.mod
└── go.sum
```

### **Code Example - E-commerce Monolith:**

#### **1. Models (Domain)**

```go
// internal/models/user.go
package models

import "time"

type User struct {
    ID        int64     `json:"id"`
    Email     string    `json:"email"`
    Password  string    `json:"-"` // Không trả về trong JSON
    Name      string    `json:"name"`
    CreatedAt time.Time `json:"created_at"`
}

// internal/models/product.go
package models

type Product struct {
    ID          int64   `json:"id"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Price       float64 `json:"price"`
    Stock       int     `json:"stock"`
}

// internal/models/order.go
package models

type Order struct {
    ID          int64     `json:"id"`
    UserID      int64     `json:"user_id"`
    TotalAmount float64   `json:"total_amount"`
    Status      string    `json:"status"` // pending, paid, shipped
    CreatedAt   time.Time `json:"created_at"`
}
```

#### **2. Repository Layer (Data Access)**

```go
// internal/repositories/user_repository.go
package repositories

import (
    "database/sql"
    "ecommerce/internal/models"
)

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{db: db}
}

// FindByEmail - Tìm user theo email
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
    user := &models.User{}
    query := `
        SELECT id, email, password, name, created_at
        FROM users
        WHERE email = $1
    `
    err := r.db.QueryRow(query, email).Scan(
        &user.ID,
        &user.Email,
        &user.Password,
        &user.Name,
        &user.CreatedAt,
    )
    if err != nil {
        return nil, err
    }
    return user, nil
}

// Create - Tạo user mới
func (r *UserRepository) Create(user *models.User) error {
    query := `
        INSERT INTO users (email, password, name)
        VALUES ($1, $2, $3)
        RETURNING id, created_at
    `
    return r.db.QueryRow(query, user.Email, user.Password, user.Name).
        Scan(&user.ID, &user.CreatedAt)
}
```

#### **3. Service Layer (Business Logic)**

```go
// internal/services/auth_service.go
package services

import (
    "errors"
    "ecommerce/internal/models"
    "ecommerce/internal/repositories"
    "golang.org/x/crypto/bcrypt"
)

type AuthService struct {
    userRepo *repositories.UserRepository
}

func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
    return &AuthService{userRepo: userRepo}
}

// Register - Đăng ký user mới
func (s *AuthService) Register(email, password, name string) (*models.User, error) {
    // 1. Validate input
    if email == "" || password == "" || name == "" {
        return nil, errors.New("tất cả fields là bắt buộc")
    }
    
    if len(password) < 8 {
        return nil, errors.New("password phải có ít nhất 8 ký tự")
    }
    
    // 2. Check email đã tồn tại chưa
    existingUser, _ := s.userRepo.FindByEmail(email)
    if existingUser != nil {
        return nil, errors.New("email đã được sử dụng")
    }
    
    // 3. Hash password
    hashedPassword, err := bcrypt.GenerateFromPassword(
        []byte(password), 
        bcrypt.DefaultCost,
    )
    if err != nil {
        return nil, err
    }
    
    // 4. Tạo user mới
    user := &models.User{
        Email:    email,
        Password: string(hashedPassword),
        Name:     name,
    }
    
    if err := s.userRepo.Create(user); err != nil {
        return nil, err
    }
    
    return user, nil
}

// Login - Đăng nhập
func (s *AuthService) Login(email, password string) (*models.User, error) {
    // 1. Tìm user
    user, err := s.userRepo.FindByEmail(email)
    if err != nil {
        return nil, errors.New("email hoặc password không đúng")
    }
    
    // 2. Verify password
    err = bcrypt.CompareHashAndPassword(
        []byte(user.Password), 
        []byte(password),
    )
    if err != nil {
        return nil, errors.New("email hoặc password không đúng")
    }
    
    return user, nil
}
```

#### **4. Handler Layer (UI - HTTP)**

```go
// internal/handlers/auth_handler.go
package handlers

import (
    "encoding/json"
    "net/http"
    "ecommerce/internal/services"
)

type AuthHandler struct {
    authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
    return &AuthHandler{authService: authService}
}

type RegisterRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
    Name     string `json:"name"`
}

// Register - HTTP Handler
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
    var req RegisterRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }
    
    user, err := h.authService.Register(req.Email, req.Password, req.Name)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Đăng ký thành công",
        "user":    user,
    })
}
```

#### **5. Main Entry Point**

```go
// cmd/server/main.go
package main

import (
    "database/sql"
    "log"
    "net/http"
    
    "ecommerce/internal/handlers"
    "ecommerce/internal/repositories"
    "ecommerce/internal/services"
    
    "github.com/gorilla/mux"
    _ "github.com/lib/pq"
)

func main() {
    // 1. Connect to database
    db, err := sql.Open("postgres", 
        "host=localhost user=postgres password=secret dbname=ecommerce sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    // 2. Initialize repositories
    userRepo := repositories.NewUserRepository(db)
    
    // 3. Initialize services
    authService := services.NewAuthService(userRepo)
    
    // 4. Initialize handlers
    authHandler := handlers.NewAuthHandler(authService)
    
    // 5. Setup routes
    router := mux.NewRouter()
    router.HandleFunc("/api/auth/register", authHandler.Register).Methods("POST")
    router.HandleFunc("/api/auth/login", authHandler.Login).Methods("POST")
    
    // 6. Start server
    log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
```

### **Luồng hoạt động:**

```
1. User gửi request:
   POST /api/auth/register
   {"email": "john@example.com", "password": "123456", "name": "John"}

2. Router → AuthHandler.Register()

3. AuthHandler → AuthService.Register()
   - Validate input
   - Check email exists
   - Hash password
   - Call UserRepository.Create()

4. UserRepository → Execute SQL:
   INSERT INTO users (email, password, name) VALUES (...)

5. Database → Trả về ID và created_at

6. Response trả về user:
   {"message": "Đăng ký thành công", "user": {...}}
```

---

## **Ưu điểm của Monolithic**

### 1. **Đơn giản trong phát triển ban đầu**

- ✅ Dễ dàng bắt đầu dự án mới
- ✅ Không cần kiến trúc phức tạp
- ✅ Setup nhanh chóng
- ✅ Phù hợp với startup và team nhỏ

### 2. **Dễ dàng testing**

- ✅ Có thể test toàn bộ ứng dụng cùng lúc
- ✅ End-to-end testing đơn giản
- ✅ Không cần mock external services

### 3. **Deployment đơn giản**

- ✅ Chỉ deploy **một artifact duy nhất**
- ✅ Không cần orchestration phức tạp (Kubernetes)
- ✅ Rollback dễ dàng

**Ví dụ:**

```bash
# Build
go build -o app ./cmd/server

# Deploy
./app

# Hoặc với Docker
docker build -t myapp:v1 .
docker run -p 8080:8080 myapp:v1
```

### 4. **Hiệu suất tốt**

- ✅ Giao tiếp qua **in-memory function calls**
- ✅ Nhanh hơn network calls (Microservices)
- ✅ Không có network latency overhead

**So sánh:**

```
Monolithic:
OrderService.CreateOrder() → PaymentService.Process()
└─> In-memory call (~nanoseconds)

Microservices:
OrderService → HTTP POST → PaymentService
└─> Network call (~milliseconds)
```

### 5. **Chi phí thấp ban đầu**

- ✅ Không cần infrastructure phức tạp
- ✅ Ít server hơn (1-2 servers)
- ✅ Không cần API Gateway, Service Mesh

---

## **Nhược điểm của Monolithic**

### 1. **Khó mở rộng (Scalability)**

- ❌ Phải scale **toàn bộ ứng dụng**, không thể scale riêng một module
- ❌ Lãng phí tài nguyên

**Ví dụ:**

```
Module Upload ảnh bị quá tải (CPU intensive)
Module Xem danh sách hoạt động bình thường

Monolithic: Phải scale TOÀN BỘ app
Microservices: Chỉ scale service Upload
```

### 2. **Deployment chậm và rủi ro cao**

- ❌ Mỗi lần deploy phải build **toàn bộ ứng dụng**
- ❌ Một lỗi nhỏ có thể làm sập toàn bộ hệ thống
- ❌ Downtime dài hơn

### 3. **Khó bảo trì khi ứng dụng lớn**

- ❌ Codebase ngày càng **phình to**
- ❌ Khó hiểu và khó sửa đổi
- ❌ Onboarding developer mới mất nhiều thời gian
- ❌ **"Big Ball of Mud"** - code lộn xộn

### 4. **Liên kết chặt chẽ (Tight Coupling)**

- ❌ Các modules phụ thuộc lẫn nhau
- ❌ Thay đổi một chỗ ảnh hưởng nhiều nơi
- ❌ Khó refactor

### 5. **Technology Lock-in**

- ❌ Toàn bộ app phải dùng **cùng một stack**
- ❌ Khó áp dụng công nghệ mới cho một phần cụ thể

---

## **Khi nào nên sử dụng Monolithic?**

### ✅ **NÊN dùng khi:**

1. **Dự án vừa và nhỏ**
    
    - Startup mới bắt đầu
    - MVP (Minimum Viable Product)
    - Timeline ngắn (< 6 tháng)
2. **Team nhỏ (< 10 người)**
    
    - Dễ coordination
    - Mọi người làm việc trên cùng codebase
3. **Yêu cầu đơn giản, ít thay đổi**
    
    - Business logic ổn định
    - Không cần scale thường xuyên
4. **Budget hạn chế**
    
    - Cần triển khai nhanh
    - Không có DevOps team

### ❌ **KHÔNG NÊN dùng khi:**

1. **Ứng dụng lớn, phức tạp**
    
    - Nhiều modules độc lập
    - Codebase > 100K lines
2. **Cần scale linh hoạt**
    
    - Traffic không đều giữa các modules
    - Load patterns khác nhau
3. **Team lớn (> 15-20 người)**
    
    - Nhiều team làm việc song song
    - Cần autonomous teams
4. **Cần deploy thường xuyên**
    
    - Multiple releases per day
    - Continuous deployment

---

## **So sánh Monolithic vs Microservices**

|Tiêu chí|Monolithic|Microservices|
|---|---|---|
|**Cấu trúc**|Một ứng dụng duy nhất|Nhiều services nhỏ|
|**Deployment**|Deploy toàn bộ cùng lúc|Deploy từng service riêng|
|**Scalability**|Scale toàn bộ app|Scale từng service|
|**Technology**|Một stack duy nhất|Mỗi service khác stack|
|**Database**|Shared database|Database per service|
|**Complexity**|Đơn giản ban đầu|Phức tạp ngay từ đầu|
|**Team Size**|Phù hợp team nhỏ|Phù hợp team lớn|
|**Communication**|In-memory calls|HTTP/gRPC calls|
|**Cost**|Thấp ban đầu|Cao hơn|

---

## **Kết luận**

### **Key Takeaways:**

1. **Monolithic không phải là "kiến trúc xấu"**
    
    - Phù hợp với nhiều use cases
    - Nhiều công ty thành công vẫn sử dụng
2. **Start simple, scale when needed**
    
    - Bắt đầu với Monolithic cho hầu hết dự án mới
    - Chuyển sang Microservices khi thực sự cần
3. **Trade-offs:**
    
    - Đơn giản ⟷ Scalability
    - Development speed ⟷ Flexibility
    - Low cost ⟷ High availability

### **Lời khuyên:**

> **"Start with a monolith, split when you need to"** — Martin Fowler

**Khi nào chuyển sang Microservices?**

- Team > 15-20 người
- Codebase > 100K lines
- Cần scale khác nhau cho từng module
- Có đủ DevOps expertise và budget