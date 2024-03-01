package repository

import (
	"errors"
	"fmt"

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
		loghelper.Errorln(ctx, fmt.Sprintf("CreateUserAccess | UserAccessMenuModel | Repository | Error cannot duplicate insert, err:%s", err.Error()))
		return userAccess, errors.New("cannot duplicate insert")
	}

	err = repository.DB.Model(&domain.UserAccessMenuModel{}).Create(&userAccess).Error
	if err != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("CreateUserAccess | UserAccessMenuModel | Repository | Error when Query builder create data, err:%s", err.Error()))
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
		loghelper.Errorln(ctx, fmt.Sprintf("GetAllUserAccess | UserAccessMenuModel | Repository | Error when Query builder list data, err:%s", err.Error()))
		return userAccess, totalRow, err
	}

	// ROW COUNT
	searchBuider := repository.DB.Model(&domain.UserAccessMenuModel{}).Where("deleted_at IS NULL")

	errCount := searchBuider.Count(&totalRow).Error

	if errCount != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("GetAllUserAccess | UserAccessMenuModel | Repository | Error when Query builder count total rows, err:%s", errCount.Error()))
		return userAccess, totalRow, errCount
	}
	return userAccess, totalRow, nil
}

func (repository *UserAccessRepositoryImpl) GetUserAccessById(ctx *gin.Context, id int) (userAccess domain.UserAccessMenuModel, err error) {
	err = repository.DB.Model(&domain.UserAccessMenuModel{}).First(&userAccess, id).Error
	if err != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("GetUserAccessById | UserAccessMenuModel | Repository | Error when Query builder get data, err:%s", err.Error()))
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
		loghelper.Errorln(ctx, fmt.Sprintf("DeleteUserAccess | UserAccessMenuModel | Repository | Error user access not found, err: %s", err.Error()))
		return errors.New("user access not found")
	}

	err = repository.DB.Unscoped().Delete(userAccess).Error

	if err != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("DeleteUserAccess | UserAccessMenuModel | Repository | Error when deleting user access, err: %s", err.Error()))
		return errors.New("failed to delete user access")
	}

	return nil
}

func (repository *UserAccessRepositoryImpl) checkIfUserAccessExists(ctx *gin.Context, id int) error {
	err := repository.DB.First(&domain.UserAccessMenuModel{}, id).Error
	if err != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("checkIfUserAccessExists | UserAccessMenuModel | Repository | Error when querying data, err:%s", err.Error()))
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
		loghelper.Errorln(ctx, fmt.Sprintf("UpdateUserAccess | UserAccessMenuModel | Repository | Error when updating data, err:%s", err.Error()))
		return errors.New("failed to update user access")
	}
	return nil
}

func (repository *UserAccessRepositoryImpl) getUpdatedUserAccess(ctx *gin.Context, id int, userAccess *domain.UserAccessMenuModel) error {
	err := repository.DB.First(userAccess, id).Error
	if err != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("getUpdatedUserAccess | UserAccessMenuModel | Repository | Error when querying updated data, err:%s", err.Error()))
		return errors.New("failed to get updated user access")
	}
	return nil
}
