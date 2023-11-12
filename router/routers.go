package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRouters(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	{
		v1.GET("/opening", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "GET Opening",
			})
		})

		v1.POST("/opening", func(context *gin.Context) {
			context.JSON(http.StatusCreated, gin.H{
				"msg": "POST Opening",
			})
		})

		v1.DELETE("/opening", func(context *gin.Context) {
			context.JSON(http.StatusNoContent, gin.H{
				"msg": "DELETE Opening",
			})
		})

		v1.PUT("/opening", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "PUT Opening",
			})
		})

		v1.GET("/openings", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "GET Openings",
			})
		})
	}
}
