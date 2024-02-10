package service

import (
	"github.com/Ganasa18/document-be/internal/crud/model/web"
	"github.com/Ganasa18/document-be/pkg/helper"
	"github.com/gin-gonic/gin"
)

type RoleService interface {
	GetRoles(ctx *gin.Context, pagination *helper.PaginationInput) ([]web.RoleMasterResponse, int64, error)
}
