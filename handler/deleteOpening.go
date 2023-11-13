package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maikondouglas/gopportunities/schemas"
)

// @BasePath /api/v1/
// @Sumary Delete opening
// @Description Delete a new job opening
// @Tags Openings
// @Accept json
// @Peoduce json
// @Param id query string true "Opening indentification"
// @Sucess 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorReponse
// @Failure 404 {object} ErrorReponse
// @Router /opening [delete]
func DeleteOpeningHandler(context *gin.Context) {
	id := context.Query("id")
	if id == "" {
		sendError(context, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	// Find Opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(context, http.StatusNotFound, fmt.Sprintf("Opening with id: %s not found", id))
		return
	}

	// Delete opening
	if err := db.Delete(&opening).Error; err != nil {
		sendError(context, http.StatusInternalServerError, fmt.Sprintf("Error deleting opening with id: %s", id))
		return
	}

	sendSuccess(context, "delete-opening", opening)
}
