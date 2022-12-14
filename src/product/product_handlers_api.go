package product

import (
	"encoding/json"
	"net/http"
	"strings"
	"users_api/src/errorss"
	"users_api/src/helpers"
	"users_api/src/users"

	"github.com/gin-gonic/gin"
)

type IProductApiHadler interface {
	findAll(c *gin.Context)
	add(c *gin.Context)
	delete(c *gin.Context)
	update(c *gin.Context)
	saveImage(c *gin.Context)
	showImage(c *gin.Context)
	addLike(c *gin.Context)
	removeLike(c *gin.Context)
	deleteImage(c *gin.Context)

	deleteComment(c *gin.Context)
	addComment(c *gin.Context)
	addResponse(c *gin.Context)
	findAllComments(c *gin.Context)
}

type ProductApiHadler struct {
}

var apiHelper = helpers.ApiHelper{}

var usrServ users.IUserService = users.UserService{}
var productServ = ProductService{}

func (ProductApiHadler) findAll(c *gin.Context) {
	defer apiHelper.HandleApiError(c)

	allProducts := productServ.findAll()
	c.JSON(http.StatusOK, allProducts)
}

func (ProductApiHadler) add(c *gin.Context) {
	defer apiHelper.HandleApiError(c)

	token := apiHelper.GetRequiredToken(c)
	if !usrServ.CheckRol([]string{"products", "admin"}, token) {
		panic(errorss.UnAuthUser)
	}

	productFromReq := getProductFromRequest(c)

	saved := productServ.save(productFromReq)

	showProduct(c, saved)
}

func (ProductApiHadler) delete(c *gin.Context) {
	defer apiHelper.HandleApiError(c)

	token := apiHelper.GetRequiredToken(c)
	if !usrServ.CheckRol([]string{"products", "admin"}, token) {
		panic(errorss.UnAuthUser)
	}

	id := apiHelper.GetIntParam(c, "id")

	deleted := productServ.delete(id)

	showProduct(c, deleted)
}

func (ProductApiHadler) update(c *gin.Context) {
	defer apiHelper.HandleApiError(c)

	token := apiHelper.GetRequiredToken(c)
	if !usrServ.CheckRol([]string{"products", "admin"}, token) {
		panic(errorss.UnAuthUser)
	}

	productFromReq := getProductFromRequest(c)

	updated := productServ.update(productFromReq)

	showProduct(c, updated)
}

func (ProductApiHadler) saveImage(c *gin.Context) {
	defer apiHelper.HandleApiError(c)

	id := apiHelper.GetIntParam(c, "id")

	image, err := c.FormFile("file")
	if err != nil {
		panic(errorss.ErrorResponseModel{HttpStatus: 500, Cause: "capturing file error"})
	}

	contentType := image.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") {
		panic(errorss.ErrorResponseModel{HttpStatus: 400, Cause: "file type invalid, image required"})
	}

	base64 := apiHelper.FileMultiPartToBase64(image)

	newImage := ProductImage{
		MimeType: contentType,
		Base64:   *base64,
	}

	saved := productServ.saveImage(id, &newImage)

	c.JSON(http.StatusOK, saved)

}

func (ProductApiHadler) showImage(c *gin.Context) {
	defer apiHelper.HandleApiError(c)

	id := apiHelper.GetIntParam(c, "id")
	finded := productServ.findImageIdImage(id)

	apiHelper.ShowImageInBase64(c, finded.Base64)
}

func (ProductApiHadler) deleteImage(c *gin.Context) {
	defer apiHelper.HandleApiError(c)

	token := apiHelper.GetRequiredToken(c)
	if !usrServ.CheckRol([]string{"products", "admin"}, token) {
		panic(errorss.UnAuthUser)
	}

	id := apiHelper.GetIntParam(c, "id")
	imageDeleted := productServ.deleteImageIdImage(id)

	c.JSON(http.StatusOK, &imageDeleted)

}

//=================

func getProductFromRequest(c *gin.Context) (p *ProductModel) {
	if err := c.ShouldBindJSON(&p); err != nil {
		if specificError, ok := err.(*json.UnmarshalTypeError); ok {
			if specificError.Field == "image" {
				return p
			}
		}

		panic(errorss.ErrorResponseModel{HttpStatus: 400, Cause: "Product json bad format"})
	}
	return p
}

func showProduct(c *gin.Context, p *ProductModel) {
	if p != nil {
		c.JSON(http.StatusOK, &p)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
	}
}

func (pah ProductApiHadler) bindComment(c *gin.Context) Comment {
	commentReceived := Comment{}
	if err := c.BindJSON(&commentReceived); err != nil {
		panic(&errorss.ErrorResponseModel{HttpStatus: 400, Cause: "bad format of product comment"})
	}

	return commentReceived
}

// likes

func (ProductApiHadler) addLike(c *gin.Context) {
	defer apiHelper.HandleApiError(c)

	token := apiHelper.GetOptionalToken(c)

	id := apiHelper.GetIntParam(c, "id")

	likesCount := productServ.addLike(id, int(token.IdUser))

	c.JSON(http.StatusOK, gin.H{"likes": likesCount})
}

func (ProductApiHadler) removeLike(c *gin.Context) {
	defer apiHelper.HandleApiError(c)

	token := apiHelper.GetOptionalToken(c)

	id := apiHelper.GetIntParam(c, "id")

	likesCount := productServ.removeLike(id, int(token.IdUser))

	c.JSON(http.StatusOK, gin.H{"likes": likesCount})
}

//Comments ============================

func (pah ProductApiHadler) addComment(c *gin.Context) {
	defer apiHelper.HandleApiError(c)

	token := apiHelper.GetRequiredToken(c)
	idTarget := apiHelper.GetIntParam(c, "id")

	commentReceived := pah.bindComment(c)
	commentReceived.IdTarget = idTarget
	commentReceived.IdUser = int(token.IdUser)
	commentReceived.cleanAndValidateNewComment(false)

	commentAdded := productServ.addComment(commentReceived)

	c.JSON(http.StatusOK, commentAdded)
}

func (ProductApiHadler) deleteComment(c *gin.Context) {
	defer apiHelper.HandleApiError(c)

	token := apiHelper.GetRequiredToken(c)
	idTarget := apiHelper.GetIntParam(c, "id")
	idComment := apiHelper.GetIntParam(c, "idComment")

	commentToDelete := Comment{
		Id:       idComment,
		IdUser:   int(token.IdUser),
		IdTarget: idTarget,
	}

	if usrServ.CheckRol([]string{"products", "admin"}, token) {
		commentToDelete.IdUser = 666777
	}

	commentDeleted := productServ.deleteComment(commentToDelete)
	c.JSON(http.StatusOK, commentDeleted)
}

func (pah ProductApiHadler) addResponse(c *gin.Context) {
	defer apiHelper.HandleApiError(c)

	token := apiHelper.GetRequiredToken(c)
	idTarget := apiHelper.GetIntParam(c, "id")
	idComment := apiHelper.GetIntParam(c, "idComment")

	responseReceived := pah.bindComment(c)
	responseReceived.IdUser = int(token.IdUser)
	responseReceived.IdTarget = idTarget
	responseReceived.ResponseTo = idComment
	responseReceived.cleanAndValidateNewComment(true)

	productServ.addResponse(&responseReceived)

	c.JSON(http.StatusOK, responseReceived)
}

func (pah ProductApiHadler) findAllComments(c *gin.Context) {
	defer apiHelper.HandleApiError(c)

	token := apiHelper.GetRequiredToken(c)
	if !usrServ.CheckRol([]string{"products", "admin"}, token) {
		panic(errorss.UnAuthUser)
	}

	allComments := productServ.findAllComments()

	c.JSON(http.StatusOK, allComments)
}
