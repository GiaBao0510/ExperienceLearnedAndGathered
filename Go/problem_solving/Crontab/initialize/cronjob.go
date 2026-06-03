package initialize

import "github.com/robfig/cron/v3"

// Khởi tạo cron job
func InitCron() *cron.Cron{
	c := cron.New(cron.WithSeconds())
	return c
}