package model

import "gorm.io/gorm"

// Tag 标签模型
type Tag struct {
	gorm.Model
	Name     string    `gorm:"size:50;not null;uniqueIndex" json:"name"`
	Products []Product `gorm:"many2many:product_tags;" json:"products,omitempty"`
} 
