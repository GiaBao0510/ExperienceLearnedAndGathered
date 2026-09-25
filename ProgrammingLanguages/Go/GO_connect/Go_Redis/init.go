package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"gopkg.in/yaml.v3"
)

// Khai báo cấu trúc để map với file config.yaml
type Config struct {
	Redis struct {
		Default struct {
			Address      string `yaml:"address"`
			Password     string `yaml:"password"`
			DB           int    `yaml:"db"`
			ReadTimeout  string `yaml:"readTimeout"`
			WriteTimeout string `yaml:"writeTimeout"`
		} `yaml:"default"`
	} `yaml:"redis"`
}

// Khai báo biến Redis Client toàn cục
var Rdb *redis.Client
var Ctx = context.Background()

func InitRedis() {
	// 1. Đọc file config.yaml
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Lỗi đọc file config: %v", err)
	}

	// 2. Parse YAML vào cấu trúc Config
	var cfg Config
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		log.Fatalf("Lỗi Parse YAML: %v", err)
	}

	// 3. Chuyển đổi string timeout sang time.Duration
	readTimeout, _ := time.ParseDuration(cfg.Redis.Default.ReadTimeout)
	writeTimeout, _ := time.ParseDuration(cfg.Redis.Default.WriteTimeout)

	// 4. Khởi tạo Redis Client
	Rdb = redis.NewClient(&redis.Options{
		Addr:         cfg.Redis.Default.Address,  // địa chỉ Redis server
		Password:     cfg.Redis.Default.Password, // mật khẩu (nếu có)
		DB:           cfg.Redis.Default.DB,       // số database (mặc định là 0)
		ReadTimeout:  readTimeout,                // thời gian chờ đọc dữ liệu
		WriteTimeout: writeTimeout,               // thời gian chờ ghi dữ liệu
	})

	// 5. Kiểm tra kết nối
	_, err = Rdb.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Không thể kết nối Redis: %v", err)
	}
	fmt.Println("✅ Đã khởi tạo và kết nối Redis thành công!")
}
