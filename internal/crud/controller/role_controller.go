package controller

import (
	"github.com/gin-gonic/gin"
)

type RoleController interface {
	GetRoles(ctx *gin.Context)
}
