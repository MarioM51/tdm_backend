package orders

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	Id        int            `json:"id"`
	IdUser    int            `json:"idUser"`
	Products  []OrderProduct `json:"products" gorm:"-:all"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (Order) TableName() string {
	return "orders"
}

type OrderProduct struct {
	IdOrder   int            `json:"idOrder"`
	IdProduct int            `json:"idProduct"`
	Amount    int            `json:"amount"`
	Name      string         `json:"name" gorm:"-:all"`
	Price     int            `json:"unitaryPrice" gorm:"-:all"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (OrderProduct) TableName() string {
	return "orders_products"
}
