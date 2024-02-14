package repository

import (
	"github.com/Ganasa18/document-be/internal/crud/model/domain"
	"github.com/Ganasa18/document-be/pkg/helper"
	"github.com/gin-gonic/gin"
)

type MenuRepository interface {
	GetAllMenu(ctx *gin.Context, pagination *helper.PaginationInput) ([]domain.MenuMasterModel, int64, error)
	GetMenuById(ctx *gin.Context, id int) (domain.MenuMasterModel, error)
	CreateMenu(ctx *gin.Context, request domain.MenuMasterModel) (domain.MenuMasterModel, error)
	UpdateMenu(ctx *gin.Context, request domain.MenuMasterModel)
	DeleteMenu(ctx *gin.Context, id int) error
}
