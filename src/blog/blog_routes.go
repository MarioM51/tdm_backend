package blog

import (
	"users_api/src/helpers"

	"github.com/gin-gonic/gin"
)

var dbHelper *helpers.DBHelper = nil
var constants helpers.Constants

func LinkDependencies(db *helpers.DBHelper, consts helpers.Constants) {
	dbHelper = db
	constants = consts
}

func AddApiRoutes(r *gin.Engine, prefix string) {

	var blogHandler IBlogHandler = BlogHandler{}

	//CRUD
	r.GET(prefix+"/blogs", blogHandler.findAll)
	r.GET(prefix+"/blogs/:id", blogHandler.findById)
	r.POST(prefix+"/blogs", blogHandler.save)
	r.PUT(prefix+"/blogs", blogHandler.update)
	r.DELETE(prefix+"/blogs/:id", blogHandler.deleteById)

	//images
	r.GET(prefix+"/blogs/:id/image", blogHandler.showThumbnail)

	//likes
	r.POST(prefix+"/blogs/:id/like", blogHandler.addLike)
	r.DELETE(prefix+"/blogs/:id/like", blogHandler.removeLike)

	//comments
	r.GET(prefix+"/blogs/comments", blogHandler.findAllComments)
	r.POST(prefix+"/blogs/:id/comment", blogHandler.addComment)
	r.DELETE(prefix+"/blogs/:id/comment/:idComment", blogHandler.deleteComment)
	r.POST(prefix+"/blogs/:id/comment/:idComment", blogHandler.addCommentResponse)

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
