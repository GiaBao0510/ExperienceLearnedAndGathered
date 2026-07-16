package config

// Struct này chứa toàn bộ cấu hình
type Config struct {
	DB     DBConfig     `mapstructure:"database"`
	Server ServerConfig `mapstructure:"server"`
}

// Cấu hình kết nối DB
type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"name"`
	SSLMode  string `mapstructure:"sslmode"`
}

// Cấu hình của server
type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}
