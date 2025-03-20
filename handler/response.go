package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, message string) {
	ctx.JSON(code, gin.H{
		"message":   message,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data any) {
	ctx.JSON(
		http.StatusOK, gin.H{
			"message": fmt.Sprintf("operation from handler: %s succeeded", op),
			"data":    data,
		},
	)
}
