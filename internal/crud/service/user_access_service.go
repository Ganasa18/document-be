package service

import (
	"github.com/Ganasa18/document-be/internal/crud/model/web"
	"github.com/Ganasa18/document-be/pkg/helper"
	"github.com/gin-gonic/gin"
)

type UserAccessService interface {
	GetAllUserAccess(ctx *gin.Context, pagination *helper.PaginationInput)
	GetUserAccessById(ctx *gin.Context)
	CreateUserAccess(ctx *gin.Context, userAccess web.UserAccessRequest) (web.UserAccessResponse, error)
	UpdateUserAccess(ctx *gin.Context)
	DeleteUserAccess(ctx *gin.Context)
}