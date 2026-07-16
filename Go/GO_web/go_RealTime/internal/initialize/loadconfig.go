package 

import "log"

func LoadConfig() {
	// Khởi tạo instance viper
	my_viper := viper.New()

	// Nhận dạng file config
	my_viper.AddConfigPath("./configs")
	my_viper.SetConfigName("local")
	my_viper.SetConfigType("yaml")

	// Đọc nội dung file congfig vào bộ nhớ
	if err := my_viper.ReadConfig(); err != nil {
		log.Fatal("Lỗi khi đọc file config: ", err)
	}

	//
}
