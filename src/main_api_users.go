package main

import (
	"time"
	"users_api/src/blog"
	"users_api/src/helpers"
	"users_api/src/product"
	"users_api/src/users"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	//Setup database
	users.CreateUserSchema()
	product.CreateProductSchema()
	blog.CreateBlogSchema()

	router := gin.Default()

	//Setup static server
	router.Static("/admin", "./public/admin-spa")
	router.Static("/static", "./public/static")

	//Setup Cors
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Content-Type, Content-Length, Accept-Encoding, Token, accept, origin, X-Requested-With"},
		ExposeHeaders: []string{"Content-Type, Content-Length, Accept-Encoding, Token, accept, origin"},

		MaxAge: 12 * time.Hour,
	}))

	//Setup Api
	const apiPrefix string = "/api"
	users.AddApiRoutes(router, apiPrefix)
	product.AddApiRoutes(router, apiPrefix)
	blog.AddApiRoutes(router, apiPrefix)

	//Setup server side rendering
	templatesM := []helpers.TemplateModel{}
	product.AddSsrRoutes(router, &templatesM)
	router.HTMLRender = helpers.CreateHTMLRenderHelper(templatesM)

	router.Run("localhost:8081")
}
