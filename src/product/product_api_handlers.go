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
}

type ProductApiHadler struct {
}

var apiHelper = helpers.ApiHelper{}

var usrServ users.IUserService = users.UserService{}
var productServ = ProductService{}

func (ProductApiHadler) findAll(c *gin.Context) {
	defer apiHelper.HandleError(c)

	allProducts := productServ.findAll()
	c.JSON(http.StatusOK, allProducts)
}

func (ProductApiHadler) add(c *gin.Context) {
	defer apiHelper.HandleError(c)

	token := apiHelper.GetRequiredToken(c)
	if !usrServ.CheckRol([]string{"products", "admin"}, token) {
		panic(errorss.UnAuthUser)
	}

	productFromReq := getProductFromRequest(c)

	saved := productServ.save(productFromReq)

	showProduct(c, saved)
}

func (ProductApiHadler) delete(c *gin.Context) {
	defer apiHelper.HandleError(c)

	token := apiHelper.GetRequiredToken(c)
	if !usrServ.CheckRol([]string{"products", "admin"}, token) {
		panic(errorss.UnAuthUser)
	}

	id := apiHelper.GetIntParam(c, "id")

	deleted := productServ.delete(id)

	showProduct(c, deleted)
}

func (ProductApiHadler) update(c *gin.Context) {
	defer apiHelper.HandleError(c)

	token := apiHelper.GetRequiredToken(c)
	if !usrServ.CheckRol([]string{"products", "admin"}, token) {
		panic(errorss.UnAuthUser)
	}

	productFromReq := getProductFromRequest(c)

	updated := productServ.update(productFromReq)

	showProduct(c, updated)
}

func (ProductApiHadler) saveImage(c *gin.Context) {
	defer apiHelper.HandleError(c)

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
	defer apiHelper.HandleError(c)

	id := apiHelper.GetIntParam(c, "id")
	finded := productServ.findImageByProductId(id)

	apiHelper.ShowImageInBase64(c, finded.Base64)
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

// likes

func (ProductApiHadler) addLike(c *gin.Context) {
	defer apiHelper.HandleError(c)

	token := apiHelper.GetOptionalToken(c)

	id := apiHelper.GetIntParam(c, "id")

	likesCount := productServ.addLike(id, int(token.IdUser))

	c.JSON(http.StatusOK, gin.H{"likes": likesCount})
}

func (ProductApiHadler) removeLike(c *gin.Context) {
	defer apiHelper.HandleError(c)

	token := apiHelper.GetOptionalToken(c)

	id := apiHelper.GetIntParam(c, "id")

	likesCount := productServ.removeLike(id, int(token.IdUser))

	c.JSON(http.StatusOK, gin.H{"likes": likesCount})
}
