package repository

import (
	"errors"

	"github.com/Ganasa18/document-be/internal/crud/model/domain"
	"github.com/Ganasa18/document-be/pkg/helper"
	"github.com/Ganasa18/document-be/pkg/loghelper"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserAccessRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserAccessRepository(db *gorm.DB) UserAccessRepository {
	return &UserAccessRepositoryImpl{
		DB: db,
	}
}

func (repository *UserAccessRepositoryImpl) CreateUserAccess(ctx *gin.Context, userAccess domain.UserAccessMenuModel) (domain.UserAccessMenuModel, error) {

	// Check existing data
	err := repository.DB.Where(&domain.UserAccessMenuModel{RoleId: userAccess.RoleId, MenuId: userAccess.MenuId}).First(&userAccess).Error

	if err != gorm.ErrRecordNotFound {
		loghelper.LogErrorRepository(ctx, "CreateUserAccess", "UserAccessMenuModel", "cannot duplicate insert", err)
		return userAccess, errors.New("cannot duplicate insert")
	}

	err = repository.DB.Model(&domain.UserAccessMenuModel{}).Create(&userAccess).Error
	if err != nil {
		loghelper.LogErrorRepository(ctx, "CreateUserAccess", "UserAccessMenuModel", "when query builder create data", err)
		return userAccess, err
	}

	return userAccess, nil

}

func (repository *UserAccessRepositoryImpl) GetAllUserAccess(ctx *gin.Context, pagination *helper.PaginationInput) (userAccess []domain.UserAccessMenuModel, totalRow int64, err error) {
	totalRow = 0

	offset := (pagination.Page - 1) * pagination.Limit
	queryBuilder := repository.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.OrderBy)

	// GET DATA
	err = queryBuilder.Model(&domain.UserAccessMenuModel{}).Preload("RoleMasterModel").Preload("MenuMasterModel").Where("deleted_at IS NULL").Find(&userAccess).Error

	if err != nil {
		loghelper.LogErrorRepository(ctx, "GetAllUserAccess", "UserAccessMenuModel", "when query builder list data", err)
		return userAccess, totalRow, err
	}

	// ROW COUNT
	searchBuider := repository.DB.Model(&domain.UserAccessMenuModel{}).Where("deleted_at IS NULL")

	errCount := searchBuider.Count(&totalRow).Error

	if errCount != nil {
		loghelper.LogErrorRepository(ctx, "GetAllUserAccess", "UserAccessMenuModel", "when query builder count total rows", err)
		return userAccess, totalRow, errCount
	}
	return userAccess, totalRow, nil
}

func (repository *UserAccessRepositoryImpl) GetUserAccessById(ctx *gin.Context, id int) (userAccess domain.UserAccessMenuModel, err error) {
	err = repository.DB.Model(&domain.UserAccessMenuModel{}).First(&userAccess, id).Error
	if err != nil {
		loghelper.LogErrorRepository(ctx, "GetUserAccessById", "UserAccessMenuModel", "when query builder get data", err)
		return userAccess, errors.New("user access not found")
	}
	return userAccess, nil
}

func (repository *UserAccessRepositoryImpl) UpdateUserAccess(ctx *gin.Context, userAccess domain.UserAccessMenuModel, id int) (domain.UserAccessMenuModel, error) {
	// Check if the user access
	if err := repository.checkIfUserAccessExists(ctx, id); err != nil {
		return userAccess, err
	}

	// Update the user access
	if err := repository.updateUserAccess(ctx, id, userAccess); err != nil {
		return userAccess, err
	}

	// Retrieve the updated user access
	if err := repository.getUpdatedUserAccess(ctx, id, &userAccess); err != nil {
		return userAccess, err
	}

	return userAccess, nil
}

func (repository *UserAccessRepositoryImpl) DeleteUserAccess(ctx *gin.Context, id int) error {
	userAccess := &domain.UserAccessMenuModel{Id: id}

	err := repository.DB.First(&domain.UserAccessMenuModel{}, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		loghelper.LogErrorRepository(ctx, "DeleteUserAccess", "UserAccessMenuModel", "user access not found", err)
		return errors.New("user access not found")
	}

	err = repository.DB.Unscoped().Delete(userAccess).Error

	if err != nil {
		loghelper.LogErrorRepository(ctx, "DeleteUserAccess", "UserAccessMenuModel", "when deleting user access", err)
		return errors.New("failed to delete user access")
	}

	return nil
}

func (repository *UserAccessRepositoryImpl) checkIfUserAccessExists(ctx *gin.Context, id int) error {
	err := repository.DB.First(&domain.UserAccessMenuModel{}, id).Error
	if err != nil {
		loghelper.LogErrorRepository(ctx, "checkIfUserAccessExists", "UserAccessMenuModel", "when querying data", err)
		return errors.New("user access not found")
	}
	return nil
}

func (repository *UserAccessRepositoryImpl) updateUserAccess(ctx *gin.Context, id int, userAccess domain.UserAccessMenuModel) error {

	updateFields := map[string]interface{}{"create": userAccess.Create, "read": userAccess.Read, "update": userAccess.Update, "delete": userAccess.Delete}

	err := repository.DB.Model(&userAccess).
		Where("id = ?", id).
		Updates(updateFields).Error
	if err != nil {
		loghelper.LogErrorRepository(ctx, "UpdateUserAccess", "UserAccessMenuModel", "when updating data", err)
		return errors.New("failed to update user access")
	}
	return nil
}

func (repository *UserAccessRepositoryImpl) getUpdatedUserAccess(ctx *gin.Context, id int, userAccess *domain.UserAccessMenuModel) error {
	err := repository.DB.First(userAccess, id).Error
	if err != nil {
		loghelper.LogErrorRepository(ctx, "getUpdatedUserAccess", "UserAccessMenuModel", "when querying updated data", err)
		return errors.New("failed to get updated user access")
	}
	return nil
}
