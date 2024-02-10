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
		// CHECK HANDLER IF AUTHORIZATION HEADER EMPTY
		if providedToken == "" {
			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: http.StatusText(http.StatusUnauthorized),
				Data:   "Token not Provided",
			}
			helper.WriteToResponseBody(ctx, http.StatusUnauthorized, webResponse)
			ctx.Abort()
			return
		}

		authSplited := strings.Split(providedToken, " ")

		// CHECK HANDLER IF AUTHORIZATION HEADER NOT HAVE BEARER
		if authSplited[0] != "Bearer" {
			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: http.StatusText(http.StatusUnauthorized),
				Data:   "Header Authorization is not valid: Bearer not found",
			}
			helper.WriteToResponseBody(ctx, http.StatusUnauthorized, webResponse)
			ctx.Abort()
			return
		}

		// CHECK LENGTH AUTHORIZATION
		if len(authSplited) < 2 || len(authSplited) > 2 {
			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: http.StatusText(http.StatusUnauthorized),
				Data:   "Header Authorization is not valid: len not valid",
			}
			helper.WriteToResponseBody(ctx, http.StatusUnauthorized, webResponse)
			ctx.Abort()
			return
		}

		tokenValid, _ := helper.ValidateToken(authSplited[1])

		// CHECK TOKEN VALID
		if tokenValid == nil {
			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: http.StatusText(http.StatusUnauthorized),
				Data:   "Not Authorization",
			}
			helper.WriteToResponseBody(ctx, http.StatusUnauthorized, webResponse)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
