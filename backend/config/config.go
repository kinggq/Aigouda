package config

type Config struct {
	MySQL struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
	}
	JWT struct {
		Secret string
		Expire int // token过期时间（小时）
	}
}

var GlobalConfig Config

func Init() {
	// 默认配置
	GlobalConfig.MySQL.Host = "localhost"
	GlobalConfig.MySQL.Port = 3306
	GlobalConfig.MySQL.User = "root"
	GlobalConfig.MySQL.Password = ""
	GlobalConfig.MySQL.DBName = "aigouda"

	GlobalConfig.JWT.Secret = "your-secret-key"
	GlobalConfig.JWT.Expire = 24
} 
