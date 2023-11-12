package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"msg": "PUT Opening",
	})
}
