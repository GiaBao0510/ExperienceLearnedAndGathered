package config

// Struct này chứa toàn bộ cấu hình
type Config struct {
	DB     DBConfig     `json:"db"`
	Server ServerConfig `json:"server"`
}

// Cấu hình kết nối DB
type DBConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"name"`
	SSLMode  string `json:"sslmode"`
}

// Cấu hình của server
type ServerConfig struct {
	Port string `json:"port"`
}
