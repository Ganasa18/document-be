package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Ganasa18/document-be/internal/auth/model/domain"
	"github.com/Ganasa18/document-be/pkg/loghelper"
	"github.com/Ganasa18/document-be/pkg/utils"
	"golang.org/x/crypto/bcrypt"
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

func (repository *AuthRepositoryImpl) LoginOrRegister(ctx context.Context, user domain.UserModel, OpenId string) (domain.UserModel, error) {
	var plainPassword string
	if OpenId != utils.OPEN_API_GOOGLE {
		plainPassword = *user.Password
	}

	err := repository.DB.Where(domain.UserModel{Email: user.Email, OpenId: user.OpenId}).First(&user).Error

	// GET USER ROLE
	_ = repository.DB.Where(domain.UserModel{Email: user.Email}).Preload("RoleMasterModel").Find(&user).Error

	// REGISTER USER
	if err != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("LoginOrRegister | Error fetching the user, err:%s", err.Error()))
		if OpenId != utils.OPEN_API_GOOGLE {

			if plainPassword == "" {
				return user, errors.New("with email must have password")
			}

			// Hashing the password with the default cost of 10
			hashedPassword, errHashedPassword := bcrypt.GenerateFromPassword([]byte(*user.Password), bcrypt.DefaultCost)
			utils.IsErrorDoPanic(errHashedPassword)
			hashedPasswordStr := string(hashedPassword)
			user.Password = &hashedPasswordStr
		}
		err = repository.DB.Create(&user).Error
		if err != nil {
			loghelper.Errorln(ctx, fmt.Sprintf("LoginOrRegister | Error creating user, err:%s", err.Error()))
			return user, err
		}
		return user, nil
	}

	if OpenId != utils.OPEN_API_GOOGLE {

		if plainPassword == "" {
			return user, errors.New("authentication failed")
		}

		storedPasswordHash := *user.Password
		// Compare passwords
		err = bcrypt.CompareHashAndPassword([]byte(storedPasswordHash), []byte(plainPassword))
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return user, errors.New("authentication failed")
		}
	}

	return user, nil

}

func (repository *AuthRepositoryImpl) ForgotLinkPassword(ctx context.Context, forgotData domain.ForgotPasswordLink, email string) error {
	var user domain.UserModel

	// Check if the user exists
	err := repository.DB.Where(domain.UserModel{Email: email}).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// User not found
			loghelper.Errorln(ctx, "ForgotLinkPassword | User not found")
			return errors.New("user not registered")
		}

		// Other database error
		loghelper.Errorln(ctx, fmt.Sprintf("ForgotLinkPassword | Error querying database, err:%s", err.Error()))
		return err
	}

	// User found, create the forgot password link
	err = repository.DB.Create(&forgotData).Error
	if err != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("ForgotLinkPassword | Error creating link, err:%s", err.Error()))
		return err
	}

	return nil
}

func (repository *AuthRepositoryImpl) ExpiredForgotPassword(ctx context.Context, forgotData domain.ForgotPasswordLink) error {
	err := repository.DB.Model(&domain.ForgotPasswordLink{}).Where("hash_id = ?", forgotData.HashId).Updates(map[string]interface{}{"is_active": false}).Error
	if err != nil {
		return err
	}
	return nil
}
