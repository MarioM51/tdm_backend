package helpers

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"strconv"
	"users_api/src/crypto"
	"users_api/src/errorss"

	"github.com/gin-gonic/gin"
)

type ApiHelper struct {
}

func (ApiHelper) GetIntParam(c *gin.Context, paramName string) int {
	idParam := c.Param(paramName)
	id, err := strconv.Atoi(idParam)
	if err != nil {
		panic(err)
	}
	return id
}

func (ApiHelper) HandleError(c *gin.Context) {
	if err := recover(); err != nil {

		if errResp, ok := err.(errorss.ErrorResponseModel); ok {
			c.JSON(errResp.HttpStatus, err)
		} else {
			fmt.Print(err)
			c.JSON(500, errorss.ErrorResponseModel{HttpStatus: 500, Cause: "Error, intente mas tarde"})
		}

	}
}

func (ApiHelper) GetRequiredToken(c *gin.Context) *crypto.TokenModel {
	headerToken := c.Request.Header["Token"]
	if len(headerToken) == 0 {
		panic(errorss.ErrorResponseModel{HttpStatus: 401, Cause: "Token header required"})
	}
	plainToken := headerToken[0]
	token := crypto.ParseRequiredToken(plainToken)
	return token
}

func (ApiHelper) GetOptionalToken(c *gin.Context) *crypto.TokenModel {
	headerToken := c.Request.Header["Token"]
	defaulToken := &crypto.TokenModel{IdUser: 0}
	if len(headerToken) == 0 {
		return defaulToken
	}
	plainToken := headerToken[0]
	token := crypto.ParseOptionalToken(plainToken)
	if token == nil {
		return defaulToken
	}

	return token
}

func (ApiHelper) FileMultiPartToBase64(file *multipart.FileHeader) *string {
	f, err := file.Open()
	if err != nil {
		panic("file to base64 open fail")
	}

	// Read entire file into byte slice.
	reader := bufio.NewReader(f)
	content, err2 := ioutil.ReadAll(reader)
	if err2 != nil {
		panic("file to base64 transform fail")
	}

	// Encode as base64.
	base64Result := base64.StdEncoding.EncodeToString(content)

	return &base64Result
}

func (ApiHelper) ShowImageInBase64(c *gin.Context, imageBase64 string) {
	imagebytes, err := base64.StdEncoding.DecodeString(imageBase64)
	if err != nil {
		panic("Error decoding image")
	}

	c.Writer.Header().Set("Content-Type", "image/jpeg")
	c.Writer.Header().Set("Content-Length", strconv.Itoa(len(imagebytes)))
	c.Writer.Header().Set("Cache-Control", "max-age=604800")
	_, err2 := c.Writer.Write(imagebytes)
	if err2 != nil {
		panic("Displaying image error")
	}
}
