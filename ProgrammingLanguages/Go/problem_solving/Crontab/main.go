package main

import (
	"github.com/GiaBao0510/scheduler-cron-api/global"
	"github.com/GiaBao0510/scheduler-cron-api/initialize"
	"github.com/GiaBao0510/scheduler-cron-api/resgistry"
)

func main() {

	// Khởi tạo Logger
	global.GO_LOGGER = initialize.InitLogger()

	// Khởi tạo cron job
	global.GO_CRON = initialize.InitCron()

	// Chạy cron job
	resgistry.RegisApiRunCronjob()

	select{}	// Giữ cho ứng dụng chạy mãi mãi
}