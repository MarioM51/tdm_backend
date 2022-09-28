package product

import (
	"fmt"
	"strconv"
	"time"
)

type ProductJSONLD struct {
	Context         string          `json:"@context,omitempty"`
	Type            string          `json:"@type"`
	Identifier      string          `json:"identifier"`
	Url             string          `json:"url"`
	Images          []string        `json:"image"`
	Name            string          `json:"name"`
	Description     string          `json:"description"`
	Likes           int             `json:"likes"`
	Comments        []CommentJSONLD `json:"review"`
	Offer           Offer           `json:"offers"`
	AggregateRating AggregateRating `json:"aggregateRating"`
}

func ProductModelToArrayJSONLD(products []ProductModel) ProductsWrapperJSONLD {

	var productListA = []ProductJSONLD{}

	for _, p := range products {
		imagesURLs := []string{}
		for i := range p.Images {
			p.Images[i].Base64 = ""
			urlTemp := fmt.Sprintf(constants.UrlProductImage, p.Images[i].ID, p.Images[i].UpdatedAt.Format(time.RFC3339))
			imagesURLs = append(imagesURLs, urlTemp)
		}

		aggrRating := getAggregateRating(&p)

		productListA = append(productListA, ProductJSONLD{
			Type:            "Product",
			Identifier:      strconv.Itoa(p.ID),
			Url:             "/nothing",
			Images:          imagesURLs,
			Name:            p.Name,
			Likes:           p.Likes,
			Description:     p.Description,
			Offer:           Offer{Type: "Offer", Price: strconv.Itoa(p.Price), PriceCurrency: "MXN"},
			AggregateRating: aggrRating,
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

func ProductModelToJSONLD(p *ProductModel, fromList bool) ProductJSONLD {
	var imgUrls []string
	for i := range p.Images {
		imgUrls = append(imgUrls,
			fmt.Sprintf(constants.UrlProductImage, p.Images[i].ID, p.Images[i].UpdatedAt.Format(time.RFC3339)),
		)
	}

	aggrRating := getAggregateRating(p)

	jsonLD := ProductJSONLD{
		Type:        "Product",
		Identifier:  strconv.Itoa(p.ID),
		Url:         "/",
		Images:      imgUrls,
		Name:        p.Name,
		Likes:       p.Likes,
		Description: p.Description,
		Offer: Offer{
			Type:          "Offer",
			Price:         strconv.Itoa(p.Price),
			PriceCurrency: "MXN",
			Url:           "/",
			Availability:  "https://schema.org/PreOrder",
		},
		AggregateRating: aggrRating,
	}

	if !fromList {
		jsonLD.Context = "https://schema.org"
	}

	commentsJsonLD := []CommentJSONLD{}
	for i := range p.Comments {
		reviewTemp := fromComment(p.Comments[i])
		commentsJsonLD = append(commentsJsonLD, reviewTemp)
	}
	jsonLD.Comments = commentsJsonLD

	return jsonLD

}

func getAggregateRating(p *ProductModel) AggregateRating {

	var aggrRating = AggregateRating{
		Type:        "AggregateRating",
		BestRating:  5,
		WorstRating: 1,
		RatingValue: p.CommentsRating,
		RatingCount: uint16(p.CommentCount),
	}

	return aggrRating

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

type Offer struct {
	Type          string `json:"@type"`
	Price         string `json:"price"`
	PriceCurrency string `json:"priceCurrency"`
	Url           string `json:"url"`
	Availability  string `json:"Availability"`
	//priceValidUntil: fecha asta donde es valido el precio

}

type CommentJSONLD struct {
	Type          string       `json:"@type"`
	Identifier    string       `json:"identifier"`
	IdUser        int          `json:"idUser"`
	Text          string       `json:"text"`
	DatePublished string       `json:"datePublished"`
	ReviewRating  ReviewRating `json:"reviewRating"`
	Author        Author       `json:"author"`
	ResponseTo    int          `json:"responseTo"`
}

type Author struct {
	Type       string `json:"@type"`
	Name       string `json:"name"`
	Identifier string `json:"identifier"`
}

type ReviewRating struct {
	Type        string `json:"@type"`
	BestRating  uint8  `json:"bestRating"`
	RatingValue uint8  `json:"ratingValue"`
	WorstRating uint8  `json:"worstRating"`
}

type AggregateRating struct {
	Type        string  `json:"@type"`
	BestRating  uint8   `json:"bestRating"`
	RatingValue float32 `json:"ratingValue"`
	RatingCount uint16  `json:"ratingCount"`
	WorstRating uint8   `json:"worstRating"`
}

func fromComment(c Comment) CommentJSONLD {
	reviewTemp := CommentJSONLD{
		Type: "review",
		ReviewRating: ReviewRating{
			Type:        "Rating",
			BestRating:  5,
			WorstRating: 1,
			RatingValue: uint8(c.Stars),
		},
		Author: Author{
			Type:       "Person",
			Name:       fmt.Sprint(c.IdUser),
			Identifier: fmt.Sprint(c.IdUser),
		},
		Identifier:    fmt.Sprint(c.Id),
		IdUser:        c.IdUser,
		Text:          c.Content,
		DatePublished: c.CreatedAt.String(),
		ResponseTo:    c.ResponseTo,
	}
	return reviewTemp
}
