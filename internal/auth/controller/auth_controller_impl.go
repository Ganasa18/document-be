package controller

import (
	"net/http"

	"github.com/Ganasa18/document-be/internal/auth/model/web"
	"github.com/Ganasa18/document-be/internal/auth/service"
	response "github.com/Ganasa18/document-be/internal/base/model/web"
	"github.com/Ganasa18/document-be/pkg/helper"
	"github.com/Ganasa18/document-be/pkg/utils"
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
		// SET ASSIGN JWT TOKEN
		tokenString, err := helper.CreateToken(registerResponse.UserUniqueId, registerResponse.Email)
		utils.PanicIfError(err)
		registerResponse.Token = &tokenString
		responseData = registerResponse
	}

	webResponse := response.WebResponse{
		Code:   statusCode,
		Status: http.StatusText(statusCode),
		Data:   responseData,
	}

	helper.WriteToResponseBody(ctx, statusCode, webResponse)

}

func (controller *AuthControllerImpl) ForgotLinkPassword(ctx *gin.Context) {

	forgotRequest := web.ForgotPasswordRequest{}

	helper.ReadFromRequestBody(ctx.Request, &forgotRequest)

	err := controller.AuthService.ForgotLinkPassword(ctx, forgotRequest)

	var statusCode int
	var responseData interface{}

	if err != nil {
		statusCode = http.StatusBadRequest
		responseData = err.Error()
	} else {
		statusCode = http.StatusOK
		responseData = "success send link"
	}

	webResponse := response.WebResponse{
		Code:   statusCode,
		Status: http.StatusText(statusCode),
		Data:   responseData,
	}

	helper.WriteToResponseBody(ctx, statusCode, webResponse)
}
