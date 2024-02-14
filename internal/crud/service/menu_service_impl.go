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

func (service *MenuServiceImpl) GetAllMenu(ctx *gin.Context, pagination *helper.PaginationInput) ([]web.MenuMasterResponse, int64, error) {
	menuResponse, totalRow, err := service.MenuRepository.GetAllMenu(ctx, pagination)
	return web.ToMenuMasterResponses(menuResponse, totalRow, err)
}

func (service *MenuServiceImpl) GetMenuById(ctx *gin.Context) web.MenuMasterResponse {
	menuId := ctx.Params.ByName("menuId")
	id, err := strconv.Atoi(menuId)
	utils.PanicIfError(err)
	menuResponse, err := service.MenuRepository.GetMenuById(ctx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return web.ToMenuMasterResponse(menuResponse)
}

func (service *MenuServiceImpl) UpdateMenu(ctx *gin.Context, request web.MenuMasterRequestEdit) (web.MenuMasterResponse, error) {
	err := service.Validate.Struct(request)
	utils.PanicIfError(err)

	menuId := ctx.Params.ByName("menuId")
	id, err := strconv.Atoi(menuId)
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

	menuResponse, err := service.MenuRepository.UpdateMenu(ctx, menu, id)

	if err != nil {
		if err.Error() == "menu not found" {
			panic(exception.NewNotFoundError(err.Error()))
		}
	}

	return web.ToMenuMasterResponseWithError(menuResponse, err)
}

func (service *MenuServiceImpl) DeleteMenu(ctx *gin.Context) error {
	menuId := ctx.Params.ByName("menuId")
	id, err := strconv.Atoi(menuId)
	utils.PanicIfError(err)
	err = service.MenuRepository.DeleteMenu(ctx, id)

	return err
}
