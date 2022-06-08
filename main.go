package main

import (
	"log"
	"net/http"
	"share/share-api/common/config"
	"share/share-api/routers"

	"github.com/gin-gonic/gin"
)

func main() {

	if config.ApplicationConfig.IsInTesting {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := routers.CreateRouter()
	addr := ":" + config.ApplicationConfig.Port

	server := &http.Server{
		Addr:           addr,
		Handler:        router,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("Server is running on port %s", addr)

	server.ListenAndServe()

}
