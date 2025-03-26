package model

type Category struct {
	Base
	Name        string     `gorm:"size:50;not null" json:"name"`
	Description string     `gorm:"size:200" json:"description"`
	ParentID    *uint      `gorm:"default:null" json:"parent_id"`
	Parent      *Category  `gorm:"foreignKey:ParentID" json:"parent,omitempty"`
	Children    []Category `gorm:"foreignKey:ParentID" json:"children,omitempty"`
	Products    []Product  `gorm:"many2many:product_categories;" json:"products,omitempty"`
} 
