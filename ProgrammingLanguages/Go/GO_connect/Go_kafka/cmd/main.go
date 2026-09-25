package main

import (
	"fmt"

	"github.com/GiaBao0510/kafka-go/config"
)

func main() {

	kafkaConfig := config.LoadKafkaConfig()


	// Hiển thị thông tin lấy từ file .env
	fmt.Printf("Broker: %s\n", kafkaConfig.Brokers[0])
	fmt.Printf("Topic: %s\n", kafkaConfig.Topic)
	fmt.Printf("GroupID: %s\n", kafkaConfig.GroupID)

}