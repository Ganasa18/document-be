package cmd

import (
	appconfig "github.com/Ganasa18/document-be/config"
	"github.com/Ganasa18/document-be/http/server"
	"github.com/Ganasa18/document-be/internal/utils"
	"gorm.io/gorm"
)

func initHTTP() error {
	appConf := appconfig.InitAppConfig()

	var gConfig *gorm.Config = &gorm.Config{}
	_, err := appconfig.NewDatabase(appConf, gConfig)

	// auth repository definition

	if err != nil {
		utils.IsErrorDoPanic(err)
	}

	// run server
	err = server.RunHttpServer(appConf)
	if err != nil {
		utils.IsErrorDoPanic(err)
	}

	return nil
}
