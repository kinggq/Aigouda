package service

import (
	"aigouda/internal/model"
	"aigouda/internal/repository"
	"errors"
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

// List 获取分类列表
func (s *CategoryService) List(page, pageSize int, search string) ([]model.Category, int64, error) {
	offset := (page - 1) * pageSize
	categories, total, err := s.categoryRepo.List(offset, pageSize, search)
	if err != nil {
		return nil, 0, err
	}
	return categories, total, nil
}

// Get 获取单个分类
func (s *CategoryService) Get(id uint) (*model.Category, error) {
	category, err := s.categoryRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

// Create 创建分类
func (s *CategoryService) Create(category *model.Category) error {
	if category.Name == "" {
		return errors.New("分类名称不能为空")
	}
	return s.categoryRepo.Create(category)
}

// Update 更新分类
func (s *CategoryService) Update(category *model.Category) error {
	if category.ID == 0 {
		return errors.New("分类ID不能为空")
	}
	if category.Name == "" {
		return errors.New("分类名称不能为空")
	}
	return s.categoryRepo.Update(category)
}

// Delete 删除分类
func (s *CategoryService) Delete(id uint) error {
	return s.categoryRepo.Delete(id)
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
