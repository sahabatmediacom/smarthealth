package router

import (
	"pamer-api/internal/router"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(api *gin.RouterGroup) {
	// Ping route
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.AuthRouter(api)

	router.HospitalRouter(api)
	router.ParamedicRouter(api)
	router.PatientRouter(api)
}
