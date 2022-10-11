package orders

import (
	"log"
	"users_api/src/helpers"

	"github.com/gin-gonic/gin"
)

var dbHelper *helpers.DBHelper = nil
var emailSender *helpers.EmailSender = nil
var logger *log.Logger = nil
var constants helpers.Constants

//
var apiHelper = helpers.ApiHelper{}
var orderServ = OrderService{}

func LinkDependencies(db *helpers.DBHelper, emailSenderIn *helpers.EmailSender, loggerIn *log.Logger, constantsIn helpers.Constants) {
	dbHelper = db
	emailSender = emailSenderIn
	logger = loggerIn
	constants = constantsIn
}

func AddApiRoutes(r *gin.Engine, prefix string) {

	orderHanderApi := OrderHandlerApi{}

	r.POST(prefix+"/orders", orderHanderApi.save)

	r.POST(prefix+"/orders/find", orderHanderApi.findByIds)

	r.DELETE(prefix+"/orders/:id", orderHanderApi.deleteById)

	r.GET(prefix+"/orders/findByUserLogged", orderHanderApi.findByUserLogged)

	r.PUT(prefix+"/orders/:id/confirm", orderHanderApi.confirm)

	r.PUT(prefix+"/orders/:id/accept", orderHanderApi.accept)

	r.GET(prefix+"/orders", orderHanderApi.findAll)

	r.GET(prefix+"/orders/paymentInfo", orderHanderApi.getPaymentInfo)

}
