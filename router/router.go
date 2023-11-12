package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize the router
	router := gin.Default()

	// Initialize the router
	initializeRouters(router)

	// Run the server
	router.Run(":8080")
}
