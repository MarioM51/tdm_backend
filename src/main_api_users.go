package main

import (
	"users_api/src/users"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	users.InitDB()

	users.AddRoutes(router)
	router.Run("localhost:8080")
}
