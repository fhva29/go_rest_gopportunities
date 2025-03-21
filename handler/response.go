package handler

import (
	"fmt"
	"gopportunities/schemas"
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

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode int    `json:"errorCode"`
}

type CreateOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}
