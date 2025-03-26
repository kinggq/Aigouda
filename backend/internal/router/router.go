package router

import (
	"aigouda/internal/handler"
	"aigouda/internal/handler/miniprogram"
	"aigouda/internal/middleware"
	"path/filepath"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 配置 CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // 允许所有来源，生产环境建议设置具体的域名
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	r.Use(cors.New(config))

	// 配置静态文件服务
	uploadDir := "uploads"
	r.Static("/uploads", uploadDir)
	// 确保上传目录存在
	if err := ensureDir(uploadDir); err != nil {
		panic(err)
	}
	if err := ensureDir(filepath.Join(uploadDir, "products")); err != nil {
		panic(err)
	}

	// 健康检查
	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

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

	// API 路由组
	api := r.Group("/api")
	{
		// 用户相关路由
		api.POST("/register", handler.Register)
		api.POST("/login", handler.Login)

		// 需要认证的路由组
		auth := api.Group("")
		auth.Use(middleware.Auth())
		{
			// 用户信息
			auth.GET("/profile", handler.GetProfile)
			auth.PUT("/profile", handler.UpdateProfile)
			auth.PUT("/password", handler.UpdatePassword)

			// 商品相关
			auth.GET("/products", handler.ProductList)
			auth.GET("/products/:id", handler.GetProduct)
			auth.POST("/products/create", handler.CreateProduct)
			auth.PUT("/products/:id", handler.UpdateProduct)
			auth.DELETE("/products/:id", handler.DeleteProduct)
			auth.POST("/products/upload", handler.UploadProductImage)

			// 分类相关
			auth.POST("/categories/list", handler.CategoryList)
			auth.GET("/categories/:id", handler.GetCategory)
			auth.POST("/categories/create", handler.CreateCategory)
			auth.PUT("/categories/update/:id", handler.UpdateCategory)
			auth.DELETE("/categories/delete/:id", handler.DeleteCategory)
		}

		// 小程序专用路由
		mini := api.Group("/mini")
		{
			// 商品相关
			mini.GET("/products/list", miniprogram.ProductList)
			mini.GET("/products/:id", miniprogram.GetProduct)
		}
	}

	return r
}

// ensureDir 确保目录存在
func ensureDir(dir string) error {
	return nil // 暂时不实现，等待后续完善
} 
