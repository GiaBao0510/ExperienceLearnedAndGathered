# 📘 HƯỚNG DẪN TÍCH HỢP DATABASE VÀO DỰ ÁN GO-REALTIME

## 📑 MỤC LỤC

1. [Phân tích kiến trúc hiện tại](#1-phân-tích-kiến-trúc-hiện-tại)
2. [Lỗi SSE hiện tại và cách khắc phục](#2-lỗi-sse-hiện-tại-và-cách-khắc-phục)
3. [Kiến trúc đề xuất sau khi tích hợp DB](#3-kiến-trúc-đề-xuất-sau-khi-tích-hợp-db)
4. [Cấu trúc thư mục mở rộng](#4-cấu-trúc-thư-mục-mở-rộng)
5. [Thiết lập kết nối Database](#5-thiết-lập-kết-nối-database)
6. [Định nghĩa Model (Domain Entity)](#6-định-nghĩa-model-domain-entity)
7. [Tầng Repository — Truy vấn DB](#7-tầng-repository--truy-vấn-db)
8. [Tầng Service — Business Logic](#8-tầng-service--business-logic)
9. [Tích hợp vào Handler](#9-tích-hợp-vào-handler)
10. [Quy trình lưu dữ liệu theo từng chức năng](#10-quy-trình-lưu-dữ-liệu-theo-từng-chức-năng)
11. [Biến môi trường và bảo mật](#11-biến-môi-trường-và-bảo-mật)
12. [Testing](#12-testing)
13. [Checklist triển khai](#13-checklist-triển-khai)

---

## 1. PHÂN TÍCH KIẾN TRÚC HIỆN TẠI

### 1.1 Cấu trúc thư mục hiện tại

```
go_RealTime/
├── cmd/server/
│   └── main.go              ← Entry point: khởi tạo Hub, đăng ký route, chạy HTTP server
├── internal/
│   ├── handler/
│   │   ├── sse_handler.go   ← SSE: gửi event 1 chiều Server → Client
│   │   └── ws_handler.go    ← WebSocket: giao tiếp 2 chiều Server ↔ Client
│   ├── hub/
│   │   └── hub.go           ← Hub: quản lý tất cả WebSocket client (register/unregister/broadcast)
│   └── model/
│       └── message.go       ← Model: Message, LiveStats, MessageType
├── static/
│   ├── sse_demo.html        ← Frontend SSE demo (Live Dashboard)
│   └── ws_demo.html         ← Frontend WebSocket demo (Mini Chat)
├── diagram/                 ← CDM diagram
├── vendor/                  ← Vendor dependencies
├── go.mod
├── go.sum
└── makefile
```

### 1.2 Luồng dữ liệu hiện tại

```
┌────────────────────────────────────────────────────────────────┐
│                         HIỆN TẠI                               │
│                                                                │
│   Client ──HTTP──▶ main.go ──route──▶ SSEHandler               │
│                                    ──route──▶ WSHandler ──▶ Hub │
│                                                                │
│   ⚠️  KHÔNG CÓ DATABASE → Dữ liệu mất khi server restart     │
│   ⚠️  SSE dùng dữ liệu giả (random)                           │
│   ⚠️  WS message chỉ broadcast, không lưu lại                 │
└────────────────────────────────────────────────────────────────┘
```

### 1.3 Những gì cần thêm

| Thành phần          | Hiện tại        | Cần thêm                                   |
| ------------------- | --------------- | ------------------------------------------ |
| Kết nối DB          | ❌ Không có      | ✅ `database/sql` + `pgx` driver            |
| Config              | ❌ Hardcode      | ✅ `.env` + biến môi trường                 |
| Repository layer    | ❌ Không có      | ✅ Interface + implement cho từng bảng      |
| Service layer       | ❌ Không có      | ✅ Business logic giữa handler và repo      |
| User management     | ❌ Không có      | ✅ CRUD users, xác thực                     |
| Message persistence | ❌ Chỉ broadcast | ✅ Lưu message vào DB trước khi broadcast   |
| Notification        | ❌ Không có      | ✅ Tạo notification + lưu user_notification |

---

## 2. LỖI SSE HIỆN TẠI VÀ CÁCH KHẮC PHỤC

### 2.1 Nguyên nhân gốc rễ

**Vấn đề**: Server gửi event nhưng Frontend không nhận được.

**Phân tích chi tiết**:

Trong file `internal/handler/sse_handler.go` (dòng 88):

```go
fmt.Fprintf(w, "event: status-update\n")  // ← Server gửi event tên "status-update"
```

Trong file `static/sse_demo.html` (dòng 174):

```javascript
eventSource.addEventListener('stats-update', (e) => {  // ← FE lắng nghe "stats-update"
```

**Hai tên event không khớp nhau**:

- Server gửi: `status-update` (stat**us**)
- Frontend nghe: `stats-update` (stat**s**)

> 💡 **Bài học**: Khi dùng SSE, tên event phải **chính xác tuyệt đối** giữa server và client. Chỉ cần sai 1 ký tự, `EventSource.addEventListener()` sẽ không bao giờ trigger callback.

### 2.2 Cách khắc phục

Bạn có **2 lựa chọn** (chọn 1 trong 2, **KHÔNG LÀM CẢ HAI**):

#### Lựa chọn A: Sửa phía Backend (Khuyến nghị ✅)

Mở file `internal/handler/sse_handler.go`, tìm dòng 88:

```go
// TRƯỚC (sai):
fmt.Fprintf(w, "event: status-update\n")

// SAU (đúng):
fmt.Fprintf(w, "event: stats-update\n")
```

**Lý do khuyến nghị**: Frontend đã đặt tên event hợp lý (`stats` = statistics/thống kê), phù hợp với ngữ cảnh Live Dashboard hiển thị số liệu thống kê.

#### Lựa chọn B: Sửa phía Frontend

Mở file `static/sse_demo.html`, tìm dòng 174:

```javascript
// TRƯỚC (không khớp):
eventSource.addEventListener('stats-update', (e) => {

// SAU (khớp với server):
eventSource.addEventListener('status-update', (e) => {
```

### 2.3 Khuyến nghị thêm: Định nghĩa constant cho event name

Để tránh lỗi typo trong tương lai, nên tạo một nơi tập trung quản lý tên event:

```go
// internal/model/event.go (file mới)
package model

// SSE Event names — dùng constant để tránh typo
const (
    SSEEventConnected   = "connected"
    SSEEventStatsUpdate = "stats-update"
)
```

Sau đó dùng trong handler:

```go
fmt.Fprintf(w, "event: %s\n", model.SSEEventStatsUpdate)
```

---

## 3. KIẾN TRÚC ĐỀ XUẤT SAU KHI TÍCH HỢP DB

### 3.1 Mô hình phân tầng (Layered Architecture)

```
┌─────────────────────────────────────────────────────────┐
│                    CLIENT (Browser)                      │
│        sse_demo.html  │  ws_demo.html  │  REST API       │
└─────────┬─────────────┴────────┬───────┴────────┬───────┘
          │ SSE                  │ WebSocket       │ HTTP
          ▼                      ▼                 ▼
┌─────────────────────────────────────────────────────────┐
│              HANDLER LAYER (internal/handler/)           │
│  Nhiệm vụ: Nhận request, validate input, trả response   │
│  SSEHandler  │  WSHandler  │  UserHandler  │ ConvHandler  │
└──────────────┴──────┬──────┴───────────────┴────────────┘
                      │ gọi service
                      ▼
┌─────────────────────────────────────────────────────────┐
│              SERVICE LAYER (internal/service/)            │
│  Nhiệm vụ: Business logic, orchestration                 │
│  MessageService │ UserService │ ConversationService       │
│                 │ NotificationService                     │
└─────────────────┴──────┬──────┴─────────────────────────┘
                         │ gọi repository
                         ▼
┌─────────────────────────────────────────────────────────┐
│           REPOSITORY LAYER (internal/repository/)        │
│  Nhiệm vụ: CRUD database, chỉ biết SQL                  │
│  MessageRepo │ UserRepo │ ConversationRepo               │
│              │ NotificationRepo                          │
└──────────────┴──────┬───┴───────────────────────────────┘
                      │ SQL query
                      ▼
┌─────────────────────────────────────────────────────────┐
│            DATABASE (PostgreSQL — Supabase)               │
│  users │ conversation │ conversation_member │ message     │
│        │ notification │ user_notification                 │
└─────────────────────────────────────────────────────────┘
```

### 3.2 Tại sao chia 3 tầng?

| Tầng           | Trách nhiệm                           | Lợi ích                                        |
| -------------- | ------------------------------------- | ---------------------------------------------- |
| **Handler**    | Parse request, validate, trả response | Dễ đổi protocol (HTTP→gRPC) mà không sửa logic |
| **Service**    | Business logic, gọi nhiều repo        | Nơi duy nhất chứa logic, dễ test               |
| **Repository** | Chỉ biết SQL/DB                       | Dễ đổi DB (Postgres→MySQL) mà không sửa logic  |

> 🎓 **Nguyên tắc quan trọng**: Mỗi tầng chỉ phụ thuộc vào tầng ngay bên dưới nó (Handler → Service → Repository). **Không bao giờ** để Handler gọi trực tiếp Repository.

---

## 4. CẤU TRÚC THƯ MỤC MỞ RỘNG

```
go_RealTime/
├── cmd/
│   └── server/
│       └── main.go                  ← Entry point (cập nhật: init DB, inject dependencies)
│
├── internal/
│   ├── config/
│   │   └── config.go                ← [MỚI] Đọc biến môi trường, cấu hình DB
│   │
│   ├── database/
│   │   └── postgres.go              ← [MỚI] Kết nối PostgreSQL, connection pool
│   │
│   ├── model/
│   │   ├── message.go               ← (giữ nguyên + bổ sung fields DB)
│   │   ├── user.go                  ← [MỚI] User entity
│   │   ├── conversation.go          ← [MỚI] Conversation + ConversationMember entity
│   │   ├── notification.go          ← [MỚI] Notification + UserNotification entity
│   │   └── event.go                 ← [MỚI] SSE event name constants
│   │
│   ├── repository/
│   │   ├── interfaces.go            ← [MỚI] Interface định nghĩa cho tất cả repo
│   │   ├── user_repo.go             ← [MỚI] CRUD users
│   │   ├── message_repo.go          ← [MỚI] CRUD message
│   │   ├── conversation_repo.go     ← [MỚI] CRUD conversation + member
│   │   └── notification_repo.go     ← [MỚI] CRUD notification
│   │
│   ├── service/
│   │   ├── user_service.go          ← [MỚI] Business logic user
│   │   ├── message_service.go       ← [MỚI] Business logic message
│   │   ├── conversation_service.go  ← [MỚI] Business logic conversation
│   │   └── notification_service.go  ← [MỚI] Business logic notification
│   │
│   ├── handler/
│   │   ├── sse_handler.go           ← (cập nhật: inject service, lấy data thật từ DB)
│   │   ├── ws_handler.go            ← (cập nhật: lưu message qua service)
│   │   ├── user_handler.go          ← [MỚI] REST API cho user
│   │   └── conversation_handler.go  ← [MỚI] REST API cho conversation
│   │
│   └── hub/
│       └── hub.go                   ← (giữ nguyên cơ bản)
│
├── static/
│   ├── sse_demo.html                ← (giữ nguyên)
│   └── ws_demo.html                 ← (giữ nguyên)
│
├── .env                             ← [MỚI] Biến môi trường (KHÔNG commit lên git)
├── .env.example                     ← [MỚI] Mẫu biến môi trường (commit lên git)
├── .gitignore                       ← [MỚI] Ignore .env, vendor, binary
├── go.mod
├── go.sum
└── makefile
```

---

## 5. THIẾT LẬP KẾT NỐI DATABASE

### 5.1 File `.env` (KHÔNG COMMIT LÊN GIT)

```env
# Database
DB_HOST=db.auuceyrifwxwmsbchxge.supabase.co
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=Ch@tRealTime123
DB_NAME=postgres
DB_SSLMODE=require

# Hoặc dùng connection string đầy đủ:
DATABASE_URL=postgresql://postgres:Ch@tRealTime123@db.auuceyrifwxwmsbchxge.supabase.co:5432/postgres?sslmode=require

# Server
SERVER_PORT=8080
```

### 5.2 File `.env.example` (COMMIT LÊN GIT)

```env
# Database
DB_HOST=your-supabase-host.supabase.co
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your-password-here
DB_NAME=postgres
DB_SSLMODE=require

# Hoặc connection string:
DATABASE_URL=postgresql://postgres:YOUR_PASSWORD@your-host.supabase.co:5432/postgres?sslmode=require

# Server
SERVER_PORT=8080
```

### 5.3 File `internal/config/config.go`

```go
package config

import (
    "fmt"
    "os"
    "strconv"
)

// Config chứa toàn bộ cấu hình ứng dụng
type Config struct {
    DB     DBConfig
    Server ServerConfig
}

// DBConfig chứa cấu hình kết nối database
type DBConfig struct {
    Host     string
    Port     int
    User     string
    Password string
    DBName   string
    SSLMode  string
}

// ServerConfig chứa cấu hình server
type ServerConfig struct {
    Port string
}

// ConnectionString trả về chuỗi kết nối PostgreSQL
func (db DBConfig) ConnectionString() string {
    return fmt.Sprintf(
        "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
        db.Host, db.Port, db.User, db.Password, db.DBName, db.SSLMode,
    )
}

// Load đọc cấu hình từ biến môi trường
func Load() (*Config, error) {
    port, err := strconv.Atoi(getEnv("DB_PORT", "5432"))
    if err != nil {
        return nil, fmt.Errorf("DB_PORT không hợp lệ: %w", err)
    }

    return &Config{
        DB: DBConfig{
            Host:     getEnv("DB_HOST", "localhost"),
            Port:     port,
            User:     getEnv("DB_USER", "postgres"),
            Password: getEnv("DB_PASSWORD", ""),
            DBName:   getEnv("DB_NAME", "postgres"),
            SSLMode:  getEnv("DB_SSLMODE", "require"),
        },
        Server: ServerConfig{
            Port: ":" + getEnv("SERVER_PORT", "8080"),
        },
    }, nil
}

// getEnv lấy biến môi trường, trả về fallback nếu không tồn tại
func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}
```

### 5.4 File `internal/database/postgres.go`

```go
package database

import (
    "context"
    "database/sql"
    "fmt"
    "log"
    "time"

    _ "github.com/jackc/pgx/v5/stdlib" // PostgreSQL driver

    "github.com/GiaBao0510/Go-Realtime/internal/config"
)

// NewPostgresDB tạo kết nối đến PostgreSQL và trả về *sql.DB
//
// Tại sao dùng database/sql thay vì pgx trực tiếp?
//   - database/sql là standard library → dễ mock, dễ test
//   - Connection pool được quản lý tự động
//   - Dễ đổi driver (pgx → lib/pq) mà không sửa code
func NewPostgresDB(cfg config.DBConfig) (*sql.DB, error) {

    // Mở kết nối (chưa thực sự connect, chỉ khởi tạo pool)
    db, err := sql.Open("pgx", cfg.ConnectionString())
    if err != nil {
        return nil, fmt.Errorf("không thể mở kết nối DB: %w", err)
    }

    // Cấu hình connection pool
    db.SetMaxOpenConns(25)              // Tối đa 25 connection đồng thời
    db.SetMaxIdleConns(10)              // Tối đa 10 connection nhàn rỗi
    db.SetConnMaxLifetime(5 * time.Minute) // Mỗi connection sống tối đa 5 phút

    // Ping để kiểm tra kết nối thực sự hoạt động
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    if err := db.PingContext(ctx); err != nil {
        return nil, fmt.Errorf("không thể ping DB: %w", err)
    }

    log.Println("✅ Kết nối PostgreSQL thành công!")
    return db, nil
}
```

### 5.5 Thêm dependency vào `go.mod`

Chạy lệnh:

```bash
go get github.com/jackc/pgx/v5
go get github.com/joho/godotenv   # Để đọc file .env
```

---

## 6. ĐỊNH NGHĨA MODEL (DOMAIN ENTITY)

### 6.1 Cập nhật `internal/model/message.go`

```go
package model

import "time"

type MessageType string

const (
    TypeChat   MessageType = "chat"
    TypeSystem MessageType = "system"
    TypeStatus MessageType = "status"
)

// Message — dùng cho WebSocket broadcast (giữ nguyên struct cũ)
type Message struct {
    Type      MessageType `json:"type"`
    Content   string      `json:"content"`
    Sender    string      `json:"sender"`
    Timestamp time.Time   `json:"timestamp"`
}

// DBMessage — entity tương ứng với bảng `message` trong DB
// Tách riêng với Message vì DB có thêm các field (message_id, conversation_id, is_edited...)
type DBMessage struct {
    MessageID      int64       `json:"message_id"`
    ConversationID int         `json:"conversation_id"`
    SenderUID      *string     `json:"sender"`       // UUID, nullable (ON DELETE SET NULL)
    Content        string      `json:"content"`
    MessageType    string      `json:"message_type"` // 'text', 'image', 'file', 'system'
    IsEdited       bool        `json:"is_edited"`
    CreatedAt      time.Time   `json:"created_at"`
}

// LiveStats — dữ liệu cho SSE demo (giữ nguyên)
type LiveStats struct {
    OnlineUsers int    `json:"online_users"`
    ServerTime  string `json:"server_time"`
    CPULoad     int    `json:"cpu_load"`
    MemoryUsage int    `json:"memory_usage"`
}
```

### 6.2 File `internal/model/user.go` [MỚI]

```go
package model

import "time"

// User — entity tương ứng với bảng `users`
type User struct {
    UID          string     `json:"uid"`
    Name         string     `json:"name"`
    Email        string     `json:"email"`
    PasswordHash string     `json:"-"`             // Không bao giờ trả về client (json:"-")
    AvatarURL    *string    `json:"avatar_url"`    // Nullable
    IsOnline     bool       `json:"is_online"`
    LastSeenAt   *time.Time `json:"last_seen_at"`  // Nullable
    CreatedAt    time.Time  `json:"created_at"`
}

// CreateUserRequest — dữ liệu cần thiết khi tạo user mới
type CreateUserRequest struct {
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

// UpdateUserRequest — dữ liệu có thể cập nhật
type UpdateUserRequest struct {
    Name      *string `json:"name"`
    AvatarURL *string `json:"avatar_url"`
}
```

### 6.3 File `internal/model/conversation.go` [MỚI]

```go
package model

import "time"

// Conversation — entity tương ứng với bảng `conversation`
type Conversation struct {
    ConversationID int        `json:"conversation_id"`
    Name           *string    `json:"name"`        // NULL nếu private
    Type           string     `json:"type"`         // 'private' hoặc 'group'
    AvatarURL      *string    `json:"avatar_url"`
    CreatedBy      *string    `json:"created_by"`   // UUID, nullable
    CreatedAt      time.Time  `json:"created_at"`
}

// ConversationMember — entity tương ứng với bảng `conversation_member`
type ConversationMember struct {
    ConversationID int       `json:"conversation_id"`
    UID            string    `json:"uid"`
    Role           string    `json:"role"`     // 'admin' hoặc 'member'
    JoinedAt       time.Time `json:"joined_at"`
}
```

### 6.4 File `internal/model/notification.go` [MỚI]

```go
package model

import "time"

// Notification — entity tương ứng với bảng `notification`
type Notification struct {
    NotificationID int64     `json:"notification_id"`
    Type           string    `json:"type"`      // 'new_message', 'group_invite', 'system'
    Content        string    `json:"content"`
    CreatedAt      time.Time `json:"created_at"`
}

// UserNotification — entity tương ứng với bảng `user_notification`
type UserNotification struct {
    UID            string     `json:"uid"`
    NotificationID int64      `json:"notification_id"`
    IsRead         bool       `json:"is_read"`
    ReadAt         *time.Time `json:"read_at"` // Nullable
}
```

### 6.5 File `internal/model/event.go` [MỚI]

```go
package model

// SSE Event names — dùng constant để tránh lỗi typo giữa server và client
const (
    SSEEventConnected   = "connected"
    SSEEventStatsUpdate = "stats-update" // ⚠️ Phải khớp với addEventListener ở FE
)

// Notification types — dùng constant cho trường `type` trong bảng `notification`
const (
    NotifTypeNewMessage  = "new_message"
    NotifTypeGroupInvite = "group_invite"
    NotifTypeSystem      = "system"
)

// Message types — dùng constant cho trường `message_type` trong bảng `message`
const (
    MsgTypeText   = "text"
    MsgTypeImage  = "image"
    MsgTypeFile   = "file"
    MsgTypeSystem = "system"
)
```

---

## 7. TẦNG REPOSITORY — TRUY VẤN DB

### 7.1 File `internal/repository/interfaces.go`

```go
package repository

import (
    "context"

    "github.com/GiaBao0510/Go-Realtime/internal/model"
)

// ══════════════════════════════════════════════
// TẠI SAO DÙNG INTERFACE?
// ══════════════════════════════════════════════
// 1. Dễ mock trong unit test (thay thế DB thật bằng fake)
// 2. Dễ đổi implementation (Postgres → MySQL → MongoDB)
// 3. Tuân thủ Dependency Inversion Principle (SOLID)
// ══════════════════════════════════════════════

// UserRepository định nghĩa các thao tác với bảng `users`
type UserRepository interface {
    Create(ctx context.Context, user *model.User) error
    GetByUID(ctx context.Context, uid string) (*model.User, error)
    GetByEmail(ctx context.Context, email string) (*model.User, error)
    UpdateOnlineStatus(ctx context.Context, uid string, isOnline bool) error
    CountOnline(ctx context.Context) (int, error)
}

// MessageRepository định nghĩa các thao tác với bảng `message`
type MessageRepository interface {
    Create(ctx context.Context, msg *model.DBMessage) error
    GetByConversation(ctx context.Context, conversationID int, limit, offset int) ([]model.DBMessage, error)
    GetLatestByConversation(ctx context.Context, conversationID int) (*model.DBMessage, error)
}

// ConversationRepository định nghĩa các thao tác với bảng `conversation`
type ConversationRepository interface {
    Create(ctx context.Context, conv *model.Conversation) (int, error)
    GetByID(ctx context.Context, id int) (*model.Conversation, error)
    GetByUserUID(ctx context.Context, uid string) ([]model.Conversation, error)
    AddMember(ctx context.Context, member *model.ConversationMember) error
    RemoveMember(ctx context.Context, conversationID int, uid string) error
    GetMembers(ctx context.Context, conversationID int) ([]model.ConversationMember, error)
}

// NotificationRepository định nghĩa các thao tác với bảng `notification`
type NotificationRepository interface {
    Create(ctx context.Context, notif *model.Notification) (int64, error)
    CreateUserNotification(ctx context.Context, uid string, notificationID int64) error
    GetUnreadByUser(ctx context.Context, uid string) ([]model.Notification, error)
    MarkAsRead(ctx context.Context, uid string, notificationID int64) error
    CountUnread(ctx context.Context, uid string) (int, error)
}
```

### 7.2 Ví dụ implement: `internal/repository/message_repo.go`

```go
package repository

import (
    "context"
    "database/sql"
    "fmt"

    "github.com/GiaBao0510/Go-Realtime/internal/model"
)

// messageRepo implement MessageRepository bằng PostgreSQL
type messageRepo struct {
    db *sql.DB
}

// NewMessageRepo tạo instance mới
func NewMessageRepo(db *sql.DB) MessageRepository {
    return &messageRepo{db: db}
}

// Create lưu message mới vào bảng `message`
func (r *messageRepo) Create(ctx context.Context, msg *model.DBMessage) error {
    query := `
        INSERT INTO message (conversation_id, sender, content, message_type)
        VALUES ($1, $2, $3, $4)
        RETURNING message_id, created_at
    `
    return r.db.QueryRowContext(ctx, query,
        msg.ConversationID,
        msg.SenderUID,
        msg.Content,
        msg.MessageType,
    ).Scan(&msg.MessageID, &msg.CreatedAt)
}

// GetByConversation lấy danh sách message theo conversation, có phân trang
func (r *messageRepo) GetByConversation(
    ctx context.Context,
    conversationID int,
    limit, offset int,
) ([]model.DBMessage, error) {
    query := `
        SELECT message_id, conversation_id, sender, content,
               message_type, is_edited, created_at
        FROM message
        WHERE conversation_id = $1
        ORDER BY created_at DESC
        LIMIT $2 OFFSET $3
    `
    rows, err := r.db.QueryContext(ctx, query, conversationID, limit, offset)
    if err != nil {
        return nil, fmt.Errorf("query messages: %w", err)
    }
    defer rows.Close()

    var messages []model.DBMessage
    for rows.Next() {
        var msg model.DBMessage
        if err := rows.Scan(
            &msg.MessageID, &msg.ConversationID, &msg.SenderUID,
            &msg.Content, &msg.MessageType, &msg.IsEdited, &msg.CreatedAt,
        ); err != nil {
            return nil, fmt.Errorf("scan message: %w", err)
        }
        messages = append(messages, msg)
    }
    return messages, rows.Err()
}

// GetLatestByConversation lấy tin nhắn mới nhất của conversation
func (r *messageRepo) GetLatestByConversation(
    ctx context.Context,
    conversationID int,
) (*model.DBMessage, error) {
    query := `
        SELECT message_id, conversation_id, sender, content,
               message_type, is_edited, created_at
        FROM message
        WHERE conversation_id = $1
        ORDER BY created_at DESC
        LIMIT 1
    `
    var msg model.DBMessage
    err := r.db.QueryRowContext(ctx, query, conversationID).Scan(
        &msg.MessageID, &msg.ConversationID, &msg.SenderUID,
        &msg.Content, &msg.MessageType, &msg.IsEdited, &msg.CreatedAt,
    )
    if err == sql.ErrNoRows {
        return nil, nil // Không có message nào
    }
    if err != nil {
        return nil, fmt.Errorf("query latest message: %w", err)
    }
    return &msg, nil
}
```

### 7.3 Ví dụ implement: `internal/repository/user_repo.go`

```go
package repository

import (
    "context"
    "database/sql"
    "fmt"
    "time"

    "github.com/GiaBao0510/Go-Realtime/internal/model"
)

type userRepo struct {
    db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepository {
    return &userRepo{db: db}
}

func (r *userRepo) Create(ctx context.Context, user *model.User) error {
    query := `
        INSERT INTO users (name, email, password_hash, avatar_url)
        VALUES ($1, $2, $3, $4)
        RETURNING uid, created_at
    `
    return r.db.QueryRowContext(ctx, query,
        user.Name, user.Email, user.PasswordHash, user.AvatarURL,
    ).Scan(&user.UID, &user.CreatedAt)
}

func (r *userRepo) GetByUID(ctx context.Context, uid string) (*model.User, error) {
    query := `
        SELECT uid, name, email, password_hash, avatar_url,
               is_online, last_seen_at, created_at
        FROM users WHERE uid = $1
    `
    var user model.User
    err := r.db.QueryRowContext(ctx, query, uid).Scan(
        &user.UID, &user.Name, &user.Email, &user.PasswordHash,
        &user.AvatarURL, &user.IsOnline, &user.LastSeenAt, &user.CreatedAt,
    )
    if err == sql.ErrNoRows {
        return nil, nil
    }
    if err != nil {
        return nil, fmt.Errorf("query user by uid: %w", err)
    }
    return &user, nil
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (*model.User, error) {
    query := `
        SELECT uid, name, email, password_hash, avatar_url,
               is_online, last_seen_at, created_at
        FROM users WHERE email = $1
    `
    var user model.User
    err := r.db.QueryRowContext(ctx, query, email).Scan(
        &user.UID, &user.Name, &user.Email, &user.PasswordHash,
        &user.AvatarURL, &user.IsOnline, &user.LastSeenAt, &user.CreatedAt,
    )
    if err == sql.ErrNoRows {
        return nil, nil
    }
    if err != nil {
        return nil, fmt.Errorf("query user by email: %w", err)
    }
    return &user, nil
}

func (r *userRepo) UpdateOnlineStatus(ctx context.Context, uid string, isOnline bool) error {
    query := `
        UPDATE users
        SET is_online = $1, last_seen_at = $2
        WHERE uid = $3
    `
    now := time.Now()
    _, err := r.db.ExecContext(ctx, query, isOnline, now, uid)
    return err
}

func (r *userRepo) CountOnline(ctx context.Context) (int, error) {
    var count int
    err := r.db.QueryRowContext(ctx,
        "SELECT COUNT(*) FROM users WHERE is_online = TRUE",
    ).Scan(&count)
    return count, err
}
```

### 7.4 Ví dụ implement: `internal/repository/notification_repo.go`

```go
package repository

import (
    "context"
    "database/sql"
    "fmt"
    "time"

    "github.com/GiaBao0510/Go-Realtime/internal/model"
)

type notificationRepo struct {
    db *sql.DB
}

func NewNotificationRepo(db *sql.DB) NotificationRepository {
    return &notificationRepo{db: db}
}

// Create tạo notification mới, trả về notification_id
func (r *notificationRepo) Create(ctx context.Context, notif *model.Notification) (int64, error) {
    query := `
        INSERT INTO notification (type, content)
        VALUES ($1, $2)
        RETURNING notification_id, created_at
    `
    err := r.db.QueryRowContext(ctx, query, notif.Type, notif.Content).
        Scan(&notif.NotificationID, &notif.CreatedAt)
    return notif.NotificationID, err
}

// CreateUserNotification liên kết notification với user
func (r *notificationRepo) CreateUserNotification(
    ctx context.Context, uid string, notificationID int64,
) error {
    query := `
        INSERT INTO user_notification (uid, notification_id)
        VALUES ($1, $2)
    `
    _, err := r.db.ExecContext(ctx, query, uid, notificationID)
    return err
}

// GetUnreadByUser lấy tất cả notification chưa đọc của user
func (r *notificationRepo) GetUnreadByUser(
    ctx context.Context, uid string,
) ([]model.Notification, error) {
    query := `
        SELECT n.notification_id, n.type, n.content, n.created_at
        FROM notification n
        INNER JOIN user_notification un ON n.notification_id = un.notification_id
        WHERE un.uid = $1 AND un.is_read = FALSE
        ORDER BY n.created_at DESC
    `
    rows, err := r.db.QueryContext(ctx, query, uid)
    if err != nil {
        return nil, fmt.Errorf("query unread notifications: %w", err)
    }
    defer rows.Close()

    var notifications []model.Notification
    for rows.Next() {
        var n model.Notification
        if err := rows.Scan(&n.NotificationID, &n.Type, &n.Content, &n.CreatedAt); err != nil {
            return nil, fmt.Errorf("scan notification: %w", err)
        }
        notifications = append(notifications, n)
    }
    return notifications, rows.Err()
}

// MarkAsRead đánh dấu notification đã đọc
func (r *notificationRepo) MarkAsRead(
    ctx context.Context, uid string, notificationID int64,
) error {
    query := `
        UPDATE user_notification
        SET is_read = TRUE, read_at = $1
        WHERE uid = $2 AND notification_id = $3
    `
    _, err := r.db.ExecContext(ctx, query, time.Now(), uid, notificationID)
    return err
}

// CountUnread đếm số notification chưa đọc
func (r *notificationRepo) CountUnread(ctx context.Context, uid string) (int, error) {
    var count int
    err := r.db.QueryRowContext(ctx,
        "SELECT COUNT(*) FROM user_notification WHERE uid = $1 AND is_read = FALSE",
        uid,
    ).Scan(&count)
    return count, err
}
```

---

## 8. TẦNG SERVICE — BUSINESS LOGIC

### 8.1 File `internal/service/message_service.go`

```go
package service

import (
    "context"
    "fmt"

    "github.com/GiaBao0510/Go-Realtime/internal/model"
    "github.com/GiaBao0510/Go-Realtime/internal/repository"
)

// MessageService xử lý business logic liên quan đến message
type MessageService struct {
    msgRepo   repository.MessageRepository
    notifRepo repository.NotificationRepository
    convRepo  repository.ConversationRepository
}

// NewMessageService tạo instance mới
func NewMessageService(
    msgRepo repository.MessageRepository,
    notifRepo repository.NotificationRepository,
    convRepo repository.ConversationRepository,
) *MessageService {
    return &MessageService{
        msgRepo:   msgRepo,
        notifRepo: notifRepo,
        convRepo:  convRepo,
    }
}

// SendMessage lưu message vào DB + tạo notification cho các thành viên conversation
//
// Luồng:
//  1. Lưu message vào bảng `message`
//  2. Lấy danh sách thành viên conversation (trừ người gửi)
//  3. Tạo notification "new_message"
//  4. Liên kết notification với từng thành viên (bảng user_notification)
func (s *MessageService) SendMessage(
    ctx context.Context,
    conversationID int,
    senderUID string,
    content string,
    messageType string,
) (*model.DBMessage, error) {

    // 1. Lưu message
    msg := &model.DBMessage{
        ConversationID: conversationID,
        SenderUID:      &senderUID,
        Content:        content,
        MessageType:    messageType,
    }
    if err := s.msgRepo.Create(ctx, msg); err != nil {
        return nil, fmt.Errorf("lưu message thất bại: %w", err)
    }

    // 2. Lấy thành viên conversation
    members, err := s.convRepo.GetMembers(ctx, conversationID)
    if err != nil {
        // Log lỗi nhưng không fail request — message đã lưu thành công
        fmt.Printf("⚠️ Lỗi lấy danh sách thành viên: %v\n", err)
        return msg, nil
    }

    // 3. Tạo notification
    notif := &model.Notification{
        Type:    model.NotifTypeNewMessage,
        Content: fmt.Sprintf("Bạn có tin nhắn mới trong conversation #%d", conversationID),
    }
    notifID, err := s.notifRepo.Create(ctx, notif)
    if err != nil {
        fmt.Printf("⚠️ Lỗi tạo notification: %v\n", err)
        return msg, nil
    }

    // 4. Liên kết notification với từng thành viên (trừ người gửi)
    for _, member := range members {
        if member.UID == senderUID {
            continue // Không thông báo cho chính mình
        }
        if err := s.notifRepo.CreateUserNotification(ctx, member.UID, notifID); err != nil {
            fmt.Printf("⚠️ Lỗi tạo user_notification cho %s: %v\n", member.UID, err)
        }
    }

    return msg, nil
}

// GetConversationMessages lấy lịch sử chat có phân trang
func (s *MessageService) GetConversationMessages(
    ctx context.Context,
    conversationID int,
    page, pageSize int,
) ([]model.DBMessage, error) {
    offset := (page - 1) * pageSize
    return s.msgRepo.GetByConversation(ctx, conversationID, pageSize, offset)
}
```

### 8.2 File `internal/service/user_service.go`

```go
package service

import (
    "context"
    "fmt"

    "github.com/GiaBao0510/Go-Realtime/internal/model"
    "github.com/GiaBao0510/Go-Realtime/internal/repository"
)

type UserService struct {
    userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
    return &UserService{userRepo: userRepo}
}

// Register tạo user mới
// ⚠️ Trong thực tế: phải hash password trước khi lưu (dùng bcrypt)
func (s *UserService) Register(ctx context.Context, req model.CreateUserRequest) (*model.User, error) {
    // Kiểm tra email đã tồn tại chưa
    existing, err := s.userRepo.GetByEmail(ctx, req.Email)
    if err != nil {
        return nil, fmt.Errorf("kiểm tra email: %w", err)
    }
    if existing != nil {
        return nil, fmt.Errorf("email '%s' đã được sử dụng", req.Email)
    }

    // TODO: Hash password bằng bcrypt
    // hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

    user := &model.User{
        Name:         req.Name,
        Email:        req.Email,
        PasswordHash: req.Password, // ⚠️ PHẢI hash trong production!
    }

    if err := s.userRepo.Create(ctx, user); err != nil {
        return nil, fmt.Errorf("tạo user: %w", err)
    }

    return user, nil
}

// SetOnline cập nhật trạng thái online/offline
func (s *UserService) SetOnline(ctx context.Context, uid string, online bool) error {
    return s.userRepo.UpdateOnlineStatus(ctx, uid, online)
}

// GetOnlineCount lấy số người đang online (dùng cho SSE LiveStats)
func (s *UserService) GetOnlineCount(ctx context.Context) (int, error) {
    return s.userRepo.CountOnline(ctx)
}
```

---

## 9. TÍCH HỢP VÀO HANDLER

### 9.1 Cập nhật `ws_handler.go` — Lưu message khi broadcast

Thay đổi chính: **Trước khi broadcast message, lưu vào DB qua MessageService**.

```go
// Trong struct WSHandler, thêm field:
type WSHandler struct {
    hub        *hub.Hub
    msgService *service.MessageService // ← THÊM MỚI
}

func NewWSHandler(hub *hub.Hub, msgService *service.MessageService) *WSHandler {
    return &WSHandler{hub: hub, msgService: msgService}
}

// Trong readPump loop, sau khi parse incoming message, thêm:
// ─── LƯU VÀO DB TRƯỚC KHI BROADCAST ───
dbMsg, err := h.msgService.SendMessage(
    ctx,
    conversationID,         // Cần xác định conversation nào
    senderUID,              // UUID của user
    incoming.Content,
    model.MsgTypeText,
)
if err != nil {
    log.Printf("[WS] ⚠️ Lưu message thất bại: %v", err)
    // Vẫn tiếp tục broadcast — không block chat
}

// Broadcast message
h.hub.Broadcast(msgBytes)
```

### 9.2 Cập nhật `sse_handler.go` — Dùng dữ liệu thật từ DB

```go
// Trong struct SSEHandler, thêm field:
type SSEHandler struct {
    userService *service.UserService // ← THÊM MỚI
}

func NewSSEHandler(userService *service.UserService) *SSEHandler {
    return &SSEHandler{userService: userService}
}

// Trong vòng lặp gửi event, thay dữ liệu random bằng dữ liệu thật:
case t := <-ticker.C:
    // Lấy số user online thật từ DB
    onlineCount, err := h.userService.GetOnlineCount(r.Context())
    if err != nil {
        log.Printf("[SSE] ⚠️ Lỗi lấy online count: %v", err)
        onlineCount = 0
    }

    status := model.LiveStats{
        OnlineUsers: onlineCount,   // ← Dữ liệu thật
        ServerTime:  t.Format("15:04:05"),
        CPULoad:     getCPULoad(),  // TODO: Implement lấy CPU thật
        MemoryUsage: getMemUsage(), // TODO: Implement lấy RAM thật
    }
```

### 9.3 Cập nhật `main.go` — Dependency Injection

```go
package main

import (
    "log"
    "net/http"

    "github.com/joho/godotenv"

    "github.com/GiaBao0510/Go-Realtime/internal/config"
    "github.com/GiaBao0510/Go-Realtime/internal/database"
    "github.com/GiaBao0510/Go-Realtime/internal/handler"
    "github.com/GiaBao0510/Go-Realtime/internal/hub"
    "github.com/GiaBao0510/Go-Realtime/internal/repository"
    "github.com/GiaBao0510/Go-Realtime/internal/service"
)

func main() {
    // ──── 1. Load config ────
    if err := godotenv.Load(); err != nil {
        log.Println("⚠️  Không tìm thấy file .env, dùng biến môi trường hệ thống")
    }

    cfg, err := config.Load()
    if err != nil {
        log.Fatalf("❌ Lỗi load config: %v", err)
    }

    // ──── 2. Kết nối Database ────
    db, err := database.NewPostgresDB(cfg.DB)
    if err != nil {
        log.Fatalf("❌ Lỗi kết nối DB: %v", err)
    }
    defer db.Close()

    // ──── 3. Khởi tạo Repository ────
    userRepo   := repository.NewUserRepo(db)
    msgRepo    := repository.NewMessageRepo(db)
    convRepo   := repository.NewConversationRepo(db)
    notifRepo  := repository.NewNotificationRepo(db)

    // ──── 4. Khởi tạo Service ────
    userService := service.NewUserService(userRepo)
    msgService  := service.NewMessageService(msgRepo, notifRepo, convRepo)

    // ──── 5. Khởi tạo Hub ────
    wsHub := hub.NewHub()
    go wsHub.Run()

    // ──── 6. Đăng ký Routes ────
    mux := http.NewServeMux()
    mux.Handle("/sse", handler.NewSSEHandler(userService))
    mux.Handle("/ws",  handler.NewWSHandler(wsHub, msgService))
    mux.Handle("/",    http.FileServer(http.Dir("./static")))

    // ──── 7. Chạy Server ────
    log.Printf("🚀 Server đang chạy tại http://localhost%s", cfg.Server.Port)
    if err := http.ListenAndServe(cfg.Server.Port, mux); err != nil {
        log.Fatalf("❌ Lỗi khởi động server: %v", err)
    }
}
```

---

## 10. QUY TRÌNH LƯU DỮ LIỆU THEO TỪNG CHỨC NĂNG

### 10.1 Khi user gửi tin nhắn (WebSocket)

```
Client gửi JSON ──▶ WSHandler.ServeHTTP()
    │
    ├──▶ Parse JSON (incomingMessage)
    │
    ├──▶ MessageService.SendMessage()
    │       ├──▶ MessageRepo.Create()           ← INSERT vào bảng `message`
    │       ├──▶ ConversationRepo.GetMembers()   ← SELECT thành viên
    │       ├──▶ NotificationRepo.Create()       ← INSERT vào bảng `notification`
    │       └──▶ NotificationRepo.CreateUserNotification()  ← INSERT user_notification
    │
    └──▶ Hub.Broadcast()  ← Gửi đến tất cả WebSocket client đang kết nối
```

### 10.2 Khi user kết nối / ngắt kết nối (WebSocket)

```
Client connect ──▶ WSHandler.ServeHTTP()
    │
    ├──▶ UserService.SetOnline(uid, true)    ← UPDATE users SET is_online = TRUE
    │
    └──▶ Hub.Register(client)

Client disconnect ──▶ defer cleanup
    │
    ├──▶ UserService.SetOnline(uid, false)   ← UPDATE users SET is_online = FALSE
    │
    └──▶ Hub.Unregister(client)
```

### 10.3 SSE push dữ liệu thống kê

```
Mỗi 5 giây ──▶ SSEHandler (ticker)
    │
    ├──▶ UserService.GetOnlineCount()    ← SELECT COUNT(*) FROM users WHERE is_online
    │
    └──▶ fmt.Fprintf(w, "event: stats-update\n")  ← Push SSE event xuống client
```

### 10.4 Tạo conversation mới (REST API)

```
POST /api/conversations ──▶ ConversationHandler
    │
    ├──▶ ConversationService.CreateGroup()
    │       ├──▶ ConversationRepo.Create()      ← INSERT conversation
    │       ├──▶ ConversationRepo.AddMember()    ← INSERT conversation_member (admin)
    │       └──▶ Lặp: AddMember() cho các thành viên khác
    │
    └──▶ Response JSON
```

### 10.5 Đọc notification (REST API)

```
GET /api/notifications ──▶ NotificationHandler
    │
    ├──▶ NotificationService.GetUnread(uid)
    │       └──▶ NotificationRepo.GetUnreadByUser()  ← SELECT ... WHERE is_read = FALSE
    │
    └──▶ Response JSON

PATCH /api/notifications/:id/read ──▶ NotificationHandler
    │
    └──▶ NotificationRepo.MarkAsRead()   ← UPDATE user_notification SET is_read = TRUE
```

---

## 11. BIẾN MÔI TRƯỜNG VÀ BẢO MẬT

### 11.1 File `.gitignore` [MỚI]

```gitignore
# Biến môi trường — KHÔNG BAO GIỜ COMMIT
.env

# Binary
go_RealTime
*.exe

# IDE
.idea/
.vscode/

# OS
.DS_Store
Thumbs.db
```

### 11.2 Các nguyên tắc bảo mật

| Nguyên tắc                     | Chi tiết                                                                       |
| ------------------------------ | ------------------------------------------------------------------------------ |
| **Không hardcode credentials** | Mọi thông tin nhạy cảm (password, key) PHẢI nằm trong `.env`                   |
| **Không commit `.env`**        | Thêm `.env` vào `.gitignore` TRƯỚC khi commit đầu tiên                         |
| **Hash password**              | Dùng `bcrypt` với cost ≥ 12 để hash password trước khi lưu                     |
| **SQL injection**              | Luôn dùng parameterized query (`$1`, `$2`) — KHÔNG BAO GIỜ nối chuỗi           |
| **CORS**                       | Cấu hình `Access-Control-Allow-Origin` cụ thể, tránh dùng `*` trong production |
| **SSL**                        | Supabase yêu cầu `sslmode=require` — đảm bảo luôn bật                          |

### 11.3 Thông tin Supabase hiện tại

> ⚠️ **CẢNH BÁO**: Thông tin bên dưới chỉ dùng cho development. KHÔNG dùng trong production.

| Thông tin         | Giá trị                                                                              |
| ----------------- | ------------------------------------------------------------------------------------ |
| Project           | Chat_Realtime_temp                                                                   |
| Host              | `db.auuceyrifwxwmsbchxge.supabase.co`                                                |
| Port              | `5432`                                                                               |
| User              | `postgres`                                                                           |
| Database          | `postgres`                                                                           |
| SSL Mode          | `require`                                                                            |
| Direct connection | `postgresql://postgres:[PASSWORD]@db.auuceyrifwxwmsbchxge.supabase.co:5432/postgres` |

---

## 12. TESTING

### 12.1 Cấu trúc test

```
internal/
├── repository/
│   ├── message_repo.go
│   └── message_repo_test.go     ← Integration test (cần DB thật)
├── service/
│   ├── message_service.go
│   └── message_service_test.go  ← Unit test (mock repository)
└── handler/
    └── ws_handler_test.go       ← Unit test (mock service)
```

### 12.2 Ví dụ unit test cho Service (mock repo)

```go
// internal/service/message_service_test.go
package service_test

import (
    "context"
    "testing"

    "github.com/GiaBao0510/Go-Realtime/internal/model"
    "github.com/GiaBao0510/Go-Realtime/internal/service"
)

// mockMessageRepo implement MessageRepository cho testing
type mockMessageRepo struct {
    messages []model.DBMessage
}

func (m *mockMessageRepo) Create(ctx context.Context, msg *model.DBMessage) error {
    msg.MessageID = int64(len(m.messages) + 1)
    m.messages = append(m.messages, *msg)
    return nil
}

// ... implement các method khác ...

func TestSendMessage_Success(t *testing.T) {
    msgRepo := &mockMessageRepo{}
    // ... setup mock repos ...

    svc := service.NewMessageService(msgRepo, notifRepo, convRepo)

    msg, err := svc.SendMessage(context.Background(), 1, "user-123", "Hello!", "text")
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if msg.MessageID == 0 {
        t.Error("expected message_id > 0")
    }
    if len(msgRepo.messages) != 1 {
        t.Errorf("expected 1 message in repo, got %d", len(msgRepo.messages))
    }
}
```

### 12.3 Chạy test

```bash
# Chạy tất cả test
go test ./...

# Chạy test với verbose
go test -v ./internal/service/...

# Chạy test với coverage
go test -cover ./...
```

---

## 13. CHECKLIST TRIỂN KHAI

Thực hiện theo thứ tự này để đảm bảo mỗi bước đều có thể compile và test:

### Phase 1: Chuẩn bị (không ảnh hưởng code hiện tại)

- [ ] Tạo file `.env` với thông tin Supabase
- [ ] Tạo file `.env.example` (không chứa password)
- [ ] Tạo file `.gitignore`
- [ ] Chạy `go get github.com/jackc/pgx/v5` và `go get github.com/joho/godotenv`
- [ ] Kiểm tra kết nối DB bằng `psql` hoặc tool khác

### Phase 2: Infrastructure layer

- [ ] Tạo `internal/config/config.go`
- [ ] Tạo `internal/database/postgres.go`
- [ ] Test kết nối DB trong `main.go` (chỉ ping, chưa dùng)

### Phase 3: Model layer

- [ ] Tạo `internal/model/user.go`
- [ ] Tạo `internal/model/conversation.go`
- [ ] Tạo `internal/model/notification.go`
- [ ] Tạo `internal/model/event.go`
- [ ] Cập nhật `internal/model/message.go` (thêm `DBMessage`)

### Phase 4: Repository layer

- [ ] Tạo `internal/repository/interfaces.go`
- [ ] Tạo `internal/repository/user_repo.go`
- [ ] Tạo `internal/repository/message_repo.go`
- [ ] Tạo `internal/repository/conversation_repo.go`
- [ ] Tạo `internal/repository/notification_repo.go`

### Phase 5: Service layer

- [ ] Tạo `internal/service/user_service.go`
- [ ] Tạo `internal/service/message_service.go`

### Phase 6: Tích hợp (cập nhật code hiện tại)

- [ ] **Sửa lỗi SSE**: Đổi `status-update` → `stats-update` trong `sse_handler.go`
- [ ] Cập nhật `SSEHandler` — inject `UserService`
- [ ] Cập nhật `WSHandler` — inject `MessageService`
- [ ] Cập nhật `main.go` — Dependency Injection đầy đủ

### Phase 7: Testing & Verification

- [ ] Viết unit test cho service layer
- [ ] Chạy `go test ./...`
- [ ] Test SSE bằng browser: `http://localhost:8080/sse_demo.html`
- [ ] Test WebSocket bằng browser: `http://localhost:8080/ws_demo.html`
- [ ] Kiểm tra DB có dữ liệu mới sau khi chat

---

## 📌 TÓM TẮT

| Vấn đề                  | Giải pháp                                                   |
| ----------------------- | ----------------------------------------------------------- |
| SSE FE không nhận event | Đổi tên event `status-update` → `stats-update` trong server |
| Dữ liệu mất khi restart | Thêm PostgreSQL (Supabase) để persist data                  |
| Code khó bảo trì        | Chia 3 tầng: Handler → Service → Repository                 |
| Khó mở rộng             | Dùng Interface + Dependency Injection                       |
| Bảo mật                 | `.env` + `.gitignore` + parameterized query + bcrypt        |
| Typo giữa FE/BE         | Dùng constant (`model.SSEEventStatsUpdate`)                 |

> 🎓 **Lời khuyên từ Senior**: Đừng cố làm tất cả cùng lúc. Hãy làm theo checklist, mỗi phase là một commit riêng biệt. Commit message rõ ràng (ví dụ: `feat: add database connection layer`, `fix: SSE event name mismatch`). Khi gặp lỗi, chỉ cần revert 1 commit thay vì refactor lại toàn bộ.
