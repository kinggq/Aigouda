package service

import (
	"aigouda/internal/model"
	"aigouda/internal/repository"
	"errors"
)

type TagService struct{}

func NewTagService() *TagService {
	return &TagService{}
}

type CreateTagRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateTagRequest struct {
	Name *string `json:"name"`
}

// Create 创建标签
func (s *TagService) Create(req *CreateTagRequest) (*model.Tag, error) {
	tag := &model.Tag{
		Name: req.Name,
	}

	if err := repository.DB.Create(tag).Error; err != nil {
		return nil, err
	}

	return tag, nil
}

// Update 更新标签
func (s *TagService) Update(id uint, req *UpdateTagRequest) (*model.Tag, error) {
	var tag model.Tag
	if err := repository.DB.First(&tag, id).Error; err != nil {
		return nil, errors.New("标签不存在")
	}

	updates := map[string]interface{}{}
	if req.Name != nil {
		updates["name"] = *req.Name
	}

	if err := repository.DB.Model(&tag).Updates(updates).Error; err != nil {
		return nil, err
	}

	return &tag, nil
}

// Delete 删除标签
func (s *TagService) Delete(id uint) error {
	// 检查是否有关联的商品
	count := repository.DB.Model(&model.Tag{}).Association("Products").Count()
	if count > 0 {
		return errors.New("该标签下有关联商品，无法删除")
	}

	return repository.DB.Delete(&model.Tag{}, id).Error
}

// List 获取标签列表
func (s *TagService) List() ([]model.Tag, error) {
	var tags []model.Tag
	if err := repository.DB.Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

// GetByID 获取标签详情
func (s *TagService) GetByID(id uint) (*model.Tag, error) {
	var tag model.Tag
	if err := repository.DB.First(&tag, id).Error; err != nil {
		return nil, errors.New("标签不存在")
	}
	return &tag, nil
}

// GetProducts 获取标签下的商品
func (s *TagService) GetProducts(id uint, page, pageSize int) ([]model.Product, int64, error) {
	var tag model.Tag
	if err := repository.DB.First(&tag, id).Error; err != nil {
		return nil, 0, errors.New("标签不存在")
	}

	var products []model.Product
	var total int64

	// 获取总数
	total = repository.DB.Model(&tag).Association("Products").Count()

	// 获取分页数据
	if err := repository.DB.Model(&tag).Association("Products").Find(&products); err != nil {
		return nil, 0, err
	}

	// 手动分页
	offset := (page - 1) * pageSize
	start := offset
	end := offset + pageSize
	if start >= len(products) {
		return []model.Product{}, total, nil
	}
	if end > len(products) {
		end = len(products)
	}

	return products[start:end], total, nil
} 
