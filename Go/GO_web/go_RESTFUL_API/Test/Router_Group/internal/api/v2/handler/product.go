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
	ProductName string `json:"name"`
	Price float64 `json:"price"`
}
 

//Constructor để tạo instance của struct Product
func NewProduct() *Product{
	return  &Product{}
} 

//Tạo slice để lưu trữ Product
var Products []Product = []Product{
	{ID: 1, ProductName: "Iphone 14 Pro Max", Price: 30000000},
	{ID: 2, ProductName: "Samsung Galaxy S23 Ultra", Price: 25000000},
	{ID: 3, ProductName: "Xiaomi Mi 12 Pro", Price: 20000000},
	{ID: 4, ProductName: "Oppo Find X5 Pro", Price: 22000000},
	{ID: 5, ProductName: "Vivo X80 Pro", Price: 21000000},
	{ID: 6, ProductName: "Realme GT 2 Pro", Price: 18000000},
	{ID: 7, ProductName: "Iphone 15 Pro Max", Price: 30000000},
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

//Tìm kiếm sản phẩm dựa trên tên sản phẩm
func (obj *Product) SearchProducts(c *gin.Context){
	
	c.JSON(http.StatusOK, gin.H{
		"message": "Search results",
		"data": "",
	})
}

//Thêm sản phẩm mới
func (obj *Product) CreateProduct(c *gin.Context) {
	
	
	c.JSON(http.StatusOK, gin.H{
		"message": "create Product successfully",
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