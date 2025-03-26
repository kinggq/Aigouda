package router

import (
	"aigouda/internal/handler"
	"aigouda/internal/middleware"
	"github.com/gin-gonic/gin"
)

// 自定义 CORS 中间件
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400") // 24小时

		if c.Request.Method == "OPTIONS" {
			c.JSON(200, gin.H{
				"code":    0,
				"message": "ok",
			})
			return
		}

		c.Next()
	}
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 使用自定义 CORS 中间件
	r.Use(corsMiddleware())

	// 配置静态文件服务
	r.Static("/uploads", "./uploads")

	// 测试路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	// 处理所有 OPTIONS 请求
	r.OPTIONS("/*path", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    0,
			"message": "ok",
		})
	})

	// 用户相关路由
	userHandler := handler.NewUserHandler()
	r.POST("/api/register", userHandler.Register)
	r.POST("/api/login", userHandler.Login)

	// 需要认证的路由组
	authorized := r.Group("/api")
	authorized.Use(middleware.AuthMiddleware())
	{
		// 用户相关
		authorized.POST("/user/profile", userHandler.GetCurrentUser)

		// 商品相关
		productHandler := handler.NewProductHandler()
		authorized.GET("/products", productHandler.List)
		authorized.GET("/products/:id", productHandler.Get)
		authorized.POST("/products/create", productHandler.Create)
		authorized.PUT("/products/:id", productHandler.Update)
		authorized.DELETE("/products/:id", productHandler.Delete)
		authorized.POST("/products/upload", productHandler.UploadImage)

		// 分类相关
		categoryHandler := handler.NewCategoryHandler()
		authorized.POST("/categories/list", categoryHandler.List)
		authorized.GET("/categories/:id", categoryHandler.Get)
		authorized.POST("/categories/create", categoryHandler.Create)
		authorized.POST("/categories/update/:id", categoryHandler.Update)
		authorized.DELETE("/categories/delete/:id", categoryHandler.Delete)
	}

	return r
} 
