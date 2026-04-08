package handler

import (
	"encoding/json"
	"es-go-crud/internal/model"
	"es-go-crud/internal/repository"
	"net/http"
	"strings"
)

//Chứa reference đến repository để gọi CRUD
type ProductHandler struct {
	repo *repository.ProductRepository
}

//Hàm khởi tạo, nhận repository đã được tạo sẵn
func NewProductHandler(repo *repository.ProductRepository) *ProductHandler {
	return &ProductHandler{repo: repo}
}

//responseJSON: là hàm tiện ích để trả về JSON response với status code và payload
func responseJSON(w http.ResponseWriter, code int, payload interface{}){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

// CREATE: Xử lý HTTP POST /products
func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var product model.Product

	// Bước 1: Giải mã JSON body thành struct Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		responseJSON(w, http.StatusBadRequest, model.APIResponse{
			Success: false, Message: "Đầu vào không hợp lệ",
		})
		return
	}

	// Bước 2 Gọi repository để tạo mới
	id, err := h.repo.Create(r.Context(), &product)
	if err != nil {
		responseJSON(w, http.StatusInternalServerError, model.APIResponse{
			Success: false, Message: "Lỗi khi tạo sản phẩm",
		})
		return
	}

	// Tạo response thành công
	responseJSON(w, http.StatusCreated, model.APIResponse{
		Success: true, 
		Message: "Sản phẩm đã được tạo", 
		Data: map[string]string{"id": id},
	})
}

// SEARCH: Xử lý HTTP POST /products/search
func (h *ProductHandler) Search(w http.ResponseWriter, r *http.Request) {
	var req model.SearchRequest
	
	// Đầu tiên Giải mã JSON body thành struct SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responseJSON(w, http.StatusBadRequest, model.APIResponse{
			Success: false, Message: "Đầu vào không hợp lệ",
		})
		return
	}

	products, err := h.repo.Search(r.Context(), req.Field, req.Query)
	if err != nil {
		responseJSON(w, http.StatusInternalServerError, model.APIResponse{
			Success: false, Message: "Lỗi khi tìm kiếm sản phẩm",
		})
		return
	}

	responseJSON(w, http.StatusOK, model.APIResponse{
		Success: true,
		Message: "Tìm kiếm thành công",
		Data: products,
	})
}

// GET ALL xử lý HTTP GET /products - Lấy tất cả sản phẩm
func (h*ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := h.repo.GetAll(r.Context())
	if err != nil {
		responseJSON(w, http.StatusInternalServerError, model.APIResponse{
			Success: false, Message: "Lỗi khi lấy sản phẩm",
		})
		return
	}
	
	responseJSON(w, http.StatusOK, model.APIResponse{
		Success: true,
		Message: "Lấy sản phẩm thành công",
		Data: products,
	})
}

// UPDATE xử lý HTTP PUT /products/{id} - Cập nhật sản phẩm theo ID
func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	// trích xuất ID từ URL path
	id := strings.TrimPrefix(r.URL.Path, "/api/products/")
	if id == "" {
		responseJSON(w, http.StatusBadRequest, model.APIResponse{
			Success: false, Message: "ID sản phẩm không được để trống",
		})
		return 
	}

	var update map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		responseJSON(w, http.StatusBadRequest, model.APIResponse{
			Success: false, Message: "Đầu vào không hợp lệ",
		})
		return
	}

	if err := h.repo.Update(r.Context(), id, update); err != nil {
		responseJSON(w, http.StatusInternalServerError, model.APIResponse{
			Success: false,
			Message: "Lỗi khi cập nhật sản phẩm",
			Data: err.Error(),			
		})
		return
	}

	responseJSON(w, http.StatusOK, model.APIResponse{
		Success: true,
		Message: "Sản phẩm đã được cập nhật",
		Data: nil,
	})
}

// DELETE xử lý HTTP DELETE /products/{id} - Xóa sản phẩm theo ID
func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// trích xuất ID từ URL path
	id := strings.TrimPrefix(r.URL.Path, "/api/products/")
	if id == "" {
		responseJSON(w, http.StatusBadRequest, model.APIResponse{
			Success: false, Message: "ID sản phẩm không được để trống",
		})
		return 
	}

	if err := h.repo.Delete(r.Context(), id); err != nil {
		responseJSON(w, http.StatusInternalServerError, model.APIResponse{
			Success: false,
			Message: "Lỗi khi xóa sản phẩm",
		})
		return
	}

	responseJSON(w, http.StatusOK, model.APIResponse{
		Success: true,
		Message: "Sản phẩm đã được xóa",
		Data: nil,
	})
}