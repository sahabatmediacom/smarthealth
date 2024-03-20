package router

import (
	"pamer-api/internal/handler"
	"pamer-api/internal/repository"
	"pamer-api/internal/service"
	"pamer-api/middleware"

	"github.com/gin-gonic/gin"
)

func PatientRouter(api *gin.RouterGroup) {
	patientRepository := repository.NewPatientRepository()
	patientService := service.NewPatientService(patientRepository)
	patientHandler := handler.NewPatientHandler(patientService)

	r := api.Group("/patients")
	r.Use(middleware.Authenticate)
	r.GET("/", patientHandler.GetAll)
	r.GET("/:id", patientHandler.GetPatientData)
	r.GET("/:id/record", patientHandler.GetPatientRecord)
	r.POST("/:id/record")

}
