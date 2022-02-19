package main

import (
	"time"
	"users_api/src/product"
	"users_api/src/users"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	users.CreateUserSchema()
	product.CreateProductSchema()

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Content-Type, Content-Length, Accept-Encoding, Token, accept, origin, X-Requested-With"},
		ExposeHeaders: []string{"Content-Type, Content-Length, Accept-Encoding, Token, accept, origin"},

		MaxAge: 12 * time.Hour,
	}))

	users.AddRoutes(router)
	product.AddRoutes(router)

	router.Run("localhost:8081")
}
