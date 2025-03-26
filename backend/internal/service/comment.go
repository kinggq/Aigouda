package service

import (
	"aigouda/internal/model"
	"aigouda/internal/repository"
	"errors"
	"gorm.io/gorm"
)

type CommentService struct{}

func NewCommentService() *CommentService {
	return &CommentService{}
}

type CreateCommentRequest struct {
	ProductID uint   `json:"product_id"`
	Content   string `json:"content"`
	Rating    int    `json:"rating"`
}

// Create 创建评论
func (s *CommentService) Create(userID uint, req *CreateCommentRequest) (*model.Comment, error) {
	// 检查商品是否存在
	var product model.Product
	if err := repository.DB.First(&product, req.ProductID).Error; err != nil {
		return nil, err
	}

	comment := &model.Comment{
		ProductID: req.ProductID,
		UserID:    userID,
		Content:   req.Content,
		Rating:    req.Rating,
	}

	if err := repository.DB.Create(comment).Error; err != nil {
		return nil, err
	}

	// 更新商品销量
	if err := repository.DB.Model(&product).UpdateColumn("sales", gorm.Expr("sales + ?", 1)).Error; err != nil {
		return nil, err
	}

	return comment, nil
}

// Delete 删除评论
func (s *CommentService) Delete(id uint, userID uint) error {
	var comment model.Comment
	if err := repository.DB.First(&comment, id).Error; err != nil {
		return err
	}

	// 检查是否是评论作者
	if comment.UserID != userID {
		return errors.New("无权删除此评论")
	}

	return repository.DB.Delete(&comment).Error
}

// List 获取商品评论列表
func (s *CommentService) List(productID uint, page, pageSize int) ([]model.Comment, int64, error) {
	var comments []model.Comment
	var total int64

	// 获取总数
	if err := repository.DB.Model(&model.Comment{}).Where("product_id = ?", productID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	offset := (page - 1) * pageSize
	if err := repository.DB.Where("product_id = ?", productID).
		Offset(offset).Limit(pageSize).
		Order("created_at DESC").
		Find(&comments).Error; err != nil {
		return nil, 0, err
	}

	return comments, total, nil
}

// GetByID 获取评论详情
func (s *CommentService) GetByID(id uint) (*model.Comment, error) {
	var comment model.Comment
	if err := repository.DB.First(&comment, id).Error; err != nil {
		return nil, err
	}
	return &comment, nil
} 
