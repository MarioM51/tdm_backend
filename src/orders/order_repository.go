package orders

import (
	"users_api/src/errorss"
	"users_api/src/helpers"
	"users_api/src/product"
)

type OrderRepository struct {
}

var dbHelper = helpers.DBHelper{}

func CreateOrderSchema() {
	dbHelper.Connect()

}

func (dbh OrderRepository) save(newOder *Order) *Order {

	tx := dbHelper.DB.Create(newOder)
	if tx.Error != nil {
		panic(errorss.ErrorResponseModel{HttpStatus: 500, Cause: "Error saving order"})
	}

	for i := range newOder.Products {
		newOrderProduct := newOder.Products[i]

		newOrderProduct.IdOrder = newOder.Id
		tx = dbHelper.DB.Create(&newOrderProduct)
		if tx.Error != nil {
			panic(errorss.ErrorResponseModel{HttpStatus: 500, Cause: "Error saving order products"})
		}
	}

	return newOder
}

func (dbh OrderRepository) findById(id int) (order *Order) {

	dbHelper.DB.First(&order, id)
	if order.Id == 0 {
		return nil
	}

	dbh.findOrderProducts(order)

	return order
}

func (dbh OrderRepository) findByIds(ids []int) (orders *[]Order) {

	tx := dbHelper.DB.Find(&orders, ids)
	if tx.Error != nil {
		panic(errorss.ErrorResponseModel{HttpStatus: 500, Cause: "Error finding orders"})
	}
	for i := range *orders {
		dbh.findOrderProducts(&(*orders)[i])
	}

	return orders
}

func (dbh OrderRepository) findAll() (orders *[]Order) {
	tx := dbHelper.DB.Find(&orders)
	if tx.Error != nil {
		panic(errorss.ErrorResponseModel{HttpStatus: 500, Cause: "Error finding orders"})
	}
	for i := range *orders {
		dbh.findOrderProducts(&(*orders)[i])
	}
	return orders
}

func (dbh OrderRepository) findOrderProducts(order *Order) {

	var orderProduct []OrderProduct
	dbHelper.DB.Find(&orderProduct, order.Id)
	order.Products = orderProduct

	for j := range order.Products {
		var product product.ProductModel
		dbHelper.DB.Find(&product, order.Products[j].IdProduct)
		order.Products[j].Name = product.Name
		order.Products[j].Price = product.Price
	}
}

func (dbh OrderRepository) deleteById(id int) {
	tx := dbHelper.DB.Delete(&Order{}, id)
	if tx.Error != nil {
		panic(errorss.ErrorResponseModel{HttpStatus: 500, Cause: "Error deleting order"})
	}

	tx = dbHelper.DB.Delete(&OrderProduct{}, id)
	if tx.Error != nil {
		panic(errorss.ErrorResponseModel{HttpStatus: 500, Cause: "Error deleting order products"})
	}
}

func (dbh OrderRepository) updateField(idOrder int, property string, value interface{}) (o *Order) {
	dbHelper.DB.First(&o, idOrder)
	if o.Id <= 0 {
		panic(errorss.ErrorResponseModel{HttpStatus: 404, Cause: "Order not found"})
	}

	tx1 := dbHelper.DB.Model(&o).Where("id = ?", o.Id).Update(property, value)
	if tx1.Error != nil {
		panic(errorss.ErrorResponseModel{HttpStatus: 500, Cause: "Error updating " + property + " order products"})
	}

	return o
}

func (dbh OrderRepository) findByUserId(idOrder uint) (orders *[]Order) {
	tx1 := dbHelper.DB.Where("id_user = ?", idOrder).Find(&orders)
	if tx1.Error != nil {
		panic(errorss.ErrorResponseModel{HttpStatus: 500, Cause: "Error finding orders by user"})
	}

	for i := range *orders {
		dbh.findOrderProducts(&(*orders)[i])
	}

	return orders
}
