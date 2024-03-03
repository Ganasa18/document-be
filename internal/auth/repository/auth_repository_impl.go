package repository

import (
	"errors"
	"os"
	"strconv"

	"github.com/Ganasa18/document-be/internal/auth/model/domain"
	crudDomain "github.com/Ganasa18/document-be/internal/crud/model/domain"
	"github.com/Ganasa18/document-be/pkg/loghelper"
	"github.com/Ganasa18/document-be/pkg/utils"
	"github.com/gin-gonic/gin"
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

func (repository *AuthRepositoryImpl) LoginOrRegister(ctx *gin.Context, user domain.UserModel, OpenId string, TypeAction string) (domain.UserModel, error) {
	var plainPassword string
	if OpenId != utils.OPEN_API_GOOGLE {
		plainPassword = *user.Password
	}

	// Check user exist
	err := repository.DB.Where(domain.UserModel{Email: user.Email, OpenId: user.OpenId}).First(&user).Error

	// Get User Role
	roleErr := repository.DB.Where(domain.UserModel{Email: user.Email}).Preload("RoleMasterModel").Find(&user).Error

	if roleErr != nil {
		loghelper.LogErrorRepository(ctx, "LoginOrRegister", "UserModel", "fetching the RoleMasterModel", err)
		return user, errors.New("failed to get role")
	}

	// Register User
	if err != nil {
		// Check user exist login type

		if TypeAction != "" && TypeAction == utils.TYPE_ACTION_LOGIN {
			if err == gorm.ErrRecordNotFound {
				loghelper.LogErrorRepository(ctx, "LoginOrRegister", "UserModel", "user not register", err)
				return user, errors.New("user not register, please register before")
			}
		}
		loghelper.LogErrorRepository(ctx, "LoginOrRegister", "UserModel", "fetching the user", err)

		if OpenId != utils.OPEN_API_GOOGLE {

			if plainPassword == "" {
				loghelper.LogErrorRepository(ctx, "LoginOrRegister", "UserModel", "with email must have password", err)
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
			loghelper.LogErrorRepository(ctx, "LoginOrRegister", "UserModel", "creating user", err)
			return user, err
		}

		// Update User Profile
		err = repository.DB.Model(&domain.UserModel{}).Where("id = ?", user.Id).Updates(map[string]interface{}{"profile_id": user.Id}).Error
		if err != nil {
			loghelper.LogErrorRepository(ctx, "LoginOrRegister", "UserModel", "updating user profile", err)
			return user, err
		}

		return user, nil
	}

	if OpenId != utils.OPEN_API_GOOGLE {

		if plainPassword == "" {
			loghelper.LogErrorRepository(ctx, "LoginOrRegister", "UserModel", "authentication failed", err)
			return user, errors.New("authentication failed")
		}

		storedPasswordHash := *user.Password
		// Compare passwords
		err = bcrypt.CompareHashAndPassword([]byte(storedPasswordHash), []byte(plainPassword))
		if err == bcrypt.ErrMismatchedHashAndPassword {
			loghelper.LogErrorRepository(ctx, "LoginOrRegister", "UserModel", "authentication compare failed", err)
			return user, errors.New("authentication failed")
		}
	}

	return user, nil

}

func (repository *AuthRepositoryImpl) CreateOrGetProfile(ctx *gin.Context, profile domain.ProfileUser) (domain.ProfileUser, error) {
	// Check if the profile exists
	err := repository.DB.Where(domain.ProfileUser{UserId: profile.UserId}).First(&profile).Error
	if err != nil {
		// If the profile does not exist, create it
		if err == gorm.ErrRecordNotFound {
			if createErr := repository.DB.Create(&profile).Error; createErr != nil {
				loghelper.LogErrorRepository(ctx, "CreateOrGetProfile", "UserModel", "creating profile", err)
				return profile, createErr
			}
		} else {
			// Error getting profile
			loghelper.LogErrorRepository(ctx, "CreateOrGetProfile", "ProfileUser", "getting profile", err)
			return profile, err
		}
	}

	return profile, nil
}

func (repository *AuthRepositoryImpl) ForgotLinkPassword(ctx *gin.Context, forgotData domain.ForgotPasswordLink, email string) error {
	var user domain.UserModel

	// Check if the user exists
	err := repository.DB.Where(domain.UserModel{Email: email}).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// User not found
			loghelper.LogErrorRepository(ctx, "ForgotLinkPassword", "UserModel", "user not found", err)
			return errors.New("user not registered")
		}

		// Other database error
		loghelper.LogErrorRepository(ctx, "ForgotLinkPassword", "UserModel", "querying database", err)
		return err
	}

	// User found, create the forgot password link
	err = repository.DB.Create(&forgotData).Error
	if err != nil {
		loghelper.LogErrorRepository(ctx, "ForgotLinkPassword", "ForgotPasswordLink", "creating link", err)
		return err
	}

	return nil
}

func (repository *AuthRepositoryImpl) ExpiredForgotPassword(ctx *gin.Context, forgotData domain.ForgotPasswordLink) error {
	err := repository.DB.Model(&domain.ForgotPasswordLink{}).Where("hash_id = ?", forgotData.HashId).Updates(map[string]interface{}{"is_active": false}).Error
	if err != nil {
		loghelper.LogErrorRepository(ctx, "ExpiredForgotPassword", "ForgotPasswordLink", "update expired forgot password", err)
		return err
	}
	return nil
}

func (repository *AuthRepositoryImpl) ResetPasswordUser(ctx *gin.Context, user domain.UserModel, hashId string) error {

	var checkValid domain.ForgotPasswordLink
	err := repository.DB.Model(&domain.ForgotPasswordLink{}).Where("hash_id = ?", hashId).First(&checkValid).Error
	if err != nil {
		loghelper.LogErrorRepository(ctx, "ResetPasswordUser", "ForgotPasswordLink", "get data forgot password link", err)
		return errors.New("failed to get data forgot password link")
	}
	if checkValid.IsActive != nil && !*checkValid.IsActive {
		loghelper.LogErrorRepository(ctx, "ResetPasswordUser", "ForgotPasswordLink", "expired reset password", err)
		return errors.New("expired reset password")
	}

	err = repository.DB.Model(&domain.UserModel{}).Where("email = ?", user.Email).Updates(map[string]interface{}{"password": user.Password, "updated_at": user.UpdatedAt}).Error
	if err != nil {
		loghelper.LogErrorRepository(ctx, "ResetPasswordUser", "ForgotPasswordLink", "update new password", err)
		return errors.New("failed update user password")
	}

	err = repository.DB.Model(&domain.ForgotPasswordLink{}).Where("hash_id = ?", hashId).Updates(map[string]interface{}{"is_active": false}).Error
	if err != nil {
		loghelper.LogErrorRepository(ctx, "ResetPasswordUser", "ForgotPasswordLink", "update status", err)
		return errors.New("failed update status")
	}

	return nil
}

func (repository *AuthRepositoryImpl) GetUserMenu(ctx *gin.Context, RoleId int) ([]crudDomain.UserAccessMenuModel, error) {
	user_access := []crudDomain.UserAccessMenuModel{}
	err := repository.DB.Model(&user_access).Where("deleted_at IS NULL").Where("role_id = ?", RoleId).Preload("RoleMasterModel").Preload("MenuMasterModel").Find(&user_access).Error

	if err != nil {
		loghelper.LogErrorRepository(ctx, "GetUserMenu", "UserAccessMenuModel", "when query builder list data", err)
		return user_access, err
	}

	return user_access, nil
}

func (repository *AuthRepositoryImpl) UpdateUserRole(ctx *gin.Context, user domain.UserModel, adminToken string) (domain.UserModel, error) {

	adminUser := domain.UserModel{}

	err := repository.DB.Where(domain.UserModel{UserUniqueId: adminToken}).Find(&adminUser).Error
	if err != nil {
		loghelper.LogErrorRepository(ctx, "UpdateUserRole", "UserModel", "get user access", err)
		return domain.UserModel{}, errors.New("failed to get user access")
	}

	// Check Admin
	idAdmin, err := strconv.Atoi(os.Getenv(utils.CONFIG_ADMIN_ID))
	if err != nil {
		loghelper.LogErrorRepository(ctx, "UpdateUserRole", "UserModel", "env admin not setted", err)
		return domain.UserModel{}, errors.New("env admin not setted")
	}

	if *adminUser.RoleId != idAdmin {
		loghelper.LogErrorRepository(ctx, "UpdateUserRole", "UserModel", "does not have privilege to update role", err)
		return domain.UserModel{}, errors.New("user does not have privilege to update role")
	}

	// Update Role
	err = repository.DB.Model(&domain.UserModel{}).Where("user_unique_id = ?", user.UserUniqueId).Updates(map[string]interface{}{"role_id": *user.RoleId}).Error

	if err != nil {
		loghelper.LogErrorRepository(ctx, "UpdateUserRole", "UserModel", "failed to update data role", err)
		return domain.UserModel{}, errors.New("failed to update data")
	}

	// Get Return Data
	err = repository.DB.Where(domain.UserModel{UserUniqueId: user.UserUniqueId}).Preload("RoleMasterModel").First(&user).Error
	if err != nil {
		loghelper.LogErrorRepository(ctx, "UpdateUserRole", "UserModel", "failed to get return user data", err)
		return domain.UserModel{}, errors.New("failed to get user data")
	}

	return user, nil
}
