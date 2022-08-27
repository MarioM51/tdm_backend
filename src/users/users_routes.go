package users

import (
	"users_api/src/helpers"

	"github.com/gin-gonic/gin"
)

var userHandler IUserHadler = UserHadler{}

var dbHelper *helpers.DBHelper = nil

func LinkDependencies(db *helpers.DBHelper) {
	dbHelper = db
}

func AddApiRoutes(r *gin.Engine, prefix string) {
	r.GET(prefix+"/users", userHandler.getAll)
	r.POST(prefix+"/users", userHandler.add)
	r.PUT(prefix+"/users", userHandler.update)
	r.GET(prefix+"/users/:id", userHandler.getById)
	r.DELETE(prefix+"/users/:id", userHandler.deleteById)
	r.GET(prefix+"/users/:id/activate/:code", userHandler.activate)
	r.POST(prefix+"/users/login", userHandler.login)
}
