package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maikondouglas/gopportunities/schemas"
)

func UpdateOpeningHandler(context *gin.Context) {
	request := UpdateOpeningRequest{}

	context.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		sendError(context, http.StatusBadRequest, err.Error())

		return
	}

	id := context.Query("id")
	if id == "" {
		sendError(context, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())

		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(context, http.StatusNotFound, "Opening not found")

		return
	}

	// Updating opening
	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	// Save opening
	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("Error updating opening: %v", err.Error())
		sendError(context, http.StatusInternalServerError, "Error updating opening")

		return
	}

	sendSuccess(context, "update-opening", opening)
}
