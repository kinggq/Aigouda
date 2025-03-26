package service

import (
	"aigouda/internal/model"
	"aigouda/internal/repository"
)

type CategoryService struct {
	categoryRepo *repository.CategoryRepository
}

func NewCategoryService() *CategoryService {
	return &CategoryService{
		categoryRepo: repository.NewCategoryRepository(),
	}
}

type CreateCategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ParentID    *uint  `json:"parent_id"`
}

type UpdateCategoryRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	ParentID    *uint   `json:"parent_id"`
}

// GetCategoryList 获取分类列表
func GetCategoryList(page, pageSize int, keyword string) ([]model.Category, int64, error) {
	offset := (page - 1) * pageSize
	var categories []model.Category
	var total int64

	query := repository.DB.Model(&model.Category{})
	if keyword != "" {
		query = query.Where("name LIKE ?", "%"+keyword+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Offset(offset).Limit(pageSize).Find(&categories).Error
	if err != nil {
		return nil, 0, err
	}

	return categories, total, nil
}

// GetCategory 获取分类详情
func GetCategory(id uint) (*model.Category, error) {
	var category model.Category
	err := repository.DB.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// CreateCategory 创建分类
func CreateCategory(category *model.Category) error {
	return repository.DB.Create(category).Error
}

// UpdateCategory 更新分类
func UpdateCategory(category *model.Category) error {
	return repository.DB.Save(category).Error
}

// DeleteCategory 删除分类
func DeleteCategory(id uint) error {
	return repository.DB.Delete(&model.Category{}, id).Error
}

// GetProducts 获取分类下的商品
func (s *CategoryService) GetProducts(categoryID uint, page, pageSize int) ([]model.Product, int64, error) {
	offset := (page - 1) * pageSize
	products, total, err := s.categoryRepo.GetProducts(categoryID, offset, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return products, total, nil
} 
