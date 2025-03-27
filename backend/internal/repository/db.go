package repository

import (
	"aigouda/config"
	"aigouda/internal/model"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB() error {
	dsn := config.GetDSN()
	log.Printf("Connecting to database with DSN: %s", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	// 设置连接池
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db

	// 自动迁移数据库表
	err = DB.AutoMigrate(
		&model.User{},
		&model.Product{},
		&model.Category{},
	)
	if err != nil {
		return err
	}

	log.Println("Database initialized successfully")
	return nil
} 
