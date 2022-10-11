package product

import (
	"strconv"
	"time"
)

type ProductModel struct {
	ID          int          `json:"id" gorm:"primaryKey"`
	Name        string       `json:"name"`
	Price       int          `json:"price"`
	Description string       `json:"description" gorm:"size:60"`
	Likes       int          `json:"likes" gorm:"-:all"`
	Image       ProductImage `json:"image,omitempty" gorm:"foreignKey:ID"`
}

type ProductImage struct {
	ID        int       `json:"id_product" gorm:"primaryKey"`
	MimeType  string    `json:"mime_type,omitempty" gorm:"size:15"`
	Base64    string    `json:"base64,omitempty"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

/* Like */

type LikeProduct struct {
	FkProduct int       `json:"fk_product"`
	FKUser    int       `json:"fk_user"`
	CreatedAt time.Time `json:"created_at"`
}

/* JSON-LD */

func ProductModelToArrayJSONLD(products []ProductModel) ProductsWrapperJSONLD {

	var productListA = []ProductJSONLD{}
	for _, p := range products {
		productListA = append(productListA, ProductJSONLD{
			Type:           "Product",
			Identifier:     strconv.Itoa(p.ID),
			Url:            "/nothing",
			Image:          "/nothing_yet",
			ImageUpdatedAt: p.Image.UpdatedAt.Format(time.RFC3339),
			Name:           p.Name,
			Description:    p.Description,
			Offer:          Offer{Type: "Offer", Price: strconv.Itoa(p.Price), PriceCurrency: "MXN"},
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
	Type           string `json:"@type"`
	Identifier     string `json:"identifier"`
	Url            string `json:"url"`
	Image          string `json:"image"`
	ImageUpdatedAt string `json:"image_updated_at"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Offer          Offer  `json:"offers"`
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
