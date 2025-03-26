package repository

import (
	"aigouda/internal/model"
)

type ProductRepository struct{}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

// List 获取商品列表
func (r *ProductRepository) List(offset, limit int, search string) ([]model.Product, int64, error) {
	var products []model.Product
	var total int64

	query := DB.Model(&model.Product{}).Preload("Category")
	if search != "" {
		query = query.Where("title LIKE ? OR description LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	if err := query.Offset(offset).Limit(limit).Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

// GetByID 获取单个商品
func (r *ProductRepository) GetByID(id uint) (*model.Product, error) {
	var product model.Product
	if err := DB.Preload("Category").First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

// Create 创建商品
func (r *ProductRepository) Create(product *model.Product) error {
	return DB.Create(product).Error
}

// Update 更新商品
func (r *ProductRepository) Update(id uint, updates map[string]interface{}) error {
	return DB.Model(&model.Product{}).Where("id = ?", id).Updates(updates).Error
}

// Delete 删除商品
func (r *ProductRepository) Delete(id uint) error {
	return DB.Delete(&model.Product{}, id).Error
}

// GetByCategoryID 获取分类下的商品
func (r *ProductRepository) GetByCategoryID(categoryID uint, offset, limit int) ([]model.Product, int64, error) {
	var products []model.Product
	var total int64

	query := DB.Model(&model.Product{}).Where("category_id = ?", categoryID).Preload("Category")

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	if err := query.Offset(offset).Limit(limit).Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, total, nil
} 
