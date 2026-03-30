# Validate Body Request (JSON) với Validator Package

## Mục lục

1. [Tổng quan về Validate Body Request](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#t%E1%BB%95ng-quan-v%E1%BB%81-validate-body-request)
2. [Đọc raw body vs Bind JSON](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#%C4%91%E1%BB%8Dc-raw-body-vs-bind-json)
3. [Validate JSON cơ bản với `ShouldBindBodyWithJSON`](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#validate-json-c%C6%A1-b%E1%BA%A3n-v%E1%BB%9Bi-shouldbindbodywithJSON)
4. [Nested Struct — Validate struct lồng nhau](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#nested-struct--validate-struct-l%E1%BB%93ng-nhau)
5. [Custom Validation cho số nguyên: `min_int` và `max_int`](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#custom-validation-cho-s%E1%BB%91-nguy%C3%AAn-min_int-v%C3%A0-max_int)
6. [Default Value với con trỏ `*bool`](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#default-value-v%E1%BB%9Bi-con-tr%E1%BB%8F-bool)
7. [Custom Validation cho phần mở rộng file: `file_extension`](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#custom-validation-cho-ph%E1%BA%A7n-m%E1%BB%9F-r%E1%BB%99ng-file-file_extension)
8. [Validate Slice với tag `dive`](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#validate-slice-v%E1%BB%9Bi-tag-dive)
9. [Giải quyết vấn đề lỗi khi slice có phần tử rỗng](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#gi%E1%BA%A3i-quy%E1%BA%BFt-v%E1%BA%A5n-%C4%91%E1%BB%81-l%E1%BB%97i-khi-slice-c%C3%B3-ph%E1%BA%A7n-t%E1%BB%AD-r%E1%BB%97ng)
10. [Validate Map và tự kiểm tra key](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#validate-map-v%C3%A0-t%E1%BB%B1-ki%E1%BB%83m-tra-key)
11. [Toàn bộ struct `Product` sau khi hoàn thiện](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#to%C3%A0n-b%E1%BB%99-struct-product-sau-khi-ho%C3%A0n-thi%E1%BB%87n)
12. [Tổng kết](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#t%E1%BB%95ng-k%E1%BA%BFt)

---

## Tổng quan về Validate Body Request

Ở các bài trước, ta validate **Path Param** (`ShouldBindUri`) và **Query Param** (`ShouldBindQuery`). Bài này tập trung vào phần còn lại: validate **Body** của request — dữ liệu JSON mà client gửi lên qua phương thức `POST`, `PUT`, `PATCH`.

```
Client gửi request:
┌─────────────────────────────────────────┐
│ POST /api/v2/products                   │
│ Content-Type: application/json          │
│                                         │
│ {                                       │  ← Body (JSON)
│   "name": "Iphone 15",                  │     cần validate
│   "price": 25000000                     │
│ }                                       │
└─────────────────────────────────────────┘
```

**Ba hàm bind tương ứng ba nguồn dữ liệu:**

|Nguồn|Hàm bind|Tag trong struct|
|---|---|---|
|Path param (`/:id`)|`ShouldBindUri`|`uri:"key"`|
|Query string (`?key=val`)|`ShouldBindQuery`|`form:"key"`|
|Request body (JSON)|`ShouldBindBodyWithJSON`|`json:"key"`|

---

## Đọc raw body vs Bind JSON

### Cách 1: Đọc raw body (chưa có validation)

```go
func (obj *Product) CreateProduct(c *gin.Context) {

    // GetRawData() đọc toàn bộ body dưới dạng []byte (mảng byte thô)
    // Không kiểm tra định dạng, không validate — đọc bất cứ thứ gì client gửi
    body, err := c.GetRawData()
    if err != nil {
        c.JSON(http.StatusBadRequest, "Error reading request body: "+err.Error())
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "create Product successfully",
        "data":    string(body), // Chuyển []byte → string để hiển thị
    })
}
```

**Kiểm thử:**

```
POST /api/v2/products
Body: Welcome my website

Output:
{
    "data": "Welcome my website",
    "message": "create Product successfully"
}
```

> 💡 **`GetRawData()` trả về `[]byte`** — mảng byte thô, chưa được parse. Dùng `string(body)` để chuyển sang chuỗi hiển thị được. Cách này hữu ích khi cần đọc body cho mục đích đặc biệt (như tính chữ ký HMAC), nhưng **không dùng cho validate JSON thông thường**.

### Cách 2: Bind và Validate JSON (chuẩn)

```go
func (obj *Product) CreateProduct(c *gin.Context) {

    var param Product
    // ShouldBindBodyWithJSON đồng thời làm 3 việc:
    // 1. Đọc body từ request
    // 2. Parse JSON → điền vào struct param
    // 3. Kiểm tra tất cả các tag binding:"..."
    if err := c.ShouldBindBodyWithJSON(&param); err != nil {
        c.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "create Product successfully",
        "data":    param,
    })
}
```

---

## Validate JSON cơ bản với `ShouldBindBodyWithJSON`

### Khai báo struct với validation tag

```go
// Product là struct đại diện cho dữ liệu sản phẩm nhận từ request body
// Tag json:"key"     → tên field trong JSON (snake_case theo chuẩn REST API)
// Tag binding:"..."  → điều kiện validation
type Product struct {
    ID    int     `json:"id"`
    // required: bắt buộc phải có
    // min=3:    tên ít nhất 3 ký tự
    // max=100:  tên tối đa 100 ký tự
    Name  string  `json:"name"    binding:"required,min=3,max=100"`
    // required: bắt buộc phải có
    // gt=0:     giá phải > 0 (không cho phép âm hoặc bằng 0)
    Price float64 `json:"price"   binding:"required,gt=0"`
    // omitempty: nếu client không gửi field này → bỏ qua, không báo lỗi
    Display bool  `json:"display" binding:"omitempty"`
}
```

**Kiểm thử:**

```json
POST /api/v2/products
Body:
{
    "id": 15,
    "name": "Ship qua 2 quả tên lửa",
    "price": 100000.001
}

Output:
{
    "data": {
        "id": 15,
        "name": "Ship qua 2 quả tên lửa",
        "price": 100000.001,
        "display": false
    },
    "message": "create Product successfully"
}
```

> 💡 **`display` trả về `false` dù client không gửi?** Đây là **zero value** của kiểu `bool` trong Go. Khi JSON không chứa field `display`, Go tự điền giá trị mặc định: `bool` → `false`, `int` → `0`, `string` → `""`. Xem phần [Default Value](https://claude.ai/chat/7b7a5cc2-9582-4e30-af9a-a46fe9e09fd4#default-value-v%E1%BB%9Bi-con-tr%E1%BB%8F-bool) để xử lý trường hợp này.

---

## Nested Struct — Validate struct lồng nhau

Validator tự động kiểm tra **đệ quy** vào các struct lồng bên trong. Không cần làm gì thêm — chỉ cần khai báo đúng tag là đủ.

```go
type Product struct {
    ID           int          `json:"id"`
    Name         string       `json:"name"          binding:"required,min=3,max=100"`
    Price        float64      `json:"price"         binding:"required,gt=0"`
    Display      bool         `json:"display"       binding:"omitempty"`
    // binding:"required" trên nested struct: đảm bảo client phải gửi object product_image
    // Validator sẽ tự động kiểm tra các field bên trong ProductImage
    ProductImage ProductImage `json:"product_image" binding:"required"`
}

type ProductImage struct {
    // url: tag có sẵn trong validator — kiểm tra chuỗi có đúng định dạng URL không
    ImageName string `json:"name" binding:"required,min=3,max=100"`
    ImageLink string `json:"link" binding:"required,url"`
}
```

**Kiểm thử thành công:**

```json
POST /api/v2/products
Body:
{
    "id": 15,
    "name": "Ship qua 2 quả tên lửa",
    "price": 100000.001,
    "product_image": {
        "name": "hinh1",
        "link": "http://localhost:8080/api/v2/products/15/hinh1"
    }
}

Output:
{
    "data": {
        "id": 15,
        "name": "Ship qua 2 quả tên lửa",
        "price": 100000.001,
        "display": false,
        "product_image": {
            "name": "hinh1",
            "link": "http://localhost:8080/api/v2/products/15/hinh1"
        }
    },
    "message": "create Product successfully"
}
```

> 💡 **Validator kiểm tra nested struct như thế nào?** Khi gặp trường kiểu struct (như `ProductImage`), validator tự động "đi sâu vào" và kiểm tra các field bên trong. Không cần tag `dive` với nested struct thông thường — `dive` chỉ cần thiết với **slice và map** (xem phần sau).

---

## Custom Validation cho số nguyên: `min_int` và `max_int`

### Tại sao cần `min_int` / `max_int`?

Tag `min` và `max` có sẵn trong validator **hoạt động khác nhau tùy theo kiểu dữ liệu:**

|Kiểu|`min=3` nghĩa là|
|---|---|
|`string`|Độ dài chuỗi ≥ 3 ký tự|
|`int`, `float64`|Giá trị số ≥ 3|
|`slice`, `map`|Số phần tử ≥ 3|

Vậy với `int`, ta có thể dùng `gte=1000` thay vì `min_int=1000`. Tuy nhiên, `min_int` / `max_int` là custom rule được tạo ra để **tên tag rõ ràng hơn** với người đọc code — ngay khi nhìn vào biết đây là giới hạn số nguyên.

### Thêm vào `utils/validation.go`

```go
// Trong RegisterValidationError():

// --- min_int: kiểm tra field >= giá trị tối thiểu ---
// fl.Param() lấy phần tham số sau dấu =, ví dụ: min_int=1000 → fl.Param() = "1000"
// strconv.ParseInt chuyển chuỗi "1000" sang số nguyên 64-bit
//   - tham số 10: hệ thập phân (base 10)
//     + 10: hệ thập phân (0-9)
//     + 16: hệ thập lục phân (0-9, a-f)
//     +  8: hệ bát phân (0-7)
//     +  0: tự động nhận diện (prefix 0x=hex, 0=octal, còn lại=decimal)
//   - tham số 64: kích thước bit của kết quả (int64)
//     +  0: int (phụ thuộc platform: 32 hoặc 64 bit)
//     +  8: int8  (-128 đến 127)
//     + 16: int16 (-32768 đến 32767)
//     + 32: int32 (-2^31 đến 2^31-1)
//     + 64: int64 (-2^63 đến 2^63-1) ← dùng 64 để an toàn nhất
v.RegisterValidation("min_int", func(fl validator.FieldLevel) bool {
    minStr := fl.Param()
    minVal, err := strconv.ParseInt(minStr, 10, 64)
    if err != nil {
        return false // Nếu param không phải số → coi như không hợp lệ
    }
    return fl.Field().Int() >= minVal
})

// --- max_int: kiểm tra field <= giá trị tối đa ---
v.RegisterValidation("max_int", func(fl validator.FieldLevel) bool {
    maxStr := fl.Param()
    maxVal, err := strconv.ParseInt(maxStr, 10, 64)
    if err != nil {
        return false
    }
    return fl.Field().Int() <= maxVal
})
```

**Cập nhật thông báo lỗi trong `HandleValidationErrors`:**

```go
case "min_int":
    errors[e.Field()] = e.Field() + " phải lớn hơn hoặc bằng " + e.Param() + "."
case "max_int":
    errors[e.Field()] = e.Field() + " phải nhỏ hơn hoặc bằng " + e.Param() + "."
```

**Áp dụng vào struct Product:**

```go
type Product struct {
    ID    int    `json:"id"`
    Name  string `json:"name"  binding:"required,min=3,max=100"`
    // Price đổi sang int, dùng min_int và max_int thay vì gt=0
    Price int    `json:"price" binding:"required,min_int=1000,max_int=100000000"`
    // ...
}
```

---

## Default Value với con trỏ `*bool`

### Vấn đề: phân biệt "không gửi" và "gửi false"

Với kiểu `bool` thông thường, Go không thể phân biệt hai trường hợp:

- Client **không gửi** field `display` → Go điền `false` (zero value)
- Client **gửi `"display": false`** → cũng là `false`

Trong cả hai trường hợp, `param.Display == false` — không thể biết client có gửi hay không.

### Giải pháp: dùng con trỏ `*bool`

```go
type Product struct {
    // ...
    // *bool thay vì bool
    // Con trỏ có 3 trạng thái:
    //   nil   → client không gửi field này
    //   &true → client gửi "display": true
    //   &false → client gửi "display": false
    Display *bool `json:"display" binding:"omitempty"`
}

func (obj *Product) CreateProduct(c *gin.Context) {
    var param Product
    if err := c.ShouldBindBodyWithJSON(&param); err != nil {
        c.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
        return
    }

    // Kiểm tra nil để biết client có gửi field hay không
    if param.Display == nil {
        // Client không gửi → áp dụng giá trị mặc định là true
        defaultDisplay := true
        param.Display = &defaultDisplay
        // Phải tạo biến trung gian (defaultDisplay) rồi lấy địa chỉ
        // Không thể viết param.Display = &true (Go không cho phép lấy địa chỉ của literal)
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "create Product successfully",
        "data":    param,
    })
}
```

> 💡 **Tại sao không thể viết `param.Display = &true`?** Trong Go, bạn không thể lấy địa chỉ (`&`) của một **literal value** (giá trị trực tiếp như `true`, `false`, `42`). Phải gán vào biến trước rồi mới lấy địa chỉ của biến đó.

> 💡 **Kỹ thuật này áp dụng cho mọi kiểu dữ liệu**, không chỉ `bool`:
> 
> - `*int` — phân biệt "không gửi" với "gửi 0"
> - `*string` — phân biệt "không gửi" với "gửi chuỗi rỗng"
> - `*float64` — phân biệt "không gửi" với "gửi 0.0"

---

## Custom Validation cho phần mở rộng file: `file_extension`

### Yêu cầu

Field `ImageLink` trong `ProductImage` cần kiểm tra URL ảnh có phần mở rộng hợp lệ (`.jpg`, `.jpeg`, `.png`, `.gif`) hay không.

### Thêm vào `utils/validation.go`

```go
import "path/filepath"

// Danh sách phần mở rộng hợp lệ mặc định (dùng khi tag không có param)
var allowedExtensions = []string{".jpg", ".jpeg", ".png", ".gif"}

// Trong RegisterValidationError():

// --- file_extension: kiểm tra phần mở rộng file ---
// Hỗ trợ hai chế độ:
//   1. Có param:  binding:"file_extension=jpg jpeg png"  → kiểm tra theo danh sách param
//   2. Không param: binding:"file_extension"             → kiểm tra theo allowedExtensions mặc định
v.RegisterValidation("file_extension", func(fl validator.FieldLevel) bool {
    fileName := fl.Field().String()

    // Nếu tên file rỗng → bỏ qua validation (kết hợp với omitempty nếu cần)
    if fileName == "" {
        return true
    }

    // filepath.Ext() trích xuất phần mở rộng kể cả dấu chấm: "hinh1.jpg" → ".jpg"
    // strings.TrimPrefix bỏ dấu chấm đầu:                                  ".jpg" → "jpg"
    // strings.ToLower chuyển về chữ thường để so sánh không phân biệt hoa thường
    ext := strings.TrimPrefix(strings.ToLower(filepath.Ext(fileName)), ".")

    allowStr := fl.Param() // Lấy danh sách phần mở rộng từ param (nếu có)

    if allowStr != "" {
        // Chế độ 1: dùng danh sách từ param
        // fl.Param() với binding:"file_extension=jpg jpeg png" → "jpg jpeg png"
        allowedExts := strings.Split(allowStr, " ")
        for _, allow := range allowedExts {
            if ext == strings.ToLower(allow) {
                return true
            }
        }
        return false
    }

    // Chế độ 2: dùng danh sách mặc định
    for _, allowExists := range allowedExtensions {
        if ext == strings.TrimPrefix(strings.ToLower(allowExists), ".") {
            return true
        }
    }
    return false
})
```

**Cập nhật thông báo lỗi:**

```go
case "file_extension":
    errors[e.Field()] = e.Field() + " phải có phần mở rộng hợp lệ: " +
        strings.Join(allowedExtensions, ", ") + "."
```

**Áp dụng vào struct:**

```go
type ProductImage struct {
    ImageName string `json:"name" binding:"required,min=3,max=100"`
    // file_extension: dùng danh sách mặc định (.jpg, .jpeg, .png, .gif)
    ImageLink string `json:"link" binding:"required,file_extension"`
}
```

**Kiểm thử:**

```
# Hợp lệ
"link": "http://localhost:8080/image/hinh1.jpg" → ✅

# Không hợp lệ
"link": "http://localhost:8080/image/hinh1.pdf" → ❌
→ { "error": { "ImageLink": "ImageLink phải có phần mở rộng hợp lệ: .jpg, .jpeg, .png, .gif." } }
```

---

## Validate Slice với tag `dive`

### Tag `gt` và `lt` với Slice

Khi áp dụng lên một **slice**, các tag `gt` và `lt` kiểm tra **số lượng phần tử**:

```go
Tags []string `json:"tags" binding:"required,gt=3,lt=5"`
// required: slice không được nil/rỗng
// gt=3:     số phần tử phải > 3 (tức là ít nhất 4 phần tử)
// lt=5:     số phần tử phải < 5 (tức là tối đa 4 phần tử)
// → Tags phải có đúng 4 phần tử
```

### Tag `dive` — đi sâu vào từng phần tử của slice

Với slice chứa **struct**, tag `dive` bảo validator **"đi vào trong" từng phần tử** để kiểm tra các field bên trong.

```go
// gt=0: slice phải có ít nhất 1 phần tử
// dive: với MỖI phần tử trong slice, kiểm tra tiếp các tag binding bên trong ProductAttribute
ProductAttribute []ProductAttribute `json:"product_attribute" binding:"required,gt=0,dive"`

type ProductAttribute struct {
    AttributeName  string `json:"attribute_name"  binding:"required"`
    AttributeValue string `json:"attribute_value" binding:"required"`
}
```

> 💡 **Không có `dive` thì sao?** Nếu không có `dive`, validator chỉ kiểm tra **bản thân slice** (số lượng phần tử, có nil không...) chứ **không kiểm tra nội dung từng phần tử** bên trong. Muốn validate từng struct trong slice → bắt buộc phải có `dive`.

---

## Giải quyết vấn đề lỗi khi slice có phần tử rỗng

### Vấn đề

Khi client gửi slice có một phần tử rỗng `{}`:

```json
"product_attribute": [
    { "attribute_name": "color", "attribute_value": "white" },
    {}
]
```

Validator phát hiện `{}` vi phạm `required` trên các field bên trong và trả lỗi:

```json
{
    "error": {
        "AttributeName": "AttributeName là bắt buộc.",
        "AttributeValue": "AttributeValue là bắt buộc."
    }
}
```

Tuy nhiên thông báo lỗi này **không chỉ rõ phần tử nào** trong slice bị lỗi — khó cho client biết cần sửa chỗ nào.

### Giải pháp: Validate thủ công từng phần tử kèm vị trí

Sau khi `ShouldBindBodyWithJSON` pass (hoặc thay thế hoàn toàn), tự duyệt slice và kiểm tra từng phần tử:

```go
func (obj *Product) CreateProduct(c *gin.Context) {
    var param Product
    if err := c.ShouldBindBodyWithJSON(&param); err != nil {
        c.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
        return
    }

    // Validate thủ công từng phần tử trong ProductAttribute
    // Mục đích: trả về thông báo lỗi kèm VỊ TRÍ (index) phần tử bị lỗi
    attributeErrors := make(map[string]string)
    for i, attr := range param.ProductAttribute {
        if attr.AttributeName == "" {
            // fmt.Sprintf("product_attribute[%d].attribute_name", i) → rõ phần tử thứ mấy
            attributeErrors[fmt.Sprintf("product_attribute[%d].attribute_name", i)] =
                "AttributeName là bắt buộc."
        }
        if attr.AttributeValue == "" {
            attributeErrors[fmt.Sprintf("product_attribute[%d].attribute_value", i)] =
                "AttributeValue là bắt buộc."
        }
    }

    if len(attributeErrors) > 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": attributeErrors})
        return
    }

    // Xử lý default value cho Display
    if param.Display == nil {
        defaultDisplay := true
        param.Display = &defaultDisplay
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "create Product successfully",
        "data":    param,
    })
}
```

**Kết quả sau khi sửa — thông báo rõ ràng hơn:**

```json
{
    "error": {
        "product_attribute[1].attribute_name":  "AttributeName là bắt buộc.",
        "product_attribute[1].attribute_value": "AttributeValue là bắt buộc."
    }
}
```

Client giờ biết chính xác phần tử **thứ 1** (index 1) trong mảng bị thiếu dữ liệu.

---

## Validate Map và tự kiểm tra key

### Tag `dive` với Map

Tương tự slice, `dive` với map bảo validator đi vào kiểm tra **value** của từng entry:

```go
// gt=0: map phải có ít nhất 1 entry
// dive: với MỖI value trong map, kiểm tra các tag binding bên trong ProductInfo
// Lưu ý: validator chỉ kiểm tra được VALUE, không kiểm tra được KEY của map
ProductInfo map[string]ProductInfo `json:"product_info" binding:"required,gt=0,dive"`

type ProductInfo struct {
    InfoKey   string `json:"info_key"   binding:"required"`
    InfoValue string `json:"info_value" binding:"required"`
}
```

### Vấn đề: validator không kiểm tra được key của map

Validator có thể kiểm tra **value** của map qua `dive`, nhưng **không thể validate key**. Nếu ta muốn key phải là UUID hợp lệ, phải tự kiểm tra thủ công:

```go
func (obj *Product) CreateProduct(c *gin.Context) {
    var param Product
    if err := c.ShouldBindBodyWithJSON(&param); err != nil {
        c.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
        return
    }

    // Xử lý default value
    if param.Display == nil {
        defaultDisplay := true
        param.Display = &defaultDisplay
    }

    // Tự kiểm tra KEY của ProductInfo phải là UUID hợp lệ
    // Validator không hỗ trợ validate map key → phải làm thủ công
    for key := range param.ProductInfo {
        if _, err := uuid.Parse(key); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Key '" + key + "' trong product_info không hợp lệ. Key phải là UUID.",
            })
            return
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "create Product successfully",
        "data":    param,
    })
}
```

**Kiểm thử với key là UUID hợp lệ:**

```json
POST /api/v2/products
Body:
{
    "id": 15,
    "name": "Ship qua 2 quả tên lửa",
    "price": 10000,
    "product_image": { "name": "hinh1", "link": ".../hinh1.jpg" },
    "tags": ["1", "2", "3", "4"],
    "product_attribute": [{ "attribute_name": "color", "attribute_value": "white" }],
    "product_info": {
        "9a463ff3-17a3-436b-9a7a-3f864282f9fc": {
            "info_key": "dbsql",
            "info_value": "123"
        }
    }
}

Output: 200 OK ✅
```

**Kiểm thử với key không phải UUID:**

```json
"product_info": { "abc": { "info_key": "dbsql", "info_value": "123" } }

Output:
{
    "error": "Key 'abc' trong product_info không hợp lệ. Key phải là UUID."
}
```

---

## Toàn bộ struct `Product` sau khi hoàn thiện

```go
// Product là struct đại diện cho body request khi tạo sản phẩm mới
type Product struct {
    ID              int                        `json:"id"`
    Name            string                     `json:"name"              binding:"required,min=3,max=100"`
    Price           int                        `json:"price"             binding:"required,min_int=1000,max_int=100000000"`
    Display         *bool                      `json:"display"           binding:"omitempty"`             // con trỏ để phân biệt nil vs false
    ProductImage    ProductImage               `json:"product_image"     binding:"required"`              // nested struct
    Tags            []string                   `json:"tags"              binding:"required,gt=3,lt=5"`   // slice: 4 phần tử
    ProductAttribute []ProductAttribute        `json:"product_attribute" binding:"required,gt=0,dive"`   // slice of struct
    ProductInfo     map[string]ProductInfo     `json:"product_info"      binding:"required,gt=0,dive"`   // map, key validate thủ công
}

type ProductImage struct {
    ImageName string `json:"name" binding:"required,min=3,max=100"`
    ImageLink string `json:"link" binding:"required,file_extension"` // custom rule
}

type ProductAttribute struct {
    AttributeName  string `json:"attribute_name"  binding:"required"`
    AttributeValue string `json:"attribute_value" binding:"required"`
}

type ProductInfo struct {
    InfoKey   string `json:"info_key"   binding:"required"`
    InfoValue string `json:"info_value" binding:"required"`
}
```

---

## Tổng kết

### Các hàm bind theo nguồn dữ liệu

|Nguồn|Hàm|Tag struct|
|---|---|---|
|Path param|`ShouldBindUri`|`uri:"key"`|
|Query string|`ShouldBindQuery`|`form:"key"`|
|JSON body|`ShouldBindBodyWithJSON`|`json:"key"`|
|Raw body|`GetRawData()`|(không dùng tag)|

### Tóm tắt các kỹ thuật trong bài

|Kỹ thuật|Cách dùng|Ghi chú|
|---|---|---|
|**Nested struct**|Khai báo trường kiểu struct, thêm `binding:"required"`|Validator tự đệ quy vào bên trong|
|**`min_int` / `max_int`**|Custom rule, dùng `strconv.ParseInt` + `fl.Field().Int()`|Rõ ràng hơn `gte`/`lte` cho số nguyên|
|**`*bool` default value**|Dùng con trỏ, kiểm tra `nil` rồi gán default|Phân biệt "không gửi" với "gửi false"|
|**`file_extension`**|Custom rule, dùng `filepath.Ext()` + `strings.TrimPrefix()`|Hỗ trợ param hoặc dùng danh sách mặc định|
|**`dive` với slice**|Thêm `dive` sau rule của slice|Bắt buộc phải có để validate từng phần tử|
|**`dive` với map**|Thêm `dive` sau rule của map|Chỉ validate value, không validate key|
|**Validate map key**|Tự duyệt `for key := range map` và kiểm tra|Validator không hỗ trợ, phải làm thủ công|
|**Lỗi slice có index**|Validate thủ công với `fmt.Sprintf("[%d].field", i)`|Thông báo lỗi rõ ràng hơn cho client|

### Những điểm quan trọng cần nhớ

1. **`dive` là bắt buộc** khi muốn validator kiểm tra nội dung từng phần tử trong slice/map — thiếu `dive` thì chỉ kiểm tra độ dài của collection.
2. **Con trỏ (`*bool`, `*int`...)** cho phép phân biệt "trường không được gửi" (`nil`) với "trường có giá trị zero" (`false`, `0`).
3. **Validator chỉ validate value của map, không validate key** — phải tự viết vòng lặp kiểm tra key.
4. **`strconv.ParseInt(str, base, bitSize)`** — base 10 là thập phân, bitSize 64 an toàn nhất cho mọi nền tảng.
5. **`filepath.Ext()`** luôn trả về phần mở rộng kèm dấu chấm (`.jpg`) — dùng `strings.TrimPrefix` để bỏ dấu chấm khi cần so sánh.