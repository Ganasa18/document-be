package controller

import (
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

func (*UserAccessControllerImpl) DeleteUserAccess(ctx *gin.Context) {
	panic("unimplemented")
}

func (*UserAccessControllerImpl) GetAllUserAccess(ctx *gin.Context) {
	panic("unimplemented")
}

func (*UserAccessControllerImpl) GetUserAccessById(ctx *gin.Context) {
	panic("unimplemented")
}

func (*UserAccessControllerImpl) UpdateUserAccess(ctx *gin.Context) {
	panic("unimplemented")
}
