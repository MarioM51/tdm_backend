package orders

import "users_api/src/errorss"

type OrderService struct {
}

var orderRepo = OrderRepository{}

func (OrderService) save(newOder *Order) *Order {
	orderSaved := orderRepo.save(newOder)

	return orderSaved
}

func (OrderService) findById(id int) *Order {
	orderFinded := orderRepo.findById(id)
	if orderFinded == nil {
		panic(errorss.ErrorResponseModel{HttpStatus: 404, Cause: "order not found"})
	}

	return orderFinded
}

func (OrderService) findByIds(ids []int) *[]Order {
	ordersFinded := orderRepo.findByIds(ids)

	return ordersFinded
}

func (oServ OrderService) deleteById(id int) *Order {
	del := oServ.findById(id)
	orderRepo.deleteById(id)
	return del
}
