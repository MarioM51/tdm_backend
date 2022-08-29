package product

import (
	"time"
)

type ProductModel struct {
	ID           int            `json:"id" gorm:"primaryKey"`
	Name         string         `json:"name" gorm:"size:60"`
	Price        int            `json:"price"`
	Description  string         `json:"description" gorm:"size:160"`
	Likes        int            `json:"likes" gorm:"-:all"`
	Images       []ProductImage `json:"images,omitempty" gorm:"foreignKey:ID"`
	Comments     []Comment      `json:"comments,omitempty" gorm:"-:all"`
	OnHomeScreen time.Time      `json:"onHomeScreen,omitempty"`
	/*
		category
			material
		weight: = KGM
		: {
			"@type": "quantitativeValue"
			"unitCode": "KGM"
			"value": 0.50
		}
			width
			height
			depth

	*/
}

func (ProductModel) TableName() string {
	return "products"
}

type ProductImage struct {
	ID        int       `json:"id_image" gorm:"primaryKey"`
	FkProduct int       `json:"fk_product,omitempty"`
	MimeType  string    `json:"mime_type,omitempty" gorm:"size:15"`
	Base64    string    `json:"base64,omitempty"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (ProductImage) TableName() string {
	return "product_images"
}

/* Like */

type LikeProduct struct {
	FkProduct int       `json:"fk_product"`
	FKUser    int       `json:"fk_user"`
	CreatedAt time.Time `json:"created_at"`
}

func (LikeProduct) TableName() string {
	return "product_likes"
}

type Comment struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	IdUser    int       `json:"idUser"`
	IdTarget  int       `json:"idTarget"`
	Content   string    `json:"content"`
	Stars     int       `json:"stars"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"-"`
}

func (Comment) TableName() string {
	return "product_comments"
}
