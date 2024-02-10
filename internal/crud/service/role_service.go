package service

import (
	"github.com/Ganasa18/document-be/internal/crud/model/web"
	"github.com/Ganasa18/document-be/pkg/helper"
	"github.com/gin-gonic/gin"
)

type RoleService interface {
	GetRoles(ctx *gin.Context, pagination *helper.PaginationInput) ([]web.RoleMasterResponse, int64, error)
	GetRoleById(ctx *gin.Context) web.RoleMasterResponse
	CreateRole(ctx *gin.Context, request web.RoleMasterRequest) (web.RoleMasterResponse, error)
	UpdateRole(ctx *gin.Context, request web.RoleMasterRequest) (web.RoleMasterResponse, error)
	DeleteRole(ctx *gin.Context) error
}
