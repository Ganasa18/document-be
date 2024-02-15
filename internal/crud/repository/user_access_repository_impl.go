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

func (*UserAccessRepositoryImpl) DeleteUserAccess(ctx *gin.Context) {
	panic("unimplemented")
}

func (*UserAccessRepositoryImpl) GetAllUserAccess(ctx *gin.Context, pagination *helper.PaginationInput) {
	panic("unimplemented")
}

func (*UserAccessRepositoryImpl) GetUserAccessById(ctx *gin.Context) {
	panic("unimplemented")
}

func (*UserAccessRepositoryImpl) UpdateUserAccess(ctx *gin.Context) {
	panic("unimplemented")
}
