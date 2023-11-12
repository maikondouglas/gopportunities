package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"msg": "GET Openings",
	})
}
