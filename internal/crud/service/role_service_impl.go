package service

import (
	"github.com/Ganasa18/document-be/internal/crud/model/web"
	"github.com/Ganasa18/document-be/internal/crud/repository"
	"github.com/Ganasa18/document-be/pkg/helper"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type RoleServiceImpl struct {
	RoleRepository repository.RoleRepository
	Validate       *validator.Validate
}

func NewRoleService(roleRepository repository.RoleRepository, validate *validator.Validate) RoleService {
	return &RoleServiceImpl{
		RoleRepository: roleRepository,
		Validate:       validate,
	}
}

func (service *RoleServiceImpl) GetRoles(ctx *gin.Context, pagination *helper.PaginationInput) (roles []web.RoleMasterResponse, totalRows int64, err error) {
	roleResponse, totalRow, err := service.RoleRepository.GetRoles(ctx, pagination)
	return web.ToRoleMasterResponses(roleResponse, totalRow, err)
}
