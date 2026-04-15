package main

import (
	"context"
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
	return []models.Order{
		{
            ID:         "ORD-001",
            CustomerID: "KH-42",
            Status:     "pending",
            CreatedAt:  time.Now(),
            TotalAmount: 350000,
            Items: []models.OrderItem{
                {ProductID: "SP-001", Name: "Cà phê phin", Quantity: 2, Price: 45000},
                {ProductID: "SP-002", Name: "Bánh mì thịt", Quantity: 3, Price: 25000},
                {ProductID: "SP-003", Name: "Nước cam", Quantity: 1, Price: 35000},
            },
        },
        {
            ID:         "ORD-002",
            CustomerID: "KH-15",
            Status:     "pending",
            CreatedAt:  time.Now(),
            TotalAmount: 199000,
            Items: []models.OrderItem{
                {ProductID: "SP-004", Name: "Trà sữa trân châu", Quantity: 2, Price: 55000},
                {ProductID: "SP-005", Name: "Bánh flan", Quantity: 1, Price: 25000},
            },
        },
        {
            ID:         "ORD-003",
            CustomerID: "KH-08",
            Status:     "pending",
            CreatedAt:  time.Now(),
            TotalAmount: 520000,
            Items: []models.OrderItem{
                {ProductID: "SP-006", Name: "Bún bò Huế", Quantity: 4, Price: 65000},
            },
        },
	}
}