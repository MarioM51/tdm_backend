package blog

import (
	"github.com/gin-gonic/gin"
)

func AddApiRoutes(r *gin.Engine, prefix string) {

	blogHandler := BlogHandler{}

	r.GET(prefix+"/api/blogs", blogHandler.findAll)
	r.GET(prefix+"/blogs/:id", blogHandler.findById)
	r.POST(prefix+"/blogs", blogHandler.save)
	r.PUT(prefix+"/blogs", blogHandler.update)
	r.DELETE(prefix+"/blogs/:id", blogHandler.deleteById)

}
