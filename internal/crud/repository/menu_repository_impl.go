package repository

import (
	"fmt"

	"github.com/Ganasa18/document-be/internal/crud/model/domain"
	"github.com/Ganasa18/document-be/pkg/helper"
	"github.com/Ganasa18/document-be/pkg/loghelper"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MenuRepositoryImpl struct {
	DB *gorm.DB
}

func NewMenuRepository(db *gorm.DB) MenuRepository {
	return &MenuRepositoryImpl{
		DB: db,
	}
}

func (repository *MenuRepositoryImpl) CreateMenu(ctx *gin.Context, menu domain.MenuMasterModel) (domain.MenuMasterModel, error) {

	err := repository.DB.Create(&menu).Error
	if err != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("CreateMenu | Error when Query builder create data, err:%s", err.Error()))
		return menu, err
	}
	return menu, nil
}

func (repository *MenuRepositoryImpl) DeleteMenu(ctx *gin.Context, id int) error {
	panic("unimplemented")
}

func (repository *MenuRepositoryImpl) GetAllMenu(ctx *gin.Context, pagination *helper.PaginationInput) (menu []domain.MenuMasterModel, totalRow int64, err error) {
	totalRow = 0

	offset := (pagination.Page - 1) * pagination.Limit
	queryBuilder := repository.DB.Limit(pagination.Limit).Offset(offset)

	if pagination.Search != "" {
		queryBuilder = queryBuilder.Where("name ILIKE ?", "%"+pagination.Search+"%")
	}

	// GET DATA
	err = queryBuilder.Model(&domain.MenuMasterModel{}).Where("deleted_at IS NULL").Find(&menu).Error

	if err != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("GetAllMenu | Error when Query builder list data, err:%s", err.Error()))
		return menu, totalRow, err
	}

	// ROW COUNT
	searchBuider := repository.DB.Model(&domain.MenuMasterModel{}).Where("deleted_at IS NULL")
	if pagination.Search != "" {
		searchBuider = searchBuider.Where("name ILIKE ?", "%"+pagination.Search+"%")
	}
	errCount := searchBuider.Count(&totalRow).Error

	if errCount != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("GetAllMenu | Error when Query builder count total rows, err:%s", errCount.Error()))
		return menu, totalRow, errCount
	}
	return menu, totalRow, nil

}

func (repository *MenuRepositoryImpl) GetMenuById(ctx *gin.Context, id int) (domain.MenuMasterModel, error) {
	panic("unimplemented")
}

func (repository *MenuRepositoryImpl) UpdateMenu(ctx *gin.Context, request domain.MenuMasterModel) {
	panic("unimplemented")
}
