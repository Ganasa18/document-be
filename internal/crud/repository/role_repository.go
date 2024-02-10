package repository

import (
	"context"

	"github.com/Ganasa18/document-be/internal/crud/model/domain"
	"github.com/Ganasa18/document-be/pkg/helper"
)

type RoleRepository interface {
	GetRoles(ctx context.Context, pagination *helper.PaginationInput) ([]domain.RoleMasterModel, int64, error)
}
