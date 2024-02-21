package repository

import (
	"context"

	"github.com/Ganasa18/document-be/internal/auth/model/domain"
	crudDomain "github.com/Ganasa18/document-be/internal/crud/model/domain"
)

type AuthRepository interface {
	LoginOrRegister(ctx context.Context, user domain.UserModel, OpenId string, TypeAction string) (domain.UserModel, error)
	CreateOrGetProfile(ctx context.Context, profile domain.ProfileUser) (domain.ProfileUser, error)
	ForgotLinkPassword(ctx context.Context, forgotData domain.ForgotPasswordLink, email string) error
	ExpiredForgotPassword(ctx context.Context, forgotData domain.ForgotPasswordLink) error
	ResetPasswordUser(ctx context.Context, user domain.UserModel, hashId string) error
	GetUserMenu(ctx context.Context, RoleId int) ([]crudDomain.UserAccessMenuModel, error)
	UpdateUserRole(ctx context.Context, user domain.UserModel, adminToken string) (domain.UserModel, error)
}
