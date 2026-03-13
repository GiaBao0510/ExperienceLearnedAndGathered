package handler

import (
	"fmt"
	"net/http"

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
}

//Lấy danh sách Product
func (obj *Product) GetProducts(c *gin.Context){
	fmt.Println("Lấy danh sách sản phẩm")
	c.JSON(http.StatusOK, gin.H{
		"message": "list Product",
		"data": Products,
	})
}

//Lấy thông tin sản phẩm dựa trên ID
func (obj *Product) GetProductByID(c *gin.Context){
	
	id := c.Param("id")
	for _,Product := range Products{
		if fmt.Sprintf("%v", Product.ID) == id{
			c.JSON(http.StatusOK, gin.H{
				"message": "Product found",
				"data": Product,
			})
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Product not found",
		"data": nil,
	})
}

//Thêm sản phẩm mới
func (obj *Product) CreateProduct(c *gin.Context) {
	fmt.Println("Nhập thông tin sản phẩm")
	
	//Đọc dữ liệu từ request body
	var newProduct Product
	if err := c.ShouldBindJSON(&newProduct); err != nil{		//Nếu có lỗi khi đọc dữ liệu, trả về lỗi
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
		})
		return
	}

	//Tạo ID mới cho Product
	newProduct.ID = len(Products) + 1

	//Thêm Product mới vào slice
	Products = append(Products, newProduct)
	
	c.JSON(http.StatusOK, gin.H{
		"message": "create Product successfully",
	})
}


//Cập nhật thông tin sản phẩm
func (obj *Product) UpdateProduct(c *gin.Context){

	id := c.Param("id")

	//Đọc dữ liệu từ request body
	var updateProduct Product
	if err := c.ShouldBindJSON(&updateProduct); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
			"error": err.Error(),
		})
		return
	}

	//Tìm ID của Product rồi cập nhật thông tin
	for idx, Product := range Products{
		if fmt.Sprintf("%v",Product.ID) == id{
			
			//Cập nhật thông tin Product
			Products[idx].ProductName = updateProduct.ProductName
			Products[idx].Price = updateProduct.Price

			//Trả về thông báo cập nhật thành công
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

//Xóa sản phẩm
func (obj *Product) DeleteProduct(c *gin.Context){

	id := c.Param("id")

	for index, Product := range Products{
		if fmt.Sprintf("%v", Product.ID) == id{
			Products = append(Products[:index], Products[index+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"message": "delete Product successfully",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Product not found",
		"data": nil,
	})
}