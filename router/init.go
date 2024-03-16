package router

import "github.com/gin-gonic/gin"

func InitializeRouter() *gin.Engine {
	r := gin.Default()

	// Define API group
	api := r.Group("/api")

	// Define routes
	initializeRoutes(api)

	return r
}
