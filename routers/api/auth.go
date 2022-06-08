package api

import (
	"net/http"
	"share/share-api/models/app"
	"share/share-api/mw"
	"share/share-api/services"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type auth struct {
	UserName string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func Auth(ctx *gin.Context) {
	appG := app.Gin{Ctx: ctx}
	valid := validation.Validation{}

	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	a := auth{username, password}
	ok, _ := valid.Valid(&a)

	if !ok {
		code := http.StatusBadRequest
		appG.Response(code, nil)
		return
	}

	userAuth := services.Auth{Username: username, Password: password}

	userDoesExist, err := userAuth.Validate()

	if err != nil {
		code := http.StatusInternalServerError
		appG.Response(code, nil)
		return
	}

	if !userDoesExist {
		code := http.StatusUnauthorized
		appG.Response(code, nil)
		return
	}

	token, err := mw.GenerateToken(username, password)
	if err != nil {
		code := http.StatusInternalServerError
		appG.Response(code, nil)
		return
	}

	appG.Response(http.StatusOK, token)

}
