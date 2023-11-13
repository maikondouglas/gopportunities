package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maikondouglas/gopportunities/schemas"
)

// @BasePath /api/v1/
// @Sumary Show opening
// @Description Show a job opening
// @Tags Openings
// @Accept json
// @Peoduce json
// @Param id query string true "Opening indentification"
// @Sucess 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorReponse
// @Failure 404 {object} ErrorReponse
// @Router /opening/:id [get]
func ShowOpeningHandler(context *gin.Context) {
	id := context.Param("id")

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

	sendSuccess(context, "show-opening", opening)
}
