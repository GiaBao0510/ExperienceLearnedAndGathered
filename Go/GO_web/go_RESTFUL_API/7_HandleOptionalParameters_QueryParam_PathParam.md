# Xử Lý Tham Số Tùy Chọn - Query Param và Path Param

## Mục lục

1. [Ôn lại: Path Param vs Query Param](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#%C3%B4n-l%E1%BA%A1i-path-param-vs-query-param)
2. [Điểm khác biệt quan trọng: Default Value](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#%C4%91i%E1%BB%83m-kh%C3%A1c-bi%E1%BB%87t-quan-tr%E1%BB%8Dng-default-value)
3. [Tình huống thực tế: API tin tức](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#t%C3%ACnh-hu%E1%BB%91ng-th%E1%BB%B1c-t%E1%BA%BF-api-tin-t%E1%BB%A9c)
4. [Vấn đề khi dùng cùng một hàm cho hai route](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#v%E1%BA%A5n-%C4%91%E1%BB%81-khi-d%C3%B9ng-c%C3%B9ng-m%E1%BB%99t-h%C3%A0m-cho-hai-route)
5. [Cài đặt handler `news.go`](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#c%C3%A0i-%C4%91%E1%BA%B7t-handler-newsgo)
6. [Cập nhật `main.go`](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#c%E1%BA%ADp-nh%E1%BA%ADt-maingo)
7. [Kiểm thử API](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#ki%E1%BB%83m-th%E1%BB%AD-api)
8. [Tổng kết](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#t%E1%BB%95ng-k%E1%BA%BFt)

---

## Ôn lại: Path Param vs Query Param

Gin cung cấp hai cách chính để client truyền tham số vào API:

### Path Param — tham số trong đường dẫn

```
GET /api/v1/news/cong-nghe-ai-2024
                 ↑
         đây là path param (:slug)
```

- Khai báo trong route bằng dấu `:` → `"/:slug"`
- Lấy giá trị bằng `c.Param("slug")`
- **Bắt buộc** — là một phần của URL, nếu thiếu sẽ không khớp route

### Query Param — tham số sau dấu `?`

```
GET /api/v1/news?slug=cong-nghe-ai-2024&page=2
                  ↑                      ↑
            query param               query param
```

- Không cần khai báo trong route
- Lấy giá trị bằng `c.Query("slug")`
- **Tùy chọn** — có thể vắng mặt, lúc đó `c.Query()` trả về chuỗi rỗng `""`

---

## Điểm khác biệt quan trọng: Default Value

Đây là điểm cốt lõi của bài này:

||Path Param|Query Param|
|---|---|---|
|Hàm lấy giá trị|`c.Param("key")`|`c.Query("key")`|
|Giá trị mặc định|❌ Không hỗ trợ|✅ `c.DefaultQuery("key", "default")`|
|Khi tham số vắng mặt|Route không khớp (404)|Trả về `""` hoặc giá trị mặc định|

### `c.DefaultQuery()` hoạt động như thế nào?

```go
// Nếu ?slug có giá trị → dùng giá trị đó
// Nếu ?slug không có hoặc rỗng → dùng "thong-tin-moi-cap-nhat"
slug := c.DefaultQuery("slug", "thong-tin-moi-cap-nhat")
```

Ví dụ minh họa:

```
GET /api/v1/news?slug=cong-nghe-ai
→ slug = "cong-nghe-ai"

GET /api/v1/news?slug=
→ slug = "thong-tin-moi-cap-nhat"   ← dùng default

GET /api/v1/news
→ slug = "thong-tin-moi-cap-nhat"   ← dùng default
```

> 💡 **`c.DefaultQuery` tương đương viết tắt của:**
> 
> ```go
> slug := c.Query("slug")
> if slug == "" {
>     slug = "thong-tin-moi-cap-nhat"
> }
> ```
> 
> Dùng `c.DefaultQuery` ngắn gọn và rõ ý định hơn.

---

## Tình huống thực tế: API tin tức

Chúng ta muốn xây dựng API tin tức hỗ trợ **hai cách truy cập**:

```
# Cách 1: Truy cập trực tiếp bằng slug trong URL (Path Param)
GET /api/v1/news/cong-nghe-ai-2024
→ Lấy bài viết có slug = "cong-nghe-ai-2024"

# Cách 2: Truy cập trang danh sách, slug là tùy chọn (Query Param)
GET /api/v1/news                          → Lấy tin mặc định
GET /api/v1/news?slug=cong-nghe-ai-2024   → Lọc theo slug
```

Để làm điều này, ta đăng ký **hai route** trỏ vào **cùng một handler**:

```go
news.GET("/", newsHandler_v1.GetNews)       // Route 1: không có slug
news.GET("/:slug", newsHandler_v1.GetNews)  // Route 2: có slug trong path
```

---

## Vấn đề khi dùng cùng một hàm cho hai route

Khi một hàm xử lý cả hai route, cần hiểu rõ Gin sẽ điều hướng như thế nào:

```
GET /api/v1/news          → khớp route "/"      → c.Param("slug") = ""
GET /api/v1/news/         → khớp route "/"      → c.Param("slug") = ""
GET /api/v1/news/ai-2024  → khớp route "/:slug" → c.Param("slug") = "ai-2024"
```

Khi `c.Param("slug")` trả về `""`, ta biết client đang truy cập route không có slug — đây là lúc dùng `c.DefaultQuery()` để lấy giá trị thay thế từ query string.

> ⚠️ **Lưu ý:** Không thể đặt cả `GET "/"` và `GET "/:slug"` trong cùng một group nếu Gin không thể phân biệt được. Thực tế Gin xử lý tốt trường hợp này vì `"/"` và `"/:slug"` là hai pattern rõ ràng khác nhau.

---

## Cài đặt handler `news.go`

Tạo file `Router_Group/internal/api/v1/handler/news.go`:

```go
package handler

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// NewsHandler là struct chứa các phương thức xử lý request liên quan đến tin tức
type NewsHandler struct{}

// NewsHandlerConstructor là Constructor — tạo instance mới của NewsHandler
func NewsHandlerConstructor() *NewsHandler {
    return &NewsHandler{}
}

// GetNews xử lý cả hai route:
//   - GET /api/v1/news/          → slug lấy từ query param (có default)
//   - GET /api/v1/news/:slug     → slug lấy từ path param
//
// Luồng xử lý:
//   1. Thử lấy slug từ path param
//   2. Nếu rỗng (route không có /:slug), lấy từ query param với giá trị mặc định
func (obj *NewsHandler) GetNews(c *gin.Context) {

    // Bước 1: Thử lấy slug từ path param
    // Nếu client gọi GET /api/v1/news/cong-nghe → slug = "cong-nghe"
    // Nếu client gọi GET /api/v1/news/          → slug = "" (chuỗi rỗng)
    slug := c.Param("slug")

    if slug == "" {
        // Bước 2a: Không có path param → lấy từ query string
        // Ví dụ: GET /api/v1/news?slug=cong-nghe → slug = "cong-nghe"
        // Ví dụ: GET /api/v1/news                → slug = "thong-tin-moi-cap-nhat" (default)
        slug = c.DefaultQuery("slug", "thong-tin-moi-cap-nhat")

        c.JSON(http.StatusOK, gin.H{
            "source":  "query param (or default)",
            "message": slug,
        })
    } else {
        // Bước 2b: Có path param → dùng trực tiếp
        c.JSON(http.StatusOK, gin.H{
            "source":  "path param",
            "message": slug,
        })
    }
}
```

> 💡 **Tại sao thêm trường `"source"` vào response?** Khi học và kiểm thử, biết slug đến từ đâu (path param hay query param hay default) giúp hiểu rõ luồng xử lý hơn. Trong production thực tế, không cần trường này.

---

## Cập nhật `main.go`

Thêm handler và route cho News vào `main.go`:

```go
package main

import (
    v1Handler "router-group/internal/api/v1/handler"
    v2Handler "router-group/internal/api/v2/handler"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Khởi tạo các Handler
    userHandler_v1     := v1Handler.NewUser()
    productHandler_v1  := v1Handler.NewProduct()
    categoryHandler_v1 := v1Handler.NewCategoryHandler()
    newsHandler_v1     := v1Handler.NewsHandlerConstructor() // Handler mới

    userHandler_v2 := v2Handler.NewUser()

    v1 := r.Group("/api/v1")
    {
        // ... các group user, product, category như bài trước ...

        // Group mới: News
        // Hai route cùng trỏ vào một handler GetNews
        news := v1.Group("/news")
        {
            news.GET("/", newsHandler_v1.GetNews)        // Không có slug → dùng query param
            news.GET("/:slug", newsHandler_v1.GetNews)   // Có slug trong path
        }
    }

    v2 := r.Group("/api/v2")
    {
        // ...
    }

    r.Run(":8080")
}
```

---

## Kiểm thử API

Khởi động server: `go run main.go`

### Trường hợp 1: Có path param

```bash
curl http://localhost:8080/api/v1/news/cong-nghe-ai-2024
```

Kết quả mong đợi:

```json
{
    "source": "path param",
    "message": "cong-nghe-ai-2024"
}
```

### Trường hợp 2: Có query param

```bash
curl "http://localhost:8080/api/v1/news?slug=the-thao-hom-nay"
```

Kết quả mong đợi:

```json
{
    "source": "query param (or default)",
    "message": "the-thao-hom-nay"
}
```

### Trường hợp 3: Không có tham số nào → dùng default

```bash
curl http://localhost:8080/api/v1/news
```

Kết quả mong đợi:

```json
{
    "source": "query param (or default)",
    "message": "thong-tin-moi-cap-nhat"
}
```

### Trường hợp 4: Query param rỗng → dùng default

```bash
curl "http://localhost:8080/api/v1/news?slug="
```

Kết quả mong đợi:

```json
{
    "source": "query param (or default)",
    "message": "thong-tin-moi-cap-nhat"
}
```

### Bảng tóm tắt kết quả kiểm thử

|URL|slug đến từ|Giá trị slug|
|---|---|---|
|`/api/v1/news/cong-nghe`|Path param|`"cong-nghe"`|
|`/api/v1/news?slug=the-thao`|Query param|`"the-thao"`|
|`/api/v1/news`|Default|`"thong-tin-moi-cap-nhat"`|
|`/api/v1/news?slug=`|Default|`"thong-tin-moi-cap-nhat"`|

---

## Tổng kết

Bài này tập trung vào sự khác biệt quan trọng giữa Path Param và Query Param khi xử lý tham số tùy chọn:

||`c.Param()`|`c.Query()`|`c.DefaultQuery()`|
|---|---|---|---|
|**Nguồn**|URL path|Query string|Query string|
|**Bắt buộc**|Có|Không|Không|
|**Khi vắng mặt**|Route không khớp|Trả về `""`|Trả về giá trị mặc định|
|**Dùng khi nào**|ID tài nguyên bắt buộc|Tham số lọc/tìm kiếm|Tham số có giá trị mặc định hợp lý|

### Khi nào nên dùng cái nào?

- **Path Param** (`/:slug`) — khi tham số **xác định tài nguyên cụ thể**, không thể thiếu. Ví dụ: `/users/:id`, `/products/:slug`.
- **Query Param** (`?page=2`) — khi tham số là **bộ lọc, tùy chọn phân trang, hoặc có thể vắng mặt**. Ví dụ: `?search=iphone`, `?page=1&limit=10`.
- **`DefaultQuery`** — khi muốn đảm bảo luôn có giá trị hợp lý dù client không truyền vào, tránh phải kiểm tra `== ""` thủ công.