package blog

import (
	"net/http"
	"strings"
	"users_api/src/errorss"
	"users_api/src/helpers"
	"users_api/src/users"

	"github.com/gin-gonic/gin"
)

type IBlogHandler interface {
	findAll(c *gin.Context)
	findById(c *gin.Context)
	showThumbnail(c *gin.Context)
	save(c *gin.Context)
	update(c *gin.Context)
	deleteById(c *gin.Context)

	addLike(c *gin.Context)
	removeLike(c *gin.Context)
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
	finded.Thumbnail = ""

	showBlog(c, finded)
}

func (BlogHandler) showThumbnail(c *gin.Context) {
	defer apiHelper.HandleError(c)

	idBlog := apiHelper.GetIntParam(c, "id")
	finded := blogS.findById(idBlog)

	cut := strings.Index(finded.Thumbnail, ",") + 1
	base := finded.Thumbnail[cut:len(finded.Thumbnail)]

	apiHelper.ShowImageInBase64(c, base)
}

func (BlogHandler) save(c *gin.Context) {
	defer apiHelper.HandleError(c)

	token := apiHelper.GetRequiredToken(c)
	if !usrServ.CheckRol([]string{"blogs", "admin"}, token) {
		panic(errorss.UnAuthUser)
	}

	newBlog := getBlogFromRequest(c)
	savedBlog := blogS.save(newBlog)
	showBlog(c, savedBlog)
}

func (BlogHandler) update(c *gin.Context) {
	defer apiHelper.HandleError(c)

	token := apiHelper.GetRequiredToken(c)
	if !usrServ.CheckRol([]string{"blogs", "admin"}, token) {
		panic(errorss.UnAuthUser)
	}

	newInfoBlog := getBlogFromRequest(c)
	updatedBlog := blogS.update(newInfoBlog)
	showBlog(c, updatedBlog)
}

func (BlogHandler) deleteById(c *gin.Context) {
	defer apiHelper.HandleError(c)

	token := apiHelper.GetRequiredToken(c)
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
		panic(errorss.ErrorResponseModel{HttpStatus: 400, Cause: "bad format of blog json"})
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

func (BlogHandler) addLike(c *gin.Context) {
	defer apiHelper.HandleError(c)

	token := apiHelper.GetOptionalToken(c)

	idBlog := apiHelper.GetIntParam(c, "id")

	likesCount := blogS.addLike(idBlog, int(token.IdUser))

	c.JSON(http.StatusNotFound, gin.H{"likes": likesCount})
}

func (BlogHandler) removeLike(c *gin.Context) {
	defer apiHelper.HandleError(c)

	token := apiHelper.GetOptionalToken(c)

	idBlog := apiHelper.GetIntParam(c, "id")

	likesCount := blogS.removeLike(idBlog, int(token.IdUser))

	c.JSON(http.StatusNotFound, gin.H{"likes": likesCount})
}
