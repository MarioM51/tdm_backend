package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IProductSsrHadler interface {
	findAll(c *gin.Context)
	findDetails(c *gin.Context)
}

type ProductSsrHadler struct {
}

func (ProductSsrHadler) findAll(c *gin.Context) {
	allProducts := productServ.findAll()
	allProductsJSONLD := ProductModelToArrayJSONLD(*allProducts)

	c.HTML(http.StatusOK, "product-list", gin.H{"PRODUCTS_JSONLD": allProductsJSONLD.Val})

}

func (ProductSsrHadler) findDetails(c *gin.Context) {
	defer apiHelper.HandleSSRError(c)
	id := apiHelper.GetLastNumberInParam(c, "name")
	productDetails := productServ.findById(id)
	productDetailsLDJson := ProductModelToJSONLD(productDetails, false)
	data := gin.H{
		"PRODUCT": productDetailsLDJson,
	}
	c.HTML(http.StatusOK, "product-details", data)
}
