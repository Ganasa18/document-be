package controller

import (
	"math"
	"net/http"

	response "github.com/Ganasa18/document-be/internal/base/model/web"
	"github.com/Ganasa18/document-be/internal/crud/model/web"
	"github.com/Ganasa18/document-be/internal/crud/service"
	"github.com/Ganasa18/document-be/pkg/helper"
	"github.com/gin-gonic/gin"
)

type MenuControllerImpl struct {
	MenuService service.MenuService
}

func NewMenuControllere(menuService service.MenuService) MenuController {
	return &MenuControllerImpl{
		MenuService: menuService,
	}
}

func (controller *MenuControllerImpl) CreateMenu(ctx *gin.Context) {
	menuRequest := web.MenuMasterRequest{}

	helper.ReadFromRequestBody(ctx.Request, &menuRequest)

	menuResponse, err := controller.MenuService.CreateMenu(ctx, menuRequest)

	var statusCode int
	var responseData interface{}
	if err != nil {
		statusCode = http.StatusBadRequest
		responseData = err.Error()
	} else {
		statusCode = http.StatusOK
		responseData = menuResponse
	}

	webResponse := response.WebResponse{
		Code:   statusCode,
		Status: http.StatusText(statusCode),
		Data:   responseData,
	}

	helper.WriteToResponseBody(ctx, statusCode, webResponse)
}

func (controller *MenuControllerImpl) GetAllMenu(ctx *gin.Context) {
	pagination := helper.Pagination(ctx)
	menuResponse, totalRow, err := controller.MenuService.GetAllMenu(ctx, &pagination)

	var statusCode int
	var responseData interface{}

	totalPage := 0
	if totalRow > 0 {
		totalPage = int(math.Ceil(float64(totalRow) / float64(pagination.Limit)))
	}

	if err != nil {
		statusCode = http.StatusBadRequest
		responseData = err.Error()
	} else {
		statusCode = http.StatusOK
		responseData = menuResponse
	}

	pageInfo := response.PageInfoResponse{
		Total:       totalRow,
		PerPage:     pagination.Limit,
		CurrentPage: pagination.Page,
		TotalPage:   totalPage,
	}

	webResponse := response.WebResponsePaginate{
		Code:     statusCode,
		Status:   http.StatusText(statusCode),
		Data:     responseData,
		PageInfo: pageInfo,
	}

	helper.WriteToResponseBody(ctx, statusCode, webResponse)
}

func (controller *MenuControllerImpl) GetMenuById(ctx *gin.Context) {
	menuResponse := controller.MenuService.GetMenuById(ctx)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   menuResponse,
	}

	helper.WriteToResponseBody(ctx, http.StatusOK, webResponse)
}

func (controller *MenuControllerImpl) UpdateMenu(ctx *gin.Context) {
	menuRequest := web.MenuMasterRequestEdit{}
	helper.ReadFromRequestBody(ctx.Request, &menuRequest)

	roleResponse, err := controller.MenuService.UpdateMenu(ctx, menuRequest)

	var statusCode int
	var responseData interface{}

	if err != nil {
		statusCode = http.StatusBadRequest
		responseData = err.Error()
	} else {
		statusCode = http.StatusOK
		responseData = roleResponse
	}

	webResponse := response.WebResponse{
		Code:   statusCode,
		Status: http.StatusText(statusCode),
		Data:   responseData,
	}

	helper.WriteToResponseBody(ctx, statusCode, webResponse)
}

func (controller *MenuControllerImpl) DeleteMenu(ctx *gin.Context) {
	err := controller.MenuService.DeleteMenu(ctx)

	var statusCode int
	var responseData interface{}

	if err != nil {
		statusCode = http.StatusBadRequest
		responseData = err.Error()
	} else {
		statusCode = http.StatusOK
		responseData = "success deleted"
	}

	webResponse := response.WebResponse{
		Code:   statusCode,
		Status: http.StatusText(statusCode),
		Data:   responseData,
	}

	helper.WriteToResponseBody(ctx, statusCode, webResponse)
}
