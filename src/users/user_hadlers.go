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
	login(c *gin.Context)
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

	newUser := getUser(c)
	userAdded := usrServ.saveUser(*newUser)
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

	//TODO: Usar validacion
	validteToken(c)

	newInfo := getUser(c)
	if newInfo.ID <= 0 {
		panic(errorss.ErrorResponseModel{HttpStatus: 400, Cause: "El Id is obligatorio"})
	}
	userUpdated := usrServ.updateUser(newInfo)
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
	var err *errorss.ErrorResponseModel = usrServ.activate(uint(id), code)
	if err != nil {
		c.JSON(err.HttpStatus, err)
	} else {
		c.JSON(http.StatusOK, "")
	}

}

func (_ UserHadler) login(c *gin.Context) {
	defer handleError(c)
	toLoggin := getUser(c)

	if toLoggin.Password == "" || toLoggin.Email == "" {
		c.JSON(400, "Email and password are required")
		return
	}

	token := usrServ.login(toLoggin)

	c.JSON(200, map[string]string{
		"token": token,
	})
}

func handleError(c *gin.Context) {
	if err := recover(); err != nil {

		if errResp, ok := err.(errorss.ErrorResponseModel); ok {
			c.JSON(errResp.HttpStatus, err)
		} else {
			fmt.Print(err)
			c.JSON(500, errorss.ErrorResponseModel{HttpStatus: 500, Cause: "Error, intente mas tarde"})
		}

	}
}

func getUser(c *gin.Context) (newInfo *UserModel) {
	if err := c.BindJSON(&newInfo); err != nil {
		panic(err)
	}
	return newInfo
}
