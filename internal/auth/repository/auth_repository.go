package repository

import (
	"context"

	"github.com/Ganasa18/document-be/internal/auth/model/domain"
)

type AuthRepository interface {
	LoginOrRegister(ctx context.Context, user domain.UserModel, OpenId string) (domain.UserModel, error)
	CreateOrGetProfile(ctx context.Context, profile domain.ProfileUser) (domain.ProfileUser, error)
	ForgotLinkPassword(ctx context.Context, forgotData domain.ForgotPasswordLink, email string) error
	ExpiredForgotPassword(ctx context.Context, forgotData domain.ForgotPasswordLink) error
	ResetPasswordUser(ctx context.Context, user domain.UserModel, hashId string) error
}
