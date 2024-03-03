package repository

import (
	"github.com/Ganasa18/document-be/internal/crud/model/domain"
	"github.com/Ganasa18/document-be/pkg/helper"
	"github.com/gin-gonic/gin"
)

type RoleRepository interface {
	GetRoles(ctx *gin.Context, pagination *helper.PaginationInput) ([]domain.RoleMasterModel, int64, error)
	GetRoleById(ctx *gin.Context, id int) (domain.RoleMasterModel, error)
	CreateRole(ctx *gin.Context, role domain.RoleMasterModel) (domain.RoleMasterModel, error)
	UpdateRole(ctx *gin.Context, role domain.RoleMasterModel, id int) (domain.RoleMasterModel, error)
	DeleteRole(ctx *gin.Context, id int) error
}
