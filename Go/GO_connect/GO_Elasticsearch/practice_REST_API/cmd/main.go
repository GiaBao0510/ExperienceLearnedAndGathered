package main

import (
	"es-go-crud/internal/config"
	"es-go-crud/internal/handler"
	"es-go-crud/internal/infrastructure/elasticsearch"
	"es-go-crud/internal/repository"
	"fmt"
	"log"
	"net/http"
)

func main() {
	
	// Đầu tiên: đọc cấu hình từ file .env
	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("Lỗi khi đọc cấu hình: %v\n", err)
		return
	}
	log.Printf("Cấu hình đã được tải: %+v", cfg)

	//Bước 2: khởi tạo Elasticsearch client
	esClient, err := elasticsearch.NewESClient(cfg.ES_Address)
	if err != nil {
		log.Fatalf("Không thể khởi tạo Elasticsearch client: %v", err)
	}
	log.Println("Kết nối Elasticsearch thành công")

	// Bước 3: Inject esClient vào Repository
	repo := repository.NewProductRepository(esClient, cfg.ES_INDEX_PRODUCTS)

	// Khởi tạo Handler với repository đã được tạo
	handler := handler.NewProductHandler(repo)

	// Bước 5: đăng ký router
	mux := http.NewServeMux()

	//POST
	mux.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost{
			handler.Create(w, r)
		}else{
			http.Error(w, "Phương thức không được hỗ trợ", http.StatusMethodNotAllowed)
		}
	})

	//SEARCH
	mux.HandleFunc("/api/products/search", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost{
			handler.Search(w, r)
		}else{
			http.Error(w, "Phương thức không được hỗ trợ", http.StatusMethodNotAllowed)
		}
	})

	// GET ALL 
	mux.HandleFunc("/api/products/get-all", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet{
			handler.GetAll(w, r)
		}else{
			http.Error(w, "Phương thức không được hỗ trợ", http.StatusMethodNotAllowed)
		}
	})

	// DELETE và UPDATE tương tự, sẽ thêm sau
	// Dùng prefix /api/products/ (có dấu /) để bắt cả /api/products/{id}
	mux.HandleFunc("/api/products/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
			case http.MethodDelete:
				handler.Delete(w, r)
			case http.MethodPut:
				handler.Update(w, r)
			default:
				http.Error(w, "Phương thức không được hỗ trợ", http.StatusMethodNotAllowed)
		}
	} )

	// Bước 6: chạy server
	addr := fmt.Sprintf(":%s", cfg.APP_Port)

	log.Printf("Server đang chạy tại http://localhost%s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Lỗi khi chạy server: %v", err)
	}
}