package routers

import (
	"github.com/gin-gonic/gin"
	"share/share-api/common/config"
	"share/share-api/mw"
	"share/share-api/routers/api"
)

func CreateRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.POST("/auth", api.Auth)
	router.POST("/auth/createUser", api.CreateUser)

	apiv1 := router.Group(config.ApplicationConfig.EndpointPrefix)
	apiv1.Use(mw.JWT())
	return router
}
