package routers

import (
	"share/share-api/common/config"
	"share/share-api/mw"
	"share/share-api/routers/api"

	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.POST("/auth", api.Auth)
	router.POST("/auth/register", api.CreateUser)

	apiv1 := router.Group(config.ApplicationConfig.EndpointPrefix)
	apiv1.Use(mw.JWT())
	apiv1.POST("/share", api.ShareNewPost)
	apiv1.GET("/posts", api.GetPosts)

	return router
}
