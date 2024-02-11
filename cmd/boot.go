package cmd

import (
	appconfig "github.com/Ganasa18/document-be/config"
	"github.com/Ganasa18/document-be/http/server"
	"github.com/Ganasa18/document-be/pkg/utils"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func initHTTP() error {

	appConf := appconfig.InitAppConfig()

	var gConfig *gorm.Config = &gorm.Config{}
	db, err := appconfig.NewDatabase(appConf, gConfig)

	validate := validator.New()

	// auth definition
	authCtrl, roleCtrl := server.InitializeModel(db, validate)

	if err != nil {
		utils.IsErrorDoPanic(err)
	}

	// run server
	err = server.RunHttpServer(appConf, authCtrl, roleCtrl)
	if err != nil {
		utils.IsErrorDoPanic(err)
	}

	return nil
}
