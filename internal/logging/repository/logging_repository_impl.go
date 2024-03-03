package repository

import (
	"errors"

	"github.com/Ganasa18/document-be/internal/logging/model/domain"
	"github.com/Ganasa18/document-be/pkg/loghelper"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoggingRepositoryImpl struct {
	DB *gorm.DB
}

func NewLoggingRepository(db *gorm.DB) LoggingRepository {
	return &LoggingRepositoryImpl{
		DB: db,
	}
}

func (repository *LoggingRepositoryImpl) AddLoginLogging(ctx *gin.Context, login domain.LoginLogModel) error {
	err := repository.DB.Create(&login).Error
	if err != nil {
		loghelper.LogErrorRepository(ctx, "AddLoginLogging", "LoginLogModel", "create history login", err)
		return errors.New("failed to create history login")
	}

	return nil
}
