package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	Ctx *gin.Context
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (g *Gin) Response(httpCode int, data interface{}) {
	g.Ctx.JSON(httpCode, Response{
		Code:    httpCode,
		Message: http.StatusText(httpCode),
		Data:    data,
	})
}
