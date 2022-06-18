package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"share/share-api/models"
	"share/share-api/models/app"
	"time"
)

func GetPosts(ctx *gin.Context) {
	appG := app.Gin{Ctx: ctx}
	appG.Response(http.StatusOK, nil)
}

func ShareNewPost(ctx *gin.Context) {
	appG := app.Gin{Ctx: ctx}
	appG.Response(http.StatusOK, models.ShareRequest{Data: "data", Time: time.Now().Unix()})
}
