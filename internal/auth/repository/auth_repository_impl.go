package repository

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"

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

func (repository *AuthRepositoryImpl) LoginOrRegister(ctx context.Context, user domain.UserModel, OpenId string, TypeAction string) (domain.UserModel, error) {
	var plainPassword string
	if OpenId != utils.OPEN_API_GOOGLE {
		plainPassword = *user.Password
	}

	// Check user exist
	err := repository.DB.Where(domain.UserModel{Email: user.Email, OpenId: user.OpenId}).First(&user).Error

	// Get User Role
	roleErr := repository.DB.Where(domain.UserModel{Email: user.Email}).Preload("RoleMasterModel").Find(&user).Error

	if roleErr != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("LoginOrRegister | UserModel | Repository | Error fetching the RoleMasterModel, err:%s", err.Error()))
		return user, errors.New("failed to get role")
	}

	// Register User
	if err != nil {
		// Check user exist login type

		if TypeAction != "" && TypeAction == utils.TYPE_ACTION_LOGIN {
			if err == gorm.ErrRecordNotFound {
				loghelper.Errorln(ctx, fmt.Sprintf("LoginOrRegister | UserModel | Repository | Error user not register, err:%s", err.Error()))
				return user, errors.New("user not register, please register before")
			}
		}

		loghelper.Errorln(ctx, fmt.Sprintf("LoginOrRegister | UserModel | Repository | Error fetching the user, err:%s", err.Error()))

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
			loghelper.Errorln(ctx, fmt.Sprintf("LoginOrRegister | UserModel | Repository | Error creating user, err:%s", err.Error()))
			return user, err
		}

		// Update User Profile
		err = repository.DB.Model(&domain.UserModel{}).Where("id = ?", user.Id).Updates(map[string]interface{}{"profile_id": user.Id}).Error
		if err != nil {
			loghelper.Errorln(ctx, fmt.Sprintf("LoginOrRegister | UserModel | Repository | Error updating user profile, err:%s", err.Error()))
			return user, err
		}

		return user, nil
	}

	if OpenId != utils.OPEN_API_GOOGLE {

		if plainPassword == "" {
			loghelper.Errorln(ctx, fmt.Sprintf("LoginOrRegister | UserModel | Repository | Error authentication failed, err:%s", err.Error()))
			return user, errors.New("authentication failed")
		}

		storedPasswordHash := *user.Password
		// Compare passwords
		err = bcrypt.CompareHashAndPassword([]byte(storedPasswordHash), []byte(plainPassword))
		if err == bcrypt.ErrMismatchedHashAndPassword {
			loghelper.Errorln(ctx, fmt.Sprintf("LoginOrRegister | UserModel | Repository | Error authentication compare failed, err:%s", err.Error()))
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
				loghelper.Errorln(ctx, fmt.Sprintf("CreateOrGetProfile | ProfileUser | Repository | Error creating profile, err:%s", createErr.Error()))
				return profile, createErr
			}
		} else {
			// Error getting profile
			loghelper.Errorln(ctx, fmt.Sprintf("CreateOrGetProfile | ProfileUser | Repository | Error getting profile err:%s", err.Error()))
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
			loghelper.Errorln(ctx, "ForgotLinkPassword | UserModel | Repository | User not found")
			return errors.New("user not registered")
		}

		// Other database error
		loghelper.Errorln(ctx, fmt.Sprintf("ForgotLinkPassword | UserModel | Repository | Error querying database, err:%s", err.Error()))
		return err
	}

	// User found, create the forgot password link
	err = repository.DB.Create(&forgotData).Error
	if err != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("ForgotLinkPassword | ForgotPasswordLink | Repository | Error creating link, err:%s", err.Error()))
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
		loghelper.Errorln(ctx, fmt.Sprintf("ResetPasswordUser | ForgotPasswordLink | Repository | Error get data forgot password link, err:%s", err.Error()))
		return errors.New("failed to get data forgot password link")
	}
	if checkValid.IsActive != nil && !*checkValid.IsActive {
		loghelper.Errorln(ctx, fmt.Sprintf("ResetPasswordUser | ForgotPasswordLink | Repository | Error expired reset password, err:%s", err.Error()))
		return errors.New("expired reset password")
	}

	err = repository.DB.Model(&domain.UserModel{}).Where("email = ?", user.Email).Updates(map[string]interface{}{"password": user.Password, "updated_at": user.UpdatedAt}).Error
	if err != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("ResetPasswordUser | ForgotPasswordLink | Repository | Error update new password, err:%s", err.Error()))
		return errors.New("failed update user password")
	}

	err = repository.DB.Model(&domain.ForgotPasswordLink{}).Where("hash_id = ?", hashId).Updates(map[string]interface{}{"is_active": false}).Error
	if err != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("ResetPasswordUser | ForgotPasswordLink | Repository | Error update status, err:%s", err.Error()))
		return errors.New("failed update status")
	}

	return nil
}

func (repository *AuthRepositoryImpl) GetUserMenu(ctx context.Context, RoleId int) ([]crudDomain.UserAccessMenuModel, error) {
	user_access := []crudDomain.UserAccessMenuModel{}
	err := repository.DB.Model(&user_access).Where("deleted_at IS NULL").Where("role_id = ?", RoleId).Preload("RoleMasterModel").Preload("MenuMasterModel").Find(&user_access).Error

	if err != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("GetUserMenu | UserAccessMenuModel | Repository | Error when Query builder list data, err:%s", err.Error()))
		return user_access, err
	}

	return user_access, nil
}

func (repository *AuthRepositoryImpl) UpdateUserRole(ctx context.Context, user domain.UserModel, adminToken string) (domain.UserModel, error) {

	adminUser := domain.UserModel{}

	err := repository.DB.Where(domain.UserModel{UserUniqueId: adminToken}).Find(&adminUser).Error
	if err != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("UpdateUserRole | UserModel | Repository | Error get user access, err:%s", err.Error()))
		return domain.UserModel{}, errors.New("failed to get user access")
	}

	// Check Admin
	idAdmin, err := strconv.Atoi(os.Getenv(utils.CONFIG_ADMIN_ID))
	if err != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("UpdateUserRole | UserModel | Repository | Error env admin not setted, err:%s", err.Error()))
		return domain.UserModel{}, errors.New("env admin not setted")
	}

	if *adminUser.RoleId != idAdmin {
		loghelper.Errorln(ctx, fmt.Sprintf("UpdateUserRole | UserModel | Repository | Error does not have privilege to update role, err:%s", err.Error()))
		return domain.UserModel{}, errors.New("user does not have privilege to update role")
	}

	// Update Role
	err = repository.DB.Model(&domain.UserModel{}).Where("user_unique_id = ?", user.UserUniqueId).Updates(map[string]interface{}{"role_id": *user.RoleId}).Error

	if err != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("UpdateUserRole | UserModel | Repository | Error failed to update data role, err:%s", err.Error()))
		return domain.UserModel{}, errors.New("failed to update data")
	}

	// Get Return Data
	err = repository.DB.Where(domain.UserModel{UserUniqueId: user.UserUniqueId}).Preload("RoleMasterModel").First(&user).Error
	if err != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("UpdateUserRole | UserModel | Repository | Error failed to get return user data, err:%s", err.Error()))
		return domain.UserModel{}, errors.New("failed to get user data")
	}

	return user, nil
}
