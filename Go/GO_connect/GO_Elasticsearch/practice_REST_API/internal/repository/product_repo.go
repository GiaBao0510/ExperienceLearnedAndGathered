package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"es-go-crud/internal/model"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type ProductRepository struct {
	client *elasticsearch.Client // Nhận từ bên ngoài, không tự khởi tạo
	index  string
}

// Hàm khởi tạo, nhận *elasticsearch.Client đã được tạo sẳn
func NewProductRepository(client *elasticsearch.Client, index string) *ProductRepository {
	return &ProductRepository{
		client: client,
		index:  index,
	}
}

// Thêm mới Document
func (r *ProductRepository) Create(ctx context.Context, product *model.Product) (string, error) {

	// Đầu tiên chuyển đổi struct Product thành JSON bytes
	data, err := json.Marshal(product)
	if err != nil {
		return "", fmt.Errorf("Lỗi khi chuyển đổi product thành JSON: %v", err)
	}

	// Bước 2: Gọi ES index API (tương tự HTTP POST/ index/_doc)
	// bytes.NewReader(data) để tạo io.Reader từ JSON bytes
	res, err := r.client.Index(
		r.index,               // Tên index đã được cấu hình
		bytes.NewReader(data), // Dữ liệu JSON dưới dạng io.Reader
	)

	if err != nil {
		return "", fmt.Errorf("Lỗi khi gọi ES Index API: %v", err)
	}
	defer res.Body.Close() // Đảm bảo đóng response body sau khi sử dụng

	// Bước 3: kiểm tra lỗi từ response của ES
	if res.IsError() {
		return "", fmt.Errorf("ES Index API trả về lỗi: %s", res.String())
	}

	// Bước 4: Phân tích kết quả trả về để lấy ID của document mới tạo
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("Lỗi khi phân tích kết quả ES: %v", err)
	}

	// ES trả về field "_id" chứa ID của document mới tạo
	id, _ := result["_id"].(string)
	return id, nil
}

// READ - Tim kiếm document (match query)
func (r *ProductRepository) Search(ctx context.Context, field, query string) ([]*model.Product, error) {
	// Bước 1: Xây dựng ES query DSL dưới dạng map
	// "Match query" để tìm kiếm full-text, cho phép gần đúng
	searchBody := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				field: query, // Tìm kiếm trên field nào và giá trị tìm kiếm là gì
			},
		},
	}

	// Bước 2: Chuyển đổi searchBody thành JSON bytes
	data, err := json.Marshal(searchBody)
	if err != nil {
		return nil, fmt.Errorf("Lỗi khi chuyển đổi searchBody thành JSON: %v", err)
	}

	// Bước 3: Gọi ES search API ( HTTP GET /index/_search với body)
	res, err := r.client.Search(
		r.client.Search.WithContext(ctx),
		r.client.Search.WithIndex(r.index), // Chỉ tiemf trong index đã được cấu hình
		r.client.Search.WithBody(bytes.NewReader(data)),
		r.client.Search.WithPretty(), // Tùy chọn để ES trả về kết quả dễ đọc hơn (dành cho debug)
	)
	if err != nil {
		return nil, fmt.Errorf("Lỗi khi gọi ES Search API: %v", err)
	}
	defer res.Body.Close()

	// Kiểm tra lỗi từ response của ES
	if res.IsError() {
		return nil, fmt.Errorf("ES Search API trả về lỗi: %s", res.String())
	}

	// Bước 4: Parse kết quả - ES thường trả về dạng hits.hits[]
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("Lỗi khi phân tích kết quả ES: %v", err)
	}

	// Trích xuất mảng hits
	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})

	var products []*model.Product
	for _, hit := range hits {
		hitMap := hit.(map[string]interface{})

		// Re-serialize _source (dữ liệu gốc của document) thành JSON bytes
		sourceBytes, _ := json.Marshal(hitMap["_source"])

		// Deserialize JSON bytes thành struct Product
		var product model.Product
		json.Unmarshal(sourceBytes, &product)

		// Gán ID của document (từ _id) vào struct Product
		product.ID = hitMap["_id"].(string)
		products = append(products, &product)
	}

	return products, nil
}

// GET ALL - Lấy tất cả document (match_all query)
func (r *ProductRepository) GetAll(ctx context.Context) ([]*model.Product, error) {
	
	// ES query để lấy tất cả document trong index
	searchBody := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}
	// Chuyển đổi searchBody thành JSON bytes
	data, err := json.Marshal(searchBody)
	if err != nil {
		return nil, fmt.Errorf("Lỗi khi chuyển đổi searchBody thành JSON: %v", err)
	}

	// Gọi ES search API
	res, err := r.client.Search(
		r.client.Search.WithContext(ctx),
		r.client.Search.WithIndex(r.index),
		r.client.Search.WithBody(bytes.NewReader(data)),
		r.client.Search.WithPretty(),
	)

	if err != nil {
		return nil, fmt.Errorf("Lỗi khi gọi ES Search API: %v", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("ES Search API trả về lỗi: %s", res.String())
	}

	// Parse kết quả tương tự như hàm Search
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("Lỗi khi phân tích kết quả ES: %v", err)
	}

	// Trích xuất mảng hits
	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})

	// Duyệt qua hits để chuyển đổi thành slice []*model.Product
	var products []*model.Product
	for _, hit := range hits {
		hitMap := hit.(map[string]interface{})

		// Re-serialize _source (dữ liệu gốc của document) thành JSON bytes
		sourceBytes, _ := json.Marshal(hitMap["_source"])

		// Deserialize JSON bytes thành struct Product
		var product model.Product
		json.Unmarshal(sourceBytes, &product)

		// Gán ID của document (từ _id) vào struct Product
		product.ID = hitMap["_id"].(string)
		products = append(products, &product)
	}

	return products, nil
}

// UPDATE - Cập nhật document (HTTP POST /index/_doc/{id})
func (r *ProductRepository) Update(ctx context.Context, id string, updates map[string]interface{}) error{
	// ES update API yêu cầu body có dạng {"doc": {field: value}}
	// Điêu này giúp ES chỉ cập nhật những field được chỉ định, không cần gửi toàn bộ document
	body := map[string]interface{}{"doc": updates}

	data, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("Lỗi khi chuyển đổi update body thành JSON: %v\n", err)
	}

	// esapi.UpdateRequest: Cấu trúc để gọi ES update API, tương tự HTTP POST /index/_doc/{id}/_update
	req := esapi.UpdateRequest{
		Index: r.index,
		DocumentID: id,
		Body: bytes.NewReader(data),
		Refresh: "true",
	}

	// req.Do: Thực thi request và nhận response từ ES
	res, err := req.Do(ctx, r.client)
	if err != nil {
		return fmt.Errorf("Lỗi khi gọi ES Update API: %v\n", err)
	}
	defer res.Body.Close()

	if res.IsError(){
		return fmt.Errorf("ES Update API trả về lỗi: %s\n", res.String())
	}

	return nil
}

// DELETE - Xóa document (HTTP DELETE /index/_doc/{id})
func (r *ProductRepository) Delete(ctx context.Context, id string) error {
	
	// ES delete API tương tự HTTP DELETE /index/_doc/{id}
	res, err := r.client.Delete(
		r.index,							// Tên index đã được cấu hình
		id,									// ID của document cần xóa
		r.client.Delete.WithContext(ctx),	// Truyền context để quản lý timeout, hủy bỏ request nếu cần
	)
	if err != nil{
		return fmt.Errorf("Lỗi khi gọi ES Delete API: %v\n", err)
	}
	defer res.Body.Close()

	if res.IsError(){
		return fmt.Errorf("ES Delete API trả về lỗi: %s\n", res.String())
	}
	return  nil
}