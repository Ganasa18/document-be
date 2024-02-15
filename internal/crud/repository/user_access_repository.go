package repository

import (
	"github.com/Ganasa18/document-be/internal/crud/model/domain"
	"github.com/Ganasa18/document-be/pkg/helper"
	"github.com/gin-gonic/gin"
)

type UserAccessRepository interface {
	GetAllUserAccess(ctx *gin.Context, pagination *helper.PaginationInput)
	GetUserAccessById(ctx *gin.Context)
	CreateUserAccess(ctx *gin.Context, request domain.UserAccessMenuModel) (domain.UserAccessMenuModel, error)
	UpdateUserAccess(ctx *gin.Context)
	DeleteUserAccess(ctx *gin.Context)
}
