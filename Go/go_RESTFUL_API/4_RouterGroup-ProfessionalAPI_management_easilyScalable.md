# Router Group - Quản Lý API Chuyên Nghiệp với Gin Framework

## Mục lục

1. [Tổng quan](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#t%E1%BB%95ng-quan)
2. [Router Group là gì?](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#router-group-l%C3%A0-g%C3%AC)
3. [Tại sao cần Router Group?](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#t%E1%BA%A1i-sao-c%E1%BA%A7n-router-group)
4. [Cấu trúc thư mục dự án](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#c%E1%BA%A5u-tr%C3%BAc-th%C6%B0-m%E1%BB%A5c-d%E1%BB%B1-%C3%A1n)
5. [Ví dụ minh họa: API quản lý User và Product](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#v%C3%AD-d%E1%BB%A5-minh-h%E1%BB%8Da-api-qu%E1%BA%A3n-l%C3%BD-user-v%C3%A0-product)
6. [Tạo Handler cho từng tài nguyên](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#t%E1%BA%A1o-handler-cho-t%E1%BB%ABng-t%C3%A0i-nguy%C3%AAn)
7. [Thiết lập Route trong main.go](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#thi%E1%BA%BFt-l%E1%BA%ADp-route-trong-maingo)
8. [Kiểm thử API bằng cURL](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#ki%E1%BB%83m-th%E1%BB%AD-api-b%E1%BA%B1ng-curl)
9. [Tổng kết](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#t%E1%BB%95ng-k%E1%BA%BFt)

---

## Tổng quan

Trong thực tế, một ứng dụng backend thường có rất nhiều API endpoint. Nếu không tổ chức cẩn thận, code sẽ trở nên khó đọc và khó bảo trì. **Router Group** trong Gin Framework là giải pháp giúp chúng ta gom nhóm các route có cùng prefix lại với nhau, làm cho code gọn gàng và dễ mở rộng hơn.

---

## Router Group là gì?

**Router Group** cho phép bạn tạo ra một nhóm các route có chung một **prefix URL** (phần đầu của đường dẫn). Thay vì phải lặp lại prefix nhiều lần, bạn chỉ cần khai báo một lần cho cả nhóm.

**Ví dụ so sánh:**

❌ **Không dùng Router Group** (lặp đi lặp lại `/api/v1`):

```go
r.GET("/api/v1/users", getUsers)
r.POST("/api/v1/users", createUser)
r.GET("/api/v1/users/:id", getUserByID)
r.PUT("/api/v1/users/:id", updateUser)
r.DELETE("/api/v1/users/:id", deleteUser)
```

✅ **Dùng Router Group** (khai báo prefix một lần):

```go
v1 := r.Group("/api/v1")
{
    users := v1.Group("/users")
    {
        users.GET("/", getUsers)
        users.POST("/", createUser)
        users.GET("/:id", getUserByID)
        users.PUT("/:id", updateUser)
        users.DELETE("/:id", deleteUser)
    }
}
```

> 💡 **Lưu ý:** Cặp ngoặc nhọn `{ }` trong ví dụ trên **không phải là cú pháp bắt buộc** của Go. Đây chỉ là **quy ước trình bày (convention)** giúp code trông rõ ràng hơn, dễ thấy phạm vi của từng group. Về mặt kỹ thuật, bỏ chúng đi code vẫn chạy bình thường.

---

## Tại sao cần Router Group?

|Vấn đề khi không dùng Group|Lợi ích khi dùng Router Group|
|---|---|
|Phải lặp lại prefix URL nhiều lần|Khai báo prefix một lần duy nhất|
|Khó thêm middleware cho một nhóm route|Dễ áp dụng middleware cho cả nhóm|
|Code khó đọc, khó bảo trì|Code rõ ràng, có cấu trúc|
|Khó quản lý versioning (v1, v2, ...)|Dễ dàng quản lý nhiều version API|

---

## Cấu trúc thư mục dự án

Dự án sử dụng kiến trúc **layered (phân tầng)** để tách biệt logic theo từng version API:

```
Router_Group/
├── main.go                          # Entry point của ứng dụng
├── go.mod                           # Khai báo module và dependencies
└── internal/
    └── api/
        ├── v1/
        │   └── handler/
        │       ├── user.go          # Handler quản lý User (version 1)
        │       └── product.go       # Handler quản lý Product (version 1)
        └── v2/
            └── handler/
                └── user.go          # Handler quản lý User (version 2)
```

> 💡 **Tại sao lại tổ chức thư mục như vậy?**
> 
> - Thư mục `internal/` chứa code nội bộ của ứng dụng, không được phép import từ bên ngoài module (đây là quy tắc của Go).
> - Phân chia theo `v1/`, `v2/` giúp dễ dàng nâng cấp API mà không phá vỡ các client đang dùng version cũ.
> - Thư mục `handler/` chứa các hàm xử lý request HTTP.

---

## Ví dụ minh họa: API quản lý User và Product

Chúng ta sẽ xây dựng một API server với **3 nhóm route** sau:

```
# Nhóm 1: Quản lý User - Version 1
GET    /api/v1/users          → Lấy danh sách tất cả users
GET    /api/v1/users/:id      → Lấy thông tin một user theo ID
POST   /api/v1/users          → Tạo user mới
PUT    /api/v1/users/:id      → Cập nhật thông tin user
DELETE /api/v1/users/:id      → Xóa user

# Nhóm 2: Quản lý Product - Version 1
GET    /api/v1/products       → Lấy danh sách tất cả sản phẩm
GET    /api/v1/products/:id   → Lấy thông tin một sản phẩm theo ID
POST   /api/v1/products       → Tạo sản phẩm mới
PUT    /api/v1/products/:id   → Cập nhật thông tin sản phẩm
DELETE /api/v1/products/:id   → Xóa sản phẩm

# Nhóm 3: Quản lý User - Version 2
GET    /api/v2/users          → Lấy danh sách tất cả users (phiên bản mới)
GET    /api/v2/users/:id      → Lấy thông tin một user theo ID (phiên bản mới)
POST   /api/v2/users          → Tạo user mới (phiên bản mới)
PUT    /api/v2/users/:id      → Cập nhật thông tin user (phiên bản mới)
DELETE /api/v2/users/:id      → Xóa user (phiên bản mới)
```

---

## Tạo Handler cho từng tài nguyên

Handler là các hàm chịu trách nhiệm xử lý từng loại request HTTP.

### 1. Handler quản lý User (`internal/api/v1/handler/user.go`)

```go
package handler

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
)

// ----- Struct và dữ liệu mẫu -----

// User định nghĩa cấu trúc dữ liệu của một người dùng
// Các tag `json:"..."` giúp Go biết cách chuyển đổi sang/từ JSON
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

// NewUser là Constructor - hàm tạo ra một instance mới của struct User
// Đây là pattern phổ biến trong Go để khởi tạo đối tượng
func NewUser() *User {
    return &User{}
}

// users là slice lưu trữ dữ liệu trong bộ nhớ (thay thế cho database)
// Trong thực tế, dữ liệu này sẽ được lưu vào database như PostgreSQL, MySQL...
var users []User = []User{
    {ID: 1, Name: "John Doe", Email: "JDoe@gmail.com"},
    {ID: 2, Name: "Jane Smith", Email: "JSmith@gmail.com"},
    {ID: 3, Name: "Bob Johnson", Email: "BJohnson@gmail.com"},
    {ID: 4, Name: "Alice Williams", Email: "AWilliams@gmail.com"},
    {ID: 5, Name: "Charlie Brown", Email: "CBrown@gmail.com"},
    {ID: 6, Name: "Nguyen Van A", Email: "NVAn@gmail.com"},
}

// ----- Các phương thức xử lý request -----

// GetUsers xử lý request GET /api/v1/users
// Trả về danh sách tất cả người dùng
func (obj *User) GetUsers(c *gin.Context) {
    fmt.Println("Lấy danh sách người dùng")
    c.JSON(http.StatusOK, gin.H{
        "message": "list user",
        "data":    users,
    })
}

// GetUserByID xử lý request GET /api/v1/users/:id
// c.Param("id") lấy giá trị của tham số :id từ URL
// Ví dụ: GET /api/v1/users/3 → id = "3"
func (obj *User) GetUserByID(c *gin.Context) {
    id := c.Param("id")
    for _, user := range users {
        if fmt.Sprintf("%v", user.ID) == id {
            c.JSON(http.StatusOK, gin.H{
                "message": "user found",
                "data":    user,
            })
            return // Quan trọng: phải return ngay sau khi tìm thấy để tránh gọi c.JSON nhiều lần
        }
    }

    c.JSON(http.StatusNotFound, gin.H{
        "message": "user not found",
        "data":    nil,
    })
}

// CreateUser xử lý request POST /api/v1/users
// c.ShouldBindJSON đọc body của request và chuyển đổi JSON thành struct User
func (obj *User) CreateUser(c *gin.Context) {
    fmt.Println("Nhập thông tin người dùng")

    var newUser User
    if err := c.ShouldBindJSON(&newUser); err != nil {
        // Nếu body không đúng định dạng JSON hoặc thiếu trường bắt buộc → trả về lỗi 400
        c.JSON(http.StatusBadRequest, gin.H{
            "message": "invalid request body",
        })
        return
    }

    // Tự động tạo ID mới dựa trên số lượng user hiện có
    // Lưu ý: Cách này chỉ phù hợp với dữ liệu in-memory. Với database thực tế nên dùng AUTO_INCREMENT hoặc UUID
    newUser.ID = len(users) + 1

    users = append(users, newUser)

    c.JSON(http.StatusOK, gin.H{
        "message": "create user successfully",
    })
}

// UpdateUser xử lý request PUT /api/v1/users/:id
// Tìm user theo ID rồi cập nhật thông tin
func (obj *User) UpdateUser(c *gin.Context) {
    id := c.Param("id")

    var updateUser User
    if err := c.ShouldBindJSON(&updateUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "message": "invalid request body",
            "error":   err.Error(),
        })
        return
    }

    for idx, user := range users {
        if fmt.Sprintf("%v", user.ID) == id {
            users[idx].Name = updateUser.Name
            users[idx].Email = updateUser.Email

            c.JSON(http.StatusOK, gin.H{
                "message": "update user successfully",
            })
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{
        "message": "user not found",
    })
}

// DeleteUser xử lý request DELETE /api/v1/users/:id
// Kỹ thuật xóa phần tử khỏi slice: nối phần trước và phần sau của phần tử cần xóa
func (obj *User) DeleteUser(c *gin.Context) {
    id := c.Param("id")

    for index, user := range users {
        if fmt.Sprintf("%v", user.ID) == id {
            // Xóa phần tử tại vị trí index:
            // users[:index]    → lấy tất cả phần tử trước vị trí index
            // users[index+1:]  → lấy tất cả phần tử sau vị trí index
            // append(...)      → ghép hai phần lại
            users = append(users[:index], users[index+1:]...)

            c.JSON(http.StatusOK, gin.H{
                "message": "delete user successfully",
            })
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{
        "message": "user not found",
        "data":    nil,
    })
}
```

> ⚠️ **Lưu ý về bug trong GetUserByID:** Code gốc ở hàm `GetUserByID` thiếu lệnh `return` sau khi tìm thấy user và gọi `c.JSON`. Điều này khiến chương trình tiếp tục chạy xuống và gọi thêm `c.JSON` lần thứ hai cho status 404, gây ra lỗi **"superfluous response.WriteHeader call"**. Code đã được sửa ở trên bằng cách thêm `return` vào đúng chỗ.

---

### 2. Handler quản lý Product (`internal/api/v1/handler/product.go`)

Cấu trúc tương tự handler User, nhưng dành cho tài nguyên Product:

```go
package handler

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
)

// ----- Struct và dữ liệu mẫu -----

// Product định nghĩa cấu trúc dữ liệu của một sản phẩm
type Product struct {
    ID          int     `json:"id"`
    ProductName string  `json:"name"`
    Price       float64 `json:"price"`
}

// NewProduct là Constructor - tạo instance mới của struct Product
func NewProduct() *Product {
    return &Product{}
}

// Products là slice lưu trữ dữ liệu sản phẩm trong bộ nhớ
var Products []Product = []Product{
    {ID: 1, ProductName: "Iphone 14 Pro Max", Price: 30000000},
    {ID: 2, ProductName: "Samsung Galaxy S23 Ultra", Price: 25000000},
    {ID: 3, ProductName: "Xiaomi Mi 12 Pro", Price: 20000000},
    {ID: 4, ProductName: "Oppo Find X5 Pro", Price: 22000000},
    {ID: 5, ProductName: "Vivo X80 Pro", Price: 21000000},
    {ID: 6, ProductName: "Realme GT 2 Pro", Price: 18000000},
}

// ----- Các phương thức xử lý request -----

// GetProducts xử lý request GET /api/v1/products
func (obj *Product) GetProducts(c *gin.Context) {
    fmt.Println("Lấy danh sách sản phẩm")
    c.JSON(http.StatusOK, gin.H{
        "message": "list Product",
        "data":    Products,
    })
}

// GetProductByID xử lý request GET /api/v1/products/:id
func (obj *Product) GetProductByID(c *gin.Context) {
    id := c.Param("id")
    for _, product := range Products {
        if fmt.Sprintf("%v", product.ID) == id {
            c.JSON(http.StatusOK, gin.H{
                "message": "Product found",
                "data":    product,
            })
            return // Thêm return để tránh gọi c.JSON nhiều lần
        }
    }

    c.JSON(http.StatusNotFound, gin.H{
        "message": "Product not found",
        "data":    nil,
    })
}

// CreateProduct xử lý request POST /api/v1/products
func (obj *Product) CreateProduct(c *gin.Context) {
    fmt.Println("Nhập thông tin sản phẩm")

    var newProduct Product
    if err := c.ShouldBindJSON(&newProduct); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "message": "invalid request body",
        })
        return
    }

    newProduct.ID = len(Products) + 1
    Products = append(Products, newProduct)

    c.JSON(http.StatusOK, gin.H{
        "message": "create Product successfully",
    })
}

// UpdateProduct xử lý request PUT /api/v1/products/:id
func (obj *Product) UpdateProduct(c *gin.Context) {
    id := c.Param("id")

    var updateProduct Product
    if err := c.ShouldBindJSON(&updateProduct); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "message": "invalid request body",
            "error":   err.Error(),
        })
        return
    }

    for idx, product := range Products {
        if fmt.Sprintf("%v", product.ID) == id {
            Products[idx].ProductName = updateProduct.ProductName
            Products[idx].Price = updateProduct.Price

            c.JSON(http.StatusOK, gin.H{
                "message": "update Product successfully",
            })
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{
        "message": "Product not found",
    })
}

// DeleteProduct xử lý request DELETE /api/v1/products/:id
func (obj *Product) DeleteProduct(c *gin.Context) {
    id := c.Param("id")

    for index, product := range Products {
        if fmt.Sprintf("%v", product.ID) == id {
            Products = append(Products[:index], Products[index+1:]...)
            c.JSON(http.StatusOK, gin.H{
                "message": "delete Product successfully",
            })
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{
        "message": "Product not found",
        "data":    nil,
    })
}
```

---

### 3. Handler User Version 2 (`internal/api/v2/handler/user.go`)

File này có cấu trúc tương tự `v1/handler/user.go`. Trong thực tế, version 2 thường có những thay đổi như thêm trường mới, thay đổi logic xử lý, hoặc cải thiện hiệu suất so với version 1.

> 💡 **Tại sao v1 và v2 có cùng package name `handler`?** Mỗi thư mục trong Go là một package riêng biệt. Dù cả hai đều đặt tên là `handler`, chúng vẫn là hai package **khác nhau** vì nằm ở hai đường dẫn khác nhau. Khi import vào `main.go`, ta dùng **alias** (tên gọi tắt) để phân biệt:
> 
> ```go
> import (
>     v1Handler "router-group/internal/api/v1/handler"
>     v2Handler "router-group/internal/api/v2/handler"
> )
> ```

---

## Thiết lập Route trong `main.go`

Đây là nơi kết nối tất cả các handler lại và định nghĩa cấu trúc URL của toàn bộ API:

```go
package main

import (
    v1Handler "router-group/internal/api/v1/handler"
    v2Handler "router-group/internal/api/v2/handler"

    "github.com/gin-gonic/gin"
)

func main() {
    // Khởi tạo Gin router với các middleware mặc định (Logger và Recovery)
    r := gin.Default()

    // ---- Khởi tạo các Handler ----
    // Sử dụng Constructor để tạo instance - đây là cách khởi tạo đối tượng trong Go
    userHandler_v1 := v1Handler.NewUser()
    productHandler_v1 := v1Handler.NewProduct()
    userHandler_v2 := v2Handler.NewUser()

    // ---- Định nghĩa Router Group cho API v1 ----
    // Tất cả các route trong block này đều có prefix /api/v1
    v1 := r.Group("/api/v1")
    {
        // Nhóm route cho User v1: prefix = /api/v1/users
        user := v1.Group("/users")
        {
            user.GET("/", userHandler_v1.GetUsers)           // GET  /api/v1/users
            user.GET("/:id", userHandler_v1.GetUserByID)     // GET  /api/v1/users/:id
            user.POST("/", userHandler_v1.CreateUser)        // POST /api/v1/users
            user.PUT("/:id", userHandler_v1.UpdateUser)      // PUT  /api/v1/users/:id
            user.DELETE("/:id", userHandler_v1.DeleteUser)   // DELETE /api/v1/users/:id
        }

        // Nhóm route cho Product v1: prefix = /api/v1/products
        product := v1.Group("/products")
        {
            product.GET("/", productHandler_v1.GetProducts)
            product.GET("/:id", productHandler_v1.GetProductByID)
            product.POST("/", productHandler_v1.CreateProduct)
            product.PUT("/:id", productHandler_v1.UpdateProduct)
            product.DELETE("/:id", productHandler_v1.DeleteProduct)
        }
    }

    // ---- Định nghĩa Router Group cho API v2 ----
    // Tất cả các route trong block này đều có prefix /api/v2
    v2 := r.Group("/api/v2")
    {
        // Nhóm route cho User v2: prefix = /api/v2/users
        user := v2.Group("/users")
        {
            user.GET("/", userHandler_v2.GetUsers)
            user.GET("/:id", userHandler_v2.GetUserByID)
            user.POST("/", userHandler_v2.CreateUser)
            user.PUT("/:id", userHandler_v2.UpdateUser)
            user.DELETE("/:id", userHandler_v2.DeleteUser)
        }
    }

    // Khởi động server và lắng nghe trên cổng 8080
    r.Run(":8080")
}
```

### Sơ đồ cấu trúc Route

```
r (gin.Engine)
├── /api/v1  (group v1)
│   ├── /users  (group user)
│   │   ├── GET    /          → GetUsers
│   │   ├── GET    /:id       → GetUserByID
│   │   ├── POST   /          → CreateUser
│   │   ├── PUT    /:id       → UpdateUser
│   │   └── DELETE /:id       → DeleteUser
│   │
│   └── /products  (group product)
│       ├── GET    /          → GetProducts
│       ├── GET    /:id       → GetProductByID
│       ├── POST   /          → CreateProduct
│       ├── PUT    /:id       → UpdateProduct
│       └── DELETE /:id       → DeleteProduct
│
└── /api/v2  (group v2)
    └── /users  (group user)
        ├── GET    /          → GetUsers (v2)
        ├── GET    /:id       → GetUserByID (v2)
        ├── POST   /          → CreateUser (v2)
        ├── PUT    /:id       → UpdateUser (v2)
        └── DELETE /:id       → DeleteUser (v2)
```

---

## Kiểm thử API bằng cURL

Sau khi chạy server (`go run main.go`), bạn có thể kiểm thử API bằng lệnh cURL trong terminal.

### Lấy danh sách users

```bash
curl -X GET http://localhost:8080/api/v1/users/
```

### Lấy user theo ID

```bash
curl -X GET http://localhost:8080/api/v1/users/1
```

### Tạo user mới

```bash
curl -X POST http://localhost:8080/api/v1/users/ \
  -H "Content-Type: application/json" \
  -d '{"name": "Tran Van B", "email": "TVB@gmail.com"}'
```

### Cập nhật user

```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe Updated", "email": "JDoeNew@gmail.com"}'
```

### Xóa user

```bash
curl -X DELETE http://localhost:8080/api/v1/users/1
```

### Lấy danh sách products (v1)

```bash
curl -X GET http://localhost:8080/api/v1/products/
```

### Lấy danh sách users (v2)

```bash
curl -X GET http://localhost:8080/api/v2/users/
```

---

## Tổng kết

Qua bài học này, chúng ta đã nắm được:

|Khái niệm|Ý nghĩa|
|---|---|
|**Router Group**|Gom nhóm các route có cùng prefix URL để code gọn hơn|
|**API Versioning**|Phân chia API theo version (v1, v2) để dễ nâng cấp mà không phá vỡ client cũ|
|**Handler**|Hàm xử lý từng loại HTTP request (GET, POST, PUT, DELETE)|
|**Constructor (`NewXxx()`)**|Hàm khởi tạo instance của struct trong Go|
|**`c.Param()`**|Lấy tham số từ URL (ví dụ: `:id`)|
|**`c.ShouldBindJSON()`**|Đọc và parse JSON từ request body|
|**`c.JSON()`**|Trả về response dạng JSON cho client|
|**Package alias**|Dùng alias khi import nhiều package có cùng tên (`v1Handler`, `v2Handler`)|

### Các điểm cần cải thiện trong thực tế

1. **Dùng database thực sự** thay cho slice in-memory (PostgreSQL, MySQL, MongoDB...)
2. **Thêm middleware** cho xác thực (Authentication) và phân quyền (Authorization)
3. **Validation chặt chẽ hơn** cho dữ liệu đầu vào (dùng thư viện `validator`)
4. **Xử lý lỗi nhất quán** bằng cách tạo một error handler chung
5. **Thêm logging** để theo dõi và debug lỗi
6. **Viết unit test** cho từng handler