package home

import (
	"users_api/src/helpers"

	"github.com/gin-gonic/gin"
)

var apiHelper = helpers.ApiHelper{}

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
