package service

import (
	"strconv"

	"github.com/Ganasa18/document-be/internal/crud/model/domain"
	"github.com/Ganasa18/document-be/internal/crud/model/web"
	"github.com/Ganasa18/document-be/internal/crud/repository"
	"github.com/Ganasa18/document-be/pkg/exception"
	"github.com/Ganasa18/document-be/pkg/helper"
	"github.com/Ganasa18/document-be/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RoleServiceImpl struct {
	RoleRepository repository.RoleRepository
	Validate       *validator.Validate
}

// UpdateRole implements RoleService.

func NewRoleService(roleRepository repository.RoleRepository, validate *validator.Validate) RoleService {
	return &RoleServiceImpl{
		RoleRepository: roleRepository,
		Validate:       validate,
	}
}

// CreateRole implements RoleService.
func (service *RoleServiceImpl) CreateRole(ctx *gin.Context, request web.RoleMasterRequest) (web.RoleMasterResponse, error) {
	err := service.Validate.Struct(request)
	utils.PanicIfError(err)

	// LOGIC
	role := domain.RoleMasterModel{
		RoleName: request.RoleName,
	}

	roleResponse, err := service.RoleRepository.CreateRole(ctx, role)

	return web.ToRoleMasterResponseWithError(roleResponse, err)
}

func (service *RoleServiceImpl) UpdateRole(ctx *gin.Context, request web.RoleMasterRequest) (web.RoleMasterResponse, error) {
	err := service.Validate.Struct(request)
	utils.PanicIfError(err)

	roleId := ctx.Params.ByName("roleId")
	id, err := strconv.Atoi(roleId)
	utils.PanicIfError(err)

	// LOGIC
	role := domain.RoleMasterModel{
		RoleName: request.RoleName,
	}

	roleResponse, err := service.RoleRepository.UpdateRole(ctx, role, id)

	if err != nil {
		if err.Error() == "role not found" {
			panic(exception.NewNotFoundError(err.Error()))
		}
	}

	return web.ToRoleMasterResponseWithError(roleResponse, err)
}

func (service *RoleServiceImpl) DeleteRole(ctx *gin.Context) error {
	roleId := ctx.Params.ByName("roleId")
	id, err := strconv.Atoi(roleId)
	utils.PanicIfError(err)
	err = service.RoleRepository.DeleteRole(ctx, id)

	return err
}

func (service *RoleServiceImpl) GetRoleById(ctx *gin.Context) web.RoleMasterResponse {
	roleId := ctx.Params.ByName("roleId")
	id, err := strconv.Atoi(roleId)
	utils.PanicIfError(err)
	roleResponse, err := service.RoleRepository.GetRoleById(ctx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return web.ToRoleMasterResponse(roleResponse)
}

func (service *RoleServiceImpl) GetRoles(ctx *gin.Context, pagination *helper.PaginationInput) (roles []web.RoleMasterResponse, totalRows int64, err error) {
	roleResponse, totalRow, err := service.RoleRepository.GetRoles(ctx, pagination)
	return web.ToRoleMasterResponses(roleResponse, totalRow, err)
}
