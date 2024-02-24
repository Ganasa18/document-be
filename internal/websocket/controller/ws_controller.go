package controller

import "github.com/gin-gonic/gin"

type WebSocketController interface {
	HandlerWebSocketController(ctx *gin.Context)
	HandleMessages()
}
