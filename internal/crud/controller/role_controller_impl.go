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

type RoleControllerImpl struct {
	RoleService service.RoleService
}

func NewRoleController(roleService service.RoleService) RoleController {
	return &RoleControllerImpl{
		RoleService: roleService,
	}
}

func (controller *RoleControllerImpl) CreateRole(ctx *gin.Context) {
	roleRequest := web.RoleMasterRequest{}
	helper.ReadFromRequestBody(ctx.Request, &roleRequest)

	roleResponse, err := controller.RoleService.CreateRole(ctx, roleRequest)

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

func (controller *RoleControllerImpl) UpdateRole(ctx *gin.Context) {
	roleRequest := web.RoleMasterRequest{}
	helper.ReadFromRequestBody(ctx.Request, &roleRequest)

	roleResponse, err := controller.RoleService.UpdateRole(ctx, roleRequest)

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

// DeleteRole implements RoleController.
func (controller *RoleControllerImpl) DeleteRole(ctx *gin.Context) {
	err := controller.RoleService.DeleteRole(ctx)

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

func (controller *RoleControllerImpl) GetRoleById(ctx *gin.Context) {

	roleResponse := controller.RoleService.GetRoleById(ctx)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   roleResponse,
	}

	helper.WriteToResponseBody(ctx, http.StatusOK, webResponse)

}

func (controller *RoleControllerImpl) GetRoles(ctx *gin.Context) {

	pagination := helper.Pagination(ctx)
	roleResponse, totalRow, err := controller.RoleService.GetRoles(ctx, &pagination)

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
		responseData = roleResponse
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
