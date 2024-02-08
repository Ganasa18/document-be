package server

import "github.com/gin-gonic/gin"

func (h *HttpServe) setupRouter() {
	v1 := h.router.Group("/api/v1")

	v1.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})
}
