package service

import (
	"github.com/Ganasa18/document-be/internal/crud/model/domain"
	"github.com/Ganasa18/document-be/internal/crud/model/web"
	"github.com/Ganasa18/document-be/internal/crud/repository"
	"github.com/Ganasa18/document-be/pkg/helper"
	"github.com/Ganasa18/document-be/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type MenuServiceImpl struct {
	MenuRepository repository.MenuRepository
	Validate       *validator.Validate
}

func NewMenuService(menuRepository repository.MenuRepository, validate *validator.Validate) MenuService {
	return &MenuServiceImpl{
		MenuRepository: menuRepository,
		Validate:       validate,
	}
}

// CreateMenu implements MenuService.
func (service *MenuServiceImpl) CreateMenu(ctx *gin.Context, request web.MenuMasterRequest) (web.MenuMasterResponse, error) {
	err := service.Validate.Struct(request)
	utils.PanicIfError(err)

	// LOGIC
	menu := domain.MenuMasterModel{
		Name:       request.Name,
		Title:      request.Title,
		Path:       request.Path,
		IconName:   request.IconName,
		IsSubMenu:  request.IsSubMenu,
		ParentName: request.ParentName,
	}

	menuResponse, err := service.MenuRepository.CreateMenu(ctx, menu)

	return web.ToMenuMasterResponseWithError(menuResponse, err)
}

func (service *MenuServiceImpl) DeleteMenu(ctx *gin.Context) {
	panic("unimplemented")
}

func (service *MenuServiceImpl) GetAllMenu(ctx *gin.Context, pagination *helper.PaginationInput) ([]web.MenuMasterResponse, int64, error) {
	menuResponse, totalRow, err := service.MenuRepository.GetAllMenu(ctx, pagination)
	return web.ToMenuMasterResponses(menuResponse, totalRow, err)
}

func (service *MenuServiceImpl) GetMenuById(ctx *gin.Context) web.MenuMasterResponse {
	panic("unimplemented")
}

func (service *MenuServiceImpl) UpdateMenu(ctx *gin.Context, request web.MenuMasterRequest) {
	panic("unimplemented")
}
