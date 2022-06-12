package api

import (
	"net/http"
	"share/share-api/models/app"
	"time"

	"github.com/gin-gonic/gin"
)

type ShareRequest struct {
	Id   int    `json:"id"`
	Data string `json:"data"`
	Time int64  `json:"time"`
}

func Share(ctx *gin.Context) {
	appG := app.Gin{Ctx: ctx}

	appG.Response(http.StatusOK, ShareRequest{Data: "data", Time: time.Now().Unix()})

}
