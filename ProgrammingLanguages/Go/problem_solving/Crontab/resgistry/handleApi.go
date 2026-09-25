package resgistry

import (
	"github.com/GiaBao0510/scheduler-cron-api/api"
	"github.com/GiaBao0510/scheduler-cron-api/global"
)

// Đăng ký các API để chạy cron job
func RegisApiRunCronjob() {
	api.SendEmailForVipUsersEvery3Seconds(global.GO_CRON)
	api.GetInformationForVipUsersEvery5Seconds(global.GO_CRON)
	//... Ngoài ra còn có các sự kiến khác

	global.GO_CRON.Start()	// Chạy cron job
}