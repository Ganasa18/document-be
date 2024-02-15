package controller

import "github.com/gin-gonic/gin"

type UserAccessController interface {
	GetAllUserAccess(ctx *gin.Context)
	GetUserAccessById(ctx *gin.Context)
	CreateUserAccess(ctx *gin.Context)
	UpdateUserAccess(ctx *gin.Context)
	DeleteUserAccess(ctx *gin.Context)
}
