# URL, Path Parameters và Query Parameters trong API

## 📋 Mục lục

1. [URL là gì?](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#1-url-l%C3%A0-g%C3%AC)
2. [Cấu trúc URL chi tiết](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#2-c%E1%BA%A5u-tr%C3%BAc-url-chi-ti%E1%BA%BFt)
3. [Endpoint trong API](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#3-endpoint-trong-api)
4. [Path Parameters](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#4-path-parameters)
5. [Query Parameters](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#5-query-parameters)
6. [So sánh Path vs Query Parameters](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#6-so-s%C3%A1nh-path-vs-query-parameters)
7. [Best Practices](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#7-best-practices)
8. [Ví dụ thực tế](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#8-v%C3%AD-d%E1%BB%A5-th%E1%BB%B1c-t%E1%BA%BF)

---

## 1. URL là gì?

**URL (Uniform Resource Locator)** là địa chỉ của một tài nguyên trên web, giống như **địa chỉ nhà** giúp chúng ta tìm đến đúng nơi cần đến.

### 🎯 Ví dụ thực tế

**Địa chỉ nhà:**

```
123 Nguyễn Huệ, Quận 1, TP.HCM, Việt Nam
```

**URL:**

```
https://api.example.com:8080/users/123?role=admin&active=true#profile
```

### 💡 URL trong API

Trong API, URL được dùng để:

- **Xác định tài nguyên**: users, products, orders
- **Xác định hành động**: GET, POST, PUT, DELETE
- **Truyền dữ liệu**: qua path hoặc query parameters

**Ví dụ:**

```bash
# Lấy danh sách sản phẩm
http://localhost:8080/products

# Tạo user mới
http://localhost:8080/users/create

# Lấy thông tin user có ID = 123
http://localhost:8080/users/123

# Tìm kiếm sản phẩm
http://localhost:8080/products?category=laptop&price_max=20000000
```

---

## 2. Cấu trúc URL chi tiết

### 2.1. Các thành phần của URL

```
https://api.example.com:8080/products/123?sort=price&order=asc#reviews

├─ Protocol: https://
├─ Domain: api.example.com
├─ Port: :8080
├─ Path: /products/123
├─ Query Parameters: ?sort=price&order=asc
└─ Fragment: #reviews
```

### 2.2. Giải thích từng thành phần

#### **1. Protocol (Giao thức)**

**Mục đích:** Quy định cách truyền dữ liệu.

```
http://   → Không mã hóa
https://  → Có mã hóa (an toàn hơn)
```

**Ví dụ:**

```
http://localhost:8080/api
https://api.example.com/users
```

#### **2. Domain Name (Tên miền)**

**Mục đích:** Xác định server chứa tài nguyên.

```
localhost        → Máy tính local
example.com      → Domain thật
api.example.com  → Subdomain cho API
192.168.1.100    → IP address
```

**Ví dụ:**

```
http://localhost:8080
https://api.facebook.com
https://graph.facebook.com
```

#### **3. Port (Cổng)**

**Mục đích:** Xác định cổng kết nối trên server.

```
:80    → HTTP mặc định
:443   → HTTPS mặc định
:8080  → Development thường dùng
:3000  → Node.js thường dùng
:5000  → Flask/Python thường dùng
```

**Ví dụ:**

```
http://localhost:8080
http://localhost:3000
https://example.com:443  (có thể bỏ :443 vì là mặc định)
```

#### **4. Path (Đường dẫn)**

**Mục đích:** Xác định tài nguyên cụ thể.

```
/products           → Tất cả sản phẩm
/products/123       → Sản phẩm có ID = 123
/users/john/orders  → Orders của user john
```

**Ví dụ:**

```
http://localhost:8080/users
http://localhost:8080/products/123
http://localhost:8080/api/v1/orders
```

#### **5. Query Parameters (Tham số truy vấn)**

**Mục đích:** Truyền thêm thông tin, lọc, sắp xếp.

**Cú pháp:**

```
?key1=value1&key2=value2&key3=value3
```

**Ví dụ:**

```
/products?category=laptop&price_max=20000000
/users?role=admin&active=true&page=2
/search?q=golang&sort=date&limit=10
```

#### **6. Fragment (Neo/Anchor)**

**Mục đích:** Trỏ đến vị trí cụ thể trong trang.

```
#section1
#profile
#comments
```

**Ví dụ:**

```
https://example.com/docs#installation
https://github.com/gin-gonic/gin#features
```

> ⚠️ **Lưu ý:** Fragment không được gửi lên server, chỉ dùng ở client.

### 2.3. Ví dụ phân tích URL đầy đủ

```
https://api.shopee.vn:443/products/123456?sort=price&page=2#reviews
```

**Phân tích:**

|Thành phần|Giá trị|Ý nghĩa|
|---|---|---|
|**Protocol**|`https://`|Kết nối bảo mật|
|**Domain**|`api.shopee.vn`|Server API của Shopee|
|**Port**|`:443`|HTTPS mặc định|
|**Path**|`/products/123456`|Sản phẩm ID = 123456|
|**Query**|`?sort=price&page=2`|Sắp xếp theo giá, trang 2|
|**Fragment**|`#reviews`|Nhảy đến phần reviews|

---

## 3. Endpoint trong API

### 3.1. Endpoint là gì?

**Endpoint** là điểm cuối (URL cụ thể) mà client gọi để thực hiện một hành động.

**Cấu trúc:**

```
Method + URL = Endpoint

GET    /users           → Lấy danh sách users
POST   /users           → Tạo user mới
GET    /users/123       → Lấy user ID = 123
PUT    /users/123       → Cập nhật user 123
DELETE /users/123       → Xóa user 123
```

### 3.2. Ví dụ cơ bản với Gin

**File:** `main.go`

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Tạo router
	r := gin.Default()

	// Endpoint 1: Trang chủ
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to API",
		})
	})

	// Endpoint 2: Demo
	r.GET("/demo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This is demo endpoint",
		})
	})

	// Endpoint 3: Lấy danh sách users
	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": []string{"John", "Jane", "Bob"},
		})
	})

	// Endpoint 4: Lấy danh sách products
	r.GET("/products", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": []string{"Laptop", "Mouse", "Keyboard"},
		})
	})

	// Chạy server ở port 8080
	r.Run(":8080")
}
```

### 3.3. Test các Endpoints

**Cách 1: Dùng trình duyệt**

```
http://localhost:8080/
http://localhost:8080/demo
http://localhost:8080/users
http://localhost:8080/products
```

**Cách 2: Dùng curl**

```bash
curl http://localhost:8080/
curl http://localhost:8080/users
curl http://localhost:8080/products
```

**Output:**

```json
// GET /users
{
    "data": ["John", "Jane", "Bob"]
}

// GET /products
{
    "data": ["Laptop", "Mouse", "Keyboard"]
}
```

**Cách 3: Dùng Postman hoặc Thunder Client**

- Method: GET
- URL: `http://localhost:8080/users`
- Send

---

## 4. Path Parameters

### 4.1. Path Parameter là gì?

**Path Parameters** là tham số **bắt buộc** nằm ngay trong đường dẫn URL, thường dùng để **xác định tài nguyên cụ thể**.

**Đặc điểm:**

- Là phần **bắt buộc** của URL
- Dùng để **định danh** tài nguyên (ID, username, slug)
- Không thể bỏ qua

**Cú pháp:**

```go
/users/:id        // :id là path parameter
/products/:id
/posts/:slug
/users/:userId/orders/:orderId
```

### 4.2. Ví dụ cơ bản

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Path Parameter: :id
	r.GET("/users/:id", func(c *gin.Context) {
		// Lấy giá trị của :id
		id := c.Param("id")
		
		c.JSON(200, gin.H{
			"message": "Get user by ID",
			"user_id": id,
		})
	})

	r.Run(":8080")
}
```

**Test:**

```bash
curl http://localhost:8080/users/123
curl http://localhost:8080/users/456
curl http://localhost:8080/users/abc
```

**Output:**

```json
// GET /users/123
{
    "message": "Get user by ID",
    "user_id": "123"
}

// GET /users/456
{
    "message": "Get user by ID",
    "user_id": "456"
}

// GET /users/abc
{
    "message": "Get user by ID",
    "user_id": "abc"
}
```

### 4.3. Nhiều Path Parameters

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Nhiều path parameters
	r.GET("/users/:userId/orders/:orderId", func(c *gin.Context) {
		userId := c.Param("userId")
		orderId := c.Param("orderId")
		
		c.JSON(200, gin.H{
			"message":  "Get order of user",
			"user_id":  userId,
			"order_id": orderId,
		})
	})

	r.Run(":8080")
}
```

**Test:**

```bash
curl http://localhost:8080/users/123/orders/789
```

**Output:**

```json
{
    "message": "Get order of user",
    "user_id": "123",
    "order_id": "789"
}
```

### 4.4. Ví dụ thực tế với Database

```go
package main

import (
	"github.com/gin-gonic/gin"
)

// Giả lập database
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{ID: "1", Name: "John Doe", Age: 25},
	{ID: "2", Name: "Jane Smith", Age: 30},
	{ID: "3", Name: "Bob Johnson", Age: 35},
}

func main() {
	r := gin.Default()

	// Get user by ID
	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		
		// Tìm user trong database
		for _, user := range users {
			if user.ID == id {
				c.JSON(200, gin.H{
					"success": true,
					"data":    user,
				})
				return
			}
		}
		
		// Không tìm thấy
		c.JSON(404, gin.H{
			"success": false,
			"message": "User not found",
		})
	})

	r.Run(":8080")
}
```

**Test:**

```bash
# Tìm thấy
curl http://localhost:8080/users/1

# Không tìm thấy
curl http://localhost:8080/users/999
```

**Output:**

```json
// Tìm thấy (Status 200)
{
    "success": true,
    "data": {
        "id": "1",
        "name": "John Doe",
        "age": 25
    }
}

// Không tìm thấy (Status 404)
{
    "success": false,
    "message": "User not found"
}
```

---

## 5. Query Parameters

### 5.1. Query Parameter là gì?

**Query Parameters** là tham số **tùy chọn** nằm sau dấu `?` trong URL, thường dùng để **lọc, tìm kiếm, sắp xếp**.

**Đặc điểm:**

- Là phần **tùy chọn** của URL
- Dùng để **lọc, tìm kiếm, phân trang, sắp xếp**
- Có thể bỏ qua

**Cú pháp:**

```
/products?category=laptop&price_max=20000000&sort=price
           ↑
         Bắt đầu query parameters
         
?key1=value1&key2=value2&key3=value3
             ↑
           Phân cách bằng &
```

### 5.2. Ví dụ cơ bản

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Query Parameters
	r.GET("/search", func(c *gin.Context) {
		// Lấy query parameters
		q := c.Query("q")           // Từ khóa tìm kiếm
		sort := c.Query("sort")     // Sắp xếp theo
		limit := c.Query("limit")   // Giới hạn kết quả
		
		c.JSON(200, gin.H{
			"query":  q,
			"sort":   sort,
			"limit":  limit,
		})
	})

	r.Run(":8080")
}
```

**Test:**

```bash
curl "http://localhost:8080/search?q=golang&sort=date&limit=10"
```

**Output:**

```json
{
    "query": "golang",
    "sort": "date",
    "limit": "10"
}
```

### 5.3. Query Parameters với giá trị mặc định

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/products", func(c *gin.Context) {
		// DefaultQuery: Nếu không có thì dùng giá trị mặc định
		category := c.DefaultQuery("category", "all")
		sort := c.DefaultQuery("sort", "name")
		page := c.DefaultQuery("page", "1")
		limit := c.DefaultQuery("limit", "20")
		
		c.JSON(200, gin.H{
			"category": category,
			"sort":     sort,
			"page":     page,
			"limit":    limit,
		})
	})

	r.Run(":8080")
}
```

**Test:**

```bash
# Không truyền parameters (dùng mặc định)
curl http://localhost:8080/products

# Truyền một số parameters
curl "http://localhost:8080/products?category=laptop&page=2"

# Truyền đầy đủ parameters
curl "http://localhost:8080/products?category=phone&sort=price&page=3&limit=50"
```

**Output:**

```json
// Không truyền (mặc định)
{
    "category": "all",
    "sort": "name",
    "page": "1",
    "limit": "20"
}

// Truyền một số
{
    "category": "laptop",
    "sort": "name",       // Mặc định
    "page": "2",
    "limit": "20"         // Mặc định
}

// Truyền đầy đủ
{
    "category": "phone",
    "sort": "price",
    "page": "3",
    "limit": "50"
}
```

### 5.4. Ví dụ thực tế: Filter và Pagination

```go
package main

import (
	"strconv"
	"github.com/gin-gonic/gin"
)

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
}

var products = []Product{
	{1, "Macbook Pro", "laptop", 35000000},
	{2, "Dell XPS 13", "laptop", 25000000},
	{3, "iPhone 15", "phone", 25000000},
	{4, "Samsung S24", "phone", 20000000},
	{5, "iPad Pro", "tablet", 22000000},
	{6, "Surface Pro", "tablet", 28000000},
}

func main() {
	r := gin.Default()

	r.GET("/products", func(c *gin.Context) {
		// Lấy query parameters
		category := c.Query("category")
		priceMaxStr := c.Query("price_max")
		sort := c.DefaultQuery("sort", "name")
		
		// Filter theo category
		filtered := []Product{}
		for _, p := range products {
			// Filter category (nếu có)
			if category != "" && p.Category != category {
				continue
			}
			
			// Filter price_max (nếu có)
			if priceMaxStr != "" {
				priceMax, err := strconv.ParseFloat(priceMaxStr, 64)
				if err == nil && p.Price > priceMax {
					continue
				}
			}
			
			filtered = append(filtered, p)
		}
		
		c.JSON(200, gin.H{
			"total":    len(filtered),
			"category": category,
			"sort":     sort,
			"data":     filtered,
		})
	})

	r.Run(":8080")
}
```

**Test:**

```bash
# Tất cả sản phẩm
curl http://localhost:8080/products

# Chỉ laptop
curl http://localhost:8080/products?category=laptop

# Sản phẩm dưới 25 triệu
curl "http://localhost:8080/products?price_max=25000000"

# Laptop dưới 30 triệu
curl "http://localhost:8080/products?category=laptop&price_max=30000000"
```

**Output:**

```json
// Chỉ laptop
{
    "total": 2,
    "category": "laptop",
    "sort": "name",
    "data": [
        {
            "id": 1,
            "name": "Macbook Pro",
            "category": "laptop",
            "price": 35000000
        },
        {
            "id": 2,
            "name": "Dell XPS 13",
            "category": "laptop",
            "price": 25000000
        }
    ]
}

// Laptop dưới 30 triệu
{
    "total": 1,
    "category": "laptop",
    "sort": "name",
    "data": [
        {
            "id": 2,
            "name": "Dell XPS 13",
            "category": "laptop",
            "price": 25000000
        }
    ]
}
```

---

## 6. So sánh Path vs Query Parameters

### 6.1. Bảng so sánh

|Khía cạnh|Path Parameters|Query Parameters|
|---|---|---|
|**Vị trí**|Trong path|Sau dấu `?`|
|**Cú pháp**|`/users/:id`|`/users?role=admin`|
|**Bắt buộc**|✅ Có|❌ Không|
|**Mục đích**|Định danh tài nguyên|Lọc, tìm kiếm, options|
|**Ví dụ**|`/users/123`|`/users?active=true`|
|**SEO**|Tốt hơn|Kém hơn|
|**Readable**|Rõ ràng hơn|Dài hơn|

### 6.2. Khi nào dùng gì?

**✅ Dùng Path Parameters khi:**

- Xác định **tài nguyên cụ thể** (ID, username, slug)
- Tham số **bắt buộc** để endpoint hoạt động
- URL **ngắn gọn** và **dễ đọc**

**Ví dụ:**

```
GET /users/123              ✅ Lấy user ID = 123
GET /products/laptop-dell   ✅ Lấy sản phẩm có slug
GET /posts/how-to-code      ✅ Lấy bài viết
```

**✅ Dùng Query Parameters khi:**

- **Lọc, tìm kiếm, sắp xếp** dữ liệu
- Tham số **tùy chọn** (có thể bỏ qua)
- **Nhiều điều kiện** lọc

**Ví dụ:**

```
GET /users?role=admin&active=true          ✅ Lọc users
GET /products?category=laptop&price_max=30 ✅ Tìm kiếm
GET /posts?sort=date&page=2&limit=20       ✅ Phân trang
```

### 6.3. Kết hợp cả hai

**Tốt nhất:** Kết hợp Path và Query Parameters!

```go
// Lấy orders của user 123, lọc theo status
GET /users/123/orders?status=pending

// Lấy comments của post 456, phân trang
GET /posts/456/comments?page=2&limit=10

// Lấy products của category laptop, sắp xếp
GET /categories/laptop/products?sort=price&order=asc
```

**Ví dụ code:**

```go
r.GET("/users/:userId/orders", func(c *gin.Context) {
	// Path parameter
	userId := c.Param("userId")
	
	// Query parameters
	status := c.Query("status")
	page := c.DefaultQuery("page", "1")
	
	c.JSON(200, gin.H{
		"user_id": userId,
		"status":  status,
		"page":    page,
	})
})
```

**Test:**

```bash
curl "http://localhost:8080/users/123/orders?status=pending&page=2"
```

**Output:**

```json
{
    "user_id": "123",
    "status": "pending",
    "page": "2"
}
```

---

## 7. Best Practices

### 7.1. Quy tắc đặt tên URL

**✅ Nên:**

```
/users              → Danh từ số nhiều
/products
/orders

/users/123          → Rõ ràng
/products/456

/users/123/orders   → Phân cấp logic
```

**❌ Không nên:**

```
/getUsers           → Không dùng động từ
/createProduct
/deleteOrder

/user               → Không dùng số ít
/product

/users_orders       → Không dùng underscore
/product-list
```

### 7.2. RESTful API conventions

```
GET    /users           → Lấy danh sách
POST   /users           → Tạo mới
GET    /users/123       → Lấy chi tiết
PUT    /users/123       → Cập nhật toàn bộ
PATCH  /users/123       → Cập nhật một phần
DELETE /users/123       → Xóa
```

### 7.3. Query Parameters best practices

**✅ Nên:**

```
?category=laptop              → Dùng lowercase
?price_max=20000000           → Dùng underscore
?sort=price&order=asc         → Rõ ràng
?page=2&limit=20              → Chuẩn phân trang
```

**❌ Không nên:**

```
?Category=Laptop              → Uppercase
?priceMax=20000000            → CamelCase
?s=p&o=a                      → Viết tắt khó hiểu
```

### 7.4. Validation

**Luôn validate input:**

```go
r.GET("/users/:id", func(c *gin.Context) {
	idStr := c.Param("id")
	
	// Validate: ID phải là số
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(400, gin.H{
			"error": "Invalid user ID",
		})
		return
	}
	
	// Tiếp tục xử lý...
})
```

### 7.5. Error handling

```go
r.GET("/products/:id", func(c *gin.Context) {
	id := c.Param("id")
	
	// Tìm product
	product := findProduct(id)
	
	if product == nil {
		// 404 Not Found
		c.JSON(404, gin.H{
			"error": "Product not found",
		})
		return
	}
	
	// 200 OK
	c.JSON(200, gin.H{
		"data": product,
	})
})
```

---

## 8. Ví dụ thực tế

### 8.1. E-commerce API hoàn chỉnh

```go
package main

import (
	"strconv"
	"github.com/gin-gonic/gin"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Description string  `json:"description"`
}

var products = []Product{
	{1, "Macbook Pro 16", "laptop", 59990000, 10, "Apple M3 Max"},
	{2, "Dell XPS 15", "laptop", 45990000, 15, "Intel i9, 32GB RAM"},
	{3, "iPhone 15 Pro Max", "phone", 32990000, 20, "1TB, Titan"},
	{4, "Samsung Galaxy S24 Ultra", "phone", 28990000, 25, "12GB RAM"},
	{5, "iPad Pro 12.9", "tablet", 27990000, 12, "M2 chip"},
	{6, "Logitech MX Master 3S", "accessory", 2590000, 50, "Wireless mouse"},
}

func main() {
	r := gin.Default()

	// 1. Lấy tất cả products (có filter)
	r.GET("/products", func(c *gin.Context) {
		category := c.Query("category")
		priceMinStr := c.Query("price_min")
		priceMaxStr := c.Query("price_max")
		sort := c.DefaultQuery("sort", "name")
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

		// Filter
		filtered := []Product{}
		for _, p := range products {
			if category != "" && p.Category != category {
				continue
			}

			if priceMinStr != "" {
				priceMin, _ := strconv.ParseFloat(priceMinStr, 64)
				if p.Price < priceMin {
					continue
				}
			}

			if priceMaxStr != "" {
				priceMax, _ := strconv.ParseFloat(priceMaxStr, 64)
				if p.Price > priceMax {
					continue
				}
			}

			filtered = append(filtered, p)
		}

		// Pagination
		start := (page - 1) * limit
		end := start + limit
		if start > len(filtered) {
			start = len(filtered)
		}
		if end > len(filtered) {
			end = len(filtered)
		}

		c.JSON(200, gin.H{
			"total":    len(filtered),
			"page":     page,
			"limit":    limit,
			"category": category,
			"sort":     sort,
			"data":     filtered[start:end],
		})
	})

	// 2. Lấy product theo ID
	r.GET("/products/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid product ID",
			})
			return
		}

		for _, p := range products {
			if p.ID == id {
				c.JSON(200, gin.H{
					"success": true,
					"data":    p,
				})
				return
			}
		}

		c.JSON(404, gin.H{
			"success": false,
			"error":   "Product not found",
		})
	})

	// 3. Tìm kiếm products
	r.GET("/search", func(c *gin.Context) {
		q := c.Query("q")
		
		if q == "" {
			c.JSON(400, gin.H{
				"error": "Query parameter 'q' is required",
			})
			return
		}

		results := []Product{}
		for _, p := range products {
			// Tìm trong tên hoặc description
			if contains(p.Name, q) || contains(p.Description, q) {
				results = append(results, p)
			}
		}

		c.JSON(200, gin.H{
			"query":   q,
			"total":   len(results),
			"results": results,
		})
	})

	// 4. Lấy products theo category
	r.GET("/categories/:category/products", func(c *gin.Context) {
		category := c.Param("category")
		sort := c.DefaultQuery("sort", "name")

		results := []Product{}
		for _, p := range products {
			if p.Category == category {
				results = append(results, p)
			}
		}

		if len(results) == 0 {
			c.JSON(404, gin.H{
				"error": "No products found in this category",
			})
			return
		}

		c.JSON(200, gin.H{
			"category": category,
			"total":    len(results),
			"sort":     sort,
			"data":     results,
		})
	})

	r.Run(":8080")
}

// Helper function
func contains(s, substr string) bool {
	// Simple case-insensitive search
	return len(s) >= len(substr) && 
		   (s == substr || len(substr) == 0)
	// Trong thực tế nên dùng strings.Contains() hoặc regex
}
```

### 8.2. Test API

```bash
# 1. Tất cả products
curl http://localhost:8080/products

# 2. Chỉ laptop
curl http://localhost:8080/products?category=laptop

# 3. Sản phẩm từ 20-30 triệu
curl "http://localhost:8080/products?price_min=20000000&price_max=30000000"

# 4. Laptop dưới 50 triệu, trang 1, 5 items
curl "http://localhost:8080/products?category=laptop&price_max=50000000&page=1&limit=5"

# 5. Chi tiết product ID = 3
curl http://localhost:8080/products/3

# 6. Tìm kiếm "Pro"
curl "http://localhost:8080/search?q=Pro"

# 7. Products của category phone
curl http://localhost:8080/categories/phone/products
```

---

## 📚 Tổng kết

### URL Components

```
https://api.example.com:8080/products/123?sort=price&page=2#reviews

├─ Protocol: https://
├─ Domain: api.example.com  
├─ Port: :8080
├─ Path: /products/123
├─ Query: ?sort=price&page=2
└─ Fragment: #reviews
```

### Path vs Query Parameters

|Feature|Path Params|Query Params|
|---|---|---|
|Bắt buộc|✅|❌|
|Mục đích|Định danh|Lọc/Tìm kiếm|
|Ví dụ|`/users/:id`|`?role=admin`|

### RESTful Endpoints

```
GET    /resources           → List
POST   /resources           → Create
GET    /resources/:id       → Get one
PUT    /resources/:id       → Update
DELETE /resources/:id       → Delete
```

### Best Practices

1. ✅ Dùng danh từ số nhiều
2. ✅ Lowercase và underscore
3. ✅ Path cho ID, Query cho filter
4. ✅ Validate input
5. ✅ Error handling đầy đủ