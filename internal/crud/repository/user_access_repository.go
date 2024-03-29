package repository

import (
	"github.com/Ganasa18/document-be/internal/crud/model/domain"
	"github.com/Ganasa18/document-be/pkg/helper"
	"github.com/gin-gonic/gin"
)

type UserAccessRepository interface {
	GetAllUserAccess(ctx *gin.Context, pagination *helper.PaginationInput) ([]domain.UserAccessMenuModel, int64, error)
	GetUserAccessById(ctx *gin.Context, id int) (domain.UserAccessMenuModel, error)
	CreateUserAccess(ctx *gin.Context, request domain.UserAccessMenuModel) (domain.UserAccessMenuModel, error)
	UpdateUserAccess(ctx *gin.Context, request domain.UserAccessMenuModel, id int) (domain.UserAccessMenuModel, error)
	DeleteUserAccess(ctx *gin.Context, id int) error
}
