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
	c.HTML(http.StatusOK, "product-details", gin.H{"Products": "hi"})
}
