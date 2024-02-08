package repository

import (
	"context"

	"github.com/Ganasa18/document-be/internal/auth/model/domain"
	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{
		DB: db,
	}
}

// LoginOrRegister implements AuthRepository.
func (repository *AuthRepositoryImpl) LoginOrRegister(ctx context.Context, user domain.UserModel) domain.UserModel {

	return domain.UserModel{}
}
