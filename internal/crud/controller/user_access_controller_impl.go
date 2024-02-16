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

type UserAccessControllerImpl struct {
	UserAccessService service.UserAccessService
}

func NewUserAccessControllere(userAccessService service.UserAccessService) UserAccessController {
	return &UserAccessControllerImpl{
		UserAccessService: userAccessService,
	}
}

// CreateUserAccess implements UserAccessController.
func (controller *UserAccessControllerImpl) CreateUserAccess(ctx *gin.Context) {
	userAccessRequest := web.UserAccessRequest{}
	helper.ReadFromRequestBody(ctx.Request, &userAccessRequest)

	userAccessResponse, err := controller.UserAccessService.CreateUserAccess(ctx, userAccessRequest)

	var statusCode int
	var responseData interface{}

	if err != nil {
		statusCode = http.StatusBadRequest
		responseData = err.Error()
	} else {
		statusCode = http.StatusOK
		responseData = userAccessResponse
	}

	webResponse := response.WebResponse{
		Code:   statusCode,
		Status: http.StatusText(statusCode),
		Data:   responseData,
	}

	helper.WriteToResponseBody(ctx, statusCode, webResponse)
}

func (controller *UserAccessControllerImpl) GetAllUserAccess(ctx *gin.Context) {
	pagination := helper.Pagination(ctx)
	menuResponse, totalRow, err := controller.UserAccessService.GetAllUserAccess(ctx, &pagination)

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

func (controller *UserAccessControllerImpl) GetUserAccessById(ctx *gin.Context) {
	userAccessResponse := controller.UserAccessService.GetUserAccessById(ctx)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   userAccessResponse,
	}

	helper.WriteToResponseBody(ctx, http.StatusOK, webResponse)
}

func (*UserAccessControllerImpl) UpdateUserAccess(ctx *gin.Context) {
	panic("unimplemented")
}

func (controller *UserAccessControllerImpl) DeleteUserAccess(ctx *gin.Context) {
	err := controller.UserAccessService.DeleteUserAccess(ctx)

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
