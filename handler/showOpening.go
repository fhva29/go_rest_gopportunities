package handler

import (
	"fmt"
	"gopportunities/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParam").Error())
		return
	}

	opening := schemas.Opening{}

	//Find opening
	if err := db.First(&opening, id).Error; err != nil {
		logger.Errorf("error finding opening: %v", err.Error())
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	sendSuccess(ctx, "show-opening", opening)
}
