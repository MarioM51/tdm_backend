package main

import (
	"fmt"
	"log"
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
	//Get enviroment
	constants := helpers.Constants{}
	constants.LoadConstants()

	//Create module dependencies
	logger := log.New(os.Stdout, "\r\n", log.LstdFlags)

	dbHelper := &helpers.DBHelper{}
	dbHelper.Connect(constants, logger)

	var emailSender *helpers.EmailSender = helpers.NewEmailSender(constants)

	//Pass dependencies to modules
	users.LinkDependencies(dbHelper)
	product.LinkDependencies(dbHelper, constants)
	blog.LinkDependencies(dbHelper, constants)
	orders.LinkDependencies(dbHelper, emailSender, logger, constants)
	home.LinkDependencies(&constants)

	if constants.IsLocal() {
		gin.SetMode(gin.DebugMode)
	} else if constants.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	} else {
		panic(fmt.Sprintf("env '%v' not defined, required: [%v]", constants.Env, constants.AvalaibleEnviroments()))
	}

	router := gin.Default()
	gin.DisableConsoleColor()
	//router.Use(gzip.DefaultHandler().Gin)
	gin.DefaultWriter = os.Stdout

	//Setup static server
	router.Use(func() gin.HandlerFunc {
		return func(c *gin.Context) {
			path := c.Request.URL.Path
			isAdminSPA := strings.HasPrefix(path, "/admin")
			isStaticResources := strings.HasPrefix(path, "/static_"+constants.StaticResourcesVersion)
			isOutOfAPI := !strings.HasPrefix(path, "/api")

			if isAdminSPA || isStaticResources {
				//keep first the header adds and then the setup of gzip, otherwise the headers will be deleted
				c.Writer.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
				gzip.DefaultHandler().Gin(c)
			} else if isOutOfAPI {
				gzip.DefaultHandler().Gin(c)
			}
		}
	}(),
	)

	router.Static("/admin", "./public/admin-spa")
	router.Static(constants.StaticFolder, "./public/static_"+constants.StaticResourcesVersion)

	//Setup Cors
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Content-Type, Content-Length, Accept-Encoding, Token, accept, origin, X-Requested-With"},
		ExposeHeaders: []string{"Content-Type, Content-Length, Accept-Encoding, Token, accept, origin"},

		MaxAge: 12 * time.Hour,
	}))

	//Setup Api
	users.AddApiRoutes(router, constants.ApiPrefix)
	product.AddApiRoutes(router, constants.ApiPrefix)
	blog.AddApiRoutes(router, constants.ApiPrefix)
	orders.AddApiRoutes(router, constants.ApiPrefix)

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
	router.HTMLRender = helpers.CreateHTMLRenderHelper(templates, constants)

	router.Run(constants.Domain + ":" + constants.Port)
}
