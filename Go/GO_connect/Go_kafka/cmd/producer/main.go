package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/GiaBao0510/kafka-go/config"
	kafkaclient "github.com/GiaBao0510/kafka-go/internal/kafka"
	"github.com/GiaBao0510/kafka-go/models"
)

// Giả lập tạo 5 đơn hàng và gửi lên Kafka
func main() {
	
	// Load cấu hình từ biến môi trường
	kafkaConfig := config.LoadKafkaConfig()
	log.Printf("Kết nối đến Kafka Broker: %s, Topic: %s\n", kafkaConfig.Brokers[0], kafkaConfig.Topic)

	// Tạo Producer
	producer := kafkaclient.NewProducer(kafkaConfig.Brokers, kafkaConfig.Topic)
	// Đảm bảo Close được gọi khi main kết thúc
	defer func() {
		if err := producer.Close(); err != nil {
			log.Printf("Lỗi khi đóng producer: %v", err)
		}
	}()

	// Giả lập tạo 5 đơn hàng và gửi lên Kafka
	orders := createSampleOrders()

	for _, order := range orders {
		// Context với timeout để đảm bảo không bị treo khi gửi message
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		
		err := producer.SendMessage(ctx, order.ID, order)
		cancel() // Đảm bảo hủy context sau khi xong

		if err != nil {
			log.Printf("Lỗi khi gửi đơn hàng %s: %v", order.ID, err)
		}

		log.Printf("Đã gửi đơn hàng: ID=%s, Khách=%s, Tổng=%.0f VNĐ\n",
            order.ID, order.CustomerID, order.TotalAmount)

		time.Sleep(500 * time.Millisecond)	// Dunngừ một chút giữa các lần gửi để dễ quan sát
	}

	log.Println("Đã gửi tất cả đơn hàng lên Kafka.")
}

// Tạo dữ liệu mẫu gửi lên
func createSampleOrders() []models.Order {

	customerIDs := []string{"KH-42", "KH-15", "KH-08", "KH-23", "KH-37"} // Nhóm khách hàng cố định
	var orders []models.Order
	
	for i := 1; i <= 10; i++ {

		// Dùng customerID theo vòng lặp để tạo các đơn hàng khác nhau
		customerID := customerIDs[(i-1)%len(customerIDs)]

		orders = append(orders, models.Order{
			ID:          fmt.Sprintf("ORD-%03d", i),
			CustomerID:  customerID,
			Status:      "pending",
			CreatedAt:   time.Now(),
			TotalAmount: float64(100000 + i*10000),
			Items: []models.OrderItem{
				{ProductID: "SP-001", Name: "Cà phê phin", Quantity: 1, Price: 45000},
			},
		})
	}
	return orders
}