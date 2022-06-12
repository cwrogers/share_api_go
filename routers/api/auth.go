package api

import (
	"net/http"
	"regexp"
	"share/share-api/common/strings"
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
		appG.Response(code, err.Error())
		return
	}

	if !userDoesExist {
		code := http.StatusUnauthorized
		appG.Response(code, strings.UserNotFound)
		return
	}

	token, refreshToken, err := mw.GenerateToken(username, password)
	if err != nil {
		code := http.StatusInternalServerError
		appG.Response(code, err.Error())
		return
	}

	appG.Response(http.StatusOK, mw.AuthenticationResponse{Token: token, RefreshToken: refreshToken})

}

func CreateUser(ctx *gin.Context) {
	appG := app.Gin{Ctx: ctx}
	valid := validation.Validation{}

	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	a := auth{username, password}
	ok, _ := valid.Valid(&a)

	if !ok || !isValidEmail(username) {
		code := http.StatusBadRequest
		appG.Response(code, nil)
		return
	}

	userAuth := services.Auth{Username: username, Password: password}

	userCreated, err := userAuth.CreateUser()

	if err != nil {
		code := http.StatusInternalServerError
		appG.Response(code, err.Error())
		return
	}

	if !userCreated {
		code := http.StatusInternalServerError
		appG.Response(code, strings.GenericErrorMessage)
		return
	}

}

// validate email
func isValidEmail(email string) bool {
	// regex email
	regex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	r := regexp.MustCompile(regex)
	return r.MatchString(email)
}
