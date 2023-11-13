package router

import (
	"github.com/gin-gonic/gin"
	"github.com/maikondouglas/gopportunities/docs"
	"github.com/maikondouglas/gopportunities/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRouters(router *gin.Engine) {
	// Initialize Handler
	basePath := "/api/v1"
	handler.InitializeHandler()

	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)

	{
		v1.GET("/opening/:id", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
	}

	// Initialize swagger client
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
