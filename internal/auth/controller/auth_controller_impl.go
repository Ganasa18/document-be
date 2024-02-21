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

func (controller *AuthControllerImpl) LoginOrRegister(ctx *gin.Context) {
	registerRequest := web.UserRegisterRequest{}

	helper.ReadFromRequestBody(ctx.Request, &registerRequest)

	registerResponse, err := controller.AuthService.LoginOrRegister(ctx, registerRequest)

	// GET USER MENU
	userMenu, _ := controller.AuthService.GetUserMenu(ctx, registerResponse.Role.Id)

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

		// Create a new cookie
		ctx.SetCookie(utils.COOKIE_TOKEN, tokenString, 24*3600, "/", "", false, true)

		registerResponse.Token = &tokenString
		registerResponse.Menu = &userMenu
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

	value, err := controller.AuthService.ForgotLinkPassword(ctx, forgotRequest)

	var statusCode int
	var responseData interface{}

	if err != nil {
		statusCode = http.StatusBadRequest
		responseData = err.Error()
	} else {
		statusCode = http.StatusOK
		responseData = web.ForgotPasswordResponseSuccess{
			Status: "success send link",
			HashId: value,
		}
	}

	webResponse := response.WebResponse{
		Code:   statusCode,
		Status: http.StatusText(statusCode),
		Data:   responseData,
	}

	helper.WriteToResponseBody(ctx, statusCode, webResponse)
}

func (controller *AuthControllerImpl) ResetPasswordUser(ctx *gin.Context) {

	passwordResetRequest := web.ResetPasswordRequest{}

	helper.ReadFromRequestBody(ctx.Request, &passwordResetRequest)

	err := controller.AuthService.ResetPasswordUser(ctx, passwordResetRequest)

	var statusCode int
	var responseData interface{}

	if err != nil {
		statusCode = http.StatusBadRequest
		responseData = err.Error()
	} else {
		statusCode = http.StatusOK
		responseData = "success reset password"
	}

	webResponse := response.WebResponse{
		Code:   statusCode,
		Status: http.StatusText(statusCode),
		Data:   responseData,
	}

	helper.WriteToResponseBody(ctx, statusCode, webResponse)
}

func (controller *AuthControllerImpl) UpdateUserRole(ctx *gin.Context) {

	requestUpdateUserAccess := web.UpdateUserAccessRequest{}
	helper.ReadFromRequestBody(ctx.Request, &requestUpdateUserAccess)

	responseUpdate, err := controller.AuthService.UpdateUserRole(ctx, requestUpdateUserAccess)

	var statusCode int
	var responseData interface{}

	if err != nil {
		statusCode = http.StatusBadRequest
		responseData = err.Error()
	} else {
		statusCode = http.StatusOK
		responseData = responseUpdate
	}

	webResponse := response.WebResponse{
		Code:   statusCode,
		Status: http.StatusText(statusCode),
		Data:   responseData,
	}

	helper.WriteToResponseBody(ctx, statusCode, webResponse)

}
