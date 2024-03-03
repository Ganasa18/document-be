package repository

import (
	"github.com/Ganasa18/document-be/internal/auth/model/domain"
	crudDomain "github.com/Ganasa18/document-be/internal/crud/model/domain"
	"github.com/gin-gonic/gin"
)

type AuthRepository interface {
	LoginOrRegister(ctx *gin.Context, user domain.UserModel, OpenId string, TypeAction string) (domain.UserModel, error)
	CreateOrGetProfile(ctx *gin.Context, profile domain.ProfileUser) (domain.ProfileUser, error)
	ForgotLinkPassword(ctx *gin.Context, forgotData domain.ForgotPasswordLink, email string) error
	ExpiredForgotPassword(ctx *gin.Context, forgotData domain.ForgotPasswordLink) error
	ResetPasswordUser(ctx *gin.Context, user domain.UserModel, hashId string) error
	GetUserMenu(ctx *gin.Context, RoleId int) ([]crudDomain.UserAccessMenuModel, error)
	UpdateUserRole(ctx *gin.Context, user domain.UserModel, adminToken string) (domain.UserModel, error)
}
