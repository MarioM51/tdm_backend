package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IUserHadler interface {
	getAll(c *gin.Context)
	add(c *gin.Context)
	getById(c *gin.Context)
	update(c *gin.Context)
	deleteById(c *gin.Context)
}

type UserHadler struct {
}

var usrServ IUserService = UserService{}

func (_ UserHadler) getAll(c *gin.Context) {
	allUsers := usrServ.findAll()
	c.JSON(http.StatusOK, allUsers)
}

func (_ UserHadler) add(c *gin.Context) {
	var newUser UserModel
	fmt.Println("Hadler user, Add user" + newUser.string())
	if err := c.BindJSON(&newUser); err != nil {
		panic(err)
	}
	userAdded := usrServ.saveUser(newUser)
	c.JSON(http.StatusCreated, userAdded)
}

func (_ UserHadler) getById(c *gin.Context) {
	id := getIntParam(c, "id")

	userFinded := usrServ.findUserById(uint(id))
	showUser(c, userFinded)
}

func (_ UserHadler) update(c *gin.Context) {
	var newInfo UserModel
	if err := c.BindJSON(&newInfo); err != nil {
		panic(err)
	}
	userUpdated := usrServ.updateUser(newInfo)
	showUser(c, userUpdated)
}

func (_ UserHadler) deleteById(c *gin.Context) {
	id := getIntParam(c, "id")
	userFinded := usrServ.deleteUser(uint(id))
	showUser(c, userFinded)
}
