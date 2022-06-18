package app

import (
	"net/http"
	"share/share-api/models"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	Ctx *gin.Context
}

func (g *Gin) Response(httpCode int, data interface{}) {

	g.Ctx.JSON(httpCode, models.Response{
		Code:    httpCode,
		Message: http.StatusText(httpCode),
		Data:    data,
	})

}
