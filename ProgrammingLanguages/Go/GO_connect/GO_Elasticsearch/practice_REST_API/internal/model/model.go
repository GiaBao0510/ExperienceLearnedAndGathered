package model

// Kiểu dữ liệu đại diện cho một sản phẩm
type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `jaon:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	InStock     bool    `json:"in_stock"`
}

// Hàm này chứa kết quả trả về từ É khi tìm kiếm
type SearchResult struct {
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		Hits []struct {
			ID     string  `json:"_id"`
			Source Product `json:"_source"`
		}
	} `json:"hits"`
}

// SearchRequest đại diện cho request tìm kếm từ client gửi lên
type SearchRequest struct {
	Query string `json:"query"`
	Field string `json:"field"`
}

// API Response đại diện cho cấu trúc phản hồi api cho client
type APIResponse struct {
	Success bool  	`json:"status"`
	Message string 	`json:"message"`
	Data interface{} `json:"data,omitempty"`
}