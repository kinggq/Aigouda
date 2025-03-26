package model

type Product struct {
	Base
	Title         string        `gorm:"size:100;not null" json:"title"`
	Price         float64       `gorm:"type:decimal(10,2);not null" json:"price"`
	OriginalPrice float64       `gorm:"type:decimal(10,2)" json:"original_price"`
	Discount      float64       `gorm:"type:decimal(10,2)" json:"discount"`
	MainImage     string        `gorm:"size:255" json:"main_image"`
	Stock         int           `gorm:"default:0" json:"stock"`
	CategoryID    uint          `gorm:"not null" json:"category_id"`
	Category      Category      `gorm:"foreignKey:CategoryID" json:"category"`
	Images        []ProductImage `gorm:"foreignKey:ProductID" json:"images"`
	Tags          []Tag         `gorm:"many2many:product_tag_relations;" json:"tags"`
	Sales         int           `gorm:"default:0" json:"sales"`
	Comments      []Comment     `gorm:"foreignKey:ProductID" json:"comments"`
}

type ProductImage struct {
	Base
	ProductID uint   `gorm:"not null" json:"product_id"`
	ImageURL  string `gorm:"size:255;not null" json:"image_url"`
}

type Comment struct {
	Base
	ProductID uint   `gorm:"not null" json:"product_id"`
	UserID    uint   `gorm:"not null" json:"user_id"`
	Content   string `gorm:"size:500" json:"content"`
	Rating    int    `gorm:"type:tinyint" json:"rating"`
}
