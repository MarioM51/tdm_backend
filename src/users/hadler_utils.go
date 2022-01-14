package users

import (
	"net/http"
	"strconv"

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
		c.JSON(http.StatusOK, &user)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
	}
}
