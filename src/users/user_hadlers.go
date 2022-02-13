package users

import (
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
var unauthUserMsg = errorss.ErrorResponseModel{HttpStatus: 403, Cause: "User Unauthorized"}

func (_ UserHadler) getAll(c *gin.Context) {
	defer handleError(c)
	if !checkRol(c, "admin") {
		panic(unauthUserMsg)
	}

	allUsers := usrServ.findAll()
	c.JSON(http.StatusOK, allUsers)
}

func (_ UserHadler) add(c *gin.Context) {
	defer handleError(c)

	newUser := getUser(c)
	userAdded := usrServ.save(*newUser)
	c.JSON(http.StatusCreated, userAdded)
}

func (_ UserHadler) getById(c *gin.Context) {
	defer handleError(c)
	if !checkRol(c, "admin") {
		panic(unauthUserMsg)
	}

	id := getIntParam(c, "id")
	userFinded := usrServ.findById(uint(id))
	showUser(c, userFinded)
}

func (_ UserHadler) update(c *gin.Context) {
	defer handleError(c)
	if !checkRol(c, "admin") {
		panic(unauthUserMsg)
	}

	newInfo := getUser(c)
	if newInfo.ID <= 0 {
		panic(errorss.ErrorResponseModel{HttpStatus: 400, Cause: "Id required"})
	}
	userUpdated := usrServ.update(newInfo)
	showUser(c, userUpdated)

}

func (_ UserHadler) deleteById(c *gin.Context) {
	defer handleError(c)
	if !checkRol(c, "admin") {
		panic(unauthUserMsg)
	}

	id := getIntParam(c, "id")
	userFinded := usrServ.delete(uint(id))
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
		c.JSON(http.StatusOK, map[string]string{
			"msg": "user activated",
		})
	}

}

func (_ UserHadler) login(c *gin.Context) {
	defer handleError(c)
	toLoggin := getUser(c)

	if toLoggin.Password == "" || toLoggin.Email == "" {
		c.JSON(400, "Email and password are required")
		return
	}

	token, user := usrServ.login(toLoggin)

	c.Writer.Header().Set("Token", token)
	showUser(c, user)
}
