package users

import (
	"github.com/gin-gonic/gin"
)

var userHandler IUserHadler = UserHadler{}

func AddRoutes(r *gin.Engine) {
	r.GET("/users", userHandler.getAll)
	r.POST("/users", userHandler.add)
	r.PUT("/users", userHandler.update)
	r.GET("/users/:id", userHandler.getById)
	r.DELETE("/users/:id", userHandler.deleteById)
}
