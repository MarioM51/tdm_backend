package orders

import (
	"github.com/gin-gonic/gin"
)

func AddApiRoutes(r *gin.Engine, prefix string) {

	orderHanderApi := OrderHandlerApi{}

	r.POST(prefix+"/orders", orderHanderApi.save)

	r.POST(prefix+"/orders/find", orderHanderApi.findByIds)

	r.DELETE(prefix+"/orders/:id", orderHanderApi.deleteById)

	r.GET(prefix+"/orders/findByUserLogged", orderHanderApi.findByUserLogged)

	r.PUT(prefix+"/orders/:id/confirm", orderHanderApi.confirm)

	r.PUT(prefix+"/orders/:id/accept", orderHanderApi.accept)

	r.GET(prefix+"/orders", orderHanderApi.findAll)

}
