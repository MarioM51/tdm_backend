package product

import (
	"users_api/src/helpers"

	"github.com/gin-gonic/gin"
)

var productApi IProductApiHadler = ProductApiHadler{}
var productSsrhandler IProductSsrHadler = ProductSsrHadler{}

func AddApiRoutes(r *gin.Engine, prefix string) {
	r.GET(prefix+"/products", productApi.findAll)
	r.POST(prefix+"/products", productApi.add)
	r.PUT(prefix+"/products", productApi.update)
	r.DELETE(prefix+"/products/:id", productApi.delete)

	r.POST(prefix+"/products/:id/images", productApi.saveImage)
	r.GET(prefix+"/products/image/:id", productApi.showImage)
}

func AddSsrRoutes(r *gin.Engine, tModels *[]helpers.TemplateModel) {
	r.GET("/products", productSsrhandler.findAll)
	r.GET("/products-details", productSsrhandler.findDetails)

	templatesProducts := []helpers.TemplateModel{
		{
			Name:       "product-list",
			PagePath:   "templates/products/product-list.gohtml",
			LayoutPath: helpers.LAYOUT_A,
		},
		{
			Name:       "product-details",
			PagePath:   "templates/products/product-details.gohtml",
			LayoutPath: helpers.LAYOUT_A,
		},
	}
	*tModels = append(*tModels, templatesProducts...)
}
