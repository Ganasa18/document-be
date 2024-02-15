package service

import (
	"fmt"

	"github.com/Ganasa18/document-be/internal/crud/model/domain"
	"github.com/Ganasa18/document-be/internal/crud/model/web"
	"github.com/Ganasa18/document-be/internal/crud/repository"
	"github.com/Ganasa18/document-be/pkg/helper"
	"github.com/Ganasa18/document-be/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserAccessServiceImpl struct {
	UserAccessRepository repository.UserAccessRepository
	Validate             *validator.Validate
}

func NewUserAccessService(userAccessRepository repository.UserAccessRepository, validate *validator.Validate) UserAccessService {
	return &UserAccessServiceImpl{
		UserAccessRepository: userAccessRepository,
		Validate:             validate,
	}
}

// CreateUserAccess implements UserAccessService.
func (service *UserAccessServiceImpl) CreateUserAccess(ctx *gin.Context, request web.UserAccessRequest) (web.UserAccessResponse, error) {
	err := service.Validate.Struct(request)
	utils.PanicIfError(err)
	// LOGIC
	userAccess := domain.UserAccessMenuModel{
		RoleId: request.RoleId,
		MenuId: request.MenuId,
		Create: request.Create,
		Read:   request.Read,
		Update: request.Update,
		Delete: request.Delete,
	}

	fmt.Println("USERACCESS", userAccess)

	userAccessResponse, err := service.UserAccessRepository.CreateUserAccess(ctx, userAccess)

	return web.ToUserAccessMasterResponseWithError(userAccessResponse, err)
}

// DeleteUserAccess implements UserAccessService.
func (*UserAccessServiceImpl) DeleteUserAccess(ctx *gin.Context) {
	panic("unimplemented")
}

// GetAllUserAccess implements UserAccessService.
func (*UserAccessServiceImpl) GetAllUserAccess(ctx *gin.Context, pagination *helper.PaginationInput) {
	panic("unimplemented")
}

// GetUserAccessById implements UserAccessService.
func (*UserAccessServiceImpl) GetUserAccessById(ctx *gin.Context) {
	panic("unimplemented")
}

// UpdateUserAccess implements UserAccessService.
func (*UserAccessServiceImpl) UpdateUserAccess(ctx *gin.Context) {
	panic("unimplemented")
}
