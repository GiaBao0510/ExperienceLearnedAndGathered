
# Components và Communication - Giải Mã Kiến Trúc Hệ Thống Phân Tán

Học cách phân tích hệ thống thành components và hiểu communication patterns. Master data flow, bottleneck detection, và sync vs async processing trong distributed systems.

# Components và Communication: Hiểu Bản Chất Của Một Hệ Thống

Lần đầu tiên tôi phải debug một production issue, tôi hoàn toàn choáng váng.

API response chậm. Nhưng chậm ở đâu? Database? Application? Network? Cache?

Senior architect ngồi bên cạnh, vẽ một diagram đơn giản:

```
Client → Load Balancer → App Server → Cache → Database
```

Rồi anh hỏi: _"Request đi qua bao nhiêu bước? Mỗi bước mất bao lâu?"_

Đó là lần đầu tiên tôi hiểu: **Hệ thống không phải là một khối code. Nó là tập hợp các components tương tác với nhau.**

## Tại Sao Phải Hiểu Components?

Khi bạn code, bạn nhìn vào functions, classes, modules. Đó là **code-level view**.

Khi bạn design systems, bạn phải nhìn vào **components và interactions**. Đó là **architecture-level view**.

Sự khác biệt?

**Code-level:**

```python
def get_user(user_id):
    return db.query("SELECT * FROM users WHERE id = ?", user_id)
```

**Architecture-level:**

```
Client → API Gateway → User Service → Database
         (50ms)         (10ms)         (200ms)
         
Total latency: 260ms
Bottleneck: Database (200ms = 77% of total time)

```

Architecture-level view giúp bạn:

-   Identify bottlenecks nhanh chóng
-   Understand failure points
-   Design for scale
-   Communicate với team

**Bạn không thể optimize cái bạn không hiểu.**

## Anatomy Của Một Hệ Thống Web

Mọi web system, dù đơn giản hay phức tạp, đều có cùng building blocks cơ bản.

### Component 1: Client (Browser/Mobile App)

**Vai trò:** User interface, gửi requests, render responses

**Đặc điểm:**

-   Không đáng tin cậy (user có thể modify code)
-   Không kiểm soát được environment (network, device)
-   Stateful (lưu data locally)

**Implication cho architect:**

```
NEVER trust client
- Không validate ở client là đủ
- Không lưu sensitive data
- Không rely vào client-side logic cho business rules

ALWAYS validate server-side
- Client validation chỉ là UX enhancement
- Server validation là security requirement

```

**Ví dụ thực tế:**

Tôi từng thấy một app validate giá sản phẩm ở client. Hacker modify JavaScript, đổi giá từ $100 thành $1, submit form.

Server tin tưởng client → Database lưu order với giá $1 → Company mất tiền.

**Bài học:** Client là presentation layer, không phải business logic layer.

### Component 2: Load Balancer

**Vai trò:** Phân phối traffic đều giữa multiple servers

**Khi nào cần?**

-   Khi 1 server không đủ xử lý traffic
-   Khi cần high availability (1 server die, còn servers khác)

**Khi nào KHÔNG cần?**

-   Startup với < 1000 users
-   Internal tools với ít traffic
-   Prototypes/MVPs

**Trade-offs:**

```
Gains:
- Distribute load evenly
- No single point of failure
- Rolling deployment (update từng server)

Costs:
- Thêm 1 layer complexity
- Load balancer itself có thể là SPOF (nếu không redundant)
- Network latency (+10-50ms)

```

**Algorithms phổ biến:**

**Round Robin:**

```
Request 1 → Server A
Request 2 → Server B  
Request 3 → Server C
Request 4 → Server A (repeat)

```

Đơn giản, phân phối đều. Nhưng không care server nào đang overload.

**Least Connections:**

```
Server A: 50 active connections
Server B: 20 active connections
Server C: 35 active connections

→ New request goes to Server B

```

Smart hơn, nhưng cần track state.

**Weighted Round Robin:**

```
Server A (16GB RAM): weight = 2
Server B (8GB RAM): weight = 1

→ A gets 2 requests for every 1 request to B

```

Useful khi servers có specs khác nhau.

**Personal insight:** Trong thực tế, tôi thấy Round Robin là sufficient cho 90% cases. Least Connections chỉ cần khi requests có execution time rất khác nhau.

### Component 3: Application Server

**Vai trò:** Business logic, process requests, orchestrate data flow

**Stateless vs Stateful - Quyết định quan trọng:**

**Stateless Server:**

```python
# Không lưu user session trong memory
def handle_request(request):
    user_id = jwt.decode(request.token)  # Get from token
    user = db.get_user(user_id)          # Get from DB
    return process(user)
    
# Benefit: Any server can handle any request
# Can scale horizontally easily

```

**Stateful Server:**

```python
# Lưu user session trong memory
sessions = {}  # In-memory storage

def handle_login(user_id):
    sessions[user_id] = UserSession(user_id)
    
def handle_request(user_id):
    session = sessions[user_id]  # Retrieve from memory
    return process(session)
    
# Problem: User phải hit same server
# Scaling is hard (need sticky sessions)

```

**Trade-off decision:**

```
Choose Stateless when:
- Need horizontal scaling
- Multiple servers behind load balancer
- Cloud environment (servers can die anytime)

Choose Stateful when:
- Need extremely low latency (no DB roundtrip)
- WebSocket connections (need persistent connection)
- Small scale (1-2 servers)

```

**Tôi recommend:** Default to stateless. Stateful chỉ khi có lý do rất tốt.

### Component 4: Cache Layer

**Vai trò:** Lưu frequently accessed data in-memory để giảm DB load

**Khi nào dùng cache?**

```
Cache khi:
- Read >> Write (ratio 10:1 or higher)
- Data không thay đổi thường xuyên
- Expensive queries (complex JOINs, aggregations)
- High traffic endpoints

Không cache khi:
- Write-heavy workload
- Data thay đổi liên tục (real-time stock prices)
- Personalized data per user (mỗi user khác nhau)

```

**Ví dụ thực tế: E-commerce Product Page**

```python
# Product info: Perfect for cache
product = cache.get(f"product:{id}")
if not product:
    product = db.query("SELECT * FROM products WHERE id = ?", id)
    cache.set(f"product:{id}", product, ttl=3600)  # 1 hour

# Inventory count: DON'T cache (changes frequently)
stock = db.query("SELECT stock FROM inventory WHERE product_id = ?", id)

# User's cart: DON'T cache (personalized)
cart = db.query("SELECT * FROM carts WHERE user_id = ?", user_id)

```

**Cache Invalidation - The Hard Problem:**

Phil Karlton nói: _"There are only two hard things in Computer Science: cache invalidation and naming things."_

Anh ấy đúng. Cache invalidation là nightmare.

**Problem:**

```
1. User updates profile → Database updated
2. Cache vẫn có old data
3. User refresh → Sees old data (cache hit)
4. User confused, files bug report

```

**Solutions:**

**Option 1: TTL (Time To Live)**

```python
cache.set("user:123", user, ttl=300)  # Expire after 5 minutes

Simple
Data stale for up to 5 minutes

```

**Option 2: Active Invalidation**

```python
def update_user(user_id, data):
    db.update("users", user_id, data)
    cache.delete(f"user:{user_id}")  # Delete immediately
    
Always fresh
Must remember to invalidate everywhere

```

**Option 3: Cache-Aside Pattern (Recommended)**

```python
def get_user(user_id):
    user = cache.get(f"user:{user_id}")
    if user:
        return user
    
    user = db.query("SELECT * FROM users WHERE id = ?", user_id)
    cache.set(f"user:{user_id}", user, ttl=3600)
    return user

def update_user(user_id, data):
    db.update("users", user_id, data)
    cache.delete(f"user:{user_id}")  # Invalidate
    # Next read will miss cache → Fetch fresh from DB → Cache again

```

**Personal rule:** Start simple với TTL. Chỉ thêm active invalidation khi data freshness thực sự critical.

### Component 5: Database

**Vai trò:** Persistent storage, source of truth

**Database là bottleneck phổ biến nhất.**

Tại sao? Vì disk I/O chậm hơn memory 100,000 lần.

```
Memory read: 100 nanoseconds
Disk read: 10 milliseconds

10ms / 100ns = 100,000x slower

```

**Implications:**

1.  **Minimize database calls**
    
    -   Use cache
    -   Batch queries
    -   Eager loading (N+1 problem)
2.  **Optimize queries**
    
    -   Add indexes
    -   Avoid SELECT *
    -   Use EXPLAIN ANALYZE
3.  **Scale database last**
    
    -   Optimize application first
    -   Add cache second
    -   Scale DB only when necessary (expensive)

**Ví dụ real-world:**

Một team complain: "Database quá chậm, cần upgrade server!"

Tôi check query logs:

```sql
-- Running 1000 times per second
SELECT * FROM users WHERE email = 'john@example.com';
-- Execution time: 500ms

```

Vấn đề? Không có index trên email column → Full table scan.

Solution:

```sql
CREATE INDEX idx_users_email ON users(email);
-- Execution time: 5ms (100x faster!)

```

**Cost:** $0 và 30 giây setup time.

Thay vì upgrade server ($500/month), fix query (free). **Always optimize before scaling.**

## Communication Patterns: Sync vs Async

Đây là một trong những decisions quan trọng nhất trong system design.

### Synchronous (Request-Response)

**How it works:**

```
Client → [HTTP Request] → Server
Client ← [Wait...] ←
Client ← [HTTP Response] ← Server

```

**Đặc điểm:**

-   Client blocks và đợi response
-   Immediate feedback
-   Tight coupling (client phụ thuộc server)

**Khi nào dùng:**

-   User cần kết quả ngay lập tức
-   Simple workflows
-   CRUD operations

**Ví dụ:**

```python
# User login
def login(email, password):
    user = authenticate(email, password)  # Must know now
    if user:
        token = generate_token(user)
        return {"token": token}           # Return immediately
    return {"error": "Invalid credentials"}

```

### Asynchronous (Message Queue)

**How it works:**

```
Client → [HTTP Request] → Server → [Add to Queue] → Return immediately
                                         ↓
                                    Background Worker → Process job

```

**Đặc điểm:**

-   Client không đợi
-   Eventual processing
-   Loose coupling (client không care who processes)

**Khi nào dùng:**

-   Process mất nhiều thời gian (> 5 seconds)
-   User không cần kết quả ngay
-   Need to handle traffic spikes

**Ví dụ:**

```python
# User uploads video
def upload_video(video_file):
    video_id = save_to_s3(video_file)              # Quick (2s)
    queue.add({
        "job": "process_video",
        "video_id": video_id,
        "tasks": ["thumbnail", "transcode", "notify"]
    })
    return {"status": "processing", "id": video_id}  # Return immediately
    
# Background worker (runs separately)
def process_video(video_id):
    generate_thumbnail(video_id)    # 5s
    transcode_720p(video_id)        # 30s
    transcode_1080p(video_id)       # 60s
    send_notification(video_id)     # 1s

```

**User experience:**

-   Upload → Gets response sau 2s (tốt)
-   Video ready → Notification sau 1-2 phút (acceptable)

Nếu làm sync → User đợi 98 giây → Timeout → Bad UX.

### Trade-off Matrix

```
Synchronous:
Simple logic
Immediate feedback
Easy debugging
Poor UX for slow operations
Timeout risk
Can't handle traffic spikes

Asynchronous:
Better UX (fast response)
Can handle spikes (queue buffers)
Retry logic built-in
Complex (need queue infrastructure)
Eventual consistency
Harder debugging (distributed tracing)

```

**Decision framework:**

```python
if user_needs_result_immediately:
    use_synchronous()
elif process_time > 5_seconds:
    use_asynchronous()
elif traffic_spiky:  # Black Friday, viral post
    use_asynchronous()
else:
    use_synchronous()  # Simpler

```

## Data Flow & Bottleneck Detection

**Golden rule:** Hệ thống chỉ nhanh bằng component chậm nhất.

### Measuring Data Flow

**Example flow: User loads homepage**

```
Step 1: DNS lookup               → 20ms
Step 2: TCP handshake           → 50ms
Step 3: TLS negotiation         → 100ms
Step 4: HTTP request to server  → 10ms
Step 5: Server processing       → 30ms
Step 6: Database query          → 500ms ← BOTTLENECK
Step 7: Server render HTML      → 20ms
Step 8: Response to client      → 10ms
Step 9: Browser parse/render    → 200ms

Total: 940ms
Bottleneck: Database (500ms = 53% of total)

```

**Action plan:**

1.  Optimize database query (add index)
2.  Add cache layer
3.  Measure again

**Sau optimization:**

```
Step 6: Cache hit               → 5ms

Total: 445ms (52% faster!)

```

**Lesson:** Đừng optimize bừa. Measure → Find bottleneck → Optimize bottleneck.

### Common Bottlenecks & Solutions

**Bottleneck 1: Database**

```
Symptoms: High query latency, CPU/IO maxed out
Solutions:
- Add indexes
- Optimize queries
- Add read replicas
- Implement caching

```

**Bottleneck 2: Network**

```
Symptoms: High latency, timeouts
Solutions:
- CDN for static assets
- Reduce payload size (compression)
- Fewer HTTP requests (bundling)

```

**Bottleneck 3: Application Logic**

```
Symptoms: High CPU usage, slow response
Solutions:
- Profile code
- Optimize algorithms (O(n²) → O(n log n))
- Add caching
- Horizontal scaling

```
