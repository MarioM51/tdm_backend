package blog

import (
	"net/http"
	"strconv"
	"users_api/src/errorss"

	"github.com/gin-gonic/gin"
)

type IBlogSSRHandler interface {
	findAll(c *gin.Context)
	findById(c *gin.Context)
}

type BlogSSRHandler struct{}

func (BlogSSRHandler) findAll(c *gin.Context) {
	defer apiHelper.HandleError(c)

	var allBlogs = blogS.findAll()

	allBlogsJSONLD := BlogModelToArrayJSONLD(*allBlogs)

	c.HTML(http.StatusOK, "blog-list", gin.H{"BLOGS_JSONLD": allBlogsJSONLD.Val})

}

func (BlogSSRHandler) findById(c *gin.Context) {
	defer apiHelper.HandleError(c)

	name := c.Param("name")
	rawId := name[len(name)-1:]
	id, err := strconv.Atoi(rawId)
	if err != nil {
		panic(errorss.ErrorResponseModel{HttpStatus: 400, Cause: "Url bad formated"})
	}

	var finded = blogS.findById(id)
	finded.Thumbnail = ""

	allBlogsJSONLD := BlogModelToJSONLD(*finded)

	c.HTML(http.StatusOK, "blog-details", gin.H{
		"BLOG_JSONLD": allBlogsJSONLD.Val,
	})

}
