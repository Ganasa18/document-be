package service

import (
	"github.com/Ganasa18/document-be/internal/auth/model/web"
	crud "github.com/Ganasa18/document-be/internal/crud/model/web"
	"github.com/gin-gonic/gin"
)

type AuthService interface {
	LoginOrRegister(ctx *gin.Context, request web.UserRegisterRequest) (web.UserBaseResponse, error)
	ForgotLinkPassword(ctx *gin.Context, request web.ForgotPasswordRequest) (string, error)
	ResetPasswordUser(ctx *gin.Context, request web.ResetPasswordRequest) error
	GetUserMenu(ctx *gin.Context, RoleId int) ([]crud.MenuMasterUserResponse, error)
	UpdateUserRole(ctx *gin.Context, request web.UpdateUserAccessRequest) (web.UserBaseResponse, error)
}
