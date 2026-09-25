package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB là biến toàn cục để kết nối với MongoDB
var DB *mongo.Database

// Thực hiện kết nối đến Database MongoDB
func ConnectDB() {

	//Lấy Connection từ biến môi trường
	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB_NAME")

	//Kiểm tra nếu biến môi trường không tồn tại
	if mongoURI == "" || dbName == "" {
		log.Fatal("MONGO_URI and MONGO_DB_NAME environment variables are not set")
	}

	// Tạo context với timeout trong 10s (nếu trong 10s không kết nối được sẽ tự động hủy)
	ctx, cancel :=  context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()  // Đảm bảo rằng context sẽ được hủy sau khi kết nối xong

	//Cấu hình kết nối
	clientOptions := options.Client().ApplyURI(mongoURI)

	//Thực hiện kết nối đến MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB: ", err)
	}

	//Kiểm tra kết nối có thành công không
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Error pinging MongoDB: ", err)
	}

	//Gán kết nối vào biến toàn cục DB
	DB = client.Database(dbName)

	log.Println("Successfully connected to MongoDB!")
}

// GetCollection trả về collection từ database
func GetCollection(collectionName string) *mongo.Collection {
	return DB.Collection(collectionName)
}