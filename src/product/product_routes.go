package product

import "github.com/gin-gonic/gin"

var productApi IProductApiHadler = ProductApiHadler{}

func AddRoutes(r *gin.Engine) {
	r.GET("/products", productApi.findAll)
	r.POST("/products", productApi.add)
	r.PUT("/products", productApi.update)
	r.DELETE("/products/:id", productApi.delete)
}
