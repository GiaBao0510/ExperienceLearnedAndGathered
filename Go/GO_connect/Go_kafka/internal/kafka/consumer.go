package kafka

import (
	"context"
	"fmt"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

// Consumer bọc kafka.Reader để nhận dữ liệu từ Kafka
type Consumer struct {
	reader *kafka.Reader
}

// Tạo struct chứa thông tin Message đã nhận
type Message struct {
	Key       string    // Key của message, thường dùng để phân phối message đến các partition
	Value     []byte    // Value của message, chứa dữ liệu thực tế được gửi từ producer
	Partition int       // Partition mà message được gửi đến, giúp xác định nơi lưu trữ message trong Kafka
	Offset    int64     // Offset của message trong partition, giúp xác định vị trí của message trong Kafka
	Timestamp time.Time // Thời gian tạo message
}

// NewConsumer khởi tạo Consumer mới kết nối đến Kafka với cấu hình đã cho
func NewConsumer(brokers []string, topic, groupID string) *Consumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: groupID, // Định nghĩa groud ID để Kafka  dùng GroupID để theo dõi offset đã commit

		// MinBytes và MaxBytes kiểm soát kích thước fetch (fetch.min.bytes / fetch.max.bytes).
		MinBytes: 10e3,		// Trả dữ liệu ngay khi có >=1 byte sẵn sàng, tránh delay chờ gom đủ 10KB
		MaxBytes: 10e6,		// 0MB - chỉ là trần bảo vệ RAM, không bắt buộc phải đạt tới mới trả dữ liệu

		// CommitInterval: tự động commit offset mỗi 1 giây.
		CommitInterval: time.Second,

		StartOffset: kafka.FirstOffset, // Bắt đầu đọc từ offset đầu tiên nếu chưa có offset nào được commit trước đó
	})

	return &Consumer{reader: reader}
}

// ReadMessage đọc một message từ Kafka và trả về struct Message đã được giải mã
func (c *Consumer) ReadMessage(ctx context.Context) (Message, error) {
	msg, err := c.reader.ReadMessage(ctx)
	if err != nil {
		return Message{}, fmt.Errorf("Đọc message thất bại:: %v", err)
	}
	return Message{
		Key:       string(msg.Key),
		Value:     msg.Value,
		Partition: msg.Partition,
		Offset:    msg.Offset,
		Timestamp: msg.Time,
	}, nil
}

// Close đóng kết nối của consumer
func (c *Consumer)Close() error{
	return c.reader.Close()
}
