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
		return userAccess, errors.New("cannot duplicate insert")
	}

	err = repository.DB.Model(&domain.UserAccessMenuModel{}).Create(&userAccess).Error
	if err != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("CreateUserAccess | Error when Query builder create data, err:%s", err.Error()))
		return userAccess, err
	}

	return userAccess, nil

}

func (repository *UserAccessRepositoryImpl) GetAllUserAccess(ctx *gin.Context, pagination *helper.PaginationInput) (userAccess []domain.UserAccessMenuModel, totalRow int64, err error) {
	totalRow = 0

	offset := (pagination.Page - 1) * pagination.Limit
	queryBuilder := repository.DB.Limit(pagination.Limit).Offset(offset)

	// if pagination.Search != "" {
	// 	queryBuilder = queryBuilder.Where("role_name ILIKE ?", "%"+pagination.Search+"%")
	// }

	// GET DATA
	err = queryBuilder.Model(&domain.UserAccessMenuModel{}).Preload("RoleMasterModel").Preload("MenuMasterModel").Where("deleted_at IS NULL").Find(&userAccess).Error

	if err != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("GetAllUserAccess | Error when Query builder list data, err:%s", err.Error()))
		return userAccess, totalRow, err
	}

	// ROW COUNT
	searchBuider := repository.DB.Model(&domain.UserAccessMenuModel{}).Where("deleted_at IS NULL")
	// if pagination.Search != "" {
	// 	searchBuider = searchBuider.Where("role_name ILIKE ?", "%"+pagination.Search+"%")
	// }
	errCount := searchBuider.Count(&totalRow).Error

	if errCount != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("GetAllUserAccess | Error when Query builder count total rows, err:%s", errCount.Error()))
		return userAccess, totalRow, errCount
	}
	return userAccess, totalRow, nil
}

func (repository *UserAccessRepositoryImpl) GetUserAccessById(ctx *gin.Context, id int) (userAccess domain.UserAccessMenuModel, err error) {
	err = repository.DB.Model(&domain.UserAccessMenuModel{}).First(&userAccess, id).Error
	if err != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("GetUserAccessById | Error when Query builder get data, err:%s", err.Error()))
		return userAccess, errors.New("user access not found")
	}
	return userAccess, nil
}

func (*UserAccessRepositoryImpl) UpdateUserAccess(ctx *gin.Context) {
	panic("unimplemented")
}

func (repository *UserAccessRepositoryImpl) DeleteUserAccess(ctx *gin.Context, id int) error {
	userAccess := &domain.UserAccessMenuModel{Id: id}

	err := repository.DB.First(&domain.UserAccessMenuModel{}, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("user access not found")
	}

	err = repository.DB.Unscoped().Delete(userAccess).Error

	if err != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("DeleteUserAccess | Error when deleting user access, err: %s", err.Error()))
		return errors.New("failed to delete user access")
	}

	return nil
}
