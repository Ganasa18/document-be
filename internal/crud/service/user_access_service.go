package service

import (
	"github.com/Ganasa18/document-be/internal/crud/model/web"
	"github.com/Ganasa18/document-be/pkg/helper"
	"github.com/gin-gonic/gin"
)

type UserAccessService interface {
	GetAllUserAccess(ctx *gin.Context, pagination *helper.PaginationInput) ([]web.UserAccessResponseJoinRoleAndMenu, int64, error)
	GetUserAccessById(ctx *gin.Context) web.UserAccessResponse
	CreateUserAccess(ctx *gin.Context, userAccess web.UserAccessRequest) (web.UserAccessResponse, error)
	UpdateUserAccess(ctx *gin.Context)
	DeleteUserAccess(ctx *gin.Context) error
}
