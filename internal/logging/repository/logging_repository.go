package repository

import (
	"github.com/Ganasa18/document-be/internal/logging/model/domain"
	"github.com/gin-gonic/gin"
)

type LoggingRepository interface {
	AddLoginLogging(ctx *gin.Context, login domain.LoginLogModel) error
}
