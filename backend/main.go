package main

import (
	"aigouda/config"
	"aigouda/internal/repository"
	"aigouda/internal/router"
	"log"
	"os"
)

func main() {
	// 加载配置
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 初始化数据库
	if err := repository.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 设置路由
	r := router.SetupRouter()

	// 获取端口
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 启动服务器
	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
} 
