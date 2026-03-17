# Tối Ưu Router Validation - Xây Dựng Hàm Validation Dùng Chung

## Mục lục

1. [Vấn đề của cách tiếp cận cũ](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#v%E1%BA%A5n-%C4%91%E1%BB%81-c%E1%BB%A7a-c%C3%A1ch-ti%E1%BA%BFp-c%E1%BA%ADn-c%C5%A9)
2. [Giải pháp: Tách Validation thành package riêng](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#gi%E1%BA%A3i-ph%C3%A1p-t%C3%A1ch-validation-th%C3%A0nh-package-ri%C3%AAng)
3. [Xây dựng package `utils/validation.go`](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#x%C3%A2y-d%E1%BB%B1ng-package-utilsvalidationgo)
4. [Cập nhật các Handler sử dụng utils](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#c%E1%BA%ADp-nh%E1%BA%ADt-c%C3%A1c-handler-s%E1%BB%AD-d%E1%BB%A5ng-utils)
5. [Cấu trúc thư mục sau khi tối ưu](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#c%E1%BA%A5u-tr%C3%BAc-th%C6%B0-m%E1%BB%A5c-sau-khi-t%E1%BB%91i-%C6%B0u)
6. [Kiểm thử API](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#ki%E1%BB%83m-th%E1%BB%AD-api)
7. [Tổng kết](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#t%E1%BB%95ng-k%E1%BA%BFt)

---

## Vấn đề của cách tiếp cận cũ

Ở bài trước, mỗi handler tự viết logic validation riêng của mình:

```go
// Trong user.go
if _, err := uuid.Parse(id); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"message": "invalid uuid format"})
    return
}

// Trong product.go
if !slugRegex.MatchString(slug) {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Slug must contain only..."})
    return
}

// Trong product.go (SearchProducts)
if search == "" {
    c.JSON(http.StatusBadRequest, gin.H{"message": "Search query cannot be empty"})
    return
}
```

Cách này dẫn đến **3 vấn đề lớn**:

|Vấn đề|Hậu quả|
|---|---|
|**Lặp code (Code duplication)**|Cùng một logic validation nhưng viết lại ở nhiều nơi|
|**Thông báo lỗi không nhất quán**|Mỗi nơi tự đặt message khác nhau, gây khó chịu cho client|
|**Khó bảo trì**|Muốn thay đổi quy tắc validation phải sửa ở nhiều file|

**Nguyên tắc DRY (Don't Repeat Yourself):** Mỗi đoạn logic chỉ nên được viết **một lần duy nhất** ở một nơi.

---

## Giải pháp: Tách Validation thành package riêng

Tạo một package `utils` chứa các hàm validation **dùng chung** cho toàn bộ ứng dụng. Các handler chỉ cần **gọi hàm**, không tự viết lại logic.

```
Trước:                          Sau:
┌─────────────┐                 ┌─────────────┐
│  user.go    │ ← tự validate   │  user.go    │ ──┐
├─────────────┤                 ├─────────────┤   │  gọi hàm
│  product.go │ ← tự validate   │  product.go │ ──┼──────────→ utils/validation.go
├─────────────┤                 ├─────────────┤   │
│ category.go │ ← tự validate   │ category.go │ ──┘
└─────────────┘                 └─────────────┘
```

---

## Xây dựng package `utils/validation.go`

Tạo file `Router_Group/utils/validation.go`:

```go
package utils

import (
    "fmt"
    "regexp"
    "strconv"

    "github.com/google/uuid"
)
```

> 💡 **Tên package là `utils`** — quy ước đặt tên cho package chứa các hàm tiện ích (utility functions) dùng chung trong Go. Các package khác import bằng đường dẫn `"router-group/utils"`.

---

### Hàm 1: `ValidationRequired` — Kiểm tra trường bắt buộc

```go
// ValidationRequired kiểm tra xem một trường có bị bỏ trống không.
// Tham số:
//   - fieldname: tên trường (dùng trong thông báo lỗi)
//   - value:     giá trị cần kiểm tra
// Trả về: error nếu value rỗng, nil nếu hợp lệ
func ValidationRequired(fieldname, value string) error {
    if value == "" {
        return fmt.Errorf("The %s field must not be blank.", fieldname)
    }
    return nil
}
```

**Ví dụ sử dụng:**

```go
err := utils.ValidationRequired("search", "")
// → error: "The search field must not be blank."

err = utils.ValidationRequired("search", "iphone")
// → nil (hợp lệ)
```

---

### Hàm 2: `ValidationLength` — Kiểm tra độ dài chuỗi

```go
// ValidationLength kiểm tra độ dài của một chuỗi có nằm trong khoảng [min, max] không.
// Tham số:
//   - fieldname: tên trường
//   - value:     chuỗi cần kiểm tra
//   - min:       độ dài tối thiểu (tính bằng byte)
//   - max:       độ dài tối đa (tính bằng byte)
func ValidationLength(fieldname, value string, min, max int) error {
    if len(value) < min || len(value) > max {
        return fmt.Errorf("The %s field must be between %d and %d characters long.", fieldname, min, max)
    }
    return nil
}
```

> ⚠️ **Lưu ý về `len()` với tiếng Việt:** Hàm `len()` trong Go đếm số **byte**, không phải số **ký tự**. Mỗi ký tự tiếng Việt (UTF-8) chiếm 2-3 byte, nên `len("iphone")` = 6 nhưng `len("điện thoại")` = 19 dù nhìn có vẻ 10 ký tự. Nếu cần đếm đúng ký tự Unicode, dùng `len([]rune(value))` thay thế.

**Ví dụ sử dụng:**

```go
err := utils.ValidationLength("search", "ip", 3, 50)
// → error: "The search field must be between 3 and 50 characters long."

err = utils.ValidationLength("search", "iphone", 3, 50)
// → nil (hợp lệ)
```

---

### Hàm 3: `ValidationRegex` — Kiểm tra định dạng bằng Regex

```go
// ValidationRegex kiểm tra xem một chuỗi có khớp với biểu thức chính quy không.
// Tham số:
//   - fieldname: tên trường
//   - value:     chuỗi cần kiểm tra
//   - reg:       con trỏ đến *regexp.Regexp đã được biên dịch sẵn
func ValidationRegex(fieldname, value string, reg *regexp.Regexp) error {
    if !reg.MatchString(value) {
        return fmt.Errorf("The %s field is not in the correct format (%s).", fieldname, reg.String())
    }
    return nil
}
```

> 💡 **Tại sao nhận `*regexp.Regexp` thay vì nhận chuỗi pattern?** Việc biên dịch (compile) một regex tốn chi phí CPU. Nếu hàm nhận chuỗi pattern và tự compile mỗi lần gọi, hiệu suất sẽ kém. Thay vào đó, ta compile một lần ở cấp package (với `regexp.MustCompile`), rồi truyền con trỏ vào hàm để tái sử dụng — nhanh hơn đáng kể.

**Ví dụ sử dụng:**

```go
var slugRegex = regexp.MustCompile(`^[a-z0-9]+(?:[.-][a-z0-9]+)*$`)

err := utils.ValidationRegex("slug", "Iphone-14", slugRegex)
// → error: "The slug field is not in the correct format (^[a-z0-9]+(?:[.-][a-z0-9]+)*$)."

err = utils.ValidationRegex("slug", "iphone-14", slugRegex)
// → nil (hợp lệ)
```

---

### Hàm 4: `ValidationPositiveInt` — Kiểm tra số nguyên dương

```go
// ValidationPositiveInt kiểm tra xem một chuỗi có thể chuyển thành số nguyên dương không.
// Tham số:
//   - fieldname: tên trường
//   - value:     chuỗi cần kiểm tra (ví dụ: tham số :id từ URL luôn là string)
func ValidationPositiveInt(fieldname, value string) error {
    v, err := strconv.Atoi(value) // Chuyển chuỗi → số nguyên

    if err != nil || v <= 0 {
        return fmt.Errorf("The requirement is that %s must be a positive integer.", fieldname)
    }
    return nil
}
```

> 💡 **Tại sao tham số `value` là `string` chứ không phải `int`?** Tất cả tham số lấy từ URL (`c.Param()`, `c.Query()`) đều là **chuỗi** (string). Go không tự động ép kiểu, nên ta cần `strconv.Atoi()` để chuyển sang số nguyên. Hàm này kết hợp cả hai bước: chuyển kiểu và kiểm tra hợp lệ.

**Ví dụ sử dụng:**

```go
err := utils.ValidationPositiveInt("id", "abc")
// → error: "The requirement is that id must be a positive integer."

err = utils.ValidationPositiveInt("id", "-5")
// → error (âm không hợp lệ)

err = utils.ValidationPositiveInt("id", "3")
// → nil (hợp lệ)
```

---

### Hàm 5: `ValidationUUID` — Kiểm tra và parse UUID

```go
// ValidationUUID kiểm tra xem một chuỗi có đúng định dạng UUID không.
// Khác các hàm khác, hàm này trả về 2 giá trị:
//   - uuid.UUID: giá trị UUID đã được parse (dùng được ngay, không cần parse lại)
//   - error:     lỗi nếu không hợp lệ, nil nếu hợp lệ
func ValidationUUID(fieldname, value string) (uuid.UUID, error) {
    uid, err := uuid.Parse(value)
    if err != nil {
        return uuid.Nil, fmt.Errorf("Field %s: '%s' is not a valid UUID.", fieldname, value)
    }
    return uid, nil
}
```

> 💡 **`uuid.Nil` là gì?** Đây là giá trị UUID "rỗng" — tất cả 128 bit đều bằng 0, biểu diễn là `00000000-0000-0000-0000-000000000000`. Dùng làm giá trị trả về mặc định khi có lỗi, tương tự như trả về `0` cho int hay `""` cho string.

> 💡 **Tại sao trả về cả `uuid.UUID` lẫn `error`?** Đây là **Go idiom** (thành ngữ Go) — hàm có thể thất bại nên trả về cả kết quả lẫn lỗi. Caller kiểm tra `error` trước, nếu `nil` thì dùng `uuid.UUID` trực tiếp mà không cần parse lại lần nữa.

**Ví dụ sử dụng:**

```go
uid, err := utils.ValidationUUID("UUID", "not-a-uuid")
// → uid = uuid.Nil, err = "Field UUID: 'not-a-uuid' is not a valid UUID."

uid, err = utils.ValidationUUID("UUID", "550e8400-e29b-41d4-a716-446655440000")
// → uid = (UUID object), err = nil
```

---

### Hàm 6: `ValidationInList` — Kiểm tra giá trị có trong danh sách

```go
// ValidationInList kiểm tra xem một giá trị có tồn tại trong danh sách cho phép không.
// Tham số:
//   - fieldName: tên trường
//   - value:     giá trị cần kiểm tra
//   - list:      map các giá trị hợp lệ (dùng map[string]bool để tra cứu O(1))
func ValidationInList(fieldName, value string, list map[string]bool) error {
    if !list[value] {
        return fmt.Errorf(
            "The value '%s' of field '%s' does not exist in the list.\nValid values: %v",
            value, fieldName, GetKeysFromMap(list),
        )
    }
    return nil
}

// GetKeysFromMap là hàm hỗ trợ — trích xuất tất cả keys từ map thành slice.
// Dùng để hiển thị danh sách giá trị hợp lệ trong thông báo lỗi.
// Lưu ý: thứ tự các key trong map không cố định (Go map là unordered).
func GetKeysFromMap(m map[string]bool) []string {
    // make([]string, 0, len(m)): tạo slice rỗng với capacity = số phần tử trong map
    // Điều này tránh việc Go phải cấp phát thêm bộ nhớ khi append
    keys := make([]string, 0, len(m))
    for key := range m {
        keys = append(keys, key)
    }
    return keys
}
```

**Ví dụ sử dụng:**

```go
categories := map[string]bool{"golang": true, "python": true, "csharp": true}

err := utils.ValidationInList("category", "ruby", categories)
// → error: "The value 'ruby' of field 'category' does not exist in the list.
//           Valid values: [golang python csharp]"

err = utils.ValidationInList("category", "golang", categories)
// → nil (hợp lệ)
```

---

### Toàn bộ file `utils/validation.go`

```go
package utils

import (
    "fmt"
    "regexp"
    "strconv"

    "github.com/google/uuid"
)

// ValidationRequired kiểm tra trường bắt buộc không được để trống.
func ValidationRequired(fieldname, value string) error {
    if value == "" {
        return fmt.Errorf("The %s field must not be blank.", fieldname)
    }
    return nil
}

// ValidationLength kiểm tra độ dài chuỗi nằm trong khoảng [min, max].
func ValidationLength(fieldname, value string, min, max int) error {
    if len(value) < min || len(value) > max {
        return fmt.Errorf("The %s field must be between %d and %d characters long.", fieldname, min, max)
    }
    return nil
}

// ValidationRegex kiểm tra chuỗi khớp với biểu thức chính quy.
func ValidationRegex(fieldname, value string, reg *regexp.Regexp) error {
    if !reg.MatchString(value) {
        return fmt.Errorf("The %s field is not in the correct format (%s).", fieldname, reg.String())
    }
    return nil
}

// ValidationPositiveInt kiểm tra chuỗi có thể chuyển thành số nguyên dương.
func ValidationPositiveInt(fieldname, value string) error {
    v, err := strconv.Atoi(value)
    if err != nil || v <= 0 {
        return fmt.Errorf("The requirement is that %s must be a positive integer.", fieldname)
    }
    return nil
}

// ValidationUUID kiểm tra và parse UUID, trả về (uuid.UUID, error).
func ValidationUUID(fieldname, value string) (uuid.UUID, error) {
    uid, err := uuid.Parse(value)
    if err != nil {
        return uuid.Nil, fmt.Errorf("Field %s: '%s' is not a valid UUID.", fieldname, value)
    }
    return uid, nil
}

// ValidationInList kiểm tra giá trị có tồn tại trong danh sách hợp lệ.
func ValidationInList(fieldName, value string, list map[string]bool) error {
    if !list[value] {
        return fmt.Errorf(
            "The value '%s' of field '%s' does not exist in the list.\nValid values: %v",
            value, fieldName, GetKeysFromMap(list),
        )
    }
    return nil
}

// GetKeysFromMap trích xuất tất cả keys từ map thành slice (dùng trong thông báo lỗi).
func GetKeysFromMap(m map[string]bool) []string {
    keys := make([]string, 0, len(m))
    for key := range m {
        keys = append(keys, key)
    }
    return keys
}
```

---

## Cập nhật các Handler sử dụng utils

### 1. Cập nhật `internal/api/v1/handler/product.go`

Thêm import `"router-group/utils"` và cập nhật hai hàm:

```go
import (
    "fmt"
    "net/http"
    "regexp"
    "strings"

    "github.com/gin-gonic/gin"
    "router-group/utils" // Import package utils dùng chung
)

var slugRegex = regexp.MustCompile(`^[a-z0-9]+(?:[.-][a-z0-9]+)*$`)

// GetProductBySlug xử lý GET /api/v1/products/:slug
// Thay vì tự viết if !slugRegex.MatchString(slug) {...}
// → gọi utils.ValidationRegex() — ngắn gọn và nhất quán hơn
func (obj *Product) GetProductBySlug(c *gin.Context) {
    slug := c.Param("slug")

    // Dùng hàm dùng chung thay vì tự validate
    if err := utils.ValidationRegex("slug", slug, slugRegex); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "message": err.Error(),
        })
        return
    }

    for _, product := range Products {
        if fmt.Sprintf("%v", product.ID) == slug {
            c.JSON(http.StatusOK, gin.H{
                "message": "Product found",
                "data":    product,
            })
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{
        "message": "Product not found",
        "data":    nil,
    })
}

// SearchProducts xử lý GET /api/v1/products?search=...
// Cải tiến: tách riêng từng validation thay vì lồng else-if
func (obj *Product) SearchProducts(c *gin.Context) {
    search := c.Query("search")

    // Validate 1: Không được rỗng
    if err := utils.ValidationRequired("search", search); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "data": nil})
        return
    }

    // Validate 2: Độ dài từ 3 đến 50 ký tự
    // Tách thành if riêng thay vì else-if — rõ ràng và dễ thêm rule mới hơn
    if err := utils.ValidationLength("search", search, 3, 50); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "data": nil})
        return
    }

    // Tìm kiếm không phân biệt hoa thường
    var searchResults []Product
    for _, product := range Products {
        if strings.Contains(strings.ToLower(product.ProductName), strings.ToLower(search)) {
            searchResults = append(searchResults, product)
        }
    }

    if len(searchResults) == 0 {
        c.JSON(http.StatusNotFound, gin.H{
            "message": "No products found",
            "data":    nil,
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Search results",
        "data":    searchResults,
    })
}
```

> 💡 **Tại sao dùng `if` riêng thay vì `if-else if-else`?** Cách cũ dùng `else if` buộc phải lồng nhau, khó đọc khi có nhiều rule. Cách mới dùng **early return** — mỗi lỗi trả về ngay, code phẳng hơn (flat), dễ thêm rule mới mà không ảnh hưởng phần còn lại. Đây là pattern phổ biến trong Go, thường gọi là **"guard clause"**.

---

### 2. Cập nhật `internal/api/v1/handler/user.go`

Thêm import `"router-group/utils"` và cập nhật hàm `GetUserByUUID`:

```go
import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
    "router-group/utils"
)

// GetUserByUUID xử lý GET /api/v1/users/:uuid
func (obj *User) GetUserByUUID(c *gin.Context) {
    id := c.Param("uuid")

    // utils.ValidationUUID trả về (uuid.UUID, error)
    // Dấu _ bỏ qua giá trị UUID vì ta so sánh bằng string bên dưới
    _, err := utils.ValidationUUID("UUID", id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "message": err.Error(),
        })
        return
    }

    for _, user := range users {
        if fmt.Sprintf("%v", user.ID) == id {
            c.JSON(http.StatusOK, gin.H{
                "message": "user found",
                "data":    user,
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

> 💡 **Tại sao response trả về `"message"` và `"error"` cùng một nội dung trong code gốc?** Đây là dư thừa — không cần thiết phải có cả hai key với cùng giá trị. Phiên bản cải thiện chỉ giữ lại `"message"` cho gọn.

---

### 3. Cập nhật `internal/api/v1/handler/category.go`

Thêm import `"router-group/utils"` và cập nhật hàm `GetCategoryByCategories`:

```go
import (
    "net/http"

    "github.com/gin-gonic/gin"
    "router-group/utils"
)

var categories = map[string]bool{
    "golang": true,
    "csharp": true,
    "python": true,
}

type Category struct{}

func NewCategoryHandler() *Category {
    return &Category{}
}

// GetCategoryByCategories xử lý GET /api/v1/categories/:category
func (obj *Category) GetCategoryByCategories(c *gin.Context) {
    category := c.Param("category")

    // utils.ValidationInList kiểm tra và tự tạo thông báo lỗi kèm danh sách hợp lệ
    if err := utils.ValidationInList("category", category, categories); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "message": err.Error(),
            "data":    nil,
        })
        return // Bắt buộc có return để dừng xử lý
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "success",
        "data":    category,
    })
}
```

---

## Cấu trúc thư mục sau khi tối ưu

```
Router_Group/
├── main.go
├── go.mod
├── utils/                           ← Package mới (dùng chung)
│   └── validation.go                ← Tất cả hàm validation tập trung tại đây
└── internal/
    └── api/
        ├── v1/
        │   └── handler/
        │       ├── user.go          ← Import utils, bỏ logic validate thủ công
        │       ├── product.go       ← Import utils, bỏ logic validate thủ công
        │       └── category.go      ← Import utils, bỏ logic validate thủ công
        └── v2/
            └── handler/
                └── user.go
```

**Lợi ích thực tế:** Nếu mai này cần thay đổi quy tắc UUID (ví dụ chuyển sang UUID v7), bạn chỉ cần sửa **một chỗ duy nhất** trong `utils/validation.go` thay vì tìm và sửa ở tất cả các handler.

---

## So sánh trước và sau khi tối ưu

```
Trước (handler tự validate):          Sau (dùng utils):
─────────────────────────────         ──────────────────────────────────
if _, err := uuid.Parse(id);          _, err := utils.ValidationUUID("UUID", id)
   err != nil {                        if err != nil {
    c.JSON(400, gin.H{                     c.JSON(400, gin.H{
        "message": "invalid uuid           "message": err.Error(),
                   format",            })
        "error": err.Error(),          return
    })                                }
    return
}
```

Logic validation giảm từ 7 dòng xuống còn 5 dòng, và quan trọng hơn: **thông báo lỗi được tạo tự động, nhất quán** ở mọi nơi.

---

## Kiểm thử API

### Kiểm tra `ValidationRequired` + `ValidationLength` (SearchProducts)

```bash
# Lỗi: search rỗng → "The search field must not be blank."
curl "http://localhost:8080/api/v1/products?search="

# Lỗi: quá ngắn (< 3 ký tự) → "The search field must be between 3 and 50 characters long."
curl "http://localhost:8080/api/v1/products?search=ip"

# Lỗi: quá dài (> 50 ký tự)
curl "http://localhost:8080/api/v1/products?search=aaaaabbbbbcccccdddddeeeeeaaaaabbbbbcccccdddddeeeee1"

# Hợp lệ
curl "http://localhost:8080/api/v1/products?search=iphone"
```

### Kiểm tra `ValidationUUID` (GetUserByUUID)

```bash
# Lỗi: sai định dạng → "Field UUID: 'abc-123' is not a valid UUID."
curl http://localhost:8080/api/v1/users/abc-123

# Hợp lệ (thay bằng UUID thực từ GET /users)
curl http://localhost:8080/api/v1/users/550e8400-e29b-41d4-a716-446655440000
```

### Kiểm tra `ValidationInList` (GetCategoryByCategories)

```bash
# Lỗi: → "The value 'ruby' of field 'category' does not exist in the list. Valid values: [golang python csharp]"
curl http://localhost:8080/api/v1/categories/ruby

# Hợp lệ
curl http://localhost:8080/api/v1/categories/golang
```

### Kiểm tra `ValidationRegex` (GetProductBySlug)

```bash
# Lỗi: chữ hoa → "The slug field is not in the correct format..."
curl http://localhost:8080/api/v1/products/Iphone-14

# Hợp lệ
curl http://localhost:8080/api/v1/products/iphone-14
```

---

## Tổng kết

Bài này áp dụng nguyên tắc **DRY (Don't Repeat Yourself)** vào tầng validation:

| Hàm                     | Kiểm tra                             | Trả về               |
| ----------------------- | ------------------------------------ | -------------------- |
| `ValidationRequired`    | Chuỗi không được rỗng                | `error`              |
| `ValidationLength`      | Độ dài trong khoảng [min, max]       | `error`              |
| `ValidationRegex`       | Khớp biểu thức chính quy             | `error`              |
| `ValidationPositiveInt` | Chuỗi là số nguyên dương             | `error`              |
| `ValidationUUID`        | Chuỗi đúng định dạng UUID            | `(uuid.UUID, error)` |
| `ValidationInList`      | Giá trị nằm trong danh sách cho phép | `error`              |

### Những điểm quan trọng cần nhớ

1. **Tách validation ra package riêng** giúp code tuân thủ nguyên tắc DRY và dễ bảo trì.
2. **Dùng `if` + `return` (guard clause)** thay vì `if-else if-else` lồng nhau — code phẳng hơn, dễ đọc hơn.
3. **Luôn có `return` sau mỗi `c.JSON` trả lỗi** — thiếu `return` là bug thầm lặng rất nguy hiểm.
4. **Hàm trả về `error`** là pattern chuẩn của Go — caller kiểm tra `err != nil` trước khi dùng kết quả.
5. **`make([]string, 0, len(m))`** khi tạo slice với kích thước biết trước — tránh cấp phát bộ nhớ nhiều lần khi `append`.