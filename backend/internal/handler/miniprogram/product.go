package miniprogram

import (
	"aigouda/internal/service/miniprogram"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ProductList 小程序获取商品列表
func ProductList(c *gin.Context) {
	// 从 Query 参数获取分页信息
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	keyword := c.Query("keyword")

	// 参数验证
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 50 {
		pageSize = 10
	}

	products, total, err := miniprogram.GetMiniProductList(page, pageSize, keyword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": gin.H{
			"list":     products,
			"total":    total,
			"page":     page,
			"pageSize": pageSize,
		},
	})
}

// GetProduct 小程序获取商品详情
func GetProduct(c *gin.Context) {
	id := c.Param("id")
	product, err := miniprogram.GetMiniProduct(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": product,
	})
} 
