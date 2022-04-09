package blog

import (
	"github.com/gin-gonic/gin"
)

func AddRoutes(r *gin.Engine) {

	blogHandler := BlogHandler{}

	r.GET("/blogs", blogHandler.findAll)
	r.GET("/blogs/:id", blogHandler.findById)
	r.POST("/blogs", blogHandler.save)
	r.PUT("/blogs", blogHandler.update)
	r.DELETE("/blogs/:id", blogHandler.deleteById)

}
