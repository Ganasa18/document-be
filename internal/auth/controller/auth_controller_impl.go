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

	registerResponse := controller.AuthService.LoginOrRegister(ctx, registerRequest)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   registerResponse,
	}

	helper.WriteToResponseBody(ctx, http.StatusOK, webResponse)

}
