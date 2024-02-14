package service

import (
	"github.com/Ganasa18/document-be/internal/crud/model/web"
	"github.com/Ganasa18/document-be/pkg/helper"
	"github.com/gin-gonic/gin"
)

type MenuService interface {
	GetAllMenu(ctx *gin.Context, pagination *helper.PaginationInput) ([]web.MenuMasterResponse, int64, error)
	GetMenuById(ctx *gin.Context) web.MenuMasterResponse
	CreateMenu(ctx *gin.Context, request web.MenuMasterRequest) (web.MenuMasterResponse, error)
	UpdateMenu(ctx *gin.Context, request web.MenuMasterRequest)
	DeleteMenu(ctx *gin.Context)
}
