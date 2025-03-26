package main

import (
	"aigouda/config"
	"aigouda/internal/router"
	"aigouda/internal/repository"
	"log"
)

func main() {
	// 初始化配置
	config.Init()

	// 初始化数据库
	if err := repository.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 设置路由
	r := router.SetupRouter()

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
} 
