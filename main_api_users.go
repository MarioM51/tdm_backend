package main

import (
	"os"
	"strings"
	"time"
	"users_api/src/blog"
	"users_api/src/helpers"
	"users_api/src/home"
	"users_api/src/orders"
	"users_api/src/product"
	"users_api/src/users"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nanmu42/gzip"
)

func main() {
	//Setup database
	users.CreateUserSchema()
	product.CreateProductSchema()
	blog.CreateBlogSchema()
	orders.CreateOrderSchema()

	router := gin.Default()
	gin.DisableConsoleColor()
	router.Use(gzip.DefaultHandler().Gin)
	gin.DefaultWriter = os.Stdout

	//Setup static server
	router.Use(func() gin.HandlerFunc {
		return func(c *gin.Context) {
			if strings.HasPrefix(c.Request.URL.Path, "/admin") || strings.HasPrefix(c.Request.URL.Path, "/static") {
				c.Writer.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
			}
		}
	}(),
	)

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
	users.AddApiRoutes(router, helpers.API_PREFIX)
	product.AddApiRoutes(router, helpers.API_PREFIX)
	blog.AddApiRoutes(router, helpers.API_PREFIX)
	orders.AddApiRoutes(router, helpers.API_PREFIX)

	//Setup server side rendering
	templates := []helpers.TemplateModel{}
	//common templates
	commonTemplates := []helpers.TemplateModel{
		{
			Name:       "not-found",
			PagePath:   "templates/common/not-found.gohtml",
			LayoutPath: helpers.LAYOUT_A,
		},
		{
			Name:       "fatal-error",
			PagePath:   "templates/common/fatal-error.gohtml",
			LayoutPath: helpers.LAYOUT_A,
		},
	}
	templates = append(templates, commonTemplates...)

	//specific templates
	product.AddSsrRoutes(router, &templates)
	blog.AddSsrRoutes(router, &templates)
	home.AddSsrRoutes(router, &templates)
	router.HTMLRender = helpers.CreateHTMLRenderHelper(templates)

	router.Run(helpers.DOMAIN + ":" + helpers.PORT)
}
