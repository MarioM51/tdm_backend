package users

import (
	"users_api/src/logs"

	"github.com/gin-gonic/gin"
)

var userHandler IUserHadler = UserHadler{}
var Logger = logs.New(false)

func AddRoutes(r *gin.Engine) {
	r.GET("/users", userHandler.getAll)
	r.POST("/users", userHandler.add)
	r.PUT("/users", userHandler.update)
	r.GET("/users/:id", userHandler.getById)
	r.DELETE("/users/:id", userHandler.deleteById)
	r.GET("/users/:id/activate/:code", userHandler.activate)
	r.POST("/users/login", userHandler.login)
}
