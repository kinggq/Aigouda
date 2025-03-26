package miniprogram

import (
	"aigouda/internal/model"
	"aigouda/internal/repository"
	"errors"
)

// GetMiniProductList 小程序获取商品列表
func GetMiniProductList(page, pageSize int, keyword string) ([]model.Product, int64, error) {
	offset := (page - 1) * pageSize
	
	// 构建查询条件
	query := repository.DB.Model(&model.Product{})
	
	// 添加关键词搜索
	if keyword != "" {
		query = query.Where("title LIKE ?", "%"+keyword+"%")
	}
	
	// 获取总数
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	
	// 获取分页数据
	var products []model.Product
	if err := query.Offset(offset).Limit(pageSize).Find(&products).Error; err != nil {
		return nil, 0, err
	}
	
	return products, total, nil
}

// GetMiniProduct 小程序获取商品详情
func GetMiniProduct(id string) (*model.Product, error) {
	var product model.Product
	if err := repository.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, errors.New("商品不存在")
	}
	return &product, nil
} 
