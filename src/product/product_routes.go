package product

import (
	"users_api/src/helpers"

	"github.com/gin-gonic/gin"
)

var productApi IProductApiHadler = ProductApiHadler{}
var productSsrhandler IProductSsrHadler = ProductSsrHadler{}
var dbHelper *helpers.DBHelper = nil
var constants helpers.Constants

func LinkDependencies(db *helpers.DBHelper, consts helpers.Constants) {
	dbHelper = db
	constants = consts
}

func AddApiRoutes(r *gin.Engine, prefix string) {
	r.GET(prefix+"/products", productApi.findAll)
	r.POST(prefix+"/products", productApi.add)
	r.PUT(prefix+"/products", productApi.update)
	r.DELETE(prefix+"/products/:id", productApi.delete)

	// image
	r.POST(prefix+"/products/:id/images", productApi.saveImage)
	r.GET(prefix+"/products/image/:id", productApi.showImage)
	r.DELETE(prefix+"/products/image/:id", productApi.deleteImage)

	//likes
	r.POST(prefix+"/products/:id/like", productApi.addLike)
	r.DELETE(prefix+"/products/:id/like", productApi.removeLike)

	//comments
	r.POST(prefix+"/products/:id/comment", productApi.addComment)
	r.DELETE(prefix+"/products/:id/comment/:idComment", productApi.deleteComment)
}

func AddSsrRoutes(r *gin.Engine, tModels *[]helpers.TemplateModel) {
	r.GET("/products", productSsrhandler.findAll)
	r.GET("/products/:name", productSsrhandler.findDetails)

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
