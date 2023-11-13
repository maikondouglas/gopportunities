package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maikondouglas/gopportunities/schemas"
)

// @BasePath /api/v1/
// @Sumary Update opening
// @Description Update a new job opening
// @Tags Openings
// @Accept json
// @Peoduce json
// @Param id query string true "Opening indentification"
// @Param opening body UpdateOpeningRequest true "Opening data to Update"
// @Sucess 200 {object} UpdateOpeningResponse
// @Failure 400 {object} ErrorReponse
// @Failure 500 {object} ErrorReponse
// @Router /opening [put]
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
