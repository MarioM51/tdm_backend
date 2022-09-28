package product

import (
	"time"
	"users_api/src/errorss"

	"gorm.io/gorm"
)

type ProductModel struct {
	ID             int            `json:"id" gorm:"primaryKey"`
	Name           string         `json:"name" gorm:"size:60"`
	Price          int            `json:"price"`
	Description    string         `json:"description" gorm:"size:160"`
	Likes          int            `json:"likes" gorm:"-:all"`
	Images         []ProductImage `json:"images,omitempty" gorm:"foreignKey:ID"`
	Comments       []Comment      `json:"comments,omitempty" gorm:"-:all"`
	OnHomeScreen   time.Time      `json:"onHomeScreen,omitempty"`
	CommentCount   int            `json:"commentCount" gorm:"-:all"`
	CommentsRating float32        `json:"commentsRating" gorm:"-:all"`
	DeletedAt      gorm.DeletedAt `json:"-"`
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
	ID        int            `json:"id_image" gorm:"primaryKey"`
	FkProduct int            `json:"fk_product,omitempty"`
	MimeType  string         `json:"mime_type,omitempty" gorm:"size:15"`
	Base64    string         `json:"base64,omitempty"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"-"`
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
	Id         int            `json:"id" gorm:"primaryKey"`
	IdUser     int            `json:"idUser"`
	IdTarget   int            `json:"idTarget"`
	Content    string         `json:"content"`
	Stars      int            `json:"stars"`
	CreatedAt  time.Time      `json:"created_at"`
	DeletedAt  gorm.DeletedAt `json:"-"`
	ResponseTo int            `json:"responseTo" gorm:"default:null"`
}

func (c *Comment) cleanAndValidateNewComment(isResponse bool) {
	//the id it's going to be enerated by the database
	c.Id = 0

	if len(c.Content) <= 5 {
		panic(errorss.ErrorResponseModel{HttpStatus: 400, Cause: "content: must have at leat 5 characters"})
	}

	if c.IdUser <= 0 {
		panic(errorss.ErrorResponseModel{HttpStatus: 400, Cause: "idUser: required"})
	}

	if c.IdTarget <= 0 {
		panic(errorss.ErrorResponseModel{HttpStatus: 400, Cause: "IdTarget: required"})
	}

	if isResponse {
		//the response can't have stars
		c.Stars = 0

		if c.ResponseTo <= 0 {
			panic(errorss.ErrorResponseModel{HttpStatus: 400, Cause: "responseTo: required"})
		}

	} else {
		c.ResponseTo = 0

		if c.IdTarget <= 0 {
			panic(errorss.ErrorResponseModel{HttpStatus: 400, Cause: "stats: required"})
		}
	}

}

func (Comment) TableName() string {
	return "product_comments"
}
