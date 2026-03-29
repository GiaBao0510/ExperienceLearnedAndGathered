## **Validate Body Request (Json) với gói Validator**

Phần này sẽ xác thực thông tin đầu vào từ phía Client gửi thông qua phương thức POST

Ví dụ ban đầu tại tệp tin `Router_Group\internal\api\v2\handler\product.go`:
```go
//...
//Thêm sản phẩm mới
func (obj *Product) CreateProduct(c *gin.Context) {
	
	body, err := c.GetRawData() 	//Lấy dữ liệu thô từ request body
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error reading request body: " + err.Error())
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "create Product successfully",
		"data": string(body),
	})
}
//...
```

_Output:_
```
POST: http://localhost:8080/api/v2/products
Body: Welcome my website

out put:
{
    "data": "Welcome my website",
    "message": "create Product successfully"
}
```

Nhưng tiếp theo là đa số client thường gửi dữ liệu dạng json, thì ta tùy chỉnh lại câu lệnh tại tệp tin `Router_Group\internal\api\v2\handler\product.go` như sau:

```go
//Tạo list Product
type Product struct{
	ID int `json:"id"`
	Name string `json:"name" binding:"required,min=3,max=100"`
	Price float64 `json:"price" binding:"required,gt=0"`
	Display bool `json:"display" binding:"omitempty"`
}
//...
//Thêm sản phẩm mới
func (obj *Product) CreateProduct(c *gin.Context) {
	
	var param Product
	if err := c.ShouldBindBodyWithJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "create Product successfully",
		"data": param,
	})
}
//...
```

_Output:_
```
POST: http://localhost:8080/api/v2/products
Body: {
    "ID": 15,
    "name": "Ship qua 2 quả tên lửa",
    "Price": 100000.001
}

out put:
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

Tiếp theo giả sử ta thêm kiểu struct vào trong struct `Product` để xem nó kiểm tra điều kiện như thế nào:
```go
//Tạo list Product
type Product struct{
	ID int `json:"id"`
	Name string `json:"name" binding:"required,min=3,max=100"`
	Price float64 `json:"price" binding:"required,gt=0"`
	Display bool `json:"display" binding:"omitempty"`
	ProductImage ProductImage `json:"product_image" binding:"required"`
}

type ProductImage struct {
	ImageName string `json:"name" binding:"required,min=3,max=100"`
	ImageLink string `json:"link" binding:"required,url"`
}
//...
//Thêm sản phẩm mới
func (obj *Product) CreateProduct(c *gin.Context) {
	
	var param Product
	if err := c.ShouldBindBodyWithJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "create Product successfully",
		"data": param,
	})
}
//...
```

_Output:_
```
POST: http://localhost:8080/api/v2/products
Body: {
    "ID": 15,
    "name": "Ship qua 2 quả tên lửa",
    "Price": 100000.001,
    "product_image":{
        "name": "hinh1",
        "link": "http://localhost:8080/api/v2/products/15/hinh1"
    }
}

out put:
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

---
## **Validate Body Request với gói Validator - Xử lý giá trị số, Default Value & File Extension**

Ở tại tệp tin `Router_Group\utils\validation.go`, chúng ta cập nhập thêm đoạn code:

```go
func HandleValidationErrors(Err error) gin.H {

	//Kiểm tra lỗi này có phải thuộc package validator hay không
	if validationErr, ok := Err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)

		for _, e := range validationErr {

			switch e.Tag() {
			//...
			case "min_int":
				errors[e.Field()] = e.Field() + " phải lớn hơn hoặc bằng " + e.Param() + "."
			case "max_int":
				errors[e.Field()] = e.Field() + " phải nhỏ hơn hoặc bằng " + e.Param() + "."
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

	//...

	//Đăng ký hàm kiểm tra giá trị số nguyên dương cho tag "min_int", và kiểm tra field có phải là string hay không, nếu không thì trả về lỗi
	v.RegisterValidation("min_int", func(fl validator.FieldLevel) bool {
		minStr := fl.Param()
		minVal, err := strconv.ParseInt(
			minStr, 
			10,        // 10: hệ thập phân; 16: hệ thập lục phân; 8: hệ bát phân; 0: tự động nhận diện
			64, 		// 0: int, 8: int8, 16: int16, 32: int32, 64: int64
		)
		if err != nil {
			return false
		}

		return fl.Field().Int() >= minVal	// Kiểm tra xem field này có lớn hoặc bằng giá trị tối thiểu hay không
	})

	v.RegisterValidation("max_int", func(fl validator.FieldLevel) bool {
		maxStr := fl.Param()
		maxVal, err := strconv.ParseInt(
			maxStr, 
			10,        // 10: hệ thập phân; 16: hệ thập lục phân; 8: hệ bát phân; 0: tự động nhận diện
			64, 		// 0: int, 8: int8, 16: int16, 32: int32, 64: int64
		)
		if err != nil {
			return false
		}

		return fl.Field().Int() <= maxVal
	})

	return nil
}
```

Ở tệp tin `Router_Group\internal\api\v2\handler\product.go`, thay đổi chút:
```go
type Product struct{
	ID int `json:"id"`
	Name string `json:"name" binding:"required,min=3,max=100"`
	Price int `json:"price" binding:"required,min_int=1000,max_int=100000000"`
	Display bool `json:"display" binding:"omitempty"`
	ProductImage ProductImage `json:"product_image" binding:"required"`
}
```

Ở đây ta thiết lập trường hợp Default value:
Ở tệp tin  `Router_Group\internal\api\v2\handler\product.go`, ta thay đổi chút ở phần Display:
```go
type Product struct{
	ID int `json:"id"`
	Name string `json:"name" binding:"required,min=3,max=100"`
	Price int `json:"price" binding:"required,min_int=1000,max_int=100000000"`
	Display *bool `json:"display" binding:"omitempty"`
	ProductImage ProductImage `json:"product_image" binding:"required"`
}

//...
func (obj *Product) CreateProduct(c *gin.Context) {
	
	var param Product
	if err := c.ShouldBindBodyWithJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	//Trường hợp giá trị Display không được cung cấp, mặc định là true
	if param.Display == nil {
		defaultDisplay := true
		param.Display = &defaultDisplay
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "create Product successfully",
		"data": param,
	})
}
//...
```


Ở đây ta thêm phần xác thực định dạng file đầu vào:
- Ở tệp tin `Router_Group\utils\validation.go`
```go
func HandleValidationErrors(Err error) gin.H {

	//Kiểm tra lỗi này có phải thuộc package validator hay không
	if validationErr, ok := Err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)

		for _, e := range validationErr {

			switch e.Tag() {
			//...
			case "file_extension":
				errors[e.Field()] = e.Field() + " phải có phần mở rộng hợp lệ " + strings.Join(allowedExtensions, ", ") + "."
			}
			log.Printf("Validation error on field '%s': %s, Tag: %+v", e.Field(), e.Error(), e.Tag())
		}

		return gin.H{"error": errors}
	}

	return gin.H{"error": "Yêu cầu không hợp lệ: " + Err.Error()}
}

var allowedExtensions = []string{".jpg", ".jpeg", ".png", ".gif"}
func RegisterValidationError() error {

	//Kiểm tra xem kiểu đầu vào có thuộc kiểu validator không
	v, ok := binding.Validator.Engine().(*validator.Validate)

	if !ok {
		return fmt.Errorf("Failed to register validation: could not get validator engine")
	}

	//...

	// File extension validation: .jpg, .jpeg, .png, .gif
	v.RegisterValidation("file_extension", func(fl validator.FieldLevel) bool {
		fileName := fl.Field().String()
		allowStr := fl.Param()

		// ✅ Nếu fileName rỗng, skip validation
		if fileName == "" {
			return true
		}

		ext := strings.TrimPrefix(strings.ToLower(filepath.Ext(fileName)), ".")

		// ✅ Nếu allowStr có param, phân tách và dùng, nếu không dùng default
		if allowStr != "" {
			allowedExts := strings.Split(allowStr, " ")
			for _, allow := range allowedExts {
				if ext == strings.ToLower(allow) {
					return true
				}
			}
			return false
		}

		// ✅ Dùng default allowedExtensions
		for _, allowExists := range allowedExtensions {
			if ext == strings.TrimPrefix(strings.ToLower(allowExists), ".") {
				return true
			}
		}

		return false
	})

	return nil
}
```

Tại tệp tin   `Router_Group\internal\api\v2\handler\product.go`, ta thay đổi chút ở phần struct ProductImage:
```go
type ProductImage struct {
    ImageName string `json:"name" binding:"required,min=3,max=100"`
    ImageLink string `json:"link" binding:"required,file_extension"`
}
```
