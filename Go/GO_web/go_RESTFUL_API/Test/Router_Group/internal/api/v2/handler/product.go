package handler

import (
	"fmt"
	"net/http"
	"router-group/utils"

	"github.com/gin-gonic/gin"
)

// ---- Tạo group quản lý Product v1  -----

//Tạo list Product
type Product struct{
	ID int `json:"id"`
	Name string `json:"name" binding:"required,min=3,max=100"`
	Price int `json:"price" binding:"required,min_int=1000,max_int=100000000"`
	Display *bool `json:"display" binding:"omitempty"`
	ProductImage ProductImage `json:"product_image" binding:"required"`
}

type ProductImage struct {
	ImageName string `json:"name" binding:"required,min=3,max=100"`
	ImageLink string `json:"link" binding:"required,file_extension"`
}
 

//Constructor để tạo instance của struct Product
func NewProduct() *Product{
	return  &Product{}
} 

//Tạo slice để lưu trữ Product
var Products []Product = []Product{
	{ID: 1, Name: "Iphone 14 Pro Max", Price: 30000000, Display: nil, ProductImage: ProductImage{ImageName: "iphone14_pro_max.jpg", ImageLink: "https://example.com/iphone14_pro_max.jpg"}},
	{ID: 2, Name: "Samsung Galaxy S23 Ultra", Price: 25000000, Display: nil, ProductImage: ProductImage{ImageName: "samsung_galaxy_s23_ultra.jpg", ImageLink: "https://example.com/samsung_galaxy_s23_ultra.jpg"}},
	{ID: 3, Name: "Xiaomi Mi 12 Pro", Price: 20000000, Display: nil, ProductImage: ProductImage{ImageName: "xiaomi_mi_12_pro.jpg", ImageLink: "https://example.com/xiaomi_mi_12_pro.jpg"}},
	{ID: 4, Name: "Oppo Find X5 Pro", Price: 22000000, Display: nil, ProductImage: ProductImage{ImageName: "oppo_find_x5_pro.jpg", ImageLink: "https://example.com/oppo_find_x5_pro.jpg"}},
	{ID: 5, Name: "Vivo X80 Pro", Price: 21000000, Display: nil, ProductImage: ProductImage{ImageName: "vivo_x80_pro.jpg", ImageLink: "https://example.com/vivo_x80_pro.jpg"}},
	{ID: 6, Name: "Realme GT 2 Pro", Price: 18000000, Display: nil, ProductImage: ProductImage{ImageName: "realme_gt_2_pro.jpg", ImageLink: "https://example.com/realme_gt_2_pro.jpg"}},
	{ID: 7, Name: "Iphone 15 Pro Max", Price: 30000000, Display: nil, ProductImage: ProductImage{ImageName: "iphone15_pro_max.jpg", ImageLink: "https://example.com/iphone15_pro_max.jpg"}},
}

//Lấy danh sách Product
func (obj *Product) GetProducts(c *gin.Context){
	fmt.Println("Lấy danh sách sản phẩm")
	c.JSON(http.StatusOK, gin.H{
		"message": "list Product",
		"data": Products,
	})
}

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

type SearchProducts struct{
	Search string `form:"search" binding:"required,search,min=3,max=100"` 
	Limit int `form:"limit" binding:"omitempty,gte=1,lte=100"`
	Email string `form:"email" binding:"email"`
	Date string `form:"date" binding:"omitempty,datetime=2006-01-02"`
}

//Tìm kiếm sản phẩm dựa trên tên sản phẩm
func (obj *Product) SearchProducts(ctx *gin.Context){
	
	var params SearchProducts
	if err := ctx.ShouldBindQuery(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	if params.Limit == 0 {
		params.Limit = 1
	}

	if params.Email == ""{
		params.Email = "No email provided"
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully search Product",
		"Search": params.Search,
		"Limit": params.Limit,
		"Email": params.Email,
		"Date": params.Date,
	})
}

//Thêm sản phẩm mới
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


//Cập nhật thông tin sản phẩm
func (obj *Product) UpdateProduct(c *gin.Context){
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Product not found",
	})
}

//Xóa sản phẩm
func (obj *Product) DeleteProduct(c *gin.Context){


	c.JSON(http.StatusNotFound, gin.H{
		"message": "Product not found",
		"data": nil,
	})
}