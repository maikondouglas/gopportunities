package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maikondouglas/gopportunities/schemas"
)

func sendError(context *gin.Context, code int, msg string) {
	context.Header("Content-Type", "application/json")
	context.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(context *gin.Context, op string, data interface{}) {
	context.Header("Content-Type", "application/json")
	context.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Operation %s successfull", op),
		"data":    data,
	})
}

type ErrorReponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreagteOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

type DeleteOpeningResponse struct {
	Message string          `json:"message"`
	Data    schemas.Opening `json:"data"`
}

type ShowOpeningResponse struct {
	Message string          `json:"message"`
	Data    schemas.Opening `json:"data"`
}

type ListOpeningResponse struct {
	Message string            `json:"message"`
	Data    []schemas.Opening `json:"data"`
}
