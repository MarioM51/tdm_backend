package blog

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IBlogSSRHandler interface {
	findAll(c *gin.Context)
	findById(c *gin.Context)
}

type BlogSSRHandler struct{}

func (BlogSSRHandler) findAll(c *gin.Context) {
	defer apiHelper.HandleApiError(c)

	var allBlogs = blogS.findAll()

	allBlogsJSONLD := BlogModelToArrayJSONLD(*allBlogs)

	c.HTML(http.StatusOK, "blog-list", gin.H{"BLOGS_JSONLD": allBlogsJSONLD.Val})

}

func (BlogSSRHandler) findById(c *gin.Context) {
	defer apiHelper.HandleApiError(c)
	id := apiHelper.GetLastNumberInParam(c, "name")
	var finded = blogS.findById(id)
	finded.Thumbnail = ""

	allBlogsJSONLD := BlogModelToJSONLDWrapped(*finded)

	c.HTML(http.StatusOK, "blog-details", gin.H{
		"BLOG_JSONLD": allBlogsJSONLD.Val,
	})

}
