package controller

import (
	"github.com/gin-gonic/gin"
)

type RoleController interface {
	GetRoles(ctx *gin.Context)
	GetRoleById(ctx *gin.Context)
	CreateRole(ctx *gin.Context)
	UpdateRole(ctx *gin.Context)
	DeleteRole(ctx *gin.Context)
}
