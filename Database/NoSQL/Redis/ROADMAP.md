## 🗺️ **LỘ TRÌNH HỌC REDIS TỪ CƠ BẢN ĐẾN NÂNG CAO**

### 🧱 Giai đoạn 1: **Kiến thức Cơ bản**

#### 1. Tổng quan về Redis

- Redis là gì? (In-memory data store)
- Redis dùng để làm gì? (Cache, pub/sub, queue, session store, real-time data store,...)
- Sự khác biệt giữa Redis và các hệ quản trị CSDL khác (MySQL, MongoDB...)

#### 2. Cài đặt Redis

- Cài Redis trên:
    - Windows (WSL / Docker / Redis Stack)
    - Linux (Ubuntu)
    - macOS (Homebrew)

- Redis CLI (`redis-cli`)
- Redis Insight (tool UI để quản lý Redis)

#### 3. Cấu trúc dữ liệu cơ bản trong Redis

- `String`
- `List`
- `Set`
- `Sorted Set
- `Hash
- `Bitmap`, `HyperLogLog` (sẽ học sau)

> 👉 _Tập làm quen với lệnh CLI: `SET`, `GET`, `INCR`, `LPUSH`, `SADD`, `ZADD`, `HSET`, `HGET`, `DEL`, `EXISTS`..._

---

### ⚙️ Giai đoạn 2: **Tính năng trung cấp**

#### 4. Quản lý dữ liệu

- TTL – Time To Live (`EXPIRE`, `TTL`, `PERSIST`)
- Atomic operations (`INCR`, `DECR`, `MSET`, `MGET`, `DEL`)
- Transaction (`MULTI`, `EXEC`, `WATCH`)
- Backup và Persistence:
    - RDB (snapshot)
    - AOF (Append Only File)
    - Hybrid persistence
- Dump và Restore dữ liệu

#### 5. Pub/Sub
- `PUBLISH`, `SUBSCRIBE`
- Dùng trong real-time notification, chat

#### 6. Lua Script
- Chạy script Lua trong Redis (`EVAL`)
- Atomic multi-operations với Lua

---

### ⚡ Giai đoạn 3: **Ứng dụng Redis trong thực tế**

#### 7. Redis làm Cache
- Cache Aside (write-through / read-through / write-behind)
- Expiration
- Ví dụ: Caching truy vấn DB, caching API response
- Dùng trong .NET / Node.js / Python...

#### 8. Redis làm Queue / Stream

- `LPUSH` / `BRPOP` — hàng đợi đơn giản
- Redis Streams (`XADD`, `XREAD`, `XACK`, `XGROUP`)
- So sánh Redis Streams vs Kafka
#### 9. Session Store

- Lưu session của user trong Redis
- Dùng với ExpressJS, ASP.NET, Django...

#### 10. Rate Limiting

- Dùng Redis để giới hạn request/user theo IP
- Thuật toán sliding window, token bucket, fixed window...

---

### 🧠 Giai đoạn 4: **Kiến thức Nâng cao**

#### 11. Redis Cluster và High Availability

- Redis Sentinel
- Redis Cluster (Sharding)
- Cơ chế tự động phân vùng dữ liệu
- Failover và Master/Replica

#### 12. Security & Performance

- Redis AUTH
- Redis ACL (Access Control List)
- Monitor, slowlog, stats (`INFO`, `MONITOR`, `SLOWLOG`)
- Redis Benchmark
- Connection pooling

#### 13. Redis với Docker & Kubernetes

- Redis image chính thức
- Redis với Helm chart
- Tự cấu hình Redis Sentinel Cluster trong môi trường production

#### 14. Redis Modules (nâng cao)

- RedisJSON (store document JSON)
- RediSearch (full-text search)
- RedisTimeSeries
- RedisBloom (filter/estimation)

---

### 🔨 Công cụ & Tài liệu hỗ trợ

- [https://redis.io/](https://redis.io/)
- Redis CLI: `redis-cli`
- RedisInsight (UI)
- [Awesome Redis GitHub](https://github.com/joeferner/redis-commander)
    
- Dự án thực hành:
    - Xây dựng hệ thống comment real-time
    - Xây dựng middleware Rate Limit
    - Caching product list trong e-commerce