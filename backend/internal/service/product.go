package service

import (
	"aigouda/internal/model"
	"aigouda/internal/repository"
	"errors"
)

type ProductService struct {
	productRepo *repository.ProductRepository
}

func NewProductService() *ProductService {
	return &ProductService{
		productRepo: repository.NewProductRepository(),
	}
}

type CreateProductRequest struct {
	Title        string   `json:"title"`
	Price        float64  `json:"price"`
	OriginalPrice float64  `json:"original_price"`
	Discount     float64  `json:"discount"`
	MainImage    string   `json:"main_image"`
	Images       []string `json:"images"`
	Tags         []string `json:"tags"`
}

type UpdateProductRequest struct {
	Title        *string   `json:"title"`
	Price        *float64  `json:"price"`
	OriginalPrice *float64  `json:"original_price"`
	Discount     *float64  `json:"discount"`
	MainImage    *string   `json:"main_image"`
	Images       *[]string `json:"images"`
	Tags         *[]string `json:"tags"`
}

// List 获取商品列表
func (s *ProductService) List(page, pageSize int, search string) ([]model.Product, int64, error) {
	offset := (page - 1) * pageSize
	products, total, err := s.productRepo.List(offset, pageSize, search)
	if err != nil {
		return nil, 0, err
	}
	return products, total, nil
}

// Get 获取单个商品
func (s *ProductService) Get(id uint) (*model.Product, error) {
	product, err := s.productRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// Create 创建商品
func (s *ProductService) Create(product *model.Product) error {
	if product.Title == "" {
		return errors.New("商品标题不能为空")
	}
	if product.Price <= 0 {
		return errors.New("商品价格必须大于0")
	}
	if product.Stock < 0 {
		return errors.New("商品库存不能小于0")
	}
	if product.CategoryID == 0 {
		return errors.New("商品分类不能为空")
	}
	return s.productRepo.Create(product)
}

// Update 更新商品
func (s *ProductService) Update(product *model.Product) error {
	if product.ID == 0 {
		return errors.New("商品ID不能为空")
	}
	if product.Title == "" {
		return errors.New("商品标题不能为空")
	}
	if product.Price <= 0 {
		return errors.New("商品价格必须大于0")
	}
	if product.Stock < 0 {
		return errors.New("商品库存不能小于0")
	}
	if product.CategoryID == 0 {
		return errors.New("商品分类不能为空")
	}

	// 只更新需要修改的字段
	updates := map[string]interface{}{
		"title":          product.Title,
		"price":          product.Price,
		"original_price": product.OriginalPrice,
		"discount":       product.Discount,
		"main_image":     product.MainImage,
		"category_id":    product.CategoryID,
		"stock":          product.Stock,
	}

	return s.productRepo.Update(product.ID, updates)
}

// Delete 删除商品
func (s *ProductService) Delete(id uint) error {
	return s.productRepo.Delete(id)
}

// GetByID 获取商品详情
func (s *ProductService) GetByID(id uint) (*model.Product, error) {
	var product model.Product
	if err := repository.DB.Preload("Images").Preload("Tags").Preload("Category").First(&product, id).Error; err != nil {
		return nil, errors.New("商品不存在")
	}
	return &product, nil
} 
