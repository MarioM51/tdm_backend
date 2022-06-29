package orders

import (
	"net/http"
	"users_api/src/errorss"
	"users_api/src/helpers"

	"github.com/gin-gonic/gin"
)

type OrderHandlerApi struct {
}

var apiHelper = helpers.ApiHelper{}
var orderServ = OrderService{}

func (OrderHandlerApi) save(c *gin.Context) {
	defer apiHelper.HandleError(c)

	/*
		token := apiHelper.GetRequiredToken(c)
		if !usrServ.CheckRol([]string{"blogs", "admin"}, token) {
			panic(errorss.UnAuthUser)
		}
	*/
	newOrder := getOrderFromRequest(c)

	token := apiHelper.GetOptionalToken(c)
	if token.IdUser <= 0 {
		newOrder.IdUser = 1
	} else {
		newOrder.IdUser = int(token.IdUser)
	}

	newOrder.Id = 0
	for i := range newOrder.Products {
		newOrder.Products[i].IdOrder = 0
	}

	savedOrder := orderServ.save(newOrder)
	showOrder(c, savedOrder)
}

func (OrderHandlerApi) findByIds(c *gin.Context) {
	defer apiHelper.HandleError(c)

	var ids []int
	if err := c.BindJSON(&ids); err != nil {
		panic(errorss.ErrorResponseModel{HttpStatus: 400, Cause: "bad request, array of integers required"})
	}

	orders := orderServ.findByIds(ids)

	showOrders(c, orders)

}

func (OrderHandlerApi) deleteById(c *gin.Context) {
	defer apiHelper.HandleError(c)

	id := apiHelper.GetIntParam(c, "id")
	orderDeleted := orderServ.deleteById(id)

	showOrder(c, orderDeleted)
}

func (OrderHandlerApi) confirm(c *gin.Context) {
	defer apiHelper.HandleError(c)

	id := apiHelper.GetIntParam(c, "id")

	token := apiHelper.GetRequiredToken(c)

	confirmed := orderServ.confirm(id, token.IdUser)

	showOrder(c, confirmed)
}

func (OrderHandlerApi) accept(c *gin.Context) {
	defer apiHelper.HandleError(c)

	id := apiHelper.GetIntParam(c, "id")

	token := apiHelper.GetRequiredToken(c)
	if !usrServ.CheckRol([]string{"admin"}, token) {
		panic(errorss.UnAuthUser)
	}

	accepted := orderServ.accept(id, token.IdUser)

	showOrder(c, accepted)
}

func (OrderHandlerApi) findAll(c *gin.Context) {
	defer apiHelper.HandleError(c)

	token := apiHelper.GetRequiredToken(c)
	if !usrServ.CheckRol([]string{"admin"}, token) {
		panic(errorss.UnAuthUser)
	}

	all := orderServ.findAll()
	showOrders(c, all)
}

func (OrderHandlerApi) findByUserLogged(c *gin.Context) {
	defer apiHelper.HandleError(c)
	token := apiHelper.GetRequiredToken(c)
	idUser := token.IdUser

	orderFinded := orderServ.findByUserId(idUser)
	showOrders(c, orderFinded)
}

//==========

func getOrderFromRequest(c *gin.Context) (o *Order) {
	if err := c.ShouldBindJSON(&o); err != nil {
		panic(errorss.ErrorResponseModel{HttpStatus: 400, Cause: "bad order json format"})
	}

	return o
}

func showOrder(c *gin.Context, p *Order) {
	if p != nil {
		c.JSON(http.StatusOK, &p)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "order not found"})
	}
}

func showOrders(c *gin.Context, orders *[]Order) {
	if orders == nil || len(*orders) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "order not found"})
		return
	}

	c.JSON(http.StatusOK, orders)
}
