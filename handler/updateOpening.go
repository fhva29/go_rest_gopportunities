package handler

import (
	"fmt"
	"gopportunities/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("error validating request: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParam").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		logger.Errorf("error finding opening: %v", err.Error())
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	err := db.Model(&opening).Updates(request).Error
	if err != nil {
		logger.Errorf("error updating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error updating opening with id: %s", id))
		return
	}

	sendSuccess(ctx, "update-opening", opening)
}
