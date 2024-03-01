package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Ganasa18/document-be/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	utils.IsErrorDoPanic(err)
}

func WriteToResponseBody(ctx *gin.Context, statusCode int, response interface{}) {
	start := time.Now()
	respStr, _ := json.Marshal(response)
	end := time.Since(start)
	logrus.Infoln(fmt.Sprintf("METHOD: %s, URL: %s, RESPONSE: %s, LATENCY: %vms", ctx.Request.Method, ctx.Request.URL, respStr, end.Milliseconds()))
	ctx.JSON(statusCode, response)

}
