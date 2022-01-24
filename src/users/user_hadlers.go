package users

import (
	"fmt"
	"net/http"
	"users_api/src/errorss"

	"github.com/gin-gonic/gin"
)

type IUserHadler interface {
	getAll(c *gin.Context)
	add(c *gin.Context)
	getById(c *gin.Context)
	update(c *gin.Context)
	deleteById(c *gin.Context)
	activate(c *gin.Context)
}

type UserHadler struct {
}

var usrServ IUserService = UserService{}

func (_ UserHadler) getAll(c *gin.Context) {
	defer handleError(c)

	allUsers := usrServ.findAll()
	c.JSON(http.StatusOK, allUsers)
}

func (_ UserHadler) add(c *gin.Context) {
	defer handleError(c)

	var newUser UserModel
	if err := c.BindJSON(&newUser); err != nil {
		panic(err)
	}
	userAdded := usrServ.saveUser(newUser)
	c.JSON(http.StatusCreated, userAdded)
}

func (_ UserHadler) getById(c *gin.Context) {
	defer handleError(c)

	id := getIntParam(c, "id")
	userFinded := usrServ.findUserById(uint(id))
	showUser(c, userFinded)
}

func (_ UserHadler) update(c *gin.Context) {
	defer handleError(c)

	var newInfo UserModel
	if err := c.BindJSON(&newInfo); err != nil {
		panic(err)
	}
	if newInfo.ID <= 0 {
		panic(errorss.ErrorResponseModel{HttpStatus: 400, Cause: "El Id is obligatorio"})
	}

	userUpdated := usrServ.updateUser(&newInfo)
	showUser(c, userUpdated)
}

func (_ UserHadler) deleteById(c *gin.Context) {
	defer handleError(c)

	id := getIntParam(c, "id")
	userFinded := usrServ.deleteUser(uint(id))
	showUser(c, userFinded)
}

func (_ UserHadler) activate(c *gin.Context) {
	defer handleError(c)

	id := getIntParam(c, "id")
	code := c.Param("code")
	err := usrServ.activate(uint(id), code)
	if err != nil {
		c.JSON(err.HttpStatus, err)
	} else {
		c.JSON(http.StatusOK, "")
	}

}

func handleError(c *gin.Context) {
	if err := recover(); err != nil {

		if errResp, ok := err.(*errorss.ErrorResponseModel); ok {
			c.JSON(errResp.HttpStatus, err)
		} else {
			fmt.Print(err)
			c.JSON(500, errorss.ErrorResponseModel{HttpStatus: 500, Cause: "Error, intente mas tarde"})
		}

	}
}
