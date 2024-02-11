package controller

import (
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	LoginOrRegister(ctx *gin.Context)
	ForgotLinkPassword(ctx *gin.Context)
}
