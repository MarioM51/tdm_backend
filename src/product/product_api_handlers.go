package product

import (
	"net/http"
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

	token := apiHelper.GetToken(c)
	if !usrServ.CheckRol([]string{"products", "admin"}, token) {
		panic(errorss.UnAuthUser)
	}

	productFromReq := getProductFromRequest(c)

	saved := productServ.save(productFromReq)

	showProduct(c, saved)
}

func (ProductApiHadler) delete(c *gin.Context) {
	defer apiHelper.HandleError(c)

	token := apiHelper.GetToken(c)
	if !usrServ.CheckRol([]string{"products", "admin"}, token) {
		panic(errorss.UnAuthUser)
	}

	id := apiHelper.GetIntParam(c, "id")

	deleted := productServ.delete(id)

	showProduct(c, deleted)
}

func (ProductApiHadler) update(c *gin.Context) {
	defer apiHelper.HandleError(c)

	token := apiHelper.GetToken(c)
	if !usrServ.CheckRol([]string{"products", "admin"}, token) {
		panic(errorss.UnAuthUser)
	}

	productFromReq := getProductFromRequest(c)

	updated := productServ.update(productFromReq)

	showProduct(c, updated)
}

//=================

func getProductFromRequest(c *gin.Context) (p *ProductModel) {
	if err := c.BindJSON(&p); err != nil {
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
