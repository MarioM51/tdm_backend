package users

import (
	"net/http"
	"users_api/src/errorss"
	"users_api/src/helpers"

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
var apiHelper = helpers.ApiHelper{}

func (UserHadler) getAll(c *gin.Context) {
	defer apiHelper.HandleError(c)

	token := apiHelper.GetRequiredToken(c)
	if !usrServ.CheckRol([]string{"admin"}, token) {
		panic(errorss.UnAuthUser)
	}

	allUsers := usrServ.findAll()
	c.JSON(http.StatusOK, allUsers)
}

func (UserHadler) add(c *gin.Context) {
	defer apiHelper.HandleError(c)

	newUser := getUser(c)
	userAdded := usrServ.save(*newUser)
	c.JSON(http.StatusCreated, userAdded)
}

func (UserHadler) getById(c *gin.Context) {
	defer apiHelper.HandleError(c)

	token := apiHelper.GetRequiredToken(c)
	if !usrServ.CheckRol([]string{"admin"}, token) {
		panic(errorss.UnAuthUser)
	}

	id := apiHelper.GetIntParam(c, "id")
	userFinded := usrServ.findById(uint(id))
	showUser(c, userFinded)
}

func (UserHadler) update(c *gin.Context) {
	defer apiHelper.HandleError(c)

	token := apiHelper.GetRequiredToken(c)
	if !usrServ.CheckRol([]string{"admin"}, token) {
		panic(errorss.UnAuthUser)
	}

	newInfo := getUser(c)
	if newInfo.ID <= 0 {
		panic(errorss.ErrorResponseModel{HttpStatus: 400, Cause: "Id required"})
	}
	userUpdated := usrServ.update(newInfo)
	showUser(c, userUpdated)

}

func (UserHadler) deleteById(c *gin.Context) {
	defer apiHelper.HandleError(c)

	token := apiHelper.GetRequiredToken(c)
	if !usrServ.CheckRol([]string{"admin"}, token) {
		panic(errorss.UnAuthUser)
	}

	id := apiHelper.GetIntParam(c, "id")
	userFinded := usrServ.delete(uint(id))
	showUser(c, userFinded)
}

func (UserHadler) activate(c *gin.Context) {
	defer apiHelper.HandleError(c)

	id := apiHelper.GetIntParam(c, "id")
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

func (UserHadler) login(c *gin.Context) {
	defer apiHelper.HandleError(c)
	toLoggin := getUser(c)

	if toLoggin.Password == "" || toLoggin.Email == "" {
		c.JSON(400, "Email and password are required")
		return
	}

	token, user := usrServ.login(toLoggin)

	c.Writer.Header().Set("Token", token)
	showUser(c, user)
}

func showUser(c *gin.Context, user *UserModel) {
	if user != nil {
		user.Password = ""
		c.JSON(http.StatusOK, &user)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
	}
}

func getUser(c *gin.Context) (newInfo *UserModel) {
	if err := c.BindJSON(&newInfo); err != nil {
		panic(err)
	}
	return newInfo
}
