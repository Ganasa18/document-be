package server

import (
	appconfig "github.com/Ganasa18/document-be/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type HttpServe struct {
	router *gin.Engine
}

func RunHttpServer(appConf *appconfig.Config) error {
	var hs HttpServe

	hs.router = gin.New()
	gin.SetMode(appConf.GinMode)

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"*"}
	hs.router.Use(cors.New(corsConfig))

	hs.router.SetTrustedProxies([]string{appConf.AppUrl})

	hs.setupRouter()

	return hs.router.Run(appConf.AppUrl + ":" + appConf.AppPort)
}
