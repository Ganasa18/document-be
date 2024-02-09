package controller

import (
	"net/http"

	"github.com/Ganasa18/document-be/internal/auth/model/web"
	"github.com/Ganasa18/document-be/internal/auth/service"
	response "github.com/Ganasa18/document-be/internal/base/model/web"
	"github.com/Ganasa18/document-be/pkg/helper"
	"github.com/gin-gonic/gin"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}

// LoginOrRegister implements AuthController.
func (controller *AuthControllerImpl) LoginOrRegister(ctx *gin.Context) {
	registerRequest := web.UserRegisterRequest{}

	helper.ReadFromRequestBody(ctx.Request, &registerRequest)

	registerResponse, err := controller.AuthService.LoginOrRegister(ctx, registerRequest)

	var statusCode int
	var responseData interface{}

	if err != nil {
		statusCode = http.StatusBadRequest
		responseData = err.Error()
	} else {
		statusCode = http.StatusOK
		responseData = registerResponse
	}

	webResponse := response.WebResponse{
		Code:   statusCode,
		Status: http.StatusText(statusCode),
		Data:   responseData,
	}

	helper.WriteToResponseBody(ctx, http.StatusOK, webResponse)

}
