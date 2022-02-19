package product

type ProductModel struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Image       string  `json:"image"`
	Description string  `json:"description"`
}
