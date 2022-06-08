package orders

import (
	"github.com/gin-gonic/gin"
)

func AddApiRoutes(r *gin.Engine, prefix string) {

	orderHanderApi := OrderHandlerApi{}

	r.POST(prefix+"/orders", orderHanderApi.save)

	r.POST(prefix+"/orders/find", orderHanderApi.findByIds)

	r.DELETE(prefix+"/orders/:id", orderHanderApi.deleteById)

}
