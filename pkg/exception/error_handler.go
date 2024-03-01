package exception

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Ganasa18/document-be/internal/base/model/web"
	"github.com/Ganasa18/document-be/pkg/helper"
	"github.com/Ganasa18/document-be/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

func ExceptionRecoveryMiddleware(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic recovered:", err)
			handleError(c, err)
			c.Abort()
		}
	}()

	c.Next()
}

func handleError(ctx *gin.Context, err interface{}) {
	if notFoundError(ctx, err) {
		return
	}

	if validationErrors(ctx, err) {
		return
	}

	internalServerError(ctx, err)
}

type ValidationError struct {
	FieldName string `json:"field_name"`
}

func validationErrors(ctx *gin.Context, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)

	if ok {
		ctx.Writer.Header().Set(utils.HEADER_CONTENT_TYPE, utils.CONTENT_TYPE_APPLICATION_JSON)
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		// Convert the validation errors into an array
		errorArray := make([]ValidationError, len(exception))
		for i, fieldError := range exception {
			errorArray[i] = ValidationError{
				FieldName: "Field " + helper.CamelToSnake(fieldError.Field()) + " " + fieldError.Tag() + " error",
			}
		}

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   map[string]interface{}{"errors": errorArray},
		}
		logrus.Errorln(fmt.Sprintf("METHOD: %s, URL: %s, message: %s", ctx.Request.Method, ctx.Request.URL, errorArray))
		helper.WriteToResponseBody(ctx, http.StatusBadRequest, webResponse)
		return true
	} else {
		return false
	}
}
func notFoundError(ctx *gin.Context, err interface{}) bool {
	start := time.Now()
	exception, ok := err.(NotFoundError)
	if ok {
		ctx.Writer.Header().Set(utils.HEADER_CONTENT_TYPE, utils.CONTENT_TYPE_APPLICATION_JSON)
		ctx.Writer.WriteHeader(http.StatusNotFound)
		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}
		end := time.Since(start)
		logrus.Infoln(fmt.Sprintf("METHOD: %s, URL: %s, RESPONSE: %s , LATENCY: %vms", ctx.Request.Method, ctx.Request.URL, exception.Error, end.Milliseconds()))
		helper.WriteToResponseBody(ctx, http.StatusNotFound, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(ctx *gin.Context, err interface{}) {
	start := time.Now()
	ctx.Writer.Header().Set(utils.HEADER_CONTENT_TYPE, utils.CONTENT_TYPE_APPLICATION_JSON)
	ctx.Writer.WriteHeader(http.StatusInternalServerError)
	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}
	end := time.Since(start)
	logrus.Infoln(fmt.Sprintf("METHOD: %s, URL: %s, RESPONSE: %s , LATENCY: %vms", ctx.Request.Method, ctx.Request.URL, err, end.Milliseconds()))
	helper.WriteToResponseBody(ctx, http.StatusInternalServerError, webResponse)
}
