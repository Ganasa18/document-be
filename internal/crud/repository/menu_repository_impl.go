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
		loghelper.Errorln(ctx, fmt.Sprintf("CreateMenu | MenuMasterModel | Repository | Error when Query builder create data, err:%s", err.Error()))
		return menu, err
	}
	return menu, nil
}

func (repository *MenuRepositoryImpl) GetAllMenu(ctx *gin.Context, pagination *helper.PaginationInput) (menu []domain.MenuMasterModel, totalRow int64, err error) {
	totalRow = 0

	offset := (pagination.Page - 1) * pagination.Limit
	queryBuilder := repository.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.OrderBy)

	if pagination.Search != "" {
		queryBuilder = queryBuilder.Where("name ILIKE ?", "%"+pagination.Search+"%")
	}

	// GET DATA
	err = queryBuilder.Model(&domain.MenuMasterModel{}).Where("deleted_at IS NULL").Find(&menu).Error

	if err != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("GetAllMenu | MenuMasterModel | Repository | Error when Query builder list data, err:%s", err.Error()))
		return menu, totalRow, err
	}

	// ROW COUNT
	searchBuider := repository.DB.Model(&domain.MenuMasterModel{}).Where("deleted_at IS NULL")
	if pagination.Search != "" {
		searchBuider = searchBuider.Where("name ILIKE ?", "%"+pagination.Search+"%")
	}
	errCount := searchBuider.Count(&totalRow).Error

	if errCount != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("GetAllMenu | MenuMasterModel | Repository | Error when Query builder count total rows, err:%s", errCount.Error()))
		return menu, totalRow, errCount
	}
	return menu, totalRow, nil

}

func (repository *MenuRepositoryImpl) GetMenuById(ctx *gin.Context, id int) (menu domain.MenuMasterModel, err error) {
	err = repository.DB.Model(&domain.MenuMasterModel{}).First(&menu, id).Error
	if err != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("GetMenuById | MenuMasterModel | Repository | Error when Query builder get data, err:%s", err.Error()))
		return menu, errors.New("menu not found")
	}
	return menu, nil
}

func (repository *MenuRepositoryImpl) UpdateMenu(ctx *gin.Context, menu domain.MenuMasterModel, id int) (domain.MenuMasterModel, error) {
	// Check if the menu with the given ID exists
	if err := repository.checkIfMenuExists(ctx, id); err != nil {
		return menu, err
	}

	// Update the menu for the given ID
	if err := repository.updateMenuFields(ctx, id, menu); err != nil {
		return menu, err
	}

	// Retrieve the updated menu
	if err := repository.getUpdatedMenu(ctx, id, &menu); err != nil {
		return menu, err
	}

	return menu, nil
}

func (repository *MenuRepositoryImpl) DeleteMenu(ctx *gin.Context, id int) error {
	menu := &domain.MenuMasterModel{Id: id}

	err := repository.DB.First(&domain.MenuMasterModel{}, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		loghelper.Errorln(ctx, fmt.Sprintf("DeleteMenu | MenuMasterModel | Repository | Error menu not found, err: %s", err.Error()))
		return errors.New("menu not found")
	}

	err = repository.DB.Delete(menu).Error

	if err != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("DeleteMenu | MenuMasterModel | Repository | Error when deleting menu, err: %s", err.Error()))
		return errors.New("failed to delete menu")
	}

	return nil
}

func (repository *MenuRepositoryImpl) checkIfMenuExists(ctx *gin.Context, id int) error {
	err := repository.DB.First(&domain.MenuMasterModel{}, id).Error
	if err != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("checkIfMenuExists | MenuMasterModel | Repository | Error when querying data, err:%s", err.Error()))
		return errors.New("menu not found")
	}
	return nil
}

func (repository *MenuRepositoryImpl) updateMenuFields(ctx *gin.Context, id int, menu domain.MenuMasterModel) error {

	updateFields := map[string]interface{}{"name": menu.Name, "title": menu.Title, "path": menu.Path, "icon_name": menu.IconName, "is_sub_menu": menu.IsSubMenu}

	err := repository.DB.Model(&menu).
		Where("id = ?", id).
		Updates(updateFields).Error
	if err != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("updateMenuFields | MenuMasterModel | Repository | Error when updating data, err:%s", err.Error()))
		return errors.New("failed to update menu")
	}
	return nil
}

func (repository *MenuRepositoryImpl) getUpdatedMenu(ctx *gin.Context, id int, menu *domain.MenuMasterModel) error {
	err := repository.DB.First(menu, id).Error
	if err != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("getUpdatedMenu | MenuMasterModel | Repository | Error when querying updated data, err:%s", err.Error()))
		return errors.New("failed to get updated menu")
	}
	return nil
}
