package initialize

import (
	"log"

	"github.com/GiaBao0510/Go-Realtime/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	// Khởi tạo instance viper
	my_viper := viper.New()

	// Nhận dạng file config
	my_viper.AddConfigPath("./configs")
	my_viper.SetConfigName("local")
	my_viper.SetConfigType("yaml")

	// Đọc nội dung file congfig vào bộ nhớ
	if err := my_viper.ReadInConfig(); err != nil {
		log.Fatal("Lỗi khi đọc file config: ", err)
	}

	//Map cấu hình vào struct Config
	if err := my_viper.Unmarshal(&global.Config); err != nil {
		log.Fatal("Lỗi khi map cấu hình vào struct Config: ", err)
	}

	log.Println("Cấu hình đã được load thành công")
}
