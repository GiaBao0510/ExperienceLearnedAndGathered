package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

// Producer bọc kafka.Writer để gửi dữ liệu lên Kafka
type Producer struct {
	writer *kafka.Writer
}

// NewProducer khởi tạo Producer mới kết nối đến Kafka với cấu hình đã cho
func NewProducer(brokers []string, topic string) *Producer {
	writer := &kafka.Writer{
		Addr:                   kafka.TCP(brokers...), // Addr là địa chỉ của Kafka broker, có thể là một hoặc nhiều broker
		Topic:                  topic,                 // Topic là tên topic mà producer sẽ gửi dữ liệu lên
		Balancer:               &kafka.LeastBytes{},   // Balancer là chiến lược phân phối message đến các partition, ở đây dùng LeastBytes để gửi đến partition có ít dữ liệu nhất
		WriteTimeout:           10 * time.Second,      // WriteTimeout: nếu không ghi trong 10s thì sẽ gặp lỗi
		ReadTimeout:            10 * time.Second,      // ReadTimeout: nếu không đọc trong 10s thì sẽ gặp lỗi
		RequiredAcks:           kafka.RequireAll,      // RequiredAcks: chờ tất cả replica xác nhận. Đảm bảo messahe không bị mất nếu broker crash
		AllowAutoTopicCreation: true,                  // AllowAutoTopicCreation: nếu topic chưa tồn tại thì tự động tạo
	}

	return &Producer{writer: writer}
}

// SendMessage gửi một object bất kỳ đến Kakfa dưới dạng JSON
func (p *Producer) SendMessage(ctx context.Context, key string, value any) error {
	// Chuyển value thành JSON bytes
	jsonBytes, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("Không thể marshal dữ liệu thành JSON: %v", err)
	}

	// Tạo message Kafka với key và value đã được marshal
	msg := kafka.Message{
		Key:   []byte(key),
		Value: jsonBytes,
	}

	// Gửi message đến Kafka
	if err := p.writer.WriteMessages(ctx, msg); err != nil {
		return fmt.Errorf("Không thể gửi message đến Kafka: %v", err)
	}

	return nil
}

// Close đóng kết nối của producer
func (p *Producer) Close() error {
	return p.writer.Close()
}
