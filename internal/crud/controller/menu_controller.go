package controller

import "github.com/gin-gonic/gin"

type MenuController interface {
	GetAllMenu(ctx *gin.Context)
	GetMenuById(ctx *gin.Context)
	CreateMenu(ctx *gin.Context)
	UpdateMenu(ctx *gin.Context)
	DeleteMenu(ctx *gin.Context)
}
