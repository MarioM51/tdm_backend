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
	newOrder.Id = 0
	for i := range newOrder.Products {
		newOrder.Products[i].IdOrder = 0
	}

	savedBlog := orderServ.save(newOrder)
	showOrder(c, savedBlog)
}

func (OrderHandlerApi) findByIds(c *gin.Context) {
	defer apiHelper.HandleError(c)

	var ids []int
	if err := c.BindJSON(&ids); err != nil {
		panic(errorss.ErrorResponseModel{HttpStatus: 400, Cause: "bad request, array of integers required"})
	}

	orders := orderServ.findByIds(ids)

	if orders == nil || len(*orders) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "order not found"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (OrderHandlerApi) deleteById(c *gin.Context) {
	defer apiHelper.HandleError(c)

	id := apiHelper.GetIntParam(c, "id")
	orderDeleted := orderServ.deleteById(id)

	showOrder(c, orderDeleted)
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
