package main

import (
	"github/GiaBao0510/error-handling/config"
	"github/GiaBao0510/error-handling/internal/controller"
	"github/GiaBao0510/error-handling/internal/repository"
	"github/GiaBao0510/error-handling/internal/router"
	"github/GiaBao0510/error-handling/internal/service"
	"log"
	"os"
)

func main() {

	// Tải biến môi trường
	if err := config.LoadEnv(); err != nil {
		log.Fatalf("Lỗi khi tải biến môi trường: %v", err)
	}

	// Kết nối đến DB
	dbConfig := config.DBConfig{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		User:     os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
		Dbname:   os.Getenv("DBNAME"),
		Sslmode:  os.Getenv("SSLMODE"),
	}
	dns := dbConfig.DNS()
	db, err := config.NewDatabase(dns)

	if err != nil {
		log.Fatalf("Lỗi khi kết nối đến cơ sở dữ liệu: %v", err)
	}
	defer db.Close()

	// Dependency Injection thủ công
	userRepo := repository.NewUserRepository(db)
	userSVC := service.NewUserSerivce(userRepo)
	userCtrl := controller.NewUserController(userSVC)

	r := router.SetupRouterForUser(userCtrl)

	log.Println("🚀Server đang chạy tại http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Lỗi khi chạy server: %v", err)
	}
}
