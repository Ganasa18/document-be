package repository

import (
	"errors"
	"strings"

	"github.com/Ganasa18/document-be/internal/crud/model/domain"
	"github.com/Ganasa18/document-be/pkg/helper"
	"github.com/Ganasa18/document-be/pkg/loghelper"
	"github.com/gin-gonic/gin"
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

func (repository *RoleRepositoryImpl) CreateRole(ctx *gin.Context, role domain.RoleMasterModel) (domain.RoleMasterModel, error) {

	RoleName := strings.ToLower(role.RoleName)
	// Check existing role
	err := repository.DB.Where(&domain.RoleMasterModel{RoleName: RoleName}).First(&role).Error

	if err != gorm.ErrRecordNotFound {
		loghelper.LogErrorRepository(ctx, "CreateRole", "RoleMasterModel", "cannot duplicate insert", err)
		return role, errors.New("cannot duplicate insert")
	}

	err = repository.DB.Model(&domain.RoleMasterModel{}).Create(&role).Error
	if err != nil {
		loghelper.LogErrorRepository(ctx, "CreateRole", "RoleMasterModel", "when query builder create data", err)
		return role, err
	}

	return role, nil
}

func (repository *RoleRepositoryImpl) UpdateRole(ctx *gin.Context, role domain.RoleMasterModel, id int) (domain.RoleMasterModel, error) {
	// Check if the role with the given ID exists
	err := repository.DB.First(&domain.RoleMasterModel{}, id).Error
	if err != nil {
		loghelper.LogErrorRepository(ctx, "UpdateRole", "RoleMasterModel", "when querying data", err)
		return role, errors.New("role not found")
	}

	// Update the role_name for the role with the given ID
	err = repository.DB.Model(&domain.RoleMasterModel{}).Where("id = ?", id).Updates(map[string]interface{}{"role_name": role.RoleName}).Error
	if err != nil {
		loghelper.LogErrorRepository(ctx, "UpdateRole", "RoleMasterModel", "when updating data", err)
		return role, errors.New("failed to update role")
	}

	// Retrieve the updated role
	err = repository.DB.First(&role, id).Error
	if err != nil {
		loghelper.LogErrorRepository(ctx, "UpdateRole", "RoleMasterModel", "when querying updated data", err)
		return role, errors.New("failed to get updated role")
	}

	return role, err
}

func (repository *RoleRepositoryImpl) DeleteRole(ctx *gin.Context, id int) error {
	role := &domain.RoleMasterModel{Id: id}

	err := repository.DB.First(&domain.RoleMasterModel{}, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		loghelper.LogErrorRepository(ctx, "DeleteRole", "RoleMasterModel", "role not found", err)
		return errors.New("role not found")
	}

	err = repository.DB.Delete(role).Error

	if err != nil {
		loghelper.LogErrorRepository(ctx, "DeleteRole", "RoleMasterModel", "when deleting role", err)
		return errors.New("failed to delete role")
	}

	return nil
}

func (repository *RoleRepositoryImpl) GetRoleById(ctx *gin.Context, id int) (role domain.RoleMasterModel, err error) {
	err = repository.DB.Model(&domain.RoleMasterModel{}).First(&role, id).Error
	if err != nil {
		loghelper.LogErrorRepository(ctx, "GetRoleById", "RoleMasterModel", "when query builder get data", err)
		return role, errors.New("role not found")
	}
	return role, nil
}

func (repository *RoleRepositoryImpl) GetRoles(ctx *gin.Context, pagination *helper.PaginationInput) (roles []domain.RoleMasterModel, totalRow int64, err error) {

	totalRow = 0

	offset := (pagination.Page - 1) * pagination.Limit
	queryBuilder := repository.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.OrderBy)

	if pagination.Search != "" {
		queryBuilder = queryBuilder.Where("role_name ILIKE ?", "%"+pagination.Search+"%")
	}

	// GET DATA
	err = queryBuilder.Model(&domain.RoleMasterModel{}).Where("deleted_at IS NULL").Find(&roles).Error

	if err != nil {
		loghelper.LogErrorRepository(ctx, "FindAll", "RoleMasterModel", "when query builder list data", err)
		return roles, totalRow, err
	}

	// ROW COUNT
	searchBuider := repository.DB.Model(&domain.RoleMasterModel{}).Where("deleted_at IS NULL")
	if pagination.Search != "" {
		searchBuider = searchBuider.Where("role_name ILIKE ?", "%"+pagination.Search+"%")
	}
	errCount := searchBuider.Count(&totalRow).Error

	if errCount != nil {
		loghelper.LogErrorRepository(ctx, "FindAll", "RoleMasterModel", "when query builder count total rows", err)
		return roles, totalRow, errCount
	}
	return roles, totalRow, nil

}
