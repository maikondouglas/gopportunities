package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maikondouglas/gopportunities/schemas"
)

// @BasePath /api/v1/
// @Sumary Create opening
// @Description Create a new job opening
// @Tags Openings
// @Accept json
// @Peoduce json
// @Param request body CreateOpeningRequest true "Request body"
// @Sucess 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorReponse
// @Failure 500 {object} ErrorReponse
// @Router /opening [post]
func CreateOpeningHandler(context *gin.Context) {
	request := CreateOpeningRequest{}

	context.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		sendError(context, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("Error creating opening: %v", err)
		sendError(context, http.StatusInternalServerError, "Error creating opening on database")
		return
	}

	sendSuccess(context, "create-opening", opening)
}
