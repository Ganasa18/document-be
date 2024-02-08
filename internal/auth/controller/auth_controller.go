package controller

import (
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	LoginOrRegister(ctx *gin.Context)
}
