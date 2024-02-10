package repository

import (
	"context"
	"fmt"

	"github.com/Ganasa18/document-be/internal/crud/model/domain"
	"github.com/Ganasa18/document-be/pkg/helper"
	"github.com/Ganasa18/document-be/pkg/loghelper"
	"gorm.io/gorm"
)

type RoleRepositoryImpl struct {
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &RoleRepositoryImpl{
		DB: db,
	}
}

func (repository *RoleRepositoryImpl) GetRoles(ctx context.Context, pagination *helper.PaginationInput) (roles []domain.RoleMasterModel, totalRow int64, err error) {

	totalRow = 0

	offset := (pagination.Page - 1) * pagination.Limit
	queryBuilder := repository.DB.Limit(pagination.Limit).Offset(offset)

	if pagination.Search != "" {
		queryBuilder = queryBuilder.Where("role_name ILIKE ?", "%"+pagination.Search+"%")
	}

	// GET DATA
	err = queryBuilder.Model(&domain.RoleMasterModel{}).Where("deleted_at IS NULL").Find(&roles).Error

	if err != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("FindAll | Error when Query builder list data, err:%s", err.Error()))
		return roles, totalRow, err
	}

	// ROW COUNT
	searchBuider := repository.DB.Model(&domain.RoleMasterModel{}).Where("deleted_at IS NULL")
	if pagination.Search != "" {
		searchBuider = searchBuider.Where("role_name ILIKE ?", "%"+pagination.Search+"%")
	}
	errCount := searchBuider.Count(&totalRow).Error

	if errCount != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("FindAll | Error when Query builder count total rows, err:%s", errCount.Error()))
		return roles, totalRow, errCount
	}
	return roles, totalRow, nil

}
