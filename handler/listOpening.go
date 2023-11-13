package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maikondouglas/gopportunities/schemas"
)

// @BasePath /api/v1/
// @Sumary List openings
// @Description List all jobs openings
// @Tags Openings
// @Accept json
// @Peoduce json
// @Sucess 200 {object} ListOpeningResponse
// @Failure 400 {object} ErrorReponse
// @Failure 404 {object} ErrorReponse
// @Router /openings [get]
func ListOpeningsHandler(context *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(context, http.StatusInternalServerError, "Error listing openings")
	}

	sendSuccess(context, "list-openings", openings)
}
