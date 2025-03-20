package handler

import (
	"gopportunities/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		logger.Errorf("error finding openings: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error finding openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
