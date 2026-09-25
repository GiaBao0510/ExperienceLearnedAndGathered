# Validator Package - Xác Thực Dữ Liệu Chuyên Nghiệp với `go-playground/validator`

## Mục lục

1. [Tổng quan](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#t%E1%BB%95ng-quan)
2. [Cài đặt thư viện](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#c%C3%A0i-%C4%91%E1%BA%B7t-th%C6%B0-vi%E1%BB%87n)
3. [Validate Path Param với `ShouldBindUri`](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#validate-path-param-v%E1%BB%9Bi-shouldbinduri)
4. [Xây dựng `HandleValidationErrors` — Xử lý lỗi tập trung](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#x%C3%A2y-d%E1%BB%B1ng-handlevalidationerrors--x%E1%BB%AD-l%C3%BD-l%E1%BB%97i-t%E1%BA%ADp-trung)
5. [Validate UUID với tag `binding:"uuid"`](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#validate-uuid-v%E1%BB%9Bi-tag-bindinguuid)
6. [Custom Validation — Tự định nghĩa rule riêng](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#custom-validation--t%E1%BB%B1-%C4%91%E1%BB%8Bnh-ngh%C4%A9a-rule-ri%C3%AAng)
7. [Validate Query Param với `ShouldBindQuery`](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#validate-query-param-v%E1%BB%9Bi-shouldbindquery)
8. [Toàn bộ file `utils/validation.go`](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#to%C3%A0n-b%E1%BB%99-file-utilsvalidationgo)
9. [Tổng kết các tag validation](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#t%E1%BB%95ng-k%E1%BA%BFt-c%C3%A1c-tag-validation)

---

## Tổng quan

Ở các bài trước, ta tự viết hàm validation trong `utils/validation.go`. Cách đó hoạt động tốt nhưng vẫn phải viết nhiều code boilerplate. Bài này giới thiệu thư viện **`go-playground/validator`** — thư viện được Gin tích hợp sẵn — cho phép khai báo rule validation ngay trong **struct tag**, giảm đáng kể lượng code cần viết.

**So sánh hai cách tiếp cận:**

```
Cách cũ (bài trước):                   Cách mới (bài này):
──────────────────────────────          ──────────────────────────────────────
id := c.Param("id")                     type Param struct {
err := utils.ValidationPositiveInt(        ID int `uri:"id" binding:"gt=0"`
    "id", id)                           }
if err != nil { ... }                   ctx.ShouldBindUri(&params)
```

Thay vì viết logic kiểm tra từng bước, ta **khai báo điều kiện trong struct tag** rồi để Gin + validator tự xử lý.

---

## Cài đặt thư viện

```bash
go get github.com/go-playground/validator/v10
```

> 💡 **Gin đã dùng validator bên trong.** Khi gọi `ShouldBindUri`, `ShouldBindQuery`, `ShouldBindJSON`... Gin tự động sử dụng `go-playground/validator` để kiểm tra các tag `binding:"..."`. Ta cài thêm package này để có thể **đọc và xử lý chi tiết lỗi** từ validator (dùng kiểu `validator.ValidationErrors`).

---

## Validate Path Param với `ShouldBindUri`

### Cơ chế hoạt động

Thay vì gọi `c.Param("id")` rồi tự kiểm tra, ta tạo một struct với tag `uri` và `binding`, rồi dùng `ctx.ShouldBindUri()`:

```
Tag uri:"id"       → ánh xạ tham số :id từ URL vào trường ID
Tag binding:"gt=0" → điều kiện: ID phải > 0
```

Nếu `:id` trong URL không thể chuyển thành `int`, hoặc giá trị `≤ 0`, Gin sẽ tự trả về lỗi mà không cần ta tự kiểm tra.

### Tạo struct và handler trong `internal/api/v2/handler/user.go`

```go
package handler

import (
    "net/http"
    "router-group/utils"

    "github.com/gin-gonic/gin"
)

// GetUserByID_V2_Param là struct mô tả điều kiện của tham số :id
// Tag uri:"id"    → tên tham số trong route (khớp với /:id trong main.go)
// Tag binding:"gt=0" → điều kiện: phải là số nguyên và phải > 0
// Lưu ý: kiểu của ID phải là int (không phải string).
// Gin sẽ tự chuyển chuỗi từ URL sang int. Nếu không chuyển được → lỗi tự động.
type GetUserByID_V2_Param struct {
    ID int `uri:"id" binding:"gt=0"`
}

// GetUserByID xử lý GET /api/v2/users/:id
func (obj *User) GetUserByID(ctx *gin.Context) {

    var params GetUserByID_V2_Param

    // ShouldBindUri đồng thời làm 2 việc:
    // 1. Lấy giá trị :id từ URL và gán vào params.ID (có chuyển kiểu string → int)
    // 2. Kiểm tra điều kiện binding:"gt=0"
    // Nếu có lỗi ở bước nào → trả về error ngay
    if err := ctx.ShouldBindUri(&params); err != nil {
        ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "message": "Get user by ID successfully",
        "user_id": params.ID,
    })
}
```

### Kết quả kiểm thử

```bash
# Hợp lệ
GET /api/v2/users/10
→ { "message": "Get user by ID successfully", "user_id": 10 }

# Không hợp lệ: giá trị âm (vi phạm gt=0)
GET /api/v2/users/-10
→ { "error": "Key: 'GetUserByID_V2_Param.ID' Error:Field validation for 'ID' failed on the 'gt' tag" }

# Không hợp lệ: không phải số (không thể chuyển sang int)
GET /api/v2/users/abc
→ { "error": "strconv.ParseInt: parsing \"abc\": invalid syntax" }
```

> 💡 **Tại sao lỗi ở hai trường hợp cuối trông khác nhau?**
> 
> - Giá trị `-10` chuyển sang `int` thành công (= -10), nhưng **vi phạm rule `gt=0`** → đây là lỗi từ `validator.ValidationErrors`.
> - Giá trị `"abc"` **không thể chuyển sang `int`** → đây là lỗi kiểu dữ liệu (parse error), không phải lỗi validation → thông báo lỗi có dạng khác.
> 
> Hàm `HandleValidationErrors` sẽ xử lý cả hai trường hợp này — xem phần tiếp theo.

---

## Xây dựng `HandleValidationErrors` — Xử lý lỗi tập trung

### Vấn đề với thông báo lỗi gốc

Lỗi trả về từ validator mặc định rất kỹ thuật, khó hiểu với người dùng:

```
"Key: 'GetUserByID_V2_Param.ID' Error:Field validation for 'ID' failed on the 'gt' tag"
```

Ta cần một hàm chuyển lỗi kỹ thuật thành thông báo **thân thiện, có ý nghĩa**.

### Thêm vào `utils/validation.go`

```go
import (
    // ... các import cũ ...
    "log"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
)

// HandleValidationErrors chuyển đổi lỗi từ validator thành response JSON thân thiện.
//
// Hàm nhận vào error và kiểm tra xem đó có phải lỗi từ validator không:
//   - Nếu đúng (validator.ValidationErrors): duyệt từng lỗi, tạo message theo tag
//   - Nếu không (parse error, type error...): trả về thông báo lỗi chung
//
// Trả về gin.H (map[string]interface{}) để dùng trực tiếp trong ctx.JSON()
func HandleValidationErrors(Err error) gin.H {

    // Type assertion: kiểm tra Err có phải kiểu validator.ValidationErrors không
    // validator.ValidationErrors là một slice — có thể có nhiều lỗi cùng lúc
    // (ví dụ: vừa vi phạm min vừa vi phạm regex)
    if validationErr, ok := Err.(validator.ValidationErrors); ok {

        errors := make(map[string]string) // key = tên field, value = thông báo lỗi

        for _, e := range validationErr {
            // e.Tag()   → tên rule bị vi phạm (gt, uuid, min, max, ...)
            // e.Field() → tên trường bị lỗi (ID, UUID, Search, ...)
            // e.Param() → tham số của rule (ví dụ: min=3 → e.Param() = "3")
            switch e.Tag() {
            case "gt":
                errors[e.Field()] = e.Field() + " phải lớn hơn giá trị tối thiểu."
            case "gte":
                errors[e.Field()] = e.Field() + " phải lớn hơn hoặc bằng " + e.Param() + "."
            case "lte":
                errors[e.Field()] = e.Field() + " phải nhỏ hơn hoặc bằng " + e.Param() + "."
            case "uuid":
                errors[e.Field()] = e.Field() + " phải là một UUID hợp lệ."
            case "slug":
                errors[e.Field()] = e.Field() + " chỉ có thể chứa: chữ thường, số, dấu gạch ngang (-) hoặc dấu chấm (.)"
            case "max":
                errors[e.Field()] = e.Field() + " độ dài tối đa là " + e.Param() + " ký tự."
            case "min":
                errors[e.Field()] = e.Field() + " độ dài tối thiểu là " + e.Param() + " ký tự."
            case "oneof":
                // e.Param() trả về các giá trị cách nhau bằng dấu cách: "php python golang"
                // strings.Split rồi Join để đổi thành "php,python,golang"
                allowedValues := strings.Join(strings.Split(e.Param(), " "), ", ")
                errors[e.Field()] = e.Field() + " phải là một trong các giá trị: " + allowedValues + "."
            case "required":
                errors[e.Field()] = e.Field() + " là bắt buộc."
            case "search":
                errors[e.Field()] = e.Field() + " chỉ được phép chứa chữ cái, số và khoảng trắng."
            case "email":
                errors[e.Field()] = e.Field() + " phải là một địa chỉ email hợp lệ."
            case "datetime":
                errors[e.Field()] = e.Field() + " phải có định dạng ngày tháng hợp lệ (YYYY-MM-DD)."
            }

            // Ghi log để debug (chỉ in ra server console, không gửi cho client)
            log.Printf("Validation error — field: '%s', tag: '%s', error: %s",
                e.Field(), e.Tag(), e.Error())
        }

        return gin.H{"error": errors}
    }

    // Lỗi không phải từ validator (ví dụ: không parse được int từ string)
    // → trả về thông báo lỗi chung kèm chi tiết gốc
    return gin.H{"error": "Yêu cầu không hợp lệ: " + Err.Error()}
}
```

> 💡 **`e.Param()` là gì?** Khi tag có tham số như `min=3`, `max=100`, `lte=50`... thì `e.Param()` trả về phần số đó (`"3"`, `"100"`, `"50"`). Nhờ vậy ta có thể tạo thông báo lỗi động: _"Độ dài tối đa là 100 ký tự"_ thay vì hard-code con số.

---

## Validate UUID với tag `binding:"uuid"`

Validator có sẵn tag `uuid` kiểm tra định dạng UUID chuẩn — không cần viết regex thủ công.

### Thêm vào `internal/api/v2/handler/user.go`

```go
// GetUserByUUID_Param mô tả điều kiện của tham số :uuid
// Tag binding:"uuid" → validator tự kiểm tra định dạng UUID chuẩn
//
// ⚠️ Lưu ý quan trọng: trường phải là chữ HOA (UUID, không phải uuid)
// để exported và ShouldBindUri có thể gán giá trị vào.
// Trường chữ thường (unexported) sẽ bị bỏ qua khi bind → UUID luôn rỗng.
type GetUserByUUID_Param struct {
    UUID string `uri:"uuid" binding:"uuid"`
}

// GetUserBy_UUID xử lý GET /api/v2/users/uuid/:uuid
func (obj *User) GetUserBy_UUID(ctx *gin.Context) {

    var params GetUserByUUID_Param
    if err := ctx.ShouldBindUri(&params); err != nil {
        ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "message": "Get user by UUID successfully",
        "UUID":    params.UUID,
    })
}
```

> ⚠️ **Bug trong code gốc — Unexported field:** Struct gốc khai báo `uuid string` (chữ thường). Trong Go, trường **chữ thường là unexported** — các thư viện bên ngoài (kể cả Gin's binding) **không thể gán giá trị** vào trường này. Kết quả: `params.uuid` luôn là `""` dù URL có UUID hợp lệ. Phải đổi thành `UUID string` (chữ hoa).

### Cập nhật `main.go` cho route UUID

```go
user := v2.Group("/users")
{
    user.GET("/", userHandler_v2.GetUsers)
    user.GET("/:id", userHandler_v2.GetUserByID)
    user.GET("/uuid/:uuid", userHandler_v2.GetUserBy_UUID) // Route riêng để tránh xung đột với /:id
    user.POST("/", userHandler_v2.CreateUser)
    user.PUT("/:id", userHandler_v2.UpdateUser)
    user.DELETE("/:id", userHandler_v2.DeleteUser)
}
```

> 💡 **Tại sao dùng `/uuid/:uuid` thay vì `/:uuid`?** Nếu dùng `/:uuid`, Gin sẽ không biết phân biệt route này với `/:id` — cả hai đều nhận một tham số trong path. Thêm prefix `/uuid/` làm rõ đây là route khác biệt, tránh conflict.

### Kết quả kiểm thử

```bash
# Không hợp lệ: chuỗi bất kỳ
GET /api/v2/users/uuid/test
→ { "error": { "UUID": "UUID phải là một UUID hợp lệ." } }

# Không hợp lệ: số không đủ định dạng UUID
GET /api/v2/users/uuid/123
→ { "error": { "UUID": "UUID phải là một UUID hợp lệ." } }

# Hợp lệ
GET /api/v2/users/uuid/9a463ff3-17a3-436b-9a7a-3f864282f9fc
→ { "message": "Get user by UUID successfully", "UUID": "9a463ff3-17a3-436b-9a7a-3f864282f9fc" }
```

---

## Custom Validation — Tự định nghĩa rule riêng

### Khi nào cần Custom Validation?

Package validator có sẵn hàng chục tag như `uuid`, `email`, `min`, `max`, `oneof`... nhưng không thể bao phủ mọi trường hợp đặc thù. Khi cần rule riêng (như `slug`, hay format nghiệp vụ đặc biệt), ta tự đăng ký tag mới.

### Cơ chế đăng ký Custom Validation

```
Gin sử dụng validator engine bên trong.
Ta lấy engine đó ra, đăng ký rule mới, rồi Gin dùng lại engine đã có rule mới.

binding.Validator.Engine() → lấy *validator.Validate
v.RegisterValidation("tag-name", func) → đăng ký rule
```

### Thêm `RegisterValidationError` vào `utils/validation.go`

```go
import (
    // ...
    "regexp"
    "github.com/gin-gonic/gin/binding"
)

// RegisterValidationError đăng ký các custom validation rule vào Gin's validator engine.
// Phải được gọi một lần duy nhất khi khởi động server (trong main.go).
//
// Hàm trả về error nếu không lấy được validator engine (rất hiếm xảy ra).
func RegisterValidationError() error {

    // Lấy validator engine mà Gin đang dùng
    // binding.Validator.Engine() trả về interface{}, cần type assertion sang *validator.Validate
    v, ok := binding.Validator.Engine().(*validator.Validate)
    if !ok {
        return fmt.Errorf("failed to register validation: could not get validator engine")
    }

    // --- Đăng ký rule "slug" ---
    // Rule: chỉ chứa chữ thường, số, dấu - hoặc .
    // fl.Field().String() lấy giá trị của trường dưới dạng string
    slugRegex := regexp.MustCompile(`^[a-z0-9]+(?:[-.][a-z0-9]+)*$`)
    v.RegisterValidation("slug", func(fl validator.FieldLevel) bool {
        return slugRegex.MatchString(fl.Field().String())
    })

    // --- Đăng ký rule "search" ---
    // Rule: chỉ chứa chữ cái (hoa/thường), số và khoảng trắng
    searchRegex := regexp.MustCompile(`^[a-zA-Z0-9\s]+$`)
    v.RegisterValidation("search", func(fl validator.FieldLevel) bool {
        return searchRegex.MatchString(fl.Field().String())
    })

    return nil
}
```

> 💡 **`fl.Field().String()` là gì?** `fl` (FieldLevel) là interface cung cấp thông tin về trường đang được validate. `fl.Field()` trả về `reflect.Value`, `.String()` lấy giá trị thực dưới dạng string. Đây là cách validator truy cập giá trị trường một cách generic, không phụ thuộc kiểu cụ thể.

### Gọi trong `main.go` — đăng ký khi khởi động

```go
func main() {
    r := gin.Default()

    // Đăng ký custom validation TRƯỚC khi khởi tạo handler và route
    // Nếu đăng ký thất bại → panic ngay để không chạy server với validation thiếu
    if err := utils.RegisterValidationError(); err != nil {
        panic(err)
    }

    // Khởi tạo handler...
    // Khai báo route...

    r.Run(":8080")
}
```

> 💡 **Tại sao dùng `panic` thay vì `log.Fatal`?** Cả hai đều dừng chương trình. `panic` thêm stack trace giúp debug dễ hơn. Việc đăng ký validation thất bại là lỗi lập trình (không phải lỗi runtime), nên dừng hard là đúng đắn.

### Tạo `internal/api/v2/handler/product.go` với Slug validation

```go
package handler

import (
    "net/http"
    "router-group/utils"

    "github.com/gin-gonic/gin"
)

type Product struct {
    ID          int     `json:"id"`
    ProductName string  `json:"name"`
    Price       float64 `json:"price"`
}

func NewProduct() *Product {
    return &Product{}
}

// GetProductBySlug_Param mô tả điều kiện của tham số :slug
// binding:"slug,min=5,max=100":
//   - "slug" → custom rule đã đăng ký (chỉ chứa chữ thường, số, - và .)
//   - "min=5" → ít nhất 5 ký tự
//   - "max=100" → tối đa 100 ký tự
// Nhiều rule ngăn cách bằng dấu phẩy, validator kiểm tra theo thứ tự từ trái qua phải
//
// ⚠️ Trường phải viết hoa (Slug) để binding hoạt động
type GetProductBySlug_Param struct {
    Slug string `uri:"slug" binding:"slug,min=5,max=100"`
}

// GetProductBySlug xử lý GET /api/v2/products/:slug
func (obj *Product) GetProductBySlug(ctx *gin.Context) {

    var params GetProductBySlug_Param
    if err := ctx.ShouldBindUri(&params); err != nil {
        ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
        return
    }

    // Trong thực tế: tìm sản phẩm theo slug trong database
    ctx.JSON(http.StatusNotFound, gin.H{
        "message": "Product not found",
        "slug":    params.Slug,
    })
}
```

### Tạo `internal/api/v2/handler/category.go` với `oneof` validation

```go
package handler

import (
    "net/http"
    "router-group/utils"

    "github.com/gin-gonic/gin"
)

type Category struct{}

func NewCategoryHandler() *Category {
    return &Category{}
}

// GetCategoryByCategoriest_Param mô tả điều kiện của tham số :category
// binding:"oneof=php python golang java" → chỉ chấp nhận một trong 4 giá trị này
// Đây là tag có sẵn trong validator, không cần custom
//
// ⚠️ Trường uri:"" bị trống trong code gốc — phải điền đúng tên tham số: uri:"category"
type GetCategoryByCategoriest_Param struct {
    Category string `uri:"category" binding:"oneof=php python golang java"`
}

// GetCategoryByCategories xử lý GET /api/v2/categories/:category
func (obj *Category) GetCategoryByCategories(c *gin.Context) {

    var params GetCategoryByCategoriest_Param
    if err := c.ShouldBindUri(&params); err != nil {
        c.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "success",
        "data":    params.Category,
    })
}
```

> ⚠️ **Bug trong code gốc:** `uri:""` bị để trống — không ánh xạ được tham số `:category` từ URL. Phải sửa thành `uri:"category"`.

---

## Validate Query Param với `ShouldBindQuery`

### `ShouldBindUri` vs `ShouldBindQuery` — Khác nhau ở đâu?

||`ShouldBindUri`|`ShouldBindQuery`|
|---|---|---|
|**Nguồn dữ liệu**|Path param (`/:id`)|Query string (`?search=...`)|
|**Tag trong struct**|`uri:"key"`|`form:"key"`|
|**Bắt buộc mặc định**|Có (là phần của URL)|Không (trừ khi có `binding:"required"`)|
|**Ví dụ**|`/users/10`|`/products?search=iphone`|

> 💡 **Tại sao Query dùng tag `form` chứ không phải `query`?** Gin xử lý query string và form data (HTML form POST) theo cùng một cơ chế — cả hai đều dùng `ShouldBindQuery`/`ShouldBind` với tag `form`. Đây là quy ước của thư viện, không phải lỗi.

### Tag `omitempty` — Trường tùy chọn

```go
Limit int `form:"limit" binding:"omitempty,gte=1,lte=100"`
```

- **Không có `omitempty`**: trường bắt buộc phải có trong request.
- **Có `omitempty`**: nếu trường **vắng mặt hoặc bằng zero value** (`0`, `""`, `false`...) → **bỏ qua** tất cả rule phía sau, không báo lỗi.
- Ví dụ trên: `limit` là tùy chọn, nhưng **nếu có** thì phải từ 1 đến 100.

### Tạo `SearchProducts` trong `internal/api/v2/handler/product.go`

```go
// SearchProducts_Param mô tả tất cả query param cho tính năng tìm kiếm
type SearchProducts_Param struct {
    // Search: bắt buộc, custom rule "search", độ dài 3-100 ký tự
    Search string `form:"search" binding:"required,search,min=3,max=100"`

    // Limit: tùy chọn (omitempty), nếu có phải từ 1 đến 100
    Limit int `form:"limit" binding:"omitempty,gte=1,lte=100"`

    // Email: bắt buộc, phải đúng định dạng email (tag có sẵn trong validator)
    Email string `form:"email" binding:"email"`

    // Date: tùy chọn, nếu có phải đúng định dạng YYYY-MM-DD
    // datetime=2006-01-02 → định dạng theo chuẩn Go (năm 2006, tháng 01, ngày 02)
    Date string `form:"date" binding:"omitempty,datetime=2006-01-02"`
}

// SearchProducts xử lý GET /api/v2/products/search?search=...&limit=...
func (obj *Product) SearchProducts(ctx *gin.Context) {

    var params SearchProducts_Param

    // ShouldBindQuery đọc từ query string (?key=value)
    // Không cần khai báo tham số trong route
    if err := ctx.ShouldBindQuery(&params); err != nil {
        ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
        return
    }

    // Xử lý giá trị mặc định cho các trường tùy chọn
    if params.Limit == 0 {
        params.Limit = 10 // Mặc định lấy 10 kết quả
    }
    if params.Email == "" {
        params.Email = "No email provided"
    }

    ctx.JSON(http.StatusOK, gin.H{
        "message": "Successfully search Product",
        "search":  params.Search,
        "limit":   params.Limit,
        "email":   params.Email,
        "date":    params.Date,
    })
}
```

> 💡 **Định dạng `datetime=2006-01-02` có gì đặc biệt?** Go dùng một ngày cụ thể làm "reference time": `Mon Jan 2 15:04:05 MST 2006`. Thay vì dùng ký tự đại diện như `YYYY`, `MM`, `DD` (kiểu Python/Java), Go dùng các số cụ thể: năm = `2006`, tháng = `01`, ngày = `02`, giờ = `15`, phút = `04`, giây = `05`. Nên `datetime=2006-01-02` nghĩa là "định dạng YYYY-MM-DD".

### Cập nhật route trong `main.go`

```go
product := v2.Group("/products")
{
    // ⚠️ Route /search phải đặt TRƯỚC /:slug
    // Nếu đặt sau, Gin sẽ nhầm "search" là giá trị của :slug
    product.GET("/search", productHandler_v2.SearchProducts)
    product.GET("/", productHandler_v2.GetProducts)
    product.GET("/:slug", productHandler_v2.GetProductBySlug)
    product.POST("/", productHandler_v2.CreateProduct)
    product.PUT("/:id", productHandler_v2.UpdateProduct)
    product.DELETE("/:id", productHandler_v2.DeleteProduct)
}
```

### Kiểm thử SearchProducts

```bash
# Hợp lệ: chỉ có search
GET /api/v2/products/search?search=iphone
→ { "search": "iphone", "limit": 10, "email": "No email provided", "date": "" }

# Hợp lệ: đầy đủ tham số
GET /api/v2/products/search?search=samsung&limit=5&email=user@gmail.com&date=2024-01-15
→ { "search": "samsung", "limit": 5, "email": "user@gmail.com", "date": "2024-01-15" }

# Lỗi: search rỗng
GET /api/v2/products/search
→ { "error": { "Search": "Search là bắt buộc." } }

# Lỗi: search có ký tự đặc biệt (vi phạm custom rule "search")
GET /api/v2/products/search?search=iphone!!!
→ { "error": { "Search": "Search chỉ được phép chứa chữ cái, số và khoảng trắng." } }

# Lỗi: limit vượt quá 100
GET /api/v2/products/search?search=iphone&limit=200
→ { "error": { "Limit": "Limit phải nhỏ hơn hoặc bằng 100." } }

# Lỗi: ngày sai định dạng
GET /api/v2/products/search?search=iphone&date=15-01-2024
→ { "error": { "Date": "Date phải có định dạng ngày tháng hợp lệ (YYYY-MM-DD)." } }
```

---

## Toàn bộ file `utils/validation.go`

```go
package utils

import (
    "fmt"
    "log"
    "regexp"
    "strconv"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/gin/binding"
    "github.com/go-playground/validator/v10"
    "github.com/google/uuid"
)

func ValidationRequired(fieldname, value string) error {
    if value == "" {
        return fmt.Errorf("The %s field must not be blank.", fieldname)
    }
    return nil
}

func ValidationLength(fieldname, value string, min, max int) error {
    if len(value) < min || len(value) > max {
        return fmt.Errorf("The %s field must be between %d and %d characters long.", fieldname, min, max)
    }
    return nil
}

func ValidationRegex(fieldname, value string, reg *regexp.Regexp) error {
    if !reg.MatchString(value) {
        return fmt.Errorf("The %s field is not in the correct format (%s).", fieldname, reg.String())
    }
    return nil
}

func ValidationPositiveInt(fieldname, value string) error {
    v, err := strconv.Atoi(value)
    if err != nil || v <= 0 {
        return fmt.Errorf("The requirement is that %s must be a positive integer.", fieldname)
    }
    return nil
}

func ValidationUUID(fieldname, value string) (uuid.UUID, error) {
    uid, err := uuid.Parse(value)
    if err != nil {
        return uuid.Nil, fmt.Errorf("Field %s: '%s' is not a valid UUID.", fieldname, value)
    }
    return uid, nil
}

func ValidationInList(fieldName, value string, list map[string]bool) error {
    if !list[value] {
        return fmt.Errorf(
            "The value '%s' of field '%s' does not exist in the list.\nValid values: %v",
            value, fieldName, GetKeysFromMap(list),
        )
    }
    return nil
}

func GetKeysFromMap(m map[string]bool) []string {
    keys := make([]string, 0, len(m))
    for key := range m {
        keys = append(keys, key)
    }
    return keys
}

// HandleValidationErrors — chuyển lỗi validator thành JSON thân thiện
func HandleValidationErrors(Err error) gin.H {
    if validationErr, ok := Err.(validator.ValidationErrors); ok {
        errors := make(map[string]string)
        for _, e := range validationErr {
            switch e.Tag() {
            case "gt":
                errors[e.Field()] = e.Field() + " phải lớn hơn giá trị tối thiểu."
            case "gte":
                errors[e.Field()] = e.Field() + " phải lớn hơn hoặc bằng " + e.Param() + "."
            case "lte":
                errors[e.Field()] = e.Field() + " phải nhỏ hơn hoặc bằng " + e.Param() + "."
            case "uuid":
                errors[e.Field()] = e.Field() + " phải là một UUID hợp lệ."
            case "slug":
                errors[e.Field()] = e.Field() + " chỉ có thể chứa: chữ thường, số, dấu gạch ngang (-) hoặc dấu chấm (.)"
            case "max":
                errors[e.Field()] = e.Field() + " độ dài tối đa là " + e.Param() + " ký tự."
            case "min":
                errors[e.Field()] = e.Field() + " độ dài tối thiểu là " + e.Param() + " ký tự."
            case "oneof":
                allowedValues := strings.Join(strings.Split(e.Param(), " "), ", ")
                errors[e.Field()] = e.Field() + " phải là một trong các giá trị: " + allowedValues + "."
            case "required":
                errors[e.Field()] = e.Field() + " là bắt buộc."
            case "search":
                errors[e.Field()] = e.Field() + " chỉ được phép chứa chữ cái, số và khoảng trắng."
            case "email":
                errors[e.Field()] = e.Field() + " phải là một địa chỉ email hợp lệ."
            case "datetime":
                errors[e.Field()] = e.Field() + " phải có định dạng ngày tháng hợp lệ (YYYY-MM-DD)."
            }
            log.Printf("Validation error — field: '%s', tag: '%s', error: %s",
                e.Field(), e.Tag(), e.Error())
        }
        return gin.H{"error": errors}
    }
    return gin.H{"error": "Yêu cầu không hợp lệ: " + Err.Error()}
}

// RegisterValidationError — đăng ký custom validation rule khi khởi động server
func RegisterValidationError() error {
    v, ok := binding.Validator.Engine().(*validator.Validate)
    if !ok {
        return fmt.Errorf("failed to register validation: could not get validator engine")
    }

    slugRegex := regexp.MustCompile(`^[a-z0-9]+(?:[-.][a-z0-9]+)*$`)
    v.RegisterValidation("slug", func(fl validator.FieldLevel) bool {
        return slugRegex.MatchString(fl.Field().String())
    })

    searchRegex := regexp.MustCompile(`^[a-zA-Z0-9\s]+$`)
    v.RegisterValidation("search", func(fl validator.FieldLevel) bool {
        return searchRegex.MatchString(fl.Field().String())
    })

    return nil
}
```

---

## Tổng kết các tag validation

### Tag có sẵn trong validator

|Tag|Ý nghĩa|Ví dụ|
|---|---|---|
|`required`|Bắt buộc phải có|`binding:"required"`|
|`gt=n`|Lớn hơn n|`binding:"gt=0"`|
|`gte=n`|Lớn hơn hoặc bằng n|`binding:"gte=1"`|
|`lte=n`|Nhỏ hơn hoặc bằng n|`binding:"lte=100"`|
|`min=n`|Độ dài tối thiểu n ký tự|`binding:"min=3"`|
|`max=n`|Độ dài tối đa n ký tự|`binding:"max=50"`|
|`uuid`|Chuỗi UUID hợp lệ|`binding:"uuid"`|
|`email`|Địa chỉ email hợp lệ|`binding:"email"`|
|`oneof=a b c`|Một trong các giá trị cho phép|`binding:"oneof=php golang python"`|
|`omitempty`|Bỏ qua nếu trường rỗng/zero|`binding:"omitempty,min=1"`|
|`datetime=layout`|Định dạng ngày tháng|`binding:"datetime=2006-01-02"`|

### Tag tự định nghĩa (bài này)

|Tag|Quy tắc|Đăng ký tại|
|---|---|---|
|`slug`|Chỉ chữ thường, số, `-`, `.`|`RegisterValidationError()`|
|`search`|Chỉ chữ cái, số, khoảng trắng|`RegisterValidationError()`|

### Cần nhớ

1. **Struct field phải là exported (chữ HOA)** — trường `uuid string` không hoạt động, phải là `UUID string`.
2. **`uri:"key"` cho path param, `form:"key"` cho query param** — dùng nhầm tag sẽ không bind được.
3. **`omitempty` đặt trước các rule khác** — `binding:"omitempty,min=3"` đúng; `binding:"min=3,omitempty"` có thể không hoạt động như mong muốn.
4. **Route cụ thể đặt trước route động** — `/search` phải đứng trước `/:slug` trong cùng một group.
5. **Gọi `RegisterValidationError()` trước khi khởi tạo handler** — custom rule phải được đăng ký trước khi bất kỳ validation nào chạy.