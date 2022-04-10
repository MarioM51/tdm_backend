package product

import "strconv"

type ProductModel struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Image       string `json:"image"`
	Description string `json:"description"`
}

/*
type params struct {
	Val ProductsJSONLD
}
*/

func ProductModelToArrayJSONLD(products []ProductModel) ProductsWrapperJSONLD {

	var productListA = []ProductJSONLD{}
	for _, p := range products {
		productListA = append(productListA, ProductJSONLD{
			Type:        "Product",
			Identifier:  strconv.Itoa(p.Id),
			Url:         "/nothing",
			Image:       p.Image,
			Name:        p.Name,
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
	Type        string `json:"@type"`
	Identifier  string `json:"identifier"`
	Url         string `json:"url"`
	Image       string `json:"image"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Offer       Offer  `json:"offers"`
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
