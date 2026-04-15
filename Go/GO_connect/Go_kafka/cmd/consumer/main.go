package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/GiaBao0510/kafka-go/config"
	kafkaClient "github.com/GiaBao0510/kafka-go/internal/kafka"
	"github.com/GiaBao0510/kafka-go/models"
)

func main() {
	// Load cấu hình từ biến môi trường
	cfg := config.LoadKafkaConfig()
	log.Printf("Kết nối đến Kafka Broker: %s, Topic: %s\n", cfg.Brokers[0], cfg.Topic)

	// Tạo consumer thuộc group "nhom-xu-ly-don-hang"
	consumer := kafkaClient.NewConsumer(cfg.Brokers, cfg.Topic, cfg.GroupID)
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Printf("Lỗi khi đóng consumer: %v", err)
		}
	}()

	// Thiết lập xử lý tín hiệu Ctrl+C để đóng consumer đúng cách
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	log.Println("Consumer đang chạy... Nhấn Ctrl+C để dừng.")

	// Vòng lặp chính để nhận và xử lý message
	for {
		msg, err := consumer.ReadMessage(ctx)
		if err != nil {
			if errors.Is(ctx.Err(), context.Canceled){
				log.Println("Nhận tín hiệu dừng, thoát chương trình...")
				break
			}
			log.Printf("Lỗi khi đọc message: %v", err)
			continue
		}

		// Xử lý message nhận được
		handleOrder(msg)
	}
}

// Hàm xử lý message nhận được từ Kafka
func handleOrder(msg kafkaClient.Message) {
	var order models.Order
	if err := json.Unmarshal(msg.Value, &order); err != nil {
		log.Printf("Lỗi khi giải mã message: %v", err)
		return
	}
	// In thông tin
	log.Printf("Nhận đơn hàng mới:")
	log.Printf("  Partition: %d | Offset: %d", msg.Partition, msg.Offset)
	log.Printf("  ID đơn hàng: %s", order.ID)
	log.Printf("  Khách hàng:  %s", order.CustomerID)
	log.Printf("  Tổng tiền:   %.0f VNĐ", order.TotalAmount)
	log.Printf("  Số món:      %d sản phẩm", len(order.Items))
	log.Printf("  Trạng thái:  %s", order.Status)
	log.Println("---")

	// Trong ứng dụng thực tế, đây là nơi bạn:
	// - Lưu vào database
	// - Gửi email xác nhận
	// - Cập nhật kho hàng
	// - v.v...
}