package controller

import (
	"math"
	"net/http"

	response "github.com/Ganasa18/document-be/internal/base/model/web"
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
