package server

import (
	appconfig "github.com/Ganasa18/document-be/config"
	authController "github.com/Ganasa18/document-be/internal/auth/controller"
	crudController "github.com/Ganasa18/document-be/internal/crud/controller"
	"github.com/Ganasa18/document-be/pkg/exception"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type HttpServe struct {
	router         *gin.Engine
	authController authController.AuthController
	roleController crudController.RoleController
	menuController crudController.MenuController
}

func RunHttpServer(appConf *appconfig.Config, authController authController.AuthController, roleController crudController.RoleController, menuController crudController.MenuController) error {
	var hs HttpServe

	hs.router = gin.New()
	gin.SetMode(appConf.GinMode)

	// Global Exception Error Handler
	hs.router.Use(exception.ExceptionRecoveryMiddleware)

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"*"}
	hs.router.Use(cors.New(corsConfig))

	hs.router.SetTrustedProxies([]string{appConf.AppUrl})

	hs.authController = authController
	hs.roleController = roleController
	hs.menuController = menuController
	hs.setupRouter()

	return hs.router.Run(appConf.AppUrl + ":" + appConf.AppPort)
}
