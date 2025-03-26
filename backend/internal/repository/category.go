package repository

import (
	"aigouda/internal/model"
	"errors"
)

type CategoryRepository struct{}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{}
}

// List 获取分类列表
func (r *CategoryRepository) List(offset, limit int, search string) ([]model.Category, int64, error) {
	var categories []model.Category
	var total int64

	query := DB.Model(&model.Category{})
	if search != "" {
		query = query.Where("name LIKE ? OR description LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	if err := query.Offset(offset).Limit(limit).Find(&categories).Error; err != nil {
		return nil, 0, err
	}

	return categories, total, nil
}

// GetByID 获取单个分类
func (r *CategoryRepository) GetByID(id uint) (*model.Category, error) {
	var category model.Category
	if err := DB.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

// Create 创建分类
func (r *CategoryRepository) Create(category *model.Category) error {
	return DB.Create(category).Error
}

// Update 更新分类
func (r *CategoryRepository) Update(category *model.Category) error {
	return DB.Save(category).Error
}

// Delete 删除分类
func (r *CategoryRepository) Delete(id uint) error {
	// 检查是否有子分类
	var count int64
	if err := DB.Model(&model.Category{}).Where("parent_id = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("该分类下有子分类，无法删除")
	}

	// 检查是否有关联的商品
	count = DB.Model(&model.Category{}).Association("Products").Count()
	if count > 0 {
		return errors.New("该分类下有关联商品，无法删除")
	}

	return DB.Delete(&model.Category{}, id).Error
}

// GetProducts 获取分类下的商品
func (r *CategoryRepository) GetProducts(categoryID uint, offset, limit int) ([]model.Product, int64, error) {
	var category model.Category
	if err := DB.First(&category, categoryID).Error; err != nil {
		return nil, 0, err
	}

	var products []model.Product
	var total int64

	// 获取总数
	total = DB.Model(&category).Association("Products").Count()

	// 获取分页数据
	if err := DB.Model(&category).Association("Products").Find(&products); err != nil {
		return nil, 0, err
	}

	// 手动分页
	start := offset
	end := offset + limit
	if start >= len(products) {
		return []model.Product{}, total, nil
	}
	if end > len(products) {
		end = len(products)
	}

	return products[start:end], total, nil
} 
