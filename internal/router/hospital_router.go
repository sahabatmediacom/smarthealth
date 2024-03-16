package router

import (
	"pamer-api/config"
	"pamer-api/internal/handler"
	"pamer-api/internal/repository"
	"pamer-api/internal/service"

	"github.com/gin-gonic/gin"
)

func HospitalRouter(api *gin.RouterGroup) {
	hospitalRepository := repository.NewHospitalRepo(config.DB)
	hospitalService := service.NewHospitalService(hospitalRepository)
	hospitalHandler := handler.NewHospitalHandler(hospitalService)

	r := api.Group("/hospitals")
	r.GET("/", hospitalHandler.GetAll)
	r.GET("/:id", hospitalHandler.Get)
	r.POST("/", hospitalHandler.Create)
	r.PUT("/:id", hospitalHandler.Update)
	r.DELETE("/:id", hospitalHandler.Delete)
}
