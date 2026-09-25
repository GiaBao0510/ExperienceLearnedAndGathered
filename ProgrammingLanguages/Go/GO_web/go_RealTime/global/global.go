package global

import (
	"database/sql"

	"github.com/GiaBao0510/Go-Realtime/pkg/setting"
)

var (
	Config setting.Config
	PostgreSQL *sql.DB
)