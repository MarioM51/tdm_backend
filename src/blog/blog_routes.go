package blog

import (
	"users_api/src/helpers"

	"github.com/gin-gonic/gin"
)

func AddApiRoutes(r *gin.Engine, prefix string) {

	blogHandler := BlogHandler{}

	r.GET(prefix+"/blogs", blogHandler.findAll)
	r.GET(prefix+"/blogs/:id", blogHandler.findById)
	r.GET(prefix+"/blogs/:id/image", blogHandler.showThumbnail)
	r.POST(prefix+"/blogs", blogHandler.save)
	r.PUT(prefix+"/blogs", blogHandler.update)
	r.DELETE(prefix+"/blogs/:id", blogHandler.deleteById)

}

func AddSsrRoutes(r *gin.Engine, tModels *[]helpers.TemplateModel) {
	blogSSRHandler := BlogSSRHandler{}

	r.GET("/blogs", blogSSRHandler.findAll)
	r.GET("/blogs/:name", blogSSRHandler.findById)

	templatesProducts := []helpers.TemplateModel{
		{
			Name:       "blog-list",
			PagePath:   "templates/blogs/blog-list.gohtml",
			LayoutPath: helpers.LAYOUT_A,
		},
		{
			Name:       "blog-details",
			PagePath:   "templates/blogs/blog-details.gohtml",
			LayoutPath: helpers.LAYOUT_A,
		},
	}
	*tModels = append(*tModels, templatesProducts...)
}
