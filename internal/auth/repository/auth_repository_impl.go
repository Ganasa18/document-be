package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Ganasa18/document-be/internal/auth/model/domain"
	crudDomain "github.com/Ganasa18/document-be/internal/crud/model/domain"
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

		// Update User Profile
		err = repository.DB.Model(&domain.UserModel{}).Where("id = ?", user.Id).Updates(map[string]interface{}{"profile_id": user.Id}).Error
		if err != nil {
			loghelper.Errorln(ctx, fmt.Sprintf("LoginOrRegister | Error updating user profile, err:%s", err.Error()))
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

func (repository *AuthRepositoryImpl) CreateOrGetProfile(ctx context.Context, profile domain.ProfileUser) (domain.ProfileUser, error) {
	// Check if the profile exists
	err := repository.DB.Where(domain.ProfileUser{UserId: profile.UserId}).First(&profile).Error
	if err != nil {
		// If the profile does not exist, create it
		if err == gorm.ErrRecordNotFound {
			if createErr := repository.DB.Create(&profile).Error; createErr != nil {
				loghelper.Errorln(ctx, fmt.Sprintf("CreateOrGetProfile | Error creating profile, err:%s", createErr.Error()))
				return profile, createErr
			}
		} else {
			// Error getting profile
			loghelper.Errorln(ctx, fmt.Sprintf("CreateOrGetProfile | Error getting profile err:%s", err.Error()))
			return profile, err
		}
	}

	return profile, nil
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

func (repository *AuthRepositoryImpl) ResetPasswordUser(ctx context.Context, user domain.UserModel, hashId string) error {

	var checkValid domain.ForgotPasswordLink
	err := repository.DB.Model(&domain.ForgotPasswordLink{}).Where("hash_id = ?", hashId).First(&checkValid).Error
	if err != nil {
		return err
	}
	if checkValid.IsActive != nil && !*checkValid.IsActive {
		return errors.New("expired reset password")
	}

	err = repository.DB.Model(&domain.UserModel{}).Where("email = ?", user.Email).Updates(map[string]interface{}{"password": user.Password, "updated_at": user.UpdatedAt}).Error
	if err != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("ResetPasswordUser | Error update new password, err:%s", err.Error()))
		return errors.New("failed update user password")
	}

	err = repository.DB.Model(&domain.ForgotPasswordLink{}).Where("hash_id = ?", hashId).Updates(map[string]interface{}{"is_active": false}).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository *AuthRepositoryImpl) GetUserMenu(ctx context.Context, RoleId int) ([]crudDomain.UserAccessMenuModel, error) {
	user_access := []crudDomain.UserAccessMenuModel{}
	err := repository.DB.Model(&user_access).Where("role_id = ?", RoleId).Preload("RoleMasterModel").Preload("MenuMasterModel").Find(&user_access).Error

	if err != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("GetUserMenu | Error when Query builder list data, err:%s", err.Error()))
		return user_access, err
	}

	return user_access, nil
}
