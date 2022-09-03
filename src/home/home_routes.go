package home

import (
	"users_api/src/blog"
	"users_api/src/helpers"
	"users_api/src/product"

	"github.com/gin-gonic/gin"
)

var apiHelper = helpers.ApiHelper{}
var productServ = product.ProductService{}
var blogS blog.IBlogService = blog.BlogService{}

func AddSsrRoutes(r *gin.Engine, tModels *[]helpers.TemplateModel) {
	homeSSRHandler := HomeSSRHandler{}

	r.GET("/", homeSSRHandler.home)

	templatesProducts := []helpers.TemplateModel{
		{
			Name:       "home",
			PagePath:   "templates/home.gohtml",
			LayoutPath: helpers.LAYOUT_A,
		},
	}
	*tModels = append(*tModels, templatesProducts...)
}
