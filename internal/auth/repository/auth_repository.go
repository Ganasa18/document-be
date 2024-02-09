package repository

import (
	"context"

	"github.com/Ganasa18/document-be/internal/auth/model/domain"
)

type AuthRepository interface {
	LoginOrRegister(ctx context.Context, user domain.UserModel, OpenId string) (domain.UserModel, error)
}
