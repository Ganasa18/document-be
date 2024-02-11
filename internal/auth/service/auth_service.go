package service

import (
	"github.com/Ganasa18/document-be/internal/auth/model/web"
	"github.com/gin-gonic/gin"
)

type AuthService interface {
	LoginOrRegister(ctx *gin.Context, request web.UserRegisterRequest) (web.UserRegisterResponse, error)
	ForgotLinkPassword(ctx *gin.Context, request web.ForgotPasswordRequest) (string, error)
	ResetPasswordUser(ctx *gin.Context, request web.ResetPasswordRequest) error
}
