package product

import (
	"strconv"
	"time"
)

type ProductModel struct {
	ID          int            `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"size:60"`
	Price       int            `json:"price"`
	Description string         `json:"description" gorm:"size:160"`
	Likes       int            `json:"likes" gorm:"-:all"`
	Images      []ProductImage `json:"images,omitempty" gorm:"foreignKey:ID"`
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

/* JSON-LD */

func ProductModelToArrayJSONLD(products []ProductModel) ProductsWrapperJSONLD {

	var productListA = []ProductJSONLD{}
	for _, p := range products {
		for i := range p.Images {
			p.Images[i].Base64 = ""
		}

		productListA = append(productListA, ProductJSONLD{
			Type:        "Product",
			Identifier:  strconv.Itoa(p.ID),
			Url:         "/nothing",
			Images:      p.Images,
			Name:        p.Name,
			Likes:       p.Likes,
			Description: p.Description,
			Offer:       Offer{Type: "Offer", Price: strconv.Itoa(p.Price), PriceCurrency: "MXN"},
		})
	}

	return ProductsWrapperJSONLD{
		Val: ProductsJSONLD{
			Context:         "https://schema.org",
			Type:            "ItemList",
			NumberOfItems:   strconv.Itoa(len(products)),
			ItemListElement: productListA,
		},
	}

}

type ProductsWrapperJSONLD struct {
	Val ProductsJSONLD
}

type ProductsJSONLD struct {
	Context         string          `json:"@context"`
	Type            string          `json:"@type"`
	NumberOfItems   string          `json:"numberOfItems"`
	ItemListElement []ProductJSONLD `json:"itemListElement"`
}

type ProductJSONLD struct {
	Type           string         `json:"@type"`
	Identifier     string         `json:"identifier"`
	Url            string         `json:"url"`
	Images         []ProductImage `json:"images"`
	ImageUpdatedAt string         `json:"image_updated_at"`
	Name           string         `json:"name"`
	Description    string         `json:"description"`
	Likes          int            `json:"likes"`
	Offer          Offer          `json:"offers"`
}

type Offer struct {
	Type          string `json:"@type"`
	Price         string `json:"price"`
	PriceCurrency string `json:"priceCurrency"`
}

/*
var productList = []ProductJSONLD{
	{
		Type:        "Product",
		Identifier:  "1",
		Url:         "products/1",
		Image:       "products/1/image",
		Name:        "Tompo de llavero",
		Description: "Trompo pequeño de madera que se puede usar como llavero",
		Offer:       Offer{Type: "Offer", Price: "25", PriceCurrency: "MXN"},
	},
*/
