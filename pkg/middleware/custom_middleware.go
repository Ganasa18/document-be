package middleware

import (
	"net/http"
	"strings"

	"github.com/Ganasa18/document-be/internal/base/model/web"
	"github.com/Ganasa18/document-be/pkg/helper"
	"github.com/Ganasa18/document-be/pkg/utils"
	"github.com/gin-gonic/gin"
)

func CustomAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		providedToken := ctx.Request.Header.Get(utils.HEADER_AUTHORIZATION)

		if providedToken == "" {
			handleUnauthorized(ctx, "Token not Provided")
			return
		}

		authSplited := strings.Split(providedToken, " ")

		if len(authSplited) != 2 || authSplited[0] != "Bearer" {
			handleUnauthorized(ctx, "Header Authorization is not valid")
			return
		}

		tokenValid, _ := helper.ValidateToken(authSplited[1])

		if tokenValid == nil {
			handleUnauthorized(ctx, "Not Authorized")
			return
		}

		ctx.Next()
	}
}

func handleUnauthorized(ctx *gin.Context, message string) {
	webResponse := web.WebResponse{
		Code:   http.StatusUnauthorized,
		Status: http.StatusText(http.StatusUnauthorized),
		Data:   message,
	}
	helper.WriteToResponseBody(ctx, http.StatusUnauthorized, webResponse)
	ctx.Abort()
}
