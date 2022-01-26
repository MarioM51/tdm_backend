package users

import (
	"fmt"
	"net/http"
	"strconv"
	"users_api/src/crypto"
	"users_api/src/errorss"

	"github.com/gin-gonic/gin"
)

func getIntParam(c *gin.Context, paramName string) int {
	idParam := c.Param(paramName)
	id, err := strconv.Atoi(idParam)
	if err != nil {
		panic(err)
	}
	return id
}

func showUser(c *gin.Context, user *UserModel) {
	if user != nil {
		user.Password = ""
		c.JSON(http.StatusOK, &user)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
	}
}

func getToken(c *gin.Context) *crypto.TokenModel {
	headerToken := c.Request.Header["Token"]
	if len(headerToken) == 0 {
		panic(errorss.ErrorResponseModel{HttpStatus: 401, Cause: "Token header required"})
	}
	plainToken := headerToken[0]
	token := crypto.ParseToken(plainToken)
	return token
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

func checkRol(c *gin.Context, rolToSearch string) bool {
	token := getToken(c)
	userLogged := usrServ.findById(token.IdUser)
	roleFounded := false
	for i := range userLogged.Roles {
		if userLogged.Roles[i].Name == rolToSearch {
			roleFounded = true
			break
		}
	}
	return roleFounded
}
