package helpers

import (
	"fmt"
	"strconv"
	"users_api/src/crypto"
	"users_api/src/errorss"

	"github.com/gin-gonic/gin"
)

type ApiHelper struct {
}

func (ApiHelper) GetIntParam(c *gin.Context, paramName string) int {
	idParam := c.Param(paramName)
	id, err := strconv.Atoi(idParam)
	if err != nil {
		panic(err)
	}
	return id
}

func (ApiHelper) HandleError(c *gin.Context) {
	if err := recover(); err != nil {

		if errResp, ok := err.(errorss.ErrorResponseModel); ok {
			c.JSON(errResp.HttpStatus, err)
		} else {
			fmt.Print(err)
			c.JSON(500, errorss.ErrorResponseModel{HttpStatus: 500, Cause: "Error, intente mas tarde"})
		}

	}
}

func (ApiHelper) GetToken(c *gin.Context) *crypto.TokenModel {
	headerToken := c.Request.Header["Token"]
	if len(headerToken) == 0 {
		panic(errorss.ErrorResponseModel{HttpStatus: 401, Cause: "Token header required"})
	}
	plainToken := headerToken[0]
	token := crypto.ParseToken(plainToken)
	return token
}
