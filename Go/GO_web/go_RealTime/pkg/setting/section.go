package setting

// Struct này chứa toàn bộ cấu hình
type Config struct {
	DB     DBConfig     `mapstructure:"Database"`
	Server ServerConfig `mapstructure:"Server"`
}

// Cấu hình kết nối DB
type DBConfig struct {
	DB_Host     string `mapstructure:"DB_Host"`
	DB_Port     int    `mapstructure:"DB_Port"`
	DB_User     string `mapstructure:"DB_User"`
	DB_Password string `mapstructure:"DB_Password"`
	DB_Name     string `mapstructure:"DB_Name"`
	DB_SSLMode  string `mapstructure:"DB_SSLMode"`
}

// Cấu hình của server
type ServerConfig struct {
	Server_Host string `mapstructure:"Server_Host"`
	Server_Port string `mapstructure:"Server_Port"`
}
