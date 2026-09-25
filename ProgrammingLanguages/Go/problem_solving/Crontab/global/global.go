package global

import (
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

// Tạo biến toàn cục để lưu trữ thông tin đối tượng và có phạm vi sử dụng trong toàn bộ ứng dụng
var (
	GO_CRON *cron.Cron
	GO_LOGGER *zap.Logger
)