
Đầu tiên tải gói này, thông qua câu lệnh sau:

```shell
> go get github.com/go-playground/validator/v10
```

Đến với nội dung bài này thì sẽ thực hiện trên "/api/v2/users". Sau đó ta dựa trên gói trên để xác thực điều kiện lại api: "/api/v2/users/:id", để kiểm tra id có hợp lệ hay không và phải là số nguyên dương

Ví dụ:
- Tại tệp tin `Router_Group\main.go`:
```go
v2 := r.Group("/api/v2")
	{
		user := v2.Group("/users")
		{
			user.GET("/", userHandler_v2.GetUsers)
			user.GET("/:id", userHandler_v2.GetUserByID)
			user.POST("/", userHandler_v2.CreateUser)
			user.PUT("/:id", userHandler_v2.UpdateUser)
			user.DELETE("/:id", userHandler_v2.DeleteUser)
		}
	}
```


```go
// Tạo struct để kiểm tra đầu vào của GetUserByID
// Điều kiện ID > 0, và nó phải là kiểu int mới sử dụng được. Ngược lại thì không
type GetUserByID_V2_Param struct{
	ID int `uri:"id" binding:"gt=0"`
}

// Lấy thông tin người dùng dựa trên ID
func (obj *User) GetUserByID(ctx *gin.Context) {

	var params GetUserByID_V2_Param
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}


	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get user by ID successfully",
		"data":    nil,
	})
}
```

Ví dụ api: `http://localhost:8080/api/v2/users/10`
output: 
```json
{
    "data": null,
    "message": "Get user by ID successfully"
}
```

Ví dụ api: `http://localhost:8080/api/v2/users/-10`
output: 
```json
{
    "error": "Key: 'GetUserByID_V2_Param.ID' Error:Field validation for 'ID' failed on the 'gt' tag"
}
```

Ví dụ api: `http://localhost:8080/api/v2/users/a`
output: 
```json
{
    "error": "strconv.ParseInt: parsing \"a\": invalid syntax"
}
```

---
Tại tệp tin `Router_Group\utils\validation.go` chúng ta sẽ bổ sung thêm một hàm `HandleValidationErrors`

```go
// Viết hàm xử lý lỗi
func HandleValidationErrors (Err error) gin.H {

	//Kiểm tra lỗi này có phải thuộc package validator hay không
	if validationErr ,ok := Err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)

		for _, e := range validationErr{

			switch e.Tag() {
			case "gt":
				errors[e.Field()] = e.Field() + " phải lớn hơn giá trị tối thiểu."
			}
			log.Printf("Validation error on field '%s': %s, Tag: %+v", e.Field(), e.Error(), e.Tag())
		}

		return gin.H{ "error": errors }
	}

	return gin.H{ "error": "Yêu cầu không hợp lệ: " + Err.Error() }
}
```

Chỉnh sửa lại một dung file để kết nối lại file `Router_Group\internal\api\v2\handler\user.go` để nó dễ hiểu hơn

```go
// Lấy thông tin người dùng dựa trên ID
func (obj *User) GetUserByID(ctx *gin.Context) {

	var params GetUserByID_V2_Param
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}


	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get user by ID successfully",
		"user_id":    params.ID,
	})
}
```

Tương tự như vậy ta thêm phần chỉnh sửa ở UUID

tại hàm `Router_Group\main.go`, ta thêm đoạn code như sau:

```go 
v2 := r.Group("/api/v2")
{
	user := v2.Group("/users")
	{
		user.GET("/", userHandler_v2.GetUsers)
		user.GET("/:id", userHandler_v2.GetUserByID)
		user.GET("/uuid/:uuid", userHandler_v2.GetUserBy_UUID)	// Add
		user.POST("/", userHandler_v2.CreateUser)
		user.PUT("/:id", userHandler_v2.UpdateUser)
		user.DELETE("/:id", userHandler_v2.DeleteUser)
	}
}
```

Ở tập tin `Router_Group\internal\api\v2\handler\user.go`, ta bổ sung thêm hàm như sau: 
```go
type GetUserByUUID_Param struct{
	uuid string `uri:"uuid" binding:"uuid"`
}

// Lấy thông tin người dùng dựa trên UUID
func (obj *User) GetUserBy_UUID(ctx *gin.Context) {

	var params GetUserByUUID_Param
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get user by UUID successfully",
		"UUID":    params.uuid,
	})
}
```

Cũng tại tệp tin `Router_Group\utils\validation.go` tùy chỉnh lại câu lệnh đôi chút:
```go
// Viết hàm xử lý lỗi
func HandleValidationErrors (Err error) gin.H {

	//Kiểm tra lỗi này có phải thuộc package validator hay không
	if validationErr ,ok := Err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)

		for _, e := range validationErr{

			switch e.Tag() {
				case "gt":
					errors[e.Field()] = e.Field() + " phải lớn hơn giá trị tối thiểu."
				
				case "uuid":
					errors[e.Field()] = e.Field() + " phải là một UUID hợp lệ."
			}
			log.Printf("Validation error on field '%s': %s, Tag: %+v", e.Field(), e.Error(), e.Tag())
		}

		return gin.H{ "error": errors }
	}

	return gin.H{ "error": "Yêu cầu không hợp lệ: " + Err.Error() }
}
```

kiểm tra và kết quả:

api: `http://localhost:8080/api/v2/users/uuid/test`
output:
```json
{
    "error": {
        "UUID": "UUID phải là một UUID hợp lệ."
    }
}
```

api: `http://localhost:8080/api/v2/users/uuid/123`
output:
```json
{
    "error": {
        "UUID": "UUID phải là một UUID hợp lệ."
    }
}
```

api: `http://localhost:8080/api/v2/users/uuid/9a463ff3-17a3-436b-9a7a-3f864282f9fc`
output:
```json
{
    "UUID": "9a463ff3-17a3-436b-9a7a-3f864282f9fc",
    "message": "Get user by UUID successfully"
}
```

---
**Custom Validation** 

Ở phần này cúng ta sẽ tự custom lại nếu như trong package Validator nó chưa đáp ưng theo ý muốn chúng ta
Bây giờ chúng ta tạo tệp tin mới như sau `Router_Group\internal\api\v2\handler\product.go`:

Tại hàm `Router_Group\main.go` chúng ta cập nhập thêm một số câu lệnh như sau: 

```go
v2 := r.Group("/api/v2")
	{
		user := v2.Group("/users")
		{
			user.GET("/", userHandler_v2.GetUsers)
			user.GET("/:id", userHandler_v2.GetUserByID)
			user.GET("/uuid/:uuid", userHandler_v2.GetUserBy_UUID)	// Add
			user.POST("/", userHandler_v2.CreateUser)
			user.PUT("/:id", userHandler_v2.UpdateUser)
			user.DELETE("/:id", userHandler_v2.DeleteUser)
		}

		product := v2.Group("/products")
		{
			product.GET("/", productHandler_v2.GetProducts)
			product.GET("/:slug", productHandler_v2.GetProductBySlug)
			product.POST("/", productHandler_v2.CreateProduct)
			product.PUT("/:id", productHandler_v2.UpdateProduct)
			product.DELETE("/:id", productHandler_v2.DeleteProduct)
		}
	}
```

tại tệp tin `Router_Group\utils\validation.go` thay đổi đôi chút và thêm hàm mới `RegisterValidationError` như sau:

```go
// Viết hàm xử lý lỗi
func HandleValidationErrors(Err error) gin.H {

	//Kiểm tra lỗi này có phải thuộc package validator hay không
	if validationErr, ok := Err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)

		for _, e := range validationErr {

			switch e.Tag() {
			case "gt":
				errors[e.Field()] = e.Field() + " phải lớn hơn giá trị tối thiểu."
			case "uuid":
				errors[e.Field()] = e.Field() + " phải là một UUID hợp lệ."
			case "slug":
				errors[e.Field()] = e.Field() + " Chỉ có thể chứa: chữ thường, số, dấu gạch ngang (-) hoặc dấu chấm (.)"
			case "max":
				errors[e.Field()] = e.Field() + " Độ dài tối đa là " + e.Param() + " ký tự."
			case "min":
				errors[e.Field()] = e.Field() + " Độ dài tối thiểu là " + e.Param() + " ký tự."
			case "oneof":
                allowedValues := strings.Join(strings.Split(e.Param(), " "), ",")
                errors[e.Field()] = e.Field() + " phải là một trong các giá trị sau: " + allowedValues + "."
			}
			log.Printf("Validation error on field '%s': %s, Tag: %+v", e.Field(), e.Error(), e.Tag())
		}

		return gin.H{"error": errors}
	}

	return gin.H{"error": "Yêu cầu không hợp lệ: " + Err.Error()}
}

// Tự tạo một hàm để kiểm tra lỗi và trả về lỗi nếu có
func RegisterValidationError() error {

	//Kiểm tra xem kiểu đầu vào có thuộc kiểu validator không
	v, ok := binding.Validator.Engine().(*validator.Validate)

	if !ok {
		return fmt.Errorf("Failed to register validation: could not get validator engine")
	}

	var slugRegex = regexp.MustCompile(`^[a-z0-9]+(?:[-.][a-z0-9]+)*$`)

	//Đăng ký hàm kiểm tra regex cho tag "slug", và kiểm tra field có phải là string hay không, nếu không thì trả về lỗi
	v.RegisterValidation("slug", func(fl validator.FieldLevel) bool {
		return slugRegex.MatchString(fl.Field().String()) // Kiểm tra xem field này có khớp với regex hay không
	})

	return nil
}
```

Sau đó chúng ta sẽ gọi hàm `RegisterValidationError` này ra file main.go để cho các file khác có thể sử dụng chung, thì chúng ta tùy chỉnh ở tệp tin main.go như sau:

```go
func main() {
	r := gin.Default()

	if err := utils.RegisterValidationError(); err != nil {
		panic(err)
	}
	
	//...
}
```

Tạo tệp tin `Router_Group\internal\api\v2\handler\product.go`

```go
package handler

import (
	"fmt"
	"net/http"
	"router-group/utils"

	"github.com/gin-gonic/gin"
)

//Tạo list Product
type Product struct{
	ID int `json:"id"`
	ProductName string `json:"name"`
	Price float64 `json:"price"`
}
 

//Constructor để tạo instance của struct Product
func NewProduct() *Product{
	return  &Product{}
} 
//...
// Tạo cấu trúc kiểu tra đầu vào của GetProductBySlug
type GetProductBySlug_Param struct{
	slug string `uri:"slug" binding:"slug, min=5, max=100"`
}

//Lấy thông tin sản phẩm dựa trên Slug
func (obj *Product) GetProductBySlug(ctx *gin.Context){
	
	var params GetProductBySlug_Param
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"message": "Product not found",
		"data": params.slug,
	})
}

//...
```

Tiếp theo ta tạo tệp tin `Router_Group\internal\api\v2\handler\category.go`:

```go
type Category struct {
}

func NewCategoryHandler() *Category {
	return &Category{}
}

type GetCategoryByCategoriest_Param struct {
	category string `uri:"" binding:"oneof=php python golang java"`
}

// Lấy Categories theo mục
func (obj *Category) GetCategoryByCategories(c *gin.Context) {

	var params GetCategoryByCategoriest_Param
	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": params.category,
	})
}
```







