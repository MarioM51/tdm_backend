package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeSSRHandler struct{}

func (HomeSSRHandler) home(c *gin.Context) {
	defer apiHelper.HandleApiError(c)
	const aa = "hello"
	c.HTML(http.StatusOK, "home", gin.H{"HI": aa})

}
