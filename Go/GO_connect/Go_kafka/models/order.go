package models

import "time"

// Kiểu dữ liệu đơn hàng
type Order struct {
	ID           string      `json:"id"`
	CustomerID   string      `json:"customer_id"`
	Items        []OrderItem `json:"items"`
	TotalAmount  float64     `json:"total_amount"`
	Status       string      `json:"status"`
	CreatedAt    time.Time   `json:"created_at"`
}

// Kiểu dữ liệu chi tiết đơn hàng
type OrderItem struct {
	ProductID string  `json:"product_id"`
	Name      string  `json:"name"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}