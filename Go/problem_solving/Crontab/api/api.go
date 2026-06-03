package api

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/GiaBao0510/scheduler-cron-api/global"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

// Gửi email cho mỗi người dùng Vip mỗi 3 giây
func SendEmailForVipUsersEvery3Seconds(cr *cron.Cron) {
	fmt.Println("... Gửi email cho mỗi người dùng Vip mỗi 3 giây ...")

	// Theo định dạng
	// * * * * * *
	// Giây, Phút, Giờ, Ngày, Tháng, Năm
	_, err := cr.AddFunc("*/3 * * * * *", func(){
		log.Println("Gửi email cho người dùng Vip")
	})

	if err != nil {
		global.GO_LOGGER.Error("Lỗi khi thêm cron job: ", zap.Error(err))	
	}
}

func GetInformationForVipUsersEvery5Seconds(cr *cron.Cron) {
	fmt.Println("... Lấy thông tin cho mỗi người dùng Vip mỗi 5 giây ...")

	_, err := cr.AddFunc("*/5 * * * * *", func(){
		log.Println("Lấy thông tin cho người dùng Vip")
	})

	rs, err := http.Get("https://httpbin.org/get")
	if err != nil {
		global.GO_LOGGER.Error("Lỗi khi gửi yêu cầu HTTP: ", zap.Error(err))
		return
	}

	body, err := io.ReadAll(rs.Body)
	if err != nil {
		global.GO_LOGGER.Error("Lỗi khi đọc phản hồi HTTP: ", zap.Error(err))
		return
	}

	log.Println("Get information for Vip users:", string(body))

	if err != nil {
		global.GO_LOGGER.Error("Lỗi khi thêm cron job: ", zap.Error(err))
	}
}