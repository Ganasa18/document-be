package loghelper

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Infoln(ctx context.Context, args ...interface{}) {
	logrus.WithContext(ctx).Infoln(args...)
}

func Errorln(ctx context.Context, args ...interface{}) {
	logrus.WithContext(ctx).Errorln(args...)
}

func Fatalln(ctx context.Context, args ...interface{}) {
	logrus.WithContext(ctx).Fatalln(args...)
}

func Warningln(ctx context.Context, args ...interface{}) {
	logrus.WithContext(ctx).Warningln(args...)
}

func LogErrorRepository(ctx *gin.Context, operation, model string, errMsg string, err error) {
	logMessage := fmt.Sprintf("%s | %s | %s | Repository | Error: %s", operation, model, errMsg, err.Error())
	Errorln(ctx, logMessage)
}
