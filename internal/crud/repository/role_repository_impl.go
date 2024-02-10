package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

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

func (repository *RoleRepositoryImpl) CreateRole(ctx context.Context, role domain.RoleMasterModel) (domain.RoleMasterModel, error) {

	RoleName := strings.ToLower(role.RoleName)
	// Check existing role
	err := repository.DB.Where(&domain.RoleMasterModel{RoleName: RoleName}).First(&role).Error

	if err != gorm.ErrRecordNotFound {
		return role, errors.New("cannot duplicate insert")
	}

	err = repository.DB.Model(&domain.RoleMasterModel{}).Create(&role).Error
	if err != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("CreateRole | Error when Query builder create data, err:%s", err.Error()))
		return role, err
	}

	return role, nil
}

func (repository *RoleRepositoryImpl) UpdateRole(ctx context.Context, role domain.RoleMasterModel, id int) (domain.RoleMasterModel, error) {
	// Check if the role with the given ID exists
	err := repository.DB.First(&domain.RoleMasterModel{}, id).Error
	if err != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("UpdateRole | Error when querying data, err:%s", err.Error()))
		return role, errors.New("role not found")
	}

	// Update the role_name for the role with the given ID
	err = repository.DB.Model(&domain.RoleMasterModel{}).Where("id = ?", id).Updates(map[string]interface{}{"role_name": role.RoleName}).Error
	if err != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("UpdateRole | Error when updating data, err:%s", err.Error()))
		return role, errors.New("failed to update role")
	}

	// Retrieve the updated role
	err = repository.DB.First(&role, id).Error
	if err != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("UpdateRole | Error when querying updated data, err:%s", err.Error()))
		return role, errors.New("failed to get updated role")
	}

	return role, err
}

func (repository *RoleRepositoryImpl) DeleteRole(ctx context.Context, id int) error {
	role := &domain.RoleMasterModel{Id: id}

	err := repository.DB.First(&domain.RoleMasterModel{}, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("role not found")
	}

	err = repository.DB.Delete(role).Error

	if err != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("DeleteRole | Error when deleting role, err: %s", err.Error()))
		return errors.New("failed to delete role")
	}

	return nil
}

func (repository *RoleRepositoryImpl) GetRoleById(ctx context.Context, id int) (role domain.RoleMasterModel, err error) {
	err = repository.DB.Model(&domain.RoleMasterModel{}).First(&role, id).Error
	if err != nil {
		loghelper.Errorln(ctx, fmt.Sprintf("GetRoleById | Error when Query builder get data, err:%s", err.Error()))
		return role, errors.New("role not found")
	}
	return role, nil
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
