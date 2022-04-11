package blog

import (
	"net/http"
	"users_api/src/errorss"
	"users_api/src/helpers"
	"users_api/src/users"

	"github.com/gin-gonic/gin"
)

type IBlogHandler interface {
	findAll(c *gin.Context)
	findById(c *gin.Context)
	save(c *gin.Context)
	update(c *gin.Context)
	deleteById(c *gin.Context)
}

type BlogHandler struct {
}

var usrServ users.IUserService = users.UserService{}
var blogS IBlogService = BlogService{}
var apiHelper = helpers.ApiHelper{}

func (BlogHandler) findAll(c *gin.Context) {
	defer apiHelper.HandleError(c)
	var allBlogs = blogS.findAll()
	c.JSON(http.StatusOK, allBlogs)
}

func (BlogHandler) findById(c *gin.Context) {
	defer apiHelper.HandleError(c)
	idBlog := apiHelper.GetIntParam(c, "id")
	finded := blogS.findById(idBlog)

	showBlog(c, finded)
}

func (BlogHandler) save(c *gin.Context) {
	defer apiHelper.HandleError(c)

	token := apiHelper.GetToken(c)
	if !usrServ.CheckRol([]string{"blogs", "admin"}, token) {
		panic(errorss.UnAuthUser)
	}

	newBlog := getBlogFromRequest(c)
	savedBlog := blogS.save(newBlog)
	showBlog(c, savedBlog)
}

func (BlogHandler) update(c *gin.Context) {
	defer apiHelper.HandleError(c)

	token := apiHelper.GetToken(c)
	if !usrServ.CheckRol([]string{"blogs", "admin"}, token) {
		panic(errorss.UnAuthUser)
	}

	newInfoBlog := getBlogFromRequest(c)
	updatedBlog := blogS.update(newInfoBlog)
	showBlog(c, updatedBlog)
}

func (BlogHandler) deleteById(c *gin.Context) {
	defer apiHelper.HandleError(c)

	token := apiHelper.GetToken(c)
	if !usrServ.CheckRol([]string{"blogs", "admin"}, token) {
		panic(errorss.UnAuthUser)
	}

	idBlog := apiHelper.GetIntParam(c, "id")
	deletedBlog := blogS.deleteById(idBlog)
	showBlog(c, deletedBlog)
}

//============================

func getBlogFromRequest(c *gin.Context) (b *BlogModel) {
	if err := c.BindJSON(&b); err != nil {
		panic(errorss.ErrorResponseModel{HttpStatus: 400, Cause: "Product json bad format"})
	}
	return b
}

func showBlog(c *gin.Context, p *BlogModel) {
	if p != nil {
		c.JSON(http.StatusOK, &p)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "blog not found"})
	}
}