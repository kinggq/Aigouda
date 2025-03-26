package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	MySQL struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
	} `yaml:"mysql"`
	JWT struct {
		Secret string `yaml:"secret"`
	} `yaml:"jwt"`
}

var GlobalConfig Config

func LoadConfig() error {
	// 优先使用环境变量
	if dbHost := os.Getenv("DB_HOST"); dbHost != "" {
		GlobalConfig.MySQL.Host = dbHost
	}
	if dbPort := os.Getenv("DB_PORT"); dbPort != "" {
		fmt.Sscanf(dbPort, "%d", &GlobalConfig.MySQL.Port)
	}
	if dbUser := os.Getenv("DB_USER"); dbUser != "" {
		GlobalConfig.MySQL.User = dbUser
	}
	if dbPassword := os.Getenv("DB_PASSWORD"); dbPassword != "" {
		GlobalConfig.MySQL.Password = dbPassword
	}
	if dbName := os.Getenv("DB_NAME"); dbName != "" {
		GlobalConfig.MySQL.DBName = dbName
	}
	if jwtSecret := os.Getenv("JWT_SECRET"); jwtSecret != "" {
		GlobalConfig.JWT.Secret = jwtSecret
	}

	// 如果环境变量未设置，则读取配置文件
	if GlobalConfig.MySQL.Host == "" {
		file, err := os.ReadFile("config/config.yaml")
		if err != nil {
			return fmt.Errorf("error reading config file: %v", err)
		}

		err = yaml.Unmarshal(file, &GlobalConfig)
		if err != nil {
			return fmt.Errorf("error parsing config file: %v", err)
		}
	}

	return nil
} 
