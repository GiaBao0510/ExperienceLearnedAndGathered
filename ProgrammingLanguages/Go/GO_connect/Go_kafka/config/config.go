package config

import "github.com/GiaBao0510/kafka-go/utils"

type KafkaConfig struct {
	// Danh sách các broker của Kafka
	// Trong môi trường dev: thì chỉ cần 1 broker
	// Trong môi trường production: thì cần nhiều broker để đảm bảo tính sẵn sàng và khả năng mở rộng
	Brokers []string

	// Tên topic mà ứng dụng sẽ gửi hoặc nhận dữ liệu
	Topic string

	// Dùng cho Consumer
	GroupID string
}

// LoadKafkaConfig đọc biến môi trường
func LoadKafkaConfig() KafkaConfig {
	broker := utils.GetEnv("KAFKA_BROKER", "localhost:9092")
	return KafkaConfig{
		Brokers: []string{broker},
		Topic:   utils.GetEnv("KAFKA_TOPIC", "don-hang"),
		GroupID: utils.GetEnv("KAFKA_GROUP_ID", "nhom-xu-ly-don-hang"),
	}
}