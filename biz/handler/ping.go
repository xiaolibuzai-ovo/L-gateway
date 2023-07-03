package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Ping .
func Ping(ctx context.Context, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
