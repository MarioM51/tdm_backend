package home

import (
	"net/http"
	"users_api/src/blog"
	"users_api/src/product"

	"github.com/gin-gonic/gin"
)

type HomeSSRHandler struct{}

func (HomeSSRHandler) home(c *gin.Context) {
	defer apiHelper.HandleApiError(c)

	allProducts := productServ.FindOnHomeScreen()
	allProductsJSONLD := product.ProductModelToArrayJSONLD(*allProducts)

	var allBlogs = blogS.FindOnHomeScreen()
	allBlogsJSONLD := blog.BlogModelToArrayJSONLD(*allBlogs)

	c.HTML(
		http.StatusOK,
		"home",
		gin.H{
			"PRODUCTS_JSONLD": allProductsJSONLD.Val,
			"BLOGS_JSONLD":    allBlogsJSONLD.Val,
			"LOCALBUS_JSONLD": localInfoLDJson,
		})

}
