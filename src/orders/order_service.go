package orders

import (
	"time"
	"users_api/src/errorss"
	"users_api/src/users"
)

type OrderService struct {
}

const newOrderMessage = `<html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
	</head>
	<body>
		<p>Nueva orden agregada</p>
	</body>
</html>`

const acceptedOrder = `<html>
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
</head>
<body>
	<p>Su pedido ha sido aceptado, siga los para pagar</p>
</body>
</html>`

var orderRepo = OrderRepository{}
var usrServ users.IUserService = users.UserService{}

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

func (OrderService) confirm(idOrder int, idUser uint) *Order {
	user := usrServ.FindById(idUser)

	if !user.CanConfirmOrder() {
		panic(errorss.ErrorResponseModel{HttpStatus: 403, Cause: "User needs to add aditional info"})
	}

	confirmed := orderRepo.updateField(idOrder, "confirmed_at", time.Now())

	if confirmed.IdUser == 1 {
		orderRepo.updateField(idOrder, "id_user", idUser)
	}

	err := emailSender.SendEmail("test@sales.com", "Se agrego un nuevo pedido", newOrderMessage)
	if err != nil {
		logger.Print("Error email not sended: " + err.Error())
	}

	return confirmed
}

func (OrderService) accept(idOrder int, idUser uint) *Order {
	user := usrServ.FindById(idUser)

	accepted := orderRepo.updateField(idOrder, "accepted_at", time.Now())

	err := emailSender.SendEmail(user.Email, "Pedido aceptado", acceptedOrder)
	if err != nil {
		logger.Print("Error email not sended: " + err.Error())
	}

	return accepted
}

func (OrderService) findAll() *[]Order {
	all := orderRepo.findAll()
	return all
}

func (OrderService) findByUserId(idUser uint) *[]Order {
	ordersOfuser := orderRepo.findByUserId(idUser)
	return ordersOfuser
}
