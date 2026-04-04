# Middleware trong Backend và Go/Gin

## Mục lục

1. [Middleware là gì?](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#middleware-l%C3%A0-g%C3%AC)
2. [Phân loại Middleware](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#ph%C3%A2n-lo%E1%BA%A1i-middleware)
3. [Tại sao phải dùng Middleware?](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#t%E1%BA%A1i-sao-ph%E1%BA%A3i-d%C3%B9ng-middleware)
4. [Middleware trong Gin — Cơ bản](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#middleware-trong-gin--c%C6%A1-b%E1%BA%A3n)
5. [Các cách áp dụng Middleware trong Gin](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#c%C3%A1c-c%C3%A1ch-%C3%A1p-d%E1%BB%A5ng-middleware-trong-gin)
6. [Thứ tự Middleware — Giải thích từng lớp](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#th%E1%BB%A9-t%E1%BB%B1-middleware--gi%E1%BA%A3i-th%C3%ADch-t%E1%BB%ABng-l%E1%BB%9Bp)
7. [Case Study: Phối hợp nhiều Middleware](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#case-study-ph%E1%BB%91i-h%E1%BB%A3p-nhi%E1%BB%81u-middleware)
8. [Cách viết Middleware thanh lịch hơn](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#c%C3%A1ch-vi%E1%BA%BFt-middleware-thanh-l%E1%BB%8Bch-h%C6%A1n)
9. [Tổng kết](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#t%E1%BB%95ng-k%E1%BA%BFt)

---

## Middleware là gì?

**Middleware** (phần mềm trung gian) là lớp xử lý nằm **giữa Request và Response**, có nhiệm vụ tiếp nhận và xử lý thông tin trước khi chuyển tiếp đến các Handler hoặc Middleware tiếp theo.

![](https://200lab.io/blog/_next/image?url=https%3A%2F%2Fstatics.cdn.200lab.io%2F2022%2F06%2Fmiddleware-request-response-200lab.png&w=640&q=75)

```
Client
  ↓ gửi request
┌─────────────────────────────────────────┐
│  Middleware 1 (Logger)                  │
│    ↓                                    │
│  Middleware 2 (Authentication)          │
│    ↓                                    │
│  Middleware 3 (Authorization)           │
│    ↓                                    │
│  Handler (Business Logic)               │
│    ↓                                    │
│  Middleware 3 (ghi kết quả)             │
│    ↓                                    │
│  Middleware 2 (ghi kết quả)             │
│    ↓                                    │
│  Middleware 1 (ghi log thời gian)       │
└─────────────────────────────────────────┘
  ↓ trả response
Client
```

> 💡 **Đặc điểm quan trọng:** Middleware hoạt động theo mô hình **"vào-ra"** (onion model — mô hình hành tây). Mỗi middleware có thể chạy code **trước** khi request đi vào handler, và **sau** khi handler đã xử lý xong. Thứ tự vào và ra là ngược nhau.

Trong các **hệ thống phân tán**, middleware còn đóng vai trò là các service trung gian kết nối các thành phần với nhau (message queue, API gateway, service mesh...).

---

## Phân loại Middleware

Theo Wikipedia, có **4 loại middleware** chính. Dưới đây là giải thích cụ thể từng loại kèm ví dụ thực tế:

### 1. Transactional Middleware — Middleware giao dịch

**Là gì:** Xử lý các nhóm thao tác phải thực hiện như một đơn vị — tất cả thành công hoặc tất cả thất bại (ACID).

**Hoạt động ra sao:**

```
Giao dịch chuyển tiền:
  Bước 1: Trừ 1.000.000đ từ tài khoản A
  Bước 2: Cộng 1.000.000đ vào tài khoản B
  → Nếu Bước 2 lỗi: tự động rollback Bước 1
  → Đảm bảo không mất tiền giữa chừng
```

**Khi nào áp dụng:** Hệ thống ngân hàng, thanh toán, đặt vé, bất kỳ nơi nào cần tính nhất quán dữ liệu.

---

### 2. Message Middleware — Middleware nhắn tin

**Là gì:** Kết nối các hệ thống với nhau qua hàng đợi tin nhắn (message queue). Bên gửi và bên nhận không cần biết nhau, không cần hoạt động cùng lúc.

**Hoạt động ra sao:**

```
Hệ thống đặt hàng:
  Service A (Order) → [Queue] → Service B (Email) gửi xác nhận
                              → Service C (Inventory) cập nhật kho
                              → Service D (Analytics) ghi thống kê
  → Nếu Service B đang bận → message chờ trong queue
  → Không cần A chờ B xử lý xong mới tiếp tục
```

**Khi nào áp dụng:** Hệ thống microservices, gửi email/notification, xử lý ảnh, bất kỳ tác vụ nào có thể xử lý bất đồng bộ. Ví dụ: RabbitMQ, Apache Kafka, AWS SQS.

---

### 3. Procedural Middleware — Middleware thủ tục

**Là gì:** Cho phép một chương trình gọi hàm/thủ tục trên máy khác (remote) như thể đang gọi hàm cục bộ (local). Che giấu sự phức tạp của giao tiếp mạng.

**Hoạt động ra sao:**

```
Service A muốn lấy thông tin user từ Service B:
  // Code trông như gọi hàm bình thường
  user := userService.GetUser(id)
  // Nhưng thực ra bên dưới:
  // A → (serialize) → mạng → (deserialize) → B → xử lý → (serialize) → mạng → A
```

**Khi nào áp dụng:** Giao tiếp giữa các service trong microservices. Ví dụ: gRPC, REST API, GraphQL.

---

### 4. Object-oriented Middleware — Middleware hướng đối tượng

**Là gì:** Tương tự Procedural nhưng áp dụng các nguyên tắc OOP — các thành phần giao tiếp qua đối tượng, kế thừa, và exceptions thay vì các hàm thủ tục đơn thuần.

**Hoạt động ra sao:**

```
// Thay vì gọi hàm đơn thuần (Procedural):
result = remoteCall("getUserById", id)

// Object-oriented: giao tiếp qua đối tượng với đầy đủ tính chất OOP
UserService proxy = getRemoteObject(UserService.class)
User user = proxy.getById(id)  // Kế thừa interface UserService
// Exceptions từ remote được truyền về như local exception
```

**Khi nào áp dụng:** Enterprise Java (EJB, CORBA), .NET Remoting, các hệ thống doanh nghiệp lớn cũ.

---

> 💡 **Tóm tắt để dễ nhớ:**
> 
> |Loại|Từ khóa|Ví dụ thực tế|
> |---|---|---|
> |Transactional|"Tất cả hoặc không có gì"|Ngân hàng, thanh toán|
> |Message|"Hàng đợi, bất đồng bộ"|Kafka, RabbitMQ|
> |Procedural|"Gọi hàm từ xa như cục bộ"|gRPC, REST|
> |Object-oriented|"Procedural + OOP"|EJB, CORBA|
> 
> Trong lập trình web/API hàng ngày, khi nói "middleware" ta thường chỉ loại **Procedural** — các hàm xử lý nằm trong pipeline request/response của web framework như Gin, Express, Laravel.

---

## Tại sao phải dùng Middleware?

Hãy hình dung một API server có 20 endpoint. Mỗi endpoint cần:

- Kiểm tra token xác thực
- Ghi log request
- Kiểm tra quyền truy cập
- Parse dữ liệu từ body

**Không dùng middleware:** Phải lặp lại 4 đoạn code trên ở **20 chỗ** — 80 đoạn code giống nhau. Muốn sửa logic logging → phải sửa 20 file.

**Dùng middleware:** Viết 4 middleware một lần, áp dụng cho tất cả endpoint. Muốn sửa → chỉ sửa 1 chỗ.

```
Không dùng middleware:              Dùng middleware:
────────────────────────            ─────────────────────────────
endpoint1:                          [Logger MW] ──→ tất cả endpoint
  checkToken()    ← lặp lại         [Auth MW]   ──→ tất cả endpoint
  log()           ← lặp lại         [CORS MW]   ──→ tất cả endpoint
  checkRole()     ← lặp lại              ↓
  businessLogic()                    endpoint1: chỉ có businessLogic()
                                     endpoint2: chỉ có businessLogic()
endpoint2:                           endpoint3: chỉ có businessLogic()
  checkToken()    ← lặp lại
  log()           ← lặp lại
  ...
```

**Nguyên tắc:** Middleware giúp áp dụng nguyên tắc **DRY (Don't Repeat Yourself)** cho các logic xuyên suốt (cross-cutting concerns).

---

## Middleware trong Gin — Cơ bản

### Định nghĩa

Trong Gin, mọi middleware đều là một hàm kiểu `gin.HandlerFunc`:

```go
type HandlerFunc func(*Context)
```

Middleware thực chất là một hàm **trả về** `gin.HandlerFunc`, cho phép truyền tham số cấu hình vào:

```go
// Middleware logger đơn giản
func logger() gin.HandlerFunc {
    return func(c *gin.Context) {

        // ── PHẦN TRƯỚC (Pre-processing) ──────────────────────
        // Code ở đây chạy TRƯỚC khi request đến handler
        log.Printf("→ Request từ %s: %s %s",
            c.ClientIP(),
            c.Request.Method,
            c.Request.URL.Path,
        )

        // c.Next() chuyển tiếp request đến middleware/handler tiếp theo
        // Sau khi handler xử lý xong, luồng sẽ quay lại đây
        c.Next()

        // ── PHẦN SAU (Post-processing) ────────────────────────
        // Code ở đây chạy SAU khi handler đã xử lý xong
        log.Printf("← Response: status %d", c.Writer.Status())
    }
}
```

### `c.Next()` vs `c.Abort()`

|Hàm|Tác dụng|Dùng khi nào|
|---|---|---|
|`c.Next()`|Tiếp tục → chuyển request đến middleware/handler kế tiếp|Mọi thứ OK, cho đi tiếp|
|`c.Abort()`|Dừng → không gọi middleware/handler nào thêm nữa|Phát hiện lỗi, từ chối request|
|`c.AbortWithStatus(code)`|Dừng + set status code|Trả lỗi không có body|
|`c.AbortWithStatusJSON(code, obj)`|Dừng + trả JSON|Trả lỗi có body JSON|

```go
// Ví dụ middleware kiểm tra API key
func requireAPIKey() gin.HandlerFunc {
    return func(c *gin.Context) {
        apiKey := c.GetHeader("X-API-Key")

        if apiKey == "" || apiKey != os.Getenv("API_KEY") {
            // Abort: dừng lại, không cho vào handler
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "error": "Invalid or missing API key",
            })
            return // return sau Abort là best practice
        }

        // Next: hợp lệ, cho đi tiếp
        c.Next()
    }
}
```

---

## Các cách áp dụng Middleware trong Gin

### Cách 1: Áp dụng cho toàn bộ ứng dụng

```go
func main() {
    router := gin.New() // gin.New() không có middleware mặc định
    // hoặc gin.Default() đã có Logger + Recovery sẵn

    // Tất cả request đều đi qua logger
    router.Use(logger())
    router.Use(requireAPIKey())

    router.GET("/ping", pingHandler)
    router.GET("/users", getUsersHandler)
    router.Run(":8080")
}
```

### Cách 2: Áp dụng cho một group route

```go
func main() {
    router := gin.New()

    // Chỉ các route /api/v1/* mới đi qua logger và auth
    v1 := router.Group("/api/v1", logger(), requireAPIKey())
    {
        v1.GET("/users", getUsersHandler)
        v1.POST("/users", createUserHandler)
    }

    // Route public không qua middleware
    router.GET("/health", healthCheckHandler)

    router.Run(":8080")
}
```

### Cách 3: Áp dụng cho từng route cụ thể

```go
func main() {
    router := gin.New()

    // Chỉ route DELETE mới cần xác thực admin
    router.GET("/products",          getProductsHandler)              // không cần auth
    router.DELETE("/products/:id",   requireAdmin(), deleteHandler)   // cần auth admin

    router.Run(":8080")
}
```

### Cách 4: Phối hợp nhiều middleware theo thứ tự

```go
func main() {
    router := gin.New()

    // Middleware chạy theo thứ tự từ trái qua phải
    router.GET("/admin/reports",
        logger(),          // 1. Ghi log
        authenticate(),    // 2. Kiểm tra token
        authorizeAdmin(),  // 3. Kiểm tra quyền admin
        getReportsHandler, // 4. Xử lý business logic
    )

    router.Run(":8080")
}
```

> 💡 **Gin sử dụng `gin.Default()` hay `gin.New()`?**
> 
> - `gin.Default()` = `gin.New()` + tự thêm 2 middleware mặc định: **Logger** (in log request) và **Recovery** (bắt panic, tránh crash server).
> - `gin.New()` = engine rỗng, không có middleware nào — dùng khi muốn toàn quyền kiểm soát middleware stack.

---

## Thứ tự Middleware — Giải thích từng lớp

Khi xây dựng một ứng dụng thực tế, middleware phải được sắp xếp theo đúng thứ tự — lớp ngoài bảo vệ và chuẩn bị môi trường cho lớp trong. Dưới đây là thứ tự **đúng và hợp lý**, kèm giải thích chi tiết từng lớp:

![](https://zalopay-oss.github.io/go-advanced/images/ch5-03-middleware_flow.png)

```
Request đến
     ↓
1.  ExceptionHandler      ← bắt mọi lỗi không ngờ tới
     ↓
2.  HSTS                  ← ép dùng HTTPS
     ↓
3.  HttpsRedirection      ← chuyển HTTP → HTTPS
     ↓
4.  CORS                  ← kiểm tra nguồn gốc request
     ↓
5.  RateLimiter           ← giới hạn số request/giây
     ↓
6.  Routing               ← xác định handler nào xử lý
     ↓
7.  Logger                ← ghi log (đã biết route)
     ↓
8.  Authentication        ← xác thực "bạn là ai?"
     ↓
9.  Authorization         ← kiểm tra "bạn được phép làm gì?"
     ↓
10. Custom Middleware      ← logic đặc thù của ứng dụng
     ↓
11. Endpoint (Handler)    ← xử lý business logic
```

### Giải thích chi tiết từng lớp

**Lớp 1 — ExceptionHandler (Bắt lỗi toàn cục)** Lớp ngoài cùng, bắt mọi exception/panic chưa được xử lý từ các lớp bên trong. Phải đứng ngoài cùng để bao phủ toàn bộ pipeline. Trong Gin đây là middleware `Recovery`.

```go
// Gin tích hợp sẵn:
router.Use(gin.Recovery())
// Khi có panic bên trong → Recovery bắt lại, trả về 500 thay vì crash server
```

**Lớp 2 — HSTS (HTTP Strict Transport Security)** Thêm header `Strict-Transport-Security` vào response, báo cho browser biết: "Lần sau phải dùng HTTPS, không dùng HTTP nữa". Phải đứng sớm để header này luôn được gửi về, kể cả khi các lớp sau có lỗi.

```
Response header: Strict-Transport-Security: max-age=31536000; includeSubDomains
```

**Lớp 3 — HttpsRedirection (Chuyển hướng HTTPS)** Nếu request đến qua HTTP (không bảo mật) → tự động chuyển hướng sang HTTPS. Phải đứng trước CORS và Authentication vì không cần xử lý gì thêm — chuyển hướng ngay.

```
HTTP request → 301 Redirect → HTTPS request
```

**Lớp 4 — CORS (Cross-Origin Resource Sharing)** Kiểm tra xem domain của client có được phép gọi API không. Phải đứng **trước RateLimiter và Authentication** vì:

- Browser gửi **preflight request** (OPTIONS) trước khi gửi request thật
- Preflight cần được phản hồi nhanh, không nên qua rate limiter hay auth

```go
router.Use(cors.New(cors.Config{
    AllowOrigins: []string{"https://myapp.com"},
    AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
}))
```

**Lớp 5 — RateLimiter (Giới hạn tần suất)** Giới hạn số request mỗi IP/user trong một khoảng thời gian. Phải đứng **trước Authentication** để:

- Bảo vệ cả endpoint login khỏi brute-force attack
- Giảm tải trước khi thực hiện các thao tác tốn kém (query DB để verify token)

```
IP 1.2.3.4 gửi 1000 request/giây → RateLimiter chặn → trả 429 Too Many Requests
```

**Lớp 6 — Routing (Định tuyến)** Xác định request này khớp với route nào và handler nào sẽ xử lý. Gin thực hiện routing tự động khi bạn khai báo `router.GET(...)`. Phải biết route trước khi ghi log (Logger cần tên route) và authenticate (một số route không cần auth).

**Lớp 7 — Logger (Ghi nhật ký)** Ghi lại thông tin request. Đặt sau Routing vì:

- Cần biết route pattern để log có ý nghĩa (`/users/:id` thay vì `/users/123`)
- Đặt sau các lớp bảo mật đầu (CORS, RateLimiter) để không log các request bị chặn sớm làm nhiễu log

```
[INFO] 2024-01-15 10:30:00 | GET /api/v1/users/123 | 200 | 15ms | 192.168.1.1
```

**Lớp 8 — Authentication (Xác thực danh tính)** Kiểm tra "bạn là ai?" — xác thực token, session, API key. Đặt sau Logger để mọi request đều được ghi lại (kể cả request không hợp lệ). Đặt **trước** Authorization vì phải biết user là ai mới kiểm tra quyền được.

```go
func authenticate() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        user, err := validateToken(token)
        if err != nil {
            c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
            return
        }
        // Lưu thông tin user vào context để các lớp sau dùng
        c.Set("currentUser", user)
        c.Next()
    }
}
```

**Lớp 9 — Authorization (Kiểm tra quyền)** Kiểm tra "bạn được làm gì?" — user đã xác thực ở lớp 8, giờ kiểm tra họ có quyền truy cập resource này không. Phải sau Authentication vì cần thông tin user (đã lưu ở `c.Set`).

```go
func authorizeAdmin() gin.HandlerFunc {
    return func(c *gin.Context) {
        user := c.MustGet("currentUser").(User)
        if user.Role != "admin" {
            c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden"})
            return
        }
        c.Next()
    }
}
```

**Lớp 10 — Custom Middleware (Logic đặc thù)** Các middleware riêng của ứng dụng: cache, request ID, đa ngôn ngữ, transform data... Đặt sau tất cả middleware hạ tầng vì chúng phụ thuộc vào thông tin đã được xử lý ở các lớp trước (user, route...).

**Lớp 11 — Endpoint (Handler)** Xử lý business logic thực sự: query database, tính toán, trả kết quả. Đây là đích đến cuối cùng sau khi đã qua tất cả các lớp bảo vệ.

---

### Tại sao thứ tự này là đúng?

```
ExceptionHandler phải ngoài cùng → bắt lỗi của TẤT CẢ lớp trong
HSTS + HTTPS trước CORS → bảo mật kết nối trước khi xử lý nguồn gốc
CORS trước RateLimiter → preflight OPTIONS không nên bị rate limit
RateLimiter trước Auth → bảo vệ endpoint login, tiết kiệm tài nguyên
Logger trước Auth → ghi lại cả request sẽ bị từ chối
Authentication trước Authorization → biết "ai" trước rồi mới biết "được làm gì"
Custom Middleware sau tất cả → có đủ thông tin từ các lớp trước
```

> ⚠️ **Nếu đảo Authentication và Authorization:** Authorization không biết user là ai vì Authentication chưa chạy → không thể kiểm tra quyền → lỗi.

> ⚠️ **Nếu đặt RateLimiter sau Authentication:** Attacker có thể gửi hàng nghìn request với token giả → server phải query DB để verify token từng lần → tốn tài nguyên → dễ bị DDoS.

---

## Case Study: Phối hợp nhiều Middleware

### Bài toán

Xây dựng hệ thống có 3 loại route với yêu cầu bảo mật khác nhau:

```
Public API   → Không cần auth
User API     → Cần đăng nhập
Admin API    → Cần đăng nhập VÀ là admin
```

### Giải pháp

```go
func main() {
    router := gin.New()
    router.Use(gin.Recovery()) // ExceptionHandler cho tất cả

    // Middleware dùng chung
    router.Use(logger())

    // ── PUBLIC: không cần auth ──────────────────────────────────
    public := router.Group("/api/v1")
    {
        public.GET("/products", getProductsHandler)
        public.POST("/login",   loginHandler)
    }

    // ── USER: cần đăng nhập ─────────────────────────────────────
    // authenticate() kiểm tra token → lưu user vào context
    user := router.Group("/api/v1", authenticate())
    {
        user.GET("/profile",       getProfileHandler)
        user.PUT("/profile",       updateProfileHandler)
        user.GET("/orders",        getOrdersHandler)
    }

    // ── ADMIN: cần đăng nhập VÀ là admin ────────────────────────
    // authenticate() chạy trước → authorizeAdmin() chạy sau
    admin := router.Group("/admin", authenticate(), authorizeAdmin())
    {
        admin.GET("/users",        getAllUsersHandler)
        admin.DELETE("/users/:id", deleteUserHandler)
        admin.GET("/reports",      getReportsHandler)
    }

    router.Run(":8080")
}
```

### Ưu điểm của cách tiếp cận này

- `authenticate()` viết một lần, dùng ở mọi nơi cần auth
- `authorizeAdmin()` viết một lần, dùng cho tất cả route admin
- Muốn thêm role mới (`authorizeManager()`) → chỉ viết thêm một middleware
- Không cần copy-paste logic vào từng handler

---

## Cách viết Middleware thanh lịch hơn

Khi số lượng middleware nhiều, việc quản lý thứ tự và thêm/bớt có thể trở nên khó khăn. Phần này giới thiệu cách xây dựng **middleware chain** (chuỗi middleware) linh hoạt hơn.

### Vấn đề với cách thông thường

```go
// Khó đọc, khó thêm/bớt
router.GET("/ping",
    middleware1(), middleware2(), middleware3(), middleware4(), handler)
```

### Giải pháp: Router với middleware chain

```go
// Cách sử dụng rõ ràng, dễ quản lý
r := NewRouter()
r.Use(logger)      // Thêm middleware
r.Use(timeout)     // Thứ tự rõ ràng
r.Use(ratelimit)   // Dễ thêm/xóa
r.Add("/", helloHandler)
```

### Hiện thực từ đầu

```go
// middleware là kiểu hàm nhận Handler và trả về Handler mới (đã wrap)
// Giống như "bọc" handler trong một lớp xử lý mới
type middleware func(http.Handler) http.Handler

// Router quản lý chuỗi middleware và ánh xạ route → handler
type Router struct {
    middlewareChain []middleware         // Danh sách middleware theo thứ tự
    mux             map[string]http.Handler // Route → Handler đã được wrap
}

func NewRouter() *Router {
    return &Router{
        mux: make(map[string]http.Handler),
    }
}

// Use thêm một middleware vào cuối chuỗi
func (r *Router) Use(m middleware) {
    r.middlewareChain = append(r.middlewareChain, m)
}

// Add đăng ký một route với handler,
// tự động wrap handler trong toàn bộ middleware chain
func (r *Router) Add(route string, h http.Handler) {
    var mergedHandler = h

    // Duyệt NGƯỢC từ cuối đến đầu để apply middleware
    // Lý do: middleware[0] phải là lớp ngoài cùng (chạy đầu tiên)
    //
    // Ví dụ: chain = [logger, auth, ratelimit], handler = h
    // i=2: mergedHandler = ratelimit(h)
    // i=1: mergedHandler = auth(ratelimit(h))
    // i=0: mergedHandler = logger(auth(ratelimit(h)))
    //
    // Khi request đến: logger chạy trước → auth → ratelimit → h
    // Đúng thứ tự như ta Use() vào
    for i := len(r.middlewareChain) - 1; i >= 0; i-- {
        mergedHandler = r.middlewareChain[i](mergedHandler)
    }

    r.mux[route] = mergedHandler
}

// ServeHTTP triển khai interface http.Handler — cho phép Router làm HTTP server
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    if handler, ok := r.mux[req.URL.Path]; ok {
        handler.ServeHTTP(w, req)
    } else {
        http.NotFound(w, req)
    }
}
```

### Minh họa tại sao duyệt ngược

```
Middleware chain: [logger, auth, ratelimit]
Handler:          h

Duyệt ngược:
  i=2: mergedHandler = ratelimit(h)
  i=1: mergedHandler = auth( ratelimit(h) )
  i=0: mergedHandler = logger( auth( ratelimit(h) ) )

Kết quả cuối: logger → auth → ratelimit → h
              ✅ Đúng thứ tự ta muốn (logger chạy đầu tiên)

Nếu duyệt xuôi:
  i=0: mergedHandler = logger(h)
  i=1: mergedHandler = auth( logger(h) )
  i=2: mergedHandler = ratelimit( auth( logger(h) ) )

Kết quả: ratelimit → auth → logger → h
         ❌ Ngược thứ tự (ratelimit chạy đầu tiên, không phải logger)
```

### Ví dụ thực tế với `net/http`

```go
// Định nghĩa các middleware
func loggerMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("→ %s %s", r.Method, r.URL.Path)
        next.ServeHTTP(w, r) // Gọi handler/middleware tiếp theo
        log.Printf("← Done")
    })
}

func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return // Không gọi next.ServeHTTP → dừng chain
        }
        next.ServeHTTP(w, r)
    })
}

// Handler cuối cùng
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, World!")
}

// Sử dụng
func main() {
    r := NewRouter()
    r.Use(loggerMiddleware)
    r.Use(authMiddleware)
    r.Add("/hello", http.HandlerFunc(helloHandler))

    http.ListenAndServe(":8080", r)
}
```

---

## Tổng kết

### Những điểm quan trọng cần nhớ

|Khái niệm|Tóm tắt|
|---|---|
|**Middleware là gì**|Lớp xử lý trung gian giữa Request và Response|
|**`c.Next()`**|Chuyển tiếp request đến middleware/handler kế tiếp|
|**`c.Abort()`**|Dừng chain, không cho request đi tiếp|
|**Phạm vi áp dụng**|Toàn app → Group → Route cụ thể|
|**Thứ tự middleware**|ExceptionHandler → HTTPS → CORS → RateLimiter → Routing → Logger → Auth → AuthZ → Custom → Endpoint|
|**Duyệt ngược khi build chain**|Để middleware đầu tiên `Use()` là lớp ngoài cùng|

### Nguyên tắc thiết kế Middleware

1. **Mỗi middleware làm một việc duy nhất** — Logger chỉ log, Auth chỉ xác thực.
2. **Luôn gọi `c.Next()` hoặc `c.Abort()`** — không để request bị "treo" giữa chừng.
3. **Dùng `c.Set()` / `c.Get()`** để truyền dữ liệu giữa các middleware (ví dụ: lưu user sau khi authenticate).
4. **Đặt middleware theo đúng thứ tự** — sai thứ tự có thể gây lỗ hổng bảo mật.
5. **Middleware nên stateless** — không lưu trạng thái bên trong middleware, dùng context để truyền dữ liệu.